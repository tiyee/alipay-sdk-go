# subservicekeyword 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/search/subservicekeyword`。本包提供 4 个支付宝 V3 REST 操作。

通过 `subservicekeyword.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SearchSubServiceKeywordApply`](#api-searchsubservicekeywordapply) | `alipay.open.search.subservicekeyword.apply` |
| [`SearchSubServiceKeywordBatchQuery`](#api-searchsubservicekeywordbatchquery) | `alipay.open.search.subservicekeyword.batchquery` |
| [`SearchSubServiceKeywordDelete`](#api-searchsubservicekeyworddelete) | `alipay.open.search.subservicekeyword.delete` |
| [`SearchSubServiceKeywordQueryStatus`](#api-searchsubservicekeywordquerystatus) | `alipay.open.search.subservicekeyword.querystatus` |

## 使用示例

- [SearchSubServiceKeywordApply](api_subservicekeyword_searchsubservicekeywordapply_example_test.go)
- [SearchSubServiceKeywordBatchQuery](api_subservicekeyword_searchsubservicekeywordbatchquery_example_test.go)
- [SearchSubServiceKeywordDelete](api_subservicekeyword_searchsubservicekeyworddelete_example_test.go)
- [SearchSubServiceKeywordQueryStatus](api_subservicekeyword_searchsubservicekeywordquerystatus_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-searchsubservicekeywordapply"></a>

### SearchSubServiceKeywordApply — 提报服务关键词

- 官方接口：[`alipay.open.search.subservicekeyword.apply`](https://opendocs.alipay.com/open-v3/183cd891_alipay.open.search.subservicekeyword.apply)
- 路由：`POST /v3/alipay/open/search/subservicekeyword/apply`
- 请求类型：[SearchSubServiceKeywordApplyRequest](subservicekeyword_request.go#L4)
- 响应类型：[SearchSubServiceKeywordApplyResponse](subservicekeyword_response.go#L8)

小程序-服务推广-提报服务关键词

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `config_id` | `ConfigId` | `*string` | JSON body | 未标注 | 关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取 |
| `key_words` | `KeyWords` | `*[]string` | JSON body | 未标注 | 服务关键词列表，每批最多传入30个，关键词长度小于12个汉字。 超过数量限制会申请失败，剩余关键词可通过 [alipay.open.search.appkeywordquota.query](https://opendocs.alipay.com/mini/062ndt?pathHash=e3e78b68&ref=api&scene=common)查询 |
| `sub_service_code` | `SubServiceCode` | `*string` | JSON body | 未标注 | 子服务code，提报服务关键词，alipay.open.app.service.list.query(服务批量查询)这个接口可以获取 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id |

<a id="api-searchsubservicekeywordbatchquery"></a>

### SearchSubServiceKeywordBatchQuery — 服务关键词批量查询接口

- 官方接口：[`alipay.open.search.subservicekeyword.batchquery`](https://opendocs.alipay.com/open-v3/6ee3720b_alipay.open.search.subservicekeyword.batchquery)
- 路由：`GET /v3/alipay/open/search/subservicekeyword/batchquery`
- 请求类型：[SearchSubServiceKeywordBatchQueryRequest](subservicekeyword_request.go#L24)
- 响应类型：[SearchSubServiceKeywordBatchQueryResponse](subservicekeyword_response.go#L14)

批量查询小程序的服务关键词

请求约束：SearchSubServiceKeywordBatchQueryRequest is the request type for alipay.open.search.subservicekeyword.batchquery.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_appid` | `TargetAppid` | `*string` | query | 否 | 小程序id |
| `status` | `Status` | `*string` | query | 否 | 审核状态，服务关键词的审核状态，枚举值，不传该字段表示查询所有状态 |
| `page_number` | `PageNumber` | `*string` | query | 否 | 分页参数，分页查询时使用 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页参数，分页查询时使用 |

<a id="api-searchsubservicekeyworddelete"></a>

### SearchSubServiceKeywordDelete — 删除服务关键词

- 官方接口：[`alipay.open.search.subservicekeyword.delete`](https://opendocs.alipay.com/open-v3/bf575dfc_alipay.open.search.subservicekeyword.delete)
- 路由：`POST /v3/alipay/open/search/subservicekeyword/delete`
- 请求类型：[SearchSubServiceKeywordDeleteRequest](subservicekeyword_request.go#L16)
- 响应类型：[SearchSubServiceKeywordDeleteResponse](subservicekeyword_response.go#L28)

删除服务关键词

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `config_id` | `ConfigId` | `*string` | JSON body | 未标注 | 关键词配置ID |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id |

<a id="api-searchsubservicekeywordquerystatus"></a>

### SearchSubServiceKeywordQueryStatus — 查询小程序服务关键词的审核工单的状态

- 官方接口：[`alipay.open.search.subservicekeyword.querystatus`](https://opendocs.alipay.com/open-v3/7acef1a0_alipay.open.search.subservicekeyword.querystatus)
- 路由：`GET /v3/alipay/open/search/subservicekeyword/querystatus`
- 请求类型：[SearchSubServiceKeywordQueryStatusRequest](subservicekeyword_request.go#L36)
- 响应类型：[SearchSubServiceKeywordQueryStatusResponse](subservicekeyword_response.go#L32)

查询小程序服务关键词的审核工单的状态

请求约束：SearchSubServiceKeywordQueryStatusRequest is the request type for alipay.open.search.subservicekeyword.querystatus.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_appid` | `TargetAppid` | `*string` | query | 否 | 小程序id |
| `apply_no` | `ApplyNo` | `*string` | query | 否 | 审核工单id，提报接口的返回值 |
