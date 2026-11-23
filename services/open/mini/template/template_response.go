// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package template

// MiniTemplateMarketingCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateMarketingCreateResponse struct {
	// DetailId 是投放详情主键id。
	DetailId *string `json:"detail_id,omitempty"`
}

// MiniTemplateUsageQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateUsageQueryResponse struct {
	// AppIds 是小程序appId列表。
	AppIds *[]string `json:"app_ids,omitempty"`
	// TemplateUsageInfoList 是模板使用信息。
	TemplateUsageInfoList *[]TemplateUsageInfo `json:"template_usage_info_list,omitempty"`
}
