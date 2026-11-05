package jointaccountbill

// EcJointAccountBillDetailBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcJointAccountBillDetailBatchQueryResponse struct {
	// BillList 是返回账单列表
	BillList *[]JointAccountBillDetail `json:"bill_list,omitempty"`
	// BizScene 是业务场景
	BizScene *string `json:"biz_scene,omitempty"`
	// EnterpriseId 是企业ID，和接口入参一致
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// HasNextPage 是是否有下一页
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// PageNum 是当期页数
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是当期页记录数
	PageSize *string `json:"page_size,omitempty"`
}
