// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package servicerecord

import (
	"encoding/json"
)

// CcmServiceRecordCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceRecordCreateResponse struct {
	// ServiceRecordId 是服务记录ID
	ServiceRecordId *string `json:"service_record_id,omitempty"`
}

// CcmServiceRecordModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceRecordModifyResponse = json.RawMessage
