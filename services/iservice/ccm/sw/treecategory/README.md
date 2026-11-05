# treecategory 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/treecategory`。本包提供 4 个支付宝 V3 REST 操作。

通过 `treecategory.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CcmSwTreeCategoryBatchQuery`](#api-ccmswtreecategorybatchquery) | `alipay.iservice.ccm.sw.treecategory.batchquery` |
| [`CcmSwTreeCategoryCreate`](#api-ccmswtreecategorycreate) | `alipay.iservice.ccm.sw.treecategory.create` |
| [`CcmSwTreeCategoryDelete`](#api-ccmswtreecategorydelete) | `alipay.iservice.ccm.sw.treecategory.delete` |
| [`CcmSwTreeCategoryModify`](#api-ccmswtreecategorymodify) | `alipay.iservice.ccm.sw.treecategory.modify` |

## 使用示例

- [CcmSwTreeCategoryBatchQuery](api_treecategory_ccmswtreecategorybatchquery_example_test.go)
- [CcmSwTreeCategoryCreate](api_treecategory_ccmswtreecategorycreate_example_test.go)
- [CcmSwTreeCategoryDelete](api_treecategory_ccmswtreecategorydelete_example_test.go)
- [CcmSwTreeCategoryModify](api_treecategory_ccmswtreecategorymodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmswtreecategorybatchquery"></a>

### CcmSwTreeCategoryBatchQuery — 知识库-节点-批量查询

- 官方接口：[`alipay.iservice.ccm.sw.treecategory.batchquery`](https://opendocs.alipay.com/open-v3/24848e83_alipay.iservice.ccm.sw.treecategory.batchquery)
- 路由：`GET /v3/alipay/iservice/ccm/sw/treecategory/batchquery`
- 请求类型：[CcmSwTreeCategoryBatchQueryRequest](treecategory_request.go#L34)
- 响应类型：[CcmSwTreeCategoryBatchQueryResponse](treecategory_response.go#L10)

知识库-节点-批量查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 子部门ID，不传为默认部门 |
| `library_id` | `LibraryId` | `*string` | query | 否 | 知识库ID |
| `name` | `Name` | `*string` | query | 否 | 节点名称 |
| `page_num` | `PageNum` | `*string` | query | 否 | 页数，page_size不能为空 |
| `page_size` | `PageSize` | `*string` | query | 否 | 页显示大小，page_num不能为空 |

<a id="api-ccmswtreecategorycreate"></a>

### CcmSwTreeCategoryCreate — 知识库-节点-创建

- 官方接口：[`alipay.iservice.ccm.sw.treecategory.create`](https://opendocs.alipay.com/open-v3/08f8ea8c_alipay.iservice.ccm.sw.treecategory.create)
- 路由：`PUT /v3/alipay/iservice/ccm/sw/treecategory/create`
- 请求类型：[CcmSwTreeCategoryCreateRequest](treecategory_request.go#L4)
- 响应类型：[CcmSwTreeCategoryCreateResponse](treecategory_response.go#L22)

知识库-节点-创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `description` | `Description` | `*string` | JSON body | 未标注 | 描述 |
| `father_id` | `FatherId` | `*int64` | JSON body | 未标注 | 父节点ID |
| `name` | `Name` | `*string` | JSON body | 未标注 | 节点名称 |
| `tags` | `Tags` | `*[]string` | JSON body | 未标注 | 标签。KNOWLEDGE（知识库）；PLATFORM（公有云工作台）；HELP（公有云帮助中心） |

<a id="api-ccmswtreecategorydelete"></a>

### CcmSwTreeCategoryDelete — 知识库-节点-删除

- 官方接口：[`alipay.iservice.ccm.sw.treecategory.delete`](https://opendocs.alipay.com/open-v3/cf84306f_alipay.iservice.ccm.sw.treecategory.delete)
- 路由：`DELETE /v3/alipay/iservice/ccm/sw/treecategory/delete`
- 请求类型：[CcmSwTreeCategoryDeleteRequest](treecategory_request.go#L48)
- 响应类型：[CcmSwTreeCategoryDeleteResponse](treecategory_response.go#L28)

知识库-节点-删除

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-ccmswtreecategorymodify"></a>

### CcmSwTreeCategoryModify — 知识库-节点-修改

- 官方接口：[`alipay.iservice.ccm.sw.treecategory.modify`](https://opendocs.alipay.com/open-v3/6fea3439_alipay.iservice.ccm.sw.treecategory.modify)
- 路由：`PATCH /v3/alipay/iservice/ccm/sw/treecategory/modify`
- 请求类型：[CcmSwTreeCategoryModifyRequest](treecategory_request.go#L18)
- 响应类型：[CcmSwTreeCategoryModifyResponse](treecategory_response.go#L31)

知识库-节点-修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `description` | `Description` | `*string` | JSON body | 未标注 | 描述 |
| `father_id` | `FatherId` | `*string` | JSON body | 未标注 | 父节点ID |
| `id` | `Id` | `*int64` | JSON body | 未标注 | 节点ID |
| `name` | `Name` | `*string` | JSON body | 未标注 | 节点名称 |
| `tags` | `Tags` | `*string` | JSON body | 未标注 | 标签。KNOWLEDGE（知识库）；PLATFORM（公有云工作台）；HELP（公有云帮助中心） |
