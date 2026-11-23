package auth

import (
	"crypto"
	"crypto/rsa"
	"fmt"
)

// PublicKeyConfig 是公钥模式的凭证配置。应用私钥以标准库
// crypto.Signer 传入（通常为 utils.ParsePrivateKeyPKCS8 解析的
// *rsa.PrivateKey）；支付宝公钥为解析后的 RSA 公钥。
// 配置不含证书字段；公钥模式不提供“失败后改用证书”的备用项。
type PublicKeyConfig struct {
	// AppSigner 是应用私钥对应的签名者。
	AppSigner crypto.Signer
	// AlipayPublicKey 是可信支付宝公钥。
	AlipayPublicKey *rsa.PublicKey
}

// PublicKeyCredentials 是公钥模式的 Credentials 实现：
// 应用 RSA2 签名器 + 固定支付宝公钥验签器。
// 实例构造后内部状态不变，可并发复用。
type PublicKeyCredentials struct {
	signer   Signer
	verifier Verifier
}

// NewPublicKeyCredentials 构造公钥凭证。双方 RSA 公钥的模数和指数
// 必须有效且不少于 2048 位；标准库 RSA 私钥还需通过参数一致性校验。
// 自定义 crypto.Signer 的内部配置由其实现负责，构造不调用 Sign。
func NewPublicKeyCredentials(config PublicKeyConfig) (*PublicKeyCredentials, error) {
	if isNilInterface(config.AppSigner) {
		return nil, nilCredentialsErr("app signer")
	}
	if config.AlipayPublicKey == nil {
		return nil, nilCredentialsErr("alipay public key")
	}
	signer, err := NewRSA2Signer(config.AppSigner)
	if err != nil {
		return nil, fmt.Errorf("alipay/auth: build app signer: %w", err)
	}
	verifier, err := NewRSA2Verifier(config.AlipayPublicKey)
	if err != nil {
		return nil, fmt.Errorf("alipay/auth: build alipay verifier: %w", err)
	}
	return &PublicKeyCredentials{signer: signer, verifier: verifier}, nil
}

// Mode 实现 Credentials，返回 ModePublicKey。
func (c *PublicKeyCredentials) Mode() Mode { return ModePublicKey }

// Signer 实现 Credentials。
func (c *PublicKeyCredentials) Signer() Signer { return c.signer }

// Verifier 实现 Credentials。
func (c *PublicKeyCredentials) Verifier() Verifier { return c.verifier }

// Validate 实现 Credentials：公钥凭证在构造期已完成材料校验，
// 此处确认组合对象仍然完整。
func (c *PublicKeyCredentials) Validate() error {
	if c.signer == nil || isNilInterface(c.signer) {
		return nilCredentialsErr("app signer")
	}
	if c.verifier == nil || isNilInterface(c.verifier) {
		return nilCredentialsErr("alipay verifier")
	}
	if c.signer.Algorithm() != SignatureAlgorithmRSA2 {
		return fmt.Errorf("alipay/auth: unsupported signer algorithm %q", c.signer.Algorithm())
	}
	return nil
}
