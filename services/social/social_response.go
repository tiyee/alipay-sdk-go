// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package social

// SocialBaseContentLibStandardContentBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SocialBaseContentLibStandardContentBatchQueryResponse struct {
	// ContentDetails 是内容详情列表
	ContentDetails *[]SourceContentInfo `json:"content_details,omitempty"`
	// ContentIds 是内容ID列表
	ContentIds *[]string `json:"content_ids,omitempty"`
	// PageNum 是当前分页
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// Total 是总记录数
	Total *int64 `json:"total,omitempty"`
}
