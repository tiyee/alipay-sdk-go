package usertrade

// FundUserTradeBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type FundUserTradeBatchQueryRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// BizScene 是查询交易数据场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// EndTime 是拉取的结算时间，格式yyyy-MM-dd HH:mm:ss，拉取逻辑包含该时刻。。
	EndTime *string `json:"end_time,omitempty"`
	// PageIndex 是分页页码，从1开始，必须大于0。
	PageIndex *string `json:"page_index,omitempty"`
	// PageSize 是分页大小，必须大于0，最大设置100。
	PageSize *string `json:"page_size,omitempty"`
	// ProductCode 是查询交易数据产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// StartTime 是拉取的起始时间，格式yyyy-MM-dd HH:mm:ss，时间必须早于拉取的截止时间，并且，接口仅限查询用户30天内交易数据，拉取逻辑包含该时刻。。
	StartTime *string `json:"start_time,omitempty"`
}
