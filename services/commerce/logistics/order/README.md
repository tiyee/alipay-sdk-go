# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/logistics/order`。本包提供 5 个支付宝 V3 REST 操作。

通过 `order.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`LogisticsOrderInstantDeliveryCancel`](#api-logisticsorderinstantdeliverycancel) | `alipay.commerce.logistics.order.instantdelivery.cancel` |
| [`LogisticsOrderInstantDeliveryCreate`](#api-logisticsorderinstantdeliverycreate) | `alipay.commerce.logistics.order.instantdelivery.create` |
| [`LogisticsOrderInstantDeliveryPreCreate`](#api-logisticsorderinstantdeliveryprecreate) | `alipay.commerce.logistics.order.instantdelivery.precreate` |
| [`LogisticsOrderIstdCancelPreConsult`](#api-logisticsorderistdcancelpreconsult) | `alipay.commerce.logistics.order.istdcancel.preconsult` |
| [`LogisticsOrderIstdRetryCreate`](#api-logisticsorderistdretrycreate) | `alipay.commerce.logistics.order.istdretry.create` |

## 使用示例

- [LogisticsOrderInstantDeliveryCancel](api_order_logisticsorderinstantdeliverycancel_example_test.go)
- [LogisticsOrderInstantDeliveryCreate](api_order_logisticsorderinstantdeliverycreate_example_test.go)
- [LogisticsOrderInstantDeliveryPreCreate](api_order_logisticsorderinstantdeliveryprecreate_example_test.go)
- [LogisticsOrderIstdCancelPreConsult](api_order_logisticsorderistdcancelpreconsult_example_test.go)
- [LogisticsOrderIstdRetryCreate](api_order_logisticsorderistdretrycreate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-logisticsorderinstantdeliverycancel"></a>

### LogisticsOrderInstantDeliveryCancel — 取消即时配送订单

- 官方接口：[`alipay.commerce.logistics.order.instantdelivery.cancel`](https://opendocs.alipay.com/pre-apis/068eft)
- 路由：`POST /v3/alipay/commerce/logistics/order/instantdelivery/cancel`
- 请求类型：[LogisticsOrderInstantDeliveryCancelRequest](order_request.go#L4)
- 响应类型：[LogisticsOrderInstantDeliveryCancelResponse](order_response.go#L4)

取消即时配送订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cancel_reason` | `CancelReason` | `*string` | JSON body | 未标注 | 取消原因，取消原因id为i_6时必填 |
| `cancel_reason_id` | `CancelReasonId` | `*string` | JSON body | 未标注 | 取消原因id |
| `logistics_code` | `LogisticsCode` | `*string` | JSON body | 未标注 | 即时配送公司编码 |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 支付宝订单流水号 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商家订单号，与order_no不能同时为空 |
| `waybill_no` | `WaybillNo` | `*string` | JSON body | 未标注 | 即时配送运单编号 |

<a id="api-logisticsorderinstantdeliverycreate"></a>

### LogisticsOrderInstantDeliveryCreate — 下即时配送订单

- 官方接口：[`alipay.commerce.logistics.order.instantdelivery.create`](https://opendocs.alipay.com/pre-apis/05lo9w)
- 路由：`POST /v3/alipay/commerce/logistics/order/instantdelivery/create`
- 请求类型：[LogisticsOrderInstantDeliveryCreateRequest](order_request.go#L20)
- 响应类型：[LogisticsOrderInstantDeliveryCreateResponse](order_response.go#L10)

下即时配送订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `consumer_id` | `ConsumerId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空 |
| `consumer_notify` | `ConsumerNotify` | [`*ConsumerNotifyIstd`](#model-consumernotifyistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `consumer_source` | `ConsumerSource` | `*string` | JSON body | 未标注 | 消费者来源,支付宝：alipay; 微信：wx |
| `goods_details` | `GoodsDetails` | [`*[]GoodsDetailIstd`](#model-goodsdetailistd) | JSON body | 未标注 | 商品明细 |
| `goods_info` | `GoodsInfo` | [`*GoodsInfoIstd`](#model-goodsinfoistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `logistics_code` | `LogisticsCode` | `*string` | JSON body | 未标注 | 即时配送公司编码 |
| `logistics_token` | `LogisticsToken` | `*string` | JSON body | 未标注 | 下即时配送订单token, 配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户openId；如果consumer_source是wx，则consumer_id可以为空 |
| `order_ext_istd` | `OrderExtIstd` | [`*OrderExtIstd`](#model-orderextistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商家订单号 |
| `receiver` | `Receiver` | [`*ReceiverIstd`](#model-receiveristd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `sender` | `Sender` | [`*SenderIstd`](#model-senderistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编号 |

<a id="api-logisticsorderinstantdeliveryprecreate"></a>

### LogisticsOrderInstantDeliveryPreCreate — 预下即时配送订单

- 官方接口：[`alipay.commerce.logistics.order.instantdelivery.precreate`](https://opendocs.alipay.com/pre-apis/05l560)
- 路由：`POST /v3/alipay/commerce/logistics/order/instantdelivery/precreate`
- 请求类型：[LogisticsOrderInstantDeliveryPreCreateRequest](order_request.go#L45)
- 响应类型：[LogisticsOrderInstantDeliveryPreCreateResponse](order_response.go#L38)

预下即时配送订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `consumer_id` | `ConsumerId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空 |
| `consumer_notify` | `ConsumerNotify` | [`*ConsumerNotifyIstd`](#model-consumernotifyistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `consumer_source` | `ConsumerSource` | `*string` | JSON body | 未标注 | 消费者来源, 支付宝：alipay, 微信：wx |
| `goods_details` | `GoodsDetails` | [`*[]GoodsDetailIstd`](#model-goodsdetailistd) | JSON body | 未标注 | 商品明细 |
| `goods_info` | `GoodsInfo` | [`*GoodsInfoIstd`](#model-goodsinfoistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `logistics_companies` | `LogisticsCompanies` | [`*[]LogisticsCompanyIstd`](#model-logisticscompanyistd) | JSON body | 未标注 | 即时配送公司配置列表 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝的openId；如果consumer_source是wx，则consumer_id可以为空 |
| `order_ext_istd` | `OrderExtIstd` | [`*OrderExtIstdForPreOrder`](#model-orderextistdforpreorder) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商家订单号 |
| `receiver` | `Receiver` | [`*ReceiverIstd`](#model-receiveristd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `sender` | `Sender` | [`*SenderIstd`](#model-senderistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编号 |

<a id="api-logisticsorderistdcancelpreconsult"></a>

### LogisticsOrderIstdCancelPreConsult — 预取消即时配送订单

- 官方接口：[`alipay.commerce.logistics.order.istdcancel.preconsult`](https://opendocs.alipay.com/pre-apis/068d52)
- 路由：`POST /v3/alipay/commerce/logistics/order/istdcancel/preconsult`
- 请求类型：[LogisticsOrderIstdCancelPreConsultRequest](order_request.go#L68)
- 响应类型：[LogisticsOrderIstdCancelPreConsultResponse](order_response.go#L44)

预取消即时配送订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cancel_reason` | `CancelReason` | `*string` | JSON body | 未标注 | 取消原因，取消原因id为i_6时必填 |
| `cancel_reason_id` | `CancelReasonId` | `*string` | JSON body | 未标注 | 取消原因id |
| `logistics_code` | `LogisticsCode` | `*string` | JSON body | 未标注 | 即时配送公司编码 |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 支付宝订单流水号 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商家订单号，与order_no不能同时为空 |
| `waybill_no` | `WaybillNo` | `*string` | JSON body | 未标注 | 即时配送运单编号 |

<a id="api-logisticsorderistdretrycreate"></a>

### LogisticsOrderIstdRetryCreate — 重下即时配送订单

- 官方接口：[`alipay.commerce.logistics.order.istdretry.create`](https://opendocs.alipay.com/pre-apis/05l561)
- 路由：`POST /v3/alipay/commerce/logistics/order/istdretry/create`
- 请求类型：[LogisticsOrderIstdRetryCreateRequest](order_request.go#L84)
- 响应类型：[LogisticsOrderIstdRetryCreateResponse](order_response.go#L52)

重下即时配送订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `consumer_id` | `ConsumerId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户uid；如果consumer_source是wx，则consumer_id可以为空 |
| `consumer_notify` | `ConsumerNotify` | [`*ConsumerNotifyIstd`](#model-consumernotifyistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `consumer_source` | `ConsumerSource` | `*string` | JSON body | 未标注 | 消费者来源, 支付宝：alipay, 微信：wx |
| `goods_details` | `GoodsDetails` | [`*[]GoodsDetailIstd`](#model-goodsdetailistd) | JSON body | 未标注 | 商品明细 |
| `goods_info` | `GoodsInfo` | [`*GoodsInfoIstd`](#model-goodsinfoistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `logistics_code` | `LogisticsCode` | `*string` | JSON body | 未标注 | 即时配送公司编码 |
| `logistics_token` | `LogisticsToken` | `*string` | JSON body | 未标注 | 下即时配送订单token, 配送公司可以返回此字段，当商家下单时候带上这个字段，保证在一段时间内运费不变 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 消费者id， 如果consumer_source是alipay，则consumer_id必须是支付宝用户openId；如果consumer_source是wx，则consumer_id可以为空 |
| `order_ext_istd` | `OrderExtIstd` | [`*OrderExtIstd`](#model-orderextistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商家订单号 |
| `receiver` | `Receiver` | [`*ReceiverIstd`](#model-receiveristd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `sender` | `Sender` | [`*SenderIstd`](#model-senderistd) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编号 |

## 嵌套请求参数

<a id="model-consumernotifyistd"></a>

### ConsumerNotifyIstd

类型定义：[ConsumerNotifyIstd](order_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_count` | `GoodsCount` | `*int64` | 未标注 | 商品数量 |
| `goods_img` | `GoodsImg` | `*string` | 未标注 | 商品缩略图url，支持格式：bmp、jpg、jpeg、png、gif |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称 |
| `merchant_mobile` | `MerchantMobile` | `*string` | 未标注 | 商家电话 |
| `merchant_name` | `MerchantName` | `*string` | 未标注 | 商家名称，tiny_app_id和merchant_name不能同时为空 |
| `tiny_app_id` | `TinyAppId` | `*string` | 未标注 | 商家小程序appid |
| `tiny_app_url` | `TinyAppUrl` | `*string` | 未标注 | 商家小程序的路径，建议为订单页面 |

<a id="model-goodsdetailistd"></a>

### GoodsDetailIstd

类型定义：[GoodsDetailIstd](order_model.go#L22)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `count` | `Count` | `*int64` | 未标注 | 商品数量 |
| `name` | `Name` | `*string` | 未标注 | 商品名称 |
| `price` | `Price` | `*string` | 未标注 | 商品单价 |
| `unit` | `Unit` | `*string` | 未标注 | 商品单位 |

<a id="model-goodsinfoistd"></a>

### GoodsInfoIstd

类型定义：[GoodsInfoIstd](order_model.go#L34)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_info` | `DeliveryInfo` | `*string` | 未标注 | 交货信息，用于骑手交付商品 |
| `first_class` | `FirstClass` | `*string` | 未标注 | 一级品类类目 |
| `height` | `Height` | `*int64` | 未标注 | 商品高度，单位为cm |
| `length` | `Length` | `*int64` | 未标注 | 商品长度，单位为cm |
| `pickup_info` | `PickupInfo` | `*string` | 未标注 | 取货信息，用于骑手到店取货 |
| `price` | `Price` | `*string` | 未标注 | 总价 |
| `second_class` | `SecondClass` | `*string` | 未标注 | 二级品类类目 |
| `weight` | `Weight` | `*string` | 未标注 | 货物重量，单位为kg，精确到小数点后两位 |
| `width` | `Width` | `*int64` | 未标注 | 商品宽度，单位为cm |

<a id="model-logisticscompanyistd"></a>

### LogisticsCompanyIstd

类型定义：[LogisticsCompanyIstd](order_model.go#L56)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `logistics_code` | `LogisticsCode` | `*string` | 未标注 | 即时配送公司编码，由支付宝分配 |
| `service_codes` | `ServiceCodes` | [`*[]ServiceCodeIstd`](#model-servicecodeistd) | 未标注 | 即时配送公司服务代码列表 |

<a id="model-orderextistd"></a>

### OrderExtIstd

类型定义：[OrderExtIstd](order_model.go#L64)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cash_on_delivery` | `CashOnDelivery` | `*string` | 未标注 | 骑手应付金额 |
| `cash_on_pickup` | `CashOnPickup` | `*string` | 未标注 | 骑手应收金额 |
| `consumer_order_time` | `ConsumerOrderTime` | `*string` | 未标注 | 消费者下单时间 |
| `delivery_direction` | `DeliveryDirection` | `*int64` | 未标注 | 物流流向，1：从门店取件送至用户；2：从用户取件送至门店 |
| `delivery_type` | `DeliveryType` | `*int64` | 未标注 | 配送类型, 0: 即时单 1 预约单 |
| `desc` | `Desc` | `*string` | 未标注 | 备注 |
| `expected_delivery_time` | `ExpectedDeliveryTime` | `*string` | 未标注 | 期望派单时间 |
| `expected_finish_time` | `ExpectedFinishTime` | `*string` | 未标注 | 期望送达时间 |
| `expected_pick_time` | `ExpectedPickTime` | `*string` | 未标注 | 期望取件时间 |
| `insure_price` | `InsurePrice` | `*string` | 未标注 | 保价金额 |
| `is_direct_delivery` | `IsDirectDelivery` | `*int64` | 未标注 | 是否直拿直送，1:直拿直送 0:非直拿直送 |
| `is_finish_code_needed` | `IsFinishCodeNeeded` | `*int64` | 未标注 | 是否需要收货码，1:需要 0:不需要 |
| `is_insured` | `IsInsured` | `*int64` | 未标注 | 是否保价，1:保价 0:不保价 |
| `is_pickup_code_needed` | `IsPickupCodeNeeded` | `*int64` | 未标注 | 是否需要取货码，1:需要 0:不需要 |
| `poi_seq` | `PoiSeq` | `*string` | 未标注 | 门店订单流水号 |
| `service_code` | `ServiceCode` | `*string` | 未标注 | 即时配送公司的服务代码 |

<a id="model-orderextistdforpreorder"></a>

### OrderExtIstdForPreOrder

类型定义：[OrderExtIstdForPreOrder](order_model.go#L100)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cash_on_delivery` | `CashOnDelivery` | `*string` | 未标注 | 骑手应付金额，单位元 |
| `cash_on_pickup` | `CashOnPickup` | `*string` | 未标注 | 骑手应收金额，单位元 |
| `consumer_order_time` | `ConsumerOrderTime` | `*string` | 未标注 | 消费者下单时间 |
| `delivery_direction` | `DeliveryDirection` | `*int64` | 未标注 | 物流流向，1：从门店取件送至用户；2：从用户取件送至门店 |
| `delivery_type` | `DeliveryType` | `*int64` | 未标注 | 配送类型, 0: 即时单 1 预约单 |
| `desc` | `Desc` | `*string` | 未标注 | 备注 |
| `expected_delivery_time` | `ExpectedDeliveryTime` | `*string` | 未标注 | 期望派单时间 |
| `expected_finish_time` | `ExpectedFinishTime` | `*string` | 未标注 | 期望送达时间 |
| `expected_pick_time` | `ExpectedPickTime` | `*string` | 未标注 | 期望取件时间 |
| `insure_price` | `InsurePrice` | `*string` | 未标注 | 保价金额 |
| `is_direct_delivery` | `IsDirectDelivery` | `*int64` | 未标注 | 是否直拿直送，1:直拿直送 0:非直拿直送 |
| `is_finish_code_needed` | `IsFinishCodeNeeded` | `*int64` | 未标注 | 是否需要收货码，1:需要 0:不需要 |
| `is_insured` | `IsInsured` | `*int64` | 未标注 | 是否保价，1:保价 0:不保价 |
| `is_pickup_code_needed` | `IsPickupCodeNeeded` | `*int64` | 未标注 | 是否需要取货码，1:需要 0:不需要 |
| `poi_seq` | `PoiSeq` | `*string` | 未标注 | 门店订单流水号 |
| `service_code` | `ServiceCode` | `*string` | 未标注 | 即时配送公司的服务代码 |

<a id="model-receiveristd"></a>

### ReceiverIstd

类型定义：[ReceiverIstd](order_model.go#L166)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 地址(街道、小区、大厦等，用于定位) |
| `address_detail` | `AddressDetail` | `*string` | 未标注 | 地址详情(楼号、单元号、层号) |
| `city` | `City` | `*string` | 未标注 | 城市名称，如杭州市 |
| `coordinate_type` | `CoordinateType` | `*int64` | 未标注 | 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标。目前只支持0:火星坐标 |
| `lat` | `Lat` | `*string` | 未标注 | 纬度 |
| `lng` | `Lng` | `*string` | 未标注 | 经度 |
| `mobile_no` | `MobileNo` | `*string` | 未标注 | 手机号 |
| `name` | `Name` | `*string` | 未标注 | 姓名 |

<a id="model-senderistd"></a>

### SenderIstd

类型定义：[SenderIstd](order_model.go#L186)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 地址(街道、小区、大厦等，用于定位) |
| `address_detail` | `AddressDetail` | `*string` | 未标注 | 地址详情(楼号、单元号、层号) |
| `city` | `City` | `*string` | 未标注 | 城市名称，如杭州市 |
| `coordinate_type` | `CoordinateType` | `*int64` | 未标注 | 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1:百度坐标， 目前只支持0 |
| `lat` | `Lat` | `*string` | 未标注 | 纬度 |
| `lng` | `Lng` | `*string` | 未标注 | 经度 |
| `mobile_no` | `MobileNo` | `*string` | 未标注 | 手机号 |
| `name` | `Name` | `*string` | 未标注 | 姓名 |

<a id="model-servicecodeistd"></a>

### ServiceCodeIstd

类型定义：[ServiceCodeIstd](order_model.go#L206)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `service_code` | `ServiceCode` | `*string` | 未标注 | 即时配送公司服务代码 |
