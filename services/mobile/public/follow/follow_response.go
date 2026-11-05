package follow

// MobilePublicFollowListResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicFollowListResponse struct {
	// Code 是返回结果码，如200，标识成功
	Code *int64 `json:"code,omitempty"`
	// Count 是当前组的值
	Count *string `json:"count,omitempty"`
	Data  *Data   `json:"data,omitempty"`
	// NextAlipayUserId 是与nextUserId对应，标准Alipay UserId
	NextAlipayUserId *string `json:"next_alipay_user_id,omitempty"`
	// NextUserId 是查询分组的userid
	NextUserId *string `json:"next_user_id,omitempty"`
}
