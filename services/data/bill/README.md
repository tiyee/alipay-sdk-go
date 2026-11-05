# bill 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/data/bill`。本包提供 15 个支付宝 V3 REST 操作。

通过 `bill.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`BillDownloadURLQuery`](#api-billdownloadurlquery) | `alipay.data.dataservice.bill.downloadurl.query` |
| [`BillBalance`](#api-billbalance) | `alipay.data.bill.balance.query` |
| [`BillBalanceHis`](#api-billbalancehis) | `alipay.data.bill.balancehis.query` |
| [`BillAccountLog`](#api-billaccountlog) | `alipay.data.bill.accountlog.query` |
| [`BillAccountBookEReceiptApply`](#api-billaccountbookereceiptapply) | `alipay.data.bill.accountbookereceipt.apply` |
| [`BillAccountBookEReceiptQuery`](#api-billaccountbookereceiptquery) | `alipay.data.bill.accountbookereceipt.query` |
| [`BillBailQuery`](#api-billbailquery) | `alipay.data.bill.bail.query` |
| [`BillBizFundAgentQuery`](#api-billbizfundagentquery) | `alipay.data.bill.bizfundagent.query` |
| [`BillBuyQuery`](#api-billbuyquery) | `alipay.data.bill.buy.query` |
| [`BillEReceiptApply`](#api-billereceiptapply) | `alipay.data.bill.ereceipt.apply` |
| [`BillEReceiptQuery`](#api-billereceiptquery) | `alipay.data.bill.ereceipt.query` |
| [`BillEReceiptAgentApply`](#api-billereceiptagentapply) | `alipay.data.bill.ereceiptagent.apply` |
| [`BillSellQuery`](#api-billsellquery) | `alipay.data.bill.sell.query` |
| [`BillTransferQuery`](#api-billtransferquery) | `alipay.data.bill.transfer.query` |
| [`BillTransferAccountBookQuery`](#api-billtransferaccountbookquery) | `alipay.data.bill.transferaccountbook.query` |

## 使用示例

- [BillDownloadURLQuery](api_bill_billdownloadurlquery_example_test.go)
- [BillBalance](api_bill_billbalance_example_test.go)
- [BillBalanceHis](api_bill_billbalancehis_example_test.go)
- [BillAccountLog](api_bill_billaccountlog_example_test.go)
- [BillAccountBookEReceiptApply](api_bill_billaccountbookereceiptapply_example_test.go)
- [BillAccountBookEReceiptQuery](api_bill_billaccountbookereceiptquery_example_test.go)
- [BillBailQuery](api_bill_billbailquery_example_test.go)
- [BillBizFundAgentQuery](api_bill_billbizfundagentquery_example_test.go)
- [BillBuyQuery](api_bill_billbuyquery_example_test.go)
- [BillEReceiptApply](api_bill_billereceiptapply_example_test.go)
- [BillEReceiptQuery](api_bill_billereceiptquery_example_test.go)
- [BillEReceiptAgentApply](api_bill_billereceiptagentapply_example_test.go)
- [BillSellQuery](api_bill_billsellquery_example_test.go)
- [BillTransferQuery](api_bill_billtransferquery_example_test.go)
- [BillTransferAccountBookQuery](api_bill_billtransferaccountbookquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-billdownloadurlquery"></a>

### BillDownloadURLQuery — 查询对账单下载地址

- 官方接口：[`alipay.data.dataservice.bill.downloadurl.query`](https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49)
- 路由：`GET /v3/alipay/data/dataservice/bill/downloadurl/query`
- 请求类型：[BillDownloadURLRequest](bill_request.go#L9)
- 响应类型：[BillDownloadURLResponse](bill_response.go#L6)

为方便商户快速查账，支持商户通过本接口获取商户离线账单下载地址

请求约束：BillDownloadURLRequest 是 alipay.data.dataservice.bill.downloadurl.query
（对账单下载地址查询，V3 GET）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bill_type` | `BillType` | `string` | query | 否 | 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型。<br>SDK 字段说明：账单类型（如 trade、signcustomer）；<br>取值依据各产品对账单说明。 |
| `bill_date` | `BillDate` | `string` | query | 否 | 账单时间： * 日账单格式为yyyy-MM-dd，最早可下载近6年的日账单。不支持下载当日账单，只能下载前一日24点前的账单数据（T+1），当日数据一般于次日 9 点前生成，特殊情况可能延迟。<br> * 月账单格式为yyyy-MM，最早可下载近6年的月账单。不支持下载当月账单，只能下载上一月账单数据，当月账单一般在次月 3 日生成，特殊情况可能延迟。<br> * 当biz_type为settlementMerge时候，时间为汇总批次结算资金到账的日期，日期格式为yyyy-MM-dd，最早可下载2023年4月17日及以后的账单。<br> * 接口调用仅支持下载近6年的账单，更多账单请前往 b.alipay.com 对账中心-账单下载页下载。<br>SDK 字段说明：账单时间：日账单 yyyy-MM-dd（T+1，不支持当日），<br>月账单 yyyy-MM（次月 3 日生成）。 |
| `smid` | `SMID` | `*string` | query | 否 | 二级商户smid，这个参数只在bill_type是trade_zft_merchant时才能使用<br>SDK 字段说明：二级商户 smid；仅 bill_type=trade_zft_merchant 时可用。 |
| `secure` | `Secure` | `*string` | query | 否 | true表示使用安全链接，即返回的下载链接为https。 非true值的情况下，统一为false，使用http链接。<br>SDK 字段说明：标识返回 https 安全链接；非 true 使用 http 链接。 |

注意事项：

- 下载地址 30 秒失效；获取后应立即下载并校验文件内容。

<a id="api-billbalance"></a>

### BillBalance — 支付宝商家账户当前余额查询

- 官方接口：[`alipay.data.bill.balance.query`](https://opendocs.alipay.com/open-v3/9b66d9d8_alipay.data.bill.balance.query)
- 路由：`GET /v3/alipay/data/bill/balance/query`
- 请求类型：[BillBalanceRequest](bill_request.go#L24)
- 响应类型：[BillBalanceResponse](bill_response.go#L15)

为支付宝商家提供支付宝账户当前余额信息，包括冻结金额，可用余额和总额，供对账使用。

请求约束：BillBalanceRequest 是 alipay.data.bill.balance.query（当前余额查询，
V3 GET）的请求；查询自身账户时无需传 bill_user_id。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bill_user_id` | `BillUserID` | `*string` | query | 否 | 目标查询账户（仅支持部分场景，查询自身时候不需要传递当前字段）。<br>SDK 字段说明：目标查询账户；仅支持部分场景。 |

<a id="api-billbalancehis"></a>

### BillBalanceHis — 支付宝商家账户历史余额查询

- 官方接口：[`alipay.data.bill.balancehis.query`](https://opendocs.alipay.com/open-v3/029a9c24_alipay.data.bill.balancehis.query)
- 路由：`GET /v3/alipay/data/bill/balancehis/query`
- 请求类型：[BillBalanceHisRequest](bill_request.go#L32)
- 响应类型：[BillBalanceHisResponse](bill_response.go#L27)

为支付宝商家提供支付宝账户历史余额信息，供对账使用。

请求约束：BillBalanceHisRequest 是 alipay.data.bill.balancehis.query
（历史余额查询，V3 GET）的请求：biz_date 与 biz_month 互斥，
传入其一。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_date` | `BizDate` | `*string` | query | 否 | 查询指定自然日的起初余额和期末余额，biz_date与biz_month传入参数互斥<br>SDK 字段说明：查询指定自然日（yyyy-MM-dd）的期初与期末余额。 |
| `biz_month` | `BizMonth` | `*string` | query | 否 | 查询指定自然月的起初余额和期末余额，biz_date与biz_month传入参数互斥<br>SDK 字段说明：查询指定自然月（yyyy-MM）的期初与期末余额。 |
| `bill_user_id` | `BillUserID` | `*string` | query | 未标注 | 目标查询账户；仅支持部分场景。<br>此字段存在于 SDK，官方快照未列出，业务支持条件尚未确认。 |

注意事项：

- biz_date 与 biz_month 互斥，必须传入其一。

<a id="api-billaccountlog"></a>

### BillAccountLog — 支付宝商家账户账务明细查询

- 官方接口：[`alipay.data.bill.accountlog.query`](https://opendocs.alipay.com/open-v3/4cce9615_alipay.data.bill.accountlog.query)
- 路由：`GET /v3/alipay/data/bill/accountlog/query`
- 请求类型：[BillAccountLogRequest](bill_request.go#L44)
- 响应类型：[BillAccountLogResponse](bill_response.go#L35)

为支付宝商家提供支付宝账户的账务流水信息，供对账使用。

请求约束：BillAccountLogRequest 是 alipay.data.bill.accountlog.query
（账务明细查询，V3 GET）的请求：时间范围与订单号二选一使用，
支付宝订单号与商户订单号互斥。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 账务流水创建时间的起始范围。只能查询一年内的记录。精确查询不需要指定<br>SDK 字段说明：账务流水创建时间的起始范围；只能查询一年内的记录。 |
| `end_time` | `EndTime` | `*string` | query | 否 | 账务流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。精确查询不需要指定<br>SDK 字段说明：账务流水创建时间的结束范围；与起始时间间隔不超过 31 天，<br>查询结果为左闭右开区间。 |
| `alipay_order_no` | `AlipayOrderNo` | `*string` | query | 否 | 支付宝订单号，通过支付宝订单号精确查询相关的流水明细。支付宝订单号和商户订单号互斥。选择支付宝订单号时，不可选商户订单号。<br>SDK 字段说明：支付宝订单号精确查询；与 MerchantOrderNo 互斥。 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | query | 否 | 商户订单号，通过商户订单号精确查询相关的流水明细。支付宝订单号和商户订单号互斥。选择商户订单号时，不可选支付宝订单号。<br>SDK 字段说明：商户订单号精确查询；与 AlipayOrderNo 互斥。 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页号，从1开始<br>SDK 字段说明：分页号，从 1 开始。 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000<br>SDK 字段说明：分页大小（1000-2000），默认 2000。 |
| `trans_code` | `TransCode` | `*string` | query | 否 | 账务的类型代码，特殊场景下使用<br>SDK 字段说明：账务的类型代码；特殊场景下使用。 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 协议授权码，特殊场景下使用<br>SDK 字段说明：协议授权码；特殊场景下使用。 |
| `agreement_product_code` | `AgreementProductCode` | `*string` | query | 否 | 协议产品码。特殊场景下使用<br>SDK 字段说明：协议产品码；特殊场景下使用。 |
| `bill_user_id` | `BillUserID` | `*string` | query | 否 | 指定用户做账单查询 |
| `open_id` | `OpenID` | `*string` | query | 否 | 用于标记支付宝用户在应用下的唯一标识<br>SDK 字段说明：用户在该应用下的唯一标识。 |

注意事项：

- 时间范围查询间隔不超过 31 天；精确查询使用订单号（互斥）。

<a id="api-billaccountbookereceiptapply"></a>

### BillAccountBookEReceiptApply — 申请子账本电子回单(incubating)

- 官方接口：[`alipay.data.bill.accountbookereceipt.apply`](https://opendocs.alipay.com/pre-apis/05wyde)
- 路由：`POST /v3/alipay/data/bill/accountbookereceipt/apply`
- 请求类型：[DataBillAccountBookEReceiptApplyRequest](bill_request.go#L109)
- 响应类型：[DataBillAccountBookEReceiptApplyResponse](bill_response.go#L47)

申请子账本电子回单（子账本业务定制接口）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 协议号 |
| `key` | `Key` | `*string` | JSON body | 未标注 | 明细凭证，传入流水号（转账接口查询结果）。汇总凭证，传入起止时间，格式yyyy-MM-dd_yyyy-MM-dd。 |
| `store_no` | `StoreNo` | `*string` | JSON body | 未标注 | 子账本号，或者子账本名称。模糊查询 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 申请的类型。可传入：FUND_PLATFORM_DETAIL - 资金明细证明，FUND_PLATFORM_SUM - 资金汇总证明 |

<a id="api-billaccountbookereceiptquery"></a>

### BillAccountBookEReceiptQuery — 查询子账本电子回单状态(incubating)

- 官方接口：[`alipay.data.bill.accountbookereceipt.query`](https://opendocs.alipay.com/open-v3/6ce2ea1b_alipay.data.bill.accountbookereceipt.query)
- 路由：`GET /v3/alipay/data/bill/accountbookereceipt/query`
- 请求类型：[DataBillAccountBookEReceiptQueryRequest](bill_request.go#L147)
- 响应类型：[DataBillAccountBookEReceiptQueryResponse](bill_response.go#L53)

电子签章申请分两个步骤： 1、使用alipay.data.bill.ereceipt.apply创建申请，并获取file_id信息。 2、等待一段时间后，使用file_id查询alipay.data.bill.ereceipt.query申请信息。 单个申请建议每5分钟查询一次状态。汇总回单建议30分钟查询一次

请求约束：DataBillAccountBookEReceiptQueryRequest 是 alipay.data.bill.accountbookereceipt.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `file_id` | `FileId` | `*string` | query | 否 | 根据申请id查询状态。申请接口可以参考alipay.data.bill.ereceiptagent.apply |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 协议号，根据不同业务协议类型，传入对应类型的协议号，用于isv授权检查并获取商户信息。如果业务类型未指定，则使用默认类型对应的协议号。 |
| `agreement_type` | `AgreementType` | `*string` | query | 否 | 根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入 |
| `secure` | `Secure` | `*string` | query | 否 | true表示使用安全链接，即返回的下载链接为https。 非true值的情况下，统一为false，使用http链接。 |

<a id="api-billbailquery"></a>

### BillBailQuery — 支付宝商家账户保证金查询

- 官方接口：[`alipay.data.bill.bail.query`](https://opendocs.alipay.com/open-v3/4421b78a_alipay.data.bill.bail.query)
- 路由：`GET /v3/alipay/data/bill/bail/query`
- 请求类型：[DataBillBailQueryRequest](bill_request.go#L159)
- 响应类型：[DataBillBailQueryResponse](bill_response.go#L63)

为支付宝商家提供支付宝账户保证金信息，供对账使用。

请求约束：DataBillBailQueryRequest 是 alipay.data.bill.bail.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 保证金流水创建时间的起始范围 |
| `end_time` | `EndTime` | `*string` | query | 否 | 保证金流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间 |
| `bail_type` | `BailType` | `*string` | query | 否 | 保证金类型，目前支持TMALL_S_BAIL-天猫保证金，TAOBAO_BAIL-淘宝保证金 |
| `trans_log_id` | `TransLogId` | `*string` | query | 否 | 保证金流水号。如果查询参数中指定流水号，则只查询流水号相关的记录 |
| `biz_orig_no` | `BizOrigNo` | `*string` | query | 否 | 业务基础订单号。如果查询参数中指定订单号，则只查询相关的记录 |

<a id="api-billbizfundagentquery"></a>

### BillBizFundAgentQuery — ISV代理商户资金业务账单查询

- 官方接口：[`alipay.data.bill.bizfundagent.query`](https://opendocs.alipay.com/open-v3/9b60baa5_alipay.data.bill.bizfundagent.query)
- 路由：`GET /v3/alipay/data/bill/bizfundagent/query`
- 请求类型：[DataBillBizFundAgentQueryRequest](bill_request.go#L173)
- 响应类型：[DataBillBizFundAgentQueryResponse](bill_response.go#L69)

用于ISV代理商户查询商户的资金业务账单

请求约束：DataBillBizFundAgentQueryRequest 是 alipay.data.bill.bizfundagent.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 账单查询时间范围 -- 起始时间 |
| `end_time` | `EndTime` | `*string` | query | 否 | 账单查询时间范围 -- 结束范围。时间范围最大不超过31天。 |
| `page_no` | `PageNo` | `*string` | query | 否 | 页码，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000 |
| `agreement_type` | `AgreementType` | `*string` | query | 否 | 根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入。 安全发协议：FUND_SAFT_SIGN_WITHHOLDING_P 专用金协议：FLEXIBLE_EMPLOYMENT_SERVICE_P |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | ISV与商户授权协议号 |
| `account_book_id` | `AccountBookId` | `*string` | query | 否 | 记账本id，服务商在帮助商户开通时候获取的唯一身份号 |

<a id="api-billbuyquery"></a>

### BillBuyQuery — 支付宝商家账户买入交易查询

- 官方接口：[`alipay.data.bill.buy.query`](https://opendocs.alipay.com/open-v3/83504d51_alipay.data.bill.buy.query)
- 路由：`GET /v3/alipay/data/bill/buy/query`
- 请求类型：[DataBillBuyQueryRequest](bill_request.go#L191)
- 响应类型：[DataBillBuyQueryResponse](bill_response.go#L81)

为支付宝商家提供支付宝账户买入交易信息，供对账使用。

请求约束：DataBillBuyQueryRequest 是 alipay.data.bill.buy.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 交易流水创建时间的起始范围 |
| `end_time` | `EndTime` | `*string` | query | 否 | 交易流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间 |
| `alipay_order_no` | `AlipayOrderNo` | `*string` | query | 否 | 支付宝交易流水号。如果查询参数中指定流水号，则只查询流水号相关的记录 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | query | 否 | 商户交易号。如果查询参数中指定交易号，则只查询相关的记录 |
| `store_no` | `StoreNo` | `*string` | query | 否 | 门店编号，模糊搜索 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页号，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000 |

<a id="api-billereceiptapply"></a>

### BillEReceiptApply — 申请电子回单(incubating)

- 官方接口：[`alipay.data.bill.ereceipt.apply`](https://opendocs.alipay.com/open-v3/46460043_alipay.data.bill.ereceipt.apply)
- 路由：`POST /v3/alipay/data/bill/ereceipt/apply`
- 请求类型：[DataBillEReceiptApplyRequest](bill_request.go#L121)
- 响应类型：[DataBillEReceiptApplyResponse](bill_response.go#L93)

支持商家下载多种类型支付宝资金凭证。包括：余额收支证明、余额收支流水证明、转入转出收支证明、收支汇总证明（日汇总）以及收支汇总证明（月汇总）。
 账单凭证申请分两个步骤：
1、使用alipay.data.bill.ereceipt.apply创建申请，并获取file_id信息。
2、使用file_id查询alipay.data.bill.ereceipt.query获取回单信息，回单生成完毕将会返回对应的下载链接，下载链接时效为30秒，过期将无法下载，需重新调用alipay.data.bill.ereceipt.query获取新的下载链接。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bill_user_id` | `BillUserId` | `*string` | JSON body | 未标注 | 授权商户pid。需调用方获取商户授权，才可使用。 |
| `key` | `Key` | `*string` | JSON body | 未标注 | 根据不同业务类型，传入不同参数。 格式如下： * ACCOUNT_LOG_DETAIL - 传入账务流水号，示例：117007*********9151 * ACCOUNT_FLOW_DETAIL - 传入流水日期区间示例：20220301-20220331 * FUND_DETAIL - 传入资金单据号，示例：2022051511007***************0731 * ACCOUNT_LOG_SUM_DAILY - 传入日期，示例：2022-01-01 * ACCOUNT_LOG_SUM_MONTHLY - 传入月份，示例：2022-01 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 申请的类型。可传入： * ACCOUNT_LOG_DETAIL - 余额收支证明 * ACCOUNT_FLOW_DETAIL - 余额收支流水证明 * FUND_DETAIL - 转入转出收支证明 * ACCOUNT_LOG_SUM_DAILY - 收支汇总证明（日汇总） * ACCOUNT_LOG_SUM_MONTHLY - 收支汇总证明（月汇总） |

<a id="api-billereceiptquery"></a>

### BillEReceiptQuery — 查询电子回单状态(incubating)

- 官方接口：[`alipay.data.bill.ereceipt.query`](https://opendocs.alipay.com/open-v3/65fde737_alipay.data.bill.ereceipt.query)
- 路由：`GET /v3/alipay/data/bill/ereceipt/query`
- 请求类型：[DataBillEReceiptQueryRequest](bill_request.go#L209)
- 响应类型：[DataBillEReceiptQueryResponse](bill_response.go#L99)

电子签章申请分两个步骤：
1、使用alipay.data.bill.ereceipt.apply创建申请，并获取file_id信息。
2、等待一段时间后，使用file_id查询alipay.data.bill.ereceipt.query申请信息。
单个申请建议每5分钟查询一次状态。批量回单建议20分钟查询一次。汇总回单建议30分钟查询一次

请求约束：DataBillEReceiptQueryRequest 是 alipay.data.bill.ereceipt.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `file_id` | `FileId` | `*string` | query | 否 | 根据申请id查询状态。通过 [alipay.data.bill.ereceipt.apply(申请电子回单(incubating))](https://opendocs.alipay.com/apis/api_15/alipay.data.bill.ereceipt.apply) 接口同步响应获取。 |

<a id="api-billereceiptagentapply"></a>

### BillEReceiptAgentApply — ISV代理商户申请电子回单

- 官方接口：[`alipay.data.bill.ereceiptagent.apply`](https://opendocs.alipay.com/open-v3/d8c0f195_alipay.data.bill.ereceiptagent.apply)
- 路由：`POST /v3/alipay/data/bill/ereceiptagent/apply`
- 请求类型：[DataBillEReceiptAgentApplyRequest](bill_request.go#L131)
- 响应类型：[DataBillEReceiptAgentApplyResponse](bill_response.go#L109)

ISV代理商户，申请商户的资金业务回单。
电子签章申请分两个步骤： 1、使用alipay.data.bill.ereceiptagent.apply创建申请，并获取file_id信息。 2、等待一段时间后，使用file_id查询alipay.data.bill.ereceipt.query申请信息。单个申请建议每5分钟查询一次状态，汇总（日/月）回单建议20分钟查询一次。月初回单生成较慢，请耐心等待（汇总回单在月初申请需要等待几个小时） 注：申请成功，并不代表明细查询的流水一定存在。如果流水不存在，在查询时将会提示。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | ISV与商户签约授权协议号 |
| `agreement_type` | `AgreementType` | `*string` | JSON body | 未标注 | 根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入。 安全发协议：FUND_SAFT_SIGN_WITHHOLDING_P 专用金协议：FLEXIBLE_EMPLOYMENT_SERVICE_P 批量有密：FUND_AUTHORIZE_FLOW |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 查询时间区间的结束时间，属于闭区间。 [start_time,end_time] |
| `key` | `Key` | `*string` | JSON body | 未标注 | 根据不同业务类型，传入不同参数。格式如下： * FUND_DETAIL - 资金业务回单。传入资金单号(fundOrderId，可通过充值、转账、提现接口获取) |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 搜索范围的起始时间 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 申请的类型（目前仅支持资金业务回单）。可传入： * FUND_DETAIL - 资金业务回单。 |

<a id="api-billsellquery"></a>

### BillSellQuery — 支付宝商家账户卖出交易查询

- 官方接口：[`alipay.data.bill.sell.query`](https://opendocs.alipay.com/open-v3/3c1a5835_alipay.data.bill.sell.query)
- 路由：`GET /v3/alipay/data/bill/sell/query`
- 请求类型：[DataBillSellQueryRequest](bill_request.go#L215)
- 响应类型：[DataBillSellQueryResponse](bill_response.go#L115)

为支付宝商家提供支付宝账户卖出交易信息，供对账使用。

请求约束：DataBillSellQueryRequest 是 alipay.data.bill.sell.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 卖出交易流水创建时间的起始范围 |
| `end_time` | `EndTime` | `*string` | query | 否 | 卖出交易流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间 |
| `alipay_order_no` | `AlipayOrderNo` | `*string` | query | 否 | 支付宝交易流水号。如果查询参数中指定流水号，则只查询流水号相关的记录 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | query | 否 | 商户交易号。如果查询参数中指定交易号，则只查询相关的记录 |
| `store_no` | `StoreNo` | `*string` | query | 否 | 门店编号，模糊搜索 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页号，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000 |

<a id="api-billtransferquery"></a>

### BillTransferQuery — 支付宝商家账户充值，转账，提现查询

- 官方接口：[`alipay.data.bill.transfer.query`](https://opendocs.alipay.com/open-v3/8ddb5fed_alipay.data.bill.transfer.query)
- 路由：`GET /v3/alipay/data/bill/transfer/query`
- 请求类型：[DataBillTransferQueryRequest](bill_request.go#L233)
- 响应类型：[DataBillTransferQueryResponse](bill_response.go#L127)

为支付宝商家提供支付宝账户充值，转账，提现信息，供对账使用。

请求约束：DataBillTransferQueryRequest 是 alipay.data.bill.transfer.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 充值、转账、提现流水业务时间的起始范围 |
| `end_time` | `EndTime` | `*string` | query | 否 | 充值、转账、提现流水业务时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间 |
| `type` | `Type` | `*string` | query | 否 | 转账类型：充值-DEPOSIT，提现-WITHDRAW，转账-TRANSFER。 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页号，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000 |

<a id="api-billtransferaccountbookquery"></a>

### BillTransferAccountBookQuery — 子账本充提转账单查询(incubating)

- 官方接口：[`alipay.data.bill.transferaccountbook.query`](https://opendocs.alipay.com/pre-apis/05wycr)
- 路由：`GET /v3/alipay/data/bill/transferaccountbook/query`
- 请求类型：[DataBillTransferAccountBookQueryRequest](bill_request.go#L247)
- 响应类型：[DataBillTransferAccountBookQueryResponse](bill_response.go#L139)

子账本充提转账单查询（子账本业务定制接口）

请求约束：DataBillTransferAccountBookQueryRequest 是 alipay.data.bill.transferaccountbook.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_time` | `StartTime` | `*string` | query | 否 | 充值、转账、提现流水业务时间的起始范围 |
| `end_time` | `EndTime` | `*string` | query | 否 | 充值、转账、提现流水业务时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间 |
| `type` | `Type` | `*string` | query | 否 | 转账类型：充值-DEPOSIT，提现-WITHDRAW，转账-TRANSFER。 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 协议号 |
| `store_no` | `StoreNo` | `*string` | query | 否 | 子账本号，或者子账本名称。模糊查询 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页号，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页大小1000-2000，默认2000 |
