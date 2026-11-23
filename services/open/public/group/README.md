# open/public/group 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/group`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `group.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicGroupBatchQuery`](#api-publicgroupbatchquery) | [`alipay.open.public.group.batchquery`](https://opendocs.alipay.com/apis/05wydu) | 查询用户分组列表 | `POST` |
| [`PublicGroupCreate`](#api-publicgroupcreate) | [`alipay.open.public.group.create`](https://opendocs.alipay.com/apis/05wukz) | 用户分组创建接口 | `POST` |
| [`PublicGroupCrowdQuery`](#api-publicgroupcrowdquery) | [`alipay.open.public.group.crowd.query`](https://opendocs.alipay.com/apis/05wyfg) | 人群数量查询 | `POST` |
| [`PublicGroupDelete`](#api-publicgroupdelete) | [`alipay.open.public.group.delete`](https://opendocs.alipay.com/apis/05wuht) | 用户分组删除接口 | `DELETE` |
| [`PublicGroupModify`](#api-publicgroupmodify) | [`alipay.open.public.group.modify`](https://opendocs.alipay.com/apis/05wcqq) | 用户分组修改接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicgroupbatchquery"></a>

### PublicGroupBatchQuery — 查询用户分组列表

- 官方接口：[`alipay.open.public.group.batchquery`](https://opendocs.alipay.com/apis/05wydu)
- 路由：`POST /v3/alipay/open/public/group/batchquery`
- 请求类型：[PublicGroupBatchQueryRequest](group_request.go#L29)
- 响应类型：[PublicGroupBatchQueryResponse](group_response.go#L5)

使用该接口查询分组列表，并配合其它分组管理接口使用，最多显示200条

请求约束：PublicGroupBatchQueryRequest is the request type for alipay.open.public.group.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publicgroupcreate"></a>

### PublicGroupCreate — 用户分组创建接口

- 官方接口：[`alipay.open.public.group.create`](https://opendocs.alipay.com/apis/05wukz)
- 路由：`POST /v3/alipay/open/public/group/create`
- 请求类型：[PublicGroupCreateRequest](group_request.go#L5)
- 响应类型：[PublicGroupCreateResponse](group_response.go#L11)

开发者可使用此接口根据指定条件从关注用户中圈定人群生成分组，分组可用于组发消息等场景，分组圈定的人群必须大于50人。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_rule` | `LabelRule` | [`*[]ComplexLabelRule`](#model-complexlabelrule) | JSON body | 未标注 | 标签规则，满足该规则的粉丝将被圈定，标签id不能重复 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 分组名称，仅支持中文、字母、数字、下划线的组合。 |

<a id="api-publicgroupcrowdquery"></a>

### PublicGroupCrowdQuery — 人群数量查询

- 官方接口：[`alipay.open.public.group.crowd.query`](https://opendocs.alipay.com/apis/05wyfg)
- 路由：`POST /v3/alipay/open/public/group/crowd/query`
- 请求类型：[PublicGroupCrowdQueryRequest](group_request.go#L13)
- 响应类型：[PublicGroupCrowdQueryResponse](group_response.go#L17)

用户可以指定自己创建的分组，调用接口查询对应的人群数量，人群必须大于50人才能正常返回，否则报错。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_rule` | `LabelRule` | [`*[]ComplexLabelRule`](#model-complexlabelrule) | JSON body | 未标注 | 用户分组的规则项列表，规则项之间元素是与的逻辑，每个规则项内部用多个值表示或的逻辑 |

<a id="api-publicgroupdelete"></a>

### PublicGroupDelete — 用户分组删除接口

- 官方接口：[`alipay.open.public.group.delete`](https://opendocs.alipay.com/apis/05wuht)
- 路由：`DELETE /v3/alipay/open/public/group/delete`
- 请求类型：[PublicGroupDeleteRequest](group_request.go#L32)
- 响应类型：[PublicGroupDeleteResponse](group_response.go#L23)

调用接口删除用户分组

请求约束：PublicGroupDeleteRequest is the request type for alipay.open.public.group.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `group_id` | `GroupId` | `*string` | query | 否 | 分组ID，通过[alipay.open.public.group.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.group.create)（用户分组创建接口) 创建用户分组获取。 |

<a id="api-publicgroupmodify"></a>

### PublicGroupModify — 用户分组修改接口

- 官方接口：[`alipay.open.public.group.modify`](https://opendocs.alipay.com/apis/05wcqq)
- 路由：`POST /v3/alipay/open/public/group/modify`
- 请求类型：[PublicGroupModifyRequest](group_request.go#L19)
- 响应类型：[PublicGroupModifyResponse](group_response.go#L27)

调用改接口修改用户分组，仅支持修改分组名称。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `group_id` | `GroupId` | `*string` | JSON body | 未标注 | 分组ID，通过[alipay.open.public.group.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.group.create)（用户分组创建接口) 创建用户分组获取。 |
| `label_rule` | `LabelRule` | [`*[]ComplexLabelRule`](#model-complexlabelrule) | JSON body | 未标注 | 标签规则，满足该规则的粉丝将被圈定，标签id不能重复 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 分组名称，仅支持中文、字母、数字、下划线的组合。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-complexlabelrule"></a>

### ComplexLabelRule

类型定义：[ComplexLabelRule](group_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | 未标注 | 标签id |
| `label_value` | `LabelValue` | `*string` | 未标注 | 标签取值，当有多个取值时用英文","分隔（比如使用in操作符时）；不允许传入下划线"_"、竖线"&#124;"或者空格" " |
| `operator` | `Operator` | `*string` | 未标注 | 目前支持EQ（等于）、NEQ（不等于）、LT（小于），GT（大于）、LTEQ（小于等于）、GTEQ（大于等于）、LIKE（匹配）、BETWEEN（范围）、IN（包含）、NOTIN（不包含）操作 |

## 使用示例

- [PublicGroupBatchQuery](api_group_publicgroupbatchquery_example_test.go)
- [PublicGroupCreate](api_group_publicgroupcreate_example_test.go)
- [PublicGroupCrowdQuery](api_group_publicgroupcrowdquery_example_test.go)
- [PublicGroupDelete](api_group_publicgroupdelete_example_test.go)
- [PublicGroupModify](api_group_publicgroupmodify_example_test.go)
