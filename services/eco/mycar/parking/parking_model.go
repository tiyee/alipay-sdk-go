package parking

// BusinessItem 是支付宝官方 V3 OpenAPI 模型。
type BusinessItem struct {
	// BusinessAppid 是业务归属appid
	BusinessAppid *string `json:"business_appid,omitempty"`
	// BusinessPayeeId 是业务收款pid，business_type为AGREEMENT_PAY时，此参数必传
	BusinessPayeeId *string `json:"business_payee_id,omitempty"`
	// BusinessPid 是业务归属pid
	BusinessPid *string `json:"business_pid,omitempty"`
	// BusinessType 是业务类型ONLINE_PAY(在线缴费) 、AGREEMENT_PAY(无感停车)	DEVICE_ONLINE_PAY(车机在线缴费)
	BusinessType *string `json:"business_type,omitempty"`
}

// InterfaceInfoList 是支付宝官方 V3 OpenAPI 模型。
type InterfaceInfoList struct {
	// InterfaceName 是H5传入参数固定值：alipay.eco.mycar.parking.userpage.query;		小程序传入参数固定值：alipay.eco.mycar.parking.isv.homepage;		协议状态变更通知传入参数固定值：	alipay.eco.mycar.parking.agreement.notify;
	InterfaceName *string `json:"interface_name,omitempty"`
	// InterfaceType 是H5、小程序传入参数固定值:interface_page；	通知接口传入参数固定值:interface_service;
	InterfaceType *string `json:"interface_type,omitempty"`
	// InterfaceUrl 是H5 SPI接口的调用地址url，协议必须为https，对整个url字符串必须进行UrlEncode编码。编码为UTF-8;		小程序SPI接口的调用地址ur格式：小程序|小程序缴费页面地址;		协议状态变更通知SPI接口的调用地址url，同H5格式;
	InterfaceUrl *string `json:"interface_url,omitempty"`
}

// ParkingLotChargingRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ParkingLotChargingRuleInfo struct {
	// ChargingPeriod 是停车收费时间段，输入24小时制的小时时间段格式: 	开始时间-结束时间,	结束时间必须大于开始时间	时间取值范围：0到24
	ChargingPeriod *string `json:"charging_period,omitempty"`
	// ChargingStrategy 是计费策略；	UP：向上取整	DOMN：向下取整
	ChargingStrategy *string `json:"charging_strategy,omitempty"`
	// ChargingUnitFee 是跳费金额 单位元
	ChargingUnitFee *string `json:"charging_unit_fee,omitempty"`
	// ChargingUnitMinutes 是跳费间隔分钟数
	ChargingUnitMinutes *string `json:"charging_unit_minutes,omitempty"`
	// EnableRolloverChargePeriod 是超时是否顺延到下个计费区间
	EnableRolloverChargePeriod *bool `json:"enable_rollover_charge_period,omitempty"`
	// FirstChargingSpanMinutes 是首段计费周期分钟数
	FirstChargingSpanMinutes *string `json:"first_charging_span_minutes,omitempty"`
	// FirstChargingUnitFee 是首段计费周期跳费金额 单位: 元
	FirstChargingUnitFee *string `json:"first_charging_unit_fee,omitempty"`
	// FirstChargingUnitMinutes 是首段计费周期跳费间隔 单位分钟
	FirstChargingUnitMinutes *string `json:"first_charging_unit_minutes,omitempty"`
	// FreeEnterMinutes 是入场免费时间 单位:分钟
	FreeEnterMinutes *string `json:"free_enter_minutes,omitempty"`
	// FreeExitMinutes 是离场免费时间 单位:分钟
	FreeExitMinutes *string `json:"free_exit_minutes,omitempty"`
	// MaxFeePerDay 是日封顶金额 单位: 元
	MaxFeePerDay *string `json:"max_fee_per_day,omitempty"`
}

// ParkingLotServiceInfo 是支付宝官方 V3 OpenAPI 模型。
type ParkingLotServiceInfo struct {
	// ServiceName 是在线缴费服务名称
	ServiceName *string `json:"service_name,omitempty"`
	// ServiceType 是停车场可提供的服务类型	发票服务：INVOICE_SERVICE	优惠套餐: DISCOUNT_PACKAGE	预约停车: RESERVED_PARKING
	ServiceType *string `json:"service_type,omitempty"`
	// ServiceUrl 是商户提供服务对应的跳转链接
	ServiceUrl *string `json:"service_url,omitempty"`
}
