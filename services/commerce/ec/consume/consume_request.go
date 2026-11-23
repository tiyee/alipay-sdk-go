package consume

// EcConsumeDetailQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type EcConsumeDetailQueryRequest struct {
	// AccountId 是共同账户ID(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AntShopId 是蚂蚁门店ID，不可单独用于查询条件，需要配合企业ID和交易号使用
	AntShopId *string `json:"ant_shop_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// PayNo 是支付宝账单号
	PayNo *string `json:"pay_no,omitempty"`
	// QueryOptions 是查询选项
	QueryOptions *[]string `json:"query_options,omitempty"`
}
