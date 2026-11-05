package library

// CcmSwLibraryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryCreateRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// InstanceCode 是知识库对应的类目树名称，与treeId不可同时为空
	InstanceCode *string `json:"instance_code,omitempty"`
	// Name 是知识库名称
	Name *string `json:"name,omitempty"`
	// TreeId 是知识库对应的类目树Id，与instanceCode不可同时为空
	TreeId *int64 `json:"tree_id,omitempty"`
}

// CcmSwLibraryModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwLibraryModifyRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Id 是知识库Id
	Id *string `json:"id,omitempty"`
	// Name 是知识库名称
	Name *string `json:"name,omitempty"`
}

// CcmSwLibraryBatchQueryRequest is the query request model for alipay.iservice.ccm.sw.library.batchquery.
type CcmSwLibraryBatchQueryRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
}

// CcmSwLibraryDeleteRequest is the request model for alipay.iservice.ccm.sw.library.delete.
type CcmSwLibraryDeleteRequest struct{}
