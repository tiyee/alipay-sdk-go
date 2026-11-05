# open/mini/qrcode 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/qrcode`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `qrcode.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_qrcode_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniQRCodeBind`](#api-miniqrcodebind) | [`alipay.open.mini.qrcode.bind`](https://opendocs.alipay.com/open-v3/01ebf6a3_alipay.open.mini.qrcode.bind) | 关联普通二维码 | `POST` |
| [`MiniQRCodeQuery`](#api-miniqrcodequery) | [`alipay.open.mini.qrcode.query`](https://opendocs.alipay.com/open-v3/c89a2391_alipay.open.mini.qrcode.query) | 查询普通二维码 | `GET` |
| [`MiniQRCodeUnbind`](#api-miniqrcodeunbind) | [`alipay.open.mini.qrcode.unbind`](https://opendocs.alipay.com/open-v3/8905c9d2_alipay.open.mini.qrcode.unbind) | 删除已关联普通二维码 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniqrcodebind"></a>

### MiniQRCodeBind — 关联普通二维码

- 官方接口：[`alipay.open.mini.qrcode.bind`](https://opendocs.alipay.com/open-v3/01ebf6a3_alipay.open.mini.qrcode.bind)
- 路由：`POST /v3/alipay/open/mini/qrcode/bind`
- 请求类型：[MiniQRCodeBindRequest](qrcode_request.go#L5)
- 响应类型：[MiniQRCodeBindResponse](qrcode_response.go#L5)

商户可不需更换线下二维码，通过该接口完成配置后，用户用支付宝扫描普通二维码时打开小程序并跳转到指定页面，普通二维码关联小程序最多可设置100条规则。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mode` | `Mode` | `*string` | JSON body | 未标注 | 匹配规则，仅支持EXACT（精确匹配）、FUZZY（模糊匹配）两个值。 精确匹配：根据填写的二维码地址精确匹配，地址完全一致时才能唤起小程序（如：配置二维码地址为https://www.alipay.com/my?id=123，当用户扫这个地址的二维码可唤起小程序）。 模糊匹配：根据填写的二维码地址模糊匹配，只要地址前缀匹配即可唤起小程序（如：配置二维码地址为https://www.alipay.com/my/，当用户扫的二维码地址为https://www.alipay.com/my/id=123,可唤起小程序）。 |
| `page_redirection` | `PageRedirection` | `*string` | JSON body | 未标注 | 小程序功能页，配置扫描二维码后打开的小程序功能页面路径 |
| `route_pattern` | `RoutePattern` | `*string` | JSON body | 未标注 | 模式，用于描述通用的路由url格式，适用于同一个host下面存在大量路由的场景，主要用于提升路由性能，其中{0}、{1} 代表路径中的变量 |
| `route_url` | `RouteUrl` | `*string` | JSON body | 未标注 | 二维码域名，须通过ICP备案验证，支持http、https、ftp开头的链接 |

<a id="api-miniqrcodequery"></a>

### MiniQRCodeQuery — 查询普通二维码

- 官方接口：[`alipay.open.mini.qrcode.query`](https://opendocs.alipay.com/open-v3/c89a2391_alipay.open.mini.qrcode.query)
- 路由：`GET /v3/alipay/open/mini/qrcode/query`
- 请求类型：[MiniQRCodeQueryRequest](qrcode_request.go#L25)
- 响应类型：[MiniQRCodeQueryResponse](qrcode_response.go#L11)

通过appId查询商户关联的普通二维码路由信息

请求约束：MiniQRCodeQueryRequest is the request type for alipay.open.mini.qrcode.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_no` | `PageNo` | `*string` | query | 否 | 页码 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页条数，最大为10 |

<a id="api-miniqrcodeunbind"></a>

### MiniQRCodeUnbind — 删除已关联普通二维码

- 官方接口：[`alipay.open.mini.qrcode.unbind`](https://opendocs.alipay.com/open-v3/8905c9d2_alipay.open.mini.qrcode.unbind)
- 路由：`POST /v3/alipay/open/mini/qrcode/unbind`
- 请求类型：[MiniQRCodeUnbindRequest](qrcode_request.go#L17)
- 响应类型：[MiniQRCodeUnbindResponse](qrcode_response.go#L19)

删除已关联普通二维码

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `route_group` | `RouteGroup` | `*string` | JSON body | 未标注 | 路由规则组，用于唯一标记一条路由规则。调用 [关联普通二维码](https://opendocs.alipay.com/apis/00rkye) 接口关联普通二维码后的返回值。 |
| `route_url` | `RouteUrl` | `*string` | JSON body | 未标注 | 路由规则URL，调用[关联普通二维码](https://opendocs.alipay.com/apis/00rkye)接口传入的参数：route_url。解除模式与二维码的绑定关系需传入 |
