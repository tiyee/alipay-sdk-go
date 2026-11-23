package orderinfo

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// OrderInfoSyncRequest 是 alipay.trade.orderinfo.sync（订单信息同步）
// 的请求。out_request_no 由支付宝做幂等控制：相同请求号再次请求
// 不会更新订单信息。
type OrderInfoSyncRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no"`
	// OutRequestNo 是外部请求号，商家自定义且保证唯一。
	OutRequestNo string `json:"out_request_no"`
	// BizType 是业务类型，具体值与支付宝约定（如 CREDIT_AUTH、
	// CREDIT_DEDUCT）。
	BizType string `json:"biz_type"`
	// OrderBizInfo 是同步信息内容，格式与支付宝约定（JSON 文本）。
	OrderBizInfo *string `json:"order_biz_info,omitempty"`
	// OrigRequestNo 是原始业务请求单号（如履约对象退款的退款请求号）。
	OrigRequestNo *string `json:"orig_request_no,omitempty"`
}

func (r OrderInfoSyncRequest) validate() error {
	if r.TradeNo == "" {
		return &core.ValidationError{Field: "trade_no", Reason: "trade_no is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if r.BizType == "" {
		return &core.ValidationError{Field: "biz_type", Reason: "biz_type is required"}
	}
	return nil
}
