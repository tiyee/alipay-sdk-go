package applyorder

// OfflineMarketApplyOrderBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketApplyOrderBatchQueryRequest struct {
	// Action 是操作动作
	Action *string `json:"action,omitempty"`
	// ApplyIds 是支付宝流水ID列表，最大不超过100个
	ApplyIds *[]string `json:"apply_ids,omitempty"`
	// BizId 是业务主体ID。根据biz_type不同可能对应shop_id或item_id。
	BizId *string `json:"biz_id,omitempty"`
	// BizType 是业务类型：SHOP-店铺，ITEM-商品。
	BizType *string `json:"biz_type,omitempty"`
	// EndTime 是查询的流水创建时间最后值。格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"end_time,omitempty"`
	// OpId 是操作用户的支付账号id
	OpId *string `json:"op_id,omitempty"`
	// OpRole 是系统集成商统一传入ISV
	OpRole *string `json:"op_role,omitempty"`
	// PageNo 是页码，留空标示第一页，默认20个结果为一页
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是每页记录数。默认20，最大100。
	PageSize *int64 `json:"page_size,omitempty"`
	// RequestIds 是请求ID列表，最大不超过100个。 	注意：暂时不支持此字段查询。
	RequestIds *[]string `json:"request_ids,omitempty"`
	// StartTime 是查询的流水创建时间起始值，只能查询近3个月数据。格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"start_time,omitempty"`
	// Status 是流水状态：INIT-初始，PROCESS-处理中，SUCCESS-成功，FAIL-失败。
	Status *string `json:"status,omitempty"`
}
