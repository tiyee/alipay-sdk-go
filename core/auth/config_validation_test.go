package auth_test

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"io"
	"math"
	"math/big"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// publicOnlySigner 模拟远端签名者：构造只可检查公钥，不可发起签名调用。
type publicOnlySigner struct {
	public crypto.PublicKey
	calls  int
}

func (s *publicOnlySigner) Public() crypto.PublicKey { return s.public }
func (s *publicOnlySigner) Sign(io.Reader, []byte, crypto.SignerOpts) ([]byte, error) {
	s.calls++
	return nil, errors.New("unexpected signing operation")
}

func TestCredentialConstructionRejectsInvalidRSAKeys(t *testing.T) {
	appKey := testutil.NewAppCryptoSigner(t)
	platformKey := testutil.NewAlipayCryptoSigner(t)
	largeExponent := math.MaxInt32
	largeExponent++
	for _, test := range []struct {
		name string
		key  *rsa.PublicKey
	}{
		{"typed nil", nil},
		{"missing modulus", &rsa.PublicKey{E: 65537}},
		{"zero modulus", &rsa.PublicKey{N: new(big.Int), E: 65537}},
		{"negative modulus", &rsa.PublicKey{N: new(big.Int).Neg(appKey.N), E: 65537}},
		{"even modulus", &rsa.PublicKey{N: new(big.Int).Add(appKey.N, big.NewInt(1)), E: 65537}},
		{"short modulus", &rsa.PublicKey{N: big.NewInt(17), E: 65537}},
		{"zero exponent", &rsa.PublicKey{N: appKey.N}},
		{"even exponent", &rsa.PublicKey{N: appKey.N, E: 2}},
		{"large exponent", &rsa.PublicKey{N: appKey.N, E: largeExponent}},
	} {
		t.Run(test.name, func(t *testing.T) {
			remote := &publicOnlySigner{public: test.key}
			credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{AppSigner: remote, AlipayPublicKey: &platformKey.PublicKey})
			if err == nil || credentials != nil || remote.calls != 0 {
				t.Fatalf("invalid signer was accepted or called: credentials=%v err=%v calls=%d", credentials, err, remote.calls)
			}
			credentials, err = auth.NewPublicKeyCredentials(auth.PublicKeyConfig{AppSigner: appKey, AlipayPublicKey: test.key})
			if err == nil || credentials != nil {
				t.Fatalf("invalid verifier was accepted: credentials=%v err=%v", credentials, err)
			}
		})
	}
	remote := &publicOnlySigner{public: &appKey.PublicKey}
	if _, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{AppSigner: remote, AlipayPublicKey: &platformKey.PublicKey}); err != nil || remote.calls != 0 {
		t.Fatalf("valid remote signer: error=%v calls=%d", err, remote.calls)
	}
	var typedNil *publicOnlySigner
	if signer, err := auth.NewRSA2Signer(typedNil); err == nil || signer != nil {
		t.Fatalf("typed nil signer accepted: %v, %v", signer, err)
	}
}

func TestCredentialConstructionRejectsInvalidPrivateParameters(t *testing.T) {
	for _, mutate := range []func(*rsa.PrivateKey){
		func(key *rsa.PrivateKey) { key.D = nil },
		func(key *rsa.PrivateKey) { key.D = new(big.Int).Add(key.D, big.NewInt(1)) },
		func(key *rsa.PrivateKey) { key.Primes = nil },
		func(key *rsa.PrivateKey) { key.Primes[0] = nil },
	} {
		key := testutil.NewAppCryptoSigner(t)
		// 解析后的密钥可能已有预计算缓存，校验不得因此忽略被修改的参数。
		key.Precompute()
		mutate(key)
		if signer, err := auth.NewRSA2Signer(key); err == nil || signer != nil {
			t.Fatalf("invalid private key accepted: %v, %v", signer, err)
		}
	}
}

func TestCertificateConstructionRejectsMissingAppPublicKey(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	suite.AppCert.PublicKey = (*rsa.PublicKey)(nil)
	credentials, err := auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner: testutil.NewAppCryptoSigner(t), AppCertificate: suite.AppCert,
		AlipayCertificates: []*x509.Certificate{suite.AlipayCert}, RootCertificates: []*x509.Certificate{suite.Root},
	})
	if err == nil || credentials != nil {
		t.Fatalf("invalid app certificate accepted: %v, %v", credentials, err)
	}
}
