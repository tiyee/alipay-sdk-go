package voucherpackage

// VoucherSendRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendRuleInfo struct {
	// MaxQuantityByDay 是设置此字段，允许指定单天最大发券数量。
	MaxQuantityByDay *int64 `json:"max_quantity_by_day,omitempty"`
	// NaturalPersonLimit 是是否开启自然人领取限制。 自然人表示按照身份证纬度进行领取限制。
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// PhoneNumberLimit 是是否开启电话号码领取限制。
	PhoneNumberLimit *bool `json:"phone_number_limit,omitempty"`
	// PhoneNumberNeedInputLimit 是下单时是否需要用户填写手机号码
	PhoneNumberNeedInputLimit *bool `json:"phone_number_need_input_limit,omitempty"`
	// PublishEndTime 是券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime *string `json:"publish_start_time,omitempty"`
	// Quantity 是发行券的数量。
	Quantity *int64 `json:"quantity,omitempty"`
	// QuantityDayLimitPerUser 是每人每日可领取数量限制，填0或者不填表示不限制。
	QuantityDayLimitPerUser *int64 `json:"quantity_day_limit_per_user,omitempty"`
	// QuantityLimitPerUser 是每人领取限制。 默认按照支付宝账号进行领取限制; 不填写或填入0表示没有领取限制.
	QuantityLimitPerUser *int64 `json:"quantity_limit_per_user,omitempty"`
	// QuantityLimitPerUserPeriodType 是周期限领配置,限制每人在固定周期内领取张数(voucher_quantity_limit_per_user),默认LIFE_CYCLE
	QuantityLimitPerUserPeriodType *string `json:"quantity_limit_per_user_period_type,omitempty"`
	// QuantityMonthLimitPerUser 是每人每月可领取数量限制，填0或者不填表示不限制
	QuantityMonthLimitPerUser *int64 `json:"quantity_month_limit_per_user,omitempty"`
	// QuantityWeekLimitPerUser 是每人每周可领取数量限制，填0或者不填表示不限制
	QuantityWeekLimitPerUser *int64 `json:"quantity_week_limit_per_user,omitempty"`
	// RealNameLimit 是限制支付宝实名用户才能领取支付券,默认为false表示不限制 枚举值 true\false
	RealNameLimit *bool `json:"real_name_limit,omitempty"`
}

// VoucherUseRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseRuleInfo struct {
	// PaymentRestrictionList 是支付方式限制
	PaymentRestrictionList *[]string `json:"payment_restriction_list,omitempty"`
	// QuantityLimitPerUser 是核销次数限制。
	QuantityLimitPerUser *int64 `json:"quantity_limit_per_user,omitempty"`
	// QuantityLimitPerUserPeriodType 是周期限领配置,限制每人在固定周期内领取张数(quantity_limit_per_user)
	QuantityLimitPerUserPeriodType *string `json:"quantity_limit_per_user_period_type,omitempty"`
	// VoucherMaxUseTimes 是券最大核销次数。
	VoucherMaxUseTimes *int64 `json:"voucher_max_use_times,omitempty"`
	// VoucherUseExtInfo 是自定义扩展信息
	VoucherUseExtInfo  *string             `json:"voucher_use_ext_info,omitempty"`
	VoucherUseTimeInfo *VoucherUseTimeInfo `json:"voucher_use_time_info,omitempty"`
}

// VoucherUseTimeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseTimeInfo struct {
	AbsolutePeriodInfo *VoucherAbsolutePeriodInfo `json:"absolute_period_info,omitempty"`
	// PeriodType 是券有效期。
	PeriodType         *string                    `json:"period_type,omitempty"`
	RelativePeriodInfo *VoucherRelativePeriodInfo `json:"relative_period_info,omitempty"`
}

// VoucherUseTimeRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseTimeRuleInfo struct {
	DateRuleInfo    *DateRuleInfo    `json:"date_rule_info,omitempty"`
	HolidayRuleInfo *HolidayRuleInfo `json:"holiday_rule_info,omitempty"`
	// RuleType 是规则类型
	RuleType     *string       `json:"rule_type,omitempty"`
	WeekRuleInfo *WeekRuleInfo `json:"week_rule_info,omitempty"`
}

// WeekRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type WeekRuleInfo struct {
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
	// WeekDay 是星期数，以“,”连接多个星期数
	WeekDay *string `json:"week_day,omitempty"`
}
