package pdeduct

// PdeductAsyncPayResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductAsyncPayResponse struct {
	// AgreementId 是支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// OutOrderNo 是商户代扣业务流水
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// ResultStatus 是扣款流水状态
	ResultStatus *string `json:"result_status,omitempty"`
}

// PdeductBillPayStatusResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductBillPayStatusResponse struct {
	// AgreementId 是支付宝协议流水
	AgreementId *string `json:"agreement_id,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// OrderNo 是支付宝流billNo
	OrderNo *string `json:"order_no,omitempty"`
	// OrderResultCode 是订单的结果码
	OrderResultCode *string `json:"order_result_code,omitempty"`
	// OrderResultMsg 是订单的结果描述
	OrderResultMsg *string `json:"order_result_msg,omitempty"`
	// OutOrderNo 是外部订单流水
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// Status 是支付宝订单支付状态。 	0：未知状态。 	1：支付成功。 	2：支付失败。
	Status *string `json:"status,omitempty"`
}

// PdeductPayResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductPayResponse struct {
	// AgreementId 是支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// BillNo 是支付宝订单流水号
	BillNo *string `json:"bill_no,omitempty"`
	// ExtendField 是扩展参数
	ExtendField *string `json:"extend_field,omitempty"`
	// OutOrderNo 是商户代扣业务流水
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// ResultCode 是针对于支付失败时，给的对应错误明细，如果判断外围的错误码是INVOKE_PAYACCEPTANCE_EXCEPTION需要近一步再结合着结果模型中的result_code, result_msg来判断
	ResultCode *string `json:"result_code,omitempty"`
	// ResultMsg 是针对于支付失败时，给的对应错误明细，如果判断外围的错误码是INVOKE_PAYACCEPTANCE_EXCEPTION需要近一步再结合着结果模型中的result_code, result_msg来判断
	ResultMsg *string `json:"result_msg,omitempty"`
	// ResultStatus 是订单支付状态。 	0：未知 	1：成功 	2：失败
	ResultStatus *string `json:"result_status,omitempty"`
}
