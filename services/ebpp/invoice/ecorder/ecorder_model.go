// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ecorder

// EcOrderItem 是支付宝官方 V3 OpenAPI 模型。
type EcOrderItem struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// BizOutNo 是外部业务号
	BizOutNo *string `json:"biz_out_no,omitempty"`
	// EmployeeId 是员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// OpenId 是用户支付宝UID
	OpenId *string `json:"open_id,omitempty"`
	// OrderContent 是订单内容，JSON格式。
	OrderContent *string `json:"order_content,omitempty"`
	// OrderId 是订单ID
	OrderId *string `json:"order_id,omitempty"`
	// OrderType 是订单类型	METRO：地铁	TAKE_AWAY：外卖	OTHER：其他	REACH_SHOP：到店
	OrderType *string `json:"order_type,omitempty"`
	// PartnerId 是合作伙伴ID
	PartnerId *string `json:"partner_id,omitempty"`
	// PayNo 是关联支付宝交易号
	PayNo *string `json:"pay_no,omitempty"`
	// UserId 是用户支付宝UID
	UserId *string `json:"user_id,omitempty"`
}
