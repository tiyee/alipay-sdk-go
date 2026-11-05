package employee

// ExpenseCtrlEmployeeRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseCtrlEmployeeRuleInfo struct {
	// Effective 是费控规则状态，1表示有效，0表示无效
	Effective *int64 `json:"effective,omitempty"`
	// EffectiveEndDate 是费控规则有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是费控规则有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// OwnerType 是费控规则作用范围，	EMPLOYEE表示员工，ENTERPRISE表示企业
	OwnerType *string `json:"owner_type,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是费控规则名称
	StandardName *string `json:"standard_name,omitempty"`
}
