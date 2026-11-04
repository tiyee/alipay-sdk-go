package trade

// 请求与响应模型。字段名严格匹配官方协议；金额使用十进制字符串。
// 可选字段使用指针，nil 表示未传；响应中未知枚举值保留原始字符串。
// 契约依据：
//   - alipay.trade.query: https://opendocs.alipay.com/open/02ekfq.md
//   - alipay.trade.precreate: https://opendocs.alipay.com/open/02np92
//   - alipay.trade.refund: https://opendocs.alipay.com/open/02ekfs.md
//   - alipay.trade.fastpay.refund.query: https://opendocs.alipay.com/open/02ekft.md
//   - alipay.trade.close: https://opendocs.alipay.com/open/02ivbu.md

// QueryRequest 是 alipay.trade.query 的请求。
// out_trade_no 与 trade_no 二选一且不能同时为空；
// 同时存在时网关优先使用 trade_no。
type QueryRequest struct {
	// OutTradeNo 是订单支付时传入的商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []QueryOption `json:"query_options,omitempty"`
}

// TradeFundBill 是交易支付使用的资金渠道明细。
type TradeFundBill struct {
	// FundChannel 是交易使用的资金渠道；完整取值见官方支付渠道列表。
	FundChannel FundChannel `json:"fund_channel"`
	// Amount 是该支付工具类型所使用的金额。
	Amount string `json:"amount"`
	// RealAmount 是渠道实际付款金额。
	RealAmount *string `json:"real_amount,omitempty"`
	// FundType 是渠道所使用的资金类型，仅银行卡渠道返回。
	FundType *FundType `json:"fund_type,omitempty"`
}

// QueryResponse 是 alipay.trade.query 的业务响应。
// code=10000 仅表示查询成功；是否支付成功以 TradeStatus 为准。
type QueryResponse struct {
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

// PrecreateRequest 是 alipay.trade.precreate（统一收单线下交易预创建）的请求。
// out_trade_no、total_amount、subject 为必选。
type PrecreateRequest struct {
	// OutTradeNo 是商户订单号：64 字符以内，仅支持字母、数字、下划线，
	// 且需保证在商户端不重复。
	OutTradeNo string `json:"out_trade_no"`
	// TotalAmount 是订单总金额，两位小数，取值范围 [0.01,100000000]。
	TotalAmount string `json:"total_amount"`
	// Subject 是订单标题；不可使用 / = & 等特殊字符。
	Subject string `json:"subject"`
	// ProductCode 是产品码，默认 FACE_TO_FACE_PAYMENT。
	ProductCode *ProductCode `json:"product_code,omitempty"`
	// SellerID 是卖家支付宝用户 ID；收款账户优先级见官方文档。
	SellerID *string `json:"seller_id,omitempty"`
	// Body 是订单附加信息，将在异步通知、对账单中原样返回。
	Body *string `json:"body,omitempty"`
	// GoodsDetail 是订单包含的商品列表，上限 150 条。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
	// TimeExpire 是订单绝对超时时间（yyyy-MM-dd HH:mm:ss）；
	// 与 TimeoutExpress 同传时优先使用本字段。
	TimeExpire *string `json:"time_expire,omitempty"`
	// TimeoutExpress 是订单相对超时时间（1m～15d）；
	// 注意预下单二维码最长有效 2 小时。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams `json:"extend_params,omitempty"`
	// BusinessParams 是与支付宝约定的直传业务信息。
	BusinessParams *BusinessParams `json:"business_params,omitempty"`
	// UndiscountableAmount 是不可打折金额。
	UndiscountableAmount *string `json:"undiscountable_amount,omitempty"`
	// DiscountableAmount 是可打折金额；
	// 与 TotalAmount 同传时须满足 TotalAmount = DiscountableAmount + UndiscountableAmount。
	DiscountableAmount *string `json:"discountable_amount,omitempty"`
	// StoreID 是商户门店编号。
	StoreID *string `json:"store_id,omitempty"`
	// OperatorID 是商户操作员编号。
	OperatorID *string `json:"operator_id,omitempty"`
	// TerminalID 是商户机具终端编号。
	TerminalID *string `json:"terminal_id,omitempty"`
	// DisablePayChannels 是禁用渠道列表（逗号分割），与 EnablePayChannels 互斥。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// EnablePayChannels 是指定支付渠道列表（逗号分割），与 DisablePayChannels 互斥。
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// MerchantOrderNo 是商户的原始订单号。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// CodeType 是码类型；为空默认为订单码。
	CodeType *CodeType `json:"code_type,omitempty"`
	// QueryOptions 定制额外返回的信息字段。
	QueryOptions []PrecreateQueryOption `json:"query_options,omitempty"`
	// BkAgentReqInfo 是间联交易下由收单机构上送的信息。
	BkAgentReqInfo *BkAgentReqInfo `json:"bkagent_req_info,omitempty"`
}

// GoodsDetail 是预下单与支付的商品明细。
type GoodsDetail struct {
	// GoodsID 是商品编号（最小粒度 ID）。
	GoodsID string `json:"goods_id"`
	// GoodsName 是商品名称。
	GoodsName string `json:"goods_name"`
	// Quantity 是商品数量。
	Quantity int64 `json:"quantity"`
	// Price 是商品单价，单位为元。
	Price string `json:"price"`
	// GoodsCategory 是商品类目。
	GoodsCategory *string `json:"goods_category,omitempty"`
	// CategoriesTree 是商品类目树，类目 id 用 | 分割。
	CategoriesTree *string `json:"categories_tree,omitempty"`
	// ShowURL 是商品的展示地址。
	ShowURL *string `json:"show_url,omitempty"`
	// OutItemID 是商家侧小程序商品 ID。
	OutItemID *string `json:"out_item_id,omitempty"`
	// OutSkuID 是商家侧小程序商品 SKU ID。
	OutSkuID *string `json:"out_sku_id,omitempty"`
}

// SettleInfo 描述结算信息；目前只支持一条结算明细。
type SettleInfo struct {
	// SettleDetailInfos 是结算详细信息，目前只支持一条。
	SettleDetailInfos []SettleDetailInfo `json:"settle_detail_infos"`
}

// SettleDetailInfo 是单条结算明细。
type SettleDetailInfo struct {
	// TransInType 是结算收款方的账户类型。
	TransInType AccountType `json:"trans_in_type"`
	// TransIn 是结算收款方；类型为 defaultSettle 时不能传值。
	TransIn *string `json:"trans_in,omitempty"`
	// Amount 是结算的金额，创建订单时必须与交易金额相同。
	Amount *string `json:"amount,omitempty"`
	// SummaryDimension 是结算汇总维度，需与 cardAliasNo 配合使用。
	SummaryDimension *string `json:"summary_dimension,omitempty"`
	// SettleEntityID 是结算主体标识。
	SettleEntityID *string `json:"settle_entity_id,omitempty"`
	// SettleEntityType 是结算主体类型：SecondMerchant 或 Store。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// ActualAmount 仅用于直付通账期模式的部分确认结算。
	ActualAmount *string `json:"actual_amount,omitempty"`
	// SettlePeriodTime 是超期自动确认结算时间（1d～365d），仅账期结算模式有效。
	SettlePeriodTime *string `json:"settle_period_time,omitempty"`
}

// ExtendParams 是预下单的业务扩展参数。
type ExtendParams struct {
	// SysServiceProviderID 是系统商编号（返佣数据提取依据）。
	SysServiceProviderID *string `json:"sys_service_provider_id,omitempty"`
	// CardType 是卡类型。
	CardType *string `json:"card_type,omitempty"`
	// SpecifiedSellerName 允许商户指定交易展示的卖家名称。
	SpecifiedSellerName *string `json:"specified_seller_name,omitempty"`
	// RoyaltyFreeze 为 true 时进行资金冻结，用于后续分账。
	RoyaltyFreeze *string `json:"royalty_freeze,omitempty"`
	// TCInstallmentOrderID 是公域商品交易分期单 ID。
	TCInstallmentOrderID *string `json:"tc_installment_order_id,omitempty"`
	// TradeComponentOrderID 是公域商品交易业务订单 ID。
	TradeComponentOrderID *string `json:"trade_component_order_id,omitempty"`
}

// BusinessParams 是商户与支付宝约定的直传业务信息。
type BusinessParams struct {
	// CampusCard 是校园卡编号。
	CampusCard *string `json:"campus_card,omitempty"`
	// CardType 是虚拟卡卡类型。
	CardType *string `json:"card_type,omitempty"`
	// ActualOrderTime 是实际订单时间（乘车码场景为刷码乘车时间）。
	ActualOrderTime *string `json:"actual_order_time,omitempty"`
	// GoodTaxes 是商户传入的交易税费。
	GoodTaxes *string `json:"good_taxes,omitempty"`
	// EnterprisePayInfo 是因公付业务信息（JSON 文本）。
	EnterprisePayInfo *string `json:"enterprise_pay_info,omitempty"`
	// EnterprisePayAmount 是因公付金额。
	EnterprisePayAmount *string `json:"enterprise_pay_amount,omitempty"`
	// MCCreateTradeIP 是商户端创建订单的 IP（ipv4/ipv6）。
	MCCreateTradeIP *string `json:"mc_create_trade_ip,omitempty"`
	// TinyAppMerchantBizType 是小程序细分业务场景。
	TinyAppMerchantBizType *string `json:"tiny_app_merchant_biz_type,omitempty"`
}

// BkAgentReqInfo 是间联交易下由收单机构上送的信息。
type BkAgentReqInfo struct {
	// AcqCode 是收单机构在清算组织登记或分配的机构代码。
	AcqCode *string `json:"acq_code,omitempty"`
	// MerchCode 是收单机构在清算组织登记的商户编码。
	MerchCode *string `json:"merch_code,omitempty"`
	// DeviceType 是终端设备类型（如 01 ATM、02 传统 POS、03 mPOS）。
	DeviceType *string `json:"device_type,omitempty"`
	// SerialNum 是终端序列号；特定设备类型下必填。
	SerialNum *string `json:"serial_num,omitempty"`
	// Location 是终端设备实时经纬度，格式为纬度/经度。
	Location *string `json:"location,omitempty"`
	// PaymentType 是支付类型（如 agent_pay 智能体 SDK 支付）。
	PaymentType *string `json:"payment_type,omitempty"`
}

// PrecreateResponse 是 alipay.trade.precreate 的业务响应。
// 二维码有效时间为 2 小时；同一订单号不能重复申请二维码。
type PrecreateResponse struct {
	// OutTradeNo 是商户的订单号。
	OutTradeNo string `json:"out_trade_no"`
	// QRCode 是预下单生成的二维码码串，有效时间 2 小时。
	QRCode *string `json:"qr_code,omitempty"`
	// ShareCode 是预下单生成的吱口令码串。
	ShareCode *string `json:"share_code,omitempty"`
	// PrepayID 是预下单 ID，用于唤起收银台支付。
	PrepayID *string `json:"prepay_id,omitempty"`
}

// RefundRequest 是 alipay.trade.refund（交易退款）的请求。
// out_trade_no 与 trade_no 二选一；out_request_no 在部分退款时必选，
// 且重试时必须保持不变以防重复退款。
type RefundRequest struct {
	// RefundAmount 是本次退款金额，不能大于订单金额。
	RefundAmount string `json:"refund_amount"`
	// OutTradeNo 是商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// TradeNo 是支付宝交易号；与 OutTradeNo 同传时优先使用本字段。
	TradeNo string `json:"trade_no,omitempty"`
	// OutRequestNo 是退款请求号，需保证在交易号下唯一；
	// 部分退款时必选；同请求号多次请求只退一次。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RefundReason 是退款原因说明。
	RefundReason *string `json:"refund_reason,omitempty"`
	// RefundGoodsDetail 是退款包含的商品列表。
	RefundGoodsDetail []RefundGoodsDetail `json:"refund_goods_detail,omitempty"`
	// RefundRoyaltyParameters 是退分账明细信息。
	RefundRoyaltyParameters []RoyaltyDetailInfo `json:"refund_royalty_parameters,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []RefundQueryOption `json:"query_options,omitempty"`
	// RelatedSettleConfirmNo 用于账期交易确认结算后退款。
	RelatedSettleConfirmNo *string `json:"related_settle_confirm_no,omitempty"`
}

// RefundGoodsDetail 是退款包含的单个商品信息。
type RefundGoodsDetail struct {
	// GoodsID 是商品编号，对应支付时传入的 goods_id。
	GoodsID string `json:"goods_id"`
	// RefundAmount 是该商品的退款总金额。
	RefundAmount string `json:"refund_amount"`
	// OutCertificateNoList 是外部商品凭证编号列表。
	OutCertificateNoList []string `json:"out_certificate_no_list,omitempty"`
	// OutItemID 是商家侧小程序商品 ID。
	OutItemID *string `json:"out_item_id,omitempty"`
	// OutSkuID 是商家侧小程序商品 SKU ID。
	OutSkuID *string `json:"out_sku_id,omitempty"`
}

// RoyaltyDetailInfo 是退分账明细（OpenApiRoyaltyDetailInfoPojo）。
type RoyaltyDetailInfo struct {
	// RoyaltyType 是分账类型；为空默认为普通分账 transfer。
	RoyaltyType *RoyaltyType `json:"royalty_type,omitempty"`
	// TransOut 是支出方账户。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutType 是支出方账户类型。
	TransOutType *AccountType `json:"trans_out_type,omitempty"`
	// TransInType 是收入方账户类型。
	TransInType *AccountType `json:"trans_in_type,omitempty"`
	// TransIn 是收入方账户。
	TransIn *string `json:"trans_in,omitempty"`
	// Amount 是分账的金额。
	Amount *string `json:"amount,omitempty"`
	// Desc 是分账描述。
	Desc *string `json:"desc,omitempty"`
	// RoyaltyScene 是分账场景（达人佣金、平台服务费等）。
	RoyaltyScene *string `json:"royalty_scene,omitempty"`
	// TransInName 是分账收款方姓名；上送则进行一致性校验。
	TransInName *string `json:"trans_in_name,omitempty"`
}

// RefundResponse 是 alipay.trade.refund 的业务响应。
// code=10000 仅代表退款请求成功；FundChange=Y 才表示退款成功，
// 否则需通过退款查询接口确认。
type RefundResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// BuyerLogonID 是用户的登录 id。
	BuyerLogonID string `json:"buyer_logon_id,omitempty"`
	// RefundFee 是该笔交易累计已经退款成功的金额。
	RefundFee string `json:"refund_fee,omitempty"`
	// FundChange 标识本次退款是否发生了资金变化；
	// Y 表示退款成功，N 或未返回需退款查询确认。
	FundChange *FundChange `json:"fund_change,omitempty"`
	// RefundDetailItemList 是退款使用的资金渠道，
	// 仅在签约指定或 query_options 指定时返回。
	RefundDetailItemList []TradeFundBill `json:"refund_detail_item_list,omitempty"`
	// StoreName 是交易在支付时候的门店名称。
	StoreName *string `json:"store_name,omitempty"`
	// BuyerUserID 是买家在支付宝的用户 id。
	BuyerUserID *string `json:"buyer_user_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// SendBackFee 是本次商户实际退回金额，需 query_options 指定后返回。
	SendBackFee *string `json:"send_back_fee,omitempty"`
	// PreAuthCancelFee 是撤销的预授权金额。
	PreAuthCancelFee *string `json:"pre_auth_cancel_fee,omitempty"`
	// RefundHybAmount 是本次请求退惠营宝金额。
	RefundHybAmount *string `json:"refund_hyb_amount,omitempty"`
	// RefundChargeInfoList 是退费信息。
	RefundChargeInfoList []RefundChargeInfo `json:"refund_charge_info_list,omitempty"`
	// RefundVoucherDetailList 是本次退款退的券信息，
	// 仅在 query_options 指定 refund_voucher_detail_list 时返回。
	RefundVoucherDetailList []VoucherDetail `json:"refund_voucher_detail_list,omitempty"`
}

// RefundChargeInfo 是单条退费信息。
type RefundChargeInfo struct {
	// RefundChargeFee 是实退费用。
	RefundChargeFee *string `json:"refund_charge_fee,omitempty"`
	// SwitchFeeRate 是签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
	// ChargeType 是手续费类型。
	ChargeType *ChargeType `json:"charge_type,omitempty"`
	// RefundSubFeeDetailList 是组合支付退费明细。
	RefundSubFeeDetailList []RefundSubFee `json:"refund_sub_fee_detail_list,omitempty"`
}

// RefundSubFee 是组合支付的退费明细子项。
type RefundSubFee struct {
	// RefundChargeFee 是实退费用。
	RefundChargeFee *string `json:"refund_charge_fee,omitempty"`
	// SwitchFeeRate 是签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
}

// VoucherDetail 是优惠券信息。
type VoucherDetail struct {
	// ID 是券 id。
	ID string `json:"id"`
	// Name 是券名称。
	Name string `json:"name"`
	// Type 是券类型；官方注明可能新增类型，未知值原样保留。
	Type VoucherType `json:"type"`
	// Amount 是优惠券面额，应等于商家出资加上其他出资方出资。
	Amount string `json:"amount"`
	// MerchantContribute 是商家出资金额。
	MerchantContribute *string `json:"merchant_contribute,omitempty"`
	// OtherContribute 是其他出资方出资金额。
	OtherContribute *string `json:"other_contribute,omitempty"`
	// Memo 是优惠券备注信息。
	Memo *string `json:"memo,omitempty"`
	// TemplateID 是券模板 id。
	TemplateID *string `json:"template_id,omitempty"`
	// OtherContributeDetail 是优惠券的其他出资方明细。
	OtherContributeDetail []ContributeDetail `json:"other_contribute_detail,omitempty"`
	// PurchaseBuyerContribute 是用户购买该券实际付款金额。
	PurchaseBuyerContribute *string `json:"purchase_buyer_contribute,omitempty"`
	// PurchaseMerchantContribute 是购买该券时商户优惠金额。
	PurchaseMerchantContribute *string `json:"purchase_merchant_contribute,omitempty"`
	// PurchaseAntContribute 是购买该券时平台优惠金额。
	PurchaseAntContribute *string `json:"purchase_ant_contribute,omitempty"`
}

// ContributeDetail 是优惠券的出资方明细。
type ContributeDetail struct {
	// ContributeType 是出资方类型；官方注明可能新增类型。
	ContributeType ContributeType `json:"contribute_type"`
	// ContributeAmount 是出资方金额。
	ContributeAmount string `json:"contribute_amount"`
}

// QueryRefundRequest 是 alipay.trade.fastpay.refund.query 的请求。
type QueryRefundRequest struct {
	// OutRequestNo 是退款请求号（必选）：退款时传入的请求号；
	// 退款时未传入则为创建交易时的商户订单号。
	OutRequestNo string `json:"out_request_no"`
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号；与 TradeNo 同传时优先使用 trade_no。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []RefundQueryOption `json:"query_options,omitempty"`
}

// RefundRoyaltyResult 是退分账结果（直付通产品返回）。
type RefundRoyaltyResult struct {
	// RefundAmount 是退分账金额。
	RefundAmount string `json:"refund_amount"`
	// ResultCode 是退分账结果码；未知值原样保留。
	ResultCode RoyaltyResultCode `json:"result_code"`
	// RoyaltyType 是分账类型；为空默认为 transfer。
	RoyaltyType *RoyaltyType `json:"royalty_type,omitempty"`
	// TransOut 是转出人支付宝账号对应用户 ID。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutEmail 是转出人支付宝账号。
	TransOutEmail *string `json:"trans_out_email,omitempty"`
	// TransIn 是转入人支付宝账号对应用户 ID。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInEmail 是转入人支付宝账号。
	TransInEmail *string `json:"trans_in_email,omitempty"`
	// OriTransOut 是商户请求的转出账号。
	OriTransOut *string `json:"ori_trans_out,omitempty"`
	// OriTransIn 是商户请求的转入账号。
	OriTransIn *string `json:"ori_trans_in,omitempty"`
}

// DepositBackInfo 是银行卡冲退信息。
type DepositBackInfo struct {
	// HasDepositBack 标识是否存在银行卡冲退信息（"true"/"false"）。
	HasDepositBack *string `json:"has_deposit_back,omitempty"`
	// DbackStatus 是银行卡冲退状态；冲退失败时资金自动转入用户余额。
	DbackStatus *DepositBackStatus `json:"dback_status,omitempty"`
	// DbackAmount 是银行卡冲退金额。
	DbackAmount *string `json:"dback_amount,omitempty"`
	// BankAckTime 是银行响应时间。
	BankAckTime *string `json:"bank_ack_time,omitempty"`
	// EstBankReceiptTime 是预估银行到账时间。
	EstBankReceiptTime *string `json:"est_bank_receipt_time,omitempty"`
}

// QueryRefundResponse 是退款查询的业务响应。
// code=10000 仅代表查询操作成功；仅当 RefundStatus 为 REFUND_SUCCESS
// 时表示退款成功，字段未返回表示退款请求未收到或退款失败。
// 官方建议退款发起后间隔 10 秒以上再查询。
type QueryRefundResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是创建交易传入的商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// OutRequestNo 是本笔退款对应的退款请求号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// TotalAmount 是该笔退款所对应交易的订单金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// RefundAmount 是本次退款请求对应的退款金额。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundStatus 是退款状态；未返回表示退款请求未收到或退款失败。
	RefundStatus *RefundStatus `json:"refund_status,omitempty"`
	// RefundRoyaltys 是退分账明细信息，当前仅在直付通产品中返回。
	RefundRoyaltys []RefundRoyaltyResult `json:"refund_royaltys,omitempty"`
	// GmtRefundPay 是退款时间，需 query_options 指定 gmt_refund_pay 才返回。
	GmtRefundPay *string `json:"gmt_refund_pay,omitempty"`
	// RefundDetailItemList 是本次退款使用的资金渠道，需 query_options 指定。
	RefundDetailItemList []TradeFundBill `json:"refund_detail_item_list,omitempty"`
	// SendBackFee 是本次商户实际退回金额，需 query_options 指定。
	SendBackFee *string `json:"send_back_fee,omitempty"`
	// DepositBackInfo 是银行卡冲退信息，需 query_options 指定 deposit_back_info。
	DepositBackInfo *DepositBackInfo `json:"deposit_back_info,omitempty"`
	// RefundVoucherDetailList 是本交易支付时使用的所有优惠券信息，
	// 需 query_options 指定 refund_voucher_detail_list。
	RefundVoucherDetailList []VoucherDetail `json:"refund_voucher_detail_list,omitempty"`
	// PreAuthCancelFee 是撤销的预授权金额。
	PreAuthCancelFee *string `json:"pre_auth_cancel_fee,omitempty"`
	// RefundHybAmount 是本次退款金额中退惠营宝的金额。
	RefundHybAmount *string `json:"refund_hyb_amount,omitempty"`
	// RefundChargeInfoList 是退费信息。
	RefundChargeInfoList []RefundChargeInfo `json:"refund_charge_info_list,omitempty"`
	// DepositBackInfoList 是银行卡冲退信息列表；文档中该字段的
	// query_options 取值存在原文不一致，以官方确认为准。
	DepositBackInfoList []DepositBackInfo `json:"deposit_back_info_list,omitempty"`
}

// CloseRequest 是 alipay.trade.close 的请求。
// trade_no 与 out_trade_no 二选一且不能同时为空；同传时以 trade_no 为准。
// 仅"等待买家付款"状态的交易可关单。
type CloseRequest struct {
	// TradeNo 是支付宝交易号（最短 16 位）。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是订单支付时传入的商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// OperatorID 是商家操作员编号 id。
	OperatorID *string `json:"operator_id,omitempty"`
}

// CloseResponse 是 alipay.trade.close 的业务响应。
type CloseResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是创建交易传入的商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
}
