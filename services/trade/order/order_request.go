package order

import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// OrderPayRequest 是 alipay.trade.order.pay（统一收单交易订单支付）的请求，
// 用于对已存在订单发起履约支付。trade_no 必选；
// out_request_no 用于幂等控制。
type OrderPayRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no"`
	// TotalAmount 是订单总金额；用于校验与已存交易中的金额是否一致。
	TotalAmount string `json:"total_amount"`
	// OutRequestNo 是商户请求号，标识一次请求的唯一 id，用于幂等控制。
	OutRequestNo string `json:"out_request_no"`
	// BuyerID 是买家的支付宝用户 id；用于校验与已存交易的买家是否一致。
	BuyerID *string `json:"buyer_id,omitempty"`
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string `json:"buyer_open_id,omitempty"`
	// BuyerPayDetail 是买家支付明细。
	BuyerPayDetail []BuyerPayDetail `json:"buyer_pay_detail,omitempty"`
	// FulfillmentAmount 是本次履约支付金额；履约支付场景传入。
	FulfillmentAmount *string `json:"fulfillment_amount,omitempty"`
	// IsAsyncPay 标识是否走异步支付。
	IsAsyncPay *bool `json:"is_async_pay,omitempty"`
	// OrderPayMode 是订单支付模式；特殊场景与支付宝约定。
	OrderPayMode *string `json:"order_pay_mode,omitempty"`
	// ProductCode 是销售产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// AdvancePaymentType 是垫资支付模式；特殊场景与支付宝约定。
	AdvancePaymentType *string `json:"advance_payment_type,omitempty"`
}

func (r OrderPayRequest) validate() error {
	if r.TradeNo == "" {
		return &core.ValidationError{Field: "trade_no", Reason: "trade_no is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return tradevalidate.OrderAmount(r.TotalAmount, "total_amount")
}

// OrderSettleRequest 是 alipay.trade.order.settle（统一收单交易结算）的请求。
// trade_no、out_request_no 必选；单独调用分账完结时可不传分账明细。
type OrderSettleRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no"`
	// OutRequestNo 是结算请求流水号：64 字符以内，仅可包含字母、数字、
	// 下划线，需保证在商户端不重复。
	OutRequestNo string `json:"out_request_no"`
	// RoyaltyParameters 是分账明细信息；单独调用分账完结时可不传，
	// 其他场景必传。
	RoyaltyParameters []trade.RoyaltyDetailInfo `json:"royalty_parameters,omitempty"`
	// RoyaltyMode 是分账模式：sync（同步执行）或 async（异步执行），
	// 默认同步。
	RoyaltyMode *RoyaltyMode `json:"royalty_mode,omitempty"`
	// OperatorID 是操作员 ID。
	OperatorID *string `json:"operator_id,omitempty"`
	// ExtendParams 是分账扩展参数（完结标记等）。
	ExtendParams *OrderSettleExtendParams `json:"extend_params,omitempty"`
}

// OrderSettleQueryRequest 是 alipay.trade.order.settle.query
// （交易分账查询，V3 GET）的请求：settle_no 或（out_request_no +
// trade_no）二选一。
type OrderSettleQueryRequest struct {
	// SettleNo 是支付宝分账请求单号；传入后无需再传外部请求号与交易号。
	SettleNo *string `query:"settle_no,omitempty"`
	// OutRequestNo 是分账接口指定的外部请求号；需与 TradeNo 一起传入。
	OutRequestNo *string `query:"out_request_no,omitempty"`
	// TradeNo 是支付宝交易号；需与 OutRequestNo 一起传入。
	TradeNo *string `query:"trade_no,omitempty"`
}

func (r OrderSettleRequest) validate() error {
	if r.TradeNo == "" {
		return &core.ValidationError{Field: "trade_no", Reason: "trade_no is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if r.RoyaltyMode != nil && !r.RoyaltyMode.IsKnown() {
		return &core.ValidationError{Field: "royalty_mode",
			Reason: "royalty_mode must be sync or async"}
	}
	return nil
}

func (r OrderSettleQueryRequest) validate() error {
	if r.SettleNo != nil && *r.SettleNo != "" {
		return nil
	}
	if r.OutRequestNo != nil && *r.OutRequestNo != "" && r.TradeNo != nil && *r.TradeNo != "" {
		return nil
	}
	return &core.ValidationError{Field: "settle_no",
		Reason: "settle_no, or out_request_no together with trade_no, is required"}
}
