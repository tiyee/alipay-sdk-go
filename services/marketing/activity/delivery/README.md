# marketing/activity/delivery 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/delivery`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `delivery.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_delivery_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityDeliveryCreate`](#api-activitydeliverycreate) | [`alipay.marketing.activity.delivery.create`](https://opendocs.alipay.com/open-v3/43363a7a_alipay.marketing.activity.delivery.create) | 创建推广计划 | `POST` |
| [`ActivityDeliveryQuery`](#api-activitydeliveryquery) | [`alipay.marketing.activity.delivery.query`](https://opendocs.alipay.com/open-v3/cab7f6c3_alipay.marketing.activity.delivery.query) | 查询推广计划 | `POST` |
| [`ActivityDeliveryStop`](#api-activitydeliverystop) | [`alipay.marketing.activity.delivery.stop`](https://opendocs.alipay.com/open-v3/dbc77a2d_alipay.marketing.activity.delivery.stop) | 停止推广计划 | `PATCH` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activitydeliverycreate"></a>

### ActivityDeliveryCreate — 创建推广计划

- 官方接口：[`alipay.marketing.activity.delivery.create`](https://opendocs.alipay.com/open-v3/43363a7a_alipay.marketing.activity.delivery.create)
- 路由：`POST /v3/alipay/marketing/delivery`
- 请求类型：[ActivityDeliveryCreateRequest](delivery_request.go#L5)
- 响应类型：[ActivityDeliveryCreateResponse](delivery_response.go#L5)

为商家提供推广内容能力。创建推广成功后，请依据推广变更消息alipay.marketing.activity.delivery.message感知推广变更状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*DeliveryAgencyMerchantInfo`](#model-deliveryagencymerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_base_info` | `DeliveryBaseInfo` | [`*DeliveryBaseInfo`](#model-deliverybaseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_booth_code` | `DeliveryBoothCode` | `*string` | JSON body | 未标注 | 投放的展位编码。 |
| `delivery_config_list` | `DeliveryConfigList` | [`*[]DeliveryConfig`](#model-deliveryconfig) | JSON body | 未标注 | [已废弃] 待创建的投放配置列表。 最大数量限制20个。 |
| `delivery_play_config` | `DeliveryPlayConfig` | [`*DeliveryPlayConfig`](#model-deliveryplayconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_target_rule` | `DeliveryTargetRule` | [`*DeliveryTargetRule`](#model-deliverytargetrule) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `idem_no` | `IdemNo` | `*string` | JSON body | 未标注 | 商品卖货推广创建幂等号，用作幂等控制。 幂等作用：请求返回与上一次相同的结果。 外部接入方需保证幂等号唯一。 |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用：请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |

<a id="api-activitydeliveryquery"></a>

### ActivityDeliveryQuery — 查询推广计划

- 官方接口：[`alipay.marketing.activity.delivery.query`](https://opendocs.alipay.com/open-v3/cab7f6c3_alipay.marketing.activity.delivery.query)
- 路由：`POST /v3/alipay/marketing/delivery/{delivery_id}/query`
- 请求类型：[ActivityDeliveryQueryRequest](delivery_request.go#L23)
- 响应类型：[ActivityDeliveryQueryResponse](delivery_response.go#L19)

查询推广计划详情信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `belong_merchant_info` | `MarketingActivityDeliveryQuery.BelongMerchantInfo` | [`*DeliveryAgencyMerchantInfo`](#model-deliveryagencymerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_config_list` | `MarketingActivityDeliveryQuery.DeliveryConfigList` | [`*[]DeliveryConfig`](#model-deliveryconfig) | JSON body | 未标注 | "[已废弃] 待查询的投放配置列表。 最大数量限制20个。" |
| `merchant_access_mode` | `MarketingActivityDeliveryQuery.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `delivery_id` | `DeliveryId` | `string` | 路径 | 是 | 推广计划id |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityDeliveryQuery](delivery_model.go#L5) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activitydeliverystop"></a>

### ActivityDeliveryStop — 停止推广计划

- 官方接口：[`alipay.marketing.activity.delivery.stop`](https://opendocs.alipay.com/open-v3/dbc77a2d_alipay.marketing.activity.delivery.stop)
- 路由：`PATCH /v3/alipay/marketing/delivery/{delivery_id}/stop`
- 请求类型：[ActivityDeliveryStopRequest](delivery_request.go#L30)
- 响应类型：[ActivityDeliveryStopResponse](delivery_response.go#L36)

停止推广计划。
注意：停止推广操作成功到推广内容停止生效，可能存在几分钟的延迟。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `belong_merchant_info` | `MarketingActivityDeliveryStop.BelongMerchantInfo` | [`*DeliveryAgencyMerchantInfo`](#model-deliveryagencymerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_config_list` | `MarketingActivityDeliveryStop.DeliveryConfigList` | [`*[]DeliveryConfig`](#model-deliveryconfig) | JSON body | 未标注 | "[已废弃] 待停止的投放配置列表。 最大数量限制20个。" |
| `merchant_access_mode` | `MarketingActivityDeliveryStop.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityDeliveryStop.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `delivery_id` | `DeliveryId` | `string` | 路径 | 是 | 推广计划id |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityDeliveryStop](delivery_model.go#L14) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-deliveryactivitycontentinfo"></a>

### DeliveryActivityContentInfo

类型定义：[DeliveryActivityContentInfo](delivery_model.go#L25)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | 未标注 | 活动id |

<a id="model-deliveryagencymerchantinfo"></a>

### DeliveryAgencyMerchantInfo

类型定义：[DeliveryAgencyMerchantInfo](delivery_model.go#L31)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_type` | `BusinessType` | `*string` | 未标注 | 合作业务类型。 |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户信息 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。 限制: 服务商代运营模式必传 如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。 因此要求该SMID必须绑定相应的支付宝PID。 |

<a id="model-deliveryavailablecitycode"></a>

### DeliveryAvailableCityCode

类型定义：[DeliveryAvailableCityCode](delivery_model.go#L41)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `all_city` | `AllCity` | `*bool` | 未标注 | 是否全国。 与city_codes二选一。只允许填true，否则不填。 |
| `city_codes` | `CityCodes` | `*[]string` | 未标注 | 城市编码。与all_city二选一。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/ ） |

<a id="model-deliveryavailablescope"></a>

### DeliveryAvailableScope

类型定义：[DeliveryAvailableScope](delivery_model.go#L49)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_available_city_code` | `DeliveryAvailableCityCode` | [`*DeliveryAvailableCityCode`](#model-deliveryavailablecitycode) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_available_scope_type` | `DeliveryAvailableScopeType` | `*string` | 未标注 | 投放可用范围类型，目前支持的类型是城市范围CITY_CODE |

<a id="model-deliverybaseinfo"></a>

### DeliveryBaseInfo

类型定义：[DeliveryBaseInfo](delivery_model.go#L56)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_begin_time` | `DeliveryBeginTime` | `*string` | 未标注 | 投放计划开始时间。 格式为：yyyy-MM-dd HH:mm:ss。 |
| `delivery_end_time` | `DeliveryEndTime` | `*string` | 未标注 | 投放计划结束时间。 格式为：yyyy-MM-dd HH:mm:ss。 |
| `delivery_material` | `DeliveryMaterial` | [`*DeliveryMaterial`](#model-deliverymaterial) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_name` | `DeliveryName` | `*string` | 未标注 | 投放计划名称。 投放计划名称不会对用户进行表达，只用于商户管理使用。长度需要大于等于3，小于20。 |

<a id="model-deliverycitycoderule"></a>

### DeliveryCityCodeRule

类型定义：[DeliveryCityCodeRule](delivery_model.go#L67)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `all_city` | `AllCity` | `*bool` | 未标注 | 是否全国。只允许填true，否则不填。 |
| `city_codes` | `CityCodes` | `*[]string` | 未标注 | 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/ ） |

<a id="model-deliveryconfig"></a>

### DeliveryConfig

类型定义：[DeliveryConfig](delivery_model.go#L75)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | 未标注 | 活动id |
| `booth_code` | `BoothCode` | `*string` | 未标注 | 展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAYMENT_RESULT |
| `channel` | `Channel` | `*string` | 未标注 | 可投放的渠道标识. |

<a id="model-deliverycontentconfig"></a>

### DeliveryContentConfig

类型定义：[DeliveryContentConfig](delivery_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_send_guide` | `DeliverySendGuide` | [`*DeliverySendGuide`](#model-deliverysendguide) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deliverycontentinfo"></a>

### DeliveryContentInfo

类型定义：[DeliveryContentInfo](delivery_model.go#L90)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_activity_content` | `DeliveryActivityContent` | [`*DeliveryActivityContentInfo`](#model-deliveryactivitycontentinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_app_content` | `DeliveryAppContent` | [`*DeliveryMiniAppContentInfo`](#model-deliveryminiappcontentinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_content_type` | `DeliveryContentType` | `*string` | 未标注 | 运营内容类型。 |
| `delivery_display_info` | `DeliveryDisplayInfo` | [`*DeliveryDisplayInfo`](#model-deliverydisplayinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_item_content` | `DeliveryItemContent` | [`*DeliveryItemContentInfo`](#model-deliveryitemcontentinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deliverycrowdrule"></a>

### DeliveryCrowdRule

类型定义：[DeliveryCrowdRule](delivery_model.go#L100)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `crowd_id` | `CrowdId` | `*string` | 未标注 | 指定人群ID |

<a id="model-deliverydisplayinfo"></a>

### DeliveryDisplayInfo

类型定义：[DeliveryDisplayInfo](delivery_model.go#L106)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `main_title` | `MainTitle` | `*string` | 未标注 | 副标题。 |
| `sub_title` | `SubTitle` | `*string` | 未标注 | 副标题 |

<a id="model-deliveryfullsendconfig"></a>

### DeliveryFullSendConfig

类型定义：[DeliveryFullSendConfig](delivery_model.go#L114)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_content_info` | `DeliveryContentInfo` | [`*DeliveryContentInfo`](#model-deliverycontentinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_floor_amount` | `DeliveryFloorAmount` | `*string` | 未标注 | 满足消费金额门槛，左闭区间，单位元。 说明：限制支付时的订单金额最少满足的金额门槛。 限制：value &gt; 0 && value &lt;= 99999。 |

<a id="model-deliveryitemcontentinfo"></a>

### DeliveryItemContentInfo

类型定义：[DeliveryItemContentInfo](delivery_model.go#L121)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | 未标注 | 支付宝侧商品id。 注意事项：支持状态为“出售中”的商品 |
| `item_type` | `ItemType` | `*string` | 未标注 | 商品模版类型 |

<a id="model-deliverymaterial"></a>

### DeliveryMaterial

类型定义：[DeliveryMaterial](delivery_model.go#L129)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_single_material` | `DeliverySingleMaterial` | [`*DeliverySingleMaterial`](#model-deliverysinglematerial) | 未标注 | 见字段类型及官方接口文档。 |
| `material_id_list` | `MaterialIdList` | `*[]string` | 未标注 | 素材ID，通过alipay.marketing.material.create接口提供的素材组 |

<a id="model-deliverymerchantinfo"></a>

### DeliveryMerchantInfo

类型定义：[DeliveryMerchantInfo](delivery_model.go#L136)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户id |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。 |

<a id="model-deliverymerchantrule"></a>

### DeliveryMerchantRule

类型定义：[DeliveryMerchantRule](delivery_model.go#L144)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_id_list` | `BrandIdList` | `*[]string` | 未标注 | 指定品牌id。 说明：如商户需选择某个品牌下维护的收款账号，请上传相关品牌id |
| `delivery_merchant_infos` | `DeliveryMerchantInfos` | [`*[]DeliveryMerchantInfo`](#model-deliverymerchantinfo) | 未标注 | 曝光商户选取列表。 说明：需要传入您期望曝光的商户的商户号，传入为空时默认使用投放优惠券活动的适用范围。 限制：曝光商户号需与投放归属商户号相同，或传入有跨主体授权关系的商户号或有弱绑定关系的M3账号 |
| `delivery_merchant_mode` | `DeliveryMerchantMode` | `*string` | 未标注 | 指定支付成功页模式。 |

<a id="model-deliveryminiappcontentinfo"></a>

### DeliveryMiniAppContentInfo

类型定义：[DeliveryMiniAppContentInfo](delivery_model.go#L154)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_delivery_type` | `MiniAppDeliveryType` | `*string` | 未标注 | 仅小程序推广可用。 |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 小程序ID |
| `service_code_list` | `ServiceCodeList` | `*[]string` | 未标注 | 小程序服务编码/home_page，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 限制: 1.只有mini_app_id有值时该值传入才会有效 2.当前只支持一个服务code 3.service_codes必须是已传入的mini_app_id下的小程序 |

<a id="model-deliveryplayconfig"></a>

### DeliveryPlayConfig

类型定义：[DeliveryPlayConfig](delivery_model.go#L164)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_full_send_config` | `DeliveryFullSendConfig` | [`*DeliveryFullSendConfig`](#model-deliveryfullsendconfig) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_single_send_config` | `DeliverySingleSendConfig` | [`*DeliverySingleSendConfig`](#model-deliverysinglesendconfig) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deliverysearchboxrule"></a>

### DeliverySearchBoxRule

类型定义：[DeliverySearchBoxRule](delivery_model.go#L170)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `box_id` | `BoxId` | `*string` | 未标注 | 指定搜索直达区域 |

<a id="model-deliverysendguide"></a>

### DeliverySendGuide

类型定义：[DeliverySendGuide](delivery_model.go#L176)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_guide_url` | `DeliveryGuideUrl` | `*string` | 未标注 | 引导链接，从支付宝公域跳转到服务商(商户)承接页的引导链接。 |

<a id="model-deliverysinglematerial"></a>

### DeliverySingleMaterial

类型定义：[DeliverySingleMaterial](delivery_model.go#L182)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_image` | `DeliveryImage` | `*string` | 未标注 | 投放计划图片素材。通过接口alipay.marketing.material.image.upload上传图片返回的resource_id。 |

<a id="model-deliverysinglesendconfig"></a>

### DeliverySingleSendConfig

类型定义：[DeliverySingleSendConfig](delivery_model.go#L188)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_content_config` | `DeliveryContentConfig` | [`*DeliveryContentConfig`](#model-deliverycontentconfig) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_content_info` | `DeliveryContentInfo` | [`*DeliveryContentInfo`](#model-deliverycontentinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deliverytargetrule"></a>

### DeliveryTargetRule

类型定义：[DeliveryTargetRule](delivery_model.go#L194)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_id_list` | `BrandIdList` | `*[]string` | 未标注 | 指定品牌id。 说明：如商户需选择某个品牌下维护的收款账号，请上传相关品牌id |
| `delivery_available_scope` | `DeliveryAvailableScope` | [`*DeliveryAvailableScope`](#model-deliveryavailablescope) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_city_code_rule` | `DeliveryCityCodeRule` | [`*DeliveryCityCodeRule`](#model-deliverycitycoderule) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_crowd_rule` | `DeliveryCrowdRule` | [`*DeliveryCrowdRule`](#model-deliverycrowdrule) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_merchant_infos` | `DeliveryMerchantInfos` | [`*[]DeliveryMerchantInfo`](#model-deliverymerchantinfo) | 未标注 | "曝光商户选取列表。 说明：需要传入您期望曝光的商户的商户号，传入为空时默认使用投放优惠券活动的适用范围。 限制：曝光商户号需与投放归属商户号相同，或传入有跨主体授权关系的商户号或有弱绑定关系的M3账号" |
| `delivery_merchant_mode` | `DeliveryMerchantMode` | `*string` | 未标注 | 指定支付成功页模式。 |
| `delivery_merchant_rule` | `DeliveryMerchantRule` | [`*DeliveryMerchantRule`](#model-deliverymerchantrule) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_promo_tags` | `DeliveryPromoTags` | `*string` | 未标注 | 推荐定向曝光订单优惠标记。 1. 创建交易订单时，上传符合的订单优惠标。在业务参数信息“business_params”的子字段“delivery_promo_tags”填充订单优惠标。“统一收单交易支付接口”详情查看： https://opendocs.alipay.com/open/02ekfp?scene=32 2. 订单优惠标，检验规则： * 单个优惠标最大长度32字节。多个标记值以英文逗号分隔。 * 单个优惠标仅支持字母/数字/下划线。 3. 订单优惠标规则和其他推荐规则同时叠加。 |
| `delivery_recall_mode` | `DeliveryRecallMode` | `*string` | 未标注 | 投放召回模式，现在支持CITY_RECALL（城市召回)和SHOP_RECALL（门店召回)两种模式。 |
| `delivery_search_box_rule` | `DeliverySearchBoxRule` | [`*DeliverySearchBoxRule`](#model-deliverysearchboxrule) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_template_rule` | `DeliveryTemplateRule` | [`*DeliveryTemplateRule`](#model-deliverytemplaterule) | 未标注 | 见字段类型及官方接口文档。 |
| `delivery_type` | `DeliveryType` | `*string` | 未标注 | 小程序推广可用。 |

<a id="model-deliverytemplaterule"></a>

### DeliveryTemplateRule

类型定义：[DeliveryTemplateRule](delivery_model.go#L216)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | 未标注 | 指定商家消息区域 |
