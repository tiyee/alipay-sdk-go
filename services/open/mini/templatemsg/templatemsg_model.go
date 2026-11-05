// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package templatemsg

// MarketingDeliveryDetail 是支付宝官方 V3 OpenAPI 模型。
type MarketingDeliveryDetail struct {
	// BindIdList 是投放详情绑定的模板id列表。
	BindIdList *[]string `json:"bind_id_list,omitempty"`
	// DetailId 是投放详情id。
	DetailId *string `json:"detail_id,omitempty"`
	// GmtEnd 是结束时间。
	GmtEnd *string `json:"gmt_end,omitempty"`
	// GmtStart 是开始时间。
	GmtStart *string `json:"gmt_start,omitempty"`
	// Name 是投放详情名称，取创建运营投放的title，未传title取券名称。
	Name *string `json:"name,omitempty"`
	// Status 是投放详情状态，INIT:下线状态，ON_LINE:上线状态。
	Status *string `json:"status,omitempty"`
}
