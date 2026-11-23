package agreement

// UserAgreementExecutionPlanModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementExecutionPlanModifyRequest struct {
	// AgreementNo 是周期性扣款产品，授权免密支付协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// DeductTime 是商户下一次扣款时间
	DeductTime *string `json:"deduct_time,omitempty"`
	// Memo 是具体修改原因
	Memo *string `json:"memo,omitempty"`
}

// UserAgreementMigrateRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementMigrateRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ）
	AgreementNo *string `json:"agreement_no,omitempty"`
	// TargetAppId 是需要迁移的目标商户应用id
	TargetAppId *string `json:"target_app_id,omitempty"`
	// TargetInvokeAppId 是需要迁移的目标商户调用应用id
	TargetInvokeAppId *string `json:"target_invoke_app_id,omitempty"`
	// TargetPartnerId 是需要迁移的目标商户id
	TargetPartnerId *string `json:"target_partner_id,omitempty"`
}

// UserAgreementPermissionCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementPermissionCreateRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ）
	AgreementNo    *string            `json:"agreement_no,omitempty"`
	BusinessParams *BusinessParamsMap `json:"business_params,omitempty"`
	// NoticeTemplate 是目前共两种类型SERVIOCE_NOTICE和DEFAULT_PERMISSION，如果是服务变更提醒，需传入SERVIOCE_NOTICE
	NoticeTemplate *string `json:"notice_template,omitempty"`
	// OutRequestNo 是商户请求号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// TotalAmount 是订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	TotalAmount *string `json:"total_amount,omitempty"`
}

// UserAgreementTransferRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementTransferRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略
	AgreementNo      *string           `json:"agreement_no,omitempty"`
	PeriodRuleParams *PeriodRuleParams `json:"period_rule_params,omitempty"`
	// TargetProductCode 是协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。这里指的是需要修改目标产品码的值
	TargetProductCode *string `json:"target_product_code,omitempty"`
}

// UserAgreementUnsignRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementUnsignRequest struct {
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略	。	本参数与 external_agreement_no 不可同时为空。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayLogonId 是用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_user_id 不可同时为空，若都填写，则以alipay_user_id 为准。
	AlipayLogonId *string `json:"alipay_logon_id,omitempty"`
	// AlipayOpenId 是用户的支付宝账号对应的支付宝唯一用户号，	本参数与alipay_logon_id 不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id。
	AlipayOpenId *string `json:"alipay_open_id,omitempty"`
	// AlipayUserId 是用户的支付宝账号对应的支付宝唯一用户号，以2088 开头的 16 位纯数字 组成; 	本参数与alipay_logon_id 不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id。
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// ExtendParams 是扩展参数
	ExtendParams *string `json:"extend_params,omitempty"`
	// ExternalAgreementNo 是代扣协议中标示用户的唯一签约号(确保在商户系统中	唯一)。	格式规则:支持大写小写字母和数字，最长 32 位。	注意：若调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时传入 external_agreement_no 则该值必填且需与签约接口传入值相同。
	ExternalAgreementNo *string `json:"external_agreement_no,omitempty"`
	// NotifyUrl 是通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OperateType 是注意：仅异步解约需传入，其余情况无需传递本参数。
	OperateType *string `json:"operate_type,omitempty"`
	// PersonalProductCode 是协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。
	PersonalProductCode *string `json:"personal_product_code,omitempty"`
	// SignScene 是签约协议场景，该值需要与系统/页面签约接口调用时传入的值保持一 致。如：周期扣款场景，需与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。	当传入商户签约号 external_agreement_no时，场景不能为空或默认值 DEFAULT|DEFAULT。
	SignScene *string `json:"sign_scene,omitempty"`
	// ThirdPartyType 是签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。
	ThirdPartyType *string `json:"third_party_type,omitempty"`
}

// UserAgreementUserVerifyApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementUserVerifyApplyRequest struct {
	// AgreementNo 是协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
}

// UserAgreementQueryRequest is the query request model for alipay.user.agreement.query.
type UserAgreementQueryRequest struct {
	// PersonalProductCode 是协议产品码，商户和支付宝签约时确定，商户可咨询技术支持。
	PersonalProductCode *string `json:"personal_product_code,omitempty" query:"personal_product_code,omitempty"`
	// AlipayUserId 是用户的支付宝账号对应 的支付宝唯一用户号，以 2088 开头的 16 位纯数字 组成。	本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
	AlipayUserId *string `json:"alipay_user_id,omitempty" query:"alipay_user_id,omitempty"`
	// AlipayOpenId 是用户的支付宝账号对应 的支付宝唯一用户号，	本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
	AlipayOpenId *string `json:"alipay_open_id,omitempty" query:"alipay_open_id,omitempty"`
	// AlipayLogonId 是用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_open_id 或 alipay_user_id 同时填写，优先按照 alipay_open_id 或 alipay_user_id 处理。
	AlipayLogonId *string `json:"alipay_logon_id,omitempty" query:"alipay_logon_id,omitempty"`
	// SignScene 是签约场景码，该值需要与系统/页面签约接口调用时传入的值保持一 致。如：周期扣款场景与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。 	注意：当传入商户签约号 external_agreement_no 时，该值不能为空或默认值 DEFAULT|DEFAULT。
	SignScene *string `json:"sign_scene,omitempty" query:"sign_scene,omitempty"`
	// ExternalAgreementNo 是代扣协议中标示用户的唯一签约号(确保在商户系统中	唯一)。	格式规则:支持大写小写字母和数字，最长 32 位。
	ExternalAgreementNo *string `json:"external_agreement_no,omitempty" query:"external_agreement_no,omitempty"`
	// ThirdPartyType 是签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。	默认为PARTNER。
	ThirdPartyType *string `json:"third_party_type,omitempty" query:"third_party_type,omitempty"`
	// AgreementNo 是支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
}

// UserAgreementUserVerifyQueryRequest is the query request model for alipay.user.agreement.userverify.query.
type UserAgreementUserVerifyQueryRequest struct {
	// AgreementNo 是协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// ApplyToken 是申请时获取的token
	ApplyToken *string `json:"apply_token,omitempty" query:"apply_token,omitempty"`
}
