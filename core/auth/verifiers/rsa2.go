// Package verifiers 提供支付宝响应签名校验的实现。
package verifiers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/tiyee/go-alipay/utils"
)

// AlgorithmRSA2 是支付宝 RSA2 签名对应的 sign_type 值。
const AlgorithmRSA2 = "RSA2"

// RSA2 使用可信支付宝公钥校验 RSA2 签名（RSA PKCS#1 v1.5 + SHA-256）。
// 公钥模式直接配置支付宝公钥；证书模式按响应 SN 选择可信证书，
// 在后续阶段扩展。实例初始化后可并发复用。
type RSA2 struct {
	publicKey *rsa.PublicKey
}

// NewRSA2 解析 X.509 SubjectPublicKeyInfo PEM 编码的支付宝公钥
// 并构建验签器。
func NewRSA2(alipayPublicKeyPEM string) (*RSA2, error) {
	publicKey, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(alipayPublicKeyPEM))
	if err != nil {
		return nil, err
	}
	return &RSA2{publicKey: publicKey}, nil
}

// Algorithm 实现 auth.Verifier，返回 "RSA2"。
func (v *RSA2) Algorithm() string { return AlgorithmRSA2 }

// Verify 实现 auth.Verifier：对 data 计算 SHA-256 摘要，
// 校验 base64 解码后的签名。签名格式或内容错误均返回错误。
func (v *RSA2) Verify(data []byte, sign string) error {
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return fmt.Errorf("alipay/verifiers: decode signature: %w", err)
	}
	digest := sha256.Sum256(data)
	if err := rsa.VerifyPKCS1v15(v.publicKey, crypto.SHA256, digest[:], signature); err != nil {
		return fmt.Errorf("alipay/verifiers: rsa2 verify: %w", err)
	}
	return nil
}
