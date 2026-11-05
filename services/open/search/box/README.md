# box 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/search/box`。本包提供 9 个支付宝 V3 REST 操作。

通过 `box.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SearchBoxApply`](#api-searchboxapply) | `alipay.open.search.box.apply` |
| [`SearchBoxBatchQuery`](#api-searchboxbatchquery) | `alipay.open.search.box.batchquery` |
| [`SearchBoxConsult`](#api-searchboxconsult) | `alipay.open.search.box.consult` |
| [`SearchBoxModify`](#api-searchboxmodify) | `alipay.open.search.box.modify` |
| [`SearchBoxOffline`](#api-searchboxoffline) | `alipay.open.search.box.offline` |
| [`SearchBoxOnline`](#api-searchboxonline) | `alipay.open.search.box.online` |
| [`SearchBoxQuery`](#api-searchboxquery) | `alipay.open.search.box.query` |
| [`SearchBoxActivityApply`](#api-searchboxactivityapply) | `alipay.open.search.boxactivity.apply` |
| [`SearchBoxActivityQuery`](#api-searchboxactivityquery) | `alipay.open.search.boxactivity.query` |

## 使用示例

- [SearchBoxApply](api_box_searchboxapply_example_test.go)
- [SearchBoxBatchQuery](api_box_searchboxbatchquery_example_test.go)
- [SearchBoxConsult](api_box_searchboxconsult_example_test.go)
- [SearchBoxModify](api_box_searchboxmodify_example_test.go)
- [SearchBoxOffline](api_box_searchboxoffline_example_test.go)
- [SearchBoxOnline](api_box_searchboxonline_example_test.go)
- [SearchBoxQuery](api_box_searchboxquery_example_test.go)
- [SearchBoxActivityApply](api_box_searchboxactivityapply_example_test.go)
- [SearchBoxActivityQuery](api_box_searchboxactivityquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-searchboxapply"></a>

### SearchBoxApply — 申请创建搜索直达配置

- 官方接口：[`alipay.open.search.box.apply`](https://opendocs.alipay.com/open-v3/622c6b11_alipay.open.search.box.apply?scene=common&pathHash=d1a9edf5)
- 路由：`POST /v3/alipay/open/search/box/apply`
- 请求类型：[SearchBoxApplyRequest](box_request.go#L4)
- 响应类型：[SearchBoxApplyResponse](box_response.go#L4)

申请创建搜索直达配置，相关概念说明及准入条件详见 [ 搜索直达 ](https://opendocs.alipay.com/mini/03fjba)
通过该接口创建搜索直达，满足准入门槛时，创建后立马生效。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_desc` | `BoxDesc` | `*string` | JSON body | 未标注 | 品牌介绍，5-15个中文字符。 小程序直达不支持设置此项 |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id，参考[ 品牌认证说明 ](https://opendocs.alipay.com/rules/029uy4) |
| `business_benefit_switch` | `BusinessBenefitSwitch` | `*bool` | JSON body | 未标注 | 可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块 |
| `business_district_ids` | `BusinessDistrictIds` | `*[]string` | JSON body | 未标注 | 小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放） |
| `custom_keywords` | `CustomKeywords` | `*[]string` | JSON body | 未标注 | 自定义触发词，最多可配置10个，限1-8个中文字符。 小程序直达不支持设置此项 |
| `image_id` | `ImageId` | `*string` | JSON body | 未标注 | 氛围图片id，调用 [支付宝文件上传接口](https://opendocs.alipay.com/mini/03hvl1?ref=api) 上传图片获取图片id(bizCode：search_box_atmosphere)。 [图片规范](https://opendocs.alipay.com/b/03al6f) 小程序直达不支持设置此项。 |
| `image_name` | `ImageName` | `*string` | JSON body | 未标注 | 氛围图片名。 小程序直达不支持设置此项 |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |
| `related_accounts` | `RelatedAccounts` | [`*[]SearchBoxAppInfo`](#model-searchboxappinfo) | JSON body | 未标注 | 关联账号信息，1-3个。 内部字段均需设置。当为品牌直达时，数组中的第1个账号会被设置为"账号1"，也就是将作为搜索直达专区头部的跳转地址 |
| `service_infos` | `ServiceInfos` | [`*[]SearchBoxServiceInfo`](#model-searchboxserviceinfo) | JSON body | 未标注 | "服务信息，服务必须审核通过才能申请搜索直达。品牌直达最多可配置同一品牌认证下的小程序4个，小程序直达最多可配置2个。 内部字段均需设置。" |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合[ 准入类目 ](https://opendocs.alipay.com/b/03al6e) |

<a id="api-searchboxbatchquery"></a>

### SearchBoxBatchQuery — 批量查询搜索直达配置列表

- 官方接口：[`alipay.open.search.box.batchquery`](https://opendocs.alipay.com/open-v3/fd1480ed_alipay.open.search.box.batchquery?scene=common&pathHash=4ca4364b)
- 路由：`POST /v3/alipay/open/search/box/batchquery`
- 请求类型：[SearchBoxBatchQueryRequest](box_request.go#L30)
- 响应类型：[SearchBoxBatchQueryResponse](box_response.go#L10)

批量查询商户已开通搜索直达的小程序及直达信息（搜索直达名称和id）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |
| `page_number` | `PageNumber` | `*int64` | JSON body | 未标注 | 分页查询的当前页号,从1开始 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页查询的数量，默认10，不得超过50 |

<a id="api-searchboxconsult"></a>

### SearchBoxConsult — 搜索直达创建预校验

- 官方接口：[`alipay.open.search.box.consult`](https://opendocs.alipay.com/open-v3/28d19e54_alipay.open.search.box.consult)
- 路由：`POST /v3/alipay/open/search/box/consult`
- 请求类型：[SearchBoxConsultRequest](box_request.go#L40)
- 响应类型：[SearchBoxConsultResponse](box_response.go#L24)

用于对搜索直达创建申请的入参和条件进行提前判断校验。
具体准入条件参考[搜索直达准入条件](https://opendocs.alipay.com/pre-open/0311q3#%E5%87%86%E5%85%A5%E6%9D%A1%E4%BB%B6)

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `area_keywords` | `AreaKeywords` | `*string` | JSON body | 未标注 | 小程序直达配置的常用服务中带有门店信息时，可添加简称触发词 |
| `box_desc` | `BoxDesc` | `*string` | JSON body | 未标注 | 品牌介绍，5-15个中文字符。 小程序直达不支持设置此项 |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id，参考[品牌认证说明](https://opendocs.alipay.com/rules/029uy4) |
| `business_benefit_switch` | `BusinessBenefitSwitch` | `*bool` | JSON body | 未标注 | 可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块 |
| `business_district_ids` | `BusinessDistrictIds` | `*[]string` | JSON body | 未标注 | 小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放） |
| `custom_keywords` | `CustomKeywords` | `*[]string` | JSON body | 未标注 | 自定义触发词，最多可配置6个，限1-8个中文字符。 小程序直达不支持设置此项 |
| `image_id` | `ImageId` | `*string` | JSON body | 未标注 | 氛围图片id，调用 [支付宝文件上传接口](https://opendocs.alipay.com/mini/03hvl1?ref=api) 上传图片获取图片id(bizCode：search_box_atmosphere)。 [图片规范](https://opendocs.alipay.com/b/03al6f) 小程序直达不支持设置此项。 |
| `image_name` | `ImageName` | `*string` | JSON body | 未标注 | 氛围图片名。 小程序直达不支持设置此项 |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |
| `related_accounts` | `RelatedAccounts` | [`*[]SearchBoxAppInfo`](#model-searchboxappinfo) | JSON body | 未标注 | 关联账号信息，1-2个。 内部字段均需设置。 |
| `service_infos` | `ServiceInfos` | [`*[]SearchBoxServiceInfo`](#model-searchboxserviceinfo) | JSON body | 未标注 | 服务信息，服务必须审核通过才能申请搜索直达。 内部字段均需设置。 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合[准入类目](https://opendocs.alipay.com/b/03al6e) |

<a id="api-searchboxmodify"></a>

### SearchBoxModify — 修改搜索直达配置

- 官方接口：[`alipay.open.search.box.modify`](https://opendocs.alipay.com/open-v3/28bd3996_alipay.open.search.box.modify?scene=common&pathHash=f07276ab)
- 路由：`POST /v3/alipay/open/search/box/modify`
- 请求类型：[SearchBoxModifyRequest](box_request.go#L68)
- 响应类型：[SearchBoxModifyResponse](box_response.go#L32)

修改搜索直达配置模块，配置对应的模块参考：[ 搜索直达模块说明 ](https://opendocs.alipay.com/mini/03fjba#%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E)

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_desc` | `BoxDesc` | `*string` | JSON body | 未标注 | 品牌介绍，5-15个中文字符。当修改品牌介绍模块(module_type=BOX_EXCLUSIVE_BASE)时传入。 小程序直达时不支持设置 |
| `box_id` | `BoxId` | `*string` | JSON body | 未标注 | 搜索直达配置id |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id，参考[ 品牌认证说明 ](https://opendocs.alipay.com/rules/029uy4) |
| `business_benefit_switch` | `BusinessBenefitSwitch` | `*bool` | JSON body | 未标注 | 可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块 |
| `business_district_ids` | `BusinessDistrictIds` | `*[]string` | JSON body | 未标注 | 小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放） |
| `custom_keywords` | `CustomKeywords` | `*string` | JSON body | 未标注 | 自定义触发词 |
| `image_id` | `ImageId` | `*string` | JSON body | 未标注 | 氛围图片id，调用 [支付宝文件上传接口](https://opendocs.alipay.com/mini/03hvl1?ref=api) 上传图片获取图片id(bizCode：search_box_atmosphere)。 [图片规范](https://opendocs.alipay.com/b/03al6f) 小程序直达不支持设置此项。 |
| `image_name` | `ImageName` | `*string` | JSON body | 未标注 | 氛围图片名，当修改氛围图模块(module_type=BOX_ATMOSPHERE_IMAGE)时传入。 小程序直达不支持设置 |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |
| `module_id` | `ModuleId` | `*string` | JSON body | 未标注 | 搜索直达模块配置id，新增模块时可不传，更新模块时必传。从alipay.open.search.box.query详情接口查询可获取已存在模块的配置id |
| `module_type` | `ModuleType` | `*string` | JSON body | 未标注 | 搜索直达模块类型，参考[ 搜索直达模块说明 ](https://opendocs.alipay.com/mini/03fjba#%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E)。 |
| `related_accounts` | `RelatedAccounts` | [`*[]SearchBoxAppInfo`](#model-searchboxappinfo) | JSON body | 未标注 | "关联账号信息，可配置1-2个。传入账号需归属于商家主体。品牌直达修改官方账号(module_type=BOX_EXCLUSIVE_ACCOUNTS)时传入，小程序直达修改常用服务模块时也需传值。 当为品牌直达时，数组中的第1个账号会被设置为"账号1"，也就是将作为搜索直达专区头部的跳转地址。" |
| `service_infos` | `ServiceInfos` | [`*[]SearchBoxServiceInfo`](#model-searchboxserviceinfo) | JSON body | 未标注 | 服务信息，服务必须审核通过才能申请搜索直达，可配置1-4个。当修改常用服务模块(module_type=BOX_EXCLUSIVE_FUNCTIONS)时传入 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合[ 准入类目 ](https://opendocs.alipay.com/b/03al6e) |

<a id="api-searchboxoffline"></a>

### SearchBoxOffline — 下架搜索直达

- 官方接口：[`alipay.open.search.box.offline`](https://opendocs.alipay.com/open-v3/4d9cc1dc_alipay.open.search.box.offline)
- 路由：`POST /v3/alipay/open/search/box/offline`
- 请求类型：[SearchBoxOfflineRequest](box_request.go#L100)
- 响应类型：[SearchBoxOfflineResponse](box_response.go#L38)

若不需要运营搜索直达时，可调用次接口下架搜索直达

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_id` | `BoxId` | `*string` | JSON body | 未标注 | 搜索直达配置id |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |

<a id="api-searchboxonline"></a>

### SearchBoxOnline — 上架搜索直达

- 官方接口：[`alipay.open.search.box.online`](https://opendocs.alipay.com/open-v3/00d18795_alipay.open.search.box.online)
- 路由：`POST /v3/alipay/open/search/box/online`
- 请求类型：[SearchBoxOnlineRequest](box_request.go#L108)
- 响应类型：[SearchBoxOnlineResponse](box_response.go#L42)

若搜索直达处于下架状态时，可通过此接口来重新上架直达，上架须满足搜索直达准入门槛，若不满足，则会上架失败并报错

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_id` | `BoxId` | `*string` | JSON body | 未标注 | 搜索直达配置id |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |

<a id="api-searchboxquery"></a>

### SearchBoxQuery — 查询搜索直达配置详情

- 官方接口：[`alipay.open.search.box.query`](https://opendocs.alipay.com/open-v3/97b73d37_alipay.open.search.box.query?scene=common&pathHash=6f2e052c)
- 路由：`GET /v3/alipay/open/search/box/query`
- 请求类型：[SearchBoxQueryRequest](box_request.go#L146)
- 响应类型：[SearchBoxQueryResponse](box_response.go#L46)

查询商户已开通搜索直达的小程序及直达信息（搜索直达名称和id），各模块说明详见 [ 搜索直达 ](https://opendocs.alipay.com/mini/03fjba)。

请求约束：SearchBoxQueryRequest is the request type for alipay.open.search.box.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_id` | `BoxId` | `*string` | query | 否 | 搜索直达配置id |
| `merchant_id` | `MerchantId` | `*string` | query | 否 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |

<a id="api-searchboxactivityapply"></a>

### SearchBoxActivityApply — 申请创建搜索直达活动配置

- 官方接口：[`alipay.open.search.boxactivity.apply`](https://opendocs.alipay.com/open-v3/33107e1b_alipay.open.search.boxactivity.apply)
- 路由：`POST /v3/alipay/open/search/boxactivity/apply`
- 请求类型：[SearchBoxActivityApplyRequest](box_request.go#L116)
- 响应类型：[SearchBoxActivityApplyResponse](box_response.go#L63)

当搜索直达处于上架状态时，通过该接口来创建banner活动。相关概念和说明详见 [ 搜索直达 ](https://opendocs.alipay.com/mini/03fjba)。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `action_url` | `ActionUrl` | `*string` | JSON body | 未标注 | 活动链接，必须为关联的小程序的页面链接（链接以 alipays 开头），可参考[ 小程序scheme链接介绍 ](https://opendocs.alipay.com/support/01rb18) |
| `background_word` | `BackgroundWord` | `*string` | JSON body | 未标注 | 创建banner活动时，可添加底纹推广词，有机会在首页搜索框内展示，提升banner活动曝光机会 |
| `box_id` | `BoxId` | `*string` | JSON body | 未标注 | 搜索直达配置id |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 活动结束时间，开始和结束时间之间不得超过90天 |
| `material_id` | `MaterialId` | `*string` | JSON body | 未标注 | 当material_type="IMAGE"时，为图片id；当material_type="VIDEO"时，为视频id。小程序直达仅支持图片类型。图片id可以通过调用[ 支付宝文件上传接口 ](https://opendocs.alipay.com/mini/03hvl1?ref=api)上传图片，获取图片id(bizCode：search_box_banner)。[ 图片规范 ](https://opendocs.alipay.com/b/03al6f) |
| `material_type` | `MaterialType` | `*string` | JSON body | 未标注 | IMAGE-图片/VIDEO-视频 |
| `merchant_id` | `MerchantId` | `*string` | JSON body | 未标注 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |
| `oriented_rule` | `OrientedRule` | [`*DeliveryOrientedRuleInfo`](#model-deliveryorientedruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 活动开始时间，开始和结束时间之间不得超过90天 |
| `target_appid` | `TargetAppid` | `*string` | JSON body | 未标注 | 跳转应用ID |
| `target_appname` | `TargetAppname` | `*string` | JSON body | 未标注 | 跳转小程序名称 |
| `target_regions` | `TargetRegions` | [`*[]DeliveryTargetRegion`](#model-deliverytargetregion) | JSON body | 未标注 | 投放目标区域列表 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 活动标题 |
| `video_info` | `VideoInfo` | [`*SearchBoxActivityVideoInfo`](#model-searchboxactivityvideoinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-searchboxactivityquery"></a>

### SearchBoxActivityQuery — 查询搜索直达活动配置详情

- 官方接口：[`alipay.open.search.boxactivity.query`](https://opendocs.alipay.com/open-v3/5b195961_alipay.open.search.boxactivity.query)
- 路由：`GET /v3/alipay/open/search/boxactivity/query`
- 请求类型：[SearchBoxActivityQueryRequest](box_request.go#L154)
- 响应类型：[SearchBoxActivityQueryResponse](box_response.go#L69)

查询单个搜索直达banner的活动信息

请求约束：SearchBoxActivityQueryRequest is the request type for alipay.open.search.boxactivity.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `box_activity_id` | `BoxActivityId` | `*string` | query | 否 | 搜索直达活动id |
| `merchant_id` | `MerchantId` | `*string` | query | 否 | 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。 |

## 嵌套请求参数

<a id="model-deliveryorientedruleinfo"></a>

### DeliveryOrientedRuleInfo

类型定义：[DeliveryOrientedRuleInfo](box_model.go#L38)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `client_platform` | `ClientPlatform` | `*string` | 未标注 | 客户端 |

<a id="model-deliverytargetregion"></a>

### DeliveryTargetRegion

类型定义：[DeliveryTargetRegion](box_model.go#L44)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `region_code` | `RegionCode` | `*string` | 未标注 | 区域编码，参考[取值说明](https://opendocs.alipay.com/pre-open/03144c) |
| `region_name` | `RegionName` | `*string` | 未标注 | 区域名，参考[取值说明](https://opendocs.alipay.com/pre-open/03144c) |
| `region_type` | `RegionType` | `*string` | 未标注 | 区域类型，1代表全国，2代表省级区域，3代表市级区域，参考[取值说明](https://opendocs.alipay.com/pre-open/03144c) |

<a id="model-searchboxactivityvideoinfo"></a>

### SearchBoxActivityVideoInfo

类型定义：[SearchBoxActivityVideoInfo](box_model.go#L64)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `video_btn_text` | `VideoBtnText` | `*string` | 未标注 | 视频按钮文案，2-4个汉字 |
| `video_coverimg_id` | `VideoCoverimgId` | `*string` | 未标注 | 视频封面图片id，图片id可以通过调用接口alipay.open.file.upload上传图片，获取图片id。图片规范：https://opendocs.alipay.com/mini/operation/atmospheredesign |
| `video_sub_title` | `VideoSubTitle` | `*string` | 未标注 | 视频副标题，2-16个汉字 |
| `video_title` | `VideoTitle` | `*string` | 未标注 | 视频主标题，2-10个汉字 |

<a id="model-searchboxappinfo"></a>

### SearchBoxAppInfo

类型定义：[SearchBoxAppInfo](box_model.go#L76)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_name` | `AppName` | `*string` | 未标注 | 应用名称 |
| `app_type` | `AppType` | `*string` | 未标注 | 应用类型，TINY_APP-小程序/PUBLIC_APP-生活号 |
| `relate_appid` | `RelateAppid` | `*string` | 未标注 | 应用ID |

<a id="model-searchboxserviceinfo"></a>

### SearchBoxServiceInfo

类型定义：[SearchBoxServiceInfo](box_model.go#L156)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `service_category_code` | `ServiceCategoryCode` | `*string` | 未标注 | 服务类目编码 |
| `service_category_name` | `ServiceCategoryName` | `*string` | 未标注 | 服务类目名称 |
| `service_code` | `ServiceCode` | `*string` | 未标注 | 服务编码 |
| `service_name` | `ServiceName` | `*string` | 未标注 | 服务名称 |
