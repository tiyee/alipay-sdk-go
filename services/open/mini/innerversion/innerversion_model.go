package innerversion

// MiniAppPluginInfo 是支付宝官方 V3 OpenAPI 模型。
type MiniAppPluginInfo struct {
	// PackageUrl 是插件包地址。
	PackageUrl *string `json:"package_url,omitempty"`
	// PluginId 是插件id，即插件的唯一标识。
	PluginId *string `json:"plugin_id,omitempty"`
	// RequireVersion 是小程序引用的插件版本，可以指定插件版本，1.2.3。也可以使用最新版本,*。推荐使用*。
	RequireVersion *string `json:"require_version,omitempty"`
}

// MiniVersionBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionBaseInfo struct {
	// AppVersion 是版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BuildSource 是小程序所属来源。
	BuildSource *string `json:"build_source,omitempty"`
	// BundleId 是com.alipay.alipaywallet:支付宝，com.amap.app:高德。
	BundleId *string `json:"bundle_id,omitempty"`
	// DevId 是小程序开发者ID。
	DevId *string `json:"dev_id,omitempty"`
	// GrayStrategy 是小程序灰度值。
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// MiniAppId 是小程序应用ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Status 是小程序状态。
	Status *string `json:"status,omitempty"`
}
