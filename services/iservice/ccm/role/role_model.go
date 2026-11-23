// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package role

// Role 是支付宝官方 V3 OpenAPI 模型。
type Role struct {
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// CreateTime 是创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// Description 是角色描述描述
	Description *string `json:"description,omitempty"`
	// FunctionIds 是角色关联额功能点id
	FunctionIds *[]string `json:"function_ids,omitempty"`
	// Id 是角色id
	Id *string `json:"id,omitempty"`
	// Name 是角色名
	Name *string `json:"name,omitempty"`
	// UpdateTime 是最后修改时间
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}
