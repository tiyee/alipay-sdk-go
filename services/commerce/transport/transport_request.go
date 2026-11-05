// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package transport

// TransportChargerChargerBindInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type TransportChargerChargerBindInfoSyncRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// BindQrcode 是绑定充电桩二维码值，用户扫码绑桩扫描的二维码
	BindQrcode *string `json:"bind_qrcode,omitempty"`
	// BindStatus 是绑定状态 1:绑定; 2:解绑;  0:未知
	BindStatus *string `json:"bind_status,omitempty"`
	// BindTime 是绑定时间，绑定成功必传
	BindTime *string `json:"bind_time,omitempty"`
	// ConnectorType 是1：家用插座	2、交流接口插座	3：交流接口插头	4、直流接口枪头	5、无线充电座	6、其他
	ConnectorType *string `json:"connector_type,omitempty"`
	// CurrentOutputType 是1：国标	2：欧标	3：美标	4：日标	5：其他
	CurrentOutputType *string `json:"current_output_type,omitempty"`
	// EquipId 是设备编号
	EquipId *string `json:"equip_id,omitempty"`
	// EquipImei 是设备IMEI号
	EquipImei *string `json:"equip_imei,omitempty"`
	// EquipName 是设备名称，对用户展示，可以是设备名称，也可以是设备SN号，设备编号
	EquipName *string `json:"equip_name,omitempty"`
	// EquipSn 是设备SN号
	EquipSn *string `json:"equip_sn,omitempty"`
	// EquipStatus 是设备状态，遵循中电联协议的设备状态。0：离网;1：空闲;2：占用未充电;3：占用充电中;4：占用（预约锁定）;255：故障
	EquipStatus *string `json:"equip_status,omitempty"`
	// EquipTypeNo 是设备型号
	EquipTypeNo *string `json:"equip_type_no,omitempty"`
	// NetworkDataBeginTime 是流量开始时间
	NetworkDataBeginTime *string `json:"network_data_begin_time,omitempty"`
	// NetworkDataEndTime 是流量结束时间
	NetworkDataEndTime *string `json:"network_data_end_time,omitempty"`
	// OperatorId 是运营商编码，一般为企业组织机构代码
	OperatorId *string `json:"operator_id,omitempty"`
	// OperatorUid 是商户内部唯一标识用户的用户标识
	OperatorUid *string `json:"operator_uid,omitempty"`
	// RatedCurrent 是额定电流,单位：A
	RatedCurrent *string `json:"rated_current,omitempty"`
	// RatedPower 是额定功率，单位：kW
	RatedPower *string `json:"rated_power,omitempty"`
	// RatedVoltage 是额定电压，单位：V
	RatedVoltage *string `json:"rated_voltage,omitempty"`
	// SBrandCode 是设备有二级品牌商，需要传二级品牌编号
	SBrandCode *string `json:"s_brand_code,omitempty"`
	// UnbindTime 是解绑时间，解绑必传
	UnbindTime *string `json:"unbind_time,omitempty"`
}

// TransportParkingEnterInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingEnterInfoSyncRequest struct {
	// AgreementQuery 是是否启用车牌代扣状态查询功能，true为启用，false为停用
	AgreementQuery *bool `json:"agreement_query,omitempty"`
	// ChargeInfoList 是本次行程收费计划清单
	ChargeInfoList *[]ParkingChargeInfo `json:"charge_info_list,omitempty"`
	// EntranceName 是车辆进入停车场的入口通道名称
	EntranceName *string `json:"entrance_name,omitempty"`
	// EntranceNumber 是车辆进入停车场的入口通道编号；按线下实际编号为准；
	EntranceNumber *int64 `json:"entrance_number,omitempty"`
	// FreeEnterMinutes 是当前停车场的入场免费时长分钟数
	FreeEnterMinutes *string `json:"free_enter_minutes,omitempty"`
	// InTime 是车辆入场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，入场时间不应晚于当前网络时间
	InTime *string `json:"in_time,omitempty"`
	// IsEncryptPlateNo 是车牌是否加密，true为加密，false为不加密，默认为false
	IsEncryptPlateNo *bool                  `json:"is_encrypt_plate_no,omitempty"`
	MembershipInfo   *ParkingMembershipInfo `json:"membership_info,omitempty"`
	// NeedCharge 是当前行程是否需要计费。true：需要，false：不需要。不传默认为true。
	NeedCharge *bool `json:"need_charge,omitempty"`
	// OpenAppid 是蚂蚁会员统一ID对应的归属应用appid
	OpenAppid *string `json:"open_appid,omitempty"`
	// OpenId 是蚂蚁会员统一ID
	OpenId *string `json:"open_id,omitempty"`
	// OutSerialNo 是外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// ParkingId 是支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create	(录入停车场信息)接口获取。
	ParkingId *string `json:"parking_id,omitempty"`
	// PlateColor 是车牌颜色，车牌颜色，枚举支持：	*BLUE：蓝。	*GREEN：绿。	*YELLOW：黄。	*WHITE：白。	*BLACK：黑。	*LIMEGREEN：黄绿色。
	PlateColor *string `json:"plate_color,omitempty"`
	// PlateNo 是车牌号（支持加密格式）
	PlateNo *string `json:"plate_no,omitempty"`
	// ServiceUrl 是停车服务页面地址。	1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F		2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex		3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex
	ServiceUrl *string `json:"service_url,omitempty"`
}

// TransportParkingExitInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingExitInfoSyncRequest struct {
	// ExitName 是车辆出场的出口通道名称；取实地通道名称；
	ExitName *string `json:"exit_name,omitempty"`
	// ExitNumber 是车辆出场的出口通道编号；与线下实地编码一致；
	ExitNumber *int64 `json:"exit_number,omitempty"`
	// IsEncryptPlateNo 是车牌是否加密，true为加密，false为不加密，默认为false
	IsEncryptPlateNo *bool `json:"is_encrypt_plate_no,omitempty"`
	// OpenAppid 是蚂蚁会员统一ID对应的归属应用appid
	OpenAppid *string `json:"open_appid,omitempty"`
	// OpenId 是蚂蚁会员统一ID
	OpenId *string `json:"open_id,omitempty"`
	// OutSerialNo 是外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// OutTime 是车辆出场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，出场时间不应晚于当前网络时间，也不当早于入场时间。
	OutTime *string `json:"out_time,omitempty"`
	// PlateColor 是车牌颜色，车牌颜色，枚举支持：	*BLUE：蓝。	*GREEN：绿。	*YELLOW：黄。	*WHITE：白。	*BLACK：黑。	*LIMEGREEN：黄绿色。
	PlateColor *string `json:"plate_color,omitempty"`
	// PlateNo 是车牌号（支持加密格式）
	PlateNo *string `json:"plate_no,omitempty"`
	// ServiceUrl 是停车服务页面地址。	1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F 2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex 3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex
	ServiceUrl *string `json:"service_url,omitempty"`
}

// TransportParkingPaymentInfoSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingPaymentInfoSyncRequest struct {
	// AlipayTradeNo 是支付宝交易号；
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty"`
	// DiscountAmount 是停车费折扣金额；单位元；
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// DiscountInformation 是停车费缴费折扣信息；
	DiscountInformation *[]ParkingPaymentDiscountInfo `json:"discount_information,omitempty"`
	// FreeExitMinutes 是停车场内缴费后，允许出场的免费时长分钟数
	FreeExitMinutes *string `json:"free_exit_minutes,omitempty"`
	// InactiveUser 是365天未使用支付宝付停车费用户
	InactiveUser *bool `json:"inactive_user,omitempty"`
	// IsEncryptPlateNo 是车牌是否加密，true为加密，false为不加密，默认为false
	IsEncryptPlateNo *bool `json:"is_encrypt_plate_no,omitempty"`
	// MobileNumber 是缴费用户的手机号；
	MobileNumber *string `json:"mobile_number,omitempty"`
	// NodeType 是表示当前请求类型：用户缴费或用户在岗亭前缴费准备离场状态
	NodeType *string `json:"node_type,omitempty"`
	// OpenAppid 是蚂蚁会员统一ID对应的归属应用appid
	OpenAppid *string `json:"open_appid,omitempty"`
	// OpenId 是蚂蚁会员统一ID
	OpenId *string `json:"open_id,omitempty"`
	// OrderCreateTime 是停车岗亭前订单创建时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，创建时不应晚于当前网络时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// OutOrderNo 是服务商停车费支付订单号；
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutSerialNo 是外部停车流水号(用于串通进场与出场信息)
	OutSerialNo *string `json:"out_serial_no,omitempty"`
	// PayFrequency 是用户在支付宝侧支付停车费次数
	PayFrequency *string `json:"pay_frequency,omitempty"`
	// PaymentAmount 是停车费实付金额；单位元；
	PaymentAmount *string `json:"payment_amount,omitempty"`
	// PaymentTime 是场内缴费时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，场内缴费不应晚于当前网络时间
	PaymentTime *string `json:"payment_time,omitempty"`
	// PaymentType 是支付方式；	ALIPAY：支付宝	WECHAT：微信	CUQP：云闪付	CASH：现金	OTHER：其他渠道	FREE: 免费
	PaymentType *string `json:"payment_type,omitempty"`
	// PaymentUserOpenId 是停车缴费用户的OPENID；
	PaymentUserOpenId *string `json:"payment_user_open_id,omitempty"`
	// PlateColor 是车牌颜色，车牌颜色，枚举支持：	*BLUE：蓝。	*GREEN：绿。	*YELLOW：黄。	*WHITE：白。	*BLACK：黑。	*LIMEGREEN：黄绿色。
	PlateColor *string `json:"plate_color,omitempty"`
	// PlateNo 是车牌号（支持加密格式）
	PlateNo *string `json:"plate_no,omitempty"`
	// ServiceUrl 是停车服务页面地址。	1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：  https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F  	2、若服务链接可沿用入场时提供的服务链接，可以留空。	3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex
	ServiceUrl *string `json:"service_url,omitempty"`
	// TotalAmount 是停车费总金额；单位元；理论应等于实付金额+折扣金额
	TotalAmount *string `json:"total_amount,omitempty"`
}
