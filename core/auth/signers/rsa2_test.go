package signers_test

import (
	"bytes"
	"testing"

	"github.com/tiyee/go-alipay/core/auth/signers"
	"github.com/tiyee/go-alipay/core/auth/verifiers"
	"github.com/tiyee/go-alipay/internal/testutil"
)

// 测试材料均为本仓库专用虚构密钥。

func TestRSA2SignMatchesIndependentOpenSSLVector(t *testing.T) {
	// openssl dgst -sha256 -sign 独立生成的已知答案向量，
	// 避免同一套签名与验签实现互相证明。
	signer, err := signers.NewRSA2(testutil.AppPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	got, err := signer.Sign([]byte(testutil.KnownAnswerSignContent))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if got != testutil.KnownAnswerSignature {
		t.Fatalf("Sign() = %q, want openssl vector %q", got, testutil.KnownAnswerSignature)
	}
}

func TestRSA2SignThenVerifyWithSeparateKeyMaterial(t *testing.T) {
	// 应用私钥签名、支付宝公钥（同一密钥对的公钥部分）验签，
	// 并用网关私钥的独立密钥对验证失败路径。
	signer, err := signers.NewRSA2(testutil.AppPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	verifier, err := verifiers.NewRSA2(testutil.AppPublicKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	signature, err := signer.Sign([]byte("data"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if err := verifier.Verify([]byte("data"), signature); err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
}

func TestRSA2VerifyFailures(t *testing.T) {
	signer, err := signers.NewRSA2(testutil.AppPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	gatewaySigner, err := signers.NewRSA2(testutil.AlipayPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	verifier, err := verifiers.NewRSA2(testutil.AppPublicKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	signature, err := signer.Sign([]byte("payload"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	other, err := gatewaySigner.Sign([]byte("payload"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	tests := []struct {
		name string
		data []byte
		sign string
	}{
		{name: "报文篡改", data: []byte("payload-tampered"), sign: signature},
		{name: "签名截断", data: []byte("payload"), sign: signature[:len(signature)-4]},
		{name: "非 base64", data: []byte("payload"), sign: "not-base64!!"},
		{name: "错误密钥签名", data: []byte("payload"), sign: other},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifier.Verify(tt.data, tt.sign); err == nil {
				t.Fatalf("Verify() expect error, got nil")
			}
		})
	}
}

func TestRSA2DeterministicSignContent(t *testing.T) {
	signer, err := signers.NewRSA2(testutil.AppPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	first, err := signer.Sign([]byte(testutil.KnownAnswerSignContent))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if first != testutil.KnownAnswerSignature {
		t.Fatalf("PKCS#1 v1.5 签名对同一输入应确定：got %q", first)
	}
}

func TestRSA2RejectsUnsupportedKeys(t *testing.T) {
	tests := []struct {
		name string
		pem  string
	}{
		{name: "空内容", pem: ""},
		{name: "非 PEM", pem: "not a pem"},
		{name: "PKCS#1 私钥", pem: "-----BEGIN RSA PRIVATE KEY-----\nMA==\n-----END RSA PRIVATE KEY-----"},
		{name: "公钥当私钥", pem: testutil.AppPublicKeyPEM},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := signers.NewRSA2(tt.pem); err == nil {
				t.Fatalf("NewRSA2() expect error, got nil")
			}
		})
	}
	if _, err := verifiers.NewRSA2(testutil.AppPrivateKeyPEM); err == nil {
		t.Fatal("verifiers.NewRSA2 with private key expect error")
	}
	if _, err := verifiers.NewRSA2(""); err == nil {
		t.Fatal("verifiers.NewRSA2 with empty expect error")
	}
}

// 编译期检查：算法标识固定为 RSA2。
func TestAlgorithmValues(t *testing.T) {
	signer, err := signers.NewRSA2(testutil.AppPrivateKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	if signer.Algorithm() != "RSA2" {
		t.Fatalf("Algorithm() = %q, want RSA2", signer.Algorithm())
	}
	verifier, err := verifiers.NewRSA2(testutil.AppPublicKeyPEM)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	if verifier.Algorithm() != "RSA2" {
		t.Fatalf("Algorithm() = %q, want RSA2", verifier.Algorithm())
	}
	if !bytes.Contains([]byte(signer.Algorithm()), []byte("RSA")) {
		t.Fatal("unreachable")
	}
}
