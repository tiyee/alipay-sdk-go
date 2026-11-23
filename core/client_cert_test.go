package core_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
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
	for _, count := range []int{1, 2} {
		for _, status := range []int{http.StatusOK, http.StatusBadRequest} {
			t.Run(fmt.Sprintf("certificates_%d_status_%d", count, status), func(t *testing.T) {
				suite := testutil.NewCertSuite(t)
				certs := []string{suite.AlipayCertPEM}
				if count == 2 {
					cert, _ := suite.NewSecondAlipayCert(t)
					certs = append(certs, testutil.CertToPEM(t, cert))
				}
				gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					node := `{"code":"invalid-parameter","message":"test","trade_no":"x"}`
					for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000012", "response-nonce", "") {
						w.Header()[key] = value
					}
					w.WriteHeader(status)
					_, _ = w.Write([]byte(node))
				}))
				defer gateway.Close()
				client := newCertClient(t, gateway.URL, suite, certs...)
				resp, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "order"}})
				var certErr *core.CertificateError
				var selectErr *auth.CertSelectError
				if !errors.As(err, &certErr) || !errors.As(err, &selectErr) {
					t.Fatalf("error = %v, want certificate error with selection cause", err)
				}
				if resp == nil || resp.Result == nil || resp.Data != nil || resp.Result.SignatureVerified {
					t.Fatalf("unexpected response: %+v", resp)
				}
			})
		}
	}
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

func TestClientRejectsMissingResponseSignatureMetadata(t *testing.T) {
	for _, missing := range []string{"alipay-timestamp", "alipay-nonce"} {
		for _, status := range []int{http.StatusOK, http.StatusBadRequest} {
			t.Run(fmt.Sprintf("%s_%d", missing, status), func(t *testing.T) {
				gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body := []byte(`{"code":"invalid-parameter","message":"test","trade_no":"order"}`)
					timestamp, nonce := "1764000000012", "response-nonce"
					if missing == "alipay-timestamp" {
						timestamp = ""
					} else {
						nonce = ""
					}
					// 即便按空字段计算的签名有效，缺少协议字段也必须报错。
					headers := testutil.V3ResponseHeaders(t, body, timestamp, nonce, "")
					headers.Del(missing)
					for name, values := range headers {
						w.Header()[name] = values
					}
					w.WriteHeader(status)
					_, _ = w.Write(body)
				}))
				defer gateway.Close()
				client := newTestClient(t, gateway.URL)
				response, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "order"}})
				var protocolErr *core.ProtocolError
				if !errors.As(err, &protocolErr) || !strings.Contains(protocolErr.Reason, missing) {
					t.Fatalf("error = %v, want missing %s protocol error", err, missing)
				}
				if response == nil || response.Data != nil || response.Result == nil || response.Result.SignatureVerified {
					t.Fatalf("unexpected response: %+v", response)
				}
			})
		}
	}
}
