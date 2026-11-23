# open/public/label 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/label`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `label.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_label_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicLabelCreate`](#api-publiclabelcreate) | [`alipay.open.public.label.create`](https://opendocs.alipay.com/pre-apis/05wupo) | 公众号标签管理-添加标签 | `POST` |
| [`PublicLabelQuery`](#api-publiclabelquery) | [`alipay.open.public.label.query`](https://opendocs.alipay.com/pre-apis/05xq3d) | 公众号标签管理-查询标签 | `GET` |
| [`PublicLabelUserQuery`](#api-publiclabeluserquery) | [`alipay.open.public.label.user.query`](https://opendocs.alipay.com/pre-apis/05wum4) | 公众号标签管理-查询用户标签 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publiclabelcreate"></a>

### PublicLabelCreate — 公众号标签管理-添加标签

- 官方接口：[`alipay.open.public.label.create`](https://opendocs.alipay.com/pre-apis/05wupo)
- 路由：`POST /v3/alipay/open/public/label/create`
- 请求类型：[PublicLabelCreateRequest](label_request.go#L5)
- 响应类型：[PublicLabelCreateResponse](label_response.go#L5)

开发者可以使用该接口创建一个标签，一个服务窗最多同时拥有500个标签，标签名不能重复。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `name` | `Name` | `*string` | JSON body | 未标注 | 标签名 |

<a id="api-publiclabelquery"></a>

### PublicLabelQuery — 公众号标签管理-查询标签

- 官方接口：[`alipay.open.public.label.query`](https://opendocs.alipay.com/pre-apis/05xq3d)
- 路由：`GET /v3/alipay/open/public/label/query`
- 请求类型：[PublicLabelQueryRequest](label_request.go#L11)
- 响应类型：[PublicLabelQueryResponse](label_response.go#L13)

开发者可以查询该服务窗所有标签值及拥有该标签的用户个数

请求约束：PublicLabelQueryRequest is the request type for alipay.open.public.label.query.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publiclabeluserquery"></a>

### PublicLabelUserQuery — 公众号标签管理-查询用户标签

- 官方接口：[`alipay.open.public.label.user.query`](https://opendocs.alipay.com/pre-apis/05wum4)
- 路由：`GET /v3/alipay/open/public/label/user/query`
- 请求类型：[PublicLabelUserQueryRequest](label_request.go#L14)
- 响应类型：[PublicLabelUserQueryResponse](label_response.go#L19)

开发者可以使用该接口查询某个用户的所有标签

请求约束：PublicLabelUserQueryRequest is the request type for alipay.open.public.label.user.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 支付宝用户的userid，2088开头长度为16位的字符串 |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝用户的唯一标识 |
