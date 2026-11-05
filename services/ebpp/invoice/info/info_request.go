// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// InvoiceInfoSendRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInfoSendRequest struct {
	// InvoiceInfoList 是发票信息列表
	InvoiceInfoList *[]InvoiceSendOpen `json:"invoice_info_list,omitempty"`
	// MShortName 是开票商户品牌简称，与商户入驻时的品牌简称保持一致。
	MShortName *string `json:"m_short_name,omitempty"`
	// SubMShortName 是开票商户门店简称，与商户入驻时的门店简称保持一致。
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}
