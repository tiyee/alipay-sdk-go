package jointaccountbill

// JointAccountBillDetail 是支付宝官方 V3 OpenAPI 模型。
type JointAccountBillDetail struct {
	// Amount 是消费金额，单位：元
	Amount *string `json:"amount,omitempty"`
	// BillNo 是账单业务号
	BillNo *string `json:"bill_no,omitempty"`
	// BizDate 是业务时间
	BizDate *string `json:"biz_date,omitempty"`
	// BizNo 是订单号
	BizNo *string `json:"biz_no,omitempty"`
	// BizOutNo 是外部单号
	BizOutNo *string `json:"biz_out_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InOut 是1-退款，2-支付
	InOut *string `json:"in_out,omitempty"`
	// OpenId 是用户支付宝uid
	OpenId *string `json:"open_id,omitempty"`
	// Title 是账单标题
	Title *string `json:"title,omitempty"`
	// UserId 是用户支付宝uid
	UserId *string `json:"user_id,omitempty"`
}
