package ordervoucher

// VoucherUseGuide 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseGuide struct {
	// MiniAppId 是券可用的小程序appId，卡包详情页可跳转到该appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppPath 是指定跳转到mini_app_id时的具体页面路径。
	MiniAppPath *string `json:"mini_app_path,omitempty"`
	// MiniAppUseGuide 是小程序核销。
	MiniAppUseGuide *bool `json:"mini_app_use_guide,omitempty"`
	// OfflineCodeUseGuide 是线下扫码核销。
	OfflineCodeUseGuide *bool `json:"offline_code_use_guide,omitempty"`
	// ServiceCodes 是小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。
	ServiceCodes *[]string `json:"service_codes,omitempty"`
}

// VoucherUseGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseGuideInfo struct {
	MiniAppUseGuideInfo *VoucherMiniAppUseGuideInfo `json:"mini_app_use_guide_info,omitempty"`
	// UseGuideMode 是使用引导模式
	UseGuideMode *[]string `json:"use_guide_mode,omitempty"`
	// UseUrl 是核销跳转链接
	UseUrl *string `json:"use_url,omitempty"`
}

// VoucherUseGuideModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseGuideModify struct {
	// MiniAppId 是券可用的小程序 appId，卡包详情页可跳转到该 appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppPath 是指定跳转到 mini_app_id 时的具体页面路径。
	MiniAppPath *string `json:"mini_app_path,omitempty"`
	// ServiceCodes 是小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。
	ServiceCodes *[]string `json:"service_codes,omitempty"`
}

// VoucherUseRule 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseRule struct {
	DiscountVoucher       *DiscountVoucher       `json:"discount_voucher,omitempty"`
	ExchangeVoucher       *ExchangeVoucher       `json:"exchange_voucher,omitempty"`
	FixVoucher            *FixVoucher            `json:"fix_voucher,omitempty"`
	SpecialVoucher        *SpecialVoucher        `json:"special_voucher,omitempty"`
	VoucherAvailableScope *VoucherAvailableScope `json:"voucher_available_scope,omitempty"`
	VoucherValidPeriod    *VoucherValidPeriod    `json:"voucher_valid_period,omitempty"`
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

// VoucherUseRuleModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseRuleModify struct {
	ExchangeVoucher       *ExchangeVoucherModify       `json:"exchange_voucher,omitempty"`
	VoucherAvailableScope *VoucherAvailableScopeModify `json:"voucher_available_scope,omitempty"`
	VoucherValidPeriod    *VoucherValidPeriodModify    `json:"voucher_valid_period,omitempty"`
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

// VoucherValidPeriod 是支付宝官方 V3 OpenAPI 模型。
type VoucherValidPeriod struct {
	// Type 是券有效期类型
	Type *string `json:"type,omitempty"`
	// ValidBeginTime 是券可使用的开始时间。格式为：yyyy-MM-dd HH:mm:ss
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// ValidDaysAfterReceive 是券生效后N天内可以使用。可以配合wait_days_after_receive字段使用。
	ValidDaysAfterReceive *int64 `json:"valid_days_after_receive,omitempty"`
	// ValidEndTime 是券可使用的结束时间。格式为yyyy-MM-dd HH:mm:ss
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// WaitDaysAfterReceive 是用户领券后需要等待N天，券才可以生效。如果该字段填入0或者不填写，则用户领券后立刻生效。
	WaitDaysAfterReceive *int64 `json:"wait_days_after_receive,omitempty"`
}

// VoucherValidPeriodModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherValidPeriodModify struct {
	// ValidBeginTime 是券可使用的开始时间。格式为：yyyy-MM-dd HH:mm:ss。
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// ValidDaysAfterReceive 是券生效后 N 天内可以使用。
	ValidDaysAfterReceive *int64 `json:"valid_days_after_receive,omitempty"`
	// ValidEndTime 是券可使用的结束时间。格式为yyyy-MM-dd HH:mm:ss。
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// WaitDaysAfterReceive 是用户领券后需要等待 N 天，券才可以生效。
	WaitDaysAfterReceive *int64 `json:"wait_days_after_receive,omitempty"`
}

// WeekRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type WeekRuleInfo struct {
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
	// WeekDay 是星期数，以“,”连接多个星期数
	WeekDay *string `json:"week_day,omitempty"`
}
