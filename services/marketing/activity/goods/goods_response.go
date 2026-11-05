// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package goods

// ActivityGoodsBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityGoodsBatchQueryResponse struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// AppItemInfos 是小程序商品信息
	AppItemInfos *[]AppItemInfo `json:"app_item_infos,omitempty"`
	// GoodsInfos 是商品编码列表
	GoodsInfos *[]ActivityGoodsInfo `json:"goods_infos,omitempty"`
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalSize 是商品编码总数量
	TotalSize *int64 `json:"total_size,omitempty"`
}
