// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ols

// CcmOlsChatRecordQueryRequest is the query request model for alipay.iservice.ccm.ols.chatrecord.query.
type CcmOlsChatRecordQueryRequest struct {
	// CcsInstanceId 是租户实例id（数据权限id），不填，则使用默认的租户实例id
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
	// Id 是在线服务记录id
	Id *string `json:"id,omitempty" query:"id,omitempty"`
}
