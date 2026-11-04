package core

import (
	"context"
	"encoding/json"
	"encoding/json/jsontext"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tiyee/go-alipay/core/auth"
)

// Signer 与 Verifier 的契约定义在 core/auth；
// 此处提供类型别名，调用方仅导入 core 即可获得完整配置契约。
type (
	Signer   = auth.Signer
	Verifier = auth.Verifier
)

// ResponseCode 是支付宝公共响应码，对应响应节点中的 code 字段。
// 值依据官方公共错误码说明（https://opendocs.alipay.com/common/02km9f.md）。
type ResponseCode string

// ResponseCodeSuccess 表示接口调用成功。
const ResponseCodeSuccess ResponseCode = "10000"

// String 实现 fmt.Stringer。
func (c ResponseCode) String() string { return string(c) }

// ParameterMode 表示业务参数进入网关请求的编码位置。
type ParameterMode uint8

const (
	// ParameterModeBizJSON 表示业务参数由编码器序列化为 biz_content JSON。
	// 大多数接口属于此模式。
	ParameterModeBizJSON ParameterMode = iota + 1
	// ParameterModeForm 表示业务参数直接作为请求参数发送，不封装进
	// biz_content；依据[官方签名说明](https://opendocs.alipay.com/common/057k53.md)，
	// 参数形式由具体接口决定。
	ParameterModeForm
)

// String 实现 fmt.Stringer。
func (m ParameterMode) String() string {
	switch m {
	case ParameterModeBizJSON:
		return "biz_json"
	case ParameterModeForm:
		return "form"
	default:
		return fmt.Sprintf("ParameterMode(%d)", uint8(m))
	}
}

func (m ParameterMode) valid() bool {
	return m == ParameterModeBizJSON || m == ParameterModeForm
}

// OperationSpec 描述一次网关接口调用的协议元信息。
// 所有字段都是协议层的确定描述；接口名、版本与响应节点名必须与官方文档一致。
type OperationSpec struct {
	// Method 是官方接口方法名，例如 "alipay.trade.query"。
	Method string
	// Version 是接口版本，例如 "1.0"。
	Version string
	// ResponseKey 是响应 JSON 中的业务节点名，例如
	// "alipay_trade_query_response"。
	ResponseKey string
	// ParameterMode 决定业务参数的编码位置。
	ParameterMode ParameterMode
	// AcceptedCodes 是视为调用成功的公共响应码；构造 Operation 时
	// 为空会默认为 [ResponseCodeSuccess]。响应验签通过后，
	// code 不在列表内将作为 APIError 返回。
	AcceptedCodes []ResponseCode
}

func (s OperationSpec) validate() error {
	switch {
	case s.Method == "":
		return &ValidationError{Field: "method", Reason: "method is required"}
	case s.Version == "":
		return &ValidationError{Field: "version", Reason: "version is required"}
	case s.ResponseKey == "":
		return &ValidationError{Field: "response_key", Reason: "response key is required"}
	case !s.ParameterMode.valid():
		return &ValidationError{Field: "parameter_mode", Reason: "parameter mode is not set"}
	}
	for _, code := range s.AcceptedCodes {
		if code == "" {
			return &ValidationError{Field: "accepted_codes", Reason: "accepted code must not be empty"}
		}
	}
	return nil
}

// clone 返回规格副本，并复制切片，避免共享可变数据。
func (s OperationSpec) clone() OperationSpec {
	if s.AcceptedCodes == nil {
		return s
	}
	codes := make([]ResponseCode, len(s.AcceptedCodes))
	copy(codes, s.AcceptedCodes)
	s.AcceptedCodes = codes
	return s
}

// EncodedParams 是编码器输出的业务参数。字段按 Mode 互斥使用：
// BizJSON 模式只允许 JSON，Form 模式只允许 Form。
// JSON 使用 jsontext.Value 保存唯一一份编码结果，同时用于签名与传输；
// Form 在协议边界使用 url.Values。
type EncodedParams struct {
	// Mode 是业务参数的编码模式。
	Mode ParameterMode
	// JSON 是 biz_content 的原始 JSON 字节，仅在 BizJSON 模式下有效。
	JSON jsontext.Value
	// Form 是直接参数，仅在 Form 模式下有效。
	Form url.Values
}

func (p EncodedParams) validate() error {
	switch p.Mode {
	case ParameterModeBizJSON:
		if len(p.JSON) == 0 {
			return &ValidationError{Field: "params", Reason: "biz json mode requires encoded json"}
		}
		if len(p.Form) != 0 {
			return &ValidationError{Field: "params", Reason: "form values must be empty in biz json mode"}
		}
	case ParameterModeForm:
		if len(p.JSON) != 0 {
			return &ValidationError{Field: "params", Reason: "json must be empty in form mode"}
		}
		for key, values := range p.Form {
			if len(values) > 1 {
				return &ValidationError{Field: "params", Reason: fmt.Sprintf("form key %q has multiple values", key)}
			}
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

// Decoder 将已通过验签的业务响应节点解码为具体响应类型。
// 解码器只处理验证后的业务节点，不接触完整原始响应体。
type Decoder[T any] interface {
	Decode(payload jsontext.Value) (T, error)
}

// EncoderFunc 支持 func 值实现 Encoder。
type EncoderFunc[T any] func(value T) (EncodedParams, error)

// Encode 实现 Encoder。
func (f EncoderFunc[T]) Encode(value T) (EncodedParams, error) { return f(value) }

// DecoderFunc 支持 func 值实现 Decoder。
type DecoderFunc[T any] func(payload jsontext.Value) (T, error)

// Decode 实现 Decoder。
func (f DecoderFunc[T]) Decode(payload jsontext.Value) (T, error) { return f(payload) }

// JSONEncoder 返回使用 encoding/json/v2 的默认编码器，
// 将请求 struct 序列化为 biz_content JSON。
func JSONEncoder[T any]() Encoder[T] {
	return EncoderFunc[T](func(value T) (EncodedParams, error) {
		raw, err := json.Marshal(value)
		if err != nil {
			return EncodedParams{}, fmt.Errorf("alipay: encode biz_content: %w", err)
		}
		return EncodedParams{Mode: ParameterModeBizJSON, JSON: jsontext.Value(raw)}, nil
	})
}

// JSONDecoder 返回使用 encoding/json/v2 的默认解码器，
// 将业务响应节点解码为具体类型。
func JSONDecoder[T any]() Decoder[T] {
	return DecoderFunc[T](func(payload jsontext.Value) (T, error) {
		var value T
		if err := json.Unmarshal(payload, &value); err != nil {
			var zero T
			return zero, fmt.Errorf("alipay: decode response node: %w", err)
		}
		return value, nil
	})
}

// Operation 将请求与响应类型在编译期配对，并绑定接口描述与编解码器。
// 字段不导出；通过 NewJSONOperation 或 NewOperation 构造，
// 零值 Operation 在发送请求前失败。
type Operation[Req, Resp any] struct {
	spec    OperationSpec
	encoder Encoder[Req]
	decoder Decoder[Resp]
}

// NewJSONOperation 使用默认 JSON 编解码器构造 Operation，
// 适用于业务参数放入 biz_content 的接口。
func NewJSONOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error) {
	return newOperation(spec, JSONEncoder[Req](), JSONDecoder[Resp]())
}

// NewOperation 使用调用方提供的编解码器构造 Operation，
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
	if len(spec.AcceptedCodes) == 0 {
		spec.AcceptedCodes = []ResponseCode{ResponseCodeSuccess}
	}
	return Operation[Req, Resp]{spec: spec.clone(), encoder: encoder, decoder: decoder}, nil
}

// Spec 返回接口描述的副本；修改副本不影响 Operation。
func (op Operation[Req, Resp]) Spec() OperationSpec {
	return op.spec.clone()
}

func (op Operation[Req, Resp]) validate() (OperationSpec, error) {
	if op.encoder == nil || op.decoder == nil {
		return OperationSpec{}, &ValidationError{Field: "operation",
			Reason: "operation is not initialized; use NewJSONOperation or NewOperation"}
	}
	return op.spec.clone(), nil
}

// CallSettings 保存仅作用于单次调用的设置。
// 代调用令牌、通知地址等调用级参数不属于 Client 共享配置，
// 以避免并发调用相互污染。
type CallSettings struct {
	// NotifyURL 是本次请求的异步通知地址（公共参数 notify_url）。
	NotifyURL string
	// AppAuthToken 是代调用时的应用授权令牌（公共参数 app_auth_token）。
	AppAuthToken string
}

// CallOption 修改单次调用的 CallSettings。
type CallOption func(*CallSettings)

// WithNotifyURL 设置本次调用的异步通知地址。
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
// Node 是已通过验签的业务响应节点原文（jsontext.Value），
// 属于敏感诊断数据：不应写入日志或错误信息，正常使用应通过
// 类型化解码获得业务数据。Client 已读取并关闭响应体，
// 本结构不保证底层响应体仍可读取。
type APIResult struct {
	// HTTPStatus 是响应 HTTP 状态码。
	HTTPStatus int
	// Header 是响应头的副本。
	Header http.Header
	// SignatureVerified 标识业务节点是否通过签名验证。
	SignatureVerified bool
	// Node 是业务响应节点的原始 JSON 字节，仅验签成功后填充。
	Node jsontext.Value
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
