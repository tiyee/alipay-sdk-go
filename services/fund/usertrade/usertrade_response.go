package usertrade

// FundUserTradeBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundUserTradeBatchQueryResponse struct {
	// EndTime 是拉取的结束时间，格式yyyy-MM-dd HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`
	// HasNext 是是否有下一页。
	HasNext *bool `json:"has_next,omitempty"`
	// PageIndex 是入参中的分页页码。
	PageIndex *string `json:"page_index,omitempty"`
	// PageSize 是入参中的分页大小。
	PageSize *string `json:"page_size,omitempty"`
	// StartTime 是拉取的起始时间，格式yyyy-MM-dd HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`
	// TotalCount 是交易时间区间内的总交易数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// TradeInfoList 是用户交易数据。
	TradeInfoList *[]UserTradeInfoDTO `json:"trade_info_list,omitempty"`
}
