package usertrade

// GoodsInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type GoodsInfoDTO struct {
	// GoodsAmount 是商品金额。
	GoodsAmount *string `json:"goods_amount,omitempty"`
	// GoodsCount 是商品购买数量。
	GoodsCount *string `json:"goods_count,omitempty"`
	// GoodsId 是商品标准69码。
	GoodsId *string `json:"goods_id,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
}

// UserTradeInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type UserTradeInfoDTO struct {
	// GoodsInfoList 是商品信息列表。
	GoodsInfoList *[]GoodsInfoDTO `json:"goods_info_list,omitempty"`
	// PartnerId 是门店ISV的PID。
	PartnerId *string `json:"partner_id,omitempty"`
	// RiskLevel 是是否是风险交易：NO_RISK-无风险；POTENTIAL_RISK-潜在风险（中等风险）；HIGH_RISK-高风险。
	RiskLevel *string `json:"risk_level,omitempty"`
	// TradeAmount 是交易总金额，单位元，精确两位小数点。
	TradeAmount *string `json:"trade_amount,omitempty"`
	// TradeNo 是交易单号。
	TradeNo *string `json:"trade_no,omitempty"`
	// TradeTime 是交易时间。
	TradeTime *string `json:"trade_time,omitempty"`
	// UnfilteredTotalGoodsCount 是商品数据会根据活动商品列表进行过滤，该字段代表未过滤的商品列表大小。
	UnfilteredTotalGoodsCount *int64 `json:"unfiltered_total_goods_count,omitempty"`
	// UserId 是消费者支付宝ID。
	UserId *string `json:"user_id,omitempty"`
}
