package tradecomplain

// TargetInfo 是支付宝官方 V3 OpenAPI 模型。
type TargetInfo struct {
	// TargetId 是应用id，例如小程序id
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是应用类型	小程序传入：APPID	生活号传入：PUBLICID
	TargetType *string `json:"target_type,omitempty"`
}
