package benefit

// CardBenefitCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitCreateResponse struct {
	// BenefitId 是权益ID
	BenefitId *string `json:"benefit_id,omitempty"`
}

// CardBenefitDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitDeleteResponse struct {
	// Result 是权益删除结果；true成功，false失败
	Result *bool `json:"result,omitempty"`
}

// CardBenefitModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitModifyResponse struct {
	// Result 是权益修改结果；true成功：false失败
	Result *bool `json:"result,omitempty"`
}

// CardBenefitQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitQueryResponse struct {
	// McardTemplateBenefitQuery 是会员卡模板外部权益列表
	McardTemplateBenefitQuery *[]McardTemplateBenefitQuery `json:"mcard_template_benefit_query,omitempty"`
}
