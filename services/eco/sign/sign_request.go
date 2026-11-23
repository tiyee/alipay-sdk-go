// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sign

// SignFlowCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type SignFlowCancelRequest struct {
	// TargetAppId 在 URL query 中传递。目标isv应用ID
	TargetAppId *string `query:"target_app_id" json:"-"`
	// FlowId 是流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。
	FlowId *string `json:"flow_id,omitempty"`
	// RevokeReason 是撤销原因，默认为"撤销"。商家可自定义。
	RevokeReason *string `json:"revoke_reason,omitempty"`
}

// SignFlowCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type SignFlowCreateRequest struct {
	// TargetAppId 在 URL query 中传递。目标isv应用ID
	TargetAppId *string `query:"target_app_id" json:"-"`
	// Attachments 是附件信息
	Attachments *[]Attachment `json:"attachments,omitempty"`
	// BusinessScene 是流程主题
	BusinessScene *string     `json:"business_scene,omitempty"`
	ConfigInfo    *ConfigInfo `json:"config_info,omitempty"`
	// TemplateInfos 是模板信息
	TemplateInfos *[]TemplateInfo `json:"template_infos,omitempty"`
}

// SignFlowFinishRequest 是支付宝官方 V3 OpenAPI 模型。
type SignFlowFinishRequest struct {
	// TargetAppId 在 URL query 中传递。目标isv应用ID
	TargetAppId *string `query:"target_app_id" json:"-"`
	// FlowId 是流程id
	FlowId *string `json:"flow_id,omitempty"`
}

// SignFlowQueryRequest is the query request model for alipay.eco.sign.flow.query.
type SignFlowQueryRequest struct {
	// TargetAppId 是目标isv应用ID
	TargetAppId *string `json:"target_app_id,omitempty" query:"target_app_id,omitempty"`
	// FlowId 是流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。
	FlowId *string `json:"flow_id,omitempty" query:"flow_id,omitempty"`
}
