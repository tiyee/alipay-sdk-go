package benefit

// CardBenefitCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitCreateRequest struct {
	McardTemplateBenefit *McardTemplateBenefit `json:"mcard_template_benefit,omitempty"`
}

// CardBenefitModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CardBenefitModifyRequest struct {
	// BenefitId 是权益ID，通过 alipay.marketing.card.benefit.create(会员卡模板外部权益创建)接口创建获取。
	BenefitId            *string               `json:"benefit_id,omitempty"`
	McardTemplateBenefit *McardTemplateBenefit `json:"mcard_template_benefit,omitempty"`
}

// CardBenefitDeleteRequest is the query request model for alipay.marketing.card.benefit.delete.
type CardBenefitDeleteRequest struct {
	// TemplateId 是会员卡模板ID，通过 alipay.marketing.card.template.create（会员卡模板创建)接口创建会员卡模板获取。
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
	// BenefitId 是权益ID，通过 alipay.marketing.card.benefit.create(会员卡模板外部权益创建)接口创建获取。
	BenefitId *string `json:"benefit_id,omitempty" query:"benefit_id,omitempty"`
}

// CardBenefitQueryRequest is the query request model for alipay.marketing.card.benefit.query.
type CardBenefitQueryRequest struct {
	// TemplateId 是会员卡模板ID，通过 alipay.marketing.card.template.create（会员卡模板创建)接口创建会员卡模板获取。
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
	// BenefitId 是权益ID，通过 alipay.marketing.card.benefit.create(会员卡模板外部权益创建)接口创建获取。
	BenefitId *string `json:"benefit_id,omitempty" query:"benefit_id,omitempty"`
}
