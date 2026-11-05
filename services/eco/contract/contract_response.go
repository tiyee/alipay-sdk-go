// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package contract

// ContractSignFlowsCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type ContractSignFlowsCreateResponse struct {
	// FlowId 是电子合同流程id（需记录到业务系统中，后续相关服务调用都基于此ID）
	FlowId *string `json:"flow_id,omitempty"`
}
