# ebpp/invoice/sycn 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/sycn`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `sycn.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_sycn_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceSycn`](#api-invoicesycn) | [`alipay.ebpp.invoice.sycn`](https://opendocs.alipay.com/apis/05la3z) | 发票信息同步至支付宝 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicesycn"></a>

### InvoiceSycn — 发票信息同步至支付宝

- 官方接口：[`alipay.ebpp.invoice.sycn`](https://opendocs.alipay.com/apis/05la3z)
- 路由：`POST /v3/alipay/ebpp/invoice/sycn`
- 请求类型：[InvoiceSycnRequest](sycn_request.go#L5)
- 响应类型：[InvoiceSycnResponse](sycn_response.go#L5)

外部商户开蓝票/冲红执行成功后将发票信息及原件（如有）同步至支付宝，由支付宝通知用户查看已开具的发票。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice_info` | `InvoiceInfo` | [`*[]InvoiceModelContent`](#model-invoicemodelcontent) | JSON body | 未标注 | 同步发票信息模型 |
| `m_short_name` | `MShortName` | `*string` | JSON body | 未标注 | 商户的品牌名称简称,该字段需要接入前向发票管家申请，<br> m_short_name+sub_m_short_name具有唯一约束<br> 如：肯德基：KFC |
| `sub_m_short_name` | `SubMShortName` | `*string` | JSON body | 未标注 | 支付宝为商户分配的商户门店简称，该字段需要接入前在发票管家申请<br> 如：肯德基-杭州西湖区文一西路店：KFC-HZ-XH001 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-invoiceitemcontent"></a>

### InvoiceItemContent

类型定义：[InvoiceItemContent](sycn_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_amount` | `ItemAmount` | `*string` | 未标注 | 价税合计。(等于sumPrice和tax之和) |
| `item_name` | `ItemName` | `*string` | 未标注 | 发票项目名称（或商品名称） |
| `item_no` | `ItemNo` | `*string` | 未标注 | 商品编号 |
| `item_price` | `ItemPrice` | `*string` | 未标注 | 单价，格式：100.00。新版电子发票，折扣行此参数不能传，非折扣行必传 |
| `item_quantity` | `ItemQuantity` | `*int64` | 未标注 | 数量。新版电子发票，折扣行此参数不能传，非折扣行必传 |
| `item_sum_price` | `ItemSumPrice` | `*string` | 未标注 | 单项总价，格式：100.00 |
| `item_tax_price` | `ItemTaxPrice` | `*string` | 未标注 | 税额 |
| `item_tax_rate` | `ItemTaxRate` | `*string` | 未标注 | 税率 |
| `item_unit` | `ItemUnit` | `*string` | 未标注 | 台 |
| `row_type` | `RowType` | `*int64` | 未标注 | 发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0 |

<a id="model-invoicemodelcontent"></a>

### InvoiceModelContent

类型定义：[InvoiceModelContent](sycn_model.go#L29)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | 未标注 | 支付宝端的申请id。如果在开票过程中，是通过支付宝提交的申请到机构端，支付宝会带上开票申请在支付宝生成的申请id，机构在回传发票的时候只需要回传这个申请id，不用获取用户的uid，支付宝可以根据申请id将发票归集到对应的用户名下 |
| `extend_fields` | `ExtendFields` | `*string` | 未标注 | key=value，每组键值对以回车分割 |
| `file_download_type` | `FileDownloadType` | `*string` | 未标注 | 下载的发票文件类型，可选值： pdf（发票原文件） ofd（发票原文件） jpg（发票原文件缩略图） |
| `file_download_url` | `FileDownloadUrl` | `*string` | 未标注 | 发票原文件下载地址 1.当tax_type=PLAIN、ALL_ELECTRONIC_GENERAL或ALL_ELECTRONIC_SPECIAL时， file_download_url必传 且file_download_type取值范围为pdf或ofd； 2.当tax_type=SPECIAL时， file_download_url必传 file_download_type可以传入pdf，ofd，jpg 3.当其他票种时，file_download_url可以不传 |
| `financial_electronic_type` | `FinancialElectronicType` | `*string` | 未标注 | 财政电子票据子类型，当tax_type=FINANCIAL_ELECTRONIC_BILL时要求必填 可选值如下： 01:非税收入通用票据 02:非税收入专用票据 03:非税收入一般缴款书 04:资金往来结算票据 05:公益事业捐赠票据 06:医疗收费票据 07:社会团体会费票据 08:社会保险基金票据 09:工会经费收入票据 99:其他财政票据 |
| `invoice_amount` | `InvoiceAmount` | `*string` | 未标注 | 发票金额，大于0且精确到小数点两位，以元为单位<br> 需要传入税价合计金额 |
| `invoice_code` | `InvoiceCode` | `*string` | 未标注 | 发票代码，为国税局生成的唯一值，全电票时为空，其他情况不可为空 |
| `invoice_content` | `InvoiceContent` | [`*[]InvoiceItemContent`](#model-invoiceitemcontent) | 未标注 | 发票内容项 |
| `invoice_date` | `InvoiceDate` | `*string` | 未标注 | 发票日期，用户填写，目前精确到日 |
| `invoice_fake_code` | `InvoiceFakeCode` | `*string` | 未标注 | 发票防伪码 |
| `invoice_file_data` | `InvoiceFileData` | `*string` | 未标注 | 原始发票PDF/OFD文件流 |
| `invoice_img_url` | `InvoiceImgUrl` | `*string` | 未标注 | 发票原始文件jpg文件地址 |
| `invoice_no` | `InvoiceNo` | `*string` | 未标注 | 发票号码，国税局生成的唯一号码，不可为空串；<br> 使用时请注意，invoice_no+invoice_code唯一，不能重复 |
| `invoice_operator` | `InvoiceOperator` | `*string` | 未标注 | 发票开具操作人 |
| `invoice_title` | `InvoiceTitle` | [`*InvoiceTitle`](#model-invoicetitle) | 未标注 | 见字段类型及官方接口文档。 |
| `invoice_type` | `InvoiceType` | `*string` | 未标注 | 发票类型，按照可选值只传入英文部分，该字段严格要求大小写<br> 可选值:<br> blue（蓝票）<br> red（红票） |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。 |
| `original_blue_invoice_code` | `OriginalBlueInvoiceCode` | `*string` | 未标注 | 仅用于同步红票，原始蓝票发票代码，同步红票时必传（全电票时为空） |
| `original_blue_invoice_no` | `OriginalBlueInvoiceNo` | `*string` | 未标注 | 仅用于同步红票，原始蓝票发票号码，同步红票时必传 |
| `out_biz_no` | `OutBizNo` | `*string` | 未标注 | 商户交易流水号，不可为空串;<br> 传入红票时请注意，此字段的值要和蓝票保持一致 |
| `out_invoice_id` | `OutInvoiceId` | `*string` | 未标注 | 商户唯一开票申请业务流水号，同一个isv下不能重复 |
| `register_address` | `RegisterAddress` | `*string` | 未标注 | 开票单位地址 |
| `register_bank_account` | `RegisterBankAccount` | `*string` | 未标注 | 开票单位开户行账号 |
| `register_bank_name` | `RegisterBankName` | `*string` | 未标注 | 开票单位开户行名称 |
| `register_name` | `RegisterName` | `*string` | 未标注 | 开票单位 |
| `register_no` | `RegisterNo` | `*string` | 未标注 | 纳税人识别号，不可为空串 |
| `register_phone_no` | `RegisterPhoneNo` | `*string` | 未标注 | 开票人电话，支持座机和手机两种格式 |
| `sum_amount` | `SumAmount` | `*string` | 未标注 | 价税合计 |
| `tax_amount` | `TaxAmount` | `*string` | 未标注 | 税额 |
| `tax_type` | `TaxType` | `*string` | 未标注 | 税种 可选值： PLAIN：增值税电子普通发票 ALL_ELECTRONIC_GENERAL：电子发票（普通发票） ALL_ELECTRONIC_SPECIAL：电子发票（增值税专用发票） SPECIAL：增值税专用发票 PLAIN_INVOICE:增值税普通发票 FINANCIAL_ELECTRONIC_BILL:财政电子票据 |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。 |

<a id="model-invoicetitle"></a>

### InvoiceTitle

类型定义：[InvoiceTitle](sycn_model.go#L94)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `is_default` | `IsDefault` | `*bool` | 未标注 | 是否为用户设置默认抬头<br> 字段值包括两种情况：<br> false（非默认）<br> true（默认抬头） |
| `logon_id` | `LogonId` | `*string` | 未标注 | 支付宝用户登录名 |
| `open_bank_account` | `OpenBankAccount` | `*string` | 未标注 | 银行账号 |
| `open_bank_name` | `OpenBankName` | `*string` | 未标注 | 开户银行 |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户id |
| `tax_register_no` | `TaxRegisterNo` | `*string` | 未标注 | 纳税人识别号 |
| `tele_phone_no` | `TelePhoneNo` | `*string` | 未标注 | 用户私人手机号 |
| `title_name` | `TitleName` | `*string` | 未标注 | 抬头名称 |
| `title_type` | `TitleType` | `*string` | 未标注 | 抬头类型 字段值有两种情况抬: PERSONAL（个人） CORPORATION（企业） |
| `user_address` | `UserAddress` | `*string` | 未标注 | 地址 |
| `user_email` | `UserEmail` | `*string` | 未标注 | 邮箱 |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户id |
| `user_mobile` | `UserMobile` | `*string` | 未标注 | 电话号码 |
