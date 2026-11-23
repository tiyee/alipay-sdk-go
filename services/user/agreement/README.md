# agreement 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/user/agreement`。本包提供 8 个支付宝 V3 REST 操作。

通过 `agreement.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AgreementExecutionPlanModify`](#api-agreementexecutionplanmodify) | `alipay.user.agreement.executionplan.modify` |
| [`AgreementMigrate`](#api-agreementmigrate) | `alipay.user.agreement.migrate` |
| [`AgreementPermissionCreate`](#api-agreementpermissioncreate) | `alipay.user.agreement.permission.create` |
| [`AgreementQuery`](#api-agreementquery) | `alipay.user.agreement.query` |
| [`AgreementTransfer`](#api-agreementtransfer) | `alipay.user.agreement.transfer` |
| [`AgreementUnsign`](#api-agreementunsign) | `alipay.user.agreement.unsign` |
| [`AgreementUserVerifyApply`](#api-agreementuserverifyapply) | `alipay.user.agreement.userverify.apply` |
| [`AgreementUserVerifyQuery`](#api-agreementuserverifyquery) | `alipay.user.agreement.userverify.query` |

## 使用示例

- [AgreementExecutionPlanModify](api_agreement_agreementexecutionplanmodify_example_test.go)
- [AgreementMigrate](api_agreement_agreementmigrate_example_test.go)
- [AgreementPermissionCreate](api_agreement_agreementpermissioncreate_example_test.go)
- [AgreementQuery](api_agreement_agreementquery_example_test.go)
- [AgreementTransfer](api_agreement_agreementtransfer_example_test.go)
- [AgreementUnsign](api_agreement_agreementunsign_example_test.go)
- [AgreementUserVerifyApply](api_agreement_agreementuserverifyapply_example_test.go)
- [AgreementUserVerifyQuery](api_agreement_agreementuserverifyquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-agreementexecutionplanmodify"></a>

### AgreementExecutionPlanModify — 周期性扣款协议执行计划修改接口

- 官方接口：[`alipay.user.agreement.executionplan.modify`](https://opendocs.alipay.com/open-v3/21227669_alipay.user.agreement.executionplan.modify?scene=common&pathHash=89395d7c)
- 路由：`POST /v3/alipay/user/agreement/executionplan/modify`
- 请求类型：[UserAgreementExecutionPlanModifyRequest](agreement_request.go#L4)
- 响应类型：[UserAgreementExecutionPlanModifyResponse](agreement_response.go#L8)

周期性扣款产品，商户延期扣款接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 周期性扣款产品，授权免密支付协议号 |
| `deduct_time` | `DeductTime` | `*string` | JSON body | 未标注 | 商户下一次扣款时间 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 具体修改原因 |

<a id="api-agreementmigrate"></a>

### AgreementMigrate — 代扣协议迁移

- 官方接口：[`alipay.user.agreement.migrate`](https://opendocs.alipay.com/open-v3/a611f0a3_alipay.user.agreement.migrate)
- 路由：`POST /v3/alipay/user/agreement/migrate`
- 请求类型：[UserAgreementMigrateRequest](agreement_request.go#L14)
- 响应类型：[UserAgreementMigrateResponse](agreement_response.go#L16)

由商户调用，将商户与用户签署的代扣协议内容进行迁移，包括协议主体迁移等

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） |
| `target_app_id` | `TargetAppId` | `*string` | JSON body | 未标注 | 需要迁移的目标商户应用id |
| `target_invoke_app_id` | `TargetInvokeAppId` | `*string` | JSON body | 未标注 | 需要迁移的目标商户调用应用id |
| `target_partner_id` | `TargetPartnerId` | `*string` | JSON body | 未标注 | 需要迁移的目标商户id |

<a id="api-agreementpermissioncreate"></a>

### AgreementPermissionCreate — 支付宝协议支付许可创建、服务变更接口

- 官方接口：[`alipay.user.agreement.permission.create`](https://opendocs.alipay.com/open-v3/87401d30_alipay.user.agreement.permission.create)
- 路由：`POST /v3/alipay/user/agreement/permission/create`
- 请求类型：[UserAgreementPermissionCreateRequest](agreement_request.go#L26)
- 响应类型：[UserAgreementPermissionCreateResponse](agreement_response.go#L19)

支付宝协议支付许可创建接口、服务变更发送消息提醒接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） |
| `business_params` | `BusinessParams` | [`*BusinessParamsMap`](#model-businessparamsmap) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `notice_template` | `NoticeTemplate` | `*string` | JSON body | 未标注 | 目前共两种类型SERVIOCE_NOTICE和DEFAULT_PERMISSION，如果是服务变更提醒，需传入SERVIOCE_NOTICE |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户请求号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。 |
| `total_amount` | `TotalAmount` | `*string` | JSON body | 未标注 | 订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。 |

<a id="api-agreementquery"></a>

### AgreementQuery — 支付宝个人代扣协议查询接口

- 官方接口：[`alipay.user.agreement.query`](https://opendocs.alipay.com/open-v3/4dd54af2_alipay.user.agreement.query)
- 路由：`GET /v3/alipay/user/agreement/query`
- 请求类型：[UserAgreementQueryRequest](agreement_request.go#L80)
- 响应类型：[UserAgreementQueryResponse](agreement_response.go#L31)

支付宝个人代扣协议查询接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `personal_product_code` | `PersonalProductCode` | `*string` | query | 否 | 协议产品码，商户和支付宝签约时确定，商户可咨询技术支持。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 用户的支付宝账号对应 的支付宝唯一用户号，以 2088 开头的 16 位纯数字 组成。 本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | query | 否 | 用户的支付宝账号对应 的支付宝唯一用户号， 本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。 |
| `alipay_logon_id` | `AlipayLogonId` | `*string` | query | 否 | 用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_open_id 或 alipay_user_id 同时填写，优先按照 alipay_open_id 或 alipay_user_id 处理。 |
| `sign_scene` | `SignScene` | `*string` | query | 否 | 签约场景码，该值需要与系统/页面签约接口调用时传入的值保持一 致。如：周期扣款场景与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。 注意：当传入商户签约号 external_agreement_no 时，该值不能为空或默认值 DEFAULT&#124;DEFAULT。 |
| `external_agreement_no` | `ExternalAgreementNo` | `*string` | query | 否 | 代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。 格式规则:支持大写小写字母和数字，最长 32 位。 |
| `third_party_type` | `ThirdPartyType` | `*string` | query | 否 | 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 默认为PARTNER。 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略 |

<a id="api-agreementtransfer"></a>

### AgreementTransfer — 协议由普通通用代扣协议产品转移到周期扣协议产品

- 官方接口：[`alipay.user.agreement.transfer`](https://opendocs.alipay.com/open-v3/360a356d_alipay.user.agreement.transfer)
- 路由：`POST /v3/alipay/user/agreement/transfer`
- 请求类型：[UserAgreementTransferRequest](agreement_request.go#L39)
- 响应类型：[UserAgreementTransferResponse](agreement_response.go#L77)

由商户调用，将商户之前通用代扣产品转移到周期扣的协议产品

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略 |
| `period_rule_params` | `PeriodRuleParams` | [`*PeriodRuleParams`](#model-periodruleparams) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `target_product_code` | `TargetProductCode` | `*string` | JSON body | 未标注 | 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。这里指的是需要修改目标产品码的值 |

<a id="api-agreementunsign"></a>

### AgreementUnsign — 支付宝个人代扣协议解约接口

- 官方接口：[`alipay.user.agreement.unsign`](https://opendocs.alipay.com/open-v3/b16bea7c_alipay.user.agreement.unsign)
- 路由：`POST /v3/alipay/user/agreement/unsign`
- 请求类型：[UserAgreementUnsignRequest](agreement_request.go#L48)
- 响应类型：[UserAgreementUnsignResponse](agreement_response.go#L93)

支付宝个人代扣协议解约接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ），如果传了该参数，其他参数会被忽略 。 本参数与 external_agreement_no 不可同时为空。 |
| `alipay_logon_id` | `AlipayLogonId` | `*string` | JSON body | 未标注 | 用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_user_id 不可同时为空，若都填写，则以alipay_user_id 为准。 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | JSON body | 未标注 | 用户的支付宝账号对应的支付宝唯一用户号， 本参数与alipay_logon_id 不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 用户的支付宝账号对应的支付宝唯一用户号，以2088 开头的 16 位纯数字 组成;<br> 本参数与alipay_logon_id 不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id。 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 扩展参数 |
| `external_agreement_no` | `ExternalAgreementNo` | `*string` | JSON body | 未标注 | 代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。 格式规则:支持大写小写字母和数字，最长 32 位。 注意：若调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时传入 external_agreement_no 则该值必填且需与签约接口传入值相同。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `operate_type` | `OperateType` | `*string` | JSON body | 未标注 | 注意：仅异步解约需传入，其余情况无需传递本参数。 |
| `personal_product_code` | `PersonalProductCode` | `*string` | JSON body | 未标注 | 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码。 |
| `sign_scene` | `SignScene` | `*string` | JSON body | 未标注 | 签约协议场景，该值需要与系统/页面签约接口调用时传入的值保持一 致。如：周期扣款场景，需与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。 当传入商户签约号 external_agreement_no时，场景不能为空或默认值 DEFAULT&#124;DEFAULT。 |
| `third_party_type` | `ThirdPartyType` | `*string` | JSON body | 未标注 | 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 |

<a id="api-agreementuserverifyapply"></a>

### AgreementUserVerifyApply — 协议用户一致性验证

- 官方接口：[`alipay.user.agreement.userverify.apply`](https://opendocs.alipay.com/open-v3/df28f556_alipay.user.agreement.userverify.apply)
- 路由：`POST /v3/alipay/user/agreement/userverify/apply`
- 请求类型：[UserAgreementUserVerifyApplyRequest](agreement_request.go#L74)
- 响应类型：[UserAgreementUserVerifyApplyResponse](agreement_response.go#L96)

验证协议对应的用户信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 协议号 |

<a id="api-agreementuserverifyquery"></a>

### AgreementUserVerifyQuery — 用户协议鉴权查询结果

- 官方接口：[`alipay.user.agreement.userverify.query`](https://opendocs.alipay.com/open-v3/a514c337_alipay.user.agreement.userverify.query)
- 路由：`GET /v3/alipay/user/agreement/userverify/query`
- 请求类型：[UserAgreementUserVerifyQueryRequest](agreement_request.go#L100)
- 响应类型：[UserAgreementUserVerifyQueryResponse](agreement_response.go#L102)

商户使用用户验证申请接口时，获取token，然后通过次接口回去验证状态

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 协议号 |
| `apply_token` | `ApplyToken` | `*string` | query | 否 | 申请时获取的token |

## 嵌套请求参数

<a id="model-businessparamsmap"></a>

### BusinessParamsMap

类型定义：[BusinessParamsMap](agreement_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `change_time` | `ChangeTime` | `*string` | 未标注 | 服务变更时间 |
| `new_amount` | `NewAmount` | `*string` | 未标注 | 服务变更后新金额 |
| `ori_amount` | `OriAmount` | `*string` | 未标注 | 原始价格 |

<a id="model-periodruleparams"></a>

### PeriodRuleParams

类型定义：[PeriodRuleParams](agreement_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `execute_time` | `ExecuteTime` | `*string` | 未标注 | 首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。精确到日，格式为yyyy-MM-dd 结合其他必填的扣款周期参数，会确定商户以后的扣款计划。发起扣款的时间需符合这里的扣款计划。 |
| `is_trial_mode` | `IsTrialMode` | `*string` | 未标注 | 若签约中传该参数，且值为 Y ，则表示本次签约有试用期。若传值为 N 或不传该参数，则表示本次签约无试用期。 |
| `period` | `Period` | `*int64` | 未标注 | 周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period=90，则扣款周期为90天。 |
| `period_type` | `PeriodType` | `*string` | 未标注 | 周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。 DAY即扣款周期按天计，MONTH代表扣款周期按自然月。 与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；period_type为MONTH，period=3，则扣款周期为3个自然月。 自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），以此避免有些月份可能不存在对应日期的情况。 |
| `scene_rule_params` | `SceneRuleParams` | [`*SceneRuleParams`](#model-sceneruleparams) | 未标注 | 见字段类型及官方接口文档。 |
| `single_amount` | `SingleAmount` | `*string` | 未标注 | 单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。 |
| `total_amount` | `TotalAmount` | `*string` | 未标注 | 总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。 |
| `total_payments` | `TotalPayments` | `*int64` | 未标注 | 总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。 |

<a id="model-sceneruleparams"></a>

### SceneRuleParams

类型定义：[SceneRuleParams](agreement_model.go#L45)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_period` | `DiscountPeriod` | `*string` | 未标注 | 在周期扣场景化模板中配置优惠类型为优惠期玩法时需要该参数，表示后续有多少期扣款可享受优惠，值为自然数代表周期 |
| `low_price_period` | `LowPricePeriod` | `*string` | 未标注 | 在周期扣场景化模板中配置优惠类型为低价玩法时需要该参数，表示代扣低价期持续的时间。单位是天，该值为自然数 |
