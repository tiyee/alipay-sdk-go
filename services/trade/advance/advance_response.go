package advance

// AdvanceConsultResponse 是订单咨询的业务响应。
type AdvanceConsultResponse struct {
	// ReferResult 标识当前时间点用户是否允许垫资。
	ReferResult *bool `json:"refer_result,omitempty"`
	// ResultCode 是结果码（如用户被注销）。
	ResultCode *string `json:"result_code,omitempty"`
	// ResultMessage 是用户不准入原因。
	ResultMessage *string `json:"result_message,omitempty"`
	// RiskLevel 是订单风险评估等级；单笔订单风险预评估时返回。
	RiskLevel *string `json:"risk_level,omitempty"`
	// UserRiskPrediction 是用户风险预测结果。
	UserRiskPrediction *UserRiskPrediction `json:"user_risk_prediction,omitempty"`
	// WaitRepaymentAmount 是用户剩余的总待还金额。
	WaitRepaymentAmount *string `json:"wait_repayment_amount,omitempty"`
	// WaitRepaymentOrderCount 是用户总的未还垫资笔数。
	WaitRepaymentOrderCount *int64 `json:"wait_repayment_order_count,omitempty"`
	// WaitRepaymentOrderInfos 是待还订单列表。
	WaitRepaymentOrderInfos []WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"`
}
