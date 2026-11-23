// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package delivery

// ActivityDeliveryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryCreateResponse struct {
	// DeliveryGuidePreviewUrl 是预览链接。	可用于预览用户在承接页面领券的流程。
	DeliveryGuidePreviewUrl *string `json:"delivery_guide_preview_url,omitempty"`
	// DeliveryId 是推广计划id
	DeliveryId *string `json:"delivery_id,omitempty"`
	// ErrorDeliveryConfigList 是[已废弃]	创建失败的投放信息列表		失败的投放配置列表中，会返回对应的deliveryConfig对象。
	ErrorDeliveryConfigList *[]ErrorDeliveryConfig `json:"error_delivery_config_list,omitempty"`
	// OutBizNo 是外部业务单号，用作该请求唯一标识。 	创建推广计划的结果，会通过异步消息方式通知，消息体包含外部业务单号，来判断该请求创建结果。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// SuccessDeliveryConfigList 是[已废弃]	创建成功的投放信息列表		成功的投放配置列表中，会返回对应的deliveryConfig对象。
	SuccessDeliveryConfigList *[]SuccessDeliveryConfig `json:"success_delivery_config_list,omitempty"`
}

// ActivityDeliveryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryQueryResponse struct {
	DeliveryBaseInfo *DeliveryBaseInfo `json:"delivery_base_info,omitempty"`
	// DeliveryBoothCode 是运营计划的展位编码。 枚举值： SERVICE_MESSAGE：商家消息（包含订单、订阅、其他消息）
	DeliveryBoothCode *string `json:"delivery_booth_code,omitempty"`
	// DeliveryErrorMsg 是投放计划错误信息描述，如投放计划审核失败时为审核失败原因。
	DeliveryErrorMsg *string `json:"delivery_error_msg,omitempty"`
	// DeliveryId 是投放计划id
	DeliveryId *string `json:"delivery_id,omitempty"`
	// DeliveryInfoList 是"[已废弃] 投放信息列表"
	DeliveryInfoList   *[]PromoDeliveryInfo `json:"delivery_info_list,omitempty"`
	DeliveryPlayConfig *DeliveryPlayConfig  `json:"delivery_play_config,omitempty"`
	// DeliveryStatus 是投放计划状态。
	DeliveryStatus     *string             `json:"delivery_status,omitempty"`
	DeliveryTargetRule *DeliveryTargetRule `json:"delivery_target_rule,omitempty"`
}

// ActivityDeliveryStopResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryStopResponse struct {
	// DeliveryId 是推广计划id。
	DeliveryId *string `json:"delivery_id,omitempty"`
	// ErrorDeliveryConfigList 是停止失败的投放配置列表
	ErrorDeliveryConfigList *[]ErrorDeliveryConfig `json:"error_delivery_config_list,omitempty"`
	// SuccessDeliveryConfigList 是停止成功的投放配置列表
	SuccessDeliveryConfigList *[]SuccessDeliveryConfig `json:"success_delivery_config_list,omitempty"`
}
