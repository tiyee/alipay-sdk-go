# auth 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/auth`。本包提供 6 个支付宝 V3 REST 操作。

通过 `auth.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AuthOperationCancel`](#api-authoperationcancel) | `alipay.fund.auth.operation.cancel` |
| [`AuthOperationDetailQuery`](#api-authoperationdetailquery) | `alipay.fund.auth.operation.detail.query` |
| [`AuthOrderAppFreeze`](#api-authorderappfreeze) | `alipay.fund.auth.order.app.freeze` |
| [`AuthOrderFreeze`](#api-authorderfreeze) | `alipay.fund.auth.order.freeze` |
| [`AuthOrderUnfreeze`](#api-authorderunfreeze) | `alipay.fund.auth.order.unfreeze` |
| [`AuthOrderVoucherCreate`](#api-authordervouchercreate) | `alipay.fund.auth.order.voucher.create` |

## 使用示例

- [AuthOperationCancel](api_auth_authoperationcancel_example_test.go)
- [AuthOperationDetailQuery](api_auth_authoperationdetailquery_example_test.go)
- [AuthOrderAppFreeze](api_auth_authorderappfreeze_example_test.go)
- [AuthOrderFreeze](api_auth_authorderfreeze_example_test.go)
- [AuthOrderUnfreeze](api_auth_authorderunfreeze_example_test.go)
- [AuthOrderVoucherCreate](api_auth_authordervouchercreate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-authoperationcancel"></a>

### AuthOperationCancel — 资金授权撤销接口

- 官方接口：[`alipay.fund.auth.operation.cancel`](https://opendocs.alipay.com/open-v3/2fecf986_alipay.fund.auth.operation.cancel)
- 路由：`POST /v3/alipay/fund/auth/operation/cancel`
- 请求类型：[FundAuthOperationCancelRequest](auth_request.go#L4)
- 响应类型：[FundAuthOperationCancelResponse](auth_response.go#L4)

只有商户由于业务系统处理超时需要终止后续业务处理或者授权结果未知时可调用撤销，其他正常授权冻结的操作如需实现相同功能请调用资金授权解冻服务。提交资金授权后调用【资金授权操作查询】，没有明确的授权结果再调用【资金授权撤销】

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_no` | `AuthNo` | `*string` | JSON body | 未标注 | 支付宝授权资金订单号。 与商户的授权资金订单号不能同时为空，二者都传入时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `operation_id` | `OperationId` | `*string` | JSON body | 未标注 | 支付宝的授权资金操作流水号。 与商户的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户的授权资金订单号。 与支付宝的授权资金订单号不能同时为空，二者都传入时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。 该值与资金冻结时 out_order_no一致。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户的授权资金操作流水号。 与支付宝的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。 该值与资金冻结时out_request_no一致 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 商户对本次撤销操作的附言描述。 长度不超过100个字母或50个汉字。 |

<a id="api-authoperationdetailquery"></a>

### AuthOperationDetailQuery — 资金授权操作查询接口

- 官方接口：[`alipay.fund.auth.operation.detail.query`](https://opendocs.alipay.com/open-v3/984254a0_alipay.fund.auth.operation.detail.query)
- 路由：`POST /v3/alipay/fund/auth/operation/detail/query`
- 请求类型：[FundAuthOperationDetailQueryRequest](auth_request.go#L20)
- 响应类型：[FundAuthOperationDetailQueryResponse](auth_response.go#L18)

通过该接口可以查询单笔明细的详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_no` | `AuthNo` | `*string` | JSON body | 未标注 | 支付宝授权资金订单号。 与商户的授权资金订单号不能同时为空，二者都传入时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。 |
| `operation_id` | `OperationId` | `*string` | JSON body | 未标注 | 支付宝的授权资金操作流水号。 与商户的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。 |
| `operation_type` | `OperationType` | `*string` | JSON body | 未标注 | 需要查询的授权资金操作类型。 未传入本参数时，如果仅查询出单笔明细则直接返回，如果查询出多笔则优先返回冻结明细、无冻结明细时返回解冻明细； 当传入本参数时，则严格按照该操作类型返回对应明细 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户的授权资金订单号。 与支付宝的授权资金订单号不能同时为空，二者都传入时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。 该值与资金冻结时 out_order_no一致。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户的授权资金操作流水号。 与支付宝的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。 查询冻结明细时，该值与发起冻结操作时传入的out_request_no一致； 查询解冻明细时，该值与发起解冻操作时传入的out_request_no一致； 查询支付明细时，该值与发起转支付操作时传入的out_trade_no一致。 |
| `query_options` | `QueryOptions` | `*[]string` | JSON body | 未标注 | 需要查询的额外信息 |

<a id="api-authorderappfreeze"></a>

### AuthOrderAppFreeze — 线上资金授权冻结接口

- 官方接口：[`alipay.fund.auth.order.app.freeze`](https://opendocs.alipay.com/open-v3/46bf87fe_alipay.fund.auth.order.app.freeze)
- 快照中的路由：`POST /v3/alipay/fund/auth/order/app/freeze`。官方接口文档演示通过 `sdkExecute` 构建 App 订单串，REST 调用方式仍待核实。
- 请求类型：[FundAuthOrderAppFreezeRequest](auth_request.go#L36)
- 响应类型：[FundAuthOrderAppFreezeResponse](auth_response.go#L85)

创建支付宝授权订单并完成资金冻结。适用于线上场景完成资金授权, 例如从商户APP端拉起支付宝收银台完成冻结。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00] |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | 业务参数，如风控参数outRiskInfo等。 |
| `deposit_product_mode` | `DepositProductMode` | `*string` | JSON body | 未标注 | 免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：[https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `disable_pay_channels` | `DisablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入 |
| `enable_pay_channels` | `EnablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入 |
| `extra_param` | `ExtraParam` | `*string` | JSON body | 未标注 | 业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见[https://opendocs.alipay.com/open/10719](https://opendocs.alipay.com/open/10719)； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 [开放平台-芝麻免押-信用服务管理](https://b.alipay.com/page/zmgaia/pre-auth/index) 创建信用服务获取，详情可查看 [创建信用服务](https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1)。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见[https://opendocs.alipay.com/open/11157/qlsxya](https://opendocs.alipay.com/open/11157/qlsxya) |
| `identity_params` | `IdentityParams` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；买家实名信息。 传入后支付宝会比对买家在支付宝端的实名信息。包含两个可选key： 1.identity_hash，买家姓名拼接身份证号后，使用SHA256摘要方式与UTF8编码后的hash值，返回的十六进制字符串，例如 “张三4566498798498498498498”对应的identity_hash应为“acc2b92ffc5ed9b472faa19748f10045c30434132784f774b00216a56b8841c6” 2.alipay_user_id，买家uid |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 订单标题。 业务订单的简单描述，如商品名称等 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。 |
| `pay_timeout` | `PayTimeout` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间。从商户客户端请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `payee_logon_id` | `PayeeLogonId` | `*string` | JSON body | 未标注 | 收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一 |
| `payee_user_id` | `PayeeUserId` | `*string` | JSON body | 未标注 | 收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一 |
| `post_payments` | `PostPayments` | [`*[]PostPayment`](#model-postpayment) | JSON body | 未标注 | 后付费项目， 有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档： [https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码。 预授权支付产品固定为 PREAUTH_PAY |
| `scene_code` | `SceneCode` | `*string` | JSON body | 未标注 | 场景码，用于区分预授权不同业务场景。如：当面预授权通用场景（O2O_AUTH_COMMON_SCENE）、支付宝预授权通用场景（ONLINE_AUTH_COMMON_SCENE）、境外当面预授权通用场景（OVERSEAS_O2O_AUTH_COMMON_SCENE）、境外支付预授权通用场景（OVERSEAS_ONLINE_AUTH_COMMON_SCENE）等 |
| `settle_currency` | `SettleCurrency` | `*string` | JSON body | 未标注 | 商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间。从商户客户端请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `trans_currency` | `TransCurrency` | `*string` | JSON body | 未标注 | 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP, 人民币：CNY |

<a id="api-authorderfreeze"></a>

### AuthOrderFreeze — 资金授权冻结接口

- 官方接口：[`alipay.fund.auth.order.freeze`](https://opendocs.alipay.com/open-v3/1ffb104d_alipay.fund.auth.order.freeze)
- 路由：`POST /v3/alipay/fund/auth/order/freeze`
- 请求类型：[FundAuthOrderFreezeRequest](auth_request.go#L80)
- 响应类型：[FundAuthOrderFreezeResponse](auth_response.go#L117)

收银员使用扫码设备读取用户支付宝钱包“付款码”后，将条码信息和订单信息通过本接口上送至支付宝发起资金冻结。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00] |
| `auth_code` | `AuthCode` | `*string` | JSON body | 未标注 | 用户付款码。 1.条码场景：25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准； 2.刷脸场景： 1）fp开头的35位字符串； 2）300-700字符的随机字符串； 注：刷脸场景考虑到未来可能拓展更多格式，建议外围不必做规则拦截，由支付宝统一做有效性校验 |
| `auth_code_type` | `AuthCodeType` | `*string` | JSON body | 未标注 | 付款码类型。 1.条码场景：bar_code 2.刷脸场景：security_code |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | 业务参数，如风控参数outRiskInfo等。 |
| `deposit_product_mode` | `DepositProductMode` | `*string` | JSON body | 未标注 | 免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：[https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `disable_pay_channels` | `DisablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入 |
| `enable_pay_channels` | `EnablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入 |
| `extra_param` | `ExtraParam` | `*string` | JSON body | 未标注 | 业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见[https://opendocs.alipay.com/open/10719](https://opendocs.alipay.com/open/10719)； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 [开放平台-芝麻免押-信用服务管理](https://b.alipay.com/page/zmgaia/pre-auth/index) 创建信用服务获取，详情可查看 [创建信用服务](https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1)。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见[https://opendocs.alipay.com/open/11157/qlsxya](https://opendocs.alipay.com/open/11157/qlsxya) |
| `identity_params` | `IdentityParams` | `*string` | JSON body | 未标注 | 用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。 姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。 identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如"张三4566498798498498498498" |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 订单标题。 业务订单的简单描述，如商品名称等 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。 |
| `pay_timeout` | `PayTimeout` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `payee_logon_id` | `PayeeLogonId` | `*string` | JSON body | 未标注 | 收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一 |
| `payee_user_id` | `PayeeUserId` | `*string` | JSON body | 未标注 | 收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一 |
| `post_payments` | `PostPayments` | [`*[]PostPayment`](#model-postpayment) | JSON body | 未标注 | 后付费项目， 有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档： [https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码。 当面资金预授权固定为 PRE_AUTH |
| `scene_code` | `SceneCode` | `*string` | JSON body | 未标注 | 场景码。 刷脸场景下传入HOTEL，其他情况下无需传入 |
| `settle_currency` | `SettleCurrency` | `*string` | JSON body | 未标注 | 商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP |
| `terminal_params` | `TerminalParams` | `*string` | JSON body | 未标注 | 机具管控sdk加签参数，参数示例 "terminal_params":"{"terminalType":"IOT","signature":"QIIAX8DqbFbNf2oe97FI1RSLAycC/tU4GVjer3bN8K4qLtAB","apdidToken":"xPA3ptuArwYc3F6Va_pjVwv7Qx7Tg5TJdrA_Jb_moYte9AqGZgEAAA==","hardToken":"","time":"1539847253","bizCode":"11000200040004000121","bizTid":"010100F01i1XyacMgpOinHerfdBw1xA9dNDocctlnqhLD8lfODr1A7Q","signedKeys":"authCode,totalAmount,apdidToken,hardToken,time,bizCode,bizTid"}" |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `trans_currency` | `TransCurrency` | `*string` | JSON body | 未标注 | 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP |

<a id="api-authorderunfreeze"></a>

### AuthOrderUnfreeze — 资金授权解冻接口

- 官方接口：[`alipay.fund.auth.order.unfreeze`](https://opendocs.alipay.com/open-v3/cb3b0ae2_alipay.fund.auth.order.unfreeze)
- 路由：`POST /v3/alipay/fund/auth/order/unfreeze`
- 请求类型：[FundAuthOrderUnfreezeRequest](auth_request.go#L130)
- 响应类型：[FundAuthOrderUnfreezeResponse](auth_response.go#L149)

当资金授权发生之后一段时间内，由于买家或者商家等其他原因需要要解冻资金，商家可通过资金授权解冻接口将授权资金进行解冻，支付宝将在收到解冻请求并验证成功后，按解冻规则将冻结资金按原路进行解冻

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00] |
| `auth_no` | `AuthNo` | `*string` | JSON body | 未标注 | 支付宝资金授权订单号。 |
| `extra_param` | `ExtraParam` | `*string` | JSON body | 未标注 | 解冻扩展信息。map&lt;string, string&gt;的json格式，目前支持如下key： unfreezeBizInfo：由芝麻消费，当前支持value如下： "bizComplete":"true"——标识本次解冻用户已履约，true表示信用单履约完结 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 解冻请求流水号。 如果是针对同一笔授权单不同的解冻请求，如第一次解冻1元，第二次解冻2元，则解冻请求流水号必须不重复； 如果是针对同一笔解冻请求的多次发起，则需要保证每次发起，解冻请求流水号和解冻金额都相同 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 商户对本次解冻操作的附言描述 |

<a id="api-authordervouchercreate"></a>

### AuthOrderVoucherCreate — 资金授权发码接口

- 官方接口：[`alipay.fund.auth.order.voucher.create`](https://opendocs.alipay.com/open-v3/6f5ccbfe_alipay.fund.auth.order.voucher.create)
- 路由：`POST /v3/alipay/fund/auth/order/voucher/create`
- 请求类型：[FundAuthOrderVoucherCreateRequest](auth_request.go#L146)
- 响应类型：[FundAuthOrderVoucherCreateResponse](auth_response.go#L171)

收银员通过收银台或商户后台调用支付宝接口，生成二维码后，展示给用户，由用户扫描二维码完成资金冻结。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00] |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | 业务参数，如风控参数outRiskInfo等。 |
| `deposit_product_mode` | `DepositProductMode` | `*string` | JSON body | 未标注 | 免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：[https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `disable_pay_channels` | `DisablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入 |
| `enable_pay_channels` | `EnablePayChannels` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入 |
| `extra_param` | `ExtraParam` | `*string` | JSON body | 未标注 | 业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见[https://opendocs.alipay.com/open/10719](https://opendocs.alipay.com/open/10719)； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 [开放平台-芝麻免押-信用服务管理](https://b.alipay.com/page/zmgaia/pre-auth/index) 创建信用服务获取，详情可查看 [创建信用服务](https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1)。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见[https://opendocs.alipay.com/open/11157/qlsxya](https://opendocs.alipay.com/open/11157/qlsxya) |
| `identity_params` | `IdentityParams` | `*string` | JSON body | 未标注 | 无特殊需要请勿传入；买家实名信息。 传入后支付宝会比对买家在支付宝端的实名信息。包含两个可选key： 1.identity_hash，买家姓名拼接身份证号后，使用SHA256摘要方式与UTF8编码后的hash值，返回的十六进制字符串，例如 “张三4566498798498498498498”对应的identity_hash应为“acc2b92ffc5ed9b472faa19748f10045c30434132784f774b00216a56b8841c6” 2.alipay_open_id，买家支付宝openId 3.alipay_user_id，买家支付宝uid |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 订单标题。 业务订单的简单描述，如商品名称等 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。 |
| `pay_timeout` | `PayTimeout` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `payee_logon_id` | `PayeeLogonId` | `*string` | JSON body | 未标注 | 收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一，二者无需同时传入 |
| `payee_user_id` | `PayeeUserId` | `*string` | JSON body | 未标注 | 收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一，二者无需同时传入 |
| `post_payments` | `PostPayments` | [`*[]PostPayment`](#model-postpayment) | JSON body | 未标注 | 后付费项目，有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档：[https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545](https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545) |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码。 当面资金预授权固定为 PRE_AUTH |
| `settle_currency` | `SettleCurrency` | `*string` | JSON body | 未标注 | 商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。 |
| `trans_currency` | `TransCurrency` | `*string` | JSON body | 未标注 | 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP |

## 嵌套请求参数

<a id="model-postpayment"></a>

### PostPayment

类型定义：[PostPayment](auth_model.go#L14)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 后付费金额，单位为：元（人民币），精确到小数点后两位。 |
| `description` | `Description` | `*string` | 未标注 | 计费说明 |
| `name` | `Name` | `*string` | 未标注 | 后付费项目名称 |
