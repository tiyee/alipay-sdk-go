package voucherpackage

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

// VoucherPackageActivityInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageActivityInfo struct {
	ActivityBaseInfo          *ActivityBaseInfo          `json:"activity_base_info,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo        `json:"voucher_use_rule_info,omitempty"`
}

// VoucherPackageBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageBaseInfo struct {
	// PurchaseEndTime 是券包购买结束时间
	PurchaseEndTime *string `json:"purchase_end_time,omitempty"`
	// PurchaseStartTime 是券包购买开始时间
	PurchaseStartTime *string `json:"purchase_start_time,omitempty"`
	// VoucherPackageId 是券包id
	VoucherPackageId *string `json:"voucher_package_id,omitempty"`
	// VoucherPackageName 是券包名称
	VoucherPackageName *string `json:"voucher_package_name,omitempty"`
	// VoucherPackageStatus 是券包状态
	VoucherPackageStatus *string `json:"voucher_package_status,omitempty"`
	// VoucherTotalAmount 是券包总价值（单位是元）： 券包下券的优惠总额（满减券指优惠面额，折扣券指优惠上限，特价券指原价减去特价）
	VoucherTotalAmount *string `json:"voucher_total_amount,omitempty"`
}

// VoucherPackageConsultResult 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageConsultResult struct {
	// ConsultResultCode 是券包购买咨询结果code：	SUCCESS：可以领取；	PRODUCT_NOT_ENABLE：产品不可用；	NOT_IN_SELL_TIME：不在销售时间内；	PRODUCT_SOLD_OUT：产品售罄；	USER_PURCHASE_LIMIT：用户购买上限；	OTHER：不可购买，未知原因
	ConsultResultCode *string `json:"consult_result_code,omitempty"`
	// VoucherPackageId 是券包id
	VoucherPackageId *string `json:"voucher_package_id,omitempty"`
}

// VoucherPackageInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageInfo struct {
	VoucherPackageBaseInfo      *VoucherPackageBaseInfo      `json:"voucher_package_base_info,omitempty"`
	VoucherPackageSalesLiteInfo *VoucherPackageSalesLiteInfo `json:"voucher_package_sales_lite_info,omitempty"`
}

// VoucherPackageModeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageModeInfo struct {
	// VoucherPackageId 是券包id，对应alipay.marketing.activity.voucherpackage.query中voucher_package_id。
	VoucherPackageId *string `json:"voucher_package_id,omitempty"`
}

// VoucherPackageSalesInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageSalesInfo struct {
	// Budget 是券包售卖预算，单位是份数
	Budget *int64 `json:"budget,omitempty"`
	// PayChannel 是券包购买支付渠道	pcredit：花呗	creditCard：信用卡	credit_group：花呗与信用卡	为空则不限渠道
	PayChannel *string `json:"pay_channel,omitempty"`
	// PurchaseUrl 是券包购买链接
	PurchaseUrl *string `json:"purchase_url,omitempty"`
	// SaleCountLimitInPeriod 是券包售卖期限内最大购买次数
	SaleCountLimitInPeriod *int64 `json:"sale_count_limit_in_period,omitempty"`
	// SalePeriodLimit 是券包购买期限类型	NO：不限制 	ALL：售卖时间内 	DAY：天 	WEEK：周 	MONTH：月
	SalePeriodLimit *string `json:"sale_period_limit,omitempty"`
	// SalePrice 是券包售卖价格，单位是元
	SalePrice *string `json:"sale_price,omitempty"`
}

// VoucherPackageSalesLiteInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageSalesLiteInfo struct {
	// Budget 是券包预算，单位是份数
	Budget *int64 `json:"budget,omitempty"`
	// SalePrice 是券包售卖价格，单位是元
	SalePrice *string `json:"sale_price,omitempty"`
}

// VoucherPackageUseRule 是支付宝官方 V3 OpenAPI 模型。
type VoucherPackageUseRule struct {
	// UseRuleDesc 是券包使用规则
	UseRuleDesc *string `json:"use_rule_desc,omitempty"`
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
