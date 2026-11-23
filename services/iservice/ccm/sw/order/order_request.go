package order

// CcmSwOrderSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwOrderSyncRequest struct {
	// Amount 是订单金额
	Amount *string `json:"amount,omitempty"`
	// LinkUrl 是订单链接
	LinkUrl *string `json:"link_url,omitempty"`
	// LogisticCount 是订单物流数量
	LogisticCount *int64 `json:"logistic_count,omitempty"`
	// Logistics 是物流信息列表
	Logistics *[]LogisticInfo `json:"logistics,omitempty"`
	// OrderCreateTime 是订单创建时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// OrderId 是订单id
	OrderId *string `json:"order_id,omitempty"`
	// OrderType 是普通订单：NORMAL	预售订单：PRE_SALE
	OrderType *string `json:"order_type,omitempty"`
	// SpuCount 是订单商品种类
	SpuCount *int64 `json:"spu_count,omitempty"`
	// Spus 是订单商品信息
	Spus *[]SPUInfo `json:"spus,omitempty"`
	// Status 是订单状态，目前支持以下几种状态	下单未支付	WAIT_PAY	支付未发货	PAIED	已发货	IN_DELIVERY	售后中	IN_AFTER_SALE	订单完成	FINISHED
	Status *string `json:"status,omitempty"`
	// SubStatus 是订单子状态
	SubStatus *string `json:"sub_status,omitempty"`
	// UserId 是下单用户id(外部系统ID)
	UserId *string `json:"user_id,omitempty"`
}
