// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseexctrl

// InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// PageNum 是页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页行数
	PageSize *int64 `json:"page_size,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleCreateRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// Address 是企业地址
	Address *string `json:"address,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BankAccount 是开户行账号
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行名称
	BankName *string `json:"bank_name,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// Telephone 是电话
	Telephone *string `json:"telephone,omitempty"`
	// TitleCode 是抬头编码
	TitleCode *string `json:"title_code,omitempty"`
	// TitleName 是企业抬头名称
	TitleName *string `json:"title_name,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleModifyRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// Address 是详细地址
	Address *string `json:"address,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BankAccount 是开户行账号
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行
	BankName *string `json:"bank_name,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// Telephone 是电话
	Telephone *string `json:"telephone,omitempty"`
	// TitleCode 是抬头编码
	TitleCode *string `json:"title_code,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
	// TitleName 是企业抬头名称
	TitleName *string `json:"title_name,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleQueryRequest is the query request model for alipay.ebpp.invoice.enterpriseexctrl.employertitle.query.
type InvoiceEnterpriseExCtrlEmployerTitleQueryRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty" query:"title_id,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty" query:"tax_register_no,omitempty"`
}
