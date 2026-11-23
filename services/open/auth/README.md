# open/auth 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/auth`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `auth.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_auth_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`AuthAppAESGet`](#api-authappaesget) | [`alipay.open.auth.app.aes.get`](https://opendocs.alipay.com/open-v3/4d82ebbf_alipay.open.auth.app.aes.get) | 授权应用aes密钥查询 | `GET` |
| [`AuthAppAESSet`](#api-authappaesset) | [`alipay.open.auth.app.aes.set`](https://opendocs.alipay.com/open-v3/306fd1a6_alipay.open.auth.app.aes.set) | 授权应用aes密钥设置 | `POST` |
| [`AuthTokenAppQuery`](#api-authtokenappquery) | [`alipay.open.auth.token.app.query`](https://opendocs.alipay.com/open-v3/b6e7c6b3_alipay.open.auth.token.app.query) | 查询某个应用授权AppAuthToken的授权信息 | `GET` |
| [`AuthUserAuthRelationshipQuery`](#api-authuserauthrelationshipquery) | [`alipay.open.auth.userauth.relationship.query`](https://opendocs.alipay.com/open-v3/5ea1017e_alipay.open.auth.userauth.relationship.query) | 用户授权关系查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-authappaesget"></a>

### AuthAppAESGet — 授权应用aes密钥查询

- 官方接口：[`alipay.open.auth.app.aes.get`](https://opendocs.alipay.com/open-v3/4d82ebbf_alipay.open.auth.app.aes.get)
- 路由：`GET /v3/alipay/open/auth/app/aes/get`
- 请求类型：[AuthAppAESGetRequest](auth_request.go#L11)
- 响应类型：[AuthAppAESGetResponse](auth_response.go#L5)

商家应用授权给ISV服务商代管理应用场景下，ISV开发者使用三方应用appId发起调用，通过该接口获取商户应用的AES密钥。接口返回的商户应用aes密钥内容为加密内容，需要ISV使用三方应用的AES密钥对接口返回进行解密。java sdk内部会自动完成解密返回明文内容。[SDK AES加解密支持](https://docs.open.alipay.com/common/104567)

请求约束：AuthAppAESGetRequest is the request type for alipay.open.auth.app.aes.get.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_app_id` | `MerchantAppId` | `*string` | query | 否 | 商家应用appId |

<a id="api-authappaesset"></a>

### AuthAppAESSet — 授权应用aes密钥设置

- 官方接口：[`alipay.open.auth.app.aes.set`](https://opendocs.alipay.com/open-v3/306fd1a6_alipay.open.auth.app.aes.set)
- 路由：`POST /v3/alipay/open/auth/app/aes/set`
- 请求类型：[AuthAppAESSetRequest](auth_request.go#L5)
- 响应类型：[AuthAppAESSetResponse](auth_response.go#L11)

商家应用授权给ISV服务商代管理应用场景下，ISV开发者使用三方应用appId发起调用，通过该接口设置商户应用的AES密钥。接口返回的商户应用aes密钥内容为加密内容，需要ISV使用三方应用的AES密钥对接口返回进行解密。java sdk内部会自动完成解密返回明文内容。[SDK AES加解密支持。](https://docs.open.alipay.com/common/104567)

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_app_id` | `MerchantAppId` | `*string` | JSON body | 未标注 | 商家应用appId |

<a id="api-authtokenappquery"></a>

### AuthTokenAppQuery — 查询某个应用授权AppAuthToken的授权信息

- 官方接口：[`alipay.open.auth.token.app.query`](https://opendocs.alipay.com/open-v3/b6e7c6b3_alipay.open.auth.token.app.query)
- 路由：`GET /v3/alipay/open/auth/token/app/query`
- 请求类型：[AuthTokenAppQueryRequest](auth_request.go#L17)
- 响应类型：[AuthTokenAppQueryResponse](auth_response.go#L17)

当商户把服务窗、店铺等接口的权限授权给ISV之后，支付宝会给ISV颁发一个app_auth_token。如若授权成功之后，ISV想知道用户的授权信息，如授权者、授权接口列表等信息，可以调用本接口查询某个app_auth_token对应的授权信息

请求约束：AuthTokenAppQueryRequest is the request type for alipay.open.auth.token.app.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_auth_token` | `AppAuthToken` | `*string` | query | 否 | 应用授权令牌 |

<a id="api-authuserauthrelationshipquery"></a>

### AuthUserAuthRelationshipQuery — 用户授权关系查询

- 官方接口：[`alipay.open.auth.userauth.relationship.query`](https://opendocs.alipay.com/open-v3/5ea1017e_alipay.open.auth.userauth.relationship.query)
- 路由：`GET /v3/alipay/open/auth/userauth/relationship/query`
- 请求类型：[AuthUserAuthRelationshipQueryRequest](auth_request.go#L23)
- 响应类型：[AuthUserAuthRelationshipQueryResponse](auth_response.go#L37)

校验本应用和用户的授权关系有效性

请求约束：AuthUserAuthRelationshipQueryRequest is the request type for alipay.open.auth.userauth.relationship.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 蚂蚁统一会员ID，如果未开启openid改造，则此值必传 |
| `open_id` | `OpenId` | `*string` | query | 否 | 蚂蚁统一会员ID，蚂蚁统一会员ID，如果已完成openid改造，则此值必传 |
| `scopes` | `Scopes` | `*string` | query | 否 | 查询授权关系所关联的scope，半角逗号分隔的列表，且不允许有重复值 |
