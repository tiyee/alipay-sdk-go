package servicemarket

// PromotionTaskRelation 是支付宝官方 V3 OpenAPI 模型。
type PromotionTaskRelation struct {
	// ApplyReason 是申请推广时填写的申请理由，自定义的内容。
	ApplyReason *string `json:"apply_reason,omitempty"`
	// CommodityId 是服务商品ID。
	CommodityId *string `json:"commodity_id,omitempty"`
	// CommodityName 是服务商品名称。
	CommodityName *string `json:"commodity_name,omitempty"`
	// CreateTime 是推广关系创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// PromoterContactName 是推广服务商联系人名称。
	PromoterContactName *string `json:"promoter_contact_name,omitempty"`
	// PromoterContactPhone 是推广服务商联系电话。
	PromoterContactPhone *string `json:"promoter_contact_phone,omitempty"`
	// PromoterName 是推广服务商的名称。
	PromoterName *string `json:"promoter_name,omitempty"`
	// PromoterPid 是推广服务商的pid。
	PromoterPid *string `json:"promoter_pid,omitempty"`
	// PromotionId 是推广任务id。
	PromotionId *string `json:"promotion_id,omitempty"`
	// PromotionName 是推广任务的名称。
	PromotionName *string `json:"promotion_name,omitempty"`
}
