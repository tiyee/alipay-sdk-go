// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// QrCodeRouteGroup 是支付宝官方 V3 OpenAPI 模型。
type QrCodeRouteGroup struct {
	// Mode 是匹配规则，EXACT（精确匹配）、FUZZY（模糊匹配）、PATTERN（模式匹配）\\（如：配置二维码地址为https://www.alipay.com/my?id=123，当用户扫这个地址的二维码可唤起小程序）。 模糊匹配：根据填写的二维码地址模糊匹配，只要地址前缀匹配即可唤起小程序（如：配置二维码地址为https://www.alipay.com/my/，当用户扫的二维码地址为https://www.alipay.com/my/id=123,可唤起小程序）。 。 模式匹配：根据填写的二维码地址可变变量进行匹配，只要地址的变量位置自定义，变量之外的部分匹配即可唤起小程序(如：配置的二维码地址为https://www.alipay.com/{0}/my/{1},当用户扫的二维码地址为：https://www.alipay.com/user/my/scan，可唤起小程序)。
	Mode *string `json:"mode,omitempty"`
	// RouteGroup 是路由组id（参数说明：该参数可用于alipay.open.mini.qrcode.unbind接口入参route_group，进行二维码解绑）。
	RouteGroup *string `json:"route_group,omitempty"`
	// RouteUrl 是规则路由地址。
	RouteUrl *string `json:"route_url,omitempty"`
}
