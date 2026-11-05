# invoicereceipt 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/boss/fnc/invoicereceipt`。本包提供 2 个支付宝 V3 REST 操作。

通过 `invoicereceipt.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`FncInvoiceReceiptBatchQuery`](#api-fncinvoicereceiptbatchquery) | `alipay.boss.fnc.invoicereceipt.batchquery` |
| [`FncInvoiceReceiptCreate`](#api-fncinvoicereceiptcreate) | `alipay.boss.fnc.invoicereceipt.create` |

## 使用示例

- [FncInvoiceReceiptBatchQuery](api_invoicereceipt_fncinvoicereceiptbatchquery_example_test.go)
- [FncInvoiceReceiptCreate](api_invoicereceipt_fncinvoicereceiptcreate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-fncinvoicereceiptbatchquery"></a>

### FncInvoiceReceiptBatchQuery — 根据月账单号集合和来源查询开票单据及金额汇总

- 官方接口：[`alipay.boss.fnc.invoicereceipt.batchquery`](https://opendocs.alipay.com/pre-apis/05xiqt)
- 路由：`POST /v3/alipay/boss/fnc/invoicereceipt/batchquery`
- 请求类型：[BossFncInvoiceReceiptBatchQueryRequest](invoicereceipt_request.go#L4)
- 响应类型：[BossFncInvoiceReceiptBatchQueryResponse](invoicereceipt_response.go#L10)

根据月账单号集合和来源查询开票单据及金额汇总

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bill_nos` | `BillNos` | `*[]string` | JSON body | 未标注 | 月账单号集合，唯一性ID集合，月账单表唯一主键ID集合 |
| `out_biz_type` | `OutBizType` | `*string` | JSON body | 未标注 | 来源类型，01：主站；02：芝麻；03：金融云；04：微贷 |

<a id="api-fncinvoicereceiptcreate"></a>

### FncInvoiceReceiptCreate — 账单变更同步接口月帐单创建消息同步

- 官方接口：[`alipay.boss.fnc.invoicereceipt.create`](https://opendocs.alipay.com/pre-apis/05xnh4)
- 路由：`POST /v3/alipay/boss/fnc/invoicereceipt/create`
- 请求类型：[BossFncInvoiceReceiptCreateRequest](invoicereceipt_request.go#L12)
- 响应类型：[BossFncInvoiceReceiptCreateResponse](invoicereceipt_response.go#L19)

月帐单创建消息同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `event_code` | `EventCode` | `*string` | JSON body | 未标注 | 账单金额变更消息事件码<br> AR_MTHLY_BILL_CREATE:应收账单创建,<br> AR_MTHLY_BILL_CHANGE:应收金额更新 |
| `event_type` | `EventType` | `*string` | JSON body | 未标注 | 账单金额变更消息事件类型<br> CREATE:创建事件,UPDATE:更新事件 |
| `monthly_bill` | `MonthlyBill` | [`*ArMonthlyBillDTO`](#model-armonthlybilldto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `msg_id` | `MsgId` | `*string` | JSON body | 未标注 | 事件全局id，可以使用uuid |
| `out_biz_type` | `OutBizType` | `*string` | JSON body | 未标注 | 消息来源 01:主站，02：芝麻，03:金融云,04:微贷 |

## 嵌套请求参数

<a id="model-armonthlybilldto"></a>

### ArMonthlyBillDTO

类型定义：[ArMonthlyBillDTO](invoicereceipt_model.go#L8)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `account_day` | `AccountDay` | `*string` | 未标注 | 出账日 |
| `accrued_date` | `AccruedDate` | `*string` | 未标注 | 应收日期(yyyyMMdd) |
| `adjust_amt` | `AdjustAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `analysis_dmsn_1` | `AnalysisDmsn1` | `*string` | 未标注 | 分析维度1 |
| `analysis_dmsn_2` | `AnalysisDmsn2` | `*string` | 未标注 | 分析维度2 |
| `analysis_dmsn_3` | `AnalysisDmsn3` | `*string` | 未标注 | 分析维度3 |
| `analysis_dmsn_4` | `AnalysisDmsn4` | `*string` | 未标注 | 分析维度4 |
| `arrangement_no` | `ArrangementNo` | `*string` | 未标注 | 合约号，商户签约产品对应的合约号 |
| `bill_amt` | `BillAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `bill_end_date` | `BillEndDate` | `*string` | 未标注 | 账单结束日期 |
| `bill_month` | `BillMonth` | `*string` | 未标注 | 账单月 |
| `bill_no` | `BillNo` | `*string` | 未标注 | 月账单号,唯一标识月账单的ID |
| `bill_start_date` | `BillStartDate` | `*string` | 未标注 | 账单开始日期 |
| `bill_type` | `BillType` | `*string` | 未标注 | 账单类型<br> 21：应收月账单<br> 目前仅支持该类型的账单类型 |
| `charge_item_code` | `ChargeItemCode` | `*string` | 未标注 | 收费项编码 |
| `charge_type` | `ChargeType` | `*string` | 未标注 | 计费类型<br> 01：云在线计费，02：离线计费，03：老主站计收费，04：芝麻在线计收费 |
| `check_status` | `CheckStatus` | `*string` | 未标注 | 出账状态<br> 01：账单生成，02：待确认出账，03：已出账，04：客户已对账，05：对账差异，06：部分出账 |
| `checked_amt` | `CheckedAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `checking_amt` | `CheckingAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `clcn_basic_amt` | `ClcnBasicAmt` | `*int64` | 未标注 | 计算依据 |
| `clcn_basic_type` | `ClcnBasicType` | `*string` | 未标注 | 计算依据类型<br> 01：交易金额，02：交易笔数，03：交易张数，04：资费依据扩展，05：其他 |
| `clcn_method` | `ClcnMethod` | `*string` | 未标注 | 计算方法<br> 01：差额累进，02：全额累进，03：按传入金额收费，04：价格包量，05：包时间周期，06：按时间全额累进，07：包时间周期差额累进，08：期间差额累进，09：期间全额累进，10：包时间周期单笔，11：标准单笔计费，12：包时间周期单笔-按照传入金额计费，13：期间靠档差额累进，靠档计息使用，14：自定义 |
| `freeze_amt` | `FreezeAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `gmt_create` | `GmtCreate` | `*string` | 未标注 | 创建时间（创建备份时需要） |
| `gmt_modified` | `GmtModified` | `*string` | 未标注 | 修改时间（创建备份时需要） |
| `gmt_pay` | `GmtPay` | `*string` | 未标注 | 付款时间 |
| `inst_id` | `InstId` | `*string` | 未标注 | 分支机构号 |
| `invoice_amt` | `InvoiceAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `invoiced_amt` | `InvoicedAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `ip_id` | `IpId` | `*string` | 未标注 | 结算对象参与者标识 |
| `ip_name` | `IpName` | `*string` | 未标注 | 参与者名字 |
| `ip_role_id` | `IpRoleId` | `*string` | 未标注 | 结算对象参与者角色标识 |
| `last_moder` | `LastModer` | `*string` | 未标注 | 最后修改人 |
| `link_invoice_amt` | `LinkInvoiceAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `pay_original` | `PayOriginal` | `*string` | 未标注 | 付款来源<br> 1：主账务，2：包量，3：预收，4：预存，5：银行，6：积分，7：统一支付，8：支付受理 |
| `pay_status` | `PayStatus` | `*string` | 未标注 | 付款状态<br> 01：待结算；02：部分结算；03：结算完成 |
| `pay_way` | `PayWay` | `*string` | 未标注 | 付款方式<br> 1：资金；2：包量 |
| `payee_account` | `PayeeAccount` | `*string` | 未标注 | 收款方账号 |
| `payer_account` | `PayerAccount` | `*string` | 未标注 | 付款方账号 |
| `payer_ip_role_id` | `PayerIpRoleId` | `*string` | 未标注 | 扣款对象pid |
| `prod_code` | `ProdCode` | `*string` | 未标注 | 产品码 |
| `received_amt` | `ReceivedAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `record_id` | `RecordId` | `*string` | 未标注 | 记录id，账单全局唯一号 |
| `service_amt` | `ServiceAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `settle_type` | `SettleType` | `*string` | 未标注 | 结算类型<br> 01：实时；02：预收；03：后收；04：周期结算；05：按日汇总结算；09：延期结算 |
| `sign_ip_id` | `SignIpId` | `*string` | 未标注 | 签约对象参与者标识 |
| `sign_ip_role_id` | `SignIpRoleId` | `*string` | 未标注 | 签约对象参与者角色标识 |
| `stl_ip_role_id` | `StlIpRoleId` | `*string` | 未标注 | 结算对象PID |
| `tax_amt` | `TaxAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
| `tax_rate` | `TaxRate` | `*int64` | 未标注 | 税率 |
| `tax_type` | `TaxType` | `*string` | 未标注 | 税收类型，01：增值税，02：营业税<br> 营业税已经在大部分行业废弃，要传入需咨询后再确认 |
| `to_writeoff_detail_count` | `ToWriteoffDetailCount` | `*int64` | 未标注 | 待核销明细数量 |
| `type` | `Type` | `*string` | 未标注 | 账单分类 1应收，2返点 |
| `writeoff_detail_dount` | `WriteoffDetailDount` | `*int64` | 未标注 | 核销明细数量 |
| `writingoff_amt` | `WritingoffAmt` | [`*MultiCurrencyMoneyOpenAPI`](../../README.md#model-multicurrencymoneyopenapi) | 未标注 | 见字段类型及官方接口文档。 |
