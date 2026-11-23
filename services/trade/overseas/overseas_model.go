package overseas

// OverseasSettleExtendParams 是跨境结算扩展参数。
type OverseasSettleExtendParams struct {
	// GoodsDetail 是商品明细列表（JSON 文本）。
	GoodsDetail *string `json:"goods_detail,omitempty"`
}
