# alipay-sdk-go

支付宝开放平台的第三方 Go SDK，提供交易、退款、转账、账单、授权等接口，以及网页支付、App 支付订单串和支付通知验签能力。**本项目非支付宝官方 SDK。**

服务端调用仅支持支付宝 V3 REST；网页和 App 支付通过本地生成签名产物接入。

当前尚未发布首版，也未完成真实沙箱验收。部分接口根据官方接口快照生成，尚未逐项验证业务规则和权限要求。接入前请确认目标接口的产品权限、认证方式并完成实际业务验证。

## 特性

- **Client 与 Service 分离**：`core.Client` 负责通信、签名验签与错误处理，`services/` 按业务领域提供请求和响应类型。
- **类型化调用**：具名请求、响应和枚举，泛型 `Operation[Req, Resp]` 保持请求与响应配对，支持在自己的包中扩展接口。
- **两种认证方式**：平台公钥与应用证书模式，统一使用 RSA2；成功响应通过验签后才解码为业务数据。
- **多种支付入口**：V3 REST 调用、网页支付表单或链接、App 支付订单串，以及普通支付通知解析与验签。
- **可控的运行配置**：支持 context 取消与超时、自定义 HTTP Client、`log/slog` 日志；默认静默、不自动重试。
- **仅依赖标准库**：默认测试离线运行，无需真实凭证。

## 目录

- [安装](#安装)
- [签名与验签](#签名与验签)：[平台公钥模式](#平台公钥模式)、[应用证书模式](#应用证书模式)
- [选择接入场景](#选择接入场景)：[常用接口速查](#常用接口速查)、[网页和 App 支付](#网页和-app-支付)、[支付通知与订单状态](#支付通知与订单状态)
- [自行接入未封装的接口](#自定义请求)
- [文件上传](#文件上传)
- [常用配置](#常用配置)
- [错误处理](#错误处理)
- [示例与文档](#示例与文档)
- [支持范围与更多资料](#支持范围与更多资料)

## 安装

最低支持 **Go 1.24.0**，仅依赖标准库。在所有支持版本上都可调用领域 Service、`core.Execute`、`notify.Parse` 和 `notify.ParseRequest`。Go 1.27 起另提供对应的泛型方法。库在编译时选择工具链支持的实现，默认构建无需设置标签或实验开关。公开的原始 JSON 类型统一为 `encoding/json.RawMessage`。

```sh
go mod init example.com/alipay-demo  # 已有 Go module 时跳过
go get github.com/tiyee/alipay-sdk-go
```

按需导入核心包与业务包：

```go
import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)
```

本文中的 **V3 指支付宝接口协议**，模块路径为 `github.com/tiyee/alipay-sdk-go`，无需追加 `/v3`。项目尚未发布首版，公开 API 仍可能调整。

## 签名与验签

按应用在开放平台配置的方式选择平台公钥模式或应用证书模式。两种方式都使用应用私钥进行 RSA2 签名，由 SDK 自动完成请求签名和响应验签；它们不对业务内容进行加密。接口内容加密是独立能力，当前 SDK 尚未支持。

| 模式 | SDK 需要的材料 | 验签方式 |
| --- | --- | --- |
| 平台公钥模式 | 应用私钥、支付宝公钥 | 使用配置的支付宝公钥验证响应和支付通知 |
| 应用证书模式 | 应用私钥、应用公钥证书、支付宝公钥证书、支付宝根证书 | 使用本地可信支付宝证书中的公钥验签 |

应用私钥用于签名，支付宝公钥用于验签。应用公钥与支付宝公钥用途不同；应用证书模式仍需要与应用公钥证书匹配的私钥，不能仅提供证书。凭证材料需与 APPID、沙箱或生产环境一致。

### 平台公钥模式

下面以查询订单为例。请先在支付宝开放平台准备好：

- 应用 APPID，并开通目标接口所需的产品权限。
- 应用 RSA 私钥，保存为 PKCS#8 PEM 文件（`BEGIN PRIVATE KEY`）。
- 支付宝公钥，保存为 PKIX PEM 文件（`BEGIN PUBLIC KEY`）。
- 当前环境中已创建的商户订单号。

示例使用沙箱环境，APPID、密钥和订单也应来自同一沙箱应用。将以下代码保存为 `main.go`：

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

func main() {
	if err := queryOrder(); err != nil {
		var apiErr *core.APIError
		if errors.As(err, &apiErr) {
			// 仅输出错误码；未验签的错误只能用于诊断。
			fmt.Fprintln(os.Stderr, "支付宝接口错误：", apiErr.Code, "已验签：", apiErr.Verified)
		} else {
			fmt.Fprintln(os.Stderr, "查询失败，请检查应用配置、密钥文件和网络连接")
		}
		os.Exit(1)
	}
}

func queryOrder() error {
	appID := os.Getenv("ALIPAY_APP_ID")
	outTradeNo := os.Getenv("ALIPAY_OUT_TRADE_NO")
	if appID == "" || outTradeNo == "" {
		return fmt.Errorf("ALIPAY_APP_ID and ALIPAY_OUT_TRADE_NO are required")
	}
	credentials, err := loadCredentials()
	if err != nil {
		return err
	}
	client, err := core.NewClient(credentials,
		option.WithAppID(appID),
		option.WithEnvironment(core.EnvironmentSandbox),
	)
	if err != nil {
		return err
	}
	service, err := trade.NewService(client)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := service.Query(ctx, trade.QueryRequest{OutTradeNo: outTradeNo})
	if err != nil {
		return err
	}
	// 查询成功后，仍需按交易状态判断订单是否已付款。
	fmt.Println("交易状态：", resp.Data.TradeStatus)
	fmt.Println("订单金额：", resp.Data.TotalAmount)
	return nil
}

func loadCredentials() (auth.Credentials, error) {
	privateKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}
	publicKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_PUBLIC_KEY"))
	if err != nil {
		return nil, err
	}
	privateKey, err := utils.ParsePrivateKeyPKCS8(string(privateKeyPEM))
	if err != nil {
		return nil, err
	}
	alipayKey, err := utils.ParsePublicKeyPKIX(string(publicKeyPEM))
	if err != nil {
		return nil, err
	}
	return auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       privateKey,
		AlipayPublicKey: alipayKey,
	})
}
```

配置环境变量后运行（密钥变量填写文件路径）：

```sh
export ALIPAY_APP_ID='your-sandbox-app-id'
export ALIPAY_APP_PRIVATE_KEY='/path/to/app-private-key.pem'
export ALIPAY_PUBLIC_KEY='/path/to/alipay-public-key.pem'
export ALIPAY_OUT_TRADE_NO='your-order-no'
go run .
```

生产环境使用 `core.EnvironmentProduction`，并同时换成生产应用的凭证与订单。Client 可并发复用，建议在应用启动时创建；不同应用或环境分别创建实例。

### 应用证书模式

使用应用证书时，准备以下 PEM 文件。这里的证书是支付宝开放平台的应用及支付宝证书材料：

| 环境变量 | 文件内容 | 用途 |
| --- | --- | --- |
| `ALIPAY_APP_PRIVATE_KEY` | 应用 RSA 私钥（PKCS#8） | 对请求签名，必须与应用公钥证书匹配 |
| `ALIPAY_APP_CERT` | 应用公钥证书 | 标识应用证书身份 |
| `ALIPAY_CERT` | 平台验签叶证书，可附带签发它的中间 CA | 对响应和支付通知验签 |
| `ALIPAY_ROOT_CERT` | 支付宝根证书文件，可能包含多张证书 | 校验支付宝证书并生成根证书标识 |

将上例的 `loadCredentials` 函数替换为下面的实现，保留其余初始化和订单查询代码，即可改为证书模式；无需再配置 `ALIPAY_PUBLIC_KEY`。

```go
func loadCredentials() (auth.Credentials, error) {
	privateKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}
	privateKey, err := utils.ParsePrivateKeyPKCS8(string(privateKeyPEM))
	if err != nil {
		return nil, err
	}
	appCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_CERT"))
	if err != nil {
		return nil, err
	}
	appCert, err := utils.ParseCertificatePEM(string(appCertPEM))
	if err != nil {
		return nil, err
	}
	alipayCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_CERT"))
	if err != nil {
		return nil, err
	}
	alipayCerts, err := utils.ParseCertificatesPEM(string(alipayCertPEM))
	if err != nil {
		return nil, err
	}
	rootCertPEM, err := os.ReadFile(os.Getenv("ALIPAY_ROOT_CERT"))
	if err != nil {
		return nil, err
	}
	rootCerts, err := utils.ParseRootCertificatesPEM(string(rootCertPEM))
	if err != nil {
		return nil, err
	}
	return auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner:          privateKey,
		AppCertificate:     appCert,
		AlipayCertificates: alipayCerts,
		RootCertificates:   rootCerts,
	})
}
```

配置当前沙箱应用的材料和订单后运行：

```sh
export ALIPAY_APP_ID='your-sandbox-app-id'
export ALIPAY_APP_PRIVATE_KEY='/path/to/app-private-key.pem'
export ALIPAY_APP_CERT='/path/to/app-cert.pem'
export ALIPAY_CERT='/path/to/alipay-cert.pem'
export ALIPAY_ROOT_CERT='/path/to/alipay-root-cert.pem'
export ALIPAY_OUT_TRADE_NO='your-order-no'
go run .
```

SDK 在构造凭证时校验证书有效期、私钥与应用证书的匹配关系及支付宝证书的可信关系，并计算所需证书标识。构造失败时请先修正材料，再创建 Client。

证书凭证也可传给 `trade.NewPaymentProducts` 和 `notify.NewHandler`，网页/App 支付和通知处理沿用对应场景的调用方式。普通支付通知不携带证书标识；通知验签当前要求凭证中只有一张可信支付宝证书，不能直接使用配置了多张支付宝证书的轮换凭证。

SDK 不自动下载或更新证书；更新本地证书后需要重建凭证及使用它的实例。轮换期间，上面的 `ALIPAY_CERT` 文件可包含新旧支付宝证书的多个 PEM 块。需要从多个文件装载证书的完整示例见[证书模式示例](examples/certmode/main.go)。遇到 `*core.CertificateError` 时，应检查本地是否缺少对应支付宝证书。

各领域的 `api_*_example_test.go` 覆盖全部 746 个 REST 方法，每个示例都包含平台证书装载、Client 与 Service 初始化、请求及响应处理。其中已有的 `TestSandboxService_*` 可加载沙箱平台证书与 JSON 请求文件执行真实调用。默认测试只编译使用示例、跳过沙箱调用，保持离线。

## 选择接入场景

| 需求 | 入口 | 使用说明 |
| --- | --- | --- |
| 查询、预下单、支付、退款、关单 | `trade.NewService(client)` | [交易接口](services/trade/README.md) |
| 电脑网站、手机网站、App 支付 | `trade.NewPaymentProducts(credentials, opts...)` | [支付产物示例](examples/pagepay/main.go) |
| 接收支付通知 | `notify.NewHandler(credentials, notify.PolicyTradeStatusSync)` | [通知接入示例](examples/tradenotify/main.go) |
| 转账、资金授权等 | `services/fund` 及其子包 | [资金接口](services/fund/README.md) |
| 账单下载地址、余额与账务明细 | `services/data` | [账单与数据接口](services/data/README.md) |
| 用户授权、应用授权 | `services/system`、`services/open` | [授权示例](examples/oauth/main.go) |
| 查找其他接口 | 按业务领域导入 `services/` 下的包 | [业务包目录](services/) |
| 调用尚未封装的 V3 接口 | 自定义请求、响应与 `core.Operation` | [自定义接口示例](examples/customapi/customapi.go) |

各接口方法的 GoDoc 提供官方文档链接。请求字段、必填条件、产品权限和业务结果请结合目标接口文档确认。

### 常用接口速查

下表列出常用官方接口与 SDK 入口；Service 方法通过 `NewService(client)` 创建的实例调用。

| 场景 | 官方接口名 | SDK 入口 |
| --- | --- | --- |
| 交易查询 | `alipay.trade.query` | `trade.Service.Query` |
| 扫码支付预下单 | `alipay.trade.precreate` | `trade.Service.PreCreate` |
| 交易支付 / 创建 / 撤销 | `alipay.trade.pay` / `alipay.trade.create` / `alipay.trade.cancel` | `trade.Service.Pay` / `Create` / `Cancel` |
| 交易关闭 | `alipay.trade.close` | `trade.Service.Close` |
| 退款 / 退款查询 | `alipay.trade.refund` / `alipay.trade.fastpay.refund.query` | `refund.Service.Refund` / `QueryRefund` |
| 电脑网站支付 | `alipay.trade.page.pay` | `trade.PaymentProducts.PagePayForm` / `PagePayURL` |
| 手机网站支付 | `alipay.trade.wap.pay` | `trade.PaymentProducts.WapPayForm` / `WapPayURL` |
| App 支付 | `alipay.trade.app.pay` | `trade.PaymentProducts.AppPayOrderString` |
| 单笔转账 / 转账查询 | `alipay.fund.trans.uni.transfer` / `alipay.fund.trans.common.query` | `trans.Service.UniTransfer` / `CommonQuery` |
| 对账单下载地址查询 | `alipay.data.dataservice.bill.downloadurl.query` | `bill.Service.BillDownloadURLQuery` |
| 用户授权令牌换取 / 刷新 | `alipay.system.oauth.token` | `system.Service.OAuthToken` |
| 应用授权令牌换取 | `alipay.open.auth.token.app` | `open.Service.AuthTokenApp` |
| 文件上传 | `alipay.open.file.upload` | `file.Service.FileUpload` |

资金转账导入 `services/fund/trans`，账单查询导入 `services/data/bill`；各子包通过 `NewService(client)` 复用同一 Client。请求字段、嵌套模型及接口对应方法见各领域文档：[交易](services/trade/README.md)、[资金](services/fund/README.md)、[账单](services/data/README.md)、[用户授权](services/system/README.md)、[应用授权](services/open/README.md)。更多业务领域和子包见[业务包目录](services/)。

### 网页和 App 支付

使用与 Client 相同的凭证、APPID 和环境创建 `trade.PaymentProducts`，按场景选择：

- 电脑网站支付：`PagePayForm` 生成 HTML 表单，或 `PagePayURL` 生成跳转链接。
- 手机网站支付：`WapPayForm` 或 `WapPayURL`。
- App 支付：`AppPayOrderString` 生成订单串，交给支付宝客户端 SDK 唤起支付。

这些方法只在本地构建签名产物。通知地址可通过 `core.WithNotifyURL` 传入产物构建方法；返回地址按请求的 `ReturnURL` 字段设置。完整参数与错误处理见[支付产物示例](examples/pagepay/main.go)。

例如，使用上文已装载的凭证生成手机网站支付链接。以下函数需导入 `context`、`core`、`core/auth`、`core/option` 和 `services/trade`：

```go
func buildWapPayURL(ctx context.Context, credentials auth.Credentials,
	appID, outTradeNo string) (string, error) {
	products, err := trade.NewPaymentProducts(credentials,
		option.WithAppID(appID),
		option.WithEnvironment(core.EnvironmentSandbox),
	)
	if err != nil {
		return "", err
	}
	returnURL := "https://merchant.example.com/pay/return"
	return products.WapPayURL(ctx, trade.PagePayRequest{
		OutTradeNo:  outTradeNo, // 从商户订单系统取得唯一订单号。
		TotalAmount: "0.01",
		Subject:     "示例商品",
		ProductCode: "QUICK_WAP_WAY",
		ReturnURL:   &returnURL,
	}, core.WithNotifyURL("https://merchant.example.com/pay/notify"))
}
```

将生成的链接用于浏览器跳转；通知地址和返回地址请替换为商户自己的地址。电脑网站支付使用 `PagePayForm`，产品码为 `FAST_INSTANT_TRADE_PAY`；App 支付使用 `AppPayRequest` 与 `AppPayOrderString`，产品码为 `QUICK_MSECURITY_PAY`。

支付页跳转和客户端返回后，应通过已验签的异步通知或订单查询确认支付结果。

### 支付通知与订单状态

使用 `core/notify` 的 Handler 接收 POST 表单，再通过 `trade.NotificationDecoder()` 解析为交易通知。完整 HTTP 处理器见[通知接入示例](examples/tradenotify/main.go)。

`notify.ParseRequest(r.Context(), handler, r, trade.NotificationDecoder())` 会限制请求体大小、解析一次表单并验签，成功后从返回值的 `Data` 读取类型化通知。若 Web 框架已解析 POST 表单，可用 `notify.Parse(r.Context(), handler, r.PostForm, trade.NotificationDecoder())`；不要再次 URL 解码，也不要混入 URL 查询参数。

通知示例启动时要求设置 `ALIPAY_SELLER_ID`。多商户应用应按订单读取预期收款方。

收到通知后，商户还需要：

1. 校验应用、收款方、商户订单号和订单金额。
2. 根据交易状态更新订单，并保证重复通知不会重复处理。
3. 在业务处理成功或确认重复通知后，返回 `success`；验签或处理失败时不应确认成功。

查询成功仅代表查询请求完成。`trade.TradeStatusSuccess` 和 `trade.TradeStatusFinished` 表示已付款；等待付款、关闭等状态需要分别处理。通知中的 `Paid()` 可判断付款状态，仍需完成上述商户校验。未收到通知时，使用 `service.Query` 确认订单状态。

金额使用十进制字符串，例如 `"12.34"`，避免浮点数计算。退款请求成功后也应确认退款状态，接口与使用示例见[退款接口](services/trade/refund/README.md)。

## 自定义请求

### 自行接入未封装的接口

SDK 未封装的 V3 REST 接口，可以在自己的项目中实现，无需修改 SDK、注册接口表或导入 `internal/`。自定义接口与内置 Service 共用同一个 Client，复用凭证、请求签名、响应验签、HTTP、日志与错误处理。

接入步骤：

1. 查阅目标接口的官方 V3 文档，确认接口名、HTTP 方法、`/v3/...` 路由、参数位置及认证要求。不能将接口名的点号直接替换成斜杠来推导路由。
2. 在自己的包中定义请求和响应 struct，按官方字段设置标签；金额使用字符串，可选字段按需使用指针，枚举使用具名类型。
3. 用 `core.NewJSONOperation[Req, Resp]` 声明接口契约，并在调用前完成必填、条件必填与业务参数校验。
4. 将请求放入 `core.Request[Req].Body`，通过已初始化的 `client.Execute` 调用；先检查错误，再读取 `resp.Data`。

下面沿用已核实的 `alipay.trade.query` 契约演示完整封装方式（[官方 V3 文档](https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query)）。该接口已有 `trade.Service.Query`；这里用于说明扩展机制，请按目标接口的真实契约替换模型与 `OperationSpec`。示例只保留按商户订单号查询所需的请求字段和部分响应字段：

```go
package customapi

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// OrderQueryRequest 是按商户订单号查询交易的请求。
type OrderQueryRequest struct {
	OutTradeNo string `json:"out_trade_no"`
}

// OrderQueryResponse 保留本例需要的交易查询响应字段。
type OrderQueryResponse struct {
	TradeNo     string            `json:"trade_no"`
	OutTradeNo  string            `json:"out_trade_no"`
	TradeStatus trade.TradeStatus `json:"trade_status"`
	TotalAmount string            `json:"total_amount"`
}

// QueryOrder 调用自定义 Operation；client 沿用应用初始化时创建的实例。
func QueryOrder(ctx context.Context, client *core.Client,
	req OrderQueryRequest, opts ...core.CallOption) (*core.Response[OrderQueryResponse], error) {
	if req.OutTradeNo == "" {
		return nil, &core.ValidationError{
			Field: "out_trade_no", Reason: "out_trade_no is required",
		}
	}
	op, err := core.NewJSONOperation[OrderQueryRequest, OrderQueryResponse](core.OperationSpec{
		Method:          "alipay.trade.query",
		HTTPMethod:      core.HTTPMethodPost,
		Path:            "/v3/alipay/trade/query",
		ParameterMode:   core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, client, op, core.Request[OrderQueryRequest]{
		Body:    req,
		Options: core.NewCallSettings(opts...),
	})
}
```

将上述代码放入自己项目的 `customapi` 包，在应用中导入该包后调用。下面的函数接收前文已创建的 Client 和商户订单号，返回交易状态（需导入 `context`、`core`、`services/trade` 及自己的 `customapi` 包）：

```go
func queryOrderStatus(ctx context.Context, client *core.Client,
	outTradeNo string) (trade.TradeStatus, error) {
	resp, err := customapi.QueryOrder(ctx, client, customapi.OrderQueryRequest{
		OutTradeNo: outTradeNo,
	})
	if err != nil {
		return "", err // 沿用本文“错误处理”中的错误分类。
	}
	return resp.Data.TradeStatus, nil // 查询成功后仍需判断是否已付款。
}
```

需要代调用时，向 `QueryOrder` 追加 `core.WithAppAuthToken(appAuthToken)`；选项仅作用于本次调用，不修改共享 Client。REST 通知地址按目标接口的请求字段设置。

### 参数编码与扩展边界

`NewJSONOperation` 的默认响应解码器为 `core.JSONDecoder[Resp]()`，请求编码由 `ParameterMode` 决定：

| 官方参数位置 / 格式 | `ParameterMode` | 请求模型写法 |
| --- | --- | --- |
| JSON 请求体 | `core.ParameterModeJSONBody` | 使用 `json:"field_name"` 标签 |
| URL 查询参数（GET 必须使用） | `core.ParameterModeQuery` | 使用 `query:"field_name"`，可选字段可用指针或 `query:"field_name,omitempty"` |
| Query 与 JSON 正文 | `core.ParameterModeJSONBodyWithQuery` | 纯查询字段使用 `query:"auth_token" json:"-"`，正文字段使用 JSON 标签 |
| multipart 文件上传 | `core.ParameterModeMultipart` | 文件用 `*core.MultipartFile` 与 `multipart:"field_name,file"`，业务 JSON 用 `multipart:"field_name,json"`；见[文件上传](#文件上传) |

HTTP 方法与参数位置均以官方契约为准，接口名包含 `query` 不代表一定使用 GET。路径参数通过 `path:"name"` 字段对应路由中的 `{name}`；JSON 请求中的纯路径字段同时标记 `json:"-"`。

`AuthRequirement` 必须显式声明：允许公钥或证书模式时用 `core.AuthRequirementAny`，要求证书模式时用 `core.AuthRequirementCertificate`；凭证不满足要求会在发送请求前失败。Operation 不会自动执行自定义模型的业务校验，应像上例一样显式校验，或在自定义编码器中实现。

需要特殊字段映射或编解码时，用 `core.NewOperation` 传入自己实现的 `core.Encoder[Req]` 和 `core.Decoder[Resp]`；编码器只输出业务参数，签名和公共身份参数由 Client 处理，解码器只接收已通过验签的业务响应。需要独立 Service 或测试替身时，注入 `core.Executor`，用 `core.Execute(ctx, executor, op, req)` 复用同一执行流程。

扩展入口仅支持 SDK 已实现的 V3 REST 传输和认证能力，不能用于传统网关服务端调用、尚未确认 V3 契约的接口或内容加密。网页/App 支付及通知沿用各自入口。完整的 GET、自定义编解码器与独立 Service 示例见[可编译示例](examples/customapi/customapi.go)。

## 文件上传

`services/open/file` 提供 `alipay.open.file.upload` 的 V3 multipart 调用入口。以下函数需导入 `context`、`os`、`core`，并用别名 `alipayfile` 导入 `github.com/tiyee/alipay-sdk-go/services/open/file`：

```go
func uploadImage(ctx context.Context, client *core.Client,
	path string) (*core.Response[alipayfile.FileUploadResponse], error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	service, err := alipayfile.NewService(client)
	if err != nil {
		return nil, err
	}
	bizCode := "content_creation"
	return service.FileUpload(ctx, alipayfile.FileUploadRequest{
		Data: &alipayfile.FileUpload{BizCode: &bizCode},
		FileContent: &core.MultipartFile{
			Filename:    "image.jpg",
			ContentType: "image/jpeg",
			Content:     content,
		},
	})
}
```

示例为内容发布场景的 JPEG 图片上传；其他场景按目标产品要求设置素材用途。文件内容会完整读入内存；调用前请按业务限制文件大小。该生成接口尚未完成逐项业务及沙箱验证，产品权限、素材格式与大小限制见[文件接口说明](services/open/file/README.md)。

## 常用配置

在 `core.NewClient(credentials, opts...)` 中传入所需选项：

| 选项 | 用途 |
| --- | --- |
| `option.WithAppID(appID)` | 设置应用 APPID |
| `option.WithEnvironment(core.EnvironmentSandbox)` | 选择沙箱；生产使用 `core.EnvironmentProduction` |
| `option.WithRequestTimeout(10*time.Second)` | 设置单次请求超时，默认 30 秒；也可用调用的 context 设置截止时间 |
| `option.WithHTTPClient(httpClient)` | 使用自定义 `*http.Client` 配置代理、连接池等 |
| `option.WithMaxResponseBytes(n)` | 限制响应体大小，默认 8 MiB |
| `option.WithLogger(logger)` | 启用 `log/slog` 日志；也可用 `WithLogHandler` 注入 Handler |

SDK 默认静默，启用日志后仅记录操作、耗时、状态码和错误分类等筛选字段。日志用法见[日志示例](examples/logging/main.go)。

例如，在初始化处同时注入 HTTP Client 与 JSON 日志（需额外导入 `net/http`、`log/slog`）：

```go
httpClient := &http.Client{Timeout: 15 * time.Second}
logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
client, err := core.NewClient(credentials,
	option.WithAppID(appID),
	option.WithEnvironment(core.EnvironmentSandbox),
	option.WithHTTPClient(httpClient),
	option.WithRequestTimeout(10*time.Second),
	option.WithLogger(logger),
)
```

按上文方式检查 `err` 后再使用 Client。需要链路跟踪时，可将接入方的跟踪 Transport 配置到自定义 `http.Client`，SDK 会传递调用 context，并对签名请求禁用自动重定向。

代商户调用时，将 `core.WithAppAuthToken(token)` 作为本次 Service 调用的额外参数传入，无需修改共享 Client：

```go
resp, err := service.Query(ctx,
	trade.QueryRequest{OutTradeNo: outTradeNo},
	core.WithAppAuthToken(appAuthToken),
)
```

`appAuthToken` 应来自应用授权令牌换取结果，并安全保存。REST 调用的通知地址应按目标接口的请求字段设置，不能使用 `core.WithNotifyURL`。

## 错误处理

先检查 `err`，再读取响应。错误支持标准错误链，用 `errors.As` 匹配错误类型（Go 1.26 起也可用 `errors.AsType[T]`），用 `errors.Is` 判断取消或超时。

| 错误类型 | 建议处理方式 |
| --- | --- |
| `*core.ConfigError` | 检查初始化配置 |
| `*core.ValidationError` | 修正请求参数；请求尚未发送 |
| `*core.TransportError` | 检查网络、取消或超时原因；写操作需查询确认是否已生效 |
| `*core.HTTPError` | 检查返回的 HTTP 状态 |
| `*core.ProtocolError` | 检查响应格式或大小限制 |
| `*core.SignatureError` | 签名或验签失败，不应信任响应或通知内容 |
| `*core.CertificateError` | 检查并更新证书配置 |
| `*core.APIError` | 按 `Code`、`Msg` 和 `Links` 排查；`Verified=false` 的错误仅供诊断 |

SDK 默认不自动重试。支付、退款、转账等写操作超时后可能已经生效，应先查询结果；确需重试时，按接口要求复用稳定的业务请求号并保证幂等。记录错误时请筛选必要字段，避免输出令牌、付款码、密钥或完整报文。

## 示例与文档

| 示例 | 内容 | 运行方式（在仓库根目录） |
| --- | --- | --- |
| [网页和 App 支付](examples/pagepay/main.go) | 生成电脑网站表单、手机网站链接和 App 订单串 | `go run ./examples/pagepay` |
| [证书模式](examples/certmode/main.go) | 装载应用证书、根证书和多张支付宝证书 | `go run ./examples/certmode` |
| [支付通知](examples/tradenotify/main.go) | HTTP 通知验签、商户校验与内存幂等演示 | `go run ./examples/tradenotify` |
| [授权](examples/oauth/main.go) | 用户授权令牌、刷新令牌及应用授权令牌换取 | `go run ./examples/oauth` |
| [日志](examples/logging/main.go) | 标准 slog、自定义 Handler 和日志故障处理 | `go run ./examples/logging`（离线） |
| [自定义接口](examples/customapi/customapi.go) | 自定义模型、编解码器与独立 Service | `go test ./examples/customapi`（离线） |

各示例文件头部列出所需配置。`pagepay`、`certmode` 和 `oauth` 的密钥变量读取文件路径；`tradenotify` 读取 PEM 文本内容，不能直接复用上文的文件路径变量。通知示例使用内存订单演示，实际接入需替换为持久化订单与幂等处理。

默认离线检查：

```sh
go test ./...
go vet ./...
```

真实沙箱调用需显式启用；不要将生产凭证或资金操作用于默认测试。

## 支持范围与更多资料

- 支持公钥和本地证书两种认证方式，使用 RSA2 签名。
- 通知解析目前仅支持普通支付通知；其他通知类型需要单独确认契约。
- 不提供传统网关服务端调用，不支持内容加密、GBK、XML 或 RSA(v1) 签名。
- [版本变更](CHANGELOG.md)：查看版本记录。
