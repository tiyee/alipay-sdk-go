// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package operation

// OperationActivityMerchantModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type OperationActivityMerchantModifyRequest struct {
	// ConsumptionThreshold 是消费门槛，单位分
	ConsumptionThreshold *int64 `json:"consumption_threshold,omitempty"`
	// DiscountAmount 是优惠金额，单位分
	DiscountAmount *int64 `json:"discount_amount,omitempty"`
	// Type 是NORMAL 常规类型，不限制, RESTRICTED 限制报名条件
	Type *string `json:"type,omitempty"`
}

// OperationActivityMerchantSignRequest 是支付宝官方 V3 OpenAPI 模型。
type OperationActivityMerchantSignRequest struct {
	// ConsumptionThreshold 是消费门槛（单位分）
	ConsumptionThreshold *int64 `json:"consumption_threshold,omitempty"`
	// DiscountAmount 是优惠金额，单位分
	DiscountAmount *int64 `json:"discount_amount,omitempty"`
	// Type 是报名类型，NORMAL 常规类型，不限制, RESTRICTED 限制报名条件
	Type *string `json:"type,omitempty"`
}

// OperationActivityMerchantUnsignRequest is the request model for alipay.commerce.operation.activity.merchant.unsign.
type OperationActivityMerchantUnsignRequest struct{}

// OperationBankActivityQueryRequest is the query request model for alipay.commerce.operation.bank.activity.query.
type OperationBankActivityQueryRequest struct {
	// OpenId 是支付宝openId
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// UserId 是蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
}
