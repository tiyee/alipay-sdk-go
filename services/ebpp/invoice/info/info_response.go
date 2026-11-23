// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// InvoiceInfoSendResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInfoSendResponse struct {
	// Url 是发票访问地址，同步红票的情况下不返回该字段，同步蓝票的情况下如果同步单张发票，返回发票详情链接地址，如果同步的是多张发票，返回发票列表地址
	Url *string `json:"url,omitempty"`
}
