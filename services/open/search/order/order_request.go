package order

// SearchBaseOrderModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBaseOrderModifyRequest struct {
	BizData *SearchBaseOrderCreateAPIRequest `json:"biz_data,omitempty"`
	// BizType 是表示修改内容的业务类型。
	BizType *string `json:"biz_type,omitempty"`
	// OptType 是操作的具体类型。
	OptType *string `json:"opt_type,omitempty"`
}

// SearchOrderDetailQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailQueryRequest struct {
	BizData *SearchOrderRequest `json:"biz_data,omitempty"`
	// BizType 是表示修改内容的业务类型。
	BizType *string `json:"biz_type,omitempty"`
	// OptType 是进行具体操作的类型比如：create、mod、cancel。
	OptType *string `json:"opt_type,omitempty"`
}

// SearchServiceOrderBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchServiceOrderBatchQueryRequest struct {
	BizData *SearchApplyPageQueryRequest `json:"biz_data,omitempty"`
	// BizType 是表示查询内容的类型。
	BizType *string `json:"biz_type,omitempty"`
	// OptType 是操作的类型，如query。
	OptType *string `json:"opt_type,omitempty"`
}

// SearchApplyPageQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchApplyPageQueryRequest struct {
	// ApplyType 是申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	ApplyType *string `json:"apply_type,omitempty"`
	// AuditStatusList 是筛选状态数组 AUDIT AGREE REJECT CANCEL EDIT。
	AuditStatusList *[]string `json:"audit_status_list,omitempty"`
	// CategoryCode 是类目编码。
	CategoryCode *string `json:"category_code,omitempty"`
	// Name 是服务名称。
	Name *string `json:"name,omitempty"`
	// PageNum 是当前页码。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是页面的显示记录条数。
	PageSize *string `json:"page_size,omitempty"`
	// ServiceCode 是服务code。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceId 是小程序appid。
	ServiceId *string `json:"service_id,omitempty"`
	// StartRow 是起始记录，起始：0。
	StartRow *string `json:"start_row,omitempty"`
	// SubServiceCode 是二级服务code。
	SubServiceCode *string `json:"sub_service_code,omitempty"`
}

// SearchBaseOrderCreateAPIRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBaseOrderCreateAPIRequest struct {
	// AccessType 是搜索直达类型 BASE：基础信息，只支持基础信息工单提报。
	AccessType *string `json:"access_type,omitempty"`
	// Appid 是小程序Id。
	Appid     *string          `json:"appid,omitempty"`
	BaseItems *SearchBaseItems `json:"base_items,omitempty"`
	// Descprise 是服务描述。
	Descprise *string `json:"descprise,omitempty"`
	// IsDraft 是是否为草稿态。
	IsDraft *bool `json:"is_draft,omitempty"`
	// OrderId 是申请单id，仅仅驳回或修改是传入。
	OrderId *string `json:"order_id,omitempty"`
	// SpecCode 是服务的类型 SP_MINI_APP 小程序 SP_PUBLIC_APP 生活号。
	SpecCode *string `json:"spec_code,omitempty"`
}

// SearchOrderRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderRequest struct {
	// ApplyId 是申请单id。
	ApplyId *string `json:"apply_id,omitempty"`
	// ApplyType 是申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	ApplyType *string `json:"apply_type,omitempty"`
	// BrandTemplateId 是品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四。
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// ServiceCode 是服务code。
	ServiceCode *string `json:"service_code,omitempty"`
	// SpecCode 是服务的类型默认使用小程序 SP_MINI_APP 小程序 SP_PUBLIC_APP 生活号。
	SpecCode *string `json:"spec_code,omitempty"`
}
