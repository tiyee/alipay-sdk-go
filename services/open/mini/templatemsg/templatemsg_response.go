// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package templatemsg

// MiniTemplateMsgMaketingBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateMsgMaketingBatchQueryResponse struct {
	// MarketingDeliveryDetailList 是投放详情列表。
	MarketingDeliveryDetailList *[]MarketingDeliveryDetail `json:"marketing_delivery_detail_list,omitempty"`
	// Total 是投放详情总条数。
	Total *int64 `json:"total,omitempty"`
}

// MiniTemplateMsgTinyPaySwitchConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateMsgTinyPaySwitchConfirmResponse struct {
}
