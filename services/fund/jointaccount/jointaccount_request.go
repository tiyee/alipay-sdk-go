package jointaccount

// FundJointAccountMemberBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberBatchQueryRequest struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是三方授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `json:"biz_scene,omitempty"`
	// Identity 是企业账号：identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号。
	Identity *string `json:"identity,omitempty"`
	// IdentityType 是账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式。
	IdentityType *string `json:"identity_type,omitempty"`
	// LastOpenId 是按游标查询，传入上次查询最后一个用户的openId。
	LastOpenId *string `json:"last_open_id,omitempty"`
	// LastUserId 是按游标查询，传入上次查询最后一个用户ID，适用于逐页查询。优先级大于page_num。
	LastUserId *string `json:"last_user_id,omitempty"`
	// PageNum 是采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据。。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是每页大小，不传的情况下默认20条，上限100；不足20条则按实际记录数返回。
	PageSize *string `json:"page_size,omitempty"`
	// ProductCode 是销售产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountMemberBindRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberBindRequest struct {
	// AccountId 是账本ID。
	AccountId    *string               `json:"account_id,omitempty"`
	AccountQuota *JointAccountQuotaDTO `json:"account_quota,omitempty"`
	// AgreementNo 是授权签约协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AgreementSignInfo 是关联代扣协议签约串。
	AgreementSignInfo *string `json:"agreement_sign_info,omitempty"`
	// BizScene 是业务场景，联系支付宝分配。
	BizScene       *string                        `json:"biz_scene,omitempty"`
	BusinessParams *InviteMemberBusinessParamsDTO `json:"business_params,omitempty"`
	// Identity 是该字段为被邀请方账号ID： 当identity_type为ALIPAY_USER_ID时，需要填支付宝ID；当identity_type为ALIPAY_LOGON_ID时，需要填支付宝登录号；当identity_type为ALIPAY_OPEN_ID时，需要填支付宝openId。
	Identity *string `json:"identity,omitempty"`
	// IdentityType 是账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、ALIPAY_OPEN_ID：支付宝openId。
	IdentityType         *string                  `json:"identity_type,omitempty"`
	IdentityVerifiedInfo *IdentityVerifiedInfoDTO `json:"identity_verified_info,omitempty"`
	// Name 是成员姓名，账号类型为ALIPAY_LOGON_ID时必填。
	Name *string `json:"name,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountMemberConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberConsultRequest struct {
	// AccountId 是企业ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BusinessParams 是扩展参数，如需使用请联系支付宝开放同学。
	BusinessParams *string `json:"business_params,omitempty"`
	// OpenIdList 是要检查的员工openId。
	OpenIdList *[]string `json:"open_id_list,omitempty"`
	// UserIdList 是要检查的员工ID。
	UserIdList *[]string `json:"user_id_list,omitempty"`
}

// FundJointAccountMemberUnbindRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberUnbindRequest struct {
	// AccountId 是账本id。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// Identity 是成员账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号。
	Identity *string `json:"identity,omitempty"`
	// IdentityType 是账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式。
	IdentityType *string `json:"identity_type,omitempty"`
	// Name 是姓名，账号类型为ALIPAY_LOGON_ID时必填。
	Name *string `json:"name,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// Status 是成员当前状态： 邀请中（PROCESSING）、正常（NORMAL）。
	Status *string `json:"status,omitempty"`
}

// FundJointAccountMemberRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountMemberRuleModifyRequest struct {
	// AccountId 是账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是业务场景。
	BizScene   *string                 `json:"biz_scene,omitempty"`
	MemberList *JointAccountMemberList `json:"member_list,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type FundJointAccountRuleModifyRequest struct {
	// AccountId 是账本ID。
	AccountId *string `json:"account_id,omitempty"`
	// AccountName 是账本名称。
	AccountName *string `json:"account_name,omitempty"`
	// AccountQuota 是账户额度。
	AccountQuota *[]JointAccountQuotaDTO `json:"account_quota,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo    *string            `json:"agreement_no,omitempty"`
	AuthorizedRule *AuthorizedRuleDTO `json:"authorized_rule,omitempty"`
	// BizScene 是场景码。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundJointAccountBillQueryRequest 是 alipay.fund.jointaccount.bill.query 的 V3 query 请求参数。
type FundJointAccountBillQueryRequest struct {
	// ProductCode 是因公付产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// PageNum 是采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据。
	PageNum *string `query:"page_num,omitempty"`
	// PageSize 是每页大小，不传的情况下默认20条，超过20条默认按20条查询；不足20条则按实际记录数返回。
	PageSize *string `query:"page_size,omitempty"`
	// UserId 是员工ID。
	UserId *string `query:"user_id,omitempty"`
	// OpenId 是员工的openId。
	OpenId *string `query:"open_id,omitempty"`
	// StartDate 是起始日期。
	StartDate *string `query:"start_date,omitempty"`
	// EndDate 是截止日期。
	EndDate *string `query:"end_date,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
}

// FundJointAccountDetailQueryRequest 是 alipay.fund.jointaccount.detail.query 的 V3 query 请求参数。
type FundJointAccountDetailQueryRequest struct {
	// ProductCode 是销售产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景码。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是合花群ID（查询方式二：通过传入account_id+agreement_no查询）<br> 补充说明：<br> -该字段可在签约接口alipay.fund.jointaccount.sign调用后，由账户变更结果通知接口alipay.fund.jointaccount.account.completed中返回。<br> -该字段可在签约接口调用后，由查询账户详情接口alipay.fund.jointaccount.detail.query中返回。。
	AccountId *string `query:"account_id,omitempty"`
	// AgreementNo 是授权协议号（查询方式二：通过传入account_id+agreement_no查询）<br> 补充说明：<br> -该字段可在签约接口alipay.fund.jointaccount.sign调用后，由账户变更结果通知接口alipay.fund.jointaccount.account.completed中返回。<br> -该字段可在签约接口调用后，由查询账户详情接口alipay.fund.jointaccount.detail.query中返回。。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// OutBizNo 是商户侧单号（查询方式一：通过传入 out_biz_no查询）<br> 补充说明：<br> -该字段与签约接口alipay.fund.jointaccount.sign请求参数中传入值保持一致。
	OutBizNo *string `query:"out_biz_no,omitempty"`
	// Channel 是业务渠道</br> 补充说明：</br> 需要商户侧必传，支付宝侧不做强校验。
	Channel *string `query:"channel,omitempty"`
}

// FundJointAccountListQueryRequest 是 alipay.fund.jointaccount.list.query 的 V3 query 请求参数。
type FundJointAccountListQueryRequest struct {
	// ProductCode 是产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `query:"biz_scene,omitempty"`
	// OperateRole 是角色：创建方(CREATOR)、参与方(PARTICIPANT)。
	OperateRole *string `query:"operate_role,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// Identity 是员工账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号。
	Identity *string `query:"identity,omitempty"`
	// IdentityType 是账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式。
	IdentityType *string `query:"identity_type,omitempty"`
}

// FundJointAccountMemberQueryRequest 是 alipay.fund.jointaccount.member.query 的 V3 query 请求参数。
type FundJointAccountMemberQueryRequest struct {
	// ProductCode 是产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// PageNum 是采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据。。
	PageNum *string `query:"page_num,omitempty"`
	// PageSize 是每页大小，不传的情况下默认20条，超过20条默认按20条查询；不足20条则按实际记录数返回。
	PageSize *string `query:"page_size,omitempty"`
	// UserId 是员工id。
	UserId *string `query:"user_id,omitempty"`
	// OpenId 是员工的openId。
	OpenId *string `query:"open_id,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
}

// FundJointAccountQuotaQueryRequest 是 alipay.fund.jointaccount.quota.query 的 V3 query 请求参数。
type FundJointAccountQuotaQueryRequest struct {
	// ProductCode 是产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// MemberId 是员工ID。
	MemberId *string `query:"member_id,omitempty"`
	// MemberOpenId 是员工的openid。
	MemberOpenId *string `query:"member_open_id,omitempty"`
	// OperateRole 是角色：创建方(CREATOR)、参与方(PARTICIPANT)。
	OperateRole *string `query:"operate_role,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
}

// FundJointAccountSignQueryRequest 是 alipay.fund.jointaccount.sign.query 的 V3 query 请求参数。
type FundJointAccountSignQueryRequest struct {
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene *string `query:"biz_scene,omitempty"`
	// OutBizNo 是外部业务号，查询对应账户状态，优先级高于account_id。
	OutBizNo *string `query:"out_biz_no,omitempty"`
	// AccountId 是企业签约账户ID,用于外部商户已获取企业签约ID，查询账户状态。
	AccountId *string `query:"account_id,omitempty"`
}

// FundJointAccountTradeQueryRequest 是 alipay.fund.jointaccount.trade.query 的 V3 query 请求参数。
type FundJointAccountTradeQueryRequest struct {
	// ProductCode 是销售产品码。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `query:"biz_scene,omitempty"`
	// MemberId 是成员ID，消费发起人。
	MemberId *string `query:"member_id,omitempty"`
	// MemberOpenId 是消费发起人的openId。
	MemberOpenId *string `query:"member_open_id,omitempty"`
	// AccountId 是企业账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// AgreementNo 是三方授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// TradeNo 是支付宝交易订单号。
	TradeNo *string `query:"trade_no,omitempty"`
	// PlatformOrderId 是外部平台订单号。使用该参数查询时，返回结果见trade_info_list。
	PlatformOrderId *string `query:"platform_order_id,omitempty"`
}
