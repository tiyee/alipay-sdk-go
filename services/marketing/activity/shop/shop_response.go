// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package shop

// ActivityShopBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityShopBatchQueryResponse struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// ShopInfos 是可用门店列表
	ShopInfos *[]ActivityShopInfo `json:"shop_infos,omitempty"`
	// TotalSize 是可用门店总数量
	TotalSize *int64 `json:"total_size,omitempty"`
}
