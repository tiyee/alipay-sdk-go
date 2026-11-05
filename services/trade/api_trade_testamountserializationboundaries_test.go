package trade_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestAmountSerializationBoundaries(t *testing.T) {
	tests := []string{"0.01", "88.88", "100", "100000000.00", "12.3"}
	for _, amount := range tests {
		t.Run(amount, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
			_, err := svc.PreCreate(context.Background(), trade.PreCreateRequest{
				OutTradeNo: "T1", TotalAmount: amount, Subject: "标题"})
			if err != nil {
				t.Fatalf("PreCreate() error = %v", err)
			}
			var biz struct {
				TotalAmount string `json:"total_amount"`
			}
			if err := json.Unmarshal(executor.req.Params.JSON, &biz); err != nil {
				t.Fatal(err)
			}
			if biz.TotalAmount != amount {
				t.Fatalf("total_amount = %q, want %q (exact decimal passthrough)", biz.TotalAmount, amount)
			}
		})
	}
}
