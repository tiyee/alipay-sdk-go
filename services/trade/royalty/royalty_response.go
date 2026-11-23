package royalty

import (
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// RoyaltyRateQueryResponse 是分账比例查询的业务响应。
type RoyaltyRateQueryResponse struct {
	// UserID 是当前商户 userId。
	UserID *string `json:"user_id,omitempty"`
	// MaxRatio 是最大分账比例，百分比整数，取值范围 0～100。
	MaxRatio *int64 `json:"max_ratio,omitempty"`
}

// RoyaltyRelationBindResponse 是分账关系绑定的业务响应。
type RoyaltyRelationBindResponse struct {
	// ResultCode 是业务结果码：SUCCESS 或 FAIL。
	ResultCode *trade.BatchResultCode `json:"result_code,omitempty"`
}

// RoyaltyRelationUnbindResponse 是分账关系解绑的业务响应。
type RoyaltyRelationUnbindResponse struct {
	// ResultCode 是业务结果码：SUCCESS 或 FAIL。
	ResultCode *trade.BatchResultCode `json:"result_code,omitempty"`
}

// RoyaltyRelationBatchQueryResponse 是分账关系查询的业务响应。
type RoyaltyRelationBatchQueryResponse struct {
	// ReceiverList 是分账收款方列表。
	ReceiverList []RoyaltyEntity `json:"receiver_list,omitempty"`
	// TotalPageNum 是总页数。
	TotalPageNum *int64 `json:"total_page_num,omitempty"`
	// TotalRecordNum 是分账关系记录总数。
	TotalRecordNum *int64 `json:"total_record_num,omitempty"`
	// CurrentPageNum 是当前页数。
	CurrentPageNum *int64 `json:"current_page_num,omitempty"`
	// CurrentPageSize 是当前页面大小。
	CurrentPageSize *int64 `json:"current_page_size,omitempty"`
	// ResultCode 是业务结果码：SUCCESS 或 FAIL。
	ResultCode *trade.BatchResultCode `json:"result_code,omitempty"`
}
