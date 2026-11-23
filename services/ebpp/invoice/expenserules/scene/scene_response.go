package scene

// InvoiceExpenseRulesSceneRuleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesSceneRuleCreateResponse struct {
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
}

// InvoiceExpenseRulesSceneRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesSceneRuleModifyResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}

// InvoiceExpenseRulesSceneRuleQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesSceneRuleQueryResponse struct {
	// PageNum 是当前页数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是当前记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// StandardRuleInfoList 是费控规则列表
	StandardRuleInfoList *[]StandardRuleInfo `json:"standard_rule_info_list,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}
