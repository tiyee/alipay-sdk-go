// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package payee

// PublicPayeeBindCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicPayeeBindCreateRequest struct {
	// LoginId 是收款账号，需要绑定的收款支付宝账号，请勿与 pid 同时传入。。
	LoginId *string `json:"login_id,omitempty"`
	// Pid 是支付宝用户id，2088开头的16位长度字符串，请勿与 login_id 同时传入。。
	Pid *string `json:"pid,omitempty"`
}

// PublicPayeeBindDeleteRequest is the request type for alipay.open.public.payee.bind.delete.
type PublicPayeeBindDeleteRequest struct {
	// LoginId 是收款账号，需要绑定的收款支付宝账号，请勿与 pid 同时传入。
	LoginId *string `json:"login_id,omitempty" query:"login_id,omitempty"`
	// Pid 是支付宝用户id，2088开头的16位长度字符串，请勿与 login_id 同时传入。
	Pid *string `json:"pid,omitempty" query:"pid,omitempty"`
}
