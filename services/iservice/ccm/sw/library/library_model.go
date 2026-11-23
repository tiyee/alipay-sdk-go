package library

// LibraryInfo 是支付宝官方 V3 OpenAPI 模型。
type LibraryInfo struct {
	// Id 是知识库ID
	Id *int64 `json:"id,omitempty"`
	// Name 是知识库名称
	Name *string `json:"name,omitempty"`
	// TreeId 是类目ID
	TreeId *int64 `json:"tree_id,omitempty"`
	// UpdateTime 是修改时间
	UpdateTime *string `json:"update_time,omitempty"`
}
