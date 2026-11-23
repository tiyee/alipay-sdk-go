# open/mini/tips 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/tips`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `tips.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniTipsDeliveryBatchQuery`](#api-minitipsdeliverybatchquery) | [`alipay.open.mini.tips.delivery.batchquery`](https://opendocs.alipay.com/open-v3/75394e37_alipay.open.mini.tips.delivery.batchquery) | 小程序收藏引导投放活动配置批量查询 | `POST` |
| [`MiniTipsDeliveryCreate`](#api-minitipsdeliverycreate) | [`alipay.open.mini.tips.delivery.create`](https://opendocs.alipay.com/open-v3/ca9b74a9_alipay.open.mini.tips.delivery.create) | 小程序收藏引导投放活动配置创建 | `POST` |
| [`MiniTipsDeliveryModify`](#api-minitipsdeliverymodify) | [`alipay.open.mini.tips.delivery.modify`](https://opendocs.alipay.com/open-v3/a4b55863_alipay.open.mini.tips.delivery.modify) | 小程序收藏引导投放活动修改 | `POST` |
| [`MiniTipsDeliveryQuery`](#api-minitipsdeliveryquery) | [`alipay.open.mini.tips.delivery.query`](https://opendocs.alipay.com/open-v3/f59fb71e_alipay.open.mini.tips.delivery.query) | 小程序收藏引导投放活动详情查询 | `GET` |
| [`MiniTipsStatisticQuery`](#api-minitipsstatisticquery) | [`alipay.open.mini.tips.statistic.query`](https://opendocs.alipay.com/open-v3/87364a7f_alipay.open.mini.tips.statistic.query) | 小程序收藏引导汇总数据查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-minitipsdeliverybatchquery"></a>

### MiniTipsDeliveryBatchQuery — 小程序收藏引导投放活动配置批量查询

- 官方接口：[`alipay.open.mini.tips.delivery.batchquery`](https://opendocs.alipay.com/open-v3/75394e37_alipay.open.mini.tips.delivery.batchquery)
- 路由：`POST /v3/alipay/open/mini/tips/delivery/batchquery`
- 请求类型：[MiniTipsDeliveryBatchQueryRequest](tips_request.go#L5)
- 响应类型：[MiniTipsDeliveryBatchQueryResponse](tips_response.go#L5)

小程序收藏引导投放活动配置批量查询接口，获取该小程序下所有活动配置及状态信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_number` | `PageNumber` | `*string` | JSON body | 未标注 | 每页记录条数，最小1，最大20 |
| `page_size` | `PageSize` | `*string` | JSON body | 未标注 | 查询的页数，从1开始，最大1000 |

<a id="api-minitipsdeliverycreate"></a>

### MiniTipsDeliveryCreate — 小程序收藏引导投放活动配置创建

- 官方接口：[`alipay.open.mini.tips.delivery.create`](https://opendocs.alipay.com/open-v3/ca9b74a9_alipay.open.mini.tips.delivery.create)
- 路由：`POST /v3/alipay/open/mini/tips/delivery/create`
- 请求类型：[MiniTipsDeliveryCreateRequest](tips_request.go#L13)
- 响应类型：[MiniTipsDeliveryCreateResponse](tips_response.go#L15)

小程序收藏引导投放活动配置创建。通过配置有效的小程序收藏引导（tips）文案，创建小程序收藏引导活动。用户在活动有效期内等于相应页面，将展示配置的个性化文案，帮助用户收藏小程序。请结合您的小程序及页面，配置个性化文案。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `delivery_content` | `DeliveryContent` | `*string` | JSON body | 未标注 | 收藏引导文案内容，不得超过14个字 |
| `delivery_name` | `DeliveryName` | `*string` | JSON body | 未标注 | 收藏引导活动名称 |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 活动结束时间 |
| `match_type` | `MatchType` | `*string` | JSON body | 未标注 | 匹配类型 |
| `match_url` | `MatchUrl` | `*string` | JSON body | 未标注 | 目标页面地址，当匹配类型为TARGETURL时，投放文案仅在当前页面生效 |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 活动开始时间 |

<a id="api-minitipsdeliverymodify"></a>

### MiniTipsDeliveryModify — 小程序收藏引导投放活动修改

- 官方接口：[`alipay.open.mini.tips.delivery.modify`](https://opendocs.alipay.com/open-v3/a4b55863_alipay.open.mini.tips.delivery.modify)
- 路由：`POST /v3/alipay/open/mini/tips/delivery/modify`
- 请求类型：[MiniTipsDeliveryModifyRequest](tips_request.go#L29)
- 响应类型：[MiniTipsDeliveryModifyResponse](tips_response.go#L21)

小程序收藏引导投放活动修改接口，目前支持活动状态的修改，包括暂停、恢复、终止单个收藏引导活动。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `delivery_id` | `DeliveryId` | `*string` | JSON body | 未标注 | 收藏引导投放活动ID |
| `operate_type` | `OperateType` | `*string` | JSON body | 未标注 | 活动投放操作类型 |

<a id="api-minitipsdeliveryquery"></a>

### MiniTipsDeliveryQuery — 小程序收藏引导投放活动详情查询

- 官方接口：[`alipay.open.mini.tips.delivery.query`](https://opendocs.alipay.com/open-v3/f59fb71e_alipay.open.mini.tips.delivery.query)
- 路由：`GET /v3/alipay/open/mini/tips/delivery/query`
- 请求类型：[MiniTipsDeliveryQueryRequest](tips_request.go#L37)
- 响应类型：[MiniTipsDeliveryQueryResponse](tips_response.go#L25)

小程序收藏引导投放活动配置详情查询接口，获取当前活动ID对应活动的配置及状态信息。

请求约束：MiniTipsDeliveryQueryRequest is the request type for alipay.open.mini.tips.delivery.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `delivery_id` | `DeliveryId` | `*string` | query | 否 | 收藏引导投放活动ID，查询收藏引导活动必填参数 |

<a id="api-minitipsstatisticquery"></a>

### MiniTipsStatisticQuery — 小程序收藏引导汇总数据查询

- 官方接口：[`alipay.open.mini.tips.statistic.query`](https://opendocs.alipay.com/open-v3/87364a7f_alipay.open.mini.tips.statistic.query)
- 路由：`GET /v3/alipay/open/mini/tips/statistic/query`
- 请求类型：[MiniTipsStatisticQueryRequest](tips_request.go#L43)
- 响应类型：[MiniTipsStatisticQueryResponse](tips_response.go#L47)

小程序收藏引导tips文案投放汇总数据查询接口。可查询小程序维度或活动维度的tips曝光uv，收藏uv，以及收藏转化率。请先配置投放活动，否则查询结果返回参数为空。

请求约束：MiniTipsStatisticQueryRequest is the request type for alipay.open.mini.tips.statistic.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `delivery_id` | `DeliveryId` | `*string` | query | 否 | 收藏引导投放活动ID，供查询收藏引导活动配置接口调用 ，当以小程序维度查询数据（query_type为app）时delivery_id为空 |
| `query_type` | `QueryType` | `*string` | query | 否 | 查询类型，表示以当前维度进行数据聚合。 |
| `start_date` | `StartDate` | `*string` | query | 否 | 查询开始日期，精度为天 |
| `end_date` | `EndDate` | `*string` | query | 否 | 查询截止日期，精度为天 |

## 使用示例

- [MiniTipsDeliveryQuery](api_tips_minitipsdeliveryquery_example_test.go)
- [MiniTipsDeliveryBatchQuery](api_tips_minitipsdeliverybatchquery_example_test.go)
- [MiniTipsDeliveryCreate](api_tips_minitipsdeliverycreate_example_test.go)
- [MiniTipsDeliveryModify](api_tips_minitipsdeliverymodify_example_test.go)
- [MiniTipsStatisticQuery](api_tips_minitipsstatisticquery_example_test.go)
