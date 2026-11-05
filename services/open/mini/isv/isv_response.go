// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

// MiniISVCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniISVCreateResponse struct {
	// OrderNo 是小程序代创建订单号。
	OrderNo *string `json:"order_no,omitempty"`
}

// MiniISVFastRegisterCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniISVFastRegisterCreateResponse struct {
	// AuthorizeUrl 是授权确认跳转url。
	AuthorizeUrl *string `json:"authorize_url,omitempty"`
	// OrderNo 是代创建试用小程序单号。
	OrderNo *string `json:"order_no,omitempty"`
}

// MiniISVFastRegisterQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniISVFastRegisterQueryResponse struct {
	// AppName 是小程序应用名。
	AppName *string `json:"app_name,omitempty"`
	// IsvAppId 是服务商appid。
	IsvAppId *string `json:"isv_app_id,omitempty"`
	// MiniAppId 是小程序appid。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// OrderNo 是代创建试用小程序单号。
	OrderNo *string `json:"order_no,omitempty"`
	// OutOrderNo 是商户传入的外部订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// Status 是WAIT_CONFIRM待确认创建， CONFIRM_TIMEOUT待创建确认授权超时， CREATE_MINI_APP_SUCCESS创建应用成功，CREATE_MINI_APP_FAILED创建应用失败，CANCELLED已注销，CERTIFY_SUCCESS认证成功，CERTIFY_FAILED认证失败，CERTIFY_OPERATOR管理员认证中，CERTIFY_MERCHANT法人认证中。
	Status *string `json:"status,omitempty"`
}

// MiniISVQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniISVQueryResponse struct {
	// MinAppId 是小程序appId，商家确认后非空，若商家未确认或超时返回空。
	MinAppId *string `json:"min_app_id,omitempty"`
	// OrderNo 是订单号。
	OrderNo *string `json:"order_no,omitempty"`
	// OutOrderNo 是外部订单号，开发者帐号+外部订单号维度来控制请求务幂等。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// Pid 是商户pid。
	Pid *string `json:"pid,omitempty"`
	// Status 是PROCESS处理中，TIMEOUT超时，AGREED同意， REJECTED拒绝。
	Status *string `json:"status,omitempty"`
}
