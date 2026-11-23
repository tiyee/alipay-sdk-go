// 本文件提供动态生成的虚构测试证书材料：根 CA、支付宝公钥证书、
// 应用公钥证书及过期/未生效/错误签发等变体。
// 叶子证书的公钥与本文件顶部既有的专用测试私钥配对，
// 便于在证书模式下复用既有签名/验签测试链路。
// 所有材料仅用于测试，不对应任何真实商户或应用，严禁用于生产环境。
package testutil

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/auth/signers"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// certCounter 保证同一进程内多套证书的主体与序列号互不相同，
// 避免 SN（MD5(issuer+serial)）在跨套件场景下意外重复。
var certCounter atomic.Int64

// CertSuite 是一套动态生成的证书材料。
type CertSuite struct {
	// Root 是自签发 RSA 根 CA 证书。
	Root *x509.Certificate
	// RootPEM 是 Root 的 PEM 编码。
	RootPEM string
	// SecondRoot 是第二张自签发根 CA，与 Root 无信任关系。
	SecondRoot *x509.Certificate
	// SecondRootPEM 是 SecondRoot 的 PEM 编码。
	SecondRootPEM string
	// AlipayCert 模拟支付宝公钥证书，由 Root 签发，
	// 公钥与 AlipayPrivateKeyPEM 配对。
	AlipayCert *x509.Certificate
	// AlipayCertPEM 是 AlipayCert 的 PEM 编码。
	AlipayCertPEM string
	// AppCert 模拟应用公钥证书，由 Root 签发，
	// 公钥与 AppPrivateKeyPEM 配对。
	AppCert *x509.Certificate
	// AppCertPEM 是 AppCert 的 PEM 编码。
	AppCertPEM string

	rootKey       *rsa.PrivateKey
	secondRootKey *rsa.PrivateKey
}

// NewCertSuite 生成一套虚构测试证书。叶子证书有效期覆盖当前时刻。
func NewCertSuite(tb testing.TB) *CertSuite {
	tb.Helper()
	suite := &CertSuite{}
	suite.rootKey = newTestRSAKey(tb)
	suite.Root = issueCert(tb, certRequest{
		subject:   "Alipay Go SDK Dynamic Test Root CA",
		publicKey: &suite.rootKey.PublicKey,
		ca:        true,
		signerKey: suite.rootKey, // 自签发
	})
	suite.secondRootKey = newTestRSAKey(tb)
	suite.SecondRoot = issueCert(tb, certRequest{
		subject:   "Alipay Go SDK Dynamic Test Root CA R2",
		publicKey: &suite.secondRootKey.PublicKey,
		ca:        true,
		signerKey: suite.secondRootKey,
	})
	suite.AlipayCert = suite.newLeafCert(tb, suite.Root, suite.rootKey, "Alipay Go SDK Dynamic Test Public Cert", AlipayPublicKeyPEM)
	suite.AppCert = suite.newLeafCert(tb, suite.Root, suite.rootKey, "Alipay Go SDK Dynamic Test App Cert", AppPublicKeyPEM)
	suite.RootPEM = CertToPEM(tb, suite.Root)
	suite.SecondRootPEM = CertToPEM(tb, suite.SecondRoot)
	suite.AlipayCertPEM = CertToPEM(tb, suite.AlipayCert)
	suite.AppCertPEM = CertToPEM(tb, suite.AppCert)
	return suite
}

// NewCredentials 返回公钥模式的测试凭证（专用虚构密钥）。
// tb 允许为 nil（Example 函数等无测试上下文场景）。
func NewCredentials(tb testing.TB) *auth.PublicKeyCredentials {
	alipayPublicKey, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(AlipayPublicKeyPEM))
	failOrFail(tb, "parse alipay public key", err)
	credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       NewAppCryptoSigner(tb),
		AlipayPublicKey: alipayPublicKey,
	})
	failOrFail(tb, "build public key credentials", err)
	return credentials
}

// NewCertificateCredentials 返回证书模式的测试凭证：
// 应用证书取自套件，支付宝证书由调用方提供 PEM（可多张，轮换场景），
// 根证书取自套件。
func (s *CertSuite) NewCertificateCredentials(tb testing.TB, alipayCertPEMs ...string) *auth.CertificateCredentials {
	if len(alipayCertPEMs) == 0 {
		alipayCertPEMs = []string{s.AlipayCertPEM}
	}
	alipayCerts := make([]*x509.Certificate, 0, len(alipayCertPEMs))
	for _, pemText := range alipayCertPEMs {
		cert, err := utils.ParseCertificatePEM(pemText)
		failOrFail(tb, "parse alipay certificate", err)
		alipayCerts = append(alipayCerts, cert)
	}
	credentials, err := auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner:          NewAppCryptoSigner(tb),
		AppCertificate:     s.AppCert,
		AlipayCertificates: alipayCerts,
		RootCertificates:   []*x509.Certificate{s.Root},
	})
	failOrFail(tb, "build certificate credentials", err)
	return credentials
}

// CertSN 返回证书的支付宝公钥证书 SN（按官方规则计算）。
// 测试辅助：测试内对已知证书直接计算期望值。
func CertSN(tb testing.TB, cert *x509.Certificate) string {
	if cert == nil {
		failOrFail(tb, "cert is nil", errors.New("nil certificate"))
	}
	sn, err := protocol.CertSN(cert)
	failOrFail(tb, "compute cert sn", err)
	return sn
}

// SignWithKey 使用给定 RSA 私钥对数据生成 base64 签名（测试辅助）。
func SignWithKey(data []byte, key *rsa.PrivateKey) (string, error) {
	signer, err := signers.NewRSA2(key)
	if err != nil {
		return "", err
	}
	return signer.Sign(data)
}

// RootCertSNOf 返回按官方规则对给定根证书集合计算的 alipay_root_cert_sn。
func RootCertSNOf(tb testing.TB, certs ...*x509.Certificate) string {
	sn, err := protocol.RootCertSN(certs)
	failOrFail(tb, "compute root cert sn", err)
	return sn
}

// TestAppID 是测试使用的虚构应用 APPID。
const TestAppID = "2021000000000001"

// CertToPEM 将证书编码为 PEM 文本。
func CertToPEM(tb testing.TB, cert *x509.Certificate) string {
	tb.Helper()
	if cert == nil {
		tb.Fatal("cert is nil")
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}))
}

// NewAlipayCertSignedBySecondRoot 签发一张公钥与 AlipayPrivateKeyPEM 配对、
// 但由 SecondRoot 签发的支付宝证书，用于“签发者不在可信根集合”场景。
func (s *CertSuite) NewAlipayCertSignedBySecondRoot(tb testing.TB) *x509.Certificate {
	tb.Helper()
	return s.newLeafCert(tb, s.SecondRoot, s.secondRootKey, "Alipay Go SDK Dynamic Test Public Cert", AlipayPublicKeyPEM)
}

// NewIntermediateChain issues a test platform leaf through an intermediate CA under Root.
func (s *CertSuite) NewIntermediateChain(tb testing.TB) (leaf, intermediate *x509.Certificate) {
	tb.Helper()
	key := newTestRSAKey(tb)
	intermediate = issueCert(tb, certRequest{
		subject: "Alipay Go SDK Test Intermediate", publicKey: &key.PublicKey, ca: true,
		signer: s.Root, signerKey: s.rootKey,
	})
	leaf = s.newLeafCert(tb, intermediate, key, "Alipay Go SDK Test Platform Leaf", AlipayPublicKeyPEM)
	return leaf, intermediate
}

// NewSecondAlipayCert 签发另一张由 Root 签发的支付宝证书，
// 公钥独立生成，用于多证书选择场景。
func (s *CertSuite) NewSecondAlipayCert(tb testing.TB) (*x509.Certificate, *rsa.PrivateKey) {
	tb.Helper()
	key := newTestRSAKey(tb)
	cert := issueCert(tb, certRequest{
		subject:   "Alipay Go SDK Dynamic Test Public Cert B",
		publicKey: &key.PublicKey,
		signer:    s.Root,
		signerKey: s.rootKey,
	})
	return cert, key
}

// NewRootWithSameSubject 生成与 Root 主体相同但密钥不同的自签发根 CA，
// 用于“签发者 DN 匹配但签名不匹配”场景。
func (s *CertSuite) NewRootWithSameSubject(tb testing.TB) *x509.Certificate {
	tb.Helper()
	key := newTestRSAKey(tb)
	return issueCert(tb, certRequest{
		subject:   s.Root.Subject.CommonName,
		publicKey: &key.PublicKey,
		ca:        true,
		signerKey: key,
	})
}

// NewECDSARootCert 生成自签发的 ECDSA（P-256）根 CA 证书，
// 用于“根证书集合中没有 RSA 签名算法证书”场景。
func NewECDSARootCert(tb testing.TB) *x509.Certificate {
	tb.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		tb.Fatalf("generate ecdsa key: %v", err)
	}
	return issueCert(tb, certRequest{
		subject:   "Alipay Go SDK Dynamic Test ECDSA Root CA",
		publicKey: &key.PublicKey,
		ca:        true,
		signerKey: key,
	})
}

// ExpiredAlipayCert 返回一张已过期的支付宝证书（由 Root 签发）。
func (s *CertSuite) ExpiredAlipayCert(tb testing.TB) *x509.Certificate {
	tb.Helper()
	return s.alipayVariant(tb, validity{notBefore: time.Now().Add(-48 * time.Hour), notAfter: time.Now().Add(-24 * time.Hour)})
}

// NotYetValidAlipayCert 返回一张尚未生效的支付宝证书（由 Root 签发）。
func (s *CertSuite) NotYetValidAlipayCert(tb testing.TB) *x509.Certificate {
	tb.Helper()
	return s.alipayVariant(tb, validity{notBefore: time.Now().Add(24 * time.Hour), notAfter: time.Now().Add(48 * time.Hour)})
}

func (s *CertSuite) alipayVariant(tb testing.TB, v validity) *x509.Certificate {
	tb.Helper()
	pub, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(AlipayPublicKeyPEM))
	if err != nil {
		tb.Fatalf("parse alipay public key: %v", err)
	}
	return issueCert(tb, certRequest{
		subject:   "Alipay Go SDK Dynamic Test Public Cert Variant",
		publicKey: pub,
		validity:  v,
		signer:    s.Root,
		signerKey: s.rootKey,
	})
}

// newLeafCert 由指定 CA 签发一张公钥来自 PEM 的叶子证书，有效期覆盖当前时刻。
func (s *CertSuite) newLeafCert(tb testing.TB, ca *x509.Certificate, caKey *rsa.PrivateKey, commonName, publicKeyPEM string) *x509.Certificate {
	tb.Helper()
	pub, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(publicKeyPEM))
	if err != nil {
		tb.Fatalf("parse public key: %v", err)
	}
	return issueCert(tb, certRequest{
		subject:   commonName,
		publicKey: pub,
		signer:    ca,
		signerKey: caKey,
	})
}

type certRequest struct {
	subject   string
	publicKey any
	ca        bool
	validity  validity
	signer    *x509.Certificate // nil 表示自签发
	signerKey crypto.Signer     // RSA 或 ECDSA 私钥
}

type validity struct {
	notBefore time.Time
	notAfter  time.Time
}

func issueCert(tb testing.TB, req certRequest) *x509.Certificate {
	tb.Helper()
	if req.validity.notBefore.IsZero() {
		req.validity.notBefore = time.Now().Add(-time.Hour)
	}
	if req.validity.notAfter.IsZero() {
		req.validity.notAfter = time.Now().Add(24 * time.Hour)
	}
	template := &x509.Certificate{
		SerialNumber:          newTestSerial(),
		Subject:               pkix.Name{Organization: []string{"Alipay Go SDK Test"}, CommonName: req.subject, Country: []string{"CN"}},
		NotBefore:             req.validity.notBefore,
		NotAfter:              req.validity.notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  req.ca,
	}
	if req.ca {
		template.KeyUsage |= x509.KeyUsageCertSign
	}
	issuer := template
	if req.signer != nil {
		issuer = req.signer
	}
	der, err := x509.CreateCertificate(rand.Reader, template, issuer, req.publicKey, req.signerKey)
	if err != nil {
		tb.Fatalf("create certificate: %v", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		tb.Fatalf("parse certificate: %v", err)
	}
	return cert
}

func newTestRSAKey(tb testing.TB) *rsa.PrivateKey {
	tb.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		tb.Fatalf("generate rsa key: %v", err)
	}
	return key
}

func newTestSerial() *big.Int {
	return big.NewInt(certCounter.Add(1))
}
