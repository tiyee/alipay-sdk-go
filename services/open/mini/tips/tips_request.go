// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package tips

// MiniTipsDeliveryBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryBatchQueryRequest struct {
	// PageNumber 是每页记录条数，最小1，最大20。
	PageNumber *string `json:"page_number,omitempty"`
	// PageSize 是查询的页数，从1开始，最大1000。
	PageSize *string `json:"page_size,omitempty"`
}

// MiniTipsDeliveryCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryCreateRequest struct {
	// DeliveryContent 是收藏引导文案内容，不得超过14个字。
	DeliveryContent *string `json:"delivery_content,omitempty"`
	// DeliveryName 是收藏引导活动名称。
	DeliveryName *string `json:"delivery_name,omitempty"`
	// EndTime 是活动结束时间。
	EndTime *string `json:"end_time,omitempty"`
	// MatchType 是匹配类型。
	MatchType *string `json:"match_type,omitempty"`
	// MatchUrl 是目标页面地址，当匹配类型为TARGETURL时，投放文案仅在当前页面生效。
	MatchUrl *string `json:"match_url,omitempty"`
	// StartTime 是活动开始时间。
	StartTime *string `json:"start_time,omitempty"`
}

// MiniTipsDeliveryModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryModifyRequest struct {
	// DeliveryId 是收藏引导投放活动ID。
	DeliveryId *string `json:"delivery_id,omitempty"`
	// OperateType 是活动投放操作类型。
	OperateType *string `json:"operate_type,omitempty"`
}

// MiniTipsDeliveryQueryRequest is the request type for alipay.open.mini.tips.delivery.query.
type MiniTipsDeliveryQueryRequest struct {
	// DeliveryId 是收藏引导投放活动ID，查询收藏引导活动必填参数
	DeliveryId *string `json:"delivery_id,omitempty" query:"delivery_id,omitempty"`
}

// MiniTipsStatisticQueryRequest is the request type for alipay.open.mini.tips.statistic.query.
type MiniTipsStatisticQueryRequest struct {
	// DeliveryId 是收藏引导投放活动ID，供查询收藏引导活动配置接口调用	，当以小程序维度查询数据（query_type为app）时delivery_id为空
	DeliveryId *string `json:"delivery_id,omitempty" query:"delivery_id,omitempty"`
	// QueryType 是查询类型，表示以当前维度进行数据聚合。
	QueryType *string `json:"query_type,omitempty" query:"query_type,omitempty"`
	// StartDate 是查询开始日期，精度为天
	StartDate *string `json:"start_date,omitempty" query:"start_date,omitempty"`
	// EndDate 是查询截止日期，精度为天
	EndDate *string `json:"end_date,omitempty" query:"end_date,omitempty"`
}
