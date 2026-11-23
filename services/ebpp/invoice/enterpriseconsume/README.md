# ebpp/invoice/enterpriseconsume 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/enterpriseconsume`。本目录提供 8 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `enterpriseconsume.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceEnterpriseConsumeConsumeBatchQuery`](#api-invoiceenterpriseconsumeconsumebatchquery) | [`alipay.ebpp.invoice.enterpriseconsume.consume.batchquery`](https://opendocs.alipay.com/apis/05jgvh) | 企业因公账单查询 | `POST` |
| [`InvoiceEnterpriseConsumeDetailQuery`](#api-invoiceenterpriseconsumedetailquery) | [`alipay.ebpp.invoice.enterpriseconsume.detail.query`](https://opendocs.alipay.com/apis/05j4ce) | 凭证详情查询接口 | `GET` |
| [`InvoiceEnterpriseConsumeEnterpriseOpenRuleCreate`](#api-invoiceenterpriseconsumeenterpriseopenrulecreate) | [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create`](https://opendocs.alipay.com/apis/05wvsl) | 新增开票规则 | `POST` |
| [`InvoiceEnterpriseConsumeEnterpriseOpenRuleModify`](#api-invoiceenterpriseconsumeenterpriseopenrulemodify) | [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify`](https://opendocs.alipay.com/apis/05wzal) | 修改开票规则 | `POST` |
| [`InvoiceEnterpriseConsumeEnterpriseOpenRuleQuery`](#api-invoiceenterpriseconsumeenterpriseopenrulequery) | [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query`](https://opendocs.alipay.com/apis/05wye3) | 查询开票规则 | `GET` |
| [`InvoiceEnterpriseConsumeOpenRuleRelationCreate`](#api-invoiceenterpriseconsumeopenrulerelationcreate) | [`alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create`](https://opendocs.alipay.com/apis/05ucs6) | 新增开票规则并绑定费控规则 | `POST` |
| [`InvoiceEnterpriseConsumeRelateDetailQuery`](#api-invoiceenterpriseconsumerelatedetailquery) | [`alipay.ebpp.invoice.enterpriseconsume.relatedetail.query`](https://opendocs.alipay.com/apis/05ip0t) | 合花账单关联详情查询 | `GET` |
| [`InvoiceEnterpriseConsumeSummaryInvoiceDetailQuery`](#api-invoiceenterpriseconsumesummaryinvoicedetailquery) | [`alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query`](https://opendocs.alipay.com/apis/05imom) | 查询汇总发票详情 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceenterpriseconsumeconsumebatchquery"></a>

### InvoiceEnterpriseConsumeConsumeBatchQuery — 企业因公账单查询

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.consume.batchquery`](https://opendocs.alipay.com/apis/05jgvh)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseconsume/consume/batchquery`
- 请求类型：[InvoiceEnterpriseConsumeConsumeBatchQueryRequest](enterpriseconsume_request.go#L5)
- 响应类型：[InvoiceEnterpriseConsumeConsumeBatchQueryResponse](enterpriseconsume_response.go#L5)

基于企业ID和员工UID，在指定时间段（时间间隔不超过24小时）内获取因公账单数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `consumption_end` | `ConsumptionEnd` | `*string` | JSON body | 未标注 | 交易结束时间，格式：yyyy-MM-dd HH:mm:ss |
| `consumption_start` | `ConsumptionStart` | `*string` | JSON body | 未标注 | 交易开始时间，格式：yyyy-MM-dd HH:mm:ss |
| `employee_list` | `EmployeeList` | `*[]string` | JSON body | 未标注 | 员工支付宝UID列表，单次传入最大员工数量为10 |
| `employee_open_ids` | `EmployeeOpenIds` | `*[]string` | JSON body | 未标注 | 员工支付宝UID列表，单次传入最大员工数量为10 |

<a id="api-invoiceenterpriseconsumedetailquery"></a>

### InvoiceEnterpriseConsumeDetailQuery — 凭证详情查询接口

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.detail.query`](https://opendocs.alipay.com/apis/05j4ce)
- 路由：`GET /v3/alipay/ebpp/invoice/enterpriseconsume/detail/query`
- 请求类型：[InvoiceEnterpriseConsumeDetailQueryRequest](enterpriseconsume_request.go#L103)
- 响应类型：[InvoiceEnterpriseConsumeDetailQueryResponse](enterpriseconsume_response.go#L11)

获取凭证基础信息结构化数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `voucher_id` | `VoucherId` | `*string` | query | 否 | 账单凭证ID |

<a id="api-invoiceenterpriseconsumeenterpriseopenrulecreate"></a>

### InvoiceEnterpriseConsumeEnterpriseOpenRuleCreate — 新增开票规则

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create`](https://opendocs.alipay.com/apis/05wvsl)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/create`
- 请求类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest](enterpriseconsume_request.go#L21)
- 响应类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateResponse](enterpriseconsume_response.go#L18)

新增一条开票规则配置

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `bill_month_day` | `BillMonthDay` | `*int64` | JSON body | 未标注 | 账期结束日期 |
| `bill_scope` | `BillScope` | `*string` | JSON body | 未标注 | 开票账单范围 |
| `combined_pay_mode` | `CombinedPayMode` | `*string` | JSON body | 未标注 | 组合支付模式 |
| `default_invoice_kind` | `DefaultInvoiceKind` | `*string` | JSON body | 未标注 | 默认开票类型 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `invoice_remark_value_rule` | `InvoiceRemarkValueRule` | `*string` | JSON body | 未标注 | 发票备注取值规则。NONE表示不使用额度发放说明；QUOTA_ISSUE_DESC表示将因公付额度发放说明作为发票备注。不传、传null或空字符串时默认NONE。 |
| `invoice_rule_name` | `InvoiceRuleName` | `*string` | JSON body | 未标注 | 开票规则名称 |
| `invoice_title_id` | `InvoiceTitleId` | `*string` | JSON body | 未标注 | 发票抬头 |
| `open_mode` | `OpenMode` | `*string` | JSON body | 未标注 | 开票模式 |
| `receive_address` | `ReceiveAddress` | `*string` | JSON body | 未标注 | 收件人地址 |
| `receive_name` | `ReceiveName` | `*string` | JSON body | 未标注 | 收件人姓名 |
| `receive_phone` | `ReceivePhone` | `*string` | JSON body | 未标注 | 收件人手机号 |
| `seller_type` | `SellerType` | `*string` | JSON body | 未标注 | 销方类型 |
| `tag` | `Tag` | `*string` | JSON body | 未标注 | 开票规则标记 |

<a id="api-invoiceenterpriseconsumeenterpriseopenrulemodify"></a>

### InvoiceEnterpriseConsumeEnterpriseOpenRuleModify — 修改开票规则

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify`](https://opendocs.alipay.com/apis/05wzal)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/modify`
- 请求类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest](enterpriseconsume_request.go#L57)
- 响应类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyResponse](enterpriseconsume_response.go#L24)

修改开票规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `bill_scope` | `BillScope` | `*string` | JSON body | 未标注 | 开票账单范围 |
| `combined_pay_mode` | `CombinedPayMode` | `*string` | JSON body | 未标注 | 组合支付模式 |
| `default_invoice_kind` | `DefaultInvoiceKind` | `*string` | JSON body | 未标注 | 默认开票类型 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `invoice_remark_value_rule` | `InvoiceRemarkValueRule` | `*string` | JSON body | 未标注 | 发票备注取值规则。NONE表示不使用额度发放说明；QUOTA_ISSUE_DESC表示将因公付额度发放说明作为发票备注。修改接口未传、传null或空字符串时保留原配置。 |
| `invoice_rule_id` | `InvoiceRuleId` | `*string` | JSON body | 未标注 | 开票规则ID |
| `invoice_rule_name` | `InvoiceRuleName` | `*string` | JSON body | 未标注 | 开票规则名称 |
| `invoice_title_id` | `InvoiceTitleId` | `*string` | JSON body | 未标注 | 发票抬头 |
| `receive_address` | `ReceiveAddress` | `*string` | JSON body | 未标注 | 收件人地址 |
| `receive_name` | `ReceiveName` | `*string` | JSON body | 未标注 | 收件人姓名 |
| `receive_phone` | `ReceivePhone` | `*string` | JSON body | 未标注 | 收件人手机号 |
| `seller_type` | `SellerType` | `*string` | JSON body | 未标注 | 销方类型 |

<a id="api-invoiceenterpriseconsumeenterpriseopenrulequery"></a>

### InvoiceEnterpriseConsumeEnterpriseOpenRuleQuery — 查询开票规则

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query`](https://opendocs.alipay.com/apis/05wye3)
- 路由：`GET /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/query`
- 请求类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest](enterpriseconsume_request.go#L113)
- 响应类型：[InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryResponse](enterpriseconsume_response.go#L30)

根据开票规则ID查询开票规则详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业ID |
| `invoice_rule_id` | `InvoiceRuleId` | `*string` | query | 否 | 开票规则ID |
| `account_id` | `AccountId` | `*string` | query | 否 | 共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |

<a id="api-invoiceenterpriseconsumeopenrulerelationcreate"></a>

### InvoiceEnterpriseConsumeOpenRuleRelationCreate — 新增开票规则并绑定费控规则

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create`](https://opendocs.alipay.com/apis/05ucs6)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseconsume/openrulerelation/create`
- 请求类型：[InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest](enterpriseconsume_request.go#L89)
- 响应类型：[InvoiceEnterpriseConsumeOpenRuleRelationCreateResponse](enterpriseconsume_response.go#L35)

新增开票规则并自动绑定到费控规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `invoice_rule_id` | `InvoiceRuleId` | `*string` | JSON body | 未标注 | 开票规则ID |
| `standard_id_list` | `StandardIdList` | `*[]string` | JSON body | 未标注 | 费控规则ID列表 |

<a id="api-invoiceenterpriseconsumerelatedetailquery"></a>

### InvoiceEnterpriseConsumeRelateDetailQuery — 合花账单关联详情查询

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.relatedetail.query`](https://opendocs.alipay.com/apis/05ip0t)
- 路由：`GET /v3/alipay/ebpp/invoice/enterpriseconsume/relatedetail/query`
- 请求类型：[InvoiceEnterpriseConsumeRelateDetailQueryRequest](enterpriseconsume_request.go#L125)
- 响应类型：[InvoiceEnterpriseConsumeRelateDetailQueryResponse](enterpriseconsume_response.go#L41)

合花账单关联详情查询，通过凭证ID获取这个账单详细信息和关联的凭证详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `voucher_id` | `VoucherId` | `*string` | query | 否 | 账单凭证ID |

<a id="api-invoiceenterpriseconsumesummaryinvoicedetailquery"></a>

### InvoiceEnterpriseConsumeSummaryInvoiceDetailQuery — 查询汇总发票详情

- 官方接口：[`alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query`](https://opendocs.alipay.com/apis/05imom)
- 路由：`GET /v3/alipay/ebpp/invoice/enterpriseconsume/summaryinvoicedetail/query`
- 请求类型：[InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest](enterpriseconsume_request.go#L135)
- 响应类型：[InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryResponse](enterpriseconsume_response.go#L46)

根据汇总开票通知拿到的开票批次ID查询汇总发票详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业 ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `summary_id` | `SummaryId` | `*string` | query | 否 | 开票批次 ID |
