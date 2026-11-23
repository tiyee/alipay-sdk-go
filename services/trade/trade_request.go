package trade

import (
	"net/url"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

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
	// OrgPID 表示银行间联模式的机构 PID。
	OrgPID *string `json:"org_pid,omitempty"`
}

// PreCreateRequest 是 alipay.trade.precreate（统一收单线下交易预创建）的请求。
// out_trade_no、total_amount、subject 为必选。
type PreCreateRequest struct {
	// NotifyURL 是本次请求接收异步通知的地址。
	NotifyURL *string `json:"notify_url,omitempty"`
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
	QueryOptions []PreCreateQueryOption `json:"query_options,omitempty"`
	// BkAgentReqInfo 是间联交易下由收单机构上送的信息。
	BkAgentReqInfo *BkAgentReqInfo `json:"bkagent_req_info,omitempty"`
	// AlipayStoreID 表示支付宝店铺编号。
	AlipayStoreID *string `json:"alipay_store_id,omitempty"`
	// BuyerLogonID 表示买家支付宝账号。
	BuyerLogonID *string `json:"buyer_logon_id,omitempty"`
	// ExtUserInfo 是指定买家的身份信息。
	ExtUserInfo *ExtUserInfo `json:"ext_user_info,omitempty"`
	// PassbackParams 表示公用回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
	// PaymentType 表示指定支付方式。
	PaymentType *PaymentType `json:"payment_type,omitempty"`
	// QRCodeTimeoutExpress 表示二维码订单相对超时时间。
	QRCodeTimeoutExpress *string `json:"qr_code_timeout_express,omitempty"`
	// RoyaltyInfo 是本次支付的分账信息。
	RoyaltyInfo *RoyaltyInfo `json:"royalty_info,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
}

// CloseRequest 是 alipay.trade.close 的请求。
// trade_no 与 out_trade_no 二选一且不能同时为空；同传时以 trade_no 为准。
// 仅"等待买家付款"状态的交易可关单。
type CloseRequest struct {
	// NotifyURL 是本次请求接收异步通知的地址。
	NotifyURL *string `json:"notify_url,omitempty"`
	// TradeNo 是支付宝交易号（最短 16 位）。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是订单支付时传入的商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// OperatorID 是商家操作员编号 id。
	OperatorID *string `json:"operator_id,omitempty"`
	// ExtendParams 是关单扩展参数。
	ExtendParams *CloseExtendParams `json:"extend_params,omitempty"`
}

// PayRequest 是 alipay.trade.pay（统一收单交易支付）的请求。
// 当面付场景：out_trade_no、total_amount、subject、scene、auth_code
// 为必选；周期扣款场景无需 scene 与 auth_code，协议号经 AgreementParams 传入。
type PayRequest struct {
	// NotifyURL 是本次请求接收异步通知的地址。
	NotifyURL *string `json:"notify_url,omitempty"`
	// OutTradeNo 是商户订单号：64 字符以内，仅支持字母、数字、下划线，
	// 且需保证在商户端不重复。
	OutTradeNo string `json:"out_trade_no"`
	// TotalAmount 是订单总金额，两位小数，取值范围 [0.01,100000000]。
	TotalAmount string `json:"total_amount"`
	// Subject 是订单标题；不可使用 / = & 等特殊字符。
	Subject string `json:"subject"`
	// Scene 是支付场景：bar_code（条码支付）或 security_code（刷脸支付）；
	// 周期扣款或代扣场景无需传入。
	Scene *PayScene `json:"scene,omitempty"`
	// AuthCode 是买家付款码（25~30 开头的 16~24 位数字）或刷脸标识串
	// （fp 开头的 35 位字符串）；当面付场景必选。
	AuthCode *string `json:"auth_code,omitempty"`
	// AuthNo 是资金预授权单号；支付宝预授权与新当面资金授权场景必填。
	AuthNo *string `json:"auth_no,omitempty"`
	// AuthConfirmMode 是预授权确认模式。
	AuthConfirmMode *string `json:"auth_confirm_mode,omitempty"`
	// AgreementParams 是周期扣款或代扣场景的协议参数。
	AgreementParams *AgreementParams `json:"agreement_params,omitempty"`
	// ProductCode 是产品码，默认 FACE_TO_FACE_PAYMENT。
	ProductCode *ProductCode `json:"product_code,omitempty"`
	// SellerID 是卖家支付宝用户 ID；收款账户优先级见官方文档。
	SellerID *string `json:"seller_id,omitempty"`
	// BuyerID 是买家支付宝用户 ID；预授权场景必填，其它场景无需传入。
	BuyerID *string `json:"buyer_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// Body 是订单附加信息，将在异步通知、对账单中原样返回。
	Body *string `json:"body,omitempty"`
	// GoodsDetail 是订单包含的商品列表，上限 150 条。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
	// OperatorID 是商户操作员编号。
	OperatorID *string `json:"operator_id,omitempty"`
	// StoreID 是商户门店编号。
	StoreID *string `json:"store_id,omitempty"`
	// TerminalID 是商户机具终端编号。
	TerminalID *string `json:"terminal_id,omitempty"`
	// TimeExpire 是订单绝对超时时间（yyyy-MM-dd HH:mm:ss）；
	// 与 TimeoutExpress 同传时优先使用本字段。
	TimeExpire *string `json:"time_expire,omitempty"`
	// TimeoutExpress 是订单相对超时时间（1m～15d）。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// RoyaltyInfo 描述分账信息。
	RoyaltyInfo *RoyaltyInfo `json:"royalty_info,omitempty"`
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams `json:"extend_params,omitempty"`
	// BusinessParams 是与支付宝约定的直传业务信息。
	BusinessParams *BusinessParams `json:"business_params,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []PayQueryOption `json:"query_options,omitempty"`
	// AdvancePaymentType 表示支付模式类型,若值为ENJOY_PAY_V2表示当前交易允许走先享后付2.0垫资。
	AdvancePaymentType *string `json:"advance_payment_type,omitempty"`
	// AlipayStoreID 表示支付宝店铺编号。
	AlipayStoreID *string `json:"alipay_store_id,omitempty"`
	// BkAgentReqInfo 是收单机构上送的间联交易信息。
	BkAgentReqInfo *BkAgentReqInfo `json:"bkagent_req_info,omitempty"`
	// DisablePayChannels 表示禁用支付渠道。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// DiscountableAmount 表示可打折金额。
	DiscountableAmount *string `json:"discountable_amount,omitempty"`
	// ExtUserInfo 是指定买家的身份信息。
	ExtUserInfo *ExtUserInfo `json:"ext_user_info,omitempty"`
	// IsAsyncPay 表示是否异步支付，传入true时，表明本次期望走异步支付，会先将支付请求受理下来，再异步推进。
	IsAsyncPay *bool `json:"is_async_pay,omitempty"`
	// MerchantOrderNo 表示商户的原始订单号。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// PassbackParams 表示公用回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
	// PayParams 是异步支付与重试参数。
	PayParams *PayParams `json:"pay_params,omitempty"`
	// PromoParams 是优惠业务参数。
	PromoParams *PromoParams `json:"promo_params,omitempty"`
	// RequestOrgPID 表示收单机构(例如银行）的标识，填写该机构在支付宝的pid。
	RequestOrgPID *string `json:"request_org_pid,omitempty"`
	// SettleCurrency 是结算币种。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
	// TerminalParams 表示IOT设备信息。
	TerminalParams *string `json:"terminal_params,omitempty"`
	// TransCurrency 表示标价币种,  total_amount 对应的币种单位。
	TransCurrency *string `json:"trans_currency,omitempty"`
	// UndiscountableAmount 表示不可打折金额。
	UndiscountableAmount *string `json:"undiscountable_amount,omitempty"`
}

// CreateRequest 是 alipay.trade.trade.create（统一收单交易创建）的请求。
// 买家确认后完成支付；out_trade_no、total_amount、subject 与买家标识
// 为必选（买家标识三选一）。
type CreateRequest struct {
	// NotifyURL 是本次请求接收异步通知的地址。
	NotifyURL *string `json:"notify_url,omitempty"`
	// OutTradeNo 是商户订单号：64 字符以内，仅支持字母、数字、下划线，
	// 且需保证在商户端不重复。
	OutTradeNo string `json:"out_trade_no"`
	// TotalAmount 是订单总金额，两位小数，取值范围 [0.01,100000000]。
	TotalAmount string `json:"total_amount"`
	// Subject 是订单标题；不可使用 / = & 等特殊字符。
	Subject string `json:"subject"`
	// BuyerID 是买家的支付宝用户 ID（与 BuyerOpenID、BuyerLogonID 三选一）。
	BuyerID *string `json:"buyer_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// BuyerLogonID 是买家支付宝登录账号。
	BuyerLogonID *string `json:"buyer_logon_id,omitempty"`
	// ProductCode 是产品码。
	ProductCode *ProductCode `json:"product_code,omitempty"`
	// SellerID 是卖家支付宝用户 ID；收款账户优先级见官方文档。
	SellerID *string `json:"seller_id,omitempty"`
	// Body 是订单附加信息，将在异步通知、对账单中原样返回。
	Body *string `json:"body,omitempty"`
	// GoodsDetail 是订单包含的商品列表，上限 150 条。
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
	// TimeExpire 是订单绝对超时时间（yyyy-MM-dd HH:mm:ss）。
	TimeExpire *string `json:"time_expire,omitempty"`
	// TimeoutExpress 是订单相对超时时间（1m～15d）。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo `json:"settle_info,omitempty"`
	// RoyaltyInfo 描述分账信息。
	RoyaltyInfo *RoyaltyInfo `json:"royalty_info,omitempty"`
	// AgreementSignParams 是支付并签约场景的签约参数；字段与必填
	// 关系按官方商家扣款场景文档核验（app.pay 与 create 复用同一模型）。
	AgreementSignParams *SignParams `json:"agreement_sign_params,omitempty"`
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams `json:"extend_params,omitempty"`
	// BusinessParams 是与支付宝约定的直传业务信息。
	BusinessParams *BusinessParams `json:"business_params,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []PayQueryOption `json:"query_options,omitempty"`
	// MerchantOrderNo 是商户的原始订单号。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// StoreID 是商户门店编号。
	StoreID *string `json:"store_id,omitempty"`
	// OperatorID 是商户操作员编号。
	OperatorID *string `json:"operator_id,omitempty"`
	// TerminalID 是商户机具终端编号。
	TerminalID *string `json:"terminal_id,omitempty"`
	// DisablePayChannels 是禁用渠道列表（逗号分割），与 EnablePayChannels 互斥。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// EnablePayChannels 是指定支付渠道列表（逗号分割）。
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// DiscountableAmount 是可打折金额；
	// 与 TotalAmount 同传时须满足 TotalAmount = DiscountableAmount + UndiscountableAmount。
	DiscountableAmount *string `json:"discountable_amount,omitempty"`
	// UndiscountableAmount 是不可打折金额。
	UndiscountableAmount *string `json:"undiscountable_amount,omitempty"`
	// AlipayStoreID 表示支付宝店铺编号。
	AlipayStoreID *string `json:"alipay_store_id,omitempty"`
	// BkAgentReqInfo 是收单机构上送的间联交易信息。
	BkAgentReqInfo *BkAgentReqInfo `json:"bkagent_req_info,omitempty"`
	// ExtUserInfo 是指定买家的身份信息。
	ExtUserInfo *ExtUserInfo `json:"ext_user_info,omitempty"`
	// LogisticsDetail 是订单物流信息。
	LogisticsDetail *LogisticsDetail `json:"logistics_detail,omitempty"`
	// OpAppID 是小程序支付必传的实际经营主体小程序应用 ID。
	OpAppID *string `json:"op_app_id,omitempty"`
	// OpBuyerOpenID 表示买家支付宝用户唯一标识（商户实际经营主体的小程序应用关联的买家open_id）。
	OpBuyerOpenID *string `json:"op_buyer_open_id,omitempty"`
	// PassbackParams 表示公用回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
	// ReceiverAddressInfo 是收货人和收货地址。
	ReceiverAddressInfo *ReceiverAddressInfo `json:"receiver_address_info,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
}

// CancelRequest 是 alipay.trade.cancel（统一收单交易撤销）的请求。
// trade_no 与 out_trade_no 二选一且不能同时为空。
type CancelRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是原支付请求的商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

func (r PayRequest) validate() error {
	if err := validatePaymentOrderNo(r.OutTradeNo); err != nil {
		return err
	}
	if err := tradevalidate.OrderAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if err := validateNotifyEcho(r.Subject, r.Body, r.PassbackParams); err != nil {
		return err
	}
	if r.Scene != nil && !r.Scene.IsKnown() {
		return &core.ValidationError{Field: "scene",
			Reason: "scene must be one of the documented values"}
	}
	if (r.Scene == nil || *r.Scene == PaySceneBarCode || *r.Scene == PaySceneSecurityCode) &&
		r.AuthCode == nil && r.AgreementParams == nil && r.AuthNo == nil {
		return &core.ValidationError{Field: "auth_code",
			Reason: "bar_code/security_code scenes require auth_code; withhold scenarios require agreement_params or auth_no"}
	}
	if r.AgreementParams != nil && r.AgreementParams.AgreementNo == "" {
		return &core.ValidationError{Field: "agreement_params.agreement_no",
			Reason: "agreement_no is required"}
	}
	for i, goods := range r.GoodsDetail {
		if err := validateGoodsDetail(goods, i); err != nil {
			return err
		}
	}
	return nil
}

func (r CreateRequest) validate() error {
	if err := validatePaymentOrderNo(r.OutTradeNo); err != nil {
		return err
	}
	if err := tradevalidate.OrderAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if err := validateNotifyEcho(r.Subject, r.Body, r.PassbackParams); err != nil {
		return err
	}
	if r.OpBuyerOpenID != nil && (r.OpAppID == nil || *r.OpAppID == "") {
		return &core.ValidationError{Field: "op_app_id", Reason: "op_app_id is required for a mini program operator's buyer"}
	}
	if r.BuyerID == nil && r.BuyerOpenID == nil && r.BuyerLogonID == nil {
		return &core.ValidationError{Field: "buyer_id",
			Reason: "one of buyer_id, buyer_open_id or buyer_logon_id is required"}
	}
	if r.DiscountableAmount != nil {
		if err := tradevalidate.OrderAmount(*r.DiscountableAmount, "discountable_amount"); err != nil {
			return err
		}
	}
	if r.UndiscountableAmount != nil {
		if err := tradevalidate.OrderAmount(*r.UndiscountableAmount, "undiscountable_amount"); err != nil {
			return err
		}
	}
	for i, goods := range r.GoodsDetail {
		if err := validateGoodsDetail(goods, i); err != nil {
			return err
		}
	}
	return nil
}

func (r CancelRequest) validate() error {
	return tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

// PagePayRequest 是电脑网站/手机网站支付的业务参数。
// out_trade_no、total_amount、subject 与产品码为必选；
// 产品码按场景选择（电脑网站 FAST_INSTANT_TRADE_PAY，
// 手机网站 QUICK_WAP_WAY）。
type PagePayRequest struct {
	// OutTradeNo 是商户订单号：64 字符以内，仅支持字母、数字、下划线，
	// 且需保证在商户端不重复。
	OutTradeNo string
	// TotalAmount 是订单总金额，两位小数，取值范围 [0.01,100000000]。
	TotalAmount string
	// Subject 是订单标题；不可使用 / = & 等特殊字符。
	Subject string
	// ProductCode 是产品码：电脑网站支付 FAST_INSTANT_TRADE_PAY，
	// 手机网站支付 QUICK_WAP_WAY。
	ProductCode ProductCode
	// Body 是订单附加信息，将在异步通知、对账单中原样返回。
	Body *string
	// GoodsDetail 是订单包含的商品列表。
	GoodsDetail []GoodsDetail
	// TimeExpire 是订单绝对超时时间（yyyy-MM-dd HH:mm:ss）。
	TimeExpire *string
	// TimeoutExpress 是订单相对超时时间（1m～15d）。
	TimeoutExpress *string
	// QRPayMode 是扫码支付方式（电脑网站支付）：0/1/2/3/4。
	QRPayMode *QRPayMode
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams
	// BusinessParams 是与支付宝约定的 JSON 文本。
	BusinessParams *string
	// ReturnURL 是支付成功后的页面跳转地址；不是支付完成的可信证据。
	ReturnURL *string
	// PassbackParams 是手机网站支付已 URL 编码的公用回传参数，供异步通知回传。
	PassbackParams *string
	// QRCodeWidth 是自定义二维码宽度，仅电脑网站支付的模式 4 使用。
	QRCodeWidth *int64
	// QuitURL 是用户中途退出手机网站支付时的返回地址。
	QuitURL *string
	// ExtUserInfo 是指定买家信息。
	ExtUserInfo *ExtUserInfo
	// SubMerchant 是电脑网站支付的二级商户信息。
	SubMerchant *SubMerchant
	// InvoiceInfo 是电脑网站支付的开票信息。
	InvoiceInfo *InvoiceInfo
	// StoreID 是电脑网站支付的商户门店编号。
	StoreID *string
	// RoyaltyInfo 是电脑网站支付的分账信息。
	RoyaltyInfo *RoyaltyInfo
}

// AppPayRequest 是 App 支付的业务参数。
// out_trade_no、total_amount、subject 必填。
// 支付并签约场景传入 AgreementSignParams。
type AppPayRequest struct {
	// OutTradeNo 是商户订单号。
	OutTradeNo string
	// TotalAmount 是订单总金额。
	TotalAmount string
	// Subject 是订单标题。
	Subject string
	// ProductCode 默认为 QUICK_MSECURITY_PAY，周期扣款使用 CYCLE_PAY_AUTH。
	ProductCode ProductCode
	// Body 是订单附加信息。
	Body *string
	// GoodsDetail 是订单包含的商品列表。
	GoodsDetail []GoodsDetail
	// TimeExpire 是订单绝对超时时间。
	TimeExpire *string
	// TimeoutExpress 是订单相对超时时间。
	TimeoutExpress *string
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams
	// AgreementSignParams 是支付并签约场景的签约参数。
	AgreementSignParams *SignParams
	// PassbackParams 是已 URL 编码的公用回传参数。
	PassbackParams *string
	// ExtUserInfo 是外部指定的买家信息。
	ExtUserInfo *ExtUserInfo
}

// artifactParams 将请求模型编码为产物业务参数：
// biz_content 为 JSON 请求体，return_url 为直接参数。
func (r PagePayRequest) artifactParams() (url.Values, error) {
	biz := pageBizContent{
		OutTradeNo:     r.OutTradeNo,
		TotalAmount:    r.TotalAmount,
		Subject:        r.Subject,
		ProductCode:    string(r.ProductCode),
		Body:           r.Body,
		GoodsDetail:    r.GoodsDetail,
		TimeExpire:     r.TimeExpire,
		TimeoutExpress: r.TimeoutExpress,
		QRPayMode:      r.QRPayMode,
		SettleInfo:     r.SettleInfo,
		ExtendParams:   r.ExtendParams,
		BusinessParams: r.BusinessParams,
		PassbackParams: r.PassbackParams,
		QRCodeWidth:    r.QRCodeWidth,
		QuitURL:        r.QuitURL,
		ExtUserInfo:    r.ExtUserInfo,
		SubMerchant:    r.SubMerchant,
		InvoiceInfo:    r.InvoiceInfo,
		StoreID:        r.StoreID,
		RoyaltyInfo:    r.RoyaltyInfo,
	}
	return encodeArtifactParams(biz, r.ReturnURL)
}

func (r AppPayRequest) artifactParams() (url.Values, error) {
	biz := appBizContent{
		OutTradeNo:          r.OutTradeNo,
		TotalAmount:         r.TotalAmount,
		Subject:             r.Subject,
		ProductCode:         r.ProductCode,
		Body:                r.Body,
		GoodsDetail:         r.GoodsDetail,
		TimeExpire:          r.TimeExpire,
		TimeoutExpress:      r.TimeoutExpress,
		ExtendParams:        r.ExtendParams,
		AgreementSignParams: r.AgreementSignParams,
		ExtUserInfo:         r.ExtUserInfo,
		PassbackParams:      r.PassbackParams,
	}
	return encodeArtifactParams(biz, nil)
}

func (r PagePayRequest) validate(method string) error {
	if err := validatePaymentOrderNo(r.OutTradeNo); err != nil {
		return err
	}
	if err := tradevalidate.OrderAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if err := validateNotifyEcho(r.Subject, r.Body, r.PassbackParams); err != nil {
		return err
	}
	if method == methodWapPay {
		if r.ProductCode != ProductCodeQuickWapWay {
			return &core.ValidationError{Field: "product_code", Reason: "wap.pay requires QUICK_WAP_WAY"}
		}
		if r.QRPayMode != nil || r.QRCodeWidth != nil || r.InvoiceInfo != nil || r.SubMerchant != nil || r.StoreID != nil || r.RoyaltyInfo != nil {
			return &core.ValidationError{Field: "biz_content", Reason: "PC payment fields cannot be sent to wap.pay"}
		}
	} else {
		if r.ProductCode != ProductCodeFastInstantTradePay {
			return &core.ValidationError{Field: "product_code", Reason: "page.pay requires FAST_INSTANT_TRADE_PAY"}
		}
		if r.QuitURL != nil {
			return &core.ValidationError{Field: "quit_url", Reason: "quit_url is only supported by wap.pay"}
		}
		if r.PassbackParams != nil {
			return &core.ValidationError{Field: "passback_params", Reason: "use wap.pay for this page payment parameter"}
		}
		if r.QRPayMode != nil {
			switch *r.QRPayMode {
			case QRPaySimple, QRPayEmbedded, QRPayRedirect, QRPayMini, QRPayCustomWidth:
			default:
				return &core.ValidationError{Field: "qr_pay_mode", Reason: "unknown QR payment mode"}
			}
		}
		if r.QRCodeWidth != nil && (r.QRPayMode == nil || *r.QRPayMode != QRPayCustomWidth || *r.QRCodeWidth <= 0) {
			return &core.ValidationError{Field: "qrcode_width", Reason: "positive width and qr_pay_mode 4 are required"}
		}
		if info := r.InvoiceInfo; info != nil && (info.Details == "" || info.KeyInfo.InvoiceMerchantName == "" || info.KeyInfo.TaxNum == "") {
			return &core.ValidationError{Field: "invoice_info", Reason: "invoice details, merchant name and tax number are required"}
		}
		if sub := r.SubMerchant; sub != nil && (sub.MerchantID == nil || *sub.MerchantID == "") {
			return &core.ValidationError{Field: "sub_merchant.merchant_id", Reason: "merchant ID is required for a sub-merchant"}
		}
	}
	return validateArtifactDetails(method, r.GoodsDetail, r.ExtendParams, r.ExtUserInfo)
}

func (r AppPayRequest) validate() error {
	if err := validatePaymentOrderNo(r.OutTradeNo); err != nil {
		return err
	}
	if err := tradevalidate.OrderAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if err := validateNotifyEcho(r.Subject, r.Body, r.PassbackParams); err != nil {
		return err
	}
	switch r.ProductCode {
	case "", ProductCodeQuickMSecurityPay:
	case ProductCodeCyclePayAuth:
		if r.AgreementSignParams == nil {
			return &core.ValidationError{Field: "agreement_sign_params", Reason: "required for cycle payments"}
		}
	default:
		return &core.ValidationError{Field: "product_code", Reason: "unsupported app.pay product code"}
	}
	if err := r.AgreementSignParams.validateAppAgreement(); err != nil {
		return err
	}
	return validateArtifactDetails(methodAppPay, r.GoodsDetail, r.ExtendParams, r.ExtUserInfo)
}

func validateArtifactDetails(method string, items []GoodsDetail, extend *ExtendParams, buyer *ExtUserInfo) error {
	if extend != nil && (extend.FqNum != nil || extend.FqChannels != nil || extend.FqSellerPercent != nil) {
		return &core.ValidationError{Field: "extend_params", Reason: "fq_num, fq_channels and fq_seller_percent are REST parameters"}
	}
	if buyer != nil && buyer.FixBuyer != nil {
		return &core.ValidationError{Field: "ext_user_info.fix_buyer", Reason: "fix_buyer is a REST parameter"}
	}
	for i, goods := range items {
		if method != methodWapPay && (goods.Body != nil || goods.GoodsEANID != nil) {
			return &core.ValidationError{Field: "goods_detail", Reason: "only wap.pay supports body and goods_ean_id on payment artifacts"}
		}
		if err := validateGoodsDetail(goods, i); err != nil {
			return err
		}
	}
	return nil
}

func (r QueryRequest) validate() error {
	return tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r CloseRequest) validate() error {
	return tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (p *SignParams) validateAppAgreement() error {
	if p == nil {
		return nil
	}
	switch {
	case p.AllowHuazhiDegrade != nil:
		return &core.ValidationError{Field: "agreement_sign_params.allow_huazhi_degrade", Reason: "this parameter is not supported by app.pay"}
	case p.ProductCode == nil || *p.ProductCode == "":
		return &core.ValidationError{Field: "agreement_sign_params.product_code", Reason: "product_code is required"}
	case p.PersonalProductCode == nil || *p.PersonalProductCode == "":
		return &core.ValidationError{Field: "agreement_sign_params.personal_product_code", Reason: "personal_product_code is required"}
	case p.SignScene == nil || *p.SignScene == "":
		return &core.ValidationError{Field: "agreement_sign_params.sign_scene", Reason: "sign_scene is required"}
	case p.AccessParams == nil || p.AccessParams.Channel == nil || *p.AccessParams.Channel == "":
		return &core.ValidationError{Field: "agreement_sign_params.access_params.channel", Reason: "signing channel is required"}
	case *p.PersonalProductCode == "CYCLE_PAY_AUTH_P" && p.PeriodRuleParams == nil:
		return &core.ValidationError{Field: "agreement_sign_params.period_rule_params", Reason: "period rules are required for cycle payments"}
	}
	return p.PeriodRuleParams.validateAppPeriod(*p.PersonalProductCode == "CYCLE_PAY_AUTH_P")
}

func (p *PeriodRuleParams) validateAppPeriod(cycle bool) error {
	if p == nil {
		return nil
	}
	const field = "agreement_sign_params.period_rule_params."
	if p.IsTrialMode != nil && *p.IsTrialMode != "Y" && *p.IsTrialMode != "N" {
		return &core.ValidationError{Field: field + "is_trial_mode", Reason: "expected Y or N"}
	}
	if p.SingleAmount == nil {
		return &core.ValidationError{Field: field + "single_amount", Reason: "single amount is required"}
	}
	if err := tradevalidate.Amount(*p.SingleAmount, field+"single_amount"); err != nil {
		return err
	}
	if !cycle {
		return nil
	}
	if p.PeriodType == nil || (*p.PeriodType != "DAY" && *p.PeriodType != "MONTH") {
		return &core.ValidationError{Field: field + "period_type", Reason: "cycle payments require DAY or MONTH"}
	}
	if p.Period == nil || *p.Period <= 0 {
		return &core.ValidationError{Field: field + "period", Reason: "cycle payments require a positive period"}
	}
	if p.ExecuteTime == nil {
		return &core.ValidationError{Field: field + "execute_time", Reason: "cycle payments require the first payment date"}
	}
	date, err := time.Parse(time.DateOnly, *p.ExecuteTime)
	if err != nil || (*p.PeriodType == "MONTH" && date.Day() > 28) {
		return &core.ValidationError{Field: field + "execute_time", Reason: "expected yyyy-MM-dd, on or before day 28 for a monthly cycle"}
	}
	return nil
}
