package invoice

// Invoice 是支付宝官方 V3 OpenAPI 模型。
type Invoice struct {
	// Email 是邮箱地址
	Email *string `json:"email,omitempty"`
	// InvoiceContent 是发票内容，配送服务费
	InvoiceContent *string `json:"invoice_content,omitempty"`
	// InvoiceFee 是发票金额
	InvoiceFee *string `json:"invoice_fee,omitempty"`
	// InvoiceTitle 是发票抬头，如支付宝（杭州）信息技术有限公司
	InvoiceTitle *string `json:"invoice_title,omitempty"`
	// InvoiceType 是发票类型，0：默认电子发票
	InvoiceType *int64 `json:"invoice_type,omitempty"`
	// Phone 是手机号码
	Phone *string `json:"phone,omitempty"`
	// Receiver 是收件人姓名
	Receiver *string `json:"receiver,omitempty"`
	// TaxNumber 是税号
	TaxNumber *string `json:"tax_number,omitempty"`
	// TitleType 是抬头类型，0：企业单位，1：个人/非企业
	TitleType *int64 `json:"title_type,omitempty"`
}

// WaybillInvoice 是支付宝官方 V3 OpenAPI 模型。
type WaybillInvoice struct {
	// WaybillAmount 是即时配送运单金额
	WaybillAmount *string `json:"waybill_amount,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// WaybillInvoiceIstd 是支付宝官方 V3 OpenAPI 模型。
type WaybillInvoiceIstd struct {
	// Reason 是失败原因，需要传单个配送单的驳回原因
	Reason *string `json:"reason,omitempty"`
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
	// WaybillAmount 是即时配送运单金额，waybill_invoce_status为1的情况下不能为空
	WaybillAmount *string `json:"waybill_amount,omitempty"`
	// WaybillInvoiceStatus 是明细的运单开票状态，1：开票成功 2：不可开票 3：可开票；整体开票状态为0的情况下，无开票明细；整体开票状态为1，明细开票状态全部是1；整体开票状态为2，明细开票状态为2或者3
	WaybillInvoiceStatus *int64 `json:"waybill_invoice_status,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// WaybillInvoiceQueryIstd 是支付宝官方 V3 OpenAPI 模型。
type WaybillInvoiceQueryIstd struct {
	// Reason 是失败原因，需要传单个即时配送运单的驳回原因
	Reason *string `json:"reason,omitempty"`
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
	// WaybillAmount 是即时配送运单金额，waybill_invoce_status为1的情况下不能为空
	WaybillAmount *string `json:"waybill_amount,omitempty"`
	// WaybillInvoiceStatus 是明细的运单开票状态，1：开票成功 2：不可开票 3：可开票；整体开票状态为0的情况下，无开票明细；整体开票状态为1，明细开票状态全部是1；整体开票状态为2，明细开票状态全部为2
	WaybillInvoiceStatus *int64 `json:"waybill_invoice_status,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}
