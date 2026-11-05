# ebpp/invoice/detail 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/detail`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `detail.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_detail_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceDetailOutputQuery`](#api-invoicedetailoutputquery) | [`alipay.ebpp.invoice.detail.output.query`](https://opendocs.alipay.com/apis/05jxvp) | 获取发票详细信息 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicedetailoutputquery"></a>

### InvoiceDetailOutputQuery — 获取发票详细信息

- 官方接口：[`alipay.ebpp.invoice.detail.output.query`](https://opendocs.alipay.com/apis/05jxvp)
- 路由：`POST /v3/alipay/ebpp/invoice/detail/output/query`
- 请求类型：[InvoiceDetailOutputQueryRequest](detail_request.go#L5)
- 响应类型：[InvoiceDetailOutputQueryResponse](detail_response.go#L5)

根据发票代码，发票号码获取发票详细信息
查询权限要求：
发票属于当前isv报销单据进行状态（支持的状态有
EXPENSE_APPLY－用户已提交申请
EXPENSE_APPROVAL_PASS －报销审核通过
EXPENSE_FINISHED－报销完结）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice_code` | `InvoiceCode` | `*string` | JSON body | 未标注 | 发票代码 长度限制（10-12位），全电票则为空 |
| `invoice_no` | `InvoiceNo` | `*string` | JSON body | 未标注 | 发票号码 长度限制（8-20位） |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户 id |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 获取发票明细应用场景。固定为 INVOICE_EXPENSE 表示发票报销。 |
| `skip_expense_progress_sync` | `SkipExpenseProgressSync` | `*bool` | JSON body | 未标注 | 是否跳过发票报销状态同步；当为true时，跳过报销状态同步校验。默认为false，需要先做报销状态同步 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
