// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package open

// AutoInfoQueryResponse is the official V3 AliosOpenAutoInfoQueryResponseModel schema.
type AutoInfoQueryResponse struct {
	// EngineNo represents engine_no. 发动机号
	EngineNo *string `json:"engine_no,omitzero"`
	// LicenseNo represents license_no. 车牌号
	LicenseNo *string `json:"license_no,omitzero"`
	// UserActivationTime represents user_activation_time. 用户激活时间
	UserActivationTime *string `json:"user_activation_time,omitzero"`
	// VehicleType represents vehicle_type. 车辆类型
	VehicleType *string `json:"vehicle_type,omitzero"`
	// Vin represents vin. 车辆识别号码
	Vin *string `json:"vin,omitzero"`
}
