// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package inner

// MiniInnerExperienceCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerExperienceCancelRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端id。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是业务小程序appid。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Oid 是oid。
	Oid *string `json:"oid,omitempty"`
	// OperatorId 是操作员id。
	OperatorId *string `json:"operator_id,omitempty"`
}

// MiniInnerMembersAddRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerMembersAddRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// DomainAccount 是域账号（不为空时会添加联系人）。
	DomainAccount *string `json:"domain_account,omitempty"`
	// LoginId 是支付宝登陆账号。
	LoginId *string `json:"login_id,omitempty"`
	// MiniAppId 是业务appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Type 是成员类型。
	Type *string `json:"type,omitempty"`
}

// MiniInnerMembersDeleteRequest is the request type for alipay.open.mini.inner.members.delete.
type MiniInnerMembersDeleteRequest struct {
	// MiniAppId 是业务小程序id
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppOrigin 是业务来源
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
	// DomainAccount 是域账号（如传入则删除联系人）
	DomainAccount *string `json:"domain_account,omitempty" query:"domain_account,omitempty"`
	// LoginId 是登陆支付宝账号
	LoginId *string `json:"login_id,omitempty" query:"login_id,omitempty"`
	// Type 是成员类型
	Type *string `json:"type,omitempty" query:"type,omitempty"`
}

// MiniInnerMembersQueryRequest is the request type for alipay.open.mini.inner.members.query.
type MiniInnerMembersQueryRequest struct {
	// MiniAppId 是业务小程序id
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppOrigin 是业务来源
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
}
