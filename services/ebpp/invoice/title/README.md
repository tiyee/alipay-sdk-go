# ebpp/invoice/title 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/title`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `title.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_title_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceTitleDynamicGet`](#api-invoicetitledynamicget) | [`alipay.ebpp.invoice.title.dynamic.get`](https://opendocs.alipay.com/apis/05v8k3) | 根据动态码查询发票抬头 | `GET` |
| [`InvoiceTitleListGet`](#api-invoicetitlelistget) | [`alipay.ebpp.invoice.title.list.get`](https://opendocs.alipay.com/apis/06864n) | 获取指定用户所有的有效抬头列表 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicetitledynamicget"></a>

### InvoiceTitleDynamicGet — 根据动态码查询发票抬头

- 官方接口：[`alipay.ebpp.invoice.title.dynamic.get`](https://opendocs.alipay.com/apis/05v8k3)
- 路由：`GET /v3/alipay/ebpp/invoice/title/dynamic/get`
- 请求类型：[InvoiceTitleDynamicGetRequest](title_request.go#L5)
- 响应类型：[InvoiceTitleDynamicGetResponse](title_response.go#L5)

支付宝钱包内的发票管家app在用户抬头展示页面会显示对应该抬头的动态条形码，通过本接口，可根据条形码中数字串获取对应的用户抬头信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bar_code` | `BarCode` | `*string` | query | 否 | 抬头动态码 |

<a id="api-invoicetitlelistget"></a>

### InvoiceTitleListGet — 获取指定用户所有的有效抬头列表

- 官方接口：[`alipay.ebpp.invoice.title.list.get`](https://opendocs.alipay.com/apis/06864n)
- 路由：`GET /v3/alipay/ebpp/invoice/title/list/get`
- 请求类型：[InvoiceTitleListGetRequest](title_request.go#L11)
- 响应类型：[InvoiceTitleListGetResponse](title_response.go#L10)

用于服务商在用户授权的情况下获取用户的发票抬头列表供用户使用。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `user_id` | `UserId` | `*string` | query | 否 | 支付宝用户id |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝用户id |
