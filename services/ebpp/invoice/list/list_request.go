// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package list

// InvoiceListExpenseSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceListExpenseSyncRequest struct {
	// BusinessTime 是报销进度变更时间
	BusinessTime *string `json:"business_time,omitempty"`
	// ExpenseOrderNo 是商户报销单据号
	ExpenseOrderNo *string `json:"expense_order_no,omitempty"`
	// ExpenseTaxNo 是报销企业税号
	ExpenseTaxNo *string `json:"expense_tax_no,omitempty"`
	// InvoiceElementList 是需要同步报销状态的发票列表
	InvoiceElementList *[]InvoiceElementStatusSyncOpen `json:"invoice_element_list,omitempty"`
	// Memo 是报销审批备注
	Memo *string `json:"memo,omitempty"`
	// OpenId 是用户uid
	OpenId *string `json:"open_id,omitempty"`
	// Status 是EXPENSE_APPLY:报销申请提交;EXPENSE_APPROVAL_PASS:报销申请审核通过;EXPENSE_FINISHED:报销申请完成;EXPENSE_CANCEL:报销申请撤回
	Status *string `json:"status,omitempty"`
	// UserId 是用户uid
	UserId *string `json:"user_id,omitempty"`
}
