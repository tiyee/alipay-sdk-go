package settle

// SettleConfirmExtendParams 是确认结算扩展参数。
type SettleConfirmExtendParams struct {
	// RoyaltyFreeze 标识是否进行资金冻结用于后续分账（true/false）。
	RoyaltyFreeze *string `json:"royalty_freeze,omitempty"`
}

// SettleEntity 是结算主体信息。
type SettleEntity struct {
	// SettleEntityType 是结算主体类型：SecondMerchant 或 Store。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// SettleEntityID 是结算主体账号。
	SettleEntityID *string `json:"settle_entity_id,omitempty"`
	// SettleEntityBizType 是结算主体资产业务类型（如 settled 已结算资产）。
	SettleEntityBizType *string `json:"settle_entity_biz_type,omitempty"`
}
