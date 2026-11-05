// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fincloud

// GeneralSaaSFaceSourceCertify is the official V3 DatadigitalFincloudGeneralsaasFaceSourceCertifyModel schema.
type GeneralSaaSFaceSourceCertify struct {
	// BizCode represents biz_code. 业务产品码
	BizCode *string `json:"biz_code,omitzero"`
	// CertName represents cert_name. 用户姓名，与身份证上的姓名相匹配
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号码，如大陆身份证号码
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 证件信息类型，大陆身份证：IDENTITY_CARD
	CertType *string `json:"cert_type,omitzero"`
	// LivenessStrategy represents liveness_strategy. 活体检测类型，当前枚举支持： OFF：关闭活体检测； CHECK：正常活体检测。 不传默认CHECK。
	LivenessStrategy *string `json:"liveness_strategy,omitzero"`
	// Nation represents nation. 国家地区代码，当证件类型为外国人永久居留证时必填。
	Nation *string `json:"nation,omitzero"`
	// OuterBizNo represents outer_biz_no. 客户业务单据号
	OuterBizNo *string `json:"outer_biz_no,omitzero"`
	// Phone represents phone. 手机号，特殊场景下，在验证过程中需用户输入手机来接受系统发送的安全验证短信。默认不传即可。
	Phone *string `json:"phone,omitzero"`
	// Reserved represents reserved. true：需要留底；false：无需留底。默认留底
	Reserved *bool `json:"reserved,omitzero"`
	// SecurityStrategy represents security_strategy. ON：开启安全拦截。OFF：关闭安全拦截。默认取值为ON。
	SecurityStrategy *string `json:"security_strategy,omitzero"`
}

// GeneralSaaSOCRServerDetect is the official V3 DatadigitalFincloudGeneralsaasOcrServerDetectModel schema.
type GeneralSaaSOCRServerDetect struct {
	// OcrType represents ocr_type. 可识别OCR类型。如，ID_CARD_FRONT 身份证含照片的一面；ID_CARD_BACK 身份证带国徽的一面。具体支持的类型以接入文档为准。
	OcrType *string `json:"ocr_type,omitzero"`
	// OuterOrderNo represents outer_order_no. 客户业务单据号，请保证幂等性。
	OuterOrderNo *string `json:"outer_order_no,omitzero"`
}

// OpenCertifyIdentifyInfo is the official V3 OpenCertifyIdentifyInfo schema.
type OpenCertifyIdentifyInfo struct {
	// CertName represents cert_name. 填入真实姓名
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 填入姓名相匹配的证件号码
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 当前仅支持IDENTITY_CARD
	CertType *string `json:"cert_type,omitzero"`
	// IdentityType represents identity_type. 当前仅支持CERT_INFO
	IdentityType *string `json:"identity_type,omitzero"`
	// PhoneNo represents phone_no. 选填手机号
	PhoneNo *string `json:"phone_no,omitzero"`
}

// OpenCertifyMerchantConfigs is the official V3 OpenCertifyMerchantConfigs schema.
type OpenCertifyMerchantConfigs struct {
	// FaceReserveStrategy represents face_reserve_strategy. 不传默认为reserve
	FaceReserveStrategy *string `json:"face_reserve_strategy,omitzero"`
	// ReturnUrl represents return_url. 认证成功后需要跳转的地址，一般为商户业务页面；若无跳转地址可填空字符""。
	ReturnUrl *string `json:"return_url,omitzero"`
}

// OpenCertifyMetaInfo is the official V3 OpenCertifyMetaInfo schema.
type OpenCertifyMetaInfo struct {
	// DeviceType represents device_type. 设备操作系统类型
	DeviceType *string `json:"device_type,omitzero"`
}
