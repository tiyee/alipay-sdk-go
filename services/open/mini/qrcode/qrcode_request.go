// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// MiniQRCodeBindRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniQRCodeBindRequest struct {
	// Mode 是匹配规则，仅支持EXACT（精确匹配）、FUZZY（模糊匹配）两个值。 精确匹配：根据填写的二维码地址精确匹配，地址完全一致时才能唤起小程序（如：配置二维码地址为https://www.alipay.com/my?id=123，当用户扫这个地址的二维码可唤起小程序）。 模糊匹配：根据填写的二维码地址模糊匹配，只要地址前缀匹配即可唤起小程序（如：配置二维码地址为https://www.alipay.com/my/，当用户扫的二维码地址为https://www.alipay.com/my/id=123,可唤起小程序）。。
	Mode *string `json:"mode,omitempty"`
	// PageRedirection 是小程序功能页，配置扫描二维码后打开的小程序功能页面路径。
	PageRedirection *string `json:"page_redirection,omitempty"`
	// RoutePattern 是模式，用于描述通用的路由url格式，适用于同一个host下面存在大量路由的场景，主要用于提升路由性能，其中{0}、{1} 代表路径中的变量。
	RoutePattern *string `json:"route_pattern,omitempty"`
	// RouteUrl 是二维码域名，须通过ICP备案验证，支持http、https、ftp开头的链接。
	RouteUrl *string `json:"route_url,omitempty"`
}

// MiniQRCodeUnbindRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniQRCodeUnbindRequest struct {
	// RouteGroup 是路由规则组，用于唯一标记一条路由规则。调用 <a href="https://opendocs.alipay.com/apis/00rkye" target="_blank">关联普通二维码</a> 接口关联普通二维码后的返回值。。
	RouteGroup *string `json:"route_group,omitempty"`
	// RouteUrl 是路由规则URL，调用<a href="https://opendocs.alipay.com/apis/00rkye" target="_blank">关联普通二维码</a>接口传入的参数：route_url。解除模式与二维码的绑定关系需传入。
	RouteUrl *string `json:"route_url,omitempty"`
}

// MiniQRCodeQueryRequest is the request type for alipay.open.mini.qrcode.query.
type MiniQRCodeQueryRequest struct {
	// PageNo 是页码
	PageNo *string `json:"page_no,omitempty" query:"page_no,omitempty"`
	// PageSize 是每页条数，最大为10
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}
