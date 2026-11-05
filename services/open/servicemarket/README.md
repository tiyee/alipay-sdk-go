# servicemarket 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/servicemarket`。本包提供 2 个支付宝 V3 REST 操作。

通过 `servicemarket.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`order`](order/README.md) | 8 | `ServiceMarketOrderAccept` / `ServiceMarketOrderCreate` / `ServiceMarketOrderItemCancel` / `ServiceMarketOrderItemComplete` / `ServiceMarketOrderItemConfirm` / `ServiceMarketOrderNotify` / `ServiceMarketOrderQuery` / `ServiceMarketOrderReject` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`ServiceMarketCommodityShopOffline`](#api-servicemarketcommodityshopoffline) | `alipay.open.servicemarket.commodity.shop.offline` |
| [`ServiceMarketPromotionTaskRelationQuery`](#api-servicemarketpromotiontaskrelationquery) | `alipay.open.servicemarket.promotiontask.relation.query` |

## 使用示例

- [ServiceMarketCommodityShopOffline](api_servicemarket_servicemarketcommodityshopoffline_example_test.go)
- [ServiceMarketPromotionTaskRelationQuery](api_servicemarket_servicemarketpromotiontaskrelationquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-servicemarketcommodityshopoffline"></a>

### ServiceMarketCommodityShopOffline — 门店插件下架操作

- 官方接口：[`alipay.open.servicemarket.commodity.shop.offline`](https://opendocs.alipay.com/open-v3/f56986eb_alipay.open.servicemarket.commodity.shop.offline)
- 路由：`POST /v3/alipay/open/servicemarket/commodity/shop/offline`
- 请求类型：[ServiceMarketCommodityShopOfflineRequest](servicemarket_request.go#L4)
- 响应类型：[ServiceMarketCommodityShopOfflineResponse](servicemarket_response.go#L4)

本接口需要商户授权服务商门店操作权限后，服务商可对商户门店上的插件进行下架操作

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_id` | `CommodityId` | `*string` | JSON body | 未标注 | 服务商户ID |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 门店ID |

<a id="api-servicemarketpromotiontaskrelationquery"></a>

### ServiceMarketPromotionTaskRelationQuery — 服务商推广关系查询

- 官方接口：[`alipay.open.servicemarket.promotiontask.relation.query`](https://opendocs.alipay.com/open-v3/76898459_alipay.open.servicemarket.promotiontask.relation.query)
- 路由：`GET /v3/alipay/open/servicemarket/promotiontask/relation/query`
- 请求类型：[ServiceMarketPromotionTaskRelationQueryRequest](servicemarket_request.go#L12)
- 响应类型：[ServiceMarketPromotionTaskRelationQueryResponse](servicemarket_response.go#L8)

用于开发服务商查询已与其达成推广关系的地推服务商信息

请求约束：ServiceMarketPromotionTaskRelationQueryRequest is the request type for alipay.open.servicemarket.promotiontask.relation.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_id` | `CommodityId` | `*string` | query | 否 | 服务商品ID，并不是uid，取值来源于前端的获取 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页数量，范围(1-100)。默认：10 |
| `page_num` | `PageNum` | `*string` | query | 否 | 当前页码，从1开始。默认：1 |
