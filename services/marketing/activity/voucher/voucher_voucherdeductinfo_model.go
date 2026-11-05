package voucher

// VoucherDeductInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherDeductInfo struct {
	CustomerDefineVoucherInfo *CustomerDefineVoucherInfo `json:"customer_define_voucher_info,omitempty"`
	DiscountVoucherInfo       *DiscountVoucherInfo       `json:"discount_voucher_info,omitempty"`
	ExchangeVoucherInfo       *ExchangeVoucherInfo       `json:"exchange_voucher_info,omitempty"`
	FixVoucherInfo            *FixVoucherInfo            `json:"fix_voucher_info,omitempty"`
	SpecialVoucherInfo        *SpecialVoucherInfo        `json:"special_voucher_info,omitempty"`
	// VoucherType 是券类型。
	VoucherType *string `json:"voucher_type,omitempty"`
}

// VoucherDeductThresholdInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherDeductThresholdInfo struct {
	// CustomerDefineVoucherThresholdDesc 是自定义门槛描述
	CustomerDefineVoucherThresholdDesc    *string                                `json:"customer_define_voucher_threshold_desc,omitempty"`
	VoucherExpressInfo                    *VoucherExpressInfo                    `json:"voucher_express_info,omitempty"`
	VoucherGoodsQuantitySpecificationInfo *VoucherGoodsQuantitySpecificationInfo `json:"voucher_goods_quantity_specification_info,omitempty"`
}

// VoucherDisplayPatternInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherDisplayPatternInfo struct {
	// BrandLogo 是商家logo 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值。
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandLogoUrl 是商家品牌logo链接
	BrandLogoUrl *string `json:"brand_logo_url,omitempty"`
	// BrandName 是商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// CustomerServiceMobile 是客服电话。
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// CustomerServiceUrl 是客服链接。
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// VoucherDescription 是券详细使用说明。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// VoucherDetailImageUrls 是券详情链接
	VoucherDetailImageUrls *[]string `json:"voucher_detail_image_urls,omitempty"`
	// VoucherDetailImages 是券详细图列表，会展示在用户支付宝卡包券详情页 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE  ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。
	VoucherDetailImages *[]string `json:"voucher_detail_images,omitempty"`
	// VoucherImage 是券详情页封面图，会展示在用户支付宝卡包券详情页。	需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值。
	VoucherImage *string `json:"voucher_image,omitempty"`
	// VoucherImageUrl 是券封面链接
	VoucherImageUrl *string `json:"voucher_image_url,omitempty"`
	// VoucherName 是券名称。
	VoucherName *string `json:"voucher_name,omitempty"`
}

// VoucherExchangeGoodsInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherExchangeGoodsInfo struct {
	// ExchangeGoodsName 是兑换商品名称
	ExchangeGoodsName *string `json:"exchange_goods_name,omitempty"`
}

// VoucherExpressInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherExpressInfo struct {
	// PayExpress 是邮费标识，属于券的服务，用于用户门槛信息展示
	PayExpress *string `json:"pay_express,omitempty"`
}

// VoucherGoodsQuantitySpecificationInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherGoodsQuantitySpecificationInfo struct {
	// DirectionalType 是定向类型用于规格设置，比如：FIX 表示 固定满X件，APPOINT 表示指定第N件
	DirectionalType *string `json:"directional_type,omitempty"`
	// SpecificationQuantity 是数量
	SpecificationQuantity *int64 `json:"specification_quantity,omitempty"`
	// SpecificationUnit 是单位
	SpecificationUnit *string `json:"specification_unit,omitempty"`
}

// VoucherInventoryInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherInventoryInfo struct {
	// SendCount 是已发放数量。
	SendCount *int64 `json:"send_count,omitempty"`
	// UseCount 是已核销数量。
	UseCount *int64 `json:"use_count,omitempty"`
}

// VoucherMiniAppSendGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherMiniAppSendGuideInfo struct {
	// MiniAppUrl 是小程序发券链接
	MiniAppUrl *string `json:"mini_app_url,omitempty"`
}

// VoucherMiniAppUseGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherMiniAppUseGuideInfo struct {
	// MiniAppServiceCodes 是小程序服务编码。取值来源：通过alipay.open.app.service.apply(小程序服务创建)接口创建服务后获取。
	MiniAppServiceCodes *[]string `json:"mini_app_service_codes,omitempty"`
	// MiniAppUrl 是小程序核销链接
	MiniAppUrl *string `json:"mini_app_url,omitempty"`
}

// VoucherPackageModeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageModeInfo struct {
	// VoucherPackageId 是券包id，对应alipay.marketing.activity.voucherpackage.query中voucher_package_id。
	VoucherPackageId *string `json:"voucher_package_id,omitempty"`
}

// VoucherRechargeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherRechargeInfo struct {
	// RechargeType 是预充值方式。
	RechargeType               *string                     `json:"recharge_type,omitempty"`
	VoucherBalanceRechargeInfo *VoucherBalanceRechargeInfo `json:"voucher_balance_recharge_info,omitempty"`
}

// VoucherRelativePeriodInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherRelativePeriodInfo struct {
	TimeRestrictInfo *TimeRestrictInfo `json:"time_restrict_info,omitempty"`
	// ValidDaysAfterReceive 是券生效后 N 天内可以使用。 可以配合wait_days_after_receive 字段使用。
	ValidDaysAfterReceive *int64 `json:"valid_days_after_receive,omitempty"`
	// ValidDaysRelativeType 是优惠券有效期的取整方式，需与 valid_days_after_receive 字段配合使用，用于确定用户领取优惠券后的具体失效时间。 	取值逻辑说明：ROUND_FORWARD（向前取整），优惠券在领取当日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 	示例： 	-若 valid_days_relative_type=ROUND_FORWARD，valid_days_after_receive=1，用户于 1 号 09:12:40 领券，则优惠券在 1 号 23:59:59 失效。 	-若 valid_days_after_receive=2，则在 2 号 23:59:59 失效。ROUND_BACKWARD（向后取整），优惠券从领取次日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 	示例： 	-若valid_days_relative_type=ROUND_BACKWARDS，valid_days_after_receive=1，领券时间是1号9点12分40秒，那么 2 号 23:59:59 失效。 	-若 valid_days_after_receive=2，则在 3 号 23:59:59 失效。
	ValidDaysRelativeType *string `json:"valid_days_relative_type,omitempty"`
	// WaitDaysAfterReceive 是用户领券后需要等待 N 天，券才可以生效。 	字段值为 0 表示用户领券后立刻生效。
	WaitDaysAfterReceive *int64 `json:"wait_days_after_receive,omitempty"`
}

// VoucherSaleModeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherSaleModeInfo struct {
	// FundCustodyMode 是资金托管方式
	FundCustodyMode *string `json:"fund_custody_mode,omitempty"`
	// OverdueRefundable 是是否支持优惠券过期后，自动退款给用户。 不填默认否，枚举值： true：是  false：否。自动退款功能需要服务商在优惠券过期时，主动调用alipay.marketing.activity.order.refund接口进行退款。  如果配置优惠券时选择了过期自动退款，但是实际券过期后，服务商没有进行退款，那么用户投诉后，需要服务商进行解决。
	OverdueRefundable *bool `json:"overdue_refundable,omitempty"`
	// OverdueRefundableNeedConfirm 是在支持过期自动退款后，支付宝发起自动退款前 是否需要向商户确认退款。 不填默认否，枚举值： true：是 false：否。		如果选择了 true:是 ，则需要商户配合接入  过期自动退款前是否需要向商户确认SPI(spi.alipay.marketing.activity.order.refundconfirm)  以及   退款确认结果API(alipay.marketing.activity.order.refundconfirmcommit)
	OverdueRefundableNeedConfirm *bool `json:"overdue_refundable_need_confirm,omitempty"`
	// PayeePid 是收款账号。 目前的结算规则是，每核销一笔优惠券，支付宝会打款到该收款账户。
	PayeePid *string `json:"payee_pid,omitempty"`
	// Refundable 是购买的优惠券是否允许退款。 不填默认否，枚举值： true：是 false：否
	Refundable *bool `json:"refundable,omitempty"`
	// RefundableNeedConfirm 是在支持随时退后，支付宝发起退款前用于判断是否需要向商户确认退款。 不填默认否，枚举值： true：是 false：否。 如果选择了 true:是  ，则需要商户配合接入退款前是否需要向商户确认SPI(spi.alipay.marketing.activity.order.refundconfirm)
	RefundableNeedConfirm *bool `json:"refundable_need_confirm,omitempty"`
	// SaleAmount 是用户购买优惠券需要支付的金额
	SaleAmount *string            `json:"sale_amount,omitempty"`
	SettleInfo *BizfundSettleInfo `json:"settle_info,omitempty"`
}

// VoucherSendGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendGuideInfo struct {
	MiniAppSendGuideInfo *VoucherMiniAppSendGuideInfo `json:"mini_app_send_guide_info,omitempty"`
	// SendGuideMode 是发放引导模式
	SendGuideMode *[]string `json:"send_guide_mode,omitempty"`
}

// VoucherSendModeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendModeInfo struct {
	// DeliveryChannelList 是推广渠道列表
	DeliveryChannelList    *[]string               `json:"delivery_channel_list,omitempty"`
	VoucherPackageModeInfo *VoucherPackageModeInfo `json:"voucher_package_mode_info,omitempty"`
	VoucherSaleModeInfo    *VoucherSaleModeInfo    `json:"voucher_sale_mode_info,omitempty"`
	// VoucherSendMode 是券发放模式。
	VoucherSendMode     *string              `json:"voucher_send_mode,omitempty"`
	VoucherSendRuleInfo *VoucherSendRuleInfo `json:"voucher_send_rule_info,omitempty"`
}

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

// VoucherSummary 是支付宝官方 V3 OpenAPI 模型。
type VoucherSummary struct {
	// PublishCount 是券已发放数量。
	PublishCount *int64 `json:"publish_count,omitempty"`
	// UsedCount 是券已核销数量。
	UsedCount *int64 `json:"used_count,omitempty"`
}

// VoucherUseGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseGuideInfo struct {
	MiniAppUseGuideInfo *VoucherMiniAppUseGuideInfo `json:"mini_app_use_guide_info,omitempty"`
	// UseGuideMode 是使用引导模式
	UseGuideMode *[]string `json:"use_guide_mode,omitempty"`
	// UseUrl 是核销跳转链接
	UseUrl *string `json:"use_url,omitempty"`
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
