// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sync

// InvoiceSyncSimpleSendRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceSyncSimpleSendRequest struct {
	InvoiceInfo *InvoicePDFSyn `json:"invoice_info,omitempty"`
	// MShortName 是开票商户品牌简称，与商户入驻时的品牌简称保持一致。详情参见 电子发票
	MShortName *string `json:"m_short_name,omitempty"`
	// SubMShortName 是开票商户门店简称，与商户入驻时的门店简称保持一致。
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}
