// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package signflows

// SignFlowsDetailQueryRequest is the query request model for alipay.eco.signflows.detail.query.
type SignFlowsDetailQueryRequest struct {
	// TargetAppId 是目标isv应用ID
	TargetAppId *string `json:"target_app_id,omitempty" query:"target_app_id,omitempty"`
	// FlowId 是流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。
	FlowId *string `json:"flow_id,omitempty" query:"flow_id,omitempty"`
}

// SignFlowsURLQueryRequest is the query request model for alipay.eco.signflows.url.query.
type SignFlowsURLQueryRequest struct {
	// TargetAppId 是目标isv应用ID
	TargetAppId *string `json:"target_app_id,omitempty" query:"target_app_id,omitempty"`
	// FlowId 是流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。
	FlowId *string `json:"flow_id,omitempty" query:"flow_id,omitempty"`
	// ThirdPartyUserId 是创建流程时指定个人唯一标识
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty" query:"third_party_user_id,omitempty"`
	// OrgThirdPartyUserId 是创建流程时指定企业唯一标识
	OrgThirdPartyUserId *string `json:"org_third_party_user_id,omitempty" query:"org_third_party_user_id,omitempty"`
}
