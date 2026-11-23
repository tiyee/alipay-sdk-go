package employee

// EcEmployeeTitleFailed 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleFailed struct {
	// EmployeeId 是员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Message 是返回每一个抬头批量处理信息
	Message *string `json:"message,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
	// TitleTag 是标识抬头是否默认DEFAULT/多抬头MULTI
	TitleTag *string `json:"title_tag,omitempty"`
}

// EcEmployeeTitleModifyTaxNoAndTitleId 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleModifyTaxNoAndTitleId struct {
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
}

// EmployeeAccountingEntityDTO 是支付宝官方 V3 OpenAPI 模型。
type EmployeeAccountingEntityDTO struct {
	// AccountingEntityCode 是核算主体编码
	AccountingEntityCode *string `json:"accounting_entity_code,omitempty"`
	// AccountingEntityId 是核算主体id
	AccountingEntityId *string `json:"accounting_entity_id,omitempty"`
	// AccountingEntityName 是核算主体名称
	AccountingEntityName *string `json:"accounting_entity_name,omitempty"`
}

// EmployeeDepartmentDTO 是支付宝官方 V3 OpenAPI 模型。
type EmployeeDepartmentDTO struct {
	// DepartmentId 是部门id
	DepartmentId *string `json:"department_id,omitempty"`
	// DepartmentName 是部门名称
	DepartmentName *string `json:"department_name,omitempty"`
}

// EmployeeInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type EmployeeInfoDTO struct {
	// AccountingEntityList 是员工所属核算主体列表
	AccountingEntityList *[]EmployeeAccountingEntityDTO `json:"accounting_entity_list,omitempty"`
	// Activate 是是否激活
	Activate *string `json:"activate,omitempty"`
	// DepartmentList 是员工所属部门列表
	DepartmentList *[]EmployeeDepartmentDTO `json:"department_list,omitempty"`
	// Email 是员工邮箱
	Email *string `json:"email,omitempty"`
	// EmployeeCertNo 是证件号
	EmployeeCertNo *string `json:"employee_cert_no,omitempty"`
	// EmployeeCertType 是证件类型
	EmployeeCertType *string `json:"employee_cert_type,omitempty"`
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeName 是员工姓名
	EmployeeName *string `json:"employee_name,omitempty"`
	// EmployeeNo 是员工编号/工号
	EmployeeNo *string `json:"employee_no,omitempty"`
	// EncryptCertNo 是加密证件号（证件号转大写后使用SHA256加密），搭配证件号使用
	EncryptCertNo *string `json:"encrypt_cert_no,omitempty"`
	// EncryptMobile 是加密手机号（使用SHA256进行加密）
	EncryptMobile *string `json:"encrypt_mobile,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是变更时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// InvoiceEmail 是员工收票邮箱
	InvoiceEmail *string `json:"invoice_email,omitempty"`
	// IotFaceStatus 是员工是否人脸在库
	IotFaceStatus *string `json:"iot_face_status,omitempty"`
	// IotUniqueId 是IOT开通刷脸支持唯一操作流水号，此处透出用于外部服务商通过该unique_id便捷调用IOT侧接口
	IotUniqueId *string `json:"iot_unique_id,omitempty"`
	// IotVid 是员工在企业人脸库的人脸唯一标识
	IotVid *string `json:"iot_vid,omitempty"`
	// JobLevelShow 是员工职级
	JobLevelShow *string `json:"job_level_show,omitempty"`
	// LabelNames 是员工标签
	LabelNames *[]string `json:"label_names,omitempty"`
	// Mobile 是手机号码
	Mobile *string `json:"mobile,omitempty"`
	// OpenId 是用户ID（绑定支付宝账号的uid）
	OpenId *string `json:"open_id,omitempty"`
	// Profiles 是个性化信息 详见文档
	Profiles *string `json:"profiles,omitempty"`
	// RoleList 是角色列表
	RoleList *[]string `json:"role_list,omitempty"`
	// TlEmployeeId 是直属主管员工ID
	TlEmployeeId *string `json:"tl_employee_id,omitempty"`
	// UserId 是用户ID（绑定支付宝账号的uid）
	UserId *string `json:"user_id,omitempty"`
}

// EmployeeTitleDTO 是支付宝官方 V3 OpenAPI 模型。
type EmployeeTitleDTO struct {
	// AccountId 是共同账户id,与enterprise_id两者必填其一
	AccountId *string `json:"account_id,omitempty"`
	// CreateBy 是创建人
	CreateBy *string `json:"create_by,omitempty"`
	// EmployeeId 是企业域定义的员工id，与user_id两者必填其一
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业域定义的企业id，与account_id两者必填其一
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ModifyBy 是修改人
	ModifyBy *string `json:"modify_by,omitempty"`
	// OpenId 是员工支付宝账号open_id,与employee_id两者必填其一
	OpenId *string `json:"open_id,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// TitleId 是抬头id
	TitleId *string `json:"title_id,omitempty"`
	// TitleTag 是关系标记
	TitleTag *string `json:"title_tag,omitempty"`
	// UserId 是员工支付宝账号2088,与employee_id两者必填其一
	UserId *string `json:"user_id,omitempty"`
}
