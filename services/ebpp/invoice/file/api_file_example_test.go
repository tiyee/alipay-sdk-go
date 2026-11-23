package file_test

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
	file "github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/file"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_InvoiceFileOutputQuery 展示 alipay.ebpp.invoice.file.output.query 的调用和响应读取。
// 示例只演示调用流程，真实请求的必填条件与业务状态以方法 GoDoc 链接的官方契约为准。
func ExampleService_InvoiceFileOutputQuery() {
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

	svc, err := file.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	openId := "074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5"
	invoiceCode := "123"
	invoiceNo := "111"
	response, err := svc.InvoiceFileOutputQuery(ctx, file.InvoiceFileOutputQueryRequest{
		OpenId:      &openId,
		InvoiceCode: &invoiceCode,
		InvoiceNo:   &invoiceNo,
	})
	if err != nil {
		fmt.Println("接口调用失败")
		return
	}
	if response.Data.FileType == nil {
		fmt.Println("响应字段缺失")
		return
	}
	fmt.Println(*response.Data.FileType)
}

// TestSandboxService_InvoiceFileOutputQuery 使用沙箱证书与外部请求参数调用 alipay.ebpp.invoice.file.output.query。
func TestSandboxService_InvoiceFileOutputQuery(t *testing.T) {
	sandbox.Run(t, "alipay.ebpp.invoice.file.output.query", file.NewService, (*file.Service).InvoiceFileOutputQuery)
}
