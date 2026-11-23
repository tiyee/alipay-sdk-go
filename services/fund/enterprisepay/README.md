# enterprisepay 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/enterprisepay`。本包提供 7 个支付宝 V3 REST 操作。

通过 `enterprisepay.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`EnterprisePayGroupAdd`](#api-enterprisepaygroupadd) | `alipay.fund.enterprisepay.group.add` |
| [`EnterprisePayGroupDelete`](#api-enterprisepaygroupdelete) | `alipay.fund.enterprisepay.group.delete` |
| [`EnterprisePayGroupModify`](#api-enterprisepaygroupmodify) | `alipay.fund.enterprisepay.group.modify` |
| [`EnterprisePayGroupQuery`](#api-enterprisepaygroupquery) | `alipay.fund.enterprisepay.group.query` |
| [`EnterprisePayMemberModify`](#api-enterprisepaymembermodify) | `alipay.fund.enterprisepay.member.modify` |
| [`EnterprisePaySign`](#api-enterprisepaysign) | `alipay.fund.enterprisepay.sign` |
| [`EnterprisePayUnsign`](#api-enterprisepayunsign) | `alipay.fund.enterprisepay.unsign` |

## 使用示例

- [EnterprisePayGroupAdd](api_enterprisepay_enterprisepaygroupadd_example_test.go)
- [EnterprisePayGroupDelete](api_enterprisepay_enterprisepaygroupdelete_example_test.go)
- [EnterprisePayGroupModify](api_enterprisepay_enterprisepaygroupmodify_example_test.go)
- [EnterprisePayGroupQuery](api_enterprisepay_enterprisepaygroupquery_example_test.go)
- [EnterprisePayMemberModify](api_enterprisepay_enterprisepaymembermodify_example_test.go)
- [EnterprisePaySign](api_enterprisepay_enterprisepaysign_example_test.go)
- [EnterprisePayUnsign](api_enterprisepay_enterprisepayunsign_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-enterprisepaygroupadd"></a>

### EnterprisePayGroupAdd — 因公付新增账户下群组

- 官方接口：[`alipay.fund.enterprisepay.group.add`](https://opendocs.alipay.com/apis/05wmyu)
- 路由：`POST /v3/alipay/fund/enterprisepay/group/add`
- 请求类型：[FundEnterprisePayGroupAddRequest](enterprisepay_request.go#L8)
- 响应类型：[FundEnterprisePayGroupAddResponse](enterprisepay_response.go#L4)

创建因公付群组，包含群组信息和出资信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业签约账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 平台和企业的三方授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码，联系支付宝分配 |
| `fund_ext_info` | `FundExtInfo` | [`*FundExtInfo`](../README.md#model-fundextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 群组名称 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，外部群组号 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码，默认值 ENTERPRISE_PAY |

<a id="api-enterprisepaygroupdelete"></a>

### EnterprisePayGroupDelete — 因公付删除账户下群组

- 官方接口：[`alipay.fund.enterprisepay.group.delete`](https://opendocs.alipay.com/apis/05wvyb)
- 路由：`DELETE /v3/alipay/fund/enterprisepay/group/delete`
- 请求类型：[FundEnterprisePayGroupDeleteRequest](enterprisepay_request.go#L103)
- 响应类型：[FundEnterprisePayGroupDeleteResponse](enterprisepay_response.go#L14)

因公付删除账户下群组

请求约束：FundEnterprisePayGroupDeleteRequest 是 alipay.fund.enterprisepay.group.delete 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码，默认值 ENTERPRISE_PAY |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 场景码，联系支付宝分配 |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业签约账户ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 平台和企业的三方授权协议号 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 外部业务号 |

<a id="api-enterprisepaygroupmodify"></a>

### EnterprisePayGroupModify — 因公付更新账户下群组

- 官方接口：[`alipay.fund.enterprisepay.group.modify`](https://opendocs.alipay.com/apis/05wmzb)
- 路由：`POST /v3/alipay/fund/enterprisepay/group/modify`
- 请求类型：[FundEnterprisePayGroupModifyRequest](enterprisepay_request.go#L25)
- 响应类型：[FundEnterprisePayGroupModifyResponse](enterprisepay_response.go#L18)

因公付更新账户下群组

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业签约账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 平台和企业的三方授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码，联系支付宝分配 |
| `fund_ext_info` | `FundExtInfo` | [`*FundExtInfo`](../README.md#model-fundextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 群组名称 |
| `operation_type_list` | `OperationTypeList` | `*[]string` | JSON body | 未标注 | 更新操作类型：多个时逗号分隔 NAME-群组名称 FUND-出资信息 字段不传时，默认全量更新 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码，默认值 ENTERPRISE_PAY |

<a id="api-enterprisepaygroupquery"></a>

### EnterprisePayGroupQuery — 因公付查询账户下群组

- 官方接口：[`alipay.fund.enterprisepay.group.query`](https://opendocs.alipay.com/apis/05wcpq)
- 路由：`GET /v3/alipay/fund/enterprisepay/group/query`
- 请求类型：[FundEnterprisePayGroupQueryRequest](enterprisepay_request.go#L117)
- 响应类型：[FundEnterprisePayGroupQueryResponse](enterprisepay_response.go#L22)

查询因公付群组，包含群组信息和对应的出资主体

请求约束：FundEnterprisePayGroupQueryRequest 是 alipay.fund.enterprisepay.group.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码，默认值 ENTERPRISE_PAY |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 场景码，联系支付宝分配 |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业签约账户ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 平台和企业的三方授权协议号 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 外部业务号，外部群组号 |

<a id="api-enterprisepaymembermodify"></a>

### EnterprisePayMemberModify — 因公付更新员工资金协议

- 官方接口：[`alipay.fund.enterprisepay.member.modify`](https://opendocs.alipay.com/apis/05wvv1)
- 路由：`POST /v3/alipay/fund/enterprisepay/member/modify`
- 请求类型：[FundEnterprisePayMemberModifyRequest](enterprisepay_request.go#L44)
- 响应类型：[FundEnterprisePayMemberModifyResponse](enterprisepay_response.go#L36)

因公付更新员工资金协议

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业签约账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 平台和企业的三方授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码，联系支付宝分配 |
| `fund_ext_info` | `FundExtInfo` | [`*FundExtInfo`](../README.md#model-fundextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `group_id_list` | `GroupIdList` | `*[]string` | JSON body | 未标注 | 员工当前关联的群组ID，支持多个 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 蚂蚁统一会员openId |
| `operation_type_list` | `OperationTypeList` | `*[]string` | JSON body | 未标注 | 更新操作类型：多个时逗号分隔 GROUP-群组信息 FUND-出资信息 字段不传时，默认全量更新 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码，默认值 ENTERPRISE_PAY |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 蚂蚁统一会员ID |

<a id="api-enterprisepaysign"></a>

### EnterprisePaySign — 因公付企业签约

- 官方接口：[`alipay.fund.enterprisepay.sign`](https://opendocs.alipay.com/apis/05udr5)
- 路由：`POST /v3/alipay/fund/enterprisepay/sign`
- 请求类型：[FundEnterprisePaySignRequest](enterprisepay_request.go#L65)
- 响应类型：[FundEnterprisePaySignResponse](enterprisepay_response.go#L40)

用于获取因公付企业签约链接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_name` | `AccountName` | `*string` | JSON body | 未标注 | 企业简称，传空采用默认规则命名，使用公司名称 |
| `appoint_sign_terminal` | `AppointSignTerminal` | `*string` | JSON body | 未标注 | 指定签约终端类型 PC-电脑浏览器 MOBILE-手机端 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景，联系支付宝分配 |
| `ext_params` | `ExtParams` | `*string` | JSON body | 未标注 | 特殊场景下扩展字段 |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 开户账号： 当identity_type是ALIPAY_USER_ID时填支付宝会员ID（2088开头）； 当identity_type是ALIPAY_LOGON_ID 时填支付宝登录号； 当identity_type是OUT_USER_ID时填外部平台的用户uid; 当identity_type是ALIPAY_OPEN_ID时填用户openId |
| `identity_name` | `IdentityName` | `*string` | JSON body | 未标注 | 开户用户名称，identity_type为ALIPAY_LOGON_ID时必填 |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、OUT_USER_ID：外部用户uid |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 商户侧单号（幂等字段） ，补充说明：幂等逻辑（商户appid+out_biz_no），商户集成时需注意，如换号重复发起，则认为是一笔新的请求 |
| `out_entity_id` | `OutEntityId` | `*string` | JSON body | 未标注 | 外部实体ID，一般表述为外部操作人操作某个外部实体。比如企业码2284号 |
| `out_source` | `OutSource` | `*string` | JSON body | 未标注 | ISV平台定义的共同账户开户来源 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |

<a id="api-enterprisepayunsign"></a>

### EnterprisePayUnsign — 因公付企业解约

- 官方接口：[`alipay.fund.enterprisepay.unsign`](https://opendocs.alipay.com/apis/05wbw2)
- 路由：`POST /v3/alipay/fund/enterprisepay/unsign`
- 请求类型：[FundEnterprisePayUnsignRequest](enterprisepay_request.go#L91)
- 响应类型：[FundEnterprisePayUnsignResponse](enterprisepay_response.go#L48)

用于因公付企业解约时获取解约链接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业账号 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码 |
