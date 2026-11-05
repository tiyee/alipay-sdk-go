package employee

// EcEmployeeAddRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeAddRequest struct {
	// AccountingEntityIds 是员工所属核算主体，核算主体可用于管控不同员工的出资方式，建议和不同出资账户关联
	AccountingEntityIds *[]string `json:"accounting_entity_ids,omitempty"`
	// CreateShareCode 是是否需要生成吱口令，默认不生成
	CreateShareCode *bool `json:"create_share_code,omitempty"`
	// DepartmentIds 是员工所属部门，支持多个部门，不传默认为根部门
	DepartmentIds *[]string `json:"department_ids,omitempty"`
	// EmployeeCertNo 是员工证件号码
	EmployeeCertNo *string `json:"employee_cert_no,omitempty"`
	// EmployeeCertType 是员工证件类型
	EmployeeCertType *string `json:"employee_cert_type,omitempty"`
	// EmployeeEmail 是员工邮箱
	EmployeeEmail *string `json:"employee_email,omitempty"`
	// EmployeeMobile 是员工手机号
	EmployeeMobile *string `json:"employee_mobile,omitempty"`
	// EmployeeName 是员工姓名
	EmployeeName *string `json:"employee_name,omitempty"`
	// EmployeeNo 是员工编号/工号
	EmployeeNo *string `json:"employee_no,omitempty"`
	// EncryptCertNo 是加密证件号，需与证件类型搭配使用，可替代证件号码(employee_cert_no)字段
	EncryptCertNo *string `json:"encrypt_cert_no,omitempty"`
	// EncryptMobile 是加密手机号，可替代手机号(employee_mobile)字段
	EncryptMobile *string `json:"encrypt_mobile,omitempty"`
	// EnterpriseId 是员工所属企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// FreeSignToken 是员工签约授权密钥；	该密钥需联系企业码业务小二申请，当传入密钥有效，则员工添加时会默认签约企业因公付，并激活员工。
	FreeSignToken *string `json:"free_sign_token,omitempty"`
	// Identity 是员工身份唯一标识
	Identity *string `json:"identity,omitempty"`
	// IdentityOpenId 是当身份类型为支付宝用户ID(identity_type=ALIPAY_USER_ID)时，传入user_id对应的open_id
	IdentityOpenId *string `json:"identity_open_id,omitempty"`
	// IdentityType 是员工身份类型
	IdentityType *string `json:"identity_type,omitempty"`
	// IotCheckType 是员工签约开通企业刷脸付时，刷脸核身的核验方式
	IotCheckType *string `json:"iot_check_type,omitempty"`
	// IotVid 是员工在企业人脸库的人脸唯一标识（该字段已作废，不支持自定义传入）(iot不支持自定义传入)
	IotVid *string `json:"iot_vid,omitempty"`
	// LabelNames 是员工标签，用于员工的打标分类，后续费控管理可使用标签进行控制，支持输入多个标签，如“差旅员工，用餐员工”等；若企业需要将离职员工统一管理且不变更员工状态，可打标“离职员工”
	LabelNames *[]string `json:"label_names,omitempty"`
	// Profiles 是个性化信息 详见文档
	Profiles *string `json:"profiles,omitempty"`
	// RoleList 是默认角色列表，目前只支持默认为 USER(目前仅支持普通员工，作为默认值，无需调用方传入)
	RoleList *[]string `json:"role_list,omitempty"`
	// SignReturnUrl 是员工签约后回跳地址；支持 alipay scheme 协议地址，不传默认跳转企业码小程序首页
	SignReturnUrl *string `json:"sign_return_url,omitempty"`
	// SignUrlCarryInfo 是签约链接页面是否需要自动回显员工的身份信息（手机或邮箱），默认不显示
	SignUrlCarryInfo *bool `json:"sign_url_carry_info,omitempty"`
	// WithholdingSignStr 是代扣签约字符串，员工需签约开通代扣时必传，否则无需传值；另外，如需参与企业码针对代扣的营销活动的服务商，必须要将该字段上传
	WithholdingSignStr *string `json:"withholding_sign_str,omitempty"`
}

// EcEmployeeDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeDeleteRequest struct {
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// EcEmployeeInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeInfoModifyRequest struct {
	// AccountingEntityIds 是员工所属核算主体，核算主体可用于管控不同员工的出资方式，建议和不同出资账户关联，一个员工只能有一个核算主体。
	AccountingEntityIds *[]string `json:"accounting_entity_ids,omitempty"`
	// DepartmentIds 是员工所属部门。如果不传值，则不更新所属部门
	DepartmentIds *[]string `json:"department_ids,omitempty"`
	// EmployeeEmail 是员工邮箱。如果不传值，则不更新邮箱。
	EmployeeEmail *string `json:"employee_email,omitempty"`
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty"`
	// EmployeeMobile 是员工手机号。如果不传值，则不更新手机号。
	EmployeeMobile *string `json:"employee_mobile,omitempty"`
	// EmployeeName 是员工姓名
	EmployeeName *string `json:"employee_name,omitempty"`
	// EmployeeNo 是员工工号。如果不传值，则不更新工号。
	EmployeeNo *string `json:"employee_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// LabelNames 是员工标签，用于员工的打标分类，后续费控管理可使用标签进行控制，支持输入多个标签，如“差旅员工，用餐员工”等；	使用场景：1. 请严格对标签分类，不要滥用员工标签，否则影响员工和费控管理； 2. 一个员工最多10个标签，如无必要请勿使用；3. 标签名只能包含字母（大小写）、数字或中文字符
	LabelNames *[]string `json:"label_names,omitempty"`
	// Profiles 是个性化信息。如果不传值，则不更新个性化信息。 详见文档
	Profiles *string `json:"profiles,omitempty"`
	// RoleList 是角色列表，目前只支持填一种角色。如果不传值，则不更新角色。
	RoleList *[]string `json:"role_list,omitempty"`
}

// EcEmployeeTitleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleCreateRequest struct {
	// EmployeeTitleList 是员工与抬头关系列表
	EmployeeTitleList *[]EmployeeTitleDTO `json:"employee_title_list,omitempty"`
}

// EcEmployeeTitleDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleDeleteRequest struct {
	// EmployeeTitleList 是员工抬头列表
	EmployeeTitleList *[]EmployeeTitleDTO `json:"employee_title_list,omitempty"`
}

// EcEmployeeTitleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleModifyRequest struct {
	// AccountId 是共同账号id，与enterprise_id两者必填其一
	AccountId *string `json:"account_id,omitempty"`
	// CreateBy 是创建人
	CreateBy *string `json:"create_by,omitempty"`
	// EmployeeId 是企业域定义的员工id，与user_id两者必填其一
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业域定义下的企业id，与accountId两者必填其一
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ModifyBy 是修改人
	ModifyBy *string `json:"modify_by,omitempty"`
	// MultiTitleList 是绑定的抬头ID列表
	MultiTitleList *[]EcEmployeeTitleModifyTaxNoAndTitleId `json:"multi_title_list,omitempty"`
	// NewTitleId 是修改后的抬头id
	NewTitleId *string `json:"new_title_id,omitempty"`
	// OldTitleId 是修改前的抬头id(数据库自行查询是否有有效抬头)
	OldTitleId *string `json:"old_title_id,omitempty"`
	// OpenId 是支付宝open_id，与employee_id两者必填其一
	OpenId *string `json:"open_id,omitempty"`
	// TaxRegisterNo 是税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// TitleTag 是员工抬头关系标
	TitleTag *string `json:"title_tag,omitempty"`
	// UserId 是支付宝id2088，与employee_id两者必填其一
	UserId *string `json:"user_id,omitempty"`
}

// EcEmployeeIDListQueryRequest is the query request model for alipay.commerce.ec.employee.idlist.query.
type EcEmployeeIDListQueryRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// DepartmentId 是部门id
	DepartmentId *string `json:"department_id,omitempty" query:"department_id,omitempty"`
	// PageNum 是查询页数
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是每页查询大小，限制最大为1000
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}

// EcEmployeeInfoQueryRequest is the query request model for alipay.commerce.ec.employee.info.query.
type EcEmployeeInfoQueryRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// OpenId 是蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// UserId 是支付宝用户ID
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty" query:"employee_id,omitempty"`
	// Mobile 是员工手机号码
	Mobile *string `json:"mobile,omitempty" query:"mobile,omitempty"`
	// AlipayLogonId 是支付宝登录号（手机号或邮箱）
	AlipayLogonId *string `json:"alipay_logon_id,omitempty" query:"alipay_logon_id,omitempty"`
	// EmployeeCertType 是员工证件类型
	EmployeeCertType *string `json:"employee_cert_type,omitempty" query:"employee_cert_type,omitempty"`
	// EmployeeCertNo 是员工证件号码
	EmployeeCertNo *string `json:"employee_cert_no,omitempty" query:"employee_cert_no,omitempty"`
	// EmployeeEmail 是员工邮箱
	EmployeeEmail *string `json:"employee_email,omitempty" query:"employee_email,omitempty"`
}

// EcEmployeeInviteQueryRequest is the query request model for alipay.commerce.ec.employee.invite.query.
type EcEmployeeInviteQueryRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// EmployeeId 是企业码员工ID，签约适用对象为指定员工时填写
	EmployeeId *string `json:"employee_id,omitempty" query:"employee_id,omitempty"`
	// PageContentCode 是用于展示不同的页面文案，使用标准页面文案是无需传值
	PageContentCode *string `json:"page_content_code,omitempty" query:"page_content_code,omitempty"`
	// WithholdingSignStr 是代扣签约串。需要签约代扣协议时，必传，返回的签约链接会拼接上代扣签约串。
	WithholdingSignStr *string `json:"withholding_sign_str,omitempty" query:"withholding_sign_str,omitempty"`
	// CreateShareCode 是是否生成签约吱口令，不传默认为N
	CreateShareCode *string `json:"create_share_code,omitempty" query:"create_share_code,omitempty"`
}
