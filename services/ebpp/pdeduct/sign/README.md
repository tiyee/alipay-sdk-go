# sign 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/pdeduct/sign`。本包提供 4 个支付宝 V3 REST 操作。

通过 `sign.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`PdeductSignAdd`](#api-pdeductsignadd) | `alipay.ebpp.pdeduct.sign.add` |
| [`PdeductSignCancel`](#api-pdeductsigncancel) | `alipay.ebpp.pdeduct.sign.cancel` |
| [`PdeductSignQuery`](#api-pdeductsignquery) | `alipay.ebpp.pdeduct.sign.query` |
| [`PdeductSignValidate`](#api-pdeductsignvalidate) | `alipay.ebpp.pdeduct.sign.validate` |

## 使用示例

- [PdeductSignAdd](api_sign_pdeductsignadd_example_test.go)
- [PdeductSignCancel](api_sign_pdeductsigncancel_example_test.go)
- [PdeductSignQuery](api_sign_pdeductsignquery_example_test.go)
- [PdeductSignValidate](api_sign_pdeductsignvalidate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-pdeductsignadd"></a>

### PdeductSignAdd — 缴费直连代扣签约

- 官方接口：[`alipay.ebpp.pdeduct.sign.add`](https://opendocs.alipay.com/open-v3/084db95a_alipay.ebpp.pdeduct.sign.add)
- 路由：`POST /v3/alipay/ebpp/pdeduct/sign/add`
- 请求类型：[PdeductSignAddRequest](sign_request.go#L4)
- 响应类型：[PdeductSignAddResponse](sign_response.go#L8)

缴费直连代扣签约

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_channel` | `AgentChannel` | `*string` | JSON body | 未标注 | 机构签约代扣来源渠道<br> PUBLICPLATFORM：服务窗 |
| `agent_code` | `AgentCode` | `*string` | JSON body | 未标注 | 从服务窗发起则为publicId的值 |
| `bill_key` | `BillKey` | `*string` | JSON body | 未标注 | 户号，机构针对于每户的水、电都会有唯一的标识户号 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 业务类型。枚举支持： *JF：公共事业缴纳。 *WUYE：物业缴费。 *HK：信用卡还款。 *TX：通讯缴费。 *IND：保险缴费。 |
| `charge_inst` | `ChargeInst` | `*string` | JSON body | 未标注 | 支付宝缴费系统中的出账机构ID |
| `deduct_prod_code` | `DeductProdCode` | `*string` | JSON body | 未标注 | 代扣产品码。 缴费业务场景中，传入INST_DIRECT_DEDUCT。保险业务场景中, 传入INSURANCE_MERCHANT_DEDUCT。 |
| `deduct_type` | `DeductType` | `*string` | JSON body | 未标注 | 签约类型可为空 |
| `ext_user_info` | `ExtUserInfo` | `*string` | JSON body | 未标注 | 外部用户实名认证相关信息, 用于做签约时的实名校验。 <br> 注： <br> name: 姓名<br> cert_type: 身份证：IDENTITY_CARD、护照：PASSPORT、军官证：OFFICER_CARD、士兵证：SOLDIER_CARD、户口本：HOKOU等<br> cert_no: 证件号码<br> need_check_info取值 T/F，只有为T时才做强制校验。<br> mobile:手机号，目前暂不使用此字段做校验<br> min_age: 允许的最小买家年龄,min_age为整数，必须大于等于0. |
| `extend_field` | `ExtendField` | `*string` | JSON body | 未标注 | 扩展字段 |
| `notify_config` | `NotifyConfig` | `*string` | JSON body | 未标注 | 通知方式设置，可为空 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户UserId在应用AppId下的唯一用户标识 |
| `out_agreement_id` | `OutAgreementId` | `*string` | JSON body | 未标注 | 外部产生的协议ID |
| `owner_name` | `OwnerName` | `*string` | JSON body | 未标注 | 户名，户主真实姓名 |
| `pay_config` | `PayConfig` | `*string` | JSON body | 未标注 | 支付工具设置，目前可为空 |
| `pay_password_token` | `PayPasswordToken` | `*string` | JSON body | 未标注 | 支付密码鉴权token，需要用户首先处于登陆态，然后访问https://ebppprod.alipay.com/deduct/enterMobileicPayPassword.htm?cb=自己指定的回跳连接地址，访问页面后会进到独立密码校验页，用户输入密码校验成功后，会生成token回调到指定的回跳地址，如果设置cb=www.alipay.com则最后回调的内容是www.alipay.com?token=312314ADFDSFAS，然后签约时直接取得地址后token的值即可。 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 商户id，商户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `sign_expire_date` | `SignExpireDate` | `*string` | JSON body | 未标注 | 签约到期时间。空表示无限期，一期固定传空。 |
| `specified_sort_channels` | `SpecifiedSortChannels` | [`*[]SpecifiedSortChannelParam`](#model-specifiedsortchannelparam) | JSON body | 未标注 | 用户设置的优先扣款渠道 |
| `sub_biz_type` | `SubBizType` | `*string` | JSON body | 未标注 | 业务子类型。业务子类型是业务类型的下一级概念。枚举支持： *WATER：缴水费。 *ELECTRIC：缴电费。 *GAS：缴燃气费。 *COMMUN：缴固话宽带费。 *CATV：缴有线电视费。 *TRAFFIC：缴交通罚款。 *WUYE：缴物业费。 *RC：定期还车贷代扣。 *RH：定期还房贷代扣。 *RR：定期还房租代扣。 *RN：定期还网贷代扣。 *CZ：手机充值代扣。 *CAR：车险。 *LIFE：寿险。 *HEALTH：健康险。 例如：WATER表示JF下面的水费。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |

<a id="api-pdeductsigncancel"></a>

### PdeductSignCancel — 缴费直连代扣取消签约

- 官方接口：[`alipay.ebpp.pdeduct.sign.cancel`](https://opendocs.alipay.com/open-v3/2ad13974_alipay.ebpp.pdeduct.sign.cancel)
- 路由：`POST /v3/alipay/ebpp/pdeduct/sign/cancel`
- 请求类型：[PdeductSignCancelRequest](sign_request.go#L48)
- 响应类型：[PdeductSignCancelResponse](sign_response.go#L26)

缴费直连代扣，用户取消签约接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_channel` | `AgentChannel` | `*string` | JSON body | 未标注 | 此值只是供代扣中心做最后的渠道统计用，并不做值是什么的强校验，只要不为空就可以 |
| `agent_code` | `AgentCode` | `*string` | JSON body | 未标注 | 标识发起方的ID，从服务窗发起则为appId的值，appId即开放平台分配给接入ISV的id，此处也可以随便真其它值，只要能标识唯一即可 |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝代扣协议ID |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户UserId在应用AppId下的唯一用户标识 |
| `pay_password_token` | `PayPasswordToken` | `*string` | JSON body | 未标注 | 需要用户首先处于登陆态，然后访问https://ebppprod.alipay.com/deduct/enterMobileicPayPassword.htm?cb=自己指定的回跳连接地址,访问页面后会进到独立密码校验页，用户输入密码校验成功后，会生成token回调到指定的回跳地址，如果设置cb=www.baidu.com则最后回调的内容是www.baidu.com?token=312314ADFDSFAS,然后签约时直接取得地址后token的值即可 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户ID |

<a id="api-pdeductsignquery"></a>

### PdeductSignQuery — 直连代扣协议查询接口

- 官方接口：[`alipay.ebpp.pdeduct.sign.query`](https://opendocs.alipay.com/apis/05x9ew)
- 路由：`GET /v3/alipay/ebpp/pdeduct/sign/query`
- 请求类型：[PdeductSignQueryRequest](sign_request.go#L100)
- 响应类型：[PdeductSignQueryResponse](sign_response.go#L36)

提供给朗新查询代扣协议状态的接口，用于进行双边对账，解决单边协议问题

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 用户ID |
| `open_id` | `OpenId` | `*string` | query | 否 | 用户UserId在应用AppId下的唯一用户标识 |
| `agreement_id` | `AgreementId` | `*string` | query | 否 | 支付宝代扣协议Id。若协议id不传递，则需要保证业务类型、子业务类型、出账机构、户号必传 |
| `biz_type` | `BizType` | `*string` | query | 否 | 业务类型。<br> JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用<br> WUYE：缴物业费<br> HK：信用卡还款<br> TX：手机充值 |
| `sub_biz_type` | `SubBizType` | `*string` | query | 否 | 业务子类型。<br> WATER：缴水费<br> ELECTRIC：缴电费<br> GAS：缴燃气费<br> COMMUN：缴固话宽带<br> CATV：缴有线电视费<br> TRAFFIC：缴交通罚款<br> WUYE：缴物业费<br> HK：信用卡还款<br> CZ：手机充值 |
| `charge_inst` | `ChargeInst` | `*string` | query | 否 | 支付宝缴费系统中的出账机构ID |
| `bill_key` | `BillKey` | `*string` | query | 否 | 户号，机构针对于每户的水、电都会有唯一的标识户号 |

<a id="api-pdeductsignvalidate"></a>

### PdeductSignValidate — 缴费直连代扣签约前置校验

- 官方接口：[`alipay.ebpp.pdeduct.sign.validate`](https://opendocs.alipay.com/apis/05wukm)
- 路由：`POST /v3/alipay/ebpp/pdeduct/sign/validate`
- 请求类型：[PdeductSignValidateRequest](sign_request.go#L64)
- 响应类型：[PdeductSignValidateResponse](sign_response.go#L56)

服务窗代扣签约的前置校验接口，签约前，可以先调用该接口进行校验是否可以签约

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_channel` | `AgentChannel` | `*string` | JSON body | 未标注 | 机构签约代扣来源渠道<br> PUBLICPLATFORM：服务窗 |
| `agent_code` | `AgentCode` | `*string` | JSON body | 未标注 | 从服务窗发起则为该服务窗的appid的值 |
| `bill_key` | `BillKey` | `*string` | JSON body | 未标注 | 户号，机构针对于每户的水、电都会有唯一的标识户号 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 业务类型。<br> JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用<br> WUYE：缴物业费<br> HK：信用卡还款<br> TX：手机充值 |
| `charge_inst` | `ChargeInst` | `*string` | JSON body | 未标注 | 支付宝缴费系统中的出账机构ID |
| `deduct_type` | `DeductType` | `*string` | JSON body | 未标注 | 签约类型,可为空。目前类型有INST_DIRECT_DEDUCT |
| `extend_field` | `ExtendField` | `*string` | JSON body | 未标注 | 扩展字段 |
| `notify_config` | `NotifyConfig` | `*string` | JSON body | 未标注 | 通知方式设置，本期预留字段，固定传空。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户UserId在应用AppId下的唯一用户标识 |
| `out_agreement_id` | `OutAgreementId` | `*string` | JSON body | 未标注 | 外部机构签约的协议id |
| `owner_name` | `OwnerName` | `*string` | JSON body | 未标注 | 户名，户主真实姓名 |
| `pay_config` | `PayConfig` | `*string` | JSON body | 未标注 | 支付工具设置，目前可为空。类型有：BALANCE;CARTOON;BIGAMOUNT_CREDIT_CARTOON;DEBIT_EXPRESS;OPTIMIZED_MOTO;PCREDIT_PAY;MONEY_FUND |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 商户签约支付宝账号的userid，2088开头16位长度的字符串，在支付宝系统中唯一 |
| `sign_expire_date` | `SignExpireDate` | `*string` | JSON body | 未标注 | 签约到期时间。空表示无限期，一期固定传空。 |
| `sub_biz_type` | `SubBizType` | `*string` | JSON body | 未标注 | 业务子类型。<br> WATER：缴水费<br> ELECTRIC：缴电费<br> GAS：缴燃气费<br> COMMUN：缴固话宽带<br> CATV：缴有线电视费<br> TRAFFIC：缴交通罚款<br> WUYE：缴物业费<br> HK：信用卡还款<br> CZ：手机充值 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户支付宝账号id，2088开头16位长度的字符串，在支付宝系统中唯一 |

## 嵌套请求参数

<a id="model-specifiedsortchannelparam"></a>

### SpecifiedSortChannelParam

类型定义：[SpecifiedSortChannelParam](sign_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `asset_type_code` | `AssetTypeCode` | `*string` | 未标注 | 资产编码 |
| `assigned_channel` | `AssignedChannel` | `*string` | 未标注 | 指定支付渠道标识 |
| `channel_full_name` | `ChannelFullName` | `*string` | 未标注 | 支付渠道名称 |
| `channel_index` | `ChannelIndex` | `*string` | 未标注 | 支付渠道索引 |
