// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package setting

// PublicSettingCategoryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicSettingCategoryQueryResponse struct {
	// PrimaryCategory 是已设置的一级行业分类名称。。
	PrimaryCategory *string `json:"primary_category,omitempty"`
	// SecondaryCategory 是已设置的二级行业分类名称。。
	SecondaryCategory *string `json:"secondary_category,omitempty"`
}
