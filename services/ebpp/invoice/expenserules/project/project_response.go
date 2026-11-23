package project

// InvoiceExpenseRulesProjectEmployeeModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectEmployeeModifyResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}

// InvoiceExpenseRulesProjectInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectInfoModifyResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}

// InvoiceExpenseRulesProjectRuleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectRuleCreateResponse struct {
	// ProjectId 是项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

// InvoiceExpenseRulesProjectRuleQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectRuleQueryResponse struct {
	// PageNum 是当前页数
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是当前记录数
	PageSize *string `json:"page_size,omitempty"`
	// ProjectRuleInfoList 是项目规则列表
	ProjectRuleInfoList *[]ProjectRuleInfo `json:"project_rule_info_list,omitempty"`
	// TotalCount 是总记录数
	TotalCount *string `json:"total_count,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *string `json:"total_page_count,omitempty"`
}

// InvoiceExpenseRulesProjectRulesModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectRulesModifyResponse struct {
	// Success 是是否成功
	Success *string `json:"success,omitempty"`
}
