// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package tips

// MiniTipsDeliveryBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryBatchQueryResponse struct {
	// TipsDeliveryList 是收藏引导活动信息列表。根据活动状态排序，顺序为：生效中，暂停，审核中，白名单生效中，初始化，审核失败，结束。相同状态优先展示状态更新时间最迟的活动。。
	TipsDeliveryList *[]TipsDelivery `json:"tips_delivery_list,omitempty"`
	// TotalPages 是总页数。
	TotalPages *int64 `json:"total_pages,omitempty"`
	// TotalSize 是收藏引导数据总数量。
	TotalSize *int64 `json:"total_size,omitempty"`
}

// MiniTipsDeliveryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryCreateResponse struct {
	// DeliveryId 是收藏引导活动ID。
	DeliveryId *string `json:"delivery_id,omitempty"`
}

// MiniTipsDeliveryModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryModifyResponse struct {
}

// MiniTipsDeliveryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsDeliveryQueryResponse struct {
	// DeliveryContent 是收藏引导文案内容，不得超过14个字。
	DeliveryContent *string `json:"delivery_content,omitempty"`
	// DeliveryId 是收藏引导投放活动ID。
	DeliveryId *string `json:"delivery_id,omitempty"`
	// DeliveryName 是收藏引导配置投放活动的名称。
	DeliveryName *string `json:"delivery_name,omitempty"`
	// EndTime 是活动结束时间。
	EndTime *string `json:"end_time,omitempty"`
	// FailReason 是文案审核驳回理由，仅AUDIT_REJECTED生效。
	FailReason *string `json:"fail_reason,omitempty"`
	// MatchType 是匹配类型。
	MatchType *string `json:"match_type,omitempty"`
	// MatchUrl 是目标页面地址，当匹配类型为TARGETURL时，投放文案仅在当前页面生效。
	MatchUrl *string `json:"match_url,omitempty"`
	// StartTime 是活动开始时间。
	StartTime *string `json:"start_time,omitempty"`
	// Status 是活动状态。
	Status *string `json:"status,omitempty"`
}

// MiniTipsStatisticQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniTipsStatisticQueryResponse struct {
	// AppCollectCnt 是昨日（接口调用当日的前一天）小程序收藏人数，时间内未创建投放活动，则响应时不传出参数。单位：个。。
	AppCollectCnt *int64 `json:"app_collect_cnt,omitempty"`
	// AppUv 是昨日（接口调用当日的前一天）小程序访问人数。当小程序日访问量为0，或以活动维度查询，或时间内未创建投放活动，则响应时不传出参数。单位：个。。
	AppUv *int64 `json:"app_uv,omitempty"`
	// DeliveryContent 是收藏文案内容（仅以app维度查询时为空）,仅活动维度查询成功时返回。。
	DeliveryContent *string `json:"delivery_content,omitempty"`
	// DeliveryId 是收藏引导投放活动ID（仅以app维度查询时为空）,仅活动维度查询成功时返回。。
	DeliveryId *string `json:"delivery_id,omitempty"`
	// TipsCollectRate 是昨日（接口调用当日的前一天）收藏引导文案收藏转化率，数据为小数字符串类型，取小数点后两位，单位为百分比。小程序曝光数为0，或时间内未创建投放活动，则响应时不传出参数。。
	TipsCollectRate *string `json:"tips_collect_rate,omitempty"`
	// TipsCollectUv 是昨日（接口调用当日的前一天）收藏引导文案收藏用户数，时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TipsCollectUv *int64 `json:"tips_collect_uv,omitempty"`
	// TipsExpoUv 是昨日（接口调用当日的前一天）收藏引导文案曝光用户数，时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TipsExpoUv *int64 `json:"tips_expo_uv,omitempty"`
	// TotalAppCollectCnt 是时间段内累计小程序收藏人数，时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TotalAppCollectCnt *int64 `json:"total_app_collect_cnt,omitempty"`
	// TotalAppUv 是时间段内累计小程序访问人数。当小程序日访问量为0，或以活动维度查询，或时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TotalAppUv *int64 `json:"total_app_uv,omitempty"`
	// TotalTipsCollectRate 是时间段内收藏引导文案收藏转化率，数据为小数字符串类型，取小数点后两位，单位为百分比。小程序曝光数为0，或时间内未创建投放活动，则响应时不传出参数。。
	TotalTipsCollectRate *string `json:"total_tips_collect_rate,omitempty"`
	// TotalTipsCollectUv 是时间段内累计收藏引导文案收藏用户数，时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TotalTipsCollectUv *int64 `json:"total_tips_collect_uv,omitempty"`
	// TotalTipsExpoUv 是时间段内累计收藏引导文案曝光用户数，仅当小程序/活动投放活动时返回。时间内未创建投放活动，则响应时不传出参数。单位：个。。
	TotalTipsExpoUv *int64 `json:"total_tips_expo_uv,omitempty"`
}
