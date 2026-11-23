# ebpp/invoice/apply 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/apply`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `apply.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_apply_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceApplyInvUpload`](#api-invoiceapplyinvupload) | [`alipay.ebpp.invoice.apply.inv.upload`](https://opendocs.alipay.com/apis/05ujox) | 发票回传 | `PUT` |
| [`InvoiceApplyResultSync`](#api-invoiceapplyresultsync) | [`alipay.ebpp.invoice.apply.result.sync`](https://opendocs.alipay.com/apis/05vwzk) | ISV向支付宝同步发票申请结果 | `POST` |
| [`InvoiceApplyStatusNotify`](#api-invoiceapplystatusnotify) | [`alipay.ebpp.invoice.apply.status.notify`](https://opendocs.alipay.com/apis/05uo9u) | 发票申请状态变更 | `PUT` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceapplyinvupload"></a>

### InvoiceApplyInvUpload — 发票回传

- 官方接口：[`alipay.ebpp.invoice.apply.inv.upload`](https://opendocs.alipay.com/apis/05ujox)
- 路由：`PUT /v3/alipay/ebpp/invoice/apply/inv/upload`
- 请求类型：[InvoiceApplyInvUploadRequest](apply_request.go#L5)
- 响应类型：[InvoiceApplyInvUploadResponse](apply_response.go#L5)

发票回传接口适用于以下场景：
1、间联开票模式：阿里发票平台向ISV提交了原始发票申请，商户在ISV后台开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台归集与交付。
2、直连开票模式：阿里发票平台请求ISV开票，ISV系统异步执行开票事务后，基于开票流水号(serial_no)回传开票结果至阿里发票平台归集与交付。
3、ISV直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `anti_fake_code` | `AntiFakeCode` | `*string` | JSON body | 未标注 | 发票防伪码/密码 |
| `apply_id` | `ApplyId` | `*string` | JSON body | 未标注 | 发票申请ID, 由阿里发票平台生成。 拆单开票后一笔申请可能对应多张发票。 source=apply时必填，其他source可为空 |
| `biz_error_code` | `BizErrorCode` | `*string` | JSON body | 未标注 | 开票失败错误码， 开票失败(create_result=fail)时必填 |
| `biz_error_msg` | `BizErrorMsg` | `*string` | JSON body | 未标注 | 开票失败错误信息， 开票失败(create_result=fail)时必填 |
| `business_type` | `BusinessType` | `*int64` | JSON body | 未标注 | 抬头类型。可选值： 0：个人 1：企业 |
| `check_code` | `CheckCode` | `*string` | JSON body | 未标注 | 校验码 |
| `create_result` | `CreateResult` | `*string` | JSON body | 未标注 | 开票结果，枚举值： success: 发票开具成功； fail: 开票失败； source=async时必填，传实际的开票结果。其他source可不传，默认为success |
| `device_no` | `DeviceNo` | `*string` | JSON body | 未标注 | 开票分机号/机器编号 |
| `invoice_amount` | `InvoiceAmount` | `*string` | JSON body | 未标注 | 合计含税金额（开票金额），格式为2位小数。 当开红票时，该字段为负数。 开票成功时(create_result=success)必填 |
| `invoice_code` | `InvoiceCode` | `*string` | JSON body | 未标注 | 发票代码。 开票成功时(create_result=success)必填 |
| `invoice_date` | `InvoiceDate` | `*string` | JSON body | 未标注 | 开票日期，格式 yyyy-MM-dd. 开票成功时(create_result=success)必填 |
| `invoice_file_content` | `InvoiceFileContent` | `*string` | JSON body | 未标注 | 发票版式文件数据，base64化的字符串； 电子发票必填 |
| `invoice_file_type` | `InvoiceFileType` | `*string` | JSON body | 未标注 | 发票板式文件类型。可选值： PDF, OFD。 开电票且开票成功时(create_result=success)必填 |
| `invoice_items` | `InvoiceItems` | [`*InvoiceItem`](#model-invoiceitem) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `invoice_kind` | `InvoiceKind` | `*int64` | JSON body | 未标注 | 开票发票类型 可选值： 0: 电票； 1：纸质普票； 2：纸质专票； 开票成功时(create_result=success)必填 |
| `invoice_memo` | `InvoiceMemo` | `*string` | JSON body | 未标注 | 发票备注，会显示在票面 |
| `invoice_no` | `InvoiceNo` | `*string` | JSON body | 未标注 | 发票号码。 开票成功时(create_result=success)必填 |
| `invoice_type` | `InvoiceType` | `*string` | JSON body | 未标注 | 发票(开票)类型，可选值： blue: 蓝票； red: 红票； 开票成功时(create_result=success)必填 |
| `levy_type` | `LevyType` | `*string` | JSON body | 未标注 | 征税方式，0普通征收，1减按征收，2差额征收 |
| `normal_invoice_code` | `NormalInvoiceCode` | `*string` | JSON body | 未标注 | 原发票代码。 冲红且开票成功时(create_result=success)必填 |
| `normal_invoice_no` | `NormalInvoiceNo` | `*string` | JSON body | 未标注 | 原发票号码。 冲红且开票成功时(create_result=success)必填 |
| `payee_address` | `PayeeAddress` | `*string` | JSON body | 未标注 | 销方地址。 开票成功时(create_result=success)必填 |
| `payee_bank_account_id` | `PayeeBankAccountId` | `*string` | JSON body | 未标注 | 销方银行账号 |
| `payee_bank_name` | `PayeeBankName` | `*string` | JSON body | 未标注 | 销方开户行名称 |
| `payee_checker` | `PayeeChecker` | `*string` | JSON body | 未标注 | 复核人 |
| `payee_name` | `PayeeName` | `*string` | JSON body | 未标注 | 销方名称。 开票成功时(create_result=success)必填 |
| `payee_operator` | `PayeeOperator` | `*string` | JSON body | 未标注 | 开票人。 开票成功时(create_result=success)必填 |
| `payee_phone` | `PayeePhone` | `*string` | JSON body | 未标注 | 销方联系电话。 开票成功时(create_result=success)必填 |
| `payee_receiver` | `PayeeReceiver` | `*string` | JSON body | 未标注 | 收款人 |
| `payee_register_no` | `PayeeRegisterNo` | `*string` | JSON body | 未标注 | 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。 开票成功时(create_result=success)必填 |
| `payer_address` | `PayerAddress` | `*string` | JSON body | 未标注 | 购方地址， 专票且开票成功时(create_result=success)必填 |
| `payer_bank_account_id` | `PayerBankAccountId` | `*string` | JSON body | 未标注 | 购方银行账号， 专票且开票成功时(create_result=success)必填 |
| `payer_bank_name` | `PayerBankName` | `*string` | JSON body | 未标注 | 购方开户行名称， 专票且开票成功时(create_result=success)必填 |
| `payer_email` | `PayerEmail` | `*string` | JSON body | 未标注 | 购方电子邮箱，需满足邮箱格式。 格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)* |
| `payer_name` | `PayerName` | `*string` | JSON body | 未标注 | 购方抬头。 开票成功时(create_result=success)必填 |
| `payer_phone` | `PayerPhone` | `*string` | JSON body | 未标注 | 购方联系电话， 专票且开票成功时(create_result=success)必填 |
| `payer_register_no` | `PayerRegisterNo` | `*string` | JSON body | 未标注 | 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。 开企业抬头时必填， 专票且开票成功时(create_result=success)必填 |
| `platform_code` | `PlatformCode` | `*string` | JSON body | 未标注 | 业务平台Code, 由发票中台分配。 用于标识需交付发票的业务平台，source=upload时必填，其他source可为空 |
| `platform_tid` | `PlatformTid` | `*string` | JSON body | 未标注 | 业务平台发票申请对应的订单号。 用于区分业务平台订单号，source=upload时必填，其他source可为空 |
| `platform_user_id` | `PlatformUserId` | `*string` | JSON body | 未标注 | 业务平台发票申请对应的订单号。 用于区分业务平台订单号，source=upload时必填，其他source可为空 |
| `qr_code` | `QrCode` | `*string` | JSON body | 未标注 | 二维码 |
| `receive_mobile` | `ReceiveMobile` | `*string` | JSON body | 未标注 | 购方手机号码，用于收票 |
| `red_notice_no` | `RedNoticeNo` | `*string` | JSON body | 未标注 | 红字通知单号； 专票冲红且开票成功时(create_result=success)必填 |
| `serial_no` | `SerialNo` | `*string` | JSON body | 未标注 | 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。 source=async时必填，其他source可为空 |
| `source` | `Source` | `*string` | JSON body | 未标注 | 发票来源，可选值： apply: 间连开票模式，ISV回传商家基于申请开具的发票； async：直连开票模式，ISV回传开票结果； upload：直接回传，进行归集与交付的发票 |
| `special_flag` | `SpecialFlag` | `*string` | JSON body | 未标注 | 特殊票种标识，可选值： “00”非特殊票种 （默认值） “02”农产品收购 “06”抵扣通行费 “07”其它通行费 “08”成品油销售 “18”成品油专票负数发票 |
| `sum_price` | `SumPrice` | `*string` | JSON body | 未标注 | 合计金额（不含税），格式为2位小数。 当开红票时，该字段为负数。 开票成功时(create_result=success)必填 |
| `sum_tax` | `SumTax` | `*string` | JSON body | 未标注 | 合计税额，格式为2位小数。 当开红票时，该字段为0或负数。 开票成功时(create_result=success)必填 |

<a id="api-invoiceapplyresultsync"></a>

### InvoiceApplyResultSync — ISV向支付宝同步发票申请结果

- 官方接口：[`alipay.ebpp.invoice.apply.result.sync`](https://opendocs.alipay.com/apis/05vwzk)
- 路由：`POST /v3/alipay/ebpp/invoice/apply/result/sync`
- 请求类型：[InvoiceApplyResultSyncRequest](apply_request.go#L104)
- 响应类型：[InvoiceApplyResultSyncResponse](apply_response.go#L11)

ISV向支付宝同步发票申请结果

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | JSON body | 未标注 | 支付宝发起开票申请的id，该id具有唯一性，该字段由支付宝向税控发起申请的时候带过去，作为支付宝向税控开票申请的唯一标志 |
| `result` | `Result` | `*string` | JSON body | 未标注 | 支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果。<br> 取值如下：<br> SUCCESS:成功;FAIL:失败 |
| `result_code` | `ResultCode` | `*string` | JSON body | 未标注 | 结果码，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果进行详细说明的结果码。<br> 取值如下：<br> 成功(SUCCESS),<br> 参数不合法(PARAMETER_ILLEGAL),<br> 商户税控设备异常(MERCHANT_TAX_DEVICE_ERROR). |
| `result_msg` | `ResultMsg` | `*string` | JSON body | 未标注 | 结果描述，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果描述信息。 |
| `tax_apply_id` | `TaxApplyId` | `*string` | JSON body | 未标注 | 该字段是税控商或ISV收到支付宝开票请求后生成的申请id，由税控商或isv自己生成，该id具有唯一性<br> 当ISV接入时是按照tax_apply_id来查询发票信息时，该字段必填。 |

<a id="api-invoiceapplystatusnotify"></a>

### InvoiceApplyStatusNotify — 发票申请状态变更

- 官方接口：[`alipay.ebpp.invoice.apply.status.notify`](https://opendocs.alipay.com/apis/05uo9u)
- 路由：`PUT /v3/alipay/ebpp/invoice/apply/status/notify`
- 请求类型：[InvoiceApplyStatusNotifyRequest](apply_request.go#L118)
- 响应类型：[InvoiceApplyStatusNotifyResponse](apply_response.go#L17)

发票申请状态变更通知

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | JSON body | 未标注 | 发票申请ID，由阿里发票平台生成。字母或数字组成。 申请单的唯一标识，幂等字段。 |
| `apply_status` | `ApplyStatus` | `*string` | JSON body | 未标注 | 申请状态，可选值： applying: 申请中，初始状态； cancelled: 申请已取消、或商户已驳回； creating_inv: 商户已受理/开票中，待发票结果回传； inv_failed: 开票失败； inv_success: 开票成功； inv_part_success: 部分成功（拆单场景下存在。举例：发票申请拆单之后有10张票，其中有1张开票成功时，此时申请状态为inv_part_success，当10张票全部成功申请状态则为inv_success） |
| `invoice_uks` | `InvoiceUks` | [`*[]InvoiceUkDTO`](#model-invoiceukdto) | JSON body | 未标注 | 该申请下所有已开具成功的发票。 状态变更为 apply_status=inv_success 时该字段必传 |
| `message` | `Message` | `*string` | JSON body | 未标注 | 说明信息：驳回或失败原因 apply_status=inv_failed 或 apply_status=cancelled 时必传 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-invoiceitem"></a>

### InvoiceItem

类型定义：[InvoiceItem](apply_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 含税总金额 (等于sum_price和tax之和)， 单位：元，格式为2位小数，精度2位小数 |
| `item_name` | `ItemName` | `*string` | 未标注 | 发票项目名称（或商品名称） |
| `item_no` | `ItemNo` | `*string` | 未标注 | 发票项目编号（或商品编号） |
| `price` | `Price` | `*string` | 未标注 | 单价（不含税），格式为2位小数。最大支持6位小数，不足2位小数时需转化为2位小数格式。 折扣行此参数不能传 |
| `quantity` | `Quantity` | `*string` | 未标注 | 数量，须大于0，最多6位小数。 折扣行此参数不能传 |
| `row_type` | `RowType` | `*int64` | 未标注 | 发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。 比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0。 |
| `specification` | `Specification` | `*string` | 未标注 | 规格型号 |
| `sum_price` | `SumPrice` | `*string` | 未标注 | 不含税总金额，格式为2位小数 单位：元，精度2位小数 |
| `tax` | `Tax` | `*string` | 未标注 | 税额， 格式为2位小数 |
| `tax_rate` | `TaxRate` | `*string` | 未标注 | 税率。格式为2位小数，如：0.00, 0.03, 0.13等等 |
| `unit` | `Unit` | `*string` | 未标注 | 单位，折扣行不能传 |
| `zero_rate_flag` | `ZeroRateFlag` | `*string` | 未标注 | 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率 |

<a id="model-invoiceukdto"></a>

### InvoiceUkDTO

类型定义：[InvoiceUkDTO](apply_model.go#L33)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `invoice_code` | `InvoiceCode` | `*string` | 未标注 | 发票代码 |
| `invoice_no` | `InvoiceNo` | `*string` | 未标注 | 发票号码 |
