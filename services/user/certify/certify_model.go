package certify

// OpenCertifyIdentityParam 是支付宝官方 V3 OpenAPI 模型。
type OpenCertifyIdentityParam struct {
	// CertDigest 是cert_digest：	身份信息摘要
	CertDigest *string `json:"cert_digest,omitempty"`
	// CertName 是cert_name：	填入真实姓名
	CertName *string `json:"cert_name,omitempty"`
	// CertNo 是cert_no：	填入姓名相匹配的证件号码
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是cert_type：	1、若为身份证，填IDENTITY_CARD；	2、若为港澳居民来往内地通行证，填HOME_VISIT_PERMIT_HK_MC；	3、若为台湾居民来往内地通行证，填HOME_VISIT_PERMIT_TAIWAN；	4、若为港澳居民居住证，填RESIDENCE_PERMIT_HK_MC；	5、若为台湾居民居住证，填RESIDENCE_PERMIT_TAIWAN；	6、外国人永久居留证，填PERMANENT_RESIDENCE_FOREIGNER；
	CertType *string `json:"cert_type,omitempty"`
	// FacialPictureFront 是facial_picture_front：	活体人脸图片base64串，主要用于离线人脸比对
	FacialPictureFront *string `json:"facial_picture_front,omitempty"`
	// IdentityType 是identity_type	1.若本人验证，使用CERT_INFO；	2.若代他人验证，使用AGENT_CERT_INFO；
	IdentityType *string `json:"identity_type,omitempty"`
	// OpenId 是open_id：	开放平台用户ID
	OpenId *string `json:"open_id,omitempty"`
	// PhoneNo 是phone_no：	填写姓名相匹配的手机号
	PhoneNo *string `json:"phone_no,omitempty"`
	// UserId 是user_id：	支付宝用户ID
	UserId *string `json:"user_id,omitempty"`
}

// OpenCertifyMerchantConfig 是支付宝官方 V3 OpenAPI 模型。
type OpenCertifyMerchantConfig struct {
	// AuthScope 是用于开放认证授权
	AuthScope *string `json:"auth_scope,omitempty"`
	// AuthType 是用于指定授权类型，与auth_scope配合使用
	AuthType *string `json:"auth_type,omitempty"`
	// FaceReserveStrategy 是不传默认为reserve
	FaceReserveStrategy *string `json:"face_reserve_strategy,omitempty"`
	// FacialPictureLevel 是若有特殊人脸等级采集要求，可指定等级
	FacialPictureLevel *string `json:"facial_picture_level,omitempty"`
	// LinkedMerchantAppId 是用于授权二级商户操作
	LinkedMerchantAppId *string `json:"linked_merchant_app_id,omitempty"`
	// LinkedMerchantLogoUrl 是用于指定展示的商户logo
	LinkedMerchantLogoUrl *string `json:"linked_merchant_logo_url,omitempty"`
	// LinkedMerchantName 是用于指定展示的商户名称
	LinkedMerchantName *string `json:"linked_merchant_name,omitempty"`
	// OutPutFacialPicture 是在拥有该权限前提下，用于商户控制是否透出活体人脸
	OutPutFacialPicture *bool `json:"out_put_facial_picture,omitempty"`
	// ReturnUrl 是认证成功后需要跳转的地址，一般为商户业务页面；若无跳转地址可填空字符"";
	ReturnUrl *string `json:"return_url,omitempty"`
}
