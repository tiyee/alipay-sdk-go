package live

// MerchantLiveChannelQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantLiveChannelQueryResponse struct {
	// ChannelContent 是渠道内容，包含主播、文章的上游信息。字符串内容为Map，需要转换
	ChannelContent *string `json:"channel_content,omitempty"`
	// ChannelIdentity 是渠道业务标识
	ChannelIdentity *string `json:"channel_identity,omitempty"`
	// ChannelSecret 是渠道密文
	ChannelSecret *string `json:"channel_secret,omitempty"`
	// ChannelType 是渠道类型
	ChannelType *string `json:"channel_type,omitempty"`
}

// MerchantLiveItemPlanModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantLiveItemPlanModifyResponse struct {
	// MsgInfo 是针对操作结果的说明
	MsgInfo *string `json:"msg_info,omitempty"`
}
