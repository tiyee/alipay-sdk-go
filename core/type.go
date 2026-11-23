package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/tiyee/alipay-sdk-go/internal/jsoncodec"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
)

// AuthRequirement 声明 Operation 对认证模式的要求。公共 Client 类型
// 不区分公钥与证书，必须证书的接口依靠执行前校验；内置 Operation 按
// 官方契约填写要求，自定义 Operation 同样必须显式声明。
type AuthRequirement uint8

const (
	// AuthRequirementAny 表示公钥与证书模式均可调用。
	AuthRequirementAny AuthRequirement = iota + 1
	// AuthRequirementCertificate 表示该接口要求证书模式调用；
	// 认证快照不满足时请求不发送。
	AuthRequirementCertificate
)

// String 实现 fmt.Stringer。
func (r AuthRequirement) String() string {
	switch r {
	case AuthRequirementAny:
		return "any"
	case AuthRequirementCertificate:
		return "certificate"
	default:
		return fmt.Sprintf("AuthRequirement(%d)", uint8(r))
	}
}

func (r AuthRequirement) valid() bool {
	return r == AuthRequirementAny || r == AuthRequirementCertificate
}

// HTTPMethod 是 V3 REST 接口的 HTTP 方法，取值与官方 OpenAPI 一致。
type HTTPMethod string

const (
	// HTTPMethodGet 是 GET 请求，业务参数置于 query。
	HTTPMethodGet HTTPMethod = "GET"
	// HTTPMethodPost 是 POST 请求，业务参数位置由 OperationSpec 指定。
	HTTPMethodPost HTTPMethod = "POST"
	// HTTPMethodPut 是 PUT 请求。
	HTTPMethodPut HTTPMethod = "PUT"
	// HTTPMethodDelete 是 DELETE 请求。
	HTTPMethodDelete HTTPMethod = "DELETE"
	// HTTPMethodPatch 是 PATCH 请求。
	HTTPMethodPatch HTTPMethod = "PATCH"
)

// String 实现 fmt.Stringer。
func (m HTTPMethod) String() string { return string(m) }

func (m HTTPMethod) valid() bool {
	switch m {
	case HTTPMethodGet, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodPatch:
		return true
	default:
		return false
	}
}

// ParameterMode 表示业务参数进入 V3 请求的编码位置或编码格式。
type ParameterMode uint8

const (
	// ParameterModeJSONBody 表示业务参数由编码器序列化为 V3 JSON
	// 请求体；GET 接口不使用该模式。
	ParameterModeJSONBody ParameterMode = iota + 1
	// ParameterModeQuery 表示业务参数置于 V3 请求 query；GET 接口
	// 必须使用该模式。
	ParameterModeQuery
	// ParameterModeMultipart 表示业务参数采用 V3 multipart/form-data 请求体。
	ParameterModeMultipart
	// ParameterModeJSONBodyWithQuery 将 query 标签字段写入 URI，独立编码 JSON 正文。
	ParameterModeJSONBodyWithQuery
)

// String 实现 fmt.Stringer。
func (m ParameterMode) String() string {
	switch m {
	case ParameterModeJSONBody:
		return "json_body"
	case ParameterModeJSONBodyWithQuery:
		return "json_body_with_query"
	case ParameterModeQuery:
		return "query"
	case ParameterModeMultipart:
		return "multipart"
	default:
		return fmt.Sprintf("ParameterMode(%d)", uint8(m))
	}
}

// OperationSpec 描述一次 V3 REST 接口调用的协议元信息。
// 所有字段都是协议层的确定描述；路径与参数位置必须与官方
// V3 OpenAPI 一致，零值或无效取值在构造 Operation 时拒绝。
type OperationSpec struct {
	// Method 是官方接口名，例如 "alipay.trade.query"，用于日志与诊断。
	Method string
	// HTTPMethod 是 V3 接口的 HTTP 方法。
	HTTPMethod HTTPMethod
	// Path 是 V3 相对路由，例如 "/v3/alipay/trade/query"；
	// 必须以 /v3/ 开头且不含域名与 query。
	Path string
	// ParameterMode 决定业务参数的编码位置。
	ParameterMode ParameterMode
	// AuthRequirement 是对认证模式的要求；必须显式声明：
	// 未设置或未知取值在构造 Operation 时拒绝。证书要求在统一
	// 执行路径校验，不因 Executor 或自定义 API 绕过。
	AuthRequirement AuthRequirement
}

func (s OperationSpec) validate() error {
	switch {
	case s.Method == "":
		return &ValidationError{Field: "method", Reason: "method is required"}
	case !s.HTTPMethod.valid():
		return &ValidationError{Field: "http_method", Reason: "http method is required"}
	case s.Path == "":
		return &ValidationError{Field: "path", Reason: "path is required"}
	case !strings.HasPrefix(s.Path, protocol.V3PathPrefix):
		return &ValidationError{Field: "path", Reason: "v3 path must start with " + protocol.V3PathPrefix}
	case strings.ContainsAny(s.Path, "?#"):
		return &ValidationError{Field: "path", Reason: "v3 path must not carry query or fragment"}
	}
	if s.HTTPMethod == HTTPMethodGet && s.ParameterMode != ParameterModeQuery {
		return &ValidationError{Field: "parameter_mode",
			Reason: "v3 GET requests carry parameters in the query"}
	}
	if _, ok := validModes[s.ParameterMode]; !ok {
		return &ValidationError{Field: "parameter_mode",
			Reason: "parameter mode is not valid for the v3 protocol"}
	}
	if !s.AuthRequirement.valid() {
		return &ValidationError{Field: "auth_requirement",
			Reason: "auth requirement must be declared; set it to a known value"}
	}
	return nil
}

// validModes 是 V3 协议允许的参数模式。
var validModes = map[ParameterMode]struct{}{
	ParameterModeJSONBody:          {},
	ParameterModeJSONBodyWithQuery: {},
	ParameterModeQuery:             {},
	ParameterModeMultipart:         {},
}

// EncodedParams 是编码器输出的业务参数。Mode 决定可用字段。
// JSONBodyWithQuery 模式同时携带 JSON 与 Form。
// JSON 使用 json.RawMessage 保存唯一一份编码结果，同时用于签名与传输；
// Form 在协议边界使用 url.Values。
type EncodedParams struct {
	// Mode 是业务参数的编码模式。
	Mode ParameterMode
	// JSON 是 V3 请求体的原始 JSON 字节，适用于两种 JSON 正文模式。
	JSON json.RawMessage
	// Form 是 V3 query 参数，适用于 Query 和 JSONBodyWithQuery 模式。
	Form url.Values
	// Body 是 multipart/form-data 模式下的完整请求体。
	Body []byte
	// SignBody 是 multipart 请求中参与签名的业务 JSON 原文；文件内容不参与签名。
	SignBody []byte
	// ContentType 是 multipart 请求体的 Content-Type，必须包含 boundary。
	ContentType string
	// PathParams 替换 OperationSpec.Path 中的 {name} 路径参数。
	// 未使用路径参数的 Operation 留空；路径模板与参数必须一一对应。
	PathParams map[string]string
}

func (p EncodedParams) validate(mode ParameterMode) error {
	if p.Mode != mode {
		return &ValidationError{Field: "parameter_mode", Reason: "encoded parameters do not match the operation"}
	}
	for name, value := range p.PathParams {
		if name == "" || value == "" {
			return &ValidationError{Field: "path_params", Reason: "path parameter names and values must not be empty"}
		}
	}
	switch p.Mode {
	case ParameterModeJSONBody, ParameterModeJSONBodyWithQuery:
		if len(p.SignBody) != 0 {
			return &ValidationError{Field: "params", Reason: "multipart sign body is only valid in multipart mode"}
		}
		if len(p.JSON) == 0 {
			return &ValidationError{Field: "params", Reason: "json body mode requires encoded json"}
		}
		if p.Mode == ParameterModeJSONBody && len(p.Form) != 0 {
			return &ValidationError{Field: "params", Reason: "form values must be empty in json body mode"}
		}
		if len(p.Body) != 0 || p.ContentType != "" {
			return &ValidationError{Field: "params", Reason: "multipart data must be empty in json body mode"}
		}
	case ParameterModeQuery:
		if len(p.SignBody) != 0 {
			return &ValidationError{Field: "params", Reason: "multipart sign body is only valid in multipart mode"}
		}
		if len(p.JSON) != 0 {
			return &ValidationError{Field: "params", Reason: "json must be empty in query mode"}
		}
		if len(p.Body) != 0 || p.ContentType != "" {
			return &ValidationError{Field: "params", Reason: "multipart data must be empty in query mode"}
		}
	case ParameterModeMultipart:
		if len(p.JSON) != 0 || len(p.Form) != 0 {
			return &ValidationError{Field: "params", Reason: "json and query values must be empty in multipart mode"}
		}
		if len(p.Body) == 0 || !strings.HasPrefix(p.ContentType, "multipart/form-data; boundary=") {
			return &ValidationError{Field: "params", Reason: "multipart body and content type with boundary are required"}
		}
	default:
		return &ValidationError{Field: "params", Reason: "parameter mode is not set"}
	}
	return nil
}

// Encoder 将具体请求类型编码为协议参数。实现必须输出确定性的编码结果：
// 同一输入产生的字节同时用于签名与传输，编码后不会重新序列化。
type Encoder[T any] interface {
	Encode(value T) (EncodedParams, error)
}

// Decoder 将已通过验签的业务响应解码为具体响应类型。
// 解码器只处理验证后的业务数据，不接触完整原始响应体。
type Decoder[T any] interface {
	Decode(payload json.RawMessage) (T, error)
}

// EncoderFunc 支持 func 值实现 Encoder。
type EncoderFunc[T any] func(value T) (EncodedParams, error)

// Encode 实现 Encoder。
func (f EncoderFunc[T]) Encode(value T) (EncodedParams, error) { return f(value) }

// DecoderFunc 支持 func 值实现 Decoder。
type DecoderFunc[T any] func(payload json.RawMessage) (T, error)

// Decode 实现 Decoder。
func (f DecoderFunc[T]) Decode(payload json.RawMessage) (T, error) { return f(payload) }

// JSONBodyEncoder 返回 V3 请求体编码器，将请求 struct 序列化为
// JSON 请求体。编码结果同时用于签名与传输，不会重新序列化。
func JSONBodyEncoder[T any]() Encoder[T] { return jsonBodyEncoder[T](ParameterModeJSONBody) }

func jsonBodyEncoder[T any](mode ParameterMode) Encoder[T] {
	return EncoderFunc[T](func(value T) (EncodedParams, error) {
		pathParams, err := extractPathParams(value)
		if err != nil {
			return EncodedParams{}, err
		}
		raw, err := json.Marshal(value)
		if err != nil {
			return EncodedParams{}, fmt.Errorf("alipay: encode request body: %w", err)
		}
		params := EncodedParams{Mode: mode, JSON: json.RawMessage(raw), PathParams: pathParams}
		if mode == ParameterModeJSONBodyWithQuery {
			params.Form, err = encodeQuery(value)
			if err != nil {
				return EncodedParams{}, err
			}
		}
		return params, nil
	})
}

// JSONDecoder 使用当前工具链对应的严格 JSON 解码器。
// 将已验签的响应数据解码为具体类型。
func JSONDecoder[T any]() Decoder[T] {
	return DecoderFunc[T](func(payload json.RawMessage) (T, error) {
		var value T
		if reflect.TypeFor[T]().Kind() == reflect.Struct && bytes.Equal(bytes.TrimSpace(payload), []byte("null")) {
			return value, fmt.Errorf("alipay: expected a response object, got null")
		}
		if err := jsoncodec.Unmarshal(payload, &value); err != nil {
			var zero T
			return zero, fmt.Errorf("alipay: decode response: %w", err)
		}
		return value, nil
	})
}

// Operation 将请求与响应类型在编译期配对，并绑定接口描述与编解码器。
// 字段不导出；通过 NewJSONOperation 或 NewOperation 构造，构造后的
// Operation 携带确定的接口描述；零值 Operation 在发送请求前失败。
type Operation[Req, Resp any] struct {
	spec    OperationSpec
	encoder Encoder[Req]
	decoder Decoder[Resp]
}

// NewJSONOperation 使用默认编解码器构造 V3 Operation：
// 解码器固定为 JSONDecoder；编码器按规格选择，GET 接口使用
// query 编码，JSON 接口序列化请求体，multipart 接口使用
// MultipartBodyEncoder。
// 规格未声明 ParameterMode 时按 HTTPMethod 推断。
func NewJSONOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error) {
	if spec.ParameterMode == 0 {
		if spec.HTTPMethod == HTTPMethodGet {
			spec.ParameterMode = ParameterModeQuery
		} else {
			spec.ParameterMode = ParameterModeJSONBody
		}
	}
	switch spec.ParameterMode {
	case ParameterModeQuery:
		return newOperation(spec, QueryEncoder[Req](), JSONDecoder[Resp]())
	case ParameterModeJSONBodyWithQuery:
		return newOperation(spec, jsonBodyEncoder[Req](spec.ParameterMode), JSONDecoder[Resp]())
	case ParameterModeJSONBody:
		return newOperation(spec, JSONBodyEncoder[Req](), JSONDecoder[Resp]())
	case ParameterModeMultipart:
		return newOperation(spec, MultipartBodyEncoder[Req](), JSONDecoder[Resp]())
	default:
		return Operation[Req, Resp]{}, &ValidationError{Field: "parameter_mode",
			Reason: "v3 spec requires a supported parameter mode"}
	}
}

// NewOperation 使用调用方提供的编解码器构造 V3 Operation，
// 适用于直接参数编码或需要自定义字段映射的接口。
// 编解码器接口不含私有方法，外部包可以直接实现。
func NewOperation[Req, Resp any](spec OperationSpec, encoder Encoder[Req], decoder Decoder[Resp]) (Operation[Req, Resp], error) {
	return newOperation(spec, encoder, decoder)
}

func newOperation[Req, Resp any](spec OperationSpec, encoder Encoder[Req], decoder Decoder[Resp]) (Operation[Req, Resp], error) {
	if encoder == nil {
		return Operation[Req, Resp]{}, &ValidationError{Field: "encoder", Reason: "encoder is required"}
	}
	if decoder == nil {
		return Operation[Req, Resp]{}, &ValidationError{Field: "decoder", Reason: "decoder is required"}
	}
	if err := spec.validate(); err != nil {
		return Operation[Req, Resp]{}, err
	}
	return Operation[Req, Resp]{spec: spec, encoder: encoder, decoder: decoder}, nil
}

// Spec 返回接口描述的副本；修改副本不影响 Operation。
func (op Operation[Req, Resp]) Spec() OperationSpec {
	return op.spec
}

func (op Operation[Req, Resp]) validate() (OperationSpec, error) {
	if op.encoder == nil || op.decoder == nil {
		return OperationSpec{}, &ValidationError{Field: "operation",
			Reason: "operation is not initialized; use NewJSONOperation or NewOperation"}
	}
	return op.spec, nil
}

// CallSettings 保存仅作用于单次调用的设置。
// 代调用令牌、通知地址等调用级参数不属于 Client 共享配置，
// 以避免并发调用相互污染。
//
// NotifyURL 供页面/唤端产物构建使用；V3 REST 请求模型以业务字段
// 携带 notify_url，REST 执行路径拒绝该选项。
type CallSettings struct {
	// NotifyURL 是本次请求的异步通知地址（公共参数 notify_url）。
	NotifyURL string
	// AppAuthToken 是代调用时的应用授权令牌（公共参数 app_auth_token）。
	AppAuthToken string
}

// CallOption 修改单次调用的 CallSettings。
type CallOption func(*CallSettings)

// WithNotifyURL 设置本次调用的异步通知地址；仅产物构建路径消费该值。
func WithNotifyURL(rawURL string) CallOption {
	return func(s *CallSettings) { s.NotifyURL = rawURL }
}

// WithAppAuthToken 设置本次调用的代调用授权令牌。
func WithAppAuthToken(token string) CallOption {
	return func(s *CallSettings) { s.AppAuthToken = token }
}

// NewCallSettings 从选项构建调用设置；重复设置时后一个生效。
func NewCallSettings(opts ...CallOption) CallSettings {
	var settings CallSettings
	for _, opt := range opts {
		if opt != nil {
			opt(&settings)
		}
	}
	return settings
}

// Request 是一次调用的类型化输入。Body 是具体业务请求 struct；
// Options 是本次调用设置，不承载签名与公共身份覆盖值。
type Request[T any] struct {
	// Body 是具体业务请求模型。
	Body T
	// Options 是本次调用的设置。
	Options CallSettings
}

// Response 是一次调用的结果封装。调用方必须先检查 error：
// 仅当 error 为 nil 时 Data 非 nil；失败时 Result 仍可能携带已获得的元信息。
type Response[T any] struct {
	// Data 是类型化业务数据，仅在调用成功时非 nil。
	Data *T
	// Result 是响应元信息，可能为 nil（例如请求未发出）。
	Result *APIResult
}

// APIResult 携带响应的协议元信息。
//
// Node 是已通过验签的响应体原文（json.RawMessage），属于敏感诊断数据：
// 不应写入日志或错误信息，正常使用应通过类型化解码获得业务数据。
// Client 已读取并关闭响应体，本结构不保证底层响应体仍可读取。
type APIResult struct {
	// HTTPStatus 是响应 HTTP 状态码。
	HTTPStatus int
	// Header 是响应头的副本。
	Header http.Header
	// SignatureVerified 标识响应体是否通过签名验证。
	SignatureVerified bool
	// Node 是响应体的原始 JSON 字节，仅验签成功后填充。
	Node json.RawMessage
}

// WireRequest 是执行器接收的协议层请求，只包含已编码的协议数据，
// 不接收动态业务对象。
type WireRequest struct {
	// Spec 是接口描述。
	Spec OperationSpec
	// Params 是已编码的业务参数。
	Params EncodedParams
	// Settings 是本次调用的设置。
	Settings CallSettings
}

// Executor 执行已编码的协议请求并返回验证后的元信息。
// 领域 Service 与测试替身依赖该非泛型接口；泛型调用入口由
// Execute 与 Client.Execute 提供，两者共享同一实现。
type Executor interface {
	Do(ctx context.Context, req WireRequest) (*APIResult, error)
}
