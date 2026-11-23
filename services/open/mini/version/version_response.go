package version

// MiniVersionAuditedCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditedCancelResponse struct {
}

// MiniVersionBuildQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionBuildQueryResponse struct {
	// BuildStatus 是构建的状态，0-构建排队中；1-正在构建；2-构建成功；3-构建失败；5-构建超时。
	BuildStatus *string `json:"build_status,omitempty"`
	// CreateStatus 是创建版本的状态，状态说明如下： 0：构建排队中； 1：正在构建； 2：构建成功； 3：构建失败； 5：构建超时； 6：版本创建成功。。
	CreateStatus *string `json:"create_status,omitempty"`
	// NeedRotation 是是否需要轮询。
	NeedRotation *string `json:"need_rotation,omitempty"`
	// VersionCreated 是是否创建版本。
	VersionCreated *string `json:"version_created,omitempty"`
}

// MiniVersionDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionDeleteResponse struct {
}

// MiniVersionDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionDetailQueryResponse struct {
	// AppDesc 是小程序应用描述。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppLogo 是小程序应用logo图标。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppName 是小程序应用名称。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是小程序应用简介，一句话描述小程序功能。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// AppVersion 是小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BaseAudit 是准入审核结果 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	BaseAudit       *string             `json:"base_audit,omitempty"`
	BaseAuditRecord *MiniAppAuditReason `json:"base_audit_record,omitempty"`
	// CanRelease 是是否可上架 true：可上架 false：不可上架 （version_status 为PROMOTE_AUDIT_REJECT有值） （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	CanRelease *string `json:"can_release,omitempty"`
	// GmtApplyAudit 是小程序版本提交审核的时间，格式为yyyy-MM-dd HH:mm:ss。
	GmtApplyAudit *string `json:"gmt_apply_audit,omitempty"`
	// GmtAuditEnd 是小程序版本审核结束时间，格式为yyyy-MM-dd HH:mm:ss。
	GmtAuditEnd *string `json:"gmt_audit_end,omitempty"`
	// GmtCreate 是小程序版本创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtOffline 是小程序版本下架时间。
	GmtOffline *string `json:"gmt_offline,omitempty"`
	// GmtOnline 是小程序版本上架时间。
	GmtOnline *string `json:"gmt_online,omitempty"`
	// GrayStrategy 是小程序灰度比例值，p10-10%的用户，p30-30%的用户，p50-50%的用户。
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// Memo 是小程序备注。
	Memo *string `json:"memo,omitempty"`
	// MiniAppCategoryInfoList 是小程序类目列表。
	MiniAppCategoryInfoList *[]MiniAppCategoryInfo `json:"mini_app_category_info_list,omitempty"`
	// MiniCategoryInfoList 是新小程序类目列表。
	MiniCategoryInfoList *[]MiniAppCategoryInfo `json:"mini_category_info_list,omitempty"`
	// PackageInfoList 是小程序功能包列表。
	PackageInfoList *[]MiniPackageInfo `json:"package_info_list,omitempty"`
	// PromoteAudit 是营销审核结果： PASS：通过 REJECT：驳回 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	PromoteAudit       *string             `json:"promote_audit,omitempty"`
	PromoteAuditRecord *MiniAppAuditReason `json:"promote_audit_record,omitempty"`
	// RejectReason 是审核驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// ScanResult 是安全扫描结果，True-安全扫描通过，False-安全扫描不通过，NAN-扫描中。
	ScanResult *string `json:"scan_result,omitempty"`
	// ScreenShotList 是小程序版本截图，多个截图以逗号隔开。
	ScreenShotList *[]string `json:"screen_shot_list,omitempty"`
	// ServiceEmail 是客服邮箱。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是客服电话。
	ServicePhone *string `json:"service_phone,omitempty"`
	// ServiceRegionInfo 是小程序服务区域信息，当服务区域类型是LOCATION时，这里明确指出具体的服务区域信息。
	ServiceRegionInfo *[]RegionInfo `json:"service_region_info,omitempty"`
	// ServiceRegionType 是小程序服务区域类型。
	ServiceRegionType *string `json:"service_region_type,omitempty"`
	// Status 是小程序版本状态，INIT-开发中，AUDITING-审核中，WAIT_RELEASE-待上架，BASE_AUDIT_PASS-准入不可营销，AUDIT_REJECT-审核驳回，RELEASE-已上架，GRAY-灰度中，OFFLINE-已下架，AUDIT_OFFLINE-已被强制下架。
	Status *string `json:"status,omitempty"`
	// VersionDesc 是小程序版本描述。
	VersionDesc *string `json:"version_desc,omitempty"`
}

// MiniVersionListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionListQueryResponse struct {
	// AppVersionInfos 是版本列表，根据版本号倒叙排列，最多呈现最近 20 条。即版本号大的在前面；如果不存在任何版本，返回空列表。。
	AppVersionInfos *[]AppVersionInfo `json:"app_version_infos,omitempty"`
	// AppVersions 是小程序支付宝端每个状态的最新版本号列表(历史返回值，已不推荐使用，逐渐废弃)。
	AppVersions *[]string `json:"app_versions,omitempty"`
}

// MiniVersionOfflineResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionOfflineResponse struct {
}

// MiniVersionOnlineResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionOnlineResponse struct {
}

// MiniVersionRollbackResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionRollbackResponse struct {
}

// MiniVersionUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionUploadResponse struct {
	// BuildStatus 是构建的状态，0-构建排队中；1-正在构建；2-构建成功；3-构建失败；5-构建超时。
	BuildStatus *string `json:"build_status,omitempty"`
	// CreateStatus 是创建版本的状态，0-构建排队中；1-正在构建；2-构建成功；3-构建失败；5-构建超时；6-版本创建成功。
	CreateStatus *string `json:"create_status,omitempty"`
	// NeedRotation 是是否需要轮询。
	NeedRotation *string `json:"need_rotation,omitempty"`
	// VersionCreated 是是否创建了版本。
	VersionCreated *string `json:"version_created,omitempty"`
}
