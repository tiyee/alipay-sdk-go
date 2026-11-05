package agreement

// FundAgreementQuotaModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAgreementQuotaModifyRequest struct {
	// AgreementQuotaModifyList 是协议额度设置信息。
	AgreementQuotaModifyList *[]AgreementQuotaModifyList `json:"agreement_quota_modify_list,omitempty"`
	// BizScene 是场景码，非固定值，使用的对应的资金场景码保持一致。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是额度产品码，非固定值，使用的对应的资金产品码保持一致。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundAgreementQuotaQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAgreementQuotaQueryRequest struct {
	// AgreementNoList 是授权协议号。
	AgreementNoList *[]string `json:"agreement_no_list,omitempty"`
	// BizScene 是场景码，非固定值，使用的对应的资金场景码保持一致。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是额度产品码，非固定值，使用的对应的资金产品码保持一致。
	ProductCode *string `json:"product_code,omitempty"`
}
