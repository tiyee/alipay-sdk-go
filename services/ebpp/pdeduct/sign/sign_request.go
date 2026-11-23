package sign

// PdeductSignAddRequest 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignAddRequest struct {
	// AgentChannel 是机构签约代扣来源渠道 	PUBLICPLATFORM：服务窗
	AgentChannel *string `json:"agent_channel,omitempty"`
	// AgentCode 是从服务窗发起则为publicId的值
	AgentCode *string `json:"agent_code,omitempty"`
	// BillKey 是户号，机构针对于每户的水、电都会有唯一的标识户号
	BillKey *string `json:"bill_key,omitempty"`
	// BizType 是业务类型。枚举支持：	*JF：公共事业缴纳。	*WUYE：物业缴费。	*HK：信用卡还款。	*TX：通讯缴费。	*IND：保险缴费。
	BizType *string `json:"biz_type,omitempty"`
	// ChargeInst 是支付宝缴费系统中的出账机构ID
	ChargeInst *string `json:"charge_inst,omitempty"`
	// DeductProdCode 是代扣产品码。 缴费业务场景中，传入INST_DIRECT_DEDUCT。保险业务场景中, 传入INSURANCE_MERCHANT_DEDUCT。
	DeductProdCode *string `json:"deduct_prod_code,omitempty"`
	// DeductType 是签约类型可为空
	DeductType *string `json:"deduct_type,omitempty"`
	// ExtUserInfo 是外部用户实名认证相关信息, 用于做签约时的实名校验。  	注：  	 	name: 姓名 	cert_type: 身份证：IDENTITY_CARD、护照：PASSPORT、军官证：OFFICER_CARD、士兵证：SOLDIER_CARD、户口本：HOKOU等 	cert_no: 证件号码 	need_check_info取值 T/F，只有为T时才做强制校验。 	mobile:手机号，目前暂不使用此字段做校验 	min_age: 允许的最小买家年龄,min_age为整数，必须大于等于0.
	ExtUserInfo *string `json:"ext_user_info,omitempty"`
	// ExtendField 是扩展字段
	ExtendField *string `json:"extend_field,omitempty"`
	// NotifyConfig 是通知方式设置，可为空
	NotifyConfig *string `json:"notify_config,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// OutAgreementId 是外部产生的协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
	// OwnerName 是户名，户主真实姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// PayConfig 是支付工具设置，目前可为空
	PayConfig *string `json:"pay_config,omitempty"`
	// PayPasswordToken 是支付密码鉴权token，需要用户首先处于登陆态，然后访问https://ebppprod.alipay.com/deduct/enterMobileicPayPassword.htm?cb=自己指定的回跳连接地址，访问页面后会进到独立密码校验页，用户输入密码校验成功后，会生成token回调到指定的回跳地址，如果设置cb=www.alipay.com则最后回调的内容是www.alipay.com?token=312314ADFDSFAS，然后签约时直接取得地址后token的值即可。
	PayPasswordToken *string `json:"pay_password_token,omitempty"`
	// Pid 是商户id，商户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	Pid *string `json:"pid,omitempty"`
	// SignExpireDate 是签约到期时间。空表示无限期，一期固定传空。
	SignExpireDate *string `json:"sign_expire_date,omitempty"`
	// SpecifiedSortChannels 是用户设置的优先扣款渠道
	SpecifiedSortChannels *[]SpecifiedSortChannelParam `json:"specified_sort_channels,omitempty"`
	// SubBizType 是业务子类型。业务子类型是业务类型的下一级概念。枚举支持：	*WATER：缴水费。	*ELECTRIC：缴电费。	*GAS：缴燃气费。	*COMMUN：缴固话宽带费。	*CATV：缴有线电视费。	*TRAFFIC：缴交通罚款。	*WUYE：缴物业费。	*RC：定期还车贷代扣。	*RH：定期还房贷代扣。	*RR：定期还房租代扣。	*RN：定期还网贷代扣。	*CZ：手机充值代扣。	*CAR：车险。	*LIFE：寿险。	*HEALTH：健康险。	例如：WATER表示JF下面的水费。
	SubBizType *string `json:"sub_biz_type,omitempty"`
	// UserId 是用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitempty"`
}

// PdeductSignCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignCancelRequest struct {
	// AgentChannel 是此值只是供代扣中心做最后的渠道统计用，并不做值是什么的强校验，只要不为空就可以
	AgentChannel *string `json:"agent_channel,omitempty"`
	// AgentCode 是标识发起方的ID，从服务窗发起则为appId的值，appId即开放平台分配给接入ISV的id，此处也可以随便真其它值，只要能标识唯一即可
	AgentCode *string `json:"agent_code,omitempty"`
	// AgreementId 是支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// PayPasswordToken 是需要用户首先处于登陆态，然后访问https://ebppprod.alipay.com/deduct/enterMobileicPayPassword.htm?cb=自己指定的回跳连接地址,访问页面后会进到独立密码校验页，用户输入密码校验成功后，会生成token回调到指定的回跳地址，如果设置cb=www.baidu.com则最后回调的内容是www.baidu.com?token=312314ADFDSFAS,然后签约时直接取得地址后token的值即可
	PayPasswordToken *string `json:"pay_password_token,omitempty"`
	// UserId 是用户ID
	UserId *string `json:"user_id,omitempty"`
}

// PdeductSignValidateRequest 是支付宝官方 V3 OpenAPI 模型。
type PdeductSignValidateRequest struct {
	// AgentChannel 是机构签约代扣来源渠道 	PUBLICPLATFORM：服务窗
	AgentChannel *string `json:"agent_channel,omitempty"`
	// AgentCode 是从服务窗发起则为该服务窗的appid的值
	AgentCode *string `json:"agent_code,omitempty"`
	// BillKey 是户号，机构针对于每户的水、电都会有唯一的标识户号
	BillKey *string `json:"bill_key,omitempty"`
	// BizType 是业务类型。 	JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用 	WUYE：缴物业费 	HK：信用卡还款 	TX：手机充值
	BizType *string `json:"biz_type,omitempty"`
	// ChargeInst 是支付宝缴费系统中的出账机构ID
	ChargeInst *string `json:"charge_inst,omitempty"`
	// DeductType 是签约类型,可为空。目前类型有INST_DIRECT_DEDUCT
	DeductType *string `json:"deduct_type,omitempty"`
	// ExtendField 是扩展字段
	ExtendField *string `json:"extend_field,omitempty"`
	// NotifyConfig 是通知方式设置，本期预留字段，固定传空。
	NotifyConfig *string `json:"notify_config,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// OutAgreementId 是外部机构签约的协议id
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
	// OwnerName 是户名，户主真实姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// PayConfig 是支付工具设置，目前可为空。类型有：BALANCE;CARTOON;BIGAMOUNT_CREDIT_CARTOON;DEBIT_EXPRESS;OPTIMIZED_MOTO;PCREDIT_PAY;MONEY_FUND
	PayConfig *string `json:"pay_config,omitempty"`
	// Pid 是商户签约支付宝账号的userid，2088开头16位长度的字符串，在支付宝系统中唯一
	Pid *string `json:"pid,omitempty"`
	// SignExpireDate 是签约到期时间。空表示无限期，一期固定传空。
	SignExpireDate *string `json:"sign_expire_date,omitempty"`
	// SubBizType 是业务子类型。 	WATER：缴水费 	ELECTRIC：缴电费 	GAS：缴燃气费 	COMMUN：缴固话宽带 	CATV：缴有线电视费 	TRAFFIC：缴交通罚款 	WUYE：缴物业费 	HK：信用卡还款 	CZ：手机充值
	SubBizType *string `json:"sub_biz_type,omitempty"`
	// UserId 是用户支付宝账号id，2088开头16位长度的字符串，在支付宝系统中唯一
	UserId *string `json:"user_id,omitempty"`
}

// PdeductSignQueryRequest is the query request model for alipay.ebpp.pdeduct.sign.query.
type PdeductSignQueryRequest struct {
	// UserId 是用户ID
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// AgreementId 是支付宝代扣协议Id。若协议id不传递，则需要保证业务类型、子业务类型、出账机构、户号必传
	AgreementId *string `json:"agreement_id,omitempty" query:"agreement_id,omitempty"`
	// BizType 是业务类型。 	JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用 	WUYE：缴物业费 	HK：信用卡还款 	TX：手机充值
	BizType *string `json:"biz_type,omitempty" query:"biz_type,omitempty"`
	// SubBizType 是业务子类型。 	WATER：缴水费 	ELECTRIC：缴电费 	GAS：缴燃气费 	COMMUN：缴固话宽带 	CATV：缴有线电视费 	TRAFFIC：缴交通罚款 	WUYE：缴物业费 	HK：信用卡还款 	CZ：手机充值
	SubBizType *string `json:"sub_biz_type,omitempty" query:"sub_biz_type,omitempty"`
	// ChargeInst 是支付宝缴费系统中的出账机构ID
	ChargeInst *string `json:"charge_inst,omitempty" query:"charge_inst,omitempty"`
	// BillKey 是户号，机构针对于每户的水、电都会有唯一的标识户号
	BillKey *string `json:"bill_key,omitempty" query:"bill_key,omitempty"`
}
