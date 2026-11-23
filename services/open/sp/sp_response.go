package sp

// SpImageUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type SpImageUploadResponse struct {
	// ImageId 是图片在文件存储平台的标识。
	ImageId *string `json:"image_id,omitempty"`
}

// SpLeadsExpandCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type SpLeadsExpandCreateResponse struct {
	// LeadsId 是商机id。
	LeadsId *string `json:"leads_id,omitempty"`
}

// SpMerchantInconsistentApproveResponse 是支付宝官方 V3 OpenAPI 模型。
type SpMerchantInconsistentApproveResponse struct {
	// ApplyNo 是申请单据编号。
	ApplyNo *string `json:"apply_no,omitempty"`
}
