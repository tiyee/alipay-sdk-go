package card

// BenefitInfoDetail 是支付宝官方 V3 OpenAPI 模型。
type BenefitInfoDetail struct {
	// Amount 是PRE_FUND：实际核销或者商户赠送的金额		DISCOUNT：实际折扣掉的金额（获取权益不支持该类型）		COUPON：实际核销或者商户赠送的券
	Amount *string `json:"amount,omitempty"`
	// BenefitType 是权益类型
	BenefitType *string `json:"benefit_type,omitempty"`
	// Count 是COUPON：当核销或者赠送券时，可以设置该值
	Count *string `json:"count,omitempty"`
	// Description 是产生核销或者赠送权益的描述
	Description *string `json:"description,omitempty"`
}

// CardFrontTextDTO 是支付宝官方 V3 OpenAPI 模型。
type CardFrontTextDTO struct {
	// Label 是文案标签
	Label *string `json:"label,omitempty"`
	// Value 是展示文案
	Value *string `json:"value,omitempty"`
}

// CardUserInfo 是支付宝官方 V3 OpenAPI 模型。
type CardUserInfo struct {
	// OpenId 是用户open_id
	OpenId *string `json:"open_id,omitempty"`
	// UserUniId 是用户唯一标识, 根据user_uni_id_type类型来定 （目前暂支持支付宝userId）		支付宝userId说明：支付宝用户号是以2088开头的16位纯数字组成
	UserUniId *string `json:"user_uni_id,omitempty"`
	// UserUniIdType 是用户唯一标识类型
	UserUniIdType *string `json:"user_uni_id_type,omitempty"`
}

// McardNotifyMessage 是支付宝官方 V3 OpenAPI 模型。
type McardNotifyMessage struct {
	// ChangeReason 是用户提醒信息，按如下格式拼装，需要ISV提供change_reason。		积分变动模板：{change_reason}，您的积分有变动		余额变动模板：{change_reason}，您的余额有变动		等级变更无需提供原因。
	ChangeReason *string `json:"change_reason,omitempty"`
	// ExtInfo 是JSON格式扩展信息，主要是发送消息中的变量
	ExtInfo *string `json:"ext_info,omitempty"`
	// MessageType 是消息类型，每种消息都定义了固定消息模板
	MessageType *string `json:"message_type,omitempty"`
}

// McardStylInfo 是支付宝官方 V3 OpenAPI 模型。
type McardStylInfo struct {
	// BackgroundId 是背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片				图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif；		尺寸不小于1020*643px；		图片不得有圆角，不得拉伸变形
	BackgroundId *string `json:"background_id,omitempty"`
	// BgColor 是背景色
	BgColor *string `json:"bg_color,omitempty"`
	// LogoId 是logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片				图片说明：1M以内，格式bmp、png、jpeg、jpg、gif；		尺寸不小于500*500px的正方形；		请优先使用商家LOGO；
	LogoId *string `json:"logo_id,omitempty"`
}

// MdCodeInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type MdCodeInfoDTO struct {
	// CodeStatus 是本次回传动态码的状态：		SUCCESS: 本次发码成功		FAIL_RETRY: 本次发码失败，且需要支付宝重试（重新通知商户发码）		FAIL_NOT_RETRY: 本次发码失败，且无需支付宝重试（无需重新通知商户发码）
	CodeStatus *string `json:"code_status,omitempty"`
	// CodeValue 是动态码的码值：		code_status为SUCCESS时必填；		基于此码值生成条形码或二维码用于扫码核销。
	CodeValue *string `json:"code_value,omitempty"`
	// ExpireTime 是当前动态码的过期（失效）时间：		 code_status为SUCCESS时必填。
	ExpireTime *string `json:"expire_time,omitempty"`
	// TimeStamp 是商户回传动态码的时间戳 (单位秒)。				即商户调接口回传动态码时刻对应的long类型时间戳，用于区分不同的发码请求。
	TimeStamp *int64 `json:"time_stamp,omitempty"`
}

// MerchantCard 是支付宝官方 V3 OpenAPI 模型。
type MerchantCard struct {
	// Balance 是资金卡余额，单位：元，精确到小数点后两位。
	Balance *string `json:"balance,omitempty"`
	// BizCardNo 是支付宝业务卡号		说明：		1、开卡成功后返回该参数，需要保存留用；		2、开卡/更新/删卡/查询卡接口请求中不需要传该参数；
	BizCardNo *string `json:"biz_card_no,omitempty"`
	// CustomAssets 是会员卡自定义资产值，只供特定业务使用，通常接入无需关注
	CustomAssets *string `json:"custom_assets,omitempty"`
	// ExternalCardNo 是用户在商户crm系统中的会员卡卡号，该参数必填。
	ExternalCardNo *string `json:"external_card_no,omitempty"`
	// FrontImageId 是卡面展示图片的图片ID，通过接口（alipay.offline.material.image.upload）上传图片				这里预期展示的是个人照片；		图片说明：1M以内，格式bmp、png、jpeg、jpg、gif；		图片尺寸为230*295px，可等比放大；
	FrontImageId *string `json:"front_image_id,omitempty"`
	// FrontTextList 是卡面文案列表，1项对应1行文案，最多只能传入4行文案； 	单行文案展现分为左右两部分，左边对应label字段，右边对应value； 	形如： 学院    新闻学院
	FrontTextList *[]CardFrontTextDTO `json:"front_text_list,omitempty"`
	// Level 是会员卡等级（由商户自定义，并可以在卡模板创建时，定义等级信息）
	Level      *string        `json:"level,omitempty"`
	MdcodeInfo *MdCodeInfoDTO `json:"mdcode_info,omitempty"`
	// OpenDate 是会员卡开卡时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。
	OpenDate *string `json:"open_date,omitempty"`
	// Point 是会员卡积分，积分必须为数字型（可为浮点型，带2位小数点）
	Point *string `json:"point,omitempty"`
	// TemplateId 是会员卡更换不同的卡模板（该参数仅用在会员卡更新接口中）
	TemplateId *string `json:"template_id,omitempty"`
	// ValidDate 是会员卡有效期结束时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。
	ValidDate *string `json:"valid_date,omitempty"`
}

// MerchantCardMsgInfo 是支付宝官方 V3 OpenAPI 模型。
type MerchantCardMsgInfo struct {
	// ChangedPoint 是本次消息通知的变动积分，积分必须为数字型（可为浮点型，带2位小数点），格式非法则不发送消息。例如：	100，则代表增加100积分。	-100，则代表减少100积分。
	ChangedPoint *string `json:"changed_point,omitempty"`
}

// MerchantMenber 是支付宝官方 V3 OpenAPI 模型。
type MerchantMenber struct {
	// Birth 是生日 yyyy-MM-dd
	Birth *string `json:"birth,omitempty"`
	// Cell 是手机号
	Cell *string `json:"cell,omitempty"`
	// Gende 是性别
	Gende *string `json:"gende,omitempty"`
	// Name 是姓名
	Name *string `json:"name,omitempty"`
}

// OpenFormFieldDO 是支付宝官方 V3 OpenAPI 模型。
type OpenFormFieldDO struct {
	// Optional 是表单可选字段配置，common_fields属性定义一个表单字段数组，表单字段有效值列表与required字段有效值列表相同。		可选字段配置中不能含有必须字段配置的有效值。
	Optional *string `json:"optional,omitempty"`
	// Required 是表单必填字段配置，common_fields属性定义一个表单字段数组，字段有效值如下列表所示：		OPEN_FORM_FIELD_MOBILE -- 手机号		OPEN_FORM_FIELD_GENDER -- 性别		OPEN_FORM_FIELD_NAME -- 姓名		OPEN_FORM_FIELD_BIRTHDAY -- 生日		OPEN_FORM_FIELD_IDCARD -- 身份证		OPEN_FORM_FIELD_EMAIL -- 邮箱		OPEN_FORM_FIELD_ADDRESS -- 地址		OPEN_FORM_FIELD_CITY -- 城市		OPEN_FORM_FIELD_IS_STUDENT -- 是否学生认证		OPEN_FORM_FIELD_MEMBER_GRADE -- 会员等级
	Required *string `json:"required,omitempty"`
}

// PaidOuterCardCycleInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardCycleInfoDTO struct {
	// AlipayDeductAgreement 是支付宝用户签约协议号
	AlipayDeductAgreement *string `json:"alipay_deduct_agreement,omitempty"`
	// AlipayDeductProductCode 是支付宝代扣产品码
	AlipayDeductProductCode *string `json:"alipay_deduct_product_code,omitempty"`
	// AlipayDeductScene 是支付宝代扣场景码
	AlipayDeductScene *string `json:"alipay_deduct_scene,omitempty"`
	// CloseReason 是关闭连续购买原因。
	CloseReason *string `json:"close_reason,omitempty"`
	// CycleType 是用户开通连续购买类型。
	CycleType *string `json:"cycle_type,omitempty"`
	// OpenStatus 是开通连续购买状态。
	OpenStatus *string `json:"open_status,omitempty"`
}

// PaidOuterCardExtraInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardExtraInfoDTO struct {
	// Action 是用户操作类型
	Action       *string                       `json:"action,omitempty"`
	CycleInfo    *PaidOuterCardCycleInfoDTO    `json:"cycle_info,omitempty"`
	PurchaseInfo *PaidOuterCardPurchaseInfoDTO `json:"purchase_info,omitempty"`
}

// PaidOuterCardPurchaseInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardPurchaseInfoDTO struct {
	// ActionDate 是用户购买、升级、降级、续费的操作时间。格式为：yyyy-MM-dd hh:mm:ss
	ActionDate *string `json:"action_date,omitempty"`
	// AlipayTradeNo 是支付宝交易号，由api: alipay.trade.create 返回
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty"`
	// OutTradeNo 是商户订单号。与创建订单api：alipay.trade.create保持一致
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// Price 是用户购买金额。当购买/升级/续费场景必填，单位元，精确到小数点后2位
	Price *string `json:"price,omitempty"`
	// Source 是用户操作来源
	Source *string `json:"source,omitempty"`
}
