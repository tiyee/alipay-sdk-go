# open/public/qrcode 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/qrcode`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `qrcode.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_qrcode_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicQRCodeCreate`](#api-publicqrcodecreate) | [`alipay.open.public.qrcode.create`](https://opendocs.alipay.com/apis/05wuhi) | 带参推广二维码接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicqrcodecreate"></a>

### PublicQRCodeCreate — 带参推广二维码接口

- 官方接口：[`alipay.open.public.qrcode.create`](https://opendocs.alipay.com/apis/05wuhi)
- 路由：`POST /v3/alipay/open/public/qrcode/create`
- 请求类型：[PublicQRCodeCreateRequest](qrcode_request.go#L5)
- 响应类型：[PublicQRCodeCreateResponse](qrcode_response.go#L5)

为了满足用户渠道推广分析的需要，公众平台提供了生成带参数二维码的接口。使用该接口可以获得多个带不同场景值的二维码，用户扫描后，公众号可以接收到事件推送。目前有2种类型的二维码，分别是临时二维码、和永久二维码，前者有过期时间，最大为1800秒。 每次创建二维码ticket需要提供一个开发者自行设定的参数（scene_id）。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `code_info` | `CodeInfo` | [`*CodeInfo`](#model-codeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `code_type` | `CodeType` | `*string` | JSON body | 未标注 | 二维码类型，目前只支持两种类型：<br> TEMP：临时的（默认）；<br> PERM：永久的 |
| `expire_second` | `ExpireSecond` | `*string` | JSON body | 未标注 | 临时码过期时间，以秒为单位，最大不超过1800秒；<br> 永久码置空 |
| `show_logo` | `ShowLogo` | `*string` | JSON body | 未标注 | 二维码中间是否显示服务窗logo，Y：显示；N：不显示（默认） |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-codeinfo"></a>

### CodeInfo

类型定义：[CodeInfo](qrcode_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goto_url` | `GotoUrl` | `*string` | 未标注 | 跳转URL，扫码关注服务窗后会直接跳转到此URL |
| `scene` | `Scene` | [`*Scene`](#model-scene) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-scene"></a>

### Scene

类型定义：[Scene](qrcode_model.go#L12)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `scene_id` | `SceneId` | `*string` | 未标注 | 场景Id，最长32位，英文字母、数字以及下划线，开发者自定义 |
