// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensescene

// InvoiceExpenseSceneMerchantQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseSceneMerchantQueryResponse struct {
	// ExpenseMerchantInfoList 是商户信息
	ExpenseMerchantInfoList *[]ExpenseMerchantInfo `json:"expense_merchant_info_list,omitempty"`
	// PageNum 是当前页数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是当前记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}
