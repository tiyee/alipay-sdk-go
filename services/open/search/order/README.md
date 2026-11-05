# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/search/order`。本包提供 3 个支付宝 V3 REST 操作。

通过 `order.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SearchBaseOrderModify`](#api-searchbaseordermodify) | `alipay.open.search.baseorder.modify` |
| [`SearchOrderDetailQuery`](#api-searchorderdetailquery) | `alipay.open.search.orderdetail.query` |
| [`SearchServiceOrderBatchQuery`](#api-searchserviceorderbatchquery) | `alipay.open.search.serviceorder.batchquery` |

## 使用示例

- [SearchBaseOrderModify](api_order_searchbaseordermodify_example_test.go)
- [SearchOrderDetailQuery](api_order_searchorderdetailquery_example_test.go)
- [SearchServiceOrderBatchQuery](api_order_searchserviceorderbatchquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-searchbaseordermodify"></a>

### SearchBaseOrderModify — 搜索运营提报基础信息工单

- 官方接口：[`alipay.open.search.baseorder.modify`](https://opendocs.alipay.com/open-v3/abbe8486_alipay.open.search.baseorder.modify)
- 路由：`POST /v3/alipay/open/search/baseorder/modify`
- 请求类型：[SearchBaseOrderModifyRequest](order_request.go#L4)
- 响应类型：[SearchBaseOrderModifyResponse](order_response.go#L4)

搜索运营提报基础信息工单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_data` | `BizData` | [`*SearchBaseOrderCreateAPIRequest`](order_request.go#L55) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 表示修改内容的业务类型 |
| `opt_type` | `OptType` | `*string` | JSON body | 未标注 | 操作的具体类型 |

<a id="api-searchorderdetailquery"></a>

### SearchOrderDetailQuery — 查询搜索服务工单的详细信息

- 官方接口：[`alipay.open.search.orderdetail.query`](https://opendocs.alipay.com/open-v3/d55fc0f0_alipay.open.search.orderdetail.query)
- 路由：`POST /v3/alipay/open/search/orderdetail/query`
- 请求类型：[SearchOrderDetailQueryRequest](order_request.go#L13)
- 响应类型：[SearchOrderDetailQueryResponse](order_response.go#L10)

查询搜索服务工单的详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_data` | `BizData` | [`*SearchOrderRequest`](order_request.go#L72) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 表示修改内容的业务类型 |
| `opt_type` | `OptType` | `*string` | JSON body | 未标注 | 进行具体操作的类型比如：create、mod、cancel |

<a id="api-searchserviceorderbatchquery"></a>

### SearchServiceOrderBatchQuery — 搜索运营服务查询接口

- 官方接口：[`alipay.open.search.serviceorder.batchquery`](https://opendocs.alipay.com/open-v3/13830d89_alipay.open.search.serviceorder.batchquery)
- 路由：`POST /v3/alipay/open/search/serviceorder/batchquery`
- 请求类型：[SearchServiceOrderBatchQueryRequest](order_request.go#L22)
- 响应类型：[SearchServiceOrderBatchQueryResponse](order_response.go#L15)

搜索运营服务查询接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_data` | `BizData` | [`*SearchApplyPageQueryRequest`](order_request.go#L31) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 表示查询内容的类型 |
| `opt_type` | `OptType` | `*string` | JSON body | 未标注 | 操作的类型，如query |
