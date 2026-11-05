package batch

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// BatchTransferRequest 是 alipay.trade.batch.transfer（请款分账）的请求。
// out_request_no 与 royalty_parameters 必选；多个分账条款要求出账方
// 只能有一个。写操作：重复请求流水号返回既有单据状态。
type BatchTransferRequest struct {
	// OutRequestNo 是结算请求外部流水号：32 字符以内，可包含字母、
	// 数字、下划线，需保证在商户端不重复。
	OutRequestNo string `json:"out_request_no"`
	// RoyaltyParameters 是分账条款明细。
	RoyaltyParameters []BatchRoyaltyClause `json:"royalty_parameters"`
	// ExtendParams 是业务扩展参数。
	ExtendParams *string `json:"extend_params,omitempty"`
}

// BatchTransferQueryRequest 是 alipay.trade.batch.transfer.query
// （请款分账查询，V3 GET）的请求：settle_no 与 out_request_no
// 必填其一，同时存在时优先取 settle_no。
type BatchTransferQueryRequest struct {
	// SettleNo 是结算单据号。
	SettleNo *string `query:"settle_no,omitempty"`
	// OutRequestNo 是批量结算时传入的外部流水号。
	OutRequestNo *string `query:"out_request_no,omitempty"`
	// ExtendParams 是扩展字段。
	ExtendParams *string `query:"extend_params,omitempty"`
}

func (r BatchTransferRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if len(r.RoyaltyParameters) == 0 {
		return &core.ValidationError{Field: "royalty_parameters", Reason: "royalty_parameters is required"}
	}
	return nil
}

func (r BatchTransferQueryRequest) validate() error {
	if (r.SettleNo == nil || *r.SettleNo == "") && (r.OutRequestNo == nil || *r.OutRequestNo == "") {
		return &core.ValidationError{Field: "settle_no",
			Reason: "settle_no or out_request_no is required"}
	}
	return nil
}

// BatchSettleRequest 是 alipay.trade.batch.settle（批量结算）的请求。
// out_request_no 与 settle_clauses 必选。写操作：重复的请求流水号
// 返回既有单据状态，SDK 不自动重试。
type BatchSettleRequest struct {
	// OutRequestNo 是结算请求外部流水号：32 字符以内，可包含字母、
	// 数字、下划线，需保证在商户端不重复；重复时返回既有单据状态。
	OutRequestNo string `json:"out_request_no"`
	// SettleClauses 是结算明细条款。
	SettleClauses []SettleClause `json:"settle_clauses"`
	// BizProduct 是收单产品码。
	BizProduct *string `json:"biz_product,omitempty"`
	// SettleType 是结算方式；目前仅支持提前放款结算 quickSettlement。
	SettleType *string `json:"settle_type,omitempty"`
	// ExtendParams 是扩展参数。
	ExtendParams *string `json:"extend_params,omitempty"`
}

// BatchSettleQueryRequest 是 alipay.trade.batch.settle.query
// （批量结算查询）的请求：settle_no 与 out_request_no 必填其一，
// 同时存在时优先取 settle_no。
type BatchSettleQueryRequest struct {
	// SettleNo 是结算单据号。
	SettleNo *string `json:"settle_no,omitempty"`
	// OutRequestNo 是批量结算时传入的外部流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// ExtendParams 是扩展参数。
	ExtendParams *string `json:"extend_params,omitempty"`
}

func (r BatchSettleRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if len(r.SettleClauses) == 0 {
		return &core.ValidationError{Field: "settle_clauses", Reason: "settle_clauses is required"}
	}
	return nil
}

func (r BatchSettleQueryRequest) validate() error {
	if (r.SettleNo == nil || *r.SettleNo == "") && (r.OutRequestNo == nil || *r.OutRequestNo == "") {
		return &core.ValidationError{Field: "settle_no",
			Reason: "settle_no or out_request_no is required"}
	}
	return nil
}
