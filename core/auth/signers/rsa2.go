// Package signers 提供支付宝请求签名的实现。
package signers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/tiyee/go-alipay/utils"
)

// AlgorithmRSA2 是支付宝 RSA2 签名对应的 sign_type 值。
const AlgorithmRSA2 = "RSA2"

// RSA2 使用标准库 RSA PKCS#1 v1.5 与 SHA-256 实现 RSA2 签名，
// 不使用 RSA-PSS 替代协议算法。实例初始化后可并发复用。
type RSA2 struct {
	privateKey *rsa.PrivateKey
}

// NewRSA2 解析 PKCS#8 PEM 编码的 RSA 私钥并构建签名器。
// 调用方不得传入生产凭证；测试与示例必须使用专用虚构密钥。
func NewRSA2(privateKeyPEM string) (*RSA2, error) {
	privateKey, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(privateKeyPEM))
	if err != nil {
		return nil, err
	}
	return &RSA2{privateKey: privateKey}, nil
}

// Algorithm 实现 auth.Signer，返回 "RSA2"。
func (s *RSA2) Algorithm() string { return AlgorithmRSA2 }

// Sign 实现 auth.Signer：对 data 计算 SHA-256 摘要并生成
// PKCS#1 v1.5 签名，返回 base64 编码结果。
func (s *RSA2) Sign(data []byte) (string, error) {
	digest := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, digest[:])
	if err != nil {
		return "", fmt.Errorf("alipay/signers: rsa2 sign: %w", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
