package order

// MerchantOrderSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantOrderSyncResponse struct {
	// DistributeResult 是"分发结果 若未分发到场景侧，则会返回具体的未分发原因"
	DistributeResult *[]OrderDataDistributeInfo `json:"distribute_result,omitempty"`
	// OrderId 是支付宝订单号
	OrderId *string `json:"order_id,omitempty"`
	// OrderStatus 是订单状态
	OrderStatus *string `json:"order_status,omitempty"`
	// RecordId 是同步订单记录id
	RecordId *string `json:"record_id,omitempty"`
	// SyncSuggestions 是订单同步优化建议，供开发者参考
	SyncSuggestions *[]OrderDataSyncSuggestion `json:"sync_suggestions,omitempty"`
}
