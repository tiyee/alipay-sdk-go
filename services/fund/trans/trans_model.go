package trans

const (
	// TransStatusSuccess 表示转账成功。
	TransStatusSuccess TransStatus = "SUCCESS"
	// TransStatusFail 表示转账失败。
	TransStatusFail TransStatus = "FAIL"
	// TransStatusInit 表示等待处理。
	TransStatusInit TransStatus = "INIT"
	// TransStatusDealing 表示处理中。
	TransStatusDealing TransStatus = "DEALING"
	// TransStatusRefund 表示退票。
	TransStatusRefund TransStatus = "REFUND"
	// TransStatusWaitPay 表示等待支付。
	TransStatusWaitPay TransStatus = "WAIT_PAY"
	// TransStatusClosed 表示订单超时关闭。
	TransStatusClosed TransStatus = "CLOSED"
)

const (
	// PayeeTypeUserID 表示支付宝唯一用户号（2088 开头 16 位数字）。
	PayeeTypeUserID PayeeAccountType = "ALIPAY_USERID"
	// PayeeTypeLogonID 表示支付宝登录号（支持邮箱和手机号）。
	PayeeTypeLogonID PayeeAccountType = "ALIPAY_LOGONID"
	// PayeeTypeOpenID 表示支付宝 openid。
	PayeeTypeOpenID PayeeAccountType = "ALIPAY_OPENID"
)

// TransStatus 是转账单据状态，取值依据官方 V3 文档；
// 未知状态原样保留，调用方按默认分支处理。
type TransStatus string

// TransferSceneReportInfo 是转账场景信息项。
type TransferSceneReportInfo struct {
	// InfoType 是转账场景信息类型，取值见官方转账场景信息字段说明。
	InfoType *string `json:"info_type,omitempty"`
	// InfoContent 是转账场景信息描述，按官方字段说明传参。
	InfoContent *string `json:"info_content,omitempty"`
}

// Participant 是转账参与方（收款方/付款方）信息。
type Participant struct {
	// Identity 是参与方的唯一标识。
	Identity *string `json:"identity,omitempty"`
	// IdentityType 是参与方的标识类型：ALIPAY_USER_ID（会员 ID）、
	// ALIPAY_LOGON_ID（登录号）、ALIPAY_OPEN_ID（openid）、
	// EXPRESS_DC_STFA（对公快捷）等。
	IdentityType *string `json:"identity_type,omitempty"`
	// Name 是参与方真实姓名；非空时校验收款账号姓名一致性；
	// identity_type=ALIPAY_LOGON_ID 时必填。
	Name *string `json:"name,omitempty"`
	// CertType 是参与方的证件类型。
	CertType *string `json:"cert_type,omitempty"`
	// CertNo 是参与方的证件号。
	CertNo *string `json:"cert_no,omitempty"`
	// BankcardExtInfo 是银行卡扩展信息（转账到银行卡场景）。
	BankcardExtInfo *BankcardExtInfo `json:"bankcard_ext_info,omitempty"`
	// ExtInfo 表示描述参与方信息的扩展属性，使用前请与支付宝工程师确认。
	ExtInfo *string `json:"ext_info,omitempty"`
	// MerchantUserInfo 表示商户侧的参与方信息。
	MerchantUserInfo *string `json:"merchant_user_info,omitempty"`
}

// BankcardExtInfo 是收款银行卡扩展信息。
type BankcardExtInfo struct {
	// BankCode 是银行支行联行号。
	BankCode *string `json:"bank_code,omitempty"`
	// InstName 是机构名称。
	InstName *string `json:"inst_name,omitempty"`
	// InstBranchName 是收款银行所属支行。
	InstBranchName *string `json:"inst_branch_name,omitempty"`
	// InstProvince 是收款银行所在省。
	InstProvince *string `json:"inst_province,omitempty"`
	// InstCity 是收款银行所在市。
	InstCity *string `json:"inst_city,omitempty"`
	// AccountType 是收款账户类型：1 对公，2 对私。
	AccountType *string `json:"account_type,omitempty"`
}

// PayeeAccountType 是单笔转账到支付宝账户的收款账户类型，
// 取值依据官方 V3 文档。
type PayeeAccountType string

// AuthInfo 是协议授权信息。
type AuthInfo struct {
	// AuthCode 是授权号，协议授权或 CIF 协议对应的授权号。
	AuthCode *string `json:"auth_code,omitempty"`
	// AuthType 是授权类型。
	AuthType *string `json:"auth_type,omitempty"`
}

// IsKnown 报告账户类型是否属于文档已知值。
func (t PayeeAccountType) IsKnown() bool {
	switch t {
	case PayeeTypeUserID, PayeeTypeLogonID, PayeeTypeOpenID:
		return true
	default:
		return false
	}
}

// MultipleCurrencyDetail 是多币种转账明细。
type MultipleCurrencyDetail struct {
	// ExtInfo 表示扩展字段。
	ExtInfo *string `json:"ext_info,omitempty"`
	// PaymentAmount 表示支付金额，单位为该币种基础计价单位，如人民币为：元。
	PaymentAmount *string `json:"payment_amount,omitempty"`
	// PaymentCurrency 表示支付币种（元）。
	PaymentCurrency *string `json:"payment_currency,omitempty"`
	// SettlementAmount 表示清算金额，单位为该币种基础计价单位，如人民币为：元。
	SettlementAmount *string `json:"settlement_amount,omitempty"`
	// SettlementCurrency 表示清算币种（元）。
	SettlementCurrency *string `json:"settlement_currency,omitempty"`
	// TransAmount 表示转账金额，单位为该币种基础计价单位，如人民币为：元。
	TransAmount *string `json:"trans_amount,omitempty"`
	// TransCurrency 表示转账币种（元）。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// SignData 是签名原始信息。
type SignData struct {
	// OriAppID 表示签名商户开放平台应用APPID。
	OriAppID *string `json:"ori_app_id,omitempty"`
	// OriCharSet 表示编码类型（大小写敏感）。
	OriCharSet *string `json:"ori_char_set,omitempty"`
	// OriOutBizNo 表示ori_out_biz_no_001。
	OriOutBizNo *string `json:"ori_out_biz_no,omitempty"`
	// OriSign 表示签名密文。
	OriSign *string `json:"ori_sign,omitempty"`
	// OriSignType 表示签名算法（大小写敏感）。
	OriSignType *string `json:"ori_sign_type,omitempty"`
	// PartnerID 表示商户支付宝账号id。
	PartnerID *string `json:"partner_id,omitempty"`
}
