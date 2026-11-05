package trans_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/fund/trans"
)

// TestFundEndToEndUniTransfer 通过真实 Client 与本地网关验证
// 转账完整链路（请求签名与响应验签）。
func TestFundEndToEndUniTransfer(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	node := `{"order_id":"202611051100700015","out_biz_no":"T1","status":"SUCCESS","pay_fund_order_id":"202611051100700016","trans_date":"2026-11-05 12:00:00"}`
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		if r.URL.Path != "/v3/alipay/fund/trans/uni/transfer" {
			t.Errorf("path = %q", r.URL.Path)
		}
		testutil.VerifyV3Authorization(t, r, body, r.Header.Get("alipay-app-auth-token"))
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000500", "fund-nonce-1", testutil.CertSN(t, suite.AlipayCert)) {
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
	svc, err := trans.NewService(client)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	resp, err := svc.UniTransfer(context.Background(), trans.UniTransferRequest{
		OutBizNo: "T1", TransAmount: "10.00", ProductCode: "TRANS_ACCOUNT_NO_PWD",
		PayeeInfo: &trans.Participant{
			Identity: ptr("2088101117955611"), IdentityType: ptr("ALIPAY_USER_ID"),
		}},
		core.WithAppAuthToken("app-auth-token-1"))
	if err != nil {
		t.Fatalf("UniTransfer() error = %v", err)
	}
	if !resp.Result.SignatureVerified || resp.Data.Status == nil || *resp.Data.Status != trans.TransStatusSuccess {
		t.Fatalf("unexpected result: %+v", resp)
	}
}

func TestUniTransferRejectsPublicKeyBeforeSending(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { t.Error("public key client sent a transfer") }))
	defer server.Close()
	client, err := core.NewClient(testutil.NewCredentials(t), option.WithAppID(testutil.TestAppID), option.WithGateway(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	service, err := trans.NewService(client)
	if err != nil {
		t.Fatal(err)
	}
	_, err = service.UniTransfer(t.Context(), trans.UniTransferRequest{OutBizNo: "offline", TransAmount: "1.00", ProductCode: "TRANS_ACCOUNT_NO_PWD", PayeeInfo: &trans.Participant{Identity: ptr("test-id"), IdentityType: ptr("ALIPAY_USER_ID")}})
	if e, ok := compat.AsType[*core.ValidationError](err); !ok || e.Field != "auth_requirement" {
		t.Fatalf("error = %v", err)
	}
}
