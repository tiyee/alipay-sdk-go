package auth_test

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// 测试目标：两种凭证满足 Credentials/CertificateProvider 契约；
// 材料校验（私钥匹配、有效期、根证书链、重复 SN）在构造期完成；
// 验签算法字段必须与 RSA2 一致；证书标识选择与通知语义符合设计。

func TestPublicKeyCredentialsContract(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	if credentials.Mode() != auth.ModePublicKey {
		t.Fatalf("Mode() = %v", credentials.Mode())
	}
	if err := credentials.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if credentials.Signer() == nil || credentials.Verifier() == nil {
		t.Fatal("Signer/Verifier must not be nil")
	}
	if credentials.Signer().Algorithm() != auth.SignatureAlgorithmRSA2 {
		t.Fatalf("Signer().Algorithm() = %q", credentials.Signer().Algorithm())
	}

	// 签名与验签闭环：应用签名器产出签名结构；
	// 验签使用支付宝公钥，签名由模拟网关私钥生成。
	signature, err := credentials.Signer().Sign(context.Background(), []byte("payload"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if signature.Algorithm != auth.SignatureAlgorithmRSA2 || signature.Value == "" {
		t.Fatalf("Signature = %+v", signature)
	}
	node := []byte(`{"code":"10000","msg":"Success"}`)
	gatewaySign, err := testutil.SignWithAlipayKey(node)
	if err != nil {
		t.Fatalf("sign with gateway key: %v", err)
	}
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: gatewaySign},
	})
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}

	// 算法字段不一致必须拒绝。
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: "RSA", Value: gatewaySign},
	})
	if err == nil {
		t.Fatal("Verify() should reject a mismatched algorithm")
	}
}

func TestPublicKeyCredentialsRejections(t *testing.T) {
	alipayPublicKey, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(testutil.AlipayPublicKeyPEM))
	if err != nil {
		t.Fatalf("parse alipay public key: %v", err)
	}
	if _, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       nil,
		AlipayPublicKey: alipayPublicKey,
	}); err == nil {
		t.Fatal("nil app signer should be rejected")
	}
	var typedNil *rsa.PrivateKey
	if _, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       typedNil,
		AlipayPublicKey: alipayPublicKey,
	}); err == nil {
		t.Fatal("typed nil app signer should be rejected")
	}
	if _, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner: testutil.NewAppCryptoSigner(t),
	}); err == nil {
		t.Fatal("nil alipay public key should be rejected")
	}
}

func TestCertificateCredentialsContract(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	credentials := suite.NewCertificateCredentials(t)

	if credentials.Mode() != auth.ModeCertificate {
		t.Fatalf("Mode() = %v", credentials.Mode())
	}
	if err := credentials.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	// 身份标识：SN 与按官方规则独立计算一致，副本可安全修改。
	identity := credentials.CertificateIdentity()
	if identity.AppCertificateSN != testutil.CertSN(t, suite.AppCert) {
		t.Fatalf("AppCertificateSN = %q, want %q", identity.AppCertificateSN, testutil.CertSN(t, suite.AppCert))
	}
	if identity.RootCertificateSN != testutil.RootCertSNOf(t, suite.Root) {
		t.Fatalf("RootCertificateSN = %q", identity.RootCertificateSN)
	}
	identity.AppCertificateSN = "mutated"
	if credentials.CertificateIdentity().AppCertificateSN == "mutated" {
		t.Fatal("CertificateIdentity() should return a copy")
	}

	// 接口断言：实现 CertificateProvider。
	var provider auth.CertificateProvider = credentials
	if provider.CertificateIdentity().RootCertificateSN == "" {
		t.Fatal("provider identity should be set")
	}

	// 同步响应语义：KeyID 命中配置的支付宝证书。
	node := []byte(`{"code":"10000","msg":"Success"}`)
	sign, err := testutil.SignWithAlipayKey(node)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
		KeyID:     testutil.CertSN(t, suite.AlipayCert),
	})
	if err != nil {
		t.Fatalf("Verify() with matching key id error = %v", err)
	}

	// 未知标识返回 CertSelectError。
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
		KeyID:     "unknown-sn",
	})
	var selErr *auth.CertSelectError
	if !errors.As(err, &selErr) {
		t.Fatalf("error = %v, want *auth.CertSelectError", err)
	}

	// 通知语义：契约不携带标识，单张证书时使用该证书验签。
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
	})
	if err != nil {
		t.Fatalf("Verify() without key id and a single certificate error = %v", err)
	}

	// 篡改报文验签失败。
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   []byte(`{"code":"10000","msg":"Tampered"}`),
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
		KeyID:     testutil.CertSN(t, suite.AlipayCert),
	})
	if err == nil {
		t.Fatal("Verify() should reject tampered payload")
	}
}

func TestCertificateCredentialsRotationSemantics(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	secondCert, secondKey := suite.NewSecondAlipayCert(t)
	credentials := suite.NewCertificateCredentials(t, suite.AlipayCertPEM, testutil.CertToPEM(t, secondCert))

	// 多证书按标识选择：由第二张证书配对私钥签名时命中第二张。
	node := []byte(`{"code":"10000","msg":"Success"}`)
	sign, err := testutil.SignWithKey(node, secondKey)
	if err != nil {
		t.Fatalf("sign with second key: %v", err)
	}
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
		KeyID:     testutil.CertSN(t, secondCert),
	})
	if err != nil {
		t.Fatalf("Verify() rotation error = %v", err)
	}

	// 多证书 + 无标识：无法唯一确定可信材料，明确失败。
	err = credentials.Verifier().Verify(context.Background(), auth.VerificationInput{
		Payload:   node,
		Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: sign},
	})
	var selErr *auth.CertSelectError
	if !errors.As(err, &selErr) || selErr.KeyID != "" {
		t.Fatalf("error = %v, want CertSelectError with empty key id", err)
	}
}

func TestCertificateCredentialsRejections(t *testing.T) {
	suite := testutil.NewCertSuite(t)

	build := func(mutate func(*auth.CertificateConfig)) error {
		config := auth.CertificateConfig{
			AppSigner:          testutil.NewAppCryptoSigner(t),
			AppCertificate:     suite.AppCert,
			AlipayCertificates: []*x509.Certificate{suite.AlipayCert},
			RootCertificates:   []*x509.Certificate{suite.Root},
		}
		mutate(&config)
		_, err := auth.NewCertificateCredentials(config)
		return err
	}

	if err := build(func(c *auth.CertificateConfig) { c.AppSigner = nil }); err == nil {
		t.Fatal("nil app signer should be rejected")
	}
	if err := build(func(c *auth.CertificateConfig) { c.AppCertificate = nil }); err == nil {
		t.Fatal("nil app certificate should be rejected")
	}
	if err := build(func(c *auth.CertificateConfig) { c.AlipayCertificates = nil }); err == nil {
		t.Fatal("empty alipay certificates should be rejected")
	}
	if err := build(func(c *auth.CertificateConfig) { c.RootCertificates = nil }); err == nil {
		t.Fatal("empty root certificates should be rejected")
	}
	// 私钥与应用证书不匹配。
	if err := build(func(c *auth.CertificateConfig) {
		c.AppSigner = testutil.NewAlipayCryptoSigner(t)
	}); err == nil {
		t.Fatal("mismatched app signer should be rejected")
	}
	// 支付宝证书已过期。
	if err := build(func(c *auth.CertificateConfig) {
		c.AlipayCertificates = []*x509.Certificate{suite.ExpiredAlipayCert(t)}
	}); err == nil {
		t.Fatal("expired alipay certificate should be rejected")
	}
	// 支付宝证书尚未生效。
	if err := build(func(c *auth.CertificateConfig) {
		c.AlipayCertificates = []*x509.Certificate{suite.NotYetValidAlipayCert(t)}
	}); err == nil {
		t.Fatal("not-yet-valid alipay certificate should be rejected")
	}
	// 签发者不在根证书集合内。
	if err := build(func(c *auth.CertificateConfig) {
		c.AlipayCertificates = []*x509.Certificate{suite.NewAlipayCertSignedBySecondRoot(t)}
	}); err == nil {
		t.Fatal("alipay certificate from an unknown root should be rejected")
	}
	// 根证书不是 CA。
	if err := build(func(c *auth.CertificateConfig) {
		c.RootCertificates = []*x509.Certificate{suite.AppCert}
	}); err == nil {
		t.Fatal("non-CA root should be rejected")
	}
	// 根证书集合中没有 RSA 签名算法证书。
	if err := build(func(c *auth.CertificateConfig) {
		c.RootCertificates = []*x509.Certificate{testutil.NewECDSARootCert(t)}
	}); err == nil {
		t.Fatal("root set without an RSA-signed certificate should be rejected")
	}
	// 重复 SN。
	if err := build(func(c *auth.CertificateConfig) {
		c.AlipayCertificates = []*x509.Certificate{suite.AlipayCert, suite.AlipayCert}
	}); err == nil {
		t.Fatal("duplicate alipay certificate SN should be rejected")
	}
}

func TestCertificateCredentialsIntermediateChain(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	leaf, ca := suite.NewIntermediateChain(t)
	config := auth.CertificateConfig{
		AppSigner: testutil.NewAppCryptoSigner(t), AppCertificate: suite.AppCert,
		RootCertificates: []*x509.Certificate{suite.Root}, AlipayCertificates: []*x509.Certificate{leaf},
	}
	if _, err := auth.NewCertificateCredentials(config); err == nil {
		t.Fatal("accepted missing intermediate")
	}
	config.AlipayCertificates = []*x509.Certificate{ca}
	if _, err := auth.NewCertificateCredentials(config); err == nil {
		t.Fatal("accepted CA as a platform signing certificate")
	}
	config.AlipayCertificates = []*x509.Certificate{ca, leaf}
	creds, err := auth.NewCertificateCredentials(config)
	if err != nil {
		t.Fatal(err)
	}
	payload := []byte("signed response")
	signature, err := testutil.SignWithAlipayKey(payload)
	if err != nil {
		t.Fatal(err)
	}
	input := auth.VerificationInput{Payload: payload, Signature: auth.Signature{Algorithm: auth.SignatureAlgorithmRSA2, Value: signature}, KeyID: testutil.CertSN(t, leaf)}
	if err := creds.Verifier().Verify(t.Context(), input); err != nil {
		t.Fatal(err)
	}
	input.KeyID = testutil.CertSN(t, ca)
	if err := creds.Verifier().Verify(t.Context(), input); err == nil {
		t.Fatal("intermediate was indexed for response verification")
	} else if _, ok := compat.AsType[*auth.CertSelectError](err); !ok {
		t.Fatalf("want unknown certificate, got %v", err)
	}
	config.RootCertificates = []*x509.Certificate{suite.SecondRoot}
	if _, err := auth.NewCertificateCredentials(config); err == nil {
		t.Fatal("accepted untrusted intermediate")
	}
}
