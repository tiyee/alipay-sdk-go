package trade_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestServiceDecodeFailureIsProtocolError(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"code":"10000","total_amount":1.00}`)
	_, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if !errors.As(err, new(*core.ProtocolError)) {
		t.Fatalf("error = %v, want ProtocolError", err)
	}
	if !strings.Contains(fmt.Sprint(err), "decode") {
		t.Errorf("error = %v", err)
	}
}
