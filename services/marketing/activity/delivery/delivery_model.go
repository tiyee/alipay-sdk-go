// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package delivery

// MarketingActivityDeliveryQuery 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityDeliveryQuery struct {
	BelongMerchantInfo *DeliveryAgencyMerchantInfo `json:"belong_merchant_info,omitempty"`
	// DeliveryConfigList 是"[已废弃] 待查询的投放配置列表。 最大数量限制20个。"
	DeliveryConfigList *[]DeliveryConfig `json:"delivery_config_list,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
}

// MarketingActivityDeliveryStop 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityDeliveryStop struct {
	BelongMerchantInfo *DeliveryAgencyMerchantInfo `json:"belong_merchant_info,omitempty"`
	// DeliveryConfigList 是"[已废弃] 待停止的投放配置列表。 最大数量限制20个。"
	DeliveryConfigList *[]DeliveryConfig `json:"delivery_config_list,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。 幂等作用： 再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// DeliveryActivityContentInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryActivityContentInfo struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// DeliveryAgencyMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryAgencyMerchantInfo struct {
	// BusinessType 是合作业务类型。
	BusinessType *string `json:"business_type,omitempty"`
	// MerchantId 是商户信息
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。 限制: 服务商代运营模式必传	如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。 因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// DeliveryAvailableCityCode 是支付宝官方 V3 OpenAPI 模型。
type DeliveryAvailableCityCode struct {
	// AllCity 是是否全国。 与city_codes二选一。只允许填true，否则不填。
	AllCity *bool `json:"all_city,omitempty"`
	// CityCodes 是城市编码。与all_city二选一。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx  表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/ ）
	CityCodes *[]string `json:"city_codes,omitempty"`
}

// DeliveryAvailableScope 是支付宝官方 V3 OpenAPI 模型。
type DeliveryAvailableScope struct {
	DeliveryAvailableCityCode *DeliveryAvailableCityCode `json:"delivery_available_city_code,omitempty"`
	// DeliveryAvailableScopeType 是投放可用范围类型，目前支持的类型是城市范围CITY_CODE
	DeliveryAvailableScopeType *string `json:"delivery_available_scope_type,omitempty"`
}

// DeliveryBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryBaseInfo struct {
	// DeliveryBeginTime 是投放计划开始时间。 格式为：yyyy-MM-dd HH:mm:ss。
	DeliveryBeginTime *string `json:"delivery_begin_time,omitempty"`
	// DeliveryEndTime 是投放计划结束时间。 格式为：yyyy-MM-dd HH:mm:ss。
	DeliveryEndTime  *string           `json:"delivery_end_time,omitempty"`
	DeliveryMaterial *DeliveryMaterial `json:"delivery_material,omitempty"`
	// DeliveryName 是投放计划名称。	投放计划名称不会对用户进行表达，只用于商户管理使用。长度需要大于等于3，小于20。
	DeliveryName *string `json:"delivery_name,omitempty"`
}

// DeliveryCityCodeRule 是支付宝官方 V3 OpenAPI 模型。
type DeliveryCityCodeRule struct {
	// AllCity 是是否全国。只允许填true，否则不填。
	AllCity *bool `json:"all_city,omitempty"`
	// CityCodes 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx  表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/ ）
	CityCodes *[]string `json:"city_codes,omitempty"`
}

// DeliveryConfig 是支付宝官方 V3 OpenAPI 模型。
type DeliveryConfig struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// BoothCode 是展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAYMENT_RESULT
	BoothCode *string `json:"booth_code,omitempty"`
	// Channel 是可投放的渠道标识.
	Channel *string `json:"channel,omitempty"`
}

// DeliveryContentConfig 是支付宝官方 V3 OpenAPI 模型。
type DeliveryContentConfig struct {
	DeliverySendGuide *DeliverySendGuide `json:"delivery_send_guide,omitempty"`
}

// DeliveryContentInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryContentInfo struct {
	DeliveryActivityContent *DeliveryActivityContentInfo `json:"delivery_activity_content,omitempty"`
	DeliveryAppContent      *DeliveryMiniAppContentInfo  `json:"delivery_app_content,omitempty"`
	// DeliveryContentType 是运营内容类型。
	DeliveryContentType *string                  `json:"delivery_content_type,omitempty"`
	DeliveryDisplayInfo *DeliveryDisplayInfo     `json:"delivery_display_info,omitempty"`
	DeliveryItemContent *DeliveryItemContentInfo `json:"delivery_item_content,omitempty"`
}

// DeliveryCrowdRule 是支付宝官方 V3 OpenAPI 模型。
type DeliveryCrowdRule struct {
	// CrowdId 是指定人群ID
	CrowdId *string `json:"crowd_id,omitempty"`
}

// DeliveryDisplayInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryDisplayInfo struct {
	// MainTitle 是副标题。
	MainTitle *string `json:"main_title,omitempty"`
	// SubTitle 是副标题
	SubTitle *string `json:"sub_title,omitempty"`
}

// DeliveryFullSendConfig 是支付宝官方 V3 OpenAPI 模型。
type DeliveryFullSendConfig struct {
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info,omitempty"`
	// DeliveryFloorAmount 是满足消费金额门槛，左闭区间，单位元。	说明：限制支付时的订单金额最少满足的金额门槛。	限制：value > 0 && value <= 99999。
	DeliveryFloorAmount *string `json:"delivery_floor_amount,omitempty"`
}

// DeliveryItemContentInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryItemContentInfo struct {
	// ItemId 是支付宝侧商品id。	注意事项：支持状态为“出售中”的商品
	ItemId *string `json:"item_id,omitempty"`
	// ItemType 是商品模版类型
	ItemType *string `json:"item_type,omitempty"`
}

// DeliveryMaterial 是支付宝官方 V3 OpenAPI 模型。
type DeliveryMaterial struct {
	DeliverySingleMaterial *DeliverySingleMaterial `json:"delivery_single_material,omitempty"`
	// MaterialIdList 是素材ID，通过alipay.marketing.material.create接口提供的素材组
	MaterialIdList *[]string `json:"material_id_list,omitempty"`
}

// DeliveryMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryMerchantInfo struct {
	// MerchantId 是商户id
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// DeliveryMerchantRule 是支付宝官方 V3 OpenAPI 模型。
type DeliveryMerchantRule struct {
	// BrandIdList 是指定品牌id。 说明：如商户需选择某个品牌下维护的收款账号，请上传相关品牌id
	BrandIdList *[]string `json:"brand_id_list,omitempty"`
	// DeliveryMerchantInfos 是曝光商户选取列表。 说明：需要传入您期望曝光的商户的商户号，传入为空时默认使用投放优惠券活动的适用范围。 限制：曝光商户号需与投放归属商户号相同，或传入有跨主体授权关系的商户号或有弱绑定关系的M3账号
	DeliveryMerchantInfos *[]DeliveryMerchantInfo `json:"delivery_merchant_infos,omitempty"`
	// DeliveryMerchantMode 是指定支付成功页模式。
	DeliveryMerchantMode *string `json:"delivery_merchant_mode,omitempty"`
}

// DeliveryMiniAppContentInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryMiniAppContentInfo struct {
	// MiniAppDeliveryType 是仅小程序推广可用。
	MiniAppDeliveryType *string `json:"mini_app_delivery_type,omitempty"`
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// ServiceCodeList 是小程序服务编码/home_page，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 限制: 1.只有mini_app_id有值时该值传入才会有效 2.当前只支持一个服务code 3.service_codes必须是已传入的mini_app_id下的小程序
	ServiceCodeList *[]string `json:"service_code_list,omitempty"`
}

// DeliveryPlayConfig 是支付宝官方 V3 OpenAPI 模型。
type DeliveryPlayConfig struct {
	DeliveryFullSendConfig   *DeliveryFullSendConfig   `json:"delivery_full_send_config,omitempty"`
	DeliverySingleSendConfig *DeliverySingleSendConfig `json:"delivery_single_send_config,omitempty"`
}

// DeliverySearchBoxRule 是支付宝官方 V3 OpenAPI 模型。
type DeliverySearchBoxRule struct {
	// BoxId 是指定搜索直达区域
	BoxId *string `json:"box_id,omitempty"`
}

// DeliverySendGuide 是支付宝官方 V3 OpenAPI 模型。
type DeliverySendGuide struct {
	// DeliveryGuideUrl 是引导链接，从支付宝公域跳转到服务商(商户)承接页的引导链接。
	DeliveryGuideUrl *string `json:"delivery_guide_url,omitempty"`
}

// DeliverySingleMaterial 是支付宝官方 V3 OpenAPI 模型。
type DeliverySingleMaterial struct {
	// DeliveryImage 是投放计划图片素材。通过接口alipay.marketing.material.image.upload上传图片返回的resource_id。
	DeliveryImage *string `json:"delivery_image,omitempty"`
}

// DeliverySingleSendConfig 是支付宝官方 V3 OpenAPI 模型。
type DeliverySingleSendConfig struct {
	DeliveryContentConfig *DeliveryContentConfig `json:"delivery_content_config,omitempty"`
	DeliveryContentInfo   *DeliveryContentInfo   `json:"delivery_content_info,omitempty"`
}

// DeliveryTargetRule 是支付宝官方 V3 OpenAPI 模型。
type DeliveryTargetRule struct {
	// BrandIdList 是指定品牌id。	说明：如商户需选择某个品牌下维护的收款账号，请上传相关品牌id
	BrandIdList            *[]string               `json:"brand_id_list,omitempty"`
	DeliveryAvailableScope *DeliveryAvailableScope `json:"delivery_available_scope,omitempty"`
	DeliveryCityCodeRule   *DeliveryCityCodeRule   `json:"delivery_city_code_rule,omitempty"`
	DeliveryCrowdRule      *DeliveryCrowdRule      `json:"delivery_crowd_rule,omitempty"`
	// DeliveryMerchantInfos 是"曝光商户选取列表。 说明：需要传入您期望曝光的商户的商户号，传入为空时默认使用投放优惠券活动的适用范围。 限制：曝光商户号需与投放归属商户号相同，或传入有跨主体授权关系的商户号或有弱绑定关系的M3账号"
	DeliveryMerchantInfos *[]DeliveryMerchantInfo `json:"delivery_merchant_infos,omitempty"`
	// DeliveryMerchantMode 是指定支付成功页模式。
	DeliveryMerchantMode *string               `json:"delivery_merchant_mode,omitempty"`
	DeliveryMerchantRule *DeliveryMerchantRule `json:"delivery_merchant_rule,omitempty"`
	// DeliveryPromoTags 是推荐定向曝光订单优惠标记。	1. 创建交易订单时，上传符合的订单优惠标。在业务参数信息“business_params”的子字段“delivery_promo_tags”填充订单优惠标。“统一收单交易支付接口”详情查看： https://opendocs.alipay.com/open/02ekfp?scene=32 	2. 订单优惠标，检验规则：	* 单个优惠标最大长度32字节。多个标记值以英文逗号分隔。	* 单个优惠标仅支持字母/数字/下划线。	3. 订单优惠标规则和其他推荐规则同时叠加。
	DeliveryPromoTags *string `json:"delivery_promo_tags,omitempty"`
	// DeliveryRecallMode 是投放召回模式，现在支持CITY_RECALL（城市召回)和SHOP_RECALL（门店召回)两种模式。
	DeliveryRecallMode    *string                `json:"delivery_recall_mode,omitempty"`
	DeliverySearchBoxRule *DeliverySearchBoxRule `json:"delivery_search_box_rule,omitempty"`
	DeliveryTemplateRule  *DeliveryTemplateRule  `json:"delivery_template_rule,omitempty"`
	// DeliveryType 是小程序推广可用。
	DeliveryType *string `json:"delivery_type,omitempty"`
}

// DeliveryTemplateRule 是支付宝官方 V3 OpenAPI 模型。
type DeliveryTemplateRule struct {
	// TemplateId 是指定商家消息区域
	TemplateId *string `json:"template_id,omitempty"`
}

// ErrorDeliveryConfig 是支付宝官方 V3 OpenAPI 模型。
type ErrorDeliveryConfig struct {
	DeliveryConfig *DeliveryConfig `json:"delivery_config,omitempty"`
	// ErrorCode 是错误码。针对单个投放，具体的错误原因。
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMsg 是详细的投放错误信息。用于判断具体的错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

// PromoDeliveryInfo 是支付宝官方 V3 OpenAPI 模型。
type PromoDeliveryInfo struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// BoothCode 是展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAYMENT_RESULT
	BoothCode *string `json:"booth_code,omitempty"`
	// Channel 是可投放的渠道标识.
	Channel *string `json:"channel,omitempty"`
	// ChannelName 是渠道名称
	ChannelName *string `json:"channel_name,omitempty"`
	// DeliveryStatus 是投放状态。
	DeliveryStatus *string `json:"delivery_status,omitempty"`
}

// SuccessDeliveryConfig 是支付宝官方 V3 OpenAPI 模型。
type SuccessDeliveryConfig struct {
	DeliveryConfig *DeliveryConfig `json:"delivery_config,omitempty"`
}
