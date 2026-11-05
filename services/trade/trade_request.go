package trade

import (
	"net/url"

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
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	if err := tradevalidate.Amount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
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
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	if err := tradevalidate.Amount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if r.BuyerID == nil && r.BuyerOpenID == nil && r.BuyerLogonID == nil {
		return &core.ValidationError{Field: "buyer_id",
			Reason: "one of buyer_id, buyer_open_id or buyer_logon_id is required"}
	}
	if r.DiscountableAmount != nil {
		if err := tradevalidate.Amount(*r.DiscountableAmount, "discountable_amount"); err != nil {
			return err
		}
	}
	if r.UndiscountableAmount != nil {
		if err := tradevalidate.Amount(*r.UndiscountableAmount, "undiscountable_amount"); err != nil {
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
	// SellerID 是卖家支付宝用户 ID。
	SellerID *string
	// Body 是订单附加信息，将在异步通知、对账单中原样返回。
	Body *string
	// GoodsDetail 是订单包含的商品列表。
	GoodsDetail []GoodsDetail
	// TimeExpire 是订单绝对超时时间（yyyy-MM-dd HH:mm:ss）。
	TimeExpire *string
	// TimeoutExpress 是订单相对超时时间（1m～15d）。
	TimeoutExpress *string
	// QRPayMode 是扫码支付方式（电脑网站支付）：0/1/2/3。
	QRPayMode *string
	// QrPayTimeLimit 是二维码可扫码支付的最长时间。
	QrPayTimeLimit *string
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams
	// BusinessParams 是与支付宝约定的直传业务信息。
	BusinessParams *BusinessParams
	// ReturnURL 是支付成功后的页面跳转地址；不是支付完成的可信证据。
	ReturnURL *string
	// PassbackParams 是公用回传参数，异步通知时原样返回。
	PassbackParams *string
}

// AppPayRequest 是 App 支付的业务参数。
// out_trade_no、total_amount、subject 与产品码为必选；
// 支付并签约场景传入 SignParams。
type AppPayRequest struct {
	// OutTradeNo 是商户订单号。
	OutTradeNo string
	// TotalAmount 是订单总金额。
	TotalAmount string
	// Subject 是订单标题。
	Subject string
	// ProductCode 是产品码：App 支付为 QUICK_MSECURITY_PAY。
	ProductCode ProductCode
	// SellerID 是卖家支付宝用户 ID。
	SellerID *string
	// Body 是订单附加信息。
	Body *string
	// GoodsDetail 是订单包含的商品列表。
	GoodsDetail []GoodsDetail
	// TimeExpire 是订单绝对超时时间。
	TimeExpire *string
	// TimeoutExpress 是订单相对超时时间。
	TimeoutExpress *string
	// SettleInfo 描述结算信息。
	SettleInfo *SettleInfo
	// RoyaltyInfo 描述分账信息。
	RoyaltyInfo *RoyaltyInfo
	// ExtendParams 是业务扩展参数。
	ExtendParams *ExtendParams
	// BusinessParams 是与支付宝约定的直传业务信息。
	BusinessParams *BusinessParams
	// SignParams 是支付并签约场景的签约参数。
	SignParams *SignParams
	// PassbackParams 是公用回传参数。
	PassbackParams *string
}

// artifactParams 将请求模型编码为产物业务参数：
// biz_content 为 JSON 请求体，return_url 为直接参数。
func (r PagePayRequest) artifactParams() (url.Values, error) {
	biz := pageBizContent{
		OutTradeNo:     r.OutTradeNo,
		TotalAmount:    r.TotalAmount,
		Subject:        r.Subject,
		ProductCode:    string(r.ProductCode),
		SellerID:       r.SellerID,
		Body:           r.Body,
		GoodsDetail:    r.GoodsDetail,
		TimeExpire:     r.TimeExpire,
		TimeoutExpress: r.TimeoutExpress,
		QRPayMode:      r.QRPayMode,
		QrPayTimeLimit: r.QrPayTimeLimit,
		SettleInfo:     r.SettleInfo,
		ExtendParams:   r.ExtendParams,
		BusinessParams: r.BusinessParams,
		PassbackParams: r.PassbackParams,
	}
	return encodeArtifactParams(biz, r.ReturnURL)
}

func (r AppPayRequest) artifactParams() (url.Values, error) {
	biz := appBizContent{
		OutTradeNo:     r.OutTradeNo,
		TotalAmount:    r.TotalAmount,
		Subject:        r.Subject,
		ProductCode:    string(r.ProductCode),
		SellerID:       r.SellerID,
		Body:           r.Body,
		GoodsDetail:    r.GoodsDetail,
		TimeExpire:     r.TimeExpire,
		TimeoutExpress: r.TimeoutExpress,
		SettleInfo:     r.SettleInfo,
		RoyaltyInfo:    r.RoyaltyInfo,
		ExtendParams:   r.ExtendParams,
		BusinessParams: r.BusinessParams,
		SignParams:     r.SignParams,
		PassbackParams: r.PassbackParams,
	}
	return encodeArtifactParams(biz, nil)
}

func (r PagePayRequest) validate() error {
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	if err := tradevalidate.Amount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if r.ProductCode == "" {
		return &core.ValidationError{Field: "product_code", Reason: "product_code is required"}
	}
	for i, goods := range r.GoodsDetail {
		if err := validateGoodsDetail(goods, i); err != nil {
			return err
		}
	}
	return nil
}

func (r AppPayRequest) validate() error {
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	if err := tradevalidate.Amount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if r.ProductCode == "" {
		return &core.ValidationError{Field: "product_code", Reason: "product_code is required"}
	}
	if r.SignParams != nil && r.SignParams.ProductCode == nil {
		return &core.ValidationError{Field: "sign_params.product_code",
			Reason: "sign_params.product_code is required for the withhold scenario"}
	}
	for i, goods := range r.GoodsDetail {
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
