// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

// PublicUserDataBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicUserDataBatchQueryRequest struct {
	// BeginDate 是数据开始日期，时间格式为 "yyyyMMdd" 。。
	BeginDate *string `json:"begin_date,omitempty"`
	// EndDate 是数据结束日期，时间格式为"yyyyMMdd"。查询数据开始日期/结束日期时间跨度最大30天。。
	EndDate *string `json:"end_date,omitempty"`
}

// PublicUserFollowQueryRequest is the request type for alipay.open.public.user.follow.query.
type PublicUserFollowQueryRequest struct {
	// UserId 是支付宝用户id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是支付宝用户唯一的标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}
