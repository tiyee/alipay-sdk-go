package servicemarket

// ServiceMarketCommodityShopOfflineResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketCommodityShopOfflineResponse struct {
}

// ServiceMarketPromotionTaskRelationQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketPromotionTaskRelationQueryResponse struct {
	// CurrentPage 是当前页码。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// PageSize 是每页数量。
	PageSize *int64 `json:"page_size,omitempty"`
	// PromotionRelations 是推广服务商信息列表。
	PromotionRelations *[]PromotionTaskRelation `json:"promotion_relations,omitempty"`
	// TotalSize 是总记录数。
	TotalSize *int64 `json:"total_size,omitempty"`
}
