package upload

// MiniInnerVersionNoBuildUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionNoBuildUploadRequest struct {
	// BuildExtraInfo 是扩展信息，比如adaptorName、tinycliVersion、tinycliName、import-module、allowPrecompile、extJson、allowInstallDependency、aggregationMainAppId，如果没有特殊要求，tinycliVersion版本请用最新的： https://registry.npm.alibaba-inc.com/@alipay/tiny-cli/huoban-prod。
	BuildExtraInfo *string `json:"build_extra_info,omitempty"`
	// BuildJsPermission 是jsapi权限文件。
	BuildJsPermission *string `json:"build_js_permission,omitempty"`
	// BuildMainUrl 是小程序页面主入口。/index.html#page/component/index，该值需要和extendInfo中的page参数保持一致。
	BuildMainUrl *string `json:"build_main_url,omitempty"`
	// BuildQcloudInfo 是打包平台扩展信息。
	BuildQcloudInfo *string `json:"build_qcloud_info,omitempty"`
	// BuildVersion 是小程序版本。
	BuildVersion *string `json:"build_version,omitempty"`
	// BuildedPackageSize 是已经构建过的amr包大小，单位是字节，建议如实填写。
	BuildedPackageSize *string `json:"builded_package_size,omitempty"`
	// BuildedPackageUrl 是已经构建的包地址，目前主要自行构建的场景使用。
	BuildedPackageUrl *string `json:"builded_package_url,omitempty"`
	// BuildedPluginSize 是构建好的插件包amr大小。
	BuildedPluginSize *string `json:"builded_plugin_size,omitempty"`
	// BuildedPluginUrl 是构建好的插件包地址。
	BuildedPluginUrl *string `json:"builded_plugin_url,omitempty"`
	// BundleId 是一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundleId入驻。
	BundleId *string `json:"bundle_id,omitempty"`
	// Components 是组件参数。
	Components *string `json:"components,omitempty"`
	// InstCode 是上传调试版的接入租户类型。
	InstCode *string `json:"inst_code,omitempty"`
	// MiniAppId 是小程序id。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// NewBuildedPackageSize 是2.0已经构建过的amr包大小，单位是字节，主要是给自行构建的场景使用。
	NewBuildedPackageSize *string `json:"new_builded_package_size,omitempty"`
	// NewBuildedPackageUrl 是2.0产物包地址。
	NewBuildedPackageUrl *string `json:"new_builded_package_url,omitempty"`
	// NewBuildedPluginSize 是构建好的插件包amr大小。
	NewBuildedPluginSize *string `json:"new_builded_plugin_size,omitempty"`
	// NewBuildedPluginUrl 是构建好的插件包地址。
	NewBuildedPluginUrl *string `json:"new_builded_plugin_url,omitempty"`
	// NoSign 是否免加签, true的时候代表传过来的是加签后的包地址。
	NoSign *bool `json:"no_sign,omitempty"`
	// PluginRefs 是小程序代码中引用的插件列表，包含插件id和插件版本信息，业务方需要自行解析源码包的app.json里面的plugins信息。
	PluginRefs *[]MiniAppPluginReference `json:"plugin_refs,omitempty"`
	// SubPackages 是构建好的分包信息。
	SubPackages *[]SubPackageInfo `json:"sub_packages,omitempty"`
}

// MiniInnerVersionTemplatedUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionTemplatedUploadRequest struct {
	// AppOrigin 是来源类型，新接入方需要向支付宝申请专用来源，否则不予接入。。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是格式为: x.y.z，其中x、y、z均为整型数字，版本不能重复构建，且新版本号必须高于旧版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// Ext 是模板的配置参数，完整内容请参见：https://docs.alipay.com/mini/isv/ortga1。
	Ext *string `json:"ext,omitempty"`
	// MiniAppId 是小程序ID，特殊场景专用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// TemplateId 是模板id，模板必须已经处于上架/下架状态。
	TemplateId *string `json:"template_id,omitempty"`
	// TemplateVersion 是模板版本，必须处于上架/下架状态。如果为空，默认取当前在架，或者最新下架的版本。。
	TemplateVersion *string `json:"template_version,omitempty"`
}

// MiniInnerVersionUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionUploadRequest struct {
	// AppOrigin 是业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BuildAppType 是IDE开发打包类型，写死tinyApp。
	BuildAppType *string `json:"build_app_type,omitempty"`
	// BuildExtraInfo 是扩展信息，比如adaptorName、tinycliVersion、tinycliName、import-module、allowPrecompile、extJson、allowInstallDependency、aggregationMainAppId，如果没有特殊要求，tinycliVersion版本请用最新的： https://registry.npm.alibaba-inc.com/@alipay/tiny-cli/huoban-prod。
	BuildExtraInfo *string `json:"build_extra_info,omitempty"`
	// BuildExtraMiniProjectConfig 是用于覆盖开发者配置字段，实现止血和灰度逻辑。
	BuildExtraMiniProjectConfig *string `json:"build_extra_mini_project_config,omitempty"`
	// BuildJsPermission 是jsapi 权限文件,钉钉外请勿传。
	BuildJsPermission *string `json:"build_js_permission,omitempty"`
	// BuildMainUrl 是小程序页面主入口。/index.html#page/component/index，该值需要和extendInfo中的page参数保持一致。
	BuildMainUrl *string `json:"build_main_url,omitempty"`
	// BuildMaxAndroidClientVersion 是最大Android客户端版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题。
	BuildMaxAndroidClientVersion *string `json:"build_max_android_client_version,omitempty"`
	// BuildMaxIosClientVersion 是最大iOS客户单版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题。
	BuildMaxIosClientVersion *string `json:"build_max_ios_client_version,omitempty"`
	// BuildMinAndroidClientVersion 是最小Android客户端版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题。
	BuildMinAndroidClientVersion *string `json:"build_min_android_client_version,omitempty"`
	// BuildMinIosClientVersion 是最小iOS客户单版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题。
	BuildMinIosClientVersion *string `json:"build_min_ios_client_version,omitempty"`
	// BuildPackageMd5 是源码包文件流的MD5字符串。
	BuildPackageMd5 *string `json:"build_package_md_5,omitempty"`
	// BuildPackageName 是包名称。
	BuildPackageName *string `json:"build_package_name,omitempty"`
	// BuildPackageStream 是base64编码后的小程序源码包文件流。
	BuildPackageStream *string `json:"build_package_stream,omitempty"`
	// BuildQcloudInfo 是打包平台扩展信息。
	BuildQcloudInfo *string `json:"build_qcloud_info,omitempty"`
	// BuildSignedPkgUrl 是已经构建加签过的包地址。
	BuildSignedPkgUrl *string `json:"build_signed_pkg_url,omitempty"`
	// BuildSourcePkgSize 是源码包大小。
	BuildSourcePkgSize *string `json:"build_source_pkg_size,omitempty"`
	// BuildSourcePkgUrl 是源码包地址，源码包大小不能超过20M且地址必须永久有效，若超20M，将限制业务来源调用。将源码文件夹整个用zip格式压缩成后缀amr文件传至文件服务器。
	BuildSourcePkgUrl *string `json:"build_source_pkg_url,omitempty"`
	// BuildSubUrl 是子入口。
	BuildSubUrl *string `json:"build_sub_url,omitempty"`
	// BuildVersion 是小程序版本。
	BuildVersion *string `json:"build_version,omitempty"`
	// BuildedPackageSize 是已经构建过的包大小，单位是字节，主要是给自行构建的场景使用。
	BuildedPackageSize *string `json:"builded_package_size,omitempty"`
	// BuildedPackageUrl 是已经构建的包地址，目前主要自行构建的场景使用。
	BuildedPackageUrl *string `json:"builded_package_url,omitempty"`
	// BundleId 是一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundleId入驻。
	BundleId *string `json:"bundle_id,omitempty"`
	// ClientType 是多端类型，除了钉钉外其他业务方禁止使用。
	ClientType *string `json:"client_type,omitempty"`
	// InstCode 是上传调试版的接入租户类型。。
	InstCode *string `json:"inst_code,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginRefs 是小程序代码中引用的插件列表，包含插件id和插件版本信息，业务方需要自行解析源码包的app.json里面的plugins信息。
	PluginRefs *[]MiniAppPluginReference `json:"plugin_refs,omitempty"`
}

// MiniInnerVersionUploadStatusQueryRequest is the request type for alipay.open.mini.innerversion.uploadstatus.query.
type MiniInnerVersionUploadStatusQueryRequest struct {
	// BuildPackageId 是包管理ID，用于查询打包状态
	BuildPackageId *string `json:"build_package_id,omitempty" query:"build_package_id,omitempty"`
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// BuildVersion 是小程序版本
	BuildVersion *string `json:"build_version,omitempty" query:"build_version,omitempty"`
	// InstCode 是租户id
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
	// BundleId 是端信息
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
}
