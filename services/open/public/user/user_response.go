// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

// PublicUserDataBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicUserDataBatchQueryResponse struct {
	// DataList 是用户分析数据。
	DataList *[]UserAnalysisData `json:"data_list,omitempty"`
}

// PublicUserFollowQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicUserFollowQueryResponse struct {
	// IsFollow 是用户是否关注。枚举值如下： *T：代表已关注。 *F：代表未关注。。
	IsFollow *string `json:"is_follow,omitempty"`
}
