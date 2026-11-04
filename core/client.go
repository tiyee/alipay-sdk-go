package core

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tiyee/go-alipay/internal/protocol"
)

// Client 实现 Executor：公共参数、编码合并、签名、HTTP 执行、
// 原始响应提取、验签与错误分类。所有领域接口与自定义接口共用该流程。
//
// Do 返回错误时，结果仍可能携带已获得的响应元信息（HTTP 状态、响应头）。
// error_response 属于未签名报文，以 Verified=false 的 *APIError 返回，
// 只能作为诊断信息。
func (c *Client) Do(ctx context.Context, req WireRequest) (*APIResult, error) {
	if ctx == nil {
		return nil, &ValidationError{Field: "ctx", Reason: "context must not be nil"}
	}
	spec := req.Spec
	if err := spec.validate(); err != nil {
		return nil, err
	}
	if err := req.Params.validate(); err != nil {
		return nil, err
	}
	form, err := c.buildParams(spec, req.Params, req.Settings)
	if err != nil {
		return nil, err
	}

	start := c.now()
	result, body, err := c.send(ctx, form)
	attrs := append(opAttrs(spec.Method), durationAttr(c.now().Sub(start)))
	if err != nil {
		c.logFailure(ctx, err, attrs)
		return result, err
	}
	result, attrs, err = c.handleResponse(spec, result, body, attrs)
	if err != nil {
		c.logFailure(ctx, err, attrs)
		return result, err
	}
	attrs = append(attrs,
		slog.Int("http_status", result.HTTPStatus),
		slog.String("code", ResponseCodeSuccess.String()))
	c.logger.LogAttrs(ctx, slog.LevelDebug, "alipay request succeeded", attrs...)
	return result, nil
}

// Execute 实现 Go 1.27 泛型方法形式的调用入口。
// 它仅委托包级 Execute，不重复实现协议流程。
func (c *Client) Execute[Req, Resp any](ctx context.Context,
	operation Operation[Req, Resp], req Request[Req]) (*Response[Resp], error) {
	return Execute(ctx, c, operation, req)
}

// buildParams 组装请求参数。基于 Client 私有的新表单构建，
// 不修改调用方传入的 EncodedParams；空值业务参数不参与传输与签名。
func (c *Client) buildParams(spec OperationSpec, params EncodedParams, settings CallSettings) (url.Values, error) {
	form := url.Values{}
	form.Set(protocol.ParamAppID, c.settings.AppID)
	form.Set(protocol.ParamMethod, spec.Method)
	form.Set(protocol.ParamFormat, protocol.FormatJSON)
	form.Set(protocol.ParamCharset, protocol.CharsetUTF8)
	form.Set(protocol.ParamSignType, c.settings.Signer.Algorithm())
	form.Set(protocol.ParamTimestamp, c.now().Format(protocol.TimestampLayout))
	form.Set(protocol.ParamVersion, spec.Version)
	switch params.Mode {
	case ParameterModeBizJSON:
		form.Set(protocol.ParamBizContent, string(params.JSON))
	case ParameterModeForm:
		for key, values := range params.Form {
			if protocol.IsProtectedParam(key) {
				return nil, &ValidationError{Field: key,
					Reason: "form key overwrites a protected public parameter"}
			}
			for _, value := range values {
				if value != "" {
					form.Add(key, value)
				}
			}
		}
	}
	if settings.NotifyURL != "" {
		form.Set(protocol.ParamNotifyURL, settings.NotifyURL)
	}
	if settings.AppAuthToken != "" {
		form.Set(protocol.ParamAppAuthToken, settings.AppAuthToken)
	}
	return form, nil
}

// send 签名并执行 HTTP 请求，读取响应体。构造与传输失败
// 不产生业务元信息；收到响应时返回状态、响应头与原始报文。
func (c *Client) send(ctx context.Context, form url.Values) (*APIResult, []byte, error) {
	signContent := protocol.BuildSignContent(form)
	signature, err := c.settings.Signer.Sign([]byte(signContent))
	if err != nil {
		return nil, nil, &SignatureError{Stage: "request", Reason: "sign request", Err: err}
	}
	form.Set(protocol.ParamSign, signature)

	if c.settings.RequestTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.settings.RequestTimeout)
		defer cancel()
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.settings.Gateway.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, &ProtocolError{Op: "build", Reason: "build http request", Err: err}
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

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

// handleResponse 提取原始节点、验签并分类错误。签名与协议失败
// 优先于业务判断；未验签的响应不进入成功路径。
func (c *Client) handleResponse(spec OperationSpec, result *APIResult, body []byte, attrs []slog.Attr) (*APIResult, []slog.Attr, error) {
	attrs = append(attrs, slog.Int("http_status", result.HTTPStatus))
	if result.HTTPStatus != http.StatusOK {
		return result, attrs, &HTTPError{Status: result.HTTPStatus, BodySnippet: bodySnippet(body)}
	}
	raw, err := protocol.ExtractRawResponse(body, spec.ResponseKey)
	if err != nil {
		return result, attrs, &ProtocolError{Op: "extract", Reason: "extract response node", Err: err}
	}
	if raw.ErrorResponse != nil {
		gw, parseErr := protocol.ParseGatewayError(raw.ErrorResponse)
		if parseErr != nil {
			return result, attrs, &ProtocolError{Op: "extract",
				Reason: "parse error_response", Err: parseErr}
		}
		attrs = append(attrs, slog.String("code", gw.Code), slog.String("sub_code", gw.SubCode))
		return result, attrs, &APIError{
			Code:    ResponseCode(gw.Code),
			Msg:     gw.Msg,
			SubCode: gw.SubCode,
			SubMsg:  gw.SubMsg,
		}
	}
	if raw.Sign == "" {
		return result, attrs, &SignatureError{Stage: "response", Reason: "missing signature"}
	}
	if err := c.settings.Verifier.Verify(raw.Node, raw.Sign); err != nil {
		return result, attrs, &SignatureError{Stage: "response", Reason: "verify signature", Err: err}
	}
	result.SignatureVerified = true
	result.Node = raw.Node

	code, err := protocol.ParseResponseCode(raw.Node)
	if err != nil {
		return result, attrs, &ProtocolError{Op: "extract", Reason: "parse response code", Err: err}
	}
	if !acceptsCode(spec.AcceptedCodes, ResponseCode(code)) {
		gw, parseErr := protocol.ParseGatewayError(raw.Node)
		if parseErr != nil {
			return result, attrs, &ProtocolError{Op: "extract",
				Reason: "parse business error", Err: parseErr}
		}
		attrs = append(attrs, slog.String("code", gw.Code), slog.String("sub_code", gw.SubCode))
		return result, attrs, &APIError{
			Code:     ResponseCode(gw.Code),
			Msg:      gw.Msg,
			SubCode:  gw.SubCode,
			SubMsg:   gw.SubMsg,
			Verified: true,
		}
	}
	return result, attrs, nil
}

func acceptsCode(codes []ResponseCode, code ResponseCode) bool {
	for _, accepted := range codes {
		if accepted == code {
			return true
		}
	}
	return false
}

// logFailure 按错误分类选择日志级别：平台业务拒绝为 Warn，
// 主动取消为 Debug，网络、协议、安全与 HTTP 失败为 Error。
// 仅输出筛选后的诊断字段，不记录报文、URL、签名原文与错误文本。
func (c *Client) logFailure(ctx context.Context, err error, attrs []slog.Attr) {
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
	c.logger.LogAttrs(ctx, level, "alipay request failed", attrs...)
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
