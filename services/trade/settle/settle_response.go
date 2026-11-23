package settle

// SettleConfirmResponse 是 alipay.trade.settle.confirm 的业务响应。
type SettleConfirmResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutRequestNo 是确认结算请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SettleAmount 是本次确认结算的实际结算金额。
	SettleAmount *string `json:"settle_amount,omitempty"`
}

// SettleReceivablesQueryResponse 是待结算查询的业务响应。
type SettleReceivablesQueryResponse struct {
	// SettledAmount 是已结算金额，单位为元。
	SettledAmount *string `json:"settled_amount,omitempty"`
	// UnsettledAmount 是待结算金额，单位为元。
	UnsettledAmount *string `json:"unsettled_amount,omitempty"`
	// OnSettleAmount 是直付通请款模式下的结算在途金额。
	OnSettleAmount *string `json:"on_settle_amount,omitempty"`
}
