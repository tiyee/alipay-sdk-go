package trade

// QueryResponse 是 alipay.trade.query 的业务响应。
// 交易状态由 TradeStatus 表示，按 QueryOptions 返回额外明细。
type QueryResponse struct {
	// TradeSettleInfo 是按查询选项返回的交易结算信息。
	TradeSettleInfo *TradeSettleInfo `json:"trade_settle_info,omitempty"`
	// VoucherDetailList 是本次交易使用的优惠券明细。
	VoucherDetailList []VoucherDetail `json:"voucher_detail_list,omitempty"`
	// DiscountGoodsDetail 是官方以字符串返回的单品券优惠明细。
	DiscountGoodsDetail *string `json:"discount_goods_detail,omitempty"`
	// MedicalInsuranceInfo 是官方以字符串返回的医保信息。
	MedicalInsuranceInfo *string `json:"medical_insurance_info,omitempty"`
	// TradeNo 是支付宝交易号；未生成真实交易时不返回。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商家订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// TradeStatus 是交易状态。
	TradeStatus TradeStatus `json:"trade_status,omitempty"`
	// TotalAmount 是交易的订单金额，为支付时传入的 total_amount。
	TotalAmount string `json:"total_amount,omitempty"`
	// FundBillList 是交易支付使用的资金渠道；仅在签约中指定或
	// query_options 指定时返回（文档标注必选但注明返回条件）。
	FundBillList []TradeFundBill `json:"fund_bill_list,omitempty"`
	// BuyerUserID 是买家在支付宝的用户 id；新商户建议使用 BuyerOpenID。
	BuyerUserID *string `json:"buyer_user_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// BuyerLogonID 是买家支付宝账号；未生成真实交易时不返回。
	BuyerLogonID *string `json:"buyer_logon_id,omitempty"`
	// BuyerUserType 是买家用户类型。
	BuyerUserType *BuyerUserType `json:"buyer_user_type,omitempty"`
	// BuyerPayAmount 是买家实付金额。
	BuyerPayAmount *string `json:"buyer_pay_amount,omitempty"`
	// ReceiptAmount 是实收金额。
	ReceiptAmount *string `json:"receipt_amount,omitempty"`
	// InvoiceAmount 是可开具发票的金额。
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// PointAmount 是积分支付的金额。
	PointAmount *string `json:"point_amount,omitempty"`
	// DiscountAmount 是平台优惠金额。
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// MDiscountAmount 是商家优惠金额。
	MDiscountAmount *string `json:"mdiscount_amount,omitempty"`
	// SendPayDate 是本次交易打款给卖家的时间。
	SendPayDate *string `json:"send_pay_date,omitempty"`
	// StoreID 是商户门店编号。
	StoreID *string `json:"store_id,omitempty"`
	// StoreName 是请求交易支付中的商户店铺名称。
	StoreName *string `json:"store_name,omitempty"`
	// TerminalID 是商户机具终端编号。
	TerminalID *string `json:"terminal_id,omitempty"`
	// ExtInfos 是交易额外信息，特殊场景下与支付宝约定返回的 JSON 文本。
	ExtInfos *string `json:"ext_infos,omitempty"`
}

// PreCreateResponse 是 alipay.trade.precreate 的业务响应。
// 二维码有效时间为 2 小时；同一订单号不能重复申请二维码。
type PreCreateResponse struct {
	// OutTradeNo 是商户的订单号。
	OutTradeNo string `json:"out_trade_no"`
	// QRCode 是预下单生成的二维码码串，有效时间 2 小时。
	QRCode *string `json:"qr_code,omitempty"`
	// ShareCode 是预下单生成的吱口令码串。
	ShareCode *string `json:"share_code,omitempty"`
	// PrepayID 是预下单 ID，用于唤起收银台支付。
	PrepayID *string `json:"prepay_id,omitempty"`
}

// CloseResponse 是 alipay.trade.close 的业务响应。
type CloseResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是创建交易传入的商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
}

// PayResponse 是 alipay.trade.pay 的业务响应。
// 异步支付场景请求成功仅代表受理成功，交易结果以异步通知或
// Query 查询为准；轮询查询语义见官方文档。
type PayResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// BuyerLogonID 是买家支付宝账号。
	BuyerLogonID *string `json:"buyer_logon_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// BuyerUserID 是买家在支付宝的用户 id。
	BuyerUserID *string `json:"buyer_user_id,omitempty"`
	// BuyerUserType 是买家用户类型。
	BuyerUserType *BuyerUserType `json:"buyer_user_type,omitempty"`
	// BuyerPayAmount 是买家实付金额。
	BuyerPayAmount *string `json:"buyer_pay_amount,omitempty"`
	// TotalAmount 是交易的订单金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// ReceiptAmount 是实收金额。
	ReceiptAmount *string `json:"receipt_amount,omitempty"`
	// InvoiceAmount 是可开具发票的金额。
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// PointAmount 是积分支付的金额。
	PointAmount *string `json:"point_amount,omitempty"`
	// DiscountAmount 是平台优惠金额。
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// MDiscountAmount 是商家优惠金额。
	MDiscountAmount *string `json:"mdiscount_amount,omitempty"`
	// GmtPayment 是交易支付时间。
	GmtPayment *string `json:"gmt_payment,omitempty"`
	// FundBillList 是交易支付使用的资金渠道；仅在签约中指定或
	// query_options 指定时返回。
	FundBillList []TradeFundBill `json:"fund_bill_list,omitempty"`
	// VoucherDetailList 是本交易支付使用的优惠券信息；
	// 仅在 query_options 指定时返回。
	VoucherDetailList []VoucherDetail `json:"voucher_detail_list,omitempty"`
	// StoreName 是发生支付交易的商户门店名称。
	StoreName *string `json:"store_name,omitempty"`
	// CardBalance 是支付宝卡余额。
	CardBalance *string `json:"card_balance,omitempty"`
	// AuthTradePayMode 是预授权支付模式；仅在信用预授权支付场景返回。
	AuthTradePayMode *string `json:"auth_trade_pay_mode,omitempty"`
	// AsyncPayApplyStatus 是异步支付受理状态；仅异步支付模式且
	// query_options 指定 async_pay_info 时返回。S 表示受理成功。
	AsyncPayApplyStatus *string `json:"async_pay_apply_status,omitempty"`
}

// CreateResponse 是 alipay.trade.create 的业务响应。
// 返回交易号仅表示订单创建成功，不代表买家已完成支付。
type CreateResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

// CancelResponse 是 alipay.trade.cancel 的业务响应。
// RetryFlag=Y 表示需要重试；Action 未返回表示未查询到交易或调用失败。
type CancelResponse struct {
	// TradeNo 是支付宝交易号；发生交易关闭或退款时返回。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// RetryFlag 标识是否需要重试。
	RetryFlag *string `json:"retry_flag,omitempty"`
	// Action 是本次撤销触发的交易动作：close（未支付，触发关闭）或
	// refund（已支付，触发退款）；未返回表示未查询到交易。
	Action *CancelAction `json:"action,omitempty"`
	// GmtRefundPay 是撤销产生退款时的退款时间。
	GmtRefundPay *string `json:"gmt_refund_pay,omitempty"`
	// RefundSettlementID 是撤销产生退款时的退款清算编号。
	RefundSettlementID *string `json:"refund_settlement_id,omitempty"`
}
