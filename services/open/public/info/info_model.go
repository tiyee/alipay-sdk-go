// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// PublicAuditStatus 是支付宝官方 V3 OpenAPI 模型。
type PublicAuditStatus struct {
	// Desc 是审核状态描述，如果审核驳回则有相关的驳回理由。
	Desc *string `json:"desc,omitempty"`
	// Status 是AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过，AUDIT_NORMAL：无审核记录。
	Status *string `json:"status,omitempty"`
	// Type 是NAME:名称、AVATAR:头像、NAME_AVATAR:名称与头像、INTRODUCTION:简介。
	Type *string `json:"type,omitempty"`
}
