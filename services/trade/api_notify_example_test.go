package trade_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/notify"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleNotificationDecoder 展示平台证书模式的普通支付通知验签和业务校验。
// ALIPAY_NOTIFY_FORM 指向收到的原始 URL 编码 POST 表单文件，不能重编码或伪造签名。
// 实际 HTTP 回调中直接将收到的 *http.Request 交给 ParseRequest；
// 校验订单、金额、收款方与 APPID，持久化业务状态并保证幂等后才应答 success。
// 示例无 Output 注释，默认 go test 只编译，不读取真实配置。
func ExampleNotificationDecoder() {
	// 从配置装载同一沙箱应用的私钥、应用证书、支付宝公钥证书和根证书。
	// 普通支付通知不携带证书标识，ALIPAY_CERT 只能包含一张可信支付宝证书。
	privateKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_PRIVATE_KEY"))
	if err != nil {
		fmt.Println("读取应用私钥失败")
		return
	}
	privateKey, err := utils.ParsePrivateKeyPKCS8(string(privateKeyPEM))
	if err != nil {
		fmt.Println("解析应用私钥失败")
		return
	}
	appCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_CERT"))
	if err != nil {
		fmt.Println("读取应用证书失败")
		return
	}
	appCert, err := utils.ParseCertificatePEM(string(appCertPEM))
	if err != nil {
		fmt.Println("解析应用证书失败")
		return
	}
	alipayCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_CERT"))
	if err != nil {
		fmt.Println("读取支付宝公钥证书失败")
		return
	}
	alipayCerts, err := utils.ParseCertificatesPEM(string(alipayCertPEM))
	if err != nil {
		fmt.Println("解析支付宝公钥证书失败")
		return
	}
	rootCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_ROOT_CERT"))
	if err != nil {
		fmt.Println("读取支付宝根证书失败")
		return
	}
	rootCerts, err := utils.ParseRootCertificatesPEM(string(rootCertPEM))
	if err != nil {
		fmt.Println("解析支付宝根证书失败")
		return
	}
	credentials, err := auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner:          privateKey,
		AppCertificate:     appCert,
		AlipayCertificates: alipayCerts,
		RootCertificates:   rootCerts,
	})
	if err != nil {
		fmt.Println("创建平台证书凭证失败，请检查证书有效期、私钥匹配与根信任")
		return
	}

	if len(alipayCerts) != 1 {
		fmt.Println("通知验签需要唯一的支付宝公钥证书")
		return
	}
	handler, err := notify.NewHandler(credentials, notify.PolicyTradeStatusSync)
	if err != nil {
		fmt.Println("创建通知处理器失败")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rawForm, err := os.ReadFile(os.Getenv("ALIPAY_NOTIFY_FORM"))
	if err != nil {
		fmt.Println("读取通知表单失败")
		return
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://merchant.example.com/alipay/notify", strings.NewReader(string(rawForm)))
	if err != nil {
		fmt.Println("构造通知请求失败")
		return
	}
	defer request.Body.Close()
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	notification, err := notify.ParseRequest(ctx, handler, request, trade.NotificationDecoder())
	if err != nil {
		fmt.Println("respond fail")
		return
	}
	// 示例内存订单；实际应用从持久化订单读取预期值。
	notifyOrder := struct{ id, amount, appID string }{
		id: "ORDER-20261105-001", amount: "88.88", appID: os.Getenv("ALIPAY_APP_ID"),
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
	// 在此以 NotifyID 去重并在数据库事务中更新订单状态。
	// 事务提交成功后向支付宝原样应答 success；失败则应答 fail。
	// 此示例尚未接入持久化处理，不能确认付款已入账。
	fmt.Println("respond fail: payment persistence required")
}
