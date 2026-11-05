// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package goods

// ActivityGoodsInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityGoodsInfo struct {
	// GoodsId 是商品编码
	GoodsId *string `json:"goods_id,omitempty"`
	// GoodsUseType 是活动单品类型。
	GoodsUseType *string `json:"goods_use_type,omitempty"`
}

// AppItemInfo 是支付宝官方 V3 OpenAPI 模型。
type AppItemInfo struct {
	// ItemId 是小程序商品-支付宝内部商品id
	ItemId *string `json:"item_id,omitempty"`
	// ItemUseType 是商品可用类型
	ItemUseType *string `json:"item_use_type,omitempty"`
	// MiniAppId 是小程序id
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// OutItemId 是外部商品ID
	OutItemId *string `json:"out_item_id,omitempty"`
}
