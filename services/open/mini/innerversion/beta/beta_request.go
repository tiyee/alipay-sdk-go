package beta

// MiniInnerVersionBetaConfirmRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionBetaConfirmRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// Memo 是拒绝理由。
	Memo *string `json:"memo,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Operate 是操作：pass/reject。接受或者拒绝。
	Operate *string `json:"operate,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
	// PluginId 是要确认邀测结果的插件appid。
	PluginId *string `json:"plugin_id,omitempty"`
}

// MiniInnerVersionBetaInfoQueryRequest is the request type for alipay.open.mini.innerversion.betainfo.query.
type MiniInnerVersionBetaInfoQueryRequest struct {
	// PluginId 是插件ID
	PluginId *string `json:"plugin_id,omitempty" query:"plugin_id,omitempty"`
	// BundleId 是端信息
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// PluginVersion 是查询版本号
	PluginVersion *string `json:"plugin_version,omitempty" query:"plugin_version,omitempty"`
	// AppOrigin 是业务来源
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
}
