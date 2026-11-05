# tradecomplain 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/merchant/tradecomplain`。本包提供 5 个支付宝 V3 REST 操作。

通过 `tradecomplain.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`TradeComplainBatchQuery`](#api-tradecomplainbatchquery) | `alipay.merchant.tradecomplain.batchquery` |
| [`TradeComplainFeedbackSubmit`](#api-tradecomplainfeedbacksubmit) | `alipay.merchant.tradecomplain.feedback.submit` |
| [`TradeComplainQuery`](#api-tradecomplainquery) | `alipay.merchant.tradecomplain.query` |
| [`TradeComplainReplySubmit`](#api-tradecomplainreplysubmit) | `alipay.merchant.tradecomplain.reply.submit` |
| [`TradeComplainSupplementSubmit`](#api-tradecomplainsupplementsubmit) | `alipay.merchant.tradecomplain.supplement.submit` |

## 使用示例

- [TradeComplainBatchQuery](api_tradecomplain_tradecomplainbatchquery_example_test.go)
- [TradeComplainFeedbackSubmit](api_tradecomplain_tradecomplainfeedbacksubmit_example_test.go)
- [TradeComplainQuery](api_tradecomplain_tradecomplainquery_example_test.go)
- [TradeComplainReplySubmit](api_tradecomplain_tradecomplainreplysubmit_example_test.go)
- [TradeComplainSupplementSubmit](api_tradecomplain_tradecomplainsupplementsubmit_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-tradecomplainbatchquery"></a>

### TradeComplainBatchQuery — 查询交易投诉列表

- 官方接口：[`alipay.merchant.tradecomplain.batchquery`](https://opendocs.alipay.com/open-v3/152a9215_alipay.merchant.tradecomplain.batchquery)
- 路由：`POST /v3/alipay/merchant/tradecomplain/batchquery`
- 请求类型：[MerchantTradeComplainBatchQueryRequest](tradecomplain_request.go#L4)
- 响应类型：[MerchantTradeComplainBatchQueryResponse](tradecomplain_response.go#L8)

开发者根据其应用id查询用户在其应用（小程序、生活号）的投诉单列表记录。投诉单列表以分页查询方式返回。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `begin_time` | `BeginTime` | `*string` | JSON body | 未标注 | 查询开始时间 时间格式：yyyy-MM-dd HH:mm:ss 若不填写，则默认当前时间30天前。 注意：begin_time和end_time时间跨度最大一年 |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 查询结束时间 时间格式：yyyy-MM-dd HH:mm:ss 若不填写，则默认当前时间。 注意：begin_time和end_time时间跨度最大一年 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 当前页，从1开始 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页条数，默认10条，最多支持20条 |
| `status` | `Status` | `*string` | JSON body | 未标注 | 状态如下 待处理：MERCHANT_PROCESSING 已处理：MERCHANT_FEEDBACKED 投诉完结：FINISHED 投诉关闭：CANCELLED 客服处理中：PLATFORM_PROCESSING 客服处理完结：PLATFORM_FINISH 投诉关闭：CLOSED |
| `target_infos` | `TargetInfos` | [`*[]TargetInfo`](#model-targetinfo) | JSON body | 未标注 | 应用信息 |

<a id="api-tradecomplainfeedbacksubmit"></a>

### TradeComplainFeedbackSubmit — 商家处理交易投诉

- 官方接口：[`alipay.merchant.tradecomplain.feedback.submit`](https://opendocs.alipay.com/open-v3/da1eaf4f_alipay.merchant.tradecomplain.feedback.submit)
- 路由：`POST /v3/alipay/merchant/tradecomplain/feedback`
- 请求类型：[MerchantTradeComplainFeedbackSubmitRequest](tradecomplain_request.go#L20)
- 响应类型：[MerchantTradeComplainFeedbackSubmitResponse](tradecomplain_response.go#L22)

商户请求交易投诉处理API，上传处理类型、处理内容描述及图片凭证。其中上传图片凭证需要先通过"商户上传处理图片"接口获取图片id，再将图片id设置到接口请求中。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `complain_event_id` | `ComplainEventId` | `*string` | JSON body | 未标注 | 支付宝侧投诉单号 |
| `feedback_code` | `FeedbackCode` | `*string` | JSON body | 未标注 | 小程序交易反馈类目ID： 02:通过其他方式退款; 03:已发货; 04:其他; 05:已完成售后服务; 06:非我方责任范围; 安心付/先享后付反馈类目ID： 194:反馈处理结果; |
| `feedback_content` | `FeedbackContent` | `*string` | JSON body | 未标注 | 反馈内容，字数不超过200个字 |
| `feedback_images` | `FeedbackImages` | `*string` | JSON body | 未标注 | 商家处理投诉时反馈凭证的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取） |
| `operator` | `Operator` | `*string` | JSON body | 未标注 | 处理投诉人，字数不超过6个字 |

<a id="api-tradecomplainquery"></a>

### TradeComplainQuery — 查询单条交易投诉详情

- 官方接口：[`alipay.merchant.tradecomplain.query`](https://opendocs.alipay.com/open-v3/cecf0505_alipay.merchant.tradecomplain.query)
- 路由：`GET /v3/alipay/merchant/tradecomplain/{complain_event_id}`
- 请求类型：[MerchantTradeComplainQueryRequest](tradecomplain_request.go#L54)
- 响应类型：[MerchantTradeComplainQueryResponse](tradecomplain_response.go#L25)

开发者根据支付宝侧投诉单号查询该笔投诉详细记录。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `complain_event_id` | `ComplainEventId` | `string` | 路径 | 是 | 支付宝侧投诉单号 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。

<a id="api-tradecomplainreplysubmit"></a>

### TradeComplainReplySubmit — 商家留言回复

- 官方接口：[`alipay.merchant.tradecomplain.reply.submit`](https://opendocs.alipay.com/open-v3/5851abbd_alipay.merchant.tradecomplain.reply.submit)
- 路由：`POST /v3/alipay/merchant/tradecomplain/reply`
- 请求类型：[MerchantTradeComplainReplySubmitRequest](tradecomplain_request.go#L34)
- 响应类型：[MerchantTradeComplainReplySubmitResponse](tradecomplain_response.go#L61)

投诉单状态在“商家待处理”或“商家已反馈”或“商家处理超时”状态下，若有内容需要补充说明，可使用此接口提交文字或图片

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `complain_event_id` | `ComplainEventId` | `*string` | JSON body | 未标注 | 支付宝侧投诉单号 |
| `reply_content` | `ReplyContent` | `*string` | JSON body | 未标注 | 回复留言内容 |
| `reply_images` | `ReplyImages` | `*string` | JSON body | 未标注 | 商家处理问题反馈时的回复凭证的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取） |

<a id="api-tradecomplainsupplementsubmit"></a>

### TradeComplainSupplementSubmit — 商家补充凭证

- 官方接口：[`alipay.merchant.tradecomplain.supplement.submit`](https://opendocs.alipay.com/open-v3/bb0a25ee_alipay.merchant.tradecomplain.supplement.submit)
- 路由：`POST /v3/alipay/merchant/tradecomplain/supplement`
- 请求类型：[MerchantTradeComplainSupplementSubmitRequest](tradecomplain_request.go#L44)
- 响应类型：[MerchantTradeComplainSupplementSubmitResponse](tradecomplain_response.go#L64)

投诉进入平台处理中状态时，为保证客服能准确判定责任方，可使用此接口补充文字或图片说明问题及处理经过

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `complain_event_id` | `ComplainEventId` | `*string` | JSON body | 未标注 | 支付宝侧投诉单号 |
| `supplement_content` | `SupplementContent` | `*string` | JSON body | 未标注 | 回复内容，最多不超过200个字 |
| `supplement_images` | `SupplementImages` | `*string` | JSON body | 未标注 | 商家补充凭证时的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取） |

## 嵌套请求参数

<a id="model-targetinfo"></a>

### TargetInfo

类型定义：[TargetInfo](tradecomplain_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `target_id` | `TargetId` | `*string` | 未标注 | 应用id，例如小程序id |
| `target_type` | `TargetType` | `*string` | 未标注 | 应用类型 小程序传入：APPID 生活号传入：PUBLICID |
