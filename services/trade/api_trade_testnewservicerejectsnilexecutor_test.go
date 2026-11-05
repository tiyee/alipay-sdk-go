package trade_test

import (
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := trade.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}
