package service

// StandardServiceBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type StandardServiceBaseInfo struct {
	// BizStatus 是服务状态。
	BizStatus *string `json:"biz_status,omitempty"`
	// CategoryId 是类目id。
	CategoryId *string `json:"category_id,omitempty"`
	// ServiceCode 是服务code。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceName 是服务名称。
	ServiceName *string `json:"service_name,omitempty"`
}
