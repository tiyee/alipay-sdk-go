# iservice/ccm/role 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/role`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `role.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CcmRoleCreate`](#api-ccmrolecreate) | [`alipay.iservice.ccm.role.create`](https://opendocs.alipay.com/open-v3/632638c3_alipay.iservice.ccm.role.create) | 创建角色 | `POST` |
| [`CcmRoleDelete`](#api-ccmroledelete) | [`alipay.iservice.ccm.role.delete`](https://opendocs.alipay.com/open-v3/8a5f6702_alipay.iservice.ccm.role.delete) | 删除角色 | `DELETE` |
| [`CcmRoleGet`](#api-ccmroleget) | [`alipay.iservice.ccm.role.get`](https://opendocs.alipay.com/apis/05wvxp) | 获取角色详情 | `GET` |
| [`CcmRoleModify`](#api-ccmrolemodify) | [`alipay.iservice.ccm.role.modify`](https://opendocs.alipay.com/open-v3/8e38792c_alipay.iservice.ccm.role.modify) | 更新角色 | `POST` |
| [`CcmRolePageQuery`](#api-ccmrolepagequery) | [`alipay.iservice.ccm.role.page.query`](https://opendocs.alipay.com/open-v3/0c530d84_alipay.iservice.ccm.role.page.query) | 分页查询角色 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmrolecreate"></a>

### CcmRoleCreate — 创建角色

- 官方接口：[`alipay.iservice.ccm.role.create`](https://opendocs.alipay.com/open-v3/632638c3_alipay.iservice.ccm.role.create)
- 路由：`POST /v3/alipay/iservice/ccm/role/create`
- 请求类型：[CcmRoleCreateRequest](role_request.go#L5)
- 响应类型：[CcmRoleCreateResponse](role_response.go#L9)

创建角色

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 部门id（即租户实例ID、数据权限ID） |
| `creator_id` | `CreatorId` | `*string` | JSON body | 未标注 | 创建人id |
| `description` | `Description` | `*string` | JSON body | 未标注 | 角色描述信息 |
| `function_ids` | `FunctionIds` | `*[]string` | JSON body | 未标注 | 角色关联额功能点id |
| `name` | `Name` | `*string` | JSON body | 未标注 | 角色名 |

<a id="api-ccmroledelete"></a>

### CcmRoleDelete — 删除角色

- 官方接口：[`alipay.iservice.ccm.role.delete`](https://opendocs.alipay.com/open-v3/8a5f6702_alipay.iservice.ccm.role.delete)
- 路由：`DELETE /v3/alipay/iservice/ccm/role/delete`
- 请求类型：[CcmRoleDeleteRequest](role_request.go#L33)
- 响应类型：[CcmRoleDeleteResponse](role_response.go#L15)

删除角色

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `id` | `Id` | `*string` | query | 否 | 角色id |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 部门id（即租户实例ID、数据权限ID），如果不传入，使用租户id |

<a id="api-ccmroleget"></a>

### CcmRoleGet — 获取角色详情

- 官方接口：[`alipay.iservice.ccm.role.get`](https://opendocs.alipay.com/apis/05wvxp)
- 路由：`GET /v3/alipay/iservice/ccm/role/get`
- 请求类型：[CcmRoleGetRequest](role_request.go#L41)
- 响应类型：[CcmRoleGetResponse](role_response.go#L18)

获取角色详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `id` | `Id` | `*string` | query | 否 | 角色id |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 部门id，不传默认用租户ID |

<a id="api-ccmrolemodify"></a>

### CcmRoleModify — 更新角色

- 官方接口：[`alipay.iservice.ccm.role.modify`](https://opendocs.alipay.com/open-v3/8e38792c_alipay.iservice.ccm.role.modify)
- 路由：`POST /v3/alipay/iservice/ccm/role/modify`
- 请求类型：[CcmRoleModifyRequest](role_request.go#L19)
- 响应类型：[CcmRoleModifyResponse](role_response.go#L40)

更新角色

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 部门id（即租户实例ID、数据权限ID），不传默认使用租户id |
| `description` | `Description` | `*string` | JSON body | 未标注 | 描述信息 |
| `function_ids` | `FunctionIds` | `*[]string` | JSON body | 未标注 | 角色关联额功能点id |
| `id` | `Id` | `*string` | JSON body | 未标注 | 角色id |
| `updater_id` | `UpdaterId` | `*string` | JSON body | 未标注 | 最后修改人id |

<a id="api-ccmrolepagequery"></a>

### CcmRolePageQuery — 分页查询角色

- 官方接口：[`alipay.iservice.ccm.role.page.query`](https://opendocs.alipay.com/open-v3/0c530d84_alipay.iservice.ccm.role.page.query)
- 路由：`GET /v3/alipay/iservice/ccm/role/page/query`
- 请求类型：[CcmRolePageQueryRequest](role_request.go#L49)
- 响应类型：[CcmRolePageQueryResponse](role_response.go#L43)

分页查询角色

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `name` | `Name` | `*string` | query | 否 | 角色名称 |
| `page_num` | `PageNum` | `*string` | query | 否 | 查询结果的页码，起始值为 1，默认值为 1 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页查询时设置的每页记录数，最大值 100 行，默认为 10 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 部门id（即租户实例ID、数据权限ID） |

## 使用示例

- [CcmRoleGet](api_role_ccmroleget_example_test.go)
- [CcmRoleCreate](api_role_ccmrolecreate_example_test.go)
- [CcmRoleDelete](api_role_ccmroledelete_example_test.go)
- [CcmRoleModify](api_role_ccmrolemodify_example_test.go)
- [CcmRolePageQuery](api_role_ccmrolepagequery_example_test.go)
