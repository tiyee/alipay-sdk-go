package iot

// MerchantIoTDeviceBindRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantIoTDeviceBindRequest struct {
	// BizTid 是设备 ID ，仅device_id_type 为 ID 时填写。
	BizTid *string `json:"biz_tid,omitempty"`
	// DeviceIdType 是可选方式。枚举支持：	*  ID：表示使用 biztid 作为设备唯一识别标识，	*  SN：表示使用 supplier_id、device_sn联合作为设备唯一识别标识。	注意：由于不同机型的supplier_id不同，推荐使用 ID 。
	DeviceIdType *string `json:"device_id_type,omitempty"`
	// DeviceSn 是设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。
	DeviceSn *string `json:"device_sn,omitempty"`
	// MerchantType 是商户类型。枚举支持：	*  direct：直连商户；	*  indirect：间连商户。
	MerchantType *string `json:"merchant_type,omitempty"`
	// Pid 是直连场景填写商户收单pid。间连场景填写smid绑定的pid（可通过代运营授权API获取）。
	Pid *string `json:"pid,omitempty"`
	// ShopId 是商户pid名下的门店 ID。可通过门店管理API获取，需注意请勿使用smid创建门店 ID。
	ShopId *string `json:"shop_id,omitempty"`
	// Smid 是直连场景不填，间连场景填写商户收单smid
	Smid *string `json:"smid,omitempty"`
	// SupplierId 是设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
	SupplierId *string `json:"supplier_id,omitempty"`
}

// MerchantIoTDeviceVerifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantIoTDeviceVerifyRequest struct {
	// BizTid 是设备 ID ，device_id_type 为 ID 时填写。
	BizTid *string `json:"biz_tid,omitempty"`
	// DeviceIdType 是可选方式 [ID,SN]。ID-使用biztid作为设备唯一识别标识；SN-使用supplier_id、device_sn联合作为设备唯一识别标识。由于不同机型的supplier_id不同，推荐使用 ID 。
	DeviceIdType *string `json:"device_id_type,omitempty"`
	// DeviceSn 是设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。
	DeviceSn *string `json:"device_sn,omitempty"`
	// MerchantType 是商户类型，直连商户填写direct，间连商户填写indirect
	MerchantType *string `json:"merchant_type,omitempty"`
	// Pid 是直连场景填写商户收单pid，间连场景不填
	Pid *string `json:"pid,omitempty"`
	// Smid 是直连场景不填，间连场景填写商户收单smid
	Smid *string `json:"smid,omitempty"`
	// SupplierId 是设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
	SupplierId *string `json:"supplier_id,omitempty"`
}

// MerchantIoTDeviceQueryRequest is the query request model for alipay.merchant.iot.device.query.
type MerchantIoTDeviceQueryRequest struct {
	// DeviceIdType 是可选方式 [ID,SN]。ID-使用biztid作为设备唯一识别标识；SN-使用supplier_id、device_sn联合作为设备唯一识别标识。由于不同机型的supplier_id不同，推荐使用 ID 。
	DeviceIdType *string `json:"device_id_type,omitempty" query:"device_id_type,omitempty"`
	// BizTid 是设备 ID ，device_id_type 为 ID 时填写。
	BizTid *string `json:"biz_tid,omitempty" query:"biz_tid,omitempty"`
	// SupplierId 是设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
	SupplierId *string `json:"supplier_id,omitempty" query:"supplier_id,omitempty"`
	// DeviceSn 是设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。
	DeviceSn *string `json:"device_sn,omitempty" query:"device_sn,omitempty"`
}
