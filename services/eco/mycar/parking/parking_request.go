package parking

// MyCarParkingChargeInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingChargeInfoSyncRequest struct {
	// BusinessHours 是营业时间
	BusinessHours *string `json:"business_hours,omitempty"`
	// CarDaylightAdvancedPrice 是小车-白天 非首个计费单元价格
	CarDaylightAdvancedPrice *string `json:"car_daylight_advanced_price,omitempty"`
	// CarDaylightBasePrice 是小车-白天 首个计费单元价格
	CarDaylightBasePrice *string `json:"car_daylight_base_price,omitempty"`
	// CarNightAdvancedPrice 是小车-夜间 非首个计费单元价格
	CarNightAdvancedPrice *string `json:"car_night_advanced_price,omitempty"`
	// CarNightBasePrice 是小车-夜间 首个计费单元价格
	CarNightBasePrice *string `json:"car_night_base_price,omitempty"`
	// CarOnetimePrice 是小车 计次收费价格 	xx元 / 次，按次收费的车场填写
	CarOnetimePrice *string `json:"car_onetime_price,omitempty"`
	// DailyPriceUpperbound 是单日停车封顶价格：具体值或“不封顶”
	DailyPriceUpperbound *string `json:"daily_price_upperbound,omitempty"`
	// DaylightBusinessHours 是车场白天（收费）营业时间段
	DaylightBusinessHours *string `json:"daylight_business_hours,omitempty"`
	// FreePeriod 是单次停车免费时长
	FreePeriod *string `json:"free_period,omitempty"`
	// HasChargingPile 是是否有充电桩: -1 未知,0 否,1 是
	HasChargingPile *string `json:"has_charging_pile,omitempty"`
	// IsCharge 是是否收费：-1 未知,0 否,1 是
	IsCharge *string `json:"is_charge,omitempty"`
	// NightBusinessHours 是车场夜间（收费）营业时间段
	NightBusinessHours *string `json:"night_business_hours,omitempty"`
	// ParkingId 是停车场ID
	ParkingId *string `json:"parking_id,omitempty"`
	// ParkingSpaceCount 是总车位数
	ParkingSpaceCount *int64 `json:"parking_space_count,omitempty"`
	// Remark 是复杂计费补充描述 	收费复杂无法格式化的车场填写，以文本形式表述
	Remark *string `json:"remark,omitempty"`
	// TruckDaylightAdvancedPrice 是大车-白天 非首个计费单元价格
	TruckDaylightAdvancedPrice *string `json:"truck_daylight_advanced_price,omitempty"`
	// TruckDaylightBasePrice 是大车-白天 首个计费单元价格
	TruckDaylightBasePrice *string `json:"truck_daylight_base_price,omitempty"`
	// TruckNightAdvancedPrice 是大车-夜间 非首个计费单元价格
	TruckNightAdvancedPrice *string `json:"truck_night_advanced_price,omitempty"`
	// TruckNightBasePrice 是大车-夜间 首个计费单元价格
	TruckNightBasePrice *string `json:"truck_night_base_price,omitempty"`
	// TruckOnetimePrice 是大车 计次收费价格 	xx元 / 次，按次收费的车场填写
	TruckOnetimePrice *string `json:"truck_onetime_price,omitempty"`
}

// MyCarParkingChargingInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingChargingInfoSyncRequest struct {
	// CarNumber 是车牌号
	CarNumber *string `json:"car_number,omitempty"`
	// IsvUrl 是智能助理当前的跳转链接
	IsvUrl *string `json:"isv_url,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// SerialNo 是支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 alipay.eco.mycar.parking.enterinfo.sync	(车辆驶入接口)接口获取。
	SerialNo *string `json:"serial_no,omitempty"`
}

// MyCarParkingConfigSetRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingConfigSetRequest struct {
	// AccountNo 是签约支付宝账号
	AccountNo *string `json:"account_no,omitempty"`
	// InterfaceInfoList 是接口信息列表，停车业务需要配置的接口列表，该值为JSON数据格式的LIST对象，现阶段只需要配置一个页面接口即可 。每次请将所有的接口配置信息都传入，未传的接口信息将会被置空。
	InterfaceInfoList *[]InterfaceInfoList `json:"interface_info_list,omitempty"`
	// MerchantLogo 是商户在停车平台首页露出的LOGO。	注意：	*该图片为PNG格式，内容为BASE64的字符串，建议图片尺寸27px*27px，大小不要超过60K。	*若为空则停车平台首页将不露出商户LOGO。
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// MerchantName 是商户简称，由开发者提供
	MerchantName *string `json:"merchant_name,omitempty"`
	// MerchantServicePhone 是商户客服电话
	MerchantServicePhone *string `json:"merchant_service_phone,omitempty"`
}

// MyCarParkingEnterInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingEnterInfoSyncRequest struct {
	// AgreementQuery 是是否启用车牌代扣状态查询功能，true为启用，false为停用
	AgreementQuery *bool `json:"agreement_query,omitempty"`
	// CarColor 是车牌颜色，车牌颜色，枚举支持：	*BLUE：蓝。	*GREEN：绿。	*YELLOW：黄。	*WHITE：白。	*BLACK：黑。	*LIMEGREEN：黄绿色。
	CarColor *string `json:"car_color,omitempty"`
	// CarNumber 是车牌号
	CarNumber *string `json:"car_number,omitempty"`
	// EntranceId 是用于识别车辆进口，多出入口车场适用
	EntranceId *string `json:"entrance_id,omitempty"`
	// FreeMinutes 是当前停车场的入场免费时长分钟数
	FreeMinutes *int64 `json:"free_minutes,omitempty"`
	// InTime 是车辆入场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，入场时间不应晚于当前网络时间
	InTime *string `json:"in_time,omitempty"`
	// IsEncryptCarNumber 是是否加密，默认为false
	IsEncryptCarNumber *bool `json:"is_encrypt_car_number,omitempty"`
	// IsvUrl 是智能助理当前的跳转链接。为ISV/商家小程序scheme地址或H5页面地址。
	IsvUrl *string `json:"isv_url,omitempty"`
	// OutSerialNo 是外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// SpaceNumber 是用于标识车辆停放车位编号，路侧适用
	SpaceNumber *string `json:"space_number,omitempty"`
	// StoreId 是与parking_id对应的外部停车场编号，在门店中维护
	StoreId *string `json:"store_id,omitempty"`
}

// MyCarParkingExitInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingExitInfoSyncRequest struct {
	// ActualAmount 是实际金额
	ActualAmount *string `json:"actual_amount,omitempty"`
	// CarColor 是车牌颜色，枚举支持：	*BLUE：蓝。	*GREEN：绿。	*YELLOW：黄。	*WHITE：白。	*BLACK：黑。	*LIMEGREEN：黄绿色。
	CarColor *string `json:"car_color,omitempty"`
	// CarNumber 是车牌号
	CarNumber *string `json:"car_number,omitempty"`
	// DiscountAmount 是折扣金额
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// ExitId 是用于识别车辆出口，多出入口车场适用
	ExitId *string `json:"exit_id,omitempty"`
	// IsEncryptCarNumber 是是否加密，默认为false
	IsEncryptCarNumber *bool `json:"is_encrypt_car_number,omitempty"`
	// IsvUrl 是智能助理当前的跳转链接
	IsvUrl *string `json:"isv_url,omitempty"`
	// OrderAmount 是订单总金额
	OrderAmount *string `json:"order_amount,omitempty"`
	// OutSerialNo 是外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// OutTime 是车辆离场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
	OutTime *string `json:"out_time,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个ISV或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// SerialNo 是支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 alipay.eco.mycar.parking.enterinfo.sync(车辆驶入接口)接口获取。
	SerialNo *string `json:"serial_no,omitempty"`
	// SpaceNumber 是用于标识车辆停放车位编号，路侧适用
	SpaceNumber *string `json:"space_number,omitempty"`
	// StoreId 是与parking_id对应的外部停车场编号，在门店中维护
	StoreId *string `json:"store_id,omitempty"`
}

// MyCarParkingOrderSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOrderSyncRequest struct {
	// AgentPid 是该笔停车交易需要返佣的对象实体PID，可能是商户，可能是ISV；只做下沉，用于离线表层面对账。
	AgentPid *string `json:"agent_pid,omitempty"`
	// CarNumber 是车牌
	CarNumber *string `json:"car_number,omitempty"`
	// CardNumber 是如果是停车卡缴费，则填入停车卡卡号，否则为'*'
	CardNumber *string `json:"card_number,omitempty"`
	// InDuration 是停车时长（以分为单位）
	InDuration *string `json:"in_duration,omitempty"`
	// InTime 是入场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
	InTime *string `json:"in_time,omitempty"`
	// OpenId 是应用用户ID，唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// OrderNo 是支付宝支付流水，系统唯一
	OrderNo *string `json:"order_no,omitempty"`
	// OrderStatus 是设备商订单状态，枚举支持：	* 0：成功。	* 1：失败。
	OrderStatus *string `json:"order_status,omitempty"`
	// OrderTime 是订单创建时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
	OrderTime *string `json:"order_time,omitempty"`
	// OutOrderNo 是设备商订单号，由ISV系统生成
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutParkingId 是ISV停车场ID，由ISV定义的停车场标识，同一个isv或商户范围内唯一。需与 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口传入值一致。
	OutParkingId *string `json:"out_parking_id,omitempty"`
	// OutTime 是出场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
	OutTime *string `json:"out_time,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// ParkingName 是停车场名称，需与alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口传入值一致。
	ParkingName *string `json:"parking_name,omitempty"`
	// ParkingRecordId 是支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 alipay.eco.mycar.parking.enterinfo.sync(车辆驶入接口)接口获取。
	ParkingRecordId *string `json:"parking_record_id,omitempty"`
	// PayMoney 是缴费金额，保留小数点后两位
	PayMoney *string `json:"pay_money,omitempty"`
	// PayScene 是支付场景：INPARKINGLOT_PAY：场内在线缴费；ENTRANCE_EXIT_PAY：出入口缴费、VEH_DEVICE_PAY：车机缴费；会根据场景判断是否发放能量，当前只有场内支付场景会发能量，需要能量发放请与服务接入支持同学提前沟通。
	PayScene *string `json:"pay_scene,omitempty"`
	// PayTime 是缴费时间, 格式"YYYYMM-DD HH:mm:ss"，24小时制
	PayTime *string `json:"pay_time,omitempty"`
	// PayType 是付款方式，枚举支持：	*1：支付宝在线缴费。
	PayType *string `json:"pay_type,omitempty"`
	// Smid 是间连商户ID
	Smid *string `json:"smid,omitempty"`
	// UserId 是停车缴费用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。	注意：ISV需保证用户 id的正确性，以免导致用户在停车平台查询不到相关的订单信息。
	UserId *string `json:"user_id,omitempty"`
}

// MyCarParkingOrderUpdateRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOrderUpdateRequest struct {
	// OpenId 是应用用户ID，唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// OrderNo 是支付宝支付流水号，系统唯一
	OrderNo *string `json:"order_no,omitempty"`
	// OrderStatus 是用户停车订单状态，枚举支持：	*0：成功。	*1：失败。
	OrderStatus *string `json:"order_status,omitempty"`
	// UserId 是停车缴费用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。	注意：ISV需保证用户 id 的正确性，以免导致用户在停车平台查询不到相关的订单信息。
	UserId *string `json:"user_id,omitempty"`
}

// MyCarParkingOvertimeChargingInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOvertimeChargingInfoSyncRequest struct {
	// CarNumber 是车牌号
	CarNumber *string `json:"car_number,omitempty"`
	// IsvUrl 是智能助理当前的跳转链接
	IsvUrl *string `json:"isv_url,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// SerialNo 是支付宝业务流水号，用于记录车辆从驶入到驶出的全流程
	SerialNo *string `json:"serial_no,omitempty"`
}

// MyCarParkingParkingLotBizTransferRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotBizTransferRequest struct {
	// BusinessIsv 是车场业务归属关系列表
	BusinessIsv *[]BusinessItem `json:"business_isv,omitempty"`
	// ParkingId 是车场创建返回的parking_id字段
	ParkingId *string `json:"parking_id,omitempty"`
}

// MyCarParkingParkingLotInfoCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotInfoCreateRequest struct {
	// AgentId 是服务商的支付宝账号ID（2088开头的16位纯数字，获取方式），由服务商提供给ISV
	AgentId *string `json:"agent_id,omitempty"`
	// BusinessIsv 是车场业务归属列表
	BusinessIsv *[]BusinessItem `json:"business_isv,omitempty"`
	// ChargeSumSpace 是停车场充电车位数 单位:个
	ChargeSumSpace *string `json:"charge_sum_space,omitempty"`
	// ChargingRule 是停车场详细计费规则
	ChargingRule *[]ParkingLotChargingRuleInfo `json:"charging_rule,omitempty"`
	// CityId 是该参数废弃
	CityId *string `json:"city_id,omitempty"`
	// ContactAlipay 是该参数废弃
	ContactAlipay *string `json:"contact_alipay,omitempty"`
	// ContactEmali 是该参数废弃
	ContactEmali *string `json:"contact_emali,omitempty"`
	// ContactMobile 是该参数废弃
	ContactMobile *string `json:"contact_mobile,omitempty"`
	// ContactName 是该参数废弃
	ContactName *string `json:"contact_name,omitempty"`
	// ContactTel 是该参数废弃
	ContactTel *string `json:"contact_tel,omitempty"`
	// ContactWeixin 是该参数废弃
	ContactWeixin *string `json:"contact_weixin,omitempty"`
	// EquipmentName 是该参数废弃
	EquipmentName *string `json:"equipment_name,omitempty"`
	// IsSupportInvoice 是是否支持电子发票。枚举支持：	0：表示不支持。	1：表示支持。
	IsSupportInvoice *string `json:"is_support_invoice,omitempty"`
	// IsvMobile 是ISV电话，传入original_isv_pid、original_isv_appid时，此处为服务商电话
	IsvMobile *string `json:"isv_mobile,omitempty"`
	// Latitude 是纬度 单位:度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度 单位:度
	Longitude *string `json:"longitude,omitempty"`
	// MchntId 是收款方的支付宝账号ID（2088开头的16位纯数字，获取方式），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领
	MchntId *string `json:"mchnt_id,omitempty"`
	// OriginalIsvAppid 是ISV的APPID,服务商调用必传，由ISV提供给服务商
	OriginalIsvAppid *string `json:"original_isv_appid,omitempty"`
	// OriginalIsvMobile 是ISV电话，此值与isv_mobile组合使用
	OriginalIsvMobile *string `json:"original_isv_mobile,omitempty"`
	// OriginalIsvName 是ISV的名称，服务商调用必传，由ISV提供给服务商
	OriginalIsvName *string `json:"original_isv_name,omitempty"`
	// OriginalIsvPid 是ISV的PID，服务商调用必传,由ISV提供给服务商
	OriginalIsvPid *string `json:"original_isv_pid,omitempty"`
	// OutParkingId 是ISV停车场ID，由ISV 自定义，同一个isv或商户范围内唯一
	OutParkingId *string `json:"out_parking_id,omitempty"`
	// ParkingAddress 是停车场地址
	ParkingAddress *string `json:"parking_address,omitempty"`
	// ParkingEndTime 是该参数废弃
	ParkingEndTime *string `json:"parking_end_time,omitempty"`
	// ParkingFeeDescription 是收费说明
	ParkingFeeDescription *string `json:"parking_fee_description,omitempty"`
	// ParkingFeeDescriptionImg 是商户在停车平台露出的停车价格图片；注意：该图片为PNG格式内容为BASE64的字符串，大小不要超过1MB
	ParkingFeeDescriptionImg *string `json:"parking_fee_description_img,omitempty"`
	// ParkingLotType 是停车场类型，COMMUNITY为居民小区、BUSINESS_AREA为商圈停车场（购物中心商业广场商场等）、ROADSIDE为路侧停车、PARK_SCENIC为公园景点（景点乐园公园老街古镇等）、OFFICE_BUILDING为商务楼宇（酒店写字楼商务楼园区等）、OTHER为其他、TRANSPORTATION为交通枢纽（机场火车站汽车站码头港口等）、PUBLIC_FACILITIES为市政设施（体育场博物图书馆医院学校等）、TERRITORY独立园区（办公工业物流园区等）、BUSINESS_PLACE经营场所（4S店、门市餐饮等
	ParkingLotType *string `json:"parking_lot_type,omitempty"`
	// ParkingMobile 是停车场客服电话
	ParkingMobile *string `json:"parking_mobile,omitempty"`
	// ParkingName 是停车场名称，由ISV定义，尽量与高德地图上的一致。
	ParkingName *string `json:"parking_name,omitempty"`
	// ParkingNumber 是该参数废弃
	ParkingNumber *string `json:"parking_number,omitempty"`
	// ParkingPoiid 是如何获取parking_poiid（高德地图唯一标标识）参考文档 https://opendocs.alipay.com/support/01rghx ；若无法成功获取高德POI时，也可选用经纬度的方式注册。
	ParkingPoiid *string `json:"parking_poiid,omitempty"`
	// ParkingStartTime 是该参数废弃
	ParkingStartTime *string `json:"parking_start_time,omitempty"`
	// ParkingType 是该参数废弃
	ParkingType *string `json:"parking_type,omitempty"`
	// PayType 是支付方式。枚举支持：	1：表示支付宝在线缴费。	2：表示支付宝代扣缴费。	3：表示当面付。	4:  现金	说明：如支持多种方式以 ',' 进行分隔。
	PayType *string `json:"pay_type,omitempty"`
	// PaymentMode 是该参数废弃
	PaymentMode *string `json:"payment_mode,omitempty"`
	// SerivceUrl 是提供用户查看停车状态、缴费功能
	SerivceUrl *string `json:"serivce_url,omitempty"`
	// ServiceList 是停车场服务信息列表
	ServiceList *[]ParkingLotServiceInfo `json:"service_list,omitempty"`
	// ShopingmallId 是商圈id
	ShopingmallId *string `json:"shopingmall_id,omitempty"`
	// SumSpace 是停车场车位数 单位:个
	SumSpace *string `json:"sum_space,omitempty"`
	// SupportAfterPay 是是否支持先离后付。枚举支持：	Y：支持。	N：不支持。	说明：默认为空不支持，此参数适用于签约信用代扣的商户。
	SupportAfterPay *string `json:"support_after_pay,omitempty"`
	// TimeOut 是用户支付未离场的超时时间(以分钟为单位)
	TimeOut *string `json:"time_out,omitempty"`
}

// MyCarParkingParkingLotInfoUpdateRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotInfoUpdateRequest struct {
	// AgentId 是服务商ID（2088开头的16位纯数字），由服务商提供给ISV
	AgentId *string `json:"agent_id,omitempty"`
	// BusinessIsv 是车场业务归属列表
	BusinessIsv *[]BusinessItem `json:"business_isv,omitempty"`
	// ChargeSumSpace 是停车场充电车位数 单位:个
	ChargeSumSpace *string `json:"charge_sum_space,omitempty"`
	// ChargingRule 是停车场详细计费规则详细说明
	ChargingRule *[]ParkingLotChargingRuleInfo `json:"charging_rule,omitempty"`
	// CityId 是该参数废弃
	CityId *string `json:"city_id,omitempty"`
	// ContactAlipay 是该参数废弃
	ContactAlipay *string `json:"contact_alipay,omitempty"`
	// ContactEmail 是该参数废弃
	ContactEmail *string `json:"contact_email,omitempty"`
	// ContactMobile 是该参数废弃
	ContactMobile *string `json:"contact_mobile,omitempty"`
	// ContactName 是该参数废弃
	ContactName *string `json:"contact_name,omitempty"`
	// ContactTel 是该参数废弃
	ContactTel *string `json:"contact_tel,omitempty"`
	// ContactWeixin 是该参数废弃
	ContactWeixin *string `json:"contact_weixin,omitempty"`
	// EquipmentName 是该参数废弃
	EquipmentName *string `json:"equipment_name,omitempty"`
	// IsSupportInvoice 是是否支持电子发票。枚举支持： 	*0：表示不支持。 	*1：表示支持。
	IsSupportInvoice *string `json:"is_support_invoice,omitempty"`
	// IsvMobile 是ISV电话，传入original_isv_pid、original_isv_appid时，此处为服务商电话
	IsvMobile *string `json:"isv_mobile,omitempty"`
	// Latitude 是纬度，单位:度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度 单位:度
	Longitude *string `json:"longitude,omitempty"`
	// MchntId 是收款方ID（2088开头的16位纯数字），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领，更新时不影响交易接口中的收款方值。
	MchntId *string `json:"mchnt_id,omitempty"`
	// OriginalIsvAppid 是ISV的APPID,服务商调用必传，由ISV提供给服务商
	OriginalIsvAppid *string `json:"original_isv_appid,omitempty"`
	// OriginalIsvMobile 是ISV电话，此值与isv_mobile组合使用
	OriginalIsvMobile *string `json:"original_isv_mobile,omitempty"`
	// OriginalIsvName 是ISV的名称，服务商调用必传，由ISV提供给服务商
	OriginalIsvName *string `json:"original_isv_name,omitempty"`
	// OriginalIsvPid 是ISV的PID，服务商调用必传,由ISV提供给服务商
	OriginalIsvPid *string `json:"original_isv_pid,omitempty"`
	// OutParkingId 是该参数废弃
	OutParkingId *string `json:"out_parking_id,omitempty"`
	// ParkingAddress 是停车场地址
	ParkingAddress *string `json:"parking_address,omitempty"`
	// ParkingEndTime 是该参数废弃
	ParkingEndTime *string `json:"parking_end_time,omitempty"`
	// ParkingFeeDescription 是收费说明
	ParkingFeeDescription *string `json:"parking_fee_description,omitempty"`
	// ParkingFeeDescriptionImg 是商户在停车平台露出的停车价格图片；注意：该图片为PNG格式内容为BASE64的字符串，大小不要超过1MB
	ParkingFeeDescriptionImg *string `json:"parking_fee_description_img,omitempty"`
	// ParkingId 是支付宝返回停车场id，系统唯一
	ParkingId *string `json:"parking_id,omitempty"`
	// ParkingLotType 是停车场类型，COMMUNITY为居民小区、BUSINESS_AREA为商圈停车场（购物中心商业广场商场等）、ROADSIDE为路侧停车、PARK_SCENIC为公园景点（景点乐园公园老街古镇等）、OFFICE_BUILDING为商务楼宇（酒店写字楼商务楼园区等）、OTHER为其他、TRANSPORTATION为交通枢纽（机场火车站汽车站码头港口等）、PUBLIC_FACILITIES为市政设施（体育场博物图书馆医院学校等）、TERRITORY独立园区（办公工业物流园区等）、BUSINESS_PLACE经营场所（4S店、门市餐饮等
	ParkingLotType *string `json:"parking_lot_type,omitempty"`
	// ParkingMobile 是停车场客服电话
	ParkingMobile *string `json:"parking_mobile,omitempty"`
	// ParkingName 是停车场名称，由ISV定义，尽量与高德地图上的一致
	ParkingName *string `json:"parking_name,omitempty"`
	// ParkingNumber 是该参数废弃
	ParkingNumber *string `json:"parking_number,omitempty"`
	// ParkingPoiid 是如何获取parking_poiid（高德地图唯一标标识）参考文档 https://opendocs.alipay.com/support/01rghx ；若无法成功获取高德POI时，也可选用经纬度的方式注册。
	ParkingPoiid *string `json:"parking_poiid,omitempty"`
	// ParkingStartTime 是该参数废弃
	ParkingStartTime *string `json:"parking_start_time,omitempty"`
	// ParkingType 是该参数废弃
	ParkingType *string `json:"parking_type,omitempty"`
	// PayType 是支付方式。枚举支持：	*1：表示支付宝在线缴费。 	*2：表示支付宝代扣缴费。 	*3：表示当面付。 	*4：表示现金。 		说明：如支持多种方式以 ',' 进行分隔。
	PayType *string `json:"pay_type,omitempty"`
	// PaymentMode 是该参数废弃
	PaymentMode *string `json:"payment_mode,omitempty"`
	// RemainSpace 是剩余车位数 单位：个
	RemainSpace *string `json:"remain_space,omitempty"`
	// SerivceUrl 是提供用户查看停车状态、缴费功能
	SerivceUrl *string `json:"serivce_url,omitempty"`
	// ServiceList 是停车场场内服务信息
	ServiceList *[]ParkingLotServiceInfo `json:"service_list,omitempty"`
	// ShopingmallId 是商圈id
	ShopingmallId *string `json:"shopingmall_id,omitempty"`
	// SumSpace 是停车场车位数 ，单位:个
	SumSpace *string `json:"sum_space,omitempty"`
	// SupportAfterPay 是是否支持先离后付。枚举支持： 	*Y：支持。 	*N：不支持。 	说明：默认为空不支持，此参数适用于签约信用代扣的商户。
	SupportAfterPay *string `json:"support_after_pay,omitempty"`
	// TimeOut 是用户支付未离场的超时时间(以分钟为单位)
	TimeOut *string `json:"time_out,omitempty"`
}

// MyCarParkingPaymentInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingPaymentInfoSyncRequest struct {
	// CarNumber 是车牌号
	CarNumber *string `json:"car_number,omitempty"`
	// IsvUrl 是停车消息透出后，消息详情的落地承接页链接，可配置ISV维度的统一值
	IsvUrl *string `json:"isv_url,omitempty"`
	// OutSerialNo 是外部业务停车流水号(用于串通进场与出场信息)。如果	out_serial_no与serial_no都传递，则一起校验。	● serial_no与out_serial_no二选一必填。	● 入参有serial_no和out_serial_no，则优先根据serial_no和out_serial_no一起匹配进出场事件	● 入参仅有serial_no，则优先根据serial_no匹配进出场事件	● 入参仅有out_serial_no，则根据out_serial_no匹配进出场事件
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// PaymentFreeMinutes 是停车场内缴费后，允许出场的免费时长分钟数
	PaymentFreeMinutes *int64 `json:"payment_free_minutes,omitempty"`
	// PaymentTime 是场内缴费时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，场内缴费不应晚于当前网络时间
	PaymentTime *string `json:"payment_time,omitempty"`
	// SerialNo 是支付宝业务流水号(用于串通进场与出场信息)。可通过 alipay.eco.mycar.parking.enterinfo.sync	(车辆驶入接口)接口获取。	● serial_no与out_serial_no二选一必填。	● 入参有serial_no和out_serial_no，则优先根据serial_no和out_serial_no一起匹配进出场事件	● 入参仅有serial_no，则优先根据serial_no匹配进出场事件	● 入参仅有out_serial_no，则根据out_serial_no匹配进出场事件
	SerialNo *string `json:"serial_no,omitempty"`
}

// MyCarParkingSpaceInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingSpaceInfoSyncRequest struct {
	// FreeChargingPile 是空闲充电桩车位数
	FreeChargingPile *int64 `json:"free_charging_pile,omitempty"`
	// FreeParkingSpace 是车场空闲可用的车位数
	FreeParkingSpace *int64 `json:"free_parking_space,omitempty"`
	// ParkingId 是停车场ID
	ParkingId *string `json:"parking_id,omitempty"`
	// ParkingSpaceStatus 是停车场实时状态: 0 拥堵,1 正常，2 空闲
	ParkingSpaceStatus *string `json:"parking_space_status,omitempty"`
}

// MyCarParkingConfigQueryRequest is the query request model for alipay.eco.mycar.parking.config.query.
type MyCarParkingConfigQueryRequest struct {
	// InterfaceName 是接口名称。H5传入参数固定值：alipay.eco.mycar.parking.userpage.query；	小程序传入参数固定值：alipay.eco.mycar.parking.isv.homepage；	协议状态变更通知传入参数固定值： alipay.eco.mycar.parking.agreement.notify。
	InterfaceName *string `json:"interface_name,omitempty" query:"interface_name,omitempty"`
	// InterfaceType 是接口类型。H5、小程序传入参数固定值：interface_page；	通知接口传入参数固定值：interface_service。
	InterfaceType *string `json:"interface_type,omitempty" query:"interface_type,omitempty"`
}

// MyCarParkingParkingLotInfoQueryRequest is the query request model for alipay.eco.mycar.parking.parkinglotinfo.query.
type MyCarParkingParkingLotInfoQueryRequest struct {
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个ISV或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。	注意：parking_id和out_parking_id不能同时为空。
	ParkingId *string `json:"parking_id,omitempty" query:"parking_id,omitempty"`
	// OutParkingId 是ISV停车场ID，由ISV定义的停车场标识，同一个ISV或商户范围内唯一。需与 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口传入值一致。	注意：parking_id和out_parking_id不能同时为空。
	OutParkingId *string `json:"out_parking_id,omitempty" query:"out_parking_id,omitempty"`
}

// MyCarParkingVehicleQueryRequest is the query request model for alipay.eco.mycar.parking.vehicle.query.
type MyCarParkingVehicleQueryRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
	// CarId 是支付宝用户车辆ID，系统唯一。（该参数会在停车平台用户点击查询缴费，跳转到ISV停车缴费查询页面时，从请求中传递）
	CarId *string `json:"car_id,omitempty" query:"car_id,omitempty"`
}
