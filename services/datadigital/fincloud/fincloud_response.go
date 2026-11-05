// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fincloud

// GeneralSaaSFaceCertifyInitializeResponse is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyInitializeResponseModel schema.
type GeneralSaaSFaceCertifyInitializeResponse struct {
	// CertifyId represents certify_id. 本次申请操作的唯一标识，商户需要记录，后续的操作都需要用到
	CertifyId *string `json:"certify_id,omitzero"`
}

// GeneralSaaSFaceCertifyQueryResponse is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyQueryResponseModel schema.
type GeneralSaaSFaceCertifyQueryResponse struct {
	// FailReason represents fail_reason. 认证错误原因，扩展字段，默认不返回
	FailReason *string `json:"fail_reason,omitzero"`
	// IdentityInfo represents identity_info. 刷脸用户的身份信息，目前返回为空
	IdentityInfo *string `json:"identity_info,omitzero"`
	// MaterialInfo represents material_info. 认证主体附件信息，主要为图片类材料，一般的认证场景都是返回空
	MaterialInfo *string `json:"material_info,omitzero"`
	// Passed represents passed. 是否通过，通过为T，不通过为F
	Passed *string `json:"passed,omitzero"`
}

// GeneralSaaSFaceCertifyVerifyResponse is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyVerifyResponseModel schema.
type GeneralSaaSFaceCertifyVerifyResponse struct {
	// CertifyUrl represents certify_url. 返回用于唤起刷脸页面的url
	CertifyUrl *string `json:"certify_url,omitzero"`
}

// GeneralSaaSFaceCheckInitializeResponse is the official V3 DatadigitalFincloudGeneralsaasFaceCheckInitializeResponseModel schema.
type GeneralSaaSFaceCheckInitializeResponse struct {
	// CertifyId represents certify_id. 人脸检测流水ID，请保留方便排查问题
	CertifyId *string `json:"certify_id,omitzero"`
	// PageUrl represents page_url. 检测页面url，需要给到终端SDK(该参数后续不再维护，建议使用该参数的商户进行如下升级： 1.在App端活体检测场景中使用了该参数的商户，新版客户端SDK不再依赖该参数，请参考产品官方文档接入指南中的<a href="https://opendocs.alipay.com/open/04pxpy?pathHash=950007fa&ref=api">App端活体检测</a>部分，进行客户端升级； 2.在H5活体检测场景中使用了该参数的商户，请替换为web_url进行认证。)
	PageUrl *string `json:"page_url,omitzero"`
	// WebUrl represents web_url. H5活体检测地址
	WebUrl *string `json:"web_url,omitzero"`
}

// GeneralSaaSFaceCheckQueryResponse is the official V3 DatadigitalFincloudGeneralsaasFaceCheckQueryResponseModel schema.
type GeneralSaaSFaceCheckQueryResponse struct {
	// AlivePhoto represents alive_photo. 如果设置了need_alive_photo的值不为N，则该字段返回人脸照片的base64编码字符串。
	AlivePhoto *string `json:"alive_photo,omitzero"`
	// AttackFlag represents attack_flag. 已废弃
	AttackFlag *string `json:"attack_flag,omitzero"`
	// CertifyState represents certify_state. PROCESSING：初始化；SUCCESS：检测成功；FAIL：检测失败
	CertifyState *string `json:"certify_state,omitzero"`
	// MetaInfo represents meta_info.
	MetaInfo *OpenCertifyMetaInfo `json:"meta_info,omitzero"`
	// Quality represents quality. 人脸图片质量分，单位为分
	Quality *string `json:"quality,omitzero"`
	// Score represents score. 人脸比对分，单位分
	Score *string `json:"score,omitzero"`
}

// GeneralSaaSFaceSourceCertifyResponse is the official V3 DatadigitalFincloudGeneralsaasFaceSourceCertifyResponseModel schema.
type GeneralSaaSFaceSourceCertifyResponse struct {
	// CertifyNo represents certify_no. 人脸认证单据号，用于认证问题排查
	CertifyNo *string `json:"certify_no,omitzero"`
	// MismatchReason represents mismatch_reason. 认证不通过原因
	MismatchReason *string `json:"mismatch_reason,omitzero"`
	// Passed represents passed. T：通过；F：不通过
	Passed *string `json:"passed,omitzero"`
	// Quality represents quality. double值，人脸图片质量分，0-100的百分制分数
	Quality *string `json:"quality,omitzero"`
	// Score represents score. double值，活体检测结果分数，0-100的百分制分数
	Score *string `json:"score,omitzero"`
}

// GeneralSaaSFaceVerificationInitializeResponse is the official V3 DatadigitalFincloudGeneralsaasFaceVerificationInitializeResponseModel schema.
type GeneralSaaSFaceVerificationInitializeResponse struct {
	// CertifyId represents certify_id. 认证单据号，请保留以便排查问题。
	CertifyId *string `json:"certify_id,omitzero"`
	// PageUrl represents page_url. 人脸核身url(该参数后续不再维护，建议使用该参数的商户进行如下升级： 1.在App端人脸核身场景中使用了该参数的商户，新版客户端SDK不再依赖该参数，请参考产品官方文档接入指南中的<a href="https://opendocs.alipay.com/open/02zlo7?pathHash=3c880197&ref=api">App端人脸核身</a>部分，进行客户端升级； 2.在H5人脸核身场景中使用了该参数的商户，请替换为web_url进行认证。)
	PageUrl *string `json:"page_url,omitzero"`
	// WebUrl represents web_url. H5人脸核身地址
	WebUrl *string `json:"web_url,omitzero"`
}

// GeneralSaaSFaceVerificationQueryResponse is the official V3 DatadigitalFincloudGeneralsaasFaceVerificationQueryResponseModel schema.
type GeneralSaaSFaceVerificationQueryResponse struct {
	// AlivePhoto represents alive_photo. base64过后的图片图片
	AlivePhoto *string `json:"alive_photo,omitzero"`
	// AttackFlag represents attack_flag. 本次认证是否存在安全风险，true：检测到安全风险；false：未检测到安全风险。
	AttackFlag *string `json:"attack_flag,omitzero"`
	// CertifyState represents certify_state. 人脸认证状态。PROCESSING：初始化；SUCCESS：认证通过；FAIL：认证不通过。
	CertifyState *string `json:"certify_state,omitzero"`
	// MetaInfo represents meta_info.
	MetaInfo *OpenCertifyMetaInfo `json:"meta_info,omitzero"`
	// Quality represents quality. double值，人脸图片质量分
	Quality *string `json:"quality,omitzero"`
	// Score represents score. double值，活体检测结果分数
	Score *string `json:"score,omitzero"`
}

// GeneralSaaSOCRMobileInitializeResponse is the official V3 DatadigitalFincloudGeneralsaasOcrMobileInitializeResponseModel schema.
type GeneralSaaSOCRMobileInitializeResponse struct {
	// CertifyId represents certify_id. OCR识别流水ID，计费依据，请保留方便排查问题
	CertifyId *string `json:"certify_id,omitzero"`
}

// GeneralSaaSOCRServerDetectResponse is the official V3 DatadigitalFincloudGeneralsaasOcrServerDetectResponseModel schema.
type GeneralSaaSOCRServerDetectResponse struct {
	// CertifyId represents certify_id. OCR识别单据号，计费依据，请保留以便排查问题。
	CertifyId *string `json:"certify_id,omitzero"`
	// OcrData represents ocr_data. OCR识别结果
	OcrData *string `json:"ocr_data,omitzero"`
}
