// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package safedomain

// MiniSafeDomainCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniSafeDomainCreateRequest struct {
	// SafeDomain 是服务器域名白名单。支付宝小程序在HTTP请求、上传文件等场景只能与域名白名单中的域名进行通讯 。 新添加域白名单后需重新上传应用版本（模板模式下小程序模板及商家应用都需重新上传、实例化版本）才会生效。老版本应用继续使用老版本名单。 每个应用最多可添加30个 URL。。
	SafeDomain *string `json:"safe_domain,omitempty"`
}

// MiniSafeDomainDeleteRequest is the request type for alipay.open.mini.safedomain.delete.
type MiniSafeDomainDeleteRequest struct {
	// SafeDomain 是服务器域名白名单。支付宝小程序在HTTP请求、上传文件等场景只能与域名白名单中的域名进行通讯 。	删除域白名单后需重新上传应用版本（模板模式下小程序模板及商家应用都需重新上传、实例化版本）才会生效。老版本应用继续使用老版本名单（即删除域名可继续使用）。
	SafeDomain *string `json:"safe_domain,omitempty" query:"safe_domain,omitempty"`
}
