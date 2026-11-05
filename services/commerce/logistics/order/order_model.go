package order

// ConsumerNotifyIstd 是支付宝官方 V3 OpenAPI 模型。
type ConsumerNotifyIstd struct {
	// GoodsCount 是商品数量
	GoodsCount *int64 `json:"goods_count,omitempty"`
	// GoodsImg 是商品缩略图url，支持格式：bmp、jpg、jpeg、png、gif
	GoodsImg *string `json:"goods_img,omitempty"`
	// GoodsName 是商品名称
	GoodsName *string `json:"goods_name,omitempty"`
	// MerchantMobile 是商家电话
	MerchantMobile *string `json:"merchant_mobile,omitempty"`
	// MerchantName 是商家名称，tiny_app_id和merchant_name不能同时为空
	MerchantName *string `json:"merchant_name,omitempty"`
	// TinyAppId 是商家小程序appid
	TinyAppId *string `json:"tiny_app_id,omitempty"`
	// TinyAppUrl 是商家小程序的路径，建议为订单页面
	TinyAppUrl *string `json:"tiny_app_url,omitempty"`
}

// GoodsDetailIstd 是支付宝官方 V3 OpenAPI 模型。
type GoodsDetailIstd struct {
	// Count 是商品数量
	Count *int64 `json:"count,omitempty"`
	// Name 是商品名称
	Name *string `json:"name,omitempty"`
	// Price 是商品单价
	Price *string `json:"price,omitempty"`
	// Unit 是商品单位
	Unit *string `json:"unit,omitempty"`
}

// GoodsInfoIstd 是支付宝官方 V3 OpenAPI 模型。
type GoodsInfoIstd struct {
	// DeliveryInfo 是交货信息，用于骑手交付商品
	DeliveryInfo *string `json:"delivery_info,omitempty"`
	// FirstClass 是一级品类类目
	FirstClass *string `json:"first_class,omitempty"`
	// Height 是商品高度，单位为cm
	Height *int64 `json:"height,omitempty"`
	// Length 是商品长度，单位为cm
	Length *int64 `json:"length,omitempty"`
	// PickupInfo 是取货信息，用于骑手到店取货
	PickupInfo *string `json:"pickup_info,omitempty"`
	// Price 是总价
	Price *string `json:"price,omitempty"`
	// SecondClass 是二级品类类目
	SecondClass *string `json:"second_class,omitempty"`
	// Weight 是货物重量，单位为kg，精确到小数点后两位
	Weight *string `json:"weight,omitempty"`
	// Width 是商品宽度，单位为cm
	Width *int64 `json:"width,omitempty"`
}

// LogisticsCompanyIstd 是支付宝官方 V3 OpenAPI 模型。
type LogisticsCompanyIstd struct {
	// LogisticsCode 是即时配送公司编码，由支付宝分配
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// ServiceCodes 是即时配送公司服务代码列表
	ServiceCodes *[]ServiceCodeIstd `json:"service_codes,omitempty"`
}

// OrderExtIstd 是支付宝官方 V3 OpenAPI 模型。
type OrderExtIstd struct {
	// CashOnDelivery 是骑手应付金额
	CashOnDelivery *string `json:"cash_on_delivery,omitempty"`
	// CashOnPickup 是骑手应收金额
	CashOnPickup *string `json:"cash_on_pickup,omitempty"`
	// ConsumerOrderTime 是消费者下单时间
	ConsumerOrderTime *string `json:"consumer_order_time,omitempty"`
	// DeliveryDirection 是物流流向，1：从门店取件送至用户；2：从用户取件送至门店
	DeliveryDirection *int64 `json:"delivery_direction,omitempty"`
	// DeliveryType 是配送类型, 0: 即时单 1 预约单
	DeliveryType *int64 `json:"delivery_type,omitempty"`
	// Desc 是备注
	Desc *string `json:"desc,omitempty"`
	// ExpectedDeliveryTime 是期望派单时间
	ExpectedDeliveryTime *string `json:"expected_delivery_time,omitempty"`
	// ExpectedFinishTime 是期望送达时间
	ExpectedFinishTime *string `json:"expected_finish_time,omitempty"`
	// ExpectedPickTime 是期望取件时间
	ExpectedPickTime *string `json:"expected_pick_time,omitempty"`
	// InsurePrice 是保价金额
	InsurePrice *string `json:"insure_price,omitempty"`
	// IsDirectDelivery 是是否直拿直送，1:直拿直送 0:非直拿直送
	IsDirectDelivery *int64 `json:"is_direct_delivery,omitempty"`
	// IsFinishCodeNeeded 是是否需要收货码，1:需要 0:不需要
	IsFinishCodeNeeded *int64 `json:"is_finish_code_needed,omitempty"`
	// IsInsured 是是否保价，1:保价 0:不保价
	IsInsured *int64 `json:"is_insured,omitempty"`
	// IsPickupCodeNeeded 是是否需要取货码，1:需要 0:不需要
	IsPickupCodeNeeded *int64 `json:"is_pickup_code_needed,omitempty"`
	// PoiSeq 是门店订单流水号
	PoiSeq *string `json:"poi_seq,omitempty"`
	// ServiceCode 是即时配送公司的服务代码
	ServiceCode *string `json:"service_code,omitempty"`
}

// OrderExtIstdForPreOrder 是支付宝官方 V3 OpenAPI 模型。
type OrderExtIstdForPreOrder struct {
	// CashOnDelivery 是骑手应付金额，单位元
	CashOnDelivery *string `json:"cash_on_delivery,omitempty"`
	// CashOnPickup 是骑手应收金额，单位元
	CashOnPickup *string `json:"cash_on_pickup,omitempty"`
	// ConsumerOrderTime 是消费者下单时间
	ConsumerOrderTime *string `json:"consumer_order_time,omitempty"`
	// DeliveryDirection 是物流流向，1：从门店取件送至用户；2：从用户取件送至门店
	DeliveryDirection *int64 `json:"delivery_direction,omitempty"`
	// DeliveryType 是配送类型, 0: 即时单 1 预约单
	DeliveryType *int64 `json:"delivery_type,omitempty"`
	// Desc 是备注
	Desc *string `json:"desc,omitempty"`
	// ExpectedDeliveryTime 是期望派单时间
	ExpectedDeliveryTime *string `json:"expected_delivery_time,omitempty"`
	// ExpectedFinishTime 是期望送达时间
	ExpectedFinishTime *string `json:"expected_finish_time,omitempty"`
	// ExpectedPickTime 是期望取件时间
	ExpectedPickTime *string `json:"expected_pick_time,omitempty"`
	// InsurePrice 是保价金额
	InsurePrice *string `json:"insure_price,omitempty"`
	// IsDirectDelivery 是是否直拿直送，1:直拿直送 0:非直拿直送
	IsDirectDelivery *int64 `json:"is_direct_delivery,omitempty"`
	// IsFinishCodeNeeded 是是否需要收货码，1:需要 0:不需要
	IsFinishCodeNeeded *int64 `json:"is_finish_code_needed,omitempty"`
	// IsInsured 是是否保价，1:保价 0:不保价
	IsInsured *int64 `json:"is_insured,omitempty"`
	// IsPickupCodeNeeded 是是否需要取货码，1:需要 0:不需要
	IsPickupCodeNeeded *int64 `json:"is_pickup_code_needed,omitempty"`
	// PoiSeq 是门店订单流水号
	PoiSeq *string `json:"poi_seq,omitempty"`
	// ServiceCode 是即时配送公司的服务代码
	ServiceCode *string `json:"service_code,omitempty"`
}

// PreCreateWaybillIstd 是支付宝官方 V3 OpenAPI 模型。
type PreCreateWaybillIstd struct {
	// CouponFee 是优惠券费用，单位元
	CouponFee *string `json:"coupon_fee,omitempty"`
	// DeliverFee 是运费，单位元
	DeliverFee *string `json:"deliver_fee,omitempty"`
	// DispatchDuration 是预计骑手还剩多久接单（单位：秒）
	DispatchDuration *int64 `json:"dispatch_duration,omitempty"`
	// Distance 是配送距离(单位：米)
	Distance *int64 `json:"distance,omitempty"`
	// Fee 是实际运费，单位元
	Fee *string `json:"fee,omitempty"`
	// InsuranceFee 是保价费用，单位元
	InsuranceFee *string `json:"insurance_fee,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsToken 是配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变
	LogisticsToken *string `json:"logistics_token,omitempty"`
	// PayAmount 是支付金额，实际扣减的费用以此字段为准，单位元
	PayAmount *string `json:"pay_amount,omitempty"`
	// ServiceCode 是即时配送公司服务代码
	ServiceCode *string `json:"service_code,omitempty"`
	// ThirdCode 是即时配送公司返回的响应码。10000表示业务处理成功，40004表示业务处理失败。
	ThirdCode *string `json:"third_code,omitempty"`
	// ThirdSubCode 是即时配送公司返回的错误码
	ThirdSubCode *string `json:"third_sub_code,omitempty"`
	// ThirdSubMsg 是即时配送公司返回的错误描述
	ThirdSubMsg *string `json:"third_sub_msg,omitempty"`
}

// ReceiverIstd 是支付宝官方 V3 OpenAPI 模型。
type ReceiverIstd struct {
	// Address 是地址(街道、小区、大厦等，用于定位)
	Address *string `json:"address,omitempty"`
	// AddressDetail 是地址详情(楼号、单元号、层号)
	AddressDetail *string `json:"address_detail,omitempty"`
	// City 是城市名称，如杭州市
	City *string `json:"city,omitempty"`
	// CoordinateType 是坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标。目前只支持0:火星坐标
	CoordinateType *int64 `json:"coordinate_type,omitempty"`
	// Lat 是纬度
	Lat *string `json:"lat,omitempty"`
	// Lng 是经度
	Lng *string `json:"lng,omitempty"`
	// MobileNo 是手机号
	MobileNo *string `json:"mobile_no,omitempty"`
	// Name 是姓名
	Name *string `json:"name,omitempty"`
}

// SenderIstd 是支付宝官方 V3 OpenAPI 模型。
type SenderIstd struct {
	// Address 是地址(街道、小区、大厦等，用于定位)
	Address *string `json:"address,omitempty"`
	// AddressDetail 是地址详情(楼号、单元号、层号)
	AddressDetail *string `json:"address_detail,omitempty"`
	// City 是城市名称，如杭州市
	City *string `json:"city,omitempty"`
	// CoordinateType 是坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1:百度坐标， 目前只支持0
	CoordinateType *int64 `json:"coordinate_type,omitempty"`
	// Lat 是纬度
	Lat *string `json:"lat,omitempty"`
	// Lng 是经度
	Lng *string `json:"lng,omitempty"`
	// MobileNo 是手机号
	MobileNo *string `json:"mobile_no,omitempty"`
	// Name 是姓名
	Name *string `json:"name,omitempty"`
}

// ServiceCodeIstd 是支付宝官方 V3 OpenAPI 模型。
type ServiceCodeIstd struct {
	// ServiceCode 是即时配送公司服务代码
	ServiceCode *string `json:"service_code,omitempty"`
}
