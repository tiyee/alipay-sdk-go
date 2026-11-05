# appkeyword 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/search/appkeyword`。本包提供 5 个支付宝 V3 REST 操作。

通过 `appkeyword.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SearchAppKeywordApply`](#api-searchappkeywordapply) | `alipay.open.search.appkeyword.apply` |
| [`SearchAppKeywordBatchQuery`](#api-searchappkeywordbatchquery) | `alipay.open.search.appkeyword.batchquery` |
| [`SearchAppKeywordDelete`](#api-searchappkeyworddelete) | `alipay.open.search.appkeyword.delete` |
| [`SearchAppKeywordQueryStatus`](#api-searchappkeywordquerystatus) | `alipay.open.search.appkeyword.querystatus` |
| [`SearchAppKeywordQuotaQuery`](#api-searchappkeywordquotaquery) | `alipay.open.search.appkeywordquota.query` |

## 使用示例

- [SearchAppKeywordApply](api_appkeyword_searchappkeywordapply_example_test.go)
- [SearchAppKeywordBatchQuery](api_appkeyword_searchappkeywordbatchquery_example_test.go)
- [SearchAppKeywordDelete](api_appkeyword_searchappkeyworddelete_example_test.go)
- [SearchAppKeywordQueryStatus](api_appkeyword_searchappkeywordquerystatus_example_test.go)
- [SearchAppKeywordQuotaQuery](api_appkeyword_searchappkeywordquotaquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-searchappkeywordapply"></a>

### SearchAppKeywordApply — 提报搜索关键词

- 官方接口：[`alipay.open.search.appkeyword.apply`](https://opendocs.alipay.com/open-v3/5ed663e2_alipay.open.search.appkeyword.apply)
- 路由：`POST /v3/alipay/open/search/appkeyword/apply`
- 请求类型：[SearchAppKeywordApplyRequest](appkeyword_request.go#L4)
- 响应类型：[SearchAppKeywordApplyResponse](appkeyword_response.go#L8)

提报搜索关键词

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `config_id` | `ConfigId` | `*string` | JSON body | 未标注 | 关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取 |
| `key_words` | `KeyWords` | `*[]string` | JSON body | 未标注 | 提报的关键词 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id |

<a id="api-searchappkeywordbatchquery"></a>

### SearchAppKeywordBatchQuery — 查询小程序已配置关键词

- 官方接口：[`alipay.open.search.appkeyword.batchquery`](https://opendocs.alipay.com/open-v3/0e3cdc41_alipay.open.search.appkeyword.batchquery)
- 路由：`GET /v3/alipay/open/search/appkeyword/batchquery`
- 请求类型：[SearchAppKeywordBatchQueryRequest](appkeyword_request.go#L24)
- 响应类型：[SearchAppKeywordBatchQueryResponse](appkeyword_response.go#L14)

查询小程序已配置关键词

请求约束：SearchAppKeywordBatchQueryRequest is the request type for alipay.open.search.appkeyword.batchquery.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_appid` | `TargetAppid` | `*string` | query | 否 | 小程序id |
| `status` | `Status` | `*string` | query | 否 | 关键词的状态，枚举值，不传该字段表示查询所有状态 |
| `page_number` | `PageNumber` | `*int64` | query | 否 | 当前页 |
| `page_size` | `PageSize` | `*int64` | query | 否 | 每页显示条数 |

<a id="api-searchappkeyworddelete"></a>

### SearchAppKeywordDelete — 删除搜索关键词

- 官方接口：[`alipay.open.search.appkeyword.delete`](https://opendocs.alipay.com/open-v3/e7b09aa2_alipay.open.search.appkeyword.delete)
- 路由：`POST /v3/alipay/open/search/appkeyword/delete`
- 请求类型：[SearchAppKeywordDeleteRequest](appkeyword_request.go#L14)
- 响应类型：[SearchAppKeywordDeleteResponse](appkeyword_response.go#L28)

删除搜索关键词

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `config_id` | `ConfigId` | `*string` | JSON body | 未标注 | 关键词配置ID |
| `key_words` | `KeyWords` | `*[]string` | JSON body | 未标注 | 需要删除的生效关键词 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id |

<a id="api-searchappkeywordquerystatus"></a>

### SearchAppKeywordQueryStatus — 查询小程序搜索关键词的审核工单的状态

- 官方接口：[`alipay.open.search.appkeyword.querystatus`](https://opendocs.alipay.com/open-v3/2028e38b_alipay.open.search.appkeyword.querystatus)
- 路由：`GET /v3/alipay/open/search/appkeyword/querystatus`
- 请求类型：[SearchAppKeywordQueryStatusRequest](appkeyword_request.go#L36)
- 响应类型：[SearchAppKeywordQueryStatusResponse](appkeyword_response.go#L32)

查询小程序搜索关键词的审核工单的状态

请求约束：SearchAppKeywordQueryStatusRequest is the request type for alipay.open.search.appkeyword.querystatus.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_appid` | `TargetAppid` | `*string` | query | 否 | 小程序id |
| `apply_no` | `ApplyNo` | `*string` | query | 否 | 审核工单id，提报接口的返回值 |

<a id="api-searchappkeywordquotaquery"></a>

### SearchAppKeywordQuotaQuery — 查询小程序可配置关键词数

- 官方接口：[`alipay.open.search.appkeywordquota.query`](https://opendocs.alipay.com/open-v3/764cf4b6_alipay.open.search.appkeywordquota.query)
- 路由：`GET /v3/alipay/open/search/appkeywordquota/query`
- 请求类型：[SearchAppKeywordQuotaQueryRequest](appkeyword_request.go#L44)
- 响应类型：[SearchAppKeywordQuotaQueryResponse](appkeyword_response.go#L38)

查询小程序可配置关键词数

请求约束：SearchAppKeywordQuotaQueryRequest is the request type for alipay.open.search.appkeywordquota.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_appid` | `TargetAppid` | `*string` | query | 否 | 小程序id |
