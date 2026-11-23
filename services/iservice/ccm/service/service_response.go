// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package service

import (
	"encoding/json"
)

// CcmServiceBuyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceBuyResponse struct {
	// ServiceInstanceId 是服务实例id
	ServiceInstanceId *int64 `json:"service_instance_id,omitempty"`
	// TenantId 是商户在CCM的租户id
	TenantId *string `json:"tenant_id,omitempty"`
}

// CcmServiceCloseResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceCloseResponse = json.RawMessage

// CcmServiceInitializeResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceInitializeResponse struct {
	// SpiIds 是spi在ccm中的唯一标识，与biz_code一一对应
	SpiIds *[]SpiResult `json:"spi_ids,omitempty"`
}
