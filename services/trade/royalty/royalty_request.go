package royalty

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// RoyaltyRateQueryRequest 是 alipay.trade.royalty.rate.query
// （分账比例查询，V3 GET）的请求。
type RoyaltyRateQueryRequest struct {
	// OutRequestNo 是外部请求号：32 字符以内，仅可包含字母、数字、
	// 下划线，需保证在商户端不重复。
	OutRequestNo string `query:"out_request_no"`
}

// RoyaltyRelationBindRequest 是 alipay.trade.royalty.relation.bind
// （分账关系绑定）的请求。
type RoyaltyRelationBindRequest struct {
	// OutRequestNo 是外部请求号：32 字符以内，仅可包含字母、数字、
	// 下划线，需保证在商户端不重复。
	OutRequestNo string `json:"out_request_no"`
	// ReceiverList 是分账接收方列表，单次最多 20 个。
	ReceiverList []RoyaltyEntity `json:"receiver_list"`
}

// RoyaltyRelationUnbindRequest 是 alipay.trade.royalty.relation.unbind
// （分账关系解绑）的请求。
type RoyaltyRelationUnbindRequest struct {
	// OutRequestNo 是外部请求号。
	OutRequestNo string `json:"out_request_no"`
	// ReceiverList 是分账接收方列表，单次最多 20 个。
	ReceiverList []RoyaltyEntity `json:"receiver_list"`
}

// RoyaltyRelationBatchQueryRequest 是 alipay.trade.royalty.relation.batchquery
// （分账关系查询）的请求。
type RoyaltyRelationBatchQueryRequest struct {
	// OutRequestNo 是外部请求号。
	OutRequestNo string `json:"out_request_no"`
	// PageNum 是页码，起始页为 1；不填默认为 1。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录数，取值范围 (0,100]；不填默认为 20。
	PageSize *int64 `json:"page_size,omitempty"`
}

func (r RoyaltyRateQueryRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return nil
}

func (r RoyaltyRelationBindRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return validateReceiverList(r.ReceiverList)
}

func (r RoyaltyRelationUnbindRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return validateReceiverList(r.ReceiverList)
}

func (r RoyaltyRelationBatchQueryRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if r.PageSize != nil && (*r.PageSize <= 0 || *r.PageSize > 100) {
		return &core.ValidationError{Field: "page_size",
			Reason: "page_size must be within (0,100]"}
	}
	return nil
}
