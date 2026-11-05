package oppor

// ISVExpandOpporDTO 是支付宝官方 V3 OpenAPI 模型。
type ISVExpandOpporDTO struct {
	// Address 是商机地址。
	Address *string `json:"address,omitempty"`
	// LeadsId 是商机Id。
	LeadsId *string `json:"leads_id,omitempty"`
	// Name 是商机名称。
	Name *string `json:"name,omitempty"`
	// OpporId 是商机作业Id。
	OpporId *string `json:"oppor_id,omitempty"`
	// OutBizNo 是外部幂等唯一键。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Phone 是商机联系电话。
	Phone *string `json:"phone,omitempty"`
	// Status 是商机作业状态。
	Status *string `json:"status,omitempty"`
}
