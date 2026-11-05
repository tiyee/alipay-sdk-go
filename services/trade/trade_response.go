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
	// AdditionalStatus 表示交易附加状态。
	AdditionalStatus *string `json:"additional_status,omitempty"`
	// AlipayStoreID 表示支付宝店铺编号。
	AlipayStoreID *string `json:"alipay_store_id,omitempty"`
	// AlipaySubMerchantID 表示间连商户在支付宝端的商户编号。
	AlipaySubMerchantID *string `json:"alipay_sub_merchant_id,omitempty"`
	// AsyncPayApplyStatus 表示异步支付受理状态，仅异步支付模式且query_options指定async_pay_info时返回。
	AsyncPayApplyStatus *string `json:"async_pay_apply_status,omitempty"`
	// AuthTradePayMode 表示预授权支付模式，该参数仅在信用预授权支付场景下返回。
	AuthTradePayMode *string `json:"auth_trade_pay_mode,omitempty"`
	// BizSettleMode 表示账期结算标识，指已完成支付的订单会进行账期管控，不会实时结算。
	BizSettleMode *BizSettleMode `json:"biz_settle_mode,omitempty"`
	// BkAgentRespInfo 是收单机构的间联交易返回信息。
	BkAgentRespInfo *BkAgentRespInfo `json:"bkagent_resp_info,omitempty"`
	// Body 表示订单描述。
	Body *string `json:"body,omitempty"`
	// BuyerUserName 表示买家名称。
	BuyerUserName *string `json:"buyer_user_name,omitempty"`
	// CashierType 表示收银台类型。
	CashierType *string `json:"cashier_type,omitempty"`
	// ChargeAmount 表示该笔交易针对收款方的收费金额；单位：元。
	ChargeAmount *string `json:"charge_amount,omitempty"`
	// ChargeFlags 表示费率活动标识。
	ChargeFlags *string `json:"charge_flags,omitempty"`
	// ChargeInfoList 表示计费信息列表。
	ChargeInfoList []ChargeInfo `json:"charge_info_list,omitempty"`
	// CreditBizOrderID 表示信用业务单号。
	CreditBizOrderID *string `json:"credit_biz_order_id,omitempty"`
	// CreditPayMode 表示信用支付模式。
	CreditPayMode *string `json:"credit_pay_mode,omitempty"`
	// EnterprisePayInfo 是因公支付信息。
	EnterprisePayInfo *EnterprisePayInfo `json:"enterprise_pay_info,omitempty"`
	// ExternalPayerInfo 是外部付款方信息。
	ExternalPayerInfo *ExternalPayerInfo `json:"external_payer_info,omitempty"`
	// FulfillmentDetailList 表示履约详情列表。
	FulfillmentDetailList []FulfillmentDetail `json:"fulfillment_detail_list,omitempty"`
	// HbFqPayInfo 是花呗分期付款信息。
	HbFqPayInfo *HbFqPayInfo `json:"hb_fq_pay_info,omitempty"`
	// HybAmount 表示惠营宝回票金额。
	HybAmount *string `json:"hyb_amount,omitempty"`
	// IndustrySpecDetail 表示行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
	IndustrySpecDetail *string `json:"industry_sepc_detail,omitempty"`
	// IndustrySpecDetailAcc 表示行业特殊信息-个账相关。
	IndustrySpecDetailAcc *string `json:"industry_sepc_detail_acc,omitempty"`
	// IndustrySpecDetailGov 表示行业特殊信息-统筹相关。
	IndustrySpecDetailGov *string `json:"industry_sepc_detail_gov,omitempty"`
	// IntactChargeInfoList 表示完整的收费信息。
	IntactChargeInfoList []IntactChargeInfo `json:"intact_charge_info_list,omitempty"`
	// OpenID 表示买家支付宝用户号，该字段将废弃，不要使用。
	OpenID *string `json:"open_id,omitempty"`
	// PassbackParams 表示公用回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
	// PayAmount 表示支付币种订单金额。
	PayAmount *string `json:"pay_amount,omitempty"`
	// PayCurrency 表示订单支付币种。
	PayCurrency *string `json:"pay_currency,omitempty"`
	// PaymentInfoWithIDList 表示带支付宝支付id的信息，暂不对外开放。
	PaymentInfoWithIDList []PaymentInfoWithID `json:"payment_info_with_id_list,omitempty"`
	// PeriodScene 表示该字段用于描述当前账期交易的场景。
	PeriodScene *string `json:"period_scene,omitempty"`
	// PreAuthPayAmount 表示当用户使用芝麻信用先享后付时，会返回该字段，代表整笔交易的原始待履约金额，单位元。
	PreAuthPayAmount *string `json:"pre_auth_pay_amount,omitempty"`
	// PurchaseEntity 表示支付宝钱包跨境支付交易的购汇主体，ALIPAYHK-境外香港购汇实体，3008境内购汇实体。
	PurchaseEntity *string `json:"purchase_entity,omitempty"`
	// ReceiptCurrencyType 表示收款资金类型，当交易收款资金为数字人民币时返回值为“DC”，否则不返回该字段。
	ReceiptCurrencyType *string `json:"receipt_currency_type,omitempty"`
	// ReqGoodsDetail 表示支付请求的商品明细列表。
	ReqGoodsDetail []GoodsDetail `json:"req_goods_detail,omitempty"`
	// SettleAmount 表示结算币种订单金额。
	SettleAmount *string `json:"settle_amount,omitempty"`
	// SettleCurrency 是结算币种。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// SettleTransRate 表示结算币种兑换标价币种汇率。
	SettleTransRate *string `json:"settle_trans_rate,omitempty"`
	// SettlementID 表示支付清算编号，用于清算对账。
	SettlementID *string `json:"settlement_id,omitempty"`
	// Subject 表示订单标题。
	Subject *string `json:"subject,omitempty"`
	// TapPayInfo 是碰一下支付信息。
	TapPayInfo *TapPayInfo `json:"tap_pay_info,omitempty"`
	// TransCurrency 是订单的交易币种。
	TransCurrency *string `json:"trans_currency,omitempty"`
	// TransPayRate 表示标价币种兑换支付币种汇率。
	TransPayRate *string `json:"trans_pay_rate,omitempty"`
	// UnpaidReason 表示有限条件下返回当前订单未支付原因。
	UnpaidReason *string `json:"unpaid_reason,omitempty"`
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
	// AdvanceAmount 表示先享后付2.0垫资金额,不返回表示没有走垫资，非空表示垫资支付的金额。
	AdvanceAmount *string `json:"advance_amount,omitempty"`
	// AsyncPaymentMode 表示异步支付模式，目前有五种值。
	AsyncPaymentMode *string `json:"async_payment_mode,omitempty"`
	// BkAgentRespInfo 是收单机构的间联交易返回信息。
	BkAgentRespInfo *BkAgentRespInfo `json:"bkagent_resp_info,omitempty"`
	// BusinessParams 表示商户传入业务信息，具体值要和支付宝约定。
	BusinessParams *string `json:"business_params,omitempty"`
	// BuyerUserName 表示买家名称。
	BuyerUserName *string `json:"buyer_user_name,omitempty"`
	// CanTurnToAppPay 表示是否可以转为app支付，仅当商户代扣失败场景才会返回该字段信息。
	CanTurnToAppPay *string `json:"can_turn_to_app_pay,omitempty"`
	// ChargeAmount 表示该笔交易针对收款方的收费金额。
	ChargeAmount *string `json:"charge_amount,omitempty"`
	// ChargeFlags 表示费率活动标识。
	ChargeFlags *string `json:"charge_flags,omitempty"`
	// ChargeInfoList 是本次交易的收费信息。
	ChargeInfoList *ChargeInfo `json:"charge_info_list,omitempty"`
	// CreditBizOrderID 表示信用业务单号。
	CreditBizOrderID *string `json:"credit_biz_order_id,omitempty"`
	// CreditBizOrderStatus 表示先享免密下单之后，芝麻信用侧的业务单状态。
	CreditBizOrderStatus *string `json:"credit_biz_order_status,omitempty"`
	// CreditPayMode 表示信用支付模式。
	CreditPayMode *string `json:"credit_pay_mode,omitempty"`
	// DiscountGoodsDetail 表示本次交易支付所使用的单品券优惠的商品优惠信息。
	DiscountGoodsDetail *string `json:"discount_goods_detail,omitempty"`
	// EnterprisePayInfo 是因公支付信息。
	EnterprisePayInfo *EnterprisePayInfo `json:"enterprise_pay_info,omitempty"`
	// HybAmount 表示惠营宝回票金额。
	HybAmount *string `json:"hyb_amount,omitempty"`
	// OpenID 表示买家支付宝用户号,该参数已废弃，请不要使用。
	OpenID *string `json:"open_id,omitempty"`
	// PayAmount 表示支付币种订单金额。
	PayAmount *string `json:"pay_amount,omitempty"`
	// PayCurrency 表示支付币种。
	PayCurrency *string `json:"pay_currency,omitempty"`
	// PreAuthPayAmount 表示当用户使用芝麻信用先享后付时，会返回该字段，代表整笔交易的原始待履约金额，单位元。
	PreAuthPayAmount *string `json:"pre_auth_pay_amount,omitempty"`
	// PurchaseEntity 表示支付宝钱包跨境支付交易的购汇主体，ALIPAYHK-境外香港购汇实体，3008境内购汇实体。
	PurchaseEntity *string `json:"purchase_entity,omitempty"`
	// ReceiptCurrencyType 表示收款资金类型，当交易收款资金为数字人民币时返回值为“DC”，否则不返回该字段。
	ReceiptCurrencyType *string `json:"receipt_currency_type,omitempty"`
	// SettleAmount 表示结算币种订单金额。
	SettleAmount *string `json:"settle_amount,omitempty"`
	// SettleCurrency 是结算币种。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// SettleTransRate 表示结算币种兑换标价币种汇率。
	SettleTransRate *string `json:"settle_trans_rate,omitempty"`
	// SettlementID 表示支付清算编号，用于清算对账使用。
	SettlementID *string `json:"settlement_id,omitempty"`
	// TapPayInfo 是碰一下支付信息。
	TapPayInfo *TapPayInfo `json:"tap_pay_info,omitempty"`
	// TransCurrency 表示标价币种, total_amount对应的币种单位。
	TransCurrency *string `json:"trans_currency,omitempty"`
	// TransPayRate 表示标价币种兑换支付币种汇率。
	TransPayRate *string `json:"trans_pay_rate,omitempty"`
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
