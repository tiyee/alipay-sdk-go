package invoice

import (
	"github.com/tiyee/alipay-sdk-go/services/boss"
)

// BossFncInvoiceApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceApplyResponse struct {
	// ResultSet 是开票申请ID，唯一性ID
	ResultSet *string `json:"result_set,omitempty"`
}

// BossFncInvoiceBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceBatchQueryResponse struct {
	Amt *boss.MultiCurrencyMoneyOpenAPI `json:"amt,omitempty"`
	// CurrentPage 是当前页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// ItemsPage 是每页条数
	ItemsPage *int64 `json:"items_page,omitempty"`
	// ResultSet 是结果对象,发票对象的集合
	ResultSet *[]ArInvoiceOpenAPIResponse `json:"result_set,omitempty"`
	// TotalItems 是查询到的结果总数
	TotalItems *int64 `json:"total_items,omitempty"`
	// TotalPages 是总页数
	TotalPages *int64 `json:"total_pages,omitempty"`
}

// BossFncInvoiceQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceQueryResponse struct {
	ResultSet *ArInvoiceOpenAPIResponse `json:"result_set,omitempty"`
}

// ArInvoiceLineOpenAPIResponse 是支付宝官方 V3 OpenAPI 模型。
type ArInvoiceLineOpenAPIResponse struct {
	Amt *boss.MultiCurrencyMoneyOpenAPI `json:"amt,omitempty"`
	// DutyFreeFlag 是免税标识
	DutyFreeFlag *string `json:"duty_free_flag,omitempty"`
	// InclTaxUnitAmt 是含税单价
	InclTaxUnitAmt *int64 `json:"incl_tax_unit_amt,omitempty"`
	// InvoiceId 是关联的发票ID
	InvoiceId *string `json:"invoice_id,omitempty"`
	// InvoiceLineId 是发票行iD
	InvoiceLineId *string `json:"invoice_line_id,omitempty"`
	// MeasurementUnit 是计量单位
	MeasurementUnit *string `json:"measurement_unit,omitempty"`
	// ProductName 是货物或劳务名称
	ProductName *string `json:"product_name,omitempty"`
	// ProductSpecification 是规格型号
	ProductSpecification *string `json:"product_specification,omitempty"`
	// Quantity 是数量
	Quantity        *int64                          `json:"quantity,omitempty"`
	TaxAmt          *boss.MultiCurrencyMoneyOpenAPI `json:"tax_amt,omitempty"`
	TaxExclusiveAmt *boss.MultiCurrencyMoneyOpenAPI `json:"tax_exclusive_amt,omitempty"`
	// TaxRate 是税率
	TaxRate *int64 `json:"tax_rate,omitempty"`
	// UnitAmt 是单价
	UnitAmt *int64 `json:"unit_amt,omitempty"`
}

// ArInvoiceOpenAPIResponse 是支付宝官方 V3 OpenAPI 模型。
type ArInvoiceOpenAPIResponse struct {
	// ArInvoiceLines 是发票行信息集合
	ArInvoiceLines *[]ArInvoiceLineOpenAPIResponse `json:"ar_invoice_lines,omitempty"`
	// BuyerAddress 是购方的地址
	BuyerAddress *string `json:"buyer_address,omitempty"`
	// BuyerBankAccount 是购方银行账号
	BuyerBankAccount *string `json:"buyer_bank_account,omitempty"`
	// BuyerBankName 是购方开户行名称
	BuyerBankName *string `json:"buyer_bank_name,omitempty"`
	// BuyerInvoiceTitle 是购方发票抬头，用于打印在发票上
	BuyerInvoiceTitle *string `json:"buyer_invoice_title,omitempty"`
	// BuyerTaxNo 是购方税号
	BuyerTaxNo *string `json:"buyer_tax_no,omitempty"`
	// BuyerTelephone 是购方的电话号码
	BuyerTelephone *string `json:"buyer_telephone,omitempty"`
	// CanLink 是是否可以关联账单
	CanLink *bool `json:"can_link,omitempty"`
	// Creator 是创建人
	Creator *string `json:"creator,omitempty"`
	// Drawer 是开票人
	Drawer *string `json:"drawer,omitempty"`
	// ElcInvoiceUrl 是电子发票在oss上的存储地址
	ElcInvoiceUrl *string `json:"elc_invoice_url,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// InstId 是对应发票所在OU的id，用于标示该发票所述的OU。
	InstId     *string                         `json:"inst_id,omitempty"`
	InvoiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"invoice_amt,omitempty"`
	// InvoiceChannel 是开票渠道，01：在线，02：线下
	InvoiceChannel *string `json:"invoice_channel,omitempty"`
	// InvoiceCode 是发票编码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是发票的开票日期，yyyyMMdd
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceId 是发票ID，唯一标示一张发票
	InvoiceId *string `json:"invoice_id,omitempty"`
	// InvoiceMaterial 是发票介质，01：电子，02：纸质
	InvoiceMaterial *string `json:"invoice_material,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceNote 是发票备注
	InvoiceNote *string `json:"invoice_note,omitempty"`
	// InvoiceStatus 是发票状态：用于标示发票当前的状态，状态类型以及相关含义 	01：待开票，02：开票中，03：已开票，04：作废中，05：已作废，06：红冲中，07：已红冲，08：部分红冲，09：无效
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// InvoiceType 是发票类型：01：专票，02：普票，03：营业税发票，04：国际形式发票，05：其它发票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// IpId 是对应发票商户的ipId，在主站标示mid，在其他环境标示ip_id。
	IpId *string `json:"ip_id,omitempty"`
	// IpRoleId 是商户结算的ipRoleId，表示商户的结算对象，在主站表示pid。
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// LastModer 是最近修改人
	LastModer *string                         `json:"last_moder,omitempty"`
	LinkedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"linked_amt,omitempty"`
	// MailStatus 是发票开出后的邮寄状态，具体状态说明 	01：待邮寄，02：已邮寄，03：退回中，04：已退回
	MailStatus *string `json:"mail_status,omitempty"`
	// Memo 是备注
	Memo *string `json:"memo,omitempty"`
	// NoBillInvoiceFlag 是是否是无账单发票 Y：是  N：不是
	NoBillInvoiceFlag *string `json:"no_bill_invoice_flag,omitempty"`
	// Online 是是否为点在发票 	true：是， false：不是
	Online *bool `json:"online,omitempty"`
	// OutBusinessNoList 是外部单据号
	OutBusinessNoList *[]string `json:"out_business_no_list,omitempty"`
	// Payee 是收款人
	Payee *string `json:"payee,omitempty"`
	// PdfFileName 是PDF附件名称
	PdfFileName *string `json:"pdf_file_name,omitempty"`
	// PdfFileUrl 是可外网访问的PDF附件的下载链接
	PdfFileUrl *string `json:"pdf_file_url,omitempty"`
	// RecentMailId 是邮件信息的主键，用于‘根据发票邮寄ID获取当前的邮寄信息’的查询输入。
	RecentMailId *string `json:"recent_mail_id,omitempty"`
	// Red 是是否红字 Y：红字， N：蓝字
	Red    *string                         `json:"red,omitempty"`
	RedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"red_amt,omitempty"`
	// RelBlueInvoices 是发票如果有关联的蓝票则显示发票关联的蓝票信息
	RelBlueInvoices *[]ArRelInvoiceOpenAPIResponse `json:"rel_blue_invoices,omitempty"`
	// RelRedInvoices 是如果发票有关联的红票则为关联的红票的信息
	RelRedInvoices *[]ArRelInvoiceOpenAPIResponse `json:"rel_red_invoices,omitempty"`
	// Reviewer 是复核人
	Reviewer *string `json:"reviewer,omitempty"`
	// SalesProductNameList 是销售产品码字符串数组
	SalesProductNameList *[]string `json:"sales_product_name_list,omitempty"`
	// SellerAddress 是销方公司地址
	SellerAddress *string `json:"seller_address,omitempty"`
	// SellerBankAccount 是销方银行账号
	SellerBankAccount *string `json:"seller_bank_account,omitempty"`
	// SellerBankName 是销方开户行名称
	SellerBankName *string `json:"seller_bank_name,omitempty"`
	// SellerCompanyName 是销方公司名称
	SellerCompanyName *string `json:"seller_company_name,omitempty"`
	// SellerTaxNo 是销方税号，用于打印在发票上
	SellerTaxNo *string `json:"seller_tax_no,omitempty"`
	// SellerTelephone 是销方公司电话
	SellerTelephone *string `json:"seller_telephone,omitempty"`
	// TntInstId 是租户id
	TntInstId *string `json:"tnt_inst_id,omitempty"`
}

// ArRelInvoiceOpenAPIResponse 是支付宝官方 V3 OpenAPI 模型。
type ArRelInvoiceOpenAPIResponse struct {
	ArInvoiceLines *ArInvoiceLineOpenAPIResponse `json:"ar_invoice_lines,omitempty"`
	// BuyerAddress 是购方的地址
	BuyerAddress *string `json:"buyer_address,omitempty"`
	// BuyerBankAccount 是购方银行账号
	BuyerBankAccount *string `json:"buyer_bank_account,omitempty"`
	// BuyerBankName 是购方开户行名称
	BuyerBankName *string `json:"buyer_bank_name,omitempty"`
	// BuyerInvoiceTitle 是购方发票抬头，用于打印在发票上
	BuyerInvoiceTitle *string `json:"buyer_invoice_title,omitempty"`
	// BuyerTaxNo 是购方税号
	BuyerTaxNo *string `json:"buyer_tax_no,omitempty"`
	// BuyerTelephone 是购方的电话号码
	BuyerTelephone *string `json:"buyer_telephone,omitempty"`
	// CanLink 是是否可以关联账单
	CanLink *bool `json:"can_link,omitempty"`
	// Creator 是创建人
	Creator *string `json:"creator,omitempty"`
	// Drawer 是开票人
	Drawer *string `json:"drawer,omitempty"`
	// ElcInvoiceUrl 是电子发票在oss上的存储地址
	ElcInvoiceUrl *string `json:"elc_invoice_url,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// InstId 是对应发票所在OU的id，用于标示该发票所述的OU。
	InstId     *string                         `json:"inst_id,omitempty"`
	InvoiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"invoice_amt,omitempty"`
	// InvoiceChannel 是开票渠道，01：在线，02：线下
	InvoiceChannel *string `json:"invoice_channel,omitempty"`
	// InvoiceCode 是发票编码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是发票的开票日期，yyyyMMdd
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceId 是发票ID，唯一标示一张发票
	InvoiceId *string `json:"invoice_id,omitempty"`
	// InvoiceMaterial 是发票介质，01：电子，02：纸质
	InvoiceMaterial *string `json:"invoice_material,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceNote 是发票备注
	InvoiceNote *string `json:"invoice_note,omitempty"`
	// InvoiceStatus 是发票状态：用于标示发票当前的状态，状态类型以及相关含义 01：待开票，02：开票中，03：已开票，04：作废中，05：已作废，06：红冲中，07：已红冲，08：部分红冲，09：无效
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// InvoiceType 是发票类型：01：专票，02：普票，03：营业税发票，04：国际形式发票，05：其它发票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// IpId 是对应发票商户的ipId，在主站标示mid，在其他环境标示ip_id。
	IpId *string `json:"ip_id,omitempty"`
	// IpRoleId 是商户结算的ipRoleId，表示商户的结算对象，在主站表示pid。
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// LastModer 是最近修改人
	LastModer *string                         `json:"last_moder,omitempty"`
	LinkedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"linked_amt,omitempty"`
	// MailStatus 是发票开出后的邮寄状态，具体状态说明 01：待邮寄，02：已邮寄，03：退回中，04：已退回
	MailStatus *string `json:"mail_status,omitempty"`
	// Memo 是备注
	Memo *string `json:"memo,omitempty"`
	// NoBillInvoiceFlag 是是否是无账单发票 Y：是 N：不是
	NoBillInvoiceFlag *string `json:"no_bill_invoice_flag,omitempty"`
	// Online 是是否为线上发票 true：是， false：不是
	Online *bool `json:"online,omitempty"`
	// Payee 是收款人
	Payee *string `json:"payee,omitempty"`
	// RecentMailId 是邮件信息的主键，用于‘根据发票邮寄ID获取当前的邮寄信息’的查询输入。
	RecentMailId *string `json:"recent_mail_id,omitempty"`
	// Red 是是否红字 Y：红字， N：蓝字
	Red    *string                         `json:"red,omitempty"`
	RedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"red_amt,omitempty"`
	// Reviewer 是复核人
	Reviewer *string `json:"reviewer,omitempty"`
	// SellerAddress 是销方公司地址
	SellerAddress *string `json:"seller_address,omitempty"`
	// SellerBankAccount 是销方银行账号
	SellerBankAccount *string `json:"seller_bank_account,omitempty"`
	// SellerBankName 是销方开户行名称
	SellerBankName *string `json:"seller_bank_name,omitempty"`
	// SellerCompanyName 是销方公司名称
	SellerCompanyName *string `json:"seller_company_name,omitempty"`
	// SellerTaxNo 是销方税号，用于打印在发票上
	SellerTaxNo *string `json:"seller_tax_no,omitempty"`
	// SellerTelephone 是销方公司电话
	SellerTelephone *string `json:"seller_telephone,omitempty"`
	// TntInstId 是租户id
	TntInstId *string `json:"tnt_inst_id,omitempty"`
}
