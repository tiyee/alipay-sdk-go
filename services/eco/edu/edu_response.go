// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package edu

// EduKtBillingModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtBillingModifyResponse struct {
	// Status 是如果成功，返回Y
	Status *string `json:"status,omitempty"`
}

// EduKtBillingQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtBillingQueryResponse struct {
	// OrderStatus 是NOT_PAY  待缴费 	PAYING  支付中 	PAY_SUCCESS 支付成功，处理中 	BILLING_SUCCESS 缴费成功 	TIMEOUT_CLOSED 逾期关闭账单 	ISV_CLOSED 账单关闭
	OrderStatus *string `json:"order_status,omitempty"`
	// OutTradeNo 是ISV发送账单时输入ISV端的账单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
}

// EduKtBillingSendResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtBillingSendResponse struct {
	// OrderNo 是支付宝－中小学－教育缴费的账单号
	OrderNo *string `json:"order_no,omitempty"`
}

// EduKtParentQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtParentQueryResponse struct {
	// Status 是如果用户已经添加过这个孩子，则返回SUCCESS ， 失败返回FAIL。
	Status *string `json:"status,omitempty"`
}

// EduKtSchoolInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtSchoolInfoModifyResponse struct {
	// SchoolNo 是支付宝分配给学校的编码，作为学校在支付宝的标识
	SchoolNo *string `json:"school_no,omitempty"`
	// Status 是Y：代表成功；N：代表失败
	Status *string `json:"status,omitempty"`
}

// EduKtStudentModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EduKtStudentModifyResponse struct {
	// Status 是Y：代表成功；N：代表失败
	Status *string `json:"status,omitempty"`
}
