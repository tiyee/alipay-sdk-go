package voucher

// ActivityBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityBaseInfo struct {
	ActivityDeliveryExtInfo *ActivityDeliveryExtInfo `json:"activity_delivery_ext_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// ActivityName 是活动名称。
	ActivityName *string `json:"activity_name,omitempty"`
	// ActivityOperationStatus 是活动操作状态。
	ActivityOperationStatus *string `json:"activity_operation_status,omitempty"`
	// ActivityProductType 是券活动产品类型
	ActivityProductType *string `json:"activity_product_type,omitempty"`
	// ActivityStatus 是活动状态 。 ACTIVE:活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。  PAUSE:活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。
	ActivityStatus     *string             `json:"activity_status,omitempty"`
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// CodeMode 是码模式。   MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放  MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。
	CodeMode *string `json:"code_mode,omitempty"`
	// OutActivityId 是外部活动ID
	OutActivityId *string `json:"out_activity_id,omitempty"`
}

// ActivityDeliveryExtInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryExtInfo struct {
	// AutoDelivery 是是否自动推广
	AutoDelivery *bool `json:"auto_delivery,omitempty"`
	// ChannelList 是推广渠道列表
	ChannelList *[]string `json:"channel_list,omitempty"`
}

// AuditResultInfo 是支付宝官方 V3 OpenAPI 模型。
type AuditResultInfo struct {
	// AuditReason 是审核驳回的原因
	AuditReason *string `json:"audit_reason,omitempty"`
	// AuditResult 是审核结果	SUBMIT：审核中	APPROVE：审核通过	REJECT：审核驳回
	AuditResult *string `json:"audit_result,omitempty"`
	// AuditType 是审核类型	UNIFY_ACCESS_CHECK：统一接入平台审核
	AuditType *string `json:"audit_type,omitempty"`
}

// BelongMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type BelongMerchantInfo struct {
	// BusinessType 是合作业务类型,默认为ISV_FOR_MERCHANT
	BusinessType *string `json:"business_type,omitempty"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// BizfundSettleInfo 是支付宝官方 V3 OpenAPI 模型。
type BizfundSettleInfo struct {
	// SettleMode 是请款|账期 模式
	SettleMode *string `json:"settle_mode,omitempty"`
	// SettleTargetType 是结算的目标对象类型
	SettleTargetType *string `json:"settle_target_type,omitempty"`
}

// CustomerDefineDeductRule 是支付宝官方 V3 OpenAPI 模型。
type CustomerDefineDeductRule struct {
	// CustomerDefineDeductRuleDesc 是自定义优惠规则描述
	CustomerDefineDeductRuleDesc *string `json:"customer_define_deduct_rule_desc,omitempty"`
}

// CustomerDefineVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type CustomerDefineVoucherInfo struct {
	DeductInfo *DeductInfo `json:"deduct_info,omitempty"`
}

// DateRangeInfo 是支付宝官方 V3 OpenAPI 模型。
type DateRangeInfo struct {
	// BeginDate 是开始日期  格式：yyyy-MM-dd
	BeginDate *string `json:"begin_date,omitempty"`
	// EndDate 是结束日期  格式：yyyy-MM-dd
	EndDate *string `json:"end_date,omitempty"`
}

// DateRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type DateRuleInfo struct {
	DateRangeInfo *DateRangeInfo `json:"date_range_info,omitempty"`
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
}

// DeductInfo 是支付宝官方 V3 OpenAPI 模型。
type DeductInfo struct {
	CustomerDefineDeductRule *CustomerDefineDeductRule `json:"customer_define_deduct_rule,omitempty"`
}

// DiscountVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type DiscountVoucherInfo struct {
	// CeilingAmount 是封顶金额。
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// Discount 是折扣率。
	Discount *string `json:"discount,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// EndTimeInfo 是支付宝官方 V3 OpenAPI 模型。
type EndTimeInfo struct {
	// EndTime 是结束时间  格式：HH:mm:ss
	EndTime *string `json:"end_time,omitempty"`
	// EndTimeType 是结束时间类型
	EndTimeType *string `json:"end_time_type,omitempty"`
}

// ExchangeVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type ExchangeVoucherInfo struct {
	// Amount 是券的价值
	Amount *string `json:"amount,omitempty"`
	// BizType 是兑换券业务类型。 注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。
	BizType           *string                   `json:"biz_type,omitempty"`
	ExchangeGoodsInfo *VoucherExchangeGoodsInfo `json:"exchange_goods_info,omitempty"`
	// FloorAmount 是优惠门槛金额，表示只有当订单金额大于等于门槛金额时券才能使用。该字段为空时表示无门槛。  门槛金额的校验由服务商(商户)核销时自行校验，支付宝侧只做展示使用。
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// FixVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type FixVoucherInfo struct {
	// Amount 是面额，每张代金券可以抵扣的金额。
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。该字段不填写，认为无门槛。
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// HolidayRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type HolidayRuleInfo struct {
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
}

// PaymentFixVoucher 是支付宝官方 V3 OpenAPI 模型。
type PaymentFixVoucher struct {
	// Amount 是面额，每张优惠券可以抵扣的金额。		限制：	1、币种为人民币，单位为元。小数点以后最多保留两位。	2、取值范围:0.1<=x<=3000
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。该字段不填写，认为无门槛。
	FloorAmount *string `json:"floor_amount,omitempty"`
}

// PaymentVoucherAlipayBalanceRechargeInfo 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherAlipayBalanceRechargeInfo struct {
	// LogonId 是出资的支付宝登录账号	限制:	1、登录账号和用户ID必须且只能二选一
	LogonId *string `json:"logon_id,omitempty"`
	// UserId 是出资的支付宝用户id	限制:	1、登录账号和用户ID必须且只能二选一
	UserId *string `json:"user_id,omitempty"`
}

// PaymentVoucherAvailableGoods 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherAvailableGoods struct {
	// GoodsDescription 是商品描述信息。 用于券面展示，向用户介绍商品
	GoodsDescription *string `json:"goods_description,omitempty"`
	// GoodsIds 是可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。
	GoodsIds *[]string `json:"goods_ids,omitempty"`
	// GoodsName 是商品名称
	GoodsName *string `json:"goods_name,omitempty"`
}

// PaymentVoucherAvailableGoodsModify 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherAvailableGoodsModify struct {
	// GoodsIds 是可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。	限制:	1、核销商品范围只能增大不能减小
	GoodsIds *[]string `json:"goods_ids,omitempty"`
}

// PaymentVoucherAvailableMerchant 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherAvailableMerchant struct {
	// AvailablePids 是优惠券可核销的直连商户PID		限制：	1、available_pids和available_smids至少二选一。
	AvailablePids *[]string `json:"available_pids,omitempty"`
	// AvailableSmids 是优惠券可核销的间连商户SMID		限制： 	1、available_pids和available_smids至少二选一。
	AvailableSmids *[]string `json:"available_smids,omitempty"`
}

// PaymentVoucherAvailableMerchantModify 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherAvailableMerchantModify struct {
	// AvailablePids 是优惠券可以核销的直连商户PID。默认不修改。	 限制： 	1、核销商户范围只能增大不能减小。
	AvailablePids *[]string `json:"available_pids,omitempty"`
	// AvailableSmids 是优惠券可以核销的间连商户SMID。默认不修改。	 限制： 	 1、核销商户范围只能增大不能减小。
	AvailableSmids *[]string `json:"available_smids,omitempty"`
}

// PaymentVoucherBelongMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherBelongMerchantInfo struct {
	// MerchantId 是商户支付宝id，默认品牌名和品牌logo将从该商户信息中获取
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户支付宝id类型。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// PaymentVoucherBudgetInfo 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherBudgetInfo struct {
	// Amount 是营销资金预算,单位元。	总预算=优惠金额*总发券张数
	Amount *string `json:"amount,omitempty"`
	// BudgetType 是营销资金预算类型	1、免充值：制券方无需提前充值资金，用户核销支付券时，直接从订单原价中扣除优惠减价金额，最终只将用户实际支付的金额结算给核销商户，商户实收少于订单原价。	2.预充值:  制券方需将优惠预算提前充值到支付宝指定营销账户中，用户核销支付券时，系统从该账户可用余额中扣除优惠减价部分对应的资金，连同用户实际支付的资金，一并结算给核销商户，不影响实收。
	BudgetType   *string                     `json:"budget_type,omitempty"`
	RechargeInfo *PaymentVoucherRechargeInfo `json:"recharge_info,omitempty"`
}

// PaymentVoucherDisplayInfo 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherDisplayInfo struct {
	// BrandName 是商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// CustomerServiceMobile 是商家客服电话，会展示在支付宝卡包券详情页。
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// VoucherDescription 是用于说明详细的活动规则，会展示在支付宝卡包券详情页。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// VoucherDetailImages 是券详细图列表，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB 限制 1.voucher_image填入，该值才能填入； 2.最多3张；
	VoucherDetailImages *[]string `json:"voucher_detail_images,omitempty"`
	// VoucherImage 是券详情页封面图，会展示在用户支付宝卡包券详情页 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值  限制: 该字段在兑换券场景下必传 上传图片尺寸670*335，支持格式：png、jpg、jpeg、bmp，大小不超过2MB	限制:	1.指定可核销商品(available_goods)时必传
	VoucherImage *string `json:"voucher_image,omitempty"`
}

// PaymentVoucherRechargeInfo 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherRechargeInfo struct {
	AlipayBalanceRechargeInfo *PaymentVoucherAlipayBalanceRechargeInfo `json:"alipay_balance_recharge_info,omitempty"`
	// RechargeType 是预充值方式
	RechargeType *string `json:"recharge_type,omitempty"`
}

// PaymentVoucherSendRule 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherSendRule struct {
	// MaxQuantityByDay 是设置此字段，允许指定单天最大发券数量。		限制:	每天发放张数*活动天数应小于等于优惠券发放总量
	MaxQuantityByDay *int64 `json:"max_quantity_by_day,omitempty"`
	// NaturalPersonLimit 是限制相同身份证号领取次数(voucher_quantity_limit_per_user)。默认false不限制。	枚举值	true：是	false：否
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// PhoneNumberLimit 是限制相同手机号领取次数(voucher_quantity_limit_per_user)。默认false不限制	枚举值	true：是	false：否
	PhoneNumberLimit *bool `json:"phone_number_limit,omitempty"`
	// RealNameLimit 是限制支付宝实名用户才能领取支付券,默认为false表示不限制	枚举值	true\false
	RealNameLimit *bool `json:"real_name_limit,omitempty"`
	// VoucherQuantity 是发行券的总数量。	限制：	1、发放总个数最少1个	2、发放总个数最多99999999个
	VoucherQuantity *int64 `json:"voucher_quantity,omitempty"`
	// VoucherQuantityLimitPerUser 是每人领取限制。 默认按照支付宝账号进行领取限制; 	不填写或填入0表示没有领取限制.
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
	// VoucherQuantityLimitPerUserPeriodType 是周期限领配置,限制每人在固定周期内领取张数(voucher_quantity_limit_per_user),默认LIFE_CYCLE
	VoucherQuantityLimitPerUserPeriodType *string `json:"voucher_quantity_limit_per_user_period_type,omitempty"`
}

// PaymentVoucherUseRule 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherUseRule struct {
	// AvailableAppIds 是可核销的支付宝小程序id		限制： 	1、必须是支付宝小程序	2、如果包含重复的小程序id会自动进行去重操作。
	AvailableAppIds   *[]string                        `json:"available_app_ids,omitempty"`
	AvailableGoods    *PaymentVoucherAvailableGoods    `json:"available_goods,omitempty"`
	AvailableMerchant *PaymentVoucherAvailableMerchant `json:"available_merchant,omitempty"`
	// AvailableStoreIds 是可核销支付门店id。	限制： 	1、available_store_ids中的门店id必须是支付宝门店id。 	2、available_store_ids如果包含重复的门店id会自动进行去重操作。
	AvailableStoreIds *[]string          `json:"available_store_ids,omitempty"`
	FixVoucher        *PaymentFixVoucher `json:"fix_voucher,omitempty"`
	// UnavailableGoodsIds 是不可核销商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，不可以使用优惠券。		限制:	1、available_goods与unavailable_goods_ids只能二选一或者同时为空
	UnavailableGoodsIds *[]string `json:"unavailable_goods_ids,omitempty"`
	// UseMode 是核销跳转方式,默认CAN_USE(详情页展示可使用按钮)
	UseMode *string `json:"use_mode,omitempty"`
	// UseUrl 是核销跳转链接	限制:	1、只有use_mode为SELF_LINK时该值有效且必传。	2、核销跳转链接必须是支付宝小程序链接。
	UseUrl *string `json:"use_url,omitempty"`
	// VoucherQuantityLimitPerUser 是每个支付宝账号的核销次数限制。 默认不限制
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
	// VoucherQuantityLimitPerUserPeriodType 是核销次数限制周期配置。默认值LIFE_CYCLE表示限制每个支付宝账号在整个活动期间核销次数。
	VoucherQuantityLimitPerUserPeriodType *string                    `json:"voucher_quantity_limit_per_user_period_type,omitempty"`
	VoucherValidPeriod                    *PaymentVoucherValidPeriod `json:"voucher_valid_period,omitempty"`
}

// PaymentVoucherUseRuleDetail 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherUseRuleDetail struct {
	FixVoucher *PaymentFixVoucher `json:"fix_voucher,omitempty"`
	// UseMode 是核销跳转方式,默认CAN_USE(详情页展示可使用按钮)
	UseMode *string `json:"use_mode,omitempty"`
	// UseUrl 是核销跳转链接 限制: 1、只有use_mode为SELF_LINK时该值有效且必传。 2、核销跳转链接必须是支付宝小程序链接。
	UseUrl *string `json:"use_url,omitempty"`
	// VoucherQuantityLimitPerUser 是每个支付宝账号的核销次数限制。 默认不限制
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
	// VoucherQuantityLimitPerUserPeriodType 是核销次数限制周期配置。默认值LIFE_CYCLE表示限制每个支付宝账号在整个活动期间核销次数。
	VoucherQuantityLimitPerUserPeriodType *string                    `json:"voucher_quantity_limit_per_user_period_type,omitempty"`
	VoucherValidPeriod                    *PaymentVoucherValidPeriod `json:"voucher_valid_period,omitempty"`
}

// PaymentVoucherUseRuleModify 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherUseRuleModify struct {
	// AvailableAppIds 是可核销的支付宝小程序id 		限制： 	1、必须是支付宝小程序 	2、如果包含重复的小程序id会自动进行去重操作。	3、可核销小程序范围只能扩大不能缩小
	AvailableAppIds   *string                                `json:"available_app_ids,omitempty"`
	AvailableGoods    *PaymentVoucherAvailableGoodsModify    `json:"available_goods,omitempty"`
	AvailableMerchant *PaymentVoucherAvailableMerchantModify `json:"available_merchant,omitempty"`
	// AvailableStoreIds 是可核销支付门店id。 		限制： 	1、available_store_ids中的门店id必须是支付宝门店id。 	2、available_store_ids如果包含重复的门店id会自动进行去重操作。	3、可核销门店范围只能扩大不能缩小
	AvailableStoreIds  *string                          `json:"available_store_ids,omitempty"`
	VoucherValidPeriod *PaymentVoucherValidPeriodModify `json:"voucher_valid_period,omitempty"`
}

// PaymentVoucherValidPeriod 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherValidPeriod struct {
	// Type 是券有效期类型
	Type *string `json:"type,omitempty"`
	// ValidBeginTime 是券可使用的开始时间。		格式为：yyyy-MM-dd HH:mm:ss		限制：	type为ABSOLUTE时该字段必填。
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// ValidDaysAfterReceive 是券生效后N天内可以使用。		可以配合wait_days_after_receive字段使用。		限制：	1.type为RELATIVE时必填。	2.valid_days_after_receive必须大于0。
	ValidDaysAfterReceive *int64 `json:"valid_days_after_receive,omitempty"`
	// ValidEndTime 是券可使用的结束时间。		格式为yyyy-MM-dd HH:mm:ss		限制：	1、type为ABSOLUTE必填。	2、券可使用的结束时间valid_end_time必须大于券的发放结束时间publish_end_time
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// WaitDaysAfterReceive 是用户领券后需要等待N天，券才可以生效。默认用户领券后立刻生效。		限制：	1、type为RELATIVE时可选。
	WaitDaysAfterReceive *int64 `json:"wait_days_after_receive,omitempty"`
}

// PaymentVoucherValidPeriodModify 是支付宝官方 V3 OpenAPI 模型。
type PaymentVoucherValidPeriodModify struct {
	// ValidDaysAfterReceive 是券生效后N天内可以使用。 限制： type为RELATIVE时可修改。 valid_days_after_receive必须大于0。  修改该时间，只允许延长，不允许缩短。
	ValidDaysAfterReceive *int64 `json:"valid_days_after_receive,omitempty"`
	// ValidEndTime 是券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss 限制： type为ABSOLUTE可修改。 券可使用的结束时间valid_end_time 必须大于 券的发放结束时间 publish_end_time 修改券可使用的结束时间，只能延长，不允许缩短。
	ValidEndTime *string `json:"valid_end_time,omitempty"`
}

// SpecialVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type SpecialVoucherInfo struct {
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// SpecialAmount 是特价，即：原价-特价=优惠金额。
	SpecialAmount              *string                     `json:"special_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// TimeRangeInfo 是支付宝官方 V3 OpenAPI 模型。
type TimeRangeInfo struct {
	// BeginTime 是开始时间  格式：HH:mm:ss
	BeginTime   *string      `json:"begin_time,omitempty"`
	EndTimeInfo *EndTimeInfo `json:"end_time_info,omitempty"`
}

// TimeRestrictInfo 是支付宝官方 V3 OpenAPI 模型。
type TimeRestrictInfo struct {
	// DisablePeriodInfo 是不可用时间段
	DisablePeriodInfo *[]VoucherUseTimeRuleInfo `json:"disable_period_info,omitempty"`
	// UsablePeriodInfo 是可用时间段
	UsablePeriodInfo *[]VoucherUseTimeRuleInfo `json:"usable_period_info,omitempty"`
}

// VoucherAbsolutePeriodInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAbsolutePeriodInfo struct {
	TimeRestrictInfo *TimeRestrictInfo `json:"time_restrict_info,omitempty"`
	// ValidBeginTime 是券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss。
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// ValidEndTime 是券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss。
	ValidEndTime *string `json:"valid_end_time,omitempty"`
}

// VoucherAvailableAccountInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableAccountInfo struct {
	// AvailablePids 是优惠券可核销的直连商户PID
	AvailablePids *[]string `json:"available_pids,omitempty"`
	// AvailableSmids 是优惠券可核销的间连商户SMID
	AvailableSmids *[]string `json:"available_smids,omitempty"`
}

// VoucherAvailableAppInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableAppInfo struct {
	// AvailableAppIds 是可核销的支付宝小程序id
	AvailableAppIds *[]string `json:"available_app_ids,omitempty"`
}

// VoucherAvailableGeographyAllShopInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyAllShopInfo struct {
	// AvailableBrandId 是全部门店可用的品牌ID	提示:	1.目前只支持传入当前券归属商户的品牌ID	2.品牌ID可以在登录b.alipay.com-账户中心-品牌管理里查看。如果没有品牌，可以先创建品牌。Sop：	注意事项	1.只支持传入单个品牌ID	2.传入品牌ID后，券的商户品牌和logo仍以voucher_display_pattern_info传入为准。
	AvailableBrandId *string `json:"available_brand_id,omitempty"`
	// ExcludeShopIds 是券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。 列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。
	ExcludeShopIds *[]string `json:"exclude_shop_ids,omitempty"`
	// MerchantIds 是全部门店可用的商户ID
	MerchantIds *[]string `json:"merchant_ids,omitempty"`
}

// VoucherAvailableGeographyCityInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyCityInfo struct {
	// AllCity 是是否全国。 选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。
	AllCity *bool `json:"all_city,omitempty"`
	// AvailableCityCodes 是城市编码。请按照表格中内容填写。（参考资料）
	AvailableCityCodes *[]string `json:"available_city_codes,omitempty"`
}

// VoucherAvailableGeographyScopeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyScopeInfo struct {
	AvailableGeographyCityInfo *VoucherAvailableGeographyCityInfo `json:"available_geography_city_info,omitempty"`
	// AvailableGeographyScopeType 是券可用地理位置类型。
	AvailableGeographyScopeType *string                            `json:"available_geography_scope_type,omitempty"`
	AvailableGeographyShopInfo  *VoucherAvailableGeographyShopInfo `json:"available_geography_shop_info,omitempty"`
}

// VoucherAvailableGeographyShopInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyShopInfo struct {
	AvailableGeographyAllShop *VoucherAvailableGeographyAllShopInfo `json:"available_geography_all_shop,omitempty"`
	// AvailableRealShopIds 是券可使用的物理门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。
	AvailableRealShopIds *[]string `json:"available_real_shop_ids,omitempty"`
	// AvailableShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。
	AvailableShopIds *[]string `json:"available_shop_ids,omitempty"`
}

// VoucherAvailableGoodsInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGoodsInfo struct {
	// AvailableGoodsSkuIds 是可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。
	AvailableGoodsSkuIds *[]string `json:"available_goods_sku_ids,omitempty"`
	// ExcludeGoodsSkuIds 是不可核销商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，不可以使用优惠券。
	ExcludeGoodsSkuIds *[]string `json:"exclude_goods_sku_ids,omitempty"`
	// GoodsDescription 是商品描述信息。 用于券面展示，向用户介绍商品
	GoodsDescription *string `json:"goods_description,omitempty"`
	// GoodsDetail 是商品详情，会展示在特定渠道(如服务市场团购业务插件的套餐明细)。
	GoodsDetail *string `json:"goods_detail,omitempty"`
	// GoodsDetailImages 是商品详情图片列表，会展示在特定渠道(如服务市场团购业务插件的头图)
	GoodsDetailImages *[]string `json:"goods_detail_images,omitempty"`
	// GoodsDetailRichDescription 是商品详情富文本描述
	GoodsDetailRichDescription *string `json:"goods_detail_rich_description,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// OriginAmount 是原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
}

// VoucherAvailableItemInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableItemInfo struct {
	// ItemId 是内部商品ID
	ItemId *string `json:"item_id,omitempty"`
	// OutItemInfo 是外部商品信息
	OutItemInfo *[]VoucherAvailableOutItemInfo `json:"out_item_info,omitempty"`
}

// VoucherAvailableOutItemInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableOutItemInfo struct {
	// ItemAppId 是小程序ID
	ItemAppId *string `json:"item_app_id,omitempty"`
	// OutItemId 是外部商品ID
	OutItemId *string `json:"out_item_id,omitempty"`
}

// VoucherAvailableScopeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableScopeInfo struct {
	VoucherAvailableAccountInfo        *VoucherAvailableAccountInfo        `json:"voucher_available_account_info,omitempty"`
	VoucherAvailableAppInfo            *VoucherAvailableAppInfo            `json:"voucher_available_app_info,omitempty"`
	VoucherAvailableGeographyScopeInfo *VoucherAvailableGeographyScopeInfo `json:"voucher_available_geography_scope_info,omitempty"`
	VoucherAvailableGoodsInfo          *VoucherAvailableGoodsInfo          `json:"voucher_available_goods_info,omitempty"`
	VoucherAvailableItemInfo           *VoucherAvailableItemInfo           `json:"voucher_available_item_info,omitempty"`
}

// VoucherBalanceRechargeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherBalanceRechargeInfo struct {
	// Amount 是支付宝余额充值金额	 限制： 	 1.币种为人民币，单位元。 	 2. 总预算=优惠金额*总发券张数
	Amount *string `json:"amount,omitempty"`
	// LogonId 是出资的商户支付宝登录账号
	LogonId *string `json:"logon_id,omitempty"`
	// PartnerId 是出资的商户支付宝ID
	PartnerId *string `json:"partner_id,omitempty"`
}

// VoucherBudgetSupplyInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherBudgetSupplyInfo struct {
	// BudgetType 是营销资金预算类型 	1、免充值：制券方无需提前充值资金，用户核销支付券时，直接从订单原价中扣除优惠减价金额，最终只将用户实际支付的金额结算给核销商户，商户实收少于订单原价。  	2.预充值:  制券方需将优惠预算提前充值到支付宝指定营销账户中，用户核销支付券时，系统从该账户可用余额中扣除优惠减价部分对应的资金，连同用户实际支付的资金，一并结算给核销商户，不影响实收。
	BudgetType          *string              `json:"budget_type,omitempty"`
	VoucherRechargeInfo *VoucherRechargeInfo `json:"voucher_recharge_info,omitempty"`
}

// VoucherCustomerGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherCustomerGuideInfo struct {
	VoucherSendGuideInfo *VoucherSendGuideInfo `json:"voucher_send_guide_info,omitempty"`
	VoucherUseGuideInfo  *VoucherUseGuideInfo  `json:"voucher_use_guide_info,omitempty"`
}
