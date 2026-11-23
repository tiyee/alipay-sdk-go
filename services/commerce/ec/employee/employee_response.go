package employee

import (
	"encoding/json"
)

// EcEmployeeAddResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeAddResponse struct {
	// EmployeeId 是企业码生成的员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// IotUniqueId 是员工开通IoT企业刷脸付的唯一流水号
	IotUniqueId *string `json:"iot_unique_id,omitempty"`
	// ShareCode 是员工签约激活吱口令
	ShareCode *string `json:"share_code,omitempty"`
	// SignUrl 是员工签约激活页面链接
	SignUrl *string `json:"sign_url,omitempty"`
}

// EcEmployeeDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeDeleteResponse = json.RawMessage

// EcEmployeeIDListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeIDListQueryResponse struct {
	// CurrentPage 是当前页数
	CurrentPage *int64 `json:"current_page,omitempty"`
	// EmployeeIdList 是员工id列表
	EmployeeIdList *[]string `json:"employee_id_list,omitempty"`
	// TotalNum 是员工总数
	TotalNum *int64 `json:"total_num,omitempty"`
	// TotalPages 是总页数
	TotalPages *int64 `json:"total_pages,omitempty"`
}

// EcEmployeeInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeInfoModifyResponse struct {
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty"`
}

// EcEmployeeInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeInfoQueryResponse struct {
	EmployeeInfo *EmployeeInfoDTO `json:"employee_info,omitempty"`
}

// EcEmployeeInviteQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeInviteQueryResponse struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// MiniAppSignUrl 是支付宝服务商小程序跳入企业码小程序签约链接
	MiniAppSignUrl *string `json:"mini_app_sign_url,omitempty"`
	// ShareCode 是签约吱口令
	ShareCode *string `json:"share_code,omitempty"`
	// SignUrl 是邀请链接
	SignUrl *string `json:"sign_url,omitempty"`
}

// EcEmployeeTitleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleCreateResponse struct {
	TitleFailedList *EcEmployeeTitleFailed `json:"title_failed_list,omitempty"`
}

// EcEmployeeTitleDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleDeleteResponse struct {
	EcEmployeeTitleFailedList *EcEmployeeTitleFailed `json:"ec_employee_title_failed_list,omitempty"`
}

// EcEmployeeTitleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEmployeeTitleModifyResponse = json.RawMessage
