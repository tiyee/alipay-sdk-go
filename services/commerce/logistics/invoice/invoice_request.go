package invoice

// LogisticsInvoiceIstdWaybillCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsInvoiceIstdWaybillCreateRequest struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// OutInvoiceRequestNo 是开票请求流水号
	OutInvoiceRequestNo *string `json:"out_invoice_request_no,omitempty"`
	// WaybillInvoices 是即时配送运单列表
	WaybillInvoices *[]WaybillInvoice `json:"waybill_invoices,omitempty"`
}

// LogisticsInvoiceIstdWaybillQueryRequest is the query request model for alipay.commerce.logistics.invoice.istdwaybill.query.
type LogisticsInvoiceIstdWaybillQueryRequest struct {
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty" query:"logistics_code,omitempty"`
	// OutInvoiceRequestNo 是开票请求流水号
	OutInvoiceRequestNo *string `json:"out_invoice_request_no,omitempty" query:"out_invoice_request_no,omitempty"`
}
