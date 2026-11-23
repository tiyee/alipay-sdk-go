package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestQueryResponseUnknownEnumPreserved(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_status":"FUTURE_NEW_STATUS","total_amount":"1.00"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if resp.Data.TradeStatus != trade.TradeStatus("FUTURE_NEW_STATUS") {
		t.Fatalf("unknown status = %q, want preserved", resp.Data.TradeStatus)
	}
	if resp.Data.TradeStatus.IsKnown() {
		t.Error("unknown status must not be reported as known")
	}
}
