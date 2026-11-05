# overseas 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/overseas`。本包提供 1 个支付宝 V3 REST 操作。

通过 `overseas.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`OverseasSettle`](#api-overseassettle) | `alipay.trade.overseas.settle` |

## 使用示例

- [OverseasSettle](api_overseas_overseassettle_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-overseassettle"></a>

### OverseasSettle — 统一境内收单跨境结算接口

- 官方接口：[`alipay.trade.overseas.settle`](https://opendocs.alipay.com/apis/05v7c4)
- 路由：`POST /v3/alipay/trade/overseas/settle`
- 请求类型：[OverseasSettleRequest](overseas_request.go#L10)
- 响应类型：[OverseasSettleResponse](overseas_response.go#L4)

境内收单跨境结算

请求约束：OverseasSettleRequest 是 alipay.trade.overseas.settle（境内收单跨境
结算）的请求。写操作：out_request_no 用于幂等控制。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 外部请求号，开发者自行生成并保证唯一性，作为业务幂等性控制<br>SDK 字段说明：外部请求号；调用方生成并保证唯一，用于幂等控制。 |
| `amount` | `Amount` | `string` | JSON body | 未标注 | 跨境结算的人民币金额，单位为元；等于交易金额 (实际为实收金额)，加上净补差金额，减去已退款金额，减去净收费金额，再减去净分账金额， |
| `foreign_settle_currency` | `ForeignSettleCurrency` | `string` | JSON body | 未标注 | 跨境结算的外币币种 |
| `extend_params` | `ExtendParams` | [`*OverseasSettleExtendParams`](#model-overseassettleextendparams) | JSON body | 未标注 | 跨境结算扩展参数。 |

注意事项：

- OutRequestNo 用于幂等控制；写操作，SDK 不自动重试。

## 嵌套请求参数

<a id="model-overseassettleextendparams"></a>

### OverseasSettleExtendParams

类型定义：[OverseasSettleExtendParams](overseas_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_detail` | `GoodsDetail` | `*string` | 未标注 | 商品明细列表<br>SDK 字段说明：商品明细列表（JSON 文本）。 |
