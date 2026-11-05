package search

// AbilityPageQueryDTO 是支付宝官方 V3 OpenAPI 模型。
type AbilityPageQueryDTO struct {
	// PageData 是搜索运营列表的分页数据。
	PageData *[]SearchAbilityOrderData `json:"page_data,omitempty"`
	// PageNum 是当前页码，默认第一页。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是每页记录数，默认10。
	PageSize *string `json:"page_size,omitempty"`
	// TotalNumber 是记录总数。
	TotalNumber *int64 `json:"total_number,omitempty"`
}

// KeyWordInfo 是支付宝官方 V3 OpenAPI 模型。
type KeyWordInfo struct {
	// ApplyNo 是审核单号。
	ApplyNo *string `json:"apply_no,omitempty"`
	// AuditKeyWord 是审核中的关键词。
	AuditKeyWord *string `json:"audit_key_word,omitempty"`
	// AuditReason 是审核人的备注信息。
	AuditReason *string `json:"audit_reason,omitempty"`
	// AuditTime 是审核时间，格式为 yyyy-MM-dd HH:mm:ss。
	AuditTime *string `json:"audit_time,omitempty"`
	// ConfigId 是关键词配置ID。
	ConfigId *string `json:"config_id,omitempty"`
	// GmtCreate 是创建时间，格式为 yyyy-MM-dd HH:mm:ss。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// KeyWord 是审核完成的关键词。
	KeyWord *string `json:"key_word,omitempty"`
	// Status 是关键词当前状态。
	Status *string `json:"status,omitempty"`
}

// SearchAbilityOrderData 是支付宝官方 V3 OpenAPI 模型。
type SearchAbilityOrderData struct {
	// AccessType 是申请单类型。
	AccessType *string `json:"access_type,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// AppStatus 是小程序状态 ON_LINE、OFF_LINE。
	AppStatus *string `json:"app_status,omitempty"`
	// Appid 是小程序id。
	Appid *string `json:"appid,omitempty"`
	// ApplyId 是申请单id。
	ApplyId *string `json:"apply_id,omitempty"`
	// ApplyType 是申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	ApplyType *string `json:"apply_type,omitempty"`
	// AuditStatus 是申请状态 AGREE、REJECT、EDIT、AUDIT。
	AuditStatus *string `json:"audit_status,omitempty"`
	// BizId 是bizId 为品牌box的boxID。
	BizId *string `json:"biz_id,omitempty"`
	// BoxStatus 是品牌box的上下架状态 ONLINE OFFLINE。
	BoxStatus *string `json:"box_status,omitempty"`
	// BrandTemplateId 是品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四。
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// Children 是二级服务信息。
	Children *[]SearchAbilityOrderInfoOpenAPI `json:"children,omitempty"`
	// DataKey 是二级服务的唯一标识。
	DataKey *string `json:"data_key,omitempty"`
	// GmtModified 是更新时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Id 是服务的唯一标识（优先使用serviceCode可忽略）。
	Id *string `json:"id,omitempty"`
	// IsOldData 是是否是老工单。
	IsOldData *string `json:"is_old_data,omitempty"`
	// MajorStatus 是服务主状态,INITIAL:初始化；VALID：已生效；INVALID：已失效。
	MajorStatus *string `json:"major_status,omitempty"`
	// OnlineTime 是上架时间。
	OnlineTime *string `json:"online_time,omitempty"`
	// OpenStatus 是服务的可见性状态。
	OpenStatus *bool `json:"open_status,omitempty"`
	// Operator 是下架操作者。
	Operator *string `json:"operator,omitempty"`
	// RejectReason 是驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// SceneCode 是场景码。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneName 是场景名称。
	SceneName *string `json:"scene_name,omitempty"`
	// ServiceCode 是服务码。
	ServiceCode *string `json:"service_code,omitempty"`
	// SubServiceDesc 是子功能描述。
	SubServiceDesc *string `json:"sub_service_desc,omitempty"`
	// SubServiceName 是子功能名称。
	SubServiceName *string `json:"sub_service_name,omitempty"`
}

// SearchAbilityOrderInfoOpenAPI 是支付宝官方 V3 OpenAPI 模型。
type SearchAbilityOrderInfoOpenAPI struct {
	// AccessType 是申请单类型。
	AccessType *string `json:"access_type,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// AppStatus 是小程序状态。
	AppStatus *string `json:"app_status,omitempty"`
	// AuditStatus 是审核状态。
	AuditStatus *string `json:"audit_status,omitempty"`
	// BizId 是运营申请单详情biz_id。
	BizId *string `json:"biz_id,omitempty"`
	// BoxStatus 是上架状态。
	BoxStatus *string `json:"box_status,omitempty"`
	// BrandTemplateId 是品牌模板id。
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// DataKey 是二级服务唯一标识。
	DataKey *string `json:"data_key,omitempty"`
	// GmtModified 是最近更新时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Id 是唯一id。
	Id *string `json:"id,omitempty"`
	// IsOldData 是是否老工单。
	IsOldData *bool `json:"is_old_data,omitempty"`
	// MajorStatus 是服务主状态。
	MajorStatus *string `json:"major_status,omitempty"`
	// OnlineTime 是上架时间。
	OnlineTime *string `json:"online_time,omitempty"`
	// OpenStatus 是小程序可见性。
	OpenStatus *bool `json:"open_status,omitempty"`
	// Operator 是下架操作者。
	Operator *string `json:"operator,omitempty"`
	// RejectReason 是驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// SceneCode 是场景码。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneName 是场景名称。
	SceneName *string `json:"scene_name,omitempty"`
	// SearchAppId 是小程序id。
	SearchAppId *string `json:"search_app_id,omitempty"`
	// SepcCode 是应用标识。
	SepcCode *string `json:"sepc_code,omitempty"`
	// ServiceCode 是服务码。
	ServiceCode *string `json:"service_code,omitempty"`
	// SubServiceDesc 是子功能描述。
	SubServiceDesc *string `json:"sub_service_desc,omitempty"`
	// SubServiceName 是子功能名称。
	SubServiceName *string `json:"sub_service_name,omitempty"`
}
