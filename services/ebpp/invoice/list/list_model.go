// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package list

// InvoiceElementStatusSyncOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceElementStatusSyncOpen struct {
	// ApplyId 是同步申请id，每次发起同步时生成，isv每次请求需要保证唯一
	ApplyId *string `json:"apply_id,omitempty"`
	// ExpenseDetailUrl 是报销详情地址，提供用户通过发票管家查看报销进度的地址 	如果报销企业入驻发票管家时需要isv传入报销详情地址，则必须提供
	ExpenseDetailUrl *string `json:"expense_detail_url,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
}
