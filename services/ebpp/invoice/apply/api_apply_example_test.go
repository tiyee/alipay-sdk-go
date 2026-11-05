package apply_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/sandbox"
	apply "github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/apply"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_InvoiceApplyResultSync 展示 alipay.ebpp.invoice.apply.result.sync 的调用和响应读取。
// 示例只演示调用流程，真实请求的必填条件与业务状态以方法 GoDoc 链接的官方契约为准。
func ExampleService_InvoiceApplyResultSync() {
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

	svc, err := apply.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	applyId := "2023000000000000000000000001"
	resultCode := "SUCCESS"
	taxApplyId := "20160707399929991001"
	response, err := svc.InvoiceApplyResultSync(ctx, apply.InvoiceApplyResultSyncRequest{
		ApplyId:    &applyId,
		ResultCode: &resultCode,
		TaxApplyId: &taxApplyId,
	})
	if err != nil {
		fmt.Println("接口调用失败")
		return
	}
	if response.Data.RetryFlag == nil {
		fmt.Println("响应字段缺失")
		return
	}
	fmt.Println(*response.Data.RetryFlag)
}

// TestSandboxService_InvoiceApplyResultSync 使用沙箱证书与外部请求参数调用 alipay.ebpp.invoice.apply.result.sync。
func TestSandboxService_InvoiceApplyResultSync(t *testing.T) {
	sandbox.Run(t, "alipay.ebpp.invoice.apply.result.sync", apply.NewService, (*apply.Service).InvoiceApplyResultSync)
}

// ExampleService_InvoiceApplyInvUpload 展示 alipay.ebpp.invoice.apply.inv.upload 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 apply.InvoiceApplyInvUploadRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_InvoiceApplyInvUpload() {
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

	svc, err := apply.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	// 使用自己的业务参数；JSON 不包含公共参数或 biz_content 包装。
	requestJSON, err := os.ReadFile(os.Getenv("ALIPAY_EXAMPLE_REQUEST"))
	if err != nil {
		fmt.Println("读取请求文件失败")
		return
	}
	var request *apply.InvoiceApplyInvUploadRequest
	decoder := json.NewDecoder(bytes.NewReader(requestJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil || !json.Valid(requestJSON) || request == nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.InvoiceApplyInvUpload(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.IsSuccess == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.IsSuccess)
}

// ExampleService_InvoiceApplyStatusNotify 展示 alipay.ebpp.invoice.apply.status.notify 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 apply.InvoiceApplyStatusNotifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_InvoiceApplyStatusNotify() {
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

	svc, err := apply.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	// 使用自己的业务参数；JSON 不包含公共参数或 biz_content 包装。
	requestJSON, err := os.ReadFile(os.Getenv("ALIPAY_EXAMPLE_REQUEST"))
	if err != nil {
		fmt.Println("读取请求文件失败")
		return
	}
	var request *apply.InvoiceApplyStatusNotifyRequest
	decoder := json.NewDecoder(bytes.NewReader(requestJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil || !json.Valid(requestJSON) || request == nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.InvoiceApplyStatusNotify(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.UpdateSuccess == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.UpdateSuccess)
}
