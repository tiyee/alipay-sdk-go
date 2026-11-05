// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package security

import (
	"encoding/json"
)

// SecurityProdSssQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SecurityProdSssQueryResponse struct {
	// Ccc 是1
	Ccc *string `json:"ccc,omitempty"`
	// CccOpenId 是1
	CccOpenId *string `json:"ccc_open_id,omitempty"`
}

// SecurityRiskCustomerRiskSendResponse 是支付宝官方 V3 OpenAPI 模型。
type SecurityRiskCustomerRiskSendResponse = json.RawMessage

// SecurityRiskVerifyIdentityMiniAppConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type SecurityRiskVerifyIdentityMiniAppConfirmResponse struct {
	// ErrorCode 是身份核实校验结果失败时，返回的失败错误码
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMsg 是身份核实校验结果失败时，返回的失败错误原因说明
	ErrorMsg *string `json:"error_msg,omitempty"`
	// VerifyResult 是身份核验结果
	VerifyResult *bool `json:"verify_result,omitempty"`
}
