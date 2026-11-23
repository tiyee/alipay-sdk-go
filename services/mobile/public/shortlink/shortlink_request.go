package shortlink

// MobilePublicShortLinkCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicShortLinkCreateRequest struct {
	// BizContent 是json串，详情请见
	BizContent *string `json:"biz_content,omitempty"`
}
