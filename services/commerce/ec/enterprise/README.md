# enterprise 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/ec/enterprise`。本包提供 9 个支付宝 V3 REST 操作。

通过 `enterprise.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`EcEnterpriseAddressAdd`](#api-ecenterpriseaddressadd) | `alipay.commerce.ec.enterprise.address.add` |
| [`EcEnterpriseAddressModify`](#api-ecenterpriseaddressmodify) | `alipay.commerce.ec.enterprise.address.modify` |
| [`EcEnterpriseAddressQuery`](#api-ecenterpriseaddressquery) | `alipay.commerce.ec.enterprise.address.query` |
| [`EcEnterpriseAgreementQuery`](#api-ecenterpriseagreementquery) | `alipay.commerce.ec.enterprise.agreement.query` |
| [`EcEnterpriseCreate`](#api-ecenterprisecreate) | `alipay.commerce.ec.enterprise.create` |
| [`EcEnterpriseDelete`](#api-ecenterprisedelete) | `alipay.commerce.ec.enterprise.delete` |
| [`EcEnterpriseInfoModify`](#api-ecenterpriseinfomodify) | `alipay.commerce.ec.enterprise.info.modify` |
| [`EcEnterpriseInfoQuery`](#api-ecenterpriseinfoquery) | `alipay.commerce.ec.enterprise.info.query` |
| [`EcEnterpriseUnsign`](#api-ecenterpriseunsign) | `alipay.commerce.ec.enterprise.unsign` |

## 使用示例

- [EcEnterpriseAddressAdd](api_enterprise_ecenterpriseaddressadd_example_test.go)
- [EcEnterpriseAddressModify](api_enterprise_ecenterpriseaddressmodify_example_test.go)
- [EcEnterpriseAddressQuery](api_enterprise_ecenterpriseaddressquery_example_test.go)
- [EcEnterpriseAgreementQuery](api_enterprise_ecenterpriseagreementquery_example_test.go)
- [EcEnterpriseCreate](api_enterprise_ecenterprisecreate_example_test.go)
- [EcEnterpriseDelete](api_enterprise_ecenterprisedelete_example_test.go)
- [EcEnterpriseInfoModify](api_enterprise_ecenterpriseinfomodify_example_test.go)
- [EcEnterpriseInfoQuery](api_enterprise_ecenterpriseinfoquery_example_test.go)
- [EcEnterpriseUnsign](api_enterprise_ecenterpriseunsign_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ecenterpriseaddressadd"></a>

### EcEnterpriseAddressAdd — 企业地址添加

- 官方接口：[`alipay.commerce.ec.enterprise.address.add`](https://opendocs.alipay.com/apis/05j04z)
- 路由：`POST /v3/alipay/commerce/ec/enterprise/address`
- 请求类型：[EcEnterpriseAddressAddRequest](enterprise_request.go#L4)
- 响应类型：[EcEnterpriseAddressAddResponse](enterprise_response.go#L10)

企业地址添加

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。 |
| `address` | `Address` | `*string` | JSON body | 未标注 | 详细地址最长50个字符 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 市(国家统一行政规划编码) |
| `city_name` | `CityName` | `*string` | JSON body | 未标注 | 市 |
| `community` | `Community` | `*string` | JSON body | 未标注 | 小区/楼宇 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度 |
| `mark` | `Mark` | `*string` | JSON body | 未标注 | 备注最长50个字符 |
| `poi_id` | `PoiId` | `*string` | JSON body | 未标注 | 高德地图poi |

<a id="api-ecenterpriseaddressmodify"></a>

### EcEnterpriseAddressModify — 企业地址修改

- 官方接口：[`alipay.commerce.ec.enterprise.address.modify`](https://opendocs.alipay.com/apis/05j1bb)
- 路由：`PUT /v3/alipay/commerce/ec/enterprise/address`
- 请求类型：[EcEnterpriseAddressModifyRequest](enterprise_request.go#L30)
- 响应类型：[EcEnterpriseAddressModifyResponse](enterprise_response.go#L16)

企业地址修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。 |
| `address` | `Address` | `*string` | JSON body | 未标注 | 详细地址最长50个字符 |
| `address_id` | `AddressId` | `*string` | JSON body | 未标注 | 地址id |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 市(国家统一行政规划编码) |
| `city_name` | `CityName` | `*string` | JSON body | 未标注 | 城市名称 |
| `community` | `Community` | `*string` | JSON body | 未标注 | 小区/楼宇 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度 |
| `mark` | `Mark` | `*string` | JSON body | 未标注 | 备注 |
| `poi_id` | `PoiId` | `*string` | JSON body | 未标注 | 高德地图poi |
| `status` | `Status` | `*string` | JSON body | 未标注 | 状态(生效/失效) |

<a id="api-ecenterpriseaddressquery"></a>

### EcEnterpriseAddressQuery — 企业地址查询

- 官方接口：[`alipay.commerce.ec.enterprise.address.query`](https://opendocs.alipay.com/apis/05iz2p)
- 路由：`GET /v3/alipay/commerce/ec/enterprise/address`
- 请求类型：[EcEnterpriseAddressQueryRequest](enterprise_request.go#L107)
- 响应类型：[EcEnterpriseAddressQueryResponse](enterprise_response.go#L22)

查询企业地址信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。 |
| `account_id` | `AccountId` | `*string` | query | 否 | 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。 |
| `address_id` | `AddressId` | `*string` | query | 否 | 地址id |
| `city_code` | `CityCode` | `*string` | query | 否 | 市 |
| `page_num` | `PageNum` | `*int64` | query | 是 | 页码从1开始 |
| `page_size` | `PageSize` | `*int64` | query | 是 | 每页数据 |

<a id="api-ecenterpriseagreementquery"></a>

### EcEnterpriseAgreementQuery — 查询企业签约企业码信息

- 官方接口：[`alipay.commerce.ec.enterprise.agreement.query`](https://opendocs.alipay.com/apis/05j2cw)
- 路由：`GET /v3/alipay/commerce/ec/enterprise/agreement`
- 请求类型：[EcEnterpriseAgreementQueryRequest](enterprise_request.go#L125)
- 响应类型：[EcEnterpriseAgreementQueryResponse](enterprise_response.go#L34)

查询企业签约企业码信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id；enterprise_id与account_id至少传一个，如果都传了优先使用enterprise_id |
| `account_id` | `AccountId` | `*string` | query | 否 | 资金共同账号id；enterprise_id与account_id至少传一个，如果都传了优先使用enterprise_id(为兼容企业码1.0老接口保留的字段，已过时，新接客户请使用enterprise_id) |

<a id="api-ecenterprisecreate"></a>

### EcEnterpriseCreate — 企业签约注册-接口已过时

- 官方接口：[`alipay.commerce.ec.enterprise.create`](https://opendocs.alipay.com/pre-apis/05l7n7)
- 路由：`POST /v3/alipay/commerce/ec/enterprise/create`
- 请求类型：[EcEnterpriseCreateRequest](enterprise_request.go#L60)
- 响应类型：[EcEnterpriseCreateResponse](enterprise_response.go#L39)

注册企业并返回因公付签约老页面链接，企业访问此链接可签约因公付、设置出资账户。
此接口已过时，仅适用于存量企业与出资场景，增量业务请勿使用！

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 用于区分资金因公付协议、以及收银台文案等差异； |
| `create_fund_account` | `CreateFundAccount` | `*bool` | JSON body | 未标注 | 是否设置管理员支付宝为企业出资账户 |
| `create_iot_group` | `CreateIotGroup` | `*bool` | JSON body | 未标注 | 是否创建企业人脸库，适用于对接团餐刷脸付、门禁刷脸通行等场景 |
| `enterprise_alias` | `EnterpriseAlias` | `*string` | JSON body | 未标注 | 企业简称 |
| `enterprise_name` | `EnterpriseName` | `*string` | JSON body | 未标注 | 企业名称 |
| `group_app_id` | `GroupAppId` | `*string` | JSON body | 未标注 | 订购【一脸通行开通插件】的小程序appId |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 管理员身份标识 |
| `identity_name` | `IdentityName` | `*string` | JSON body | 未标注 | 管理员姓名 |
| `identity_open_id` | `IdentityOpenId` | `*string` | JSON body | 未标注 | 管理员身份openId |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 管理员身份类型 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号 |
| `sign_return_url` | `SignReturnUrl` | `*string` | JSON body | 未标注 | 企业码签约后回跳地址。不传默认为空，签约后停留在当前页。 |
| `sign_terminal` | `SignTerminal` | `*string` | JSON body | 未标注 | 管理员签约终端类型(字段已废弃，不需要服务商传入) |

注意事项：

- 官方快照标记此接口已过时或废弃；接入前请确认替代接口与当前开放状态。

<a id="api-ecenterprisedelete"></a>

### EcEnterpriseDelete — 企业注销

- 官方接口：[`alipay.commerce.ec.enterprise.delete`](https://opendocs.alipay.com/apis/05j82m)
- 路由：`POST /v3/alipay/commerce/ec/enterprise/delete`
- 请求类型：[EcEnterpriseDeleteRequest](enterprise_request.go#L90)
- 响应类型：[EcEnterpriseDeleteResponse](enterprise_response.go#L51)

企业注销，清理企业数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |

<a id="api-ecenterpriseinfomodify"></a>

### EcEnterpriseInfoModify — 修改企业基础信息

- 官方接口：[`alipay.commerce.ec.enterprise.info.modify`](https://opendocs.alipay.com/apis/05ip0v)
- 路由：`PUT /v3/alipay/commerce/ec/enterprise`
- 请求类型：[EcEnterpriseInfoModifyRequest](enterprise_request.go#L96)
- 响应类型：[EcEnterpriseInfoModifyResponse](enterprise_response.go#L57)

修改企业基础信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_alias` | `EnterpriseAlias` | `*string` | JSON body | 未标注 | 企业简称 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `enterprise_name` | `EnterpriseName` | `*string` | JSON body | 未标注 | 企业名称 |
| `reliable_profiles` | `ReliableProfiles` | [`*ReliableEnterpriseProfilesDTO`](#model-reliableenterpriseprofilesdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-ecenterpriseinfoquery"></a>

### EcEnterpriseInfoQuery — 查询企业详情

- 官方接口：[`alipay.commerce.ec.enterprise.info.query`](https://opendocs.alipay.com/apis/05iwc6)
- 路由：`GET /v3/alipay/commerce/ec/enterprise/info`
- 请求类型：[EcEnterpriseInfoQueryRequest](enterprise_request.go#L133)
- 响应类型：[EcEnterpriseInfoQueryResponse](enterprise_response.go#L60)

查询企业详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id,从企业注册等接口获取，和out_biz_no参数二选一，都传优先取企业id |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 服务商生成的请求唯一流水号/业务幂等号，和enterprise_id参数二选一，都传优先取企业id |

<a id="api-ecenterpriseunsign"></a>

### EcEnterpriseUnsign — 企业解约

- 官方接口：[`alipay.commerce.ec.enterprise.unsign`](https://opendocs.alipay.com/apis/05jach)
- 路由：`GET /v3/alipay/commerce/ec/enterprise/unsignurl`
- 请求类型：[EcEnterpriseUnsignRequest](enterprise_request.go#L141)
- 响应类型：[EcEnterpriseUnsignResponse](enterprise_response.go#L65)

企业解约，获取解约链接，点击链接跳转到解约页面，进行解约协议

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。 |
| `account_id` | `AccountId` | `*string` | query | 否 | 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。(为兼容企业码1.0老接口的参数，已过时，新接客户请传企业ID) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。(为兼容企业码1.0老接口的参数，已过时，新接客户请传企业ID) |

## 嵌套请求参数

<a id="model-reliableenterpriseprofilesdto"></a>

### ReliableEnterpriseProfilesDTO

类型定义：[ReliableEnterpriseProfilesDTO](enterprise_model.go#L44)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `access_channel` | `AccessChannel` | `*string` | 未标注 | 企业注册来源渠道 |
| `activity_level` | `ActivityLevel` | `*string` | 未标注 | 企业在来源渠道（如钉钉）的活跃度 |
| `auth_type` | `AuthType` | `*string` | 未标注 | 企业在来源渠道（如钉钉）的认证类型 |
| `biz_license_no` | `BizLicenseNo` | `*string` | 未标注 | 企业营业执照编码，适用于未三证合一的企业 |
| `biz_license_url` | `BizLicenseUrl` | `*string` | 未标注 | 企业在来源渠道（如钉钉）的营业执照照片URL |
| `count_employee` | `CountEmployee` | `*string` | 未标注 | 企业当前的员工数规模区间 |
| `legal_person_cert_no` | `LegalPersonCertNo` | `*string` | 未标注 | 企业法人的身份证号码 |
| `platform_gmt_create` | `PlatformGmtCreate` | `*string` | 未标注 | 企业在来源渠道（如钉钉）的注册时间 |
| `sign_fund_inst_id` | `SignFundInstId` | `*string` | 未标注 | 机构id |
| `tax_no` | `TaxNo` | `*string` | 未标注 | 企业纳税人识别号，适用于未三证合一的企业 |
| `utm_chain` | `UtmChain` | `*string` | 未标注 | 用于标识企业入驻链路 |
| `utm_content` | `UtmContent` | `*string` | 未标注 | 具体流量内容区分，如推荐码、推荐人、推荐服务商名称等。由渠道侧自定义 |
| `utm_source` | `UtmSource` | `*string` | 未标注 | 渠道侧流量来源，由渠道侧自定义 |
