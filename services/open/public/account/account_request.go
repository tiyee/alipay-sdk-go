// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package account

// PublicAccountCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicAccountCreateRequest struct {
	// AgreementId 是账户添加成功，在支付宝与其对应的协议号。如果账户重复添加，接口保证幂等依然视为添加成功，返回此前该账户在支付宝对应的协议号。其他异常该字段不存在。。
	AgreementId *string `json:"agreement_id,omitempty"`
	// BindAccountNo 是绑定帐号，建议在开发者的系统中保持唯一性。
	BindAccountNo *string `json:"bind_account_no,omitempty"`
	// DisplayName 是开发者期望在服务窗首页看到的关于该用户的显示信息，最长10个字符。
	DisplayName *string `json:"display_name,omitempty"`
	// FromUserId 是要绑定的商户会员对应的支付宝userid，2088开头长度为16位的字符串。
	FromUserId *string `json:"from_user_id,omitempty"`
	// OpenId 是支付宝用户的唯一标识。
	OpenId *string `json:"open_id,omitempty"`
	// RealName 是要绑定的商户会员的真实姓名，最长10个汉字。
	RealName *string `json:"real_name,omitempty"`
	// Remark 是备注信息，开发者可以通过该字段纪录其他的额外信息。
	Remark *string `json:"remark,omitempty"`
}

// PublicAccountDeleteRequest is the request type for alipay.open.public.account.delete.
type PublicAccountDeleteRequest struct {
	// AgreementId 是协议号，商户会员在支付宝服务窗账号中的唯一标识，与bind_account_no不能同时为空
	AgreementId *string `json:"agreement_id,omitempty" query:"agreement_id,omitempty"`
	// BindAccountNo 是绑定帐号，建议在开发者的系统中保持唯一性，与agreement_id不能同时为空
	BindAccountNo *string `json:"bind_account_no,omitempty" query:"bind_account_no,omitempty"`
	// FromUserId 是绑定用户的支付宝userid，2088开头16位长度的字符串，与agreementId不能同时为空
	FromUserId *string `json:"from_user_id,omitempty" query:"from_user_id,omitempty"`
	// OpenId 是支付宝用户的唯一标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}

// PublicAccountQueryRequest is the request type for alipay.open.public.account.query.
type PublicAccountQueryRequest struct {
	// UserId 是用户的支付宝用户号，2088开头。
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是支付宝用户的唯一标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}
