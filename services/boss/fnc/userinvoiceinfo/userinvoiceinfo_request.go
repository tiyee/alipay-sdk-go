package userinvoiceinfo

// BossFncUserInvoiceInfoCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncUserInvoiceInfoCreateRequest struct {
	// AcceptElectronic 是是否接受电子票 true:是，false:否
	AcceptElectronic *bool `json:"accept_electronic,omitempty"`
	// Address 是公司注册地址，公司营业执照上登记的住址，一般纳税人必须填写，小规模纳税人无需填写
	Address *string `json:"address,omitempty"`
	// Addressing 是开票资料pid寻址方式 true:pid寻址mid,优先以mid维度保存开票资料,若无mid则以pid维度保存开票资料 false:不寻址mid,以pid维度保存开票资料;非主站商户体系都是以pid（ipRoleId）维度保存开票资料，该值直接填为false
	Addressing *bool `json:"addressing,omitempty"`
	// Auto 是是否自动申请开票 true:是，false:否
	Auto *bool `json:"auto,omitempty"`
	// BankAccount 是银行账户，公司银行账号 	一般纳税人必须填写，小规模纳税人无需填写
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行，办理银行开户手续的营业网点 	一般纳税人必须填写，小规模纳税人无需填写
	BankName *string `json:"bank_name,omitempty"`
	// BusinessLicenceUrl 是营业执照地址，营业执照上传oss上的地址
	BusinessLicenceUrl *string `json:"business_licence_url,omitempty"`
	// Hold 是是否暂停开票，商户添加的开票资料就是不暂停开票，如果后续要暂停开票，只能去结算中台处理。
	Hold *bool `json:"hold,omitempty"`
	// IpRoleId 是商户的pid（ipRoleId）
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// OpenAccountPermitUrl 是银行开户许可证附件，银行开户许可证上传oss上的地址
	OpenAccountPermitUrl *string `json:"open_account_permit_url,omitempty"`
	// Operator 是当前操作人
	Operator *string `json:"operator,omitempty"`
	// OperatorType 是当前操作人类型,01:商户 02:管理员(小二)
	OperatorType *string `json:"operator_type,omitempty"`
	// OtherQualificationUrl 是其它资质证明地址，其他资质证明附件上传oss上的地址
	OtherQualificationUrl *string `json:"other_qualification_url,omitempty"`
	// Ou 是外部只允许添加商户开票资料，该值填写为false
	Ou *bool `json:"ou,omitempty"`
	// TaxNo 是纳税人识别号，税务登记证上的号码。一般纳税人必须填写，小规模纳税人如果是商户也必须填写，个人无需填写
	TaxNo *string `json:"tax_no,omitempty"`
	// TaxPayerQualiValid 是纳税人资格开始时间 （格式：时间戳） 	一般纳税人必须填写，小规模纳税人无需填写
	TaxPayerQualiValid *string `json:"tax_payer_quali_valid,omitempty"`
	// TaxPayerQualification 是纳税人资格种类，01：一般纳税人；02：小规模纳税人；03：国际商户
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty"`
	// TaxQualificationUrl 是一般纳税人资格证书地址，一般纳税人资格证书上传oss上的地址
	TaxQualificationUrl *string `json:"tax_qualification_url,omitempty"`
	// TaxRegCertUrl 是税务登记证地址，税务登记证上传oss上的地址
	TaxRegCertUrl *string `json:"tax_reg_cert_url,omitempty"`
	// Telephone 是公司注册电话（手机号和座机均可） 	一般纳税人必须填写，小规模纳税人无需填写
	Telephone *string `json:"telephone,omitempty"`
	// Title 是发票抬头，票面信息上的抬头信息
	Title *string `json:"title,omitempty"`
	// UserMailInfoOrderList 是收件人列表,若不修改此项可以不填 	选择非电子票时，邮寄信息必填，且保证邮寄信息的ipRoleId和开票资料的ipRoleId相等
	UserMailInfoOrderList *[]UserMailInfoOrder `json:"user_mail_info_order_list,omitempty"`
}

// BossFncUserInvoiceInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncUserInvoiceInfoModifyRequest struct {
	// AcceptElectronic 是是否接受电子票 true:是，false:否
	AcceptElectronic *bool `json:"accept_electronic,omitempty"`
	// Address 是公司注册地址，公司营业执照上登记的住址，一般纳税人必须填写，小规模纳税人无需填写
	Address *string `json:"address,omitempty"`
	// Auto 是是否自动申请开票 true:是；false:否
	Auto *bool `json:"auto,omitempty"`
	// BankAccount 是银行账户，公司银行账号 	一般纳税人必须填写，小规模纳税人无需填写
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行，办理银行开户手续的营业网点 	一般纳税人必须填写，小规模纳税人无需填写
	BankName *string `json:"bank_name,omitempty"`
	// BusinessLicenceUrl 是营业执照地址，营业执照上传oss上的地址
	BusinessLicenceUrl *string `json:"business_licence_url,omitempty"`
	// Hold 是是否暂停开票，商户修改无效，如果要暂停开票，只能去结算中台处理。
	Hold *bool `json:"hold,omitempty"`
	// Id 是开票资料ID
	Id *string `json:"id,omitempty"`
	// OpenAccountPermitUrl 是银行开户许可证附件，银行开户许可证上传oss上的地址
	OpenAccountPermitUrl *string `json:"open_account_permit_url,omitempty"`
	// Operator 是当前操作人
	Operator *string `json:"operator,omitempty"`
	// OperatorType 是当前操作人类型,01:商户；02:管理员(小二)
	OperatorType *string `json:"operator_type,omitempty"`
	// OtherQualificationUrl 是其它资质证明地址，其他资质证明附件上传oss上的地址
	OtherQualificationUrl *string `json:"other_qualification_url,omitempty"`
	// TaxNo 是纳税人识别号，税务登记证上的号码。一般纳税人必须填写，小规模纳税人如果是商户也必须填写，个人无需填写
	TaxNo *string `json:"tax_no,omitempty"`
	// TaxPayerQualiValid 是纳税人资格开始时间 （格式：时间戳） 	一般纳税人必须填写，小规模纳税人无需填写
	TaxPayerQualiValid *string `json:"tax_payer_quali_valid,omitempty"`
	// TaxPayerQualification 是纳税人资格种类，01：一般纳税人；02：小规模纳税人；03：国际商户
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty"`
	// TaxQualificationUrl 是一般纳税人资格证书地址，一般纳税人资格证书上传oss上的地址
	TaxQualificationUrl *string `json:"tax_qualification_url,omitempty"`
	// TaxRegCertUrl 是税务登记证地址，税务登记证上传oss上的地址
	TaxRegCertUrl *string `json:"tax_reg_cert_url,omitempty"`
	// Telephone 是公司注册电话 	一般纳税人必须填写，小规模纳税人无需填写
	Telephone *string `json:"telephone,omitempty"`
	// Title 是发票抬头，票面信息上的抬头信息
	Title *string `json:"title,omitempty"`
	// UserMailInfoOrderList 是收件人列表,若不修改此项可以不填 	选择非电子票时，邮寄信息必填，且保证邮寄信息的ipRoleId和开票资料的ipRoleId相等
	UserMailInfoOrderList *[]UserMailInfoOrder `json:"user_mail_info_order_list,omitempty"`
}

// BossFncUserInvoiceInfoQueryRequest is the query request model for alipay.boss.fnc.userinvoiceinfo.query.
type BossFncUserInvoiceInfoQueryRequest struct {
	// Pid 是开票pid/mid/ou，唯一标识商户信息/集团用户信息的ID
	Pid *string `json:"pid,omitempty" query:"pid,omitempty"`
}
