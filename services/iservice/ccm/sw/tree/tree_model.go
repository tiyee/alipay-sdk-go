package tree

// TreeInfo 是支付宝官方 V3 OpenAPI 模型。
type TreeInfo struct {
	// Id 是类目树ID
	Id *int64 `json:"id,omitempty"`
	// InstanceCode 是知识库对应的类目树名称，与treeId不可同时为空
	InstanceCode *string `json:"instance_code,omitempty"`
	// Name 是类目名称
	Name *string `json:"name,omitempty"`
}
