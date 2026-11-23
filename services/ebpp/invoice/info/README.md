# ebpp/invoice/info 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/info`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `info.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_info_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceInfoSend`](#api-invoiceinfosend) | [`alipay.ebpp.invoice.info.send`](https://opendocs.alipay.com/apis/05l6nx) | 发票信息回传接口（新版） | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceinfosend"></a>

### InvoiceInfoSend — 发票信息回传接口（新版）

- 官方接口：[`alipay.ebpp.invoice.info.send`](https://opendocs.alipay.com/apis/05l6nx)
- 路由：`POST /v3/alipay/ebpp/invoice/info/send`
- 请求类型：[InvoiceInfoSendRequest](info_request.go#L5)
- 响应类型：[InvoiceInfoSendResponse](info_response.go#L5)

商户或ISV将发票信息发送至发票管家。例如：商户为用户开出发票之后，通过该接口将发票发送到买家支付宝的发票管家中，买家可在发票管家中查看发票。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice_info_list` | `InvoiceInfoList` | [`*[]InvoiceSendOpen`](#model-invoicesendopen) | JSON body | 未标注 | 发票信息列表 |
| `m_short_name` | `MShortName` | `*string` | JSON body | 未标注 | 开票商户品牌简称，与商户入驻时的品牌简称保持一致。 |
| `sub_m_short_name` | `SubMShortName` | `*string` | JSON body | 未标注 | 开票商户门店简称，与商户入驻时的门店简称保持一致。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-einvtrade"></a>

### EinvTrade

类型定义：[EinvTrade](info_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `bill_no` | `BillNo` | `*string` | 未标注 | 订单编号 |
| `bill_time` | `BillTime` | `*string` | 未标注 | 下单时间 |
| `city_name` | `CityName` | `*string` | 未标注 | 商户所在城市(经营地址) |
| `detail_json` | `DetailJson` | `*string` | 未标注 | 账单明细信息，酒店水单信息，行程单信息，餐饮小票信息 |
| `download_url` | `DownloadUrl` | `*string` | 未标注 | 账单明细信息，酒店水单，行程单，餐饮小票等pdf原件链接 |
| `extend_map` | `ExtendMap` | `*string` | 未标注 | 扩展参数 不同组的k-v通过换行符区分 |
| `merchant_name` | `MerchantName` | `*string` | 未标注 | 商家名称（显示名称，非企业名称，餐饮店、酒店、打车平台名称） |
| `out_json` | `OutJson` | `*string` | 未标注 | 透传字段，不做处理，用于isv向后续报销税控方传递特殊信息标记 |
| `payee_name` | `PayeeName` | `*string` | 未标注 | 销方名称 |
| `payment_amount` | `PaymentAmount` | `*string` | 未标注 | 支付金额，单位（元）， 对应账单中的交易金额 |
| `payment_time` | `PaymentTime` | `*string` | 未标注 | 支付时间 对应账单中的账单日期 |
| `souce` | `Souce` | `*string` | 未标注 | 交易类型来源 需要按照枚举映射 |
| `trade_type` | `TradeType` | `*string` | 未标注 | 交易类型/账单分类 |

<a id="model-invoiceitemopen"></a>

### InvoiceItemOpen

类型定义：[InvoiceItemOpen](info_model.go#L35)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_ex_tax_amount` | `ItemExTaxAmount` | `*string` | 未标注 | 明细行不含税金额，单位元，保留两位小数 |
| `item_name` | `ItemName` | `*string` | 未标注 | 开票项目：货物或应税劳务、服务名称 |
| `item_no` | `ItemNo` | `*string` | 未标注 | 国税局制定的商品税收编码，必须是最末级 |
| `item_quantity` | `ItemQuantity` | `*int64` | 未标注 | 数量； 1.当row_type=0或2且item_unit_price为空，可空； 2.可精确到小数点后6位 |
| `item_spec` | `ItemSpec` | `*string` | 未标注 | 规格型号 |
| `item_sum_amount` | `ItemSumAmount` | `*string` | 未标注 | 明细行价税合计，单位元，必须保证item_sum_amount=item_ex_tax_amount+item_tax_amount。 |
| `item_tax_amount` | `ItemTaxAmount` | `*string` | 未标注 | 明细行税额，单位元，保留两位小数，无税或者免税情况下输入：0.00。 |
| `item_tax_rate` | `ItemTaxRate` | `*string` | 未标注 | 明细行税率，无税或者免税情况下输入：0.00。 |
| `item_unit` | `ItemUnit` | `*string` | 未标注 | 单位 |
| `item_unit_price` | `ItemUnitPrice` | `*string` | 未标注 | 不含税单价（元）; 1.当row_type=0或2且item_quantity为空，可空 2.可精确到小数点后8位； |
| `row_type` | `RowType` | `*string` | 未标注 | 发票行性质：0表示正常行，1表示折扣行，2表示被折扣行。 |

<a id="model-invoicesendopen"></a>

### InvoiceSendOpen

类型定义：[InvoiceSendOpen](info_model.go#L61)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | 未标注 | 支付宝端的申请id。如果在开票过程中，是通过支付宝提交的申请到机构端，支付宝会带上开票申请在支付宝生成的申请id，机构在回传发票的时候只需要回传这个申请id，不用获取用户的uid，支付宝可以根据申请id将发票归集到对应的用户名下 |
| `check_code` | `CheckCode` | `*string` | 未标注 | 票面上的校验码信息。 |
| `checker` | `Checker` | `*string` | 未标注 | 票面上复核人信息。 |
| `clerk` | `Clerk` | `*string` | 未标注 | 票面上开票员信息。 |
| `ex_tax_amount` | `ExTaxAmount` | `*string` | 未标注 | 不含税金额，票面上的不含税金额，单位（元），保留两位小数。 |
| `extend_fields` | `ExtendFields` | `*string` | 未标注 | 预留的扩展字段，格式如：key1=value1\nkey2=value2\nkey3=value3，字段之间以\n分隔。 |
| `file_download_type` | `FileDownloadType` | `*string` | 未标注 | 下载版式文件类型，需要回传文件下载链接的情况下必传 |
| `file_download_url` | `FileDownloadUrl` | `*string` | 未标注 | 发票版式文件下载地址。 |
| `financial_electronic_type` | `FinancialElectronicType` | `*string` | 未标注 | 财政电子票据子类型 |
| `invoice_code` | `InvoiceCode` | `*string` | 未标注 | 发票代码，该字段与invoice_no构成一张发票的唯一标识。 |
| `invoice_content` | `InvoiceContent` | [`*[]InvoiceItemOpen`](#model-invoiceitemopen) | 未标注 | 发票明细项。 |
| `invoice_date` | `InvoiceDate` | `*string` | 未标注 | 开票日期，格式如：YYYY-MM-DD。 |
| `invoice_kind` | `InvoiceKind` | `*string` | 未标注 | 标识发票的类型， PLAIN:增值税电子普通发票; SPECIAL:增值税专用发票; ALL_ELECTRONIC_GENERAL:电子发票（普通发票）; ALL_ELECTRONIC_SPECIAL:电子发票（增值税专用发票）; PLAIN_INVOICE:增值税普通发票; FINANCIAL_ELECTRONIC_BILL:财政电子票据; |
| `invoice_memo` | `InvoiceMemo` | `*string` | 未标注 | 票面上备注信息。 |
| `invoice_no` | `InvoiceNo` | `*string` | 未标注 | 发票号码，该字段与invoice_code构成一张发票的唯一标识。 |
| `invoice_title` | `InvoiceTitle` | [`*InvoiceTitleOpen`](#model-invoicetitleopen) | 未标注 | 见字段类型及官方接口文档。 |
| `invoice_type` | `InvoiceType` | `*string` | 未标注 | 标识是红票还是蓝票，RED:红票;BLUE:蓝票。 |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户userId，当发送红字发票时，即invoice_type＝RED时，可选填；或者apply_id不为空的时候，可选填；其他情况必填。 |
| `ori_blue_inv_code` | `OriBlueInvCode` | `*string` | 未标注 | 红票情况下，必须填入原始蓝票的发票代码（全电票时为空）。 |
| `ori_blue_inv_no` | `OriBlueInvNo` | `*string` | 未标注 | 红票情况下，必须填入原始蓝票的发票号码。 |
| `out_invoice_id` | `OutInvoiceId` | `*string` | 未标注 | 外部发票id，isv生成发票之后所生成的发票流水号，主要用做存储。 |
| `out_trade_no` | `OutTradeNo` | `*string` | 未标注 | 该发票基于的实际交易流水号。 |
| `payee` | `Payee` | `*string` | 未标注 | 票面上收款人信息。 |
| `payee_address_tel` | `PayeeAddressTel` | `*string` | 未标注 | 票面上销售方地址、电话信息。 |
| `payee_bank_name_account` | `PayeeBankNameAccount` | `*string` | 未标注 | 票面上销售方开户行及账号。 |
| `payee_register_name` | `PayeeRegisterName` | `*string` | 未标注 | 票面上销售方企业名称。 |
| `payee_register_no` | `PayeeRegisterNo` | `*string` | 未标注 | 票面上销售方税号信息。 |
| `sum_amount` | `SumAmount` | `*string` | 未标注 | 价税合计，票面上的总开票金额，单位（元），保留两位小数，sum_amount=ex_tax_amount+tax_amount。 |
| `tax_amount` | `TaxAmount` | `*string` | 未标注 | 合计税额，票面上的合计税额，单位（元），保留两位小数。 |
| `trade_list` | `TradeList` | [`*[]EinvTrade`](#model-einvtrade) | 未标注 | 交易内容，非必填但是不可以传空值，如果无需回传trade_list则不需要传该对象 |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户userId，当发送红字发票时，即invoice_type＝RED时，可选填；或者apply_id不为空的时候，可选填；其他情况必填。 |

<a id="model-invoicetitleopen"></a>

### InvoiceTitleOpen

类型定义：[InvoiceTitleOpen](info_model.go#L126)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `payer_address_tel` | `PayerAddressTel` | `*string` | 未标注 | 票面上的购买方地址、电话 |
| `payer_bank_name_account` | `PayerBankNameAccount` | `*string` | 未标注 | 票面上的购买方开户行及账户 |
| `payer_register_no` | `PayerRegisterNo` | `*string` | 未标注 | 票面上的购买方纳税人识别号 |
| `title_name` | `TitleName` | `*string` | 未标注 | 票面上的购买方名称 |
