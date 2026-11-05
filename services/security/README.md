# security 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/security`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../README.md)创建 `*core.Client`，再通过 `security.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_security_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ProdSssQuery`](#api-prodsssquery) | [`alipay.security.prod.sss.query`](https://opendocs.alipay.com/apis/0557cf) | 金柚测试3.0接口 | `POST` |
| [`RiskCustomerRiskSend`](#api-riskcustomerrisksend) | [`alipay.security.risk.customerrisk.send`](https://opendocs.alipay.com/open-v3/cfacc9de_alipay.security.risk.customerrisk.send) | 商户数据同步 | `POST` |
| [`RiskVerifyIdentityMiniAppConfirm`](#api-riskverifyidentityminiappconfirm) | [`alipay.security.risk.verifyidentity.miniapp.confirm`](https://opendocs.alipay.com/apis/05vidu) | 小程序核验服务结果确认接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-prodsssquery"></a>

### ProdSssQuery — 金柚测试3.0接口

- 官方接口：[`alipay.security.prod.sss.query`](https://opendocs.alipay.com/apis/0557cf)
- 路由：`POST /v3/alipay/security/prod/sss/query/{bbb}`
- 请求类型：[SecurityProdSssQueryRequest](security_request.go#L49)
- 响应类型：[SecurityProdSssQueryResponse](security_response.go#L9)

金柚测试3.0接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `aaa_open_id` | `AaaOpenId` | `*string` | query | 否 | 1 |
| `aaa` | `Aaa` | `*string` | query | 否 | 1 |
| `tesst` | `SecurityProdSssQuery.Tesst` | [`*[]PromiseDetail`](#model-promisedetail) | JSON body | 未标注 | stst |
| `xxx` | `SecurityProdSssQuery.Xxx` | [`*JinyouTestFive`](#model-jinyoutestfive) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `bbb` | `Bbb` | `string` | 路径 | 是 | activity_id 词条描述 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [SecurityProdSssQuery](security_model.go#L5) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-riskcustomerrisksend"></a>

### RiskCustomerRiskSend — 商户数据同步

- 官方接口：[`alipay.security.risk.customerrisk.send`](https://opendocs.alipay.com/open-v3/cfacc9de_alipay.security.risk.customerrisk.send)
- 路由：`POST /v3/alipay/security/risk/customerrisk/send`
- 请求类型：[SecurityRiskCustomerRiskSendRequest](security_request.go#L5)
- 响应类型：[SecurityRiskCustomerRiskSendResponse](security_response.go#L17)

判断商户有问题时，服务商将商户信息同步到支付宝

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bank_card_no` | `BankCardNo` | `*string` | JSON body | 未标注 | 交易对应的银行卡的卡号（支付宝直连和间连合作伙伴均可回传此参数） |
| `business_license_no` | `BusinessLicenseNo` | `*string` | JSON body | 未标注 | 该账号对应的营业执照号码（支付宝直连和间连合作伙伴均可回传此参数） |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 该用户对应的身份证号码（支付宝直连和间连合作伙伴均可回传此参数） |
| `email_address` | `EmailAddress` | `*string` | JSON body | 未标注 | 账户注册时填写的邮箱地址(适用于支付宝直连商户回传此信息) |
| `external_id` | `ExternalId` | `*string` | JSON body | 未标注 | 适用于间连业务场景下，商户在银行被分配的ID（适用于支付宝间连合作伙伴回传此信息） |
| `merch_name` | `MerchName` | `*string` | JSON body | 未标注 | 该笔风险交易购买的具体商品名称（适用于支付宝直连商户回传此信息） |
| `mobile` | `Mobile` | `*string` | JSON body | 未标注 | 手机（支付宝直连和间连合作伙伴均可回传此参数） |
| `mobile_ip` | `MobileIp` | `*string` | JSON body | 未标注 | 该账号登录合作伙伴平台时的IP地址（适用于支付宝直连商户回传此信息） |
| `order_ip` | `OrderIp` | `*string` | JSON body | 未标注 | 该账户在商户平台下单时的MAC地址或IP地址（适用于支付宝直连商户回传此信息） |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 适用于直连商户场景下的支付宝合作伙伴ID（适用于支付宝直连合作伙伴回传此信息） |
| `plat_account` | `PlatAccount` | `*string` | JSON body | 未标注 | 指第三方在商户平台注册成功后，平台给予的账户号（适用于支付宝直连商户回传此信息） |
| `process_code` | `ProcessCode` | `*string` | JSON body | 未标注 | 直连场景下，商户对该账户采取的措施：<br> 多种处理结果可调用多次，每次回传一种处理结果。<br> （01：暂停发货；<br> 02：延迟结算；<br> 03：关停账户；<br> 04：暂停发货+关停账户；<br> 05：延迟结算+关停账户；<br> 06：其他；<br> 07：平台进行退款退订；<br> 08：平台跟用户沟通后，用户撤诉。<br> 间连场景下，商户回传处理结果：<br> 对商户的处理情况：<br> 01：已延迟结算，02：关停商户，99：其他<br> 多重处理结果可调用多次，每次回传一种处理结果 |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 指间连交易场景下商家被支付宝分配的ID（适用于支付宝间连合作伙伴回传此信息） |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝推送的风险消息中的交易号（支付宝直连和间连合作伙伴均需回传此参数） |

<a id="api-riskverifyidentityminiappconfirm"></a>

### RiskVerifyIdentityMiniAppConfirm — 小程序核验服务结果确认接口

- 官方接口：[`alipay.security.risk.verifyidentity.miniapp.confirm`](https://opendocs.alipay.com/apis/05vidu)
- 路由：`POST /v3/alipay/security/risk/verifyidentity/miniapp/confirm`
- 请求类型：[SecurityRiskVerifyIdentityMiniAppConfirmRequest](security_request.go#L37)
- 响应类型：[SecurityRiskVerifyIdentityMiniAppConfirmResponse](security_response.go#L20)

确认身份核验结果接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `challenge` | `Challenge` | `*string` | JSON body | 未标注 | 业务的业务流水号，例如订单id |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 蚂蚁统一会员ID，验证时会与做身份认证时候会话中的openid做比对，如不相符核验结果不通过 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 蚂蚁统一会员ID，验证时会与做身份认证时候会话中的userId做比对，如不相符核验结果不通过 |
| `verify_id` | `VerifyId` | `*string` | JSON body | 未标注 | 核身校验id，是一次核身校验服务中唯一性的id |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-jinyoutestfive"></a>

### JinyouTestFive

类型定义：[JinyouTestFive](security_model.go#L12)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `f_1_f` | `F1F` | [`*JinyouTestFour`](#model-jinyoutestfour) | 未标注 | 见字段类型及官方接口文档。 |
| `f_2_f` | `F2F` | [`*JinyouTestThree`](#model-jinyoutestthree) | 未标注 | 见字段类型及官方接口文档。 |
| `f_3_openid` | `F3Openid` | `*string` | 未标注 | 1 |
| `f_3_y` | `F3Y` | `*string` | 未标注 | 1 |
| `f_4_n` | `F4N` | `*string` | 未标注 | 3 |

<a id="model-jinyoutestfour"></a>

### JinyouTestFour

类型定义：[JinyouTestFour](security_model.go#L24)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `f_1_openid` | `F1Openid` | `*string` | 未标注 | 1 |
| `f_1_y` | `F1Y` | `*string` | 未标注 | 1 |
| `f_2_f` | `F2F` | [`*JinyouTestThree`](#model-jinyoutestthree) | 未标注 | 见字段类型及官方接口文档。 |
| `f_3_openid` | `F3Openid` | `*string` | 未标注 | 1 |
| `f_3_y` | `F3Y` | `*int64` | 未标注 | 1 |

<a id="model-jinyoutestone"></a>

### JinyouTestOne

类型定义：[JinyouTestOne](security_model.go#L37)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `o_1_n` | `O1N` | `*string` | 未标注 | 1 |
| `o_2_openid` | `O2Openid` | `*string` | 未标注 | o_2_y |
| `o_2_y` | `O2Y` | `*string` | 未标注 | o_2_y |
| `o_3_openid` | `O3Openid` | `*string` | 未标注 | o_3_y |
| `o_3_y` | `O3Y` | `*int64` | 未标注 | o_3_y |

<a id="model-jinyoutestthree"></a>

### JinyouTestThree

类型定义：[JinyouTestThree](security_model.go#L51)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `th_1_f` | `Th1F` | [`*JinyouTestTwo`](#model-jinyoutesttwo) | 未标注 | 见字段类型及官方接口文档。 |
| `th_2_n` | `Th2N` | `*string` | 未标注 | 无枚举值 |
| `th_3_n` | `Th3N` | `*bool` | 未标注 | 1 |

<a id="model-jinyoutesttwo"></a>

### JinyouTestTwo

类型定义：[JinyouTestTwo](security_model.go#L60)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `t_1_openid` | `T1Openid` | `*string` | 未标注 | 1 |
| `t_1_y` | `T1Y` | `*string` | 未标注 | 1 |
| `t_2_f` | `T2F` | [`*JinyouTestOne`](#model-jinyoutestone) | 未标注 | 见字段类型及官方接口文档。 |
| `t_3_n` | `T3N` | `*string` | 未标注 | 1 |

<a id="model-promisedetail"></a>

### PromiseDetail

类型定义：[PromiseDetail](security_model.go#L71)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_status` | `AuthStatus` | `*bool` | 未标注 | 授权状态 |
| `create_time` | `CreateTime` | `*string` | 未标注 | 任务创建时间 |
| `end_time` | `EndTime` | `*string` | 未标注 | 任务结束时间 |
| `final_time` | `FinalTime` | `*string` | 未标注 | 任务到达终态的时间 |
| `finish_periods` | `FinishPeriods` | `*int64` | 未标注 | 任务完成期数 |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 芝麻侧的商户id |
| `merchant_logo` | `MerchantLogo` | `*string` | 未标注 | 商户logo |
| `merchant_name` | `MerchantName` | `*string` | 未标注 | 商户名称 |
| `out_biz_no` | `OutBizNo` | `*string` | 未标注 | 生活记录加入时的外部业务号 |
| `period_type` | `PeriodType` | `*string` | 未标注 | 周期类型 |
| `promise_name` | `PromiseName` | `*string` | 未标注 | 生活记录模板名称 |
| `record_id` | `RecordId` | `*string` | 未标注 | 生活记录主记录id |
| `record_status` | `RecordStatus` | `*string` | 未标注 | 主任务状态 |
| `start_time` | `StartTime` | `*string` | 未标注 | 任务开始时间 |
| `sub_record_status` | `SubRecordStatus` | `*string` | 未标注 | 子记录状态 |
| `sub_title` | `SubTitle` | `*string` | 未标注 | 副标题 |
| `template_id` | `TemplateId` | `*string` | 未标注 | 生活记录模板id |
| `total_periods` | `TotalPeriods` | `*int64` | 未标注 | 任务总期数 |
