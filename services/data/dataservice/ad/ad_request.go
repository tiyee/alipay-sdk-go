package ad

// DataDataServiceAdConversionUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type DataDataServiceAdConversionUploadRequest struct {
	// BizToken 是代理商访问灯火平台的token。
	BizToken *string `json:"biz_token,omitempty"`
	// ConversionDataList 是转化数据列表, 最多1000条。
	ConversionDataList *[]ConversionData `json:"conversion_data_list,omitempty"`
}

// DataDataServiceAdPromotePageBatchQueryRequest 是 alipay.data.dataservice.ad.promotepage.batchquery 的 V3 query 请求参数。
type DataDataServiceAdPromotePageBatchQueryRequest struct {
	// BizToken 是代理商访问灯火平台的token。
	BizToken *string `query:"biz_token,omitempty"`
	// PrincipalTag 是商家标志。
	PrincipalTag *string `query:"principal_tag,omitempty"`
	// Type 是推广页类型：COLLECT_INFO - 免费留资；TRADE - 付费留资； OPERATION_PAID - 运营商付费留资；待扩展。
	Type *string `query:"type,omitempty"`
	// PageNo 是分页参数之页数，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页参数之每页大小，最大1000。
	PageSize *string `query:"page_size,omitempty"`
}

// DataDataServiceAdPromotePageDownloadRequest 是 alipay.data.dataservice.ad.promotepage.download 的 V3 query 请求参数。
type DataDataServiceAdPromotePageDownloadRequest struct {
	// StartDate 是留资开始日期，格式：yyyy-mm-dd 或 yyyy-mm-dd hh:mm:ss;不能早于30天前。
	StartDate *string `query:"start_date,omitempty"`
	// EndDate 是留资结束日期，格式：yyyy-mm-dd 或 yyyy-mm-dd hh:mm:ss;不能晚于当天。
	EndDate *string `query:"end_date,omitempty"`
	// PageNo 是分页参数之页数，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页参数之每页大小，最大1000。
	PageSize *string `query:"page_size,omitempty"`
	// BizToken 是代理商访问灯火平台的token。
	BizToken *string `query:"biz_token,omitempty"`
	// PrincipalTag 是商家标志。
	PrincipalTag *string `query:"principal_tag,omitempty"`
	// PromotePageId 是推广页id。
	PromotePageId *string `query:"promote_page_id,omitempty"`
}
