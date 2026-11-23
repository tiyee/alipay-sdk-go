package follow

// Data 是支付宝官方 V3 OpenAPI 模型。
type Data struct {
	// AlipayOpenIdList 是支付宝用户唯一标识列表
	AlipayOpenIdList *[]string `json:"alipay_open_id_list,omitempty"`
	// AlipayUserIdList 是对应userIdList，标准Alipay UserId列表
	AlipayUserIdList *[]string `json:"alipay_user_id_list,omitempty"`
	// OpenIdList 是支付宝用户唯一标识列表
	OpenIdList *[]string `json:"open_id_list,omitempty"`
	// UserIdList 是用户id列表
	UserIdList *[]string `json:"user_id_list,omitempty"`
}
