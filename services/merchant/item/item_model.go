package item

// MerchantItemFileUpload 是支付宝官方 V3 OpenAPI 模型。
type MerchantItemFileUpload struct {
	// Scene 是业务场景描述。	小程序订单中心场景固定为 SYNC_ORDER。
	Scene *string `json:"scene,omitempty"`
}
