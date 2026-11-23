package indirect

// MerchantIndirectAuthOrderCloseRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderCloseRequest struct {
	// OrderNo 是意愿申请单号，参数二选一
	OrderNo *string `json:"order_no,omitempty"`
	// OutBizNo 是外部业务号，参数二选一，业务自定义，保证唯一
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// MerchantIndirectAuthOrderCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderCreateRequest struct {
	AuthIdentityInfo  *AuthIdentityInfo          `json:"auth_identity_info,omitempty"`
	BenefitPersonInfo *IndirectBenefitPersonInfo `json:"benefit_person_info,omitempty"`
	ContactPersonInfo *IndirectContactPersonInfo `json:"contact_person_info,omitempty"`
	ExtraCredentials  *IndirectExtraCredentials  `json:"extra_credentials,omitempty"`
	LegalPersonInfo   *IndirectLegalPersonInfo   `json:"legal_person_info,omitempty"`
	// OutBizNo 是外部业务号。支持out_biz_no和outBizNo两种格式
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Source 是1. 收单机构调用API提交申请单时，可选择是否指定单个服务商范围。非收单机构无需填写此字段。	2. 此字段填写单个服务商pid信息：填写（即：单服务商提交认证方式），查询申请单返回的认证二维码qr_code和填写服务商对应，仅能认证填写服务商下的商户；不填写（即：全服务商提交认证方式），查询申请单返回的认证二维码qr_code和收单机构对应，可认证收单机构下全部商户。
	Source *string `json:"source,omitempty"`
}

// MerchantIndirectAuthOrderQueryStatusRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderQueryStatusRequest struct {
	// OrderNo 是商家认证申请单号，参数二选一
	OrderNo *string `json:"order_no,omitempty"`
	// OutBizNo 是外部业务号，参数二选一，业务自定义，保证唯一
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Source 是1. 收单机构调用API提交申请单时，可选择是否指定单个服务商范围。非收单机构无需填写此字段。	2. 此字段填写单个服务商pid信息：填写（即：单服务商提交认证方式），查询申请单返回的认证二维码qr_code和填写服务商对应，仅能认证填写服务商下的商户；不填写（即：全服务商提交认证方式），查询申请单返回的认证二维码qr_code和收单机构对应，可认证收单机构下全部商户。
	Source *string `json:"source,omitempty"`
}

// MerchantIndirectSMIDBindQueryRequest is the query request model for alipay.merchant.indirect.smidbind.query.
type MerchantIndirectSMIDBindQueryRequest struct {
	// SubMerchantId 是支付宝商户id
	SubMerchantId *string `json:"sub_merchant_id,omitempty" query:"sub_merchant_id,omitempty"`
}
