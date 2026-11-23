// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package role

// CcmRoleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleCreateRequest struct {
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// Description 是角色描述信息
	Description *string `json:"description,omitempty"`
	// FunctionIds 是角色关联额功能点id
	FunctionIds *[]string `json:"function_ids,omitempty"`
	// Name 是角色名
	Name *string `json:"name,omitempty"`
}

// CcmRoleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleModifyRequest struct {
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID），不传默认使用租户id
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Description 是描述信息
	Description *string `json:"description,omitempty"`
	// FunctionIds 是角色关联额功能点id
	FunctionIds *[]string `json:"function_ids,omitempty"`
	// Id 是角色id
	Id *string `json:"id,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// CcmRoleDeleteRequest is the query request model for alipay.iservice.ccm.role.delete.
type CcmRoleDeleteRequest struct {
	// Id 是角色id
	Id *string `json:"id,omitempty" query:"id,omitempty"`
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID），如果不传入，使用租户id
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
}

// CcmRoleGetRequest is the query request model for alipay.iservice.ccm.role.get.
type CcmRoleGetRequest struct {
	// Id 是角色id
	Id *string `json:"id,omitempty" query:"id,omitempty"`
	// CcsInstanceId 是部门id，不传默认用租户ID
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
}

// CcmRolePageQueryRequest is the query request model for alipay.iservice.ccm.role.page.query.
type CcmRolePageQueryRequest struct {
	// Name 是角色名称
	Name *string `json:"name,omitempty" query:"name,omitempty"`
	// PageNum 是查询结果的页码，起始值为 1，默认值为 1
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是分页查询时设置的每页记录数，最大值 100 行，默认为 10
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
}
