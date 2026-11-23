// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package title

// InvoiceTitleDynamicGetResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTitleDynamicGetResponse struct {
	Title *InvoiceTitle `json:"title,omitempty"`
}

// InvoiceTitleListGetResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTitleListGetResponse struct {
	// TitleList 是抬头列表
	TitleList *[]InvoiceTitle `json:"title_list,omitempty"`
}
