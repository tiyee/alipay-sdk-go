package live

// MerchantLiveItemPlanModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantLiveItemPlanModifyRequest struct {
	// Content 是投放内容，数组，内容为具体条目
	Content *[]ItemPlanContentDO `json:"content,omitempty"`
	// EndTime 是可空，计划截止时间
	EndTime *string `json:"end_time,omitempty"`
	// OpType 是针对投放内容的操作， 包含3种类型：	覆盖:OVERRIDE 	追加:APPEND 	删除:DELETE
	OpType *string `json:"op_type,omitempty"`
	// OperatorAppid 是投放方标识，一般为小程序id
	OperatorAppid *string `json:"operator_appid,omitempty"`
	// OperatorType 是投放方类型，生活号：LIFE_APP 达人：CONTENT_USER 小程序：TINYAPP
	OperatorType *string `json:"operator_type,omitempty"`
	// PlanName 是投放计划名，可空
	PlanName *string `json:"plan_name,omitempty"`
	// StartTime 是可空，计划生效时间
	StartTime *string `json:"start_time,omitempty"`
	// Status 是计划状态 0待生效 1生效中 -1已作废
	Status *int64 `json:"status,omitempty"`
	// TargetId 是投放目标，一般为生活号id，多个逗号隔开
	TargetId *string `json:"target_id,omitempty"`
	// TargetRange 是投放范围， 指主体内的具体子渠道，为空则任意子渠道
	TargetRange *string `json:"target_range,omitempty"`
	// TargetType 是投放目标类型	生活号：LIFE_APP 	达人：CONTENT_USER 	小程序:TINYAPP
	TargetType *string `json:"target_type,omitempty"`
}

// MerchantLiveChannelQueryRequest is the query request model for alipay.merchant.live.channel.query.
type MerchantLiveChannelQueryRequest struct {
	// Secret 是渠道下单参数(orderStartId)
	Secret *string `json:"secret,omitempty" query:"secret,omitempty"`
}
