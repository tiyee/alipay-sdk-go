// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package experience

// MiniExperienceCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniExperienceCancelResponse struct {
}

// MiniExperienceCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniExperienceCreateResponse struct {
}

// MiniExperienceQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniExperienceQueryResponse struct {
	// ExpQrCodeUrl 是小程序体验版二维码地址。
	ExpQrCodeUrl *string `json:"exp_qr_code_url,omitempty"`
	// ExpSchemaUrl 是体验版schema。
	ExpSchemaUrl *string `json:"exp_schema_url,omitempty"`
	// Status 是体验版打包状态。。
	Status *string `json:"status,omitempty"`
}
