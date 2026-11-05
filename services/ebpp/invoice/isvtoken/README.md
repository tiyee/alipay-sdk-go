# ebpp/invoice/isvtoken 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/isvtoken`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `isvtoken.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_isvtoken_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceISVTokenReimApply`](#api-invoiceisvtokenreimapply) | [`alipay.ebpp.invoice.isvtoken.reim.apply`](https://opendocs.alipay.com/apis/05xa7p) | 获取报销授权令牌 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceisvtokenreimapply"></a>

### InvoiceISVTokenReimApply — 获取报销授权令牌

- 官方接口：[`alipay.ebpp.invoice.isvtoken.reim.apply`](https://opendocs.alipay.com/apis/05xa7p)
- 路由：`POST /v3/alipay/ebpp/invoice/isvtoken/reim/apply`
- 请求类型：[InvoiceISVTokenReimApplyRequest](isvtoken_request.go#L5)
- 响应类型：[InvoiceISVTokenReimApplyResponse](isvtoken_response.go#L5)

当报销软件跳转支付宝报销助手，需要通过授权令牌进行身份认证，该接口用于ISV获取报销助手的授权令牌

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `isv_app_code` | `IsvAppCode` | `*string` | JSON body | 未标注 | 报销软件app对应的isv app code，由支付宝进行分配。详情参见 ["推"模式发票报销](https://opendocs.alipay.com/open/017fwh) |
