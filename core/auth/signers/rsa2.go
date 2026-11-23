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
	"math"
	"reflect"
)

// AlgorithmRSA2 是支付宝 RSA2 签名对应的 sign_type 值。
const AlgorithmRSA2 = "RSA2"

// RSA2 基于标准库 crypto.Signer（通常为解析后的 RSA 私钥）实现
// RSA2 签名：SHA-256 摘要 + RSA PKCS#1 v1.5，签名值 base64 编码。
// 不使用 RSA-PSS 替代协议算法。实例初始化后可并发复用。
type RSA2 struct {
	signer crypto.Signer
}

// NewRSA2 基于 crypto.Signer 构建签名器。签名者必须提供有效的
// RSA 公钥且不少于 2048 位；传入 *rsa.PrivateKey 时同时校验私钥参数。
// 标准库 Signer 本身没有 context 取消保证，构造不执行签名或访问远端。
func NewRSA2(signer crypto.Signer) (*RSA2, error) {
	if signer == nil {
		return nil, fmt.Errorf("alipay/signers: nil crypto signer")
	}
	value := reflect.ValueOf(signer)
	switch value.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		if value.IsNil() {
			return nil, fmt.Errorf("alipay/signers: nil crypto signer")
		}
	}
	public := signer.Public()
	publicKey, ok := public.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("alipay/signers: crypto signer public key is %T, want RSA", public)
	}
	if publicKey == nil || publicKey.N == nil || publicKey.N.Sign() <= 0 || publicKey.N.Bit(0) == 0 || publicKey.N.BitLen() < 2048 {
		return nil, fmt.Errorf("alipay/signers: rsa modulus must be positive, odd and at least 2048 bits")
	}
	if publicKey.E < 3 || publicKey.E%2 == 0 || publicKey.E > math.MaxInt32 {
		return nil, fmt.Errorf("alipay/signers: invalid rsa public exponent")
	}
	if privateKey, ok := signer.(*rsa.PrivateKey); ok {
		// 忽略已有的预计算缓存，校验实际传入的参数，且不修改调用方对象。
		key := rsa.PrivateKey{
			PublicKey: rsa.PublicKey{N: privateKey.N, E: privateKey.E},
			D:         privateKey.D,
			Primes:    privateKey.Primes,
		}
		if key.D == nil || key.D.Sign() <= 0 {
			return nil, fmt.Errorf("alipay/signers: missing or invalid rsa private exponent")
		}
		for _, prime := range key.Primes {
			if prime == nil || prime.Sign() <= 0 {
				return nil, fmt.Errorf("alipay/signers: missing or invalid rsa prime")
			}
		}
		if err := key.Validate(); err != nil {
			return nil, fmt.Errorf("alipay/signers: invalid rsa private key: %w", err)
		}
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
