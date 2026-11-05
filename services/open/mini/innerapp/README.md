# open/mini/innerapp 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/innerapp`。本目录提供 10 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `innerapp.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniInnerAppCreate`](#api-miniinnerappcreate) | [`alipay.open.mini.innerapp.create`](https://opendocs.alipay.com/pre-apis/05x9i7) | 内部链路小程序创建 | `POST` |
| [`MiniInnerAppPluginOrder`](#api-miniinnerapppluginorder) | [`alipay.open.mini.innerapp.plugin.order`](https://opendocs.alipay.com/pre-apis/05xng4) | 小程序代商户订购插件 | `POST` |
| [`MiniInnerAppPluginRelationQuery`](#api-miniinnerapppluginrelationquery) | [`alipay.open.mini.innerapp.pluginrelation.query`](https://opendocs.alipay.com/pre-apis/05xhsh) | 查询插件引用关系 | `POST` |
| [`MiniInnerAppPluginServicePublish`](#api-miniinnerapppluginservicepublish) | [`alipay.open.mini.innerapp.pluginservice.publish`](https://opendocs.alipay.com/pre-apis/05xe0v) | 插件发布能力中心 | `POST` |
| [`MiniInnerAppPluginSyncModeModify`](#api-miniinnerapppluginsyncmodemodify) | [`alipay.open.mini.innerapp.pluginsyncmode.modify`](https://opendocs.alipay.com/pre-apis/05xo5c) | 修改插件继承模式 | `POST` |
| [`MiniInnerAppPluginUseConfigCancel`](#api-miniinnerapppluginuseconfigcancel) | [`alipay.open.mini.innerapp.pluginuseconfig.cancel`](https://opendocs.alipay.com/pre-apis/05xgfv) | 撤销插件端版本配置灰度 | `POST` |
| [`MiniInnerAppPluginUseConfigOnline`](#api-miniinnerapppluginuseconfigonline) | [`alipay.open.mini.innerapp.pluginuseconfig.online`](https://opendocs.alipay.com/pre-apis/05xacn) | 上线插件使用端版本配置 | `POST` |
| [`MiniInnerAppPluginUseConfigUpgrade`](#api-miniinnerapppluginuseconfigupgrade) | [`alipay.open.mini.innerapp.pluginuseconfig.upgrade`](https://opendocs.alipay.com/pre-apis/05xo8d) | 灰度端版本配置 | `POST` |
| [`MiniInnerAppServicePublish`](#api-miniinnerappservicepublish) | [`alipay.open.mini.innerapp.service.publish`](https://opendocs.alipay.com/pre-apis/05xndu) | 小程序模板发布服务市场 | `POST` |
| [`MiniInnerAppServiceQuery`](#api-miniinnerappservicequery) | [`alipay.open.mini.innerapp.service.query`](https://opendocs.alipay.com/pre-apis/05xk3b) | 小程序服务市场搜索 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerappcreate"></a>

### MiniInnerAppCreate — 内部链路小程序创建

- 官方接口：[`alipay.open.mini.innerapp.create`](https://opendocs.alipay.com/pre-apis/05x9i7)
- 路由：`POST /v3/alipay/open/mini/innerapp/create`
- 请求类型：[MiniInnerAppCreateRequest](innerapp_request.go#L5)
- 响应类型：[MiniInnerAppCreateResponse](innerapp_response.go#L5)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | JSON body | 未标注 | 小程序类目，二级类目与三级类目下划线隔开，最多可以选择3个类目，类目之间以;隔开 |
| `app_desc` | `AppDesc` | `*string` | JSON body | 未标注 | 小程序描述，20-200个字符 |
| `app_english_name` | `AppEnglishName` | `*string` | JSON body | 未标注 | 小程序英文名称，插件不填 |
| `app_logo` | `AppLogo` | `*string` | JSON body | 未标注 | 小程序logo |
| `app_name` | `AppName` | `*string` | JSON body | 未标注 | 小程序名称 |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源的业务方,接入文档中申请的业务来源 |
| `app_slogan` | `AppSlogan` | `*string` | JSON body | 未标注 | 小程序简介，在搜索透出,插件不需要填 |
| `app_sub_type` | `AppSubType` | `*string` | JSON body | 未标注 | 应用子类型 |
| `app_type` | `AppType` | `*string` | JSON body | 未标注 | 应用类型 |
| `isv_app_id` | `IsvAppId` | `*string` | JSON body | 未标注 | 三方应用ID，仅在创建模板小程序时需要传入 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数 |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | JSON body | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list。使用后不再读取app_category_ids值，老前台类目将废弃 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 创建的小程序所属的PID |
| `service_email` | `ServiceEmail` | `*string` | JSON body | 未标注 | 邮箱，与客服电话二选一 |
| `service_phone` | `ServicePhone` | `*string` | JSON body | 未标注 | 客服电话，插件不需要 |

<a id="api-miniinnerapppluginorder"></a>

### MiniInnerAppPluginOrder — 小程序代商户订购插件

- 官方接口：[`alipay.open.mini.innerapp.plugin.order`](https://opendocs.alipay.com/pre-apis/05xng4)
- 路由：`POST /v3/alipay/open/mini/innerapp/plugin/order`
- 请求类型：[MiniInnerAppPluginOrderRequest](innerapp_request.go#L39)
- 响应类型：[MiniInnerAppPluginOrderResponse](innerapp_response.go#L11)

提供给二方场景下，代商家小程序订购、授权与签约插件

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `merchandise_id` | `MerchandiseId` | `*string` | JSON body | 未标注 | 订购的服务商品ID |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 一二方支持传入appId |
| `request_id` | `RequestId` | `*string` | JSON body | 未标注 | 业务幂等号，规则为商品id+appOrigin+requestId 作为幂等流水号 非必填,不传则默认生成一个8位requestId |

<a id="api-miniinnerapppluginrelationquery"></a>

### MiniInnerAppPluginRelationQuery — 查询插件引用关系

- 官方接口：[`alipay.open.mini.innerapp.pluginrelation.query`](https://opendocs.alipay.com/pre-apis/05xhsh)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginrelation/query`
- 请求类型：[MiniInnerAppPluginRelationQueryRequest](innerapp_request.go#L51)
- 响应类型：[MiniInnerAppPluginRelationQueryResponse](innerapp_response.go#L17)

查询插件引用关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id_list` | `MiniAppIdList` | `*[]string` | JSON body | 未标注 | 应用ID列表 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 页码 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 页大小 |
| `plugin_id` | `PluginId` | `*string` | JSON body | 未标注 | 插件ID |
| `plugin_relation_status_list` | `PluginRelationStatusList` | `*[]string` | JSON body | 未标注 | 插件使用关系状态，取值包括WORKING/WAIT_WORKING/STOP_WORKING/EXECUTING |
| `run_model_type` | `RunModelType` | `*string` | JSON body | 未标注 | 运行类型，取值包括ONLINE/TRIAL/REVIEW/DEBUG |
| `show_beta_info` | `ShowBetaInfo` | `*bool` | JSON body | 未标注 | 是否展示邀测信息 |

<a id="api-miniinnerapppluginservicepublish"></a>

### MiniInnerAppPluginServicePublish — 插件发布能力中心

- 官方接口：[`alipay.open.mini.innerapp.pluginservice.publish`](https://opendocs.alipay.com/pre-apis/05xe0v)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginservice/publish`
- 请求类型：[MiniInnerAppPluginServicePublishRequest](innerapp_request.go#L69)
- 响应类型：[MiniInnerAppPluginServicePublishResponse](innerapp_response.go#L29)

插件发布能力中心

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ability_type_list` | `AbilityTypeList` | `*[]string` | JSON body | 未标注 | 功能类型，目前支持的有：1039支付, 1040会员, 1041基础, 1056资金, 1058信用, 1111口碑, 330120安全, 360101营销 |
| `app_logo` | `AppLogo` | `*string` | JSON body | 未标注 | 服务发布logo |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `description` | `Description` | `*string` | JSON body | 未标注 | 服务描述 |
| `detail_for_client` | `DetailForClient` | `*string` | JSON body | 未标注 | 移动端详情，用于能力中心展示 |
| `detail_for_pc` | `DetailForPc` | `*string` | JSON body | 未标注 | pc端详情，用于能力中心展示 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 插件id |
| `sell_crowd` | `SellCrowd` | `*string` | JSON body | 未标注 | 可订购人群，1003个人, 1004企业，-1无限制 |
| `show_type` | `ShowType` | `*string` | JSON body | 未标注 | 发布后是否展示，01展示（默认）、02隐藏 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 服务标题 |
| `visit_url_for_pc` | `VisitUrlForPc` | `*string` | JSON body | 未标注 | pc端url地址，不需要则为空 |

<a id="api-miniinnerapppluginsyncmodemodify"></a>

### MiniInnerAppPluginSyncModeModify — 修改插件继承模式

- 官方接口：[`alipay.open.mini.innerapp.pluginsyncmode.modify`](https://opendocs.alipay.com/pre-apis/05xo5c)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginsyncmode/modify`
- 请求类型：[MiniInnerAppPluginSyncModeModifyRequest](innerapp_request.go#L95)
- 响应类型：[MiniInnerAppPluginSyncModeModifyResponse](innerapp_response.go#L35)

修改插件继承模式，模板和插件使用关系是否自动继承到宿主小程序上

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序应用 ID， |
| `plugin_id` | `PluginId` | `*string` | JSON body | 未标注 | 插件id |
| `sync_mode` | `SyncMode` | `*string` | JSON body | 未标注 | 模式，SYNC/NONE |

<a id="api-miniinnerapppluginuseconfigcancel"></a>

### MiniInnerAppPluginUseConfigCancel — 撤销插件端版本配置灰度

- 官方接口：[`alipay.open.mini.innerapp.pluginuseconfig.cancel`](https://opendocs.alipay.com/pre-apis/05xgfv)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginuseconfig/cancel`
- 请求类型：[MiniInnerAppPluginUseConfigCancelRequest](innerapp_request.go#L107)
- 响应类型：[MiniInnerAppPluginUseConfigCancelResponse](innerapp_response.go#L39)

结束插件端版本配置灰度

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端id |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序应用 ID |
| `plugin_dev_version` | `PluginDevVersion` | `*string` | JSON body | 未标注 | 插件研发版本 |
| `plugin_id` | `PluginId` | `*string` | JSON body | 未标注 | 插件id |

<a id="api-miniinnerapppluginuseconfigonline"></a>

### MiniInnerAppPluginUseConfigOnline — 上线插件使用端版本配置

- 官方接口：[`alipay.open.mini.innerapp.pluginuseconfig.online`](https://opendocs.alipay.com/pre-apis/05xacn)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginuseconfig/online`
- 请求类型：[MiniInnerAppPluginUseConfigOnlineRequest](innerapp_request.go#L121)
- 响应类型：[MiniInnerAppPluginUseConfigOnlineResponse](innerapp_response.go#L43)

小程序使用插件版本全量上线

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端id |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序应用 ID |
| `plugin_dev_version` | `PluginDevVersion` | `*string` | JSON body | 未标注 | 插件研发版本 |
| `plugin_id` | `PluginId` | `*string` | JSON body | 未标注 | 插件id |

<a id="api-miniinnerapppluginuseconfigupgrade"></a>

### MiniInnerAppPluginUseConfigUpgrade — 灰度端版本配置

- 官方接口：[`alipay.open.mini.innerapp.pluginuseconfig.upgrade`](https://opendocs.alipay.com/pre-apis/05xo8d)
- 路由：`POST /v3/alipay/open/mini/innerapp/pluginuseconfig/upgrade`
- 请求类型：[MiniInnerAppPluginUseConfigUpgradeRequest](innerapp_request.go#L135)
- 响应类型：[MiniInnerAppPluginUseConfigUpgradeResponse](innerapp_response.go#L47)

升级/降级 端版本配置

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端id |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序应用 ID |
| `plugin_dev_version` | `PluginDevVersion` | `*string` | JSON body | 未标注 | 插件研发版本 |
| `plugin_id` | `PluginId` | `*string` | JSON body | 未标注 | 插件id |
| `strategy_value` | `StrategyValue` | `*string` | JSON body | 未标注 | 灰度值 |

<a id="api-miniinnerappservicepublish"></a>

### MiniInnerAppServicePublish — 小程序模板发布服务市场

- 官方接口：[`alipay.open.mini.innerapp.service.publish`](https://opendocs.alipay.com/pre-apis/05xndu)
- 路由：`POST /v3/alipay/open/mini/innerapp/service/publish`
- 请求类型：[MiniInnerAppServicePublishRequest](innerapp_request.go#L151)
- 响应类型：[MiniInnerAppServicePublishResponse](innerapp_response.go#L51)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | `*string` | JSON body | 未标注 | 服务发布logo |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序模板所属PID |
| `sub_title` | `SubTitle` | `*string` | JSON body | 未标注 | 服务子标题 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 服务名称 |

<a id="api-miniinnerappservicequery"></a>

### MiniInnerAppServiceQuery — 小程序服务市场搜索

- 官方接口：[`alipay.open.mini.innerapp.service.query`](https://opendocs.alipay.com/pre-apis/05xk3b)
- 路由：`GET /v3/alipay/open/mini/innerapp/service/query`
- 请求类型：[MiniInnerAppServiceQueryRequest](innerapp_request.go#L167)
- 响应类型：[MiniInnerAppServiceQueryResponse](innerapp_response.go#L57)

小程序服务市场搜索，包括搜索模板和插件

请求约束：MiniInnerAppServiceQueryRequest is the request type for alipay.open.mini.innerapp.service.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 插件id或模板id |
| `app_sub_type` | `AppSubType` | `*string` | query | 否 | 小程序应用子类型 |
| `show_type` | `ShowType` | `*string` | query | 否 | 是否在服务市场透出，SHOW展示、HIDE隐藏 |
| `keyword` | `Keyword` | `*string` | query | 否 | 关键词 |
| `page_num` | `PageNum` | `*string` | query | 否 | 页码，默认第一页 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页数，默认每页10个 |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang |
| `include_offline` | `IncludeOffline` | `*string` | query | 否 | 是否需要展示offline状态插件，不传默认为false（不展示） |
