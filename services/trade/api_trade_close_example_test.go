package trade_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/sandbox"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_Close 展示未支付订单关闭：先用查询确认交易仍处于
// WAIT_BUYER_PAY，再发起关单；已支付订单绝不进入关闭流程
// （已支付交易应走退款）。
func ExampleService_Close() {
	// 从配置装载同一沙箱应用的私钥、应用证书、支付宝公钥证书和根证书。
	// ALIPAY_CERT 指向 PEM 文件；REST 证书轮换期可包含多张支付宝证书。
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

	// Client 初始化后可并发复用；生产接入需同时更换环境与全部凭证。
	client, err := core.NewClient(credentials,
		option.WithAppID(os.Getenv("ALIPAY_APP_ID")),
		option.WithEnvironment(core.EnvironmentSandbox),
	)
	if err != nil {
		fmt.Println("创建 Client 失败")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 第一步：查询订单当前状态。

	svc, err := trade.NewService(client)
	if err != nil {
		fmt.Println("操作失败")
		return
	}
	query, err := svc.Query(ctx, trade.QueryRequest{
		OutTradeNo: "6823789339978248",
	})
	if err != nil {
		fmt.Println("操作失败")
		return
	}
	if query.Data.TradeStatus != trade.TradeStatusWaitBuyerPay {
		fmt.Println("not closable:", query.Data.TradeStatus)
		return
	}

	// 第二步：仅未支付订单发起关单。

	resp, err := svc.Close(ctx, trade.CloseRequest{
		OutTradeNo: "6823789339978248",
	})
	if err != nil {
		fmt.Println("操作失败")
		return
	}
	fmt.Println("closed", resp.Result.HTTPStatus)
}

// TestSandboxService_Close 使用沙箱证书与外部请求参数调用 alipay.trade.close。
func TestSandboxService_Close(t *testing.T) {
	sandbox.Run(t, "alipay.trade.close", trade.NewService, (*trade.Service).Close)
}
