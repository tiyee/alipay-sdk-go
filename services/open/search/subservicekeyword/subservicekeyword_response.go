package subservicekeyword

import (
	"github.com/tiyee/alipay-sdk-go/services/open/search"
)

// SearchSubServiceKeywordApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordApplyResponse struct {
	// ApplyNo 是审核工单id，唯一，提报服务关键词，提报服务关键词返回该id。
	ApplyNo *string `json:"apply_no,omitempty"`
}

// SearchSubServiceKeywordBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordBatchQueryResponse struct {
	// PageNumber 是第几页。
	PageNumber *int64 `json:"page_number,omitempty"`
	// PageSize 是每页条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// SubserviceKeywordInfo 是关键词明细。
	SubserviceKeywordInfo *[]SubServiceKeyWordInfo `json:"subservice_keyword_info,omitempty"`
	// TotalCount 是总条数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// TotalPageCount 是总页数。
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}

// SearchSubServiceKeywordDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordDeleteResponse struct {
}

// SearchSubServiceKeywordQueryStatusResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordQueryStatusResponse struct {
	// KeyWords 是关键词工单审核状态返回值。
	KeyWords *[]search.KeyWordInfo `json:"key_words,omitempty"`
}
