// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package title

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
