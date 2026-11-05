// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isvtoken

// InvoiceISVTokenReimApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceISVTokenReimApplyResponse struct {
	// IsvToken 是支付宝分配给报销软件的授权token，当前设置有效时间为20分钟
	IsvToken *string `json:"isv_token,omitempty"`
	// SerialNo 是该次请求的唯一标志，当前设置有效时间为20分钟
	SerialNo *string `json:"serial_no,omitempty"`
}
