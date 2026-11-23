package gray

// MiniInnerVersionGrayFinishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionGrayFinishRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionGrayPublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionGrayPublishRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BetaAppIdList 是待邀测应用列表，灰度插件时使用，如无邀测必要请不要传入。
	BetaAppIdList *[]string `json:"beta_app_id_list,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// GrayStrategy 是灰度值。
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
}
