package box

// BackgroundWordInfo 是支付宝官方 V3 OpenAPI 模型。
type BackgroundWordInfo struct {
	// BackgroundWord 是底纹词。
	BackgroundWord *string `json:"background_word,omitempty"`
	// FailReason 是审核失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// Status 是底纹词状态。
	Status *string `json:"status,omitempty"`
}

// BoxBusinessDistrictInfo 是支付宝官方 V3 OpenAPI 模型。
type BoxBusinessDistrictInfo struct {
	// AppName 是应用名。
	AppName *string `json:"app_name,omitempty"`
	// BusinessDistrictId 是商圈ID。
	BusinessDistrictId *string `json:"business_district_id,omitempty"`
	// BusinessDistrictName 是商圈名称。
	BusinessDistrictName *string `json:"business_district_name,omitempty"`
	// RelateAppid 是商圈所属小程序对应的服务编码。
	RelateAppid *string `json:"relate_appid,omitempty"`
	// ServiceCode 是商圈所属小程序对应的服务编码。
	ServiceCode *string `json:"service_code,omitempty"`
}

// BoxBusinessDistrictModule 是支付宝官方 V3 OpenAPI 模型。
type BoxBusinessDistrictModule struct {
	// BusinessDistrictInfos 是商圈选择。
	BusinessDistrictInfos *[]BoxBusinessDistrictInfo `json:"business_district_infos,omitempty"`
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
}

// DeliveryOrientedRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryOrientedRuleInfo struct {
	// ClientPlatform 是客户端。
	ClientPlatform *string `json:"client_platform,omitempty"`
}

// DeliveryTargetRegion 是支付宝官方 V3 OpenAPI 模型。
type DeliveryTargetRegion struct {
	// RegionCode 是区域编码，参考<a href="https://opendocs.alipay.com/pre-open/03144c" target="_blank">取值说明</a>。
	RegionCode *string `json:"region_code,omitempty"`
	// RegionName 是区域名，参考<a href="https://opendocs.alipay.com/pre-open/03144c" target="_blank">取值说明</a>。
	RegionName *string `json:"region_name,omitempty"`
	// RegionType 是区域类型，1代表全国，2代表省级区域，3代表市级区域，参考<a href="https://opendocs.alipay.com/pre-open/03144c" target="_blank">取值说明</a>。
	RegionType *string `json:"region_type,omitempty"`
}

// SearchBoxAccountModule 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxAccountModule struct {
	// AppInfos 是应用信息列表。
	AppInfos *[]SearchBoxAppInfo `json:"app_infos,omitempty"`
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
}

// SearchBoxActivityVideoInfo 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxActivityVideoInfo struct {
	// VideoBtnText 是视频按钮文案，2-4个汉字。
	VideoBtnText *string `json:"video_btn_text,omitempty"`
	// VideoCoverimgId 是视频封面图片id，图片id可以通过调用接口alipay.open.file.upload上传图片，获取图片id。图片规范：https://opendocs.alipay.com/mini/operation/atmospheredesign。
	VideoCoverimgId *string `json:"video_coverimg_id,omitempty"`
	// VideoSubTitle 是视频副标题，2-16个汉字。
	VideoSubTitle *string `json:"video_sub_title,omitempty"`
	// VideoTitle 是视频主标题，2-10个汉字。
	VideoTitle *string `json:"video_title,omitempty"`
}

// SearchBoxAppInfo 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxAppInfo struct {
	// AppName 是应用名称。
	AppName *string `json:"app_name,omitempty"`
	// AppType 是应用类型，TINY_APP-小程序/PUBLIC_APP-生活号。
	AppType *string `json:"app_type,omitempty"`
	// RelateAppid 是应用ID。
	RelateAppid *string `json:"relate_appid,omitempty"`
}

// SearchBoxBasicInfo 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxBasicInfo struct {
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// BrandId 是品牌id。
	BrandId *string `json:"brand_id,omitempty"`
	// Name 是搜索直达名称。
	Name *string `json:"name,omitempty"`
	// Status 是搜索直达配置状态，INITIAL-初始/ONLINE-已上架/EXPIRE-已失效/OFFLINE-已下架。
	Status *string `json:"status,omitempty"`
	// TargetAppid 是小程序id。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchBoxBasicInfoModule 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxBasicInfoModule struct {
	// ApplyNo 是申请单号。
	ApplyNo *string `json:"apply_no,omitempty"`
	// BoxDesc 是介绍，已生效的数据。
	BoxDesc *string `json:"box_desc,omitempty"`
	// FailReason 是审核失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// GmtModified 是修改时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// LatestAuditBoxDesc 是介绍，最近一次审核内容。
	LatestAuditBoxDesc *string `json:"latest_audit_box_desc,omitempty"`
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
	// Status 是状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效。
	Status *string `json:"status,omitempty"`
}

// SearchBoxImageModule 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxImageModule struct {
	// ApplyNo 是申请单号。
	ApplyNo *string `json:"apply_no,omitempty"`
	// FailReason 是审核失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// GmtModified 是修改时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
	// Status 是状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效。
	Status *string `json:"status,omitempty"`
}

// SearchBoxKeyWordModule 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxKeyWordModule struct {
	// ApplyNo 是申请单号。
	ApplyNo *string `json:"apply_no,omitempty"`
	// FailReason 是审核失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// GmtModified 是修改时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Keywords 是自定义触发词(生效中)。
	Keywords *[]string `json:"keywords,omitempty"`
	// LatestAuditKeywords 是自定义触发词(最近一次审核内容)。
	LatestAuditKeywords *[]string `json:"latest_audit_keywords,omitempty"`
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
	// Status 是状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效。
	Status *string `json:"status,omitempty"`
}

// SearchBoxServiceInfo 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxServiceInfo struct {
	// ServiceCategoryCode 是服务类目编码。
	ServiceCategoryCode *string `json:"service_category_code,omitempty"`
	// ServiceCategoryName 是服务类目名称。
	ServiceCategoryName *string `json:"service_category_name,omitempty"`
	// ServiceCode 是服务编码。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceName 是服务名称。
	ServiceName *string `json:"service_name,omitempty"`
}

// SearchBoxServiceModule 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxServiceModule struct {
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型。
	ModuleType *string `json:"module_type,omitempty"`
	// ServiceInfos 是服务信息列表。
	ServiceInfos *[]SearchBoxServiceInfo `json:"service_infos,omitempty"`
}
