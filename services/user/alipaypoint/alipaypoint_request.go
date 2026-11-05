package alipaypoint

// UserAlipayPointSendRequest 是支付宝官方 V3 OpenAPI 模型。
type UserAlipayPointSendRequest struct {
	// BudgetCode 是签约商户的集分宝的预算库，扣除此预算库的集分宝发放给用户。会校验budgetcode和业务方appId的签约商户pid的关联关系，若无关则发放失败。
	BudgetCode *string `json:"budget_code,omitempty"`
	// Memo 是商户关于该笔发放的描述或者信息补充，仅存储，无实际校验功能，该信息会在"集分宝"小程序的"集分宝明细"中展示给用户。
	Memo *string `json:"memo,omitempty"`
	// OpenId 是被发放集分宝用户ID，商户app_id维度下的用户标识，与user_account字段二选一。 建议优先使该字段指定发放用户，接口性能更佳
	OpenId *string `json:"open_id,omitempty"`
	// PartnerBizNo 是接入方自己交易的唯一流水ID号，不同交易请保证幂等号唯一性，集分宝服务将依据该字段来进行幂等控制，重试请求不要更换幂等号，需要慎重传递，否则可能会造成损失。
	PartnerBizNo *string `json:"partner_biz_no,omitempty"`
	// PointAmount 是发放给用户的集分宝个数。个数区间为 [1,10000000]，需为整数。
	PointAmount *int64 `json:"point_amount,omitempty"`
	// UserAccount 是被发放集分宝用户的支付宝登录号，邮箱地址或者手机号均可，与user_id字段二选一，在有user_id时，优先使用user_id字段。
	UserAccount *string `json:"user_account,omitempty"`
	// UserId 是被发放集分宝用户的蚂蚁统一会员ID，与user_account字段二选一。	建议优先使用user_id，接口性能更佳，user_id和user_account都传递时，系统优先使用本字段。
	UserId *string `json:"user_id,omitempty"`
}

// UserAlipayPointBudgetLibQueryRequest is the query request model for alipay.user.alipaypoint.budgetlib.query.
type UserAlipayPointBudgetLibQueryRequest struct {
	// BudgetCode 是集分宝预算库编码
	BudgetCode *string `json:"budget_code,omitempty" query:"budget_code,omitempty"`
}
