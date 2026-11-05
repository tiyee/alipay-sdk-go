package certify

// UserCertifyOpenInitializeRequest 是支付宝官方 V3 OpenAPI 模型。
type UserCertifyOpenInitializeRequest struct {
	// BizCode 是认证场景码。入参支持的认证场景码和商户签约的认证场景相关，取值如下:	FACE：多因子人脸认证	CERT_PHOTO：多因子证照认证	CERT_PHOTO_FACE ：多因子证照和人脸认证	SMART_FACE：多因子快捷认证
	BizCode *string `json:"biz_code,omitempty"`
	// FaceContrastPicture 是自定义人脸比对图片的base64编码格式的string字符串
	FaceContrastPicture *string                    `json:"face_contrast_picture,omitempty"`
	IdentityParam       *OpenCertifyIdentityParam  `json:"identity_param,omitempty"`
	MerchantConfig      *OpenCertifyMerchantConfig `json:"merchant_config,omitempty"`
	// OuterOrderNo 是商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	OuterOrderNo *string `json:"outer_order_no,omitempty"`
}

// UserCertifyOpenQueryRequest is the query request model for alipay.user.certify.open.query.
type UserCertifyOpenQueryRequest struct {
	// CertifyId 是本次申请操作的唯一标识，通过alipay.user.certify.open.initialize(身份认证初始化服务)接口同步响应获取。
	CertifyId *string `json:"certify_id,omitempty" query:"certify_id,omitempty"`
}
