// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sycn

// InvoiceItemContent 是支付宝官方 V3 OpenAPI 模型。
type InvoiceItemContent struct {
	// ItemAmount 是价税合计。(等于sumPrice和tax之和)
	ItemAmount *string `json:"item_amount,omitempty"`
	// ItemName 是发票项目名称（或商品名称）
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是商品编号
	ItemNo *string `json:"item_no,omitempty"`
	// ItemPrice 是单价，格式：100.00。新版电子发票，折扣行此参数不能传，非折扣行必传
	ItemPrice *string `json:"item_price,omitempty"`
	// ItemQuantity 是数量。新版电子发票，折扣行此参数不能传，非折扣行必传
	ItemQuantity *int64 `json:"item_quantity,omitempty"`
	// ItemSumPrice 是单项总价，格式：100.00
	ItemSumPrice *string `json:"item_sum_price,omitempty"`
	// ItemTaxPrice 是税额
	ItemTaxPrice *string `json:"item_tax_price,omitempty"`
	// ItemTaxRate 是税率
	ItemTaxRate *string `json:"item_tax_rate,omitempty"`
	// ItemUnit 是台
	ItemUnit *string `json:"item_unit,omitempty"`
	// RowType 是发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0
	RowType *int64 `json:"row_type,omitempty"`
}

// InvoiceModelContent 是支付宝官方 V3 OpenAPI 模型。
type InvoiceModelContent struct {
	// ApplyId 是支付宝端的申请id。如果在开票过程中，是通过支付宝提交的申请到机构端，支付宝会带上开票申请在支付宝生成的申请id，机构在回传发票的时候只需要回传这个申请id，不用获取用户的uid，支付宝可以根据申请id将发票归集到对应的用户名下
	ApplyId *string `json:"apply_id,omitempty"`
	// ExtendFields 是key=value，每组键值对以回车分割
	ExtendFields *string `json:"extend_fields,omitempty"`
	// FileDownloadType 是下载的发票文件类型，可选值：	pdf（发票原文件）	ofd（发票原文件）	jpg（发票原文件缩略图）
	FileDownloadType *string `json:"file_download_type,omitempty"`
	// FileDownloadUrl 是发票原文件下载地址	1.当tax_type=PLAIN、ALL_ELECTRONIC_GENERAL或ALL_ELECTRONIC_SPECIAL时，	file_download_url必传	且file_download_type取值范围为pdf或ofd；	2.当tax_type=SPECIAL时，	file_download_url必传	file_download_type可以传入pdf，ofd，jpg	3.当其他票种时，file_download_url可以不传
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// FinancialElectronicType 是财政电子票据子类型，当tax_type=FINANCIAL_ELECTRONIC_BILL时要求必填	可选值如下：	01:非税收入通用票据 	02:非税收入专用票据	03:非税收入一般缴款书	04:资金往来结算票据	05:公益事业捐赠票据	06:医疗收费票据	07:社会团体会费票据	08:社会保险基金票据	09:工会经费收入票据	99:其他财政票据
	FinancialElectronicType *string `json:"financial_electronic_type,omitempty"`
	// InvoiceAmount 是发票金额，大于0且精确到小数点两位，以元为单位 	需要传入税价合计金额
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// InvoiceCode 是发票代码，为国税局生成的唯一值，全电票时为空，其他情况不可为空
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceContent 是发票内容项
	InvoiceContent *[]InvoiceItemContent `json:"invoice_content,omitempty"`
	// InvoiceDate 是发票日期，用户填写，目前精确到日
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceFakeCode 是发票防伪码
	InvoiceFakeCode *string `json:"invoice_fake_code,omitempty"`
	// InvoiceFileData 是原始发票PDF/OFD文件流
	InvoiceFileData *string `json:"invoice_file_data,omitempty"`
	// InvoiceImgUrl 是发票原始文件jpg文件地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// InvoiceNo 是发票号码，国税局生成的唯一号码，不可为空串； 	使用时请注意，invoice_no+invoice_code唯一，不能重复
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceOperator 是发票开具操作人
	InvoiceOperator *string       `json:"invoice_operator,omitempty"`
	InvoiceTitle    *InvoiceTitle `json:"invoice_title,omitempty"`
	// InvoiceType 是发票类型，按照可选值只传入英文部分，该字段严格要求大小写 	可选值: 	blue（蓝票） 	red（红票）
	InvoiceType *string `json:"invoice_type,omitempty"`
	// OpenId 是支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。
	OpenId *string `json:"open_id,omitempty"`
	// OriginalBlueInvoiceCode 是仅用于同步红票，原始蓝票发票代码，同步红票时必传（全电票时为空）
	OriginalBlueInvoiceCode *string `json:"original_blue_invoice_code,omitempty"`
	// OriginalBlueInvoiceNo 是仅用于同步红票，原始蓝票发票号码，同步红票时必传
	OriginalBlueInvoiceNo *string `json:"original_blue_invoice_no,omitempty"`
	// OutBizNo 是商户交易流水号，不可为空串; 	传入红票时请注意，此字段的值要和蓝票保持一致
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// OutInvoiceId 是商户唯一开票申请业务流水号，同一个isv下不能重复
	OutInvoiceId *string `json:"out_invoice_id,omitempty"`
	// RegisterAddress 是开票单位地址
	RegisterAddress *string `json:"register_address,omitempty"`
	// RegisterBankAccount 是开票单位开户行账号
	RegisterBankAccount *string `json:"register_bank_account,omitempty"`
	// RegisterBankName 是开票单位开户行名称
	RegisterBankName *string `json:"register_bank_name,omitempty"`
	// RegisterName 是开票单位
	RegisterName *string `json:"register_name,omitempty"`
	// RegisterNo 是纳税人识别号，不可为空串
	RegisterNo *string `json:"register_no,omitempty"`
	// RegisterPhoneNo 是开票人电话，支持座机和手机两种格式
	RegisterPhoneNo *string `json:"register_phone_no,omitempty"`
	// SumAmount 是价税合计
	SumAmount *string `json:"sum_amount,omitempty"`
	// TaxAmount 是税额
	TaxAmount *string `json:"tax_amount,omitempty"`
	// TaxType 是税种	可选值：	PLAIN：增值税电子普通发票	ALL_ELECTRONIC_GENERAL：电子发票（普通发票）	ALL_ELECTRONIC_SPECIAL：电子发票（增值税专用发票）	SPECIAL：增值税专用发票	PLAIN_INVOICE:增值税普通发票	FINANCIAL_ELECTRONIC_BILL:财政电子票据
	TaxType *string `json:"tax_type,omitempty"`
	// UserId 是支付宝用户id，支付宝端的申请id存在的时候也不需要传，其他情况下，当同步的是蓝票时，必传，红票时不需传。
	UserId *string `json:"user_id,omitempty"`
}

// InvoiceTitle 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTitle struct {
	// IsDefault 是是否为用户设置默认抬头 	字段值包括两种情况： 	false（非默认） 	true（默认抬头）
	IsDefault *bool `json:"is_default,omitempty"`
	// LogonId 是支付宝用户登录名
	LogonId *string `json:"logon_id,omitempty"`
	// OpenBankAccount 是银行账号
	OpenBankAccount *string `json:"open_bank_account,omitempty"`
	// OpenBankName 是开户银行
	OpenBankName *string `json:"open_bank_name,omitempty"`
	// OpenId 是支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// TaxRegisterNo 是纳税人识别号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// TelePhoneNo 是用户私人手机号
	TelePhoneNo *string `json:"tele_phone_no,omitempty"`
	// TitleName 是抬头名称
	TitleName *string `json:"title_name,omitempty"`
	// TitleType 是抬头类型	字段值有两种情况抬:	PERSONAL（个人） 	CORPORATION（企业）
	TitleType *string `json:"title_type,omitempty"`
	// UserAddress 是地址
	UserAddress *string `json:"user_address,omitempty"`
	// UserEmail 是邮箱
	UserEmail *string `json:"user_email,omitempty"`
	// UserId 是支付宝用户id
	UserId *string `json:"user_id,omitempty"`
	// UserMobile 是电话号码
	UserMobile *string `json:"user_mobile,omitempty"`
}
