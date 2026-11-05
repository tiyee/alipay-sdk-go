# open/public/account 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/account`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `account.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_account_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicAccountCreate`](#api-publicaccountcreate) | [`alipay.open.public.account.create`](https://opendocs.alipay.com/apis/05wyjj) | 添加绑定商户会员号 | `POST` |
| [`PublicAccountDelete`](#api-publicaccountdelete) | [`alipay.open.public.account.delete`](https://opendocs.alipay.com/apis/05wuq3) | 解除绑定商户会员号 | `DELETE` |
| [`PublicAccountQuery`](#api-publicaccountquery) | [`alipay.open.public.account.query`](https://opendocs.alipay.com/apis/05wupq) | 查询绑定商户会员号 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicaccountcreate"></a>

### PublicAccountCreate — 添加绑定商户会员号

- 官方接口：[`alipay.open.public.account.create`](https://opendocs.alipay.com/apis/05wyjj)
- 路由：`POST /v3/alipay/open/public/account/create`
- 请求类型：[PublicAccountCreateRequest](account_request.go#L5)
- 响应类型：[PublicAccountCreateResponse](account_response.go#L5)

当用户成为商户的关注用户后，可以在商户的服务窗中点击“添加绑定商户会员号”功能，支付宝系统收到操作请求后将该动作通知给商户（使用用户发送信息到商户接口，其中eventType（事件类型）为click，actionParam（按钮标识）为authentication），商户根据此通知调用商户回复消息接口（其中须包含Url链接地址），支付宝收到商户的回复消息中的链接地址后，自动跳转至商户平台的上商户会员绑定界面中，让用户完成账户绑定。 当用户有效完成账户绑定后，商户调用本接口，把绑定结果数据通知给支付宝。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 账户添加成功，在支付宝与其对应的协议号。如果账户重复添加，接口保证幂等依然视为添加成功，返回此前该账户在支付宝对应的协议号。其他异常该字段不存在。 |
| `bind_account_no` | `BindAccountNo` | `*string` | JSON body | 未标注 | 绑定帐号，建议在开发者的系统中保持唯一性 |
| `display_name` | `DisplayName` | `*string` | JSON body | 未标注 | 开发者期望在服务窗首页看到的关于该用户的显示信息，最长10个字符 |
| `from_user_id` | `FromUserId` | `*string` | JSON body | 未标注 | 要绑定的商户会员对应的支付宝userid，2088开头长度为16位的字符串 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户的唯一标识 |
| `real_name` | `RealName` | `*string` | JSON body | 未标注 | 要绑定的商户会员的真实姓名，最长10个汉字 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 备注信息，开发者可以通过该字段纪录其他的额外信息 |

<a id="api-publicaccountdelete"></a>

### PublicAccountDelete — 解除绑定商户会员号

- 官方接口：[`alipay.open.public.account.delete`](https://opendocs.alipay.com/apis/05wuq3)
- 路由：`DELETE /v3/alipay/open/public/account/delete`
- 请求类型：[PublicAccountDeleteRequest](account_request.go#L23)
- 响应类型：[PublicAccountDeleteResponse](account_response.go#L11)

在支付宝服务窗中目前一共有三种解除绑定商户会员号的场景，具体如下：
用户取消关注服务窗
用户在服务窗内手动执行解除绑定操作
开发者调用解除绑定商户会员号接口解除绑定

请求约束：PublicAccountDeleteRequest is the request type for alipay.open.public.account.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | query | 否 | 协议号，商户会员在支付宝服务窗账号中的唯一标识，与bind_account_no不能同时为空 |
| `bind_account_no` | `BindAccountNo` | `*string` | query | 否 | 绑定帐号，建议在开发者的系统中保持唯一性，与agreement_id不能同时为空 |
| `from_user_id` | `FromUserId` | `*string` | query | 否 | 绑定用户的支付宝userid，2088开头16位长度的字符串，与agreementId不能同时为空 |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝用户的唯一标识 |

<a id="api-publicaccountquery"></a>

### PublicAccountQuery — 查询绑定商户会员号

- 官方接口：[`alipay.open.public.account.query`](https://opendocs.alipay.com/apis/05wupq)
- 路由：`GET /v3/alipay/open/public/account/query`
- 请求类型：[PublicAccountQueryRequest](account_request.go#L35)
- 响应类型：[PublicAccountQueryResponse](account_response.go#L15)

当用户成为商户的关注用户之后，则商户可以通过本接口查询关注者的绑定账户，以便补全异常情况下的单边账户数据。

请求约束：PublicAccountQueryRequest is the request type for alipay.open.public.account.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 用户的支付宝用户号，2088开头。 |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝用户的唯一标识 |
