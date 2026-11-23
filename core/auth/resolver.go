package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/tiyee/alipay-sdk-go/core/auth/verifiers"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
)

// trustedCertificateSet 是本地可信支付宝证书的只读索引：
// 按官方规则计算每张证书的 SN，建立 SN 到验签器的映射。
// 构造后不再变化，可并发复用。
type trustedCertificateSet struct {
	bySN map[string]*verifiers.RSA2
	sns  []string
}

// newTrustedCertificateSet 基于可信支付宝公钥证书集合构建索引。
// 空集合、nil 证书、非 RSA 公钥或重复 SN 均返回错误。
func newTrustedCertificateSet(alipayCerts []*x509.Certificate) (*trustedCertificateSet, error) {
	if len(alipayCerts) == 0 {
		return nil, fmt.Errorf("alipay/auth: at least one alipay certificate is required")
	}
	set := &trustedCertificateSet{bySN: make(map[string]*verifiers.RSA2, len(alipayCerts))}
	for _, cert := range alipayCerts {
		if cert == nil {
			return nil, fmt.Errorf("alipay/auth: nil alipay certificate")
		}
		publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("alipay/auth: alipay certificate %q public key is %T, want RSA",
				cert.Subject, cert.PublicKey)
		}
		sn, err := protocol.CertSN(cert)
		if err != nil {
			return nil, fmt.Errorf("alipay/auth: compute alipay certificate SN: %w", err)
		}
		if _, exists := set.bySN[sn]; exists {
			return nil, fmt.Errorf("alipay/auth: duplicate alipay certificate SN %q", sn)
		}
		verifier, err := verifiers.NewRSA2FromKey(publicKey)
		if err != nil {
			return nil, err
		}
		set.bySN[sn] = verifier
		set.sns = append(set.sns, sn)
	}
	return set, nil
}

// SNs 返回可信证书 SN 的副本，用于诊断与测试。
func (s *trustedCertificateSet) SNs() []string {
	out := make([]string, len(s.sns))
	copy(out, s.sns)
	return out
}

// verify 按 keyID 选择可信证书后校验签名：
//   - keyID 非空：按 SN 查找；未匹配返回 *CertSelectError。
//   - keyID 为空：协议未提供标识。仅当本地恰好配置一张证书时
//     使用该证书（与官方对通知的验签方式一致）；配置多张时无法
//     唯一确定可信材料，明确失败，不盲试未确认密钥。
//
// V3 REST 在 core 层强制要求 alipay-sn；空标识只用于普通支付通知契约。
func (s *trustedCertificateSet) verify(payload []byte, signature, keyID string) error {
	if keyID == "" {
		if len(s.sns) != 1 {
			return &CertSelectError{KeyID: ""}
		}
		keyID = s.sns[0]
	}
	verifier, ok := s.bySN[keyID]
	if !ok {
		return &CertSelectError{KeyID: keyID}
	}
	return verifier.Verify(payload, signature)
}
