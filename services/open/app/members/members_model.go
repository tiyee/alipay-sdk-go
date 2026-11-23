package members

// AppMemberInfo 是支付宝官方 V3 OpenAPI 模型。
type AppMemberInfo struct {
	// GmtInvite 是邀请时间，格式为yyyy-MM-dd HH:mm:ss。
	GmtInvite *string `json:"gmt_invite,omitempty"`
	// GmtJoin 是加入时间，格式为yyyy-MM-dd。
	GmtJoin *string `json:"gmt_join,omitempty"`
	// LogonId 是支付宝登录账号。
	LogonId *string `json:"logon_id,omitempty"`
	// NickName 是昵称。
	NickName *string `json:"nick_name,omitempty"`
	// OpenId 是支付宝用户openId。
	OpenId *string `json:"open_id,omitempty"`
	// Portrait 是头像地址fileId(完整地址需自行拼接)。
	Portrait *string `json:"portrait,omitempty"`
	// Role 是角色类型。
	Role *string `json:"role,omitempty"`
	// Status 是成员的状态。
	Status *string `json:"status,omitempty"`
	// UserId 是支付宝用户id。
	UserId *string `json:"user_id,omitempty"`
}
