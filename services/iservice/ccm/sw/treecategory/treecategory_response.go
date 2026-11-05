package treecategory

import (
	"encoding/json"

	"github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw"
)

// CcmSwTreeCategoryBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryBatchQueryResponse struct {
	// Categories 是节点集合
	Categories *[]sw.ArticleCategoryInfo `json:"categories,omitempty"`
	// PageNum 是页数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是页显示大小
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总行数
	TotalCount *int64 `json:"total_count,omitempty"`
}

// CcmSwTreeCategoryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryCreateResponse struct {
	// Id 是节点ID
	Id *int64 `json:"id,omitempty"`
}

// CcmSwTreeCategoryDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryDeleteResponse = json.RawMessage

// CcmSwTreeCategoryModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryModifyResponse = json.RawMessage
