// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package robot

// CcmRobotAvatarBaseQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmRobotAvatarBaseQueryRequest struct {
	// CommodityCode 是开通的数字人的商品码
	CommodityCode *string `json:"commodity_code,omitempty"`
	// MethodParams 是method_params:对应target_method请求方法的具体传参值k-v map
	MethodParams *string `json:"method_params,omitempty"`
	// TargetMethod 是targetMethod；数字人目标方法调用
	TargetMethod *string `json:"target_method,omitempty"`
	// TenantCode 是租户code&租户id，对于子部门客户需要上传子部门id
	TenantCode *string `json:"tenant_code,omitempty"`
}
