package beta

// MiniInnerVersionBetaConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionBetaConfirmResponse struct {
}

// MiniInnerVersionBetaInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionBetaInfoQueryResponse struct {
	// PluginBetaItemList 是邀测结果列表。
	PluginBetaItemList *[]PluginBetaItemInfo `json:"plugin_beta_item_list,omitempty"`
	// Status 是邀测状态，待确认WAITCHECK/确认CHECKED/拒绝REJECT。
	Status *string `json:"status,omitempty"`
}
