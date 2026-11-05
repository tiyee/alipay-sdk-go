# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/merchant/order`。本包提供 1 个支付宝 V3 REST 操作。

通过 `order.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`OrderSync`](#api-ordersync) | `alipay.merchant.order.sync` |

## 使用示例

- [OrderSync](api_order_ordersync_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ordersync"></a>

### OrderSync — 订单数据同步接口

- 官方接口：[`alipay.merchant.order.sync`](https://opendocs.alipay.com/open-v3/be2462f6_alipay.merchant.order.sync)
- 路由：`POST /v3/alipay/merchant/order/sync`
- 请求类型：[MerchantOrderSyncRequest](order_request.go#L4)
- 响应类型：[MerchantOrderSyncResponse](order_response.go#L4)

商户可以调用此接口同步交易对应的订单数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 订单总金额：某笔交易订单优惠前的总金额，单位为【元】 [实际案例一](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例二](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例三](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original) |
| `buyer_id` | `BuyerId` | `*string` | JSON body | 未标注 | 买家userId |
| `buyer_info` | `BuyerInfo` | [`*UserInfomation`](#model-userinfomation) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `buyer_open_id` | `BuyerOpenId` | `*string` | JSON body | 未标注 | 用户open_id |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 标准服务类目 |
| `discount_amount` | `DiscountAmount` | `*string` | JSON body | 未标注 | 商户总计优惠金额：代表商户侧给予用户的总计优惠金额 （不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】。 [实际案例一](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例二](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例三](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original) |
| `discount_info_list` | `DiscountInfoList` | [`*[]DiscountInfoData`](#model-discountinfodata) | JSON body | 未标注 | 订单优惠信息 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | JSON body | 未标注 | 扩展信息，请参见 [小程序订单中心模板](https://opendocs.alipay.com/mini/04zsxt?pathHash=654d4f10) |
| `item_order_list` | `ItemOrderList` | [`*[]ItemOrderInfo`](#model-itemorderinfo) | JSON body | 未标注 | 商品信息列表 |
| `journey_order_list` | `JourneyOrderList` | [`*[]OrderJourneyInfo`](#model-orderjourneyinfo) | JSON body | 未标注 | 行程信息 |
| `logistics_info_list` | `LogisticsInfoList` | [`*[]OrderLogisticsInformationRequest`](order_request.go#L67) | JSON body | 未标注 | 物流信息列表 |
| `order_auth_code` | `OrderAuthCode` | `*string` | JSON body | 未标注 | 本字段已废弃，无需填写！(本字段已废弃 不再使用) |
| `order_create_time` | `OrderCreateTime` | `*string` | JSON body | 未标注 | 订单创建时间 |
| `order_modified_time` | `OrderModifiedTime` | `*string` | JSON body | 未标注 | 订单修改时间 |
| `order_pay_time` | `OrderPayTime` | `*string` | JSON body | 未标注 | 订单支付时间 |
| `order_type` | `OrderType` | `*string` | JSON body | 未标注 | 订单类型 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部订单号 out_biz_no唯一对应一笔订单，相同的订单需传入相同的out_biz_no |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 交易对应的签约商户userId(注意：该字段自2020-02-16日起，可以不传入) |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 用户应付金额 ：用户最终结算时需要支付金额（不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】 [实际案例一](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*UuuWRpmekegAAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例二](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*kBkOTZpqP40AAAAAAAAAAAAAAQAAAQ/original) <br>[实际案例三](https://mdn.alipayobjects.com/portal_ykdvdu/afts/img/A*ZpkjTJQlFVAAAAAAAAAAAAAAAQAAAQ/original) |
| `pay_timeout_express` | `PayTimeoutExpress` | `*string` | JSON body | 未标注 | 支付超时时间，超过时间支付宝自行关闭订单 |
| `record_id` | `RecordId` | `*string` | JSON body | 未标注 | 商户订单同步记录id(仅部分存量接入和行业模板需要外，其他情况可以不传入) |
| `seller_id` | `SellerId` | `*string` | JSON body | 未标注 | 卖家userId(注意：该字段自2020-02-16日起，可以不传入) |
| `send_msg` | `SendMsg` | `*string` | JSON body | 未标注 | 是否需要小程序订单代理发送模版消息。不传默认不发送 |
| `service_code` | `ServiceCode` | `*string` | JSON body | 未标注 | 服务code：传入小程序后台提报的服务id，将订单与服务关联，有利于提高服务曝光机会；入参服务id的类目须与订单类型相符，若不相符将会报错；如订单类型为“外卖”，则入参的服务ID所对应的服务类目也必须得是”外卖“；service_code 通过 [alipay.open.app.service.apply](https://opendocs.alipay.com/mini/e1c835a1_alipay.open.app.service.apply) ，(服务提报申请)接口创建服务后获取。 |
| `shop_info` | `ShopInfo` | [`*OrderShopInfo`](#model-ordershopinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `source_app` | `SourceApp` | `*string` | JSON body | 未标注 | 用于区分用户下单的订单来源 |
| `sync_content` | `SyncContent` | `*string` | JSON body | 未标注 | 同步内容 |
| `ticket_info` | `TicketInfo` | [`*TicketInfo`](#model-ticketinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `ticket_order_list` | `TicketOrderList` | [`*[]TicketOrderInfo`](#model-ticketorderinfo) | JSON body | 未标注 | 凭证信息 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 订单所对应的支付宝交易号 |
| `trade_type` | `TradeType` | `*string` | JSON body | 未标注 | 交易号类型 |

## 嵌套请求参数

<a id="model-discountinfodata"></a>

### DiscountInfoData

类型定义：[DiscountInfoData](order_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 优惠金额 |
| `discount_name` | `DiscountName` | `*string` | 未标注 | 优惠名称 |
| `discount_page_link` | `DiscountPageLink` | `*string` | 未标注 | 优惠跳转链接地址 |
| `discount_quantity` | `DiscountQuantity` | `*int64` | 未标注 | 优惠数量 |
| `external_discount_id` | `ExternalDiscountId` | `*string` | 未标注 | 外部优惠id |

<a id="model-functionalservice"></a>

### FunctionalService

类型定义：[FunctionalService](order_model.go#L18)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `content` | `Content` | `*string` | 未标注 | 功能内容 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `function_code` | `FunctionCode` | `*string` | 未标注 | 功能码 支持的功能码请与产品或对应技术确认 |
| `function_name` | `FunctionName` | `*string` | 未标注 | 功能名称 |
| `function_type` | `FunctionType` | `*string` | 未标注 | 功能类型 |
| `function_url` | `FunctionUrl` | `*string` | 未标注 | 功能入口 |
| `memo` | `Memo` | `*string` | 未标注 | 备注信息 |

<a id="model-itemorderinfo"></a>

### ItemOrderInfo

类型定义：[ItemOrderInfo](order_model.go#L36)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息，请参见产品文档。小程序订单助手业务中，扩展参数必须传递素材id；其他业务场景参见对应的产品文档。 |
| `item_id` | `ItemId` | `*string` | 未标注 | 商品ID |
| `item_name` | `ItemName` | `*string` | 未标注 | 商品名称 |
| `quantity` | `Quantity` | `*int64` | 未标注 | 商品数量（单位：自拟） |
| `sku_id` | `SkuId` | `*string` | 未标注 | 商品 sku id |
| `status` | `Status` | `*string` | 未标注 | 商品状态枚举 |
| `status_desc` | `StatusDesc` | `*string` | 未标注 | 商品状态描述，默认无需传入，如需使用请联系业务负责人 |
| `unit` | `Unit` | `*string` | 未标注 | 商品规格 |
| `unit_price` | `UnitPrice` | `*string` | 未标注 | 商品单价（单位：元） |

<a id="model-journeylocation"></a>

### JourneyLocation

类型定义：[JourneyLocation](order_model.go#L58)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `aoi_id` | `AoiId` | `*string` | 未标注 | 支付宝域内aoiId |
| `city` | `City` | `*string` | 未标注 | 城市名 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `location_id` | `LocationId` | `*string` | 未标注 | 支付宝侧地点id |
| `location_id_type` | `LocationIdType` | `*string` | 未标注 | 地点id类型 |
| `merchant_division_id` | `MerchantDivisionId` | `*string` | 未标注 | 商户侧行政区划代码 |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | (废弃)商户侧地点id |
| `merchant_poi` | `MerchantPoi` | `*string` | 未标注 | 商户侧poi信息 |
| `name` | `Name` | `*string` | 未标注 | 地点名称 |
| `poi_id` | `PoiId` | `*string` | 未标注 | 支付宝域内poiId |

<a id="model-journeymerchantinfo"></a>

### JourneyMerchantInfo

类型定义：[JourneyMerchantInfo](order_model.go#L82)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `logo` | `Logo` | `*string` | 未标注 | 商户logo链接 |
| `name` | `Name` | `*string` | 未标注 | 商家名称 |
| `short_name` | `ShortName` | `*string` | 未标注 | 商家简称 |

<a id="model-journeyservicechangeinfo"></a>

### JourneyServiceChangeInfo

类型定义：[JourneyServiceChangeInfo](order_model.go#L94)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `change_status` | `ChangeStatus` | `*string` | 未标注 | 变更状态 |
| `detail_url` | `DetailUrl` | `*string` | 未标注 | 详情url |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `remind_content` | `RemindContent` | `*string` | 未标注 | 提醒内容 |

<a id="model-orderextinfo"></a>

### OrderExtInfo

类型定义：[OrderExtInfo](order_model.go#L124)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_key` | `ExtKey` | `*string` | 未标注 | 键值 |
| `ext_value` | `ExtValue` | `*string` | 未标注 | 值 |

<a id="model-orderjourneyelement"></a>

### OrderJourneyElement

类型定义：[OrderJourneyElement](order_model.go#L132)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `arrival` | `Arrival` | [`*JourneyLocation`](#model-journeylocation) | 未标注 | 见字段类型及官方接口文档。 |
| `departure` | `Departure` | [`*JourneyLocation`](#model-journeylocation) | 未标注 | 见字段类型及官方接口文档。 |
| `duration` | `Duration` | `*string` | 未标注 | 行程时长 |
| `end_time` | `EndTime` | `*string` | 未标注 | 结束时间 |
| `end_time_desc` | `EndTimeDesc` | `*string` | 未标注 | 结束时间描述（非结构化） |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `functional_services` | `FunctionalServices` | [`*[]FunctionalService`](#model-functionalservice) | 未标注 | 功能服务列表 |
| `passagers` | `Passagers` | [`*[]UserInfomation`](#model-userinfomation) | 未标注 | 出行人 |
| `service_change_info` | `ServiceChangeInfo` | [`*JourneyServiceChangeInfo`](#model-journeyservicechangeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `service_provider` | `ServiceProvider` | [`*JourneyMerchantInfo`](#model-journeymerchantinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `start_time` | `StartTime` | `*string` | 未标注 | 开始时间 |
| `start_time_desc` | `StartTimeDesc` | `*string` | 未标注 | 开始时间描述（非结构化） |

<a id="model-orderjourneyinfo"></a>

### OrderJourneyInfo

类型定义：[OrderJourneyInfo](order_model.go#L156)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action` | `Action` | `*string` | 未标注 | 操作动作 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `journey_create_time` | `JourneyCreateTime` | `*string` | 未标注 | 行程创建时间 |
| `journey_desc` | `JourneyDesc` | `*string` | 未标注 | 行程描述 |
| `journey_elements` | `JourneyElements` | [`*[]OrderJourneyElement`](#model-orderjourneyelement) | 未标注 | 行程元素列表 |
| `journey_index` | `JourneyIndex` | `*string` | 未标注 | 描述本行程为整个行程中的第几程 |
| `journey_modify_time` | `JourneyModifyTime` | `*string` | 未标注 | 行程修改时间 |
| `merchant_journey_no` | `MerchantJourneyNo` | `*string` | 未标注 | 商户行程单号 |
| `status` | `Status` | `*string` | 未标注 | 行程状态 |
| `status_desc` | `StatusDesc` | `*string` | 未标注 | 行程状态描述 |
| `sub_type` | `SubType` | `*string` | 未标注 | 行程子类型 |
| `title` | `Title` | `*string` | 未标注 | 行程标题 |
| `type` | `Type` | `*string` | 未标注 | 行程类型 |

<a id="model-orderparticipantinfo"></a>

### OrderParticipantInfo

类型定义：[OrderParticipantInfo](order_model.go#L186)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `logo_material_id` | `LogoMaterialId` | `*string` | 未标注 | 参与方logo素材id 开发者通过 alipay.merchant.item.file.upload(商品文件上传接口)上传图片，获取到对应的素材 ID( material_id ) |
| `name` | `Name` | `*string` | 未标注 | 参与方名称 |
| `short_name` | `ShortName` | `*string` | 未标注 | 参与方简称 |
| `type` | `Type` | `*string` | 未标注 | 参与方类型 |
| `uid` | `Uid` | `*string` | 未标注 | 参与方支付宝uid |

<a id="model-ordershopinfo"></a>

### OrderShopInfo

类型定义：[OrderShopInfo](order_model.go#L202)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 店铺地址 |
| `alipay_shop_id` | `AlipayShopId` | `*string` | 未标注 | 蚂蚁门店shop_id |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 门店其他业务属性，不同业务场景KEY枚举值不同，使用前请参考产品文档 |
| `merchant_shop_id` | `MerchantShopId` | `*string` | 未标注 | 商户门店id 支持英文、数字的组合 |
| `merchant_shop_link_page` | `MerchantShopLinkPage` | `*string` | 未标注 | 店铺详情链接地址 |
| `name` | `Name` | `*string` | 未标注 | 店铺名称 |
| `phone_num` | `PhoneNum` | `*string` | 未标注 | 联系电话-支持固话或手机号 仅支持数字、+、- 。例如 手机：1380***1111、固话：021-888**888 |
| `type` | `Type` | `*string` | 未标注 | 仅当alipay_shop_id字段值为非标准蚂蚁门店时使用，其他场景无需传入 |

<a id="model-ordervehicleinfo"></a>

### OrderVehicleInfo

类型定义：[OrderVehicleInfo](order_model.go#L222)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `license_plate_no` | `LicensePlateNo` | `*string` | 未标注 | 交通工具牌照号，如车牌号等 |
| `memo` | `Memo` | `*string` | 未标注 | 备注信息 |
| `shift_no` | `ShiftNo` | `*string` | 未标注 | 班次 |

<a id="model-seatinfo"></a>

### SeatInfo

类型定义：[SeatInfo](order_model.go#L232)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `seat_class` | `SeatClass` | `*string` | 未标注 | 座位等级 |
| `seat_no` | `SeatNo` | `*string` | 未标注 | 座位号 |

<a id="model-ticketinfo"></a>

### TicketInfo

类型定义：[TicketInfo](order_model.go#L240)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 店铺地址 |
| `end_time` | `EndTime` | `*string` | 未标注 | 截止时间 |
| `resv_type` | `ResvType` | `*string` | 未标注 | 凭证预约类型，一般不需要传入。可选值为INSTANT/RESERVATION，其中INSTANT代表是实时凭证，RESERVATION代表是预约凭证，不传入默认为实时凭证。区别在于预约凭证一般不是当场可取，而是用户下单后的很多天之后才可以凭借凭证提取 |
| `shop` | `Shop` | `*string` | 未标注 | 凭证可核销门店/货品自提门店，如果自提门店与购买门店不一致，可传入该字段提示用户自提门店 |
| `ticket_no` | `TicketNo` | `*string` | 未标注 | 单据号 不同类型下单据号含义不同。若类型为MEAL_NUM，则ticket_no表示取餐号；若类型为PICKUP_CODE，则ticket_no表示取件码；其他类型查看小程序订单产品文档 |
| `time` | `Time` | `*string` | 未标注 | 时间 |
| `type` | `Type` | `*string` | 未标注 | 凭证类型 具体类型查看产品文档 |

<a id="model-ticketorderinfo"></a>

### TicketOrderInfo

类型定义：[TicketOrderInfo](order_model.go#L258)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `arrival` | `Arrival` | [`*JourneyLocation`](#model-journeylocation) | 未标注 | 见字段类型及官方接口文档。 |
| `content` | `Content` | `*string` | 未标注 | 凭证内容 |
| `departure` | `Departure` | [`*JourneyLocation`](#model-journeylocation) | 未标注 | 见字段类型及官方接口文档。 |
| `effective_num` | `EffectiveNum` | `*int64` | 未标注 | 凭证有效数 |
| `end_time` | `EndTime` | `*string` | 未标注 | 凭证结束时间 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `face_value` | `FaceValue` | `*string` | 未标注 | 票面价 |
| `functional_services` | `FunctionalServices` | [`*[]FunctionalService`](#model-functionalservice) | 未标注 | 功能服务信息 |
| `invalid_reason` | `InvalidReason` | `*string` | 未标注 | 凭证失效原因 |
| `locations` | `Locations` | [`*[]JourneyLocation`](#model-journeylocation) | 未标注 | 凭证可使用地点 |
| `merchant_ticket_no` | `MerchantTicketNo` | `*string` | 未标注 | 商户侧凭证ID或核销码 |
| `seat_infos` | `SeatInfos` | [`*[]SeatInfo`](#model-seatinfo) | 未标注 | 座位信息 |
| `service_provider` | `ServiceProvider` | [`*OrderParticipantInfo`](#model-orderparticipantinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `shops` | `Shops` | [`*[]OrderShopInfo`](#model-ordershopinfo) | 未标注 | 凭证可使用门店列表 |
| `start_time` | `StartTime` | `*string` | 未标注 | 凭证开始时间 |
| `status` | `Status` | `*string` | 未标注 | 凭证状态 |
| `ticket_create_time` | `TicketCreateTime` | `*string` | 未标注 | 凭证创建时间 |
| `ticket_modify_time` | `TicketModifyTime` | `*string` | 未标注 | 凭证修改时间 |
| `ticket_num` | `TicketNum` | `*int64` | 未标注 | 凭证购买数 |
| `ticket_users` | `TicketUsers` | [`*[]UserInfomation`](#model-userinfomation) | 未标注 | 凭证使用者信息 |
| `title` | `Title` | `*string` | 未标注 | 凭证标题 |
| `vehicle_info` | `VehicleInfo` | [`*OrderVehicleInfo`](#model-ordervehicleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-userinfomation"></a>

### UserInfomation

类型定义：[UserInfomation](order_model.go#L302)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cert_no` | `CertNo` | `*string` | 未标注 | 证件号 |
| `cert_type` | `CertType` | `*string` | 未标注 | 证件类型。 |
| `ext_info` | `ExtInfo` | [`*[]OrderExtInfo`](#model-orderextinfo) | 未标注 | 扩展信息 |
| `mobile` | `Mobile` | `*string` | 未标注 | 手机号 |
| `name` | `Name` | `*string` | 未标注 | 姓名 |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户openId |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户ID，可以通过支付结果通知、用户授权等渠道获取 |
