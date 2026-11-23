package enterprise

// AddressInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type AddressInfoDTO struct {
	// AccountId 是共同账户id，当入参传入共同账户id时返回该值
	AccountId *string `json:"account_id,omitempty"`
	// Address 是详细地址
	Address *string `json:"address,omitempty"`
	// AddressId 是地址id
	AddressId *string `json:"address_id,omitempty"`
	// CityCode 是市(国家统一行政规划编码)
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是城市
	CityName *string `json:"city_name,omitempty"`
	// Community 是小区/楼宇
	Community *string `json:"community,omitempty"`
	// EnterpriseId 是企业id，当入参传入企业id时返回该值
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Latitude 是纬度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度
	Longitude *string `json:"longitude,omitempty"`
	// Mark 是备注
	Mark *string `json:"mark,omitempty"`
	// PoiId 是地图poi
	PoiId *string `json:"poi_id,omitempty"`
}

// EnterpriseAgreementDTO 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseAgreementDTO struct {
	// AgreementId 是协议id
	AgreementId *string `json:"agreement_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// EnterpriseName 是企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// SignDate 是签约时间
	SignDate *string `json:"sign_date,omitempty"`
	// SignStatus 是签约状态
	SignStatus *string `json:"sign_status,omitempty"`
}

// ReliableEnterpriseProfilesDTO 是支付宝官方 V3 OpenAPI 模型。
type ReliableEnterpriseProfilesDTO struct {
	// AccessChannel 是企业注册来源渠道
	AccessChannel *string `json:"access_channel,omitempty"`
	// ActivityLevel 是企业在来源渠道（如钉钉）的活跃度
	ActivityLevel *string `json:"activity_level,omitempty"`
	// AuthType 是企业在来源渠道（如钉钉）的认证类型
	AuthType *string `json:"auth_type,omitempty"`
	// BizLicenseNo 是企业营业执照编码，适用于未三证合一的企业
	BizLicenseNo *string `json:"biz_license_no,omitempty"`
	// BizLicenseUrl 是企业在来源渠道（如钉钉）的营业执照照片URL
	BizLicenseUrl *string `json:"biz_license_url,omitempty"`
	// CountEmployee 是企业当前的员工数规模区间
	CountEmployee *string `json:"count_employee,omitempty"`
	// LegalPersonCertNo 是企业法人的身份证号码
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty"`
	// PlatformGmtCreate 是企业在来源渠道（如钉钉）的注册时间
	PlatformGmtCreate *string `json:"platform_gmt_create,omitempty"`
	// SignFundInstId 是机构id
	SignFundInstId *string `json:"sign_fund_inst_id,omitempty"`
	// TaxNo 是企业纳税人识别号，适用于未三证合一的企业
	TaxNo *string `json:"tax_no,omitempty"`
	// UtmChain 是用于标识企业入驻链路
	UtmChain *string `json:"utm_chain,omitempty"`
	// UtmContent 是具体流量内容区分，如推荐码、推荐人、推荐服务商名称等。由渠道侧自定义
	UtmContent *string `json:"utm_content,omitempty"`
	// UtmSource 是渠道侧流量来源，由渠道侧自定义
	UtmSource *string `json:"utm_source,omitempty"`
}
