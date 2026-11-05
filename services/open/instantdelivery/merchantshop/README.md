# merchantshop 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/instantdelivery/merchantshop`。本包提供 4 个支付宝 V3 REST 操作。

通过 `merchantshop.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`InstantDeliveryMerchantShopBatchQuery`](#api-instantdeliverymerchantshopbatchquery) | `alipay.open.instantdelivery.merchantshop.batchquery` |
| [`InstantDeliveryMerchantShopCreate`](#api-instantdeliverymerchantshopcreate) | `alipay.open.instantdelivery.merchantshop.create` |
| [`InstantDeliveryMerchantShopModify`](#api-instantdeliverymerchantshopmodify) | `alipay.open.instantdelivery.merchantshop.modify` |
| [`InstantDeliveryMerchantShopQuery`](#api-instantdeliverymerchantshopquery) | `alipay.open.instantdelivery.merchantshop.query` |

## 使用示例

- [InstantDeliveryMerchantShopBatchQuery](api_merchantshop_instantdeliverymerchantshopbatchquery_example_test.go)
- [InstantDeliveryMerchantShopCreate](api_merchantshop_instantdeliverymerchantshopcreate_example_test.go)
- [InstantDeliveryMerchantShopModify](api_merchantshop_instantdeliverymerchantshopmodify_example_test.go)
- [InstantDeliveryMerchantShopQuery](api_merchantshop_instantdeliverymerchantshopquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-instantdeliverymerchantshopbatchquery"></a>

### InstantDeliveryMerchantShopBatchQuery — 即时配送商家门店分页查询

- 官方接口：[`alipay.open.instantdelivery.merchantshop.batchquery`](https://opendocs.alipay.com/pre-apis/05vn66)
- 路由：`POST /v3/alipay/open/instantdelivery/merchantshop/batchquery`
- 请求类型：[InstantDeliveryMerchantShopBatchQueryRequest](merchantshop_request.go#L4)
- 响应类型：[InstantDeliveryMerchantShopBatchQueryResponse](merchantshop_response.go#L4)

即时配送场景，商家门店信息分页查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_no` | `PageNo` | `*int64` | JSON body | 未标注 | 当前页码 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 分页数量, 最大50。 |
| `shop_name` | `ShopName` | `*string` | JSON body | 未标注 | 门店名称 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编码。 |

<a id="api-instantdeliverymerchantshopcreate"></a>

### InstantDeliveryMerchantShopCreate — 即时配送商家门店创建

- 官方接口：[`alipay.open.instantdelivery.merchantshop.create`](https://opendocs.alipay.com/pre-apis/05vluh)
- 路由：`POST /v3/alipay/open/instantdelivery/merchantshop/create`
- 请求类型：[InstantDeliveryMerchantShopCreateRequest](merchantshop_request.go#L16)
- 响应类型：[InstantDeliveryMerchantShopCreateResponse](merchantshop_response.go#L16)

帮商家在即时配送公司创建门店。名称+地址必须全局唯一；pid下的shopNo必须唯一。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `contact_name` | `ContactName` | `*string` | JSON body | 未标注 | 联系人姓名 |
| `detail_address` | `DetailAddress` | `*string` | JSON body | 未标注 | 地址。商户详细经营地址或人员所在地点 |
| `enterprise_city` | `EnterpriseCity` | `*string` | JSON body | 未标注 | 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `enterprise_district` | `EnterpriseDistrict` | `*string` | JSON body | 未标注 | 区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `enterprise_province` | `EnterpriseProvince` | `*string` | JSON body | 未标注 | 省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号 |
| `phone` | `Phone` | `*string` | JSON body | 未标注 | 联系人电话/手机号 |
| `poiid` | `Poiid` | `*string` | JSON body | 未标注 | 高德poiid |
| `shop_category` | `ShopCategory` | `*string` | JSON body | 未标注 | 店铺类目，取值参见文件https://mif-pub.alipayobjects.com/ShopCategory.xlsx 中的三级门店类目 |
| `shop_name` | `ShopName` | `*string` | JSON body | 未标注 | 门店名称，最长不超过256个字符。名称+地址需要全局唯一。 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编码，可自定义，但必须唯一 |

<a id="api-instantdeliverymerchantshopmodify"></a>

### InstantDeliveryMerchantShopModify — 即时配送商家门店更新

- 官方接口：[`alipay.open.instantdelivery.merchantshop.modify`](https://opendocs.alipay.com/pre-apis/068gkz)
- 路由：`POST /v3/alipay/open/instantdelivery/merchantshop/modify`
- 请求类型：[InstantDeliveryMerchantShopModifyRequest](merchantshop_request.go#L46)
- 响应类型：[InstantDeliveryMerchantShopModifyResponse](merchantshop_response.go#L22)

帮商家更新在即时配送公司的门店

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `contact_name` | `ContactName` | `*string` | JSON body | 未标注 | 联系人姓名 |
| `logistics_codes` | `LogisticsCodes` | `*[]string` | JSON body | 未标注 | 要更新的配送公司列表，必须包含已经创建过该门店的配送公司，可以新增开通了账户的配送公司 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号 |
| `phone` | `Phone` | `*string` | JSON body | 未标注 | 联系人电话/手机号 |
| `shop_no` | `ShopNo` | `*string` | JSON body | 未标注 | 商家门店编码 |

<a id="api-instantdeliverymerchantshopquery"></a>

### InstantDeliveryMerchantShopQuery — 即时配送商家门店详情查询

- 官方接口：[`alipay.open.instantdelivery.merchantshop.query`](https://opendocs.alipay.com/pre-apis/068ja3)
- 路由：`GET /v3/alipay/open/instantdelivery/merchantshop/query`
- 请求类型：[InstantDeliveryMerchantShopQueryRequest](merchantshop_request.go#L60)
- 响应类型：[InstantDeliveryMerchantShopQueryResponse](merchantshop_response.go#L28)

即时配送场景，查询商家门店详情信息。

请求约束：InstantDeliveryMerchantShopQueryRequest is the request type for alipay.open.instantdelivery.merchantshop.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `shop_no` | `ShopNo` | `*string` | query | 否 | 商家门店编码。 |
