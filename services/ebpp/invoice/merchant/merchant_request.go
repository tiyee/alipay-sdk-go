// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// InvoiceMerchantEnterStatusQueryRequest is the query request model for alipay.ebpp.invoice.merchant.enterstatus.query.
type InvoiceMerchantEnterStatusQueryRequest struct {
	// MShortName 是商户品牌简称，只能由大写字母、下划线、数字组成，且必须以大写字母开头。
	MShortName *string `json:"m_short_name,omitempty" query:"m_short_name,omitempty"`
	// ProductCode 是商户门店入驻的产品码,STANDARD_INVOICE:扫码开票，INVOICE_RETURN:企业自建-发票回传，INVOICE_EXPENSE:发票报销，PAYMENT_OPEN:支付即开票，PAY_FEE_OPEN:缴费后开票。
	ProductCode *string `json:"product_code,omitempty" query:"product_code,omitempty"`
	// ProcessId 是工单流水号（如果不填则默认查询最近一条工单）
	ProcessId *string `json:"process_id,omitempty" query:"process_id,omitempty"`
}
