package auth

import (
	"context"
	"crypto/x509"
	"fmt"
	"time"
)

// CertValidityError 表示证书尚未生效或已经过期。
// Field 标明需要更新的配置项；更新证书后应重建凭证及使用它的实例。
type CertValidityError struct {
	// Field 是 CertificateConfig 配置项及其数组下标。
	Field string
	// NotBefore 是证书生效时间。
	NotBefore time.Time
	// NotAfter 是证书到期时间。
	NotAfter time.Time
}

// Error 实现 error 接口。
func (e *CertValidityError) Error() string {
	return fmt.Sprintf("alipay/auth: certificate %s is outside its validity period [%s, %s]; update certificates and rebuild credentials",
		e.Field, e.NotBefore.UTC().Format(time.RFC3339), e.NotAfter.UTC().Format(time.RFC3339))
}

// certificateValidity 只保存时间值的快照，调用方修改证书对象不会延长有效期。
// 每次使用均检查全部配置材料，由调用方移除或更新失效证书。
type certificateValidity []CertValidityError

func (v *certificateValidity) add(field string, cert *x509.Certificate) {
	*v = append(*v, CertValidityError{Field: field, NotBefore: cert.NotBefore, NotAfter: cert.NotAfter})
}

func (v certificateValidity) check(now time.Time) error {
	for _, cert := range v {
		if now.Before(cert.NotBefore) || now.After(cert.NotAfter) {
			return &cert
		}
	}
	return nil
}

type certificateSigner struct {
	Signer
	validity certificateValidity
}

func (s certificateSigner) Sign(ctx context.Context, payload []byte) (Signature, error) {
	if err := s.validity.check(time.Now()); err != nil {
		return Signature{}, err
	}
	return s.Signer.Sign(ctx, payload)
}
