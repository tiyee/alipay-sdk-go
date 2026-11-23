package trade_test

import (
	"context"
	"crypto/rsa"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/notify"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// 测试目标：证书模式复用同一 trade.Service 与 core/notify 流程，
// 不重新实现任何 Service 或验签逻辑。
// 请求侧：Authorization 携带 app_cert_sn 并签名，另发 alipay-root-cert-sn 头；
// 响应侧：按 alipay-sn 选择可信证书验签；
// 通知侧：官方对通知不按 SN 选证书，直接使用支付宝证书公钥验签。

func TestTradeServiceWorksWithCertificateClient(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	alipaySN := testutil.CertSN(t, suite.AlipayCert)

	var gotRootCertSN string
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		testutil.VerifyV3Authorization(t, r, body, "")
		if !strings.Contains(r.Header.Get("Authorization"), "app_cert_sn=") {
			t.Errorf("certificate mode Authorization must carry app_cert_sn: %q", r.Header.Get("Authorization"))
		}
		gotRootCertSN = r.Header.Get("alipay-root-cert-sn")
		if gotRootCertSN == "" {
			t.Error("alipay-root-cert-sn header is missing")
		}
		node := `{"trade_no":"2026110522001400001234567890","out_trade_no":"ORDER-1","trade_status":"TRADE_SUCCESS","total_amount":"20.00"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000200", "cert-nonce-1", alipaySN) {
			w.Header()[key] = value
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client, err := core.NewClient(
		suite.NewCertificateCredentials(t),
		option.WithAppID(testutil.TestAppID),
		option.WithGateway(gateway.URL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	service, err := trade.NewService(client)
	if err != nil {
		t.Fatalf("trade.NewService() error = %v", err)
	}

	resp, err := service.Query(context.Background(), trade.QueryRequest{
		OutTradeNo: "ORDER-1",
	})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if resp.Data == nil || resp.Data.TradeStatus != trade.TradeStatusSuccess {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
	if !resp.Result.SignatureVerified {
		t.Fatal("SignatureVerified = false")
	}
	if gotRootCertSN == "" {
		t.Error("root cert sn was not sent")
	}

	// V3 契约下 notify_url 属于请求模型业务字段，调用选项必须被拒绝。
	_, err = service.Query(context.Background(), trade.QueryRequest{OutTradeNo: "ORDER-2"},
		core.WithNotifyURL("https://merchant.example.com/notify"))
	var valErr *core.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("notify_url option must be rejected on v3, got %v", err)
	}
}

func TestNotifyParseWithCertificateVerifier(t *testing.T) {
	suite := testutil.NewCertSuite(t)

	// 证书模式下通知验签使用支付宝证书公钥（官方 rsaCertCheckV1 行为），
	// Handler 复用凭证的模式对应验签器。
	notifyCredentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       testutil.NewAppCryptoSigner(t),
		AlipayPublicKey: alipayPublicKeyFromCert(t, suite.AlipayCertPEM),
	})
	if err != nil {
		t.Fatalf("build notify credentials: %v", err)
	}
	handler, err := notify.NewHandler(notifyCredentials, notify.PolicyTradeStatusSync)
	if err != nil {
		t.Fatalf("notify.NewHandler() error = %v", err)
	}

	form := baseNotifyForm()
	signed := testutil.SignNotify(t, form)

	n, err := notify.Parse(context.Background(), handler, signed, trade.NotificationDecoder())
	if err != nil {
		t.Fatalf("notify parse error = %v", err)
	}
	if n.Data == nil || n.Data.TradeStatus != trade.TradeStatusSuccess {
		t.Fatalf("unexpected notification: %+v", n.Data)
	}

	// 错误证书公钥（应用证书公钥）验签失败。
	wrongCredentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       testutil.NewAppCryptoSigner(t),
		AlipayPublicKey: alipayPublicKeyFromCert(t, suite.AppCertPEM),
	})
	if err != nil {
		t.Fatalf("build wrong credentials: %v", err)
	}
	wrongHandler, err := notify.NewHandler(wrongCredentials, notify.PolicyTradeStatusSync)
	if err != nil {
		t.Fatalf("build wrong handler: %v", err)
	}
	if _, err := notify.Parse(context.Background(), wrongHandler, signed, trade.NotificationDecoder()); err == nil {
		t.Fatal("notify should fail with a mismatched certificate")
	}
}

// alipayPublicKeyFromCert 从证书 PEM 提取 RSA 公钥。
func alipayPublicKeyFromCert(t *testing.T, certPEM string) *rsa.PublicKey {
	t.Helper()
	cert, err := utils.ParseCertificatePEM(certPEM)
	if err != nil {
		t.Fatalf("parse certificate: %v", err)
	}
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		t.Fatalf("certificate public key is %T, want RSA", cert.PublicKey)
	}
	return publicKey
}
