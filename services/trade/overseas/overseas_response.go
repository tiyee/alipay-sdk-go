package overseas

// OverseasSettleResponse 是跨境结算的业务响应。
type OverseasSettleResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutRequestNo 是请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SettleAmount 是本次跨境结算的实际结算人民币金额。
	SettleAmount *string `json:"settle_amount,omitempty"`
	// ForeignSettleAmount 是本次跨境结算的实际结算外币金额。
	ForeignSettleAmount *string `json:"foreign_settle_amount,omitempty"`
	// ForeignSettleCurrency 是实际结算的外币币种。
	ForeignSettleCurrency *string `json:"foreign_settle_currency,omitempty"`
	// ExchangeRate 是本笔跨境结算使用的汇率，精度 8 位小数。
	ExchangeRate *string `json:"exchange_rate,omitempty"`
}
