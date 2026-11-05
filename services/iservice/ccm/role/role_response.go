// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package role

import (
	"encoding/json"
)

// CcmRoleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleCreateResponse struct {
	// Id 是角色id
	Id *string `json:"id,omitempty"`
}

// CcmRoleDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleDeleteResponse = json.RawMessage

// CcmRoleGetResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleGetResponse struct {
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// CreateTime 是创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// Description 是描述信息
	Description *string `json:"description,omitempty"`
	// FunctionIds 是角色关联的功能点id
	FunctionIds *[]string `json:"function_ids,omitempty"`
	// Id 是角色id
	Id *string `json:"id,omitempty"`
	// Name 是角色名称
	Name *string `json:"name,omitempty"`
	// UpdateTime 是最后修改时间
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// CcmRoleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmRoleModifyResponse = json.RawMessage

// CcmRolePageQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmRolePageQueryResponse struct {
	// PageNum 是查询结果的页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// Roles 是角色列表
	Roles *[]Role `json:"roles,omitempty"`
	// TotalCount 是总条目数
	TotalCount *int64 `json:"total_count,omitempty"`
}
