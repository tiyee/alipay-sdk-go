package expenserules

// ExpenseCtrRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseCtrRuleInfo struct {
	// RuleFactor 是费控维度
	RuleFactor *string `json:"rule_factor,omitempty"`
	// RuleId 是费控条件ID
	RuleId *string `json:"rule_id,omitempty"`
	// RuleName 是费控条件名称
	RuleName *string `json:"rule_name,omitempty"`
	// RuleOperator 是费控条件操作符	枚举值：	LT("=","大于等于")	GT(">","大于")
	RuleOperator *string `json:"rule_operator,omitempty"`
	// RuleValue 是费控条件值
	RuleValue *string `json:"rule_value,omitempty"`
}
