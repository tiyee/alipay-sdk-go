# jointaccount 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/jointaccount`。本包提供 13 个支付宝 V3 REST 操作。

通过 `jointaccount.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`JointAccountBillQuery`](#api-jointaccountbillquery) | `alipay.fund.jointaccount.bill.query` |
| [`JointAccountDetailQuery`](#api-jointaccountdetailquery) | `alipay.fund.jointaccount.detail.query` |
| [`JointAccountListQuery`](#api-jointaccountlistquery) | `alipay.fund.jointaccount.list.query` |
| [`JointAccountMemberBatchQuery`](#api-jointaccountmemberbatchquery) | `alipay.fund.jointaccount.member.batchquery` |
| [`JointAccountMemberBind`](#api-jointaccountmemberbind) | `alipay.fund.jointaccount.member.bind` |
| [`JointAccountMemberConsult`](#api-jointaccountmemberconsult) | `alipay.fund.jointaccount.member.consult` |
| [`JointAccountMemberQuery`](#api-jointaccountmemberquery) | `alipay.fund.jointaccount.member.query` |
| [`JointAccountMemberUnbind`](#api-jointaccountmemberunbind) | `alipay.fund.jointaccount.member.unbind` |
| [`JointAccountMemberRuleModify`](#api-jointaccountmemberrulemodify) | `alipay.fund.jointaccount.memberrule.modify` |
| [`JointAccountQuotaQuery`](#api-jointaccountquotaquery) | `alipay.fund.jointaccount.quota.query` |
| [`JointAccountRuleModify`](#api-jointaccountrulemodify) | `alipay.fund.jointaccount.rule.modify` |
| [`JointAccountSignQuery`](#api-jointaccountsignquery) | `alipay.fund.jointaccount.sign.query` |
| [`JointAccountTradeQuery`](#api-jointaccounttradequery) | `alipay.fund.jointaccount.trade.query` |

## 使用示例

- [JointAccountBillQuery](api_jointaccount_jointaccountbillquery_example_test.go)
- [JointAccountDetailQuery](api_jointaccount_jointaccountdetailquery_example_test.go)
- [JointAccountListQuery](api_jointaccount_jointaccountlistquery_example_test.go)
- [JointAccountMemberBatchQuery](api_jointaccount_jointaccountmemberbatchquery_example_test.go)
- [JointAccountMemberBind](api_jointaccount_jointaccountmemberbind_example_test.go)
- [JointAccountMemberConsult](api_jointaccount_jointaccountmemberconsult_example_test.go)
- [JointAccountMemberQuery](api_jointaccount_jointaccountmemberquery_example_test.go)
- [JointAccountMemberUnbind](api_jointaccount_jointaccountmemberunbind_example_test.go)
- [JointAccountMemberRuleModify](api_jointaccount_jointaccountmemberrulemodify_example_test.go)
- [JointAccountQuotaQuery](api_jointaccount_jointaccountquotaquery_example_test.go)
- [JointAccountRuleModify](api_jointaccount_jointaccountrulemodify_example_test.go)
- [JointAccountSignQuery](api_jointaccount_jointaccountsignquery_example_test.go)
- [JointAccountTradeQuery](api_jointaccount_jointaccounttradequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-jointaccountbillquery"></a>

### JointAccountBillQuery — 因公付账单查询

- 官方接口：[`alipay.fund.jointaccount.bill.query`](https://opendocs.alipay.com/apis/05wuhp)
- 路由：`GET /v3/alipay/fund/jointaccount/bill/query`
- 请求类型：[FundJointAccountBillQueryRequest](jointaccount_request.go#L115)
- 响应类型：[FundJointAccountBillQueryResponse](jointaccount_response.go#L4)

因公付账单查询

请求约束：FundJointAccountBillQueryRequest 是 alipay.fund.jointaccount.bill.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 因公付产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景 |
| `account_id` | `AccountId` | `*string` | query | 否 | 账户ID |
| `page_num` | `PageNum` | `*string` | query | 否 | 采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页大小，不传的情况下默认20条，超过20条默认按20条查询；不足20条则按实际记录数返回 |
| `user_id` | `UserId` | `*string` | query | 否 | 员工ID |
| `open_id` | `OpenId` | `*string` | query | 否 | 员工的openId |
| `start_date` | `StartDate` | `*string` | query | 否 | 起始日期 |
| `end_date` | `EndDate` | `*string` | query | 否 | 截止日期 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权协议号 |

<a id="api-jointaccountdetailquery"></a>

### JointAccountDetailQuery — 共享资金-查询账户详情

- 官方接口：[`alipay.fund.jointaccount.detail.query`](https://opendocs.alipay.com/apis/05imoi)
- 路由：`GET /v3/alipay/fund/jointaccount/detail/query`
- 请求类型：[FundJointAccountDetailQueryRequest](jointaccount_request.go#L139)
- 响应类型：[FundJointAccountDetailQueryResponse](jointaccount_response.go#L20)

发起人在商户上点击查看账户详情，商户服务端进行受理，并调用支付宝进行处理。处理结果，同步返回给商户。

请求约束：FundJointAccountDetailQueryRequest 是 alipay.fund.jointaccount.detail.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 销售产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景码 |
| `account_id` | `AccountId` | `*string` | query | 否 | 合花群ID（查询方式二：通过传入account_id+agreement_no查询）<br> 补充说明：<br> -该字段可在签约接口alipay.fund.jointaccount.sign调用后，由账户变更结果通知接口alipay.fund.jointaccount.account.completed中返回。<br> -该字段可在签约接口调用后，由查询账户详情接口alipay.fund.jointaccount.detail.query中返回。 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权协议号（查询方式二：通过传入account_id+agreement_no查询）<br> 补充说明：<br> -该字段可在签约接口alipay.fund.jointaccount.sign调用后，由账户变更结果通知接口alipay.fund.jointaccount.account.completed中返回。<br> -该字段可在签约接口调用后，由查询账户详情接口alipay.fund.jointaccount.detail.query中返回。 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 商户侧单号（查询方式一：通过传入 out_biz_no查询）<br> 补充说明：<br> -该字段与签约接口alipay.fund.jointaccount.sign请求参数中传入值保持一致 |
| `channel` | `Channel` | `*string` | query | 否 | 业务渠道 补充说明： 需要商户侧必传，支付宝侧不做强校验 |

<a id="api-jointaccountlistquery"></a>

### JointAccountListQuery — 企业查询代付账户列表

- 官方接口：[`alipay.fund.jointaccount.list.query`](https://opendocs.alipay.com/apis/05wcqw)
- 路由：`GET /v3/alipay/fund/jointaccount/list/query`
- 请求类型：[FundJointAccountListQueryRequest](jointaccount_request.go#L155)
- 响应类型：[FundJointAccountListQueryResponse](jointaccount_response.go#L57)

企业查询代付账户列表

请求约束：FundJointAccountListQueryRequest 是 alipay.fund.jointaccount.list.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景 |
| `operate_role` | `OperateRole` | `*string` | query | 否 | 角色：创建方(CREATOR)、参与方(PARTICIPANT) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权协议号 |
| `identity` | `Identity` | `*string` | query | 否 | 员工账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号 |
| `identity_type` | `IdentityType` | `*string` | query | 否 | 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 |

<a id="api-jointaccountmemberbatchquery"></a>

### JointAccountMemberBatchQuery — 企业批量查询员工信息（分页）

- 官方接口：[`alipay.fund.jointaccount.member.batchquery`](https://opendocs.alipay.com/apis/05wfae)
- 路由：`POST /v3/alipay/fund/jointaccount/member/batchquery`
- 请求类型：[FundJointAccountMemberBatchQueryRequest](jointaccount_request.go#L4)
- 响应类型：[FundJointAccountMemberBatchQueryResponse](jointaccount_response.go#L67)

企业查询员工列表，支持分页，支持按登录号精确查询，支持按会员号精确查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 三方授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景 |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 企业账号：identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号 |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 |
| `last_open_id` | `LastOpenId` | `*string` | JSON body | 未标注 | 按游标查询，传入上次查询最后一个用户的openId |
| `last_user_id` | `LastUserId` | `*string` | JSON body | 未标注 | 按游标查询，传入上次查询最后一个用户ID，适用于逐页查询。优先级大于page_num |
| `page_num` | `PageNum` | `*string` | JSON body | 未标注 | 采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据。 |
| `page_size` | `PageSize` | `*string` | JSON body | 未标注 | 每页大小，不传的情况下默认20条，上限100；不足20条则按实际记录数返回 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码 |

<a id="api-jointaccountmemberbind"></a>

### JointAccountMemberBind — 因公付邀请员工

- 官方接口：[`alipay.fund.jointaccount.member.bind`](https://opendocs.alipay.com/apis/05wvt3)
- 路由：`POST /v3/alipay/fund/jointaccount/member/bind`
- 请求类型：[FundJointAccountMemberBindRequest](jointaccount_request.go#L28)
- 响应类型：[FundJointAccountMemberBindResponse](jointaccount_response.go#L83)

因公付邀请员工

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 账本ID |
| `account_quota` | `AccountQuota` | [`*JointAccountQuotaDTO`](#model-jointaccountquotadto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `agreement_sign_info` | `AgreementSignInfo` | `*string` | JSON body | 未标注 | 关联代扣协议签约串 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景，联系支付宝分配 |
| `business_params` | `BusinessParams` | [`*InviteMemberBusinessParamsDTO`](#model-invitememberbusinessparamsdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 该字段为被邀请方账号ID： 当identity_type为ALIPAY_USER_ID时，需要填支付宝ID；当identity_type为ALIPAY_LOGON_ID时，需要填支付宝登录号；当identity_type为ALIPAY_OPEN_ID时，需要填支付宝openId |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、ALIPAY_OPEN_ID：支付宝openId |
| `identity_verified_info` | `IdentityVerifiedInfo` | [`*IdentityVerifiedInfoDTO`](#model-identityverifiedinfodto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 成员姓名，账号类型为ALIPAY_LOGON_ID时必填 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |

<a id="api-jointaccountmemberconsult"></a>

### JointAccountMemberConsult — 员企关系咨询

- 官方接口：[`alipay.fund.jointaccount.member.consult`](https://opendocs.alipay.com/apis/05wupc)
- 路由：`POST /v3/alipay/fund/jointaccount/member/consult`
- 请求类型：[FundJointAccountMemberConsultRequest](jointaccount_request.go#L51)
- 响应类型：[FundJointAccountMemberConsultResponse](jointaccount_response.go#L89)

返回输入的员工列表是否在指定企业组织中

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 协议号 |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | 扩展参数，如需使用请联系支付宝开放同学 |
| `open_id_list` | `OpenIdList` | `*[]string` | JSON body | 未标注 | 要检查的员工openId |
| `user_id_list` | `UserIdList` | `*[]string` | JSON body | 未标注 | 要检查的员工ID |

<a id="api-jointaccountmemberquery"></a>

### JointAccountMemberQuery — 企业查询员工列表

- 官方接口：[`alipay.fund.jointaccount.member.query`](https://opendocs.alipay.com/apis/05wujm)
- 路由：`GET /v3/alipay/fund/jointaccount/member/query`
- 请求类型：[FundJointAccountMemberQueryRequest](jointaccount_request.go#L171)
- 响应类型：[FundJointAccountMemberQueryResponse](jointaccount_response.go#L101)

企业查询员工列表

请求约束：FundJointAccountMemberQueryRequest 是 alipay.fund.jointaccount.member.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景 |
| `account_id` | `AccountId` | `*string` | query | 否 | 账户ID |
| `page_num` | `PageNum` | `*string` | query | 否 | 采用分页查询，本参数为空或0默认显示第一页。如果输入的值大于总页数，则支付宝返回最后一页数据。 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页大小，不传的情况下默认20条，超过20条默认按20条查询；不足20条则按实际记录数返回 |
| `user_id` | `UserId` | `*string` | query | 否 | 员工id |
| `open_id` | `OpenId` | `*string` | query | 否 | 员工的openId |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权协议号 |

<a id="api-jointaccountmemberunbind"></a>

### JointAccountMemberUnbind — 企业删除员工

- 官方接口：[`alipay.fund.jointaccount.member.unbind`](https://opendocs.alipay.com/apis/05x1fk)
- 路由：`POST /v3/alipay/fund/jointaccount/member/unbind`
- 请求类型：[FundJointAccountMemberUnbindRequest](jointaccount_request.go#L65)
- 响应类型：[FundJointAccountMemberUnbindResponse](jointaccount_response.go#L119)

企业删除员工

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 账本id |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码 |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 成员账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号 |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 姓名，账号类型为ALIPAY_LOGON_ID时必填 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |
| `status` | `Status` | `*string` | JSON body | 未标注 | 成员当前状态： 邀请中（PROCESSING）、正常（NORMAL） |

<a id="api-jointaccountmemberrulemodify"></a>

### JointAccountMemberRuleModify — 企业设置员工规则

- 官方接口：[`alipay.fund.jointaccount.memberrule.modify`](https://opendocs.alipay.com/apis/05wygg)
- 路由：`POST /v3/alipay/fund/jointaccount/memberrule/modify`
- 请求类型：[FundJointAccountMemberRuleModifyRequest](jointaccount_request.go#L85)
- 响应类型：[FundJointAccountMemberRuleModifyResponse](jointaccount_response.go#L123)

企业设置员工规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景 |
| `member_list` | `MemberList` | [`*JointAccountMemberList`](#model-jointaccountmemberlist) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |

<a id="api-jointaccountquotaquery"></a>

### JointAccountQuotaQuery — 查询账户额度详情

- 官方接口：[`alipay.fund.jointaccount.quota.query`](https://opendocs.alipay.com/apis/05wuhk)
- 路由：`GET /v3/alipay/fund/jointaccount/quota/query`
- 请求类型：[FundJointAccountQuotaQueryRequest](jointaccount_request.go#L191)
- 响应类型：[FundJointAccountQuotaQueryResponse](jointaccount_response.go#L127)

查询账户额度详情

请求约束：FundJointAccountQuotaQueryRequest 是 alipay.fund.jointaccount.quota.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景 |
| `account_id` | `AccountId` | `*string` | query | 否 | 账户ID |
| `member_id` | `MemberId` | `*string` | query | 否 | 员工ID |
| `member_open_id` | `MemberOpenId` | `*string` | query | 否 | 员工的openid |
| `operate_role` | `OperateRole` | `*string` | query | 否 | 角色：创建方(CREATOR)、参与方(PARTICIPANT) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权协议号 |

<a id="api-jointaccountrulemodify"></a>

### JointAccountRuleModify — 企业账户规则设置

- 官方接口：[`alipay.fund.jointaccount.rule.modify`](https://opendocs.alipay.com/apis/05wghc)
- 路由：`POST /v3/alipay/fund/jointaccount/rule/modify`
- 请求类型：[FundJointAccountRuleModifyRequest](jointaccount_request.go#L98)
- 响应类型：[FundJointAccountRuleModifyResponse](jointaccount_response.go#L143)

企业账户规则设置

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 账本ID |
| `account_name` | `AccountName` | `*string` | JSON body | 未标注 | 账本名称 |
| `account_quota` | `AccountQuota` | [`*[]JointAccountQuotaDTO`](#model-jointaccountquotadto) | JSON body | 未标注 | 账户额度 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权协议号 |
| `authorized_rule` | `AuthorizedRule` | [`*AuthorizedRuleDTO`](#model-authorizedruledto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |

<a id="api-jointaccountsignquery"></a>

### JointAccountSignQuery — 企业签约结果查询

- 官方接口：[`alipay.fund.jointaccount.sign.query`](https://opendocs.alipay.com/apis/05wulj)
- 路由：`GET /v3/alipay/fund/jointaccount/sign/query`
- 请求类型：[FundJointAccountSignQueryRequest](jointaccount_request.go#L209)
- 响应类型：[FundJointAccountSignQueryResponse](jointaccount_response.go#L147)

通过外部订单号（out_biz_no）主动查询企业签约结果，可用于未收到签约结果通知的补偿操作

请求约束：FundJointAccountSignQueryRequest 是 alipay.fund.jointaccount.sign.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码，默认值 ENTERPRISE_PAY |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 场景码，联系支付宝分配 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 外部业务号，查询对应账户状态，优先级高于account_id |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业签约账户ID,用于外部商户已获取企业签约ID，查询账户状态 |

<a id="api-jointaccounttradequery"></a>

### JointAccountTradeQuery — 共同账户交易查询

- 官方接口：[`alipay.fund.jointaccount.trade.query`](https://opendocs.alipay.com/apis/05wykp)
- 路由：`GET /v3/alipay/fund/jointaccount/trade/query`
- 请求类型：[FundJointAccountTradeQueryRequest](jointaccount_request.go#L221)
- 响应类型：[FundJointAccountTradeQueryResponse](jointaccount_response.go#L157)

共同账户交易查询

请求约束：FundJointAccountTradeQueryRequest 是 alipay.fund.jointaccount.trade.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 销售产品码 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 业务场景 |
| `member_id` | `MemberId` | `*string` | query | 否 | 成员ID，消费发起人 |
| `member_open_id` | `MemberOpenId` | `*string` | query | 否 | 消费发起人的openId |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业账户ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 三方授权协议号 |
| `trade_no` | `TradeNo` | `*string` | query | 否 | 支付宝交易订单号 |
| `platform_order_id` | `PlatformOrderId` | `*string` | query | 否 | 外部平台订单号。使用该参数查询时，返回结果见trade_info_list |

## 嵌套请求参数

<a id="model-authorizedetaildto"></a>

### AuthorizeDetailDTO

类型定义：[AuthorizeDetailDTO](jointaccount_model.go#L8)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `authorized_principal_id` | `AuthorizedPrincipalId` | `*string` | 未标注 | 收单主体PID |

<a id="model-authorizedruledto"></a>

### AuthorizedRuleDTO

类型定义：[AuthorizedRuleDTO](jointaccount_model.go#L14)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `authorized_detail_list` | `AuthorizedDetailList` | [`*[]AuthorizeDetailDTO`](#model-authorizedetaildto) | 未标注 | 授权商户列表 |
| `authorized_type` | `AuthorizedType` | `*string` | 未标注 | WHITELIST-白名单 ALL-全部 |

<a id="model-identityverifiedinfodto"></a>

### IdentityVerifiedInfoDTO

类型定义：[IdentityVerifiedInfoDTO](jointaccount_model.go#L22)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cert_no` | `CertNo` | `*string` | 未标注 | MD5加密处理后的证件号 |
| `cert_type` | `CertType` | `*string` | 未标注 | 证件类型。暂仅支持 IDENTITY_CARD （身份证） |
| `user_name` | `UserName` | `*string` | 未标注 | 用户的真实姓名 |

<a id="model-invitememberbusinessparamsdto"></a>

### InviteMemberBusinessParamsDTO

类型定义：[InviteMemberBusinessParamsDTO](jointaccount_model.go#L32)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `employee_id` | `EmployeeId` | `*string` | 未标注 | 企业员工工卡 |
| `fund_ext_info` | `FundExtInfo` | [`*FundExtInfo`](../README.md#model-fundextinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `group_id_list` | `GroupIdList` | `*[]string` | 未标注 | 员工关联的群组ID，当前系统仅支持一个 |

<a id="model-jointaccountmemberlist"></a>

### JointAccountMemberList

类型定义：[JointAccountMemberList](jointaccount_model.go#L114)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `account_quota` | `AccountQuota` | [`*[]JointAccountQuotaDTO`](#model-jointaccountquotadto) | 未标注 | 额度模型 |
| `open_id` | `OpenId` | `*string` | 未标注 | 成员openId |
| `user_id` | `UserId` | `*string` | 未标注 | 成员ID |

<a id="model-jointaccountquotadto"></a>

### JointAccountQuotaDTO

类型定义：[JointAccountQuotaDTO](jointaccount_model.go#L138)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `custom_begin_date` | `CustomBeginDate` | `*string` | 未标注 | 自定义周期起始日期，精确到分钟，yyyy-MM-dd HH:mm |
| `custom_end_date` | `CustomEndDate` | `*string` | 未标注 | 自定义周期结束日期，精确到分钟，yyyy-MM-dd HH:mm |
| `quota_dimension` | `QuotaDimension` | `*string` | 未标注 | 额度维度 单笔上限：ONCE 日：DAY 月：MONTH 季度：QUARTER 年：YEAR 终身：LIFETIME 一次性使用额度：ONLYONE |
| `quota_total` | `QuotaTotal` | `*string` | 未标注 | 协议额度 |
