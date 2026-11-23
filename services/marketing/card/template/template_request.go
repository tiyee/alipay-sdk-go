package template

// CardTemplateBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateBatchQueryRequest struct {
	// PageNo 是查询起始页码，如第1页，单位：/页
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是每页数据大小，默认 20条，最大支持20个，单位：/个
	PageSize *int64 `json:"page_size,omitempty"`
}

// CardTemplateCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateCreateRequest struct {
	// AccessVersion 是接入版本
	AccessVersion *string `json:"access_version,omitempty"`
	// BizNoPrefix 是业务卡号前缀，由商户指定		支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoPrefix *string `json:"biz_no_prefix,omitempty"`
	// BizNoSuffixLen 是业务卡号后缀长度，与biz_no_prefix配合，扣除系统预留2位，剩下对应seq长度。在生成卡号时，若seq位数不足前置补0，若seq位数超出则以实际为准。举例：设为10，其中8位用于生成seq，可覆盖1亿用户，即使seq超过1亿，生成卡号也不报错，但总长度不得超过32位。建议按需设置合适的值，以获得长度一致的业务卡号，建议长度20，性能更好。
	BizNoSuffixLen *string `json:"biz_no_suffix_len,omitempty"`
	// CardActionList 是"卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多定制3个行动点。如果各项的配置内容完全一样，展示时会进行去重。"
	CardActionList *[]TemplateActionInfoDTO `json:"card_action_list,omitempty"`
	// CardLevelConf 是卡级别配置
	CardLevelConf *[]TemplateCardLevelConfDTO `json:"card_level_conf,omitempty"`
	// CardSpecTag 是卡特定标签，只供特定业务使用，通常接入无需关注
	CardSpecTag *string `json:"card_spec_tag,omitempty"`
	// CardType 是卡类型
	CardType *string `json:"card_type,omitempty"`
	// ColumnInfoList 是栏位信息
	ColumnInfoList *[]TemplateColumnInfoDTO `json:"column_info_list,omitempty"`
	// FieldRuleList 是"字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE"
	FieldRuleList    *[]TemplateFieldRuleDTO      `json:"field_rule_list,omitempty"`
	MdcodeNotifyConf *TemplateMdcodeNotifyConfDTO `json:"mdcode_notify_conf,omitempty"`
	// NotifyUrl 是通知地址
	NotifyUrl         *string                       `json:"notify_url,omitempty"`
	OpenCardConf      *TemplateOpenCardConfDTO      `json:"open_card_conf,omitempty"`
	PaidOuterCardConf *PaidOuterCardTemplateConfDTO `json:"paid_outer_card_conf,omitempty"`
	// PubChannels 是卡模板投放渠道
	PubChannels *[]PubChannelDTO `json:"pub_channels,omitempty"`
	// RequestId 是请求ID，商家自定义且并保证唯一性。
	RequestId *string `json:"request_id,omitempty"`
	// ServiceLabelList 是服务标签列表
	ServiceLabelList *[]string `json:"service_label_list,omitempty"`
	// ShopIds 是会员卡上架门店id（支付宝门店id），即发放会员卡的商家门店id
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// SpiAppId 是spi应用id，为实现spi.alipay.user.opencard.get接口的app_id。若是第三方代理模式，可以设置成服务商的 APPID 或者是商家自己的 APPID；	若不是第三方代理模式，只能设置商家自己的 APPID。
	SpiAppId *string `json:"spi_app_id,omitempty"`
	// TemplateBenefitInfo 是"权益信息， 1、在卡包的卡详情页面会自动添加“会员权益”栏位，展现此字段配置的会员卡权益信息， 2、如果添加门店渠道，则可在门店页展现会员卡的权益"
	TemplateBenefitInfo *[]TemplateBenefitInfoDTO `json:"template_benefit_info,omitempty"`
	TemplateFormConfig  *TemplateFormConfig       `json:"template_form_config,omitempty"`
	TemplateStyleInfo   *TemplateStyleInfoDTO     `json:"template_style_info,omitempty"`
	// WriteOffType 是卡包详情页面中展现出的卡码（可用于扫码核销）		(1) 静态码	qrcode: 二维码，扫码得商户开卡传入的external_card_no	barcode: 条形码，扫码得商户开卡传入的external_card_no		(2) 动态码-支付宝生成码值(动态码会在2分钟左右后过期)	dqrcode: 动态二维码，扫码得到的码值可配合会员卡查询接口使用	dbarcode: 动态条形码，扫码得到的码值可配合会员卡查询接口使用		(3) 动态码-商家自主生成码值（码值、时效性都由商户控制）	mdqrcode: 商户动态二维码，扫码得商户自主传入的码值	mdbarcode: 商户动态条码，扫码得商户自主传入的码值	如需使用商户动态码，对接详情参见 商户动态发码。
	WriteOffType *string `json:"write_off_type,omitempty"`
}

// CardTemplateModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CardTemplateModifyRequest struct {
	// AccessVersion 是接口接入版本
	AccessVersion *string `json:"access_version,omitempty"`
	// BizNoPrefix 是业务卡号前缀，由商户指定		支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀
	BizNoPrefix *string `json:"biz_no_prefix,omitempty"`
	// BizNoSuffixLen 是业务卡号后缀长度，与biz_no_prefix配合，扣除系统预留2位，剩下对应seq长度。在生成卡号时，若seq位数不足前置补0，若seq位数超出则以实际为准。举例：设为10，其中8位用于生成seq，可覆盖1亿用户，即使seq超过1亿，生成卡号也不报错，但总长度不得超过32位。建议按需设置合适的值，以获得长度一致的业务卡号，建议长度20，性能更好。单位：/位
	BizNoSuffixLen *string `json:"biz_no_suffix_len,omitempty"`
	// CardActionList 是"卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多定制4个行动点。"
	CardActionList *[]TemplateActionInfoDTO `json:"card_action_list,omitempty"`
	// CardLevelConf 是卡级别配置
	CardLevelConf *[]TemplateCardLevelConfDTO `json:"card_level_conf,omitempty"`
	// CardSpecTag 是卡特定标签，只供特定业务使用，通常接入无需关注
	CardSpecTag *string `json:"card_spec_tag,omitempty"`
	// ColumnInfoList 是栏位信息（卡包详情页面展现的栏位）
	ColumnInfoList *[]TemplateColumnInfoDTO `json:"column_info_list,omitempty"`
	// FieldRuleList 是"字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE 注意：商家会员卡场景不支持修改该内容。 "
	FieldRuleList     *[]TemplateFieldRuleDTO       `json:"field_rule_list,omitempty"`
	MdcodeNotifyConf  *TemplateMdcodeNotifyConfDTO  `json:"mdcode_notify_conf,omitempty"`
	OpenCardConf      *TemplateOpenCardConfDTO      `json:"open_card_conf,omitempty"`
	PaidOuterCardConf *PaidOuterCardTemplateConfDTO `json:"paid_outer_card_conf,omitempty"`
	// PubChannels 是卡模板投放渠道
	PubChannels *[]PubChannelDTO `json:"pub_channels,omitempty"`
	// RequestId 是请求 ID，商户自定义且需确保每次请求唯一。
	RequestId *string `json:"request_id,omitempty"`
	// ShopIds 是会员卡上架门店 id（支付宝门店id），即发放会员卡的商家门店id。不传则保持创建模板时信息。	注意：不支持删除已有shop_id，仅支持新增。
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// SpiAppId 是spi应用id，为实现spi.alipay.user.opencard.get接口的app_id。若是第三方代理模式，可以设置成服务商的 APPID 或者是商家自己的 APPID；	若不是第三方代理模式，只能设置商家自己的 APPID。
	SpiAppId *string `json:"spi_app_id,omitempty"`
	// TemplateBenefitInfo 是"权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益"
	TemplateBenefitInfo *[]TemplateBenefitInfoDTO `json:"template_benefit_info,omitempty"`
	TemplateFormConfig  *TemplateFormConfig       `json:"template_form_config,omitempty"`
	// TemplateId 是会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。
	TemplateId        *string               `json:"template_id,omitempty"`
	TemplateStyleInfo *TemplateStyleInfoDTO `json:"template_style_info,omitempty"`
	// WriteOffType 是卡包详情页面中展现出的卡码（可用于扫码核销）		(1) 静态码	qrcode: 二维码，扫码得商户开卡传入的external_card_no	barcode: 条形码，扫码得商户开卡传入的external_card_no		(2) 动态码-支付宝生成码值(动态码会在2分钟左右后过期)	dqrcode: 动态二维码，扫码得到的码值可配合会员卡查询接口使用	dbarcode: 动态条形码，扫码得到的码值可配合会员卡查询接口使用		(3) 动态码-商家自主生成码值（码值、时效性都由商户控制）	mdqrcode: 商户动态二维码，扫码得商户自主传入的码值	mdbarcode: 商户动态条码，扫码得商户自主传入的码值。	如需使用商户动态码，请联系支付宝技术支持获取相关文档。
	WriteOffType *string `json:"write_off_type,omitempty"`
}

// CardTemplateQueryRequest is the query request model for alipay.marketing.card.template.query.
type CardTemplateQueryRequest struct {
	// TemplateId 是会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
}
