package subservicekeyword

// SearchSubServiceKeywordApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordApplyRequest struct {
	// ConfigId 是关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取。
	ConfigId *string `json:"config_id,omitempty"`
	// KeyWords 是服务关键词列表，每批最多传入30个，关键词长度小于12个汉字。 超过数量限制会申请失败，剩余关键词可通过 <a href="https://opendocs.alipay.com/mini/062ndt?pathHash=e3e78b68&ref=api&scene=common">alipay.open.search.appkeywordquota.query</a>查询。
	KeyWords *[]string `json:"key_words,omitempty"`
	// SubServiceCode 是子服务code，提报服务关键词，alipay.open.app.service.list.query(服务批量查询)这个接口可以获取。
	SubServiceCode *string `json:"sub_service_code,omitempty"`
	// TargetAppid 是小程序id。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchSubServiceKeywordDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchSubServiceKeywordDeleteRequest struct {
	// ConfigId 是关键词配置ID。
	ConfigId *string `json:"config_id,omitempty"`
	// TargetAppid 是小程序id。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchSubServiceKeywordBatchQueryRequest is the request type for alipay.open.search.subservicekeyword.batchquery.
type SearchSubServiceKeywordBatchQueryRequest struct {
	// TargetAppid 是小程序id
	TargetAppid *string `json:"target_appid,omitempty" query:"target_appid,omitempty"`
	// Status 是审核状态，服务关键词的审核状态，枚举值，不传该字段表示查询所有状态
	Status *string `json:"status,omitempty" query:"status,omitempty"`
	// PageNumber 是分页参数，分页查询时使用
	PageNumber *int64 `json:"page_number,omitempty" query:"page_number,required"`
	// PageSize 是分页参数，分页查询时使用
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}

// SearchSubServiceKeywordQueryStatusRequest is the request type for alipay.open.search.subservicekeyword.querystatus.
type SearchSubServiceKeywordQueryStatusRequest struct {
	// TargetAppid 是小程序id
	TargetAppid *string `json:"target_appid,omitempty" query:"target_appid,omitempty"`
	// ApplyNo 是审核工单id，提报接口的返回值
	ApplyNo *string `json:"apply_no,omitempty" query:"apply_no,omitempty"`
}
