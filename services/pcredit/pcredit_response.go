// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package pcredit

// PcreditHuabeiAuthAccumulationQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthAccumulationQueryResponse struct {
	// PayAmount 是本周期内支付宝端根据用户消费情况试算出的应付费用，仅供参考使用。
	PayAmount *string `json:"pay_amount,omitempty"`
	// TotalDiscountAmount 是本周期内用户累计享受的优惠金额
	TotalDiscountAmount *string `json:"total_discount_amount,omitempty"`
	// TotalPayCount 是本周期内用户总的消费次数
	TotalPayCount *string `json:"total_pay_count,omitempty"`
	// TotalRealPayAmount 是本周期内用户累计支付宝付款金额
	TotalRealPayAmount *string `json:"total_real_pay_amount,omitempty"`
}

// PcreditHuabeiAuthAgreementCloseResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthAgreementCloseResponse struct {
	// OutRequestNo 是外部请求号
	OutRequestNo *string `json:"out_request_no,omitempty"`
}

// PcreditHuabeiAuthAgreementQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthAgreementQueryResponse struct {
	// AgreementName 是协议名称
	AgreementName *string `json:"agreement_name,omitempty"`
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AgreementStatus 是协议状态。Y表示状态有效，N表示状态失效
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// AlipayUserId 是支付宝用户userId
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// AuthScene 是花芝轻会员签约场景，商户和支付宝签约时确定，商户接入时需要咨询技术支持。
	AuthScene *string `json:"auth_scene,omitempty"`
	// ExternalLogonId 是用户在商户网站的登录账号
	ExternalLogonId *string `json:"external_logon_id,omitempty"`
	// GmtSign 是签约时间。如果是在签状态，返回签约时间。
	GmtSign *string `json:"gmt_sign,omitempty"`
	// GmtUnsign 是花芝轻会员协议解约时间。如果是解约状态，才返回该字段。
	GmtUnsign *string `json:"gmt_unsign,omitempty"`
	// OpenId 是支付宝用户userId
	OpenId *string `json:"open_id,omitempty"`
	// OutBizType 是外部业务类型，用于区分本次业务交互在商户侧具体的业务类别，提供给商家进行特定逻辑的处理
	OutBizType *string `json:"out_biz_type,omitempty"`
	// OutSignNo 是外部签约号，由商户提供，花芝轻会员协议中标识用户的唯一签约号（确保在商户系统中唯一）。由商户签约时传入，最终返回给商户。
	OutSignNo *string `json:"out_sign_no,omitempty"`
	// RestFreezeAmount 是用户在本花芝轻会员协议中，剩余的总冻结额度（资金池总余额），数值能实时准确，可以用于核对。两位小数，单位元。
	RestFreezeAmount *string `json:"rest_freeze_amount,omitempty"`
}

// PcreditHuabeiAuthOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthOrderQueryResponse struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayUserId 是支付宝用户userId
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// AuthOptId 是支付宝侧花呗冻结解冻操作单据id
	AuthOptId *string `json:"auth_opt_id,omitempty"`
	// GmtTrans 是业务操作完成时间
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OpenId 是支付宝用户userId
	OpenId *string `json:"open_id,omitempty"`
	// OrderTitle 是商户业务订单的简单描述，如商品名称等，长度不超过100个字母或50个汉字
	OrderTitle *string `json:"order_title,omitempty"`
	// OutRequestNo 是商户本次操作的请求流水号
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RestFreezeAmount 是完成本次操作时，用户资金池余额快照。仅作提示用，请勿用于核对，并发情况下数值有可能不准确。两位小数，单位元。
	RestFreezeAmount *string `json:"rest_freeze_amount,omitempty"`
	// SellerId 是商户的支付宝用户id
	SellerId *string `json:"seller_id,omitempty"`
	// TransAmount 是冻结或者解冻金额
	TransAmount *string `json:"trans_amount,omitempty"`
	// TransStatus 是业务操作状态，Y表示成功结束；N表示失败结束；I表示进行中。
	TransStatus *string `json:"trans_status,omitempty"`
	// TransType 是操作类型，FREEZE表示冻结操作；UNFREEZE表示解冻操作；PAY表示支付操作；SETTLE表示结算操作
	TransType *string `json:"trans_type,omitempty"`
}

// PcreditHuabeiAuthOrderUnfreezeResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthOrderUnfreezeResponse struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AgreementStatus 是协议状态。Y表示状态有效，N表示状态失效
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// AuthOptId 是支付宝侧花呗冻结解冻操作单据id
	AuthOptId *string `json:"auth_opt_id,omitempty"`
	// GmtTrans 是解冻成功时间
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty"`
	// OutRequestNo 是商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。由商户传入，最终返回给商户。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RestFreezeAmount 是完成本次操作时，用户资金池余额快照。仅作提示用，请勿用于核对，并发情况下数值有可能不准确。两位小数，单位元。
	RestFreezeAmount *string `json:"rest_freeze_amount,omitempty"`
	// SellerId 是商户的支付宝用户id
	SellerId *string `json:"seller_id,omitempty"`
	// TransStatus 是业务操作状态，Y表示成功；N表示失败。
	TransStatus *string `json:"trans_status,omitempty"`
	// UnfreezeAmount 是解冻金额
	UnfreezeAmount *string `json:"unfreeze_amount,omitempty"`
}

// PcreditHuabeiAuthSettleApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthSettleApplyResponse struct {
	// FailReason 是根据商户传入的轻会员协议结算周期内累计消费数据与支付宝端轻会员协议结算周期内消费数据比对之后，若数据不一致，则会返回对应的异常数据比对结果。
	FailReason *string `json:"fail_reason,omitempty"`
	// OutRequestNo 是商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
}
