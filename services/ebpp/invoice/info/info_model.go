// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// EinvTrade 是支付宝官方 V3 OpenAPI 模型。
type EinvTrade struct {
	// BillNo 是订单编号
	BillNo *string `json:"bill_no,omitempty"`
	// BillTime 是下单时间
	BillTime *string `json:"bill_time,omitempty"`
	// CityName 是商户所在城市(经营地址)
	CityName *string `json:"city_name,omitempty"`
	// DetailJson 是账单明细信息，酒店水单信息，行程单信息，餐饮小票信息
	DetailJson *string `json:"detail_json,omitempty"`
	// DownloadUrl 是账单明细信息，酒店水单，行程单，餐饮小票等pdf原件链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// ExtendMap 是扩展参数		不同组的k-v通过换行符区分
	ExtendMap *string `json:"extend_map,omitempty"`
	// MerchantName 是商家名称（显示名称，非企业名称，餐饮店、酒店、打车平台名称）
	MerchantName *string `json:"merchant_name,omitempty"`
	// OutJson 是透传字段，不做处理，用于isv向后续报销税控方传递特殊信息标记
	OutJson *string `json:"out_json,omitempty"`
	// PayeeName 是销方名称
	PayeeName *string `json:"payee_name,omitempty"`
	// PaymentAmount 是支付金额，单位（元），	对应账单中的交易金额
	PaymentAmount *string `json:"payment_amount,omitempty"`
	// PaymentTime 是支付时间		对应账单中的账单日期
	PaymentTime *string `json:"payment_time,omitempty"`
	// Souce 是交易类型来源 需要按照枚举映射
	Souce *string `json:"souce,omitempty"`
	// TradeType 是交易类型/账单分类
	TradeType *string `json:"trade_type,omitempty"`
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

// InvoiceSendOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceSendOpen struct {
	// ApplyId 是支付宝端的申请id。如果在开票过程中，是通过支付宝提交的申请到机构端，支付宝会带上开票申请在支付宝生成的申请id，机构在回传发票的时候只需要回传这个申请id，不用获取用户的uid，支付宝可以根据申请id将发票归集到对应的用户名下
	ApplyId *string `json:"apply_id,omitempty"`
	// CheckCode 是票面上的校验码信息。
	CheckCode *string `json:"check_code,omitempty"`
	// Checker 是票面上复核人信息。
	Checker *string `json:"checker,omitempty"`
	// Clerk 是票面上开票员信息。
	Clerk *string `json:"clerk,omitempty"`
	// ExTaxAmount 是不含税金额，票面上的不含税金额，单位（元），保留两位小数。
	ExTaxAmount *string `json:"ex_tax_amount,omitempty"`
	// ExtendFields 是预留的扩展字段，格式如：key1=value1\nkey2=value2\nkey3=value3，字段之间以\n分隔。
	ExtendFields *string `json:"extend_fields,omitempty"`
	// FileDownloadType 是下载版式文件类型，需要回传文件下载链接的情况下必传
	FileDownloadType *string `json:"file_download_type,omitempty"`
	// FileDownloadUrl 是发票版式文件下载地址。
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// FinancialElectronicType 是财政电子票据子类型
	FinancialElectronicType *string `json:"financial_electronic_type,omitempty"`
	// InvoiceCode 是发票代码，该字段与invoice_no构成一张发票的唯一标识。
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceContent 是发票明细项。
	InvoiceContent *[]InvoiceItemOpen `json:"invoice_content,omitempty"`
	// InvoiceDate 是开票日期，格式如：YYYY-MM-DD。
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceKind 是标识发票的类型，	PLAIN:增值税电子普通发票;	SPECIAL:增值税专用发票;	ALL_ELECTRONIC_GENERAL:电子发票（普通发票）;	ALL_ELECTRONIC_SPECIAL:电子发票（增值税专用发票）;	PLAIN_INVOICE:增值税普通发票;	FINANCIAL_ELECTRONIC_BILL:财政电子票据;
	InvoiceKind *string `json:"invoice_kind,omitempty"`
	// InvoiceMemo 是票面上备注信息。
	InvoiceMemo *string `json:"invoice_memo,omitempty"`
	// InvoiceNo 是发票号码，该字段与invoice_code构成一张发票的唯一标识。
	InvoiceNo    *string           `json:"invoice_no,omitempty"`
	InvoiceTitle *InvoiceTitleOpen `json:"invoice_title,omitempty"`
	// InvoiceType 是标识是红票还是蓝票，RED:红票;BLUE:蓝票。
	InvoiceType *string `json:"invoice_type,omitempty"`
	// OpenId 是支付宝用户userId，当发送红字发票时，即invoice_type＝RED时，可选填；或者apply_id不为空的时候，可选填；其他情况必填。
	OpenId *string `json:"open_id,omitempty"`
	// OriBlueInvCode 是红票情况下，必须填入原始蓝票的发票代码（全电票时为空）。
	OriBlueInvCode *string `json:"ori_blue_inv_code,omitempty"`
	// OriBlueInvNo 是红票情况下，必须填入原始蓝票的发票号码。
	OriBlueInvNo *string `json:"ori_blue_inv_no,omitempty"`
	// OutInvoiceId 是外部发票id，isv生成发票之后所生成的发票流水号，主要用做存储。
	OutInvoiceId *string `json:"out_invoice_id,omitempty"`
	// OutTradeNo 是该发票基于的实际交易流水号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// Payee 是票面上收款人信息。
	Payee *string `json:"payee,omitempty"`
	// PayeeAddressTel 是票面上销售方地址、电话信息。
	PayeeAddressTel *string `json:"payee_address_tel,omitempty"`
	// PayeeBankNameAccount 是票面上销售方开户行及账号。
	PayeeBankNameAccount *string `json:"payee_bank_name_account,omitempty"`
	// PayeeRegisterName 是票面上销售方企业名称。
	PayeeRegisterName *string `json:"payee_register_name,omitempty"`
	// PayeeRegisterNo 是票面上销售方税号信息。
	PayeeRegisterNo *string `json:"payee_register_no,omitempty"`
	// SumAmount 是价税合计，票面上的总开票金额，单位（元），保留两位小数，sum_amount=ex_tax_amount+tax_amount。
	SumAmount *string `json:"sum_amount,omitempty"`
	// TaxAmount 是合计税额，票面上的合计税额，单位（元），保留两位小数。
	TaxAmount *string `json:"tax_amount,omitempty"`
	// TradeList 是交易内容，非必填但是不可以传空值，如果无需回传trade_list则不需要传该对象
	TradeList *[]EinvTrade `json:"trade_list,omitempty"`
	// UserId 是支付宝用户userId，当发送红字发票时，即invoice_type＝RED时，可选填；或者apply_id不为空的时候，可选填；其他情况必填。
	UserId *string `json:"user_id,omitempty"`
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
