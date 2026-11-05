package qrcode

// MobilePublicQRCodeCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicQRCodeCreateRequest struct {
	// BizContent 是json串，详情请见
	BizContent *string `json:"biz_content,omitempty"`
}
