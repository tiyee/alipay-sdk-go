// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseexctrl

// EnterpriseTitleInfo 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseTitleInfo struct {
	// Address 是详细地址
	Address *string `json:"address,omitempty"`
	// BankAccount 是开户行账号
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行地址
	BankName *string `json:"bank_name,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// Telephone 是电话
	Telephone *string `json:"telephone,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
	// TitleName 是企业抬头名称
	TitleName *string `json:"title_name,omitempty"`
}
