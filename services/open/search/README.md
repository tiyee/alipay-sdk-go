# search 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/search`。本包提供 1 个支付宝 V3 REST 操作。

通过 `search.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`appkeyword`](appkeyword/README.md) | 5 | `SearchAppKeywordApply` / `SearchAppKeywordBatchQuery` / `SearchAppKeywordDelete` / `SearchAppKeywordQueryStatus` / `SearchAppKeywordQuotaQuery` |
| [`box`](box/README.md) | 9 | `SearchBoxApply` / `SearchBoxBatchQuery` / `SearchBoxConsult` / `SearchBoxModify` / `SearchBoxOffline` / `SearchBoxOnline` / `SearchBoxQuery` / `SearchBoxActivityApply` / `SearchBoxActivityQuery` |
| [`order`](order/README.md) | 3 | `SearchBaseOrderModify` / `SearchOrderDetailQuery` / `SearchServiceOrderBatchQuery` |
| [`subservicekeyword`](subservicekeyword/README.md) | 4 | `SearchSubServiceKeywordApply` / `SearchSubServiceKeywordBatchQuery` / `SearchSubServiceKeywordDelete` / `SearchSubServiceKeywordQueryStatus` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SearchAbilityBatchQuery`](#api-searchabilitybatchquery) | `alipay.open.search.ability.batchquery` |

## 使用示例

- [SearchAbilityBatchQuery](api_search_searchabilitybatchquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-searchabilitybatchquery"></a>

### SearchAbilityBatchQuery — 搜索运营列表分页查询

- 官方接口：[`alipay.open.search.ability.batchquery`](https://opendocs.alipay.com/pre-apis/05wggi)
- 路由：`POST /v3/alipay/open/search/ability/batchquery`
- 请求类型：[SearchAbilityBatchQueryRequest](search_request.go#L4)
- 响应类型：[SearchAbilityBatchQueryResponse](search_response.go#L4)

搜索运营列表分页查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_data` | `BizData` | [`*SearchOperPageQueryRequest`](search_request.go#L13) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 表示修改内容的业务类型 |
| `opt_type` | `OptType` | `*string` | JSON body | 未标注 | 进行具体操作的类型比如：query |
