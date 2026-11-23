// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package deliverychannel

// ActivityDeliveryChannelQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryChannelQueryResponse struct {
	// DeliveryChannelInfoList 是"可投放的渠道信息列表。 用于表达当前可以投放的渠道列表信息"
	DeliveryChannelInfoList *[]DeliveryChannelInfo `json:"delivery_channel_info_list,omitempty"`
	// PageNum 是查询的页码。 特别说明： 页码从1开始。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是每页查询个数
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是可返回的渠道总数
	TotalSize *string `json:"total_size,omitempty"`
}
