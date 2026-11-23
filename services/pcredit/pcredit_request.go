// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package pcredit

// PcreditHuabeiAuthAgreementCloseRequest 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthAgreementCloseRequest struct {
	// AgreementNo 是花芝协议号，支付宝系统中用以唯一标识用户签约记录的编号，即花芝协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty"`
	// OutRequestNo 是外部请求号
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PartnerId 是商户ID
	PartnerId *string `json:"partner_id,omitempty"`
	// QuitType 是用户主动意愿退出：USER_CANCEL_QUIT; 商户结算退出：SETTLE_APPLY_QUIT; 默认值为SETTLE_APPLY_QUIT；这个字段会影响用户在芝麻信用合约的状态
	QuitType *string `json:"quit_type,omitempty"`
}

// PcreditHuabeiAuthOrderUnfreezeRequest 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthOrderUnfreezeRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// NeedTerminated 是true表示需要解约。false或者不填写表示仅解冻。
	NeedTerminated *string `json:"need_terminated,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty"`
	// OrderTitle 是商户业务订单的简单描述，如商品名称等，长度不超过100个字母或50个汉字
	OrderTitle *string `json:"order_title,omitempty"`
	// OutRequestNo 是商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// SellerId 是商户的支付宝用户id。如果该值为空，则默认为商户签约账号对应的支付宝用户ID。
	SellerId *string `json:"seller_id,omitempty"`
	// UnfreezeAmount 是需要解冻的金额，单位为：元（人民币），精确到小数点后两位
	UnfreezeAmount *string `json:"unfreeze_amount,omitempty"`
}

// PcreditHuabeiAuthSettleApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type PcreditHuabeiAuthSettleApplyRequest struct {
	// ActionType 是操作类型，默认为 MERCHANT_SETTLE（结算）。枚举值如下：	*  MERCHANT_SETTLE：结算（需要主动退出在发起结算）;	*  QUIT_SETTLE：退出协议（解约协议）;	*  PERIOD_SETTLE：分阶段结算（不解约协议）。
	ActionType *string `json:"action_type,omitempty"`
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayUserId 是用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	AlipayUserId *string      `json:"alipay_user_id,omitempty"`
	ExtendParams *ExtraParams `json:"extend_params,omitempty"`
	// NeedTerminated 是是否解约，该参数控制结算完成之后的操作，为空则表示结算后解约。枚举值如下：	*  true：代表解约，默认为true。	*  false：代表不解约。
	NeedTerminated *string `json:"need_terminated,omitempty"`
	// OpenId 是用户在支付宝的唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// OutRequestNo 是商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayAmount 是需要支付的金额，单位为：元（人民币），精确到小数点后两位
	PayAmount *string `json:"pay_amount,omitempty"`
	// SellerId 是支付宝商家ID，即商家账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。	说明：若该值为空，则默认为商户签约账号对应的支付宝用户ID。
	SellerId *string `json:"seller_id,omitempty"`
}

// PcreditHuabeiAuthAccumulationQueryRequest is the query request model for alipay.pcredit.huabei.auth.accumulation.query.
type PcreditHuabeiAuthAccumulationQueryRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// Period 是协议结算周期，使用支付宝端回传的周期信息。
	Period *string `json:"period,omitempty" query:"period,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty" query:"alipay_user_id,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}

// PcreditHuabeiAuthAgreementQueryRequest is the query request model for alipay.pcredit.huabei.auth.agreement.query.
type PcreditHuabeiAuthAgreementQueryRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员。传入该参数，会忽略其它所有参数。
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty" query:"alipay_user_id,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// OutSignNo 是外部签约号，由商户提供，花芝轻会员协议中标识用户的唯一签约号（确保在商户系统中唯一）。
	OutSignNo *string `json:"out_sign_no,omitempty" query:"out_sign_no,omitempty"`
	// AuthScene 是花芝轻会员签约场景，商户和支付宝签约时确定。该值需要与系统/页面签约接口调用时传入的值保持一致。
	AuthScene *string `json:"auth_scene,omitempty" query:"auth_scene,omitempty"`
}

// PcreditHuabeiAuthOrderQueryRequest is the query request model for alipay.pcredit.huabei.auth.order.query.
type PcreditHuabeiAuthOrderQueryRequest struct {
	// AuthOptId 是支付宝侧花呗冻结、解冻操作单据id。在原先的冻结或者解冻接口调用中同步返回给商户，或者通过商户通知返回给商户。按订单号查询时，此字段不可为空。
	AuthOptId *string `json:"auth_opt_id,omitempty" query:"auth_opt_id,omitempty"`
	// AlipayUserId 是买家在支付宝的用户id。通过userid+请求流水号组合查询时，此字段不可为空。
	AlipayUserId *string `json:"alipay_user_id,omitempty" query:"alipay_user_id,omitempty"`
	// OpenId 是买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// OutRequestNo 是商户原先调用冻结、解冻接口传入的请求流水号。按照流水号查询订单时，此字段不能为空。
	OutRequestNo *string `json:"out_request_no,omitempty" query:"out_request_no,omitempty"`
}
