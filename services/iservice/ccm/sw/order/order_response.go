package order

// CcmSwOrderSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwOrderSyncResponse struct {
	// OrderId 是订单id
	OrderId *string `json:"order_id,omitempty"`
}
