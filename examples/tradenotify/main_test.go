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
