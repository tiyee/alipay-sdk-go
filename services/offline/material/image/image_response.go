package image

// OfflineMaterialImageUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMaterialImageUploadResponse struct {
	// ImageId 是图片/视频在商家中心的唯一标识
	ImageId *string `json:"image_id,omitempty"`
	// ImageUrl 是图片/视频的访问地址（为了防止盗链，该地址不允许嵌在其他页面展示，只能在新页面展示）
	ImageUrl *string `json:"image_url,omitempty"`
}
