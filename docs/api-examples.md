# 接口使用示例

`services/` 各领域的 `api_*_example_test.go` 覆盖全部 746 个 REST 方法。每个 `ExampleService_<方法名>` 都独立展示以下流程，不依赖模拟执行器或内部测试凭证：

1. 读取并解析应用私钥、应用公钥证书、支付宝公钥证书及根证书。
2. 创建 `auth.CertificateCredentials`，校验证书材料。
3. 创建 `core.Client`，显式选择沙箱环境并设置应用 APPID。
4. 创建领域 Service，设置调用上下文、构造具体请求类型并调用接口。
5. 处理错误，读取类型化响应中的业务字段或 HTTP 状态。

示例默认使用支付宝开放平台的证书模式。支付宝公钥证书用于验签，应用私钥仍用于请求签名；无需配置 `ALIPAY_PUBLIC_KEY`。生产接入时，将环境改为 `core.EnvironmentProduction`，并同时更换 APPID、全部凭证及业务参数。Client 可在应用启动时创建并并发复用。

## 配置与请求参数

证书示例使用以下环境变量，值均来自同一沙箱应用：

```sh
export ALIPAY_APP_ID='your-sandbox-app-id'
export ALIPAY_APP_PRIVATE_KEY='/absolute/path/app-private-key.pem'
export ALIPAY_APP_CERT='/absolute/path/appPublicCert.crt'
export ALIPAY_CERT='/absolute/path/alipayPublicCert.crt'
export ALIPAY_ROOT_CERT='/absolute/path/alipayRootCert.crt'
```

`ALIPAY_APP_PRIVATE_KEY` 文件使用 PKCS#8 RSA 私钥格式；REST 示例的 `ALIPAY_CERT` 可包含多个 PEM 证书块，用于证书轮换。通知验签示例要求该文件只包含一张可信支付宝证书。

已有的请求构造示例保留具名 struct、指针字段和业务状态判断，例如[交易查询与预下单](../services/trade/README.md#使用示例)、[退款](../services/trade/refund/api_refund_example_test.go)与[转账查询](../services/fund/trans/api_trans_orderquery_example_test.go)。其中的业务参数仅作演示，接入时替换为自己的订单、授权信息和产品配置。

其余方法的完整示例通过 `ALIPAY_EXAMPLE_REQUEST` 指定请求 JSON 文件，解码为该方法的具体请求类型。示例不猜测产品权限、条件必填字段或枚举；按本领域 README 的请求表及方法 GoDoc 链接的官方契约填写参数。解码拒绝未知字段、重复键、类型错误和 `null` 请求。

例如，`trade.Service.Pay` 的请求文件可按自己的付款场景填写 `trade.PayRequest` 的字段；请求文件不包含公共参数或 `biz_content` 包装。JSON 字段名使用请求类型的 `json` 标签；只有 `query` 或 `path` 标签的字段使用 Go 字段名。混合请求中标记为 `json:"-"` 的字段使用 `query` 或 `path` 标签名。例如电子合同撤销请求可写成 `{"target_app_id":"app","flow_id":"flow"}`，其中 `target_app_id` 随 URL 发送。multipart 文件参数按 `core.MultipartFile` 的字段填写，其中 `Content` 使用 base64 编码。

```sh
export ALIPAY_EXAMPLE_REQUEST='/absolute/path/request.json'
```

这些示例也展示通过 `ALIPAY_APP_AUTH_TOKEN` 向单次调用传入商户授权令牌。是否需要令牌取决于具体产品及授权场景。凭证、令牌和请求文件保存在仓库外。

## 编译与真实调用

所有 `Example*` 均不含 `Output` 注释，`go test ./...` 只编译这些示例，不执行其函数体，也不读取证书或访问支付宝。`go test -run Example` 同样不会执行这些没有输出断言的示例。

要在自己的程序中使用，可复制目标示例函数体和它使用的公开包 import 到 `package main` 的 `main` 函数，配置证书并提供自己的业务参数。示例文件中的 `testing` 与 `internal/sandbox` 仅用于同文件已有的沙箱测试，不需要复制到应用中。

已有的 `TestSandboxService_*` 提供显式启用的真实沙箱测试入口，配置的是 `ALIPAY_SANDBOX_REQUEST`，与普通示例使用的 `ALIPAY_EXAMPLE_REQUEST` 分开。运行方式见[沙箱验收指引](sandbox-verification.md#通过各领域-api-示例测试调用沙箱)。沙箱支持范围与产品权限需单独确认，编译通过不代表真实接口已验收。

## 页面支付和通知

[页面与 App 支付示例](../services/trade/api_payment_example_test.go)分别展示电脑网站、手机网站的 URL/HTML 表单和 App 订单串构建。它们装载相同证书材料，通过 `trade.NewPaymentProducts` 创建产物构建器，仅在本地签名。调用方将产物交给浏览器或支付宝客户端 SDK，支付结果通过通知或查询确认。

[支付通知示例](../services/trade/api_notify_example_test.go)装载平台证书并创建 `notify.Handler`，从 `ALIPAY_NOTIFY_FORM` 指向的原始 POST 表单文件演示 `ParseRequest`。实际 HTTP 回调中直接传入收到的请求，不重新编码或伪造签名。通知验签之后仍需校验订单、金额、APPID 和收款方，以持久化事务保证幂等；示例未接入付款持久化，因此不会确认付款已入账。完整 HTTP 接入方式见[通知程序示例](../examples/tradenotify/main.go)。
