# commerce/iot 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/iot`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `iot.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_iot_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`IoTDeviceReportUpload`](#api-iotdevicereportupload) | [`alipay.commerce.iot.device.report.upload`](https://opendocs.alipay.com/pre-apis/05xk46) | 设备检测报告上传 | `POST` |
| [`IoTSdarttoolMessageQuery`](#api-iotsdarttoolmessagequery) | [`alipay.commerce.iot.sdarttool.message.query`](https://opendocs.alipay.com/pre-apis/06bpqm) | 消息查询 | `GET` |
| [`IoTSdarttoolMessageSend`](#api-iotsdarttoolmessagesend) | [`alipay.commerce.iot.sdarttool.message.send`](https://opendocs.alipay.com/pre-apis/06bzwt) | IOT消息服务 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-iotdevicereportupload"></a>

### IoTDeviceReportUpload — 设备检测报告上传

- 官方接口：[`alipay.commerce.iot.device.report.upload`](https://opendocs.alipay.com/pre-apis/05xk46)
- 路由：`POST /v3/alipay/commerce/iot/device/report/upload`
- 请求类型：[IoTDeviceReportUploadRequest](iot_request.go#L5)
- 响应类型：[IoTDeviceReportUploadResponse](iot_response.go#L5)

设备检测报告上传

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `report_content` | `ReportContent` | `*string` | JSON body | 未标注 | 设备检测报告详情 |
| `sn` | `Sn` | `*string` | JSON body | 未标注 | 设备唯一标识 |

<a id="api-iotsdarttoolmessagequery"></a>

### IoTSdarttoolMessageQuery — 消息查询

- 官方接口：[`alipay.commerce.iot.sdarttool.message.query`](https://opendocs.alipay.com/pre-apis/06bpqm)
- 路由：`GET /v3/alipay/commerce/iot/sdarttool/message/query`
- 请求类型：[IoTSdarttoolMessageQueryRequest](iot_request.go#L41)
- 响应类型：[IoTSdarttoolMessageQueryResponse](iot_response.go#L11)

isv通过消息id查询消息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `message_no` | `MessageNo` | `*string` | query | 否 | 消息ID |

<a id="api-iotsdarttoolmessagesend"></a>

### IoTSdarttoolMessageSend — IOT消息服务

- 官方接口：[`alipay.commerce.iot.sdarttool.message.send`](https://opendocs.alipay.com/pre-apis/06bzwt)
- 路由：`POST /v3/alipay/commerce/iot/sdarttool/message/send`
- 请求类型：[IoTSdarttoolMessageSendRequest](iot_request.go#L13)
- 响应类型：[IoTSdarttoolMessageSendResponse](iot_response.go#L37)

ISV通过此接口向IOT设备推送消息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bi_da` | `BiDa` | `*bool` | JSON body | 未标注 | 消息是否必达(immediate_msg为false时此字段设置有效)，消息过期时间顺延3*24h. 建议使用expire_time设置消息过期时间 |
| `device_query_type` | `DeviceQueryType` | `*string` | JSON body | 未标注 | 设备查询条件类型 ；SUPPLIERID_SN:supplierid+sn ; ITEMID_SN:itemid + sn; SN:sn |
| `immediate_msg` | `ImmediateMsg` | `*bool` | JSON body | 未标注 | 是否即时消息; true-校验设备是否在线，false-不校验设备是否在线 |
| `item_id` | `ItemId` | `*string` | JSON body | 未标注 | 产品ID |
| `msg_content` | `MsgContent` | `*string` | JSON body | 未标注 | 消息内容 |
| `msg_content_type` | `MsgContentType` | `*string` | JSON body | 未标注 | 消息内容格式 |
| `msg_expire` | `MsgExpire` | `*int64` | JSON body | 未标注 | 消息过期时间戳(ms)， 默认为当前时间顺延24h有效。当设备在线时消息服务过期之前尝试推送。最大过期时间顺延3*24h |
| `msg_priority` | `MsgPriority` | `*int64` | JSON body | 未标注 | 消息优先级（immediate_msg为false此字段设置有效）数字越大优先发送。 |
| `msg_type` | `MsgType` | `*string` | JSON body | 未标注 | 消息类型： xpaas_common-小程序消息； ruyi_ordermsg-如意订单消息 |
| `service_id` | `ServiceId` | `*string` | JSON body | 未标注 | 服务ID，当小程序消费消息时必传(传入的是小程序ID) |
| `sn` | `Sn` | `*string` | JSON body | 未标注 | 设备sn号 |
| `supplier_id` | `SupplierId` | `*string` | JSON body | 未标注 | 厂商ID |
