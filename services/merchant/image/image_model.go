package image

// MerchantImageUpload 是支付宝官方 V3 OpenAPI 模型。
type MerchantImageUpload struct {
	// ImageType 是图片格式，支持格式：jpg、jpeg、png
	ImageType *string `json:"image_type,omitempty"`
}
