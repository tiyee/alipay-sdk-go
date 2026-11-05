package applyorder

// OfflineMarketApplyOrderBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketApplyOrderBatchQueryResponse struct {
	// BizOrderInfos 是支付宝操作流水信息列表
	BizOrderInfos *[]BizOrderQueryResponse `json:"biz_order_infos,omitempty"`
	// CurrentPageNo 是当前页码
	CurrentPageNo *int64 `json:"current_page_no,omitempty"`
	// PageSize 是每页记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalItems 是总记录数
	TotalItems *int64 `json:"total_items,omitempty"`
	// TotalPageNo 是总页码数目
	TotalPageNo *int64 `json:"total_page_no,omitempty"`
}

// BizOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type BizOrderQueryResponse struct {
	// Action 是操作动作。 	CREATE_SHOP-创建门店， 	MODIFY_SHOP-修改门店， 	CREATE_ITEM-创建商品， 	MODIFY_ITEM-修改商品， 	EFFECTIVE_ITEM-上架商品， 	INVALID_ITEM-下架商品， 	RESUME_ITEM-暂停售卖商品， 	PAUSE_ITEM-恢复售卖商品
	Action *string `json:"action,omitempty"`
	// ActionMode 是操作模式：NORMAL-普通开店
	ActionMode *string `json:"action_mode,omitempty"`
	// ApplyId 是支付宝流水ID
	ApplyId *string `json:"apply_id,omitempty"`
	// BizContextInfo 是流水上下文信息，JSON格式。根据action不同对应的结构也不同，JSON字段与含义可参考各个接口的请求参数。
	BizContextInfo *string `json:"biz_context_info,omitempty"`
	// BizId 是业务主体ID。根据biz_type不同可能对应shop_id或item_id。 	特别注意对于门店创建，当流水status=SUCCESS时，此字段才为shop_id，其他状态时为0或空。
	BizId *string `json:"biz_id,omitempty"`
	// BizType 是业务类型：SHOP-店铺，ITEM-商品
	BizType *string `json:"biz_type,omitempty"`
	// CreateTime 是创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// OpId 是操作用户的支付账号id
	OpId *string `json:"op_id,omitempty"`
	// RequestId 是注意：此字段并非外部商户请求时传入的request_id，暂时代表支付宝内部字段，请勿用。
	RequestId *string `json:"request_id,omitempty"`
	// ResultCode 是流水处理结果码 	点此查看
	ResultCode *string `json:"result_code,omitempty"`
	// ResultDesc 是流水处理结果描述
	ResultDesc *string `json:"result_desc,omitempty"`
	// Status 是流水状态：INIT-初始，PROCESS-处理中，SUCCESS-成功，FAIL-失败。
	Status *string `json:"status,omitempty"`
	// SubStatus 是流水子状态：WAIT_CERTIFY-等待认证，LICENSE_AUDITING-证照审核中，RISK_AUDITING-风控审核中，WAIT_SIGN-等待签约，FINISH-终结。
	SubStatus *string `json:"sub_status,omitempty"`
	// UpdateTime 是更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}
