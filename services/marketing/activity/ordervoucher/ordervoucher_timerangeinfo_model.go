package ordervoucher

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

// VoucherAvailableGeographyAllShopResultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyAllShopResultInfo struct {
	// FailExcludeShopInfos 是请求失败的不可用门店详情。
	FailExcludeShopInfos *[]VoucherFailShopInfo `json:"fail_exclude_shop_infos,omitempty"`
	// SuccessExcludeShopIds 是请求成功的不可用门店列表。
	SuccessExcludeShopIds *[]string `json:"success_exclude_shop_ids,omitempty"`
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

// VoucherAvailableGeographyScopeResultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyScopeResultInfo struct {
	AvailableGeographyShopResultInfo *VoucherAvailableGeographyShopResultInfo `json:"available_geography_shop_result_info,omitempty"`
}

// VoucherAvailableGeographyShopInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyShopInfo struct {
	AvailableGeographyAllShop *VoucherAvailableGeographyAllShopInfo `json:"available_geography_all_shop,omitempty"`
	// AvailableRealShopIds 是券可使用的物理门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。
	AvailableRealShopIds *[]string `json:"available_real_shop_ids,omitempty"`
	// AvailableShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。
	AvailableShopIds *[]string `json:"available_shop_ids,omitempty"`
}

// VoucherAvailableGeographyShopResultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableGeographyShopResultInfo struct {
	AvailableGeographyAllShopResultInfo *VoucherAvailableGeographyAllShopResultInfo `json:"available_geography_all_shop_result_info,omitempty"`
	// FailAvailableRealShopInfos 是失败的物理门店详情。
	FailAvailableRealShopInfos *[]VoucherFailShopInfo `json:"fail_available_real_shop_infos,omitempty"`
	// FailAvailableShopInfos 是失败的可用门店。
	FailAvailableShopInfos *[]VoucherFailShopInfo `json:"fail_available_shop_infos,omitempty"`
	// SuccessAvailableRealShopIds 是代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id  接口参数是列表类型。
	SuccessAvailableRealShopIds *[]string `json:"success_available_real_shop_ids,omitempty"`
	// SuccessAvailableShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	SuccessAvailableShopIds *[]string `json:"success_available_shop_ids,omitempty"`
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

// VoucherAvailableScope 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableScope struct {
	OrderVoucherAvailableCityCode *OrderVoucherAvailableCityCode `json:"order_voucher_available_city_code,omitempty"`
	OrderVoucherAvailableShop     *OrderVoucherAvailableShop     `json:"order_voucher_available_shop,omitempty"`
	// VoucherAvailableType 是可用范围类型。
	VoucherAvailableType *string `json:"voucher_available_type,omitempty"`
}

// VoucherAvailableScopeInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableScopeInfo struct {
	VoucherAvailableAccountInfo        *VoucherAvailableAccountInfo        `json:"voucher_available_account_info,omitempty"`
	VoucherAvailableAppInfo            *VoucherAvailableAppInfo            `json:"voucher_available_app_info,omitempty"`
	VoucherAvailableGeographyScopeInfo *VoucherAvailableGeographyScopeInfo `json:"voucher_available_geography_scope_info,omitempty"`
	VoucherAvailableGoodsInfo          *VoucherAvailableGoodsInfo          `json:"voucher_available_goods_info,omitempty"`
	VoucherAvailableItemInfo           *VoucherAvailableItemInfo           `json:"voucher_available_item_info,omitempty"`
}

// VoucherAvailableScopeModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableScopeModify struct {
	// ModifyType 是可用范围修改类型。
	ModifyType                    *string                              `json:"modify_type,omitempty"`
	OrderVoucherAvailableCityCode *OrderVoucherAvailableCityCodeModify `json:"order_voucher_available_city_code,omitempty"`
	OrderVoucherAvailableShop     *OrderVoucherAvailableShopModify     `json:"order_voucher_available_shop,omitempty"`
}

// VoucherAvailableScopeResultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherAvailableScopeResultInfo struct {
	VoucherAvailableGeographyScopeResultInfo *VoucherAvailableGeographyScopeResultInfo `json:"voucher_available_geography_scope_result_info,omitempty"`
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

// VoucherDisplayInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherDisplayInfo struct {
	// BrandLogo 是商家logo。需要通过	alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是商户品牌名称。如果不设置品牌名称，则默认使用支付宝商家系统中维护的商家别名。如果没有维护商家别名，则使用默认名称：商家优惠。用户领取优惠券后，品牌名称会对用户进行展示
	BrandName *string `json:"brand_name,omitempty"`
	// CustomerServiceMobile 是客服电话
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// CustomerServiceUrl 是客服链接
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// VoucherComment 是券补充说明，该字段为可运营字段，可描述商品特色	会展示在用户支付宝卡包券详情页
	VoucherComment *string `json:"voucher_comment,omitempty"`
	// VoucherDescription 是券详细使用说明。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// VoucherDetailImages 是券详细图列表，会展示在用户支付宝卡包券详情页。需要通过	alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE	,接口返回的resource_id即为该参数的值。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。
	VoucherDetailImages *[]string `json:"voucher_detail_images,omitempty"`
	// VoucherImage 是券详情页封面图，会展示在用户支付宝卡包券详情页。需要通过	alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值
	VoucherImage *string `json:"voucher_image,omitempty"`
}

// VoucherDisplayInfoModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherDisplayInfoModify struct {
	// BrandLogo 是商家logo 	需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是商户品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// CustomerServiceMobile 是客服电话
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// CustomerServiceUrl 是客服链接
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// VoucherComment 是券补充说明，该字段为可运营字段，可描述商品特色 会展示在用户支付宝卡包券详情页。
	VoucherComment *string `json:"voucher_comment,omitempty"`
	// VoucherDescription 是券详细使用说明。
	VoucherDescription *string `json:"voucher_description,omitempty"`
	// VoucherDetailImages 是券详细图列表，会展示在用户支付宝卡包券详情页 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE  ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。
	VoucherDetailImages *[]string `json:"voucher_detail_images,omitempty"`
	// VoucherImage 是券详情页封面图，会展示在用户支付宝卡包券详情页 需要通过  alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值。
	VoucherImage *string `json:"voucher_image,omitempty"`
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

// VoucherFailShopInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherFailShopInfo struct {
	// FailMessage 是请求失败提示信息。
	FailMessage *string `json:"fail_message,omitempty"`
	// FailReasons 是请求失败的原因。
	FailReasons *[]string `json:"fail_reasons,omitempty"`
	// RealShopId 是请求失败的物理门店 id。
	RealShopId *string `json:"real_shop_id,omitempty"`
	// ShopId 是支付宝侧蚂蚁店铺 id。
	ShopId *string `json:"shop_id,omitempty"`
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

// VoucherRefundDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherRefundDetailInfo struct {
	// VoucherRefundAmount 是取消核销优惠金额, 单位:元
	VoucherRefundAmount *string `json:"voucher_refund_amount,omitempty"`
	// VoucherRefundUseTimes 是本次取消核销次数
	VoucherRefundUseTimes *int64 `json:"voucher_refund_use_times,omitempty"`
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

// VoucherSendGuide 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendGuide struct {
	// VoucherDetailUrl 是领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
}

// VoucherSendGuideInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendGuideInfo struct {
	MiniAppSendGuideInfo *VoucherMiniAppSendGuideInfo `json:"mini_app_send_guide_info,omitempty"`
	// SendGuideMode 是发放引导模式
	SendGuideMode *[]string `json:"send_guide_mode,omitempty"`
}

// VoucherSendGuideModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendGuideModify struct {
	// VoucherDetailUrl 是领（购）券详情页链接，从支付宝公域跳转到服务商（商户）自定义领(购)券详情页。说明：当 voucher_type=EXCHANGE_VOUCHER 时，该字段可修改，其他不允许修改。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
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

// VoucherSendRuleDetail 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendRuleDetail struct {
	// NaturalPersonLimit 是是否开启自然人领取限制。自然人表示按照身份证纬度进行领取限制。
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// PhoneNumberLimit 是是否开启电话号码领取限制。
	PhoneNumberLimit *bool `json:"phone_number_limit,omitempty"`
	// VoucherQuantity 是发行券的数量。
	VoucherQuantity *int64 `json:"voucher_quantity,omitempty"`
	// VoucherQuantityLimitPerUser 是每人领取限制。默认按照支付宝uid进行领取限制;	不填写或填入0，默认没有领取限制。
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
}

// VoucherSendRuleDetailModify 是支付宝官方 V3 OpenAPI 模型。
type VoucherSendRuleDetailModify struct {
	// NaturalPersonLimit 是是否开启自然人领取限制。 自然人表示按照身份证纬度进行领取限制。
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// PhoneNumberLimit 是是否开启电话号码领取限制。
	PhoneNumberLimit *string `json:"phone_number_limit,omitempty"`
	// VoucherQuantityLimitPerUser 是每人领取限制。 默认按照支付宝uid进行领取限制;
	VoucherQuantityLimitPerUser *int64 `json:"voucher_quantity_limit_per_user,omitempty"`
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

// VoucherUseDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseDetailInfo struct {
	// VoucherUseAmount 是核销优惠金额, 单位:元
	VoucherUseAmount *string `json:"voucher_use_amount,omitempty"`
	// VoucherUseTimes 是本次核销次数（不包含撤销核销的次数）
	VoucherUseTimes *int64 `json:"voucher_use_times,omitempty"`
}

// VoucherUseDetailResultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherUseDetailResultInfo struct {
	// VoucherMaxUnUseTimes 是剩余可核销次数
	VoucherMaxUnUseTimes *int64 `json:"voucher_max_un_use_times,omitempty"`
}
