package merchant_test

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
	"github.com/tiyee/alipay-sdk-go/services/ant/merchant"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_ExpandAppRecommendAccountQuery_withAuthorization 展示单次调用传入商户授权令牌。
func ExampleService_ExpandAppRecommendAccountQuery_withAuthorization() {
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

	svc, err := merchant.NewService(client)
	if err != nil {
		fmt.Println("创建 Service 失败")
		return
	}
	pageSize := int64(10)
	response, err := svc.ExpandAppRecommendAccountQuery(ctx,
		merchant.ExpandAppRecommendAccountQueryRequest{PageSize: &pageSize},
		core.WithAppAuthToken("example-app-auth-token"))
	if err != nil {
		fmt.Println("接口调用失败")
		return
	}
	fmt.Println(response.Result.HTTPStatus)
}

// TestSandboxService_ExpandAppRecommendAccountQuery_withAuthorization 使用沙箱证书与外部请求参数调用 ant.merchant.expand.apprecommend.account.query。
func TestSandboxService_ExpandAppRecommendAccountQuery_withAuthorization(t *testing.T) {
	sandbox.Run(t, "ant.merchant.expand.apprecommend.account.query", merchant.NewService, (*merchant.Service).ExpandAppRecommendAccountQuery)
}

// ExampleService_ExpandAppRecommendAccountCreate 展示 ant.merchant.expand.apprecommend.account.create. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandAppRecommendAccountCreateRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandAppRecommendAccountCreate() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandAppRecommendAccountCreateRequest
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
	response, err := svc.ExpandAppRecommendAccountCreate(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandAppRecommendAccountDelete 展示 ant.merchant.expand.apprecommend.account.delete. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandAppRecommendAccountDeleteRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandAppRecommendAccountDelete() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandAppRecommendAccountDeleteRequest
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
	response, err := svc.ExpandAppRecommendAccountDelete(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandAppRecommendAvailableQuery 展示 ant.merchant.expand.apprecommend.available.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandAppRecommendAvailableQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandAppRecommendAvailableQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandAppRecommendAvailableQueryRequest
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
	response, err := svc.ExpandAppRecommendAvailableQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandInfoQuery 展示 ant.merchant.expand.info.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandInfoQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandInfoQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandInfoQueryRequest
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
	response, err := svc.ExpandInfoQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemCreate 展示 ant.merchant.expand.item.create. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemCreateRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemCreate() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemCreateRequest
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
	response, err := svc.ExpandItemCreate(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemDelete 展示 ant.merchant.expand.item.delete. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemDeleteRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemDelete() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemDeleteRequest
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
	response, err := svc.ExpandItemDelete(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemModify 展示 ant.merchant.expand.item.modify. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemModifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemModify() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemModifyRequest
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
	response, err := svc.ExpandItemModify(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemOpenBatchQuery 展示 ant.merchant.expand.item.open.batchquery. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemOpenBatchQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemOpenBatchQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemOpenBatchQueryRequest
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
	response, err := svc.ExpandItemOpenBatchQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemOpenCreate 展示 ant.merchant.expand.item.open.create. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemOpenCreateRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemOpenCreate() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemOpenCreateRequest
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
	response, err := svc.ExpandItemOpenCreate(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemOpenDelete 展示 ant.merchant.expand.item.open.delete. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemOpenDeleteRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemOpenDelete() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemOpenDeleteRequest
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
	response, err := svc.ExpandItemOpenDelete(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemOpenModify 展示 ant.merchant.expand.item.open.modify. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemOpenModifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemOpenModify() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemOpenModifyRequest
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
	response, err := svc.ExpandItemOpenModify(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemOpenQuery 展示 ant.merchant.expand.item.open.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemOpenQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemOpenQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemOpenQueryRequest
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
	response, err := svc.ExpandItemOpenQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemQuery 展示 ant.merchant.expand.item.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemQueryRequest
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
	response, err := svc.ExpandItemQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandItemStatusModify 展示 ant.merchant.expand.item.status.modify. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandItemStatusModifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandItemStatusModify() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandItemStatusModifyRequest
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
	response, err := svc.ExpandItemStatusModify(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandShopConsult 展示 ant.merchant.expand.shop.consult. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopConsultRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopConsult() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopConsultRequest
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
	response, err := svc.ExpandShopConsult(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandShopReceiptAccountSave 展示 ant.merchant.expand.shop.receiptaccount.save. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopReceiptAccountSaveRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopReceiptAccountSave() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopReceiptAccountSaveRequest
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
	response, err := svc.ExpandShopReceiptAccountSave(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandIndirectImageUpload 展示 ant.merchant.expand.indirect.image.upload. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandIndirectImageUploadRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandIndirectImageUpload() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandIndirectImageUploadRequest
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
	response, err := svc.ExpandIndirectImageUpload(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandMCCQuery 展示 ant.merchant.expand.mcc.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandMCCQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandMCCQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandMCCQueryRequest
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
	response, err := svc.ExpandMCCQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandOrderQuery 展示 ant.merchant.expand.order.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandOrderQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandOrderQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandOrderQueryRequest
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
	response, err := svc.ExpandOrderQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.Status == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.Status)
}

// ExampleService_ExpandShopQuery 展示 ant.merchant.expand.shop.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopQueryRequest
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
	response, err := svc.ExpandShopQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.ShopInfoStatus == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.ShopInfoStatus)
}

// ExampleService_ExpandShopModify 展示 ant.merchant.expand.shop.modify. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopModifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopModify() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopModifyRequest
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
	response, err := svc.ExpandShopModify(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandShopCreate 展示 ant.merchant.expand.shop.create. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopCreateRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopCreate() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopCreateRequest
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
	response, err := svc.ExpandShopCreate(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandShopClose 展示 ant.merchant.expand.shop.close. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopCloseRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopClose() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopCloseRequest
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
	response, err := svc.ExpandShopClose(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}

// ExampleService_ExpandShopPageQuery 展示 ant.merchant.expand.shop.page.query. 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 merchant.ExpandShopPageQueryRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_ExpandShopPageQuery() {
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

	svc, err := merchant.NewService(client)
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
	var request *merchant.ExpandShopPageQueryRequest
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
	response, err := svc.ExpandShopPageQuery(ctx, *request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	fmt.Println(response.Result.HTTPStatus)
}
