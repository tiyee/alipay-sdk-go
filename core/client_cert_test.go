package core_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// V3 证书模式测试：响应按 alipay-sn 头选择本地可信支付宝证书验签；
// 未知 SN、缺失 SN、错钥与轮换集合按失败处理，不退化为无验签调用。

func newCertClient(t *testing.T, gatewayURL string, suite *testutil.CertSuite, alipayCertPEMs ...string) *core.Client {
	t.Helper()
	credentials := suite.NewCertificateCredentials(t, alipayCertPEMs...)
	client, err := core.NewClient(credentials,
		option.WithAppID(testAppID), option.WithGateway(gatewayURL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func TestClientCertModeEndToEnd(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	alipaySN := testutil.CertSN(t, suite.AlipayCert)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		testutil.VerifyV3Authorization(t, r, body, "")
		if sn := r.Header.Get("alipay-root-cert-sn"); sn == "" {
			t.Error("certificate mode must send alipay-root-cert-sn header")
		}
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "app_cert_sn=") {
			t.Errorf("certificate mode Authorization must carry app_cert_sn: %q", authHeader)
		}
		node := `{"trade_no":"cert-1","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000006", "resp-nonce-7", alipaySN) {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newCertClient(t, gateway.URL, suite, suite.AlipayCertPEM)
	resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050010"},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !resp.Result.SignatureVerified || resp.Data.TradeNo != "cert-1" {
		t.Fatalf("unexpected result: %+v", resp)
	}
}

func TestClientCertModeUnknownSNFails(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000007", "resp-nonce-8", "unknown-sn") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newCertClient(t, gateway.URL, suite, suite.AlipayCertPEM)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050011"},
	})
	var certErr *core.CertificateError
	if !errors.As(err, &certErr) {
		t.Fatalf("expected *core.CertificateError, got %v", err)
	}
}

func TestClientCertModeMissingSNHeader(t *testing.T) {
	// 契约（core/auth/resolver.go）：证书模式响应缺 alipay-sn 头时，
	// 本地仅一张可信证书则按该证书验签（与官方对通知的验签方式一致）；
	// 轮换集合（多张证书）无法唯一确定材料，明确失败，不退化为无验签成功。
	// 是否与官方 V3 REST 对 alipay-sn 的强制性要求一致，待沙箱核验。
	t.Run("单张可信证书时回退验签成功", func(t *testing.T) {
		suite := testutil.NewCertSuite(t)
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
			for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000012", "resp-nonce-i", "") {
				w.Header()[key] = value
			}
			_, _ = w.Write([]byte(node))
		}))
		defer gateway.Close()

		client := newCertClient(t, gateway.URL, suite, suite.AlipayCertPEM)
		resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "202611050012"},
		})
		if err != nil || !resp.Result.SignatureVerified {
			t.Fatalf("single trusted cert must verify: %v, %+v", err, resp)
		}
	})
	t.Run("轮换集合缺 SN 时明确失败", func(t *testing.T) {
		suite := testutil.NewCertSuite(t)
		newCert, _ := suite.NewSecondAlipayCert(t)
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
			for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000016", "resp-nonce-m", "") {
				w.Header()[key] = value
			}
			_, _ = w.Write([]byte(node))
		}))
		defer gateway.Close()

		client := newCertClient(t, gateway.URL, suite,
			suite.AlipayCertPEM, testutil.CertToPEM(t, newCert))
		_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "202611050016"},
		})
		var certErr *core.CertificateError
		if !errors.As(err, &certErr) {
			t.Fatalf("expected *core.CertificateError, got %v", err)
		}
	})
}

func TestClientCertModeTamperedResponseFails(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	alipaySN := testutil.CertSN(t, suite.AlipayCert)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signed := []byte(`{"trade_no":"honest","trade_status":"TRADE_SUCCESS"}`)
		tampered := `{"trade_no":"tampered","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, signed, "1764000000013", "resp-nonce-j", alipaySN) {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(tampered))
	}))
	defer gateway.Close()

	client := newCertClient(t, gateway.URL, suite, suite.AlipayCertPEM)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050013"},
	})
	var sigErr *core.SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("expected *core.SignatureError, got %v", err)
	}
}

func TestClientCertModeSelectsCertAmongRotationSet(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	// 轮换场景：本地同时信任新旧两张支付宝证书，响应由新证书配对
	// 私钥签名并携带新证书 SN（alipay-sn 头）。
	newCert, newKey := suite.NewSecondAlipayCert(t)
	newSN := testutil.CertSN(t, newCert)

	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := []byte(`{"trade_no":"rotate-1","trade_status":"TRADE_SUCCESS"}`)
		for key, value := range testutil.V3ResponseHeadersWithKey(t, node, "1764000000014", "resp-nonce-k", newKey, newSN) {
			w.Header()[key] = value
		}
		_, _ = w.Write(node)
	}))
	defer gateway.Close()

	client := newCertClient(t, gateway.URL, suite,
		suite.AlipayCertPEM, testutil.CertToPEM(t, newCert))
	resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050014"},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !resp.Result.SignatureVerified || resp.Data.TradeNo != "rotate-1" {
		t.Fatalf("unexpected result: %+v", resp)
	}
}

func TestClientCertModeSignedErrorClassified(t *testing.T) {
	// 非 2xx 且带 alipay-sn 的错误响应：先验签再分类，Verified=true。
	suite := testutil.NewCertSuite(t)
	alipaySN := testutil.CertSN(t, suite.AlipayCert)
	body := `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在"}`
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range testutil.V3ResponseHeaders(t, []byte(body), "1764000000015", "resp-nonce-l", alipaySN) {
			w.Header()[key] = value
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(body))
	}))
	defer gateway.Close()

	client := newCertClient(t, gateway.URL, suite, suite.AlipayCertPEM)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050015"},
	})
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) || !apiErr.Verified || apiErr.Code != "ACQ.TRADE_NOT_EXIST" {
		t.Fatalf("error = %v, want verified ACQ.TRADE_NOT_EXIST", err)
	}
}
