package utils

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/tiyee/alipay-sdk-go/internal/protocol"
)

// timeLayout 是错误信息中使用的时间格式。
const certTimeLayout = "2006-01-02 15:04:05"

// ParseCertificatePEM 解析恰好包含一张 X.509 证书的 PEM 文本
// （例如应用公钥证书）。
// 包含多张证书或无证书时返回错误。
func ParseCertificatePEM(pemData string) (*x509.Certificate, error) {
	certs, err := ParseCertificatesPEM(pemData)
	if err != nil {
		return nil, err
	}
	if len(certs) != 1 {
		return nil, fmt.Errorf("alipay/utils: expected exactly one certificate, got %d", len(certs))
	}
	return certs[0], nil
}

// ParseCertificatesPEM 解析 PEM 文本中的全部 X.509 证书，
// 顺序与文件中的证书顺序一致。支付宝混合根证书包使用 ParseRootCertificatesPEM。
// 不包含任何证书时返回错误。
func ParseCertificatesPEM(pemData string) ([]*x509.Certificate, error) {
	return parseCertificatesPEM(pemData, false)
}

// ParseRootCertificatesPEM 按文件顺序读取支付宝 RSA2 根证书包中的 RSA 签名证书。
// 根证书包可同时含国密、ECDSA 与 RSA。普通应用证书和平台证书应使用严格的 ParseCertificatePEM 或 ParseCertificatesPEM。
// SN 筛选规则见 https://opendocs.alipay.com/common/02kdne.md。
func ParseRootCertificatesPEM(pemData string) ([]*x509.Certificate, error) {
	return parseCertificatesPEM(pemData, true)
}

func parseCertificatesPEM(pemData string, rsaRootsOnly bool) ([]*x509.Certificate, error) {
	var certs []*x509.Certificate
	rest := []byte(pemData)
	for {
		block, remainder := pem.Decode(rest)
		if block == nil {
			break
		}
		rest = remainder
		if block.Type != "CERTIFICATE" {
			return nil, fmt.Errorf("alipay/utils: unsupported PEM type %q; want %q",
				block.Type, "CERTIFICATE")
		}
		if rsaRootsOnly {
			rsaSigned, err := protocol.IsRSASignedCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("alipay/utils: inspect root certificate: %w", err)
			}
			if !rsaSigned {
				continue
			}
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("alipay/utils: parse certificate: %w", err)
		}
		certs = append(certs, cert)
	}
	if len(certs) == 0 {
		if rsaRootsOnly {
			return nil, fmt.Errorf("alipay/utils: no RSA-signed roots in certificate bundle")
		}
		return nil, fmt.Errorf("alipay/utils: no certificate block found in certificate data")
	}
	// 证书文件应只包含证书块；存在残留非空白内容说明文件被截断或夹杂，
	// 拒绝装载而不是静默忽略。
	if trailing := strings.TrimSpace(string(rest)); trailing != "" {
		return nil, fmt.Errorf("alipay/utils: unexpected trailing data after certificates")
	}
	return certs, nil
}

// CheckCertValidity 校验证书在 now 时刻处于有效期内。
// 期限由证书 NotBefore/NotAfter 决定，本函数不做其他语义判断。
func CheckCertValidity(cert *x509.Certificate, now time.Time) error {
	if cert == nil {
		return fmt.Errorf("alipay/utils: certificate is nil")
	}
	if now.Before(cert.NotBefore) {
		return fmt.Errorf("alipay/utils: certificate not valid before %s",
			cert.NotBefore.UTC().Format(certTimeLayout))
	}
	if now.After(cert.NotAfter) {
		return fmt.Errorf("alipay/utils: certificate expired at %s",
			cert.NotAfter.UTC().Format(certTimeLayout))
	}
	return nil
}

// VerifyIssuedByRoots 校验证书是否由可信根证书集合中的某张证书直接签发：
// 按签发者 DN 匹配根证书，再用该根证书的公钥校验证书签名。
// 签发者不在集合内或签名不匹配均返回错误。
// 与官方证书可信校验（AntCertificationUtil.verifyCert）对单张证书的
// 规则一致。
func VerifyIssuedByRoots(cert *x509.Certificate, roots []*x509.Certificate) error {
	if cert == nil {
		return fmt.Errorf("alipay/utils: certificate is nil")
	}
	if len(roots) == 0 {
		return fmt.Errorf("alipay/utils: no trusted root certificates configured")
	}
	var signatureErr error
	for _, root := range roots {
		// 与官方实现一致：按 DER 编码的 DN 字节比较签发者与根证书主体。
		if root == nil || !bytes.Equal(root.RawSubject, cert.RawIssuer) {
			continue
		}
		if err := cert.CheckSignatureFrom(root); err != nil {
			signatureErr = err
			continue
		}
		return nil
	}
	if signatureErr != nil {
		return fmt.Errorf("alipay/utils: certificate signature does not verify against issuer: %w", signatureErr)
	}
	return fmt.Errorf("alipay/utils: certificate issuer is not in the trusted root set")
}
