package agreement

// FundAgreementQuotaModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAgreementQuotaModifyResponse struct {
	// QuotaModifyDetailList 是额度变更结果集合。
	QuotaModifyDetailList *[]QuotaModifyDetail `json:"quota_modify_detail_list,omitempty"`
}

// FundAgreementQuotaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAgreementQuotaQueryResponse struct {
	// QuotaQueryResponseList 是单个协议的额度查询结果。
	QuotaQueryResponseList *[]QuotaQueryResponse `json:"quota_query_response_list,omitempty"`
}

// QuotaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type QuotaQueryResponse struct {
	// AgreementNo 是协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ErrorCode 是单个协议查询结果错误码。
	ErrorCode *string `json:"error_code,omitempty"`
	// FailReason 是单个协议响应错误描述。
	FailReason   *string             `json:"fail_reason,omitempty"`
	QuotaDetails *AccountQuotaDetail `json:"quota_details,omitempty"`
	// Success 是单个协议查询结果是否成功。
	Success *string `json:"success,omitempty"`
}
