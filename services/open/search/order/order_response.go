package order

// SearchBaseOrderModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBaseOrderModifyResponse struct {
	// ApplyId 是工单id。
	ApplyId *string `json:"apply_id,omitempty"`
}

// SearchOrderDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchOrderDetailQueryResponse struct {
	Data *SearchOrderDetailData `json:"data,omitempty"`
}

// SearchServiceOrderBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchServiceOrderBatchQueryResponse struct {
	Data *OrderPageQueryDTO `json:"data,omitempty"`
}
