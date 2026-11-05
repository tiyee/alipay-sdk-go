package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// TestNewAPIResponseMapping 验证撤销与异步退款的状态语义映射。
func TestNewAPIResponseMapping(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"trade_no":"T1","out_trade_no":"O1","retry_flag":"Y","action":"refund"}`)
	resp, err := svc.Cancel(context.Background(), trade.CancelRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Cancel() error = %v", err)
	}
	if resp.Data.Action == nil || *resp.Data.Action != trade.CancelActionRefund {
		t.Fatalf("Action = %+v", resp.Data.Action)
	}
	if resp.Data.RetryFlag == nil || *resp.Data.RetryFlag != "Y" {
		t.Fatalf("RetryFlag = %+v", resp.Data.RetryFlag)
	}

}
