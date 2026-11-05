package spu

// SPUAttribute 是支付宝官方 V3 OpenAPI 模型。
type SPUAttribute struct {
	// Name 是商品属性名
	Name *string `json:"name,omitempty"`
	// Value 是商品属性值
	Value *[]string `json:"value,omitempty"`
}
