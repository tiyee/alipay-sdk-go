// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package menu

// PublicMenuBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuBatchQueryResponse struct {
	// Count 是菜单数量，包括默认菜单和个性化菜单。
	Count *string `json:"count,omitempty"`
	// Menus 是菜单列表。
	Menus *[]QueryMenu `json:"menus,omitempty"`
}

// PublicMenuCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuCreateResponse struct {
	// MenuKey 是默认菜单菜单key，文本菜单为“default”，icon菜单为“iconDefault”。
	MenuKey *string `json:"menu_key,omitempty"`
}

// PublicMenuDataBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuDataBatchQueryResponse struct {
	// DataList 是菜单分析数据列表。
	DataList *[]MenuAnalysisData `json:"data_list,omitempty"`
}

// PublicMenuModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuModifyResponse struct {
}

// PublicMenuQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMenuQueryResponse struct {
	// MenuContent 是一级菜单数组，个数应为1~4个。
	MenuContent *string `json:"menu_content,omitempty"`
}
