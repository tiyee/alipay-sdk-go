package balance

// EcBalanceDownloadURLQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcBalanceDownloadURLQueryResponse struct {
	// BillDownloadUrl 是对账单下载地址链接
	BillDownloadUrl *string `json:"bill_download_url,omitempty"`
}

// EcBalancePeriodModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcBalancePeriodModifyResponse struct {
	// Result 是修改结果
	Result *bool `json:"result,omitempty"`
}
