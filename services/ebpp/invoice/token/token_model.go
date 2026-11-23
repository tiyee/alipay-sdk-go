// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package token

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

// InvoiceElement 是支付宝官方 V3 OpenAPI 模型。
type InvoiceElement struct {
	// ExpenseStatus 是发票报销状态
	ExpenseStatus *string `json:"expense_status,omitempty"`
	// ExtendFields 是扩展字段
	ExtendFields *string `json:"extend_fields,omitempty"`
	// FakeCode 是防伪校验码
	FakeCode *string `json:"fake_code,omitempty"`
	// FileType 是发票原始文件文件类型
	FileType *string `json:"file_type,omitempty"`
	// HasPdfFile 是发票是否有pdf文件
	HasPdfFile *bool `json:"has_pdf_file,omitempty"`
	// HasRisk 是该发票可能存在异常，请核实后使用 	true:无异常 	false:存在异常
	HasRisk *bool `json:"has_risk,omitempty"`
	// InvoiceAmount 是发票金额，含税，单位元
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceImgUrl 是发票pdf文件转换后jpg预览地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// InvoiceKind 是发票类型
	InvoiceKind *string `json:"invoice_kind,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceStatus 是发票状态
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// IsvContact 是服务商联系方式
	IsvContact *string `json:"isv_contact,omitempty"`
	// IsvName 是服务商名称
	IsvName *string `json:"isv_name,omitempty"`
	// LogoUrl 是logo地址
	LogoUrl *string `json:"logo_url,omitempty"`
	// MName 是商户全称
	MName *string `json:"m_name,omitempty"`
	// OutTaxAmount 是发票金额，不含税，单位元
	OutTaxAmount *string `json:"out_tax_amount,omitempty"`
	// PayeeName 是销方名称
	PayeeName *string `json:"payee_name,omitempty"`
	// PayeeTaxNo 是销方税号
	PayeeTaxNo *string `json:"payee_tax_no,omitempty"`
	// PayerName 是购方名称
	PayerName *string `json:"payer_name,omitempty"`
	// PayerTaxNo 是购方税号
	PayerTaxNo *string `json:"payer_tax_no,omitempty"`
	// PdfUrl 是PDF的下载链接
	PdfUrl *string `json:"pdf_url,omitempty"`
	// Source 是表示发票来源，由发票回传方带入。例如：bz_gd，bz_ele，bz_tmall等
	Source *string `json:"source,omitempty"`
	// TradeList 是该发票对应的交易
	TradeList *[]EinvTrade `json:"trade_list,omitempty"`
	// TradeMatchResult 是交易匹配结果
	TradeMatchResult *string `json:"trade_match_result,omitempty"`
}
