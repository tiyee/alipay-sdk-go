package trade_test

import (
	"encoding/json"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/services/trade/refund"
)

func TestRESTNotifyURLInJSON(t *testing.T) {
	callback := "https://merchant.example/notify?id=a+b&branch=2"
	checkNotifyURL(t, trade.PayRequest{NotifyURL: &callback}, callback)
	checkNotifyURL(t, trade.CreateRequest{NotifyURL: &callback}, callback)
	checkNotifyURL(t, trade.PreCreateRequest{NotifyURL: &callback}, callback)
	checkNotifyURL(t, trade.CloseRequest{NotifyURL: &callback}, callback)
	checkNotifyURL(t, refund.RefundApplyRequest{NotifyURL: &callback}, callback)
	checkNotifyURL(t, trade.PreCreateRequest{}, "")
}

func checkNotifyURL[T any](t *testing.T, req T, want string) {
	t.Helper()
	p, err := core.JSONBodyEncoder[T]().Encode(req)
	if err != nil {
		t.Fatal(err)
	}
	var body struct {
		NotifyURL *string `json:"notify_url"`
	}
	if err := json.Unmarshal(p.JSON, &body); err != nil {
		t.Fatal(err)
	}
	if want == "" && body.NotifyURL != nil || want != "" && (body.NotifyURL == nil || *body.NotifyURL != want) {
		t.Errorf("%T encoded notify_url incorrectly: %s", req, p.JSON)
	}
}
