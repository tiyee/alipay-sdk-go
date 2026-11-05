package employee

// InvoiceExpenseRulesEmployeeQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesEmployeeQueryResponse struct {
	// EmployeeList 是切换open_id前请使用此字段：员工列表
	EmployeeList *[]string `json:"employee_list,omitempty"`
	// EmployeeOpenIdList 是切换open_id后请使用此字段：员工open_id列表
	EmployeeOpenIdList *[]string `json:"employee_open_id_list,omitempty"`
	// PageNum 是当前页数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是当前记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}

// InvoiceExpenseRulesEmployeeRulesQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesEmployeeRulesQueryResponse struct {
	// EmployeeRules 是员工已绑定费控规则列表
	EmployeeRules *[]ExpenseCtrlEmployeeRuleInfo `json:"employee_rules,omitempty"`
	// PageNum 是客户端输入的页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是客户端输入的每页行数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是查询到的实例总数
	TotalCount *int64 `json:"total_count,omitempty"`
}

// InvoiceExpenseRulesGroupEmployeeModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesGroupEmployeeModifyResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}
