//go:build go1.25

package auth_test

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	"testing"
	"testing/synctest"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// 证书均为独立生成的虚构测试材料，以真实 DER 验证各角色的到期行为。
func TestCertificateCredentialsRuntimeValidity(t *testing.T) {
	for _, field := range []string{"app_certificate", "root_certificates[0]", "alipay_certificates[0]", "alipay_certificates[1]", "alipay_certificates[2]"} {
		t.Run(field, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				now := time.Now()
				appKey := testutil.NewAppCryptoSigner(t)
				platformKey := testutil.NewAlipayCryptoSigner(t)
				issue := func(name string, serial int64, public crypto.PublicKey, parent *x509.Certificate, ca bool) *x509.Certificate {
					expiry := now.Add(24 * time.Hour)
					if name == field {
						expiry = now.Add(time.Hour)
					}
					template := &x509.Certificate{
						SerialNumber: big.NewInt(serial), Subject: pkix.Name{CommonName: name},
						NotBefore: now.Add(-time.Hour), NotAfter: expiry,
						IsCA: ca, BasicConstraintsValid: true, KeyUsage: x509.KeyUsageDigitalSignature,
					}
					if ca {
						template.KeyUsage |= x509.KeyUsageCertSign
					}
					if parent == nil {
						parent = template
					}
					raw, err := x509.CreateCertificate(rand.Reader, template, parent, public, appKey)
					if err != nil {
						t.Fatal(err)
					}
					cert, err := x509.ParseCertificate(raw)
					if err != nil {
						t.Fatal(err)
					}
					return cert
				}
				root := issue("root_certificates[0]", 1, appKey.Public(), nil, true)
				intermediate := issue("alipay_certificates[1]", 2, appKey.Public(), root, true)
				leaf := issue("alipay_certificates[0]", 3, platformKey.Public(), intermediate, false)
				spare := issue("alipay_certificates[2]", 4, platformKey.Public(), intermediate, false)
				app := issue("app_certificate", 5, appKey.Public(), root, false)
				credentials, err := auth.NewCertificateCredentials(auth.CertificateConfig{
					AppSigner: appKey, AppCertificate: app, RootCertificates: []*x509.Certificate{root},
					AlipayCertificates: []*x509.Certificate{leaf, intermediate, spare},
				})
				if err != nil {
					t.Fatal(err)
				}
				payload := []byte("test response")
				signature, err := testutil.SignWithAlipayKey(payload)
				if err != nil {
					t.Fatal(err)
				}
				input := auth.VerificationInput{
					Payload: payload, KeyID: testutil.CertSN(t, leaf),
					Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: signature},
				}
				if err := credentials.Verifier().Verify(t.Context(), input); err != nil {
					t.Fatal(err)
				}
				if _, err := credentials.Signer().Sign(t.Context(), payload); err != nil {
					t.Fatal(err)
				}
				// 调用方原地改日期不能延长凭证内部快照的有效期。
				for _, cert := range []*x509.Certificate{app, root, intermediate, leaf, spare} {
					cert.NotAfter = now.Add(48 * time.Hour)
				}
				time.Sleep(time.Hour + time.Second)
				check := func(err error) {
					t.Helper()
					var validityErr *auth.CertValidityError
					if !errors.As(err, &validityErr) || validityErr.Field != field {
						t.Fatalf("error = %v, want validity error for %s", err, field)
					}
				}
				check(credentials.Validate())
				result, err := credentials.Signer().Sign(t.Context(), payload)
				check(err)
				if result.Value != "" {
					t.Fatal("expired credentials returned a signature")
				}
				check(credentials.Verifier().Verify(t.Context(), input))
			})
		})
	}
}
