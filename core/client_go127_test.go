//go:build go1.27

package core_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

func TestClientExecuteMatchesPackageExecute(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := []byte(`{"trade_no":"T1","trade_status":"TRADE_SUCCESS"}`)
		for key, value := range testutil.V3ResponseHeaders(t, node, "176400000000c", "resp-nonce-d", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write(node)
	}))
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newQueryOperation(t)
	req := core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "T1"}}

	viaMethod, err1 := client.Execute(context.Background(), op, req)
	viaPackage, err2 := core.Execute(context.Background(), client, op, req)
	if err1 != nil || err2 != nil {
		t.Fatalf("errors: %v / %v", err1, err2)
	}
	if viaMethod.Data.TradeNo != viaPackage.Data.TradeNo {
		t.Fatal("two entry points diverge")
	}
}
