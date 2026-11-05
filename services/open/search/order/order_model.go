package order

// ApplyOrderData 是支付宝官方 V3 OpenAPI 模型。
type ApplyOrderData struct {
	// AccessType 是申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	AccessType *string `json:"access_type,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// Appid 是小程序id。
	Appid *string `json:"appid,omitempty"`
	// AuditReason 是审核原因。
	AuditReason *string `json:"audit_reason,omitempty"`
	// BrandTemplateId 是品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四。
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// CreateTime 是创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// IsOldData 是标识老工单。
	IsOldData *bool `json:"is_old_data,omitempty"`
	// Latest 是是不是最新的工单。
	Latest *bool `json:"latest,omitempty"`
	// MajorStatus 是服务主状态,INITIAL:初始化；VALID：已生效；INVALID：已失效。
	MajorStatus *string `json:"major_status,omitempty"`
	// OrderId 是申请单ID。
	OrderId *string `json:"order_id,omitempty"`
	// PartAgreeInfo 是审核提效部分剔除数据的信息。
	PartAgreeInfo *[]SearchPartAgreeInfo `json:"part_agree_info,omitempty"`
	// SceneCode 是场景编码。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneName 是场景名称。
	SceneName *string `json:"scene_name,omitempty"`
	// ServiceCode 是服务码。
	ServiceCode *string `json:"service_code,omitempty"`
	// Status 是申请状态 AGREE、REJECT、EDIT、AUDIT。
	Status *string `json:"status,omitempty"`
	// SubServiceName 是子服务名称。
	SubServiceName *string `json:"sub_service_name,omitempty"`
}

// BrandDetail 是支付宝官方 V3 OpenAPI 模型。
type BrandDetail struct {
	// BlockType 是搜索工单详情数据block_type。
	BlockType *string `json:"block_type,omitempty"`
	// BoxVersion 是版本号。
	BoxVersion *string `json:"box_version,omitempty"`
	// Desc 是描述信息。
	Desc *string `json:"desc,omitempty"`
	// Img 是img地址。
	Img *string `json:"img,omitempty"`
	// Index 是索引信息。
	Index *string `json:"index,omitempty"`
	// Key 是工单详情数据key。
	Key *string `json:"key,omitempty"`
	// MaterialId 是工单详情数据material_id。
	MaterialId *string `json:"material_id,omitempty"`
	// Name 是应用名称。
	Name *string `json:"name,omitempty"`
	// PublicBindAppId 是搜索运营工单详情数据public_bind_app_id。
	PublicBindAppId *string `json:"public_bind_app_id,omitempty"`
	// ServiceCode 是服务码。
	ServiceCode *string `json:"service_code,omitempty"`
}

// OrderPageQueryDTO 是支付宝官方 V3 OpenAPI 模型。
type OrderPageQueryDTO struct {
	// PageData 是审核单数据信息。
	PageData *[]ApplyOrderData `json:"page_data,omitempty"`
	// PageNum 是当前页码，默认第一页。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是每页记录数，默认10。
	PageSize *string `json:"page_size,omitempty"`
	// TotalNumber 是记录总数。
	TotalNumber *int64 `json:"total_number,omitempty"`
}

// OrderStatusData 是支付宝官方 V3 OpenAPI 模型。
type OrderStatusData struct {
	// OrderId 是申请单ID。
	OrderId *string `json:"order_id,omitempty"`
	// RejectReason 是驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// Status 是申请状态 AGREE、REJECT、EDIT、AUDIT。
	Status *string `json:"status,omitempty"`
}

// SearchBaseItems 是支付宝官方 V3 OpenAPI 模型。
type SearchBaseItems struct {
	// CanSearch 是搜索可见性。
	CanSearch *bool `json:"can_search,omitempty"`
	// KeyWords 是提报关键词列表。
	KeyWords *[]string `json:"key_words,omitempty"`
}

// SearchOrderBrandDetail 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderBrandDetail struct {
	BrandDetail *BrandDetail `json:"brand_detail,omitempty"`
	// Key 是工单详情key值。
	Key *string `json:"key,omitempty"`
}

// SearchOrderDetailData 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailData struct {
	// AccessType 是申请单类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	AccessType *string `json:"access_type,omitempty"`
	// AppCategory 是小程序类目名。
	AppCategory *string `json:"app_category,omitempty"`
	// AppIcon 是小程序图标。
	AppIcon *string `json:"app_icon,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// AppStatus 是小程序状态。
	AppStatus *string `json:"app_status,omitempty"`
	// Appid 是小程序ID。
	Appid           *string          `json:"appid,omitempty"`
	BaseOrderStatus *OrderStatusData `json:"base_order_status,omitempty"`
	// BizId 是bizid 业务id。
	BizId          *string          `json:"biz_id,omitempty"`
	BoxOrderStatus *OrderStatusData `json:"box_order_status,omitempty"`
	// BoxStatus 是上下架状态 ONLINE OFFLINE。
	BoxStatus *string `json:"box_status,omitempty"`
	// BrandTemplateId 是品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四。
	BrandTemplateId     *string          `json:"brand_template_id,omitempty"`
	BrandboxOrderstatus *OrderStatusData `json:"brandbox_orderstatus,omitempty"`
	// CanModifyKeyword 是是否可修改关键词。
	CanModifyKeyword *bool `json:"can_modify_keyword,omitempty"`
	// Describe 是服务描述。
	Describe           *string                            `json:"describe,omitempty"`
	DetailBaseItems    *SearchOrderDetailDataBaseItems    `json:"detail_base_items,omitempty"`
	DetailBrandItems   *SearchOrderDetailDataBrandItems   `json:"detail_brand_items,omitempty"`
	DetailServiceItems *SearchOrderDetailDataServiceItems `json:"detail_service_items,omitempty"`
	// FinalStatus 是最终状态。
	FinalStatus *string `json:"final_status,omitempty"`
	// Gmtmodified 是最后修改时间。
	Gmtmodified *string `json:"gmtmodified,omitempty"`
	// KeywordGmtModified 是关键词的修改时间。
	KeywordGmtModified *string `json:"keyword_gmt_modified,omitempty"`
	// OnlineTime 是上架时间。
	OnlineTime *string `json:"online_time,omitempty"`
	// OperatorType 是操作人员类型。
	OperatorType *string `json:"operator_type,omitempty"`
	// OrderId 是申请单id。
	OrderId *string `json:"order_id,omitempty"`
	// PartnerType 是查询服务提供者生态类型 一二三方 1 | 2 | 3。
	PartnerType *string `json:"partner_type,omitempty"`
	// RejectNode 是驳回节点。
	RejectNode *string `json:"reject_node,omitempty"`
	// RejectReason 是驳回原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// SceneCode 是场景码。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneName 是场景名称。
	SceneName *string `json:"scene_name,omitempty"`
	// ServiceCode 是服务code。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceName 是服务名。
	ServiceName *string `json:"service_name,omitempty"`
	// Status 是申请单状态 AGREE、REJECT、AUDIT。
	Status *string `json:"status,omitempty"`
	// SubServiceCode 是二级服务code。
	SubServiceCode *string `json:"sub_service_code,omitempty"`
	// TemplateId 是模板Id。
	TemplateId *string `json:"template_id,omitempty"`
	// TemplateName 是模板名称。
	TemplateName *string `json:"template_name,omitempty"`
}

// SearchOrderDetailDataBaseItems 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailDataBaseItems struct {
	// CanSearch 是搜索可见性。
	CanSearch *string `json:"can_search,omitempty"`
	// Desc 是描述信息。
	Desc *string `json:"desc,omitempty"`
	// Img 是img地址。
	Img *string `json:"img,omitempty"`
	// KeyWord 是关键词。
	KeyWord *string `json:"key_word,omitempty"`
	// Name 是应用名称。
	Name *string `json:"name,omitempty"`
	// Region 是基础工单详情region。
	Region *string `json:"region,omitempty"`
	// ServCanSearch 是搜索可见性。
	ServCanSearch *string `json:"serv_can_search,omitempty"`
	// ServSearchKeywords 是搜索关键词。
	ServSearchKeywords *string `json:"serv_search_keywords,omitempty"`
	// TemplateId 是模板id。
	TemplateId *string `json:"template_id,omitempty"`
	// Url 是跳转url链接。
	Url *string `json:"url,omitempty"`
}

// SearchOrderDetailDataBrandItems 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailDataBrandItems struct {
	// BizId 是工单详情bizid。
	BizId *string `json:"biz_id,omitempty"`
	// BoxStatus 是上下架状态。
	BoxStatus *string `json:"box_status,omitempty"`
	// BrandBoxKeywords 是关键词信息。
	BrandBoxKeywords *string `json:"brand_box_keywords,omitempty"`
	// BrandDetailList 是工单详情数据信息。
	BrandDetailList *[]SearchOrderBrandDetail `json:"brand_detail_list,omitempty"`
	// BrandTemplateId 是品牌展示模板类型。
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// Channel 是工单详情数据channel。
	Channel *string `json:"channel,omitempty"`
	// MerchantType 是工单详情数据merchant_type。
	MerchantType *string `json:"merchant_type,omitempty"`
	// TemplateId 是模板id。
	TemplateId *string `json:"template_id,omitempty"`
}

// SearchOrderDetailDataServiceItems 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailDataServiceItems struct {
	// AreaCodes 是服务工单详情area_codes。
	AreaCodes *string `json:"area_codes,omitempty"`
	// CarrierCode 是carrier_code。
	CarrierCode *string `json:"carrier_code,omitempty"`
	// CarrierList 是服务工单数据carrier_list。
	CarrierList *string `json:"carrier_list,omitempty"`
	// CategoryAttributeValue 是服务工单category_attribute_value。
	CategoryAttributeValue *string `json:"category_attribute_value,omitempty"`
	// CategoryCode 是类目编码。
	CategoryCode *string `json:"category_code,omitempty"`
	// CategoryIds 是类目id。
	CategoryIds *string `json:"category_ids,omitempty"`
	// ChannelType 是渠道类型。
	ChannelType *string `json:"channel_type,omitempty"`
	// Desc 是描述信息。
	Desc *string `json:"desc,omitempty"`
	// Img 是img地址。
	Img *string `json:"img,omitempty"`
	// KeyWord 是关键词。
	KeyWord *string `json:"key_word,omitempty"`
	// KeyWordList 是关键词列表。
	KeyWordList *string `json:"key_word_list,omitempty"`
	// Logo 是logo图片地址。
	Logo *string `json:"logo,omitempty"`
	// Name 是应用名称。
	Name *string `json:"name,omitempty"`
	// ParentServiceCode 是服务编码。
	ParentServiceCode *string `json:"parent_service_code,omitempty"`
	// Region 是服务工单详情region。
	Region *string `json:"region,omitempty"`
	// ServSearchCatalogs 是serv_search_catalogs。
	ServSearchCatalogs *string `json:"serv_search_catalogs,omitempty"`
	// ServSearchKeywords 是关键词。
	ServSearchKeywords *string `json:"serv_search_keywords,omitempty"`
	// ServiceCode 是2020112321000021082348。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceName 是服务名称。
	ServiceName *string `json:"service_name,omitempty"`
	// ServiceTimeEnd 是服务结束时间。
	ServiceTimeEnd *string `json:"service_time_end,omitempty"`
	// ServiceTimeStart 是服务开始时间。
	ServiceTimeStart *string `json:"service_time_start,omitempty"`
	// ShortDesc 是服务短描述。
	ShortDesc *string `json:"short_desc,omitempty"`
	// SpecCode 是应用标识。
	SpecCode *string `json:"spec_code,omitempty"`
	// TemplateId 是模板id。
	TemplateId *string `json:"template_id,omitempty"`
}

// SearchPartAgreeInfo 是支付宝官方 V3 OpenAPI 模型。
type SearchPartAgreeInfo struct {
	// AuditInfo 是剔除数据名称。
	AuditInfo *string `json:"audit_info,omitempty"`
	// AuditOperator 是审核operator信息。
	AuditOperator *string `json:"audit_operator,omitempty"`
	// AuditReason 是审核原因。
	AuditReason *string `json:"audit_reason,omitempty"`
	// AuditType 是剔除数据类型。
	AuditType *string `json:"audit_type,omitempty"`
}
