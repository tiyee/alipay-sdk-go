// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package invite

// InviteOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InviteOrderQueryResponse struct {
	// MerchantPid 是商家支付宝账号对应的ID，2088开头。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// SignStatusList 是产品签约状态对象。
	SignStatusList *[]ProductInviteStatusInfo `json:"sign_status_list,omitempty"`
}
