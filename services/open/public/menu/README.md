# open/public/menu 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/menu`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `menu.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicMenuBatchQuery`](#api-publicmenubatchquery) | [`alipay.open.public.menu.batchquery`](https://opendocs.alipay.com/apis/05wol5) | 菜单列表查询接口 | `POST` |
| [`PublicMenuCreate`](#api-publicmenucreate) | [`alipay.open.public.menu.create`](https://opendocs.alipay.com/apis/05wvxt) | 默认菜单创建接口 | `POST` |
| [`PublicMenuDataBatchQuery`](#api-publicmenudatabatchquery) | [`alipay.open.public.menu.data.batchquery`](https://opendocs.alipay.com/apis/05wyh5) | 菜单分析数据查询接口 | `POST` |
| [`PublicMenuModify`](#api-publicmenumodify) | [`alipay.open.public.menu.modify`](https://opendocs.alipay.com/apis/05xgci) | 默认菜单更新接口 | `POST` |
| [`PublicMenuQuery`](#api-publicmenuquery) | [`alipay.open.public.menu.query`](https://opendocs.alipay.com/pre-apis/05wuni) | 公众服务平台查询菜单 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicmenubatchquery"></a>

### PublicMenuBatchQuery — 菜单列表查询接口

- 官方接口：[`alipay.open.public.menu.batchquery`](https://opendocs.alipay.com/apis/05wol5)
- 路由：`POST /v3/alipay/open/public/menu/batchquery`
- 请求类型：[PublicMenuBatchQueryRequest](menu_request.go#L29)
- 响应类型：[PublicMenuBatchQueryResponse](menu_response.go#L5)

开发者可调用该接口查询自己创建的菜单列表

请求约束：PublicMenuBatchQueryRequest is the request type for alipay.open.public.menu.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publicmenucreate"></a>

### PublicMenuCreate — 默认菜单创建接口

- 官方接口：[`alipay.open.public.menu.create`](https://opendocs.alipay.com/apis/05wvxt)
- 路由：`POST /v3/alipay/open/public/menu/create`
- 请求类型：[PublicMenuCreateRequest](menu_request.go#L5)
- 响应类型：[PublicMenuCreateResponse](menu_response.go#L13)

开发者可调用该接口为生活号创建一套默认菜单，当用户进入生活号主页时，若没有命中任何个性化菜单，则会展示该默认菜单，每个生活号必须有一套默认菜单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `button` | `Button` | [`*[]ButtonObject`](#model-buttonobject) | JSON body | 未标注 | 一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。 |

<a id="api-publicmenudatabatchquery"></a>

### PublicMenuDataBatchQuery — 菜单分析数据查询接口

- 官方接口：[`alipay.open.public.menu.data.batchquery`](https://opendocs.alipay.com/apis/05wyh5)
- 路由：`POST /v3/alipay/open/public/menu/data/batchquery`
- 请求类型：[PublicMenuDataBatchQueryRequest](menu_request.go#L13)
- 响应类型：[PublicMenuDataBatchQueryResponse](menu_response.go#L19)

开发者可以根据本接口查询30天内菜单的相关数据，包括菜单点击人数和次数

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `begin_date` | `BeginDate` | `*string` | JSON body | 未标注 | 数据开始日期，时间格式为 "yyyyMMdd" 。 |
| `end_date` | `EndDate` | `*string` | JSON body | 未标注 | 数据结束日期，时间格式为"yyyyMMdd"。查询数据开始日期/结束日期时间跨度最大30天。 |

<a id="api-publicmenumodify"></a>

### PublicMenuModify — 默认菜单更新接口

- 官方接口：[`alipay.open.public.menu.modify`](https://opendocs.alipay.com/apis/05xgci)
- 路由：`POST /v3/alipay/open/public/menu/modify`
- 请求类型：[PublicMenuModifyRequest](menu_request.go#L21)
- 响应类型：[PublicMenuModifyResponse](menu_response.go#L25)

开发者可调用该接口更新默认菜单，菜单更新成功后实时生效

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `button` | `Button` | [`*[]ButtonObject`](#model-buttonobject) | JSON body | 未标注 | 一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。 |

<a id="api-publicmenuquery"></a>

### PublicMenuQuery — 公众服务平台查询菜单

- 官方接口：[`alipay.open.public.menu.query`](https://opendocs.alipay.com/pre-apis/05wuni)
- 路由：`GET /v3/alipay/open/public/menu/query`
- 请求类型：[PublicMenuQueryRequest](menu_request.go#L32)
- 响应类型：[PublicMenuQueryResponse](menu_response.go#L29)

开发者通过调用该接口查询菜单。

请求约束：PublicMenuQueryRequest is the request type for alipay.open.public.menu.query.

本接口无业务请求参数；传入空的请求结构体即可。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-buttonobject"></a>

### ButtonObject

类型定义：[ButtonObject](menu_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_param` | `ActionParam` | `*string` | 未标注 | 当actionType为link时，该参数为url链接；<br> 当actionType为out时，该参数为用户自定义参数；<br> 当actionType为tel时，该参数为电话号码。<br> 当action_type为map时，该参数为查看地图的关键字。<br> 当action_type为consumption时，该参数可不传。<br> 该参数最长255个字符，不允许冒号等特殊字符。 |
| `action_type` | `ActionType` | `*string` | 未标注 | 菜单类型：<br> out——事件型菜单；<br> link——链接型菜单；<br> tel——点击拨打电话；<br> map——点击查看地图；<br> consumption——点击查看用户与生活号管理员账号之间的消费记录 |
| `icon` | `Icon` | `*string` | 未标注 | icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `name` | `Name` | `*string` | 未标注 | 菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK |
| `sub_button` | `SubButton` | [`*[]SubButton`](#model-subbutton) | 未标注 | 二级菜单数组，若sub_button为空，则一级菜单必须指定action_type和action_param的值，二级菜单个数可以为1~5个。 |

<a id="model-subbutton"></a>

### SubButton

类型定义：[SubButton](menu_model.go#L61)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_param` | `ActionParam` | `*string` | 未标注 | 当actionType为link时，该参数为url链接； <br> 当actionType为out时，该参数为用户自定义参数； <br> 当actionType为tel时，该参数为电话号码。 <br> 当action_type为map时，该参数为查看地图的关键字。<br> 当action_type为consumption时，该参数可不传。 <br> 该参数最长255个字符，不允许冒号等特殊字符。 |
| `action_type` | `ActionType` | `*string` | 未标注 | 菜单类型：<br> out——事件型菜单；<br> link——链接型菜单；<br> tel——点击拨打电话；<br> map——点击查看地图；<br> consumption——点击查看用户与生活号管理员账号之间的消费记录 |
| `icon` | `Icon` | `*string` | 未标注 | icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `name` | `Name` | `*string` | 未标注 | 菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK |

## 使用示例

- [PublicMenuBatchQuery](api_menu_publicmenubatchquery_example_test.go)
- [PublicMenuCreate](api_menu_publicmenucreate_example_test.go)
- [PublicMenuDataBatchQuery](api_menu_publicmenudatabatchquery_example_test.go)
- [PublicMenuModify](api_menu_publicmenumodify_example_test.go)
- [PublicMenuQuery](api_menu_publicmenuquery_example_test.go)
