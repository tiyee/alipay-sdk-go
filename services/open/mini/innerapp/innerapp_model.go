// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerapp

// MiniAppServiceInfo 是支付宝官方 V3 OpenAPI 模型。
type MiniAppServiceInfo struct {
	// BizStatus 是插件发布状态码，暂存100，风控审核200，运营审核300，等待上架400，已预发上架500，已上架501，已下架600，已驳回700。
	BizStatus *string `json:"biz_status,omitempty"`
	// IsInner 是是否是内部标，true/false。
	IsInner *bool `json:"is_inner,omitempty"`
	// IsOrder 是是否订购，true/false。
	IsOrder *bool `json:"is_order,omitempty"`
	// IsvAppId 是三方应用appid。
	IsvAppId *string `json:"isv_app_id,omitempty"`
	// MiniAppId 是应用id。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppName 是测试插件。
	MiniAppName *string `json:"mini_app_name,omitempty"`
	// SellerId 是卖家pid。
	SellerId *string `json:"seller_id,omitempty"`
	// SellerName 是卖家名。
	SellerName *string `json:"seller_name,omitempty"`
	// ServiceCode 是商品CODE。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceLogo 是服务图标。
	ServiceLogo *string `json:"service_logo,omitempty"`
	// ServiceName 是服务名。
	ServiceName *string `json:"service_name,omitempty"`
	// ServiceSlogan 是服务简介。
	ServiceSlogan *string `json:"service_slogan,omitempty"`
	// ShowType 是是否在服务市场透出，SHOW展示、HIDE隐藏。
	ShowType *string `json:"show_type,omitempty"`
}

// PluginUseConfigInfo 是支付宝官方 V3 OpenAPI 模型。
type PluginUseConfigInfo struct {
	// BundleId 是端。
	BundleId *string `json:"bundle_id,omitempty"`
	// GrayCode 是灰度策略。
	GrayCode *string `json:"gray_code,omitempty"`
	// PluginDevelopVersion 是插件研发版本。
	PluginDevelopVersion *string `json:"plugin_develop_version,omitempty"`
	// Status 是GRAY(灰度中)/RELEASE(在线状态)。
	Status *string `json:"status,omitempty"`
}

// PluginUseRelationInfo 是支付宝官方 V3 OpenAPI 模型。
type PluginUseRelationInfo struct {
	// BetaMemo 是邀测驳回原因。
	BetaMemo *string `json:"beta_memo,omitempty"`
	// BetaPluginVersion 是邀测插件版本号。
	BetaPluginVersion *string `json:"beta_plugin_version,omitempty"`
	// BetaQrCodeUrl 是邀测二维码。
	BetaQrCodeUrl *string `json:"beta_qr_code_url,omitempty"`
	// BetaStatus 是WAITCHECK-待确认;CHECKED-确认;REJECT-拒绝。
	BetaStatus *string `json:"beta_status,omitempty"`
	// GmtActive 是激活时间。
	GmtActive *string `json:"gmt_active,omitempty"`
	// GmtCreate 是订购时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtInvalid 是插件失效时间。
	GmtInvalid *string `json:"gmt_invalid,omitempty"`
	// MiniAppId 是应用ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginDeployVersion 是插件构建版本。
	PluginDeployVersion *string `json:"plugin_deploy_version,omitempty"`
	// PluginId 是插件ID。
	PluginId *string `json:"plugin_id,omitempty"`
	// PluginStatus 是插件状态，取值包括EXECUTING/WAIT_WORKING/WORKING/STOP_WORKING/WAIT_BUY。
	PluginStatus *string `json:"plugin_status,omitempty"`
	// PluginUseConfigInfoList 是分端版本配置信息列表。
	PluginUseConfigInfoList *[]PluginUseConfigInfo `json:"plugin_use_config_info_list,omitempty"`
	// PluginVersion 是插件版本。
	PluginVersion *string `json:"plugin_version,omitempty"`
	// RunModeType 是插件运行状态，取值包括ONLINE/TRIAL/REVIEW/DEBUG。
	RunModeType *string `json:"run_mode_type,omitempty"`
	// SourceFrom 是渠道来源，取值包括SHOP_MINI/PLUGIN_DEBUG/PLUGIN_TRIAL/PLUGIN_AUDIT/GENERAL_SHOP_ID。
	SourceFrom *string `json:"source_from,omitempty"`
}
