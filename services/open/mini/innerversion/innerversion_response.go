package innerversion

// MiniInnerVersionBackDevPublishResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionBackDevPublishResponse struct {
}

// MiniInnerVersionConditionBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionConditionBatchQueryResponse struct {
	// MiniVersionBaseInfoList 是小程序版本信息。
	MiniVersionBaseInfoList *[]MiniVersionBaseInfo `json:"mini_version_base_info_list,omitempty"`
}

// MiniInnerVersionContentRollbackResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionContentRollbackResponse struct {
}

// MiniInnerVersionInfoDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionInfoDeleteResponse struct {
}

// MiniInnerVersionInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionInfoQueryResponse struct {
	// AppVersion 是版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BuildSource 是版本构建类型。
	BuildSource *string `json:"build_source,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// CreateTime 是版本创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// Description 是版本描述。
	Description *string `json:"description,omitempty"`
	// ExtJson 是实例化模板参数。
	ExtJson *string `json:"ext_json,omitempty"`
	// GmtModified 是版本更新时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// GrayStartTime 是版本灰度时间。
	GrayStartTime *string `json:"gray_start_time,omitempty"`
	// GrayStrategy 是小程序灰度比例值。
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// OfflineTime 是版本下架时间。
	OfflineTime *string `json:"offline_time,omitempty"`
	// PackageUrl 是包地址。
	PackageUrl *string `json:"package_url,omitempty"`
	// PluginRefs 是引用插件信息。
	PluginRefs *[]MiniAppPluginInfo `json:"plugin_refs,omitempty"`
	// ProcessTaskId 是流程id。
	ProcessTaskId *string `json:"process_task_id,omitempty"`
	// RejectReason 是审核失败原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// RollbackTime 是版本回滚时间。
	RollbackTime *string `json:"rollback_time,omitempty"`
	// ScanResult 是安全扫描结果。
	ScanResult *string `json:"scan_result,omitempty"`
	// ScreenShotList 是小程序版本截图，多个截图以逗号隔开。
	ScreenShotList *string `json:"screen_shot_list,omitempty"`
	// ShelfTime 是版本上架时间。
	ShelfTime *string `json:"shelf_time,omitempty"`
	// SourceUrl 是源码包地址。
	SourceUrl *string `json:"source_url,omitempty"`
	// Status 是版本状态。
	Status *string `json:"status,omitempty"`
	// TemplateExtra 是模板构建参数。
	TemplateExtra *string `json:"template_extra,omitempty"`
	// TemplateVersion 是模板版本号。
	TemplateVersion *string `json:"template_version,omitempty"`
}

// MiniInnerVersionInstantiationQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionInstantiationQueryResponse struct {
	// AppVersion 是小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// TemplateId 是模板id。
	TemplateId *string `json:"template_id,omitempty"`
	// TemplateVersion 是模板的版本号。
	TemplateVersion *string `json:"template_version,omitempty"`
}

// MiniInnerVersionLastQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionLastQueryResponse struct {
	// AppDesc 是小程序应用描述。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppName 是小程序的名称。
	AppName *string `json:"app_name,omitempty"`
	// AppVersion 是版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// CategoryIds 是小程序所属类目。
	CategoryIds *string `json:"category_ids,omitempty"`
	// EnglishName 是小程序英文名称。
	EnglishName *string `json:"english_name,omitempty"`
	// LogoUrl 是小程序应用logo图标。
	LogoUrl *string `json:"logo_url,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// ServicePhone 是小程序客服电话。
	ServicePhone *string `json:"service_phone,omitempty"`
	// Slogan 是小程序应用简介，一句话描述小程序功能。
	Slogan *string `json:"slogan,omitempty"`
	// Status 是版本状态。
	Status *string `json:"status,omitempty"`
	// SubApplicationType 是小程序子类型，自研、模板类型。
	SubApplicationType *string `json:"sub_application_type,omitempty"`
}

// MiniInnerVersionModelForOfflineQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionModelForOfflineQueryResponse struct {
	// ModelJson 是线上版本包数据。
	ModelJson *string `json:"model_json,omitempty"`
	// SyncId 是同步ID，同步失败时便于去线上追溯。
	SyncId *string `json:"sync_id,omitempty"`
}

// MiniInnerVersionOfflinePublishResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOfflinePublishResponse struct {
}

// MiniInnerVersionOnlinePublishResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOnlinePublishResponse struct {
}

// MiniInnerVersionOnlineResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOnlineResponse struct {
}

// MiniInnerVersionPackageCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionPackageCreateResponse struct {
}

// MiniInnerVersionSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionSyncResponse struct {
}
