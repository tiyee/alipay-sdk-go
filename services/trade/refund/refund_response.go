package refund

import (
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

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
	RefundDetailItemList []trade.TradeFundBill `json:"refund_detail_item_list,omitempty"`
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
	RefundVoucherDetailList []trade.VoucherDetail `json:"refund_voucher_detail_list,omitempty"`
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
	RefundDetailItemList []trade.TradeFundBill `json:"refund_detail_item_list,omitempty"`
	// SendBackFee 是本次商户实际退回金额，需 query_options 指定。
	SendBackFee *string `json:"send_back_fee,omitempty"`
	// DepositBackInfo 是银行卡冲退信息，需 query_options 指定 deposit_back_info。
	DepositBackInfo *DepositBackInfo `json:"deposit_back_info,omitempty"`
	// RefundVoucherDetailList 是本交易支付时使用的所有优惠券信息，
	// 需 query_options 指定 refund_voucher_detail_list。
	RefundVoucherDetailList []trade.VoucherDetail `json:"refund_voucher_detail_list,omitempty"`
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

// RefundApplyResponse 是 alipay.trade.refund.apply 的业务响应。
// RefundStatus：REFUND_PROCESSING 处理中、REFUND_SUCCESS 成功、
// REFUND_FAIL 失败；处理中不是终态，需继续查询。
type RefundApplyResponse struct {
	// TradeNo 是支付宝交易号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是商户订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// OutRequestNo 是本笔退款对应的退款请求号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RefundAmount 是本次退款请求对应的退款金额。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundStatus 是退款受理状态。
	RefundStatus *RefundApplyStatus `json:"refund_status,omitempty"`
}
