package core

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
)

// Client 是支付宝 V3 REST OpenAPI 的通用客户端：构造相对路由请求、
// JSON 请求体或 query 参数、Authorization 头签名、HTTP 执行、
// 响应头验签与错误分类。它实现 Executor，所有领域接口与自定义
// V3 接口共用该执行流程；签名与验签使用构造期绑定的认证快照，
// 不回调 Credentials getter。
//
// V3 的成功语义由 HTTP 状态决定；非 2xx 响应按官方契约解析为
// {code, message, links} 错误体。V3 请求模型的 notify_url 属于业务
// 字段，本 Client 拒绝 WithNotifyURL 调用选项。
//
// Client 不解释商户订单状态，不包含业务字段。
// 初始化后可并发复用；每次调用生成独立的时间戳、nonce 与签名。
type Client struct {
	settings Settings
	snapshot authSnapshot
	http     *http.Client
	logger   *slog.Logger
	now      func() time.Time
	nonce    func() string
}

// NewClient 以 Credentials 为第一个必选参数构建 V3 客户端。
// 构造顺序：校验凭证并绑定认证快照 → 处理公共选项 → 校验配置。
// 构造过程检查全部本地可验证的配置，不进行网络 I/O；
// 配置缺失或非法时返回 nil 与 *ConfigError，不返回可运行的 Client。
// 应用是否开通、凭证是否被平台接受及网络可达性由实际 API 调用确认。
func NewClient(credentials auth.Credentials, opts ...ClientOption) (*Client, error) {
	snapshot, err := bindCredentials(credentials)
	if err != nil {
		return nil, err
	}
	var settings Settings
	for _, opt := range opts {
		if opt == nil {
			return nil, &ConfigError{Field: "option", Reason: "nil option"}
		}
		if err := opt(&settings); err != nil {
			return nil, err
		}
	}
	if settings.Mode != auth.ModeUnset {
		return nil, &ConfigError{Field: "auth_mode", Reason: "auth mode is derived from credentials and cannot be set by options"}
	}
	settings.Mode = snapshot.mode
	settings.applyDefaults()
	if err := settings.validate(); err != nil {
		return nil, err
	}
	if err := settings.validateGateway(); err != nil {
		return nil, err
	}
	if err := settings.validateHTTPClient(); err != nil {
		return nil, err
	}
	settings.Gateway = compat.CloneURL(settings.Gateway)
	settings.ArtifactGateway = compat.CloneURL(settings.ArtifactGateway)
	return &Client{
		settings: settings,
		snapshot: snapshot,
		http:     signedHTTPClient(settings.HTTPClient),
		logger:   settings.Logger,
		now:      time.Now,
		nonce:    newNonce,
	}, nil
}

// newNonce 生成 128 位随机 hex 串，用作请求 nonce 与 request id。
func newNonce() string {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		// crypto/rand 失败属于系统级故障，返回空串让请求路径拒绝发送。
		return ""
	}
	return hex.EncodeToString(buf[:])
}

// Settings 返回公共配置的快照，用于诊断；不包含签名器、验签器或
// 凭证本身。修改快照不影响 Client，读取快照不写回字段。
func (c *Client) Settings() Settings {
	settings := c.settings
	settings.Gateway = compat.CloneURL(c.settings.Gateway)
	settings.ArtifactGateway = compat.CloneURL(c.settings.ArtifactGateway)
	return settings
}

// Mode 返回构造期绑定的认证模式。
func (c *Client) Mode() auth.Mode { return c.snapshot.mode }

// Do 实现 Executor：规格与认证要求校验 → 构建请求与签名头 →
// HTTP 执行 → 响应验签 → 错误分类。签名与协议失败优先于业务判断；
// 未验签的响应不进入成功路径。
//
// Do 返回错误时，结果仍可能携带已获得的响应元信息（HTTP 状态、响应头）。
// 非 2xx 且未带签名的错误响应以 Verified=false 的 *APIError 返回，
// 只能作为诊断信息。
func (c *Client) Do(ctx context.Context, req WireRequest) (*APIResult, error) {
	if ctx == nil {
		return nil, &ValidationError{Field: "ctx", Reason: "context must not be nil"}
	}
	spec := req.Spec
	if err := spec.validate(); err != nil {
		return nil, err
	}
	// 必须证书的接口在统一执行路径校验，自定义 Operation 无法绕过。
	if spec.AuthRequirement == AuthRequirementCertificate && c.snapshot.mode != auth.ModeCertificate {
		return nil, &ValidationError{Field: "auth_requirement",
			Reason: "operation requires certificate credentials; the client is configured for " + c.snapshot.mode.String()}
	}
	if err := req.Params.validate(spec.ParameterMode); err != nil {
		return nil, err
	}
	if req.Settings.NotifyURL != "" {
		return nil, &ValidationError{Field: "notify_url",
			Reason: "v3 request models carry notify_url as a business field; the notify_url call option only applies to page artifact building"}
	}

	start := c.now()
	result, body, err := c.send(ctx, spec, req.Params, req.Settings)
	attrs := append(opAttrs(spec.Method), durationAttr(c.now().Sub(start)))
	if err != nil {
		logRequestFailure(c.logger, ctx, err, attrs)
		return result, err
	}
	result, attrs, err = c.handleResponse(ctx, result, body, attrs)
	if err != nil {
		logRequestFailure(c.logger, ctx, err, attrs)
		return result, err
	}
	attrs = append(attrs,
		slog.Int("http_status", result.HTTPStatus))
	c.logger.LogAttrs(ctx, slog.LevelDebug, "alipay request succeeded", attrs...)
	return result, nil
}

// v3WireRequest 汇总签名与传输共用的请求要素：
// URI 与请求体是同一份序列化结果，签名后不再重新编码。
type v3WireRequest struct {
	method      string
	url         string
	body        []byte
	contentType string
	headers     http.Header
}

// buildWireRequest 组装请求要素与签名头。基于调用方输入的新对象构建，
// 不修改传入的 EncodedParams；签名使用调用方 ctx，可随请求取消。
func (c *Client) buildWireRequest(ctx context.Context, spec OperationSpec, params EncodedParams, settings CallSettings) (*v3WireRequest, error) {
	var body []byte
	var contentType string
	switch params.Mode {
	case ParameterModeJSONBody, ParameterModeJSONBodyWithQuery:
		if spec.HTTPMethod == HTTPMethodGet {
			return nil, &ValidationError{Field: "parameter_mode",
				Reason: "v3 GET requests must not carry a json body"}
		}
		body = []byte(params.JSON)
		contentType = "application/json"
	case ParameterModeQuery:
	case ParameterModeMultipart:
		if spec.HTTPMethod == http.MethodGet {
			return nil, &ValidationError{Field: "parameter_mode", Reason: "v3 GET requests must not carry a multipart body"}
		}
		body = params.Body
		contentType = params.ContentType
	default:
		return nil, &ValidationError{Field: "parameter_mode",
			Reason: "v3 request parameter mode is not supported"}
	}
	path, err := expandPathParams(spec.Path, params.PathParams)
	if err != nil {
		return nil, err
	}
	target, err := url.Parse(strings.TrimRight(c.settings.Gateway.String(), "/") + path)
	if err != nil {
		return nil, &ProtocolError{Op: "build", Reason: "invalid request path", Err: err}
	}
	target.RawQuery = params.Form.Encode()

	nonce := c.nonce()
	if nonce == "" {
		return nil, &ProtocolError{Op: "sign", Reason: "generate request nonce"}
	}
	timestamp := strconv.FormatInt(c.now().UnixMilli(), 10)
	authString := protocol.BuildV3AuthString(c.settings.AppID, c.snapshot.identity.AppCertificateSN, nonce, timestamp)
	signBody := body
	if params.Mode == ParameterModeMultipart {
		signBody = params.SignBody
	}
	content := protocol.BuildV3SignContent(authString, spec.HTTPMethod.String(), target.RequestURI(), string(signBody), settings.AppAuthToken)
	signature, err := c.snapshot.signer.Sign(ctx, content)
	if err != nil {
		if _, ok := compat.AsType[*auth.CertValidityError](err); ok {
			return nil, &CertificateError{Reason: "request certificate validity check failed", Err: err}
		}
		return nil, &SignatureError{Stage: "request", Reason: "sign request", Err: err}
	}

	headers := http.Header{}
	if len(body) != 0 {
		headers.Set("Content-Type", contentType)
	}
	headers.Set(protocol.V3HeaderAuthorization, protocol.V3AuthScheme+" "+authString+",sign="+signature.Value)
	if c.snapshot.mode == auth.ModeCertificate {
		headers.Set(protocol.V3HeaderRootCertSN, c.snapshot.identity.RootCertificateSN)
	}
	if settings.AppAuthToken != "" {
		headers.Set(protocol.V3HeaderAppAuthToken, settings.AppAuthToken)
	}
	headers.Set(protocol.V3HeaderRequestID, nonce)
	return &v3WireRequest{
		method:      spec.HTTPMethod.String(),
		url:         target.String(),
		body:        body,
		contentType: contentType,
		headers:     headers,
	}, nil
}

// expandPathParams replaces {name} path templates and rejects missing or
// unused values so a request cannot silently target an incomplete route.
func expandPathParams(path string, params map[string]string) (string, error) {
	var builder strings.Builder
	used := make(map[string]struct{}, len(params))
	for i := 0; i < len(path); {
		if path[i] != '{' {
			builder.WriteByte(path[i])
			i++
			continue
		}
		end := strings.IndexByte(path[i+1:], '}')
		if end < 0 {
			return "", &ValidationError{Field: "path", Reason: "path contains an unterminated parameter"}
		}
		end += i + 1
		name := path[i+1 : end]
		if name == "" {
			return "", &ValidationError{Field: "path", Reason: "path contains an empty parameter name"}
		}
		value, ok := params[name]
		if !ok || value == "" {
			return "", &ValidationError{Field: "path_params", Reason: "missing value for path parameter " + name}
		}
		builder.WriteString(url.PathEscape(value))
		used[name] = struct{}{}
		i = end + 1
	}
	for name := range params {
		if _, ok := used[name]; !ok {
			return "", &ValidationError{Field: "path_params", Reason: "unused path parameter " + name}
		}
	}
	return builder.String(), nil
}

// send 在请求超时约束下执行 HTTP 请求并读取响应体。
// 构造与传输失败不产生业务元信息；收到响应时返回状态、响应头与原始报文。
func (c *Client) send(ctx context.Context, spec OperationSpec, params EncodedParams, settings CallSettings) (*APIResult, []byte, error) {
	wire, err := c.buildWireRequest(ctx, spec, params, settings)
	if err != nil {
		return nil, nil, err
	}
	if c.settings.RequestTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.settings.RequestTimeout)
		defer cancel()
	}
	httpReq, err := http.NewRequestWithContext(ctx, wire.method, wire.url, bytes.NewReader(wire.body))
	if err != nil {
		return nil, nil, &ProtocolError{Op: "build", Reason: "build http request", Err: err}
	}
	for key, values := range wire.headers {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, nil, &TransportError{Op: "send", Err: err}
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, c.settings.MaxResponseBytes+1))
	if err != nil {
		return nil, nil, &TransportError{Op: "read", Err: err}
	}
	if int64(len(raw)) > c.settings.MaxResponseBytes {
		return nil, nil, &ProtocolError{Op: "read", Reason: "response exceeds size limit"}
	}
	result := &APIResult{
		HTTPStatus: resp.StatusCode,
		Header:     resp.Header.Clone(),
	}
	return result, raw, nil
}

// handleResponse 验签并分类响应。成功（2xx）必须携带
// alipay-signature 头；非 2xx 按官方契约解析错误体，响应带签名时
// 先验签，未带签名以未验签错误返回。
func (c *Client) handleResponse(ctx context.Context, result *APIResult, body []byte, attrs []slog.Attr) (*APIResult, []slog.Attr, error) {
	attrs = append(attrs, slog.Int("http_status", result.HTTPStatus))
	if result.HTTPStatus < 200 || result.HTTPStatus > 299 {
		return c.handleErrorResponse(ctx, result, body, attrs)
	}
	if len(body) == 0 {
		return result, attrs, &ProtocolError{Op: "extract", Reason: "empty response body"}
	}
	if err := c.verifyResponse(ctx, result, body); err != nil {
		return result, attrs, err
	}
	result.Node = append([]byte(nil), body...)
	return result, attrs, nil
}

// verifyResponse 校验 V3 响应头与签名，成功与错误响应共用此入口。
// 本 SDK 要求证书模式明确提供 alipay-sn。
// V3 验签字段见 https://opendocs.alipay.com/open-v3/054d0z。
func (c *Client) verifyResponse(ctx context.Context, result *APIResult, body []byte) error {
	signature := result.Header.Get(protocol.V3HeaderSignature)
	if signature == "" {
		return &SignatureError{Stage: "response", Reason: "missing alipay-signature header"}
	}
	for _, name := range []string{protocol.V3HeaderTimestamp, protocol.V3HeaderNonce} {
		if result.Header.Get(name) == "" {
			return &ProtocolError{Op: "verify", Reason: "missing " + name + " header"}
		}
	}
	keyID := result.Header.Get(protocol.V3HeaderCertSN)
	if c.snapshot.mode == auth.ModeCertificate && keyID == "" {
		return &CertificateError{Reason: "missing alipay-sn header", Err: &auth.CertSelectError{}}
	}
	input := auth.VerificationInput{
		Payload: protocol.BuildV3VerifyContent(
			result.Header.Get(protocol.V3HeaderTimestamp),
			result.Header.Get(protocol.V3HeaderNonce),
			body,
		),
		Signature: auth.Signature{
			Algorithm: auth.SignatureAlgorithmRSA2,
			Value:     signature,
		},
		// 公钥模式忽略 KeyID；证书模式按 alipay-sn 选择本地可信证书。
		KeyID: keyID,
	}
	if err := c.snapshot.verifier.Verify(ctx, input); err != nil {
		if _, ok := compat.AsType[*auth.CertSelectError](err); ok {
			return &CertificateError{Reason: "select trusted certificate by alipay-sn", Err: err}
		}
		if _, ok := compat.AsType[*auth.CertValidityError](err); ok {
			return &CertificateError{Reason: "response certificate validity check failed", Err: err}
		}
		return &SignatureError{Stage: "response", Reason: "verify signature", Err: err}
	}
	result.SignatureVerified = true
	return nil
}

// handleErrorResponse 处理非 2xx 响应：带签名时先验签（验签失败即
// 失败），再解析官方错误体；无签名或解析失败按未验签诊断处理。
func (c *Client) handleErrorResponse(ctx context.Context, result *APIResult, body []byte, attrs []slog.Attr) (*APIResult, []slog.Attr, error) {
	if result.Header.Get(protocol.V3HeaderSignature) != "" {
		if err := c.verifyResponse(ctx, result, body); err != nil {
			return result, attrs, err
		}
	}
	gw, err := protocol.ParseV3Error(body)
	if err != nil {
		return result, attrs, &HTTPError{Status: result.HTTPStatus, BodySnippet: bodySnippet(body)}
	}
	attrs = append(attrs, slog.String("code", gw.Code))
	return result, attrs, &APIError{
		Code:     ErrorCode(gw.Code),
		Msg:      gw.Message,
		Links:    gw.Links,
		Verified: result.SignatureVerified,
	}
}

// logRequestFailure 是统一的失败日志出口。
func logRequestFailure(logger *slog.Logger, ctx context.Context, err error, attrs []slog.Attr) {
	level := slog.LevelError
	if kind, ok := err.(ErrorKinder); ok {
		attrs = append(attrs, slog.String("error_kind", kind.Kind().String()))
		switch kind.Kind() {
		case ErrorKindAPI:
			level = slog.LevelWarn
		case ErrorKindCanceled:
			level = slog.LevelDebug
		}
	}
	logger.LogAttrs(ctx, level, "alipay request failed", attrs...)
}

func opAttrs(method string) []slog.Attr {
	return []slog.Attr{
		slog.String("sdk", "alipay"),
		slog.String("operation", method),
	}
}

func durationAttr(d time.Duration) slog.Attr {
	return slog.Duration("duration", d)
}

// bodySnippet 生成有界的诊断片段：截断到 512 字节并压缩空白，
// 避免输出超长报文。
func bodySnippet(body []byte) string {
	const limit = 512
	if len(body) > limit {
		body = body[:limit]
	}
	return strings.Join(strings.Fields(string(body)), " ")
}

// ErrorKinder 供错误类型暴露稳定分类；SDK 错误类型均已实现。
type ErrorKinder interface {
	Kind() ErrorKind
}

// 编译期保证 Client 满足执行器契约。
var _ Executor = (*Client)(nil)
