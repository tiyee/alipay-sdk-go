# shop 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/offline/market/shop`。本包提供 6 个支付宝 V3 REST 操作。

通过 `shop.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MarketShopBatchQuery`](#api-marketshopbatchquery) | `alipay.offline.market.shop.batchquery` |
| [`MarketShopCategoryQuery`](#api-marketshopcategoryquery) | `alipay.offline.market.shop.category.query` |
| [`MarketShopCreate`](#api-marketshopcreate) | `alipay.offline.market.shop.create` |
| [`MarketShopModify`](#api-marketshopmodify) | `alipay.offline.market.shop.modify` |
| [`MarketShopQueryDetail`](#api-marketshopquerydetail) | `alipay.offline.market.shop.querydetail` |
| [`MarketShopSummaryBatchQuery`](#api-marketshopsummarybatchquery) | `alipay.offline.market.shop.summary.batchquery` |

## 使用示例

- [MarketShopBatchQuery](api_shop_marketshopbatchquery_example_test.go)
- [MarketShopCategoryQuery](api_shop_marketshopcategoryquery_example_test.go)
- [MarketShopCreate](api_shop_marketshopcreate_example_test.go)
- [MarketShopModify](api_shop_marketshopmodify_example_test.go)
- [MarketShopQueryDetail](api_shop_marketshopquerydetail_example_test.go)
- [MarketShopSummaryBatchQuery](api_shop_marketshopsummarybatchquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-marketshopbatchquery"></a>

### MarketShopBatchQuery — 查询商户的门店编号列表

- 官方接口：[`alipay.offline.market.shop.batchquery`](https://opendocs.alipay.com/apis/05xf6f)
- 路由：`POST /v3/alipay/offline/market/shop/batchquery`
- 请求类型：[OfflineMarketShopBatchQueryRequest](shop_request.go#L4)
- 响应类型：[OfflineMarketShopBatchQueryResponse](shop_response.go#L4)

系统商通过该接口可以查询对应APPID下所有的门店编号（支付宝口碑门店编号）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_channel` | `BizChannel` | `*string` | JSON body | 未标注 | 表示接口查询门店的业务渠道限制：ALL、 POS、不传。不传代表只查询普通门店，传入POS代表只查询简易门店，传入ALL代表查询普通门店和简易门店。 |
| `page_no` | `PageNo` | `*string` | JSON body | 未标注 | 页码，第一页传入"1"，默认500个结果为一页。此参数必须是大于0的正整数，为0时将查询报错。 |

<a id="api-marketshopcategoryquery"></a>

### MarketShopCategoryQuery — 门店类目配置查询接口

- 官方接口：[`alipay.offline.market.shop.category.query`](https://opendocs.alipay.com/apis/05xdx8)
- 路由：`GET /v3/alipay/offline/market/shop/category/query`
- 请求类型：[OfflineMarketShopCategoryQueryRequest](shop_request.go#L236)
- 响应类型：[OfflineMarketShopCategoryQueryResponse](shop_response.go#L14)

用于查询可用于开店的类目，以及类目上的配置约束信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `category_id` | `CategoryId` | `*string` | query | 否 | 类目ID，如果为空则查询全部类目。 |
| `op_role` | `OpRole` | `*string` | query | 否 | 表示接口业务的调用方身份,默认不填标识为ISV。 |

<a id="api-marketshopcreate"></a>

### MarketShopCreate — 创建门店信息

- 官方接口：[`alipay.offline.market.shop.create`](https://opendocs.alipay.com/apis/05woku)
- 路由：`POST /v3/alipay/offline/market/shop/create`
- 请求类型：[OfflineMarketShopCreateRequest](shop_request.go#L12)
- 响应类型：[OfflineMarketShopCreateResponse](shop_response.go#L20)

系统商需要通过该接口在口碑平台帮助商户创建门店信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | JSON body | 未标注 | 门店详细地址，地址字符长度在4-50个字符，注：不含省市区。门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层，“四川北路1552号欢乐广场1楼”。 |
| `audit_images` | `AuditImages` | `*string` | JSON body | 未标注 | 门店审核时需要的图片；至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过；多个图片之间以英文逗号分隔。 |
| `auth_letter` | `AuthLetter` | `*string` | JSON body | 未标注 | 门店授权函,营业执照与签约账号主体不一致时需要。 |
| `biz_version` | `BizVersion` | `*string` | JSON body | 未标注 | 店铺接口业务版本号，新接入的ISV，请统一传入2.0。 |
| `box` | `Box` | `*string` | JSON body | 未标注 | 门店是否有包厢，T表示有，F表示没有，不传在客户端不作展示。 |
| `branch_shop_name` | `BranchShopName` | `*string` | JSON body | 未标注 | 分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。 |
| `brand_logo` | `BrandLogo` | `*string` | JSON body | 未标注 | 品牌LOGO; 图片ID，不填写则默认为门店首图main_image。 |
| `brand_name` | `BrandName` | `*string` | JSON body | 未标注 | 品牌名，不填写则默认为“其它品牌”。 |
| `business_certificate` | `BusinessCertificate` | `*string` | JSON body | 未标注 | 许可证，各行业所需的证照资质参见[商户入驻要求](https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。 |
| `business_certificate_expires` | `BusinessCertificateExpires` | `*string` | JSON body | 未标注 | 许可证有效期，格式：2020-03-20或长期。严格按照格式填写。 |
| `business_time` | `BusinessTime` | `*string` | JSON body | 未标注 | 请严格按"周一-周五 09:00-20:00,周六-周日 10:00-22:00"的格式进行填写，时间段不能重复，最多支持两个时间段，24小时营业请填写"00:00-23:59" |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 类目id，请参考[商户入驻要求](https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码，国标码，详见国家统计局数据 [点此下载](https://gw.alipayobjects.com/os/basement_prod/92d874c4-39a9-4977-b710-d55dad374b7f.csv)。 |
| `contact_number` | `ContactNumber` | `*string` | JSON body | 未标注 | 门店电话号码；支持座机和手机，只支持数字和+-号，在客户端对用户展现， 支持多个电话， 以英文逗号分隔。 |
| `creator` | `Creator` | `*string` | JSON body | 未标注 | （支付宝内部参数）小二的支付宝账号。 |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区县编码，国标码，详见国家统计局数据 [点此下载](https://gw.alipayobjects.com/os/basement_prod/92d874c4-39a9-4977-b710-d55dad374b7f.csv)。 |
| `enterprise_logon_id` | `EnterpriseLogonId` | `*string` | JSON body | 未标注 | （支付宝内部参数）企业支付宝账号。 |
| `enterprise_name` | `EnterpriseName` | `*string` | JSON body | 未标注 | （支付宝内部参数）企业支付宝账户名称。 |
| `implement_id` | `ImplementId` | `*string` | JSON body | 未标注 | 机具号，多个之间以英文逗号分隔。 |
| `is_operating_online` | `IsOperatingOnline` | `*string` | JSON body | 未标注 | 是否在其他平台开店，T表示有开店，F表示未开店。 |
| `isv_uid` | `IsvUid` | `*string` | JSON body | 未标注 | ISV返佣id，门店创建、或者门店交易的返佣将通过此账号反给ISV，如果有口碑签订了返佣协议，则该字段作为返佣数据提取的依据。此字段必须是个合法uid，2088开头的16位支付宝会员账号，如果传入错误将无法创建门店。 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `leads_id` | `LeadsId` | `*string` | JSON body | 未标注 | （支付宝内部参数）leads的编号。 |
| `licence` | `Licence` | `*string` | JSON body | 未标注 | 门店营业执照图片，各行业所需的证照资质参见[商户入驻要求](https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。 |
| `licence_code` | `LicenceCode` | `*string` | JSON body | 未标注 | 门店营业执照编号，只支持输入中文，英文和数字，营业执照信息与is_operating_online至少填一项。 |
| `licence_expires` | `LicenceExpires` | `*string` | JSON body | 未标注 | 营业执照过期时间。格式：2020-10-20或长期。严格按照格式填写。 |
| `licence_name` | `LicenceName` | `*string` | JSON body | 未标注 | 门店营业执照名称。 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `main_image` | `MainImage` | `*string` | JSON body | 未标注 | 门店首图，非常重要，推荐尺寸2000*1500。 |
| `main_shop_name` | `MainShopName` | `*string` | JSON body | 未标注 | 主门店名 比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。 |
| `no_smoking` | `NoSmoking` | `*string` | JSON body | 未标注 | 是否有无烟区，T表示有无烟区，F表示没有无烟区，不传在客户端不展示。 |
| `notify_mobile` | `NotifyMobile` | `*string` | JSON body | 未标注 | 门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息， 不在客户端展示。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `online_image` | `OnlineImage` | `*string` | JSON body | 未标注 | 废弃字段，请使用online_url字段替代。 |
| `online_url` | `OnlineUrl` | `*string` | JSON body | 未标注 | 其他平台开店的店铺链接url，多个url使用英文逗号隔开,isv迁移到新接口使用此字段，与is_operating_online=T配套使用。 |
| `op_role` | `OpRole` | `*string` | JSON body | 未标注 | 表示以系统集成商的身份开店，开放平台现在统一传入ISV。 |
| `operate_notify_url` | `OperateNotifyUrl` | `*string` | JSON body | 未标注 | 当商户的门店审核状态发生变化时，会向该地址推送消息。 |
| `other_authorization` | `OtherAuthorization` | `*string` | JSON body | 未标注 | 其他资质。用于上传营业证照、许可证照外的其他资质，除已上传许可证外的其他许可证也可以在该字段上传。 |
| `parking` | `Parking` | `*string` | JSON body | 未标注 | 门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | （支付宝内部参数）服务商要操作的商户PID。 |
| `pay_type` | `PayType` | `*string` | JSON body | 未标注 | （支付宝内部参数）付款方式：code_scanned_pay：付款码 <br> online_pay：在线买单。ISV不可以指定此字段，ISV泛行业开店默认为在线买单，云纵开店可以指定支付方式。 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份编码，国标码，详见国家统计局数据 [点此下载](https://gw.alipayobjects.com/os/basement_prod/92d874c4-39a9-4977-b710-d55dad374b7f.csv)。 |
| `ref_apply_id` | `RefApplyId` | `*string` | JSON body | 未标注 | （支付宝内部参数）leads开店重试支付宝流水ID。 |
| `request_id` | `RequestId` | `*string` | JSON body | 未标注 | 支持英文字母和数字，由开发者自行定义（不允许重复），在门店notify消息中也会带有该参数，以此标明本次notify消息是对哪个请求的回应。 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。 |
| `value_added` | `ValueAdded` | `*string` | JSON body | 未标注 | 门店其他的服务，门店与用户线下兑现。 |
| `version` | `Version` | `*string` | JSON body | 未标注 | 废弃字段，使用biz_version字段替代。 |
| `wifi` | `Wifi` | `*string` | JSON body | 未标注 | 门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。 |

<a id="api-marketshopmodify"></a>

### MarketShopModify — 修改门店信息

- 官方接口：[`alipay.offline.market.shop.modify`](https://opendocs.alipay.com/apis/05wx1o)
- 路由：`POST /v3/alipay/offline/market/shop/modify`
- 请求类型：[OfflineMarketShopModifyRequest](shop_request.go#L112)
- 响应类型：[OfflineMarketShopModifyResponse](shop_response.go#L40)

系统商需要通过该接口在口碑平台帮助商户修改门店信息，所有的选项都为非必填项，需要修改什么信息就传入什么参数，不传入的参数，则保持不变。修改门店名、品牌名、品牌logo、首图、门店图片、营业执照、营业执照编号、营业执照名称、经营许可证、经营许可证有效期、门店授权函、是否在其它平台开店、在其它平台开店照片，会重新触发风控审核，修改其他字段不会触发风控审核。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | JSON body | 未标注 | 门店详细地址，地址字符长度在4-50个字符。门店详细地址，格式（不含省市区）：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层；注：门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名 |
| `audit_images` | `AuditImages` | `*string` | JSON body | 未标注 | 门店审核时需要的图片; 至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过， 多个以英文逗号分隔。 |
| `auth_letter` | `AuthLetter` | `*string` | JSON body | 未标注 | 门店授权函。 |
| `avg_price` | `AvgPrice` | `*string` | JSON body | 未标注 | 人均消费价格，最少1元，最大不超过99999元，请按实际情况填写，单位元。 |
| `biz_version` | `BizVersion` | `*string` | JSON body | 未标注 | 店铺接口业务版本号，新接入的ISV，请统一传入2.0。 |
| `box` | `Box` | `*string` | JSON body | 未标注 | 包厢支持，T表示有包厢，F表示无包厢；不传值默认F。 |
| `branch_shop_name` | `BranchShopName` | `*string` | JSON body | 未标注 | 分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。分店名长度需在2~20个字符之间。 |
| `brand_logo` | `BrandLogo` | `*string` | JSON body | 未标注 | 品牌LOGO; 图片ID，不填写则默认为门店首图main_image。 |
| `brand_name` | `BrandName` | `*string` | JSON body | 未标注 | 品牌名称；不填写则默认为“其它品牌”。 |
| `business_certificate` | `BusinessCertificate` | `*string` | JSON body | 未标注 | 许可证，各行业所需的证照资质参见[商户入驻要求](https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。 |
| `business_certificate_expires` | `BusinessCertificateExpires` | `*string` | JSON body | 未标注 | 许可证有效期，格式：2020-03-20或长期。严格按照格式填写。 |
| `business_time` | `BusinessTime` | `*string` | JSON body | 未标注 | 请严格按"周一-周五 09:00-20:00,周六-周日 10:00-22:00"的格式进行填写，时间段不能重复，最多支持两个时间段，24小时营业请填写"00:00-23:59" |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 废弃字段，不支持修改类目。<br> 类目id，请参考[商户入驻要求](https://doc.open.alipay.com/doc2/detail.htm?treeId=205&articleId=104497&docType=1)。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码，国标码，详见国家统计局数据 [点此下载](http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls)。 |
| `contact_number` | `ContactNumber` | `*string` | JSON body | 未标注 | 门店电话号码；支持座机和手机，在客户端对用户展现，支持多个电话，以英文逗号分隔。 |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区县编码，国标码，详见国家统计局数据 [点此下载](http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls)。 |
| `implement_id` | `ImplementId` | `*string` | JSON body | 未标注 | 店铺使用的机具编号，多个以英文逗号分隔。 |
| `is_operating_online` | `IsOperatingOnline` | `*string` | JSON body | 未标注 | 是否在其他平台开店，T表示有开店，F表示未开店。 |
| `is_show` | `IsShow` | `*string` | JSON body | 未标注 | 废弃字段，T表示显示，F表示隐藏，默认为T。 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `licence` | `Licence` | `*string` | JSON body | 未标注 | 门店营业执照图片，各行业所需的证照资质参见：https://doc.open.alipay.com/doc2/detail.htm?spm=a219a.7629140.0.0.NBvQVP&treeId=78&articleId=104497&docType=1。 |
| `licence_code` | `LicenceCode` | `*string` | JSON body | 未标注 | 门店营业执照编号。只支持输入中文，英文和数字。 |
| `licence_expires` | `LicenceExpires` | `*string` | JSON body | 未标注 | 营业执照过期时间。格式：2020-10-20或长期。严格按照格式填写。 |
| `licence_name` | `LicenceName` | `*string` | JSON body | 未标注 | 门店营业执照名称。 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `main_image` | `MainImage` | `*string` | JSON body | 未标注 | 门店首图；非常重要，推荐尺寸2000*1500。 |
| `main_shop_name` | `MainShopName` | `*string` | JSON body | 未标注 | 主门店名 比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。【xxx店】、(xxx店)、（xxx店）、[xxx店]、&lt;xxx店&gt;、xxx店，等类似的主店名都是不合法的，如果需要录入分店，请填写到branch_shop_name字段中。 |
| `no_smoking` | `NoSmoking` | `*string` | JSON body | 未标注 | 无烟区支持，T表示禁烟，F表示不禁烟；不传值默认F。 |
| `notify_mobile` | `NotifyMobile` | `*string` | JSON body | 未标注 | 门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息，不在客户端展示；多个以引文逗号分隔。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `online_image` | `OnlineImage` | `*string` | JSON body | 未标注 | 废弃字段，请使用online_url字段替代。 |
| `online_url` | `OnlineUrl` | `*string` | JSON body | 未标注 | 其他平台开店的店铺链接url,多个url使用英文逗号隔开，isv迁移到新接口使用此字段，与is_operating_online=T配套使用。 |
| `op_id` | `OpId` | `*string` | JSON body | 未标注 | （支付宝内部参数）操作员的支付账号ID(服务商ID、城市经理ID)。 |
| `op_role` | `OpRole` | `*string` | JSON body | 未标注 | 表示以系统集成商的身份开店，开放平台现在统一传入ISV。 |
| `operate_notify_url` | `OperateNotifyUrl` | `*string` | JSON body | 未标注 | 通知发送url;当商户的门店审核状态发生变化时，会向该地址推送消息。 |
| `other_authorization` | `OtherAuthorization` | `*string` | JSON body | 未标注 | 其他资质。用于上传营业证照、许可证照外的其他资质，除已上传许可证外的其他许可证也可以在该字段上传。 |
| `parking` | `Parking` | `*string` | JSON body | 未标注 | 门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | （支付宝内部参数）服务商要操作的商户PID。 |
| `pay_type` | `PayType` | `*string` | JSON body | 未标注 | （支付宝内部参数）付款方式：code_scanned_pay：付款码 <br> online_pay：在线买单。ISV不可以指定此字段，ISV泛行业开店默认为在线买单，云纵开店可以指定支付方式。 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份编码，国标码，详见国家统计局数据 [点此下载”](http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/2016.xls)。 |
| `request_id` | `RequestId` | `*string` | JSON body | 未标注 | 外部请求ID; 标识ISV本次修改的请求，由开发者自定义，不同的请求使用不同的ID，在门店notify消息中也会带有该参数，以此标明本次notify消息是对哪个请求的回应 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 支付宝门店ID。 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。 |
| `value_added` | `ValueAdded` | `*string` | JSON body | 未标注 | 门店其他的服务，门店与用户线下兑现。 |
| `version` | `Version` | `*string` | JSON body | 未标注 | 废弃字段，使用biz_version字段替代。 |
| `wifi` | `Wifi` | `*string` | JSON body | 未标注 | 门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。 |

<a id="api-marketshopquerydetail"></a>

### MarketShopQueryDetail — 查询单个门店信息接口

- 官方接口：[`alipay.offline.market.shop.querydetail`](https://opendocs.alipay.com/apis/05wuly)
- 路由：`GET /v3/alipay/offline/market/shop/querydetail`
- 请求类型：[OfflineMarketShopQueryDetailRequest](shop_request.go#L244)
- 响应类型：[OfflineMarketShopQueryDetailResponse](shop_response.go#L54)

系统商通过该接口可以查询单个门店的详细信息，包括门店基础信息，门店状态等信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `shop_id` | `ShopId` | `*string` | query | 否 | 支付宝门店ID |
| `op_role` | `OpRole` | `*string` | query | 否 | 服务商及商户调用情况下务必传递。操作人角色，默认商户操作:MERCHANT；服务商操作：PROVIDER；ISV: 不需要填写 |

<a id="api-marketshopsummarybatchquery"></a>

### MarketShopSummaryBatchQuery — 门店摘要信息批量查询接口

- 官方接口：[`alipay.offline.market.shop.summary.batchquery`](https://opendocs.alipay.com/apis/05wx2m)
- 路由：`POST /v3/alipay/offline/market/shop/summary/batchquery`
- 请求类型：[OfflineMarketShopSummaryBatchQueryRequest](shop_request.go#L208)
- 响应类型：[OfflineMarketShopSummaryBatchQueryResponse](shop_response.go#L160)

用于进行门店摘要信息批量查询接口。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_channel` | `BizChannel` | `*string` | JSON body | 未标注 | 表示接口查询门店的业务渠道限制：ALL、 POS、不传。不传代表只查询普通门店，传入POS代表只查询简易门店，传入ALL代表查询普通门店和简易门店。 |
| `brand_name` | `BrandName` | `*string` | JSON body | 未标注 | 品牌名 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码，国标码，如：120100表示天津市 |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区域编码，如：120104表示南开区 |
| `op_role` | `OpRole` | `*string` | JSON body | 未标注 | 表示接口业务的调用方身份：ISV、 服务商身份标识。传入ISV代表系统集成商身份。传入PROVIDER代表服务商。 |
| `page_no` | `PageNo` | `*int64` | JSON body | 未标注 | 页码，留空标示第一页，默认 20个结果为一页 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页记录数，默认20，最大 100 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份编码，国标码，如：120000表示天津 |
| `query_type` | `QueryType` | `*string` | JSON body | 未标注 | 门店数据查询类型，根据类型可以返回指定的门店数据，目前支持的类型如下：<br> BRAND_RELATION ： 品牌商关联店铺<br> MALL_SELF ：MALL自己的门店<br> MALL_RELATION：MALL关联下的门店<br> MERCHANT_SELF:商户自己的门店<br> KB_PROMOTER：口碑客推广者 |
| `related_partner_id` | `RelatedPartnerId` | `*string` | JSON body | 未标注 | query_type查询类型下所关联的商户PID |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 门店ID |
| `shop_status` | `ShopStatus` | `*string` | JSON body | 未标注 | 门店状态，传入多个状态，多个状态使用英文逗号隔开，例如：PAUSED,OPEN<br> 店铺状态：OPEN（营业）、PAUSED（暂停）、INIT（初始）、FREEZE（冻结）、CLOSED（关店） |
