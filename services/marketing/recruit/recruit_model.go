// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package recruit

// RecruitEnrollInfo 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollInfo struct {
	// Cities 是报名提交的活动城市码，详见活动城市码
	Cities         *[]string              `json:"cities,omitempty"`
	EnrollMerchant *RecruitEnrollMerchant `json:"enroll_merchant,omitempty"`
	// Materials 是报名提交的素材
	Materials *[]RecruitMaterial `json:"materials,omitempty"`
	// MiniApps 是报名提交的小程序信息，是否必选取决于方案要求
	MiniApps *[]RecruitMiniApp `json:"mini_apps,omitempty"`
	// Vouchers 是报名提交的券信息，是否必选取决于方案要求
	Vouchers *[]RecruitVoucher `json:"vouchers,omitempty"`
}

// RecruitEnrollMerchant 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollMerchant struct {
	// MerchantUid 是商户唯一ID(商户pid)
	MerchantUid *string `json:"merchant_uid,omitempty"`
}

// RecruitEnrollRule 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollRule struct {
	// MaxSize 是最大数量
	MaxSize *int64 `json:"max_size,omitempty"`
	// MinSize 是最小数量
	MinSize *int64 `json:"min_size,omitempty"`
	// Required 是是否必填
	Required *bool                  `json:"required,omitempty"`
	RuleData *RecruitEnrollRuleData `json:"rule_data,omitempty"`
	// Type 是招商报名时提交的内容类型:	券:VOUCHER;	小程序:MINI_APP;	报名商户:ENROLL_MERCHANT;	素材:MATERIAL;	活动城市:CITY;
	Type *string `json:"type,omitempty"`
}

// RecruitEnrollRuleData 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollRuleData struct {
	// RecruitVoucherRules 是招商方案可提报的券的规则列表，不同的券有不同的规则。
	RecruitVoucherRules *[]RecruitVoucherRule `json:"recruit_voucher_rules,omitempty"`
	// Schema 是素材的要求，json字符串，使用时需要把此字符串解析成json对象。	field：提报的字段key	label：字段说明	required：是否必填	type：字段类型。包含：图片（IMAGE）、文本（TEXT）、数据源（DATASOURCE）等	rules：字段约束规则	  图片（IMAGE）：	    image_size：图片宽高，单位是PX	    file_type：图片类型	    file_size：表示图片大小的最大值，单位是KB	    image_aspect_ratio: 图片宽高比例	  文本（TEXT）：	    unicode_length：文本长度，单位为Byte	  数据源（DATASOURCE）：	    unicode_length：文本长度，单位为Byte
	Schema *string `json:"schema,omitempty"`
}

// RecruitMaterial 是支付宝官方 V3 OpenAPI 模型。
type RecruitMaterial struct {
	// Data 是提报的素材内容，json字符串，使用时需要把此字符串解析成json对象。	字段来源于“招商方案查询接口（https://opendocs.alipay.com/pre-open/02bhl7）”中返回的素材字段field。如：activityImage、contentId等。	字段类型mediaType包含：	 IMAGE，图片类型	 TEXT，文本类型	 DATASOURCE，数据源类型， 如：内容ID（contentId）字段类型为DATASOURCE	字段取值：	 图片（IMAGE）：	 aftsFileId：“营销图片资源上传接口（https://opendocs.alipay.com/pre-open/02bhlj）”返回的图片唯一资源ID	 文本（TEXT）：	 text：文本内容	 数据源（DATASOURCE）：	 text：当前仅支持内容ID值，可通过“内容批量查询接口（https://opendocs.alipay.com/pre-apis/02lu12，仅供受邀用户使用）”查询内容ID。
	Data *string `json:"data,omitempty"`
	// Description 是对这个素材的描述，用于后续管理该素材
	Description *string `json:"description,omitempty"`
	// MaterialAuditInfo 是素材审核信息，展示给商家的素材审核不通过的意见。素材审核通过时该字段为空。	报名详情查询接口不用填写该字段。
	MaterialAuditInfo *string `json:"material_audit_info,omitempty"`
	// Name 是素材名名称
	Name *string `json:"name,omitempty"`
	// SubjectId 是关联的报名主体ID。	enroll_scene_type是VOUCHER时，报名主体ID是activity_id；enroll_scene_type是MINI_APP时，报名主体ID是mini_app_id
	SubjectId *string `json:"subject_id,omitempty"`
}

// RecruitMiniApp 是支付宝官方 V3 OpenAPI 模型。
type RecruitMiniApp struct {
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// RecruitPlanLight 是支付宝官方 V3 OpenAPI 模型。
type RecruitPlanLight struct {
	// Description 是招商方案描述：富文本内容，可能包含html标签
	Description *string `json:"description,omitempty"`
	// EnrollEndTime 是报名结束时间
	EnrollEndTime *string `json:"enroll_end_time,omitempty"`
	// EnrollSceneType 是招商方案所属报名场景	枚举值：	VOUCHER：券报名场景；	MINI_APP：小程序报名场景；
	EnrollSceneType *string `json:"enroll_scene_type,omitempty"`
	// EnrollStartTime 是报名开始时间
	EnrollStartTime *string `json:"enroll_start_time,omitempty"`
	// Logo 是方案头图链接，150x150 以上尺寸的方形图片
	Logo *string `json:"logo,omitempty"`
	// PlanId 是招商方案ID
	PlanId *string `json:"plan_id,omitempty"`
	// PlanName 是招商方案名称
	PlanName *string `json:"plan_name,omitempty"`
	// Status 是招商方案状态
	Status *string `json:"status,omitempty"`
}

// RecruitVoucher 是支付宝官方 V3 OpenAPI 模型。
type RecruitVoucher struct {
	// ActivityId 是券活动id。支持商家券（创建商家券请参考https://opendocs.alipay.com/apis/01xm17）和支付券（创建支付券请参考https://opendocs.alipay.com/pre-apis/027185，仅供受邀用户使用），取值均为返回结果的activity_id字段。
	ActivityId *string `json:"activity_id,omitempty"`
}

// RecruitVoucherRule 是支付宝官方 V3 OpenAPI 模型。
type RecruitVoucherRule struct {
	// AmountMax 是券面额（每张代金券可以抵扣的金额）的最大值。 币种为人民币，单位为元。小数点以后最多保留两位。 该字段为空时表示不限制。
	AmountMax *string `json:"amount_max,omitempty"`
	// AmountMin 是券面额（每张代金券可以抵扣的金额）的最小值。 币种为人民币，单位为元。小数点以后最多保留两位。 该字段为空时表示不限制。
	AmountMin *string `json:"amount_min,omitempty"`
	// DenominationPercentMax 是券优惠比例的最大值。20代表优惠比例最多是20%。券优惠券比例=券优惠面额/门槛金额。	浮点类型，取值范围为[1,100]，左右均是闭区间，小数点以后最多保留两位。	该字段为空时表示不限制。
	DenominationPercentMax *string `json:"denomination_percent_max,omitempty"`
	// DenominationPercentMin 是券优惠比例的最小值。10代表优惠比例最少是10%。券优惠券比例=券优惠面额/门槛金额。	浮点类型，取值范围为[1,100]，左右均是闭区间，小数点以后最多保留两位。	该字段为空时表示不限制。
	DenominationPercentMin *string `json:"denomination_percent_min,omitempty"`
	// FloorAmountMax 是券门槛金额的最大值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	FloorAmountMax *string `json:"floor_amount_max,omitempty"`
	// FloorAmountMin 是券门槛金额的最小值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	FloorAmountMin *string `json:"floor_amount_min,omitempty"`
	// OriginAmountMax 是券原价的最大值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	OriginAmountMax *string `json:"origin_amount_max,omitempty"`
	// OriginAmountMin 是券原价的最小值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	OriginAmountMin *string `json:"origin_amount_min,omitempty"`
	// PublishEndTimeMax 是券发放结束时间的最大值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	PublishEndTimeMax *string `json:"publish_end_time_max,omitempty"`
	// PublishEndTimeMin 是券发放结束时间的最小值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	PublishEndTimeMin *string `json:"publish_end_time_min,omitempty"`
	// PublishStartTimeMax 是券发放开始时间的最大值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	PublishStartTimeMax *string `json:"publish_start_time_max,omitempty"`
	// PublishStartTimeMin 是券发放开始时间的最小值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	PublishStartTimeMin *string `json:"publish_start_time_min,omitempty"`
	// RefundType 是退券类型要求，列表，总共有两种类型：	过期退OVERDUE_REFUND	随时退CAN_REFUND	该字段为空时表示不限制。
	RefundType *[]string `json:"refund_type,omitempty"`
	// SaleAmountMax 是用户购买优惠券需要支付的金额的最大值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	SaleAmountMax *string `json:"sale_amount_max,omitempty"`
	// SaleAmountMin 是用户购买优惠券需要支付的金额的最小值。	币种为人民币，单位为元。小数点以后最多保留两位。	该字段为空时表示不限制。
	SaleAmountMin *string `json:"sale_amount_min,omitempty"`
	// UseChannel 是券核销渠道要求，列表，总共有两个渠道：	门店SHOP	小程序MINI_APP	该字段为空时表示不限制。
	UseChannel *[]string `json:"use_channel,omitempty"`
	// ValidDaysAfterReceiveMin 是券生效后至少N天内可以使用。比如：valid_days_after_receive_min=20代表	券生效后至少20天内可以使用。	该字段为空时表示不限制。
	ValidDaysAfterReceiveMin *int64 `json:"valid_days_after_receive_min,omitempty"`
	// VoucherActivityType 是券活动类型。支持七种商家券类型和两种支付券类型。		枚举值：	商家券类型：	ALL_FIX_ORDER_VOUCHER 全场满减券；	ITEM_FIX_ORDER_VOUCHER 单品满减券；	ALL_DISCOUNT_ORDER_VOUCHER 全场折扣券；	ITEM_DISCOUNT_ORDER_VOUCHER 单品折扣券；	ITEM_SPE_ORDER_VOUCHER 单品特价券；	EXCHANGE_GROUP_BUY_ORDER_VOUCHER 兑换团购券；	EXCHANGE_FIX_ORDER_VOUCHER 兑换代金券；		支付券类型：	ALL_FIX_VOUCHER 全场满减券；	ITEM_FIX_VOUCHER 单品满减券；		创建商家券参考https://opendocs.alipay.com/apis/01xm17	创建支付券参考https://opendocs.alipay.com/pre-apis/027185（仅供受邀用户使用）	不同的创建券的参数创建出不同类型的券：	 1. 根据voucher_type区分满减券、折扣券、特价券、兑换券	 2. 根据goods_name是否为空区分单品券、全场券（当voucher_type为满减券、折扣券、特价券）	 3. 根据voucher_use_rule.exchange_voucher.biz_type区分团购券、代金券（当voucher_type为兑换券）
	VoucherActivityType *string `json:"voucher_activity_type,omitempty"`
	// VoucherQuantityLimitPerUserMax 是每人领取限制的最大值。	默认按照支付宝uid进行领取限制。	该字段为空时表示不限制。
	VoucherQuantityLimitPerUserMax *int64 `json:"voucher_quantity_limit_per_user_max,omitempty"`
	// VoucherQuantityLimitPerUserMin 是每人领取限制的最小值。	默认按照支付宝uid进行领取限制。	该字段为空时表示不限制。
	VoucherQuantityLimitPerUserMin *int64 `json:"voucher_quantity_limit_per_user_min,omitempty"`
	// VoucherQuantityMax 是券库存数量的最大值。	该字段为空时表示不限制。
	VoucherQuantityMax *int64 `json:"voucher_quantity_max,omitempty"`
	// VoucherQuantityMin 是券库存数量的最小值。	该字段为空时表示不限制。
	VoucherQuantityMin *int64 `json:"voucher_quantity_min,omitempty"`
	// VoucherValidBeginTimeMin 是券可使用的开始时间的最小值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	VoucherValidBeginTimeMin *string `json:"voucher_valid_begin_time_min,omitempty"`
	// VoucherValidEndTimeMax 是券可使用的结束时间的最大值。	格式为：yyyy-MM-dd HH:mm:ss	该字段为空时表示不限制。
	VoucherValidEndTimeMax *string `json:"voucher_valid_end_time_max,omitempty"`
}
