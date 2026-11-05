package menu

// MobilePublicMenuAddResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMenuAddResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Msg 是成功
	Msg *string `json:"msg,omitempty"`
}

// MobilePublicMenuGetResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMenuGetResponse struct {
	// Code 是success
	Code *string `json:"code,omitempty"`
	// MenuContent 是菜单内容
	MenuContent *string `json:"menu_content,omitempty"`
	// Msg 是成功
	Msg *string `json:"msg,omitempty"`
}

// MobilePublicMenuUpdateResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMenuUpdateResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Msg 是成功
	Msg *string `json:"msg,omitempty"`
}
