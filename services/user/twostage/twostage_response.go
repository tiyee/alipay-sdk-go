package twostage

// UserTwoStageCommonUseResponse 是支付宝官方 V3 OpenAPI 模型。
type UserTwoStageCommonUseResponse struct {
	// OpenId 是open_id
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户userId信息，因为用户已经在客户端给商户的小程序授权了，并且商户要通过userId信息挂接优惠券信息，所以可以无需脱敏返回给商户。
	UserId           *string           `json:"user_id,omitempty"`
	UserIdentityInfo *UserIdentityInfo `json:"user_identity_info,omitempty"`
}

// UserTwoStageIndirectUseResponse 是支付宝官方 V3 OpenAPI 模型。
type UserTwoStageIndirectUseResponse struct {
	// UserId 是蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty"`
}
