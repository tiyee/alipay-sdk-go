package settle

import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// SettleConfirmRequest 是 alipay.trade.settle.confirm（统一收单确认结算）
// 的请求。trade_no、out_request_no、settle_info 必选。
type SettleConfirmRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no"`
	// OutRequestNo 是确认结算请求流水号；由调用方生成并保证唯一，
	// 用于业务幂等控制。
	OutRequestNo string `json:"out_request_no"`
	// SettleInfo 是结算信息。
	SettleInfo trade.SettleInfo `json:"settle_info"`
	// ExtendParams 是确认结算扩展参数。
	ExtendParams *SettleConfirmExtendParams `json:"extend_params,omitempty"`
}

// SettleReceivablesQueryRequest 是 alipay.trade.settle.receivables.query
// （待结算查询）的请求：直付通账期模式查询单笔交易时 trade_no 必传；
// 其他场景传入商户信息与产品码。
type SettleReceivablesQueryRequest struct {
	// TradeNo 是支付宝交易号；直付通账期模式查询单笔交易时必传。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutRequestNo 是外部请求号：32 字符以内，可包含字母、数字、下划线。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// BizProduct 是收单产品码。
	BizProduct *string `json:"biz_product,omitempty"`
	// MerchantInfo 是结算主体信息。
	MerchantInfo *SettleEntity `json:"merchant_info,omitempty"`
	// QueryHisDate 是查询历史日期（yyyyMMdd，昨日起 30 日内）；
	// trade_no 不为空时该字段不生效。
	QueryHisDate *string `json:"query_his_date,omitempty"`
}

func (r SettleConfirmRequest) validate() error {
	if r.TradeNo == "" {
		return &core.ValidationError{Field: "trade_no", Reason: "trade_no is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if len(r.SettleInfo.SettleDetailInfos) == 0 {
		return &core.ValidationError{Field: "settle_info.settle_detail_infos",
			Reason: "settle_detail_infos is required"}
	}
	return nil
}

func (r SettleReceivablesQueryRequest) validate() error {
	if (r.TradeNo == nil || *r.TradeNo == "") && (r.OutRequestNo == nil || *r.OutRequestNo == "") {
		return &core.ValidationError{Field: "trade_no",
			Reason: "trade_no or out_request_no is required"}
	}
	return nil
}
