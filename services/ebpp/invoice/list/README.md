# ebpp/invoice/list 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/list`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `list.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_list_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceListExpenseSync`](#api-invoicelistexpensesync) | [`alipay.ebpp.invoice.list.expense.sync`](https://opendocs.alipay.com/apis/05jve4) | 报销状态批量同步开放接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicelistexpensesync"></a>

### InvoiceListExpenseSync — 报销状态批量同步开放接口

- 官方接口：[`alipay.ebpp.invoice.list.expense.sync`](https://opendocs.alipay.com/apis/05jve4)
- 路由：`POST /v3/alipay/ebpp/invoice/list/expense/sync`
- 请求类型：[InvoiceListExpenseSyncRequest](list_request.go#L5)
- 响应类型：[InvoiceListExpenseSyncResponse](list_response.go#L9)

报销状态节点状态批量同步时调用该接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `business_time` | `BusinessTime` | `*string` | JSON body | 未标注 | 报销进度变更时间 |
| `expense_order_no` | `ExpenseOrderNo` | `*string` | JSON body | 未标注 | 商户报销单据号 |
| `expense_tax_no` | `ExpenseTaxNo` | `*string` | JSON body | 未标注 | 报销企业税号 |
| `invoice_element_list` | `InvoiceElementList` | [`*[]InvoiceElementStatusSyncOpen`](#model-invoiceelementstatussyncopen) | JSON body | 未标注 | 需要同步报销状态的发票列表 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 报销审批备注 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户uid |
| `status` | `Status` | `*string` | JSON body | 未标注 | EXPENSE_APPLY:报销申请提交;EXPENSE_APPROVAL_PASS:报销申请审核通过;EXPENSE_FINISHED:报销申请完成;EXPENSE_CANCEL:报销申请撤回 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户uid |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-invoiceelementstatussyncopen"></a>

### InvoiceElementStatusSyncOpen

类型定义：[InvoiceElementStatusSyncOpen](list_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | 未标注 | 同步申请id，每次发起同步时生成，isv每次请求需要保证唯一 |
| `expense_detail_url` | `ExpenseDetailUrl` | `*string` | 未标注 | 报销详情地址，提供用户通过发票管家查看报销进度的地址<br> 如果报销企业入驻发票管家时需要isv传入报销详情地址，则必须提供 |
| `invoice_code` | `InvoiceCode` | `*string` | 未标注 | 发票代码 |
| `invoice_no` | `InvoiceNo` | `*string` | 未标注 | 发票号码 |
