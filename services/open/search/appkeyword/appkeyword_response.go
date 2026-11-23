package appkeyword

import (
	"github.com/tiyee/alipay-sdk-go/services/open/search"
)

// SearchAppKeywordApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordApplyResponse struct {
	// ApplyNo 是提报之后的后台审核工单的id。
	ApplyNo *string `json:"apply_no,omitempty"`
}

// SearchAppKeywordBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordBatchQueryResponse struct {
	// KeyWordInfo 是审核通过关键词明细。
	KeyWordInfo *[]search.KeyWordInfo `json:"key_word_info,omitempty"`
	// PageNumber 是第几页。
	PageNumber *int64 `json:"page_number,omitempty"`
	// PageSize 是每页条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总条数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// TotalPageCount 是总页数。
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}

// SearchAppKeywordDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordDeleteResponse struct {
}

// SearchAppKeywordQueryStatusResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordQueryStatusResponse struct {
	// KeyWords 是关键词工单审核状态返回值。
	KeyWords *[]search.KeyWordInfo `json:"key_words,omitempty"`
}

// SearchAppKeywordQuotaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordQuotaQueryResponse struct {
	// QuotaNum 是剩余可配置额度数量，返回具体数字。
	QuotaNum *string `json:"quota_num,omitempty"`
}
