package template

// CardTemplateBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateBatchQueryResponse struct {
	// McardTemplate 是会员卡模板基本信息
	McardTemplate *[]McardTemplate `json:"mcard_template,omitempty"`
	// TemplateTotal 是会员卡模板总数,如10个卡模版，单位：/个
	TemplateTotal *int64 `json:"template_total,omitempty"`
}

// CardTemplateCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateCreateResponse struct {
	// TemplateId 是支付宝卡模板ID
	TemplateId *string `json:"template_id,omitempty"`
}

// CardTemplateModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateModifyResponse struct {
	// TemplateId 是模板ID
	TemplateId *string `json:"template_id,omitempty"`
}

// CardTemplateQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateQueryResponse struct {
	// AccessVersion 是接入版本，接入的是基础版本一定会有值
	AccessVersion *string `json:"access_version,omitempty"`
	// BizNoPrefix 是业务卡号前缀，由商户指定		支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoPrefix *string `json:"biz_no_prefix,omitempty"`
	// BizNoSuffixLen 是业务卡号后缀的长度		支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoSuffixLen *string `json:"biz_no_suffix_len,omitempty"`
	// CardActionList 是"卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多4个行动点。"
	CardActionList *[]TemplateActionInfoDTO `json:"card_action_list,omitempty"`
	// CardLevelConfs 是卡等级配置
	CardLevelConfs *[]TemplateCardLevelConfDTO `json:"card_level_confs,omitempty"`
	// CardSpecTag 是卡特定标签，只供特定业务使用，通常接入无需关注
	CardSpecTag *string `json:"card_spec_tag,omitempty"`
	// CardType 是会员卡类型： 	OUT_MEMBER_CARD：外部权益卡
	CardType *string `json:"card_type,omitempty"`
	// ColumnInfoList 是栏位信息（卡包详情页面展现的栏位）
	ColumnInfoList *[]TemplateColumnInfoDTO `json:"column_info_list,omitempty"`
	// FieldRuleList 是"字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE"
	FieldRuleList     *[]TemplateFieldRuleDTO       `json:"field_rule_list,omitempty"`
	MdcodeNotifyConf  *TemplateMdcodeNotifyConfDTO  `json:"mdcode_notify_conf,omitempty"`
	OpenCardConf      *TemplateOpenCardConfDTO      `json:"open_card_conf,omitempty"`
	PaidOuterCardConf *PaidOuterCardTemplateConfDTO `json:"paid_outer_card_conf,omitempty"`
	// PubChannels 是卡模板投放渠道
	PubChannels *[]PubChannelDTO `json:"pub_channels,omitempty"`
	// ServiceLabelList 是服务标签列表
	ServiceLabelList *[]string `json:"service_label_list,omitempty"`
	// ShopIds 是支付宝门店id
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// SpiAppId 是spi应用id，设置的实现spi.alipay.user.opencard.get接口的app_id
	SpiAppId *string `json:"spi_app_id,omitempty"`
	// TemplateBenefitInfo 是"权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益"
	TemplateBenefitInfo *[]TemplateBenefitInfoDTO `json:"template_benefit_info,omitempty"`
	TemplateFormConfig  *TemplateFormConfig       `json:"template_form_config,omitempty"`
	TemplateStyleInfo   *TemplateStyleInfoDTO     `json:"template_style_info,omitempty"`
}
