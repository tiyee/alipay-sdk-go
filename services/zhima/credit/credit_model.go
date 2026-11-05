// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package credit

// AmountTypeData is the official V3 AmountTypeData schema.
type AmountTypeData struct {
	// Name represents name. 累计类型为金额类型的累计名称。为必传入参。
	Name *string `json:"name,omitzero"`
	// OutDiscountInfos represents out_discount_infos. 优惠信息。为List类型，支持传入多笔优惠，大部分场景传入一笔。包含：优惠名称，优惠类型，优惠金额（单位元）。 字段必传场景： 1.当累计模式为【交易绑定模式】时，为必传入参。 2.当累计模式为【交易解耦模式】，同时需要累计优惠信息时，为必传入参。
	OutDiscountInfos *[]ZMGoOutDiscountInfo `json:"out_discount_infos,omitzero"`
	// TradeInfo represents trade_info.
	TradeInfo *ZMGoTradeInfo `json:"trade_info,omitzero"`
}

// TaskTypeData is the official V3 TaskTypeData schema.
type TaskTypeData struct {
	// Name represents name. 商户数据回传的任务名称，供C端展示。当累计的数据类型为任务类型时，此为必传字段
	Name *string `json:"name,omitzero"`
}

// ExtTemplateConf is the official V3 ExtTemplateConf schema.
type ExtTemplateConf struct {
	// BuyerId represents buyer_id. 淘宝uid
	BuyerId *string `json:"buyer_id,omitzero"`
	// XxhmInfoId represents xxhm_info_id. 先享后买id
	XxhmInfoId *string `json:"xxhm_info_id,omitzero"`
}

// PreOrderExtInfo is the official V3 PreOrderExtInfo schema.
type PreOrderExtInfo struct {
	// BrandId represents brand_id. 品牌id
	BrandId *string `json:"brand_id,omitzero"`
	// BrandName represents brand_name. 品牌名称
	BrandName *string `json:"brand_name,omitzero"`
	// Channel represents channel. 请求渠道
	Channel *string `json:"channel,omitzero"`
	// OrderFeature represents order_feature. 单据属性,风控系统使用
	OrderFeature *string `json:"order_feature,omitzero"`
}

// SettleExtraParams is the official V3 SettleExtraParams schema.
type SettleExtraParams struct {
	// QuitType represents quit_type. quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT
	QuitType *string `json:"quit_type,omitzero"`
	// ScenePayLinkInfo represents scene_pay_link_info. 商链通权益抵扣信息, 当前只有运营商使用, 未使用权益抵扣忽略该字段 fundRuleList为对应的出资信息，目前该列表不支持多对象，长度限制为1。其中fundAgreementNo为本笔交易实际出资的协议，由(商链通二方通用版本接入文档)2.3.2获取，fundAmount为该出资协议出资的金额。 bizScene和subBizScene为(商链通二方通用版本接入文档)2.1.3中支付宝侧分配的信息。 bizMode=MERCHANT_ORDER为固定值。
	ScenePayLinkInfo *string `json:"scene_pay_link_info,omitzero"`
	// SettleAdjustReason represents settle_adjust_reason. action_type选择PAY_TO_ZERO时必填， "SERVICE_CANCELED":"服务已取消", "OTHER_CHANNEL_PERFORMANCE":"户已通过其他方式履约"
	SettleAdjustReason *string `json:"settle_adjust_reason,omitzero"`
}

// UnfreezeExtendParams is the official V3 UnfreezeExtendParams schema.
type UnfreezeExtendParams struct {
	// QuitType represents quit_type. 退出类型
	QuitType *string `json:"quit_type,omitzero"`
	// TotalDiscountAmount represents total_discount_amount. 芝麻Go开始到到期过程中享受的优惠总金额
	TotalDiscountAmount *string `json:"total_discount_amount,omitzero"`
	// TotalRealPayAmount represents total_real_pay_amount. 芝麻Go开始到到期过程中任务相关支付的总金额。
	TotalRealPayAmount *string `json:"total_real_pay_amount,omitzero"`
	// TotalTaskCount represents total_task_count. 芝麻Go开始到到期过程中和任务相关的完成次数
	TotalTaskCount *string `json:"total_task_count,omitzero"`
}

// ZMGoOutDiscountInfo is the official V3 ZMGoOutDiscountInfo schema.
type ZMGoOutDiscountInfo struct {
	// DiscountAmount represents discount_amount. 商户回传的优惠金额，如用户享受的红包金额，单位元
	DiscountAmount *string `json:"discount_amount,omitzero"`
	// DiscountName represents discount_name. 商户数据回传的优惠信息的名称
	DiscountName *string `json:"discount_name,omitzero"`
	// DiscountType represents discount_type. 优惠信息的类型。为枚举值，商户可任意选择一个值作为入参传入。可取值： ("exclusiveBenefit","专享优惠"), ("exclusiveDiscount","专享折扣"), ("appreciationBenefit", "增值权益"), ("memberPoint", "会员积分")
	DiscountType *string `json:"discount_type,omitzero"`
}

// ZMGoTradeInfo is the official V3 ZMGoTradeInfo schema.
type ZMGoTradeInfo struct {
	// Amount represents amount. 用户和商户发生交易的交易单金额，单位元
	Amount *string `json:"amount,omitzero"`
	// TradeNo represents trade_no. 用户和商户发生交易的交易单号，仅在商户需要回传交易信息时传入。
	TradeNo *string `json:"trade_no,omitzero"`
}
