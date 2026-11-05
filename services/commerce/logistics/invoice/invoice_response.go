package invoice

// LogisticsInvoiceIstdWaybillCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsInvoiceIstdWaybillCreateResponse struct {
	// InvoiceFee 是开票金额
	InvoiceFee *string `json:"invoice_fee,omitempty"`
	// Status 是整体的运单开票状态，0：处理中 1：开票成功 2：开票失败
	Status *int64 `json:"status,omitempty"`
	// WaybillInvoices 是即时配送运单列表
	WaybillInvoices *[]WaybillInvoiceIstd `json:"waybill_invoices,omitempty"`
}

// LogisticsInvoiceIstdWaybillQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsInvoiceIstdWaybillQueryResponse struct {
	// Status 是整体的运单开票状态，0：处理中 1：开票成功 2：开票失败
	Status *int64 `json:"status,omitempty"`
	// WaybillInvoices 是即时配送运单列表
	WaybillInvoices *[]WaybillInvoiceQueryIstd `json:"waybill_invoices,omitempty"`
}
