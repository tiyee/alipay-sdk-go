// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerbaseinfo

// MiniInnerBaseInfoAppLogoUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerBaseInfoAppLogoUploadResponse struct {
	// AppLogo 是上传图片结果地址。 图片要求： 1、不能超过256kb 2、支持的图片格式为jpg、jpeg、png 3、长宽比例必须为为1:1。
	AppLogo *string `json:"app_logo,omitempty"`
}

// MiniInnerBaseInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerBaseInfoQueryResponse struct {
	// AppAliasName 是小程序别名，简称。
	AppAliasName *string `json:"app_alias_name,omitempty"`
	// AppCategoryIds 是类目Id列表。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序应用描述。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序英文名称。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppKey 是手淘开放平台鉴权key，支付宝不需要。
	AppKey *string `json:"app_key,omitempty"`
	// AppLogo 是小程序应用logo图标。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是小程序简介。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// AppSubType 是小程序类型，TINYAPP_TEMPLATE，TINYAPP_NORMAL，TINYAPP_PLUGIN。
	AppSubType *string `json:"app_sub_type,omitempty"`
	// DevId 是小程序所属主体信息。
	DevId *string `json:"dev_id,omitempty"`
	// GmtCreate 是小程序主体创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是小程序主体更新时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// MiniAppId 是小程序Id。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，一级与二级、三级用下划线隔开，最多可以选四个类目，类目之间;隔开。使用后不再读取app_category_ids值，老前台类目将废弃。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// Origin 是应用创建来源，alipay = 支付宝，taobao = 淘宝。
	Origin *string `json:"origin,omitempty"`
	// OwnerEntity 是小程序主体信息。
	OwnerEntity *string `json:"owner_entity,omitempty"`
	// OwnerPortrait 是小程序主体头像信息。
	OwnerPortrait *string `json:"owner_portrait,omitempty"`
	// PluginRefreshType 是AUTO/自动更新、MANUAL/手动更新。
	PluginRefreshType *string `json:"plugin_refresh_type,omitempty"`
	// ServiceEmail 是小程序客服邮箱。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话。
	ServicePhone *string `json:"service_phone,omitempty"`
}
