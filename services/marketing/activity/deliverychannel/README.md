# marketing/activity/deliverychannel 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/deliverychannel`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `deliverychannel.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_deliverychannel_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityDeliveryChannelQuery`](#api-activitydeliverychannelquery) | [`alipay.marketing.activity.deliverychannel.query`](https://opendocs.alipay.com/apis/05v7cm) | 查询可推广渠道 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activitydeliverychannelquery"></a>

### ActivityDeliveryChannelQuery — 查询可推广渠道

- 官方接口：[`alipay.marketing.activity.deliverychannel.query`](https://opendocs.alipay.com/apis/05v7cm)
- 路由：`POST /v3/alipay/marketing/activity/deliverychannel/query`
- 请求类型：[ActivityDeliveryChannelQueryRequest](deliverychannel_request.go#L5)
- 响应类型：[ActivityDeliveryChannelQueryResponse](deliverychannel_response.go#L5)

查询可推广渠道。该接口会返回可以投放的所有渠道的详细信息。
其中每个渠道的boothCode含义：boothCode表达的是某个渠道可以投放的展位码。例如：支付结果页PAY_RESULT
其中每个渠道的channel含义：表达的是某个固定的渠道。例如：在boothCode = PAY_RESULT 的情况下， channel = 商户的pid。 该channel就是一个可以投放的渠道，投放后可以在对应的商户的支付结果页看到优惠券。后续创建投放时根据返回的 channel 选择投放渠道

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*DeliveryAgencyMerchantInfo`](#model-deliveryagencymerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `booth_code` | `BoothCode` | `*string` | JSON body | 未标注 | 展位码。 boothCode含义：boothCode表达的是某个渠道可以投放的展位码。例如：支付结果页PAY_RESULT 枚举值： PAY_RESULT：支付结果页。 后续新增可投放的展位后，会在文档中新增描述信息 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 查询的页码。 特别说明： 页码从1开始。 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页查询的个数，最大值为 100。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-deliveryagencymerchantinfo"></a>

### DeliveryAgencyMerchantInfo

类型定义：[DeliveryAgencyMerchantInfo](deliverychannel_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_type` | `BusinessType` | `*string` | 未标注 | 合作业务类型。 |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户信息 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。 限制: 服务商代运营模式必传 如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。 因此要求该SMID必须绑定相应的支付宝PID。 |
