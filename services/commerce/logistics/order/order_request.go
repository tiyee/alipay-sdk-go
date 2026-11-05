package order

// LogisticsOrderInstantDeliveryCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryCancelRequest struct {
	// CancelReason 是取消原因，取消原因id为i_6时必填
	CancelReason *string `json:"cancel_reason,omitempty"`
	// CancelReasonId 是取消原因id
	CancelReasonId *string `json:"cancel_reason_id,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// OrderNo 是支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty"`
	// OutOrderNo 是商家订单号，与order_no不能同时为空
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// LogisticsOrderInstantDeliveryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryCreateRequest struct {
	// ConsumerId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空
	ConsumerId     *string             `json:"consumer_id,omitempty"`
	ConsumerNotify *ConsumerNotifyIstd `json:"consumer_notify,omitempty"`
	// ConsumerSource 是消费者来源,支付宝：alipay; 微信：wx
	ConsumerSource *string `json:"consumer_source,omitempty"`
	// GoodsDetails 是商品明细
	GoodsDetails *[]GoodsDetailIstd `json:"goods_details,omitempty"`
	GoodsInfo    *GoodsInfoIstd     `json:"goods_info,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsToken 是下即时配送订单token, 配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变
	LogisticsToken *string `json:"logistics_token,omitempty"`
	// OpenId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户openId；如果consumer_source是wx，则consumer_id可以为空
	OpenId       *string       `json:"open_id,omitempty"`
	OrderExtIstd *OrderExtIstd `json:"order_ext_istd,omitempty"`
	// OutOrderNo 是商家订单号
	OutOrderNo *string       `json:"out_order_no,omitempty"`
	Receiver   *ReceiverIstd `json:"receiver,omitempty"`
	Sender     *SenderIstd   `json:"sender,omitempty"`
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
}

// LogisticsOrderInstantDeliveryPreCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryPreCreateRequest struct {
	// ConsumerId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空
	ConsumerId     *string             `json:"consumer_id,omitempty"`
	ConsumerNotify *ConsumerNotifyIstd `json:"consumer_notify,omitempty"`
	// ConsumerSource 是消费者来源, 支付宝：alipay, 微信：wx
	ConsumerSource *string `json:"consumer_source,omitempty"`
	// GoodsDetails 是商品明细
	GoodsDetails *[]GoodsDetailIstd `json:"goods_details,omitempty"`
	GoodsInfo    *GoodsInfoIstd     `json:"goods_info,omitempty"`
	// LogisticsCompanies 是即时配送公司配置列表
	LogisticsCompanies *[]LogisticsCompanyIstd `json:"logistics_companies,omitempty"`
	// OpenId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝的openId；如果consumer_source是wx，则consumer_id可以为空
	OpenId       *string                  `json:"open_id,omitempty"`
	OrderExtIstd *OrderExtIstdForPreOrder `json:"order_ext_istd,omitempty"`
	// OutOrderNo 是商家订单号
	OutOrderNo *string       `json:"out_order_no,omitempty"`
	Receiver   *ReceiverIstd `json:"receiver,omitempty"`
	Sender     *SenderIstd   `json:"sender,omitempty"`
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
}

// LogisticsOrderIstdCancelPreConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderIstdCancelPreConsultRequest struct {
	// CancelReason 是取消原因，取消原因id为i_6时必填
	CancelReason *string `json:"cancel_reason,omitempty"`
	// CancelReasonId 是取消原因id
	CancelReasonId *string `json:"cancel_reason_id,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// OrderNo 是支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty"`
	// OutOrderNo 是商家订单号，与order_no不能同时为空
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// LogisticsOrderIstdRetryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderIstdRetryCreateRequest struct {
	// ConsumerId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空
	ConsumerId     *string             `json:"consumer_id,omitempty"`
	ConsumerNotify *ConsumerNotifyIstd `json:"consumer_notify,omitempty"`
	// ConsumerSource 是消费者来源,  支付宝：alipay, 微信：wx
	ConsumerSource *string `json:"consumer_source,omitempty"`
	// GoodsDetails 是商品明细
	GoodsDetails *[]GoodsDetailIstd `json:"goods_details,omitempty"`
	GoodsInfo    *GoodsInfoIstd     `json:"goods_info,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsToken 是下即时配送订单token, 配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变
	LogisticsToken *string `json:"logistics_token,omitempty"`
	// OpenId 是消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户openId；如果consumer_source是wx，则consumer_id可以为空
	OpenId       *string       `json:"open_id,omitempty"`
	OrderExtIstd *OrderExtIstd `json:"order_ext_istd,omitempty"`
	// OutOrderNo 是商家订单号
	OutOrderNo *string       `json:"out_order_no,omitempty"`
	Receiver   *ReceiverIstd `json:"receiver,omitempty"`
	Sender     *SenderIstd   `json:"sender,omitempty"`
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
}
