# open/mini/resource 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/resource`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `resource.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_resource_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniResourcePromotionSourceNotify`](#api-miniresourcepromotionsourcenotify) | [`alipay.open.mini.resource.promotionsource.notify`](https://opendocs.alipay.com/pre-apis/05xhr9) | 提供给淘联盟推广位信息接口 | `POST` |
| [`MiniResourceRecordNotify`](#api-miniresourcerecordnotify) | [`alipay.open.mini.resource.record.notify`](https://opendocs.alipay.com/pre-apis/05xqyn) | 提供给淘联盟备案信息接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniresourcepromotionsourcenotify"></a>

### MiniResourcePromotionSourceNotify — 提供给淘联盟推广位信息接口

- 官方接口：[`alipay.open.mini.resource.promotionsource.notify`](https://opendocs.alipay.com/pre-apis/05xhr9)
- 路由：`POST /v3/alipay/open/mini/resource/promotionsource/notify`
- 请求类型：[MiniResourcePromotionSourceNotifyRequest](resource_request.go#L5)
- 响应类型：[MiniResourcePromotionSourceNotifyResponse](resource_response.go#L5)

淘联盟定制回调接口，用户流量插件账号打通场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `author_id` | `AuthorId` | `*string` | JSON body | 未标注 | 媒体唤起时传入的支付宝id |
| `params` | `Params` | `*string` | JSON body | 未标注 | 媒体唤起时提供的额外参数值列表，这里可能有多个值，打平以后拼入。即url_decode(authcbparams) |
| `promotion_id` | `PromotionId` | `*string` | JSON body | 未标注 | 推广位id |
| `promotion_name` | `PromotionName` | `*string` | JSON body | 未标注 | 推广位名称 |
| `source` | `Source` | `*string` | JSON body | 未标注 | 媒体来源，标识调用方 |

<a id="api-miniresourcerecordnotify"></a>

### MiniResourceRecordNotify — 提供给淘联盟备案信息接口

- 官方接口：[`alipay.open.mini.resource.record.notify`](https://opendocs.alipay.com/pre-apis/05xqyn)
- 路由：`POST /v3/alipay/open/mini/resource/record/notify`
- 请求类型：[MiniResourceRecordNotifyRequest](resource_request.go#L19)
- 响应类型：[MiniResourceRecordNotifyResponse](resource_response.go#L9)

淘联盟定制回调接口，用户流量插件账号打通场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `author_id` | `AuthorId` | `*string` | JSON body | 未标注 | 媒体唤起时传入的支付宝id |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 支付宝appid，代表一个小程序的唯一标识 |
| `params` | `Params` | `*string` | JSON body | 未标注 | 媒体唤起时提供的额外参数值列表，这里可能有多个值，打平以后拼入。即url_decode(authcbparams) |
| `site_id` | `SiteId` | `*string` | JSON body | 未标注 | 回调的siteid 淘宝账号唯一 |
| `source` | `Source` | `*string` | JSON body | 未标注 | 媒体来源，标识调用方 |
| `taobao_id` | `TaobaoId` | `*string` | JSON body | 未标注 | 淘宝id |
| `taobao_nick` | `TaobaoNick` | `*string` | JSON body | 未标注 | 淘宝昵称 |
