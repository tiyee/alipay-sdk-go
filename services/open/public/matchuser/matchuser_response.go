// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package matchuser

// PublicMatchUserLabelCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMatchUserLabelCreateResponse struct {
	// ErrorCount 是用户打标失败数量。
	ErrorCount *int64 `json:"error_count,omitempty"`
	// ErrorMatchers 是出错的匹配器列表。
	ErrorMatchers *[]ErrorMatcher `json:"error_matchers,omitempty"`
}

// PublicMatchUserLabelDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMatchUserLabelDeleteResponse struct {
	// ErrorCount 是用户打标失败数量。
	ErrorCount *int64 `json:"error_count,omitempty"`
	// ErrorMatchers 是出错的匹配器列表。
	ErrorMatchers *[]ErrorMatcher `json:"error_matchers,omitempty"`
}
