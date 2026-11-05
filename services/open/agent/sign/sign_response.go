package sign

// AgentCommonSignResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentCommonSignResponse struct {
}

// AgentCommonSignConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentCommonSignConfirmResponse struct {
	// OrderNo 是签约单号。
	OrderNo *string `json:"order_no,omitempty"`
}

// AgentFaceToFaceSignResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentFaceToFaceSignResponse struct {
}

// AgentMobilePaySignResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentMobilePaySignResponse struct {
}

// AgentOfflinePaymentSignResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentOfflinePaymentSignResponse struct {
}

// AgentSignStatusQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentSignStatusQueryResponse struct {
	// SignStatusList 是商户产品签约状态结果对象数组，返回本次查询的产品的签约状态。
	SignStatusList *[]ProductSignStatusInfo `json:"sign_status_list,omitempty"`
}
