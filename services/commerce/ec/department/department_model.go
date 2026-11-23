package department

// DepartmentInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type DepartmentInfoDTO struct {
	// DepartmentCode 是部门编码
	DepartmentCode *string `json:"department_code,omitempty"`
	// DepartmentId 是部门id
	DepartmentId *string `json:"department_id,omitempty"`
	// DepartmentName 是部门名称
	DepartmentName *string `json:"department_name,omitempty"`
	// DeptChargeEmployeeId 是部门负责人员工ID
	DeptChargeEmployeeId *string `json:"dept_charge_employee_id,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// ParentDepartmentId 是上级部门id，特殊值-1表示根部门
	ParentDepartmentId *string `json:"parent_department_id,omitempty"`
}
