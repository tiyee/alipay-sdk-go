# open/public/personalized 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/personalized`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `personalized.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicPersonalizedExtensionBatchQuery`](#api-publicpersonalizedextensionbatchquery) | [`alipay.open.public.personalized.extension.batchquery`](https://opendocs.alipay.com/apis/05wcq6) | 扩展区批量查询接口 | `POST` |
| [`PublicPersonalizedExtensionCreate`](#api-publicpersonalizedextensioncreate) | [`alipay.open.public.personalized.extension.create`](https://opendocs.alipay.com/apis/05wcpn) | 个性化扩展区创建接口 | `POST` |
| [`PublicPersonalizedExtensionDelete`](#api-publicpersonalizedextensiondelete) | [`alipay.open.public.personalized.extension.delete`](https://opendocs.alipay.com/apis/05wydg) | 扩展区删除接口 | `DELETE` |
| [`PublicPersonalizedMenuCreate`](#api-publicpersonalizedmenucreate) | [`alipay.open.public.personalized.menu.create`](https://opendocs.alipay.com/apis/05x9eg) | 个性化菜单创建 | `POST` |
| [`PublicPersonalizedMenuDelete`](#api-publicpersonalizedmenudelete) | [`alipay.open.public.personalized.menu.delete`](https://opendocs.alipay.com/apis/05xa6t) | 个性化菜单删除 | `DELETE` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicpersonalizedextensionbatchquery"></a>

### PublicPersonalizedExtensionBatchQuery — 扩展区批量查询接口

- 官方接口：[`alipay.open.public.personalized.extension.batchquery`](https://opendocs.alipay.com/apis/05wcq6)
- 路由：`POST /v3/alipay/open/public/personalized/extension/batchquery`
- 请求类型：[PublicPersonalizedExtensionBatchQueryRequest](personalized_request.go#L27)
- 响应类型：[PublicPersonalizedExtensionBatchQueryResponse](personalized_response.go#L5)

开发者可使用此接口批量查询扩展区列表，包括默认扩展区和个性化扩展区

请求约束：PublicPersonalizedExtensionBatchQueryRequest is the request type for alipay.open.public.personalized.extension.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publicpersonalizedextensioncreate"></a>

### PublicPersonalizedExtensionCreate — 个性化扩展区创建接口

- 官方接口：[`alipay.open.public.personalized.extension.create`](https://opendocs.alipay.com/apis/05wcpn)
- 路由：`POST /v3/alipay/open/public/personalized/extension/create`
- 请求类型：[PublicPersonalizedExtensionCreateRequest](personalized_request.go#L5)
- 响应类型：[PublicPersonalizedExtensionCreateResponse](personalized_response.go#L13)

开发者可使用该接口创建一套个性化扩展区，符合该扩展区规则的用户进入生活号主页后会看到该个性化扩展区

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `areas` | `Areas` | [`*[]ExtensionArea`](#model-extensionarea) | JSON body | 未标注 | 扩展区列表，最大条数为3 |
| `label_rule` | `LabelRule` | [`*[]LabelRule`](#model-labelrule) | JSON body | 未标注 | 标签规则，目前限定只能传入1条，在扩展区上线后，满足该标签规则的用户进入生活号首页，将看到该套扩展区。生成标签及打标详情参见[标签管理](https://opendocs.alipay.com/fw/api/106877)。 |

<a id="api-publicpersonalizedextensiondelete"></a>

### PublicPersonalizedExtensionDelete — 扩展区删除接口

- 官方接口：[`alipay.open.public.personalized.extension.delete`](https://opendocs.alipay.com/apis/05wydg)
- 路由：`DELETE /v3/alipay/open/public/personalized/extension/delete`
- 请求类型：[PublicPersonalizedExtensionDeleteRequest](personalized_request.go#L30)
- 响应类型：[PublicPersonalizedExtensionDeleteResponse](personalized_response.go#L19)

开发者可调用该接口删除创建的一套扩展区，包括个性化扩展区和默认扩展区

请求约束：PublicPersonalizedExtensionDeleteRequest is the request type for alipay.open.public.personalized.extension.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `extension_key` | `ExtensionKey` | `*string` | query | 否 | 一套扩展区的key，删除默认扩展区时传入default ，查询扩展区列表可以获得每套扩展区的key |

<a id="api-publicpersonalizedmenucreate"></a>

### PublicPersonalizedMenuCreate — 个性化菜单创建

- 官方接口：[`alipay.open.public.personalized.menu.create`](https://opendocs.alipay.com/apis/05x9eg)
- 路由：`POST /v3/alipay/open/public/personalized/menu/create`
- 请求类型：[PublicPersonalizedMenuCreateRequest](personalized_request.go#L13)
- 响应类型：[PublicPersonalizedMenuCreateResponse](personalized_response.go#L23)

开发者可使用该接口创建个性化菜单，当生活号用户标签信息命中该菜单的标签规则时，用户将会看到该套菜单。个性化菜单创建完成后立即生效，用户支付宝客户端版本低于10.0.12时，若创建的是icon菜单，也会以文本菜单的形式展示。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `button` | `Button` | [`*[]ButtonObject`](#model-buttonobject) | JSON body | 未标注 | 一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。 |
| `group_id` | `GroupId` | `*string` | JSON body | 未标注 | 人群分组id。分组创建及管理接入详情参见 [分组管理](https://opendocs.alipay.com/fw/api/106931)。 注意：group_id 与 label_rule（标签规则）不能同时为空。 |
| `label_rule` | `LabelRule` | [`*[]LabelRule`](#model-labelrule) | JSON body | 未标注 | 标签规则，目前限定只能传入1条，在个性化菜单创建成功后，满足该标签规则的用户进入生活号首页，将看到该套菜单。生成标签及打标详情参见[标签管理](https://opendocs.alipay.com/fw/api/106877)。 注意：group_id（人群分组id） 与 label_rule 不能同时为空。 |
| `mobile_client_type` | `MobileClientType` | `*string` | JSON body | 未标注 | 手机客户端类型，枚举支持： *iphone； *android； *wp； 说明：不填为不区分机型。 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。 |

<a id="api-publicpersonalizedmenudelete"></a>

### PublicPersonalizedMenuDelete — 个性化菜单删除

- 官方接口：[`alipay.open.public.personalized.menu.delete`](https://opendocs.alipay.com/apis/05xa6t)
- 路由：`DELETE /v3/alipay/open/public/personalized/menu/delete`
- 请求类型：[PublicPersonalizedMenuDeleteRequest](personalized_request.go#L36)
- 响应类型：[PublicPersonalizedMenuDeleteResponse](personalized_response.go#L29)

开发者可使用该接口删除个性化菜单

请求约束：PublicPersonalizedMenuDeleteRequest is the request type for alipay.open.public.personalized.menu.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `menu_key` | `MenuKey` | `*string` | query | 否 | 要删除的个性化菜单key。通过 [alipay.open.public.personalized.menu.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.personalized.menu.create)(个性化菜单创建)接口创建个性化菜单后获取。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-buttonobject"></a>

### ButtonObject

类型定义：[ButtonObject](personalized_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_param` | `ActionParam` | `*string` | 未标注 | 当actionType为link时，该参数为url链接；<br> 当actionType为out时，该参数为用户自定义参数；<br> 当actionType为tel时，该参数为电话号码。<br> 当action_type为map时，该参数为查看地图的关键字。<br> 当action_type为consumption时，该参数可不传。<br> 该参数最长255个字符，不允许冒号等特殊字符。 |
| `action_type` | `ActionType` | `*string` | 未标注 | 菜单类型：<br> out——事件型菜单；<br> link——链接型菜单；<br> tel——点击拨打电话；<br> map——点击查看地图；<br> consumption——点击查看用户与生活号管理员账号之间的消费记录 |
| `icon` | `Icon` | `*string` | 未标注 | icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `name` | `Name` | `*string` | 未标注 | 菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK |
| `sub_button` | `SubButton` | [`*[]SubButton`](#model-subbutton) | 未标注 | 二级菜单数组，若sub_button为空，则一级菜单必须指定action_type和action_param的值，二级菜单个数可以为1~5个。 |

<a id="model-extensionarea"></a>

### ExtensionArea

类型定义：[ExtensionArea](personalized_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goto_url` | `GotoUrl` | `*string` | 未标注 | 跳转链接，当type为"image"时必传，必须是https或alipays开头的url链接。已废弃 |
| `height` | `Height` | `*int64` | 未标注 | 扩展区高度，当type值为"h5"时必填，取值范围为200-500的整数,单位是px |
| `name` | `Name` | `*string` | 未标注 | 扩展区名字 |
| `type` | `Type` | `*string` | 未标注 | 扩展区类型，当前支持的值为h5，表示扩展区内容类型为h5。 |
| `url` | `Url` | `*string` | 未标注 | 扩展区url，传入h5页面url，必须是https开头的链接 |

<a id="model-labelrule"></a>

### LabelRule

类型定义：[LabelRule](personalized_model.go#L33)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | 未标注 | 标签id |
| `label_value` | `LabelValue` | `*string` | 未标注 | 标签值，当有多个取值时用英文","分隔，不允许传入下划线"_"、竖线"&#124;"或者空格" "和方括号"["、"]" |
| `operator` | `Operator` | `*string` | 未标注 | 目前支持EQ（等于）、BETWEEN（范围）、IN（包含）三种操作符；每个标签支持的运算符可以通过[标签列表查询接口](https://docs.open.alipay.com/api_6/alipay.open.public.life.label.batchquery#sintq)获得。该字段允许为空，默认运算符为IN |

<a id="model-subbutton"></a>

### SubButton

类型定义：[SubButton](personalized_model.go#L67)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_param` | `ActionParam` | `*string` | 未标注 | 当actionType为link时，该参数为url链接； <br> 当actionType为out时，该参数为用户自定义参数； <br> 当actionType为tel时，该参数为电话号码。 <br> 当action_type为map时，该参数为查看地图的关键字。<br> 当action_type为consumption时，该参数可不传。 <br> 该参数最长255个字符，不允许冒号等特殊字符。 |
| `action_type` | `ActionType` | `*string` | 未标注 | 菜单类型：<br> out——事件型菜单；<br> link——链接型菜单；<br> tel——点击拨打电话；<br> map——点击查看地图；<br> consumption——点击查看用户与生活号管理员账号之间的消费记录 |
| `icon` | `Icon` | `*string` | 未标注 | icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `name` | `Name` | `*string` | 未标注 | 菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK |

## 使用示例

- [PublicPersonalizedExtensionBatchQuery](api_personalized_publicpersonalizedextensionbatchquery_example_test.go)
- [PublicPersonalizedExtensionCreate](api_personalized_publicpersonalizedextensioncreate_example_test.go)
- [PublicPersonalizedExtensionDelete](api_personalized_publicpersonalizedextensiondelete_example_test.go)
- [PublicPersonalizedMenuCreate](api_personalized_publicpersonalizedmenucreate_example_test.go)
- [PublicPersonalizedMenuDelete](api_personalized_publicpersonalizedmenudelete_example_test.go)
