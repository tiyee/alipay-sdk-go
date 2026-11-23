# open/mini/inner 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/inner`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `inner.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_inner_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniInnerExperienceCancel`](#api-miniinnerexperiencecancel) | [`alipay.open.mini.inner.experience.cancel`](https://opendocs.alipay.com/pre-apis/05xe27) | 内部链路取消体验版 | `POST` |
| [`MiniInnerMembersAdd`](#api-miniinnermembersadd) | [`alipay.open.mini.inner.members.add`](https://opendocs.alipay.com/pre-apis/05xhp2) | 内部链路添加开发成员 | `POST` |
| [`MiniInnerMembersDelete`](#api-miniinnermembersdelete) | [`alipay.open.mini.inner.members.delete`](https://opendocs.alipay.com/pre-apis/05xqzc) | 内部链路删除开发成员 | `DELETE` |
| [`MiniInnerMembersQuery`](#api-miniinnermembersquery) | [`alipay.open.mini.inner.members.query`](https://opendocs.alipay.com/pre-apis/05xaad) | 内部链路查询开发成员 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerexperiencecancel"></a>

### MiniInnerExperienceCancel — 内部链路取消体验版

- 官方接口：[`alipay.open.mini.inner.experience.cancel`](https://opendocs.alipay.com/pre-apis/05xe27)
- 路由：`POST /v3/alipay/open/mini/inner/experience/cancel`
- 请求类型：[MiniInnerExperienceCancelRequest](inner_request.go#L5)
- 响应类型：[MiniInnerExperienceCancelResponse](inner_response.go#L5)

提供给内部一二方用于取消体验版

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本号 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端id |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 业务小程序appid |
| `oid` | `Oid` | `*string` | JSON body | 未标注 | oid |
| `operator_id` | `OperatorId` | `*string` | JSON body | 未标注 | 操作员id |

<a id="api-miniinnermembersadd"></a>

### MiniInnerMembersAdd — 内部链路添加开发成员

- 官方接口：[`alipay.open.mini.inner.members.add`](https://opendocs.alipay.com/pre-apis/05xhp2)
- 路由：`POST /v3/alipay/open/mini/inner/members/add`
- 请求类型：[MiniInnerMembersAddRequest](inner_request.go#L21)
- 响应类型：[MiniInnerMembersAddResponse](inner_response.go#L9)

添加开发成员

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `domain_account` | `DomainAccount` | `*string` | JSON body | 未标注 | 域账号（不为空时会添加联系人） |
| `login_id` | `LoginId` | `*string` | JSON body | 未标注 | 支付宝登陆账号 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 业务appId |
| `type` | `Type` | `*string` | JSON body | 未标注 | 成员类型 |

<a id="api-miniinnermembersdelete"></a>

### MiniInnerMembersDelete — 内部链路删除开发成员

- 官方接口：[`alipay.open.mini.inner.members.delete`](https://opendocs.alipay.com/pre-apis/05xqzc)
- 路由：`DELETE /v3/alipay/open/mini/inner/members/delete`
- 请求类型：[MiniInnerMembersDeleteRequest](inner_request.go#L35)
- 响应类型：[MiniInnerMembersDeleteResponse](inner_response.go#L13)

删除开发成员

请求约束：MiniInnerMembersDeleteRequest is the request type for alipay.open.mini.inner.members.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 业务小程序id |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务来源 |
| `domain_account` | `DomainAccount` | `*string` | query | 否 | 域账号（如传入则删除联系人） |
| `login_id` | `LoginId` | `*string` | query | 否 | 登陆支付宝账号 |
| `type` | `Type` | `*string` | query | 否 | 成员类型 |

<a id="api-miniinnermembersquery"></a>

### MiniInnerMembersQuery — 内部链路查询开发成员

- 官方接口：[`alipay.open.mini.inner.members.query`](https://opendocs.alipay.com/pre-apis/05xaad)
- 路由：`GET /v3/alipay/open/mini/inner/members/query`
- 请求类型：[MiniInnerMembersQueryRequest](inner_request.go#L49)
- 响应类型：[MiniInnerMembersQueryResponse](inner_response.go#L17)

内部链路查询开发成员

请求约束：MiniInnerMembersQueryRequest is the request type for alipay.open.mini.inner.members.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 业务小程序id |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务来源 |
