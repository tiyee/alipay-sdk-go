package consume

// EcConsumeDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcConsumeDetailQueryResponse struct {
	ConsumeInfo      *EcConsumeInfo `json:"consume_info,omitempty"`
	RelatedOrderInfo *EcOrderInfo   `json:"related_order_info,omitempty"`
	// RelatedRefundList 是关联账单列表
	RelatedRefundList *[]EcConsumeInfo `json:"related_refund_list,omitempty"`
	// RelatedVoucherList 是关联凭证详情列表
	RelatedVoucherList *[]EcVoucherInfo `json:"related_voucher_list,omitempty"`
}
