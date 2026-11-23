# invoice 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/boss/fnc/invoice`。本包提供 3 个支付宝 V3 REST 操作。

通过 `invoice.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`FncInvoiceApply`](#api-fncinvoiceapply) | `alipay.boss.fnc.invoice.apply` |
| [`FncInvoiceBatchQuery`](#api-fncinvoicebatchquery) | `alipay.boss.fnc.invoice.batchquery` |
| [`FncInvoiceQuery`](#api-fncinvoicequery) | `alipay.boss.fnc.invoice.query` |

## 使用示例

- [FncInvoiceApply](api_invoice_fncinvoiceapply_example_test.go)
- [FncInvoiceBatchQuery](api_invoice_fncinvoicebatchquery_example_test.go)
- [FncInvoiceQuery](api_invoice_fncinvoicequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-fncinvoiceapply"></a>

### FncInvoiceApply — 应收开票申请(根据账单申请)

- 官方接口：[`alipay.boss.fnc.invoice.apply`](https://opendocs.alipay.com/pre-apis/05xo5f)
- 路由：`POST /v3/alipay/boss/fnc/invoice/apply`
- 请求类型：[BossFncInvoiceApplyRequest](invoice_request.go#L8)
- 响应类型：[BossFncInvoiceApplyResponse](invoice_response.go#L8)

根据账单申请应收发票开票

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_id` | `BizId` | `*string` | JSON body | 未标注 | 调用方身份识别标记(应用名称) 与业务单据号联合唯一 |
| `biz_no` | `BizNo` | `*string` | JSON body | 未标注 | 业务单据号，业务流水号 |
| `invoice_amt` | `InvoiceAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `invoice_type` | `InvoiceType` | `*string` | JSON body | 未标注 | 申请开票类型，01：专票；02：普票 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注 |
| `mthtly_bill_nos` | `MthtlyBillNos` | `*[]string` | JSON body | 未标注 | 需开票的应收月账单号 |
| `operator` | `Operator` | `*string` | JSON body | 未标注 | 申请人的uid号 |
| `out_biz_type` | `OutBizType` | `*string` | JSON body | 未标注 | 01主站，02芝麻，03金融云 |

<a id="api-fncinvoicebatchquery"></a>

### FncInvoiceBatchQuery — 应收发票综合查询

- 官方接口：[`alipay.boss.fnc.invoice.batchquery`](https://opendocs.alipay.com/pre-apis/05xdyx)
- 路由：`POST /v3/alipay/boss/fnc/invoice/batchquery`
- 请求类型：[BossFncInvoiceBatchQueryRequest](invoice_request.go#L27)
- 响应类型：[BossFncInvoiceBatchQueryResponse](invoice_response.go#L14)

根据综合查询条件查询对应发票信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `apply_order_id` | `ApplyOrderId` | `*string` | JSON body | 未标注 | 开票申请ID，唯一标识开票申请表的ID |
| `buyer_invoice_title` | `BuyerInvoiceTitle` | `*string` | JSON body | 未标注 | 买方发票抬头，用于打印在发票上 |
| `buyer_ipid` | `BuyerIpid` | `*string` | JSON body | 未标注 | 对应发票商户的ipId，在主站标示mid，在其他环境标示ip_id。 |
| `buyer_ipids` | `BuyerIpids` | `*[]string` | JSON body | 未标注 | 购方商户ID集合 |
| `buyer_iprole_id` | `BuyerIproleId` | `*string` | JSON body | 未标注 | 商户结算的ipRoleId，表示商户的结算对象，在主站表示pid。 |
| `buyer_iprole_ids` | `BuyerIproleIds` | `*[]string` | JSON body | 未标注 | 购方PID集合 |
| `inst_id` | `InstId` | `*string` | JSON body | 未标注 | 对应发票所在OU的id，用于标示该发票所述的OU。 |
| `invoice_channel` | `InvoiceChannel` | `*string` | JSON body | 未标注 | 开票渠道 01 线上;02 线下 |
| `invoice_code` | `InvoiceCode` | `*string` | JSON body | 未标注 | 发票代码，税务部门给予发票的编码 |
| `invoice_create_date_begin` | `InvoiceCreateDateBegin` | `*string` | JSON body | 未标注 | 发票创建日期【起始日期，格式yyyyMMdd】 |
| `invoice_create_date_end` | `InvoiceCreateDateEnd` | `*string` | JSON body | 未标注 | 发票创建日期【结束日期，格式yyyyMMdd】 |
| `invoice_ids` | `InvoiceIds` | `*[]string` | JSON body | 未标注 | 发票ID集合，发票ID，唯一标示一张发票的ID |
| `invoice_material` | `InvoiceMaterial` | `*string` | JSON body | 未标注 | 开票介质 01电子 02纸质 |
| `invoice_no` | `InvoiceNo` | `*string` | JSON body | 未标注 | 发票号码，税务部门给予发票的编码 |
| `invoice_open_date_begin` | `InvoiceOpenDateBegin` | `*string` | JSON body | 未标注 | 发票开票日期【起始日期，格式yyyyMMdd】 |
| `invoice_open_date_end` | `InvoiceOpenDateEnd` | `*string` | JSON body | 未标注 | 发票开票日期【结束日期，格式yyyyMMdd】 |
| `invoice_status` | `InvoiceStatus` | `*[]string` | JSON body | 未标注 | 发票状态：用于标示发票当前的状态，状态类型以及相关含义：01 待开票;02 开票中;03 已开票;04 作废中;05 已作废;06 红冲中;07 已红冲;08 部分红冲;09 无效 |
| `invoice_types` | `InvoiceTypes` | `*[]string` | JSON body | 未标注 | 发票类型,01专票 02普票 03营业税发票 04国际形式发票 05其它发票 |
| `is_online` | `IsOnline` | `*string` | JSON body | 未标注 | 是否能在线开票 Y能 N不能 |
| `is_red` | `IsRed` | `*string` | JSON body | 未标注 | 是否红字 Y红字 N蓝字 |
| `mail_status` | `MailStatus` | `*[]string` | JSON body | 未标注 | 邮寄状态 01 待邮寄;02 已邮寄;03 退回中;04 已退回 |
| `monthly_bill_no` | `MonthlyBillNo` | `*string` | JSON body | 未标注 | 月账单号 |
| `page_no` | `PageNo` | `*int64` | JSON body | 未标注 | 起始页，分页时必须提供 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页显示数，分页时必须提供 |
| `tracking_no` | `TrackingNo` | `*string` | JSON body | 未标注 | 快递单号 |

<a id="api-fncinvoicequery"></a>

### FncInvoiceQuery — 根据发票ID查询发票信息

- 官方接口：[`alipay.boss.fnc.invoice.query`](https://opendocs.alipay.com/pre-apis/05xhqn)
- 路由：`GET /v3/alipay/boss/fnc/invoice/query`
- 请求类型：[BossFncInvoiceQueryRequest](invoice_request.go#L81)
- 响应类型：[BossFncInvoiceQueryResponse](invoice_response.go#L29)

根据发票综合查询接口（alipay.boss.fnc.invoice.batchquery）查询到发票列表后，再根据每个发票对应的发票id查询发票的详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice_id` | `InvoiceId` | `*string` | query | 否 | 查询ID：用于标示一张发票，通过发票综合查询可以查询到该ID。 |
