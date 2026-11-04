package trade_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/auth/verifiers"
	"github.com/tiyee/go-alipay/core/notify"
	"github.com/tiyee/go-alipay/internal/testutil"
	"github.com/tiyee/go-alipay/services/trade"
)

// 测试目标：通知模型严格映射官方《异步通知说明》参数表
// （https://opendocs.alipay.com/open/194/103296）；必选参数缺失与
// 非法 JSON 明细报错；未知枚举值保留原值；与 core/notify 组合完成
// 端到端签名验签链路。

// baseNotifyForm 返回官方参数表中全部必选参数与代表性可选参数。
func baseNotifyForm() url.Values {
	return url.Values{
		"notify_time":    {"2026-11-05 10:00:00"},
		"notify_type":    {"trade_status_sync"},
		"notify_id":      {"ac05099524730693a8b330c5ecf72da9786"},
		"trade_no":       {"2026110522001400001234567890"},
		"app_id":         {"2021000000000001"},
		"auth_app_id":    {"2021000000000001"},
		"out_trade_no":   {"ORDER-20261105-001"},
		"trade_status":   {"TRADE_SUCCESS"},
		"total_amount":   {"20.00"},
		"receipt_amount": {"15.00"},
	}
}

func TestNotificationDecoderFullMapping(t *testing.T) {
	values := baseNotifyForm()
	values.Set("out_biz_no", "HZRF001")
	values.Set("buyer_id", "2088102177824415")
	values.Set("buyer_open_id", "openid-001")
	values.Set("buyer_logon_id", "180****0062")
	values.Set("seller_id", "2088101106499276")
	values.Set("seller_email", "zhuzxxxxxx@alitest.com")
	values.Set("invoice_amount", "10.00")
	values.Set("buyer_pay_amount", "13.88")
	values.Set("point_amount", "12.00")
	values.Set("refund_fee", "2.58")
	values.Set("send_back_fee", "2.08")
	values.Set("subject", "XXX交易")
	values.Set("body", "XXX交易内容")
	values.Set("gmt_create", "2015-04-27 15:45:57")
	values.Set("gmt_payment", "2015-04-27 15:45:57")
	values.Set("gmt_refund", "2015-04-28 15:45:57.320")
	values.Set("gmt_close", "2015-04-29 15:45:57")
	values.Set("fund_bill_list", `[{"amount":"15.00","fundChannel":"ALIPAYACCOUNT"}]`)
	values.Set("voucher_detail_list",
		`[{"voucherId":"2015102600073002039000002D5O","templateId":"20171030000730015359000EMZP0",`+
			`"name":"5 元代金券","type":"ALIPAY_BIZ_VOUCHER","amount":"10.00",`+
			`"merchantContribute":"9.00","otherContribute":"1.00",`+
			`"otherContributeDetail":[{"contributeType":"PLATFORM","contributeAmount":"0.18"}],`+
			`"memo":"学生专用优惠"}]`)
	values.Set("biz_settle_mode", "PERIOD")

	n, err := trade.NotificationDecoder().Decode(values)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	checks := []struct {
		name string
		got  string
		want string
	}{
		{"notify_time", n.NotifyTime, "2026-11-05 10:00:00"},
		{"notify_id", n.NotifyID, "ac05099524730693a8b330c5ecf72da9786"},
		{"trade_no", n.TradeNo, "2026110522001400001234567890"},
		{"app_id", n.AppID, "2021000000000001"},
		{"auth_app_id", n.AuthAppID, "2021000000000001"},
		{"out_trade_no", n.OutTradeNo, "ORDER-20261105-001"},
		{"out_biz_no", *n.OutBizNo, "HZRF001"},
		{"buyer_id", *n.BuyerID, "2088102177824415"},
		{"buyer_open_id", *n.BuyerOpenID, "openid-001"},
		{"buyer_logon_id", *n.BuyerLogonID, "180****0062"},
		{"seller_id", *n.SellerID, "2088101106499276"},
		{"seller_email", *n.SellerEmail, "zhuzxxxxxx@alitest.com"},
		{"total_amount", n.TotalAmount, "20.00"},
		{"receipt_amount", n.ReceiptAmount, "15.00"},
		{"invoice_amount", *n.InvoiceAmount, "10.00"},
		{"buyer_pay_amount", *n.BuyerPayAmount, "13.88"},
		{"point_amount", *n.PointAmount, "12.00"},
		{"refund_fee", *n.RefundFee, "2.58"},
		{"send_back_fee", *n.SendBackFee, "2.08"},
		{"subject", *n.Subject, "XXX交易"},
		{"body", *n.Body, "XXX交易内容"},
		{"gmt_create", *n.GmtCreate, "2015-04-27 15:45:57"},
		{"gmt_payment", *n.GmtPayment, "2015-04-27 15:45:57"},
		{"gmt_refund", *n.GmtRefund, "2015-04-28 15:45:57.320"},
		{"gmt_close", *n.GmtClose, "2015-04-29 15:45:57"},
	}
	for _, check := range checks {
		if check.got != check.want {
			t.Errorf("%s = %q, want %q", check.name, check.got, check.want)
		}
	}
	if n.NotifyType != trade.NotifyTypeTradeStatusSync {
		t.Errorf("notify_type = %q", n.NotifyType)
	}
	if n.TradeStatus != trade.TradeStatusSuccess {
		t.Errorf("trade_status = %q", n.TradeStatus)
	}
	if n.BizSettleMode == nil || *n.BizSettleMode != trade.BizSettleModePeriod {
		t.Errorf("biz_settle_mode = %v", n.BizSettleMode)
	}
	if len(n.FundBillList) != 1 || n.FundBillList[0].FundChannel != "ALIPAYACCOUNT" ||
		n.FundBillList[0].Amount != "15.00" {
		t.Errorf("fund_bill_list = %+v", n.FundBillList)
	}
	if len(n.VoucherDetailList) != 1 {
		t.Fatalf("voucher_detail_list = %+v", n.VoucherDetailList)
	}
	v := n.VoucherDetailList[0]
	if v.VoucherID != "2015102600073002039000002D5O" || v.Name != "5 元代金券" ||
		v.Type != trade.NotifyVoucherTypeBiz || v.Amount != "10.00" ||
		*v.MerchantContribute != "9.00" || *v.OtherContribute != "1.00" ||
		*v.TemplateID != "20171030000730015359000EMZP0" || *v.Memo != "学生专用优惠" {
		t.Errorf("voucher detail = %+v", v)
	}
	if len(v.OtherContributeDetail) != 1 ||
		v.OtherContributeDetail[0].ContributeType != trade.ContributeTypePlatform ||
		v.OtherContributeDetail[0].ContributeAmount != "0.18" {
		t.Errorf("other contribute detail = %+v", v.OtherContributeDetail)
	}
	if !n.Paid() {
		t.Error("TRADE_SUCCESS should be paid")
	}
}

func TestNotificationDecoderUnknownValuesRetained(t *testing.T) {
	values := baseNotifyForm()
	values.Set("trade_status", "TRADE_FROZEN")
	values.Set("biz_settle_mode", "FUTURE_MODE")
	values.Set("fund_bill_list", `[{"amount":"1.00","fundChannel":"FUTURE_CHANNEL"}]`)
	values.Set("voucher_detail_list",
		`[{"voucherId":"v1","name":"新券","type":"ALIPAY_FUTURE_VOUCHER","amount":"1.00"}]`)

	n, err := trade.NotificationDecoder().Decode(values)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if n.TradeStatus != "TRADE_FROZEN" {
		t.Errorf("trade_status = %q, want original value", n.TradeStatus)
	}
	if n.Paid() {
		t.Error("unknown trade status must not be treated as paid")
	}
	if n.BizSettleMode == nil || *n.BizSettleMode != "FUTURE_MODE" {
		t.Errorf("biz_settle_mode = %v", n.BizSettleMode)
	}
	if len(n.FundBillList) != 1 || n.FundBillList[0].FundChannel != "FUTURE_CHANNEL" {
		t.Errorf("fund_bill_list = %+v", n.FundBillList)
	}
	if len(n.VoucherDetailList) != 1 || n.VoucherDetailList[0].Type != "ALIPAY_FUTURE_VOUCHER" {
		t.Errorf("voucher_detail_list = %+v", n.VoucherDetailList)
	}
}

func TestNotificationDecoderPaidSemantics(t *testing.T) {
	cases := []struct {
		status trade.TradeStatus
		paid   bool
	}{
		{trade.TradeStatusSuccess, true},
		{trade.TradeStatusFinished, true},
		{trade.TradeStatusWaitBuyerPay, false},
		{trade.TradeStatusClosed, false},
		{"UNKNOWN_NEW_STATUS", false},
	}
	for _, tc := range cases {
		values := baseNotifyForm()
		values.Set("trade_status", string(tc.status))
		n, err := trade.NotificationDecoder().Decode(values)
		if err != nil {
			t.Fatalf("decode %s: %v", tc.status, err)
		}
		if n.Paid() != tc.paid {
			t.Errorf("Paid() for %s = %v, want %v", tc.status, n.Paid(), tc.paid)
		}
	}
}

func TestNotificationDecoderMissingRequired(t *testing.T) {
	for _, name := range []string{
		"notify_time", "notify_type", "notify_id", "trade_no", "app_id", "auth_app_id",
		"out_trade_no", "trade_status", "total_amount", "receipt_amount",
	} {
		t.Run(name, func(t *testing.T) {
			values := baseNotifyForm()
			values.Del(name)
			_, err := trade.NotificationDecoder().Decode(values)
			valErr, ok := errors.AsType[*core.ValidationError](err)
			if !ok || valErr.Field != name {
				t.Fatalf("want validation error for %q, got %v", name, err)
			}
		})
	}
}

func TestNotificationDecoderInvalidJSONList(t *testing.T) {
	cases := []struct{ name, raw string }{
		{"fund_bill_list", "not-json"},
		{"voucher_detail_list", `{"not":"an-array"}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			values := baseNotifyForm()
			values.Set(tc.name, tc.raw)
			_, err := trade.NotificationDecoder().Decode(values)
			valErr, ok := errors.AsType[*core.ValidationError](err)
			if !ok || valErr.Field != tc.name {
				t.Fatalf("want validation error for %q, got %v", tc.name, err)
			}
		})
	}
}

// 端到端：网关侧签名 → core/notify 验签 → trade 领域映射。
func TestNotifyEndToEnd(t *testing.T) {
	ctx := context.Background()
	verifier := testutil.NewAlipayVerifier(t)
	form := baseNotifyForm()
	form.Set("subject", "测试 单据+&=A")
	form.Set("fund_bill_list", `[{"amount":"15.00","fundChannel":"ALIPAYACCOUNT"}]`)
	signed := testutil.SignNotify(t, form)

	got, err := notify.Parse(ctx, notify.PolicyTradeStatusSync, verifier,
		trade.NotificationDecoder(), signed)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got.Data.OutTradeNo != "ORDER-20261105-001" || !got.Data.Paid() {
		t.Fatalf("unexpected notification: %+v", *got.Data)
	}
	if got.Data.Subject == nil || *got.Data.Subject != "测试 单据+&=A" {
		t.Errorf("subject = %v", got.Data.Subject)
	}

	t.Run("tampered form fails", func(t *testing.T) {
		tampered := signed.Clone()
		tampered.Set("total_amount", "0.01")
		if _, err := notify.Parse(ctx, notify.PolicyTradeStatusSync, verifier,
			trade.NotificationDecoder(), tampered); err == nil {
			t.Fatal("tampered notification must fail")
		}
	})

	t.Run("wrong public key fails", func(t *testing.T) {
		wrongVerifier, err := verifiers.NewRSA2(testutil.AppPublicKeyPEM)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := notify.Parse(ctx, notify.PolicyTradeStatusSync, wrongVerifier,
			trade.NotificationDecoder(), signed); err == nil {
			t.Fatal("wrong public key must fail")
		}
	})
}
