package core

import (
	"context"
	"errors"
	"fmt"
)

// ErrorKind 标识 SDK 错误的稳定分类，供日志字段 error_kind 与监控使用。
// 调用方应通过 errors.As 匹配具体错误类型，ErrorKind 仅用于分类统计。
type ErrorKind string

const (
	// ErrorKindConfig 表示 Client 配置无效。
	ErrorKindConfig ErrorKind = "config"
	// ErrorKindValidation 表示请求参数未通过本地校验，请求未发送。
	ErrorKindValidation ErrorKind = "validation"
	// ErrorKindTransport 表示网络传输失败，保留原始错误链。
	ErrorKindTransport ErrorKind = "transport"
	// ErrorKindCanceled 表示调用方主动取消。
	ErrorKindCanceled ErrorKind = "canceled"
	// ErrorKindHTTP 表示网关返回非成功 HTTP 状态。
	ErrorKindHTTP ErrorKind = "http"
	// ErrorKindProtocol 表示响应格式异常、预期节点缺失或响应超限。
	ErrorKindProtocol ErrorKind = "protocol"
	// ErrorKindSignature 表示请求签名失败或响应缺签、验签失败。
	ErrorKindSignature ErrorKind = "signature"
	// ErrorKindCertificate 表示证书装载、匹配或选择失败。
	ErrorKindCertificate ErrorKind = "certificate"
	// ErrorKindAPI 表示支付宝网关返回的业务错误（含未验签的错误响应）。
	ErrorKindAPI ErrorKind = "api"
)

// String 实现 fmt.Stringer。
func (k ErrorKind) String() string { return string(k) }

// ConfigError 表示 Client 初始化配置无效。
// 错误信息只指出无效字段与原因，不输出字段值，避免泄露密钥等敏感配置。
type ConfigError struct {
	// Field 是无效的配置项名称。
	Field string
	// Reason 说明无效原因。
	Reason string
	// Err 是底层错误，可以为 nil。
	Err error
}

// Error 实现 error 接口。
func (e *ConfigError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("alipay: invalid config field %q: %s: %v", e.Field, e.Reason, e.Err)
	}
	return fmt.Sprintf("alipay: invalid config field %q: %s", e.Field, e.Reason)
}

// Unwrap 支持标准错误链。
func (e *ConfigError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder。
func (e *ConfigError) Kind() ErrorKind { return ErrorKindConfig }

// ValidationError 表示请求参数未通过本地校验，请求不会发送到网关。
type ValidationError struct {
	// Field 是无效的参数或请求字段名称。
	Field string
	// Reason 说明无效原因。
	Reason string
	// Err 是底层错误，可以为 nil。
	Err error
}

// Error 实现 error 接口。
func (e *ValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("alipay: invalid request field %q: %s: %v", e.Field, e.Reason, e.Err)
	}
	return fmt.Sprintf("alipay: invalid request field %q: %s", e.Field, e.Reason)
}

// Unwrap 支持标准错误链。
func (e *ValidationError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder。
func (e *ValidationError) Kind() ErrorKind { return ErrorKindValidation }

// TransportError 包装网络传输失败，通过 Unwrap 保留
// context.Canceled、context.DeadlineExceeded 与底层网络错误的原因链。
type TransportError struct {
	// Op 简要描述失败的操作，例如 "send" 或 "read"。
	Op string
	// Err 是底层错误。
	Err error
}

// Error 实现 error 接口。
func (e *TransportError) Error() string {
	return fmt.Sprintf("alipay: transport %s failed", e.Op)
}

// Unwrap 支持标准错误链。
func (e *TransportError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder；调用方主动取消归类为 ErrorKindCanceled。
func (e *TransportError) Kind() ErrorKind {
	if errors.Is(e.Err, context.Canceled) {
		return ErrorKindCanceled
	}
	return ErrorKindTransport
}

// HTTPError 表示网关返回非成功 HTTP 状态。
// BodySnippet 是有界的响应片段，仅用于诊断；不作为业务数据使用。
type HTTPError struct {
	// Status 是响应 HTTP 状态码。
	Status int
	// BodySnippet 是有界且去除控制字符的响应片段，可能为空。
	BodySnippet string
}

// Error 实现 error 接口；不输出响应片段，片段通过字段显式访问。
func (e *HTTPError) Error() string {
	return fmt.Sprintf("alipay: unexpected http status %d", e.Status)
}

// Kind 实现 ErrorKinder。
func (e *HTTPError) Kind() ErrorKind { return ErrorKindHTTP }

// ProtocolError 表示协议处理失败，包括响应格式异常、
// 预期节点缺失、响应超限与编解码失败。
type ProtocolError struct {
	// Op 简要描述失败的处理阶段，例如 "extract" 或 "decode"。
	Op string
	// Reason 说明失败原因。
	Reason string
	// Err 是底层错误，可以为 nil。
	Err error
}

// Error 实现 error 接口。
func (e *ProtocolError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("alipay: protocol %s: %s: %v", e.Op, e.Reason, e.Err)
	}
	return fmt.Sprintf("alipay: protocol %s: %s", e.Op, e.Reason)
}

// Unwrap 支持标准错误链。
func (e *ProtocolError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder。
func (e *ProtocolError) Kind() ErrorKind { return ErrorKindProtocol }

// SignatureError 表示请求签名失败，或响应/异步通知缺少签名、验签失败。
// 出现该错误时响应内容不能作为可信业务结果。
type SignatureError struct {
	// Stage 是失败阶段："request"、"response" 或 "notify"。
	Stage string
	// Reason 说明失败原因。
	Reason string
	// Err 是底层错误，可以为 nil。
	Err error
}

// Error 实现 error 接口。
func (e *SignatureError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("alipay: signature %s: %s: %v", e.Stage, e.Reason, e.Err)
	}
	return fmt.Sprintf("alipay: signature %s: %s", e.Stage, e.Reason)
}

// Unwrap 支持标准错误链。
func (e *SignatureError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder。
func (e *SignatureError) Kind() ErrorKind { return ErrorKindSignature }

// CertificateError 表示证书装载、匹配或选择失败。
type CertificateError struct {
	// Reason 说明失败原因，不包含证书内容。
	Reason string
	// Err 是底层错误，可以为 nil。
	Err error
}

// Error 实现 error 接口。
func (e *CertificateError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("alipay: certificate: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("alipay: certificate: %s", e.Reason)
}

// Unwrap 支持标准错误链。
func (e *CertificateError) Unwrap() error { return e.Err }

// Kind 实现 ErrorKinder。
func (e *CertificateError) Kind() ErrorKind { return ErrorKindCertificate }

// ErrorCode 是 V3 错误响应中的错误码原文，例如网关级
// "invalid-signature" 或业务级 "ACQ.TRADE_NOT_EXIST"。
type ErrorCode string

// String 实现 fmt.Stringer。
func (c ErrorCode) String() string { return string(c) }

// APIError 表示支付宝网关返回的业务错误。
//
// Verified 标识该响应是否通过了签名验证：已验签的业务错误为 true；
// 未带签名的 V3 错误响应属于未验签报文，其字段只能作为诊断信息，
// 不能作为可信业务结果。
type APIError struct {
	// Code 是 V3 错误响应中的错误码原文。
	Code ErrorCode
	// Msg 是平台返回的消息原文，可能包含业务自由文本。
	Msg string
	// Links 是 V3 错误响应附带的官方解决方案链接；可能为空。
	Links string
	// Verified 标识响应是否通过验签。
	Verified bool
}

// Error 实现 error 接口；仅输出码值，不输出平台自由文本，
// Msg 通过字段显式访问。
func (e *APIError) Error() string {
	return fmt.Sprintf("alipay: api error code=%s", e.Code)
}

// Kind 实现 ErrorKinder。
func (e *APIError) Kind() ErrorKind { return ErrorKindAPI }
