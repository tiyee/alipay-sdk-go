package template

// MobilePublicTemplateMessageDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicTemplateMessageDeleteRequest struct {
	// TemplateId 是模板id
	TemplateId *string `json:"template_id,omitempty"`
}

// MobilePublicTemplateMessageQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicTemplateMessageQueryRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// Template 是模板
	Template *string `json:"template,omitempty"`
	// TemplateId 是模板id
	TemplateId *string `json:"template_id,omitempty"`
}

// MobilePublicTemplateMessageGetRequest is the query request model for alipay.mobile.public.template.message.get.
type MobilePublicTemplateMessageGetRequest struct {
	// TemplateId 是消息母板id
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
}
