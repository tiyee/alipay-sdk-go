package image

// MerchantImageUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantImageUploadResponse struct {
	// ImageId 是图片在文件存储平台的标识
	ImageId *string `json:"image_id,omitempty"`
}
