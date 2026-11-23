// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ecorder

// InvoiceEcOrderOrderQueryRequest is the query request model for alipay.ebpp.invoice.ecorder.order.query.
type InvoiceEcOrderOrderQueryRequest struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// OrderId 是订单ID
	OrderId *string `json:"order_id,omitempty" query:"order_id,omitempty"`
	// OrderType 是订单类型
	OrderType *string `json:"order_type,omitempty" query:"order_type,omitempty"`
}
