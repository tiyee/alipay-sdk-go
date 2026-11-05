package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestQueryResponseMapping(t *testing.T) {
	// 响应示例依据查询文档正常示例。
	node := `{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248",` +
		`"buyer_logon_id":"159****5620","trade_status":"TRADE_CLOSED",` +
		`"total_amount":"88.88","send_pay_date":"2014-11-27 15:45:57",` +
		`"buyer_pay_amount":"8.88","point_amount":"10","receipt_amount":"15.25",` +
		`"store_id":"NJ_S_001","terminal_id":"NJ_T_001",` +
		`"fund_bill_list":[{"fund_channel":"ALIPAYACCOUNT","amount":"10","real_amount":"11.21"}],` +
		`"store_name":"证大五道口店","buyer_user_id":"2088101117955611",` +
		`"buyer_user_type":"PRIVATE","mdiscount_amount":"88.88"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	data := resp.Data
	if data.TradeStatus != trade.TradeStatusClosed {
		t.Errorf("TradeStatus = %q", data.TradeStatus)
	}
	if data.TotalAmount != "88.88" || data.ReceiptAmount == nil || *data.ReceiptAmount != "15.25" {
		t.Errorf("amounts = %+v", data)
	}
	if len(data.FundBillList) != 1 || data.FundBillList[0].FundChannel != "ALIPAYACCOUNT" {
		t.Errorf("FundBillList = %+v", data.FundBillList)
	}
	if data.BuyerUserType == nil || *data.BuyerUserType != trade.BuyerUserTypePrivate {
		t.Errorf("BuyerUserType = %+v", data.BuyerUserType)
	}
	if data.StoreName == nil || *data.StoreName != "证大五道口店" {
		t.Errorf("StoreName = %+v", data.StoreName)
	}
	if data.TradeStatus.IsKnown() != true {
		t.Error("known status expected")
	}
}

func TestQueryOptionalDetails(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"voucher_detail_list":[{"id":"v1","amount":"0.20"}],"trade_settle_info":{"trade_unsettled_amount":"1.10","trade_settle_detail_list":[{"operation_type":"transfer","amount":"2.30"}]},"medical_insurance_info":"{\"medicareCardType\":\"1\"}","discount_goods_detail":"[]"}`)
	resp, err := svc.Query(t.Context(), trade.QueryRequest{TradeNo: "offline", QueryOptions: []trade.QueryOption{trade.QueryOptionTradeSettleInfo, trade.QueryOptionVoucherDetailList, trade.QueryOptionMedicalInsuranceInfo, trade.QueryOptionDiscountGoodsDetail}})
	if err != nil {
		t.Fatal(err)
	}
	d := resp.Data
	if len(d.VoucherDetailList) != 1 || d.VoucherDetailList[0].ID != "v1" || d.MedicalInsuranceInfo == nil || *d.MedicalInsuranceInfo != `{"medicareCardType":"1"}` || d.DiscountGoodsDetail == nil || *d.DiscountGoodsDetail != "[]" {
		t.Fatalf("optional fields = %+v", d)
	}
	s := d.TradeSettleInfo
	if s == nil || s.TradeUnsettledAmount == nil || *s.TradeUnsettledAmount != "1.10" || len(s.TradeSettleDetailList) != 1 || s.TradeSettleDetailList[0].OperationType == nil || *s.TradeSettleDetailList[0].OperationType != trade.SettleOperationTransfer {
		t.Fatalf("settlement = %+v", s)
	}
}
