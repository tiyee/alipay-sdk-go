package invoicereceipt

// BossFncInvoiceReceiptBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceReceiptBatchQueryRequest struct {
	// BillNos 是月账单号集合，唯一性ID集合，月账单表唯一主键ID集合
	BillNos *[]string `json:"bill_nos,omitempty"`
	// OutBizType 是来源类型，01：主站；02：芝麻；03：金融云；04：微贷
	OutBizType *string `json:"out_biz_type,omitempty"`
}

// BossFncInvoiceReceiptCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceReceiptCreateRequest struct {
	// EventCode 是账单金额变更消息事件码 	AR_MTHLY_BILL_CREATE:应收账单创建, 	AR_MTHLY_BILL_CHANGE:应收金额更新
	EventCode *string `json:"event_code,omitempty"`
	// EventType 是账单金额变更消息事件类型 	CREATE:创建事件,UPDATE:更新事件
	EventType   *string           `json:"event_type,omitempty"`
	MonthlyBill *ArMonthlyBillDTO `json:"monthly_bill,omitempty"`
	// MsgId 是事件全局id，可以使用uuid
	MsgId *string `json:"msg_id,omitempty"`
	// OutBizType 是消息来源 01:主站，02：芝麻，03:金融云,04:微贷
	OutBizType *string `json:"out_biz_type,omitempty"`
}
