package treecategory

// CcmSwTreeCategoryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryCreateRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Description 是描述
	Description *string `json:"description,omitempty"`
	// FatherId 是父节点ID
	FatherId *int64 `json:"father_id,omitempty"`
	// Name 是节点名称
	Name *string `json:"name,omitempty"`
	// Tags 是标签。KNOWLEDGE（知识库）；PLATFORM（公有云工作台）；HELP（公有云帮助中心）
	Tags *[]string `json:"tags,omitempty"`
}

// CcmSwTreeCategoryModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCategoryModifyRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Description 是描述
	Description *string `json:"description,omitempty"`
	// FatherId 是父节点ID
	FatherId *string `json:"father_id,omitempty"`
	// Id 是节点ID
	Id *int64 `json:"id,omitempty"`
	// Name 是节点名称
	Name *string `json:"name,omitempty"`
	// Tags 是标签。KNOWLEDGE（知识库）；PLATFORM（公有云工作台）；HELP（公有云帮助中心）
	Tags *string `json:"tags,omitempty"`
}

// CcmSwTreeCategoryBatchQueryRequest is the query request model for alipay.iservice.ccm.sw.treecategory.batchquery.
type CcmSwTreeCategoryBatchQueryRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
	// LibraryId 是知识库ID
	LibraryId *int64 `json:"library_id,omitempty" query:"library_id,required"`
	// Name 是节点名称
	Name *string `json:"name,omitempty" query:"name,omitempty"`
	// PageNum 是页数，page_size不能为空
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是页显示大小，page_num不能为空
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,omitempty"`
}

// CcmSwTreeCategoryDeleteRequest is the request model for alipay.iservice.ccm.sw.treecategory.delete.
type CcmSwTreeCategoryDeleteRequest struct{}
