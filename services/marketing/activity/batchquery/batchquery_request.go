// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package batchquery

// ActivityBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityBatchQueryRequest struct {
	// ActivityStatus 是活动状态 。 ACTIVE:活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。 PAUSE:活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。
	ActivityStatus *string `json:"activity_status,omitempty"`
	// MerchantAccessMode 是商户接入模式。
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数。
	PageSize *int64 `json:"page_size,omitempty"`
}
