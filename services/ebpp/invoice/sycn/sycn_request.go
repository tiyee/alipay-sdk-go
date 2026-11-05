// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sycn

// InvoiceSycnRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceSycnRequest struct {
	// InvoiceInfo 是同步发票信息模型
	InvoiceInfo *[]InvoiceModelContent `json:"invoice_info,omitempty"`
	// MShortName 是商户的品牌名称简称,该字段需要接入前向发票管家申请， 	m_short_name+sub_m_short_name具有唯一约束 	如：肯德基：KFC
	MShortName *string `json:"m_short_name,omitempty"`
	// SubMShortName 是支付宝为商户分配的商户门店简称，该字段需要接入前在发票管家申请 	如：肯德基-杭州西湖区文一西路店：KFC-HZ-XH001
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}
