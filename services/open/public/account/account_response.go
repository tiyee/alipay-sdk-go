// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package account

// PublicAccountCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAccountCreateResponse struct {
	// AgreementId 是协议号，商户会员在支付宝服务窗账号中的唯一标识。
	AgreementId *string `json:"agreement_id,omitempty"`
}

// PublicAccountDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAccountDeleteResponse struct {
}

// PublicAccountQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAccountQueryResponse struct {
	// PublicBindAccounts 是绑定账户列表。
	PublicBindAccounts *[]StdPublicBindAccount `json:"public_bind_accounts,omitempty"`
}
