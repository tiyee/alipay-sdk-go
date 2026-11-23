package order

// MerchantOrderSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantOrderSyncRequest struct {
	// Amount 是订单总金额：某笔交易订单优惠前的总金额，单位为【元】	实际案例一	实际案例二	实际案例三
	Amount *string `json:"amount,omitempty"`
	// BuyerId 是买家userId
	BuyerId   *string         `json:"buyer_id,omitempty"`
	BuyerInfo *UserInfomation `json:"buyer_info,omitempty"`
	// BuyerOpenId 是用户open_id
	BuyerOpenId *string `json:"buyer_open_id,omitempty"`
	// CategoryId 是标准服务类目
	CategoryId *string `json:"category_id,omitempty"`
	// DiscountAmount 是商户总计优惠金额：代表商户侧给予用户的总计优惠金额 （不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】。	实际案例一	实际案例二	实际案例三
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// DiscountInfoList 是订单优惠信息
	DiscountInfoList *[]DiscountInfoData `json:"discount_info_list,omitempty"`
	// ExtInfo 是扩展信息，请参见 小程序订单中心模板
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// ItemOrderList 是商品信息列表
	ItemOrderList *[]ItemOrderInfo `json:"item_order_list,omitempty"`
	// JourneyOrderList 是行程信息
	JourneyOrderList *[]OrderJourneyInfo `json:"journey_order_list,omitempty"`
	// LogisticsInfoList 是物流信息列表
	LogisticsInfoList *[]OrderLogisticsInformationRequest `json:"logistics_info_list,omitempty"`
	// OrderAuthCode 是本字段已废弃，无需填写！(本字段已废弃 不再使用)
	OrderAuthCode *string `json:"order_auth_code,omitempty"`
	// OrderCreateTime 是订单创建时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// OrderModifiedTime 是订单修改时间
	OrderModifiedTime *string `json:"order_modified_time,omitempty"`
	// OrderPayTime 是订单支付时间
	OrderPayTime *string `json:"order_pay_time,omitempty"`
	// OrderType 是订单类型
	OrderType *string `json:"order_type,omitempty"`
	// OutBizNo 是外部订单号	out_biz_no唯一对应一笔订单，相同的订单需传入相同的out_biz_no
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PartnerId 是交易对应的签约商户userId(注意：该字段自2020-02-16日起，可以不传入)
	PartnerId *string `json:"partner_id,omitempty"`
	// PayAmount 是用户应付金额 ：用户最终结算时需要支付金额（不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】	实际案例一	实际案例二	实际案例三
	PayAmount *string `json:"pay_amount,omitempty"`
	// PayTimeoutExpress 是支付超时时间，超过时间支付宝自行关闭订单
	PayTimeoutExpress *string `json:"pay_timeout_express,omitempty"`
	// RecordId 是商户订单同步记录id(仅部分存量接入和行业模板需要外，其他情况可以不传入)
	RecordId *string `json:"record_id,omitempty"`
	// SellerId 是卖家userId(注意：该字段自2020-02-16日起，可以不传入)
	SellerId *string `json:"seller_id,omitempty"`
	// SendMsg 是是否需要小程序订单代理发送模版消息。不传默认不发送
	SendMsg *string `json:"send_msg,omitempty"`
	// ServiceCode 是服务code：传入小程序后台提报的服务id，将订单与服务关联，有利于提高服务曝光机会；入参服务id的类目须与订单类型相符，若不相符将会报错；如订单类型为“外卖”，则入参的服务ID所对应的服务类目也必须得是”外卖“；service_code 通过 alipay.open.app.service.apply ，(服务提报申请)接口创建服务后获取。
	ServiceCode *string        `json:"service_code,omitempty"`
	ShopInfo    *OrderShopInfo `json:"shop_info,omitempty"`
	// SourceApp 是用于区分用户下单的订单来源
	SourceApp *string `json:"source_app,omitempty"`
	// SyncContent 是同步内容
	SyncContent *string     `json:"sync_content,omitempty"`
	TicketInfo  *TicketInfo `json:"ticket_info,omitempty"`
	// TicketOrderList 是凭证信息
	TicketOrderList *[]TicketOrderInfo `json:"ticket_order_list,omitempty"`
	// TradeNo 是订单所对应的支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
	// TradeType 是交易号类型
	TradeType *string `json:"trade_type,omitempty"`
}

// OrderLogisticsInformationRequest 是支付宝官方 V3 OpenAPI 模型。
type OrderLogisticsInformationRequest struct {
	// LogisticsCode 是物流公司编号。物流公司编号值请查看产品文档
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// PhoneNo 是手机号后四位
	PhoneNo *string `json:"phone_no,omitempty"`
	// TrackingNo 是物流单号
	TrackingNo *string `json:"tracking_no,omitempty"`
}
