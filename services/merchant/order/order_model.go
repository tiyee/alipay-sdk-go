package order

// DiscountInfoData 是支付宝官方 V3 OpenAPI 模型。
type DiscountInfoData struct {
	// DiscountAmount 是优惠金额
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// DiscountName 是优惠名称
	DiscountName *string `json:"discount_name,omitempty"`
	// DiscountPageLink 是优惠跳转链接地址
	DiscountPageLink *string `json:"discount_page_link,omitempty"`
	// DiscountQuantity 是优惠数量
	DiscountQuantity *int64 `json:"discount_quantity,omitempty"`
	// ExternalDiscountId 是外部优惠id
	ExternalDiscountId *string `json:"external_discount_id,omitempty"`
}

// FunctionalService 是支付宝官方 V3 OpenAPI 模型。
type FunctionalService struct {
	// Content 是功能内容
	Content *string `json:"content,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// FunctionCode 是功能码	支持的功能码请与产品或对应技术确认
	FunctionCode *string `json:"function_code,omitempty"`
	// FunctionName 是功能名称
	FunctionName *string `json:"function_name,omitempty"`
	// FunctionType 是功能类型
	FunctionType *string `json:"function_type,omitempty"`
	// FunctionUrl 是功能入口
	FunctionUrl *string `json:"function_url,omitempty"`
	// Memo 是备注信息
	Memo *string `json:"memo,omitempty"`
}

// ItemOrderInfo 是支付宝官方 V3 OpenAPI 模型。
type ItemOrderInfo struct {
	// ExtInfo 是扩展信息，请参见产品文档。小程序订单助手业务中，扩展参数必须传递素材id；其他业务场景参见对应的产品文档。
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// ItemId 是商品ID
	ItemId *string `json:"item_id,omitempty"`
	// ItemName 是商品名称
	ItemName *string `json:"item_name,omitempty"`
	// Quantity 是商品数量（单位：自拟）
	Quantity *int64 `json:"quantity,omitempty"`
	// SkuId 是商品 sku id
	SkuId *string `json:"sku_id,omitempty"`
	// Status 是商品状态枚举
	Status *string `json:"status,omitempty"`
	// StatusDesc 是商品状态描述，默认无需传入，如需使用请联系业务负责人
	StatusDesc *string `json:"status_desc,omitempty"`
	// Unit 是商品规格
	Unit *string `json:"unit,omitempty"`
	// UnitPrice 是商品单价（单位：元）
	UnitPrice *string `json:"unit_price,omitempty"`
}

// JourneyLocation 是支付宝官方 V3 OpenAPI 模型。
type JourneyLocation struct {
	// AoiId 是支付宝域内aoiId
	AoiId *string `json:"aoi_id,omitempty"`
	// City 是城市名
	City *string `json:"city,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// LocationId 是支付宝侧地点id
	LocationId *string `json:"location_id,omitempty"`
	// LocationIdType 是地点id类型
	LocationIdType *string `json:"location_id_type,omitempty"`
	// MerchantDivisionId 是商户侧行政区划代码
	MerchantDivisionId *string `json:"merchant_division_id,omitempty"`
	// MerchantId 是(废弃)商户侧地点id
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantPoi 是商户侧poi信息
	MerchantPoi *string `json:"merchant_poi,omitempty"`
	// Name 是地点名称
	Name *string `json:"name,omitempty"`
	// PoiId 是支付宝域内poiId
	PoiId *string `json:"poi_id,omitempty"`
}

// JourneyMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type JourneyMerchantInfo struct {
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// Logo 是商户logo链接
	Logo *string `json:"logo,omitempty"`
	// Name 是商家名称
	Name *string `json:"name,omitempty"`
	// ShortName 是商家简称
	ShortName *string `json:"short_name,omitempty"`
}

// JourneyServiceChangeInfo 是支付宝官方 V3 OpenAPI 模型。
type JourneyServiceChangeInfo struct {
	// ChangeStatus 是变更状态
	ChangeStatus *string `json:"change_status,omitempty"`
	// DetailUrl 是详情url
	DetailUrl *string `json:"detail_url,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// RemindContent 是提醒内容
	RemindContent *string `json:"remind_content,omitempty"`
}

// OrderDataDistributeInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderDataDistributeInfo struct {
	// NotDistributeReason 是未分发到场景的具体原因。开发者可根据具体原因定位解决问题后发起重试。
	NotDistributeReason *string `json:"not_distribute_reason,omitempty"`
	// SceneCode 是分发场景code。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneName 是分发场景名，对应scene_code。
	SceneName *string `json:"scene_name,omitempty"`
}

// OrderDataSyncSuggestion 是支付宝官方 V3 OpenAPI 模型。
type OrderDataSyncSuggestion struct {
	// Message 是同步建议内容
	Message *string `json:"message,omitempty"`
	// Type 是同步建议类型
	Type *string `json:"type,omitempty"`
}

// OrderExtInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderExtInfo struct {
	// ExtKey 是键值
	ExtKey *string `json:"ext_key,omitempty"`
	// ExtValue 是值
	ExtValue *string `json:"ext_value,omitempty"`
}

// OrderJourneyElement 是支付宝官方 V3 OpenAPI 模型。
type OrderJourneyElement struct {
	Arrival   *JourneyLocation `json:"arrival,omitempty"`
	Departure *JourneyLocation `json:"departure,omitempty"`
	// Duration 是行程时长
	Duration *string `json:"duration,omitempty"`
	// EndTime 是结束时间
	EndTime *string `json:"end_time,omitempty"`
	// EndTimeDesc 是结束时间描述（非结构化）
	EndTimeDesc *string `json:"end_time_desc,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// FunctionalServices 是功能服务列表
	FunctionalServices *[]FunctionalService `json:"functional_services,omitempty"`
	// Passagers 是出行人
	Passagers         *[]UserInfomation         `json:"passagers,omitempty"`
	ServiceChangeInfo *JourneyServiceChangeInfo `json:"service_change_info,omitempty"`
	ServiceProvider   *JourneyMerchantInfo      `json:"service_provider,omitempty"`
	// StartTime 是开始时间
	StartTime *string `json:"start_time,omitempty"`
	// StartTimeDesc 是开始时间描述（非结构化）
	StartTimeDesc *string `json:"start_time_desc,omitempty"`
}

// OrderJourneyInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderJourneyInfo struct {
	// Action 是操作动作
	Action *string `json:"action,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// JourneyCreateTime 是行程创建时间
	JourneyCreateTime *string `json:"journey_create_time,omitempty"`
	// JourneyDesc 是行程描述
	JourneyDesc *string `json:"journey_desc,omitempty"`
	// JourneyElements 是行程元素列表
	JourneyElements *[]OrderJourneyElement `json:"journey_elements,omitempty"`
	// JourneyIndex 是描述本行程为整个行程中的第几程
	JourneyIndex *string `json:"journey_index,omitempty"`
	// JourneyModifyTime 是行程修改时间
	JourneyModifyTime *string `json:"journey_modify_time,omitempty"`
	// MerchantJourneyNo 是商户行程单号
	MerchantJourneyNo *string `json:"merchant_journey_no,omitempty"`
	// Status 是行程状态
	Status *string `json:"status,omitempty"`
	// StatusDesc 是行程状态描述
	StatusDesc *string `json:"status_desc,omitempty"`
	// SubType 是行程子类型
	SubType *string `json:"sub_type,omitempty"`
	// Title 是行程标题
	Title *string `json:"title,omitempty"`
	// Type 是行程类型
	Type *string `json:"type,omitempty"`
}

// OrderParticipantInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderParticipantInfo struct {
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// LogoMaterialId 是参与方logo素材id	开发者通过 alipay.merchant.item.file.upload(商品文件上传接口)上传图片，获取到对应的素材 ID( material_id )
	LogoMaterialId *string `json:"logo_material_id,omitempty"`
	// Name 是参与方名称
	Name *string `json:"name,omitempty"`
	// ShortName 是参与方简称
	ShortName *string `json:"short_name,omitempty"`
	// Type 是参与方类型
	Type *string `json:"type,omitempty"`
	// Uid 是参与方支付宝uid
	Uid *string `json:"uid,omitempty"`
}

// OrderShopInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderShopInfo struct {
	// Address 是店铺地址
	Address *string `json:"address,omitempty"`
	// AlipayShopId 是蚂蚁门店shop_id
	AlipayShopId *string `json:"alipay_shop_id,omitempty"`
	// ExtInfo 是门店其他业务属性，不同业务场景KEY枚举值不同，使用前请参考产品文档
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// MerchantShopId 是商户门店id	支持英文、数字的组合
	MerchantShopId *string `json:"merchant_shop_id,omitempty"`
	// MerchantShopLinkPage 是店铺详情链接地址
	MerchantShopLinkPage *string `json:"merchant_shop_link_page,omitempty"`
	// Name 是店铺名称
	Name *string `json:"name,omitempty"`
	// PhoneNum 是联系电话-支持固话或手机号 仅支持数字、+、- 。例如 手机：1380***1111、固话：021-888**888
	PhoneNum *string `json:"phone_num,omitempty"`
	// Type 是仅当alipay_shop_id字段值为非标准蚂蚁门店时使用，其他场景无需传入
	Type *string `json:"type,omitempty"`
}

// OrderVehicleInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderVehicleInfo struct {
	// LicensePlateNo 是交通工具牌照号，如车牌号等
	LicensePlateNo *string `json:"license_plate_no,omitempty"`
	// Memo 是备注信息
	Memo *string `json:"memo,omitempty"`
	// ShiftNo 是班次
	ShiftNo *string `json:"shift_no,omitempty"`
}

// SeatInfo 是支付宝官方 V3 OpenAPI 模型。
type SeatInfo struct {
	// SeatClass 是座位等级
	SeatClass *string `json:"seat_class,omitempty"`
	// SeatNo 是座位号
	SeatNo *string `json:"seat_no,omitempty"`
}

// TicketInfo 是支付宝官方 V3 OpenAPI 模型。
type TicketInfo struct {
	// Address 是店铺地址
	Address *string `json:"address,omitempty"`
	// EndTime 是截止时间
	EndTime *string `json:"end_time,omitempty"`
	// ResvType 是凭证预约类型，一般不需要传入。可选值为INSTANT/RESERVATION，其中INSTANT代表是实时凭证，RESERVATION代表是预约凭证，不传入默认为实时凭证。区别在于预约凭证一般不是当场可取，而是用户下单后的很多天之后才可以凭借凭证提取
	ResvType *string `json:"resv_type,omitempty"`
	// Shop 是凭证可核销门店/货品自提门店，如果自提门店与购买门店不一致，可传入该字段提示用户自提门店
	Shop *string `json:"shop,omitempty"`
	// TicketNo 是单据号	不同类型下单据号含义不同。若类型为MEAL_NUM，则ticket_no表示取餐号；若类型为PICKUP_CODE，则ticket_no表示取件码；其他类型查看小程序订单产品文档
	TicketNo *string `json:"ticket_no,omitempty"`
	// Time 是时间
	Time *string `json:"time,omitempty"`
	// Type 是凭证类型	具体类型查看产品文档
	Type *string `json:"type,omitempty"`
}

// TicketOrderInfo 是支付宝官方 V3 OpenAPI 模型。
type TicketOrderInfo struct {
	Arrival *JourneyLocation `json:"arrival,omitempty"`
	// Content 是凭证内容
	Content   *string          `json:"content,omitempty"`
	Departure *JourneyLocation `json:"departure,omitempty"`
	// EffectiveNum 是凭证有效数
	EffectiveNum *int64 `json:"effective_num,omitempty"`
	// EndTime 是凭证结束时间
	EndTime *string `json:"end_time,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// FaceValue 是票面价
	FaceValue *string `json:"face_value,omitempty"`
	// FunctionalServices 是功能服务信息
	FunctionalServices *[]FunctionalService `json:"functional_services,omitempty"`
	// InvalidReason 是凭证失效原因
	InvalidReason *string `json:"invalid_reason,omitempty"`
	// Locations 是凭证可使用地点
	Locations *[]JourneyLocation `json:"locations,omitempty"`
	// MerchantTicketNo 是商户侧凭证ID或核销码
	MerchantTicketNo *string `json:"merchant_ticket_no,omitempty"`
	// SeatInfos 是座位信息
	SeatInfos       *[]SeatInfo           `json:"seat_infos,omitempty"`
	ServiceProvider *OrderParticipantInfo `json:"service_provider,omitempty"`
	// Shops 是凭证可使用门店列表
	Shops *[]OrderShopInfo `json:"shops,omitempty"`
	// StartTime 是凭证开始时间
	StartTime *string `json:"start_time,omitempty"`
	// Status 是凭证状态
	Status *string `json:"status,omitempty"`
	// TicketCreateTime 是凭证创建时间
	TicketCreateTime *string `json:"ticket_create_time,omitempty"`
	// TicketModifyTime 是凭证修改时间
	TicketModifyTime *string `json:"ticket_modify_time,omitempty"`
	// TicketNum 是凭证购买数
	TicketNum *int64 `json:"ticket_num,omitempty"`
	// TicketUsers 是凭证使用者信息
	TicketUsers *[]UserInfomation `json:"ticket_users,omitempty"`
	// Title 是凭证标题
	Title       *string           `json:"title,omitempty"`
	VehicleInfo *OrderVehicleInfo `json:"vehicle_info,omitempty"`
}

// UserInfomation 是支付宝官方 V3 OpenAPI 模型。
type UserInfomation struct {
	// CertNo 是证件号
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是证件类型。
	CertType *string `json:"cert_type,omitempty"`
	// ExtInfo 是扩展信息
	ExtInfo *[]OrderExtInfo `json:"ext_info,omitempty"`
	// Mobile 是手机号
	Mobile *string `json:"mobile,omitempty"`
	// Name 是姓名
	Name *string `json:"name,omitempty"`
	// OpenId 是支付宝用户openId
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户ID，可以通过支付结果通知、用户授权等渠道获取
	UserId *string `json:"user_id,omitempty"`
}
