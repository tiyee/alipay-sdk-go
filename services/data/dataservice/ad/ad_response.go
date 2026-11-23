package ad

// DataDataServiceAdConversionUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type DataDataServiceAdConversionUploadResponse struct {
}

// DataDataServiceAdPromotePageBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataDataServiceAdPromotePageBatchQueryResponse struct {
	// List 是自建推广页详情列表。
	List *[]PromotePageDetail `json:"list,omitempty"`
	// PageNo 是分页参数之页数。
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是分页参数之每页大小。
	PageSize *int64 `json:"page_size,omitempty"`
	// Total 是分页参数之总数据量。
	Total *int64 `json:"total,omitempty"`
}

// DataDataServiceAdPromotePageDownloadResponse 是支付宝官方 V3 OpenAPI 模型。
type DataDataServiceAdPromotePageDownloadResponse struct {
	// ConversionId 是转化事件ID。
	ConversionId *string `json:"conversion_id,omitempty"`
	// List 是留资属性实例数据列表。
	List *[]PromotePageData `json:"list,omitempty"`
	// PageNo 是分页参数之页数，从1开始。
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是分页参数之每页大小，最大1000。
	PageSize *int64 `json:"page_size,omitempty"`
	// Total 是分页参数之总数据量。
	Total *int64 `json:"total,omitempty"`
}
