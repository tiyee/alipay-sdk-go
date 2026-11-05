// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerapp

// MiniInnerAppCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppCreateResponse struct {
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// MiniInnerAppPluginOrderResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginOrderResponse struct {
	// OrderId 是订单号。
	OrderId *string `json:"order_id,omitempty"`
}

// MiniInnerAppPluginRelationQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginRelationQueryResponse struct {
	// Page 是页码。
	Page *int64 `json:"page,omitempty"`
	// PageSize 是页大小。
	PageSize *int64 `json:"page_size,omitempty"`
	// PluginRelationInfoList 是插件关联关系信息。
	PluginRelationInfoList *[]PluginUseRelationInfo `json:"plugin_relation_info_list,omitempty"`
	// TotalCount 是总数。
	TotalCount *int64 `json:"total_count,omitempty"`
}

// MiniInnerAppPluginServicePublishResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginServicePublishResponse struct {
	// AbilityId 是PL2012211102012056。
	AbilityId *string `json:"ability_id,omitempty"`
}

// MiniInnerAppPluginSyncModeModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginSyncModeModifyResponse struct {
}

// MiniInnerAppPluginUseConfigCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigCancelResponse struct {
}

// MiniInnerAppPluginUseConfigOnlineResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigOnlineResponse struct {
}

// MiniInnerAppPluginUseConfigUpgradeResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppPluginUseConfigUpgradeResponse struct {
}

// MiniInnerAppServicePublishResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppServicePublishResponse struct {
	// MerchandiseId 是服务ID。
	MerchandiseId *string `json:"merchandise_id,omitempty"`
}

// MiniInnerAppServiceQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerAppServiceQueryResponse struct {
	// MiniAppServiceInfoList 是小程序服务信息。
	MiniAppServiceInfoList *[]MiniAppServiceInfo `json:"mini_app_service_info_list,omitempty"`
	// Page 是当前页数。
	Page *int64 `json:"page,omitempty"`
	// PageSize 是页大小。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总数。
	TotalCount *int64 `json:"total_count,omitempty"`
}
