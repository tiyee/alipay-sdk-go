package project_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/examples/requestfile"
	"github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules/project"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// ExampleService_InvoiceExpenseRulesProjectInfoModify 展示 alipay.ebpp.invoice.expenserules.projectinfo.modify 的平台证书接入流程。
// ALIPAY_EXAMPLE_REQUEST 指向按 project.InvoiceExpenseRulesProjectInfoModifyRequest 编写的 JSON 请求文件；
// 字段、条件必填与业务状态请参照本包 README 及方法 GoDoc 中的官方契约。
// 示例没有 Output 注释，go test 只编译，不读取配置或执行真实调用。
func ExampleService_InvoiceExpenseRulesProjectInfoModify() {
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

	svc, err := project.NewService(client)
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
	request, err := requestfile.Decode[project.InvoiceExpenseRulesProjectInfoModifyRequest](requestJSON)
	if err != nil {
		fmt.Println("请求必须为与请求类型匹配的 JSON 对象")
		return
	}
	var callOptions []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		callOptions = append(callOptions, core.WithAppAuthToken(token))
	}
	response, err := svc.InvoiceExpenseRulesProjectInfoModify(ctx, request, callOptions...)
	if err != nil {
		// 不直接输出原始错误或报文，以免泄露业务参数与令牌。
		fmt.Println("接口调用失败")
		return
	}
	// Client 已验证响应签名；response.Data 是此接口的具名响应类型。
	// HTTP 成功不代表业务终态，按接口契约读取 response.Data 的业务字段。
	if response.Data.Success == nil {
		fmt.Println("响应未包含业务状态或统计字段")
		return
	}
	fmt.Println(*response.Data.Success)
}
