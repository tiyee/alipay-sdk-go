// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package authentication

// FaceExtInfo is the official V3 FaceExtInfo schema.
type FaceExtInfo struct {
	// MaxAge represents max_age. 年龄区间判断的上限，闭区间
	MaxAge *string `json:"max_age,omitzero"`
	// MinAge represents min_age. 年龄区间判断的下限，闭区间
	MinAge *string `json:"min_age,omitzero"`
	// QueryType represents query_type. query_type不填, 返回uid query_type=1, 返回手机号 query_type=2, 返回图片
	QueryType *string `json:"query_type,omitzero"`
}

// ZhubUidTelPair is the official V3 ZhubUidTelPair schema.
type ZhubUidTelPair struct {
	// OpenId represents open_id. 支付宝用户open_id
	OpenId *string `json:"open_id,omitzero"`
	// Phone represents phone. 手机号
	Phone *string `json:"phone,omitzero"`
	// UserId represents user_id. 支付宝uid
	UserId *string `json:"user_id,omitzero"`
}

// FaceExtParams is the official V3 FaceExtParams schema.
type FaceExtParams struct {
	// BizType represents biz_type. 业务类型：7，基于1:N人脸搜索的刷脸支付场景；8，基于姓名和身份证号的刷脸支付场景。
	BizType *string `json:"biz_type,omitzero"`
}

// FaceMachineInfo is the official V3 FaceMachineInfo schema.
type FaceMachineInfo struct {
	// CameraDriveVer represents camera_drive_ver. 摄像头驱动版本号
	CameraDriveVer *string `json:"camera_drive_ver,omitzero"`
	// CameraModel represents camera_model. 摄像头型号
	CameraModel *string `json:"camera_model,omitzero"`
	// CameraName represents camera_name. 摄像头名称
	CameraName *string `json:"camera_name,omitzero"`
	// CameraVer represents camera_ver. 摄像头版本号
	CameraVer *string `json:"camera_ver,omitzero"`
	// Ext represents ext. 扩展信息
	Ext *string `json:"ext,omitzero"`
	// MachineCode represents machine_code. 机具编码
	MachineCode *string `json:"machine_code,omitzero"`
	// MachineModel represents machine_model. 机具型号
	MachineModel *string `json:"machine_model,omitzero"`
	// MachineVer represents machine_ver. 机具版本号
	MachineVer *string `json:"machine_ver,omitzero"`
}

// FaceMerchantInfo is the official V3 FaceMerchantInfo schema.
type FaceMerchantInfo struct {
	// AreaCode represents area_code. 区域编码
	AreaCode *string `json:"area_code,omitzero"`
	// BrandCode represents brand_code. 品牌编码
	BrandCode *string `json:"brand_code,omitzero"`
	// DeviceMac represents device_mac. 机具Mac地址
	DeviceMac *string `json:"device_mac,omitzero"`
	// DeviceNum represents device_num. 机具编码
	DeviceNum *string `json:"device_num,omitzero"`
	// Geo represents geo. 经纬度
	Geo *string `json:"geo,omitzero"`
	// Group represents group. 机具分组编码
	Group *string `json:"group,omitzero"`
	// MerchantId represents merchant_id. 商户ID
	MerchantId *string `json:"merchant_id,omitzero"`
	// PartnerId represents partner_id. ISV ID
	PartnerId *string `json:"partner_id,omitzero"`
	// StoreCode represents store_code. 门店编码
	StoreCode *string `json:"store_code,omitzero"`
	// Wifimac represents wifimac. WI-FI Mac地址
	Wifimac *string `json:"wifimac,omitzero"`
	// Wifiname represents wifiname. WI-FI 名称
	Wifiname *string `json:"wifiname,omitzero"`
}
