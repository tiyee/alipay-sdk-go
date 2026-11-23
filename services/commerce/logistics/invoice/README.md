# invoice 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/logistics/invoice`。本包提供 2 个支付宝 V3 REST 操作。

通过 `invoice.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`LogisticsInvoiceIstdWaybillCreate`](#api-logisticsinvoiceistdwaybillcreate) | `alipay.commerce.logistics.invoice.istdwaybill.create` |
| [`LogisticsInvoiceIstdWaybillQuery`](#api-logisticsinvoiceistdwaybillquery) | `alipay.commerce.logistics.invoice.istdwaybill.query` |

## 使用示例

- [LogisticsInvoiceIstdWaybillCreate](api_invoice_logisticsinvoiceistdwaybillcreate_example_test.go)
- [LogisticsInvoiceIstdWaybillQuery](api_invoice_logisticsinvoiceistdwaybillquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-logisticsinvoiceistdwaybillcreate"></a>

### LogisticsInvoiceIstdWaybillCreate — 开即时配送订单发票

- 官方接口：[`alipay.commerce.logistics.invoice.istdwaybill.create`](https://opendocs.alipay.com/pre-apis/068d50)
- 路由：`POST /v3/alipay/commerce/logistics/invoice/istdwaybill/create`
- 请求类型：[LogisticsInvoiceIstdWaybillCreateRequest](invoice_request.go#L4)
- 响应类型：[LogisticsInvoiceIstdWaybillCreateResponse](invoice_response.go#L4)

开即时配送订单发票

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice` | `Invoice` | [`*Invoice`](#model-invoice) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `logistics_code` | `LogisticsCode` | `*string` | JSON body | 未标注 | 即时配送公司编码 |
| `out_invoice_request_no` | `OutInvoiceRequestNo` | `*string` | JSON body | 未标注 | 开票请求流水号 |
| `waybill_invoices` | `WaybillInvoices` | [`*[]WaybillInvoice`](#model-waybillinvoice) | JSON body | 未标注 | 即时配送运单列表 |

<a id="api-logisticsinvoiceistdwaybillquery"></a>

### LogisticsInvoiceIstdWaybillQuery — 查询即时配送订单的开票结果

- 官方接口：[`alipay.commerce.logistics.invoice.istdwaybill.query`](https://opendocs.alipay.com/pre-apis/068lff)
- 路由：`GET /v3/alipay/commerce/logistics/invoice/istdwaybill/query`
- 请求类型：[LogisticsInvoiceIstdWaybillQueryRequest](invoice_request.go#L15)
- 响应类型：[LogisticsInvoiceIstdWaybillQueryResponse](invoice_response.go#L14)

查询即时配送订单的开票结果

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `logistics_code` | `LogisticsCode` | `*string` | query | 否 | 即时配送公司编码 |
| `out_invoice_request_no` | `OutInvoiceRequestNo` | `*string` | query | 否 | 开票请求流水号 |

## 嵌套请求参数

<a id="model-invoice"></a>

### Invoice

类型定义：[Invoice](invoice_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `email` | `Email` | `*string` | 未标注 | 邮箱地址 |
| `invoice_content` | `InvoiceContent` | `*string` | 未标注 | 发票内容，配送服务费 |
| `invoice_fee` | `InvoiceFee` | `*string` | 未标注 | 发票金额 |
| `invoice_title` | `InvoiceTitle` | `*string` | 未标注 | 发票抬头，如支付宝（杭州）信息技术有限公司 |
| `invoice_type` | `InvoiceType` | `*int64` | 未标注 | 发票类型，0：默认电子发票 |
| `phone` | `Phone` | `*string` | 未标注 | 手机号码 |
| `receiver` | `Receiver` | `*string` | 未标注 | 收件人姓名 |
| `tax_number` | `TaxNumber` | `*string` | 未标注 | 税号 |
| `title_type` | `TitleType` | `*int64` | 未标注 | 抬头类型，0：企业单位，1：个人/非企业 |

<a id="model-waybillinvoice"></a>

### WaybillInvoice

类型定义：[WaybillInvoice](invoice_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `waybill_amount` | `WaybillAmount` | `*string` | 未标注 | 即时配送运单金额 |
| `waybill_no` | `WaybillNo` | `*string` | 未标注 | 即时配送运单编号 |
