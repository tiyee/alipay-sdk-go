// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecomsue

// InvoiceExpenseComsueOutsourceNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseComsueOutsourceNotifyRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Amount 是消费金额
	Amount *int64 `json:"amount,omitempty"`
	// DealTime 是交易发生时间
	DealTime *string `json:"deal_time,omitempty"`
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeIdType 是员工账号类型
	EmployeeIdType *int64 `json:"employee_id_type,omitempty"`
	// EmployeeOpenId 是员工开放id
	EmployeeOpenId *string `json:"employee_open_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Extend 是拓展参数
	Extend *string `json:"extend,omitempty"`
	// IsOffSet 是交易类型
	IsOffSet *int64 `json:"is_off_set,omitempty"`
	// OutSourceId 是外部交易支付单号
	OutSourceId *string `json:"out_source_id,omitempty"`
	// Platform 是外部平台编码
	Platform *string `json:"platform,omitempty"`
	// RelateNo 是外部交易退款单号
	RelateNo *string `json:"relate_no,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
}
