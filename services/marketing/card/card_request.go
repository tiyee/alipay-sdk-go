package card

// CardActivateURLApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type CardActivateURLApplyRequest struct {
	// Callback 是会员卡开卡表单提交后回调地址。要求必须是单纯的服务端接收回调。	说明：	1.该地址不可带参数，如需回传参数，可设置out_string入参。
	Callback *string `json:"callback,omitempty"`
	// FollowAppId 是需要关注的生活号AppId。若需要在领卡页面展示“关注生活号”提示，需开通生活号并绑定会员卡。生活号快速接入详见：https://doc.open.alipay.com/docs/doc.htm?treeId=193&articleId=105933&docType=1
	FollowAppId *string `json:"follow_app_id,omitempty"`
	// OutString 是该值为商家拉起开卡组件的传递的 out_string  值。通常可用于区分不同业务场景，禁止将该字段作为开卡流程的必要字段，否则会导致会员推广场景下的入会失败
	OutString *string `json:"out_string,omitempty"`
	// TemplateId 是会员卡模板id。使用会员卡模板创建接口(alipay.marketing.card.template.create)返回的结果
	TemplateId *string `json:"template_id,omitempty"`
}

// CardConsumeSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type CardConsumeSyncRequest struct {
	// ActPayAmount 是用户实际付的现金金额		1.针对预付卡面额的核销金额在use_benefit_list展现，作为权益金额		2.权益金额不叠加在该金额上
	ActPayAmount *string       `json:"act_pay_amount,omitempty"`
	CardInfo     *MerchantCard `json:"card_info,omitempty"`
	// GainBenefitList 是获取权益列表
	GainBenefitList *[]BenefitInfoDetail `json:"gain_benefit_list,omitempty"`
	// Memo 是备注信息，现有直接填写门店信息
	Memo *string `json:"memo,omitempty"`
	// ShopCode 是门店编号
	ShopCode *string `json:"shop_code,omitempty"`
	// SwipeCertType 是产生该笔交易时，用户出具的凭证类型。枚举支持：	*ALIPAY：支付宝电子卡；	*ENTITY：实体卡；	*OTHER：其他。
	SwipeCertType *string `json:"swipe_cert_type,omitempty"`
	// TargetCardNo 是支付宝业务卡号，即通过alipay.marketing.card.open(会员卡开卡)接口开卡后获取的  card_info.biz_card_no 值。
	TargetCardNo *string `json:"target_card_no,omitempty"`
	// TargetCardNoType 是卡号ID类型。支持：	*BIZ_CARD：支付宝业务卡号（商户会员卡场景使用）。
	TargetCardNoType *string `json:"target_card_no_type,omitempty"`
	// TradeAmount 是交易金额：本次交易的实际总金额（可认为标价金额）
	TradeAmount *string `json:"trade_amount,omitempty"`
	// TradeName 是交易名称。为空时支付宝将根据交易类型提供默认名称。
	TradeName *string `json:"trade_name,omitempty"`
	// TradeNo 是支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
	// TradeTime 是线下交易时间，为用户付款的交易时间。	说明：当交易时间晚于上次消费记录同步时间时，将变更会员卡信息。
	TradeTime *string `json:"trade_time,omitempty"`
	// TradeType 是交易类型。枚举支持：	*开卡：OPEN；	*消费：TRADE；	*充值：DEPOSIT；	*退卡：RETURN。
	TradeType *string `json:"trade_type,omitempty"`
	// UseBenefitList 是实际消耗的权益
	UseBenefitList *[]BenefitInfoDetail `json:"use_benefit_list,omitempty"`
}

// CardFormTemplateSetRequest 是支付宝官方 V3 OpenAPI 模型。
type CardFormTemplateSetRequest struct {
	Fields *OpenFormFieldDO `json:"fields,omitempty"`
	// TemplateId 是会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。
	TemplateId *string `json:"template_id,omitempty"`
}

// CardMessageNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CardMessageNotifyRequest struct {
	// NotifyInfo 是消息通知详情，包含消息通知需要的信息，map 格式存储。	积分过期提醒：	expired_point 过期积分值	expired_time 过期时间
	NotifyInfo *string `json:"notify_info,omitempty"`
	// NotifyNo 是消息通知幂等标识，业务流水号
	NotifyNo *string `json:"notify_no,omitempty"`
	// NotifyType 是消息通知类型,	POINT_REMIND 积分过期提醒
	NotifyType *string `json:"notify_type,omitempty"`
	// OccurTime 是标识业务发生时间，时间格式为"yyyy-MM-dd HH:mm:ss"。
	OccurTime *string `json:"occur_time,omitempty"`
	// TargetCardNo 是支付宝业务卡号，即通过 alipay.marketing.card.open 接口开卡后获取的 card_info.biz_card_no 值。
	TargetCardNo *string `json:"target_card_no,omitempty"`
	// TargetCardNoType 是卡号ID类型
	TargetCardNoType *string `json:"target_card_no_type,omitempty"`
}

// CardOpenRequest 是支付宝官方 V3 OpenAPI 模型。
type CardOpenRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken   *string       `query:"auth_token" json:"-"`
	CardExtInfo *MerchantCard `json:"card_ext_info,omitempty"`
	// CardTemplateId 是支付宝分配的卡模板Id（卡模板创建接口返回的模板ID）
	CardTemplateId *string         `json:"card_template_id,omitempty"`
	CardUserInfo   *CardUserInfo   `json:"card_user_info,omitempty"`
	MemberExtInfo  *MerchantMenber `json:"member_ext_info,omitempty"`
	// OpenCardChannel 是领卡渠道，用于记录外部商户端领卡来源的渠道信息，渠道值可自行定义（仅限数字、字母、下划线） 		可直接标识领卡渠道，也可配合open_card_channel_id标识领卡渠道类型：		例如：		线下门店领取:20161534000000000008863（直接标识领卡渠道，门店shopId）		线下扫二维码领取:QR（标识领卡类型）;		线下活动领取:20170522000000000003609（直接标识领卡渠道，商户活动ID）
	OpenCardChannel *string `json:"open_card_channel,omitempty"`
	// OpenCardChannelId 是领卡来源的渠道id，注意区别于open_card_channel领卡渠道；		一般使用场景：		open_card_channel用于区分渠道类型，例如取值为"SHOP"（门店），"ACTIVITY"（活动）；		则open_card_channel_id可用于区分同渠道的不同实体，对应取各门店ID或各活动的标识ID等；
	OpenCardChannelId *string `json:"open_card_channel_id,omitempty"`
	// OutSerialNo 是外部商户流水号。由商户自定义，需保证商户系统中唯一。
	OutSerialNo       *string                    `json:"out_serial_no,omitempty"`
	PaidOuterCardInfo *PaidOuterCardExtraInfoDTO `json:"paid_outer_card_info,omitempty"`
}

// CardQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CardQueryRequest struct {
	CardUserInfo *CardUserInfo `json:"card_user_info,omitempty"`
	// ExtInfo 是扩展信息，暂时没有
	ExtInfo *string `json:"ext_info,omitempty"`
	// TargetCardNo 是操作卡号。	若target_card_no_type为  BIZ_CARD，则该值为业务卡号，即alipay.marketing.card.open(会员卡开卡)返回结果中的card_info#biz_card_no。	若target_card_no_type为 D_QR_CODE 或D_BAR_CODE，则该值为动态码值。	若target_card_no_type为 EXTERNAL_CARD，则该值为外部卡号。
	TargetCardNo *string `json:"target_card_no,omitempty"`
	// TargetCardNoType 是卡号ID类型
	TargetCardNoType *string `json:"target_card_no_type,omitempty"`
	// TemplateId 是会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。
	TemplateId *string `json:"template_id,omitempty"`
}

// CardUpdateRequest 是支付宝官方 V3 OpenAPI 模型。
type CardUpdateRequest struct {
	CardInfo *MerchantCard `json:"card_info,omitempty"`
	// ExtInfo 是扩展信息(暂时无用)
	ExtInfo             *string              `json:"ext_info,omitempty"`
	McardStyleInfo      *McardStylInfo       `json:"mcard_style_info,omitempty"`
	MerchantCardMsgInfo *MerchantCardMsgInfo `json:"merchant_card_msg_info,omitempty"`
	// NotifyMessages 是注意：此字段已废弃。	卡信息变更通知消息	1、在列表中定义的消息，才会发送给用户，消息格式一定。	2、根据卡信息变更情况，一次可发送多个消息
	NotifyMessages *[]McardNotifyMessage `json:"notify_messages,omitempty"`
	// OccurTime 是业务发生时间，时间格式为"yyyy-MM-dd HH:mm:ss"。
	OccurTime         *string                    `json:"occur_time,omitempty"`
	PaidOuterCardInfo *PaidOuterCardExtraInfoDTO `json:"paid_outer_card_info,omitempty"`
	// TargetCardNo 是支付宝业务卡号：1、3.0（历史版本）版本来源于alipay.marketing.card.open(会员卡开卡)接口开卡后获取的 card_info.biz_card_no 值；2、极速版来源于spi.alipay.user.opencard.get(会员卡开通，获取会员卡信息)请求参数中的biz_card_no值。
	TargetCardNo *string `json:"target_card_no,omitempty"`
	// TargetCardNoType 是卡号ID类型
	TargetCardNoType *string `json:"target_card_no_type,omitempty"`
}

// CardActivateFormQueryRequest is the query request model for alipay.marketing.card.activateform.query.
type CardActivateFormQueryRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
	// BizType 是开放表单信息查询业务类型
	BizType *string `json:"biz_type,omitempty" query:"biz_type,omitempty"`
	// TemplateId 是会员卡模板id。使用会员卡模板创建接口(alipay.marketing.card.template.create)返回的结果
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
	// RequestId 是查询用户表单提交信息的请求id。		网页移动应用&小程序应用：与auth_code  一起通过alipay.marketing.card.activateurl.apply(获取会员卡领卡投放链接)接口callback地址回传。		注意：auth_code（用户授权码）用于换取用户授权令牌 accessToken，后续调用接口需传入accessToken。
	RequestId *string `json:"request_id,omitempty" query:"request_id,omitempty"`
}

// CardDeleteRequest is the query request model for alipay.marketing.card.delete.
type CardDeleteRequest struct {
	// OutSerialNo 是商户端删卡业务流水号，商户自定义且需确保流水号唯一性。
	OutSerialNo *string `json:"out_serial_no,omitempty" query:"out_serial_no,omitempty"`
	// TargetCardNo 是支付宝业务卡号。即开卡接口返回结果中的card_info.biz_card_no。
	TargetCardNo *string `json:"target_card_no,omitempty" query:"target_card_no,omitempty"`
	// TargetCardNoType 是卡号ID类型
	TargetCardNoType *string `json:"target_card_no_type,omitempty" query:"target_card_no_type,omitempty"`
	// ReasonCode 是删卡原因		USER_UNBUND：用户解绑（可以重新绑定）		CANCEL：销户（完成销户后，就不能再重新绑定）		PRESENT：转赠（可以重新绑定）
	ReasonCode *string `json:"reason_code,omitempty" query:"reason_code,omitempty"`
	// ExtInfo 是删卡扩展参数，json格式。		用于商户的特定业务信息的传递，只有商户与支付宝约定了传递此参数且约定了参数含义，此参数才有效。		目前支持如下key：		new_card_no：新卡号		donee_user_id：受赠人userId
	ExtInfo *string `json:"ext_info,omitempty" query:"ext_info,omitempty"`
}
