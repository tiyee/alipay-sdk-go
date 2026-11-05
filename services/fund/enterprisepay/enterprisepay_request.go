package enterprisepay

import (
	"github.com/tiyee/alipay-sdk-go/services/fund"
)

// FundEnterprisePayGroupAddRequest 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupAddRequest struct {
	// AccountId 是企业签约账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是平台和企业的三方授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene    *string           `json:"biz_scene,omitempty"`
	FundExtInfo *fund.FundExtInfo `json:"fund_ext_info,omitempty"`
	// Name 是群组名称。
	Name *string `json:"name,omitempty"`
	// OutBizNo 是外部业务号，外部群组号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundEnterprisePayGroupModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupModifyRequest struct {
	// AccountId 是企业签约账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是平台和企业的三方授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene    *string           `json:"biz_scene,omitempty"`
	FundExtInfo *fund.FundExtInfo `json:"fund_ext_info,omitempty"`
	// Name 是群组名称。
	Name *string `json:"name,omitempty"`
	// OperationTypeList 是更新操作类型：多个时逗号分隔 NAME-群组名称 FUND-出资信息 字段不传时，默认全量更新。
	OperationTypeList *[]string `json:"operation_type_list,omitempty"`
	// OutBizNo 是外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundEnterprisePayMemberModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayMemberModifyRequest struct {
	// AccountId 是企业签约账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是平台和企业的三方授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene    *string           `json:"biz_scene,omitempty"`
	FundExtInfo *fund.FundExtInfo `json:"fund_ext_info,omitempty"`
	// GroupIdList 是员工当前关联的群组ID，支持多个。
	GroupIdList *[]string `json:"group_id_list,omitempty"`
	// OpenId 是蚂蚁统一会员openId。
	OpenId *string `json:"open_id,omitempty"`
	// OperationTypeList 是更新操作类型：多个时逗号分隔 GROUP-群组信息 FUND-出资信息 字段不传时，默认全量更新。
	OperationTypeList *[]string `json:"operation_type_list,omitempty"`
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `json:"product_code,omitempty"`
	// UserId 是蚂蚁统一会员ID。
	UserId *string `json:"user_id,omitempty"`
}

// FundEnterprisePaySignRequest 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePaySignRequest struct {
	// AccountName 是企业简称，传空采用默认规则命名，使用公司名称。
	AccountName *string `json:"account_name,omitempty"`
	// AppointSignTerminal 是指定签约终端类型 PC-电脑浏览器 MOBILE-手机端。
	AppointSignTerminal *string `json:"appoint_sign_terminal,omitempty"`
	// BizScene 是业务场景，联系支付宝分配。
	BizScene *string `json:"biz_scene,omitempty"`
	// ExtParams 是特殊场景下扩展字段。
	ExtParams *string `json:"ext_params,omitempty"`
	// Identity 是开户账号： 当identity_type是ALIPAY_USER_ID时填支付宝会员ID（2088开头）； 当identity_type是ALIPAY_LOGON_ID 时填支付宝登录号； 当identity_type是OUT_USER_ID时填外部平台的用户uid; 当identity_type是ALIPAY_OPEN_ID时填用户openId。
	Identity *string `json:"identity,omitempty"`
	// IdentityName 是开户用户名称，identity_type为ALIPAY_LOGON_ID时必填。
	IdentityName *string `json:"identity_name,omitempty"`
	// IdentityType 是账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、OUT_USER_ID：外部用户uid。
	IdentityType *string `json:"identity_type,omitempty"`
	// OutBizNo 是商户侧单号（幂等字段） ，补充说明：幂等逻辑（商户appid+out_biz_no），商户集成时需注意，如换号重复发起，则认为是一笔新的请求。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// OutEntityId 是外部实体ID，一般表述为外部操作人操作某个外部实体。比如企业码2284号。
	OutEntityId *string `json:"out_entity_id,omitempty"`
	// OutSource 是ISV平台定义的共同账户开户来源。
	OutSource *string `json:"out_source,omitempty"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundEnterprisePayUnsignRequest 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayUnsignRequest struct {
	// AccountId 是企业账号。
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是业务场景。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是销售产品码。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundEnterprisePayGroupDeleteRequest 是 alipay.fund.enterprisepay.group.delete 的 V3 query 请求参数。
type FundEnterprisePayGroupDeleteRequest struct {
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是企业签约账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// AgreementNo 是平台和企业的三方授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// OutBizNo 是外部业务号。
	OutBizNo *string `query:"out_biz_no,omitempty"`
}

// FundEnterprisePayGroupQueryRequest 是 alipay.fund.enterprisepay.group.query 的 V3 query 请求参数。
type FundEnterprisePayGroupQueryRequest struct {
	// ProductCode 是产品码，默认值 ENTERPRISE_PAY。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是场景码，联系支付宝分配。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountId 是企业签约账户ID。
	AccountId *string `query:"account_id,omitempty"`
	// AgreementNo 是平台和企业的三方授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// OutBizNo 是外部业务号，外部群组号。
	OutBizNo *string `query:"out_biz_no,omitempty"`
}
