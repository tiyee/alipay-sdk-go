package order

// OrderPayResponse 是 alipay.trade.order.pay 的业务响应。
// 异步支付场景以异步通知或查询确认最终结果。
type OrderPayResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// OutRequestNo 是商户请求号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// TotalAmount 是订单总金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// GmtPayment 是交易支付时间。
	GmtPayment *string `json:"gmt_payment,omitempty"`
	// AsyncPaymentMode 是异步支付模式；使用异步支付模式时才有值。
	AsyncPaymentMode *string `json:"async_payment_mode,omitempty"`
	// FulfillmentAmount 是实际履约金额；仅履约场景返回。
	FulfillmentAmount *string `json:"fulfillment_amount,omitempty"`
}

// OrderSettleResponse 是 alipay.trade.order.settle 的业务响应。
type OrderSettleResponse struct {
	// SettleNo 是支付宝分账单号，可用于查询单次分账请求执行结果。
	SettleNo *string `json:"settle_no,omitempty"`
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
}

// OrderSettleQueryResponse 是 alipay.trade.order.settle.query 的业务响应。
type OrderSettleQueryResponse struct {
	// OutRequestNo 是商户分账请求单号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// OperationDt 是分账受理时间。
	OperationDt *string `json:"operation_dt,omitempty"`
	// RoyaltyDetailList 是分账明细。
	RoyaltyDetailList []RoyaltyDetail `json:"royalty_detail_list,omitempty"`
}
