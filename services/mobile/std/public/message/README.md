# message 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/mobile/std/public/message`。本包提供 1 个支付宝 V3 REST 操作。

通过 `message.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`StdPublicMessageCustomSend`](#api-stdpublicmessagecustomsend) | `alipay.mobile.std.public.message.custom.send` |

## 使用示例

- [StdPublicMessageCustomSend](api_message_stdpublicmessagecustomsend_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-stdpublicmessagecustomsend"></a>

### StdPublicMessageCustomSend — 异步单发消息

- 官方接口：[`alipay.mobile.std.public.message.custom.send`](https://opendocs.alipay.com/pre-apis/05wult)
- 路由：`POST /v3/alipay/mobile/std/public/message/custom/send`
- 请求类型：[MobileStdPublicMessageCustomSendRequest](message_request.go#L4)
- 响应类型：[MobileStdPublicMessageCustomSendResponse](message_response.go#L8)

商户给单个用户发送消息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_content` | `BizContent` | `*string` | JSON body | 未标注 | 业务内容，其中包括消息类型msgType、消息体和消息接受人toUserId三大块，具体参见“表1-2 服务窗单发客服消息的biz_content参数说明”。 |
