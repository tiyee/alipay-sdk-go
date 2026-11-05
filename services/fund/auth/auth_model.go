package auth

// CreditRiskInfo 是支付宝官方 V3 OpenAPI 模型。
type CreditRiskInfo struct {
	// CanRetry 是查询失败的情况下，是否可以通过重试解决。true 表示可以适当重试一定次数；false 表示有明确的失败原因，无需重试，。
	CanRetry *bool `json:"can_retry,omitempty"`
	// Code 是用户信用数据查询状态。
	Code *string `json:"code,omitempty"`
	// Data 是用户信用风险信息，如：租物场景风险等级、履约历史等级、共租行为等级。
	Data *string `json:"data,omitempty"`
}

// PostPayment 是支付宝官方 V3 OpenAPI 模型。
type PostPayment struct {
	// Amount 是后付费金额，单位为：元（人民币），精确到小数点后两位。。
	Amount *string `json:"amount,omitempty"`
	// Description 是计费说明。
	Description *string `json:"description,omitempty"`
	// Name 是后付费项目名称。
	Name *string `json:"name,omitempty"`
}
