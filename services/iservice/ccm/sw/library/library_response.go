package library

import (
	"encoding/json"
)

// CcmSwLibraryBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryBatchQueryResponse struct {
	// Libraries 是知识库集合
	Libraries *[]LibraryInfo `json:"libraries,omitempty"`
}

// CcmSwLibraryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryCreateResponse struct {
	// Id 是知识库Id
	Id *int64 `json:"id,omitempty"`
}

// CcmSwLibraryDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryDeleteResponse = json.RawMessage

// CcmSwLibraryModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryModifyResponse = json.RawMessage
