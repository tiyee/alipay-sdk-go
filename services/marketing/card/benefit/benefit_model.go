package benefit

// McardTemplateBenefit 是支付宝官方 V3 OpenAPI 模型。
type McardTemplateBenefit struct {
	// BenefitDesc 是权益描述信息
	BenefitDesc *[]string `json:"benefit_desc,omitempty"`
	// EndDate 是权益结束时间。		注：在权益开始时间和结束时间范围内的权益才会认为是有效权益进行展示。
	EndDate *string `json:"end_date,omitempty"`
	// ExtInfo 是会员卡模板权益扩展信息：JSON格式; openUrl 说明：跳转到商户的优惠活动页面
	ExtInfo *string `json:"ext_info,omitempty"`
	// StartDate 是权益开始时间
	StartDate *string `json:"start_date,omitempty"`
	// TemplateId 是会员卡模板ID
	TemplateId *string `json:"template_id,omitempty"`
	// Title 是权益标题
	Title *string `json:"title,omitempty"`
}

// McardTemplateBenefitQuery 是支付宝官方 V3 OpenAPI 模型。
type McardTemplateBenefitQuery struct {
	// BenefitDesc 是权益描述信息
	BenefitDesc *string `json:"benefit_desc,omitempty"`
	// BenefitId 是权益ID
	BenefitId *string `json:"benefit_id,omitempty"`
	// EndDate 是权益结束时间
	EndDate *string `json:"end_date,omitempty"`
	// ExtInfo 是会员卡模板权益扩展信息：JSON格式; openUrl 说明：跳转到商户的优惠活动页面
	ExtInfo *string `json:"ext_info,omitempty"`
	// StartDate 是权益开始时间
	StartDate *string `json:"start_date,omitempty"`
	// TemplateId 是会员卡模板ID
	TemplateId *string `json:"template_id,omitempty"`
	// Title 是权益标题
	Title *string `json:"title,omitempty"`
}
