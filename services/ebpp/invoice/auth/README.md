# ebpp/invoice/auth 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/auth`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `auth.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_auth_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceAuthSign`](#api-invoiceauthsign) | [`alipay.ebpp.invoice.auth.sign`](https://opendocs.alipay.com/pre-apis/05jtxa) | 发票授权关系签约 | `POST` |
| [`InvoiceAuthUnsign`](#api-invoiceauthunsign) | [`alipay.ebpp.invoice.auth.unsign`](https://opendocs.alipay.com/apis/05jwd5) | 发票授权关系解约 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceauthsign"></a>

### InvoiceAuthSign — 发票授权关系签约

- 官方接口：[`alipay.ebpp.invoice.auth.sign`](https://opendocs.alipay.com/pre-apis/05jtxa)
- 路由：`POST /v3/alipay/ebpp/invoice/auth/sign`
- 请求类型：[InvoiceAuthSignRequest](auth_request.go#L5)
- 响应类型：[InvoiceAuthSignResponse](auth_response.go#L9)

用户在ISV端签约了发票相关的授权关系之后，ISV使用该接口将授权关系同步到发票管家。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `authorization_type` | `AuthorizationType` | `*string` | JSON body | 未标注 | 发票授权类型，可选值：INVOICE_AUTO_SYNC（发票自动回传） INVOICE_TWOWAY_AUTO_SYNC（发票双向自动回传） |
| `extend_fields` | `ExtendFields` | `*string` | JSON body | 未标注 | 扩展字段，格式为：KEY1=VALUE1,KEY2=VALUE2,KEY3=VALUE3 邮箱地址KEY为EMAIL_ADDRESS |
| `m_short_name` | `MShortName` | `*string` | JSON body | 未标注 | 开票商户品牌简称，与商户入驻时的品牌简称保持一致。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户userId |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 支付宝用户userId |

<a id="api-invoiceauthunsign"></a>

### InvoiceAuthUnsign — 发票授权关系解约

- 官方接口：[`alipay.ebpp.invoice.auth.unsign`](https://opendocs.alipay.com/apis/05jwd5)
- 路由：`POST /v3/alipay/ebpp/invoice/auth/unsign`
- 请求类型：[InvoiceAuthUnsignRequest](auth_request.go#L19)
- 响应类型：[InvoiceAuthUnsignResponse](auth_response.go#L12)

用户在ISV端解约了发票相关的授权关系之后，ISV使用该接口将授权关系在发票管家端解约。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `authorization_type` | `AuthorizationType` | `*string` | JSON body | 未标注 | 发票授权类型，可选值：INVOICE_AUTO_SYNC（发票自动回传） INVOICE_TWOWAY_AUTO_SYNC（发票双向自动回传） |
| `extend_fields` | `ExtendFields` | `*string` | JSON body | 未标注 | 扩展字段，格式为：KEY1=VALUE1,KEY2=VALUE2,KEY3=VALUE3 邮箱地址KEY为EMAIL_ADDRESS |
| `m_short_name` | `MShortName` | `*string` | JSON body | 未标注 | 开票商户品牌简称，与商户入驻时的品牌简称保持一致。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户userId |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 支付宝用户userId |
