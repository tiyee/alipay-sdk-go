package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/notify"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestSignedNotificationsRequireExpectedSeller(t *testing.T) {
	verifier, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ seller, want string }{{"", "fail"}, {"other", "fail"}, {"merchant", "success"}} {
		t.Run("seller="+tc.seller, func(t *testing.T) {
			h := notifyHandler{verifier: verifier, appID: "app", sellerID: "merchant", orders: newMemoryOrders(), logger: slog.New(slog.DiscardHandler), notifySvc: trade.NotificationDecoder()}
			form := url.Values{"notify_type": {"trade_status_sync"}, "notify_time": {"2026-11-05 12:00:00"}, "notify_id": {"id"}, "app_id": {"app"}, "auth_app_id": {"app"}, "out_trade_no": {"ORDER-20261105-001"}, "trade_no": {"trade"}, "trade_status": {"TRADE_SUCCESS"}, "total_amount": {"88.88"}, "receipt_amount": {"88.88"}}
			if tc.seller != "" {
				form.Set("seller_id", tc.seller)
			}
			signed := testutil.SignNotify(t, form)
			w := httptest.NewRecorder()
			h.serveNotify(w, httptest.NewRequest(http.MethodPost, "/notify", strings.NewReader(signed.Encode())))
			if w.Body.String() != tc.want || h.orders.processed["id"] != (tc.want == "success") || (h.orders.paidOrders["ORDER-20261105-001"] != "") != (tc.want == "success") {
				t.Fatalf("ack=%q, orders=%+v", w.Body.String(), h.orders)
			}
		})
	}
}

func TestAmbiguousNotificationDoesNotPayEitherOrder(t *testing.T) {
	verifier, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync)
	if err != nil {
		t.Fatal(err)
	}
	orders := newMemoryOrders()
	orders.amounts["original"] = "0.01"
	h := notifyHandler{verifier: verifier, appID: "app", sellerID: "merchant", orders: orders, logger: slog.New(slog.DiscardHandler), notifySvc: trade.NotificationDecoder()}
	fields := func(order, tradeNo, amount string) url.Values {
		return url.Values{
			"notify_type": {"trade_status_sync"}, "notify_time": {"2026-11-05 12:00:00"}, "notify_id": {"id"},
			"out_trade_no": {order}, "receipt_amount": {amount}, "seller_id": {"merchant"}, "total_amount": {amount},
			"trade_no": {tradeNo}, "trade_status": {"TRADE_SUCCESS"},
		}
	}
	original := fields("original", "tradeA", "0.01")
	changed := fields("ORDER-20261105-001", "tradeB", "88.88")
	originalSuffix := testutil.NotifySignContent(original)
	original.Set("body", "memo&"+testutil.NotifySignContent(changed)+"&z=end")
	changed.Set("body", "memo")
	changed.Set("z", "end&"+originalSuffix)
	for _, f := range []url.Values{original, changed} {
		f.Set("app_id", "app")
		f.Set("auth_app_id", "app")
	}
	if testutil.NotifySignContent(original) != testutil.NotifySignContent(changed) {
		t.Fatal("fixture requires identical signed bytes")
	}
	signed := testutil.SignNotify(t, original)
	changed.Set("sign", signed.Get("sign"))
	changed.Set("sign_type", "RSA2")
	for _, f := range []url.Values{signed, changed} {
		w := httptest.NewRecorder()
		h.serveNotify(w, httptest.NewRequest(http.MethodPost, "/notify", strings.NewReader(f.Encode())))
		if w.Body.String() != "fail" || len(orders.paidOrders) != 0 || len(orders.processed) != 0 {
			t.Fatalf("ack %q, orders %+v", w.Body.String(), orders)
		}
	}
}
