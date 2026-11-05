// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package miniapp

// AuthorizeInfo 是支付宝官方 V3 OpenAPI 模型。
type AuthorizeInfo struct {
	// AuthIds 是授权方身份证明(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片,可以上传多个授权方身份证明)。
	AuthIds *[]string `json:"auth_ids,omitempty"`
	// AuthMaterials 是授权有效资质(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片，可以上传多个授权有效资质)。
	AuthMaterials *[]string `json:"auth_materials,omitempty"`
	// AuthName 是被授权人名称。
	AuthName *string `json:"auth_name,omitempty"`
	// EndDate 是授权有效期(结束日期)。
	EndDate *string `json:"end_date,omitempty"`
	// StartDate 是授权有效期(开始日期)。
	StartDate *string `json:"start_date,omitempty"`
}

// BrandRegistrationInfo 是支付宝官方 V3 OpenAPI 模型。
type BrandRegistrationInfo struct {
	// EndDate 是商标有效期(结束日期)。
	EndDate *string `json:"end_date,omitempty"`
	// RegMaterials 是商标注册证书信息(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片,可以上传多个商标注册证书信息)。
	RegMaterials *[]string `json:"reg_materials,omitempty"`
	// RegNumber 是商标注册号。
	RegNumber *string `json:"reg_number,omitempty"`
	// Registrant 是商标注册人。
	Registrant *string `json:"registrant,omitempty"`
	// StartDate 是商标有效期(开始日期)。
	StartDate *string `json:"start_date,omitempty"`
}

// BrandResult 是支付宝官方 V3 OpenAPI 模型。
type BrandResult struct {
	// BrandId 是已有的品牌的id。
	BrandId *string `json:"brand_id,omitempty"`
	// BrandName 是已有品牌的名称。
	BrandName *string `json:"brand_name,omitempty"`
}

// MerchantBrandListResult 是支付宝官方 V3 OpenAPI 模型。
type MerchantBrandListResult struct {
	// BrandListResult 是商户已有品牌。
	BrandListResult *[]BrandResult `json:"brand_list_result,omitempty"`
}

// MiniAppBrandAuditResult 是支付宝官方 V3 OpenAPI 模型。
type MiniAppBrandAuditResult struct {
	// BrandId 是小程序提交品牌认证时的品牌id。
	BrandId *string `json:"brand_id,omitempty"`
	// BrandName 是小程序提交品牌认证时的品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// BrandStatus 是小程序提交品牌认证后的认证状态。
	BrandStatus *string `json:"brand_status,omitempty"`
	// HasBrand 是小程序提交品牌认证是否认证成功。
	HasBrand *string `json:"has_brand,omitempty"`
	// InvalidReason 是失效原因。
	InvalidReason *string `json:"invalid_reason,omitempty"`
	// RejectReason 是驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
}
