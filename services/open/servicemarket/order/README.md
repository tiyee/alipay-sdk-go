# order 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/servicemarket/order`。本包提供 8 个支付宝 V3 REST 操作。

通过 `order.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`ServiceMarketOrderAccept`](#api-servicemarketorderaccept) | `alipay.open.servicemarket.order.accept` |
| [`ServiceMarketOrderCreate`](#api-servicemarketordercreate) | `alipay.open.servicemarket.order.create` |
| [`ServiceMarketOrderItemCancel`](#api-servicemarketorderitemcancel) | `alipay.open.servicemarket.order.item.cancel` |
| [`ServiceMarketOrderItemComplete`](#api-servicemarketorderitemcomplete) | `alipay.open.servicemarket.order.item.complete` |
| [`ServiceMarketOrderItemConfirm`](#api-servicemarketorderitemconfirm) | `alipay.open.servicemarket.order.item.confirm` |
| [`ServiceMarketOrderNotify`](#api-servicemarketordernotify) | `alipay.open.servicemarket.order.notify` |
| [`ServiceMarketOrderQuery`](#api-servicemarketorderquery) | `alipay.open.servicemarket.order.query` |
| [`ServiceMarketOrderReject`](#api-servicemarketorderreject) | `alipay.open.servicemarket.order.reject` |

## 使用示例

- [ServiceMarketOrderAccept](api_order_servicemarketorderaccept_example_test.go)
- [ServiceMarketOrderCreate](api_order_servicemarketordercreate_example_test.go)
- [ServiceMarketOrderItemCancel](api_order_servicemarketorderitemcancel_example_test.go)
- [ServiceMarketOrderItemComplete](api_order_servicemarketorderitemcomplete_example_test.go)
- [ServiceMarketOrderItemConfirm](api_order_servicemarketorderitemconfirm_example_test.go)
- [ServiceMarketOrderNotify](api_order_servicemarketordernotify_example_test.go)
- [ServiceMarketOrderQuery](api_order_servicemarketorderquery_example_test.go)
- [ServiceMarketOrderReject](api_order_servicemarketorderreject_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-servicemarketorderaccept"></a>

### ServiceMarketOrderAccept — 服务商接单操作

- 官方接口：[`alipay.open.servicemarket.order.accept`](https://opendocs.alipay.com/open-v3/added906_alipay.open.servicemarket.order.accept)
- 路由：`POST /v3/alipay/open/servicemarket/order/accept`
- 请求类型：[ServiceMarketOrderAcceptRequest](order_request.go#L8)
- 响应类型：[ServiceMarketOrderAcceptResponse](order_response.go#L4)

商户订购服务插件之后，需要服务商接单，调用此接口可完成接单操作

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_order_id` | `CommodityOrderId` | `*string` | JSON body | 未标注 | 服务商品订单ID |

<a id="api-servicemarketordercreate"></a>

### ServiceMarketOrderCreate — 服务市场代订购小程序

- 官方接口：[`alipay.open.servicemarket.order.create`](https://opendocs.alipay.com/pre-apis/05xo6s)
- 路由：`POST /v3/alipay/open/servicemarket/order/create`
- 请求类型：[ServiceMarketOrderCreateRequest](order_request.go#L60)
- 响应类型：[ServiceMarketOrderCreateResponse](order_response.go#L8)

第三方服务商可以创建服务市场的订单，并代为创建小程序应用，小程序应用授权给服务商

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*ServiceMarketOrderCreate`](#model-servicemarketordercreate) | multipart | 未标注 | 见字段类型及官方接口文档。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-servicemarketorderitemcancel"></a>

### ServiceMarketOrderItemCancel — 服务订单明细实施项单项取消

- 官方接口：[`alipay.open.servicemarket.order.item.cancel`](https://opendocs.alipay.com/open-v3/97757082_alipay.open.servicemarket.order.item.cancel)
- 路由：`POST /v3/alipay/open/servicemarket/order/item/cancel`
- 请求类型：[ServiceMarketOrderItemCancelRequest](order_request.go#L14)
- 响应类型：[ServiceMarketOrderItemCancelResponse](order_response.go#L16)

服务商可以单个取消订单明细实施项

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cancel_reason` | `CancelReason` | `*string` | JSON body | 未标注 | 当前门店区域不支持实施 |
| `commodity_order_id` | `CommodityOrderId` | `*string` | JSON body | 未标注 | 订购服务订单ID，通过支付宝发送的 [服务市场商户确认订购通知](https://opendocs.alipay.com/apis/api_3/alipay.open.servicemarket.order.notify) 通知获取。 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一门店的ID。若该订单： 为支付宝服务市场发布的服务所产生订单则该参数可选。 |

<a id="api-servicemarketorderitemcomplete"></a>

### ServiceMarketOrderItemComplete — 服务商完成订单内单个明细实施项

- 官方接口：[`alipay.open.servicemarket.order.item.complete`](https://opendocs.alipay.com/open-v3/e31fe80d_alipay.open.servicemarket.order.item.complete)
- 路由：`POST /v3/alipay/open/servicemarket/order/item/complete`
- 请求类型：[ServiceMarketOrderItemCompleteRequest](order_request.go#L26)
- 响应类型：[ServiceMarketOrderItemCompleteResponse](order_response.go#L20)

商户在订购插件后，服务商完成实施操作，此接口会检测当前订单必须所属当前操作服务商

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_order_id` | `CommodityOrderId` | `*string` | JSON body | 未标注 | 订购服务插件订单号 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一门店的ID。若该订单： 为支付宝服务市场发布的服务所产生订单则该参数可选。 |

<a id="api-servicemarketorderitemconfirm"></a>

### ServiceMarketOrderItemConfirm — 服务商代商家确认实施完成

- 官方接口：[`alipay.open.servicemarket.order.item.confirm`](https://opendocs.alipay.com/open-v3/13c55349_alipay.open.servicemarket.order.item.confirm)
- 路由：`POST /v3/alipay/open/servicemarket/order/item/confirm`
- 请求类型：[ServiceMarketOrderItemConfirmRequest](order_request.go#L36)
- 响应类型：[ServiceMarketOrderItemConfirmResponse](order_response.go#L24)

此接口需商家授权服务商应用权限后，服务商可代商家进行实施确认完成动作

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_order_id` | `CommodityOrderId` | `*string` | JSON body | 未标注 | 订购服务插件订单号 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 商家订购服务选择的某一门店的ID。若该订单： 为支付宝(蚂蚁)服务市场发布的服务所产生订单则该参数可选。 |

<a id="api-servicemarketordernotify"></a>

### ServiceMarketOrderNotify — 服务市场商户确认订购通知

- 官方接口：[`alipay.open.servicemarket.order.notify`](https://opendocs.alipay.com/open-v3/2233a7ba_alipay.open.servicemarket.order.notify)
- 路由：`POST /v3/alipay/open/servicemarket/order/notify`
- 请求类型：[ServiceMarketOrderNotifyRequest](order_request.go#L54)
- 响应类型：[ServiceMarketOrderNotifyResponse](order_response.go#L28)

服务市场当商户选择服务商提供产品并订购确认时，通知服务商订单消息。服务商可以通过通知的消息内容回查该订单明细。回查接口（alipay.open.servicemarket.order.query）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |

<a id="api-servicemarketorderquery"></a>

### ServiceMarketOrderQuery — 服务市场订单明细查询

- 官方接口：[`alipay.open.servicemarket.order.query`](https://opendocs.alipay.com/open-v3/a59abd06_alipay.open.servicemarket.order.query)
- 路由：`GET /v3/alipay/open/servicemarket/order/query`
- 请求类型：[ServiceMarketOrderQueryRequest](order_request.go#L66)
- 响应类型：[ServiceMarketOrderQueryResponse](order_response.go#L32)

第三方服务商提供服务产品被商户订购后，服务市场会推送订单信息给服务商，服务商根据订单号回查该订单明细信息。

请求约束：ServiceMarketOrderQueryRequest is the request type for alipay.open.servicemarket.order.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_order_id` | `CommodityOrderId` | `*string` | query | 否 | 订购服务订单ID，通过支付宝发送的 [服务市场商户确认订购通知](https://opendocs.alipay.com/isv/b7f7864a_alipay.open.servicemarket.order.notify?scene=common&pathHash=af6ed436) 获取。 |
| `start_page` | `StartPage` | `*string` | query | 否 | 查询开始页码，1 为第一页。 |

<a id="api-servicemarketorderreject"></a>

### ServiceMarketOrderReject — 服务商拒绝接单

- 官方接口：[`alipay.open.servicemarket.order.reject`](https://opendocs.alipay.com/open-v3/f5d452ab_alipay.open.servicemarket.order.reject)
- 路由：`POST /v3/alipay/open/servicemarket/order/reject`
- 请求类型：[ServiceMarketOrderRejectRequest](order_request.go#L46)
- 响应类型：[ServiceMarketOrderRejectResponse](order_response.go#L50)

服务商可以拒接商户购买服务插件订单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `commodity_order_id` | `CommodityOrderId` | `*string` | JSON body | 未标注 | 订购服务商品订单ID |
| `reject_reason` | `RejectReason` | `*string` | JSON body | 未标注 | 拒绝接单原因 |

## 嵌套请求参数

<a id="model-servicemarketordercreate"></a>

### ServiceMarketOrderCreate

类型定义：[ServiceMarketOrderCreate](order_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | 未标注 | 11_12;12_13。小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考https://docs.alipay.com/isv/10325 |
| `app_desc` | `AppDesc` | `*string` | 未标注 | 小程序官方示例Demo，展示已支持的接口能力及组件。 |
| `app_english_name` | `AppEnglishName` | `*string` | 未标注 | 小程序应用英文名称 |
| `app_name` | `AppName` | `*string` | 未标注 | 小程序应用名称 |
| `app_origin` | `AppOrigin` | `*string` | 未标注 | 来源的业务方，需要申请 |
| `app_slogan` | `AppSlogan` | `*string` | 未标注 | 小程序应用简介，一句话描述小程序功能 |
| `market_code` | `MarketCode` | `*string` | 未标注 | 订购的服务商品ID所在的市场编码。新接入场景必须传递，具体值请联系产品分配。 |
| `merchandise_id` | `MerchandiseId` | `*string` | 未标注 | 订购的服务商品ID |
| `merchant_pid` | `MerchantPid` | `*string` | 未标注 | 商户PID |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 一二方支持传入appId |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。使用后不再读取app_category_ids值，老前台类目将废弃 |
| `out_biz_no` | `OutBizNo` | `*string` | 未标注 | 比如某种业务标准外部订单号,比如交易外部订单号，代表商户端自己订单号 |
| `service_email` | `ServiceEmail` | `*string` | 未标注 | 小程序客服邮箱 |
| `service_phone` | `ServicePhone` | `*string` | 未标注 | 小程序客服电话 |
