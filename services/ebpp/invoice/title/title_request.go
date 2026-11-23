// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package title

// InvoiceTitleDynamicGetRequest is the query request model for alipay.ebpp.invoice.title.dynamic.get.
type InvoiceTitleDynamicGetRequest struct {
	// BarCode 是抬头动态码
	BarCode *string `json:"bar_code,omitempty" query:"bar_code,omitempty"`
}

// InvoiceTitleListGetRequest is the query request model for alipay.ebpp.invoice.title.list.get.
type InvoiceTitleListGetRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
	// UserId 是支付宝用户id
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是支付宝用户id
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}
