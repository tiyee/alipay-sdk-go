// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerapp

// MiniInnerAppCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppCreateRequest struct {
	// AppCategoryIds 是小程序类目，二级类目与三级类目下划线隔开，最多可以选择3个类目，类目之间以;隔开。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序描述，20-200个字符。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序英文名称，插件不填。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppLogo 是小程序logo。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// AppOrigin 是来源的业务方,接入文档中申请的业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppSlogan 是小程序简介，在搜索透出,插件不需要填。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// AppSubType 是应用子类型。
	AppSubType *string `json:"app_sub_type,omitempty"`
	// AppType 是应用类型。
	AppType *string `json:"app_type,omitempty"`
	// IsvAppId 是三方应用ID，仅在创建模板小程序时需要传入。
	IsvAppId *string `json:"isv_app_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list。使用后不再读取app_category_ids值，老前台类目将废弃。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// Pid 是创建的小程序所属的PID。
	Pid *string `json:"pid,omitempty"`
	// ServiceEmail 是邮箱，与客服电话二选一。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是客服电话，插件不需要。
	ServicePhone *string `json:"service_phone,omitempty"`
}

// MiniInnerAppPluginOrderRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginOrderRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// MerchandiseId 是订购的服务商品ID。
	MerchandiseId *string `json:"merchandise_id,omitempty"`
	// MiniAppId 是一二方支持传入appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// RequestId 是业务幂等号，规则为商品id+appOrigin+requestId 作为幂等流水号 非必填,不传则默认生成一个8位requestId。
	RequestId *string `json:"request_id,omitempty"`
}

// MiniInnerAppPluginRelationQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginRelationQueryRequest struct {
	// MiniAppIdList 是应用ID列表。
	MiniAppIdList *[]string `json:"mini_app_id_list,omitempty"`
	// PageNum 是页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是页大小。
	PageSize *int64 `json:"page_size,omitempty"`
	// PluginId 是插件ID。
	PluginId *string `json:"plugin_id,omitempty"`
	// PluginRelationStatusList 是插件使用关系状态，取值包括WORKING/WAIT_WORKING/STOP_WORKING/EXECUTING。
	PluginRelationStatusList *[]string `json:"plugin_relation_status_list,omitempty"`
	// RunModelType 是运行类型，取值包括ONLINE/TRIAL/REVIEW/DEBUG。
	RunModelType *string `json:"run_model_type,omitempty"`
	// ShowBetaInfo 是是否展示邀测信息。
	ShowBetaInfo *bool `json:"show_beta_info,omitempty"`
}

// MiniInnerAppPluginServicePublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginServicePublishRequest struct {
	// AbilityTypeList 是功能类型，目前支持的有：1039支付, 1040会员, 1041基础, 1056资金, 1058信用, 1111口碑, 330120安全, 360101营销。
	AbilityTypeList *[]string `json:"ability_type_list,omitempty"`
	// AppLogo 是服务发布logo。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// Description 是服务描述。
	Description *string `json:"description,omitempty"`
	// DetailForClient 是移动端详情，用于能力中心展示。
	DetailForClient *string `json:"detail_for_client,omitempty"`
	// DetailForPc 是pc端详情，用于能力中心展示。
	DetailForPc *string `json:"detail_for_pc,omitempty"`
	// MiniAppId 是插件id。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// SellCrowd 是可订购人群，1003个人, 1004企业，-1无限制。
	SellCrowd *string `json:"sell_crowd,omitempty"`
	// ShowType 是发布后是否展示，01展示（默认）、02隐藏。
	ShowType *string `json:"show_type,omitempty"`
	// Title 是服务标题。
	Title *string `json:"title,omitempty"`
	// VisitUrlForPc 是pc端url地址，不需要则为空。
	VisitUrlForPc *string `json:"visit_url_for_pc,omitempty"`
}

// MiniInnerAppPluginSyncModeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginSyncModeModifyRequest struct {
	// AppOrigin 是来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// MiniAppId 是小程序应用 ID，。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginId 是插件id。
	PluginId *string `json:"plugin_id,omitempty"`
	// SyncMode 是模式，SYNC/NONE。
	SyncMode *string `json:"sync_mode,omitempty"`
}

// MiniInnerAppPluginUseConfigCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigCancelRequest struct {
	// AppOrigin 是来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序应用 ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginDevVersion 是插件研发版本。
	PluginDevVersion *string `json:"plugin_dev_version,omitempty"`
	// PluginId 是插件id。
	PluginId *string `json:"plugin_id,omitempty"`
}

// MiniInnerAppPluginUseConfigOnlineRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigOnlineRequest struct {
	// AppOrigin 是来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序应用 ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginDevVersion 是插件研发版本。
	PluginDevVersion *string `json:"plugin_dev_version,omitempty"`
	// PluginId 是插件id。
	PluginId *string `json:"plugin_id,omitempty"`
}

// MiniInnerAppPluginUseConfigUpgradeRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigUpgradeRequest struct {
	// AppOrigin 是来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序应用 ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PluginDevVersion 是插件研发版本。
	PluginDevVersion *string `json:"plugin_dev_version,omitempty"`
	// PluginId 是插件id。
	PluginId *string `json:"plugin_id,omitempty"`
	// StrategyValue 是灰度值。
	StrategyValue *string `json:"strategy_value,omitempty"`
}

// MiniInnerAppServicePublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppServicePublishRequest struct {
	// AppLogo 是服务发布logo。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序模板所属PID。
	Pid *string `json:"pid,omitempty"`
	// SubTitle 是服务子标题。
	SubTitle *string `json:"sub_title,omitempty"`
	// Title 是服务名称。
	Title *string `json:"title,omitempty"`
}

// MiniInnerAppServiceQueryRequest is the request type for alipay.open.mini.innerapp.service.query.
type MiniInnerAppServiceQueryRequest struct {
	// MiniAppId 是插件id或模板id
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppSubType 是小程序应用子类型
	AppSubType *string `json:"app_sub_type,omitempty" query:"app_sub_type,omitempty"`
	// ShowType 是是否在服务市场透出，SHOW展示、HIDE隐藏
	ShowType *string `json:"show_type,omitempty" query:"show_type,omitempty"`
	// Keyword 是关键词
	Keyword *string `json:"keyword,omitempty" query:"keyword,omitempty"`
	// PageNum 是页码，默认第一页
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是分页数，默认每页10个
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
	// AppOrigin 是业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
	// IncludeOffline 是是否需要展示offline状态插件，不传默认为false（不展示）
	IncludeOffline *string `json:"include_offline,omitempty" query:"include_offline,omitempty"`
}
