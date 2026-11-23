package iot

import (
	"encoding/json"
)

// MerchantIoTDeviceBindResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIoTDeviceBindResponse = json.RawMessage

// MerchantIoTDeviceQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIoTDeviceQueryResponse struct {
	// MerchantType 是商户类型。direct表示直连商户，indirect表示间连商户。
	MerchantType *string `json:"merchant_type,omitempty"`
	// Pid 是设备绑定的商户支付宝pid。直连场景为商户收单pid，间连场景为商户smid绑定的pid。
	Pid *string `json:"pid,omitempty"`
	// ShopId 是设备绑定的门店ID。
	ShopId *string `json:"shop_id,omitempty"`
	// Smid 是商户smid。直连场景为空。
	Smid *string `json:"smid,omitempty"`
}

// MerchantIoTDeviceVerifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIoTDeviceVerifyResponse = json.RawMessage
