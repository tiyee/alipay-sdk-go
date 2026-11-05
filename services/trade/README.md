# trade 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade`。本包提供 6 个支付宝 V3 REST 操作。

通过 `trade.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`advance`](advance/README.md) | 1 | `AdvanceConsult` |
| [`batch`](batch/README.md) | 4 | `BatchTransfer` / `BatchTransferQuery` / `BatchSettle` / `BatchSettleQuery` |
| [`customs`](customs/README.md) | 2 | `CustomsDeclare` / `CustomsQuery` |
| [`order`](order/README.md) | 3 | `OrderPay` / `OrderSettle` / `OrderSettleQuery` |
| [`orderinfo`](orderinfo/README.md) | 1 | `OrderInfoSync` |
| [`overseas`](overseas/README.md) | 1 | `OverseasSettle` |
| [`refund`](refund/README.md) | 3 | `RefundApply` / `Refund` / `QueryRefund` |
| [`royalty`](royalty/README.md) | 4 | `RoyaltyRateQuery` / `RoyaltyRelationBind` / `RoyaltyRelationUnbind` / `RoyaltyRelationBatchQuery` |
| [`settle`](settle/README.md) | 2 | `SettleConfirm` / `SettleReceivablesQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`Pay`](#api-pay) | `alipay.trade.pay` |
| [`Create`](#api-create) | `alipay.trade.create` |
| [`Cancel`](#api-cancel) | `alipay.trade.cancel` |
| [`Query`](#api-query) | `alipay.trade.query` |
| [`PreCreate`](#api-precreate) | `alipay.trade.precreate` |
| [`Close`](#api-close) | `alipay.trade.close` |

## 使用示例

- [Pay](api_trade_pay_example_test.go)
- [Create](api_trade_create_example_test.go)
- [Cancel](api_trade_cancel_example_test.go)
- [Query](api_trade_query_example_test.go)
- [PreCreate](api_trade_precreate_example_test.go)
- [Close](api_trade_close_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-pay"></a>

### Pay — 统一收单交易支付接口

- 官方接口：[`alipay.trade.pay`](https://opendocs.alipay.com/open-v3/08c7f9f8_alipay.trade.pay)
- 路由：`POST /v3/alipay/trade/pay`
- 请求类型：[PayRequest](trade_request.go#L96)
- 响应类型：[PayResponse](trade_response.go#L81)

收银员使用扫码设备读取用户手机支付宝“付款码”获取设备（如扫码枪）读取用户手机支付宝的付款码信息后，将二维码或条码信息通过本接口上送至支付宝发起支付。

请求约束：PayRequest 是 alipay.trade.pay（统一收单交易支付）的请求。
当面付场景：out_trade_no、total_amount、subject、scene、auth_code
为必选；周期扣款场景无需 scene 与 auth_code，协议号经 AgreementParams 传入。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyURL` | `*string` | JSON body | 未标注 | 本次请求的异步通知地址。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 商户订单号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。<br>SDK 字段说明：商户订单号：64 字符以内，仅支持字母、数字、下划线，<br>且需保证在商户端不重复。 |
| `total_amount` | `TotalAmount` | `string` | JSON body | 未标注 | 订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。<br>SDK 字段说明：订单总金额，两位小数，取值范围 [0.01,100000000]。 |
| `subject` | `Subject` | `string` | JSON body | 未标注 | 订单标题。 注意：256个字符以内，不可使用特殊字符，如 /，=，& 等。<br>SDK 字段说明：订单标题；不可使用 / = & 等特殊字符。 |
| `scene` | `Scene` | [`*PayScene`](trade_model.go#L247) | JSON body | 未标注 | 支付场景。枚举值： bar_code：当面付条码支付场景； security_code：当面付刷脸支付场景，对应的auth_code为fp开头的刷脸标识串； 周期扣款或代扣场景无需传入，协议号通过agreement_params参数传递； 支付宝预授权和新当面资金授权场景无需传入，授权订单号通过 auth_no字段传入。 默认值为bar_code。<br>SDK 字段说明：支付场景：bar_code（条码支付）或 security_code（刷脸支付）；<br>周期扣款或代扣场景无需传入。 |
| `auth_code` | `AuthCode` | `*string` | JSON body | 未标注 | 支付授权码。 当面付场景传买家的付款码（25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准）或者刷脸标识串（fp开头的35位字符串）； 周期扣款或代扣场景无需传入，协议号通过agreement_params参数传递； 支付宝预授权和新当面资金授权场景无需传入，授权订单号通过 auth_no字段传入。 注：交易的买家与卖家不能相同。<br>SDK 字段说明：买家付款码（25~30 开头的 16~24 位数字）或刷脸标识串<br>（fp 开头的 35 位字符串）；当面付场景必选。 |
| `auth_no` | `AuthNo` | `*string` | JSON body | 未标注 | 资金预授权单号。 支付宝预授权和新当面资金授权场景下必填。<br>SDK 字段说明：资金预授权单号；支付宝预授权与新当面资金授权场景必填。 |
| `auth_confirm_mode` | `AuthConfirmMode` | `*string` | JSON body | 未标注 | 预授权确认模式。 适用于支付宝预授权和新当面资金授权场景。枚举值： COMPLETE：转交易完成后解冻剩余冻结金额； NOT_COMPLETE：转交易完成后不解冻剩余冻结金额； 默认值为NOT_COMPLETE。 |
| `agreement_params` | `AgreementParams` | [`*AgreementParams`](#model-agreementparams) | JSON body | 未标注 | 周期扣款或代扣场景的协议参数。 |
| `product_code` | `ProductCode` | [`*ProductCode`](trade_model.go#L164) | JSON body | 未标注 | 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）： [FACE_TO_FACE_PAYMENT](https://opensupport.alipay.com/support/codelab/detail/486/487)：当面付产品； [CYCLE_PAY_AUTH](https://opensupport.alipay.com/support/codelab/detail/807/1419)：周期扣款产品； GENERAL_WITHHOLDING：代扣产品； [PRE_AUTH_ONLINE](https://opensupport.alipay.com/support/codelab/detail/712/1471)：支付宝预授权产品； PRE_AUTH：新当面资金授权产品； 默认值为FACE_TO_FACE_PAYMENT。 注意：非当面付产品使用本接口时，本参数必填。请传入对应产品码。<br>SDK 字段说明：产品码，默认 FACE_TO_FACE_PAYMENT。 |
| `seller_id` | `SellerID` | `*string` | JSON body | 未标注 | 卖家支付宝用户ID。 当需要指定收款账号时，通过该参数传入，如果该值为空，则默认为商户签约账号对应的支付宝用户ID。 收款账号优先级规则：门店绑定的收款账户&gt;请求传入的seller_id&gt;商户签约账号对应的支付宝用户ID； 注：直付通和机构间联场景下seller_id无需传入或者保持跟pid一致； 如果传入的seller_id与pid不一致，需要联系支付宝小二配置收款关系； 支付宝预授权和新当面资金授权场景下必填。<br>SDK 字段说明：卖家支付宝用户 ID；收款账户优先级见官方文档。 |
| `buyer_id` | `BuyerID` | `*string` | JSON body | 未标注 | 买家支付宝用户ID。 支付宝预授权和新当面资金授权场景下必填，其它场景不需要传入。<br>SDK 字段说明：买家支付宝用户 ID；预授权场景必填，其它场景无需传入。 |
| `buyer_open_id` | `BuyerOpenID` | `*string` | JSON body | 未标注 | 买家支付宝用户唯一标识 |
| `body` | `Body` | `*string` | JSON body | 未标注 | 订单附加信息。 如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示<br>SDK 字段说明：订单附加信息，将在异步通知、对账单中原样返回。 |
| `goods_detail` | `GoodsDetail` | [`[]GoodsDetail`](#model-goodsdetail) | JSON body | 未标注 | 订单包含的商品列表信息，json格式。<br>SDK 字段说明：订单包含的商品列表，上限 150 条。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商户操作员编号。 |
| `store_id` | `StoreID` | `*string` | JSON body | 未标注 | 商户门店编号。 指商户创建门店时输入的门店编号。 |
| `terminal_id` | `TerminalID` | `*string` | JSON body | 未标注 | 商户机具终端编号。 |
| `time_expire` | `TimeExpire` | `*string` | JSON body | 未标注 | 订单绝对超时时间。 格式为yyyy-MM-dd HH:mm:ss。超时时间范围：1m~15d。 注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。<br>SDK 字段说明：订单绝对超时时间（yyyy-MM-dd HH:mm:ss）；<br>与 TimeoutExpress 同传时优先使用本字段。 |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 订单相对超时时间。从交易创建时间开始计算。 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 当面付场景默认值为3h； 其它场景默认值为15d;<br>SDK 字段说明：订单相对超时时间（1m～15d）。 |
| `settle_info` | `SettleInfo` | [`*SettleInfo`](#model-settleinfo) | JSON body | 未标注 | 结算信息。 |
| `royalty_info` | `RoyaltyInfo` | [`*RoyaltyInfo`](#model-royaltyinfo) | JSON body | 未标注 | 分账信息。 |
| `extend_params` | `ExtendParams` | [`*ExtendParams`](#model-extendparams) | JSON body | 未标注 | 业务扩展参数。 |
| `business_params` | `BusinessParams` | [`*BusinessParams`](#model-businessparams) | JSON body | 未标注 | 与支付宝约定的直传业务信息。 |
| `query_options` | `QueryOptions` | [`[]PayQueryOption`](trade_model.go#L263) | JSON body | 未标注 | 返回参数选项。 商户通过传递该参数来定制同步需要额外返回的信息字段，数组格式。包括但不限于：["fund_bill_list","voucher_detail_list","enterprise_pay_info","discount_goods_detail","discount_amount","mdiscount_amount"]<br>SDK 字段说明：同步额外返回的信息字段。 |

注意事项：

- 当面付条码/刷脸支付需传入 Scene 与 AuthCode；周期扣款场景经
AgreementParams 传入协议号。写操作：SDK 不自动重试，超时后的
最终状态以查询接口为准。

<a id="api-create"></a>

### Create — 统一收单交易创建接口

- 官方接口：[`alipay.trade.create`](https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create)
- 路由：`POST /v3/alipay/trade/create`
- 请求类型：[CreateRequest](trade_request.go#L156)
- 响应类型：[CreateResponse](trade_response.go#L129)

商户通过该接口进行交易的创建下单

请求约束：CreateRequest 是 alipay.trade.trade.create（统一收单交易创建）的请求。
买家确认后完成支付；out_trade_no、total_amount、subject 与买家标识
为必选（买家标识三选一）。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyURL` | `*string` | JSON body | 未标注 | 本次请求的异步通知地址。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 商户订单号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。<br>SDK 字段说明：商户订单号：64 字符以内，仅支持字母、数字、下划线，<br>且需保证在商户端不重复。 |
| `total_amount` | `TotalAmount` | `string` | JSON body | 未标注 | 订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。<br>SDK 字段说明：订单总金额，两位小数，取值范围 [0.01,100000000]。 |
| `subject` | `Subject` | `string` | JSON body | 未标注 | 订单标题。 注意：256个字符以内，不可使用特殊字符，如 /，=，& 等。<br>SDK 字段说明：订单标题；不可使用 / = & 等特殊字符。 |
| `buyer_id` | `BuyerID` | `*string` | JSON body | 未标注 | 买家支付宝用户ID。 2088开头的16位纯数字，小程序场景下获取用户ID请参考：[用户授权](https://opendocs.alipay.com/mini/introduce/authcode); 其它场景下获取用户ID请参考：[网页授权获取用户信息](https://opendocs.alipay.com/open/284/106001/#s4); 注：交易的买家与卖家不能相同。<br>SDK 字段说明：买家的支付宝用户 ID（与 BuyerOpenID、BuyerLogonID 三选一）。 |
| `buyer_open_id` | `BuyerOpenID` | `*string` | JSON body | 未标注 | 买家支付宝用户唯一标识 |
| `buyer_logon_id` | `BuyerLogonID` | `*string` | JSON body | 未标注 | 买家支付宝登录账号。 建议通过open_id来传递买家信息。 |
| `product_code` | `ProductCode` | [`*ProductCode`](trade_model.go#L164) | JSON body | 未标注 | 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）： [FACE_TO_FACE_PAYMENT](https://opensupport.alipay.com/support/codelab/detail/486/487)：当面付产品； 默认值为FACE_TO_FACE_PAYMENT。 |
| `seller_id` | `SellerID` | `*string` | JSON body | 未标注 | 卖家支付宝用户ID。 当需要指定收款账号时，通过该参数传入，如果该值为空，则默认为商户签约账号对应的支付宝用户ID。 收款账号优先级规则：门店绑定的收款账户&gt;请求传入的seller_id&gt;商户签约账号对应的支付宝用户ID； 注：直付通和机构间联场景下seller_id无需传入或者保持跟pid一致； 如果传入的seller_id与pid不一致，需要联系支付宝小二配置收款关系；<br>SDK 字段说明：卖家支付宝用户 ID；收款账户优先级见官方文档。 |
| `body` | `Body` | `*string` | JSON body | 未标注 | 订单附加信息。 如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示<br>SDK 字段说明：订单附加信息，将在异步通知、对账单中原样返回。 |
| `goods_detail` | `GoodsDetail` | [`[]GoodsDetail`](#model-goodsdetail) | JSON body | 未标注 | 订单包含的商品列表信息，json格式。<br>SDK 字段说明：订单包含的商品列表，上限 150 条。 |
| `time_expire` | `TimeExpire` | `*string` | JSON body | 未标注 | 订单绝对超时时间。 格式为yyyy-MM-dd HH:mm:ss。 注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。<br>SDK 字段说明：订单绝对超时时间（yyyy-MM-dd HH:mm:ss）。 |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 订单相对超时时间。从交易创建时间开始计算。 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 当面付场景默认值为3h，如需指定，推荐设置5m及以上。 注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。<br>SDK 字段说明：订单相对超时时间（1m～15d）。 |
| `settle_info` | `SettleInfo` | [`*SettleInfo`](#model-settleinfo) | JSON body | 未标注 | 结算信息。 |
| `royalty_info` | `RoyaltyInfo` | [`*RoyaltyInfo`](#model-royaltyinfo) | JSON body | 未标注 | 分账信息。 |
| `agreement_sign_params` | `AgreementSignParams` | [`*SignParams`](#model-signparams) | JSON body | 未标注 | 支付并签约场景的签约参数；字段与必填<br>关系按官方商家扣款场景文档核验（app.pay 与 create 复用同一模型）。 |
| `extend_params` | `ExtendParams` | [`*ExtendParams`](#model-extendparams) | JSON body | 未标注 | 业务扩展参数。 |
| `business_params` | `BusinessParams` | [`*BusinessParams`](#model-businessparams) | JSON body | 未标注 | 与支付宝约定的直传业务信息。 |
| `query_options` | `QueryOptions` | [`[]PayQueryOption`](trade_model.go#L263) | JSON body | 未标注 | 返回参数选项。 商户通过传递该参数来定制需要额外返回的信息字段，数组格式。包括但不限于：["enterprise_pay_info","hyb_amount"]<br>SDK 字段说明：同步额外返回的信息字段。 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | JSON body | 未标注 | 商户的原始订单号 |
| `store_id` | `StoreID` | `*string` | JSON body | 未标注 | 商户门店编号。 指商户创建门店时输入的门店编号。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商户操作员编号。 |
| `terminal_id` | `TerminalID` | `*string` | JSON body | 未标注 | 商户机具终端编号。 |
| `disable_pay_channels` | `DisablePayChannels` | `*string` | JSON body | 未标注 | 禁用渠道,用户不可用指定渠道支付，多个渠道以逗号分割<br> 注，与enable_pay_channels互斥<br> [渠道列表](https://docs.open.alipay.com/common/wifww7)<br>SDK 字段说明：禁用渠道列表（逗号分割），与 EnablePayChannels 互斥。 |
| `enable_pay_channels` | `EnablePayChannels` | `*string` | JSON body | 未标注 | 指定支付渠道。 用户只能使用指定的渠道进行支付，多个渠道以逗号分割。 与disable_pay_channels互斥，支持传入的值：[渠道列表](https://docs.open.alipay.com/common/wifww7)。 注：如果传入了指定支付渠道，则用户只能用指定内的渠道支付，包括营销渠道也要指定才能使用。该参数可能导致用户支付受限，慎用。<br>SDK 字段说明：指定支付渠道列表（逗号分割）。 |
| `discountable_amount` | `DiscountableAmount` | `*string` | JSON body | 未标注 | 可打折金额。 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入了【可打折金额】、【不可打折金额】和【订单总金额】，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】。 如果订单金额全部参与优惠计算，则【可打折金额】和【不可打折金额】都无需传入。<br>SDK 字段说明：可打折金额；<br>与 TotalAmount 同传时须满足 TotalAmount = DiscountableAmount + UndiscountableAmount。 |
| `undiscountable_amount` | `UndiscountableAmount` | `*string` | JSON body | 未标注 | 不可打折金额。 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入了【可打折金额】、【不可打折金额】和【订单总金额】，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】。 如果订单金额全部参与优惠计算，则【可打折金额】和【不可打折金额】都无需传入。 |

注意事项：

- 返回交易号仅表示订单创建成功；买家支付结果以异步通知或查询为准。

<a id="api-cancel"></a>

### Cancel — 统一收单交易撤销接口

- 官方接口：[`alipay.trade.cancel`](https://opendocs.alipay.com/open-v3/cd7d54d2_alipay.trade.cancel?scene=common&pathHash=cb83add6)
- 路由：`POST /v3/alipay/trade/cancel`
- 请求类型：[CancelRequest](trade_request.go#L218)
- 响应类型：[CancelResponse](trade_response.go#L138)

支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。
注意：只有发生支付系统超时或者支付结果未知时可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】。

请求约束：CancelRequest 是 alipay.trade.cancel（统一收单交易撤销）的请求。
trade_no 与 out_trade_no 二选一且不能同时为空。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号，和商户订单号不能同时为空 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 原支付请求的商户订单号,和支付宝交易号不能同时为空 |

注意事项：

- 未支付交易触发关闭，已支付交易触发退款；RetryFlag=Y 表示需要重试。
撤销是写操作，SDK 不自动重试。

<a id="api-query"></a>

### Query — 统一收单交易查询

- 官方接口：[`alipay.trade.query`](https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query)
- 路由：`POST /v3/alipay/trade/query`
- 请求类型：[QueryRequest](trade_request.go#L13)
- 响应类型：[QueryResponse](trade_response.go#L5)

该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。
需要调用查询接口的情况：
当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；
调用支付接口后，返回系统错误或未知交易状态情况；
调用alipay.trade.pay，返回INPROCESS的状态；
调用alipay.trade.cancel之前，需确认支付状态；

请求约束：QueryRequest 是 alipay.trade.query 的请求。
out_trade_no 与 trade_no 二选一且不能同时为空；
同时存在时网关优先使用 trade_no。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。<br> trade_no,out_trade_no如果同时存在优先取trade_no |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号，和商户订单号不能同时为空 |
| `query_options` | `QueryOptions` | [`[]QueryOption`](trade_model.go#L185) | JSON body | 未标注 | 查询选项，商户传入该参数可定制本接口同步响应额外返回的信息字段，数组格式。<br>SDK 字段说明：同步额外返回的信息字段。 |

注意事项：

- 请求成功不代表支付成功，业务结果以响应中的 TradeStatus 为准；
未生成真实交易时 TradeNo 与 BuyerLogonID 不返回。

<a id="api-precreate"></a>

### PreCreate — 统一收单线下交易预创建

- 官方接口：[`alipay.trade.precreate`](https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate)
- 路由：`POST /v3/alipay/trade/precreate`
- 请求类型：[PreCreateRequest](trade_request.go#L24)
- 响应类型：[PreCreateResponse](trade_response.go#L59)

收银员通过收银台或商户后台调用支付宝接口，生成二维码后，展示给用户，由用户扫描二维码完成订单支付。

请求约束：PreCreateRequest 是 alipay.trade.precreate（统一收单线下交易预创建）的请求。
out_trade_no、total_amount、subject 为必选。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyURL` | `*string` | JSON body | 未标注 | 本次请求的异步通知地址。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 商户订单号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。<br>SDK 字段说明：商户订单号：64 字符以内，仅支持字母、数字、下划线，<br>且需保证在商户端不重复。 |
| `total_amount` | `TotalAmount` | `string` | JSON body | 未标注 | 订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。<br>SDK 字段说明：订单总金额，两位小数，取值范围 [0.01,100000000]。 |
| `subject` | `Subject` | `string` | JSON body | 未标注 | 订单标题。 注意：256个字符以内，不可使用特殊字符，如 /，=，& 等。<br>SDK 字段说明：订单标题；不可使用 / = & 等特殊字符。 |
| `product_code` | `ProductCode` | [`*ProductCode`](trade_model.go#L164) | JSON body | 未标注 | 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）： [FACE_TO_FACE_PAYMENT](https://opensupport.alipay.com/support/codelab/detail/486/487)：当面付产品； 默认值为FACE_TO_FACE_PAYMENT。<br>SDK 字段说明：产品码，默认 FACE_TO_FACE_PAYMENT。 |
| `seller_id` | `SellerID` | `*string` | JSON body | 未标注 | 卖家支付宝用户ID。 当需要指定收款账号时，通过该参数传入，如果该值为空，则默认为商户签约账号对应的支付宝用户ID。 收款账号优先级规则：门店绑定的收款账户&gt;请求传入的seller_id&gt;商户签约账号对应的支付宝用户ID； 注：直付通和机构间联场景下seller_id无需传入或者保持跟pid一致； 如果传入的seller_id与pid不一致，需要联系支付宝小二配置收款关系；<br>SDK 字段说明：卖家支付宝用户 ID；收款账户优先级见官方文档。 |
| `body` | `Body` | `*string` | JSON body | 未标注 | 订单附加信息。 如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示<br>SDK 字段说明：订单附加信息，将在异步通知、对账单中原样返回。 |
| `goods_detail` | `GoodsDetail` | [`[]GoodsDetail`](#model-goodsdetail) | JSON body | 未标注 | 订单包含的商品列表信息.json格式. 其它说明详见：“商品明细说明”<br>SDK 字段说明：订单包含的商品列表，上限 150 条。 |
| `time_expire` | `TimeExpire` | `*string` | JSON body | 未标注 | 订单绝对超时时间。 格式为yyyy-MM-dd HH:mm:ss。 注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。<br>SDK 字段说明：订单绝对超时时间（yyyy-MM-dd HH:mm:ss）；<br>与 TimeoutExpress 同传时优先使用本字段。 |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 订单相对超时时间。 从预下单请求时间开始计算。 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 当面付场景默认值为3h，如需指定，推荐设置5m及以上 注： 1. 二维码最长有效期是2小时，不管该参数传递的值是多少，超过2小时后二维码都将失效不能再进行扫码支付。 2. time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。<br>SDK 字段说明：订单相对超时时间（1m～15d）；<br>注意预下单二维码最长有效 2 小时。 |
| `settle_info` | `SettleInfo` | [`*SettleInfo`](#model-settleinfo) | JSON body | 未标注 | 结算信息。 |
| `extend_params` | `ExtendParams` | [`*ExtendParams`](#model-extendparams) | JSON body | 未标注 | 业务扩展参数。 |
| `business_params` | `BusinessParams` | [`*BusinessParams`](#model-businessparams) | JSON body | 未标注 | 与支付宝约定的直传业务信息。 |
| `undiscountable_amount` | `UndiscountableAmount` | `*string` | JSON body | 未标注 | 不可打折金额。 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入了【可打折金额】、【不可打折金额】和【订单总金额】，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】。 如果订单金额全部参与优惠计算，则【可打折金额】和【不可打折金额】都无需传入。 |
| `discountable_amount` | `DiscountableAmount` | `*string` | JSON body | 未标注 | 可打折金额。 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入了【可打折金额】、【不可打折金额】和【订单总金额】，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】。 如果订单金额全部参与优惠计算，则【可打折金额】和【不可打折金额】都无需传入。<br>SDK 字段说明：可打折金额；<br>与 TotalAmount 同传时须满足 TotalAmount = DiscountableAmount + UndiscountableAmount。 |
| `store_id` | `StoreID` | `*string` | JSON body | 未标注 | 商户门店编号。 指商户创建门店时输入的门店编号。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商户操作员编号。 |
| `terminal_id` | `TerminalID` | `*string` | JSON body | 未标注 | 商户机具终端编号。 |
| `disable_pay_channels` | `DisablePayChannels` | `*string` | JSON body | 未标注 | 禁用渠道,用户不可用指定渠道支付，多个渠道以逗号分割 注，与enable_pay_channels互斥 [渠道列表](https://docs.open.alipay.com/common/wifww7)<br>SDK 字段说明：禁用渠道列表（逗号分割），与 EnablePayChannels 互斥。 |
| `enable_pay_channels` | `EnablePayChannels` | `*string` | JSON body | 未标注 | 指定支付渠道。 用户只能使用指定的渠道进行支付，多个渠道以逗号分割。 与disable_pay_channels互斥，支持传入的值：[渠道列表](https://docs.open.alipay.com/common/wifww7)。 注：如果传入了指定支付渠道，则用户只能用指定内的渠道支付，包括营销渠道也要指定才能使用。该参数可能导致用户支付受限，慎用。<br>SDK 字段说明：指定支付渠道列表（逗号分割），与 DisablePayChannels 互斥。 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | JSON body | 未标注 | 商户的原始订单号 |
| `code_type` | `CodeType` | [`*CodeType`](trade_model.go#L172) | JSON body | 未标注 | 码类型，如果请求时传入了share_code，将在响应参数返回吱口令；如果为空，则默认为订单码<br>SDK 字段说明：码类型；为空默认为订单码。 |
| `query_options` | `QueryOptions` | [`[]PreCreateQueryOption`](trade_model.go#L204) | JSON body | 未标注 | 返回参数选项。 商户通过传递该参数来定制需要额外返回的信息字段，数组格式。包括但不限于：["hyb_amount","enterprise_pay_info"] |
| `bkagent_req_info` | `BkAgentReqInfo` | [`*BkAgentReqInfo`](#model-bkagentreqinfo) | JSON body | 未标注 | 间联交易下由收单机构上送的信息。 |

注意事项：

- 二维码有效时间固定 2 小时；同一订单号不能重复申请二维码；
qr_code 缺失时按业务响应处理，不代表预下单失败或成功。

<a id="api-close"></a>

### Close — 统一收单交易关闭接口

- 官方接口：[`alipay.trade.close`](https://opendocs.alipay.com/open-v3/48ea518b_alipay.trade.close?scene=common&pathHash=2c6cdbea)
- 路由：`POST /v3/alipay/trade/close`
- 请求类型：[CloseRequest](trade_request.go#L82)
- 响应类型：[CloseResponse](trade_response.go#L71)

用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。

请求约束：CloseRequest 是 alipay.trade.close 的请求。
trade_no 与 out_trade_no 二选一且不能同时为空；同传时以 trade_no 为准。
仅"等待买家付款"状态的交易可关单。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyURL` | `*string` | JSON body | 未标注 | 本次请求的异步通知地址。 |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 该交易在支付宝系统中的交易流水号。最短 16 位，最长 64 位。和out_trade_no不能同时为空，如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。<br>SDK 字段说明：支付宝交易号（最短 16 位）。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商家操作员编号 id，由商家自定义。 |

注意事项：

- 仅"等待买家付款"状态的交易可关单；SDK 不依据本地状态自动发起关单。

## 页面支付与 App 唤端产物

通过 `trade.NewPaymentProducts(credentials, opts...)` 创建产物构建器；此处仅在本地生成签名 URL、HTML 表单或订单串，不发起服务端 REST 请求。

| 官方接口 | 构建方法 | 请求类型 | 返回产物 |
| --- | --- | --- | --- |
| [alipay.trade.page.pay](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay) | `PagePayURL` / `PagePayForm` | `PagePayRequest` | 跳转 URL / 自动提交 HTML 表单 |
| [alipay.trade.wap.pay](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay) | `WapPayURL` / `WapPayForm` | `PagePayRequest` | 跳转 URL / 自动提交 HTML 表单 |
| [alipay.trade.app.pay](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay) | `AppPayOrderString` | `AppPayRequest` | 交支付宝客户端 SDK 的订单串 |

以下四个业务字段必填：`out_trade_no`、`total_amount`、`subject`、`product_code`。电脑网站支付产品码为 `FAST_INSTANT_TRADE_PAY`，手机网站支付为 `QUICK_WAP_WAY`，App 支付为 `QUICK_MSECURITY_PAY`。

通知地址通过 `core.WithNotifyURL` 传入；返回页面或客户端回调不能作为已付款凭据。金额、订单与收款方必须结合验签后的通知或订单查询核对。支付并签约使用 `AppPayRequest.SignParams`，依据[官方支付并签约文档](https://opendocs.alipay.com/open-v3/0jf33m)。

### PagePayRequest

类型定义：[PagePayRequest](trade_request.go#L296)。

| 协议字段 | Go 字段 | Go 类型 | 说明 |
| --- | --- | --- | --- |
| `out_trade_no` | `OutTradeNo` | `string` | 商户订单号：64 字符以内，仅支持字母、数字、下划线，<br>且需保证在商户端不重复。 |
| `total_amount` | `TotalAmount` | `string` | 订单总金额，两位小数，取值范围 [0.01,100000000]。 |
| `subject` | `Subject` | `string` | 订单标题；不可使用 / = & 等特殊字符。 |
| `product_code` | `ProductCode` | [`ProductCode`](trade_model.go#L164) | 产品码：电脑网站支付 FAST_INSTANT_TRADE_PAY，<br>手机网站支付 QUICK_WAP_WAY。 |
| `seller_id` | `SellerID` | `*string` | 卖家支付宝用户 ID。 |
| `body` | `Body` | `*string` | 订单附加信息，将在异步通知、对账单中原样返回。 |
| `goods_detail` | `GoodsDetail` | [`[]GoodsDetail`](#model-goodsdetail) | 订单包含的商品列表。 |
| `time_expire` | `TimeExpire` | `*string` | 订单绝对超时时间（yyyy-MM-dd HH:mm:ss）。 |
| `timeout_express` | `TimeoutExpress` | `*string` | 订单相对超时时间（1m～15d）。 |
| `qr_pay_mode` | `QRPayMode` | `*string` | 扫码支付方式（电脑网站支付）：0/1/2/3。 |
| `qr_pay_time_limit` | `QrPayTimeLimit` | `*string` | 二维码可扫码支付的最长时间。 |
| `settle_info` | `SettleInfo` | [`*SettleInfo`](#model-settleinfo) | 结算信息。 |
| `extend_params` | `ExtendParams` | [`*ExtendParams`](#model-extendparams) | 业务扩展参数。 |
| `business_params` | `BusinessParams` | [`*BusinessParams`](#model-businessparams) | 与支付宝约定的直传业务信息。 |
| `return_url` | `ReturnURL` | `*string` | 支付成功后的页面跳转地址；不是支付完成的可信证据。 |
| `passback_params` | `PassbackParams` | `*string` | 公用回传参数，异步通知时原样返回。 |

### AppPayRequest

类型定义：[AppPayRequest](trade_request.go#L336)。

| 协议字段 | Go 字段 | Go 类型 | 说明 |
| --- | --- | --- | --- |
| `out_trade_no` | `OutTradeNo` | `string` | 商户订单号。 |
| `total_amount` | `TotalAmount` | `string` | 订单总金额。 |
| `subject` | `Subject` | `string` | 订单标题。 |
| `product_code` | `ProductCode` | [`ProductCode`](trade_model.go#L164) | 产品码：App 支付为 QUICK_MSECURITY_PAY。 |
| `seller_id` | `SellerID` | `*string` | 卖家支付宝用户 ID。 |
| `body` | `Body` | `*string` | 订单附加信息。 |
| `goods_detail` | `GoodsDetail` | [`[]GoodsDetail`](#model-goodsdetail) | 订单包含的商品列表。 |
| `time_expire` | `TimeExpire` | `*string` | 订单绝对超时时间。 |
| `timeout_express` | `TimeoutExpress` | `*string` | 订单相对超时时间。 |
| `settle_info` | `SettleInfo` | [`*SettleInfo`](#model-settleinfo) | 结算信息。 |
| `royalty_info` | `RoyaltyInfo` | [`*RoyaltyInfo`](#model-royaltyinfo) | 分账信息。 |
| `extend_params` | `ExtendParams` | [`*ExtendParams`](#model-extendparams) | 业务扩展参数。 |
| `business_params` | `BusinessParams` | [`*BusinessParams`](#model-businessparams) | 与支付宝约定的直传业务信息。 |
| `sign_params` | `SignParams` | [`*SignParams`](#model-signparams) | 支付并签约场景的签约参数。 |
| `passback_params` | `PassbackParams` | `*string` | 公用回传参数。 |

## 支付异步通知

普通支付通知通过 `core/notify` 验签，并使用 `trade.NotificationDecoder` 解码为 `trade.Notification`；当前仅适用于 `notify_type=trade_status_sync`。商户仍须校验应用、收款方、订单与金额，完成幂等业务处理后再确认通知。

通知字段见 [Notification](trade_model.go#L347)；示例见 [api_notify_example_test.go](api_notify_example_test.go) 与[完整通知示例](../../examples/tradenotify/main.go)。协议依据：[支付通知说明](https://opendocs.alipay.com/open/194/103296)、[通知验签规则](https://opendocs.alipay.com/common/02mse7.md)。

## 嵌套请求参数

<a id="model-agreementparams"></a>

### AgreementParams

类型定义：[AgreementParams](trade_decodenotification_model.go#L276)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `string` | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ）<br>SDK 字段说明：支付宝系统中唯一标识用户签约记录的协议号。 |
| `apply_token` | `ApplyToken` | `*string` | 未标注 | 鉴权申请token，其格式和内容，由支付宝定义。在需要做支付鉴权校验时，该参数不能为空。<br>SDK 字段说明：鉴权申请 token；需要支付鉴权校验时不能为空。 |
| `auth_confirm_no` | `AuthConfirmNo` | `*string` | 未标注 | 鉴权确认码，在需要做支付鉴权校验时，该参数不能为空<br>SDK 字段说明：鉴权确认码；需要支付鉴权校验时不能为空。 |
| `deduct_permission` | `DeductPermission` | `*string` | 未标注 | 商户代扣扣款许可 |

<a id="model-bkagentreqinfo"></a>

### BkAgentReqInfo

类型定义：[BkAgentReqInfo](trade_decodenotification_model.go#L198)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `acq_code` | `AcqCode` | `*string` | 未标注 | 收单机构在清算组织登记或分配的机构代码 |
| `merch_code` | `MerchCode` | `*string` | 未标注 | 收单机构在清算组织登记的商户编码 |
| `device_type` | `DeviceType` | `*string` | 未标注 | 终端设备类型，受理方可参考终端注册时的设备类型填写。<br>SDK 字段说明：终端设备类型（如 01 ATM、02 传统 POS、03 mPOS）。 |
| `serial_num` | `SerialNum` | `*string` | 未标注 | 终端类型填写为 02、 03、04、05、06、08、09 或 10 时，必 须填写终端序列号。<br>SDK 字段说明：终端序列号；特定设备类型下必填。 |
| `location` | `Location` | `*string` | 未标注 | 终端设备实时经纬度信息，格式为纬度/经度，+表示北纬、东经，-表示南纬、西经。<br>SDK 字段说明：终端设备实时经纬度，格式为纬度/经度。 |
| `payment_type` | `PaymentType` | `*string` | 未标注 | 支付类型（如 agent_pay 智能体 SDK 支付）。 |

<a id="model-businessparams"></a>

### BusinessParams

类型定义：[BusinessParams](trade_decodenotification_model.go#L178)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `campus_card` | `CampusCard` | `*string` | 未标注 | 校园卡编号 |
| `card_type` | `CardType` | `*string` | 未标注 | 虚拟卡卡类型 |
| `actual_order_time` | `ActualOrderTime` | `*string` | 未标注 | 实际订单时间，在乘车码场景，传入的是用户刷码乘车时间<br>SDK 字段说明：实际订单时间（乘车码场景为刷码乘车时间）。 |
| `good_taxes` | `GoodTaxes` | `*string` | 未标注 | 商户传入的交易税费。需要落地风控使用 |
| `enterprise_pay_info` | `EnterprisePayInfo` | `*string` | 未标注 | 因公付业务信息<br>SDK 字段说明：因公付业务信息（JSON 文本）。 |
| `enterprise_pay_amount` | `EnterprisePayAmount` | `*string` | 未标注 | 因公付金额，单位：元。 |
| `mc_create_trade_ip` | `MCCreateTradeIP` | `*string` | 未标注 | 商户端创建订单的 IP，须上传正确的用户端外网 IP，支持 ipv4/ipv6 格式； mc_create_trade_ip和mcCreateTradeIp（旧）参数描述相同，首选mc_create_trade_ip入参，请勿重复入参； 如已入参mcCreateTradeIp（旧），无需新增入参mc_create_trade_ip。<br>SDK 字段说明：商户端创建订单的 IP（ipv4/ipv6）。 |
| `tiny_app_merchant_biz_type` | `TinyAppMerchantBizType` | `*string` | 未标注 | 当需要在 商家平台（b.alipay.com）—数据中心 中查看对应业务场景下的数据明细可以通过该参数传入 支持传入的值：[小程序细分业务场景](https://opendocs.alipay.com/mini/08by8k?pathHash=96e03412) |

<a id="model-extendparams"></a>

### ExtendParams

类型定义：[ExtendParams](trade_decodenotification_model.go#L162)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `sys_service_provider_id` | `SysServiceProviderID` | `*string` | 未标注 | 系统商编号<br> 该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID<br>SDK 字段说明：系统商编号（返佣数据提取依据）。 |
| `card_type` | `CardType` | `*string` | 未标注 | 卡类型 |
| `specified_seller_name` | `SpecifiedSellerName` | `*string` | 未标注 | 特殊场景下，允许商户指定交易展示的卖家名称 |
| `royalty_freeze` | `RoyaltyFreeze` | `*string` | 未标注 | 是否进行资金冻结，用于后续分账，true表示资金冻结，false或不传表示资金不冻结<br>SDK 字段说明：为 true 时进行资金冻结，用于后续分账。 |
| `tc_installment_order_id` | `TCInstallmentOrderID` | `*string` | 未标注 | 公域商品交易分期单ID，小程序交易组件订单特殊场景使用，请传入 订单分期接口(alipay.open.mini.order.installment.create)中返回的installment_order_id<br>SDK 字段说明：公域商品交易分期单 ID。 |
| `trade_component_order_id` | `TradeComponentOrderID` | `*string` | 未标注 | 公域商品交易业务订单ID<br>SDK 字段说明：公域商品交易业务订单 ID。 |

<a id="model-goodsdetail"></a>

### GoodsDetail

类型定义：[GoodsDetail](trade_decodenotification_model.go#L114)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_id` | `GoodsID` | `string` | 未标注 | 商品的编号，该参数传入支付券上绑定商品goods_id, 倘若无支付券需要消费，该字段传入商品最小粒度的商品ID（如：若商品有sku粒度，则传商户sku粒度的ID）<br>SDK 字段说明：商品编号（最小粒度 ID）。 |
| `goods_name` | `GoodsName` | `string` | 未标注 | 商品名称 |
| `quantity` | `Quantity` | `int64` | 未标注 | 商品数量 |
| `price` | `Price` | `string` | 未标注 | 商品单价，单位为元 |
| `goods_category` | `GoodsCategory` | `*string` | 未标注 | 商品类目 |
| `categories_tree` | `CategoriesTree` | `*string` | 未标注 | 商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用&#124;分割<br>SDK 字段说明：商品类目树，类目 id 用 &#124; 分割。 |
| `show_url` | `ShowURL` | `*string` | 未标注 | 商品的展示地址 |
| `out_item_id` | `OutItemID` | `*string` | 未标注 | 商家侧小程序商品ID，指商家提报给小程序商品库的商品。当前接口的extend_params.trade_component_order_id字段不为空时该字段必填，且与交易组件订单参数保持一致。了解小程序商品请参考：[https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7](https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7)<br>SDK 字段说明：商家侧小程序商品 ID。 |
| `out_sku_id` | `OutSkuID` | `*string` | 未标注 | 商家侧小程序商品ID，指商家提报给小程序商品库的商品。当前接口的extend_params.trade_component_order_id字段不为空时该字段必填，且与交易组件订单参数保持一致。了解小程序商品请参考：[https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7](https://opendocs.alipay.com/mini/06uila?pathHash=63b6fba7)<br>SDK 字段说明：商家侧小程序商品 SKU ID。 |

<a id="model-periodruleparams"></a>

### PeriodRuleParams

类型定义：[PeriodRuleParams](trade_decodenotification_model.go#L327)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `period_type` | `PeriodType` | `*string` | 未标注 | 周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。 DAY即扣款周期按天计，MONTH代表扣款周期按自然月。 与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；period_type为MONTH，period=3，则扣款周期为3个自然月。 自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），以此避免有些月份可能不存在对应日期的情况。<br>SDK 字段说明：周期类型：DAY 或 MONTH。 |
| `period` | `Period` | `*int64` | 未标注 | 周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period=90，则扣款周期为90天。<br>SDK 字段说明：周期数；与 PeriodType 组合确定扣款周期。 |
| `execute_time` | `ExecuteTime` | `*string` | 未标注 | 首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。精确到日，格式为yyyy-MM-dd 结合其他必填的扣款周期参数，会确定商户以后的扣款计划。发起扣款的时间需符合这里的扣款计划。<br>SDK 字段说明：首次执行时间（yyyy-MM-dd）；周期扣款产品必填。 |
| `single_amount` | `SingleAmount` | `*string` | 未标注 | 单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。<br>SDK 字段说明：单次扣款最大金额，单位为元。 |
| `total_amount` | `TotalAmount` | `*string` | 未标注 | 总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。 |
| `total_payments` | `TotalPayments` | `*int64` | 未标注 | 总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。 |
| `is_trial_mode` | `IsTrialMode` | `*string` | 未标注 | 若签约中传该参数，且值为 Y ，则表示本次签约有试用期。若传值为 N 或不传该参数，则表示本次签约无试用期。<br>SDK 字段说明：标识本次签约是否有试用期（Y/N）。 |

<a id="model-royaltydetailinfo"></a>

### RoyaltyDetailInfo

类型定义：[RoyaltyDetailInfo](trade_decodenotification_model.go#L214)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `royalty_type` | `RoyaltyType` | [`*RoyaltyType`](trade_model.go#L140) | 未标注 | 分账类型.<br>SDK 字段说明：分账类型；为空默认为普通分账 transfer。 |
| `trans_out` | `TransOut` | `*string` | 未标注 | 支出方账户。如果支出方账户类型为userId，本参数为支出方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果支出方类型为loginName，本参数为支出方的支付宝登录号。 泛金融类商户分账时，该字段不要上送。 |
| `trans_out_type` | `TransOutType` | [`*AccountType`](trade_model.go#L150) | 未标注 | 支出方账户类型。 |
| `trans_in_type` | `TransInType` | [`*AccountType`](trade_model.go#L150) | 未标注 | 收入方账户类型。 |
| `trans_in` | `TransIn` | `*string` | 未标注 | 收入方账户。如果收入方账户类型为userId，本参数为收入方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果收入方类型为cardAliasNo，本参数为收入方在支付宝绑定的卡编号；如果收入方类型为loginName，本参数为收入方的支付宝登录号；如果收入方类型为loginName，本参数为收入方的支付宝登录号；如果收入方类型为openId，本参数为收入方的支付宝openId信息；如果收入方类型为alipayWallet，本参数为收入方在支付宝绑定的场景钱包id |
| `trans_in_id` | `TransInID` | `*string` | 未标注 | 分账转入方的支付宝userId<br>SDK 字段说明：分账转入方的支付宝 userId。 |
| `amount` | `Amount` | `*string` | 未标注 | 分账的金额，单位为元 |
| `amount_percentage` | `AmountPercentage` | `*int64` | 未标注 | 分账信息中分账百分比。取值范围为大于0，少于或等于100的整数。<br>SDK 字段说明：分账百分比，取值范围 (0,100] 的整数。 |
| `desc` | `Desc` | `*string` | 未标注 | 分账描述 |
| `royalty_scene` | `RoyaltyScene` | `*string` | 未标注 | 可选值：达人佣金、平台服务费、技术服务费、其他<br>SDK 字段说明：分账场景（达人佣金、平台服务费等）。 |
| `trans_in_name` | `TransInName` | `*string` | 未标注 | 分账收款方姓名，上送则进行姓名与支付宝账号的一致性校验，校验不一致则分账失败。不上送则不进行姓名校验<br>SDK 字段说明：分账收款方姓名；上送则进行一致性校验。 |

<a id="model-royaltyinfo"></a>

### RoyaltyInfo

类型定义：[RoyaltyInfo](trade_decodenotification_model.go#L288)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `royalty_type` | `RoyaltyType` | [`*RoyaltyType`](trade_model.go#L140) | 未标注 | 分账类型<br> 卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）。<br>SDK 字段说明：分账类型；目前只支持 ROYALTY。 |
| `royalty_detail_infos` | `RoyaltyDetailInfos` | [`[]RoyaltyDetailInfo`](#model-royaltydetailinfo) | 未标注 | 分账明细的信息，可以描述多条分账指令，json数组。<br>SDK 字段说明：分账明细信息，可描述多条分账指令。 |

<a id="model-settledetailinfo"></a>

### SettleDetailInfo

类型定义：[SettleDetailInfo](trade_decodenotification_model.go#L142)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `trans_in_type` | `TransInType` | [`AccountType`](trade_model.go#L150) | 未标注 | 结算收款方的账户类型。 |
| `trans_in` | `TransIn` | `*string` | 未标注 | 结算收款方。当结算收款方类型是cardAliasNo时，本参数为用户在支付宝绑定的卡编号；结算收款方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；当结算收款方类型是loginName时，本参数为用户的支付宝登录号；当结算收款方类型是defaultSettle时，本参数不能传值，保持为空。<br>SDK 字段说明：结算收款方；类型为 defaultSettle 时不能传值。 |
| `amount` | `Amount` | `*string` | 未标注 | 结算的金额，单位为元。在创建订单和支付接口时必须和交易金额相同。在结算确认接口时必须等于交易金额减去已退款金额。直付通账期模式下，如使用部分结算能力、传递了actual_amount字段，则忽略本字段的校验、可不传。<br>SDK 字段说明：结算的金额，创建订单时必须与交易金额相同。 |
| `summary_dimension` | `SummaryDimension` | `*string` | 未标注 | 结算汇总维度，按照这个维度汇总成批次结算，由商户指定。 目前需要和结算收款方账户类型为cardAliasNo配合使用<br>SDK 字段说明：结算汇总维度，需与 cardAliasNo 配合使用。 |
| `settle_entity_id` | `SettleEntityID` | `*string` | 未标注 | 结算主体标识。当结算主体类型为SecondMerchant时，为二级商户的SecondMerchantID；当结算主体类型为Store时，为门店的外标。 |
| `settle_entity_type` | `SettleEntityType` | `*string` | 未标注 | 结算主体类型。<br>SDK 字段说明：结算主体类型：SecondMerchant 或 Store。 |
| `actual_amount` | `ActualAmount` | `*string` | 未标注 | 仅在直付通账期模式下，当一笔交易需要分多次发起部分确认结算时使用，表示本次确认结算的实际结算金额。传递本字段后，原amount字段不再生效，结算金额以本字段为准。如已经发生过部分确认结算、不传递本字段则默认按剩余待结算金额一次性结算。<br>SDK 字段说明：仅用于直付通账期模式的部分确认结算。 |
| `settle_period_time` | `SettlePeriodTime` | `*string` | 未标注 | 超期自动确认结算时间（1d～365d），仅账期结算模式有效。 |

<a id="model-settleinfo"></a>

### SettleInfo

类型定义：[SettleInfo](trade_decodenotification_model.go#L136)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `settle_detail_infos` | `SettleDetailInfos` | [`[]SettleDetailInfo`](#model-settledetailinfo) | 未标注 | 结算详细信息，json数组，目前只支持一条。<br>SDK 字段说明：结算详细信息，目前只支持一条。 |

<a id="model-signaccessparams"></a>

### SignAccessParams

类型定义：[SignAccessParams](trade_decodenotification_model.go#L320)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `channel` | `Channel` | `*string` | 未标注 | 传参示例： 1. ALIPAYAPP （钱包h5页面签约） 2. QRCODE(扫码签约) 3. QRCODEORSMS(扫码签约或者短信签约)<br>SDK 字段说明：签约渠道，如 ALIPAYAPP（钱包 h5 页面签约）、<br>QRCODE（扫码签约）、QRCODEORSMS（扫码或短信签约）。 |

<a id="model-signparams"></a>

### SignParams

类型定义：[SignParams](trade_decodenotification_model.go#L297)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | 未标注 | 商家和支付宝签约的产品码。 商家扣款产品传入固定值：GENERAL_WITHHOLDING<br>SDK 字段说明：商家和支付宝签约的产品码；商家扣款场景传<br>GENERAL_WITHHOLDING。 |
| `personal_product_code` | `PersonalProductCode` | `*string` | 未标注 | 个人签约产品码，商户和支付宝签约时确定。 |
| `sign_scene` | `SignScene` | `*string` | 未标注 | 协议签约场景，商户和支付宝签约时确定，商户可咨询技术支持。 |
| `external_agreement_no` | `ExternalAgreementNo` | `*string` | 未标注 | 商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。 格式规则：支持大写小写字母和数字，最长32位。 商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。 |
| `external_logon_id` | `ExternalLogonID` | `*string` | 未标注 | 用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示 |
| `access_params` | `AccessParams` | [`*SignAccessParams`](#model-signaccessparams) | 未标注 | 签约接入参数（如 ALIPAYAPP、QRCODE）。 |
| `period_rule_params` | `PeriodRuleParams` | [`*PeriodRuleParams`](#model-periodruleparams) | 未标注 | 周期扣款规则参数。 |
| `sign_notify_url` | `SignNotifyURL` | `*string` | 未标注 | 签约成功后商户用于接收异步通知的地址。如果不传入，签约与支付的异步通知都会发到外层notify_url参数传入的地址；如果外层也未传入，签约与支付的异步通知都会发到商户appid配置的网关地址。<br>SDK 字段说明：签约成功后接收异步通知的地址。 |
| `effect_time` | `EffectTime` | `*string` | 未标注 | 设置签约请求的有效时间，单位为秒。如传入600，商户发起签约请求到用户进入支付宝签约页面的时间差不能超过10分钟。<br>SDK 字段说明：签约请求的有效时间（秒）。 |
