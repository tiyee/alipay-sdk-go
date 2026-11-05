package alipaypoint

// UserAlipayPointBudgetLibQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAlipayPointBudgetLibQueryResponse struct {
	// BudgetCode 是集分宝预算库编码
	BudgetCode *string `json:"budget_code,omitempty"`
	// BudgetDesc 是预算库的中文描述
	BudgetDesc *string `json:"budget_desc,omitempty"`
	// CumulativeAmount 是预算库累计积分数
	CumulativeAmount *int64 `json:"cumulative_amount,omitempty"`
	// Enabled 是预算库是否为生效状态标识，true-生效中，false-已失效
	Enabled *bool `json:"enabled,omitempty"`
	// EndTime 是预算库结束时间, yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"end_time,omitempty"`
	// RemainAmount 是预算库当前剩余积分数
	RemainAmount *int64 `json:"remain_amount,omitempty"`
	// StartTime 是预算库开始时间, yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"start_time,omitempty"`
}

// UserAlipayPointSendResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAlipayPointSendResponse struct {
	// RecordId 是发放记录号。
	RecordId *string `json:"record_id,omitempty"`
	// TransactionId 是对账中心关联订单，一般场景无需关注
	TransactionId *string `json:"transaction_id,omitempty"`
}
