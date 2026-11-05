package orderinfo

// OrderInfoSyncResponse 是订单信息同步的业务响应。
type OrderInfoSyncResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// BuyerUserID 是买家在支付宝的用户 id。
	BuyerUserID *string `json:"buyer_user_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
}
