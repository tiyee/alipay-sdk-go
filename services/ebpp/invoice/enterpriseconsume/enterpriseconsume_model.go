// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseconsume

// CompleteVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type CompleteVoucherInfo struct {
	// AntiFakeCode 是防伪码
	AntiFakeCode *string `json:"anti_fake_code,omitempty"`
	// CheckSum 是校验码
	CheckSum *string `json:"check_sum,omitempty"`
	// DeviceNo 是开票机器码
	DeviceNo *string `json:"device_no,omitempty"`
	// FileDownloadUrl 是板式文件下载地址
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// FileType 是板式文件的类型
	FileType *string `json:"file_type,omitempty"`
	// InvoiceAmount 是发票金额（元）
	InvoiceAmount *int64 `json:"invoice_amount,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceContentList 是汇总发票详细信息(目前该字段为空)
	InvoiceContentList *[]InvoiceContentInfo `json:"invoice_content_list,omitempty"`
	// InvoiceDate 是发票日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceKind 是发票类型 0：增值税普通电子发票 1：增值税普通发票 2：增值税专用发票 3：增值税专用电子发票 4：财政电子发票
	InvoiceKind *int64 `json:"invoice_kind,omitempty"`
	// InvoiceMemo 是备注
	InvoiceMemo *string `json:"invoice_memo,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceSource 是发票来源，取值范围 SUMMARY_INVOICE_AUTO_OPEN 汇总开票系统自动开票，SUMMARY_INVOICE_MERCHANT_UPLOAD 汇总开票商户手动上传
	InvoiceSource *string `json:"invoice_source,omitempty"`
	// InvoiceTitle 是发票抬头
	InvoiceTitle *string `json:"invoice_title,omitempty"`
	// InvoiceType 是发票类型(红、蓝)
	InvoiceType *string `json:"invoice_type,omitempty"`
	// PayeeAddress 是销售方地址
	PayeeAddress *string `json:"payee_address,omitempty"`
	// PayeeBankAccount 是销售方银行账号
	PayeeBankAccount *string `json:"payee_bank_account,omitempty"`
	// PayeeBankName 是销售方银行名称
	PayeeBankName *string `json:"payee_bank_name,omitempty"`
	// PayeeChecker 是复核人
	PayeeChecker *string `json:"payee_checker,omitempty"`
	// PayeeMobile 是销售方电话
	PayeeMobile *string `json:"payee_mobile,omitempty"`
	// PayeeName 是销售方名称
	PayeeName *string `json:"payee_name,omitempty"`
	// PayeeOperator 是开票人
	PayeeOperator *string `json:"payee_operator,omitempty"`
	// PayeeReceiver 是收款人
	PayeeReceiver *string `json:"payee_receiver,omitempty"`
	// PayeeRegisterNo 是销售方税号
	PayeeRegisterNo *string `json:"payee_register_no,omitempty"`
	// PayerAddress 是购买方地址
	PayerAddress *string `json:"payer_address,omitempty"`
	// PayerBankAccount 是购买方银行账号
	PayerBankAccount *string `json:"payer_bank_account,omitempty"`
	// PayerBankName 是购买方银行名称
	PayerBankName *string `json:"payer_bank_name,omitempty"`
	// PayerMobile 是购买方电话
	PayerMobile *string `json:"payer_mobile,omitempty"`
	// PayerName 是购买方名称
	PayerName *string `json:"payer_name,omitempty"`
	// PayerRegisterNo 是购买方税号
	PayerRegisterNo *string `json:"payer_register_no,omitempty"`
	// SumPrice 是不含税金额（元）
	SumPrice *int64 `json:"sum_price,omitempty"`
	// SumTax 是合计税额（元）
	SumTax *int64 `json:"sum_tax,omitempty"`
}

// EnterpriseOpenRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseOpenRuleInfo struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// EnterpriseOpenRuleRecordInfoList 是当前生效和下次生效的开票规则记录
	EnterpriseOpenRuleRecordInfoList *[]EnterpriseOpenRuleRecordInfo `json:"enterprise_open_rule_record_info_list,omitempty"`
	// EnterpriseOpenRuleRelationInfoList 是开票规则绑定关系
	EnterpriseOpenRuleRelationInfoList *[]EnterpriseOpenRuleRelationInfo `json:"enterprise_open_rule_relation_info_list,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// InvoiceRuleName 是开票规则名称
	InvoiceRuleName *string `json:"invoice_rule_name,omitempty"`
	// OwnerId 是所有者ID（企业情况下即为企业ID）。
	OwnerId *string `json:"owner_id,omitempty"`
	// SellerType 是"销方类型：TP开票、商户开票、商户优先枚举定义：MERCHANT（商户）、TP（TP）"
	SellerType *string `json:"seller_type,omitempty"`
}

// EnterpriseOpenRuleRecordInfo 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseOpenRuleRecordInfo struct {
	// BillMonthDay 是开票规则账单日
	BillMonthDay *int64 `json:"bill_month_day,omitempty"`
	// BillScope 是开票账单范围
	BillScope *string `json:"bill_scope,omitempty"`
	// CombinedPayMode 是组合支付模式
	CombinedPayMode *string `json:"combined_pay_mode,omitempty"`
	// DefaultInvoiceKind 是默认开票类型
	DefaultInvoiceKind *string `json:"default_invoice_kind,omitempty"`
	// EffectiveStart 是开票规则生效日期
	EffectiveStart *string `json:"effective_start,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// InvoiceRemarkValueRule 是开票规则的发票备注取值规则。NONE表示不使用额度发放说明；QUOTA_ISSUE_DESC表示将因公付额度发放说明作为发票备注。历史规则未配置时按NONE理解。
	InvoiceRemarkValueRule *string `json:"invoice_remark_value_rule,omitempty"`
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// InvoiceRuleRecordId 是开票规则记录ID
	InvoiceRuleRecordId *string `json:"invoice_rule_record_id,omitempty"`
	// InvoiceTitleId 是发票抬头ID
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// OpenApplyer 是开票申请方
	OpenApplyer *string `json:"open_applyer,omitempty"`
	// OpenMode 是开票模式
	OpenMode *string `json:"open_mode,omitempty"`
	// OpenType 是开票申请类型
	OpenType *string `json:"open_type,omitempty"`
	// OwnerId 是所有者ID（企业情况下即为企业ID）。
	OwnerId *string `json:"owner_id,omitempty"`
	// Tag 是开票规则标记
	Tag *string `json:"tag,omitempty"`
}

// EnterpriseOpenRuleRelationInfo 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseOpenRuleRelationInfo struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// OwnerId 是所有者ID（企业情况下即为企业ID）。
	OwnerId *string `json:"owner_id,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
}

// ExpenseConsumeInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseConsumeInfo struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// ActualAccountNumber 是实际出资企业支付宝账号
	ActualAccountNumber *string `json:"actual_account_number,omitempty"`
	// BatchId 是汇总批次ID(废弃)
	BatchId *string `json:"batch_id,omitempty"`
	// BillNo 是支付流水号
	BillNo *string `json:"bill_no,omitempty"`
	// BillType 是账单类型（CONSUME-消费账单、REFUND-退款、TRANSFER-转账）
	BillType *string `json:"bill_type,omitempty"`
	// CategoryName 是账单行业分类（交通出行，餐饮美食，商业服务，日用百货等等）
	CategoryName *string `json:"category_name,omitempty"`
	// ConsumeAmount 是账单金额（单位：元）
	ConsumeAmount *string `json:"consume_amount,omitempty"`
	// ConsumeDate 是交易日期
	ConsumeDate *string `json:"consume_date,omitempty"`
	// ConsumeTitle 是账单展示标题
	ConsumeTitle *string `json:"consume_title,omitempty"`
	// EmployeeId 是所属员工支付宝uid
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeOpenId 是所属员工支付宝uid
	EmployeeOpenId *string `json:"employee_open_id,omitempty"`
	// MerchantId 是商户id
	MerchantId *string `json:"merchant_id,omitempty"`
	// MshopId 是门店id
	MshopId *string `json:"mshop_id,omitempty"`
	// OpenModel 是开票模式（汇总开：ENTERPRISE_AUTO_BATCH）
	OpenModel *string `json:"open_model,omitempty"`
	// OriginalVoucherId 是退款交易关联正交易单号(对应正交易的bill_no)
	OriginalVoucherId *string `json:"original_voucher_id,omitempty"`
	// OutBizNo 是外部流水号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PPayAmount 是企业支付金额（单位：元）
	PPayAmount *string `json:"p_pay_amount,omitempty"`
	// PayeeName 是商家名称
	PayeeName *string `json:"payee_name,omitempty"`
	// ProjectId 是费控规则中的项目ID
	ProjectId *string `json:"project_id,omitempty"`
	// ProjiectId 是支付使用的项目号（成本中心）此字段废弃
	ProjiectId *string `json:"projiect_id,omitempty"`
	// RefundAmount 是消费账单的退款金额（元）
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundStatus 是消费账单的退款状态
	RefundStatus *string `json:"refund_status,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
	// StoreId 是外部门店id
	StoreId *string `json:"store_id,omitempty"`
	// SummaryId 是批次id
	SummaryId *string `json:"summary_id,omitempty"`
	// TpSign 是出资类型：	PERSONAL:个人垫付	ENTERPRISE:企业垫资	TP:三方合作伙伴垫资	ACCOUNTING:记账
	TpSign *string `json:"tp_sign,omitempty"`
	// VoucherId 是凭证ID
	VoucherId *string `json:"voucher_id,omitempty"`
}

// ExpenseCtrlConsumeInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseCtrlConsumeInfo struct {
	ExpenseConsumeInfo *ExpenseConsumeInfo `json:"expense_consume_info,omitempty"`
	// RelatedEnterpriseConsumeList 是关联凑票账单列表（开票限额，多次消费后合并开票）
	RelatedEnterpriseConsumeList *[]ExpenseConsumeInfo `json:"related_enterprise_consume_list,omitempty"`
	// RelatedEnterpriseInvoiceList 是关联发票凭证对象
	RelatedEnterpriseInvoiceList *[]ExpenseInvoiceInfo `json:"related_enterprise_invoice_list,omitempty"`
}

// ExpenseInvoiceInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseInvoiceInfo struct {
	// EmployeeId 是员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeOpenId 是员工ID
	EmployeeOpenId    *string            `json:"employee_open_id,omitempty"`
	InvoiceOutputInfo *InvoiceOutputInfo `json:"invoice_output_info,omitempty"`
	OCRNormalScanInfo *OCRNormalScanInfo `json:"ocr_normal_scan_info,omitempty"`
	OCRPlaneScanInfo  *OCRPlaneScanInfo  `json:"ocr_plane_scan_info,omitempty"`
	OCRTaxiScanInfo   *OCRTaxiScanInfo   `json:"ocr_taxi_scan_info,omitempty"`
	OCRTrainScanInfo  *OCRTrainScanInfo  `json:"ocr_train_scan_info,omitempty"`
	VoucherFileInfo   *VoucherFileInfo   `json:"voucher_file_info,omitempty"`
	// VoucherId 是凭证ID
	VoucherId *string `json:"voucher_id,omitempty"`
}

// ExpenseVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseVoucherInfo struct {
	// AccountId 是企业支付宝用户id(对应凭证ownerid)
	AccountId *string `json:"account_id,omitempty"`
	// ConsumptionDate 是消费时间
	ConsumptionDate *string `json:"consumption_date,omitempty"`
	// EmployeeId 是员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeOpenId 是员工ID
	EmployeeOpenId *string `json:"employee_open_id,omitempty"`
	// Extension 是扩展预留
	Extension *string `json:"extension,omitempty"`
	// GmtCreate 是数据创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是数据更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Industry 是行业属性值（从支付宝主账单复制）
	Industry *string `json:"industry,omitempty"`
	// IsOffSet 是是否冲抵凭证：0 否（蓝票）；1 是(如:红票)
	IsOffSet *string `json:"is_off_set,omitempty"`
	// Medium 是凭证介质：纸or电子(PAPER,ELECTRON)
	Medium *string `json:"medium,omitempty"`
	// OuterSourceId 是外部唯一ID（和凭证类型有关，如果发票类型为发票号码+发票编码，如果是账单，则为账单号）
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// ParentType 是凭证父类型（INVOICE-发票，TRAVEL-出行凭证，HTL_ORDER-酒店水单，CONSUME-账单，财政票夹）
	ParentType *string `json:"parent_type,omitempty"`
	// VoucherAmount 是交易金额（元）
	VoucherAmount *string `json:"voucher_amount,omitempty"`
	// VoucherDate 是凭证创建时间
	VoucherDate *string `json:"voucher_date,omitempty"`
	// VoucherId 是凭证ID
	VoucherId *string `json:"voucher_id,omitempty"`
	// VoucherNo 是一类凭证，唯一号码，有必须传，没有可不传。	发票、账单必须有	如部分餐饮小票，没有唯一号
	VoucherNo *string `json:"voucher_no,omitempty"`
	// VoucherState 是凭证状态(0-无效，1 有效)
	VoucherState *string `json:"voucher_state,omitempty"`
	// VoucherType 是凭证类型	（FINANCIAL_BILLS-财政电子票	MEDICAL_INVOICE-医疗票据	DONATION_INVOICE-公益捐赠电子票据	SETTLEMENT_INVOICE-往来结算票据	DUES_INVOICE-社会团体会费收据	INVOICE-增值税票	PLAIN-增值税电子普通发票	SPECIAL-增值税专用发票	PLAIN_INVOICE-增值税普通发票	PAPER_INVOICE-增值税普通发票(卷式)	CONSUME-账单	ALIPAY_CONSUME-支付宝账单	TRAVEL-出行凭证	TAXI-出租车票	BUS_TICKET-汽车票	TRAIN-火车票	TOLL-过路费	BOARD_PASS-登机牌	HTL_ORDER-酒店水单	MEMO-酒店水单	OTHERS-杂票	FIXED-定额发票	LIST-小票	COMMON_INVOICE-通用发票）
	VoucherType *string `json:"voucher_type,omitempty"`
}

// InvoiceContentInfo 是支付宝官方 V3 OpenAPI 模型。
type InvoiceContentInfo struct {
	// Amount 是含税金额（元）
	Amount *string `json:"amount,omitempty"`
	// ItemName 是发票项目名称/货物名称
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是商品编码/税收分类编码
	ItemNo *string `json:"item_no,omitempty"`
	// ItemUnit 是单位
	ItemUnit *string `json:"item_unit,omitempty"`
	// Price 是单价（元）
	Price *string `json:"price,omitempty"`
	// Quantity 是商品数量
	Quantity *string `json:"quantity,omitempty"`
	// RowType 是发票行性质。0表示正常行，1表示折扣行，2表示被折扣行
	RowType *string `json:"row_type,omitempty"`
	// Specification 是规格型号
	Specification *string `json:"specification,omitempty"`
	// SumPrice 是不含税金额（元）
	SumPrice *string `json:"sum_price,omitempty"`
	// Tax 是税额
	Tax *string `json:"tax,omitempty"`
	// TaxRate 是税率
	TaxRate *string `json:"tax_rate,omitempty"`
	// ZeroRateFlag 是税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率
	ZeroRateFlag *string `json:"zero_rate_flag,omitempty"`
}

// InvoiceItemOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceItemOpen struct {
	// ItemExTaxAmount 是明细行不含税金额，单位元，保留两位小数
	ItemExTaxAmount *string `json:"item_ex_tax_amount,omitempty"`
	// ItemName 是开票项目：货物或应税劳务、服务名称
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是国税局制定的商品税收编码，必须是最末级
	ItemNo *string `json:"item_no,omitempty"`
	// ItemQuantity 是数量；	1.当row_type=0或2且item_unit_price为空，可空；	2.可精确到小数点后6位
	ItemQuantity *int64 `json:"item_quantity,omitempty"`
	// ItemSpec 是规格型号
	ItemSpec *string `json:"item_spec,omitempty"`
	// ItemSumAmount 是明细行价税合计，单位元，必须保证item_sum_amount=item_ex_tax_amount+item_tax_amount。
	ItemSumAmount *string `json:"item_sum_amount,omitempty"`
	// ItemTaxAmount 是明细行税额，单位元，保留两位小数，无税或者免税情况下输入：0.00。
	ItemTaxAmount *string `json:"item_tax_amount,omitempty"`
	// ItemTaxRate 是明细行税率，无税或者免税情况下输入：0.00。
	ItemTaxRate *string `json:"item_tax_rate,omitempty"`
	// ItemUnit 是单位
	ItemUnit *string `json:"item_unit,omitempty"`
	// ItemUnitPrice 是不含税单价（元）;	1.当row_type=0或2且item_quantity为空，可空	2.可精确到小数点后8位；
	ItemUnitPrice *string `json:"item_unit_price,omitempty"`
	// RowType 是发票行性质：0表示正常行，1表示折扣行，2表示被折扣行。
	RowType *string `json:"row_type,omitempty"`
}

// InvoiceOutputInfo 是支付宝官方 V3 OpenAPI 模型。
type InvoiceOutputInfo struct {
	// CheckCode 是发票防伪码，验证码
	CheckCode *string `json:"check_code,omitempty"`
	// Checker 是票面上复核人信息。
	Checker *string `json:"checker,omitempty"`
	// Clerk 是票面上开票员信息。
	Clerk *string `json:"clerk,omitempty"`
	// ExTaxAmount 是不含税金额，票面上的不含税金额，单位（元），保留两位小数。
	ExTaxAmount *string `json:"ex_tax_amount,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceContent 是发票明细项。
	InvoiceContent *[]InvoiceItemOpen `json:"invoice_content,omitempty"`
	// InvoiceDate 是开票日期，格式如：YYYY-MM-DD。
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceImgUrl 是发票预览图片地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// InvoiceKind 是标识发票的类型，PLAIN:增值税电子普通发票;SPECIAL:增值税专用发票;PLAIN_INVOICE:增值税普通发票;SALSE_INVOICE:机动车销售统一发票 ;PAPER_INVOICE:增值税普通发票(卷式)。
	InvoiceKind *string `json:"invoice_kind,omitempty"`
	// InvoiceMemo 是票面上备注信息。
	InvoiceMemo *string `json:"invoice_memo,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo    *string           `json:"invoice_no,omitempty"`
	InvoiceTitle *InvoiceTitleOpen `json:"invoice_title,omitempty"`
	// InvoiceType 是标识是红票还是蓝票，RED:红票;BLUE:蓝票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// OriBlueInvCode 是红票情况下，必须填入原始蓝票的发票代码。
	OriBlueInvCode *string `json:"ori_blue_inv_code,omitempty"`
	// OriBlueInvNo 是红票情况下，必须填入原始蓝票的发票号码。
	OriBlueInvNo *string `json:"ori_blue_inv_no,omitempty"`
	// Payee 是票面上收款人信息。
	Payee *string `json:"payee,omitempty"`
	// PayeeAddressTel 是票面上销售方地址、电话信息。
	PayeeAddressTel *string `json:"payee_address_tel,omitempty"`
	// PayeeBankNameAccount 是票面上销售方开户行及账号。
	PayeeBankNameAccount *string `json:"payee_bank_name_account,omitempty"`
	// PayeeName 是票面上销售方企业名称。
	PayeeName *string `json:"payee_name,omitempty"`
	// PayeeRegisterNo 是票面上销售方税号信息。
	PayeeRegisterNo *string `json:"payee_register_no,omitempty"`
	// SumAmount 是价税合计，票面上的总开票金额，单位（元），保留两位小数，sum_amount=ex_tax_amount+tax_amount。
	SumAmount *string `json:"sum_amount,omitempty"`
	// TaxAmount 是合计税额，票面上的合计税额，单位（元），保留两位小数。
	TaxAmount *string `json:"tax_amount,omitempty"`
}

// InvoiceTitleOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTitleOpen struct {
	// PayerAddressTel 是票面上的购买方地址、电话
	PayerAddressTel *string `json:"payer_address_tel,omitempty"`
	// PayerBankNameAccount 是票面上的购买方开户行及账户
	PayerBankNameAccount *string `json:"payer_bank_name_account,omitempty"`
	// PayerRegisterNo 是票面上的购买方纳税人识别号
	PayerRegisterNo *string `json:"payer_register_no,omitempty"`
	// TitleName 是票面上的购买方名称
	TitleName *string `json:"title_name,omitempty"`
}

// OCRNormalScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRNormalScanInfo struct {
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票时间
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
}

// OCRPlaneScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRPlaneScanInfo struct {
	// Destination 是目的地
	Destination *string `json:"destination,omitempty"`
	// FlightNo 是航班号
	FlightNo *string `json:"flight_no,omitempty"`
	// InvoiceDate 是乘机日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// Origin 是出发地
	Origin *string `json:"origin,omitempty"`
	// Passenger 是乘客
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
	// SeatClass 是飞机舱位
	SeatClass *string `json:"seat_class,omitempty"`
}

// OCRTaxiScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRTaxiScanInfo struct {
	// GetOffTime 是下车时间
	GetOffTime *string `json:"get_off_time,omitempty"`
	// GetOnTime 是上车时间
	GetOnTime *string `json:"get_on_time,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票时间
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// Passenger 是乘客人
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// TravelDist 是乘车距离
	TravelDist *string `json:"travel_dist,omitempty"`
}

// OCRTrainScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRTrainScanInfo struct {
	// Destination 是目的地
	Destination *string `json:"destination,omitempty"`
	// InvoiceDate 是乘车日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// Origin 是出发地
	Origin *string `json:"origin,omitempty"`
	// Passenger 是乘客
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
	// SeatClass 是座次
	SeatClass *string `json:"seat_class,omitempty"`
	// TrainNo 是车次
	TrainNo *string `json:"train_no,omitempty"`
}

// VoucherFileInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherFileInfo struct {
	// FileDownloadUrl 是发票文件下载
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// FileName 是文件名
	FileName *string `json:"file_name,omitempty"`
	// FileType 是文件类型
	FileType *string `json:"file_type,omitempty"`
}
