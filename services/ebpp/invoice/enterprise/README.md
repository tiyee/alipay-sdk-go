# ebpp/invoice/enterprise 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/enterprise`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `enterprise.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_enterprise_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceEnterpriseMerchantRelationCreate`](#api-invoiceenterprisemerchantrelationcreate) | [`alipay.ebpp.invoice.enterprise.merchantrelation.create`](https://opendocs.alipay.com/apis/05psjg) | 创建企业商户关系 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceenterprisemerchantrelationcreate"></a>

### InvoiceEnterpriseMerchantRelationCreate — 创建企业商户关系

- 官方接口：[`alipay.ebpp.invoice.enterprise.merchantrelation.create`](https://opendocs.alipay.com/apis/05psjg)
- 路由：`POST /v3/alipay/ebpp/invoice/enterprise/merchantrelation/create`
- 请求类型：[InvoiceEnterpriseMerchantRelationCreateRequest](enterprise_request.go#L5)
- 响应类型：[InvoiceEnterpriseMerchantRelationCreateResponse](enterprise_response.go#L5)

根据交易流水号创建企业商户绑定关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业ID |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `operator_id` | `OperatorId` | `*string` | JSON body | 未标注 | 传入操作人员标识 |
| `shop_info` | `ShopInfo` | [`*EcShopInfo`](#model-ecshopinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 交易流水号 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-ecshopinfo"></a>

### EcShopInfo

类型定义：[EcShopInfo](enterprise_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | 未标注 | 门店地址 |
| `city_id` | `CityId` | `*string` | 未标注 | 城市编码 |
| `city_name` | `CityName` | `*string` | 未标注 | 城市名称 |
| `district_id` | `DistrictId` | `*string` | 未标注 | 区编码 |
| `district_name` | `DistrictName` | `*string` | 未标注 | 区名称 |
| `latitude` | `Latitude` | `*string` | 未标注 | 纬度 |
| `longitude` | `Longitude` | `*string` | 未标注 | 经度 |
| `mcc_code_1` | `MccCode1` | `*string` | 未标注 | 一级mcc编码 |
| `mcc_code_2` | `MccCode2` | `*string` | 未标注 | 二级mcc编码 |
| `mcc_name_1` | `MccName1` | `*string` | 未标注 | 一级mcc名称 |
| `mcc_name_2` | `MccName2` | `*string` | 未标注 | 二级mcc名称 |
| `outdoor_img_url` | `OutdoorImgUrl` | `*string` | 未标注 | 门头照地址 |
| `poi_id` | `PoiId` | `*string` | 未标注 | 高德poiId |
| `province_id` | `ProvinceId` | `*string` | 未标注 | 省份编码 |
| `province_name` | `ProvinceName` | `*string` | 未标注 | 省份名称 |
| `shop_name` | `ShopName` | `*string` | 未标注 | 门店名称 |
