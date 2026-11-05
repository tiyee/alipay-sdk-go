// Command pagepay 演示页面与 App 支付产物构建：构建过程不发起
// HTTP 请求，产物由用户浏览器或支付宝客户端 SDK 提交。
//
// 链路：凭证装载 → 组装业务参数并签名（页面/唤端产物契约：
// method / version=1.0 / biz_content，提交到产物网关 gateway.do，
// 不走 V3 REST）→ 产出三类产物：
//
//	page.pay  POST：自动提交的 HTML Form（推荐用于电脑网站支付）
//	wap.pay   GET ：跳转 URL（手机网站支付）
//	app.pay   产物：订单串 orderStr，交给支付宝客户端 SDK 唤起支付
//
// 产物签名可使用开放平台密钥工具的"签名"功能独立核对。
// 注意：跳转页与客户端返回不构成支付完成的可信证据，支付结果以
// 异步通知与 alipay.trade.query 查询为准。
//
// 配置全部来自环境变量，示例不包含任何真实凭证：
//
//	ALIPAY_APP_ID           本应用 APPID（必填）
//	ALIPAY_APP_PRIVATE_KEY  应用私钥 PKCS#8 PEM 文件路径（必填）
//	ALIPAY_PUBLIC_KEY       支付宝公钥 PEM 文件路径（必填）
//	ALIPAY_ENV              production 或 sandbox（默认 production）
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appID := os.Getenv("ALIPAY_APP_ID")
	privateKeyPath := os.Getenv("ALIPAY_APP_PRIVATE_KEY")
	publicKeyPath := os.Getenv("ALIPAY_PUBLIC_KEY")
	if appID == "" || privateKeyPath == "" || publicKeyPath == "" {
		logger.Error("ALIPAY_APP_ID, ALIPAY_APP_PRIVATE_KEY and ALIPAY_PUBLIC_KEY are required")
		os.Exit(2)
	}
	privateKeyPEM, err := os.ReadFile(privateKeyPath)
	if err != nil {
		logger.Error("read app private key", slog.String("err", err.Error()))
		os.Exit(2)
	}
	publicKeyPEM, err := os.ReadFile(publicKeyPath)
	if err != nil {
		logger.Error("read alipay public key", slog.String("err", err.Error()))
		os.Exit(2)
	}
	signer, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(string(privateKeyPEM)))
	if err != nil {
		logger.Error("parse app private key", slog.String("err", err.Error()))
		os.Exit(2)
	}
	verifier, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(string(publicKeyPEM)))
	if err != nil {
		logger.Error("parse alipay public key", slog.String("err", err.Error()))
		os.Exit(2)
	}
	credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       signer,
		AlipayPublicKey: verifier,
	})
	if err != nil {
		logger.Error("build credentials", slog.String("err", err.Error()))
		os.Exit(2)
	}

	opts := []core.ClientOption{option.WithAppID(appID)}
	switch os.Getenv("ALIPAY_ENV") {
	case "", "production":
		// 环境须显式选择；WithEnvironment 成对设置 V3 REST 基地址与
		// 产物网关，产物提交到对应环境的 gateway.do。
		opts = append(opts, option.WithEnvironment(core.EnvironmentProduction))
	case "sandbox":
		opts = append(opts, option.WithEnvironment(core.EnvironmentSandbox))
	default:
		logger.Error("ALIPAY_ENV must be production or sandbox")
		os.Exit(2)
	}
	products, err := trade.NewPaymentProducts(credentials, opts...)
	if err != nil {
		logger.Error("build payment products", slog.String("err", err.Error()))
		os.Exit(2)
	}

	ctx := context.Background()
	pageReq := trade.PagePayRequest{
		OutTradeNo:  "demo-page-0001",
		TotalAmount: "0.01",
		Subject:     "示例商品",
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		ReturnURL:   strPtr("https://merchant.example.com/return"),
	}
	form, err := products.PagePayForm(ctx, pageReq)
	if err != nil {
		logger.Error("build page pay form", slog.String("err", err.Error()))
		os.Exit(1)
	}
	fmt.Println("== page.pay HTML Form ==")
	fmt.Println(form)

	wapURL, err := products.WapPayURL(ctx, trade.PagePayRequest{
		OutTradeNo:  "demo-wap-0001",
		TotalAmount: "0.01",
		Subject:     "示例商品",
		ProductCode: "QUICK_WAP_WAY",
	})
	if err != nil {
		logger.Error("build wap pay url", slog.String("err", err.Error()))
		os.Exit(1)
	}
	fmt.Println("== wap.pay redirect URL ==")
	fmt.Println(wapURL)

	orderStr, err := products.AppPayOrderString(ctx, trade.AppPayRequest{
		OutTradeNo:  "demo-app-0001",
		TotalAmount: "0.01",
		Subject:     "示例商品",
		ProductCode: "QUICK_MSECURITY_PAY",
	})
	if err != nil {
		logger.Error("build app pay order string", slog.String("err", err.Error()))
		os.Exit(1)
	}
	fmt.Println("== app.pay orderStr ==")
	fmt.Println(orderStr)
}

func strPtr(v string) *string { return &v }
