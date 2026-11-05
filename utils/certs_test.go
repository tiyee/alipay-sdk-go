package utils_test

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/internal/protocol"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/utils"
)

func TestParseCertificatePEM(t *testing.T) {
	suite := testutil.NewCertSuite(t)

	cert, err := utils.ParseCertificatePEM(suite.AlipayCertPEM)
	if err != nil {
		t.Fatalf("ParseCertificatePEM() error = %v", err)
	}
	if !cert.Equal(suite.AlipayCert) {
		t.Fatal("parsed certificate mismatch")
	}

	// 多张证书时拒绝（单证书入口）。
	if _, err := utils.ParseCertificatePEM(suite.RootPEM + suite.SecondRootPEM); err == nil {
		t.Fatal("ParseCertificatePEM() should reject multi-cert input")
	}

	// 无证书内容时拒绝。
	if _, err := utils.ParseCertificatePEM("not a pem"); err == nil {
		t.Fatal("ParseCertificatePEM() should reject non-PEM input")
	}
}

func TestParseCertificatesPEMKeepsFileOrder(t *testing.T) {
	suite := testutil.NewCertSuite(t)

	certs, err := utils.ParseCertificatesPEM(suite.RootPEM + suite.SecondRootPEM)
	if err != nil {
		t.Fatalf("ParseCertificatesPEM() error = %v", err)
	}
	if len(certs) != 2 {
		t.Fatalf("len(certs) = %d, want 2", len(certs))
	}
	if !certs[0].Equal(suite.Root) || !certs[1].Equal(suite.SecondRoot) {
		t.Fatal("certificate order should follow the file order")
	}

	if _, err := utils.ParseCertificatesPEM(suite.RootPEM + "garbage\n"); err == nil {
		t.Fatal("ParseCertificatesPEM() should reject invalid certificate bytes")
	}
}

func TestVerifyIssuedByRoots(t *testing.T) {
	suite := testutil.NewCertSuite(t)

	// 签发者在可信根集合内，签名匹配。
	if err := utils.VerifyIssuedByRoots(suite.AlipayCert, []*x509.Certificate{suite.Root}); err != nil {
		t.Fatalf("VerifyIssuedByRoots() error = %v", err)
	}

	// 多根集合命中其一。
	if err := utils.VerifyIssuedByRoots(suite.AlipayCert, []*x509.Certificate{suite.SecondRoot, suite.Root}); err != nil {
		t.Fatalf("VerifyIssuedByRoots() multi-root error = %v", err)
	}

	// 签发者不在根集合内（由另一张根签发）。
	other := suite.NewAlipayCertSignedBySecondRoot(t)
	if err := utils.VerifyIssuedByRoots(other, []*x509.Certificate{suite.Root}); err == nil {
		t.Fatal("VerifyIssuedByRoots() should reject an unknown issuer")
	}

	// 签发者 DN 匹配但签名不匹配（同主体、不同密钥的根）。
	forgedRoot := suite.NewRootWithSameSubject(t)
	if err := utils.VerifyIssuedByRoots(suite.AlipayCert, []*x509.Certificate{forgedRoot}); err == nil {
		t.Fatal("VerifyIssuedByRoots() should reject a signature mismatch")
	}

	// 根集合为空时拒绝。
	if err := utils.VerifyIssuedByRoots(suite.AlipayCert, nil); err == nil {
		t.Fatal("VerifyIssuedByRoots() should reject an empty root set")
	}
}

func TestCheckCertValidity(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	now := time.Now()

	if err := utils.CheckCertValidity(suite.AlipayCert, now); err != nil {
		t.Fatalf("CheckCertValidity() error = %v", err)
	}
	if err := utils.CheckCertValidity(suite.ExpiredAlipayCert(t), now); err == nil {
		t.Fatal("CheckCertValidity() should reject an expired certificate")
	}
	if err := utils.CheckCertValidity(suite.NotYetValidAlipayCert(t), now); err == nil {
		t.Fatal("CheckCertValidity() should reject a not-yet-valid certificate")
	}
}

// Public root bundle from the official SDK, used only for offline interoperability tests.
// https://github.com/alipay/alipay-sdk-nodejs-all/blob/e1676390871eb7a69b4f81aa15de8327ac15738a/test/fixtures/alipayRootCert.crt
func TestMixedRootBundle(t *testing.T) {
	data, err := os.ReadFile("testdata/alipay-root.crt")
	if err != nil {
		t.Fatal(err)
	}
	roots, err := utils.ParseRootCertificatesPEM(string(data))
	if err != nil || len(roots) != 2 {
		t.Fatalf("RSA roots = %d, error = %v", len(roots), err)
	}
	// RFC2253 issuer and decimal serial independently extracted with OpenSSL.
	sn, err := protocol.RootCertSN(roots)
	if err != nil || sn != "687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6" {
		t.Fatalf("root SN = %s, error = %v", sn, err)
	}
	first, _ := pem.Decode(data)
	if _, err := utils.ParseRootCertificatesPEM(string(pem.EncodeToMemory(first))); err == nil {
		t.Fatal("accepted bundle without RSA roots")
	}
	if _, err := utils.ParseCertificatesPEM(string(data)); err == nil {
		t.Fatal("strict leaf parser silently skipped a certificate")
	}

}
