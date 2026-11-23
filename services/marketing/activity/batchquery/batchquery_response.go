// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package batchquery

// ActivityBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityBatchQueryResponse struct {
	// ActivityLiteInfos 是活动摘要信息
	ActivityLiteInfos *[]ActivityLiteInfo `json:"activity_lite_infos,omitempty"`
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalSize 是经过条件筛选，查询到的活动总数量
	TotalSize *string `json:"total_size,omitempty"`
}
