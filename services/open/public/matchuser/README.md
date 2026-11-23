# open/public/matchuser 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/matchuser`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `matchuser.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_matchuser_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicMatchUserLabelCreate`](#api-publicmatchuserlabelcreate) | [`alipay.open.public.matchuser.label.create`](https://opendocs.alipay.com/apis/05wycu) | 用户打标接口 | `POST` |
| [`PublicMatchUserLabelDelete`](#api-publicmatchuserlabeldelete) | [`alipay.open.public.matchuser.label.delete`](https://opendocs.alipay.com/apis/05wz8y) | 用户取消标签接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicmatchuserlabelcreate"></a>

### PublicMatchUserLabelCreate — 用户打标接口

- 官方接口：[`alipay.open.public.matchuser.label.create`](https://opendocs.alipay.com/apis/05wycu)
- 路由：`POST /v3/alipay/open/public/matchuser/label/create`
- 请求类型：[PublicMatchUserLabelCreateRequest](matchuser_request.go#L5)
- 响应类型：[PublicMatchUserLabelCreateResponse](matchuser_response.go#L5)

开发者可以使用此接口向匹配到的用户添加指定标签值，每个用户每个标签只能有一个标签值，当向同一个用户添加多个标签值时，最新的标签值会覆盖之前的标签值

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | JSON body | 未标注 | 标签 id，只支持生活号自定义标签。通过 [alipay.open.public.life.label.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create)(创建标签接口)创建自定义标签后获取。 |
| `label_value` | `LabelValue` | `*string` | JSON body | 未标注 | 标签值，由开发者自主指定，标签值类型要满足 [alipay.open.public.life.label.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create)(创建标签接口)中data_type参数的限定。 |
| `matchers` | `Matchers` | [`*[]Matcher`](#model-matcher) | JSON body | 未标注 | 支付宝用户匹配器列表，最多传入10条 |

<a id="api-publicmatchuserlabeldelete"></a>

### PublicMatchUserLabelDelete — 用户取消标签接口

- 官方接口：[`alipay.open.public.matchuser.label.delete`](https://opendocs.alipay.com/apis/05wz8y)
- 路由：`POST /v3/alipay/open/public/matchuser/label/delete`
- 请求类型：[PublicMatchUserLabelDeleteRequest](matchuser_request.go#L15)
- 响应类型：[PublicMatchUserLabelDeleteResponse](matchuser_response.go#L13)

开发者可以使用此接口删除匹配到的支付宝用户的指定标签值

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | JSON body | 未标注 | 标签 id，只支持生活号自定义标签。通过 [alipay.open.public.life.label.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create)(创建标签接口)创建自定义标签后获取。 |
| `matchers` | `Matchers` | [`*[]Matcher`](#model-matcher) | JSON body | 未标注 | 支付宝用户匹配器列表，最多传入10条 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-matcher"></a>

### Matcher

类型定义：[Matcher](matchuser_model.go#L12)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `identity_card` | `IdentityCard` | `*string` | 未标注 | 身份证号码，与user_id、mobile_no不能同时为空 |
| `mobile_no` | `MobileNo` | `*string` | 未标注 | 手机号码，与user_id、identity_card不能同时为空 |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户openid |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户id，2088开头16位长度的字符串，与mobile_no、identity_card不能同时为空 |
