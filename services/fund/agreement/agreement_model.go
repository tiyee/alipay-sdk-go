package agreement

// AccountQuotaDetail 是支付宝官方 V3 OpenAPI 模型。
type AccountQuotaDetail struct {
	// QuotaAmount 是周期内的额度上限。
	QuotaAmount *int64 `json:"quota_amount,omitempty"`
	// QuotaDimension 是额度周期，有限枚举。
	QuotaDimension *string `json:"quota_dimension,omitempty"`
	// Role 是角色。
	Role *string `json:"role,omitempty"`
}

// AgreementQuotaModifyList 是支付宝官方 V3 OpenAPI 模型。
type AgreementQuotaModifyList struct {
	// AgreementNo 是授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// QuotaDetails 是设置额度明细。
	QuotaDetails *[]ModifyQuotaDetails `json:"quota_details,omitempty"`
}

// ModifyQuotaDetails 是支付宝官方 V3 OpenAPI 模型。
type ModifyQuotaDetails struct {
	// QuotaAmount 是具体额度，单位：元。如金额为空，表示删除已有的额度设置，无已有额度设置则忽略。。
	QuotaAmount *string `json:"quota_amount,omitempty"`
	// QuotaDimension 是额度维度 MONTH/DAY/SINGLE 分别代表月、日、单次。
	QuotaDimension *string `json:"quota_dimension,omitempty"`
	// Role 是PAYER/PAYEE 额度管控的角色，收or付款方，目前只支持付款方。
	Role *string `json:"role,omitempty"`
}

// QuotaModifyDetail 是支付宝官方 V3 OpenAPI 模型。
type QuotaModifyDetail struct {
	// AgreementNo 是授权协议号，和入参的协议号对应。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ErrorCode 是明细处理错误码，当success为false时有值。
	ErrorCode *string `json:"error_code,omitempty"`
	// FailReason 是错误描述，当success为false时有值。
	FailReason *string `json:"fail_reason,omitempty"`
	// Success 是当前协议下的明细处理结果。
	Success *bool `json:"success,omitempty"`
}
