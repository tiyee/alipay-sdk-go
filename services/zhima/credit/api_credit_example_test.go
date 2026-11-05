package credit_test

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
	credit "github.com/tiyee/alipay-sdk-go/services/zhima/credit"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_PayAfterUseCreditBizOrderQuery 展示 zhima.credit.payafteruse.creditbizorder.query 的调用和响应读取。
// 示例只演示调用流程，真实请求的必填条件与业务状态以方法 GoDoc 链接的官方契约为准。
func ExampleService_PayAfterUseCreditBizOrderQuery() {
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

	svc, err := credit.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	creditBizOrderId := "ZMCB20201021200102012"
	outOrderNo := "20202010212001384"
	response, err := svc.PayAfterUseCreditBizOrderQuery(ctx, credit.PayAfterUseCreditBizOrderQueryRequest{
		CreditBizOrderId: &creditBizOrderId,
		OutOrderNo:       &outOrderNo,
	})
	if err != nil {
		fmt.Println("接口调用失败")
		return
	}
	if response.Data.OrderStatus == nil {
		fmt.Println("响应字段缺失")
		return
	}
	fmt.Println(*response.Data.OrderStatus)
}

// TestSandboxService_PayAfterUseCreditBizOrderQuery 使用沙箱证书与外部请求参数调用 zhima.credit.payafteruse.creditbizorder.query。
func TestSandboxService_PayAfterUseCreditBizOrderQuery(t *testing.T) {
	sandbox.Run(t, "zhima.credit.payafteruse.creditbizorder.query", credit.NewService, (*credit.Service).PayAfterUseCreditBizOrderQuery)
}

// ExampleService_PayAfterUseCreditAgreementQuery 展示 zhima.credit.payafteruse.creditagreement.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PayAfterUseCreditAgreementQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PayAfterUseCreditAgreementQuery() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PayAfterUseCreditAgreementQueryRequest
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
	response, err := svc.PayAfterUseCreditAgreementQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.AgreementStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.AgreementStatus)
}

// ExampleService_PayAfterUseCreditAgreementTransfer 展示 zhima.credit.payafteruse.creditagreement.transfer. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PayAfterUseCreditAgreementTransferRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PayAfterUseCreditAgreementTransfer() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PayAfterUseCreditAgreementTransferRequest
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
	response, err := svc.PayAfterUseCreditAgreementTransfer(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PayAfterUseCreditBizOrderFinish 展示 zhima.credit.payafteruse.creditbizorder.finish. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PayAfterUseCreditBizOrderFinishRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PayAfterUseCreditBizOrderFinish() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PayAfterUseCreditBizOrderFinishRequest
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
	response, err := svc.PayAfterUseCreditBizOrderFinish(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.OrderStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.OrderStatus)
}

// ExampleService_PayAfterUseCreditBizOrderOrder 展示 zhima.credit.payafteruse.creditbizorder.order. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PayAfterUseCreditBizOrderOrderRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PayAfterUseCreditBizOrderOrder() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PayAfterUseCreditBizOrderOrderRequest
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
	response, err := svc.PayAfterUseCreditBizOrderOrder(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoAgreementQuery 展示 zhima.credit.pe.zmgo.agreement.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoAgreementQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoAgreementQuery() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoAgreementQueryRequest
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
	response, err := svc.PeZmGoAgreementQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.AgreementStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.AgreementStatus)
}

// ExampleService_PeZmGoAgreementUnsign 展示 zhima.credit.pe.zmgo.agreement.unsign. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoAgreementUnsignRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoAgreementUnsign() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoAgreementUnsignRequest
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
	response, err := svc.PeZmGoAgreementUnsign(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoBizOptClose 展示 zhima.credit.pe.zmgo.bizopt.close. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoBizOptCloseRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoBizOptClose() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoBizOptCloseRequest
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
	response, err := svc.PeZmGoBizOptClose(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoCumulationSync 展示 zhima.credit.pe.zmgo.cumulation.sync. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoCumulationSyncRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoCumulationSync() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoCumulationSyncRequest
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
	response, err := svc.PeZmGoCumulationSync(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoPreOrderCreate 展示 zhima.credit.pe.zmgo.preorder.create. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoPreOrderCreateRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoPreOrderCreate() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoPreOrderCreateRequest
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
	response, err := svc.PeZmGoPreOrderCreate(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoSettleApply 展示 zhima.credit.pe.zmgo.settle.apply. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoSettleApplyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoSettleApply() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoSettleApplyRequest
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
	response, err := svc.PeZmGoSettleApply(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.SettleStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.SettleStatus)
}

// ExampleService_PeZmGoSettleRefund 展示 zhima.credit.pe.zmgo.settle.refund. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoSettleRefundRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoSettleRefund() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoSettleRefundRequest
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
	response, err := svc.PeZmGoSettleRefund(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_PeZmGoSettleUnfreeze 展示 zhima.credit.pe.zmgo.settle.unfreeze. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 credit.PeZmGoSettleUnfreezeRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_PeZmGoSettleUnfreeze() {
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

	svc, err := credit.NewService(client)
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
	var request *credit.PeZmGoSettleUnfreezeRequest
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
	response, err := svc.PeZmGoSettleUnfreeze(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.UnfreezeStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.UnfreezeStatus)
}
