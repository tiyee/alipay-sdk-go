// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package issuerule

// InvoiceIssueRuleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleCreateResponse struct {
	// IssueRuleId 是创建的发放规则ID
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
}

// InvoiceIssueRuleDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleDeleteResponse struct {
	// Result 是删除结果
	Result *bool `json:"result,omitempty"`
}

// InvoiceIssueRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleModifyResponse struct {
	// Result 是修改结果
	Result *bool `json:"result,omitempty"`
}
