// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package apply

// InvoiceApplyInvUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyInvUploadResponse struct {
	// IsSuccess 是处理是否成功
	IsSuccess *bool `json:"is_success,omitempty"`
}

// InvoiceApplyResultSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyResultSyncResponse struct {
	// RetryFlag 是标注是否需要调用方重试。 	当结果返回失败时（code 不等于 10000 且 msg 不等于 SUCCESS 	），如果该字段返回true表明该失败的情况通过重试补偿可解决，为false表明失败情况通过重试无法解决，可以停止重试，根据返回的错误码尝试解决。
	RetryFlag *bool `json:"retry_flag,omitempty"`
}

// InvoiceApplyStatusNotifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceApplyStatusNotifyResponse struct {
	// MissingInvoices 是update_success=false时返回。	该申请下存在未归集成功的发票。即：attach_invoices 中携带的发票存在未回传成功的发票。	此时调用方，需先调用发票回传接口，完成发票回传。然后再重试此接口。
	MissingInvoices *[]InvoiceUkDTO `json:"missing_invoices,omitempty"`
	// UpdateSuccess 是事务执行结果，true成功，false失败
	UpdateSuccess *bool `json:"update_success,omitempty"`
}
