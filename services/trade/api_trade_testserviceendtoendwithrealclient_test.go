package trade_test

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
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestServiceEndToEndWithRealClient(t *testing.T) {
	// 通过 core.Client 与本地网关验证完整链路：Authorization 签名请求、
	// 响应头验签、成功解码与业务错误分类（V3 业务错误为非 2xx + 错误体）。
	successNode := `{"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248","trade_status":"TRADE_SUCCESS","total_amount":"88.88"}`
	errorBody := `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在"}`
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		if r.URL.Path != "/v3/alipay/trade/query" {
			t.Errorf("path = %q", r.URL.Path)
		}
		testutil.VerifyV3Authorization(t, r, body, "")
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(string(body), `"T3"`) {
			// 业务错误分支：非 2xx 且响应带签名。
			for key, value := range testutil.V3ResponseHeaders(t, []byte(errorBody), "1764000000101", "trade-nonce-2", "") {
				w.Header()[key] = value
			}
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errorBody))
			return
		}
		for key, value := range testutil.V3ResponseHeaders(t, []byte(successNode), "1764000000100", "trade-nonce-1", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(successNode))
	}))
	defer gateway.Close()

	client, err := core.NewClient(
		testutil.NewCredentials(t),
		option.WithAppID(testutil.TestAppID),
		option.WithGateway(gateway.URL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	svc, err := trade.NewService(client)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}

	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if !resp.Result.SignatureVerified || resp.Data.TradeStatus != trade.TradeStatusSuccess {
		t.Fatalf("unexpected result: %+v", resp)
	}

	_, bizErr := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T3"})
	if bizErr == nil {
		t.Fatal("expect APIError for business failure")
	}
	var apiErr *core.APIError
	if !errors.As(bizErr, &apiErr) || apiErr.Code != "ACQ.TRADE_NOT_EXIST" || !apiErr.Verified {
		t.Fatalf("error = %v", bizErr)
	}
}
