// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package delivery

// ActivityDeliveryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryCreateRequest struct {
	BelongMerchantInfo *DeliveryAgencyMerchantInfo `json:"belong_merchant_info,omitempty"`
	DeliveryBaseInfo   *DeliveryBaseInfo           `json:"delivery_base_info,omitempty"`
	// DeliveryBoothCode 是投放的展位编码。
	DeliveryBoothCode *string `json:"delivery_booth_code,omitempty"`
	// DeliveryConfigList 是[已废弃]	待创建的投放配置列表。	最大数量限制20个。
	DeliveryConfigList *[]DeliveryConfig   `json:"delivery_config_list,omitempty"`
	DeliveryPlayConfig *DeliveryPlayConfig `json:"delivery_play_config,omitempty"`
	DeliveryTargetRule *DeliveryTargetRule `json:"delivery_target_rule,omitempty"`
	// IdemNo 是商品卖货推广创建幂等号，用作幂等控制。 幂等作用：请求返回与上一次相同的结果。 外部接入方需保证幂等号唯一。
	IdemNo *string `json:"idem_no,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。 幂等作用：请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// ActivityDeliveryQueryRequest is the request model for alipay.marketing.activity.delivery.query.
type ActivityDeliveryQueryRequest struct {
	MarketingActivityDeliveryQuery `json:",inline"`
	// DeliveryId 是推广计划id
	DeliveryId string `json:"-" path:"delivery_id"`
}

// ActivityDeliveryStopRequest is the request model for alipay.marketing.activity.delivery.stop.
type ActivityDeliveryStopRequest struct {
	MarketingActivityDeliveryStop `json:",inline"`
	// DeliveryId 是推广计划id
	DeliveryId string `json:"-" path:"delivery_id"`
}
