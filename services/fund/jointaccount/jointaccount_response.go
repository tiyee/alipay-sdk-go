package jointaccount

// FundJointAccountBillQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountBillQueryResponse struct {
	// BillList 是返回账单列表。账单列表为空时，返回为空。
	BillList *[]JointAccountBillDetailDTO `json:"bill_list,omitempty"`
	// BizScene 是场景。
	BizScene *string `json:"biz_scene,omitempty"`
	// HasNextPage 是是否有下一页。
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// PageNum 是当期页数。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是当期页记录数。
	PageSize *string `json:"page_size,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountDetailQueryResponse struct {
	// AccountId 是合花群ID（支付宝侧生成）。
	AccountId *string `json:"account_id,omitempty"`
	// AccountName 是账户名称（支付宝侧生成）。
	AccountName *string `json:"account_name,omitempty"`
	// AccountQuota 是额度模型。
	AccountQuota *[]JointAccountQuotaRespDTO `json:"account_quota,omitempty"`
	// AccountStatus 是账户状态<br> -NORMAL：正常<br> -RELEASING：注销中<br> -RELEASED：已注销<br> - FREEZE：冻结。
	AccountStatus *string `json:"account_status,omitempty"`
	// AgreementNo 是授权协议号（支付宝侧生成）。
	AgreementNo    *string            `json:"agreement_no,omitempty"`
	AuthorizedRule *AuthorizedRuleDTO `json:"authorized_rule,omitempty"`
	// AvailableBalance 是当前可用金额（单位为元，必须大于0且最多小数点后两位）。
	AvailableBalance *string `json:"available_balance,omitempty"`
	// BizScene 是业务场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// CreatorId 是（创建人）支付宝侧用户唯一标识。
	CreatorId *string `json:"creator_id,omitempty"`
	// CreatorOpenId 是（创建人）支付宝侧用户唯一标识。
	CreatorOpenId *string `json:"creator_open_id,omitempty"`
	// CreatorOutId 是（创建人）商户侧用户唯一标识<br> 补充说明：<br> -如果签约时，发起人标识传递的是商户侧用户唯一标识，则该字段会返回<br> -如果签约时，发起人标识传递的是支付宝侧用户唯一标识，则该字段为空。
	CreatorOutId *string `json:"creator_out_id,omitempty"`
	// FreezeBalance 是当前冻结金额（单位为元，必须大于0且最多小数点后两位）。
	FreezeBalance *string `json:"freeze_balance,omitempty"`
	// InviteResultList 是签约时邀请的成员列表（快照）。
	InviteResultList *[]InviteResultDTO `json:"invite_result_list,omitempty"`
	// MemberList 是已加入合花群的成员列表。
	MemberList *[]JointAccountMemberInfoRespDTO `json:"member_list,omitempty"`
	// ProductCode 是销售产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// ProfitStatus 是生息状态：</br> - MAKING ：开启中</br> - MADE. ：已开启</br> - CLEARING：关闭中</br> - NONE. ：未生息。
	ProfitStatus *string `json:"profit_status,omitempty"`
	// YesterdayProfit 是昨日收益。
	YesterdayProfit *string `json:"yesterday_profit,omitempty"`
}

// FundJointAccountListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountListQueryResponse struct {
	// AccountList 是共同账户列表。
	AccountList *[]JointAccountDTO `json:"account_list,omitempty"`
	// BizScene 是场景。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountMemberBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberBatchQueryResponse struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// MemberList 是员工列表。
	MemberList *[]JointAccountMemberDTO `json:"member_list,omitempty"`
	// PageNum 是当前页数，使用游标方式查询时，没有该字段。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是当前记录数。
	PageSize *string `json:"page_size,omitempty"`
	// TotalCount 是总条数。
	TotalCount *string `json:"total_count,omitempty"`
	// TotalPageCount 是总页数。
	TotalPageCount *string `json:"total_page_count,omitempty"`
}

// FundJointAccountMemberBindResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberBindResponse struct {
	// ReplyUrl 是员工回应邀请链接。
	ReplyUrl *string `json:"reply_url,omitempty"`
}

// FundJointAccountMemberConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberConsultResponse struct {
	// AccountId 是企业ID。
	AccountId *string `json:"account_id,omitempty"`
	// Code 是返回码 10000-成功 其他都是失败。
	Code *string `json:"code,omitempty"`
	// MemberList 是员工列表。
	MemberList *[]MemberStatusRespDTO `json:"member_list,omitempty"`
	// Msg 是描述信息。
	Msg *string `json:"msg,omitempty"`
}

// FundJointAccountMemberQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberQueryResponse struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// BizScene 是场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// MemberList 是成员列表。
	MemberList *[]JointAccountMemberRespDTO `json:"member_list,omitempty"`
	// PageNum 是页码。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是页面记录数。
	PageSize *string `json:"page_size,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// TotalPageCount 是总页数。
	TotalPageCount *string `json:"total_page_count,omitempty"`
}

// FundJointAccountMemberUnbindResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberUnbindResponse struct {
}

// FundJointAccountMemberRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberRuleModifyResponse struct {
}

// FundJointAccountQuotaQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountQuotaQueryResponse struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AccountQuota 是共同账户额度模型。
	AccountQuota *[]JointAccountQuotaRespDTO `json:"account_quota,omitempty"`
	// BizScene 是场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// MemberId 是成员ID（如果查询员工才返回）。
	MemberId *string `json:"member_id,omitempty"`
	// MemberOpenId 是员工openId。
	MemberOpenId *string `json:"member_open_id,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountRuleModifyResponse struct {
}

// FundJointAccountSignQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountSignQueryResponse struct {
	// AccountId 是企业签约的账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是企业授权平台的三方授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Status 是账户状态，NORMAL-生效中，RELEASED-已销户。
	Status *string `json:"status,omitempty"`
}

// FundJointAccountTradeQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountTradeQueryResponse struct {
	// BuyerId 是买家ID。
	BuyerId *string `json:"buyer_id,omitempty"`
	// BuyerOpenId 是买家的openId。
	BuyerOpenId *string `json:"buyer_open_id,omitempty"`
	// CreateTime 是交易创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// TotalAmount 是订单总金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// TradeFundBillList 是资金单明细。
	TradeFundBillList *[]TradeFundBillDetail `json:"trade_fund_bill_list,omitempty"`
	// TradeInfoList 是交易信息集合。查询为批量查询时，使用该结果返回。
	TradeInfoList *[]TradeInfoDTO `json:"trade_info_list,omitempty"`
	// TradeNo 是交易单号。
	TradeNo *string `json:"trade_no,omitempty"`
	// TradeStatus 是交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）。
	TradeStatus *string `json:"trade_status,omitempty"`
}
