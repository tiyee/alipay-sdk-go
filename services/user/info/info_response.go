package info

// UserInfoShareResponse 是支付宝官方 V3 OpenAPI 模型。
type UserInfoShareResponse struct {
	// Address 是详细地址。
	Address *string `json:"address,omitempty"`
	// Age 是用户年龄，实名认证为T，且证件类型为身份证，则输出年龄字段，否则不输出
	Age *string `json:"age,omitempty"`
	// Area 是区县名称。
	Area *string `json:"area,omitempty"`
	// Avatar 是用户头像地址。	注意：如果没有数据（用户未设置）时不会返回该信息，请做好容错。
	Avatar *string `json:"avatar,omitempty"`
	// BusinessScope 是经营/业务范围（用户类型是公司类型时才有此字段）。
	BusinessScope *string `json:"business_scope,omitempty"`
	// CertNo 是【证件号码】结合证件类型使用.【注意】只is_certified为T的时候才有意义，否则不保证准确性.
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是【证件类型】0:身份证 1:护照 2:军官证 3:士兵证 4:回乡证 5:临时身份证 6:户口簿 7:警官证 8:台胞证 9:营业执照10:其它证件 11:港澳居民来往内地通行证 12:台湾居民来往大陆通行证 13:台湾居民居住证 14:港澳居民居住证	 15:港澳身份证 16:外国人永久居住证【注意】只有is_certified为T的时候才有意义，否则不保证准确性.
	CertType *string `json:"cert_type,omitempty"`
	// City 是市名称。
	City *string `json:"city,omitempty"`
	// CollegeName 是学信网返回的学校名称，有可能为空。
	CollegeName *string `json:"college_name,omitempty"`
	// CountryCode 是国家码
	CountryCode *string `json:"country_code,omitempty"`
	// Degree 是学信网返回的学历/学位信息，数据质量一般，请谨慎使用，取值包括：博士研究生、硕士研究生、高升本、专科、博士、硕士、本科、夜大电大函大普通班、专科(高职)、第二学士学位。
	Degree *string `json:"degree,omitempty"`
	// DeliverAddresses 是收货地址列表
	DeliverAddresses *[]UserDeliverAddress `json:"deliver_addresses,omitempty"`
	// DisplayName 是展示名称，用于展示和识别用户。返回的值是支付宝账户的可用外标之一，根据展示名称产品规则可能返回手机号/邮箱/支付宝号其中的一个。同一个账户在不同的时期返回的展示名称可能随规则变化而变化。
	DisplayName *string `json:"display_name,omitempty"`
	// Email 是优先获取email登录名，如果不存在，则返回mobile登录名
	Email *string `json:"email,omitempty"`
	// EnrollmentTime 是入学时间，yyyy-mm-dd格式
	EnrollmentTime *string `json:"enrollment_time,omitempty"`
	// EntLicenseAddress 是企业证照地址的详细地址（用户类型是公司用户才有该字段）【注意】只有is_certified为T的时候才有意义，否则不保证准确性.
	EntLicenseAddress *string `json:"ent_license_address,omitempty"`
	// EntLicenseArea 是企业证照地址所在地区或县（用户类型是公司用户才有该字段）【注意】只有is_certified为T的时候才有意义，否则不保证准确性，目前企业营业执照地址不含县(区)，area为空。
	EntLicenseArea *string `json:"ent_license_area,omitempty"`
	// EntLicenseCity 是企业证照地址所在市（用户类型是公司用户才有该字段）【注意】只有is_certified为T的时候才有意义，否则不保证准确性.
	EntLicenseCity *string `json:"ent_license_city,omitempty"`
	// EntLicenseProvince 是企业证照地址所在省份（用户类型是公司用户才有该字段）【注意】只有is_certified为T的时候才有意义，否则不保证准确性.
	EntLicenseProvince *string `json:"ent_license_province,omitempty"`
	// FirmAgentPersonCertExpiryDate 是企业代理人证件有效期（用户类型是公司类型时才有此字段）。
	FirmAgentPersonCertExpiryDate *string `json:"firm_agent_person_cert_expiry_date,omitempty"`
	// FirmAgentPersonCertNo 是企业代理人证件号码（用户类型是公司类型时才有此字段）。
	FirmAgentPersonCertNo *string `json:"firm_agent_person_cert_no,omitempty"`
	// FirmAgentPersonCertType 是企业代理人证件类型, 返回值参考cert_type字段（用户类型是公司类型时才有此字段）。
	FirmAgentPersonCertType *string `json:"firm_agent_person_cert_type,omitempty"`
	// FirmAgentPersonName 是企业代理人姓名（用户类型是公司类型时才有此字段）。
	FirmAgentPersonName *string `json:"firm_agent_person_name,omitempty"`
	// FirmLegalPersonCertExpiryDate 是企业法人证件有效期（用户类型是公司类型时才有此字段）。
	FirmLegalPersonCertExpiryDate *string `json:"firm_legal_person_cert_expiry_date,omitempty"`
	// FirmLegalPersonCertNo 是法人证件号码（用户类型是公司类型时才有此字段）。
	FirmLegalPersonCertNo *string `json:"firm_legal_person_cert_no,omitempty"`
	// FirmLegalPersonCertType 是企业法人证件类型, 返回值参考cert_type字段（用户类型是公司类型时才有此字段）。
	FirmLegalPersonCertType *string `json:"firm_legal_person_cert_type,omitempty"`
	// FirmLegalPersonName 是企业法人名称（用户类型是公司类型时才有此字段）。
	FirmLegalPersonName *string `json:"firm_legal_person_name,omitempty"`
	// FirmLegalPersonPictures 是企业法人证件图片（用户类型是公司类型时才有此字段）。
	FirmLegalPersonPictures *[]UserPicture `json:"firm_legal_person_pictures,omitempty"`
	// FirmPictures 是企业相关证件图片，包含图片地址（目前需要调用alipay.user.certify.image.fetch转换一下）及类型（用户类型是公司类型时才有此字段）。
	FirmPictures *[]UserPicture `json:"firm_pictures,omitempty"`
	// FirmType 是公司类型，包括（用户类型是公司类型时才有此字段）：		CO(公司)		INST(事业单位),		COMM(社会团体),		NGO(民办非企业组织),		STATEORGAN(党政国家机关)
	FirmType *string `json:"firm_type,omitempty"`
	// Gender 是性别。枚举值如下：	F：女性；	M：男性。
	Gender *string `json:"gender,omitempty"`
	// GraduationTime 是预期毕业时间，不保证准确性，yyyy-mm-dd格式。
	GraduationTime *string `json:"graduation_time,omitempty"`
	// IdentityCardAddress 是身份证地址信息
	IdentityCardAddress *string `json:"identity_card_address,omitempty"`
	// IdentityCardArea 是身份证地区信息
	IdentityCardArea *string `json:"identity_card_area,omitempty"`
	// IdentityCardCity 是身份证市信息
	IdentityCardCity *string `json:"identity_card_city,omitempty"`
	// IdentityCardProvince 是身份证省信息
	IdentityCardProvince *string `json:"identity_card_province,omitempty"`
	// InstOrCorp 是是否是金融机构或特殊单位，"I"表示金融机构, "C"表示无余额账户的单位账户,  "V", 表示虚拟主体(类似机构仅内部户无余额户),  "N"表示非金融机构且非无余额账户的单位账户,"?"表示匿名用户
	InstOrCorp *string `json:"inst_or_corp,omitempty"`
	// IsAdult 是T/F, 实名认证为T，且证件类型为身份证，则输出是否成年字段，否则不输出
	IsAdult *string `json:"is_adult,omitempty"`
	// IsBalanceFrozen 是余额账户是否被冻结。		T--被冻结；F--未冻结
	IsBalanceFrozen *string `json:"is_balance_frozen,omitempty"`
	// IsBlocked 是账户是否被冻结，T/F，T冻结，F未冻结
	IsBlocked *string `json:"is_blocked,omitempty"`
	// IsCertified 是是否通过实名认证。T是通过 F是没有实名认证。
	IsCertified *string `json:"is_certified,omitempty"`
	// IsStudentCertified 是是否是学生
	IsStudentCertified *string `json:"is_student_certified,omitempty"`
	// LicenseExpiryDate 是营业执照有效期，yyyyMMdd或长期，（用户类型是公司类型时才有此字段）。
	LicenseExpiryDate *string `json:"license_expiry_date,omitempty"`
	// LicenseNo 是企业执照号码（用户类型是公司类型时才有此字段）。
	LicenseNo *string `json:"license_no,omitempty"`
	// MemberGrade 是支付宝会员等级
	MemberGrade *string `json:"member_grade,omitempty"`
	// Mobile 是手机号码。
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是用户昵称。	注意：如果没有数据（用户未设置）时不会返回该信息，请做好容错。
	NickName *string `json:"nick_name,omitempty"`
	// OpenId 是支付宝用户的openId。
	OpenId *string `json:"open_id,omitempty"`
	// OrganizationCode 是组织机构代码（用户类型是公司类型时才有此字段）。
	OrganizationCode *string `json:"organization_code,omitempty"`
	// PersonBirthday 是个人用户生日。
	PersonBirthday *string `json:"person_birthday,omitempty"`
	// PersonBirthdayWithoutYear 是生日。不包含具体年份，格式MMdd
	PersonBirthdayWithoutYear *string `json:"person_birthday_without_year,omitempty"`
	// PersonCertExpiryDate 是证件有效期（用户类型是个人的时候才有此字段）。
	PersonCertExpiryDate *string `json:"person_cert_expiry_date,omitempty"`
	// PersonCertIssueDate 是证件起始日期（用户类型是个人的时候才可能有此字段，不保证准确，同时有可能为空）。
	PersonCertIssueDate *string `json:"person_cert_issue_date,omitempty"`
	// PersonPictures 是个人证件图片（用户类型是个人的时候才有此字段）。
	PersonPictures *[]UserPicture `json:"person_pictures,omitempty"`
	// Phone 是电话号码。
	Phone *string `json:"phone,omitempty"`
	// Profession 是职业
	Profession *string `json:"profession,omitempty"`
	// Province 是省份名称。
	Province *string `json:"province,omitempty"`
	// TaobaoId 是淘宝id
	TaobaoId *string `json:"taobao_id,omitempty"`
	// UserId 是支付宝用户的userId。
	UserId *string `json:"user_id,omitempty"`
	// UserName 是若用户是个人用户，则是用户的真实姓名；若是企业用户，则是企业名称。【注意】只有is_certified为T的时候才有意义，否则不保证准确性.
	UserName *string `json:"user_name,omitempty"`
	// UserNation 是身份证民族信息
	UserNation *string `json:"user_nation,omitempty"`
	// UserStatus 是用户状态（Q/T/B/W）。	Q代表快速注册用户	T代表正常用户	B代表被冻结账户	W代表已注册，未激活的账户
	UserStatus *string `json:"user_status,omitempty"`
	// UserType 是用户类型（1/2） 	1代表公司账户2代表个人账户
	UserType *string `json:"user_type,omitempty"`
	// Zip 是邮政编码。
	Zip *string `json:"zip,omitempty"`
}
