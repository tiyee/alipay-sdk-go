// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package group

// PublicGroupBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupBatchQueryResponse struct {
	// Groups 是用户分组列表，包含每个分组的id、name、以及规则模型。
	Groups *[]QueryGroup `json:"groups,omitempty"`
}

// PublicGroupCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupCreateResponse struct {
	// GroupId 是分组id。
	GroupId *string `json:"group_id,omitempty"`
}

// PublicGroupCrowdQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupCrowdQueryResponse struct {
	// Count 是分组圈出的人群数量。
	Count *string `json:"count,omitempty"`
}

// PublicGroupDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupDeleteResponse struct {
}

// PublicGroupModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupModifyResponse struct {
}
