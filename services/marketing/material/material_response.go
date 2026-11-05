// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package material

// MaterialCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MaterialCreateResponse struct {
	// MaterialId 是素材ID。
	MaterialId *string `json:"material_id,omitempty"`
}

// MaterialImageUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MaterialImageUploadResponse struct {
	// ResourceEnhance 是图片是否优化过。系统会根据要求，若上传的图片不符合要求的尺寸，会对上传的图片做尺寸的优化。
	ResourceEnhance *bool `json:"resource_enhance,omitempty"`
	// ResourceId 是图片唯一资源ID
	ResourceId *string `json:"resource_id,omitempty"`
	// ResourceUrl 是上传后图片的url地址
	ResourceUrl *string `json:"resource_url,omitempty"`
}
