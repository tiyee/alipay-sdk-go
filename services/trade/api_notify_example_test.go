package trade_test

import (
	"context"
	"fmt"
	"net/url"

	"github.com/tiyee/go-alipay/core/auth/verifiers"
	"github.com/tiyee/go-alipay/core/notify"
	"github.com/tiyee/go-alipay/internal/testutil"
	"github.com/tiyee/go-alipay/services/trade"
)

// notifyOrder 是示例内存订单；真实系统替换为数据库与唯一约束。
var notifyOrder = struct {
	id     string
	amount string
	appID  string
}{
	id:     "ORDER-20261105-001",
	amount: "88.88",
	appID:  "2021000000000001",
}

// ExampleNotificationDecoder 展示普通支付通知的接入方式：
// 验签与映射交给 SDK，随后必须完成官方要求的业务校验
// （订单号、金额、收款方、APPID），并按 NotifyID 幂等处理，
// 全部通过后才应答 success。示例密钥为仓库专用虚构材料，
// 生产环境应从配置装载真实支付宝公钥。
func ExampleNotificationDecoder() {
	verifier, err := verifiers.NewRSA2(testutil.AlipayPublicKeyPEM)
	if err != nil {
		panic(err)
	}

	// 模拟支付宝 POST 到商户回调端点的表单（测试辅助签名）。
	form := url.Values{
		"notify_time":    {"2026-11-05 10:00:00"},
		"notify_type":    {"trade_status_sync"},
		"notify_id":      {"ac05099524730693a8b330c5ecf72da9786"},
		"trade_no":       {"2026110522001400001234567890"},
		"app_id":         {notifyOrder.appID},
		"auth_app_id":    {notifyOrder.appID},
		"out_trade_no":   {notifyOrder.id},
		"trade_status":   {"TRADE_SUCCESS"},
		"total_amount":   {notifyOrder.amount},
		"receipt_amount": {"88.88"},
		"seller_id":      {"2088101106499276"},
	}
	sign, err := testutil.SignWithAlipayKey([]byte(testutil.NotifySignContent(form)))
	if err != nil {
		panic(err)
	}
	form.Set("sign_type", "RSA2")
	form.Set("sign", sign)

	notification, err := notify.Parse(context.Background(),
		notify.PolicyTradeStatusSync, verifier, trade.NotificationDecoder(), form)
	if err != nil {
		// 验签失败等情形：应答 fail，支付宝会重发通知。
		fmt.Println("respond fail:", err)
		return
	}
	data := notification.Data

	// 以下业务校验是接入位置：验签只证明通知来自支付宝，
	// 不能证明通知对应本应用、本订单或本收款账户。
	switch {
	case data.AppID != notifyOrder.appID:
		fmt.Println("respond fail: unknown app_id")
		return
	case data.SellerID == nil || *data.SellerID != "2088101106499276":
		fmt.Println("respond fail: unknown seller")
		return
	case data.OutTradeNo != notifyOrder.id:
		fmt.Println("respond fail: unknown order")
		return
	case data.TotalAmount != notifyOrder.amount:
		fmt.Println("respond fail: amount mismatch")
		return
	}
	// 幂等接入位置：以 data.NotifyID（或订单状态机）去重，
	// 重复通知直接应答 success，不重复入账。
	if !data.Paid() {
		// WAIT_BUYER_PAY、TRADE_CLOSED 等状态按业务语义处理。
		fmt.Println("respond success:", data.TradeStatus)
		return
	}
	// 处理完成后必须原样应答 success（小写、无多余字符），
	// 否则支付宝会按 4m、10m、1h……的间隔重发通知。
	fmt.Println("respond success:", data.TradeStatus, data.TotalAmount)
	// Output:
	// respond success: TRADE_SUCCESS 88.88
}
