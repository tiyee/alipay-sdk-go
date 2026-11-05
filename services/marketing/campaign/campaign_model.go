// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package campaign

// AssetChannelInfo 是支付宝官方 V3 OpenAPI 模型。
type AssetChannelInfo struct {
	// AssetTypeCode 是资产类型编码
	AssetTypeCode *string `json:"asset_type_code,omitempty"`
	// CardType 是卡类型：DC借记卡，CC贷记卡
	CardType *string `json:"card_type,omitempty"`
	// Channel 是银行卡所属的银行渠道,格式为:机构ID^渠道码,例如:ICBC^BIG_AMOUNT_CREDIT_CARTOON，花呗 INST_ALIPAY^PCREDIT_PAY)
	Channel *string `json:"channel,omitempty"`
	// ChannelCode 是渠道编码
	ChannelCode *string `json:"channel_code,omitempty"`
	// ChannelTypeCode 是渠道标识
	ChannelTypeCode *string `json:"channel_type_code,omitempty"`
	// InstId 是中国农业银行
	InstId *string `json:"inst_id,omitempty"`
}

// CashCampaignInfo 是支付宝官方 V3 OpenAPI 模型。
type CashCampaignInfo struct {
	// CampStatus 是活动状态
	CampStatus *string `json:"camp_status,omitempty"`
	// CouponName 是现金红包名称
	CouponName *string `json:"coupon_name,omitempty"`
	// CrowdNo 是现金红包活动号
	CrowdNo *string `json:"crowd_no,omitempty"`
	// OriginCrowdNo 是原始活动号,商户进行问题排查时提供
	OriginCrowdNo *string `json:"origin_crowd_no,omitempty"`
}

// ItemConsultInfo 是支付宝官方 V3 OpenAPI 模型。
type ItemConsultInfo struct {
	// ItemId 是商品id
	ItemId *string `json:"item_id,omitempty"`
	// PromoAmount 是商品优惠金额，单位为元
	PromoAmount *string `json:"promo_amount,omitempty"`
	// PromoCount 是商品优惠数量
	PromoCount *string `json:"promo_count,omitempty"`
}

// ItemPromoInfo 是支付宝官方 V3 OpenAPI 模型。
type ItemPromoInfo struct {
	// ItemConsultList 是商品咨询列表
	ItemConsultList *[]ItemConsultInfo `json:"item_consult_list,omitempty"`
	// ItemCoverPic 是单品封面图
	ItemCoverPic *string `json:"item_cover_pic,omitempty"`
	// ItemDesc 是单品描述
	ItemDesc *string `json:"item_desc,omitempty"`
	// ItemDetailPicList 是单品详情图列表
	ItemDetailPicList *[]string `json:"item_detail_pic_list,omitempty"`
	// ItemName 是单品名称
	ItemName *string `json:"item_name,omitempty"`
}

// QuotaInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type QuotaInfoDTO struct {
	// LimitMoney 是总限额
	LimitMoney *string `json:"limit_money,omitempty"`
	// LimitTimes 是总限次
	LimitTimes *string `json:"limit_times,omitempty"`
	// RemainMoney 是剩余额度
	RemainMoney *string `json:"remain_money,omitempty"`
	// RemainTimes 是剩余次数
	RemainTimes *string `json:"remain_times,omitempty"`
	// TradeNo 是交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// VoucherConsultInfo 是支付宝官方 V3 OpenAPI 模型。
type VoucherConsultInfo struct {
	// ActiveTime 是开始时间，立减为活动开始时间，券为券生效时间
	ActiveTime *string `json:"active_time,omitempty"`
	// AssetType 是资产优惠类型，用于区分不同优惠
	AssetType *string `json:"asset_type,omitempty"`
	// AssetsCode 是资产编码
	AssetsCode *string `json:"assets_code,omitempty"`
	// CeilingAmount 是封顶优惠金额，单位为元
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// ExpiredTime 是结束时间，立减为活动结束时间，券为券失效时间
	ExpiredTime   *string        `json:"expired_time,omitempty"`
	ItemPromoInfo *ItemPromoInfo `json:"item_promo_info,omitempty"`
	// Optimal 是是否为可使用的互斥券中最优的（目前返回的记录该值均为true，即可忽略该值判断）
	Optimal *bool `json:"optimal,omitempty"`
	// PreConsultId 是预咨询ID
	PreConsultId *string `json:"pre_consult_id,omitempty"`
	// PromoAmount 是优惠总金额，单位为元
	PromoAmount *string `json:"promo_amount,omitempty"`
	// PromoText 是优惠文案
	PromoText *string `json:"promo_text,omitempty"`
	// PromoType 是优惠类型： 代金(FIX),折扣(DISCOUNT),减至(SPECIAL，即特价券)
	PromoType *string `json:"promo_type,omitempty"`
	// QuotaInfoList 是额度信息
	QuotaInfoList *[]QuotaInfoDTO `json:"quota_info_list,omitempty"`
	// ReductionAmount 是满减金额，单位为元，仅针对代金券
	ReductionAmount *string `json:"reduction_amount,omitempty"`
	// ReductionRatio 是优惠额度，小数表示，0.7 表示 30% 的折扣
	ReductionRatio *string `json:"reduction_ratio,omitempty"`
	// SpecifiedAmount 是减至金额，单位为元，仅针对减至券
	SpecifiedAmount *string `json:"specified_amount,omitempty"`
	// ThresholdAmount 是优惠门槛，单位为元（部分券没有门槛金额）
	ThresholdAmount *string        `json:"threshold_amount,omitempty"`
	VoucherDetail   *VoucherDetail `json:"voucher_detail,omitempty"`
	// VoucherId 是券id
	VoucherId *string `json:"voucher_id,omitempty"`
	// VoucherName 是券名称
	VoucherName *string `json:"voucher_name,omitempty"`
	// VoucherType 是券类型： 1）全场券(ALL) 2）单品券(ITEM)
	VoucherType *string `json:"voucher_type,omitempty"`
}

// VoucherDetail 是支付宝官方 V3 OpenAPI 模型。
type VoucherDetail struct {
	// CanUseAssetChannelInfos 是可用渠道信息，可用于指定渠道优先展示
	CanUseAssetChannelInfos *[]AssetChannelInfo `json:"can_use_asset_channel_infos,omitempty"`
}
