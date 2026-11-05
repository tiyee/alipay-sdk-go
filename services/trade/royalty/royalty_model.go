package royalty

// RoyaltyEntityType 是分账接收方类型，取值依据官方 V3 文档。
type RoyaltyEntityType string

const (
	// RoyaltyEntityTypeUserID 表示支付宝用户号。
	RoyaltyEntityTypeUserID RoyaltyEntityType = "userId"
	// RoyaltyEntityTypeLoginName 表示支付宝登录号。
	RoyaltyEntityTypeLoginName RoyaltyEntityType = "loginName"
	// RoyaltyEntityTypeOpenID 表示用户在该应用下的 OpenId。
	RoyaltyEntityTypeOpenID RoyaltyEntityType = "openId"
)

// IsKnown 报告接收方类型是否属于文档已知值。
func (t RoyaltyEntityType) IsKnown() bool {
	switch t {
	case RoyaltyEntityTypeUserID, RoyaltyEntityTypeLoginName, RoyaltyEntityTypeOpenID:
		return true
	default:
		return false
	}
}

// RoyaltyEntity 是分账收款方（接受分账的账户）。
type RoyaltyEntity struct {
	// Type 是分账接收方类型。
	Type RoyaltyEntityType `json:"type"`
	// Account 是分账接收方账号：type=userId 时为 2088 开头的 16 位
	// 用户号，type=loginName 时为登录号，type=openId 时为 openId。
	Account *string `json:"account,omitempty"`
	// AccountOpenID 是分账接收方在该应用下的 OpenId；仅查询接口返回。
	AccountOpenID *string `json:"account_open_id,omitempty"`
	// Name 是分账接收方真实姓名；type=loginName 绑定时必传。
	Name *string `json:"name,omitempty"`
	// Memo 是分账关系描述。
	Memo *string `json:"memo,omitempty"`
	// BindLoginName 是绑定时的登录号；仅查询接口返回。
	BindLoginName *string `json:"bind_login_name,omitempty"`
	// LoginName 是当前 userId 对应的登录号；仅查询接口返回。
	LoginName *string `json:"login_name,omitempty"`
}
