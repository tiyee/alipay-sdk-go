# orderinfo 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/orderinfo`。本包提供 1 个支付宝 V3 REST 操作。

通过 `orderinfo.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`OrderInfoSync`](#api-orderinfosync) | `alipay.trade.orderinfo.sync` |

## 使用示例

- [OrderInfoSync](api_orderinfo_orderinfosync_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-orderinfosync"></a>

### OrderInfoSync — 支付宝订单信息同步接口

- 官方接口：[`alipay.trade.orderinfo.sync`](https://opendocs.alipay.com/open-v3/f8b82729_alipay.trade.orderinfo.sync)
- 路由：`POST /v3/alipay/trade/orderinfo/sync`
- 请求类型：[OrderInfoSyncRequest](orderinfo_request.go#L10)
- 响应类型：[OrderInfoSyncResponse](orderinfo_response.go#L4)

该接口用于商户向支付宝同步该笔订单相关业务信息

请求约束：OrderInfoSyncRequest 是 alipay.trade.orderinfo.sync（订单信息同步）
的请求。out_request_no 由支付宝做幂等控制：相同请求号再次请求
不会更新订单信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 外部请求号，商家自定义且保证商家系统中唯一。需要注意的是，支付宝会对外部请求号做幂等控制，如果一笔交易再次使用相同的外部请求号发起请求，订单信息不会更新。<br>SDK 字段说明：外部请求号，商家自定义且保证唯一。 |
| `biz_type` | `BizType` | `string` | JSON body | 未标注 | 交易信息同步对应的业务类型，具体值与支付宝约定；<br> 信用授权场景下传CREDIT_AUTH<br> 信用代扣场景下传CREDIT_DEDUCT<br>SDK 字段说明：业务类型，具体值与支付宝约定（如 CREDIT_AUTH、<br>CREDIT_DEDUCT）。 |
| `order_biz_info` | `OrderBizInfo` | `*string` | JSON body | 未标注 | 商户传入同步信息，具体值要和支付宝约定；用于芝麻信用租车、单次授权等信息同步场景，格式为json格式。 状态枚举如下： COMPLETE：同步用户已履约 适用场景：发起扣款后，芝麻生成待履约记录，如果用户通过其他方式完成订单支付，请反馈该状态，芝麻将完结待履约记录对用户形成一条良好履约记录；同步该状态时需要同步调用 取消扣款 接口关闭交易订单。 CLOSED： 同步履约已取消 适用场景：发起扣款后，芝麻生成待履约记录，如果发现该笔扣款无效需要取消，请反馈该状态，芝麻将取消用户待履约记录；同步该状态时需要同步调用 取消扣款 接口关闭交易订单。<br>SDK 字段说明：同步信息内容，格式与支付宝约定（JSON 文本）。 |
| `orig_request_no` | `OrigRequestNo` | `*string` | JSON body | 未标注 | 原始业务请求单号。如对某一次退款进行履约时，该字段传退款时的退款请求号<br>SDK 字段说明：原始业务请求单号（如履约对象退款的退款请求号）。 |

注意事项：

- 相同 OutRequestNo 重复请求不会更新订单信息（网关幂等）。
