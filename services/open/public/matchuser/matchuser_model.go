// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package matchuser

// ErrorMatcher 是支付宝官方 V3 OpenAPI 模型。
type ErrorMatcher struct {
	// ErrorMsg 是失败原因。
	ErrorMsg *string  `json:"error_msg,omitempty"`
	Matcher  *Matcher `json:"matcher,omitempty"`
}

// Matcher 是支付宝官方 V3 OpenAPI 模型。
type Matcher struct {
	// IdentityCard 是身份证号码，与user_id、mobile_no不能同时为空。
	IdentityCard *string `json:"identity_card,omitempty"`
	// MobileNo 是手机号码，与user_id、identity_card不能同时为空。
	MobileNo *string `json:"mobile_no,omitempty"`
	// OpenId 是支付宝用户openid。
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户id，2088开头16位长度的字符串，与mobile_no、identity_card不能同时为空。
	UserId *string `json:"user_id,omitempty"`
}
