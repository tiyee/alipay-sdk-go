package appkeyword

// SearchAppKeywordApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordApplyRequest struct {
	// ConfigId 是关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取。
	ConfigId *string `json:"config_id,omitempty"`
	// KeyWords 是提报的关键词。
	KeyWords *[]string `json:"key_words,omitempty"`
	// TargetAppid 是小程序id。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchAppKeywordDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchAppKeywordDeleteRequest struct {
	// ConfigId 是关键词配置ID。
	ConfigId *string `json:"config_id,omitempty"`
	// KeyWords 是需要删除的生效关键词。
	KeyWords *[]string `json:"key_words,omitempty"`
	// TargetAppid 是小程序id。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchAppKeywordBatchQueryRequest is the request type for alipay.open.search.appkeyword.batchquery.
type SearchAppKeywordBatchQueryRequest struct {
	// TargetAppid 是小程序id
	TargetAppid *string `json:"target_appid,omitempty" query:"target_appid,omitempty"`
	// Status 是关键词的状态，枚举值，不传该字段表示查询所有状态
	Status *string `json:"status,omitempty" query:"status,omitempty"`
	// PageNumber 是当前页
	PageNumber *string `json:"page_number,omitempty" query:"page_number,omitempty"`
	// PageSize 是每页显示条数
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}

// SearchAppKeywordQueryStatusRequest is the request type for alipay.open.search.appkeyword.querystatus.
type SearchAppKeywordQueryStatusRequest struct {
	// TargetAppid 是小程序id
	TargetAppid *string `json:"target_appid,omitempty" query:"target_appid,omitempty"`
	// ApplyNo 是审核工单id，提报接口的返回值
	ApplyNo *string `json:"apply_no,omitempty" query:"apply_no,omitempty"`
}

// SearchAppKeywordQuotaQueryRequest is the request type for alipay.open.search.appkeywordquota.query.
type SearchAppKeywordQuotaQueryRequest struct {
	// TargetAppid 是小程序id
	TargetAppid *string `json:"target_appid,omitempty" query:"target_appid,omitempty"`
}
