// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package app

// ActivityAppBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityAppBatchQueryResponse struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// AppInfos 是可用小程序列表
	AppInfos *[]ActivityAppInfo `json:"app_infos,omitempty"`
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalSize 是可用小程序总数量
	TotalSize *int64 `json:"total_size,omitempty"`
}
