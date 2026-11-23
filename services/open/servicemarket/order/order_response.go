package order

// ServiceMarketOrderAcceptResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderAcceptResponse struct {
}

// ServiceMarketOrderCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderCreateResponse struct {
	// CommodityOrderId 是订单号。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// MiniAppId 是支付宝分配给开发者的应用ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// ServiceMarketOrderItemCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemCancelResponse struct {
}

// ServiceMarketOrderItemCompleteResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemCompleteResponse struct {
}

// ServiceMarketOrderItemConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemConfirmResponse struct {
}

// ServiceMarketOrderNotifyResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderNotifyResponse struct {
}

// ServiceMarketOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderQueryResponse struct {
	// AutoUpgrade 是true：开启；false：关闭；需要同步校验commodity_id，如果没有查询到订购信息的话，忽略该字段返回值。
	AutoUpgrade *bool `json:"auto_upgrade,omitempty"`
	// CommodityId 是订购服务商品ID。
	CommodityId *string `json:"commodity_id,omitempty"`
	// CurrentPage 是当前查询页（本接口支持最多查询100条记录）。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// OrderItems 是订单明细列表。
	OrderItems *[]OrderItem `json:"order_items,omitempty"`
	// Specifications 是用于区分同一个服务的不同版本。
	Specifications *string `json:"specifications,omitempty"`
	// Status 是MERCHANT_ORDED（待服务商接单）。
	Status *string `json:"status,omitempty"`
	// TotalSize 是总记录数，例如1000条。
	TotalSize *int64 `json:"total_size,omitempty"`
}

// ServiceMarketOrderRejectResponse 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderRejectResponse struct {
}
