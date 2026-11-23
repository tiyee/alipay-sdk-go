// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package menu

// PublicMenuCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuCreateRequest struct {
	// Button 是一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。。
	Button *[]ButtonObject `json:"button,omitempty"`
	// Type 是菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。。
	Type *string `json:"type,omitempty"`
}

// PublicMenuDataBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuDataBatchQueryRequest struct {
	// BeginDate 是数据开始日期，时间格式为 "yyyyMMdd" 。。
	BeginDate *string `json:"begin_date,omitempty"`
	// EndDate 是数据结束日期，时间格式为"yyyyMMdd"。查询数据开始日期/结束日期时间跨度最大30天。。
	EndDate *string `json:"end_date,omitempty"`
}

// PublicMenuModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuModifyRequest struct {
	// Button 是一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。。
	Button *[]ButtonObject `json:"button,omitempty"`
	// Type 是菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。。
	Type *string `json:"type,omitempty"`
}

// PublicMenuBatchQueryRequest is the request type for alipay.open.public.menu.batchquery.
type PublicMenuBatchQueryRequest struct{}

// PublicMenuQueryRequest is the request type for alipay.open.public.menu.query.
type PublicMenuQueryRequest struct{}
