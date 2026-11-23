package knowledge

// SPUDetail 是支付宝官方 V3 OpenAPI 模型。
type SPUDetail struct {
	// Brand 是商品品牌
	Brand *string `json:"brand,omitempty"`
	// Category 是商品类目
	Category *string `json:"category,omitempty"`
	// Dimension 是商品规格
	Dimension *string `json:"dimension,omitempty"`
	// Icon 是商品图片地址
	Icon *string `json:"icon,omitempty"`
	// LinkUrl 是商品链接
	LinkUrl *string `json:"link_url,omitempty"`
	// SpuId 是商品ID
	SpuId *string `json:"spu_id,omitempty"`
	// Title 是商品名称
	Title *string `json:"title,omitempty"`
}
