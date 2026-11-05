// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package order

// InvoiceOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceOrderQueryResponse struct {
	InvoiceElementModel *InvoiceElement `json:"invoice_element_model,omitempty"`
	// OrderNo 是开票申请时传入订单号（支持主单号、子单号），不限是否为支付宝体内交易单号
	OrderNo *string `json:"order_no,omitempty"`
}
