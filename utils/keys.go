// Package utils 提供密钥与证书解析等可独立复用的工具。
// 本包不依赖 core 或 services；错误信息不包含密钥内容。
package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

// ParsePrivateKeyPKCS8 解析 PKCS#8（"PRIVATE KEY"）PEM 编码的 RSA 私钥。
// 支付宝密钥工具默认输出该格式；其他格式按需求显式增加并测试。
func ParsePrivateKeyPKCS8(pemData string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("alipay/utils: no PEM block found in private key data")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("alipay/utils: unsupported private key PEM type %q; RSA2 requires PKCS#8 %q",
			block.Type, "PRIVATE KEY")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("alipay/utils: parse PKCS#8 private key: %w", err)
	}
	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("alipay/utils: private key is %T, want RSA", key)
	}
	return rsaKey, nil
}

// ParsePublicKeyPKIX 解析 X.509 SubjectPublicKeyInfo（"PUBLIC KEY"）
// PEM 编码的 RSA 公钥，即支付宝公钥的格式。
func ParsePublicKeyPKIX(pemData string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("alipay/utils: no PEM block found in public key data")
	}
	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("alipay/utils: unsupported public key PEM type %q; want PKIX %q",
			block.Type, "PUBLIC KEY")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("alipay/utils: parse PKIX public key: %w", err)
	}
	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("alipay/utils: public key is %T, want RSA", key)
	}
	return rsaKey, nil
}

// TrimKeyText 去除密钥文本的首尾空白，便于接受从配置文件
// 或环境变量读取的带换行内容。不改变密钥内容本身。
func TrimKeyText(text string) string {
	return strings.TrimSpace(text)
}
