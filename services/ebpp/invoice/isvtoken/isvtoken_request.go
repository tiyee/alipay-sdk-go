// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isvtoken

// InvoiceISVTokenReimApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceISVTokenReimApplyRequest struct {
	// IsvAppCode 是报销软件app对应的isv app code，由支付宝进行分配。详情参见 "推"模式发票报销
	IsvAppCode *string `json:"isv_app_code,omitempty"`
}
