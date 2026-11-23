package userinvoiceinfo

import (
	"encoding/json"
)

// BossFncUserInvoiceInfoCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncUserInvoiceInfoCreateResponse struct {
	// ResultSet 是开票资料id，唯一性ID
	ResultSet *string `json:"result_set,omitempty"`
}

// BossFncUserInvoiceInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncUserInvoiceInfoModifyResponse = json.RawMessage

// BossFncUserInvoiceInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncUserInvoiceInfoQueryResponse struct {
	ResultSet *UserInvoiceInfoOpenAPIResponse `json:"result_set,omitempty"`
}

// UserInvoiceInfoOpenAPIResponse 是支付宝官方 V3 OpenAPI 模型。
type UserInvoiceInfoOpenAPIResponse struct {
	// AcceptElectronic 是(AR开票使用)是否接受电子票 true:是，false:否
	AcceptElectronic *bool `json:"accept_electronic,omitempty"`
	// Address 是公司注册地址
	Address *string `json:"address,omitempty"`
	// Auto 是(AR开票使用)是否自动申请开票 true:是，false:否
	Auto *bool `json:"auto,omitempty"`
	// BankAccount 是银行账户
	BankAccount *string `json:"bank_account,omitempty"`
	// BankName 是开户行
	BankName *string `json:"bank_name,omitempty"`
	// BusinessLicenceUrl 是营业执照地址
	BusinessLicenceUrl *string `json:"business_licence_url,omitempty"`
	// Creator 是资料创建人
	Creator *string `json:"creator,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Hold 是(AR开票使用)是否暂停开票 	true:暂停开票，false:可开票
	Hold *bool `json:"hold,omitempty"`
	// Id 是开票资料ID
	Id *string `json:"id,omitempty"`
	// IpId 是商户PID, 	type=01时表示OU用户,填充的是InstId机构代码,例如Z50, 	type=02时表示商户MID, 	type=03时表示商户PID
	IpId *string `json:"ip_id,omitempty"`
	// LastModifier 是最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`
	// OpenAccountPermitUrl 是银行开户许可证地址
	OpenAccountPermitUrl *string `json:"open_account_permit_url,omitempty"`
	// OtherQualificationUrl 是其它资质证明地址
	OtherQualificationUrl *string `json:"other_qualification_url,omitempty"`
	// Status 是资料状态，01：待审，02：已审(有效)， 03：已作废
	Status *string `json:"status,omitempty"`
	// TaskDate 是上一次任务执行时间(针对自动开票场景),格式:YYYYMMDD
	TaskDate *string `json:"task_date,omitempty"`
	// TaxNo 是纳税人识别号（购方税号）
	TaxNo *string `json:"tax_no,omitempty"`
	// TaxPayerQualification 是纳税人开票资格种类 	01：一般纳税人，02：小规模纳税人
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty"`
	// TaxQualificationUrl 是一般纳税人资格证书地址
	TaxQualificationUrl *string `json:"tax_qualification_url,omitempty"`
	// TaxRegCertUrl 是税务登记证地址
	TaxRegCertUrl *string `json:"tax_reg_cert_url,omitempty"`
	// TaxpayerQualiValid 是纳税人资格开始时间 （时间戳）
	TaxpayerQualiValid *string `json:"taxpayer_quali_valid,omitempty"`
	// Telephone 是公司注册电话
	Telephone *string `json:"telephone,omitempty"`
	// Title 是发票抬头
	Title *string `json:"title,omitempty"`
	// TntInstId 是租户ID
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// Type 是开票资料用户类型  	01：OU，02：商户，03：C用户
	Type *string `json:"type,omitempty"`
	// UserMailInfoList 是收件人地址列表
	UserMailInfoList *[]UserMailInfoVO `json:"user_mail_info_list,omitempty"`
}
