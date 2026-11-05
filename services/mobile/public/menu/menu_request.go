package menu

// MobilePublicMenuAddRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMenuAddRequest struct {
	// BizContent 是json串，详情请见
	BizContent *string `json:"biz_content,omitempty"`
}

// MobilePublicMenuUpdateRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMenuUpdateRequest struct {
	// BizContent 是json串，详情请见
	BizContent *string `json:"biz_content,omitempty"`
}

// MobilePublicMenuGetRequest is the request model for alipay.mobile.public.menu.get.
type MobilePublicMenuGetRequest struct{}
