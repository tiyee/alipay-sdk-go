# iservice/ccm/agent 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/agent`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `agent.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CcmAgentCreate`](#api-ccmagentcreate) | [`alipay.iservice.ccm.agent.create`](https://opendocs.alipay.com/open-v3/d09eb15f_alipay.iservice.ccm.agent.create) | 新增客服接口 | `POST` |
| [`CcmAgentDelete`](#api-ccmagentdelete) | [`alipay.iservice.ccm.agent.delete`](https://opendocs.alipay.com/open-v3/dc156fee_alipay.iservice.ccm.agent.delete) | 删除客服接口 | `DELETE` |
| [`CcmAgentGet`](#api-ccmagentget) | [`alipay.iservice.ccm.agent.get`](https://opendocs.alipay.com/open-v3/4a14056c_alipay.iservice.ccm.agent.get) | 查询客服详情 | `GET` |
| [`CcmAgentModify`](#api-ccmagentmodify) | [`alipay.iservice.ccm.agent.modify`](https://opendocs.alipay.com/open-v3/e5964f19_alipay.iservice.ccm.agent.modify) | 更新客服接口 | `POST` |
| [`CcmAgentQuery`](#api-ccmagentquery) | [`alipay.iservice.ccm.agent.query`](https://opendocs.alipay.com/open-v3/992aed93_alipay.iservice.ccm.agent.query) | 客服查询接口 | `POST` |
| [`CcmAgentScheduleLogQuery`](#api-ccmagentschedulelogquery) | [`alipay.iservice.ccm.agent.schedulelog.query`](https://opendocs.alipay.com/open-v3/be76c956_alipay.iservice.ccm.agent.schedulelog.query) | 查询客服状态变更流水日志 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmagentcreate"></a>

### CcmAgentCreate — 新增客服接口

- 官方接口：[`alipay.iservice.ccm.agent.create`](https://opendocs.alipay.com/open-v3/d09eb15f_alipay.iservice.ccm.agent.create)
- 路由：`POST /v3/alipay/iservice/ccm/agent/create`
- 请求类型：[CcmAgentCreateRequest](agent_request.go#L5)
- 响应类型：[CcmAgentCreateResponse](agent_response.go#L9)

新增一个CCM客服

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `answering_mode` | `AnsweringMode` | `*string` | JSON body | 未标注 | 热线接入方式，取值0，1，2 |
| `ccs_instance_ids` | `CcsInstanceIds` | `*[]string` | JSON body | 未标注 | 部门id列表 |
| `chat_configs` | `ChatConfigs` | [`*[]AgentChatInfo`](#model-agentchatinfo) | JSON body | 未标注 | 在线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应 |
| `creator_id` | `CreatorId` | `*string` | JSON body | 未标注 | 操作人客服id |
| `email` | `Email` | `*string` | JSON body | 未标注 | 客服邮箱 |
| `external_user_id` | `ExternalUserId` | `*string` | JSON body | 未标注 | 客服账号id，当user_channel=ALIPAY时，external_user_id为支付宝账号id |
| `hotline_configs` | `HotlineConfigs` | [`*[]AgentHotlineInfo`](#model-agenthotlineinfo) | JSON body | 未标注 | 热线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应 |
| `job_number` | `JobNumber` | `*string` | JSON body | 未标注 | 客服工号，新增后不可变更。4位数字，不可重复 |
| `mobile` | `Mobile` | `*string` | JSON body | 未标注 | 客服手机号 |
| `nick_name` | `NickName` | `*string` | JSON body | 未标注 | 客服昵称 |
| `real_name` | `RealName` | `*string` | JSON body | 未标注 | 客服姓名 |
| `role_ids` | `RoleIds` | [`*[]RoleId`](#model-roleid) | JSON body | 未标注 | 角色id列表 |
| `user_channel` | `UserChannel` | `*string` | JSON body | 未标注 | 客服账号渠道 ALIPAY:支付宝账号 |

<a id="api-ccmagentdelete"></a>

### CcmAgentDelete — 删除客服接口

- 官方接口：[`alipay.iservice.ccm.agent.delete`](https://opendocs.alipay.com/open-v3/dc156fee_alipay.iservice.ccm.agent.delete)
- 路由：`DELETE /v3/alipay/iservice/ccm/agent/delete`
- 请求类型：[CcmAgentDeleteRequest](agent_request.go#L107)
- 响应类型：[CcmAgentDeleteResponse](agent_response.go#L15)

管理员删除一个客服

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `id` | `Id` | `*string` | query | 否 | 客服id |

<a id="api-ccmagentget"></a>

### CcmAgentGet — 查询客服详情

- 官方接口：[`alipay.iservice.ccm.agent.get`](https://opendocs.alipay.com/open-v3/4a14056c_alipay.iservice.ccm.agent.get)
- 路由：`GET /v3/alipay/iservice/ccm/agent/get`
- 请求类型：[CcmAgentGetRequest](agent_request.go#L113)
- 响应类型：[CcmAgentGetResponse](agent_response.go#L18)

ISV或商户查询获取某个客服的详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `id` | `Id` | `*string` | query | 否 | 客服id |
| `job_number` | `JobNumber` | `*string` | query | 否 | 客服工号 |
| `user_channel` | `UserChannel` | `*string` | query | 否 | 客服账号渠道, ALIPAY:支付宝账号 user_channel和external_user_id必须配对使用 |
| `external_user_id` | `ExternalUserId` | `*string` | query | 否 | 客服账号id，当user_channel=ALIPAY时，此id为支付宝2088id external_user_id和user_channel必须配对使用 |

<a id="api-ccmagentmodify"></a>

### CcmAgentModify — 更新客服接口

- 官方接口：[`alipay.iservice.ccm.agent.modify`](https://opendocs.alipay.com/open-v3/e5964f19_alipay.iservice.ccm.agent.modify)
- 路由：`POST /v3/alipay/iservice/ccm/agent/modify`
- 请求类型：[CcmAgentModifyRequest](agent_request.go#L35)
- 响应类型：[CcmAgentModifyResponse](agent_response.go#L66)

管理员修改客服信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `answering_mode` | `AnsweringMode` | `*string` | JSON body | 未标注 | 热线接入方式，取值0，1，2 |
| `ccs_instance_ids` | `CcsInstanceIds` | `*[]string` | JSON body | 未标注 | 部门id列表 |
| `chat_configs` | `ChatConfigs` | [`*[]AgentChatInfo`](#model-agentchatinfo) | JSON body | 未标注 | 在线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应 |
| `email` | `Email` | `*string` | JSON body | 未标注 | 客服邮箱 |
| `hotline_configs` | `HotlineConfigs` | [`*[]AgentHotlineInfo`](#model-agenthotlineinfo) | JSON body | 未标注 | 热线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应 |
| `id` | `Id` | `*string` | JSON body | 未标注 | 客服id |
| `mobile` | `Mobile` | `*string` | JSON body | 未标注 | 客服手机号 |
| `nick_name` | `NickName` | `*string` | JSON body | 未标注 | 客服昵称 |
| `real_name` | `RealName` | `*string` | JSON body | 未标注 | 客服姓名 |
| `role_ids` | `RoleIds` | [`*[]RoleId`](#model-roleid) | JSON body | 未标注 | 客服角色id列表 |
| `updater_id` | `UpdaterId` | `*string` | JSON body | 未标注 | 操作人客服id |

<a id="api-ccmagentquery"></a>

### CcmAgentQuery — 客服查询接口

- 官方接口：[`alipay.iservice.ccm.agent.query`](https://opendocs.alipay.com/open-v3/992aed93_alipay.iservice.ccm.agent.query)
- 路由：`POST /v3/alipay/iservice/ccm/agent/query`
- 请求类型：[CcmAgentQueryRequest](agent_request.go#L61)
- 响应类型：[CcmAgentQueryResponse](agent_response.go#L69)

根据指定的查询条件查询客服列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `answering_mode` | `AnsweringMode` | `*string` | JSON body | 未标注 | 热线接入方式 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 部门id |
| `chat_ext_group_ids` | `ChatExtGroupIds` | `*[]string` | JSON body | 未标注 | 在线扩展技能组id |
| `chat_group_ids` | `ChatGroupIds` | `*[]string` | JSON body | 未标注 | 在线技能组id |
| `chat_level_ids` | `ChatLevelIds` | `*[]string` | JSON body | 未标注 | 在线服务等级 |
| `email` | `Email` | `*string` | JSON body | 未标注 | 客服邮箱 |
| `hotline_group_ids` | `HotlineGroupIds` | `*[]string` | JSON body | 未标注 | 热线扩展技能组id |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 查询结果的页码，起始值为 1，默认值为 1 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 分页查询时设置的每页记录数，最大值 100 行，默认为 10 |
| `real_name` | `RealName` | `*string` | JSON body | 未标注 | 客服姓名 |
| `role_ids` | `RoleIds` | `*[]string` | JSON body | 未标注 | 客服角色id |

<a id="api-ccmagentschedulelogquery"></a>

### CcmAgentScheduleLogQuery — 查询客服状态变更流水日志

- 官方接口：[`alipay.iservice.ccm.agent.schedulelog.query`](https://opendocs.alipay.com/open-v3/be76c956_alipay.iservice.ccm.agent.schedulelog.query)
- 路由：`POST /v3/alipay/iservice/ccm/agent/schedulelog/query`
- 请求类型：[CcmAgentScheduleLogQueryRequest](agent_request.go#L87)
- 响应类型：[CcmAgentScheduleLogQueryResponse](agent_response.go#L81)

ISV或商户批量查询获取客服状态变更流水日志详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_ids` | `AgentIds` | `*[]string` | JSON body | 未标注 | 客服id集合,限制最多100个id,不传查询部门下所有客服的状态变更流水日志 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 部门id（即租户实例ID、数据权限ID） |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 客服状态变更结束时间,开始时间与结束时间间隔不能超过1天,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z' |
| `limit` | `Limit` | `*int64` | JSON body | 未标注 | 查询条数,最大100, 不传默认100 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 查询结果的页码，起始值为 1，默认值为 1 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 分页查询时设置的每页记录数，最大值 100 行，默认为 100 |
| `start_id` | `StartId` | `*int64` | JSON body | 未标注 | 起始id,分页导出的起始客服状态变更流水id,第一页传0,翻页时传上一页结果的最大id |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 客服状态变更开始时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z' |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-agentchatinfo"></a>

### AgentChatInfo

类型定义：[AgentChatInfo](agent_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | 未标注 | 数据权限id(租户实例id) |
| `extended_group_ids` | `ExtendedGroupIds` | `*[]string` | 未标注 | 在线扩展技能组id列表 |
| `group_id` | `GroupId` | `*string` | 未标注 | 在线技能组id |
| `level_id` | `LevelId` | `*string` | 未标注 | 在线服务等级 |

<a id="model-agenthotlineinfo"></a>

### AgentHotlineInfo

类型定义：[AgentHotlineInfo](agent_model.go#L17)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | 未标注 | 数据权限id(租户实例id) |
| `group_ids` | `GroupIds` | `*[]string` | 未标注 | 热线技能组id |

<a id="model-roleid"></a>

### RoleId

类型定义：[RoleId](agent_model.go#L97)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | 未标注 | 部门id（即租户实例ID、数据权限ID） |
| `role_id` | `RoleId` | `*string` | 未标注 | 角色ID |
