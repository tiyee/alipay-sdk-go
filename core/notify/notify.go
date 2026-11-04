// Package notify 实现支付宝异步通知的解析、验签与类型化映射。
// 处理链路：限制请求体大小 → 解析一次表单 → 按通知策略构建待验签串 →
// 使用可信验签器校验 → 形成可信参数 → 映射领域通知模型。
// 本包不依赖业务领域包，不包含商户订单、幂等或状态机逻辑；
// 领域通知模型由 services/<domain> 定义并通过 Decoder 接入。
//
// 首期仅实现普通支付通知（交易状态同步）的验签策略；
// 生活号等通知对 sign_type 的剔除规则不同，属于未支持策略，
// 本包不会自动猜测通知类型。
package notify

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/auth"
	"github.com/tiyee/go-alipay/internal/protocol"
)

// Policy 标识异步通知的验签策略。不同通知类型构建待验签串时的
// 参数剔除规则不同，调用方必须显式选择策略；零值与其他未实现的
// 策略在解析前即被拒绝。
type Policy uint8

const (
	// PolicyTradeStatusSync 是统一收单交易支付完成异步通知
	// （普通支付通知，notify_type=trade_status_sync）的验签策略：
	// 构建待验签串时剔除 sign 与 sign_type 两个参数，其余参数
	// （含空值参数）全部参与验签。
	// 依据[自行实现验签](https://opendocs.alipay.com/common/02mse7.md)。
	// 生活号等通知需在待验签串中保留 sign_type，属于不同策略，
	// 本包未实现，不得使用本策略校验其他类型的通知。
	PolicyTradeStatusSync Policy = iota + 1
)

// String 实现 fmt.Stringer，返回策略对应的官方 notify_type 值。
func (p Policy) String() string {
	if t := p.notifyType(); t != "" {
		return t
	}
	return fmt.Sprintf("Policy(%d)", uint8(p))
}

func (p Policy) valid() bool { return p == PolicyTradeStatusSync }

// notifyType 返回策略对应的官方 notify_type 枚举值；未实现策略返回空串。
func (p Policy) notifyType() string {
	if p == PolicyTradeStatusSync {
		return "trade_status_sync"
	}
	return ""
}

// excludedParams 返回构建待验签串时剔除的参数名。
func (p Policy) excludedParams() []string {
	if p == PolicyTradeStatusSync {
		return []string{protocol.ParamSign, protocol.ParamSignType}
	}
	return nil
}

// Decoder 将已通过验签的通知参数映射为具体领域通知模型。
// 实现只应读取官方通知参数表中的字段；必选字段缺失时返回错误。
// 接口不含私有方法，领域包与调用方均可实现。
type Decoder[T any] interface {
	// Decode 接收解码一次且已通过验签的通知参数。
	Decode(values url.Values) (T, error)
}

// DecoderFunc 支持 func 值实现 Decoder。
type DecoderFunc[T any] func(values url.Values) (T, error)

// Decode 实现 Decoder。
func (f DecoderFunc[T]) Decode(values url.Values) (T, error) { return f(values) }

// Notification 是通过验签的异步通知结果。仅当 Parse 返回 error 为 nil
// 时 Data 非 nil；解析、验签或映射失败不输出可信业务数据。
type Notification[T any] struct {
	// Data 是类型化领域通知模型。
	Data *T
	// Params 是解码一次后的通知参数副本（不含 sign 与 sign_type）。
	// 领域模型尚未覆盖的新增通知字段可从这里读取；
	// 已知字段应使用 Data 的具名模型，不将 Params 作为业务数据的主要出口。
	Params url.Values
}

// DefaultMaxFormBytes 是通知表单的默认大小上限。
// 这是工程默认值，不是支付宝协议限制。
const DefaultMaxFormBytes = 64 << 10

// ParseOption 修改通知解析设置。
type ParseOption func(*parseSettings) error

type parseSettings struct {
	logger       *slog.Logger
	maxFormBytes int64
}

func newParseSettings(opts ...ParseOption) (parseSettings, error) {
	settings := parseSettings{
		logger:       slog.New(slog.DiscardHandler),
		maxFormBytes: DefaultMaxFormBytes,
	}
	for _, opt := range opts {
		if opt == nil {
			return parseSettings{}, &core.ValidationError{Field: "option", Reason: "nil option"}
		}
		if err := opt(&settings); err != nil {
			return parseSettings{}, err
		}
	}
	return settings, nil
}

// WithMaxFormBytes 设置通知表单大小上限，仅影响 ParseRequest 的请求体
// 读取；必须为正数。
func WithMaxFormBytes(n int64) ParseOption {
	return func(s *parseSettings) error {
		if n <= 0 {
			return &core.ValidationError{Field: "max_form_bytes", Reason: "must be positive"}
		}
		s.maxFormBytes = n
		return nil
	}
}

// WithLogger 注入结构化日志 Logger；显式传入 nil 视为配置错误。
// 未配置时静默。日志仅包含筛选后的诊断字段（sdk、operation、error_kind），
// 不记录通知参数值、签名原文或个人信息。
func WithLogger(logger *slog.Logger) ParseOption {
	return func(s *parseSettings) error {
		if logger == nil {
			return &core.ValidationError{Field: "logger",
				Reason: "nil logger; omit the option for silence"}
		}
		s.logger = logger
		return nil
	}
}

// Parse 基于已解码一次的通知表单参数完成验签与领域映射。
// values 应来自对原始请求体的一次解析（例如框架已解码的表单）；
// 本函数不会再次解码，也不修改调用方传入的 values。
// 重复参数属于歧义输入，一律拒绝，不静默选择首项或末项。
//
// 验证顺序：策略与依赖检查 → 重复参数 → charset → notify_type →
// sign_type → sign 存在性 → 验签 → 领域映射。任何一步失败都不输出
// 可信业务数据；验签失败返回 *core.SignatureError。
//
// 空值参数参与待验签串（与请求签名不同，官方验签文档未要求通知剔除空值）；
// charset 仅支持 utf-8；sign_type 与验签器算法不一致时拒绝，
// 缺失时按验签器算法（RSA2）验签。ctx 用于日志传递调用上下文。
func Parse[T any](ctx context.Context, policy Policy, verifier auth.Verifier,
	decoder Decoder[T], values url.Values, opts ...ParseOption) (Notification[T], error) {
	if ctx == nil {
		return Notification[T]{}, &core.ValidationError{Field: "ctx", Reason: "context must not be nil"}
	}
	settings, err := newParseSettings(opts...)
	if err != nil {
		return Notification[T]{}, err
	}
	data, params, err := parse(ctx, policy, verifier, decoder, values, settings)
	if err != nil {
		logFailure(ctx, settings.logger, err)
		return Notification[T]{}, err
	}
	settings.logger.LogAttrs(ctx, slog.LevelDebug, "alipay notify verified",
		slog.String("sdk", "alipay"), slog.String("operation", "notify"))
	return Notification[T]{Data: &data, Params: params}, nil
}

// ParseRequest 从原始 HTTP 请求解析异步通知：要求 POST 请求，
// 有界读取请求体并解码一次表单，然后执行与 Parse 相同的验签与映射流程。
// 调用方须保证请求体未被框架消费；框架已解析表单时应改用 Parse，
// 避免二次 URL 解码。通知参数只从请求体读取，不合并 URL 查询参数。
func ParseRequest[T any](ctx context.Context, r *http.Request, policy Policy,
	verifier auth.Verifier, decoder Decoder[T], opts ...ParseOption) (Notification[T], error) {
	if ctx == nil {
		return Notification[T]{}, &core.ValidationError{Field: "ctx", Reason: "context must not be nil"}
	}
	if r == nil {
		return Notification[T]{}, &core.ValidationError{Field: "request", Reason: "request is required"}
	}
	settings, err := newParseSettings(opts...)
	if err != nil {
		return Notification[T]{}, err
	}
	if r.Method != http.MethodPost {
		return Notification[T]{}, &core.ProtocolError{Op: "parse",
			Reason: "notify must be sent with POST"}
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, settings.maxFormBytes+1))
	if err != nil {
		return Notification[T]{}, &core.TransportError{Op: "read", Err: err}
	}
	if int64(len(body)) > settings.maxFormBytes {
		return Notification[T]{}, &core.ProtocolError{Op: "parse",
			Reason: "notify form exceeds size limit"}
	}
	form, err := url.ParseQuery(string(body))
	if err != nil {
		return Notification[T]{}, &core.ProtocolError{Op: "parse",
			Reason: "parse notify form", Err: err}
	}
	data, params, err := parse(ctx, policy, verifier, decoder, form, settings)
	if err != nil {
		logFailure(ctx, settings.logger, err)
		return Notification[T]{}, err
	}
	settings.logger.LogAttrs(ctx, slog.LevelDebug, "alipay notify verified",
		slog.String("sdk", "alipay"), slog.String("operation", "notify"))
	return Notification[T]{Data: &data, Params: params}, nil
}

// parse 执行验证与映射的主体流程，返回类型化数据与剔除签名参数后的
// 表单副本。失败路径不产出任何可信数据。
func parse[T any](ctx context.Context, policy Policy, verifier auth.Verifier,
	decoder Decoder[T], values url.Values, settings parseSettings) (T, url.Values, error) {
	var zero T
	if !policy.valid() {
		return zero, nil, &core.ValidationError{Field: "policy",
			Reason: "unknown notify policy; only the ordinary payment notification rule is supported"}
	}
	if verifier == nil {
		return zero, nil, &core.ValidationError{Field: "verifier", Reason: "verifier is required"}
	}
	if decoder == nil {
		return zero, nil, &core.ValidationError{Field: "decoder", Reason: "decoder is required"}
	}
	// Clone 返回深拷贝：后续读取与返回的 Params 不受调用方后续修改影响。
	form := values.Clone()
	var duplicated []string
	for key, vs := range form {
		if len(vs) > 1 {
			duplicated = append(duplicated, key)
		}
	}
	if len(duplicated) > 0 {
		slices.Sort(duplicated)
		return zero, nil, &core.ProtocolError{Op: "parse",
			Reason: fmt.Sprintf("duplicate notify param %q", duplicated[0])}
	}
	if cs := form.Get(protocol.ParamCharset); cs != "" && !strings.EqualFold(cs, protocol.CharsetUTF8) {
		return zero, nil, &core.ProtocolError{Op: "parse",
			Reason: "unsupported charset; only utf-8 is supported"}
	}
	if notifyType := form.Get(protocol.ParamNotifyType); notifyType == "" {
		return zero, nil, &core.ValidationError{Field: "notify_type",
			Reason: "notify_type is required"}
	} else if notifyType != policy.notifyType() {
		return zero, nil, &core.ValidationError{Field: "notify_type",
			Reason: "notify type does not match the selected policy"}
	}
	if signType := form.Get(protocol.ParamSignType); signType != "" && signType != verifier.Algorithm() {
		return zero, nil, &core.SignatureError{Stage: "notify",
			Reason: "unsupported sign type; only " + verifier.Algorithm() + " is supported"}
	}
	sign := form.Get(protocol.ParamSign)
	if sign == "" {
		return zero, nil, &core.SignatureError{Stage: "notify", Reason: "missing signature"}
	}
	content := protocol.BuildNotifySignContent(form, policy.excludedParams()...)
	if err := verifier.Verify([]byte(content), sign); err != nil {
		return zero, nil, &core.SignatureError{Stage: "notify",
			Reason: "verify notify signature", Err: err}
	}
	delete(form, protocol.ParamSign)
	delete(form, protocol.ParamSignType)
	data, err := decoder.Decode(form)
	if err != nil {
		return zero, nil, &core.ProtocolError{Op: "decode",
			Reason: "decode notify params", Err: err}
	}
	return data, form, nil
}

// logFailure 按统一字段约定记录通知处理失败。
// 安全与协议失败均为 Error 级别；仅输出分类字段，不记录参数值与签名。
func logFailure(ctx context.Context, logger *slog.Logger, err error) {
	attrs := []slog.Attr{
		slog.String("sdk", "alipay"),
		slog.String("operation", "notify"),
	}
	if kind, ok := err.(core.ErrorKinder); ok {
		attrs = append(attrs, slog.String("error_kind", kind.Kind().String()))
	}
	logger.LogAttrs(ctx, slog.LevelError, "alipay notify failed", attrs...)
}
