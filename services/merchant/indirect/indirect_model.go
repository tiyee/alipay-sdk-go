package indirect

// AuthIdentityInfo 是支付宝官方 V3 OpenAPI 模型。
type AuthIdentityInfo struct {
	CertificateInfo *IndirectCertificateInfo `json:"certificate_info,omitempty"`
	// CertificateType 是主体类型为企业/个体户/政府机关/事业单位/社会组织时，必填；	证照类型：营业执照(BUSINESS_CERT)/登记证书(REGISTER_CERT)	主体为政府机关/事业单位/社会组织时，填登记证书；	主体类型为企业/个体户时，填营业执照；
	CertificateType *string `json:"certificate_type,omitempty"`
	// EmployerLetterImg 是单位证明函照片（使用图片上传接口）主体类型为政府机关/事业单位时，单位证明函照片必填
	EmployerLetterImg *string                   `json:"employer_letter_img,omitempty"`
	FinancialOrgInfo  *IndirectFinancialOrgInfo `json:"financial_org_info,omitempty"`
	// IdentityType 是主体类型，枚举定义：企业(ENTERPRISE)、个体工商户(IND_BIZ)、事业单位(INST)、党政机关(GOV)、社会组织(ORG)、小微商户(MSE)
	IdentityType *string `json:"identity_type,omitempty"`
	// IsFinancialOrg 是是否金融机构
	IsFinancialOrg *bool `json:"is_financial_org,omitempty"`
	// QualificationInfoList 是经营许可证列表，填写特殊行业的经营许可证信息，一个主体最多5个行业
	QualificationInfoList *[]IndirectQualificationInfo `json:"qualification_info_list,omitempty"`
	SupportCredentials    *IndirectSupportCredentials  `json:"support_credentials,omitempty"`
}

// IndirectAuthOrderFailedReason 是支付宝官方 V3 OpenAPI 模型。
type IndirectAuthOrderFailedReason struct {
	// FailParam 是审核失败字段
	FailParam *string `json:"fail_param,omitempty"`
	// FailReason 是描述申请单审核失败原因
	FailReason *string `json:"fail_reason,omitempty"`
}

// IndirectBenefitPersonInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectBenefitPersonInfo struct {
	// CardBackImg 是受益人证件反面照（使用图片上传接口的image_id）。当证件类型为身份证时，此字段必传；当证件类型为非身份证时，此字段非必传。
	CardBackImg *string `json:"card_back_img,omitempty"`
	// CardFrontImg 是受益人证件正面照（使用图片上传接口的image_id）
	CardFrontImg *string `json:"card_front_img,omitempty"`
	// CardNo 是证件号码
	CardNo *string `json:"card_no,omitempty"`
	// CardType 是证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型，政府机关、小微商户仅支持身份证类型。
	CardType *string `json:"card_type,omitempty"`
	// EffectTime 是证件生效时间
	EffectTime *string `json:"effect_time,omitempty"`
	// ExpireTime 是证件过期时间
	ExpireTime *string `json:"expire_time,omitempty"`
	// PersonName 是受益人姓名
	PersonName *string `json:"person_name,omitempty"`
}

// IndirectCertificateInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectCertificateInfo struct {
	// CertImage 是证照图片（使用图片上传接口）
	CertImage *string `json:"cert_image,omitempty"`
	// CertNo 是证件编号
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是登记证书类型时必填，枚举：统一社会信用代码证书（UNIT_SOCIAL_CREDIT）、慈善组织公开募捐资格证书（CHARITY_ORG_SOLICIT），社会团体法人登记证书（SOCIAL_ORG_LEGAL），民办非企业单位登记证书（CIVIL_UN_ENT），基金会法人登记证书（FOUNDATION_LEGAL_PERSON），农民专业合作社法人营业执照（FARMERS_COOPERATE），宗教活动场所登记证（RELIGION_PLACES），其他证书/批文/证明（OTHER_REG_CERT）
	CertType *string `json:"cert_type,omitempty"`
	// EffectTime 是证照生效时间
	EffectTime *string `json:"effect_time,omitempty"`
	// ExpireTime 是证照过期时间，yyyy-mm-dd格式，如果是长期有效时传“forever”
	ExpireTime *string `json:"expire_time,omitempty"`
	// LegalPersonName 是证照法人姓名
	LegalPersonName *string `json:"legal_person_name,omitempty"`
	// MerchantName 是证照商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// RegisterAddress 是证照注册地址
	RegisterAddress *string `json:"register_address,omitempty"`
}

// IndirectContactPersonInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectContactPersonInfo struct {
	// ContactCardNo 是联系人身份证件号。	可传身份证、港澳居民来往内地通行证、港澳居民居住证、台湾居民来往大陆通行证、台湾居民居住证、护照、外国人永久居住证等证件号码，号码枚举见contact_card_type。若传身份证，需保证18位，且英文字母大写
	ContactCardNo *string `json:"contact_card_no,omitempty"`
	// ContactCertType 是联系人证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型。小微商户仅支持身份证类型。
	ContactCertType *string `json:"contact_cert_type,omitempty"`
	// ContactName 是联系人姓名
	ContactName *string `json:"contact_name,omitempty"`
	// ContactPhoneNo 是联系人手机
	ContactPhoneNo *string `json:"contact_phone_no,omitempty"`
}

// IndirectExtraCredentials 是支付宝官方 V3 OpenAPI 模型。
type IndirectExtraCredentials struct {
	// SmidList 是待确认商户号列表，最多20个
	SmidList *[]string `json:"smid_list,omitempty"`
}

// IndirectFinancialOrgInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectFinancialOrgInfo struct {
	// FinancialOrgCertImg 是金融机构许可证图片，最多五张（使用图片上传接口）
	FinancialOrgCertImg *[]string `json:"financial_org_cert_img,omitempty"`
	// FinancialOrgType 是银行业(BANK)、支付机构(PAYMENT)、保险业(INSURANCE)、交易及结算类金融机构(SETTLE)、OTHER(OTHER)
	FinancialOrgType *string `json:"financial_org_type,omitempty"`
}

// IndirectLegalPersonInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectLegalPersonInfo struct {
	// AuthLetterImg 是授权函照片（使用图片上传接口），当证件持有人类型为 经办人时，必须上传。其他情况，无需上传
	AuthLetterImg *string `json:"auth_letter_img,omitempty"`
	// CardBackImg 是法人/经营者证件反面照（使用图片上传接口的image_id）。当证件类型为身份证时，此字段必传；当证件类型为非身份证时，此字段非必传。
	CardBackImg *string `json:"card_back_img,omitempty"`
	// CardFrontImg 是法人/经营者证件正面照（使用图片上传接口返回的image_id）
	CardFrontImg *string `json:"card_front_img,omitempty"`
	// CardNo 是证件号码
	CardNo *string `json:"card_no,omitempty"`
	// CardType 是证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型，政府机关、小微商户仅支持身份证类型。
	CardType *string `json:"card_type,omitempty"`
	// EffectTime 是证件生效时间
	EffectTime *string `json:"effect_time,omitempty"`
	// ExpireTime 是证件过期时间
	ExpireTime *string `json:"expire_time,omitempty"`
	// IsBenefitPerson 是经营者/法人是否为受益人
	IsBenefitPerson *bool `json:"is_benefit_person,omitempty"`
	// LegalType 是证件持有人类型：	- 法人  LEGAL_PERSON	- 经办人  AGENT_PERSON	主体类型为政府机关、事业单位时选传：	（1）若上传的是法人证件，则不需要上传该字段。	（2）若因特殊情况，无法提供法人证件时，可上传经办人。	主体类型为企业、个体户、社会组织、小微商户时，默认为经营者/法人，不需要上传该字段。
	LegalType *string `json:"legal_type,omitempty"`
	// PersonName 是法人姓名
	PersonName *string `json:"person_name,omitempty"`
}

// IndirectQualificationInfo 是支付宝官方 V3 OpenAPI 模型。
type IndirectQualificationInfo struct {
	// ImageList 是行业经营许可证资质照片，一个行业类目下最多上传6张资质照片（使用图片上传接口）
	ImageList *[]string `json:"image_list,omitempty"`
	// MccCode 是行业类目编号，支付宝商家行业二级类目code
	MccCode *string `json:"mcc_code,omitempty"`
}

// IndirectSupportCredentials 是支付宝官方 V3 OpenAPI 模型。
type IndirectSupportCredentials struct {
	// City 是门店城市(该值已废弃，以city_code字段映射的城市名称为准)
	City *string `json:"city,omitempty"`
	// CityCode 是门店市行政区号（数字）
	CityCode *string `json:"city_code,omitempty"`
	// District 是门店所在区县。(该值已废弃，以district_code字段映射的区县名称为准)
	District *string `json:"district,omitempty"`
	// DistrictCode 是门店街道区号（数字）
	DistrictCode *string `json:"district_code,omitempty"`
	// MerchantType 是小微商户经营类型，枚举值：门店场所（STORE）、流动经营（STALL）
	MerchantType *string `json:"merchant_type,omitempty"`
	// Province 是门店省份(该值已废弃，以province_code字段映射的省份名称为准)
	Province *string `json:"province,omitempty"`
	// ProvinceCode 是门店省行政区号（数字
	ProvinceCode *string `json:"province_code,omitempty"`
	// StoreAddress 是门店场所填写门店详细地址，流动经营类型填“无”
	StoreAddress *string `json:"store_address,omitempty"`
	// StoreDoorImg 是门店门头照信息或摊位照（使用图片上传接口）
	StoreDoorImg *string `json:"store_door_img,omitempty"`
	// StoreInnerImg 是门店店内照片或者摊位照侧面（使用图片上传接口）
	StoreInnerImg *string `json:"store_inner_img,omitempty"`
	// StoreName 是门店名称
	StoreName *string `json:"store_name,omitempty"`
}
