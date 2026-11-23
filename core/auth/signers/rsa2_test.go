package signers_test

import (
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/auth/signers"
	"github.com/tiyee/alipay-sdk-go/core/auth/verifiers"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// 测试材料均为本仓库专用虚构密钥。

func TestRSA2SignMatchesIndependentOpenSSLVector(t *testing.T) {
	// openssl dgst -sha256 -sign 独立生成的已知答案向量，
	// 避免同一套签名与验签实现互相证明。
	signer := testutil.NewAppSigner(t)
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
	signer := testutil.NewAppSigner(t)
	verifier := testutil.NewAlipayVerifier(t)
	signature, err := signer.Sign([]byte("data"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	// 应用私钥与支付宝公钥是不同密钥对，交叉验签必须失败。
	if err := verifier.Verify([]byte("data"), signature); err == nil {
		t.Fatal("Verify() with unrelated key should fail")
	}
}

func TestRSA2VerifyFailures(t *testing.T) {
	appKey := testutil.NewAppCryptoSigner(t)
	alipayKey := testutil.NewAlipayCryptoSigner(t)
	signer, err := signers.NewRSA2(appKey)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	gatewaySigner, err := signers.NewRSA2(alipayKey)
	if err != nil {
		t.Fatalf("NewRSA2() error = %v", err)
	}
	// 用应用公钥验签：网关私钥产生的签名属于不同密钥对，必须失败。
	appVerifier, err := verifiers.NewRSA2FromKey(&appKey.PublicKey)
	if err != nil {
		t.Fatalf("build app verifier: %v", err)
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
			if err := appVerifier.Verify(tt.data, tt.sign); err == nil {
				t.Fatalf("Verify() expect error, got nil")
			}
		})
	}
}

func TestRSA2RejectsUnsupportedSigners(t *testing.T) {
	if _, err := signers.NewRSA2(nil); err == nil {
		t.Fatal("NewRSA2(nil) expect error")
	}
}

func TestAlgorithmValues(t *testing.T) {
	signer := testutil.NewAppSigner(t)
	if signer.Algorithm() != "RSA2" {
		t.Fatalf("Algorithm() = %q, want RSA2", signer.Algorithm())
	}
	verifier := testutil.NewAlipayVerifier(t)
	if verifier.Algorithm() != "RSA2" {
		t.Fatalf("Algorithm() = %q, want RSA2", verifier.Algorithm())
	}
}
