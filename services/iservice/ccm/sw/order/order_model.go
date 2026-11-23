package order

// LogisticInfo 是支付宝官方 V3 OpenAPI 模型。
type LogisticInfo struct {
	// Channel 是物流公司名称
	Channel *string `json:"channel,omitempty"`
	// Detail 是物流详情
	Detail *string `json:"detail,omitempty"`
	// LogisticId 是物流id
	LogisticId *string `json:"logistic_id,omitempty"`
	// ShipArea 是发货地
	ShipArea *string `json:"ship_area,omitempty"`
	// ShipPeriod 是发货时效
	ShipPeriod *string `json:"ship_period,omitempty"`
	// Status 是物流状态
	Status *string `json:"status,omitempty"`
	// StopUpdateTime 是物流停更时间
	StopUpdateTime *string `json:"stop_update_time,omitempty"`
}

// SPUInfo 是支付宝官方 V3 OpenAPI 模型。
type SPUInfo struct {
	// Brand 是品牌名称
	Brand *string `json:"brand,omitempty"`
	// Category 是类目
	Category *string `json:"category,omitempty"`
	// Count 是商品数量
	Count *int64 `json:"count,omitempty"`
	// Icon 是图片链接
	Icon *string `json:"icon,omitempty"`
	// Price 是商品单价(单位:元)
	Price *string `json:"price,omitempty"`
	// Provider 是商品提供方，店铺或品牌方
	Provider *string `json:"provider,omitempty"`
	// SpuId 是商品ID
	SpuId *string `json:"spu_id,omitempty"`
}
