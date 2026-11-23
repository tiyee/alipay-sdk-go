package batch

import (
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// BatchTransferResponse 是请款分账的业务响应。
type BatchTransferResponse struct {
	// OutRequestNo 是分账请求外部流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SettleNo 是分账单据号。
	SettleNo *string `json:"settle_no,omitempty"`
	// ResultCode 是业务结果码。
	ResultCode *trade.BatchResultCode `json:"result_code,omitempty"`
	// GmtCreate 是分账受理时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
}

// BatchTransferQueryResponse 是请款分账查询的业务响应。
type BatchTransferQueryResponse struct {
	// SettleNo 是结算单据号。
	SettleNo *string `json:"settle_no,omitempty"`
	// OutRequestNo 是批量结算时传入的外部流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RoyaltyDetail 是批量结算详情。
	RoyaltyDetail []BatchRoyaltyDetail `json:"royalty_detail,omitempty"`
}

// BatchSettleResponse 是 alipay.trade.batch.settle 的业务响应。
// ResultCode 为受理结果：ACCEPT_SUCCESS/SUCCESS 之外的状态需要
// 结合查询接口确认，不自动重试。
type BatchSettleResponse struct {
	// OutRequestNo 是结算请求外部流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SettleNo 是结算单据号；仅受理成功或处理成功时有效。
	SettleNo *string `json:"settle_no,omitempty"`
	// ResultCode 是业务结果码。
	ResultCode *trade.BatchResultCode `json:"result_code,omitempty"`
	// GmtCreate 是结算受理时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
}

// BatchSettleQueryResponse 是批量结算查询的业务响应。
type BatchSettleQueryResponse struct {
	// SettleNo 是结算单据号。
	SettleNo *string `json:"settle_no,omitempty"`
	// OutRequestNo 是批量结算时传入的外部流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SettleDetail 是批量结算详情。
	SettleDetail []BatchSettleDetail `json:"settle_detail,omitempty"`
}
