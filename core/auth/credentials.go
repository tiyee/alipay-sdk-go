// Package auth 定义支付宝客户端的认证凭证契约与两种内置实现。
//
// Credentials 是 NewClient 的第一个必选参数：公钥模式与证书模式
// 分别通过 NewPublicKeyCredentials 与 NewCertificateCredentials 构造，
// 外部包也可以实现 Credentials 接入自定义签名与验签。
// 具体的签名与验签算法实现位于 signers 与 verifiers 子包；
// 本包负责组合、模式派生与证书身份能力。
//
// 认证层不反向依赖 core.Client；认证失败由 Client 包装为配置、
// 签名、证书或验签错误。
package auth

import (
	"context"
	"crypto"
	"crypto/rsa"
	"fmt"
	"reflect"

	"github.com/tiyee/alipay-sdk-go/core/auth/signers"
	"github.com/tiyee/alipay-sdk-go/core/auth/verifiers"
)

// Mode 标识认证模式。模式由具体 Credentials 实现派生，
// 是只读信息，不是调用方可设置的策略开关。
type Mode uint8

const (
	// ModeUnset 表示未定义模式，构造 Client 时拒绝。
	ModeUnset Mode = iota
	// ModePublicKey 表示公钥模式。
	ModePublicKey
	// ModeCertificate 表示本地证书模式。
	ModeCertificate
)

// String 实现 fmt.Stringer。
func (m Mode) String() string {
	switch m {
	case ModePublicKey:
		return "public_key"
	case ModeCertificate:
		return "certificate"
	default:
		return "unset"
	}
}

// SignatureAlgorithm 标识签名算法。支付宝协议目前仅支持 RSA2。
type SignatureAlgorithm string

// SignatureAlgorithmRSA2 对应 SHA256WithRSA（RSA PKCS#1 v1.5）。
const SignatureAlgorithmRSA2 SignatureAlgorithm = "RSA2"

// String 实现 fmt.Stringer。
func (a SignatureAlgorithm) String() string { return string(a) }

// Signature 是签名结果：算法与按协议编码的签名值（RSA2 为 base64）。
type Signature struct {
	// Algorithm 是生成签名使用的算法。
	Algorithm SignatureAlgorithm
	// Value 是按协议编码的签名值。
	Value string
}

// VerificationInput 是一次验签所需的全部输入。
// Payload 是待验签原文（同步响应体或通知待验签串的原始字节）；
// KeyID 是协议模块从响应或通知中提取的可信材料标识
// （V3 REST 为响应头 alipay-sn），仅用于在本地已信任材料中查找，
// 不作为远端下载地址或任意公钥。是否必需由具体协议明确规定。
type VerificationInput struct {
	// Payload 是待验签原文。
	Payload []byte
	// Signature 是待校验的签名。
	Signature Signature
	// KeyID 是协议提供的可信材料标识；协议未提供时为空。
	KeyID string
}

// Signer 生成协议要求的请求签名。实现必须并发安全；
// 标准库 crypto.Signer 没有取消保证，ctx 仅用于传递调用上下文。
type Signer interface {
	// Algorithm 返回签名算法，例如 SignatureAlgorithmRSA2。
	Algorithm() SignatureAlgorithm
	// Sign 对待签名原文生成签名。
	Sign(ctx context.Context, payload []byte) (Signature, error)
}

// Verifier 使用本地可信材料校验签名。实现必须并发安全，
// 且必须校验 input.Signature.Algorithm 与自身算法一致，
// 不得因签名值自带算法字段就接受任意算法。
type Verifier interface {
	// Verify 校验验签输入；失败时返回错误，成功返回 nil。
	Verify(ctx context.Context, input VerificationInput) error
}

// Credentials 是 NewClient 的认证凭证契约：提供模式、签名器、
// 验签器与装载校验。实现必须并发安全，返回的对象在多次调用间保持稳定。
type Credentials interface {
	// Mode 返回凭证派生的认证模式。
	Mode() Mode
	// Signer 返回请求签名器。
	Signer() Signer
	// Verifier 返回响应/通知验签器。
	Verifier() Verifier
	// Validate 检查已装载状态与材料一致性；不隐式下载证书，
	// 不调用远程接口。
	Validate() error
}

// CertSelectError 表示验签所需的材料标识未匹配本地可信证书，
// 或协议未提供必需的标识。出现该错误时报文内容不能作为可信结果。
type CertSelectError struct {
	// KeyID 是未匹配的证书标识；为空表示协议未提供标识
	// 且本地可信材料无法唯一确定。
	KeyID string
}

// Error 实现 error 接口。
func (e *CertSelectError) Error() string {
	if e.KeyID == "" {
		return "alipay/auth: no certificate key id provided and the trusted material is not unique"
	}
	return fmt.Sprintf("alipay/auth: key id %q does not match any trusted local certificate", e.KeyID)
}

// rsa2Signer 将 signers.RSA2 适配为 Signer 契约。
type rsa2Signer struct {
	inner *signers.RSA2
}

// NewRSA2Signer 基于 crypto.Signer 构建 RSA2 签名器。
// 构造时检查公钥为 RSA 且不少于 2048 位。
// 导出供凭证实现与外部自定义 Credentials 复用同一适配。
func NewRSA2Signer(signer crypto.Signer) (Signer, error) {
	inner, err := signers.NewRSA2(signer)
	if err != nil {
		return nil, err
	}
	return rsa2Signer{inner: inner}, nil
}

// Algorithm 实现 Signer。
func (s rsa2Signer) Algorithm() SignatureAlgorithm { return SignatureAlgorithmRSA2 }

// Sign 实现 Signer。标准库签名不可取消，ctx 仅透传。
func (s rsa2Signer) Sign(ctx context.Context, payload []byte) (Signature, error) {
	value, err := s.inner.Sign(payload)
	if err != nil {
		return Signature{}, err
	}
	return Signature{Algorithm: SignatureAlgorithmRSA2, Value: value}, nil
}

// rsa2Verifier 将 verifiers.RSA2 适配为 Verifier 契约：
// 固定可信公钥，忽略 KeyID。
type rsa2Verifier struct {
	inner *verifiers.RSA2
}

// NewRSA2Verifier 基于已解析的 RSA 公钥构建固定公钥验签器。
func NewRSA2Verifier(publicKey *rsa.PublicKey) (Verifier, error) {
	inner, err := verifiers.NewRSA2FromKey(publicKey)
	if err != nil {
		return nil, err
	}
	return rsa2Verifier{inner: inner}, nil
}

// NewRSA2VerifierFromPEM 基于 PEM 编码的支付宝公钥构建验签器。
func NewRSA2VerifierFromPEM(alipayPublicKeyPEM string) (Verifier, error) {
	inner, err := verifiers.NewRSA2FromPEM(alipayPublicKeyPEM)
	if err != nil {
		return nil, err
	}
	return rsa2Verifier{inner: inner}, nil
}

// Verify 实现 Verifier：算法必须与 RSA2 一致，KeyID 不参与固定公钥验签。
func (v rsa2Verifier) Verify(ctx context.Context, input VerificationInput) error {
	if input.Signature.Algorithm != SignatureAlgorithmRSA2 {
		return fmt.Errorf("alipay/auth: unsupported signature algorithm %q; only %q is supported",
			input.Signature.Algorithm, SignatureAlgorithmRSA2)
	}
	return v.inner.Verify(input.Payload, input.Signature.Value)
}

// isNilInterface 报告接口值是否为 nil 或持有 nil 指针（typed nil）。
func isNilInterface(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
}

// nilCredentialsErr 返回统一的凭证错误。
func nilCredentialsErr(what string) error {
	return fmt.Errorf("alipay/auth: %s is nil", what)
}

// 编译期保证内置凭证满足契约。
var (
	_ Credentials         = (*PublicKeyCredentials)(nil)
	_ CertificateProvider = (*CertificateCredentials)(nil)
)
