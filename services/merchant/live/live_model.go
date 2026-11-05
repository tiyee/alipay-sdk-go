package live

// ItemPlanContentDO 是支付宝官方 V3 OpenAPI 模型。
type ItemPlanContentDO struct {
	// Codes 是对应服务库编码
	Codes *[]string `json:"codes,omitempty"`
	// Type 是条目类型，必须， 商品：ITEM
	Type *string `json:"type,omitempty"`
}
