// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package token

// InvoiceTokenBatchQueryRequest is the query request model for alipay.ebpp.invoice.token.batchquery.
type InvoiceTokenBatchQueryRequest struct {
	// InvoiceToken 是查询请求令牌。通过唤起支付宝钱包后用户进入发票管家选择发票列表后创建，并通过isv接收url回传给isv。详见 "选"模式发票报销。
	InvoiceToken *string `json:"invoice_token,omitempty" query:"invoice_token,omitempty"`
	// Scene 是发票要素获取应用场景。固定为 INVOICE_EXPENSE 表示发票报销。
	Scene *string `json:"scene,omitempty" query:"scene,omitempty"`
}
