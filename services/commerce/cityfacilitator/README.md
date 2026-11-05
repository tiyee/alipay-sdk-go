# commerce/cityfacilitator 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/cityfacilitator`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `cityfacilitator.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_cityfacilitator_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CityFacilitatorStationQuery`](#api-cityfacilitatorstationquery) | [`alipay.commerce.cityfacilitator.station.query`](https://opendocs.alipay.com/apis/05wvuo) | 地铁购票站点数据查询 | `GET` |
| [`CityFacilitatorVoucherBatchQuery`](#api-cityfacilitatorvoucherbatchquery) | [`alipay.commerce.cityfacilitator.voucher.batchquery`](https://opendocs.alipay.com/apis/05wol9) | 地铁购票订单批量查询 | `POST` |
| [`CityFacilitatorVoucherGenerate`](#api-cityfacilitatorvouchergenerate) | [`alipay.commerce.cityfacilitator.voucher.generate`](https://opendocs.alipay.com/apis/05wz9v) | 地铁购票核销码发码 | `POST` |
| [`CityFacilitatorVoucherRefund`](#api-cityfacilitatorvoucherrefund) | [`alipay.commerce.cityfacilitator.voucher.refund`](https://opendocs.alipay.com/apis/05wujr) | 地铁购票发码退款 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-cityfacilitatorstationquery"></a>

### CityFacilitatorStationQuery — 地铁购票站点数据查询

- 官方接口：[`alipay.commerce.cityfacilitator.station.query`](https://opendocs.alipay.com/apis/05wvuo)
- 路由：`GET /v3/alipay/commerce/cityfacilitator/station/query`
- 请求类型：[CityFacilitatorStationQueryRequest](cityfacilitator_request.go#L41)
- 响应类型：[CityFacilitatorStationQueryResponse](cityfacilitator_response.go#L9)

商户App获取地铁购票开放购票的站点和禁止到达的站点

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `city_code` | `CityCode` | `*string` | query | 否 | 城市编码请参考查询 [中华人民共和国行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)。 已支持城市：广州 440100，深圳 440300，杭州330100。 |

<a id="api-cityfacilitatorvoucherbatchquery"></a>

### CityFacilitatorVoucherBatchQuery — 地铁购票订单批量查询

- 官方接口：[`alipay.commerce.cityfacilitator.voucher.batchquery`](https://opendocs.alipay.com/apis/05wol9)
- 路由：`POST /v3/alipay/commerce/cityfacilitator/voucher/batchquery`
- 请求类型：[CityFacilitatorVoucherBatchQueryRequest](cityfacilitator_request.go#L5)
- 响应类型：[CityFacilitatorVoucherBatchQueryResponse](cityfacilitator_response.go#L15)

商户APP调该接口批量查询订单状态

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码请参考查询 [中华人民共和国行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)。 已支持城市：广州 440100，深圳 440300，杭州330100。 |
| `trade_nos` | `TradeNos` | `*[]string` | JSON body | 未标注 | 支付宝交易号列表 |

<a id="api-cityfacilitatorvouchergenerate"></a>

### CityFacilitatorVoucherGenerate — 地铁购票核销码发码

- 官方接口：[`alipay.commerce.cityfacilitator.voucher.generate`](https://opendocs.alipay.com/apis/05wz9v)
- 路由：`POST /v3/alipay/commerce/cityfacilitator/voucher/generate`
- 请求类型：[CityFacilitatorVoucherGenerateRequest](cityfacilitator_request.go#L13)
- 响应类型：[CityFacilitatorVoucherGenerateResponse](cityfacilitator_response.go#L21)

商户APP调快捷支付付款后，获取地铁购票的核销码

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码请参考查询 [中华人民共和国行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)。 已支持城市：广州 440100，深圳 440300，杭州330100。 |
| `site_begin` | `SiteBegin` | `*string` | JSON body | 未标注 | 起点站站点编码 |
| `site_end` | `SiteEnd` | `*string` | JSON body | 未标注 | 终点站站点编码 |
| `ticket_num` | `TicketNum` | `*string` | JSON body | 未标注 | 地铁票购票数量 |
| `ticket_price` | `TicketPrice` | `*string` | JSON body | 未标注 | 单张票价，元为单价 |
| `ticket_type` | `TicketType` | `*string` | JSON body | 未标注 | 地铁票种类，枚举支持： *oneway。 |
| `total_fee` | `TotalFee` | `*string` | JSON body | 未标注 | 订单总金额，元为单位 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝交易号（交易支付时，必须通过指定sellerId：2088121612215201，将钱支付到指定的中间户中） |

<a id="api-cityfacilitatorvoucherrefund"></a>

### CityFacilitatorVoucherRefund — 地铁购票发码退款

- 官方接口：[`alipay.commerce.cityfacilitator.voucher.refund`](https://opendocs.alipay.com/apis/05wujr)
- 路由：`POST /v3/alipay/commerce/cityfacilitator/voucher/refund`
- 请求类型：[CityFacilitatorVoucherRefundRequest](cityfacilitator_request.go#L33)
- 响应类型：[CityFacilitatorVoucherRefundResponse](cityfacilitator_response.go#L31)

地铁购票，商户发码后，调该接口从中间户退款

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码请参考查询 [中华人民共和国行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)。 已支持城市：广州 440100，深圳 440300，杭州330100。 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝交易号 |
