// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package operation

import (
	"encoding/json"
)

// OperationActivityMerchantModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type OperationActivityMerchantModifyResponse = json.RawMessage

// OperationActivityMerchantSignResponse 是支付宝官方 V3 OpenAPI 模型。
type OperationActivityMerchantSignResponse = json.RawMessage

// OperationActivityMerchantUnsignResponse 是支付宝官方 V3 OpenAPI 模型。
type OperationActivityMerchantUnsignResponse = json.RawMessage

// OperationBankActivityQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type OperationBankActivityQueryResponse struct {
	// ActivityInfoList 是活动信息列表
	ActivityInfoList *[]BankActivityInfo `json:"activity_info_list,omitempty"`
	// MerchantTag 是商户标签
	MerchantTag *string `json:"merchant_tag,omitempty"`
}
