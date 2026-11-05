package members

// AppMembersCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type AppMembersCreateRequest struct {
	// LogonId 是支付宝登录账号。。
	LogonId *string `json:"logon_id,omitempty"`
	// Role 是为成员添加的角色类型。
	Role *string `json:"role,omitempty"`
}

// AppMembersDeleteRequest is the request type for alipay.open.app.members.delete.
type AppMembersDeleteRequest struct {
	// UserId 是被删除成员的支付宝账户唯一标识，以2088开头。
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是被删除成员的openId。
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// Role 是被删除成员的角色类型
	Role *string `json:"role,omitempty" query:"role,omitempty"`
}

// AppMembersQueryRequest is the request type for alipay.open.app.members.query.
type AppMembersQueryRequest struct {
	// Role 是查询的成员角色类型
	Role *string `json:"role,omitempty" query:"role,omitempty"`
}
