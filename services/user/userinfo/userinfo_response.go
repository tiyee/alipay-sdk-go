package userinfo

// UserUserInfoShareResponse 是支付宝官方 V3 OpenAPI 模型。
type UserUserInfoShareResponse struct {
	// Address 是详细地址。
	Address *string `json:"address,omitempty"`
	// AddressCode 是区域编码，暂时不返回值
	AddressCode *string `json:"address_code,omitempty"`
	// AlipayUserId 是支付宝用户ID
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// Area 是区县名称。
	Area *string `json:"area,omitempty"`
	// Avatar 是用户头像
	Avatar *string `json:"avatar,omitempty"`
	// BalanceFreezeType 是【注意】当is_balance_frozen为“F”时，改字段不会返回.		CTU ---- CTU冻结，允许用户开启		ALIBABA ---- ALIBABA冻结，允许用户开启		SERVER ---- 后台冻结，允许用户开启		USER ----  用户冻结 		CTU_N---- CTU冻结，不允许用户开启		ALIBABA_N ---- ALIBABA冻结，不允许用户开启		SERVER_N ---- 后台冻结，不允许用户开启		UNKNOWN ---- 降级、或查询超时
	BalanceFreezeType *string `json:"balance_freeze_type,omitempty"`
	// Birthday 是用户生日
	Birthday *string `json:"birthday,omitempty"`
	// CertNo 是证件号码
	CertNo *string `json:"cert_no,omitempty"`
	// CertTypeValue 是0:身份证		1:护照		2:军官证		3:士兵证		4:回乡证		5:临时身份证		6:户口簿		7:警官证		8:台胞证		9:营业执照		10其它证件
	CertTypeValue *string `json:"cert_type_value,omitempty"`
	// City 是市名称。
	City *string `json:"city,omitempty"`
	// DefaultDeliverAddress 是是否默认收货地址，暂时不返回值
	DefaultDeliverAddress *string `json:"default_deliver_address,omitempty"`
	// DeliverAddressList 是收货人地址列表
	DeliverAddressList *[]DeliverAddress `json:"deliver_address_list,omitempty"`
	// DeliverArea 是收货人所在区县
	DeliverArea *string `json:"deliver_area,omitempty"`
	// DeliverCity 是收货人所在城市
	DeliverCity *string `json:"deliver_city,omitempty"`
	// DeliverFullname 是收货人全称
	DeliverFullname *string `json:"deliver_fullname,omitempty"`
	// DeliverMobile 是收货地址的联系人移动电话。
	DeliverMobile *string `json:"deliver_mobile,omitempty"`
	// DeliverPhone 是收货地址的联系人固定电话。
	DeliverPhone *string `json:"deliver_phone,omitempty"`
	// DeliverProvince 是收货人所在省份
	DeliverProvince *string `json:"deliver_province,omitempty"`
	// Email 是用户支付宝账号绑定的邮箱地址
	Email *string `json:"email,omitempty"`
	// FamilyName 是姓氏，取的是realName中的首个字符，对非中文、中文复姓支持较差。
	FamilyName *string `json:"family_name,omitempty"`
	// FirmName 是公司名称（用户类型是公司类型时公司名称才有此字段）。
	FirmName *string `json:"firm_name,omitempty"`
	// Gender 是性别（F：女性；M：男性）
	Gender *string `json:"gender,omitempty"`
	// IsBalanceFrozen 是T--被冻结；F--未冻结
	IsBalanceFrozen *string `json:"is_balance_frozen,omitempty"`
	// IsBankAuth 是T为是银行卡认证，F为非银行卡认证。
	IsBankAuth *string `json:"is_bank_auth,omitempty"`
	// IsCertified 是是否通过实名认证。T是通过 F是没有实名认证
	IsCertified *string `json:"is_certified,omitempty"`
	// IsCertifyGradeA 是T：表示A类实名认证；F：表示非A类实名认证
	IsCertifyGradeA *string `json:"is_certify_grade_a,omitempty"`
	// IsIdAuth 是T为是身份证认证，F为非身份证认证。
	IsIdAuth *string `json:"is_id_auth,omitempty"`
	// IsLicenceAuth 是T为通过营业执照认证，F为没有通过
	IsLicenceAuth *string `json:"is_licence_auth,omitempty"`
	// IsMobileAuth 是T为是手机认证，F为非手机认证。
	IsMobileAuth *string `json:"is_mobile_auth,omitempty"`
	// IsStudentCertified 是是否是学生
	IsStudentCertified *string `json:"is_student_certified,omitempty"`
	// Mobile 是手机号码。
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是用户昵称
	NickName *string `json:"nick_name,omitempty"`
	// Phone 是电话号码。
	Phone *string `json:"phone,omitempty"`
	// Province 是省份名称。
	Province *string `json:"province,omitempty"`
	// RealName 是用户的真实姓名。
	RealName *string `json:"real_name,omitempty"`
	// ReducedBirthday 是生日的月和日，MMdd格式
	ReducedBirthday *string `json:"reduced_birthday,omitempty"`
	// UserId 是用户的userId
	UserId *string `json:"user_id,omitempty"`
	// UserStatus 是用户状态（Q/T/B/W）。 	Q代表快速注册用户 	T代表已认证用户 	B代表被冻结账户 	W代表已注册，未激活的账户
	UserStatus *string `json:"user_status,omitempty"`
	// UserTypeValue 是用户类型（1/2）		1代表公司账户2代表个人账户
	UserTypeValue *string `json:"user_type_value,omitempty"`
	// Zip 是邮政编码。
	Zip *string `json:"zip,omitempty"`
}
