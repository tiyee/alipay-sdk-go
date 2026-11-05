# commerce/educate 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/educate`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `educate.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_educate_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`EducateCampusInstitutionsAdd`](#api-educatecampusinstitutionsadd) | [`alipay.commerce.educate.campus.institutions.add`](https://opendocs.alipay.com/open-v3/2f573ce0_alipay.commerce.educate.campus.institutions.add) | 入驻学校信息 | `POST` |
| [`EducateCampusInstitutionsQuery`](#api-educatecampusinstitutionsquery) | [`alipay.commerce.educate.campus.institutions.query`](https://opendocs.alipay.com/open-v3/b340d586_alipay.commerce.educate.campus.institutions.query) | 查询学校信息 | `GET` |
| [`EducateSchoolCardOrderSync`](#api-educateschoolcardordersync) | [`alipay.commerce.educate.schoolcard.order.sync`](https://opendocs.alipay.com/open-v3/b4497dd8_alipay.commerce.educate.schoolcard.order.sync) | 校园卡支付订单数据同步 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-educatecampusinstitutionsadd"></a>

### EducateCampusInstitutionsAdd — 入驻学校信息

- 官方接口：[`alipay.commerce.educate.campus.institutions.add`](https://opendocs.alipay.com/open-v3/2f573ce0_alipay.commerce.educate.campus.institutions.add)
- 路由：`POST /v3/alipay/commerce/educate/campus/institutions/add`
- 请求类型：[EducateCampusInstitutionsAddRequest](educate_request.go#L5)
- 响应类型：[EducateCampusInstitutionsAddResponse](educate_response.go#L5)

学校信息入驻

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `card_pict_url` | `CardPictUrl` | `*string` | JSON body | 未标注 | 事业单位法人证书或办学许可证的图片。入驻学校时，如果流入人工审核，会根据此图片进行辅助验证。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 市 |
| `inst_name` | `InstName` | `*string` | JSON body | 未标注 | 入驻的学校名称，必须是完整的学校全称 |
| `inst_std_code` | `InstStdCode` | `*string` | JSON body | 未标注 | 学校外标，使用统一社会信用编码 |
| `learning_stage` | `LearningStage` | `*string` | JSON body | 未标注 | 办学阶段。 枚举值如下： KINDERGARTEN（幼儿园） PRIMARY_SCHOOL（小学） MIDDLE_SCHOOL（初中） HIGH_SCHOOL（高中） SECONDARY_VOCATIONAL_SCHOOL（中职中专） 注意：如果学校兼有多种属性，可以连写用英文逗号拆分，如：MIDDLE_SCHOOL,HIGH_SCHOOL 代表兼有初中部和高中部； |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份 |
| `school_property` | `SchoolProperty` | `*string` | JSON body | 未标注 | 学校性质.枚举值如下： 1：公立 2：民办 |

<a id="api-educatecampusinstitutionsquery"></a>

### EducateCampusInstitutionsQuery — 查询学校信息

- 官方接口：[`alipay.commerce.educate.campus.institutions.query`](https://opendocs.alipay.com/open-v3/b340d586_alipay.commerce.educate.campus.institutions.query)
- 路由：`GET /v3/alipay/commerce/educate/campus/institutions/query`
- 请求类型：[EducateCampusInstitutionsQueryRequest](educate_request.go#L72)
- 响应类型：[EducateCampusInstitutionsQueryResponse](educate_response.go#L15)

可以通过学校库产品查询支付宝系统中是否已存在对应的学校，避免重复入驻，以及通过该接口查询入驻的进度：入驻成功则会返回学校信息，失败则不会返回

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `inst_name` | `InstName` | `*string` | query | 否 | 学校名称 |
| `province_code` | `ProvinceCode` | `*string` | query | 否 | 省份编码，编码标准为中华人民共和国行政区划代码，参见 2020年12月中华人民共和国县以上行政区划代码。 |
| `city_code` | `CityCode` | `*string` | query | 否 | 城市编码，编码标准为中华人民共和国行政区划代码，参见 2020年12月中华人民共和国县以上行政区划代码。 |
| `inst_id` | `InstId` | `*string` | query | 否 | 学校内标：支付宝内部学校唯一编号。使用内标查询时，默认使用内标进行精确匹配。 |
| `inst_std_code` | `InstStdCode` | `*string` | query | 否 | 学校外标：统一社会信用编码或教育部提供的学校标识码。使用学校外标查询时，默认使用外标进行精确匹配。 |
| `like_property` | `LikeProperty` | `*string` | query | 否 | 是否使用学校名称模糊匹配进行查询，默认精确匹配。 枚举值如下： 1：精确匹配； 0：模糊匹配 |

<a id="api-educateschoolcardordersync"></a>

### EducateSchoolCardOrderSync — 校园卡支付订单数据同步

- 官方接口：[`alipay.commerce.educate.schoolcard.order.sync`](https://opendocs.alipay.com/open-v3/b4497dd8_alipay.commerce.educate.schoolcard.order.sync)
- 路由：`POST /v3/alipay/commerce/educate/schoolcard/order/sync`
- 请求类型：[EducateSchoolCardOrderSyncRequest](educate_request.go#L23)
- 响应类型：[EducateSchoolCardOrderSyncResponse](educate_response.go#L22)

校园卡支付订单数据同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `actual_amount` | `ActualAmount` | `*string` | JSON body | 未标注 | 实际金额（总支付金额），单位为【元】 |
| `applet_app_id` | `AppletAppId` | `*string` | JSON body | 未标注 | 小程序appid |
| `card_balance` | `CardBalance` | `*string` | JSON body | 未标注 | 校园卡余额 |
| `card_no` | `CardNo` | `*string` | JSON body | 未标注 | 128 |
| `card_type` | `CardType` | `*string` | JSON body | 未标注 | 校园卡类型 |
| `create_time` | `CreateTime` | `*string` | JSON body | 未标注 | 该笔订单真实的创建时间，需精确到毫秒。 |
| `discount_amount` | `DiscountAmount` | `*string` | JSON body | 未标注 | 优惠金额 |
| `goods_orders` | `GoodsOrders` | [`*GoodsOrder`](#model-goodsorder) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_name` | `MerchantName` | `*string` | JSON body | 未标注 | 商家名称，不传默认展示学校名称 |
| `modified_time` | `ModifiedTime` | `*string` | JSON body | 未标注 | 订单修改时间 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户open_id |
| `order_amount` | `OrderAmount` | `*string` | JSON body | 未标注 | 订单金额 |
| `order_detail_url` | `OrderDetailUrl` | `*string` | JSON body | 未标注 | 订单详情链接 |
| `order_status` | `OrderStatus` | `*string` | JSON body | 未标注 | 订单状态 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，由商家自定义，128个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。 |
| `pay_address` | `PayAddress` | `*string` | JSON body | 未标注 | 支付地点 |
| `pay_mode` | `PayMode` | `*string` | JSON body | 未标注 | 付款方式，不传默认展示学校名称+校园卡+（卡号后四位） |
| `rake_back_pid` | `RakeBackPid` | `*string` | JSON body | 未标注 | 系统商编号。该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID |
| `school_id` | `SchoolId` | `*string` | JSON body | 未标注 | 学校内标，录入学校接口返回的参数 |
| `school_pid` | `SchoolPid` | `*string` | JSON body | 未标注 | 学校收款账号 |
| `service_provider_name` | `ServiceProviderName` | `*string` | JSON body | 未标注 | 服务提供者名称 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 业务类型 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 买家支付宝用户ID。 2088开头的16位纯数字，小程序场景下获取用户ID请参考：用户授权; 其它场景下获取用户ID请参考：网页授权获取用户信息; |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-goodsorder"></a>

### GoodsOrder

类型定义：[GoodsOrder](educate_model.go#L17)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称 |
| `goods_picture_id` | `GoodsPictureId` | `*string` | 未标注 | 商品图片素材Id |
