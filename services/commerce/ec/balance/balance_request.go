package balance

// EcBalancePeriodModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcBalancePeriodModifyRequest struct {
	// AccountId 是共同账户id - 适用于在企业码小程序创建的共同账户，和agreement_no搭配使用
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BillDay 是月账单账期
	BillDay *int64 `json:"bill_day,omitempty"`
	// EnterpriseId 是企业ID - 适用于在企业码PC端创建的企业账号
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// EcBalanceDownloadURLQueryRequest is the query request model for alipay.commerce.ec.balance.downloadurl.query.
type EcBalanceDownloadURLQueryRequest struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// BillType 是账单类型
	BillType *string `json:"bill_type,omitempty" query:"bill_type,omitempty"`
	// BillDate 是账单时间
	BillDate *string `json:"bill_date,omitempty" query:"bill_date,omitempty"`
}
