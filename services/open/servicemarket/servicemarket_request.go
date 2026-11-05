package servicemarket

// ServiceMarketCommodityShopOfflineRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketCommodityShopOfflineRequest struct {
	// CommodityId 是服务商户ID。
	CommodityId *string `json:"commodity_id,omitempty"`
	// ShopId 是门店ID。
	ShopId *string `json:"shop_id,omitempty"`
}

// ServiceMarketPromotionTaskRelationQueryRequest is the request type for alipay.open.servicemarket.promotiontask.relation.query.
type ServiceMarketPromotionTaskRelationQueryRequest struct {
	// CommodityId 是服务商品ID，并不是uid，取值来源于前端的获取
	CommodityId *string `json:"commodity_id,omitempty" query:"commodity_id,omitempty"`
	// PageSize 是每页数量，范围(1-100)。默认：10
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,omitempty"`
	// PageNum 是当前页码，从1开始。默认：1
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,omitempty"`
}
