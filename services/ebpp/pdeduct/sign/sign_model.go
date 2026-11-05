package sign

// SpecifiedSortChannelParam 是支付宝官方 V3 OpenAPI 模型。
type SpecifiedSortChannelParam struct {
	// AssetTypeCode 是资产编码
	AssetTypeCode *string `json:"asset_type_code,omitempty"`
	// AssignedChannel 是指定支付渠道标识
	AssignedChannel *string `json:"assigned_channel,omitempty"`
	// ChannelFullName 是支付渠道名称
	ChannelFullName *string `json:"channel_full_name,omitempty"`
	// ChannelIndex 是支付渠道索引
	ChannelIndex *string `json:"channel_index,omitempty"`
}
