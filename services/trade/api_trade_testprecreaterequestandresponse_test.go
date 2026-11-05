package trade_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestPreCreateRequestAndResponse(t *testing.T) {
	qr := "https://qr.alipay.com/bavh4wjlxf12tper3a"
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(fmt.Sprintf(
		`{"code":"10000","msg":"Success","out_trade_no":"20150320010101001","qr_code":%q}`, qr))
	resp, err := svc.PreCreate(context.Background(), trade.PreCreateRequest{
		OutTradeNo:  "20150320010101001",
		TotalAmount: "88.88",
		Subject:     "Iphone6 16G",
		GoodsDetail: []trade.GoodsDetail{{
			GoodsID: "apple-01", GoodsName: "ipad", Quantity: 1, Price: "2000",
		}},
	})
	if err != nil {
		t.Fatalf("PreCreate() error = %v", err)
	}
	if resp.Data.QRCode == nil || *resp.Data.QRCode != qr {
		t.Fatalf("QRCode = %+v", resp.Data.QRCode)
	}
	if resp.Data.OutTradeNo != "20150320010101001" {
		t.Fatalf("OutTradeNo = %q", resp.Data.OutTradeNo)
	}
	var biz map[string]json.RawMessage
	if err := json.Unmarshal(executor.req.Params.JSON, &biz); err != nil {
		t.Fatal(err)
	}
	if got := string(biz["total_amount"]); got != `"88.88"` {
		t.Errorf("total_amount = %s, want exact decimal string", got)
	}
	if got := string(biz["subject"]); got != `"Iphone6 16G"` {
		t.Errorf("subject = %s", got)
	}
	goods := `[{"goods_id":"apple-01","goods_name":"ipad","quantity":1,"price":"2000"}]`
	if got := string(biz["goods_detail"]); got != goods {
		t.Errorf("goods_detail = %s, want %s", got, goods)
	}
	if _, ok := biz["product_code"]; ok {
		t.Error("unset optional fields must be omitted")
	}
}

func TestPreCreatePaymentAndSettlementFieldPlacement(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult("{}")
	_, err := svc.PreCreate(t.Context(), trade.PreCreateRequest{
		OutTradeNo: "order", TotalAmount: "1.00", Subject: "item", PaymentType: ptr(trade.PaymentTypeAgentPay),
		BkAgentReqInfo: &trade.BkAgentReqInfo{AcqCode: ptr("agent")},
		SettleInfo:     &trade.SettleInfo{SettlePeriodTime: ptr("7d"), SettleDetailInfos: []trade.SettleDetailInfo{{TransInType: "defaultSettle"}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]json.RawMessage
	if err := json.Unmarshal(executor.req.Params.JSON, &body); err != nil {
		t.Fatal(err)
	}
	if string(body["payment_type"]) != `"agent_pay"` || string(body["bkagent_req_info"]) != `{"acq_code":"agent"}` {
		t.Fatalf("bad payment fields %s", executor.req.Params.JSON)
	}
	var settlement struct {
		Time    string            `json:"settle_period_time"`
		Details []json.RawMessage `json:"settle_detail_infos"`
	}
	if err := json.Unmarshal(body["settle_info"], &settlement); err != nil {
		t.Fatal(err)
	}
	if settlement.Time != "7d" || len(settlement.Details) != 1 || string(settlement.Details[0]) != `{"trans_in_type":"defaultSettle"}` {
		t.Fatalf("bad settlement: %s", body["settle_info"])
	}
}
