package jointaccountbill

// EcJointAccountBillDetailBatchQueryRequest is the query request model for alipay.commerce.ec.jointaccountbill.detail.batchquery.
type EcJointAccountBillDetailBatchQueryRequest struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// BizScene 是业务场景。不传默认为ISV_DEFAULT
	BizScene *string `json:"biz_scene,omitempty" query:"biz_scene,omitempty"`
	// UserId 是员工支付宝UID
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是员工支付宝UID
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// EmployeeId 是员工id
	EmployeeId *string `json:"employee_id,omitempty" query:"employee_id,omitempty"`
	// PageNum 是页码
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是每页大小
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
	// StartDate 是起始日期
	StartDate *string `json:"start_date,omitempty" query:"start_date,omitempty"`
	// EndDate 是截止日期
	EndDate *string `json:"end_date,omitempty" query:"end_date,omitempty"`
}
