# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/order`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `order.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_order_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`OrderPay`](#api-orderpay) | [`alipay.trade.order.pay`](https://opendocs.alipay.com/open-v3/c2751e3b_alipay.trade.order.pay) | 统一收单交易订单支付接口 | `POST` |
| [`OrderSettle`](#api-ordersettle) | [`alipay.trade.order.settle`](https://opendocs.alipay.com/open-v3/34402227_alipay.trade.order.settle) | 统一收单交易结算接口 | `POST` |
| [`OrderSettleQuery`](#api-ordersettlequery) | [`alipay.trade.order.settle.query`](https://opendocs.alipay.com/open-v3/3af46ab7_alipay.trade.order.settle.query) | 交易分账查询接口 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-orderpay"></a>

### OrderPay — 统一收单交易订单支付接口

- 官方接口：[`alipay.trade.order.pay`](https://opendocs.alipay.com/open-v3/c2751e3b_alipay.trade.order.pay)
- 路由：`POST /v3/alipay/trade/order/pay`
- 请求类型：[OrderPayRequest](order_request.go#L12)
- 响应类型：[OrderPayResponse](order_response.go#L5)

商户用已创建的交易进行支付（需要使用该接口请联系接口owner）

请求约束：OrderPayRequest 是 alipay.trade.order.pay（统一收单交易订单支付）的请求，
用于对已存在订单发起履约支付。trade_no 必选；
out_request_no 用于幂等控制。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号 |
| `total_amount` | `TotalAmount` | `string` | JSON body | 未标注 | 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]<br> 注：<br> 1.用于校验与已存交易中的金额是否相等<br>SDK 字段说明：订单总金额；用于校验与已存交易中的金额是否一致。 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 商户请求号，标识一次请求的唯一id，用于幂等控制。<br>SDK 字段说明：商户请求号，标识一次请求的唯一 id，用于幂等控制。 |
| `buyer_id` | `BuyerID` | `*string` | JSON body | 未标注 | 买家的支付宝用户id<br> 注：<br> 1.用于校验与已存交易中的买家是否相等<br>SDK 字段说明：买家的支付宝用户 id；用于校验与已存交易的买家是否一致。 |
| `buyer_open_id` | `BuyerOpenID` | `*string` | JSON body | 未标注 | 买家支付宝用户唯一标识 注： 1.用于校验与已存交易中的买家是否相等 |
| `buyer_pay_detail` | `BuyerPayDetail` | [`[]BuyerPayDetail`](#model-buyerpaydetail) | JSON body | 未标注 | 买家支付明细<br> 目前支持的支付渠道为：<br> - offline_pos：本次买家使用的是pos刷卡支付<br> 注：<br> 各支付工具金额总和=订单总金额 |
| `fulfillment_amount` | `FulfillmentAmount` | `*string` | JSON body | 未标注 | 本次履约支付金额，单位为元，精确到小数点后两位。履约支付场景才需要传入。<br>SDK 字段说明：本次履约支付金额；履约支付场景传入。 |
| `is_async_pay` | `IsAsyncPay` | `*bool` | JSON body | 未标注 | 是否异步支付，传入true时，表明本次期望走异步支付，会先将支付请求受理下来，再异步推进。商户可以通过交易的异步通知或者轮询交易的状态来确定最终的交易结果<br>SDK 字段说明：标识是否走异步支付。 |
| `order_pay_mode` | `OrderPayMode` | `*string` | JSON body | 未标注 | 订单支付模式。特殊支付场景才需要传入。具体传参需与支付宝约定。 CREDIT_FULFILLMENT_ZM表示基于芝麻授信的履约支付模式，比如芝麻先用后付产品。<br>SDK 字段说明：订单支付模式；特殊场景与支付宝约定。 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 销售产品码 |
| `advance_payment_type` | `AdvancePaymentType` | `*string` | JSON body | 未标注 | 垫资支付模式。支付时需要垫资的场景才传入。具体传参需与支付宝约定。 CREDIT_FULFILLMENT_ZM 表示先用后付产品履约动作支持芝麻垫资。<br>SDK 字段说明：垫资支付模式；特殊场景与支付宝约定。 |

注意事项：

- OutRequestNo 用于幂等控制；异步支付以异步通知或查询确认结果。

<a id="api-ordersettle"></a>

### OrderSettle — 统一收单交易结算接口

- 官方接口：[`alipay.trade.order.settle`](https://opendocs.alipay.com/open-v3/34402227_alipay.trade.order.settle)
- 路由：`POST /v3/alipay/trade/order/settle`
- 请求类型：[OrderSettleRequest](order_request.go#L49)
- 响应类型：[OrderSettleResponse](order_response.go#L23)

用于在卖家交易成功之后，基于交易订单，进行卖家与第三方（如供应商或平台商）的资金再分配。一般用于第三方从卖家抽佣场景。

请求约束：OrderSettleRequest 是 alipay.trade.order.settle（统一收单交易结算）的请求。
trade_no、out_request_no 必选；单独调用分账完结时可不传分账明细。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝订单号<br>SDK 字段说明：支付宝交易号。 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 结算请求流水号，由商家自定义。64个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复。<br>SDK 字段说明：结算请求流水号：64 字符以内，仅可包含字母、数字、<br>下划线，需保证在商户端不重复。 |
| `royalty_parameters` | `RoyaltyParameters` | [`[]RoyaltyDetailInfo`](../README.md#model-royaltydetailinfo) | JSON body | 未标注 | 分账明细信息。单独调用分账完结时，可以不传此参数。其他场景必传。 注意：商家分账场景下分账收入方 trans_in 只支持支付宝账户，不支持使用 cardAliasNo 卡编号。<br>SDK 字段说明：分账明细信息；单独调用分账完结时可不传，<br>其他场景必传。 |
| `royalty_mode` | `RoyaltyMode` | [`*RoyaltyMode`](order_model.go#L17) | JSON body | 未标注 | 分账模式，目前有两种分账同步执行sync，分账异步执行async，不传默认同步执行<br>SDK 字段说明：分账模式：sync（同步执行）或 async（异步执行），<br>默认同步。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 操作员 ID，商家自定义操作员编号。 |
| `extend_params` | `ExtendParams` | [`*OrderSettleExtendParams`](#model-ordersettleextendparams) | JSON body | 未标注 | 分账扩展参数（完结标记等）。 |

注意事项：

- OutRequestNo 用于幂等；分账结果以 SettleNo 配合查询接口确认。

<a id="api-ordersettlequery"></a>

### OrderSettleQuery — 交易分账查询接口

- 官方接口：[`alipay.trade.order.settle.query`](https://opendocs.alipay.com/open-v3/3af46ab7_alipay.trade.order.settle.query)
- 路由：`GET /v3/alipay/trade/order/settle/query`
- 请求类型：[OrderSettleQueryRequest](order_request.go#L70)
- 响应类型：[OrderSettleQueryResponse](order_response.go#L31)

根据分账请求号查询交易分账结果

请求约束：OrderSettleQueryRequest 是 alipay.trade.order.settle.query
（交易分账查询，V3 GET）的请求：settle_no 或（out_request_no +
trade_no）二选一。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `settle_no` | `SettleNo` | `*string` | query | 否 | 支付宝分账请求单号，传入该字段，无需再传外部请求号和支付宝交易号<br>SDK 字段说明：支付宝分账请求单号；传入后无需再传外部请求号与交易号。 |
| `out_request_no` | `OutRequestNo` | `*string` | query | 否 | 调用分账接口时指定的外部请求号。分账查询时需要和支付宝交易号一起传入<br>SDK 字段说明：分账接口指定的外部请求号；需与 TradeNo 一起传入。 |
| `trade_no` | `TradeNo` | `*string` | query | 否 | 支付宝交易号，传入该字段，需要和外部请求号一起传入<br>SDK 字段说明：支付宝交易号；需与 OutRequestNo 一起传入。 |

注意事项：

- settle_no 或（out_request_no + trade_no）二选一。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-buyerpaydetail"></a>

### BuyerPayDetail

类型定义：[BuyerPayDetail](order_model.go#L32)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `pay_channel` | `PayChannel` | `string` | 未标注 | 支付渠道<br>SDK 字段说明：支付渠道（如 offline_pos）。 |
| `amount` | `Amount` | `string` | 未标注 | 买家支付金额，单位：元。<br>SDK 字段说明：买家支付金额，单位为元。 |

<a id="model-ordersettleextendparams"></a>

### OrderSettleExtendParams

类型定义：[OrderSettleExtendParams](order_model.go#L40)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `royalty_finish` | `RoyaltyFinish` | `*string` | 未标注 | "冻结分账场景下生效，其他场景传入无效。代表该交易分账是否完结，可选值：true/false，不传默认为false。true：代表分账完结，则本次分账处理完成后会把该笔交易的剩余冻结金额全额解冻。false：代表分账未完结。"<br>SDK 字段说明：标识该交易分账是否完结（true/false）；<br>true 表示本次分账处理完成后解冻剩余冻结金额。 |
| `royalty_finish_amount` | `RoyaltyFinishAmount` | `*string` | 未标注 | 部分分账完结的金额，单位为元。若该参数不为空，则会按照该参数金额进行部分解冻。注：该参数不为空时，royalty_finish需要传入true<br>SDK 字段说明：部分分账完结的金额；传入时 RoyaltyFinish 须为 true。 |
