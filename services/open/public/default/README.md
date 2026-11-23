# open/public/default 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/default`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `defaultapi.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_default_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicDefaultExtensionCreate`](#api-publicdefaultextensioncreate) | [`alipay.open.public.default.extension.create`](https://opendocs.alipay.com/apis/05wvst) | 默认扩展区创建接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicdefaultextensioncreate"></a>

### PublicDefaultExtensionCreate — 默认扩展区创建接口

- 官方接口：[`alipay.open.public.default.extension.create`](https://opendocs.alipay.com/apis/05wvst)
- 路由：`POST /v3/alipay/open/public/default/extension/create`
- 请求类型：[PublicDefaultExtensionCreateRequest](defaultapi_request.go#L5)
- 响应类型：[PublicDefaultExtensionCreateResponse](defaultapi_response.go#L5)

开发者可使用此接口创建默认扩展区

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `areas` | `Areas` | [`*[]ExtensionArea`](#model-extensionarea) | JSON body | 未标注 | 默认扩展区列表，最多包含3个扩展区 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-extensionarea"></a>

### ExtensionArea

类型定义：[ExtensionArea](defaultapi_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goto_url` | `GotoUrl` | `*string` | 未标注 | 跳转链接，当type为"image"时必传，必须是https或alipays开头的url链接。已废弃 |
| `height` | `Height` | `*int64` | 未标注 | 扩展区高度，当type值为"h5"时必填，取值范围为200-500的整数,单位是px |
| `name` | `Name` | `*string` | 未标注 | 扩展区名字 |
| `type` | `Type` | `*string` | 未标注 | 扩展区类型，当前支持的值为h5，表示扩展区内容类型为h5。 |
| `url` | `Url` | `*string` | 未标注 | 扩展区url，传入h5页面url，必须是https开头的链接 |
