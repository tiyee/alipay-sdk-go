package tree

import (
	"encoding/json"
)

// CcmSwTreeBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeBatchQueryResponse struct {
	// Trees 是类目集合
	Trees *[]TreeInfo `json:"trees,omitempty"`
}

// CcmSwTreeCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCreateResponse struct {
	// Id 是类目ID
	Id *int64 `json:"id,omitempty"`
}

// CcmSwTreeDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeDeleteResponse = json.RawMessage

// CcmSwTreeGetResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeGetResponse struct {
	// Tree 是类目树字符串
	Tree *string `json:"tree,omitempty"`
}

// CcmSwTreeModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeModifyResponse = json.RawMessage

// CcmSwTreePublishResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreePublishResponse = json.RawMessage
