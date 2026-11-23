// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

// MiniISVCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniISVCreateRequest struct {
	CreateMiniRequest *CreateMiniRequest `json:"create_mini_request,omitempty"`
}

// MiniISVFastRegisterCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniISVFastRegisterCreateRequest struct {
	// AppName 是小程序名称，昵称半自动设定，强制后缀“的试用小程序”。且该参数会进行关键字检查，如果命中品牌关键字则会报错。 如遇到品牌大客户要用试用小程序，建议用户先换个名字，认证后再修改成品牌名。。
	AppName *string `json:"app_name,omitempty"`
	// AuthNotifyUrl 是授权确认成功后回调h5链接或小程序页面。
	AuthNotifyUrl *string `json:"auth_notify_url,omitempty"`
	// OutOrderNo 是开发者外部订单号，通过开发者账号+外部订单号做业务幂等。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// Uid 是商户uid。
	Uid *string `json:"uid,omitempty"`
}

// CreateMiniRequest 是支付宝官方 V3 OpenAPI 模型。
type CreateMiniRequest struct {
	// AlipayAccount 是商家登录支付宝的邮箱帐号或手机号。（1）默认只支持企业账号类型；（2）将is_individual设置为true，支持个体工商户类型的账号（同时也兼容企业账号）（3）将is_individual设置为true且上传营业执照照片，将同时对具备个体工商户营业执照的商家个人账号认证成为个体工商户账号（同时也兼容企业账号、个体工商户账号）。
	AlipayAccount *string `json:"alipay_account,omitempty"`
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// CertName 是营业执照企业名称，如果是“无主体名称个体工商户”则填“个体户+法人姓名”，例如“个体户张三”。
	CertName *string `json:"cert_name,omitempty"`
	// CertNo 是营业执照编码。
	CertNo *string `json:"cert_no,omitempty"`
	// ContactName 是商家联系人名称。
	ContactName *string `json:"contact_name,omitempty"`
	// ContactPhone 是商家联系人手机电话。
	ContactPhone *string `json:"contact_phone,omitempty"`
	// IsIndividual 是默认为 false。当设置为 true 时，支持个体工商户的账号类型（同时兼容企业账号）。
	IsIndividual *bool `json:"is_individual,omitempty"`
	// LegalPersonalName 是商家法人名称。
	LegalPersonalName *string `json:"legal_personal_name,omitempty"`
	// LicensePic 是营业执照图片的Base64编码字符串，图片大小不能超过2M。将is_individual设置为true，当传入该参数后，如果商家账号不是个体工商户类型，将同时为商家升级账号，无需商家额外操作升级（同时支持企业账号、个体工商户账号）。
	LicensePic *string `json:"license_pic,omitempty"`
	// OutOrderNo 是开发者外部订单号，自定义传入，通过开发者账号+outOrderNo做业务幂等。（1）该字段会在商家确认创建后通过应用授权通知出参notify_context以out_biz_no属性返回给ISV，应用授权通知订阅参考：<a href="https://opendocs.alipay.com/isv/01fejh">https://opendocs.alipay.com/isv/01fejh</a>。（2）该字段也会通过<a href="https://opendocs.alipay.com/mini/03l3f1?pathHash=5c81f4fb&scene=common">alipay.open.mini.merchant.confirmed</a>返回给订阅方。
	OutOrderNo *string `json:"out_order_no,omitempty"`
}

// MiniISVFastRegisterQueryRequest is the request type for alipay.open.mini.isv.fastregister.query.
type MiniISVFastRegisterQueryRequest struct {
	// OrderNo 是代创建试用小程序单号
	OrderNo *string `json:"order_no,omitempty" query:"order_no,omitempty"`
}

// MiniISVQueryRequest is the request type for alipay.open.mini.isv.query.
type MiniISVQueryRequest struct {
	// OrderNo 是服务商代商家创建小程序，由支付宝开放平台返回的订单号用于后续查询处理进展。
	OrderNo *string `json:"order_no,omitempty" query:"order_no,omitempty"`
}
