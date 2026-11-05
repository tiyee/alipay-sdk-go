package advance_test

import (
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade/advance"
)

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := advance.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}
