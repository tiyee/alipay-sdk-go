# system 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/system`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../README.md)创建 `*core.Client`，再通过 `system.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_oauth_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`OAuthToken`](#api-oauthtoken) | [`alipay.system.oauth.token`](https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token) | 换取授权访问令牌 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

授权码、访问令牌与刷新令牌均为敏感凭证；由调用方安全保存和轮换，避免写入日志。SDK 不缓存令牌。

## 接口说明与请求参数

<a id="api-oauthtoken"></a>

### OAuthToken — 换取授权访问令牌

- 官方接口：[`alipay.system.oauth.token`](https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token)
- 路由：`POST /v3/alipay/system/oauth/token`
- 请求类型：[OAuthTokenRequest](system_request.go#L10)
- 响应类型：[OAuthTokenResponse](system_response.go#L5)

换取授权访问令牌

请求约束：OAuthTokenRequest 是 alipay.system.oauth.token（换取授权访问令牌）
的请求。grant_type=authorization_code 时 code 必填；
grant_type=refresh_token 时 refresh_token 必填。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `grant_type` | `GrantType` | [`GrantType`](system_model.go#L4) | JSON body | 未标注 | 授权方式 |
| `code` | `Code` | `*string` | JSON body | 未标注 | 授权码，用户对应用授权后得到。本参数在 grant_type 为 authorization_code 时必填；为 refresh_token 时不填。<br>SDK 字段说明：授权码，用户对应用授权后得到；<br>grant_type 为 authorization_code 时必填。 |
| `refresh_token` | `RefreshToken` | `*string` | JSON body | 未标注 | 刷新令牌，上次换取访问令牌时得到。本参数在 grant_type 为 authorization_code 时不填；为 refresh_token 时必填，且该值来源于此接口的返回值 app_refresh_token（即至少需要通过 grant_type=authorization_code 调用此接口一次才能获取）。<br>SDK 字段说明：刷新令牌，上次换取访问令牌时得到；<br>grant_type 为 refresh_token 时必填。 |

注意事项：

- 访问令牌与刷新令牌按调用方令牌管理策略保存与轮换，SDK 不缓存。
