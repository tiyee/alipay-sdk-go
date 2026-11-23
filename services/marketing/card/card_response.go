package card

import (
	"encoding/json"
)

// CardActivateFormQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CardActivateFormQueryResponse struct {
	// Infos 是表单提交信息各个字段的值JSON数组				通用表单字段名称如下示例：		OPEN_FORM_FIELD_MOBILE – 手机号		OPEN_FORM_FIELD_GENDER – 性别		OPEN_FORM_FIELD_NAME – 姓名		OPEN_FORM_FIELD_BIRTHDAY – 生日		OPEN_FORM_FIELD_IDCARD – 身份证		OPEN_FORM_FIELD_EMAIL – 邮箱		OPEN_FORM_FIELD_ADDRESS – 地址				详细字段名称列表见会员卡开卡表单模板配置接口：alipay.marketing.card.formtemplate.set				注：		1. 证件类型字段(OPEN_FORM_FIELD_CERT_TYPE)返回结果取值如下:		    0 -- 身份证		    1 -- 护照		    2 -- 港澳居民通行证		    3 -- 台湾居民通行证
	Infos *string `json:"infos,omitempty"`
}

// CardActivateURLApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type CardActivateURLApplyResponse struct {
	// ApplyCardUrl 是会员卡领卡链接(使用时需URLDecode解码)。商户获取此链接后可投放到服务窗消息、店铺二维码等。
	ApplyCardUrl *string `json:"apply_card_url,omitempty"`
}

// CardConsumeSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type CardConsumeSyncResponse struct {
	// ExternalCardNo 是外部卡号
	ExternalCardNo *string `json:"external_card_no,omitempty"`
}

// CardDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CardDeleteResponse struct {
	// BizSerialNo 是支付宝端删卡业务流水号
	BizSerialNo *string `json:"biz_serial_no,omitempty"`
}

// CardFormTemplateSetResponse 是支付宝官方 V3 OpenAPI 模型。
type CardFormTemplateSetResponse = json.RawMessage

// CardMessageNotifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CardMessageNotifyResponse struct {
	// ResultCode 是二级错误处理结果（如果公用返回结果为false，则可以看这个接口判断明细原因） 如果公用返回为true，则该字段为空
	ResultCode *string `json:"result_code,omitempty"`
}

// CardOpenResponse 是支付宝官方 V3 OpenAPI 模型。
type CardOpenResponse struct {
	CardInfo *MerchantCard `json:"card_info,omitempty"`
	// OpenCardChannel 是实际记录的领卡渠道（可能跟商户传入值不同）；		可直接标识领卡渠道，也可配合open_card_channel_id标识领卡渠道类型：		例如：		线下门店领取:20161534000000000008863（直接标识领卡渠道，门店shopId）		线下扫二维码领取:QR（标识领卡类型）;		线下活动领取:20170522000000000003609（直接标识领卡渠道，商户活动ID）
	OpenCardChannel *string `json:"open_card_channel,omitempty"`
	// OpenCardChannelId 是实际记录的领卡来源渠道id（可能跟商户传入值不同）；		区别于open_card_channel领卡渠道；		一般使用场景：		open_card_channel用于区分渠道类型，例如取值为"SHOP"（门店），"ACTIVITY"（活动）；		则open_card_channel_id可用于区分同渠道的不同实体，对应取各门店ID或各活动的标识ID等；
	OpenCardChannelId *string `json:"open_card_channel_id,omitempty"`
}

// CardQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CardQueryResponse struct {
	CardInfo          *MerchantCard              `json:"card_info,omitempty"`
	PaidOuterCardInfo *PaidOuterCardExtraInfoDTO `json:"paid_outer_card_info,omitempty"`
	// PassId 是商家会员卡页面跳转到支付宝卡券详情页的pass_id，对应schema地址中的参数p，	主要用于小程序跳往会员卡详情的链接拼接
	PassId *string `json:"pass_id,omitempty"`
	// SchemaUrl 是商家会员卡页面跳转到支付宝卡券详情页面的schema地址
	SchemaUrl *string `json:"schema_url,omitempty"`
}

// CardUpdateResponse 是支付宝官方 V3 OpenAPI 模型。
type CardUpdateResponse struct {
	// ResultCode 是二级错误处理结果（如果公用返回结果为false，则可以看这个接口判断明细原因） 	如果公用返回为true，则该字段为空
	ResultCode *string `json:"result_code,omitempty"`
}
