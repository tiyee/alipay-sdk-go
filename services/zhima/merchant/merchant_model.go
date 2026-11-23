// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// CumulateDataDetail is the official V3 CumulateDataDetail schema.
type CumulateDataDetail struct {
	// ActionType represents action_type. 数据回传的动作类型，正向或逆向。
	ActionType *string `json:"action_type,omitzero"`
	// BizTime represents biz_time. 回传数据发生的实际时间
	BizTime *string `json:"biz_time,omitzero"`
	// DataType represents data_type. 回传数据类型。
	DataType *string `json:"data_type,omitzero"`
	// DiscountAmount represents discount_amount. 优惠金额
	DiscountAmount *string `json:"discount_amount,omitzero"`
	// DiscountDesc represents discount_desc. 优惠描述
	DiscountDesc *string `json:"discount_desc,omitzero"`
	// OutBizNo represents out_biz_no. 数据回传时传入的外部业务号。
	OutBizNo *string `json:"out_biz_no,omitzero"`
	// ReferOutBizNo represents refer_out_biz_no. 数据回传时逆向单据记录的对应正向单据的外部业务号。
	ReferOutBizNo *string `json:"refer_out_biz_no,omitzero"`
	// SubDataType represents sub_data_type. 回传数据子类型。
	SubDataType *string `json:"sub_data_type,omitzero"`
	// TaskAmount represents task_amount. 任务金额
	TaskAmount *string `json:"task_amount,omitzero"`
	// TaskDesc represents task_desc. 任务描述
	TaskDesc *string `json:"task_desc,omitzero"`
	// TaskTimes represents task_times. 任务次数
	TaskTimes *int64 `json:"task_times,omitzero"`
}

// AmountTypeSyncData is the official V3 AmountTypeSyncData schema.
type AmountTypeSyncData struct {
	// DiscountAmount represents discount_amount. 商户回传的优惠金额，如用户享受的红包金额，单位元
	DiscountAmount *string `json:"discount_amount,omitzero"`
	// DiscountDesc represents discount_desc. 商户数据回传的优惠信息的名称。
	DiscountDesc *string `json:"discount_desc,omitzero"`
	// HasAlipayTrade represents has_alipay_trade. 是否为支付宝交易，默认false，表示是否要做支付宝交易关联。此时tradeNo必须传支付宝交易号。
	HasAlipayTrade *bool `json:"has_alipay_trade,omitzero"`
	// TaskAmount represents task_amount. 用户和商户发生交易的交易单金额，单位元。
	TaskAmount *string `json:"task_amount,omitzero"`
	// TaskDesc represents task_desc. 任务描述
	TaskDesc *string `json:"task_desc,omitzero"`
	// TradeNo represents trade_no. 用户和商户发生金额类交易的交易单号。
	TradeNo *string `json:"trade_no,omitzero"`
}

// DiscountTypeSyncData is the official V3 DiscountTypeSyncData schema.
type DiscountTypeSyncData struct {
	// DiscountAmount represents discount_amount. 商户回传的优惠金额，如用户享受的红包金额，单位元
	DiscountAmount *string `json:"discount_amount,omitzero"`
	// DiscountDesc represents discount_desc. 商户数据回传的优惠信息的名称。
	DiscountDesc *string `json:"discount_desc,omitzero"`
}

// TimesTypeSyncData is the official V3 TimesTypeSyncData schema.
type TimesTypeSyncData struct {
	// DiscountAmount represents discount_amount. 商户回传的优惠金额，如用户享受的红包金额，单位元
	DiscountAmount *string `json:"discount_amount,omitzero"`
	// DiscountDesc represents discount_desc. 商户数据回传的优惠信息的名称。
	DiscountDesc *string `json:"discount_desc,omitzero"`
	// TaskAmount represents task_amount. 用户和商户发生交易的交易单金额，单位元。
	TaskAmount *string `json:"task_amount,omitzero"`
	// TaskDesc represents task_desc. 任务描述
	TaskDesc *string `json:"task_desc,omitzero"`
	// TaskTimes represents task_times. 当为次数型任务时必须传。
	TaskTimes *int64 `json:"task_times,omitzero"`
}

// ZMGOBasicConfig is the official V3 ZMGOBasicConfig schema.
type ZMGOBasicConfig struct {
	// BizType represents biz_type. 商户在芝麻GO配置的业务身份编码
	BizType *string `json:"biz_type,omitzero"`
	// Contact represents contact. 商家客服电话
	Contact *string `json:"contact,omitzero"`
	// IsvPid represents isv_pid. 运营商商户支付宝ID。若非ISV代理模式，也就是商户自运营模式，此属性取值与partner_id一致。
	IsvPid *string `json:"isv_pid,omitzero"`
	// MerchantCustomLogo represents merchant_custom_logo. 商户LOGO
	MerchantCustomLogo *string `json:"merchant_custom_logo,omitzero"`
	// OutBizNo represents out_biz_no. 外部业务单号，供幂等使用，需保证每次请求的值都不同
	OutBizNo *string `json:"out_biz_no,omitzero"`
	// PartnerId represents partner_id. 商户的支付宝ID，即为此商户创建芝麻GO模板
	PartnerId *string `json:"partner_id,omitzero"`
	// TemplateName represents template_name. 芝麻GO模板名称
	TemplateName *string `json:"template_name,omitzero"`
	// TemplateNo represents template_no. 模板编号
	TemplateNo *string `json:"template_no,omitzero"`
}

// ZMGOExtConfig is the official V3 ZMGOExtConfig schema.
type ZMGOExtConfig struct {
	// TextContentFillRuleId represents text_content_fill_rule_id. 文本内容填充规则ID
	TextContentFillRuleId *string `json:"text_content_fill_rule_id,omitzero"`
	// TextContentFillVariable represents text_content_fill_variable. 文本内容填充变量，变量key由运营同学提供
	TextContentFillVariable *string `json:"text_content_fill_variable,omitzero"`
}

// ZMGOObligationConfig is the official V3 ZMGOObligationConfig schema.
type ZMGOObligationConfig struct {
	// BenefitUrl represents benefit_url. 芝麻GO管理页做任务按钮对应链接
	BenefitUrl *string `json:"benefit_url,omitzero"`
	// ObligationAmount represents obligation_amount. 完成任务所需的最少额度，当obligation_template取值为：OBTL002时必传
	ObligationAmount *string `json:"obligation_amount,omitzero"`
	// ObligationTemplate represents obligation_template. 承诺模式下，义务的类型，包含次数类型和金额类型： OBTL001("OBTL001", "在x时间内完成n次"), OBTL002("OBTL002", "在x时间内达到n元金额")
	ObligationTemplate *string `json:"obligation_template,omitzero"`
	// ObligationTimes represents obligation_times. 完成任务所需的最少次数，当obligation_template取值为：OBTL001时必填
	ObligationTimes *string `json:"obligation_times,omitzero"`
	// PromiseTypeDesc represents promise_type_desc. 任务描述
	PromiseTypeDesc *string `json:"promise_type_desc,omitzero"`
	// TaskProgressRedirectSchema represents task_progress_redirect_schema. 芝麻GO管理页任务进度重定向链接
	TaskProgressRedirectSchema *string `json:"task_progress_redirect_schema,omitzero"`
}

// ZMGOOpenConfig is the official V3 ZMGOOpenConfig schema.
type ZMGOOpenConfig struct {
	// ApplyButtonDesc represents apply_button_desc. 开通页按钮文案描述;传入内容后开通页点击开通按钮展示文案会使用该字段传入内容
	ApplyButtonDesc *string `json:"apply_button_desc,omitzero"`
	// AppointDate represents appoint_date. 固定到期时间或每月的第几日。 （1）当period_mode为APPOINT_DATE时，取值固定日期，如"2022-06-01 00:00:00"，表示协议在6月1号0点到期; （2）当period_mode为FIX_DATE，取值"3"，表示协议在period_time个月后的3号到期。
	AppointDate *string `json:"appoint_date,omitzero"`
	// CardColorScheme represents card_color_scheme. 芝麻GO卡颜色。 C_BLACK("C_BLACK", "黑色"), C_RED("C_RED", "红色"), C_GREEN("C_GREEN", "绿色"), C_LIGHT_BLUE("C_LIGHT_BLUE", "浅蓝色"), C_BLUE("C_BLUE", "蓝色"), C_ORANGE("C_ORANGE","橙色"), C_PURPLE("C_PURPLE", "紫色");
	CardColorScheme *string `json:"card_color_scheme,omitzero"`
	// CustomOpenTipList represents custom_open_tip_list. 未完成任务规则详情
	CustomOpenTipList *string `json:"custom_open_tip_list,omitzero"`
	// CustomOpenTips represents custom_open_tips. 规则详情
	CustomOpenTips *string `json:"custom_open_tips,omitzero"`
	// FreezeAmount represents freeze_amount. 担保金额额度，单位:元
	FreezeAmount *string `json:"freeze_amount,omitzero"`
	// MinSignInterval represents min_sign_interval. 用户签约间隔时间，单位天
	MinSignInterval *string `json:"min_sign_interval,omitzero"`
	// PeriodMode represents period_mode. 芝麻GO义务到期的计算方式。包括： RELATIVE_DATE("RELATIVE_DATE", "相对日期"), ABSOLATE_DATE("ABSOLATE_DATE", "绝对日期(当日开始计算)"), ABSOLUTE_PLUS_1_DATE("ABSOLUTE_PLUS_1_DATE", "绝对日期(次日开始计算)"), FIX_DATE("FIX_DATE", "固定日期"), APPOINT_DATE("APPOINT_DATE", "指定日期"),
	PeriodMode *string `json:"period_mode,omitzero"`
	// PeriodTime represents period_time. 周期时长 当period_mode为RELATIVE_DATE、ABSOLATE_DATE、ABSOLUTE_PLUS_1_DATE、FIX_DATE时必填
	PeriodTime *string `json:"period_time,omitzero"`
	// ShowSignSuccessPage represents show_sign_success_page. 是否展示签约完成页
	ShowSignSuccessPage *bool `json:"show_sign_success_page,omitzero"`
	// SignAgainSchema represents sign_again_schema. 用户签约商户活动到期
	SignAgainSchema *string `json:"sign_again_schema,omitzero"`
	// SignSuccessTaskButtonDesc represents sign_success_task_button_desc. 签约完成页任务按钮描述，当展示签约完成页=true的时候 改内容必填
	SignSuccessTaskButtonDesc *string `json:"sign_success_task_button_desc,omitzero"`
	// SupportExpireDeferral represents support_expire_deferral. 是否支持延期特权，开启后用户可以在任务到期前进行延期。(该功能已停用，不建议开启使用。)
	SupportExpireDeferral *bool `json:"support_expire_deferral,omitzero"`
}

// ZMGOQuitConfig is the official V3 ZMGOQuitConfig schema.
type ZMGOQuitConfig struct {
	// QuitDesc represents quit_desc. 退出提示
	QuitDesc *string `json:"quit_desc,omitzero"`
	// QuitJumpUrl represents quit_jump_url. 退出跳转商家页地址，当quit_type取值为：QUIT_MERCHANT_JUMP，必传
	QuitJumpUrl *string `json:"quit_jump_url,omitzero"`
	// QuitType represents quit_type. 退出芝麻GO方式，取值： （1）QUIT_COMMON("QUIT_COMMON", "标准退出"), （2）QUIT_MERCHANT_JUMP("QUIT_MERCHANT_JUMP", "商家页退出"), （3）QUIT_ONLY_TIPS("QUIT_ONLY_TIPS", "仅退出提示")
	QuitType *string `json:"quit_type,omitzero"`
}

// ZMGORightConfig is the official V3 ZMGORightConfig schema.
type ZMGORightConfig struct {
	// CumulativePreferentialRedirectSchema represents cumulative_preferential_redirect_schema. 芝麻GO管理页已享优惠进度的重定向链接
	CumulativePreferentialRedirectSchema *string `json:"cumulative_preferential_redirect_schema,omitzero"`
	// CustomBenefitDesc represents custom_benefit_desc. 权益描述
	CustomBenefitDesc *string `json:"custom_benefit_desc,omitzero"`
	// CustomSubBenefitDesc represents custom_sub_benefit_desc. 权益描述子标题
	CustomSubBenefitDesc *string `json:"custom_sub_benefit_desc,omitzero"`
}

// ZMGOSettlementConfig is the official V3 ZMGOSettlementConfig schema.
type ZMGOSettlementConfig struct {
	// CustomFeeName represents custom_fee_name. 会员费扣费名称
	CustomFeeName *string `json:"custom_fee_name,omitzero"`
	// CycleFlexWithholdConfig represents cycle_flex_withhold_config.
	CycleFlexWithholdConfig *ZMGOCycleFlexConfig `json:"cycle_flex_withhold_config,omitzero"`
	// CycleWithholdConfig represents cycle_withhold_config.
	CycleWithholdConfig *ZMGOCycleWithholdConfig `json:"cycle_withhold_config,omitzero"`
	// ExpStopDelayDays represents exp_stop_delay_days. 预计结算延迟天数，当exp_stop_time_mode取值RELATIVE_DATE、ABSOLATE_DATE、ABSOLUTE_PLUS_1_DATE时，必传
	ExpStopDelayDays *int64 `json:"exp_stop_delay_days,omitzero"`
	// ExpStopTime represents exp_stop_time. 预期结算时间，当exp_stop_time_mode取值APPOINT_DATE时，必传
	ExpStopTime *string `json:"exp_stop_time,omitzero"`
	// ExpStopTimeMode represents exp_stop_time_mode. 预计结算时间模式，取值：（1）RELATIVE_DATE("RELATIVE_DATE", "相对日期"), （2）ABSOLATE_DATE("ABSOLATE_DATE", "绝对日期(当日开始计算)"), （3）ABSOLUTE_PLUS_1_DATE("ABSOLUTE_PLUS_1_DATE", "绝对日期(次日开始计算)"), （4）APPOINT_DATE("APPOINT_DATE", "指定日期")
	ExpStopTimeMode *string `json:"exp_stop_time_mode,omitzero"`
	// FulfilltimesCustomSettlementPlan represents fulfilltimes_custom_settlement_plan. 按履约次数灵活扣款计划。当settlement_type取值为fulfillTimesCustomSettlement时，此属性必传
	FulfilltimesCustomSettlementPlan *string `json:"fulfilltimes_custom_settlement_plan,omitzero"`
	// SettlementType represents settlement_type. 结算类型，取值： BENEFIT_SETTLEMENT("benefitSettlement", "按权益金额结算"), FULFILL_TIMES_CUSTOM_SETTLEMENT("fulfillTimesCustomSettlement","按履约次数扣款")
	SettlementType *string `json:"settlement_type,omitzero"`
}

// ZMGOCycleFlexConfig is the official V3 ZMGOCycleFlexConfig schema.
type ZMGOCycleFlexConfig struct {
	// CycleFlexWithholdFeeName represents cycle_flex_withhold_fee_name. 周期灵活扣的费用名称
	CycleFlexWithholdFeeName *string `json:"cycle_flex_withhold_fee_name,omitzero"`
	// CycleFlexWithholdMaxPrice represents cycle_flex_withhold_max_price. 周期灵活扣单期扣款的最大额度
	CycleFlexWithholdMaxPrice *string `json:"cycle_flex_withhold_max_price,omitzero"`
	// CycleFlexWithholdTotalPeriodCount represents cycle_flex_withhold_total_period_count. 周期灵活扣的总期数
	CycleFlexWithholdTotalPeriodCount *int64 `json:"cycle_flex_withhold_total_period_count,omitzero"`
}

// ZMGOCycleWithholdConfig is the official V3 ZMGOCycleWithholdConfig schema.
type ZMGOCycleWithholdConfig struct {
	// DeductPlan represents deduct_plan. 扣款计划
	DeductPlan *[]string `json:"deduct_plan,omitzero"`
	// ExemptionPeriod represents exemption_period. 豁免期天数
	ExemptionPeriod *string `json:"exemption_period,omitzero"`
	// Period represents period. 周期
	Period *string `json:"period,omitzero"`
	// PeriodType represents period_type. 扣款周期类型，支持DAY-天模式，MONTH-月模式
	PeriodType *string `json:"period_type,omitzero"`
	// SupportCycleWithholdHighMode represents support_cycle_withhold_high_mode. 是否支持周期扣高级模式，高级模式下扣款计划中内一起的扣款金额可以进行定制
	SupportCycleWithholdHighMode *bool `json:"support_cycle_withhold_high_mode,omitzero"`
	// SupportExemptionPeriod represents support_exemption_period. 是否支持豁免期，如果支持豁免期，用户开通后再豁免期内退出无需扣款
	SupportExemptionPeriod *bool `json:"support_exemption_period,omitzero"`
	// WithholdMode represents withhold_mode. 设置扣款模式，目前支持每月固定日期或固定时间间隔
	WithholdMode *string `json:"withhold_mode,omitzero"`
}
