package refund

import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

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
	RefundRoyaltyParameters []trade.RoyaltyDetailInfo `json:"refund_royalty_parameters,omitempty"`
	// QueryOptions 定制同步额外返回的信息字段。
	QueryOptions []RefundQueryOption `json:"query_options,omitempty"`
	// RelatedSettleConfirmNo 用于账期交易确认结算后退款。
	RelatedSettleConfirmNo *string `json:"related_settle_confirm_no,omitempty"`
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

// RefundApplyRequest 是 alipay.trade.refund.apply（异步退款受理）的请求。
// out_request_no 标识一次退款请求，同一笔交易多次退款需保证唯一；
// 请求成功仅代表受理，结果以 RefundStatus 与退款查询为准。
type RefundApplyRequest struct {
	// NotifyURL 是本次请求接收异步通知的地址。
	NotifyURL *string `json:"notify_url,omitempty"`
	// TradeNo 是支付宝交易号（与 OutTradeNo 二选一）。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// OutRequestNo 是退款请求号；同一笔交易多次退款需保证唯一，
	// 部分退款必传。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RefundAmount 是退款金额，不能大于订单金额。
	RefundAmount string `json:"refund_amount"`
	// RefundReason 是退款原因说明。
	RefundReason *string `json:"refund_reason,omitempty"`
	// OperatorID 是商户操作员编号。
	OperatorID *string `json:"operator_id,omitempty"`
	// StoreID 是商户门店编号。
	StoreID *string `json:"store_id,omitempty"`
	// TerminalID 是商户终端编号。
	TerminalID *string `json:"terminal_id,omitempty"`
	// RefundFundDetails 是退款资金明细；仅 SETTLE_OFFLINE_GROUP
	// 产品支持。
	RefundFundDetails []RefundFundDetail `json:"refund_fund_details,omitempty"`
}

func (r RefundApplyRequest) validate() error {
	if err := tradevalidate.Amount(r.RefundAmount, "refund_amount"); err != nil {
		return err
	}
	if err := tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no"); err != nil {
		return err
	}
	if r.OutRequestNo == nil || *r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no",
			Reason: "out_request_no is required for refund apply"}
	}
	return nil
}

func (r QueryRefundRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r RefundRequest) validate() error {
	if err := tradevalidate.Amount(r.RefundAmount, "refund_amount"); err != nil {
		return err
	}
	return tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}
