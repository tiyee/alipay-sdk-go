package service

// AppServiceApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type AppServiceApplyResponse struct {
	// ServiceCode 是服务id。
	ServiceCode *string `json:"service_code,omitempty"`
}

// AppServiceDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type AppServiceDeleteResponse struct {
}

// AppServiceListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppServiceListQueryResponse struct {
	// PageNum 是页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录数。
	PageSize *int64 `json:"page_size,omitempty"`
	// Records 是服务列表。
	Records *[]StandardServiceBaseInfo `json:"records,omitempty"`
	// Total 是总记录数。
	Total *int64 `json:"total,omitempty"`
}

// AppServiceQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppServiceQueryResponse struct {
	// RejectReason 是驳回理由。
	RejectReason *string `json:"reject_reason,omitempty"`
	// ServiceCode 是服务id。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceXml 是服务描述xml格式。
	ServiceXml *string `json:"service_xml,omitempty"`
	// Status 是服务状态。
	Status *string `json:"status,omitempty"`
	// TemplateType 是服务模板类型。
	TemplateType *string `json:"template_type,omitempty"`
	// UserServiceDeliveryType 是服务履约类型。
	UserServiceDeliveryType *string `json:"user_service_delivery_type,omitempty"`
}

// AppServiceSchemaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppServiceSchemaQueryResponse struct {
	// CategoryId 是类目id。
	CategoryId *string `json:"category_id,omitempty"`
	// SchemaVersion 是服务schema版本。
	SchemaVersion *string `json:"schema_version,omitempty"`
	// SchemaXml 是服务schema。
	SchemaXml *string `json:"schema_xml,omitempty"`
	// TemplateType 是服务模版类型，默认值: DEFAULT。
	TemplateType *string `json:"template_type,omitempty"`
	// UserServiceDeliveryType 是商户提报服务履约类型。
	UserServiceDeliveryType *string `json:"user_service_delivery_type,omitempty"`
}
