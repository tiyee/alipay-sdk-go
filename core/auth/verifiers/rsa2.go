// Package verifiers 提供支付宝 RSA2 响应验签的底层实现。
// 实现不依赖 core/auth 契约类型，由 core/auth 中的凭证适配为
// auth.Verifier；本包也可以独立复用（例如测试网关侧核对请求签名）。
package verifiers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/tiyee/alipay-sdk-go/utils"
)

// AlgorithmRSA2 是支付宝 RSA2 签名对应的 sign_type 值。
const AlgorithmRSA2 = "RSA2"

// RSA2 使用固定可信 RSA 公钥校验 RSA2 签名（RSA PKCS#1 v1.5 + SHA-256）。
// 公钥模式使用支付宝公钥；证书模式由 core/auth 的可信证书索引按
// 标识选择公钥后复用本实现。实例初始化后可并发复用。
type RSA2 struct {
	publicKey *rsa.PublicKey
}

// NewRSA2FromKey 基于已解析的 RSA 公钥构建验签器。
func NewRSA2FromKey(publicKey *rsa.PublicKey) (*RSA2, error) {
	if publicKey == nil || publicKey.N == nil {
		return nil, fmt.Errorf("alipay/verifiers: nil rsa public key")
	}
	return &RSA2{publicKey: publicKey}, nil
}

// NewRSA2FromPEM 解析 X.509 SubjectPublicKeyInfo（"PUBLIC KEY"）PEM
// 编码的 RSA 公钥并构建验签器，即支付宝公钥的格式。
func NewRSA2FromPEM(alipayPublicKeyPEM string) (*RSA2, error) {
	publicKey, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(alipayPublicKeyPEM))
	if err != nil {
		return nil, err
	}
	return NewRSA2FromKey(publicKey)
}

// NewRSA2FromCertificate 从单张 X.509 证书 PEM 中提取公钥构建验签器，
// 用于证书模式下的异步通知验签：官方对通知不按证书标识选择证书，
// 直接使用本地配置的支付宝证书公钥（rsaCertCheckV1 行为）。
func NewRSA2FromCertificate(alipayCertPEM string) (*RSA2, error) {
	cert, err := utils.ParseCertificatePEM(alipayCertPEM)
	if err != nil {
		return nil, err
	}
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("alipay/verifiers: alipay certificate public key is %T, want RSA",
			cert.PublicKey)
	}
	return NewRSA2FromKey(publicKey)
}

// Algorithm 返回期望的签名算法 sign_type 值 "RSA2"。
func (v *RSA2) Algorithm() string { return AlgorithmRSA2 }

// Verify 校验 payload 与 base64 编码的签名值。签名格式或内容错误
// 均返回错误。
func (v *RSA2) Verify(payload []byte, sign string) error {
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return fmt.Errorf("alipay/verifiers: decode signature: %w", err)
	}
	digest := sha256.Sum256(payload)
	if err := rsa.VerifyPKCS1v15(v.publicKey, crypto.SHA256, digest[:], signature); err != nil {
		return fmt.Errorf("alipay/verifiers: rsa2 verify: %w", err)
	}
	return nil
}
