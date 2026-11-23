// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fincloud

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// GeneralSaaSFaceCertifyInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyInitializeModel schema.
type GeneralSaaSFaceCertifyInitializeRequest struct {
	// BizCode represents biz_code. H5人脸核身场景码。入参支持的场景码。
	BizCode *string `json:"biz_code,omitzero"`
	// FaceContrastPicture represents face_contrast_picture. 自定义人脸比对图片的base64编码格式的string字符串
	FaceContrastPicture *string `json:"face_contrast_picture,omitzero"`
	// IdentityParam represents identity_param.
	IdentityParam *OpenCertifyIdentifyInfo `json:"identity_param,omitzero"`
	// MerchantConfig represents merchant_config.
	MerchantConfig *OpenCertifyMerchantConfigs `json:"merchant_config,omitzero"`
	// OuterOrderNo represents outer_order_no. 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	OuterOrderNo *string `json:"outer_order_no,omitzero"`
}

// GeneralSaaSFaceCertifyVerifyRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyVerifyModel schema.
type GeneralSaaSFaceCertifyVerifyRequest struct {
	// CertifyId represents certify_id. 本次申请操作的唯一标识，由H5人脸核身初始化接口调用后生成，后续的操作都需要用到
	CertifyId *string `json:"certify_id,omitzero"`
}

// GeneralSaaSFaceCheckInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCheckInitializeModel schema.
type GeneralSaaSFaceCheckInitializeRequest struct {
	// BizCode represents biz_code. DATA_DIGITAL_BIZ_CODE_FACE_CHECK_LIVE，代表活体检测。
	BizCode *string `json:"biz_code,omitzero"`
	// OuterOrderNo represents outer_order_no. 客户业务单据号。
	OuterOrderNo *string `json:"outer_order_no,omitzero"`
}

// GeneralSaaSFaceVerificationInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel schema.
type GeneralSaaSFaceVerificationInitializeRequest struct {
	// BizCode represents biz_code. 人脸核身具体类型目前仅支持：DATA_DIGITAL_BIZ_CODE_FACE_VERIFICATION
	BizCode *string `json:"biz_code,omitzero"`
	// CertName represents cert_name. 真实姓名
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 证件类型，当前枚举支持： IDENTITY_CARD：身份证 RESIDENCE_HK_MC：港澳居民居住证 RESIDENCE_TAIWAN：台湾居民居住证
	CertType *string `json:"cert_type,omitzero"`
	// FaceReserveStrategy represents face_reserve_strategy. 人脸保存策略，非必填。具体取值为：reserve(保存活体人脸)/never(不保存活体人脸)，不传默认为reserve
	FaceReserveStrategy *string `json:"face_reserve_strategy,omitzero"`
	// IdentityType represents identity_type. 认证类型，固定值为：CERT_INFO
	IdentityType *string `json:"identity_type,omitzero"`
	// Nation represents nation. 国家地区代码，当证件类型为外国人永久居留证时必填。
	Nation *string `json:"nation,omitzero"`
	// OuterOrderNo represents outer_order_no. 商户请求的唯一标识，商户要保证其唯一性，值为64位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	OuterOrderNo *string `json:"outer_order_no,omitzero"`
	// PhoneNo represents phone_no. 手机号码
	PhoneNo *string `json:"phone_no,omitzero"`
}

// GeneralSaaSOCRMobileInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasOcrMobileInitializeModel schema.
type GeneralSaaSOCRMobileInitializeRequest struct {
	// BizCode represents biz_code. bizCode，代表当前使用的能力类型。
	BizCode *string `json:"biz_code,omitzero"`
	// OuterOrderNo represents outer_order_no. 客户业务单据号。请保持唯一。
	OuterOrderNo *string `json:"outer_order_no,omitzero"`
}

// GeneralSaaSFaceCertifyQueryRequest contains parameters for datadigital.fincloud.generalsaas.face.certify.query.
type GeneralSaaSFaceCertifyQueryRequest struct {
	// CertifyId represents certify_id. 本次申请操作的唯一标识，通过datadigital.fincloud.generalsaas.face.certify.initialize 接口同步响应获取。
	CertifyId *string `query:"certify_id"`
}

// GeneralSaaSFaceCheckQueryRequest contains parameters for datadigital.fincloud.generalsaas.face.check.query.
type GeneralSaaSFaceCheckQueryRequest struct {
	// CertifyId represents certify_id. 传入初始化接口datadigital.fincloud.generalsaas.face.check.initialize 得到的certify_id
	CertifyId *string `query:"certify_id"`
	// NeedAlivePhoto represents need_alive_photo. 是否需要返回活体结果图片，needAlivePhoto：Y_O （原始图片）、Y_M（虚化，背景马赛克）、N（不返图）。非必填，默认为N
	NeedAlivePhoto *string `query:"need_alive_photo"`
	// NeedScore represents need_score. 该入参已经废弃。
	NeedScore *string `query:"need_score"`
	// NeedQualityScore represents need_quality_score. 是否需要返回图片质量分，非必填，默认值为N。 N表示不需要，Y表示需要
	NeedQualityScore *string `query:"need_quality_score"`
	// NeedAttackResult represents need_attack_result. 该参数已废弃
	NeedAttackResult *string `query:"need_attack_result"`
}

// GeneralSaaSFaceSourceCertifyRequest contains parameters for datadigital.fincloud.generalsaas.face.source.certify.
type GeneralSaaSFaceSourceCertifyRequest struct {
	// Data is the data multipart field.
	Data *GeneralSaaSFaceSourceCertify `multipart:"data,json"`
	// FileContent is the file_content multipart field.
	FileContent *core.MultipartFile `multipart:"file_content,file"`
}

// GeneralSaaSFaceVerificationQueryRequest contains parameters for datadigital.fincloud.generalsaas.face.verification.query.
type GeneralSaaSFaceVerificationQueryRequest struct {
	// CertifyId represents certify_id. 填入人脸核身初始化阶段获取到的certify_id
	CertifyId *string `query:"certify_id"`
	// NeedAlivePhoto represents need_alive_photo. 是否需要返回活体结果图片，默认为N（不返图）
	NeedAlivePhoto *string `query:"need_alive_photo"`
}

// GeneralSaaSOCRServerDetectRequest contains parameters for datadigital.fincloud.generalsaas.ocr.server.detect.
type GeneralSaaSOCRServerDetectRequest struct {
	// Data is the data multipart field.
	Data *GeneralSaaSOCRServerDetect `multipart:"data,json"`
	// FileContent is the file_content multipart field.
	FileContent *core.MultipartFile `multipart:"file_content,file"`
}
