// Package signers 提供支付宝 RSA2 请求签名的底层实现。
// 实现不依赖 core/auth 契约类型，由 core/auth 中的凭证适配为
// auth.Signer；本包也可以独立复用（例如测试网关侧签名）。
package signers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// AlgorithmRSA2 是支付宝 RSA2 签名对应的 sign_type 值。
const AlgorithmRSA2 = "RSA2"

// RSA2 基于标准库 crypto.Signer（通常为解析后的 RSA 私钥）实现
// RSA2 签名：SHA-256 摘要 + RSA PKCS#1 v1.5，签名值 base64 编码。
// 不使用 RSA-PSS 替代协议算法。实例初始化后可并发复用。
type RSA2 struct {
	signer crypto.Signer
}

// NewRSA2 基于 crypto.Signer 构建签名器。签名者公钥必须是 RSA；
// 其他算法返回错误。标准库 Signer 本身没有 context 取消保证。
func NewRSA2(signer crypto.Signer) (*RSA2, error) {
	if signer == nil {
		return nil, fmt.Errorf("alipay/signers: nil crypto signer")
	}
	publicKey, ok := signer.Public().(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("alipay/signers: crypto signer public key is %T, want RSA", signer.Public())
	}
	if publicKey.N == nil || publicKey.N.BitLen() < 2048 {
		return nil, fmt.Errorf("alipay/signers: rsa key must be at least 2048 bits")
	}
	return &RSA2{signer: signer}, nil
}

// Algorithm 返回 sign_type 值 "RSA2"。
func (s *RSA2) Algorithm() string { return AlgorithmRSA2 }

// Sign 对待签名串计算 SHA-256 摘要并通过底层 crypto.Signer 生成
// PKCS#1 v1.5 签名，返回 base64 编码结果。
func (s *RSA2) Sign(payload []byte) (string, error) {
	digest := sha256.Sum256(payload)
	signature, err := s.signer.Sign(rand.Reader, digest[:], crypto.SHA256)
	if err != nil {
		return "", fmt.Errorf("alipay/signers: rsa2 sign: %w", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
