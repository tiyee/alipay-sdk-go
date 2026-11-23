package logisticscompany

// LogisticsCompanyResult 是支付宝官方 V3 OpenAPI 模型。
type LogisticsCompanyResult struct {
	// LogisticsCode 是物流公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsName 是物流公司名称
	LogisticsName *string `json:"logistics_name,omitempty"`
}
