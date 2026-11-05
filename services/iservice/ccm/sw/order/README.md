# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/order`。本包提供 1 个支付宝 V3 REST 操作。

通过 `order.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CcmSwOrderSync`](#api-ccmswordersync) | `alipay.iservice.ccm.sw.order.sync` |

## 使用示例

- [CcmSwOrderSync](api_order_ccmswordersync_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmswordersync"></a>

### CcmSwOrderSync — 商品订单同步

- 官方接口：[`alipay.iservice.ccm.sw.order.sync`](https://opendocs.alipay.com/apis/0630z3)
- 路由：`POST /v3/alipay/iservice/ccm/sw/order/sync`
- 请求类型：[CcmSwOrderSyncRequest](order_request.go#L4)
- 响应类型：[CcmSwOrderSyncResponse](order_response.go#L4)

商品订单同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 订单金额 |
| `link_url` | `LinkUrl` | `*string` | JSON body | 未标注 | 订单链接 |
| `logistic_count` | `LogisticCount` | `*int64` | JSON body | 未标注 | 订单物流数量 |
| `logistics` | `Logistics` | [`*[]LogisticInfo`](#model-logisticinfo) | JSON body | 未标注 | 物流信息列表 |
| `order_create_time` | `OrderCreateTime` | `*string` | JSON body | 未标注 | 订单创建时间 |
| `order_id` | `OrderId` | `*string` | JSON body | 未标注 | 订单id |
| `order_type` | `OrderType` | `*string` | JSON body | 未标注 | 普通订单：NORMAL 预售订单：PRE_SALE |
| `spu_count` | `SpuCount` | `*int64` | JSON body | 未标注 | 订单商品种类 |
| `spus` | `Spus` | [`*[]SPUInfo`](#model-spuinfo) | JSON body | 未标注 | 订单商品信息 |
| `status` | `Status` | `*string` | JSON body | 未标注 | 订单状态，目前支持以下几种状态 下单未支付 WAIT_PAY 支付未发货 PAIED 已发货 IN_DELIVERY 售后中 IN_AFTER_SALE 订单完成 FINISHED |
| `sub_status` | `SubStatus` | `*string` | JSON body | 未标注 | 订单子状态 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 下单用户id(外部系统ID) |

## 嵌套请求参数

<a id="model-logisticinfo"></a>

### LogisticInfo

类型定义：[LogisticInfo](order_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `channel` | `Channel` | `*string` | 未标注 | 物流公司名称 |
| `detail` | `Detail` | `*string` | 未标注 | 物流详情 |
| `logistic_id` | `LogisticId` | `*string` | 未标注 | 物流id |
| `ship_area` | `ShipArea` | `*string` | 未标注 | 发货地 |
| `ship_period` | `ShipPeriod` | `*string` | 未标注 | 发货时效 |
| `status` | `Status` | `*string` | 未标注 | 物流状态 |
| `stop_update_time` | `StopUpdateTime` | `*string` | 未标注 | 物流停更时间 |

<a id="model-spuinfo"></a>

### SPUInfo

类型定义：[SPUInfo](order_model.go#L22)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand` | `Brand` | `*string` | 未标注 | 品牌名称 |
| `category` | `Category` | `*string` | 未标注 | 类目 |
| `count` | `Count` | `*int64` | 未标注 | 商品数量 |
| `icon` | `Icon` | `*string` | 未标注 | 图片链接 |
| `price` | `Price` | `*string` | 未标注 | 商品单价(单位:元) |
| `provider` | `Provider` | `*string` | 未标注 | 商品提供方，店铺或品牌方 |
| `spu_id` | `SpuId` | `*string` | 未标注 | 商品ID |
