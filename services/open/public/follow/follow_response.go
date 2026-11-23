// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package follow

// PublicFollowBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicFollowBatchQueryResponse struct {
	// Count 是本次调用获取的userId个数，最大值为10000。
	Count *string `json:"count,omitempty"`
	// NextUserId 是查询分组的userid。
	NextUserId *string `json:"next_user_id,omitempty"`
	// OpenId 是支付宝用户的。
	OpenId *string `json:"open_id,omitempty"`
	// OpenIdList 是支付宝用户的唯一标识列表。
	OpenIdList *[]string `json:"open_id_list,omitempty"`
	// UserIdList 是用户的userId列表。
	UserIdList *[]string `json:"user_id_list,omitempty"`
}
