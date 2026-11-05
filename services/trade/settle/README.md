# settle 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/settle`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `settle.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_settle_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`SettleConfirm`](#api-settleconfirm) | [`alipay.trade.settle.confirm`](https://opendocs.alipay.com/apis/0hp5md) | 统一收单确认结算接口 | `POST` |
| [`SettleReceivablesQuery`](#api-settlereceivablesquery) | [`alipay.trade.settle.receivables.query`](https://opendocs.alipay.com/pre-apis/05vlu7) | 待结算查询 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-settleconfirm"></a>

### SettleConfirm — 统一收单确认结算接口

- 官方接口：[`alipay.trade.settle.confirm`](https://opendocs.alipay.com/apis/0hp5md)
- 路由：`POST /v3/alipay/trade/settle/confirm`
- 请求类型：[SettleConfirmRequest](settle_request.go#L10)
- 响应类型：[SettleConfirmResponse](settle_response.go#L4)

用户付款后，平台上的卖家发货，用户收到商品确认收货后，平台通过确认结算接口，将交易款结算给指定的卖家账户。
单笔结算到卖家余额的场景支持分账冻结。在确认结算时，传入冻结标识 royalty_freeze=true，资金结算到卖家支付宝余额后，会先冻结在不可用余额中，后续分账将自动从冻结余额划出，平台商可调用分账完结释放冻结资金，系统默认在冻结30天后自动解冻。解冻后，仍可以发起分账。为保证分账成功率，请尽可能在分账冻结期内分账。

请求约束：SettleConfirmRequest 是 alipay.trade.settle.confirm（统一收单确认结算）
的请求。trade_no、out_request_no、settle_info 必选。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 确认结算请求流水号，开发者自行生成并保证唯一性，作为业务幂等性控制<br>SDK 字段说明：确认结算请求流水号；由调用方生成并保证唯一，<br>用于业务幂等控制。 |
| `settle_info` | `SettleInfo` | [`trade.SettleInfo`](../trade_decodenotification_model.go#L144) | JSON body | 未标注 | 结算信息。 |
| `extend_params` | `ExtendParams` | [`*SettleConfirmExtendParams`](#model-settleconfirmextendparams) | JSON body | 未标注 | 确认结算扩展参数。 |

注意事项：

- OutRequestNo 用于幂等控制；写操作，SDK 不自动重试。

<a id="api-settlereceivablesquery"></a>

### SettleReceivablesQuery — 待结算查询

- 官方接口：[`alipay.trade.settle.receivables.query`](https://opendocs.alipay.com/pre-apis/05vlu7)
- 路由：`POST /v3/alipay/trade/settle/receivables/query`
- 请求类型：[SettleReceivablesQueryRequest](settle_request.go#L25)
- 响应类型：[SettleReceivablesQueryResponse](settle_response.go#L14)

对外提供待结算资金查询的能力。

请求约束：SettleReceivablesQueryRequest 是 alipay.trade.settle.receivables.query
（待结算查询）的请求：直付通账期模式查询单笔交易时 trade_no 必传；
其他场景传入商户信息与产品码。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝交易号，当该笔交易为直付通账期模式，查询该笔交易待确认结算金额时必传<br>SDK 字段说明：支付宝交易号；直付通账期模式查询单笔交易时必传。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 外部请求号，32个字符以内，可包含字母、数字、下划线。<br>SDK 字段说明：外部请求号：32 字符以内，可包含字母、数字、下划线。 |
| `biz_product` | `BizProduct` | `*string` | JSON body | 未标注 | 收单产品码，商家和支付宝签约的产品码 |
| `merchant_info` | `MerchantInfo` | [`*SettleEntity`](#model-settleentity) | JSON body | 未标注 | 结算主体信息。 |
| `query_his_date` | `QueryHisDate` | `*string` | JSON body | 未标注 | 当trade_no不为空时，该字段不生效。 查询历史日期，格式为 yyyyMMdd ，取值范围为昨日起至往前30日内； 不传入时，查询实时待结算余额返回； 传入过去某一天日期，查询对应日期的日终待结算余额返回（注意：日常场景下，昨日日终待结算余额只可在当天 02:00 后查询，在当天 02:00 前查询返回查询错误；大促场景下昨日日终可查时间会适当延后）； 传入过去某一天非近30天内，返回参数错误。<br>SDK 字段说明：查询历史日期（yyyyMMdd，昨日起 30 日内）；<br>trade_no 不为空时该字段不生效。 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 扩展参数。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-settleconfirmextendparams"></a>

### SettleConfirmExtendParams

类型定义：[SettleConfirmExtendParams](settle_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `royalty_freeze` | `RoyaltyFreeze` | `*string` | 未标注 | 是否进行资金冻结，用于后续分账，true表示冻结，false或不传表示不冻结<br>SDK 字段说明：标识是否进行资金冻结用于后续分账（true/false）。 |

<a id="model-settleentity"></a>

### SettleEntity

类型定义：[SettleEntity](settle_model.go#L10)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `settle_entity_type` | `SettleEntityType` | `*string` | 未标注 | 结算主体类型 SecondMerchant：结算主体为二级商户；Store：结算主体为商户门店<br>SDK 字段说明：结算主体类型：SecondMerchant 或 Store。 |
| `settle_entity_id` | `SettleEntityID` | `*string` | 未标注 | 结算主体账号 当结算主体类型为SecondMerchant，本参数为二级商户的SecondMerchantID；当结算主体类型为MerchantStore，本参数为商户门店ID。 |
| `settle_entity_biz_type` | `SettleEntityBizType` | `*string` | 未标注 | 结算主体资产业务类型，settled 表示已结算资产<br>SDK 字段说明：结算主体资产业务类型（如 settled 已结算资产）。 |
