// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package follow

// PublicFollowBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicFollowBatchQueryRequest struct {
	// NextUserId 是当关注者数量超过10000时使用，本次拉取数据中第一个用户的userId，从上次接口调用返回值中获取。第一次调用置空。
	NextUserId *string `json:"next_user_id,omitempty"`
	// OpenId 是支付宝用户的唯一标识。
	OpenId *string `json:"open_id,omitempty"`
}
