// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package miniapp

// MiniMiniAppBrandCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandCancelResponse struct {
}

// MiniMiniAppBrandQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandQueryResponse struct {
	MerchantBrandListResult *MerchantBrandListResult `json:"merchant_brand_list_result,omitempty"`
	MiniAppBrandAuditResult *MiniAppBrandAuditResult `json:"miniapp_brand_audit_result,omitempty"`
}

// MiniMiniAppBrandSubmitResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandSubmitResponse struct {
}

// MiniMiniAppBrandUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandUploadResponse struct {
	// FileKey 是图片上传成功后的图片存储键(可用于填写调用alipay.open.mini.miniapp.brand.submit时需要的图片类型的字段)。
	FileKey *string `json:"file_key,omitempty"`
}

// MiniMiniAppServiceConfigModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppServiceConfigModifyResponse struct {
}
