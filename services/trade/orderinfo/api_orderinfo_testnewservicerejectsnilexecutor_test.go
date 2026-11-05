package orderinfo_test

import (
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade/orderinfo"
)

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := orderinfo.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}
