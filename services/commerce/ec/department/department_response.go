package department

import (
	"encoding/json"
)

// EcDepartmentCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentCreateResponse struct {
	// DepartmentId 是新创建的部门id
	DepartmentId *string `json:"department_id,omitempty"`
}

// EcDepartmentDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentDeleteResponse = json.RawMessage

// EcDepartmentInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentInfoModifyResponse = json.RawMessage

// EcDepartmentInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentInfoQueryResponse struct {
	DepartmentInfo *DepartmentInfoDTO `json:"department_info,omitempty"`
}

// EcDepartmentSubListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentSubListQueryResponse struct {
	// DepartmentIdList 是部门id列表
	DepartmentIdList *[]string `json:"department_id_list,omitempty"`
}
