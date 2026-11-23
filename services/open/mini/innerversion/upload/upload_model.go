package upload

// MiniAppPluginReference 是支付宝官方 V3 OpenAPI 模型。
type MiniAppPluginReference struct {
	// Lazy 是是否懒加载。
	Lazy *bool `json:"lazy,omitempty"`
	// PluginId 是插件id，即插件的唯一标识，形如小程序的appId。
	PluginId *string `json:"plugin_id,omitempty"`
	// PluginLazy 是是否懒加载。
	PluginLazy *bool `json:"plugin_lazy,omitempty"`
	// PluginVersion 是小程序引用的插件版本，可以指定插件版本，如：1.0.11；也可以填*，表示引用最新的插件版本，此时，当插件发布新版本后，小程序内的插件会自动更新到最新的版本。。
	PluginVersion *string `json:"plugin_version,omitempty"`
}

// SubPackageInfo 是支付宝官方 V3 OpenAPI 模型。
type SubPackageInfo struct {
	// BuildedPackageUrl 是构建好的结果地址。
	BuildedPackageUrl *string `json:"builded_package_url,omitempty"`
	// NewBuildedPackageUrl 是构建好的分包地址。
	NewBuildedPackageUrl *string `json:"new_builded_package_url,omitempty"`
	// NewSize 是2.0分包大小，单位字节。
	NewSize *string `json:"new_size,omitempty"`
	// Path 是分包路径。
	Path *string `json:"path,omitempty"`
	// Size 是分包大小，单位字节。
	Size *string `json:"size,omitempty"`
	// SourceUrl 是源码地址。
	SourceUrl *string `json:"source_url,omitempty"`
	// Type 是分包类型： MAIN ｜ SUB。
	Type *string `json:"type,omitempty"`
}
