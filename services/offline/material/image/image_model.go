package image

// OfflineMaterialImageUpload 是支付宝官方 V3 OpenAPI 模型。
type OfflineMaterialImageUpload struct {
	// ImageName 是图片/视频名称
	ImageName *string `json:"image_name,omitempty"`
	// ImagePid 是用于显示指定图片/视频所属的partnerId（支付宝内部使用，外部商户无需填写此字段）
	ImagePid *string `json:"image_pid,omitempty"`
	// ImageType 是图片/视频格式
	ImageType *string `json:"image_type,omitempty"`
}
