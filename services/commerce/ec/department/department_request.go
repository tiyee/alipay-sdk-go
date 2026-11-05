package department

// EcDepartmentCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentCreateRequest struct {
	// DepartmentCode 是部门编码
	DepartmentCode *string `json:"department_code,omitempty"`
	// DepartmentName 是部门名称
	DepartmentName *string `json:"department_name,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ParentDepartmentId 是上级部门id
	ParentDepartmentId *string `json:"parent_department_id,omitempty"`
}

// EcDepartmentDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentDeleteRequest struct {
	// DepartmentId 是部门id
	DepartmentId *string `json:"department_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// EcDepartmentInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcDepartmentInfoModifyRequest struct {
	// DepartmentCode 是部门编码，不同部门的编码不能相同，为空则代表不修改
	DepartmentCode *string `json:"department_code,omitempty"`
	// DepartmentId 是待修改部门的部门id
	DepartmentId *string `json:"department_id,omitempty"`
	// DepartmentName 是部门名称，为空则代表不修改
	DepartmentName *string `json:"department_name,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ParentDepartmentId 是上级部门id，为空则代表不修改
	ParentDepartmentId *string `json:"parent_department_id,omitempty"`
}

// EcDepartmentInfoQueryRequest is the query request model for alipay.commerce.ec.department.info.query.
type EcDepartmentInfoQueryRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// DepartmentId 是部门id
	DepartmentId *string `json:"department_id,omitempty" query:"department_id,omitempty"`
}

// EcDepartmentSubListQueryRequest is the query request model for alipay.commerce.ec.department.sublist.query.
type EcDepartmentSubListQueryRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// DepartmentId 是部门id，特殊值 -1 表示查询根部门id
	DepartmentId *string `json:"department_id,omitempty" query:"department_id,omitempty"`
}
