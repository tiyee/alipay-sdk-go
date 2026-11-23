// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// InvoiceFileOutputQueryRequest is the query request model for alipay.ebpp.invoice.file.output.query.
type InvoiceFileOutputQueryRequest struct {
	// UserId 是发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// InvoiceCode 是发票代码	字段长度（10-12位），全电票时为空
	InvoiceCode *string `json:"invoice_code,omitempty" query:"invoice_code,omitempty"`
	// InvoiceNo 是发票号码	字段长度（8-10位），全电票时为20位
	InvoiceNo *string `json:"invoice_no,omitempty" query:"invoice_no,omitempty"`
	// Scene 是发票pdf文件下载应用场景。固定为 INVOICE_EXPENSE  表示发票报销。
	Scene *string `json:"scene,omitempty" query:"scene,omitempty"`
	// SkipExpenseProgressSync 是是否跳过发票报销状态同步；当为true时，跳过报销状态同步校验。默认为false，需要先做报销状态同步
	SkipExpenseProgressSync *bool `json:"skip_expense_progress_sync,omitempty" query:"skip_expense_progress_sync,omitempty"`
}
