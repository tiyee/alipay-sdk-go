package sign

import (
	"encoding/json"
)

// PdeductSignAddResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignAddResponse struct {
	// AgreementId 是支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// AgreementStatus 是支付宝协议状态。签约成功则返回success
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// ExtendField 是扩展参数，可为空
	ExtendField *string `json:"extend_field,omitempty"`
	// NotifyConfig 是通知方式设置。
	NotifyConfig *string `json:"notify_config,omitempty"`
	// OutAgreementId 是商户生成的代扣协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
	// PayConfig 是支付方式设置
	PayConfig *[]string `json:"pay_config,omitempty"`
	// SignDate 是签约时间
	SignDate *string `json:"sign_date,omitempty"`
}

// PdeductSignCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignCancelResponse struct {
	// AgreementId 是支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// AgreementStatus 是支付宝协议状态。解约成功则返回success
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// OutAgreementId 是商户代扣协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
}

// PdeductSignQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignQueryResponse struct {
	// AgreementId 是协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// BillKey 是户号
	BillKey *string `json:"bill_key,omitempty"`
	// ChargeInst 是出账机构
	ChargeInst *string `json:"charge_inst,omitempty"`
	// OutAgreementId 是朗新协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
	// SignDate 是签约时间
	SignDate *string `json:"sign_date,omitempty"`
	// SignExpireDate 是协议过期时间
	SignExpireDate *string `json:"sign_expire_date,omitempty"`
	// SpecifiedSortChannels 是用户代扣协议设置的优先扣款渠道
	SpecifiedSortChannels *[]SpecifiedSortChannelParam `json:"specified_sort_channels,omitempty"`
	// UserId 是用户ID
	UserId *string `json:"user_id,omitempty"`
}

// PdeductSignValidateResponse 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignValidateResponse = json.RawMessage
