package knowledge

// CcmSwKnowledgeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwKnowledgeModifyRequest struct {
	// CategoryId 是所属类目ID
	CategoryId *int64 `json:"category_id,omitempty"`
	// CategoryName 是知识点所属类目
	CategoryName *string `json:"category_name,omitempty"`
	// Content 是知识点答案内容
	Content *string `json:"content,omitempty"`
	// ExtId 是外部知识点ID
	ExtId *string `json:"ext_id,omitempty"`
	// ExtendTitles 是扩展标题（问法）
	ExtendTitles *[]string `json:"extend_titles,omitempty"`
	// Icon 是问答关联图片附件地址
	Icon *string `json:"icon,omitempty"`
	// Id 是知识点ID
	Id *int64 `json:"id,omitempty"`
	// IsDelete 是默认为false，为true且id非空时，表示删除知识点
	IsDelete *bool `json:"is_delete,omitempty"`
	// IsSearchable 是是否被检索到，默认false，系统希望为true
	IsSearchable *bool `json:"is_searchable,omitempty"`
	// LibraryId 是知识库ID
	LibraryId *int64 `json:"library_id,omitempty"`
	// LibraryName 是知识库名称
	LibraryName *string    `json:"library_name,omitempty"`
	Spu         *SPUDetail `json:"spu,omitempty"`
	// Tags 是以逗号分割的字符串列表
	Tags *[]string `json:"tags,omitempty"`
	// Title 是知识点标题
	Title *string `json:"title,omitempty"`
}
