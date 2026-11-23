package logisticscompany

// LogisticsLogisticsCompanyInstantDeliveryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsLogisticsCompanyInstantDeliveryQueryResponse struct {
	// LogisticsCompanies 是即时配送公司列表
	LogisticsCompanies *[]LogisticsCompanyResult `json:"logistics_companies,omitempty"`
}
