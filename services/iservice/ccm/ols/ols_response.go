// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ols

// CcmOlsChatRecordQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmOlsChatRecordQueryResponse struct {
	// AgentId 是客服id
	AgentId *string `json:"agent_id,omitempty"`
	// AgentName 是客服昵称
	AgentName *string `json:"agent_name,omitempty"`
	// Categories 是多个层级类目信息
	Categories *string `json:"categories,omitempty"`
	// GroupId 是在线技能组id
	GroupId *string `json:"group_id,omitempty"`
	// GroupName 是在线技能组名称
	GroupName *string `json:"group_name,omitempty"`
	// Memo 是备注信息
	Memo *string `json:"memo,omitempty"`
	// Messages 是聊天消息列表
	Messages *[]ChatMessage `json:"messages,omitempty"`
	// Satisfaction 是满意度。 0：非常满意 1：满意 2：一般 3：不满意 4：非常不满意
	Satisfaction *string `json:"satisfaction,omitempty"`
	// Status 是服务记录状态。1：待处理 2：暂存 3：完结 4：废除 5：无效 6：未知
	Status *string `json:"status,omitempty"`
	// TalkDuration 是在线服务时长，单位秒
	TalkDuration *int64 `json:"talk_duration,omitempty"`
	// VisitorProvince 是访客归属地
	VisitorProvince *string `json:"visitor_province,omitempty"`
}
