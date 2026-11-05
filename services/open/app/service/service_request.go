package service

// AppServiceApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type AppServiceApplyRequest struct {
	// CategoryId 是行业类目id,获取请参考<a href="https://opendocs.alipay.com/mini/03ci0w?pathHash=ed3c875c">各个行业场景服务接入资料</a>。
	CategoryId *string `json:"category_id,omitempty"`
	// OutBizNo 是外部业务编号,平台会基于appId+out_biz_no做幂等控制，如果出现幂等，会返回幂等的service_code。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// SchemaVersion 是服务schema版本。
	SchemaVersion *string `json:"schema_version,omitempty"`
	// ServiceCode 是服务编码, 首次提报时设置空值，支付宝侧返回ServiceCode；当传入serviceCode，则对已提报服务做编辑操作。。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceXml 是服务xml,格式请参考<a href="https://opendocs.alipay.com/mini/03cj40?pathHash=a5de4bd9">Schema 规则介绍</a>。
	ServiceXml *string `json:"service_xml,omitempty"`
	// TemplateType 是默认值:DEFAULT。。
	TemplateType *string `json:"template_type,omitempty"`
	// UserServiceDeliveryType 是商户提报服务履约类型。
	UserServiceDeliveryType *string `json:"user_service_delivery_type,omitempty"`
}

// AppServiceDeleteRequest is the request type for alipay.open.app.service.delete.
type AppServiceDeleteRequest struct {
	// ServiceCode 是服务id
	ServiceCode *string `json:"service_code,omitempty" query:"service_code,omitempty"`
}

// AppServiceListQueryRequest is the request type for alipay.open.app.service.list.query.
type AppServiceListQueryRequest struct {
	// ServiceName 是服务名称
	ServiceName *string `json:"service_name,omitempty" query:"service_name,omitempty"`
	// CategoryId 是类目
	CategoryId *string `json:"category_id,omitempty" query:"category_id,omitempty"`
	// PageSize 是每页记录数
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,omitempty"`
	// PageNum 是页码
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,omitempty"`
}

// AppServiceQueryRequest is the request type for alipay.open.app.service.query.
type AppServiceQueryRequest struct {
	// ServiceCode 是服务id
	ServiceCode *string `json:"service_code,omitempty" query:"service_code,omitempty"`
}

// AppServiceSchemaQueryRequest is the request type for alipay.open.app.service.schema.query.
type AppServiceSchemaQueryRequest struct {
	// CategoryId 是类目id, 获取请参考各个行业场景服务接入资料
	CategoryId *string `json:"category_id,omitempty" query:"category_id,omitempty"`
	// TemplateType 是服务模版类型，默认值:DEFAULT
	TemplateType *string `json:"template_type,omitempty" query:"template_type,omitempty"`
}
