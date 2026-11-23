package beta

// PluginBetaItemInfo 是支付宝官方 V3 OpenAPI 模型。
type PluginBetaItemInfo struct {
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// BetaStatus 是被邀测状态，待确认WAITCHECK/确认CHECKED/拒绝REJECT。
	BetaStatus *string `json:"beta_status,omitempty"`
	// Memo 是拒绝原因。
	Memo *string `json:"memo,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginId 是插件id。
	PluginId *string `json:"plugin_id,omitempty"`
	// PluginVersion 是插件版本。
	PluginVersion *string `json:"plugin_version,omitempty"`
	// QrCodeUrl 是邀测二维码。
	QrCodeUrl *string `json:"qr_code_url,omitempty"`
}
