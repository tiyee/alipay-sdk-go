package trade_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestPaymentRoyaltyWireTypes(t *testing.T) {
	svc, e := newRecordingService(t)
	e.result = verifiedResult("{}")
	info := trade.RoyaltyInfo{RoyaltyDetailInfos: []trade.PaymentRoyaltyDetail{{AmountPercentage: ptr("20"), SerialNo: ptr(int64(1)), BatchNo: ptr("batch"), OutRelationID: ptr("relation")}}}
	_, err := svc.Pay(t.Context(), trade.PayRequest{OutTradeNo: "A", TotalAmount: "1.00", Subject: "title", AuthCode: ptr("code"), RoyaltyInfo: &info})
	if err != nil {
		t.Fatal(err)
	}
	check := func() {
		var body struct {
			Info struct {
				Details []struct {
					AmountPercentage string `json:"amount_percentage"`
					SerialNo         int64  `json:"serial_no"`
					BatchNo          string `json:"batch_no"`
					OutRelationID    string `json:"out_relation_id"`
				} `json:"royalty_detail_infos"`
			} `json:"royalty_info"`
		}
		if err := json.Unmarshal(e.req.Params.JSON, &body); err != nil {
			t.Fatal(err)
		}
		if len(body.Info.Details) != 1 {
			t.Fatalf("missing detail: %s", e.req.Params.JSON)
		}
		d := body.Info.Details[0]
		if d.AmountPercentage != "20" || d.SerialNo != 1 || d.BatchNo != "batch" || d.OutRelationID != "relation" {
			t.Fatalf("unexpected detail %+v", d)
		}
	}
	check()
	_, err = svc.Create(t.Context(), trade.CreateRequest{OutTradeNo: "A", TotalAmount: "1.00", Subject: "title", BuyerID: ptr("buyer"), RoyaltyInfo: &info})
	if err != nil {
		t.Fatal(err)
	}
	check()
	refund, err := json.Marshal(trade.RoyaltyDetailInfo{AmountPercentage: ptr(int64(20))})
	if err != nil || !bytes.Equal(refund, []byte(`{"amount_percentage":20}`)) {
		t.Fatalf("unexpected refund detail %s, %v", refund, err)
	}
}

func TestPayValidatesAmountAndNotificationTextBeforeExecution(t *testing.T) {
	for _, tc := range []struct{ amount, body, field string }{
		{"100000000", "a&b", ""}, {"100000000.00", "", ""}, {"99999999.99", "", ""}, {"0.01", "", ""},
		{"100000000.01", "", "total_amount"}, {"100000001", "", "total_amount"}, {"100000000000000000000", "", "total_amount"},
		{".01", "", "total_amount"}, {"1.", "", "total_amount"}, {"1", "note&buyer_id=another", "body"},
	} {
		svc, e := newRecordingService(t)
		e.result = verifiedResult("{}")
		_, err := svc.Pay(t.Context(), trade.PayRequest{OutTradeNo: "order", Subject: "title", TotalAmount: tc.amount, AuthCode: ptr("code"), Body: ptr(tc.body)})
		if tc.field == "" {
			if err != nil || !e.called {
				t.Fatalf("valid amount %q: %v", tc.amount, err)
			}
			continue
		}
		var val *core.ValidationError
		if !errors.As(err, &val) || val.Field != tc.field || e.called {
			t.Fatalf("%+v: error %v, called %v", tc, err, e.called)
		}
	}
}
