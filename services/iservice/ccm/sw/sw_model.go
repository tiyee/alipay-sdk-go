package sw

// ArticleCategoryInfo 是支付宝官方 V3 OpenAPI 模型。
type ArticleCategoryInfo struct {
	// Children 是子节点
	Children *[]string `json:"children,omitempty"`
	// Description 是节点描述
	Description *string `json:"description,omitempty"`
	// FatherId 是父节点ID
	FatherId *int64 `json:"father_id,omitempty"`
	// Id 是类目ID
	Id *int64 `json:"id,omitempty"`
	// Name 是节点名称
	Name *string `json:"name,omitempty"`
	// Tags 是标签
	Tags *[]string `json:"tags,omitempty"`
}
