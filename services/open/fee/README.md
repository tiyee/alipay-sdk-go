# open/fee 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/fee`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `fee.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_fee_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`FeeAdjustApply`](#api-feeadjustapply) | [`alipay.open.fee.adjust.apply`](https://opendocs.alipay.com/open-v3/fe77ed62_alipay.open.fee.adjust.apply) | 特殊费率申请 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-feeadjustapply"></a>

### FeeAdjustApply — 特殊费率申请

- 官方接口：[`alipay.open.fee.adjust.apply`](https://opendocs.alipay.com/open-v3/fe77ed62_alipay.open.fee.adjust.apply)
- 路由：`POST /v3/alipay/open/fee/adjust/apply`
- 请求类型：[FeeAdjustApplyRequest](fee_request.go#L9)
- 响应类型：[FeeAdjustApplyResponse](fee_response.go#L5)

服务商替代签约或交易传参的商户申请费率

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `attachment` | `Attachment` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `cert_pic` | `CertPic` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*FeeAdjustApply`](#model-feeadjustapply) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `shop_scene_pic` | `ShopScenePic` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `shop_sign_pic` | `ShopSignPic` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-feeadjustapply"></a>

### FeeAdjustApply

类型定义：[FeeAdjustApply](fee_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `account` | `Account` | `*string` | 未标注 | 服务商代申请特殊费率的商家支付宝账户 |
| `application_fee` | `ApplicationFee` | `*string` | 未标注 | 服务商代商家申请特殊费率的费率值（%），如0.38表示0.38% |
| `cert_date_limitation` | `CertDateLimitation` | `*string` | 未标注 | 证件有效期，格式：yyyy-MM-dd，长期有效传入：9999-12-31 |
| `cert_no` | `CertNo` | `*string` | 未标注 | 对应证件类型（certType）的证件号，如营业执照为营业执照号 |
| `cert_type` | `CertType` | `*string` | 未标注 | 证件类型，枚举值为 餐饮服务许可证 FOOD_SERVICE_LICENSE， 食品卫生许可证 FOOD_HEALTH_LICENSE， 食品经营许可证 FOOD_BUSINESS_LICENSE， 食品流通许可证 FOOD_CIRCULATE_LICENSE， 食品生产许可证 FOOD_PRODUCTION_LICENS， 营业执照 DEFAULT_LICENSE， 烟草资质 TOBACCO_PRODUCTION_LICENSE |
| `city_code` | `CityCode` | `*string` | 未标注 | 市 |
| `detail_address` | `DetailAddress` | `*string` | 未标注 | 详细地址 |
| `district_code` | `DistrictCode` | `*string` | 未标注 | 区/县 |
| `product_code` | `ProductCode` | `*string` | 未标注 | 服务商代商家申请调整费率的产品码 |
| `province_code` | `ProvinceCode` | `*string` | 未标注 | 省份 |
