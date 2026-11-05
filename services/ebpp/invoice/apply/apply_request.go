// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package apply

// InvoiceApplyInvUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyInvUploadRequest struct {
	// AntiFakeCode 是发票防伪码/密码
	AntiFakeCode *string `json:"anti_fake_code,omitempty"`
	// ApplyId 是发票申请ID, 由阿里发票平台生成。 拆单开票后一笔申请可能对应多张发票。 source=apply时必填，其他source可为空
	ApplyId *string `json:"apply_id,omitempty"`
	// BizErrorCode 是开票失败错误码， 开票失败(create_result=fail)时必填
	BizErrorCode *string `json:"biz_error_code,omitempty"`
	// BizErrorMsg 是开票失败错误信息， 开票失败(create_result=fail)时必填
	BizErrorMsg *string `json:"biz_error_msg,omitempty"`
	// BusinessType 是抬头类型。可选值： 0：个人 1：企业
	BusinessType *int64 `json:"business_type,omitempty"`
	// CheckCode 是校验码
	CheckCode *string `json:"check_code,omitempty"`
	// CreateResult 是开票结果，枚举值： success: 发票开具成功； fail: 开票失败； source=async时必填，传实际的开票结果。其他source可不传，默认为success
	CreateResult *string `json:"create_result,omitempty"`
	// DeviceNo 是开票分机号/机器编号
	DeviceNo *string `json:"device_no,omitempty"`
	// InvoiceAmount 是合计含税金额（开票金额），格式为2位小数。 当开红票时，该字段为负数。 开票成功时(create_result=success)必填
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// InvoiceCode 是发票代码。 开票成功时(create_result=success)必填
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票日期，格式 yyyy-MM-dd. 开票成功时(create_result=success)必填
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceFileContent 是发票版式文件数据，base64化的字符串； 电子发票必填
	InvoiceFileContent *string `json:"invoice_file_content,omitempty"`
	// InvoiceFileType 是发票板式文件类型。可选值： PDF, OFD。 开电票且开票成功时(create_result=success)必填
	InvoiceFileType *string      `json:"invoice_file_type,omitempty"`
	InvoiceItems    *InvoiceItem `json:"invoice_items,omitempty"`
	// InvoiceKind 是开票发票类型 可选值： 0: 电票； 1：纸质普票； 2：纸质专票； 开票成功时(create_result=success)必填
	InvoiceKind *int64 `json:"invoice_kind,omitempty"`
	// InvoiceMemo 是发票备注，会显示在票面
	InvoiceMemo *string `json:"invoice_memo,omitempty"`
	// InvoiceNo 是发票号码。 开票成功时(create_result=success)必填
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceType 是发票(开票)类型，可选值： blue: 蓝票； red: 红票； 开票成功时(create_result=success)必填
	InvoiceType *string `json:"invoice_type,omitempty"`
	// LevyType 是征税方式，0普通征收，1减按征收，2差额征收
	LevyType *string `json:"levy_type,omitempty"`
	// NormalInvoiceCode 是原发票代码。 冲红且开票成功时(create_result=success)必填
	NormalInvoiceCode *string `json:"normal_invoice_code,omitempty"`
	// NormalInvoiceNo 是原发票号码。 冲红且开票成功时(create_result=success)必填
	NormalInvoiceNo *string `json:"normal_invoice_no,omitempty"`
	// PayeeAddress 是销方地址。 开票成功时(create_result=success)必填
	PayeeAddress *string `json:"payee_address,omitempty"`
	// PayeeBankAccountId 是销方银行账号
	PayeeBankAccountId *string `json:"payee_bank_account_id,omitempty"`
	// PayeeBankName 是销方开户行名称
	PayeeBankName *string `json:"payee_bank_name,omitempty"`
	// PayeeChecker 是复核人
	PayeeChecker *string `json:"payee_checker,omitempty"`
	// PayeeName 是销方名称。 开票成功时(create_result=success)必填
	PayeeName *string `json:"payee_name,omitempty"`
	// PayeeOperator 是开票人。 开票成功时(create_result=success)必填
	PayeeOperator *string `json:"payee_operator,omitempty"`
	// PayeePhone 是销方联系电话。 开票成功时(create_result=success)必填
	PayeePhone *string `json:"payee_phone,omitempty"`
	// PayeeReceiver 是收款人
	PayeeReceiver *string `json:"payee_receiver,omitempty"`
	// PayeeRegisterNo 是销方税务登记证号。由大写字母或数字组成，长度要求15~20位。 开票成功时(create_result=success)必填
	PayeeRegisterNo *string `json:"payee_register_no,omitempty"`
	// PayerAddress 是购方地址， 专票且开票成功时(create_result=success)必填
	PayerAddress *string `json:"payer_address,omitempty"`
	// PayerBankAccountId 是购方银行账号， 专票且开票成功时(create_result=success)必填
	PayerBankAccountId *string `json:"payer_bank_account_id,omitempty"`
	// PayerBankName 是购方开户行名称， 专票且开票成功时(create_result=success)必填
	PayerBankName *string `json:"payer_bank_name,omitempty"`
	// PayerEmail 是购方电子邮箱，需满足邮箱格式。 格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
	PayerEmail *string `json:"payer_email,omitempty"`
	// PayerName 是购方抬头。 开票成功时(create_result=success)必填
	PayerName *string `json:"payer_name,omitempty"`
	// PayerPhone 是购方联系电话， 专票且开票成功时(create_result=success)必填
	PayerPhone *string `json:"payer_phone,omitempty"`
	// PayerRegisterNo 是购方税务登记证号，由大写字母或数字组成，长度要求15~20位。 开企业抬头时必填， 专票且开票成功时(create_result=success)必填
	PayerRegisterNo *string `json:"payer_register_no,omitempty"`
	// PlatformCode 是业务平台Code, 由发票中台分配。 用于标识需交付发票的业务平台，source=upload时必填，其他source可为空
	PlatformCode *string `json:"platform_code,omitempty"`
	// PlatformTid 是业务平台发票申请对应的订单号。 用于区分业务平台订单号，source=upload时必填，其他source可为空
	PlatformTid *string `json:"platform_tid,omitempty"`
	// PlatformUserId 是业务平台发票申请对应的订单号。 用于区分业务平台订单号，source=upload时必填，其他source可为空
	PlatformUserId *string `json:"platform_user_id,omitempty"`
	// QrCode 是二维码
	QrCode *string `json:"qr_code,omitempty"`
	// ReceiveMobile 是购方手机号码，用于收票
	ReceiveMobile *string `json:"receive_mobile,omitempty"`
	// RedNoticeNo 是红字通知单号； 专票冲红且开票成功时(create_result=success)必填
	RedNoticeNo *string `json:"red_notice_no,omitempty"`
	// SerialNo 是开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。 source=async时必填，其他source可为空
	SerialNo *string `json:"serial_no,omitempty"`
	// Source 是发票来源，可选值： apply: 间连开票模式，ISV回传商家基于申请开具的发票； async：直连开票模式，ISV回传开票结果； upload：直接回传，进行归集与交付的发票
	Source *string `json:"source,omitempty"`
	// SpecialFlag 是特殊票种标识，可选值： “00”非特殊票种 （默认值） “02”农产品收购 “06”抵扣通行费 “07”其它通行费 “08”成品油销售 “18”成品油专票负数发票
	SpecialFlag *string `json:"special_flag,omitempty"`
	// SumPrice 是合计金额（不含税），格式为2位小数。 当开红票时，该字段为负数。 开票成功时(create_result=success)必填
	SumPrice *string `json:"sum_price,omitempty"`
	// SumTax 是合计税额，格式为2位小数。 当开红票时，该字段为0或负数。 开票成功时(create_result=success)必填
	SumTax *string `json:"sum_tax,omitempty"`
}

// InvoiceApplyResultSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyResultSyncRequest struct {
	// ApplyId 是支付宝发起开票申请的id，该id具有唯一性，该字段由支付宝向税控发起申请的时候带过去，作为支付宝向税控开票申请的唯一标志
	ApplyId *string `json:"apply_id,omitempty"`
	// Result 是支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果。 	取值如下： 	SUCCESS:成功;FAIL:失败
	Result *string `json:"result,omitempty"`
	// ResultCode 是结果码，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果进行详细说明的结果码。 	取值如下： 	成功(SUCCESS), 	参数不合法(PARAMETER_ILLEGAL), 	商户税控设备异常(MERCHANT_TAX_DEVICE_ERROR).
	ResultCode *string `json:"result_code,omitempty"`
	// ResultMsg 是结果描述，支付宝向税控商或ISV发起发票申请后，对应这笔申请的发票开具结果描述信息。
	ResultMsg *string `json:"result_msg,omitempty"`
	// TaxApplyId 是该字段是税控商或ISV收到支付宝开票请求后生成的申请id，由税控商或isv自己生成，该id具有唯一性 	当ISV接入时是按照tax_apply_id来查询发票信息时，该字段必填。
	TaxApplyId *string `json:"tax_apply_id,omitempty"`
}

// InvoiceApplyStatusNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyStatusNotifyRequest struct {
	// ApplyId 是发票申请ID，由阿里发票平台生成。字母或数字组成。	申请单的唯一标识，幂等字段。
	ApplyId *string `json:"apply_id,omitempty"`
	// ApplyStatus 是申请状态，可选值：	applying: 申请中，初始状态；	cancelled: 申请已取消、或商户已驳回；	creating_inv: 商户已受理/开票中，待发票结果回传；	inv_failed: 开票失败；	inv_success: 开票成功；	inv_part_success: 部分成功（拆单场景下存在。举例：发票申请拆单之后有10张票，其中有1张开票成功时，此时申请状态为inv_part_success，当10张票全部成功申请状态则为inv_success）
	ApplyStatus *string `json:"apply_status,omitempty"`
	// InvoiceUks 是该申请下所有已开具成功的发票。	状态变更为 apply_status=inv_success 时该字段必传
	InvoiceUks *[]InvoiceUkDTO `json:"invoice_uks,omitempty"`
	// Message 是说明信息：驳回或失败原因	apply_status=inv_failed 或 apply_status=cancelled 时必传
	Message *string `json:"message,omitempty"`
}
