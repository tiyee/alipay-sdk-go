// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package token

// InvoiceTokenBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTokenBatchQueryResponse struct {
	// InvoiceElementList 是发票要素列表
	InvoiceElementList *[]InvoiceElement `json:"invoice_element_list,omitempty"`
	// OpenId 是支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户id
	UserId *string `json:"user_id,omitempty"`
}
