package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestCloseResponseMapping(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(
		`{"code":"10000","msg":"Success","trade_no":"T1","out_trade_no":"O1"}`)
	resp, err := svc.Close(context.Background(), trade.CloseRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if resp.Data.TradeNo == nil || *resp.Data.TradeNo != "T1" {
		t.Fatalf("TradeNo = %+v", resp.Data.TradeNo)
	}
}
