# open/mini/miniapp 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/miniapp`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `miniapp.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniMiniAppBrandCancel`](#api-miniminiappbrandcancel) | [`alipay.open.mini.miniapp.brand.cancel`](https://opendocs.alipay.com/open-v3/ac02eafe_alipay.open.mini.miniapp.brand.cancel) | 小程序品牌提交认证后取消品牌认证 | `POST` |
| [`MiniMiniAppBrandQuery`](#api-miniminiappbrandquery) | [`alipay.open.mini.miniapp.brand.query`](https://opendocs.alipay.com/open-v3/6187351c_alipay.open.mini.miniapp.brand.query) | 小程序品牌提交认证后查询品牌审核结果以及商户已有品牌 | `GET` |
| [`MiniMiniAppBrandSubmit`](#api-miniminiappbrandsubmit) | [`alipay.open.mini.miniapp.brand.submit`](https://opendocs.alipay.com/open-v3/24509d98_alipay.open.mini.miniapp.brand.submit) | 小程序品牌提交认证 | `POST` |
| [`MiniMiniAppBrandUpload`](#api-miniminiappbrandupload) | [`alipay.open.mini.miniapp.brand.upload`](https://opendocs.alipay.com/open-v3/e77e0097_alipay.open.mini.miniapp.brand.upload) | 小程序品牌提交认证时上传用户资质以及身份信息 | `POST` |
| [`MiniMiniAppServiceConfigModify`](#api-miniminiappserviceconfigmodify) | [`alipay.open.mini.miniapp.serviceconfig.modify`](https://opendocs.alipay.com/open-v3/95e9abe6_alipay.open.mini.miniapp.serviceconfig.modify) | 小程序设置客服方式 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniminiappbrandcancel"></a>

### MiniMiniAppBrandCancel — 小程序品牌提交认证后取消品牌认证

- 官方接口：[`alipay.open.mini.miniapp.brand.cancel`](https://opendocs.alipay.com/open-v3/ac02eafe_alipay.open.mini.miniapp.brand.cancel)
- 路由：`POST /v3/alipay/open/mini/miniapp/brand/cancel`
- 请求类型：[MiniMiniAppBrandCancelRequest](miniapp_request.go#L36)
- 响应类型：[MiniMiniAppBrandCancelResponse](miniapp_response.go#L5)

小程序品牌提交认证后取消品牌认证

请求约束：MiniMiniAppBrandCancelRequest is the request type for alipay.open.mini.miniapp.brand.cancel.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-miniminiappbrandquery"></a>

### MiniMiniAppBrandQuery — 小程序品牌提交认证后查询品牌审核结果以及商户已有品牌

- 官方接口：[`alipay.open.mini.miniapp.brand.query`](https://opendocs.alipay.com/open-v3/6187351c_alipay.open.mini.miniapp.brand.query)
- 路由：`GET /v3/alipay/open/mini/miniapp/brand/query`
- 请求类型：[MiniMiniAppBrandQueryRequest](miniapp_request.go#L39)
- 响应类型：[MiniMiniAppBrandQueryResponse](miniapp_response.go#L9)

小程序品牌提交认证后查询品牌审核结果以及商户已有品牌

请求约束：MiniMiniAppBrandQueryRequest is the request type for alipay.open.mini.miniapp.brand.query.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-miniminiappbrandsubmit"></a>

### MiniMiniAppBrandSubmit — 小程序品牌提交认证

- 官方接口：[`alipay.open.mini.miniapp.brand.submit`](https://opendocs.alipay.com/open-v3/24509d98_alipay.open.mini.miniapp.brand.submit)
- 路由：`POST /v3/alipay/open/mini/miniapp/brand/submit`
- 请求类型：[MiniMiniAppBrandSubmitRequest](miniapp_request.go#L9)
- 响应类型：[MiniMiniAppBrandSubmitResponse](miniapp_response.go#L15)

小程序品牌提交认证

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `apply_type` | `ApplyType` | `*string` | JSON body | 未标注 | 申请的资质类型 SELF_BRAND_APPLY 自有品牌申请(当需要一个新的品牌时选择此项) AUTHORIZE_BRAND_APPLY 授权品牌申请(当需要授权一个已有的品牌给商户时选择此项) EXIST_BRAND_BIND 品牌绑定，适用于品牌已申请成功的情况(当需要将已有品牌绑定到无品牌小程序时选择此项) |
| `authorize_info` | `AuthorizeInfo` | [`*AuthorizeInfo`](#model-authorizeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `brand_id` | `BrandId` | `*string` | JSON body | 未标注 | 商户已有品牌Id(只有申请的资质类型为EXIST_BRAND_BIND 时需要填写) |
| `brand_name` | `BrandName` | `*string` | JSON body | 未标注 | 品牌名称 |
| `brand_registration_info` | `BrandRegistrationInfo` | [`*BrandRegistrationInfo`](#model-brandregistrationinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `id_materials` | `IdMaterials` | `*[]string` | JSON body | 未标注 | 申请人身份证明材料(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片，可以上传多个身份证明材料) |

<a id="api-miniminiappbrandupload"></a>

### MiniMiniAppBrandUpload — 小程序品牌提交认证时上传用户资质以及身份信息

- 官方接口：[`alipay.open.mini.miniapp.brand.upload`](https://opendocs.alipay.com/open-v3/e77e0097_alipay.open.mini.miniapp.brand.upload)
- 路由：`POST /v3/alipay/open/mini/miniapp/brand/upload`
- 请求类型：[MiniMiniAppBrandUploadRequest](miniapp_request.go#L31)
- 响应类型：[MiniMiniAppBrandUploadResponse](miniapp_response.go#L19)

小程序品牌提交认证时上传用户资质以及身份信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `file_content` | `FileContent` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-miniminiappserviceconfigmodify"></a>

### MiniMiniAppServiceConfigModify — 小程序设置客服方式

- 官方接口：[`alipay.open.mini.miniapp.serviceconfig.modify`](https://opendocs.alipay.com/open-v3/95e9abe6_alipay.open.mini.miniapp.serviceconfig.modify)
- 路由：`POST /v3/alipay/open/mini/miniapp/serviceconfig/modify`
- 请求类型：[MiniMiniAppServiceConfigModifyRequest](miniapp_request.go#L23)
- 响应类型：[MiniMiniAppServiceConfigModifyResponse](miniapp_response.go#L25)

小程序设置客服方式

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `home_open` | `HomeOpen` | `*bool` | JSON body | 未标注 | 云客服是否在小程序首页透出，true-开启，false-不开启 |
| `service_config` | `ServiceConfig` | `*string` | JSON body | 未标注 | 客服方式，目前支持ANTCLOUD-云客服 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-authorizeinfo"></a>

### AuthorizeInfo

类型定义：[AuthorizeInfo](miniapp_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_ids` | `AuthIds` | `*[]string` | 未标注 | 授权方身份证明(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片,可以上传多个授权方身份证明) |
| `auth_materials` | `AuthMaterials` | `*[]string` | 未标注 | 授权有效资质(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片，可以上传多个授权有效资质) |
| `auth_name` | `AuthName` | `*string` | 未标注 | 被授权人名称 |
| `end_date` | `EndDate` | `*string` | 未标注 | 授权有效期(结束日期) |
| `start_date` | `StartDate` | `*string` | 未标注 | 授权有效期(开始日期) |

<a id="model-brandregistrationinfo"></a>

### BrandRegistrationInfo

类型定义：[BrandRegistrationInfo](miniapp_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `end_date` | `EndDate` | `*string` | 未标注 | 商标有效期(结束日期) |
| `reg_materials` | `RegMaterials` | `*[]string` | 未标注 | 商标注册证书信息(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片,可以上传多个商标注册证书信息) |
| `reg_number` | `RegNumber` | `*string` | 未标注 | 商标注册号 |
| `registrant` | `Registrant` | `*string` | 未标注 | 商标注册人 |
| `start_date` | `StartDate` | `*string` | 未标注 | 商标有效期(开始日期) |

## 使用示例

- [MiniMiniAppBrandQuery](api_miniapp_miniminiappbrandquery_example_test.go)
- [MiniMiniAppBrandCancel](api_miniapp_miniminiappbrandcancel_example_test.go)
- [MiniMiniAppBrandSubmit](api_miniapp_miniminiappbrandsubmit_example_test.go)
- [MiniMiniAppBrandUpload](api_miniapp_miniminiappbrandupload_example_test.go)
- [MiniMiniAppServiceConfigModify](api_miniapp_miniminiappserviceconfigmodify_example_test.go)
