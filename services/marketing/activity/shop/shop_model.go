// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package shop

// ActivityShopInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityShopInfo struct {
	// Latitude 是纬度，浮点型,小数点后最多保留6位
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度，浮点型, 小数点后最多保留6位。
	Longitude *string `json:"longitude,omitempty"`
	// ShopId 是支付宝侧蚂蚁店铺 id。
	ShopId *string `json:"shop_id,omitempty"`
	// ShopName 是店铺名称。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopType 是门店类型	PAYMENT_SHOP 支付门店	REAL_SHOP 代运营商业关系门店
	ShopType *string `json:"shop_type,omitempty"`
}
