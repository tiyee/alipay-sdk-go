package item

// MerchantItemFileUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantItemFileUploadResponse struct {
	// MaterialId 是文件在商品中心的素材标识（素材ID长期有效）
	MaterialId *string `json:"material_id,omitempty"`
	// MaterialKey 是文件在商品中心的素材标示，创建/更新商品时使用
	MaterialKey *string `json:"material_key,omitempty"`
}
