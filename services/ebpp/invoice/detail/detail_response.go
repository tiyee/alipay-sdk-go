// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package detail

// InvoiceDetailOutputQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceDetailOutputQueryResponse struct {
	// AntiFakeCode 是防伪码，发票校验码
	AntiFakeCode *string `json:"anti_fake_code,omitempty"`
	// ExTaxAmount 是不含税金额（废弃），单位元
	ExTaxAmount *string `json:"ex_tax_amount,omitempty"`
	// ExpenseStatus 是发票报销状态 	取值范围： 	WAIT_EXPENSE－未报销 	EXPENSE_PROCESSING－报销中 	EXPENSE_FINISHED－已报销
	ExpenseStatus *string `json:"expense_status,omitempty"`
	// HasRisk 是该发票可能存在异常，请核实后使用 	true:无异常 	false:存在异常
	HasRisk *bool `json:"has_risk,omitempty"`
	// HasUploadPdf 是是否已上传发票pdf文件 	false－未上传 	true－已上传
	HasUploadPdf *string `json:"has_upload_pdf,omitempty"`
	// InvoiceAmount 是发票金额（价税合计金额），单位元
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// InvoiceCode 是发票代码	（全电票为空）
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceImgUrl 是发票缩略图地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// InvoiceItemContentList 是发票内容项
	InvoiceItemContentList *[]InvoiceItemContent `json:"invoice_item_content_list,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceStatus 是发票状态 	SUCCEED－成功 	EXPIRED－已失效 	TRANSFERRED－已转交
	InvoiceStatus *string `json:"invoice_status,omitempty"`
	// InvoiceType 是发票类型 	值有两种情况：blue－蓝票 red－红票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// PayeeAddress 是销售方地址
	PayeeAddress *string `json:"payee_address,omitempty"`
	// PayeeBankAccount 是销售方开户银行账号
	PayeeBankAccount *string `json:"payee_bank_account,omitempty"`
	// PayeeBankName 是销售方开户银行名称
	PayeeBankName *string `json:"payee_bank_name,omitempty"`
	// PayeePhone 是销售方电话
	PayeePhone *string `json:"payee_phone,omitempty"`
	// PayeeRegisterName 是销售方名称
	PayeeRegisterName *string `json:"payee_register_name,omitempty"`
	// PayeeTaxNo 是销售方税号
	PayeeTaxNo *string `json:"payee_tax_no,omitempty"`
	// PayerAddress 是购买方地址
	PayerAddress *string `json:"payer_address,omitempty"`
	// PayerBankAccount 是购买方开户银行账号
	PayerBankAccount *string `json:"payer_bank_account,omitempty"`
	// PayerBankName 是购买方开户银行名称
	PayerBankName *string `json:"payer_bank_name,omitempty"`
	// PayerName 是购买方名称
	PayerName *string `json:"payer_name,omitempty"`
	// PayerPhone 是购买方电话
	PayerPhone *string `json:"payer_phone,omitempty"`
	// PayerTaxNo 是购买方税号
	PayerTaxNo *string `json:"payer_tax_no,omitempty"`
	// SumTaxAmount 是合计税额（废弃）
	SumTaxAmount *string `json:"sum_tax_amount,omitempty"`
	// TaxType 是票种	PLAIN：增值税电子普通发票	SPECIAL：增值税专用发票	ALL_ELECTRONIC_GENERAL：电子发票（普通发票）	ALL_ELECTRONIC_SPECIAL：电子发票（增值税专用发票）	PLAIN_INVOICE:增值税普通发票	PAPER_INVOICE:增值税普通发票（卷式）	SALSE_INVOICE:机动车销售统一发票
	TaxType *string `json:"tax_type,omitempty"`
}
