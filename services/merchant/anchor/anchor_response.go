package anchor

// MerchantAnchorInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantAnchorInfoQueryResponse struct {
	// AnchorId 是主播ID+唯一+天猫好房带货+根据uid获取
	AnchorId *string `json:"anchor_id,omitempty"`
}
