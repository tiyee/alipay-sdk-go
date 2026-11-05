// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package invite

// InviteOrderQueryRequest is the request type for alipay.open.invite.order.query.
type InviteOrderQueryRequest struct {
	// IsvBizId 是isv业务系统单据编号
	IsvBizId *string `json:"isv_biz_id,omitempty" query:"isv_biz_id,omitempty"`
	// AlipayLogonId 是商家的支付宝账号，如果调用alipay.open.invite.order.create接口时有传入支付宝账号，则这里查询签约状态时必须要传
	AlipayLogonId *string `json:"alipay_logon_id,omitempty" query:"alipay_logon_id,omitempty"`
}
