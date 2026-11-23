// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package label

// PublicLabelCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicLabelCreateRequest struct {
	// Name 是标签名。
	Name *string `json:"name,omitempty"`
}

// PublicLabelQueryRequest is the request type for alipay.open.public.label.query.
type PublicLabelQueryRequest struct{}

// PublicLabelUserQueryRequest is the request type for alipay.open.public.label.user.query.
type PublicLabelUserQueryRequest struct {
	// UserId 是支付宝用户的userid，2088开头长度为16位的字符串
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是支付宝用户的唯一标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}
