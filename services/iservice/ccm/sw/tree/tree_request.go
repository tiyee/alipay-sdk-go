package tree

// CcmSwTreeCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeCreateRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Name 是类目名称
	Name *string `json:"name,omitempty"`
}

// CcmSwTreeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreeModifyRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Id 是类目ID
	Id *int64 `json:"id,omitempty"`
	// Name 是类目名称
	Name *string `json:"name,omitempty"`
}

// CcmSwTreePublishRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwTreePublishRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Id 是类目ID
	Id *int64 `json:"id,omitempty"`
}

// CcmSwTreeBatchQueryRequest is the query request model for alipay.iservice.ccm.sw.tree.batchquery.
type CcmSwTreeBatchQueryRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
}

// CcmSwTreeDeleteRequest is the request model for alipay.iservice.ccm.sw.tree.delete.
type CcmSwTreeDeleteRequest struct{}

// CcmSwTreeGetRequest is the query request model for alipay.iservice.ccm.sw.tree.get.
type CcmSwTreeGetRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
	// Id 是类目ID
	Id *int64 `json:"id,omitempty" query:"id,required"`
}
