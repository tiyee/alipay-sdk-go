# ant/merchant 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ant/merchant`。本目录提供 25 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `merchant.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ExpandAppRecommendAccountCreate`](#api-expandapprecommendaccountcreate) | [`ant.merchant.expand.apprecommend.account.create`](https://opendocs.alipay.com/open-v3/32d136f0_ant.merchant.expand.apprecommend.account.create) | 关联账号和小程序 | `POST` |
| [`ExpandAppRecommendAccountDelete`](#api-expandapprecommendaccountdelete) | [`ant.merchant.expand.apprecommend.account.delete`](https://opendocs.alipay.com/open-v3/130a15de_ant.merchant.expand.apprecommend.account.delete) | 取消关联账号和小程序 | `DELETE` |
| [`ExpandAppRecommendAccountQuery`](#api-expandapprecommendaccountquery) | [`ant.merchant.expand.apprecommend.account.query`](https://opendocs.alipay.com/open-v3/cb07737c_ant.merchant.expand.apprecommend.account.query) | 查询已关联指定APP的账号列表 | `GET` |
| [`ExpandAppRecommendAvailableQuery`](#api-expandapprecommendavailablequery) | [`ant.merchant.expand.apprecommend.available.query`](https://opendocs.alipay.com/open-v3/fc63d6c8_ant.merchant.expand.apprecommend.available.query) | 查询可关联APP的账号列表 | `GET` |
| [`ExpandIndirectImageUpload`](#api-expandindirectimageupload) | [`ant.merchant.expand.indirect.image.upload`](https://opendocs.alipay.com/open-v3/af8e3a42_ant.merchant.expand.indirect.image.upload) | 图片上传 | `POST` |
| [`ExpandInfoQuery`](#api-expandinfoquery) | [`ant.merchant.expand.info.query`](https://opendocs.alipay.com/apis/05vb3x) | 蚂蚁商户信息查询接口 | `GET` |
| [`ExpandItemCreate`](#api-expanditemcreate) | [`ant.merchant.expand.item.create`](https://opendocs.alipay.com/open-v3/bf0fb8ad_ant.merchant.expand.item.create) | 商品创建接口 | `POST` |
| [`ExpandItemDelete`](#api-expanditemdelete) | [`ant.merchant.expand.item.delete`](https://opendocs.alipay.com/open-v3/ac2de50f_ant.merchant.expand.item.delete) | 商品删除接口 | `DELETE` |
| [`ExpandItemModify`](#api-expanditemmodify) | [`ant.merchant.expand.item.modify`](https://opendocs.alipay.com/open-v3/60d35f53_ant.merchant.expand.item.modify) | 商品修改接口 | `POST` |
| [`ExpandItemOpenBatchQuery`](#api-expanditemopenbatchquery) | [`ant.merchant.expand.item.open.batchquery`](https://opendocs.alipay.com/pre-apis/05vjno) | 批量查询商品接口 | `POST` |
| [`ExpandItemOpenCreate`](#api-expanditemopencreate) | [`ant.merchant.expand.item.open.create`](https://opendocs.alipay.com/pre-apis/05vn67) | 创建商品接口 | `POST` |
| [`ExpandItemOpenDelete`](#api-expanditemopendelete) | [`ant.merchant.expand.item.open.delete`](https://opendocs.alipay.com/pre-apis/05v7ca) | 删除商品接口 | `DELETE` |
| [`ExpandItemOpenModify`](#api-expanditemopenmodify) | [`ant.merchant.expand.item.open.modify`](https://opendocs.alipay.com/pre-apis/05vda6) | 修改商品接口 | `POST` |
| [`ExpandItemOpenQuery`](#api-expanditemopenquery) | [`ant.merchant.expand.item.open.query`](https://opendocs.alipay.com/pre-apis/05vlue) | 查询商品接口 | `GET` |
| [`ExpandItemQuery`](#api-expanditemquery) | [`ant.merchant.expand.item.query`](https://opendocs.alipay.com/open-v3/aad82813_ant.merchant.expand.item.query) | 商品查询接口 | `GET` |
| [`ExpandItemStatusModify`](#api-expanditemstatusmodify) | [`ant.merchant.expand.item.status.modify`](https://opendocs.alipay.com/open-v3/5cb8e04e_ant.merchant.expand.item.status.modify) | 商品状态修改接口 | `POST` |
| [`ExpandMCCQuery`](#api-expandmccquery) | [`ant.merchant.expand.mcc.query`](https://opendocs.alipay.com/open-v3/b00ab2d2_ant.merchant.expand.mcc.query) | 商户mcc信息查询 | `GET` |
| [`ExpandOrderQuery`](#api-expandorderquery) | [`ant.merchant.expand.order.query`](https://opendocs.alipay.com/open-v3/dd6e9adb_ant.merchant.expand.order.query) | 商户申请单查询 | `GET` |
| [`ExpandShopClose`](#api-expandshopclose) | [`ant.merchant.expand.shop.close`](https://opendocs.alipay.com/open-v3/4950631b_ant.merchant.expand.shop.close) | 蚂蚁门店关闭 | `PATCH` |
| [`ExpandShopConsult`](#api-expandshopconsult) | [`ant.merchant.expand.shop.consult`](https://opendocs.alipay.com/apis/05vda1) | 蚂蚁店铺创建咨询 | `POST` |
| [`ExpandShopCreate`](#api-expandshopcreate) | [`ant.merchant.expand.shop.create`](https://opendocs.alipay.com/open-v3/99b91969_ant.merchant.expand.shop.create) | 蚂蚁门店创建 | `POST` |
| [`ExpandShopModify`](#api-expandshopmodify) | [`ant.merchant.expand.shop.modify`](https://opendocs.alipay.com/open-v3/b1a37213_ant.merchant.expand.shop.modify) | 修改蚂蚁门店 | `PATCH` |
| [`ExpandShopPageQuery`](#api-expandshoppagequery) | [`ant.merchant.expand.shop.page.query`](https://opendocs.alipay.com/open-v3/fe443839_ant.merchant.expand.shop.page.query) | 店铺分页查询接口 | `GET` |
| [`ExpandShopQuery`](#api-expandshopquery) | [`ant.merchant.expand.shop.query`](https://opendocs.alipay.com/open-v3/b80b5948_ant.merchant.expand.shop.query) | 门店查询接口 | `GET` |
| [`ExpandShopReceiptAccountSave`](#api-expandshopreceiptaccountsave) | [`ant.merchant.expand.shop.receiptaccount.save`](https://opendocs.alipay.com/open-v3/74728d59_ant.merchant.expand.shop.receiptaccount.save) | 店铺增加收单账号 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-expandapprecommendaccountcreate"></a>

### ExpandAppRecommendAccountCreate — 关联账号和小程序

- 官方接口：[`ant.merchant.expand.apprecommend.account.create`](https://opendocs.alipay.com/open-v3/32d136f0_ant.merchant.expand.apprecommend.account.create)
- 路由：`POST /v3/ant/merchant/expand/apprecommend/account/create`
- 请求类型：[ExpandAppRecommendAccountCreateRequest](merchant_request.go#L9)
- 响应类型：[ExpandAppRecommendAccountCreateResponse](merchant_response.go#L5)

将账号与小程序进行绑定，用于支付推荐场景

请求约束：ExpandAppRecommendAccountCreateRequest is the official V3 AntMerchantExpandApprecommendAccountCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `acc_no` | `AccNo` | `*string` | JSON body | 未标注 | 待绑定账号PID |
| `app_no` | `AppNo` | `*string` | JSON body | 未标注 | 待绑定小程序的app_id |

<a id="api-expandapprecommendaccountdelete"></a>

### ExpandAppRecommendAccountDelete — 取消关联账号和小程序

- 官方接口：[`ant.merchant.expand.apprecommend.account.delete`](https://opendocs.alipay.com/open-v3/130a15de_ant.merchant.expand.apprecommend.account.delete)
- 路由：`DELETE /v3/ant/merchant/expand/apprecommend/account/delete`
- 请求类型：[ExpandAppRecommendAccountDeleteRequest](merchant_request.go#L357)
- 响应类型：[ExpandAppRecommendAccountDeleteResponse](merchant_response.go#L9)

取消账号与小程序的绑定，用于支付推荐场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_no` | `AppNo` | `*string` | query | 否 | 待解绑小程序的app_id |
| `acc_no` | `AccNo` | `*string` | query | 否 | 待解绑账号PID |

<a id="api-expandapprecommendaccountquery"></a>

### ExpandAppRecommendAccountQuery — 查询已关联指定APP的账号列表

- 官方接口：[`ant.merchant.expand.apprecommend.account.query`](https://opendocs.alipay.com/open-v3/cb07737c_ant.merchant.expand.apprecommend.account.query)
- 路由：`GET /v3/ant/merchant/expand/apprecommend/account/query`
- 请求类型：[ExpandAppRecommendAccountQueryRequest](merchant_request.go#L365)
- 响应类型：[ExpandAppRecommendAccountQueryResponse](merchant_response.go#L13)

查询已关联指定APP的账号列表，用于支付推荐场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_no` | `AppNo` | `*string` | query | 否 | 已关联小程序的app_id |
| `page_size` | `PageSize` | `*int64` | query | 否 | 单页行数，不能超过100行，单位：行 |
| `page_number` | `PageNumber` | `*int64` | query | 否 | 页码 |

<a id="api-expandapprecommendavailablequery"></a>

### ExpandAppRecommendAvailableQuery — 查询可关联APP的账号列表

- 官方接口：[`ant.merchant.expand.apprecommend.available.query`](https://opendocs.alipay.com/open-v3/fc63d6c8_ant.merchant.expand.apprecommend.available.query)
- 路由：`GET /v3/ant/merchant/expand/apprecommend/available/query`
- 请求类型：[ExpandAppRecommendAvailableQueryRequest](merchant_request.go#L375)
- 响应类型：[ExpandAppRecommendAvailableQueryResponse](merchant_response.go#L27)

查询可关联APP的账号列表，用于支付推荐场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_size` | `PageSize` | `*int64` | query | 否 | 单页行数，不能超过100行 |
| `page_number` | `PageNumber` | `*int64` | query | 否 | 页码 |

<a id="api-expandindirectimageupload"></a>

### ExpandIndirectImageUpload — 图片上传

- 官方接口：[`ant.merchant.expand.indirect.image.upload`](https://opendocs.alipay.com/open-v3/af8e3a42_ant.merchant.expand.indirect.image.upload)
- 路由：`POST /v3/ant/merchant/indirect/image`
- 请求类型：[ExpandIndirectImageUploadRequest](merchant_request.go#L429)
- 响应类型：[ExpandIndirectImageUploadResponse](merchant_response.go#L117)

图片上传

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*ExpandIndirectImageUpload`](#model-expandindirectimageupload) | multipart | 未标注 | the data multipart field. |
| `image_content` | `ImageContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | the image_content multipart field. |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-expandinfoquery"></a>

### ExpandInfoQuery — 蚂蚁商户信息查询接口

- 官方接口：[`ant.merchant.expand.info.query`](https://opendocs.alipay.com/apis/05vb3x)
- 路由：`GET /v3/ant/merchant/expand/info/query`
- 请求类型：[ExpandInfoQueryRequest](merchant_request.go#L383)
- 响应类型：[ExpandInfoQueryResponse](merchant_response.go#L41)

提供商户号查询商户基础信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_no` | `MerchantNo` | `*string` | query | 否 | 商户号，目前是商户pid |

<a id="api-expanditemcreate"></a>

### ExpandItemCreate — 商品创建接口

- 官方接口：[`ant.merchant.expand.item.create`](https://opendocs.alipay.com/open-v3/bf0fb8ad_ant.merchant.expand.item.create)
- 路由：`POST /v3/ant/merchant/expand/item/create`
- 请求类型：[ExpandItemCreateRequest](merchant_request.go#L17)
- 响应类型：[ExpandItemCreateResponse](merchant_response.go#L47)

用于服务商或商户创建商品

请求约束：ExpandItemCreateRequest is the official V3 AntMerchantExpandItemCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `detail_url` | `DetailUrl` | `*string` | JSON body | 未标注 | 详情地址 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 商品扩展信息：可以解析成 Map&lt;String, String&gt; 的 json string |
| `external_item_id` | `ExternalItemId` | `*string` | JSON body | 未标注 | 外部商品ID |
| `front_category_id` | `FrontCategoryId` | `*string` | JSON body | 未标注 | 前台类目id：target_type + target_id 和 front_category_id 二选一 |
| `label_list` | `LabelList` | [`*[]ItemLabelCreateInfo`](#model-itemlabelcreateinfo) | JSON body | 未标注 | 商品标签列表 |
| `main_pic` | `MainPic` | `*string` | JSON body | 未标注 | 主图地址 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 商品名称 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 场景：GAS_CHARGE（加油） |
| `sku_list` | `SkuList` | [`*[]ItemSkuCreateInfo`](#model-itemskucreateinfo) | JSON body | 未标注 | 商品sku列表，至少有一个 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 商户归属主体id |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 商品归属主体类型：target_type + target_id 和 front_category_id 二选一<br> 商品归属主体类型:<br> 5: 店铺<br> 4: 主站MID<br> 3: 参与者<br> 2: 角色<br> 1: 联系人 |

<a id="api-expanditemdelete"></a>

### ExpandItemDelete — 商品删除接口

- 官方接口：[`ant.merchant.expand.item.delete`](https://opendocs.alipay.com/open-v3/ac2de50f_ant.merchant.expand.item.delete)
- 路由：`DELETE /v3/ant/merchant/expand/item/delete`
- 请求类型：[ExpandItemDeleteRequest](merchant_request.go#L389)
- 响应类型：[ExpandItemDeleteResponse](merchant_response.go#L53)

用于服务商或商户删除商品

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | query | 否 | 商品ID（item_id不为空则以item_id作为删除KEY,item_id为空则以external_item_id作为删除KEY,两者不能同时为空） |
| `external_item_id` | `ExternalItemId` | `*string` | query | 否 | 外部商品ID（item_id不为空则以item_id作为删除KEY,item_id为空则以external_item_id作为删除KEY,两者不能同时为空） |

<a id="api-expanditemmodify"></a>

### ExpandItemModify — 商品修改接口

- 官方接口：[`ant.merchant.expand.item.modify`](https://opendocs.alipay.com/open-v3/60d35f53_ant.merchant.expand.item.modify)
- 路由：`POST /v3/ant/merchant/expand/item/modify`
- 请求类型：[ExpandItemModifyRequest](merchant_request.go#L43)
- 响应类型：[ExpandItemModifyResponse](merchant_response.go#L59)

用于服务商或商户修改商品

请求约束：ExpandItemModifyRequest is the official V3 AntMerchantExpandItemModifyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `detail_url` | `DetailUrl` | `*string` | JSON body | 未标注 | 商品详情地址 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 商品扩展信息：可以解析成 Map&lt;String, String&gt; 的 json string |
| `external_item_id` | `ExternalItemId` | `*string` | JSON body | 未标注 | 外部商品ID（item_id不为空时作为普通更新项，item_id为空时作为更新KEY） |
| `front_category_id` | `FrontCategoryId` | `*string` | JSON body | 未标注 | 前台类目id, null表示未分类 |
| `item_id` | `ItemId` | `*string` | JSON body | 未标注 | 商品id(和external_item_id不能同时为空) |
| `label_list` | `LabelList` | [`*[]ItemLabelModifyInfo`](#model-itemlabelmodifyinfo) | JSON body | 未标注 | 商品标签列表，null表示不修改标签列表，空list表示清空标签 |
| `main_pic` | `MainPic` | `*string` | JSON body | 未标注 | 商品主图 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 商品名称 |
| `op_timestamp` | `OpTimestamp` | `*int64` | JSON body | 未标注 | 更新时间戳（只处理时间戳最大的一次请求） |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 场景：GAS_CHARGE（加油） |
| `sku_list` | `SkuList` | [`*[]ItemSkuModifyInfo`](#model-itemskumodifyinfo) | JSON body | 未标注 | SKU列表（会和商品已存在SKU做差异化比较后做增删改操作） |
| `status` | `Status` | `*string` | JSON body | 未标注 | 商品状态：<br> EFFECT（有效）、 INVALID（无效） |

<a id="api-expanditemopenbatchquery"></a>

### ExpandItemOpenBatchQuery — 批量查询商品接口

- 官方接口：[`ant.merchant.expand.item.open.batchquery`](https://opendocs.alipay.com/pre-apis/05vjno)
- 路由：`POST /v3/ant/merchant/expand/item/open/batchquery`
- 请求类型：[ExpandItemOpenBatchQueryRequest](merchant_request.go#L71)
- 响应类型：[ExpandItemOpenBatchQueryResponse](merchant_response.go#L65)

用于ISV或商户批量查询商品

请求约束：ExpandItemOpenBatchQueryRequest is the official V3 AntMerchantExpandItemOpenBatchqueryModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `item_id_list` | `ItemIdList` | `*[]string` | JSON body | 未标注 | 商品ID列表 |

<a id="api-expanditemopencreate"></a>

### ExpandItemOpenCreate — 创建商品接口

- 官方接口：[`ant.merchant.expand.item.open.create`](https://opendocs.alipay.com/pre-apis/05vn67)
- 路由：`POST /v3/ant/merchant/expand/item/open/create`
- 请求类型：[ExpandItemOpenCreateRequest](merchant_request.go#L77)
- 响应类型：[ExpandItemOpenCreateResponse](merchant_response.go#L71)

用于ISV或商户创建商品

请求约束：ExpandItemOpenCreateRequest is the official V3 AntMerchantExpandItemOpenCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `description` | `Description` | `*string` | JSON body | 未标注 | 商品描述 |
| `ext_info` | `ExtInfo` | [`*[]ItemExtInfo`](#model-itemextinfo) | JSON body | 未标注 | 商品扩展信息（具体KEY请参见产品文档）。 小程序订单中心场景接入参见 [接入指南](https://opendocs.alipay.com/mini/024hj4)。 |
| `material_list` | `MaterialList` | [`*[]MaterialCreateInfo`](#model-materialcreateinfo) | JSON body | 未标注 | 素材列表（最多3个），素材内容为素材key（素材key为[alipay.merchant.item.file.upload](https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload) 接口返回的 material_key） |
| `name` | `Name` | `*string` | JSON body | 未标注 | 商品名称 |
| `property_list` | `PropertyList` | [`*[]ItemPropertyInfo`](#model-itempropertyinfo) | JSON body | 未标注 | 商品属性列表 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 业务场景码。 |
| `sku_list` | `SkuList` | [`*[]SkuCreateInfo`](#model-skucreateinfo) | JSON body | 未标注 | 商品SKU列表（至少1个，最多20个） |
| `standard_category_id` | `StandardCategoryId` | `*string` | JSON body | 未标注 | 商品所属标准类目ID，标品及非标品都需传入。 请填入 [小程序商品类目表](https://opendocs.alipay.com/mini/011lxt) 中三级类目ID。 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 商品类型。枚举支持： STANDARD_GOODS：标品，一般是具有明确、标准规格、型号、参数的商品，如：手机、数码产品、大多数的家电 。 NON_STANDARD_GOODS：非标品，则是在这些方面没有统一的行业标准和参数规格的商品，如：服装、鞋袜等。 |

<a id="api-expanditemopendelete"></a>

### ExpandItemOpenDelete — 删除商品接口

- 官方接口：[`ant.merchant.expand.item.open.delete`](https://opendocs.alipay.com/pre-apis/05v7ca)
- 路由：`DELETE /v3/ant/merchant/expand/item/open/delete`
- 请求类型：[ExpandItemOpenDeleteRequest](merchant_request.go#L397)
- 响应类型：[ExpandItemOpenDeleteResponse](merchant_response.go#L77)

用于ISV或商户删除商品

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | query | 否 | 商品ID |

<a id="api-expanditemopenmodify"></a>

### ExpandItemOpenModify — 修改商品接口

- 官方接口：[`ant.merchant.expand.item.open.modify`](https://opendocs.alipay.com/pre-apis/05vda6)
- 路由：`POST /v3/ant/merchant/expand/item/open/modify`
- 请求类型：[ExpandItemOpenModifyRequest](merchant_request.go#L103)
- 响应类型：[ExpandItemOpenModifyResponse](merchant_response.go#L81)

用于ISV或商户修改商品

请求约束：ExpandItemOpenModifyRequest is the official V3 AntMerchantExpandItemOpenModifyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `description` | `Description` | `*string` | JSON body | 未标注 | 商品描述 |
| `ext_info` | `ExtInfo` | [`*[]ItemExtInfo`](#model-itemextinfo) | JSON body | 未标注 | 商品扩展信息（覆盖商品已存在扩展信息、具体KEY请参见产品文档） |
| `item_id` | `ItemId` | `*string` | JSON body | 未标注 | 商品ID |
| `material_list` | `MaterialList` | [`*[]MaterialModifyInfo`](#model-materialmodifyinfo) | JSON body | 未标注 | 商品素材列表（会和商品已存在素材做差异化比较后做增删改操作） |
| `name` | `Name` | `*string` | JSON body | 未标注 | 商品名称 |
| `property_list` | `PropertyList` | [`*[]ItemPropertyInfo`](#model-itempropertyinfo) | JSON body | 未标注 | 商品属性列表（覆盖商品已存在属性） |
| `sku_list` | `SkuList` | [`*[]SkuModifyInfo`](#model-skumodifyinfo) | JSON body | 未标注 | SKU列表（会和商品已存在SKU做差异化比较后做增删改操作） |
| `standard_category_id` | `StandardCategoryId` | `*string` | JSON body | 未标注 | 商品所属标准类目ID，当 type 为STANDARD_GOODS（标品）时必填。 请填入 https://opendocs.alipay.com/mini/011lxt 小程序商品类目表中三级类目ID。 |
| `status` | `Status` | `*string` | JSON body | 未标注 | 商品状态 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 商品类型：STANDARD_GOODS（标品）、NON_STANDARD_GOODS（非标品） |

<a id="api-expanditemopenquery"></a>

### ExpandItemOpenQuery — 查询商品接口

- 官方接口：[`ant.merchant.expand.item.open.query`](https://opendocs.alipay.com/pre-apis/05vlue)
- 路由：`GET /v3/ant/merchant/expand/item/open/query`
- 请求类型：[ExpandItemOpenQueryRequest](merchant_request.go#L403)
- 响应类型：[ExpandItemOpenQueryResponse](merchant_response.go#L85)

用于ISV或商户查询商品

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_id` | `TargetId` | `*string` | query | 否 | 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID |
| `scene` | `Scene` | `*string` | query | 否 | 场景码（具体值请参见产品文档）。 |
| `target_type` | `TargetType` | `*string` | query | 否 | 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。 |
| `status` | `Status` | `*string` | query | 否 | 商品状态 |

<a id="api-expanditemquery"></a>

### ExpandItemQuery — 商品查询接口

- 官方接口：[`ant.merchant.expand.item.query`](https://opendocs.alipay.com/open-v3/aad82813_ant.merchant.expand.item.query)
- 路由：`GET /v3/ant/merchant/expand/item/query`
- 请求类型：[ExpandItemQueryRequest](merchant_request.go#L415)
- 响应类型：[ExpandItemQueryResponse](merchant_response.go#L91)

用于服务商或商户查询其商品信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_type` | `TargetType` | `*string` | query | 否 | 商品归属主体类型:<br> 5: 店铺<br> 4: 主站MID<br> 3: 参与者<br> 2: 角色<br> 1: 联系人 |
| `target_id` | `TargetId` | `*string` | query | 否 | 商品归属主体id |
| `undefined_category` | `UndefinedCategory` | `*bool` | query | 否 | 未分类类目: true 的时候 front_category_id 必须不填（null） |
| `front_category_id` | `FrontCategoryId` | `*string` | query | 否 | 商品前台类目id: null 表示所有分类 |
| `status` | `Status` | `*string` | query | 否 | 商品状态：EFFECT、INVALID、PAUSE；不填返回所有状态商品 |

<a id="api-expanditemstatusmodify"></a>

### ExpandItemStatusModify — 商品状态修改接口

- 官方接口：[`ant.merchant.expand.item.status.modify`](https://opendocs.alipay.com/open-v3/5cb8e04e_ant.merchant.expand.item.status.modify)
- 路由：`POST /v3/ant/merchant/expand/item/status/modify`
- 请求类型：[ExpandItemStatusModifyRequest](merchant_request.go#L127)
- 响应类型：[ExpandItemStatusModifyResponse](merchant_response.go#L97)

用于服务商或商户修改商品状态

请求约束：ExpandItemStatusModifyRequest is the official V3 AntMerchantExpandItemStatusModifyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | JSON body | 未标注 | 商品ID |
| `status` | `Status` | `*string` | JSON body | 未标注 | 商品状态：EFFECT、INVALID |

<a id="api-expandmccquery"></a>

### ExpandMCCQuery — 商户mcc信息查询

- 官方接口：[`ant.merchant.expand.mcc.query`](https://opendocs.alipay.com/open-v3/b00ab2d2_ant.merchant.expand.mcc.query)
- 路由：`GET /v3/ant/merchant/mcc/query`
- 请求类型：[ExpandMCCQueryRequest](merchant_request.go#L437)
- 响应类型：[ExpandMCCQueryResponse](merchant_response.go#L123)

用于查询商户的新版本mcc行业信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mcc_code_list` | `MccCodeList` | `*string` | query | 否 | 传入需要查询的code列表，以英文逗号分隔 |

<a id="api-expandorderquery"></a>

### ExpandOrderQuery — 商户申请单查询

- 官方接口：[`ant.merchant.expand.order.query`](https://opendocs.alipay.com/open-v3/dd6e9adb_ant.merchant.expand.order.query)
- 路由：`GET /v3/ant/merchant/order/{order_id}`
- 请求类型：[ExpandOrderQueryRequest](merchant_request.go#L443)
- 响应类型：[ExpandOrderQueryResponse](merchant_response.go#L129)

服务商根据申请单id，查询自己提交的商户进件、管理等申请单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderId` | `string` | 路径 | 是 | 申请单 id。 1.蚂蚁门店管理场景：通过 [ ant.merchant.expand.shop.create](https://opendocs.alipay.com/open/05afbc4a_ant.merchant.expand.shop.create?scene=common&pathHash=bf443b73)(蚂蚁门店创建)接口进件时返回的order_id。 2.直付通二级商户进件场景：通过[ant.merchant.expand.indirect.zft.create](https://opendocs.alipay.com/open/028xr0?pathHash=c9ca6d68)(直付通二级商户创建)接口进件时返回的order_id。 |
| `biz_source` | `BizSource` | `*string` | query | 否 | 业务来源 团购门店：APLUS_SHOP |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。

<a id="api-expandshopclose"></a>

### ExpandShopClose — 蚂蚁门店关闭

- 官方接口：[`ant.merchant.expand.shop.close`](https://opendocs.alipay.com/open-v3/4950631b_ant.merchant.expand.shop.close)
- 路由：`PATCH /v3/ant/merchant/shop/close`
- 请求类型：[ExpandShopCloseRequest](merchant_request.go#L345)
- 响应类型：[ExpandShopCloseResponse](merchant_response.go#L229)

通过shop_id，关闭蚂蚁门店。

请求约束：ExpandShopCloseRequest is the official V3 AntMerchantExpandShopCloseModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ext_infos` | `ExtInfos` | [`*[]ShopExtInfo`](#model-shopextinfo) | JSON body | 未标注 | 扩展信息列表。key值需要向对应行业的bd进行申请。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。 |
| `ip_role_id` | `IpRoleId` | `*string` | JSON body | 未标注 | 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。若未传入shop_id 则本参数与store_id均必填。 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 支付宝侧蚂蚁门店 id。传入本参数后可不填 store_id 及 ip_role_id。 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 商户侧门店编号。表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。若未传入 shop_id 则本参数与与ip_role_id均必填。 |

<a id="api-expandshopconsult"></a>

### ExpandShopConsult — 蚂蚁店铺创建咨询

- 官方接口：[`ant.merchant.expand.shop.consult`](https://opendocs.alipay.com/apis/05vda1)
- 路由：`POST /v3/ant/merchant/expand/shop/consult`
- 请求类型：[ExpandShopConsultRequest](merchant_request.go#L135)
- 响应类型：[ExpandShopConsultResponse](merchant_response.go#L103)

蚂蚁店铺创建咨询

请求约束：ExpandShopConsultRequest is the official V3 AntMerchantExpandShopConsultModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_cards` | `BizCards` | [`*[]SettleCardInfo`](#model-settlecardinfo) | JSON body | 未标注 | "门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。" |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id，非加油站等特殊门店无需关注。 |
| `business_address` | `BusinessAddress` | [`*AddressInfo`](#model-addressinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `business_time` | `BusinessTime` | [`*[]ShopBusinessTime`](#model-shopbusinesstime) | JSON body | 未标注 | 店铺经营时间。 |
| `cert_image` | `CertImage` | `*string` | JSON body | 未标注 | 营业执照图片KEY。通过[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id |
| `cert_name` | `CertName` | `*string` | JSON body | 未标注 | 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。 |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 证件号码，请填写店铺营业执照号。 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 证件类型 |
| `contact_infos` | `ContactInfos` | [`*[]ContactInfo`](#model-contactinfo) | JSON body | 未标注 | 联系人信息。如果填写，其中名称必填，手机、固话、email 三选一必填。 |
| `contact_mobile` | `ContactMobile` | `*string` | JSON body | 未标注 | 店铺联系手机，与店铺联系固话二选一必填 |
| `contact_phone` | `ContactPhone` | `*string` | JSON body | 未标注 | 店铺的联系固话，和店铺联系手机二选一必填 |
| `ext_infos` | `ExtInfos` | [`*[]ShopExtInfo`](#model-shopextinfo) | JSON body | 未标注 | 扩展信息列表。key值需要向对应行业的bd进行申请。 |
| `ip_role_id` | `IpRoleId` | `*string` | JSON body | 未标注 | 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。特别说明：IoT设备三绑定场景统一填写商户pid |
| `legal_cert_no` | `LegalCertNo` | `*string` | JSON body | 未标注 | 法人身份证号 |
| `legal_name` | `LegalName` | `*string` | JSON body | 未标注 | 法人名称。 |
| `license_auth_letter_image` | `LicenseAuthLetterImage` | `*string` | JSON body | 未标注 | 营业执照授权函。需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 蚂蚁门店管理场景可选。 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注 |
| `out_door_images` | `OutDoorImages` | `*string` | JSON body | 未标注 | 门头照 id，需传入 [ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id |
| `qualifications` | `Qualifications` | [`*[]IndustryQualificationInfo`](#model-industryqualificationinfo) | JSON body | 未标注 | 行业特殊资质。 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 场景，非加油站等特殊门店无需关注。 |
| `settle_alipay_logon_id` | `SettleAlipayLogonId` | `*string` | JSON body | 未标注 | 结算支付宝账号的登录号，需传入小程序归属账号。若需增加收款账号，可在创建门店后登录[商家中心](https://b.alipay.com/index2.htm) ，在 账户中心 &gt; 经营信息 &gt; 门店管理 中选择对应门店选择 更多 &gt; 编辑 &gt; 新增收款账号。 |
| `shop_category` | `ShopCategory` | `*string` | JSON body | 未标注 | 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 [支付宝门店类目-最新](https://ur.alipay.com/4cmn0n) 表格。 |
| `shop_name` | `ShopName` | `*string` | JSON body | 未标注 | 店铺名称。直连开店要保证全局店铺名称+地址唯一，间连开店要保证服务商pid下店铺名称+地址唯一 |
| `shop_type` | `ShopType` | `*string` | JSON body | 未标注 | 店铺经营类型，01表示直营，02表示加盟 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号 |

<a id="api-expandshopcreate"></a>

### ExpandShopCreate — 蚂蚁门店创建

- 官方接口：[`ant.merchant.expand.shop.create`](https://opendocs.alipay.com/open-v3/99b91969_ant.merchant.expand.shop.create)
- 路由：`POST /v3/ant/merchant/shop`
- 请求类型：[ExpandShopCreateRequest](merchant_request.go#L265)
- 响应类型：[ExpandShopCreateResponse](merchant_response.go#L221)

创建蚂蚁门店

请求约束：ExpandShopCreateRequest is the official V3 AntMerchantExpandShopCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_cards` | `BizCards` | [`*[]SettleCardInfo`](#model-settlecardinfo) | JSON body | 未标注 | 门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。 |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id，非加油站等特殊门店无需关注。通过BP站品牌工具箱可以获取品牌id |
| `business_address` | `BusinessAddress` | [`*AddressInfo`](#model-addressinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `business_time` | `BusinessTime` | [`*[]ShopBusinessTime`](#model-shopbusinesstime) | JSON body | 未标注 | 门店经营时间。 |
| `cert_image` | `CertImage` | `*string` | JSON body | 未标注 | 营业执照图片KEY |
| `cert_name` | `CertName` | `*string` | JSON body | 未标注 | 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。 |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 证件号码，请填写门店营业执照号。 |
| `cert_start_expire_time` | `CertStartExpireTime` | `*string` | JSON body | 未标注 | 营业执照的开始时间 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 证件类型 |
| `contact_infos` | `ContactInfos` | [`*[]ContactInfo`](#model-contactinfo) | JSON body | 未标注 | 联系人信息 |
| `contact_mobile` | `ContactMobile` | `*string` | JSON body | 未标注 | 门店联系手机号 |
| `contact_phone` | `ContactPhone` | `*string` | JSON body | 未标注 | 门店的联系固定电话 |
| `cover` | `Cover` | `*string` | JSON body | 未标注 | 需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 |
| `expire_time` | `ExpireTime` | `*string` | JSON body | 未标注 | 营业执照过期时间 |
| `ext_infos` | `ExtInfos` | [`*[]ShopExtInfo`](#model-shopextinfo) | JSON body | 未标注 | 扩展信息列表。key值需要向对应行业的bd进行申请。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。 |
| `fund_proofs` | `FundProofs` | `*[]string` | JSON body | 未标注 | 需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 |
| `fund_type` | `FundType` | `*string` | JSON body | 未标注 | 资金关系说明 |
| `identity_back_url` | `IdentityBackUrl` | `*string` | JSON body | 未标注 | 需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 |
| `identity_end_time` | `IdentityEndTime` | `*string` | JSON body | 未标注 | 身份证有效结束时间,国徽面信息 |
| `identity_front_url` | `IdentityFrontUrl` | `*string` | JSON body | 未标注 | 需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 |
| `identity_start_time` | `IdentityStartTime` | `*string` | JSON body | 未标注 | 身份证的有效开始时间，国徽面信息 |
| `industry_info` | `IndustryInfo` | [`*[]MerchantShopIndustryInfo`](#model-merchantshopindustryinfo) | JSON body | 未标注 | 行业信息，例如所属集团、特殊标签 |
| `industry_license` | `IndustryLicense` | [`*[]ShopIndustryLicense`](#model-shopindustrylicense) | JSON body | 未标注 | 门店资质信息 |
| `ip_role_id` | `IpRoleId` | `*string` | JSON body | 未标注 | 当前门店归属的商户角色 ID，表示当前进件的门店的归属商户。对于直连开店场景，填写商户 PID。对于间连开店场景（线上、线下、直付通），商户SMID |
| `legal_cert_no` | `LegalCertNo` | `*string` | JSON body | 未标注 | 法人身份证号 |
| `legal_name` | `LegalName` | `*string` | JSON body | 未标注 | 法人名称 |
| `license_auth_letter_image` | `LicenseAuthLetterImage` | `*string` | JSON body | 未标注 | 营业执照授权函图片id，通过 ant.merchant.expand.indirect.image.upload 接口上传图片后得到的 image_id |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注 |
| `new_business_time` | `NewBusinessTime` | [`*[]ComplexBusinessTime`](#model-complexbusinesstime) | JSON body | 未标注 | 门店营业时间，支持一周7天，支持1天内多时间段 |
| `out_door_images` | `OutDoorImages` | `*[]string` | JSON body | 未标注 | 需传入[ant.merchant.expand.indirect.image.upload](https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload) 接口上传图片后得到的 image_id。 |
| `qualifications` | `Qualifications` | [`*[]IndustryQualificationInfo`](#model-industryqualificationinfo) | JSON body | 未标注 | 行业特殊资质。 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 场景，非加油站/酒店等特殊门店无需关注。 |
| `settle_alipay_logon_id` | `SettleAlipayLogonId` | `*string` | JSON body | 未标注 | 门店结算收款的支付宝账号 |
| `shop_category` | `ShopCategory` | `*string` | JSON body | 未标注 | 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 [支付宝门店类目-最新](https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx) 表格。 |
| `shop_main_type` | `ShopMainType` | `*string` | JSON body | 未标注 | 主要针对医疗行业门店主体类型进件使用，不同门店认证主体需要相应的主体资质证书。 企业营业执照：ENTERPRISE; 事业单位法人证书：INST_RGST_CTF; 民办非企业单位登记证书：PRIVATE_NON_ENTERPRISE; 企业（澳门）：AOMEN_ENTERPRISE |
| `shop_name` | `ShopName` | `*string` | JSON body | 未标注 | 门店名称, 由商户定义在支付宝内的门店名称 |
| `shop_type` | `ShopType` | `*string` | JSON body | 未标注 | 门店经营类型，01表示直营，02表示加盟 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号 |

<a id="api-expandshopmodify"></a>

### ExpandShopModify — 修改蚂蚁门店

- 官方接口：[`ant.merchant.expand.shop.modify`](https://opendocs.alipay.com/open-v3/b1a37213_ant.merchant.expand.shop.modify)
- 路由：`PATCH /v3/ant/merchant/shop`
- 请求类型：[ExpandShopModifyRequest](merchant_request.go#L199)
- 响应类型：[ExpandShopModifyResponse](merchant_response.go#L215)

修改蚂蚁门店，按信息项修改。若无特殊说明，如果某项存在但是没填写，则不会覆盖掉原来的值

请求约束：ExpandShopModifyRequest is the official V3 AntMerchantExpandShopModifyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_cards` | `BizCards` | [`*[]SettleCardInfo`](#model-settlecardinfo) | JSON body | 未标注 | "门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。" |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 品牌id。若无需更新本信息项，可以不填写。可以从B/P站品牌工具箱获取品牌id |
| `business_address` | `BusinessAddress` | [`*AddressInfo`](#model-addressinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `business_time` | `BusinessTime` | [`*[]ShopBusinessTime`](#model-shopbusinesstime) | JSON body | 未标注 | 门店经营时间。若无需更新本信息项，可以不填写 |
| `cert_image` | `CertImage` | `*string` | JSON body | 未标注 | 营业执照图片KEY |
| `cert_name` | `CertName` | `*string` | JSON body | 未标注 | 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。 |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 证件号码。请填写门店营业执照号。 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 证件类型 |
| `contact_infos` | `ContactInfos` | [`*ContactInfo`](#model-contactinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `contact_mobile` | `ContactMobile` | `*string` | JSON body | 未标注 | 门店联系手机 |
| `contact_phone` | `ContactPhone` | `*string` | JSON body | 未标注 | 门店的联系固话 |
| `cover` | `Cover` | `*string` | JSON body | 未标注 | 封面图，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写。若填写则会整体覆盖原有的封面图 |
| `expire_time` | `ExpireTime` | `*string` | JSON body | 未标注 | 证件过期时间,当是神券团购门店（ext_infos传了biz_source=APLUS_SHOP）并且license_type=3118或者A521时，必填 |
| `ext_infos` | `ExtInfos` | [`*[]ShopExtInfo`](#model-shopextinfo) | JSON body | 未标注 | 扩展信息列表。key值需要向对应行业的bd进行申请。传入本项时，按key进行更新或补充。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。 |
| `industry_info` | `IndustryInfo` | [`*[]MerchantShopIndustryInfo`](#model-merchantshopindustryinfo) | JSON body | 未标注 | 行业信息，例如所属集团、特殊标签 |
| `industry_license` | `IndustryLicense` | [`*[]ShopIndustryLicense`](#model-shopindustrylicense) | JSON body | 未标注 | 门店资质信息 |
| `ip_role_id` | `IpRoleId` | `*string` | JSON body | 未标注 | 商户角色id，表示将要修改的门店属于哪个商户角色，和store_id一起定位待修改的门店。对于直连场景，填写商户pid；对于间连场景（线上、线下、直付通），填写商户smid。本信息项不可修改 |
| `legal_cert_no` | `LegalCertNo` | `*string` | JSON body | 未标注 | 法人身份证号。如果要更新，必须和legal_name成对传入 |
| `legal_name` | `LegalName` | `*string` | JSON body | 未标注 | 法人名称。如果要更新，必须和legal_cert_no成对传入 |
| `license_auth_letter_image` | `LicenseAuthLetterImage` | `*string` | JSON body | 未标注 | 营业执照授权函。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注 |
| `new_business_time` | `NewBusinessTime` | [`*[]ComplexBusinessTime`](#model-complexbusinesstime) | JSON body | 未标注 | 门店营业时间，支持一周7天，支持1天内多时间段 |
| `out_door_images` | `OutDoorImages` | `*[]string` | JSON body | 未标注 | 门头照，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写。若填写则会整体覆盖原有的门头照列表 |
| `qualifications` | `Qualifications` | [`*[]IndustryQualificationInfo`](#model-industryqualificationinfo) | JSON body | 未标注 | 行业特殊资质。若无需更新本信息项，可以不填写。填写会按类型进行更新或补充 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 场景。若无需更新本信息项，可以不填写 |
| `settle_alipay_logon_id` | `SettleAlipayLogonId` | `*string` | JSON body | 未标注 | 结算支付宝账号的登录号。若无需更新本信息项，可以不填写 |
| `shop_category` | `ShopCategory` | `*string` | JSON body | 未标注 | 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 [支付宝门店类目-最新](https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx) 表格。 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 蚂蚁门店id，用于查出待修改的门店。填写本参数的话，store_id和ip_role_id可以不填。本信息项不可修改 |
| `shop_main_type` | `ShopMainType` | `*string` | JSON body | 未标注 | 主要针对医疗行业门店主体类型进件使用，不同门店认证主体需要相应的主体资质证书。 企业营业执照：ENTERPRISE; 事业单位法人证书：INST_RGST_CTF; 民办非企业单位登记证书：PRIVATE_NON_ENTERPRISE; 企业（澳门）：AOMEN_ENTERPRISE |
| `shop_name` | `ShopName` | `*string` | JSON body | 未标注 | 门店名称。直连开店要保证全局门店名称+地址唯一，间连开店要保证服务商pid下门店名称+地址唯一。若无需更新本信息项，可以不填写 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。用于查出待修改的门店。 |

<a id="api-expandshoppagequery"></a>

### ExpandShopPageQuery — 店铺分页查询接口

- 官方接口：[`ant.merchant.expand.shop.page.query`](https://opendocs.alipay.com/open-v3/fe443839_ant.merchant.expand.shop.page.query)
- 路由：`GET /v3/ant/merchant/shop/pagequery`
- 请求类型：[ExpandShopPageQueryRequest](merchant_request.go#L473)
- 响应类型：[ExpandShopPageQueryResponse](merchant_response.go#L233)

用于服务商或商户查询其自己的店铺信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ip_role_id` | `IpRoleId` | `*string` | query | 否 | 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。特别说明：IoT设备三绑定场景统一填写商户pid |
| `page_num` | `PageNum` | `*int64` | query | 否 | 查询页数 |
| `page_size` | `PageSize` | `*int64` | query | 否 | 每页查询大小，限制100以内 |
| `address_version` | `AddressVersion` | `*string` | query | 否 | 行政区划版本，当前可传空值(取默认版本)、2022Q2、UPTODATE(取最新版本)，其中空值默认为：2020Q1版本(address_version=''或null)，想要查看版本是2022年2季度版本则传入:(address_version='2022Q2')，想要获取最新版本则传入:(address_version ='UPTODATE') |
| `province_code` | `ProvinceCode` | `*string` | query | 否 | 省份编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |
| `city_code` | `CityCode` | `*string` | query | 否 | 城市编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |
| `district_code` | `DistrictCode` | `*string` | query | 否 | 区县编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |

<a id="api-expandshopquery"></a>

### ExpandShopQuery — 门店查询接口

- 官方接口：[`ant.merchant.expand.shop.query`](https://opendocs.alipay.com/open-v3/b80b5948_ant.merchant.expand.shop.query)
- 路由：`GET /v3/ant/merchant/shop`
- 请求类型：[ExpandShopQueryRequest](merchant_request.go#L451)
- 响应类型：[ExpandShopQueryResponse](merchant_response.go#L143)

用于服务商或商户查询其自己的门店信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `shop_id` | `ShopId` | `*string` | query | 否 | 蚂蚁门店id |
| `store_id` | `StoreId` | `*string` | query | 否 | 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号 |
| `ip_role_id` | `IpRoleId` | `*string` | query | 否 | 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid |
| `address_version` | `AddressVersion` | `*string` | query | 否 | 行政区划版本，当前可传空值(取默认版本)、2022Q2、UPTODATE(取最新版本)，其中空值默认为：2020Q1版本（ address_version=''或null），想要查看版本是2022年2季度版本则传入:(address_version='2022Q2')，想要获取最新版本则传入:(address_version ='UPTODATE') |
| `need_recommend` | `NeedRecommend` | `*string` | query | 否 | 门店不置信时，是否需要返回shop_recommend_info |
| `need_industry_info` | `NeedIndustryInfo` | `*string` | query | 否 | need_industry_info=0时不返回行业信息；need_industry_info=1时返回不需要审核的行业信息、审核通过的行业信息 |
| `need_industry_license` | `NeedIndustryLicense` | `*string` | query | 否 | need_industry_license=0时不返回行业资质；need_industry_license=1时返回审核通过的行业资质 |
| `shop_type` | `ShopType` | `*string` | query | 否 | 门店经营类型，01表示直营，02表示加盟 |
| `biz_source` | `BizSource` | `*string` | query | 否 | 团购门店业务类型来源 |

<a id="api-expandshopreceiptaccountsave"></a>

### ExpandShopReceiptAccountSave — 店铺增加收单账号

- 官方接口：[`ant.merchant.expand.shop.receiptaccount.save`](https://opendocs.alipay.com/open-v3/74728d59_ant.merchant.expand.shop.receiptaccount.save)
- 路由：`POST /v3/ant/merchant/expand/shop/receiptaccount/save`
- 请求类型：[ExpandShopReceiptAccountSaveRequest](merchant_request.go#L189)
- 响应类型：[ExpandShopReceiptAccountSaveResponse](merchant_response.go#L113)

店铺增加收单账号

请求约束：ExpandShopReceiptAccountSaveRequest is the official V3 AntMerchantExpandShopReceiptaccountSaveModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `promise` | `Promise` | `*string` | JSON body | 未标注 | 是否承诺收单账号信息准确。具体承诺信息可查看 [门店信息承诺函](https://gw.alipayobjects.com/os/bmw-prod/922bafa8-a712-4f79-aa32-6f08d7359a5c.docx)。 Y 是 N 否 |
| `receipt_account_id` | `ReceiptAccountId` | `*string` | JSON body | 未标注 | 收单账号 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 店铺ID |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-addressinfo"></a>

### AddressInfo

类型定义：[AddressInfo](merchant_model.go#L237)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 地址。商户详细经营地址或人员所在地点 |
| `city_code` | `CityCode` | `*string` | 未标注 | 城市编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |
| `district_code` | `DistrictCode` | `*string` | 未标注 | 区县编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |
| `latitude` | `Latitude` | `*string` | 未标注 | 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。[高德经纬度查询](http://lbs.amap.com/console/show/picker) |
| `longitude` | `Longitude` | `*string` | 未标注 | 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。[高德经纬度查询](http://lbs.amap.com/console/show/picker) |
| `poiid` | `Poiid` | `*string` | 未标注 | 高德poiid |
| `province_code` | `ProvinceCode` | `*string` | 未标注 | 省份编码。 蚂蚁店铺请按照[蚂蚁店铺地区码](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 表格中填写。 直付通商户请按照[直付通商户地区码](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx ) 表格中内容填写。 |
| `type` | `Type` | `*string` | 未标注 | 地址类型。取值范围：BUSINESS_ADDRESS：经营地址（默认） |

<a id="model-complexbusinesstime"></a>

### ComplexBusinessTime

类型定义：[ComplexBusinessTime](merchant_model.go#L339)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_period` | `TimePeriod` | [`*[]ComplexBusinessTimePeriod`](#model-complexbusinesstimeperiod) | 未标注 | 营业时间段，支持多个时间段，例如08:00-11:00，14:00-16:00 |
| `week_day` | `WeekDay` | `*int64` | 未标注 | 本对象表示周几的营业时间。1~6表示周一到周六，7表示周日 |

<a id="model-complexbusinesstimeperiod"></a>

### ComplexBusinessTimePeriod

类型定义：[ComplexBusinessTimePeriod](merchant_model.go#L459)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `close_time` | `CloseTime` | `*string` | 未标注 | 关门时间 格式：HH:mm |
| `open_time` | `OpenTime` | `*string` | 未标注 | 开门时间 格式：HH:mm |

<a id="model-contactinfo"></a>

### ContactInfo

类型定义：[ContactInfo](merchant_model.go#L267)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `email` | `Email` | `*string` | 未标注 | 电子邮箱 |
| `id_card_no` | `IdCardNo` | `*string` | 未标注 | 身份证号 |
| `mobile` | `Mobile` | `*string` | 未标注 | 手机号 |
| `name` | `Name` | `*string` | 未标注 | 联系人名字 |
| `phone` | `Phone` | `*string` | 未标注 | 电话 |
| `tag` | `Tag` | `*[]string` | 未标注 | 商户联系人业务标识枚举，表示商户联系人的职责 |
| `type` | `Type` | `*string` | 未标注 | 联系人类型，取值范围：LEGAL_PERSON：法人；CONTROLLER：实际控制人；AGENT：代理人；OTHER：其他 |

<a id="model-expandindirectimageupload"></a>

### ExpandIndirectImageUpload

类型定义：[ExpandIndirectImageUpload](merchant_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `image_type` | `ImageType` | `*string` | 未标注 | 图片格式，支持格式：bmp、jpg、jpeg、png、gif. |

<a id="model-industryqualificationinfo"></a>

### IndustryQualificationInfo

类型定义：[IndustryQualificationInfo](merchant_model.go#L293)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `industry_qualification_image` | `IndustryQualificationImage` | `*string` | 未标注 | 商户行业资质图片。其值为通过ant.merchant.expand.indirect.image.upload上传图片得到的image_id |
| `industry_qualification_type` | `IndustryQualificationType` | `*string` | 未标注 | [商户行业资质类型，具体选值参见文档](https://gw.alipayobjects.com/os/bmw-prod/7aa3a36b-2bc2-4d57-815f-08edd55ef67e.xlsx) |

<a id="model-itemextinfo"></a>

### ItemExtInfo

类型定义：[ItemExtInfo](merchant_model.go#L119)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_key` | `ExtKey` | `*string` | 未标注 | 扩展信息的key |
| `ext_value` | `ExtValue` | `*string` | 未标注 | 扩展信息的值 |

<a id="model-itemlabelcreateinfo"></a>

### ItemLabelCreateInfo

类型定义：[ItemLabelCreateInfo](merchant_model.go#L33)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `label_key` | `LabelKey` | `*string` | 未标注 | 标签键<br> OIL_NUM：油号 |
| `label_value` | `LabelValue` | `*string` | 未标注 | 标签值：<br> OIL_NUM对应的值为：95# 92# 98# 等等 |

<a id="model-itemlabelmodifyinfo"></a>

### ItemLabelModifyInfo

类型定义：[ItemLabelModifyInfo](merchant_model.go#L57)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `label_key` | `LabelKey` | `*string` | 未标注 | 标签键<br> OIL_NUM：油号 |
| `label_value` | `LabelValue` | `*string` | 未标注 | 标签值：<br> OIL_NUM对应的值为：95# 92# 98# 等等 |

<a id="model-itempropertyinfo"></a>

### ItemPropertyInfo

类型定义：[ItemPropertyInfo](merchant_model.go#L135)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `property_key` | `PropertyKey` | `*string` | 未标注 | 属性名称 |
| `property_value_list` | `PropertyValueList` | `*[]string` | 未标注 | 属性值列表 |

<a id="model-itemskucreateinfo"></a>

### ItemSkuCreateInfo

类型定义：[ItemSkuCreateInfo](merchant_model.go#L41)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cost_price` | `CostPrice` | `*int64` | 未标注 | 成本价，单位分 |
| `ext_info` | `ExtInfo` | [`*[]ItemExtInfo`](#model-itemextinfo) | 未标注 | SKU扩展信息 |
| `external_sku_id` | `ExternalSkuId` | `*string` | 未标注 | 外部SKU ID |
| `inventory` | `Inventory` | `*int64` | 未标注 | 库存 |
| `original_price` | `OriginalPrice` | `*int64` | 未标注 | 标价，单位分 |
| `price` | `Price` | `*int64` | 未标注 | 售价，单位分 |

<a id="model-itemskumodifyinfo"></a>

### ItemSkuModifyInfo

类型定义：[ItemSkuModifyInfo](merchant_model.go#L65)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cost_price` | `CostPrice` | `*int64` | 未标注 | 成本价，单位分 |
| `ext_info` | `ExtInfo` | [`*[]ItemExtInfo`](#model-itemextinfo) | 未标注 | SKU扩展信息 |
| `external_sku_id` | `ExternalSkuId` | `*string` | 未标注 | 外部skuId（sku_id不为空时作为普通更新项,sku_id为空时作为更新key,sku_id和该字段同时为空则新增该条数据） |
| `inventory` | `Inventory` | `*int64` | 未标注 | 库存 |
| `original_price` | `OriginalPrice` | `*int64` | 未标注 | 标价，单位分 |
| `price` | `Price` | `*int64` | 未标注 | 售价，单位分 |
| `sku_id` | `SkuId` | `*string` | 未标注 | skuId |
| `status` | `Status` | `*string` | 未标注 | 商品SKU状态：EFFECT、INVALID |

<a id="model-itemskupropertyinfo"></a>

### ItemSkuPropertyInfo

类型定义：[ItemSkuPropertyInfo](merchant_model.go#L429)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `property_key` | `PropertyKey` | `*string` | 未标注 | 属性 |
| `property_value` | `PropertyValue` | `*string` | 未标注 | 属性值 |

<a id="model-materialcreateinfo"></a>

### MaterialCreateInfo

类型定义：[MaterialCreateInfo](merchant_model.go#L127)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `content` | `Content` | `*string` | 未标注 | 素材内容（素材地址或素材KEY） |
| `type` | `Type` | `*string` | 未标注 | 素材类型 |

<a id="model-materialmodifyinfo"></a>

### MaterialModifyInfo

类型定义：[MaterialModifyInfo](merchant_model.go#L157)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `content` | `Content` | `*string` | 未标注 | 素材内容（素材地址或素材KEY） |
| `type` | `Type` | `*string` | 未标注 | 素材类型 |

<a id="model-merchantshopindustryinfo"></a>

### MerchantShopIndustryInfo

类型定义：[MerchantShopIndustryInfo](merchant_model.go#L319)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `info_code` | `InfoCode` | `*string` | 未标注 | 行业信息CODE |
| `info_value` | `InfoValue` | `*[]string` | 未标注 | 行业信息值，具体值根据info_code决定。 当info_code=SHOP_MEDICAL_LOGO/SHOP_MEDICAL_IN_PIC这类文件信息时， 需要传入ant.merchant.expand.indirect.image.upload接口上传的文件key；当info_code= SHOP_MEDICAL_LEVEL_TAG时，需要[医院类型](https://mdn.alipayobjects.com/mrcheshop/afts/file/ZuDmQKgBifgAAAAAAAAAAAAADiWmAQBr?af_filename=医院类型.xlsx)中CODE；其他场景传入自定义文本即可。 |

<a id="model-settlecardinfo"></a>

### SettleCardInfo

类型定义：[SettleCardInfo](merchant_model.go#L213)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `account_branch_name` | `AccountBranchName` | `*string` | 未标注 | 开户支行名 |
| `account_holder_name` | `AccountHolderName` | `*string` | 未标注 | 卡户名 |
| `account_inst_city` | `AccountInstCity` | `*string` | 未标注 | 开户行所在地-市 |
| `account_inst_id` | `AccountInstId` | `*string` | 未标注 | 开户行简称缩写 |
| `account_inst_name` | `AccountInstName` | `*string` | 未标注 | 银行名称 |
| `account_inst_province` | `AccountInstProvince` | `*string` | 未标注 | 开户行所在地-省 |
| `account_no` | `AccountNo` | `*string` | 未标注 | 银行卡号 |
| `account_type` | `AccountType` | `*string` | 未标注 | 卡类型<br> 借记卡-DC<br> 信用卡-CC |
| `bank_code` | `BankCode` | `*string` | 未标注 | 联行号，详细需调用[‘联行号关联分支行查询’](https://opendocs.alipay.com/apis/0853v1)API查看查询结果. |
| `usage_type` | `UsageType` | `*string` | 未标注 | 账号使用类型<br> 对公-01<br> 对私-02 |

<a id="model-shopbusinesstime"></a>

### ShopBusinessTime

类型定义：[ShopBusinessTime](merchant_model.go#L257)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `close_time` | `CloseTime` | `*string` | 未标注 | 关门时间 格式：HH:mm |
| `open_time` | `OpenTime` | `*string` | 未标注 | 开门时间 格式：HH:mm |
| `week_day` | `WeekDay` | `*int64` | 未标注 | 本对象表示周几的营业时间。1~6表示周一到周六，7表示周日 |

<a id="model-shopextinfo"></a>

### ShopExtInfo

类型定义：[ShopExtInfo](merchant_model.go#L285)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `key_name` | `KeyName` | `*string` | 未标注 | KV数据对的key，表示该kv对象表示什么含义。 |
| `value` | `Value` | `*string` | 未标注 | kv对象的值 |

<a id="model-shopindustrylicense"></a>

### ShopIndustryLicense

类型定义：[ShopIndustryLicense](merchant_model.go#L327)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `expire_time` | `ExpireTime` | `*string` | 未标注 | 证件过期时间 |
| `license_no` | `LicenseNo` | `*string` | 未标注 | 资质证件编号 |
| `license_pic` | `LicensePic` | `*string` | 未标注 | 资质图片信息 |
| `license_type` | `LicenseType` | `*string` | 未标注 | 行业资质类型，需使用[资质列表](https://mdn.alipayobjects.com/mrcheshop/afts/file/mS30TKabvCMAAAAAQKAAAAgAeiWmAQBr?af_fileName=%E8%B5%84%E8%B4%A8%E5%88%97%E8%A1%A8.xlsx)中类型 |

<a id="model-skucreateinfo"></a>

### SkuCreateInfo

类型定义：[SkuCreateInfo](merchant_model.go#L143)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `inventory` | `Inventory` | `*int64` | 未标注 | 库存 |
| `material_list` | `MaterialList` | [`*[]MaterialCreateInfo`](#model-materialcreateinfo) | 未标注 | SKU素材列表（最多3个） |
| `original_price` | `OriginalPrice` | `*int64` | 未标注 | 标价，单位分 |
| `price` | `Price` | `*int64` | 未标注 | 售价，单位分 |
| `property_list` | `PropertyList` | [`*[]ItemSkuPropertyInfo`](#model-itemskupropertyinfo) | 未标注 | SKU属性列表 |

<a id="model-skumodifyinfo"></a>

### SkuModifyInfo

类型定义：[SkuModifyInfo](merchant_model.go#L165)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `inventory` | `Inventory` | `*int64` | 未标注 | 库存 |
| `material_list` | `MaterialList` | [`*[]MaterialModifyInfo`](#model-materialmodifyinfo) | 未标注 | SKU素材列表（会和SKU已存在素材做差异化比较后做增删改操作） |
| `original_price` | `OriginalPrice` | `*int64` | 未标注 | 标价，单位分 |
| `price` | `Price` | `*int64` | 未标注 | 售价，单位分 |
| `property_list` | `PropertyList` | [`*[]ItemSkuPropertyInfo`](#model-itemskupropertyinfo) | 未标注 | SKU属性列表（覆盖SKU已存在属性） |
| `sku_id` | `SkuId` | `*string` | 未标注 | SKU ID（传空表示新增对应SKU） |
