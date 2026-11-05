// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// MiniQRCodeBindResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniQRCodeBindResponse struct {
	// RouteGroup 是路由规则组，用于唯一标记一条路由规则。
	RouteGroup *string `json:"route_group,omitempty"`
}

// MiniQRCodeQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniQRCodeQueryResponse struct {
	// QrCodeRouteGroups 是规则路由数据列表。
	QrCodeRouteGroups *[]QrCodeRouteGroup `json:"qr_code_route_groups,omitempty"`
	// TotalItems 是数据总数。
	TotalItems *int64 `json:"total_items,omitempty"`
}

// MiniQRCodeUnbindResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniQRCodeUnbindResponse struct {
}
