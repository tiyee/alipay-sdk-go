package jointaccount

import (
	"github.com/tiyee/alipay-sdk-go/services/fund"
)

// AuthorizeDetailDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthorizeDetailDTO struct {
	// AuthorizedPrincipalId 是收单主体PID。
	AuthorizedPrincipalId *string `json:"authorized_principal_id,omitempty"`
}

// AuthorizedRuleDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthorizedRuleDTO struct {
	// AuthorizedDetailList 是授权商户列表。
	AuthorizedDetailList *[]AuthorizeDetailDTO `json:"authorized_detail_list,omitempty"`
	// AuthorizedType 是WHITELIST-白名单 ALL-全部。
	AuthorizedType *string `json:"authorized_type,omitempty"`
}

// IdentityVerifiedInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type IdentityVerifiedInfoDTO struct {
	// CertNo 是MD5加密处理后的证件号。
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是证件类型。暂仅支持 IDENTITY_CARD （身份证）。
	CertType *string `json:"cert_type,omitempty"`
	// UserName 是用户的真实姓名。
	UserName *string `json:"user_name,omitempty"`
}

// InviteMemberBusinessParamsDTO 是支付宝官方 V3 OpenAPI 模型。
type InviteMemberBusinessParamsDTO struct {
	// EmployeeId 是企业员工工卡。
	EmployeeId  *string           `json:"employee_id,omitempty"`
	FundExtInfo *fund.FundExtInfo `json:"fund_ext_info,omitempty"`
	// GroupIdList 是员工关联的群组ID，当前系统仅支持一个。
	GroupIdList *[]string `json:"group_id_list,omitempty"`
}

// InviteResultDTO 是支付宝官方 V3 OpenAPI 模型。
type InviteResultDTO struct {
	// Identity 是（被邀请人）用户唯一标识<br> 补充说明：<br> - 该字段与签约接口alipay.fund.jointaccount.sign请求参数中传入值保持一致。
	Identity *string `json:"identity,omitempty"`
	// IdentityType 是（被邀请人）账号类型，参考值如下：<br> - OUT_USER_ID：商户侧用户唯一标识<br> 补充说明： - 该字段与签约接口alipay.fund.jointaccount.sign请求参数中传入值保持一致。
	IdentityType *string `json:"identity_type,omitempty"`
	// ProcessId 是审批单号，不同被邀请人之间的审批单号不同（支付宝侧生成）。
	ProcessId *string `json:"process_id,omitempty"`
}

// JointAccountBillDetailDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountBillDetailDTO struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// Amount 是消费金额。
	Amount *string `json:"amount,omitempty"`
	// BillNo 是账单业务号。
	BillNo *string `json:"bill_no,omitempty"`
	// BizDate 是业务时间。
	BizDate *string `json:"biz_date,omitempty"`
	// BizNo 是订单号。
	BizNo *string `json:"biz_no,omitempty"`
	// InOut 是1-退款，2-支付。
	InOut *string `json:"in_out,omitempty"`
	// OpenId 是共同账户消费的成员openid。
	OpenId *string `json:"open_id,omitempty"`
	// OutTradeNo 是外部交易单号，正向支付为外部交易单号，逆向退款是为外部退款单号。
	OutTradeNo     *string          `json:"out_trade_no,omitempty"`
	PayerAssetInfo *UserAssetInfoVO `json:"payer_asset_info,omitempty"`
	// SellerFullName 是间连商户显示二级商户全名，直连、直付通则显示一级商户全名。企业商户该字段不脱敏，非企业商户该字段会脱敏。
	SellerFullName *string `json:"seller_full_name,omitempty"`
	// SellerLogonId 是收款方登录号信息。间连商户显示二级商户的登录号信息；直连、直付通显示一级商户的登录号信息；该字段脱敏。
	SellerLogonId *string `json:"seller_logon_id,omitempty"`
	// Title 是账单标题。
	Title *string `json:"title,omitempty"`
	// UserId 是共同账户消费的成员ID。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountDTO struct {
	// AccountId 是共同账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AccountName 是共同账户账本名称。
	AccountName *string `json:"account_name,omitempty"`
	// UserId 是账本创建人会员号。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountMemberDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountMemberDTO struct {
	AccountQuota  *JointAccountQuotaDTO `json:"account_quota,omitempty"`
	MemberExtInfo *MemberExtInfo        `json:"member_ext_info,omitempty"`
	// Name 是姓名，脱敏输出。
	Name *string `json:"name,omitempty"`
	// OpenId 是用户openId。
	OpenId *string `json:"open_id,omitempty"`
	// Status 是成员状态。
	Status *string `json:"status,omitempty"`
	// UserId 是用户会员号。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountMemberInfoRespDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountMemberInfoRespDTO struct {
	// OpenId 是（群成员）支付宝侧用户唯一标识。
	OpenId *string `json:"open_id,omitempty"`
	// OperateRole 是成员角色：<br> -MASTER：创建人<br> -ADMIN：管理员<br> -MEMBER：群成员<br>。
	OperateRole *string `json:"operate_role,omitempty"`
	// UserId 是（群成员）支付宝侧用户唯一标识。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountMemberList 是支付宝官方 V3 OpenAPI 模型。
type JointAccountMemberList struct {
	// AccountQuota 是额度模型。
	AccountQuota *[]JointAccountQuotaDTO `json:"account_quota,omitempty"`
	// OpenId 是成员openId。
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是成员ID。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountMemberRespDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountMemberRespDTO struct {
	// AccountQuota 是额度模型。
	AccountQuota *[]JointAccountQuotaDTO `json:"account_quota,omitempty"`
	// LogonId 是成员支付宝登录号。
	LogonId *string `json:"logon_id,omitempty"`
	// Name 是姓名。
	Name *string `json:"name,omitempty"`
	// Status 是员工当前状态： 邀请中（PROCESSING）、正常（NORMAL）。
	Status *string `json:"status,omitempty"`
	// UserId 是用户支付宝会员号。
	UserId *string `json:"user_id,omitempty"`
}

// JointAccountQuotaDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountQuotaDTO struct {
	// CustomBeginDate 是自定义周期起始日期，精确到分钟，yyyy-MM-dd HH:mm。
	CustomBeginDate *string `json:"custom_begin_date,omitempty"`
	// CustomEndDate 是自定义周期结束日期，精确到分钟，yyyy-MM-dd HH:mm。
	CustomEndDate *string `json:"custom_end_date,omitempty"`
	// QuotaDimension 是额度维度 单笔上限：ONCE 日：DAY 月：MONTH 季度：QUARTER 年：YEAR 终身：LIFETIME 一次性使用额度：ONLYONE。
	QuotaDimension *string `json:"quota_dimension,omitempty"`
	// QuotaTotal 是协议额度。
	QuotaTotal *string `json:"quota_total,omitempty"`
}

// JointAccountQuotaRespDTO 是支付宝官方 V3 OpenAPI 模型。
type JointAccountQuotaRespDTO struct {
	// CustomBeginDate 是额度生效起始日期，精确到分钟。 格式：yyyy-MM-dd HH:mm。
	CustomBeginDate *string `json:"custom_begin_date,omitempty"`
	// CustomEndDate 是额度失效结束日期，精确到分钟。 格式：yyyy-MM-dd HH:mm。
	CustomEndDate *string `json:"custom_end_date,omitempty"`
	// QuotaDimension 是额度维度。
	QuotaDimension *string `json:"quota_dimension,omitempty"`
	// QuotaRemain 是剩余额度。
	QuotaRemain *string `json:"quota_remain,omitempty"`
	// QuotaRemainCount 是展示额度剩余可用次数，-1表示无限次。
	QuotaRemainCount *string `json:"quota_remain_count,omitempty"`
	// QuotaTotal 是协议额度。
	QuotaTotal *string `json:"quota_total,omitempty"`
	// QuotaUsed 是已用额度。
	QuotaUsed *string `json:"quota_used,omitempty"`
}

// MemberExtInfo 是支付宝官方 V3 OpenAPI 模型。
type MemberExtInfo struct {
	// EmployeeId 是员工工号信息。
	EmployeeId *string `json:"employee_id,omitempty"`
}

// MemberStatusRespDTO 是支付宝官方 V3 OpenAPI 模型。
type MemberStatusRespDTO struct {
	// OpenId 是用户支付宝会员openid。
	OpenId *string `json:"open_id,omitempty"`
	// Status 是员工当前状态： 邀请中（PROCESSING）、正常（NORMAL）、不在企业中（OUT）。
	Status *string `json:"status,omitempty"`
	// UserId 是用户支付宝2088会员号。
	UserId *string `json:"user_id,omitempty"`
}

// TradeFundBillDetail 是支付宝官方 V3 OpenAPI 模型。
type TradeFundBillDetail struct {
	// Amount 是交易金额。
	Amount *string `json:"amount,omitempty"`
	// AssetTypeCode 是资产类型编码。
	AssetTypeCode *string `json:"asset_type_code,omitempty"`
	// AssetUserId 是资产的属主。
	AssetUserId *string `json:"asset_user_id,omitempty"`
	// AssetUserOpenId 是资产的属主openid。
	AssetUserOpenId *string `json:"asset_user_open_id,omitempty"`
	// BizPayType 是支付业务类型 PAY-支付 REFUND-退款。
	BizPayType *string `json:"biz_pay_type,omitempty"`
	// CreateTime 是创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// PaymentNo 是支付单据号。
	PaymentNo *string `json:"payment_no,omitempty"`
}

// TradeInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type TradeInfoDTO struct {
	// BuyerId 是买家ID。
	BuyerId *string `json:"buyer_id,omitempty"`
	// CreateTime 是交易创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// PlatformOrderId 是外部平台订单号。
	PlatformOrderId *string `json:"platform_order_id,omitempty"`
	// TotalAmount 是订单总金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// TradeAmount 是订单总金额。
	TradeAmount *string `json:"trade_amount,omitempty"`
	// TradeFundBillList 是资金单明细。
	TradeFundBillList *[]TradeFundBillDetail `json:"trade_fund_bill_list,omitempty"`
	// TradeNo 是交易单号。
	TradeNo *string `json:"trade_no,omitempty"`
	// TradeStatus 是交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）。
	TradeStatus *string `json:"trade_status,omitempty"`
}

// UserAssetInfoVO 是支付宝官方 V3 OpenAPI 模型。
type UserAssetInfoVO struct {
	// AssetId 是资产ID。
	AssetId *string `json:"asset_id,omitempty"`
	// AssetType 是资产类型。
	AssetType *string `json:"asset_type,omitempty"`
	// UserId 是支付宝用户ID。
	UserId *string `json:"user_id,omitempty"`
	// UserOpenId 是支付宝用户开放ID。
	UserOpenId *string `json:"user_open_id,omitempty"`
}
