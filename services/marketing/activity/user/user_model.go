package user

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

// ActivityDiscountVoucher 是支付宝官方 V3 OpenAPI 模型。
type ActivityDiscountVoucher struct {
	// CeilingAmount 是封顶金额。
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// Discount 是折扣率。
	Discount *string `json:"discount,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。说明：该字段可不填，填入 origin_amount 则必填。
	GoodsName *string `json:"goods_name,omitempty"`
	// OriginAmount 是原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
}

// ActivityExchangeVoucher 是支付宝官方 V3 OpenAPI 模型。
type ActivityExchangeVoucher struct {
	// Amount 是券的价值。
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// OverdueRefundable 是是否支持优惠券过期后，自动退款给用户。 不填默认否，枚举值： true：是 false：否  自动退款功能需要服务商在优惠券过期时，主动调用alipay.marketing.activity.order.refund接口进行退款。  如果配置优惠券时选择了过期自动退款，但是实际券过期后，服务商没有进行退款，那么用户投诉后，需要服务商进行解决。
	OverdueRefundable *bool `json:"overdue_refundable,omitempty"`
	// Refundable 是购买的优惠券是否允许退款。 不填默认否，枚举值： true：是 false：否
	Refundable *bool `json:"refundable,omitempty"`
	// SaleAmount 是用户购买优惠券需要支付的金额。
	SaleAmount *string `json:"sale_amount,omitempty"`
}

// ActivityFixVoucher 是支付宝官方 V3 OpenAPI 模型。
type ActivityFixVoucher struct {
	// Amount 是面额，每张代金券可以抵扣的金额。
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。该字段不填写，认为无门槛。
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
}

// ActivitySpecialVoucher 是支付宝官方 V3 OpenAPI 模型。
type ActivitySpecialVoucher struct {
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。说明：该字段可不填，填入 origin_amount 则必填。
	GoodsName *string `json:"goods_name,omitempty"`
	// OriginAmount 是原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
	// SpecialAmount 是特价，即：原价-特价=优惠金额。
	SpecialAmount *string `json:"special_amount,omitempty"`
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

// CommonVoucherDisplayInfo 是支付宝官方 V3 OpenAPI 模型。
type CommonVoucherDisplayInfo struct {
	// BrandLogo 是商家品牌 logo 链接
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// VoucherDescription 是用于说明详细的活动规则，会展示在支付宝卡包券详情页。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// VoucherDetailImages 是券详情图片链接。
	VoucherDetailImages *[]string `json:"voucher_detail_images,omitempty"`
	// VoucherImage 是券封面图链接。
	VoucherImage *string `json:"voucher_image,omitempty"`
}

// CommonVoucherDisplayLiteInfo 是支付宝官方 V3 OpenAPI 模型。
type CommonVoucherDisplayLiteInfo struct {
	// BrandLogo 是商户品牌 logo 链接。
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是商户品牌名称。 如果不设置品牌名称，则默认使用支付宝商家系统中维护的商家别名。 如果没有维护商家别名，则使用默认名称：商家优惠。  用户领取优惠券后，品牌名称会对用户进行展示。
	BrandName *string `json:"brand_name,omitempty"`
}

// CommonVoucherSendRule 是支付宝官方 V3 OpenAPI 模型。
type CommonVoucherSendRule struct {
	// MaxQuantityByDay 是设置此字段，允许指定单天最大发券数量。
	MaxQuantityByDay *int64 `json:"max_quantity_by_day,omitempty"`
	// NaturalPersonLimit 是限制相同身份证号领取次数(voucher_quantity_limit_per_user)。默认false不限制。 枚举值 true：是  false：否
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// PhoneNumberLimit 是限制相同手机号领取次数(voucher_quantity_limit_per_user)。默认false不限制 枚举值 true：是 false：否
	PhoneNumberLimit *bool `json:"phone_number_limit,omitempty"`
	// RealNameLimit 是限制支付宝实名用户才能领取支付券，默认为false表示不限制。枚举值 true\false
	RealNameLimit *bool `json:"real_name_limit,omitempty"`
	// VoucherQuantity 是发行券的总数量。
	VoucherQuantity *int64 `json:"voucher_quantity,omitempty"`
	// VoucherQuantityLimitPerUser 是每人领取限制。 默认按照支付宝账号进行领取限制; 不填写或填入0表示没有领取限制。
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
	// VoucherQuantityLimitPerUserPeriodType 是周期限领配置,限制每人在固定周期内领取张数（voucher_quantity_limit_per_user），默认LIFE_CYCLE。
	VoucherQuantityLimitPerUserPeriodType *string `json:"voucher_quantity_limit_per_user_period_type,omitempty"`
}

// CommonVoucherUseRule 是支付宝官方 V3 OpenAPI 模型。
type CommonVoucherUseRule struct {
	DiscountVoucher *ActivityDiscountVoucher `json:"discount_voucher,omitempty"`
	ExchangeVoucher *ActivityExchangeVoucher `json:"exchange_voucher,omitempty"`
	FixVoucher      *ActivityFixVoucher      `json:"fix_voucher,omitempty"`
	SpecialVoucher  *ActivitySpecialVoucher  `json:"special_voucher,omitempty"`
}

// CommonVoucherUseRuleLiteInfo 是支付宝官方 V3 OpenAPI 模型。
type CommonVoucherUseRuleLiteInfo struct {
	DiscountVoucher *ActivityDiscountVoucher `json:"discount_voucher,omitempty"`
	ExchangeVoucher *ActivityExchangeVoucher `json:"exchange_voucher,omitempty"`
	FixVoucher      *ActivityFixVoucher      `json:"fix_voucher,omitempty"`
	SpecialVoucher  *ActivitySpecialVoucher  `json:"special_voucher,omitempty"`
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

// UserVoucherBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type UserVoucherBaseInfo struct {
	// AssociateTradeNo 是若商家券操作过关联商户订单信息，则该字段返回商家券已关联的商户订单号。
	AssociateTradeNo *string `json:"associate_trade_no,omitempty"`
	// BelongMerchantId 是券归属商户
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// CreateTime 是领券时间
	CreateTime *string `json:"create_time,omitempty"`
	// ValidBeginTime 是券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// ValidEndTime 是券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// VoucherCode 是用户领取的券码code,支付宝商家券活动才会返回券码，其他优惠券活动该值为空
	VoucherCode *string `json:"voucher_code,omitempty"`
	// VoucherId 是用户券 id。支付宝为用户优惠券唯一分配的 id。
	VoucherId *string `json:"voucher_id,omitempty"`
	// VoucherMaxUseTimes 是券最大核销次数
	VoucherMaxUseTimes *int64 `json:"voucher_max_use_times,omitempty"`
	// VoucherName 是对消费者展示的券(商品)名称。
	VoucherName *string `json:"voucher_name,omitempty"`
	// VoucherStatus 是券状态。SENDED:可用 USED:已核销(即完全使用), EXPIRED:已过期(只能查询6个月内数据)
	VoucherStatus *string `json:"voucher_status,omitempty"`
	// VoucherUsedTimes 是券已核销次数
	VoucherUsedTimes *int64 `json:"voucher_used_times,omitempty"`
}

// UserVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type UserVoucherInfo struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// AvailableBeginTime 是券可用开始时间。
	AvailableBeginTime *string `json:"available_begin_time,omitempty"`
	// AvailableEndTime 是券可用结束时间。
	AvailableEndTime *string `json:"available_end_time,omitempty"`
	// BelongMerchantId 是归属商户PID
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// CreateTime 是用户领券时间。
	CreateTime                *string                       `json:"create_time,omitempty"`
	UserVoucherBaseInfo       *UserVoucherBaseInfo          `json:"user_voucher_base_info,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo    `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo     `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo            `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayLiteInfo    *CommonVoucherDisplayLiteInfo `json:"voucher_display_lite_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo    `json:"voucher_display_pattern_info,omitempty"`
	// VoucherId 是用户券 id。支付宝为用户优惠券唯一分配的 id。
	VoucherId *string `json:"voucher_id,omitempty"`
	// VoucherName 是对消费者展示的券(商品)名称。
	VoucherName         *string              `json:"voucher_name,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo `json:"voucher_send_mode_info,omitempty"`
	// VoucherStatus 是券状态。
	VoucherStatus *string `json:"voucher_status,omitempty"`
	// VoucherType 是券类型
	VoucherType            *string                       `json:"voucher_type,omitempty"`
	VoucherUseRuleInfo     *VoucherUseRuleInfo           `json:"voucher_use_rule_info,omitempty"`
	VoucherUseRuleLiteInfo *CommonVoucherUseRuleLiteInfo `json:"voucher_use_rule_lite_info,omitempty"`
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

// VoucherCustomerGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherCustomerGuideInfo struct {
	VoucherSendGuideInfo *VoucherSendGuideInfo `json:"voucher_send_guide_info,omitempty"`
	VoucherUseGuideInfo  *VoucherUseGuideInfo  `json:"voucher_use_guide_info,omitempty"`
}

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
