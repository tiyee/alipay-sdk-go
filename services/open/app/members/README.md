# members 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/app/members`。本包提供 3 个支付宝 V3 REST 操作。

通过 `members.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AppMembersCreate`](#api-appmemberscreate) | `alipay.open.app.members.create` |
| [`AppMembersDelete`](#api-appmembersdelete) | `alipay.open.app.members.delete` |
| [`AppMembersQuery`](#api-appmembersquery) | `alipay.open.app.members.query` |

## 使用示例

- [AppMembersCreate](api_members_appmemberscreate_example_test.go)
- [AppMembersDelete](api_members_appmembersdelete_example_test.go)
- [AppMembersQuery](api_members_appmembersquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-appmemberscreate"></a>

### AppMembersCreate — 应用添加成员

- 官方接口：[`alipay.open.app.members.create`](https://opendocs.alipay.com/open-v3/8b6e0190_alipay.open.app.members.create)
- 路由：`POST /v3/alipay/open/app/members/create`
- 请求类型：[AppMembersCreateRequest](members_request.go#L4)
- 响应类型：[AppMembersCreateResponse](members_response.go#L4)

应用添加成员，目前只支持小程序类型的应用使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `logon_id` | `LogonId` | `*string` | JSON body | 未标注 | 支付宝登录账号。 |
| `role` | `Role` | `*string` | JSON body | 未标注 | 为成员添加的角色类型 |

<a id="api-appmembersdelete"></a>

### AppMembersDelete — 应用删除成员

- 官方接口：[`alipay.open.app.members.delete`](https://opendocs.alipay.com/open-v3/26ef7be4_alipay.open.app.members.delete)
- 路由：`DELETE /v3/alipay/open/app/members/delete`
- 请求类型：[AppMembersDeleteRequest](members_request.go#L12)
- 响应类型：[AppMembersDeleteResponse](members_response.go#L8)

删除应用成员，目前只支持小程序类型的应用

请求约束：AppMembersDeleteRequest is the request type for alipay.open.app.members.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 被删除成员的支付宝账户唯一标识，以2088开头。 |
| `open_id` | `OpenId` | `*string` | query | 否 | 被删除成员的openId。 |
| `role` | `Role` | `*string` | query | 否 | 被删除成员的角色类型 |

<a id="api-appmembersquery"></a>

### AppMembersQuery — 应用查询成员列表

- 官方接口：[`alipay.open.app.members.query`](https://opendocs.alipay.com/open-v3/6a8b8874_alipay.open.app.members.query)
- 路由：`GET /v3/alipay/open/app/members/query`
- 请求类型：[AppMembersQueryRequest](members_request.go#L22)
- 响应类型：[AppMembersQueryResponse](members_response.go#L12)

应用查询成员列表，目前只支持小程序类型的应用

请求约束：AppMembersQueryRequest is the request type for alipay.open.app.members.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `role` | `Role` | `*string` | query | 否 | 查询的成员角色类型 |
