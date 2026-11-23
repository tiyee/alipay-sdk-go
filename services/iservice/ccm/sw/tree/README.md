# tree 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/tree`。本包提供 6 个支付宝 V3 REST 操作。

通过 `tree.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CcmSwTreeBatchQuery`](#api-ccmswtreebatchquery) | `alipay.iservice.ccm.sw.tree.batchquery` |
| [`CcmSwTreeCreate`](#api-ccmswtreecreate) | `alipay.iservice.ccm.sw.tree.create` |
| [`CcmSwTreeDelete`](#api-ccmswtreedelete) | `alipay.iservice.ccm.sw.tree.delete` |
| [`CcmSwTreeGet`](#api-ccmswtreeget) | `alipay.iservice.ccm.sw.tree.get` |
| [`CcmSwTreeModify`](#api-ccmswtreemodify) | `alipay.iservice.ccm.sw.tree.modify` |
| [`CcmSwTreePublish`](#api-ccmswtreepublish) | `alipay.iservice.ccm.sw.tree.publish` |

## 使用示例

- [CcmSwTreeBatchQuery](api_tree_ccmswtreebatchquery_example_test.go)
- [CcmSwTreeCreate](api_tree_ccmswtreecreate_example_test.go)
- [CcmSwTreeDelete](api_tree_ccmswtreedelete_example_test.go)
- [CcmSwTreeGet](api_tree_ccmswtreeget_example_test.go)
- [CcmSwTreeModify](api_tree_ccmswtreemodify_example_test.go)
- [CcmSwTreePublish](api_tree_ccmswtreepublish_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmswtreebatchquery"></a>

### CcmSwTreeBatchQuery — 知识库-类目-批量查询

- 官方接口：[`alipay.iservice.ccm.sw.tree.batchquery`](https://opendocs.alipay.com/open-v3/7ce6cb6e_alipay.iservice.ccm.sw.tree.batchquery)
- 路由：`GET /v3/alipay/iservice/ccm/sw/tree/batchquery`
- 请求类型：[CcmSwTreeBatchQueryRequest](tree_request.go#L30)
- 响应类型：[CcmSwTreeBatchQueryResponse](tree_response.go#L8)

知识库-类目-批量查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 子部门ID，不传为默认部门 |

<a id="api-ccmswtreecreate"></a>

### CcmSwTreeCreate — 知识库-类目-创建

- 官方接口：[`alipay.iservice.ccm.sw.tree.create`](https://opendocs.alipay.com/open-v3/c2bb833d_alipay.iservice.ccm.sw.tree.create)
- 路由：`PUT /v3/alipay/iservice/ccm/sw/tree/create`
- 请求类型：[CcmSwTreeCreateRequest](tree_request.go#L4)
- 响应类型：[CcmSwTreeCreateResponse](tree_response.go#L14)

知识库-类目-创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 类目名称 |

<a id="api-ccmswtreedelete"></a>

### CcmSwTreeDelete — 知识库-类目-删除

- 官方接口：[`alipay.iservice.ccm.sw.tree.delete`](https://opendocs.alipay.com/open-v3/10547684_alipay.iservice.ccm.sw.tree.delete)
- 路由：`DELETE /v3/alipay/iservice/ccm/sw/tree/delete`
- 请求类型：[CcmSwTreeDeleteRequest](tree_request.go#L36)
- 响应类型：[CcmSwTreeDeleteResponse](tree_response.go#L20)

知识库-类目-删除

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-ccmswtreeget"></a>

### CcmSwTreeGet — 知识库-类目-获取

- 官方接口：[`alipay.iservice.ccm.sw.tree.get`](https://opendocs.alipay.com/open-v3/930c8404_alipay.iservice.ccm.sw.tree.get)
- 路由：`GET /v3/alipay/iservice/ccm/sw/tree/get`
- 请求类型：[CcmSwTreeGetRequest](tree_request.go#L39)
- 响应类型：[CcmSwTreeGetResponse](tree_response.go#L23)

知识库-类目-获取

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 子部门ID，不传为默认部门 |
| `id` | `Id` | `*int64` | query | 是 | 类目ID |

<a id="api-ccmswtreemodify"></a>

### CcmSwTreeModify — 知识库-类目-修改

- 官方接口：[`alipay.iservice.ccm.sw.tree.modify`](https://opendocs.alipay.com/open-v3/b89ce966_alipay.iservice.ccm.sw.tree.modify)
- 路由：`PATCH /v3/alipay/iservice/ccm/sw/tree/modify`
- 请求类型：[CcmSwTreeModifyRequest](tree_request.go#L12)
- 响应类型：[CcmSwTreeModifyResponse](tree_response.go#L29)

知识库-类目-修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `id` | `Id` | `*int64` | JSON body | 未标注 | 类目ID |
| `name` | `Name` | `*string` | JSON body | 未标注 | 类目名称 |

<a id="api-ccmswtreepublish"></a>

### CcmSwTreePublish — 知识库-类目-发布

- 官方接口：[`alipay.iservice.ccm.sw.tree.publish`](https://opendocs.alipay.com/open-v3/3aa3bbe6_alipay.iservice.ccm.sw.tree.publish)
- 路由：`PUT /v3/alipay/iservice/ccm/sw/tree/publish`
- 请求类型：[CcmSwTreePublishRequest](tree_request.go#L22)
- 响应类型：[CcmSwTreePublishResponse](tree_response.go#L32)

知识库-类目-发布

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `id` | `Id` | `*int64` | JSON body | 未标注 | 类目ID |
