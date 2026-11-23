// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package violation

// ViolationViolationEventBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type ViolationViolationEventBatchQueryRequest struct {
	// BeginTime 是查询开始时间 时间格式：yyyy-MM-dd HH:mm:ss 若不填写，则默认当前时间30天前。 注意：begin_time和end_time时间跨度最大一年。
	BeginTime *string `json:"begin_time,omitempty"`
	// EndTime 是查询结束时间 时间格式：yyyy-MM-dd HH:mm:ss 若不填写，则默认当前时间。 注意：begin_time和end_time时间跨度最大一年。
	EndTime *string `json:"end_time,omitempty"`
	// Status 是违规工单状态。
	Status *string `json:"status,omitempty"`
}

// ViolationViolationDetailQueryRequest is the request type for alipay.open.violation.violationdetail.query.
type ViolationViolationDetailQueryRequest struct {
	// ViolationRecordId 是支付宝侧生成的违规记录唯一标识
	ViolationRecordId *string `json:"violation_record_id,omitempty" query:"violation_record_id,omitempty"`
}
