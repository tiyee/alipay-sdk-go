// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sign

import (
	"encoding/json"
)

// SignFlowCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowCancelResponse = json.RawMessage

// SignFlowCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowCreateResponse struct {
	// FlowId 是流程Id
	FlowId *string `json:"flow_id,omitempty"`
}

// SignFlowFinishResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowFinishResponse = json.RawMessage

// SignFlowQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowQueryResponse struct {
	// BusinessScene 是文件主题
	BusinessScene *string `json:"business_scene,omitempty"`
	// ContractValidity 是文件有效截止日期
	ContractValidity *int64 `json:"contract_validity,omitempty"`
	// FlowDesc 是流程描述, 如果流程已拒签或已撤回, 并且存在拒签或撤回原因, 流程描述显示为原因, 否则默认为流程状态描述
	FlowDesc *string `json:"flow_desc,omitempty"`
	// FlowEndTime 是流程结束时间
	FlowEndTime *int64 `json:"flow_end_time,omitempty"`
	// FlowId 是流程Id
	FlowId *string `json:"flow_id,omitempty"`
	// FlowStartTime 是流程开始时间
	FlowStartTime *int64 `json:"flow_start_time,omitempty"`
	// FlowStatus 是流程状态,0-草稿 1-签署中 2-完成 3-撤销 4-终止 5-过期 6-删除 7-拒签
	FlowStatus *int64 `json:"flow_status,omitempty"`
	// NoticeDeveloperUrl 是通知开发者地址
	NoticeDeveloperUrl *string `json:"notice_developer_url,omitempty"`
	// SignValidity 是签署有效截止日期
	SignValidity *int64 `json:"sign_validity,omitempty"`
	// Signers 是签署人列表及签署状态
	Signers *[]FlowSigner `json:"signers,omitempty"`
}
