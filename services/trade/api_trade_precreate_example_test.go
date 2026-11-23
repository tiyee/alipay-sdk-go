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

// ExampleService_PreCreate 展示当面付预下单：二维码有效时间 2 小时，
// qr_code 缺失时按业务响应处理。V3 REST 调用不提供 notify_url
// 调用选项，支付结果以查询接口与异步通知为准。
func ExampleService_PreCreate() {
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

	svc, err := trade.NewService(client)
	if err != nil {
		fmt.Println("操作失败")
		return
	}
	resp, err := svc.PreCreate(ctx, trade.PreCreateRequest{
		OutTradeNo:  "20150320010101001",
		TotalAmount: "88.88",
		Subject:     "Iphone6 16G",
	})
	if err != nil {
		fmt.Println("操作失败")
		return
	}
	if resp.Data.QRCode == nil {
		fmt.Println("响应未包含二维码")
		return
	}
	fmt.Println(*resp.Data.QRCode)
}

// TestSandboxService_PreCreate 使用沙箱证书与外部请求参数调用 alipay.trade.precreate。
func TestSandboxService_PreCreate(t *testing.T) {
	sandbox.Run(t, "alipay.trade.precreate", trade.NewService, (*trade.Service).PreCreate)
}
