// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecontrol

// InvoiceExpenseControlAggregationCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlAggregationCreateResponse struct {
	// AggregationId 是费控聚合关系ID
	AggregationId *string `json:"aggregation_id,omitempty"`
}

// InvoiceExpenseControlIssueBatchCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlIssueBatchCancelResponse struct {
	// Result 是是否成功
	Result *bool `json:"result,omitempty"`
}

// InvoiceExpenseControlIssueBatchCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlIssueBatchCreateResponse struct {
	// IssueBatchId 是发放批次id
	IssueBatchId *string `json:"issue_batch_id,omitempty"`
	// IssueQuotaCheckFailedList 是校验失败的数据
	IssueQuotaCheckFailedList *[]IssueQuotaCheckInfo `json:"issue_quota_check_failed_list,omitempty"`
}

// InvoiceExpenseControlQuotaCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaCreateResponse struct {
	// IssueBatchId 是发放批次id
	IssueBatchId *string `json:"issue_batch_id,omitempty"`
	// IssueQuotaCheckFailedList 是额度发放明细检查错误列表，可根据报错信息修改发放明细列表
	IssueQuotaCheckFailedList *[]IssueQuotaCheckInfo `json:"issue_quota_check_failed_list,omitempty"`
	// QuotaId 是创建的额度ID
	QuotaId *string `json:"quota_id,omitempty"`
}

// InvoiceExpenseControlQuotaModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaModifyResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}

// InvoiceExpenseControlQuotaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaQueryResponse struct {
	// ExpenseQuotaInfoList 是额度列表
	ExpenseQuotaInfoList *[]ExpenseQuotaInfo `json:"expense_quota_info_list,omitempty"`
	// PageNum 是当前页数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是当前记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}
