// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package template

// MiniTemplateMarketingCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateMarketingCreateRequest struct {
	// ActivityId 是营销活动id。
	ActivityId *string `json:"activity_id,omitempty"`
	// GmtEnd 是活动结束时间。
	GmtEnd *string `json:"gmt_end,omitempty"`
	// GmtStart 是活动开始时间。
	GmtStart *string `json:"gmt_start,omitempty"`
	// TemplateIds 是消息模板id列表，最多50个模板id。
	TemplateIds *[]string `json:"template_ids,omitempty"`
	// Title 是消息运营位名称，不填默认使用券名称。
	Title *string `json:"title,omitempty"`
}

// MiniTemplateUsageQueryRequest is the request type for alipay.open.mini.template.usage.query.
type MiniTemplateUsageQueryRequest struct {
	// TemplateId 是小程序模板APPID
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
	// PageNum 是查询的页数，起始为 1（第一页） 。默认第一页。
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是每页的数量，最多查询50个，默认查询10个
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,omitempty"`
	// TemplateVersion 是模板小程序的版本号
	TemplateVersion *string `json:"template_version,omitempty" query:"template_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端：	com.alipay.alipaywallet：支付宝端；	com.alibaba.android.rimet：DINGDING端；	com.amap.app：高德端；	com.alibaba.ailabs.genie.webapps：天猫精灵端；	com.alipay.iot.xpaas：支付宝IoT端。	如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
}
