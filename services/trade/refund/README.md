# refund 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/refund`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `refund.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_refund_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`QueryRefund`](#api-queryrefund) | [`alipay.trade.fastpay.refund.query`](https://opendocs.alipay.com/open-v3/b9ef37bd_alipay.trade.fastpay.refund.query?scene=common&pathHash=f071344e) | 统一收单交易退款查询 | `POST` |
| [`Refund`](#api-refund) | [`alipay.trade.refund`](https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund?scene=common&pathHash=00ab12b7) | 统一收单交易退款接口 | `POST` |
| [`RefundApply`](#api-refundapply) | [`alipay.trade.refund.apply`](https://opendocs.alipay.com/apis/05wbvf) | 异步退款受理接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-queryrefund"></a>

### QueryRefund — 统一收单交易退款查询

- 官方接口：[`alipay.trade.fastpay.refund.query`](https://opendocs.alipay.com/open-v3/b9ef37bd_alipay.trade.fastpay.refund.query?scene=common&pathHash=f071344e)
- 路由：`POST /v3/alipay/trade/fastpay/refund/query`
- 请求类型：[QueryRefundRequest](refund_request.go#L55)
- 响应类型：[QueryRefundResponse](refund_response.go#L68)

商户可使用该接口查询自已通过alipay.trade.refund提交的退款请求是否执行成功。

请求约束：QueryRefundRequest 是 alipay.trade.fastpay.refund.query 的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 退款请求号。 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的商户订单号。<br>SDK 字段说明：退款请求号（必选）：退款时传入的请求号；<br>退款时未传入则为创建交易时的商户订单号。 |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号。 和商户订单号不能同时为空 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 商户订单号。 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no<br>SDK 字段说明：商户订单号；与 TradeNo 同传时优先使用 trade_no。 |
| `query_options` | `QueryOptions` | [`[]RefundQueryOption`](refund_model.go#L64) | JSON body | 未标注 | 查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。枚举支持： refund_detail_item_list：本次退款使用的资金渠道； gmt_refund_pay：退款执行成功的时间； deposit_back_info：银行卡冲退信息；<br>SDK 字段说明：同步额外返回的信息字段。 |
| `org_pid` | `OrgPID` | `*string` | JSON body | 未标注 | 银行间联模式的机构 PID。 |

注意事项：

- 请求成功仅代表查询操作成功；仅当响应中 RefundStatus 为
REFUND_SUCCESS 时表示退款成功，字段未返回表示退款请求未收到
或退款失败。官方建议退款发起后间隔 10 秒以上再查询。

<a id="api-refund"></a>

### Refund — 统一收单交易退款接口

- 官方接口：[`alipay.trade.refund`](https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund?scene=common&pathHash=00ab12b7)
- 路由：`POST /v3/alipay/trade/refund`
- 请求类型：[RefundRequest](refund_request.go#L12)
- 响应类型：[RefundResponse](refund_response.go#L10)

当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，支付宝将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。
交易超过约定时间（签约时设置的可退款时间）的订单无法进行退款。
支付宝退款支持单笔交易分多次退款，多次退款需要提交原支付订单的订单号和设置不同的退款请求号。一笔退款失败后重新提交，要保证重试时退款请求号不能变更，防止该笔交易重复退款。
同一笔交易累计提交的退款金额不能超过原始交易总金额。

请求约束：RefundRequest 是 alipay.trade.refund（交易退款）的请求。
out_trade_no 与 trade_no 二选一；out_request_no 在部分退款时必选，
且重试时必须保持不变以防重复退款。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `refund_amount` | `RefundAmount` | `string` | JSON body | 未标注 | 退款金额。 需要退款的金额，该金额不能大于订单金额，单位为元，支持两位小数。 注：如果正向交易使用了营销，该退款金额包含营销金额，支付宝会按业务规则分配营销和买家自有资金分别退多少，默认优先退买家的自有资金。如交易总金额100元，用户支付时使用了80元自有资金和20元无资金流的营销券，商家实际收款80元。如果首次请求退款60元，则60元全部从商家收款资金扣除退回给用户自有资产；如果再请求退款40元，则从商家收款资金扣除20元退回用户资产以及把20元的营销券退回给用户（券是否可再使用取决于券的规则配置）。<br>SDK 字段说明：本次退款金额，不能大于订单金额。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 商户订单号。 订单支付时传入的商户订单号，商家自定义且保证商家系统中唯一。与支付宝交易号 trade_no 不能同时为空。 |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号。 和商户订单号 out_trade_no 不能同时为空，两者同时存在时，优先取值trade_no<br>SDK 字段说明：支付宝交易号；与 OutTradeNo 同传时优先使用本字段。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 退款请求号。 标识一次退款请求，需要保证在交易号下唯一，如需部分退款，则此参数必传。 注：针对同一次退款请求，如果调用接口失败或异常了，重试时需要保证退款请求号不能变更，防止该笔交易重复退款。支付宝会保证同样的退款请求号多次请求只会退一次。<br>SDK 字段说明：退款请求号，需保证在交易号下唯一；<br>部分退款时必选；同请求号多次请求只退一次。 |
| `refund_reason` | `RefundReason` | `*string` | JSON body | 未标注 | 退款原因说明。 商家自定义，将在会在商户和用户的pc退款账单详情中展示 |
| `refund_goods_detail` | `RefundGoodsDetail` | [`[]RefundGoodsDetail`](#model-refundgoodsdetail) | JSON body | 未标注 | 退款包含的商品列表信息 |
| `refund_royalty_parameters` | `RefundRoyaltyParameters` | [`[]trade.RoyaltyDetailInfo`](../trade_decodenotification_model.go#L240) | JSON body | 未标注 | 退分账明细信息。 |
| `query_options` | `QueryOptions` | [`[]RefundQueryOption`](refund_model.go#L64) | JSON body | 未标注 | 查询选项。 商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。<br>SDK 字段说明：同步额外返回的信息字段。 |
| `related_settle_confirm_no` | `RelatedSettleConfirmNo` | `*string` | JSON body | 未标注 | 针对账期交易，在确认结算后退款的话，需要指定确认结算时的结算单号。<br>SDK 字段说明：用于账期交易确认结算后退款。 |
| `goods_detail` | `GoodsDetail` | [`[]trade.GoodsDetail`](../trade_decodenotification_model.go#L116) | JSON body | 未标注 | 退款包含的商品列表信息，Json格式。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商户的操作员编号。 |
| `org_pid` | `OrgPID` | `*string` | JSON body | 未标注 | 银行间联模式的机构 PID。 |
| `refund_advance_account` | `RefundAdvanceAccount` | `*string` | JSON body | 未标注 | 指定垫资退款账号主体，可能是用户id或登录名。 |
| `refund_advance_account_type` | `RefundAdvanceAccountType` | `*string` | JSON body | 未标注 | 指定垫资退款账号的类型，取值范围：default，userId，loginName。 |
| `refund_currency` | `RefundCurrency` | `*string` | JSON body | 未标注 | 订单退款币种信息。 |
| `refund_trans_out` | `RefundTransOut` | `*string` | JSON body | 未标注 | 指定退款账号，对应指定账号退款合约中的约定账号。 |
| `refund_trans_out_type` | `RefundTransOutType` | `*string` | JSON body | 未标注 | 指定退款账号类型。 |
| `store_id` | `StoreID` | `*string` | JSON body | 未标注 | 商户门店编号，由商家自定义。 |
| `terminal_id` | `TerminalID` | `*string` | JSON body | 未标注 | 商户的终端编号。 |

注意事项：

- 请求成功仅代表退款请求受理：FundChange=Y 表示退款成功，
N 或未返回时需通过 QueryRefund 进一步确认。
重试必须保持 OutRequestNo 与 RefundAmount 不变，防止重复退款；
SDK 不自动重试退款。

<a id="api-refundapply"></a>

### RefundApply — 异步退款受理接口

- 官方接口：[`alipay.trade.refund.apply`](https://opendocs.alipay.com/apis/05wbvf)
- 路由：`POST /v3/alipay/trade/refund/apply`
- 请求类型：[RefundApplyRequest](refund_request.go#L72)
- 响应类型：[RefundApplyResponse](refund_response.go#L130)

商户通过此接口提交异步退款请求，支付宝受理后，将会异步进行退款处理。退款完成后异步通知商户退款结果，商户可以通过alipay.trade.fastpay.refund.query接口查询退款状态。

请求约束：RefundApplyRequest 是 alipay.trade.refund.apply（异步退款受理）的请求。
out_request_no 标识一次退款请求，同一笔交易多次退款需保证唯一；
请求成功仅代表受理，结果以 RefundStatus 与退款查询为准。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyURL` | `*string` | JSON body | 未标注 | 本次请求的异步通知地址。 |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号，和商户订单号不能同时为空<br>SDK 字段说明：支付宝交易号（与 OutTradeNo 二选一）。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 订单支付时传入的商户订单号,不能和 trade_no同时为空 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。<br>SDK 字段说明：退款请求号；同一笔交易多次退款需保证唯一，<br>部分退款必传。 |
| `refund_amount` | `RefundAmount` | `string` | JSON body | 未标注 | 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数<br>SDK 字段说明：退款金额，不能大于订单金额。 |
| `refund_reason` | `RefundReason` | `*string` | JSON body | 未标注 | 退款的原因说明<br>SDK 字段说明：退款原因说明。 |
| `operator_id` | `OperatorID` | `*string` | JSON body | 未标注 | 商户的操作员编号<br>SDK 字段说明：商户操作员编号。 |
| `store_id` | `StoreID` | `*string` | JSON body | 未标注 | 商户的门店编号<br>SDK 字段说明：商户门店编号。 |
| `terminal_id` | `TerminalID` | `*string` | JSON body | 未标注 | 商户的终端编号<br>SDK 字段说明：商户终端编号。 |
| `refund_fund_details` | `RefundFundDetails` | [`[]RefundFundDetail`](#model-refundfunddetail) | JSON body | 未标注 | 退款资金明细<br> 注：目前只有SETTLE_OFFLINE_GROUP产品的退款支持该字段<br>SDK 字段说明：退款资金明细；仅 SETTLE_OFFLINE_GROUP<br>产品支持。 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 商户传入业务信息，具体值要和支付宝约定，格式为json格式。 |

注意事项：

- RefundStatus 为 REFUND_PROCESSING 时退款仍在处理中，需通过
QueryRefund 确认终态；SDK 不自动重试。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-refundfunddetail"></a>

### RefundFundDetail

类型定义：[RefundFundDetail](refund_model.go#L160)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `type` | `Type` | `*string` | 未标注 | 描述资金明细类型<br> 详：若type为paySerialNo<br> 则funds参数中体现的都为对应支付流水的退款明细<br>SDK 字段说明：资金明细类型（如 paySerialNo）。 |
| `funds` | `Funds` | `[]string` | 未标注 | 退款资金明细<br> 详：<br> 若type为paySerialNo则funds为<br> [{"paySerialNo":"支付流水1","refundFee":"退款金额1"},{"paySerialNo":"支付流水2","refundFee":"退款金额2"}]<br>SDK 字段说明：退款资金明细内容；格式依 Type 而定。 |
| `trans_in` | `TransIn` | `*string` | 未标注 | 收入方账户<br> 为空则原路退回<br>SDK 字段说明：收入方账户；为空则原路退回。 |
| `trans_in_type` | `TransInType` | `*string` | 未标注 | 收入方账户类型 |

<a id="model-refundgoodsdetail"></a>

### RefundGoodsDetail

类型定义：[RefundGoodsDetail](refund_model.go#L90)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_id` | `GoodsID` | `string` | 未标注 | 商品编号。 对应支付时传入的goods_id<br>SDK 字段说明：商品编号，对应支付时传入的 goods_id。 |
| `refund_amount` | `RefundAmount` | `string` | 未标注 | 该商品的退款总金额，单位为元 |
| `out_certificate_no_list` | `OutCertificateNoList` | `[]string` | 未标注 | 外部商品凭证编号列表 |
| `out_item_id` | `OutItemID` | `*string` | 未标注 | 商家侧小程序商品ID，对应支付时传入的out_item_id<br>SDK 字段说明：商家侧小程序商品 ID。 |
| `out_sku_id` | `OutSkuID` | `*string` | 未标注 | 商家侧小程序商品sku ID，对应支付时传入的out_sku_id<br>SDK 字段说明：商家侧小程序商品 SKU ID。 |
