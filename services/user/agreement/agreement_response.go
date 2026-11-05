package agreement

import (
	"encoding/json"
)

// UserAgreementExecutionPlanModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementExecutionPlanModifyResponse struct {
	// AgreementNo 是周期性扣款产品，授权免密支付协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// DeductTime 是商户下一次扣款时间，格式 "yyyy-MM-dd"。	例如：用户在1月1日开通了连续包月，使用了10天又另行购买了“季度包”，如果此时商户希望“季度包”立即优先生效，在季度包结束后能继续使用连续包月，那么原定的周期就被延后了。此时可以通过本接口将预计扣款时间推后“季度包”的时长。
	DeductTime *string `json:"deduct_time,omitempty"`
}

// UserAgreementMigrateResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementMigrateResponse = json.RawMessage

// UserAgreementPermissionCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementPermissionCreateResponse struct {
	// DeductPermission 是商户代扣扣款许可
	DeductPermission *string `json:"deduct_permission,omitempty"`
	// EffectiveTimeEnd 是商户代扣扣款许可生效结束时间
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`
	// EffectiveTimeStart 是商户代扣扣款许可生效开始时间
	EffectiveTimeStart *string `json:"effective_time_start,omitempty"`
	// IsSuccessSend 是消息是否发送成功
	IsSuccessSend *bool `json:"is_success_send,omitempty"`
}

// UserAgreementQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementQueryResponse struct {
	// AgreementNo 是用户签约成功后的协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayLogonId 是返回脱敏的支付宝账号
	AlipayLogonId *string `json:"alipay_logon_id,omitempty"`
	// CreditAuthMode 是授信模式，目前只在花芝代扣（即花芝go）协议时才会返回
	CreditAuthMode *string `json:"credit_auth_mode,omitempty"`
	// DeviceId 是设备Id
	DeviceId *string `json:"device_id,omitempty"`
	// ExecutionPlans 是还款计划列表
	ExecutionPlans *[]ExecutionPlan `json:"execution_plans,omitempty"`
	// ExternalAgreementNo 是代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)。当入参中传了此参数时返回。
	ExternalAgreementNo *string `json:"external_agreement_no,omitempty"`
	// ExternalLogonId 是外部登录Id。当入参中传了此参数时返回。
	ExternalLogonId *string `json:"external_logon_id,omitempty"`
	// InvalidTime 是协议失效时间，格式为 yyyy-MM-dd HH:mm:ss。
	InvalidTime *string `json:"invalid_time,omitempty"`
	// LastDeductTime 是周期扣协议，上次扣款成功时间
	LastDeductTime *string `json:"last_deduct_time,omitempty"`
	// NextDeductTime 是周期扣协议，预计下次扣款时间
	NextDeductTime *string `json:"next_deduct_time,omitempty"`
	// PersonalProductCode 是协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。
	PersonalProductCode *string `json:"personal_product_code,omitempty"`
	// PricipalType 是签约主体类型。
	PricipalType *string `json:"pricipal_type,omitempty"`
	// PrincipalId 是签约主体标识。	当principal_type为CARD 时，该字段为支付宝用户号;
	PrincipalId *string `json:"principal_id,omitempty"`
	// PrincipalOpenId 是签约主体标识。	当principal_type为CARD 时，该字段为支付宝用户号;
	PrincipalOpenId *string `json:"principal_open_id,omitempty"`
	// SignScene 是签约协议的场景。
	SignScene *string `json:"sign_scene,omitempty"`
	// SignTime 是协议签约时间，格式为 yyyy-MM-dd HH:mm:ss。
	SignTime *string `json:"sign_time,omitempty"`
	// SingleQuota 是单笔代扣额度
	SingleQuota *string `json:"single_quota,omitempty"`
	// Status 是协议当前状态	1. TEMP：协议未生效，可以重新签约；	2. NORMAL：协议生效
	Status *string `json:"status,omitempty"`
	// ThirdPartyType 是签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。	默认为PARTNER。
	ThirdPartyType *string `json:"third_party_type,omitempty"`
	// ValidTime 是协议生效时间，格式为 yyyy-MM-dd HH:mm:ss。
	ValidTime *string `json:"valid_time,omitempty"`
	// ZmOpenId 是用户的芝麻信用 openId，供商 户查询用户芝麻信用使用。
	ZmOpenId *string `json:"zm_open_id,omitempty"`
}

// UserAgreementTransferResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementTransferResponse struct {
	// Amount 是单次金额限制，单位为元
	Amount *string `json:"amount,omitempty"`
	// ExecuteTime 是执行时间
	ExecuteTime *string `json:"execute_time,omitempty"`
	// Period 是周期
	Period *string `json:"period,omitempty"`
	// PeriodType 是周期类型
	PeriodType *string `json:"period_type,omitempty"`
	// TotalAmount 是总金额限制，单位为元
	TotalAmount *string `json:"total_amount,omitempty"`
	// TotalPayments 是总支付次数
	TotalPayments *string `json:"total_payments,omitempty"`
}

// UserAgreementUnsignResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementUnsignResponse = json.RawMessage

// UserAgreementUserVerifyApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementUserVerifyApplyResponse struct {
	// ApplyToken 是用户验证申请token，其格式和内容，由支付宝定义。在鉴权确认的过程中，商户可根据申请操作成功时返回的申请token，获取后续的买家结果
	ApplyToken *string `json:"apply_token,omitempty"`
}

// UserAgreementUserVerifyQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type UserAgreementUserVerifyQueryResponse struct {
	// UserMatch 是用户信息是否匹配
	UserMatch *bool `json:"user_match,omitempty"`
}
