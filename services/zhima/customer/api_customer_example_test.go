package customer_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/examples/requestfile"
	"github.com/tiyee/alipay-sdk-go/internal/sandbox"
	customer "github.com/tiyee/alipay-sdk-go/services/zhima/customer"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_JobWorthCloudResumeQuery 展示 zhima.customer.jobworth.cloudresume.query 的调用和响应读取。
// 示例只演示调用流程，真实请求的必填条件与业务状态以方法 GoDoc 链接的官方契约为准。
func ExampleService_JobWorthCloudResumeQuery() {
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

	svc, err := customer.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}

	connKey := "example-token"
	onceToken := "example-token"
	response, err := svc.JobWorthCloudResumeQuery(ctx, customer.JobWorthCloudResumeQueryRequest{
		ConnKey:   &connKey,
		OnceToken: &onceToken,
	})
	if err != nil {
		fmt.Println("接口调用失败")
		return
	}
	if response.Data.IntentionStatus == nil {
		fmt.Println("响应字段缺失")
		return
	}
	fmt.Println(*response.Data.IntentionStatus)
}

// TestSandboxService_JobWorthCloudResumeQuery 使用沙箱证书与外部请求参数调用 zhima.customer.jobworth.cloudresume.query。
func TestSandboxService_JobWorthCloudResumeQuery(t *testing.T) {
	sandbox.Run(t, "zhima.customer.jobworth.cloudresume.query", customer.NewService, (*customer.Service).JobWorthCloudResumeQuery)
}

// ExampleService_JobWorthAuthenticationQuery 展示 zhima.customer.jobworth.authentication.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 customer.JobWorthAuthenticationQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_JobWorthAuthenticationQuery() {
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

	svc, err := customer.NewService(client)
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
	request, err := requestfile.Decode[customer.JobWorthAuthenticationQueryRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.JobWorthAuthenticationQuery(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.TokenStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.TokenStatus)
}

// ExampleService_JobWorthInfoQuery 展示 zhima.customer.jobworth.info.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 customer.JobWorthInfoQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_JobWorthInfoQuery() {
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

	svc, err := customer.NewService(client)
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
	request, err := requestfile.Decode[customer.JobWorthInfoQueryRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.JobWorthInfoQuery(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_JobWorthJobDataAdd 展示 zhima.customer.jobworth.jobdata.add. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 customer.JobWorthJobDataAddRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_JobWorthJobDataAdd() {
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

	svc, err := customer.NewService(client)
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
	request, err := requestfile.Decode[customer.JobWorthJobDataAddRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.JobWorthJobDataAdd(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_JobWorthPictureUpload 展示 zhima.customer.jobworth.picture.upload. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 customer.JobWorthPictureUploadRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_JobWorthPictureUpload() {
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

	svc, err := customer.NewService(client)
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
	request, err := requestfile.Decode[customer.JobWorthPictureUploadRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.JobWorthPictureUpload(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_JobWorthPositionAdd 展示 zhima.customer.jobworth.position.add. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 customer.JobWorthPositionAddRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_JobWorthPositionAdd() {
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

	svc, err := customer.NewService(client)
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
	request, err := requestfile.Decode[customer.JobWorthPositionAddRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.JobWorthPositionAdd(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}
