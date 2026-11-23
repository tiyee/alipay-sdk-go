// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package tips

// TipsDelivery 是支付宝官方 V3 OpenAPI 模型。
type TipsDelivery struct {
	// DeliveryContent 是收藏引导文案内容，不得超过14个字。
	DeliveryContent *string `json:"delivery_content,omitempty"`
	// DeliveryId 是收藏引导投放活动ID。
	DeliveryId *string `json:"delivery_id,omitempty"`
	// DeliveryName 是收藏引导活动名称。
	DeliveryName *string `json:"delivery_name,omitempty"`
	// EndTime 是活动结束时间。
	EndTime *string `json:"end_time,omitempty"`
	// FailReason 是文案审核驳回理由，仅状态为AUDIT_REJECTED时有效。
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
