// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expense

// InvoiceExpenseProgressSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseProgressSyncRequest struct {
	// ApplyId 是同步申请id，每次发起同步时生成，isv每次请求需要保证唯一
	ApplyId *string `json:"apply_id,omitempty"`
	// BusinessTime 是报销状态变更执行时间
	BusinessTime *string `json:"business_time,omitempty"`
	// ExpenseDetailUrl 是报销详情地址，提供用户通过发票管家查看报销进度的地址 	如果报销企业入驻发票管家时需要isv传入报销详情地址，则必须提供
	ExpenseDetailUrl *string `json:"expense_detail_url,omitempty"`
	// ExpenseOrderNo 是报销单据号
	ExpenseOrderNo *string `json:"expense_order_no,omitempty"`
	// ExpenseTaxNo 是报销企业税号
	ExpenseTaxNo *string `json:"expense_tax_no,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// Memo 是备注
	Memo *string `json:"memo,omitempty"`
	// OpenId 是open_id，支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// Status 是报销操作
	Status *string `json:"status,omitempty"`
	// UserId 是发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitempty"`
}
