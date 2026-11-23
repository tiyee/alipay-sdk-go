// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package service

// CcmServiceBuyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceBuyRequest struct {
	// ServiceCode 是服务代码。通过alipay.iservice.ccm.service.initialize(isv服务初始化接口) 获取。
	ServiceCode *string `json:"service_code,omitempty"`
}

// CcmServiceCloseRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceCloseRequest struct {
	// ServiceCode 是服务代码。通过alipay.iservice.ccm.service.initialize(isv服务初始化接口) 获取。
	ServiceCode *string `json:"service_code,omitempty"`
}

// CcmServiceInitializeRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceInitializeRequest struct {
	// Description 是服务描述
	Description *string `json:"description,omitempty"`
	// Icon 是服务图标url。支持 http(s) 链接。
	Icon *string `json:"icon,omitempty"`
	// ServiceCode 是服务编码。可在服务上架后，在服务市场上查询此信息。
	ServiceCode *string `json:"service_code,omitempty"`
	// ServiceName 是服务名称。可在服务上架后，在服务市场上查询此信息。
	ServiceName *string `json:"service_name,omitempty"`
	// ServiceOrderUrl 是服务订购二维码链接
	ServiceOrderUrl *string `json:"service_order_url,omitempty"`
	// Spis 是spi定义列表
	Spis *[]ISVSpiDefinition `json:"spis,omitempty"`
}
