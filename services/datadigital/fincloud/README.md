# datadigital/fincloud 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/datadigital/fincloud`。本目录提供 10 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `fincloud.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_fincloud_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`GeneralSaaSFaceCertifyInitialize`](#api-generalsaasfacecertifyinitialize) | [`datadigital.fincloud.generalsaas.face.certify.initialize`](https://opendocs.alipay.com/open-v3/645bdc9c_datadigital.fincloud.generalsaas.face.certify.initialize) | 跳转支付宝人脸核身初始化 | `POST` |
| [`GeneralSaaSFaceCertifyQuery`](#api-generalsaasfacecertifyquery) | [`datadigital.fincloud.generalsaas.face.certify.query`](https://opendocs.alipay.com/open-v3/10a1ccfd_datadigital.fincloud.generalsaas.face.certify.query) | 跳转支付宝人脸核身查询记录 | `GET` |
| [`GeneralSaaSFaceCertifyVerify`](#api-generalsaasfacecertifyverify) | [`datadigital.fincloud.generalsaas.face.certify.verify`](https://opendocs.alipay.com/open-v3/d3f37a88_datadigital.fincloud.generalsaas.face.certify.verify) | 跳转支付宝人脸核身开始认证 | `POST` |
| [`GeneralSaaSFaceCheckInitialize`](#api-generalsaasfacecheckinitialize) | [`datadigital.fincloud.generalsaas.face.check.initialize`](https://opendocs.alipay.com/open-v3/503aa13e_datadigital.fincloud.generalsaas.face.check.initialize) | 活体检测初始化 | `POST` |
| [`GeneralSaaSFaceCheckQuery`](#api-generalsaasfacecheckquery) | [`datadigital.fincloud.generalsaas.face.check.query`](https://opendocs.alipay.com/open-v3/4b93abbc_datadigital.fincloud.generalsaas.face.check.query) | 活体检测结果查询 | `GET` |
| [`GeneralSaaSFaceSourceCertify`](#api-generalsaasfacesourcecertify) | [`datadigital.fincloud.generalsaas.face.source.certify`](https://opendocs.alipay.com/open-v3/21280890_datadigital.fincloud.generalsaas.face.source.certify) | 纯服务端人脸核身 | `POST` |
| [`GeneralSaaSFaceVerificationInitialize`](#api-generalsaasfaceverificationinitialize) | [`datadigital.fincloud.generalsaas.face.verification.initialize`](https://opendocs.alipay.com/open-v3/34fb02dc_datadigital.fincloud.generalsaas.face.verification.initialize) | 人脸核身初始化 | `POST` |
| [`GeneralSaaSFaceVerificationQuery`](#api-generalsaasfaceverificationquery) | [`datadigital.fincloud.generalsaas.face.verification.query`](https://opendocs.alipay.com/open-v3/1ef1e2ed_datadigital.fincloud.generalsaas.face.verification.query) | 人脸核身结果查询 | `GET` |
| [`GeneralSaaSOCRMobileInitialize`](#api-generalsaasocrmobileinitialize) | [`datadigital.fincloud.generalsaas.ocr.mobile.initialize`](https://opendocs.alipay.com/open-v3/74ea0784_datadigital.fincloud.generalsaas.ocr.mobile.initialize) | App端OCR初始化 | `POST` |
| [`GeneralSaaSOCRServerDetect`](#api-generalsaasocrserverdetect) | [`datadigital.fincloud.generalsaas.ocr.server.detect`](https://opendocs.alipay.com/open-v3/a6f20b9d_datadigital.fincloud.generalsaas.ocr.server.detect) | 服务端OCR | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-generalsaasfacecertifyinitialize"></a>

### GeneralSaaSFaceCertifyInitialize — 跳转支付宝人脸核身初始化

- 官方接口：[`datadigital.fincloud.generalsaas.face.certify.initialize`](https://opendocs.alipay.com/open-v3/645bdc9c_datadigital.fincloud.generalsaas.face.certify.initialize)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/face/certify/initialize`
- 请求类型：[GeneralSaaSFaceCertifyInitializeRequest](fincloud_request.go#L9)
- 响应类型：[GeneralSaaSFaceCertifyInitializeResponse](fincloud_response.go#L5)

人脸核身初始化服务，用于创建认证流程。供外部客户使用该项服务。

请求约束：GeneralSaaSFaceCertifyInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | JSON body | 未标注 | H5人脸核身场景码。入参支持的场景码。 |
| `face_contrast_picture` | `FaceContrastPicture` | `*string` | JSON body | 未标注 | 自定义人脸比对图片的base64编码格式的string字符串 |
| `identity_param` | `IdentityParam` | [`*OpenCertifyIdentifyInfo`](#model-opencertifyidentifyinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_config` | `MerchantConfig` | [`*OpenCertifyMerchantConfigs`](#model-opencertifymerchantconfigs) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | JSON body | 未标注 | 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列 |

<a id="api-generalsaasfacecertifyquery"></a>

### GeneralSaaSFaceCertifyQuery — 跳转支付宝人脸核身查询记录

- 官方接口：[`datadigital.fincloud.generalsaas.face.certify.query`](https://opendocs.alipay.com/open-v3/10a1ccfd_datadigital.fincloud.generalsaas.face.certify.query)
- 路由：`GET /v3/datadigital/fincloud/generalsaas/face/certify/query`
- 请求类型：[GeneralSaaSFaceCertifyQueryRequest](fincloud_request.go#L67)
- 响应类型：[GeneralSaaSFaceCertifyQueryResponse](fincloud_response.go#L11)

商户在用户使用支付宝完成刷脸之后，调用本接口查询核身结果状态及相关数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `certify_id` | `CertifyId` | `*string` | query | 否 | 本次申请操作的唯一标识，通过datadigital.fincloud.generalsaas.face.certify.initialize 接口同步响应获取。 |

<a id="api-generalsaasfacecertifyverify"></a>

### GeneralSaaSFaceCertifyVerify — 跳转支付宝人脸核身开始认证

- 官方接口：[`datadigital.fincloud.generalsaas.face.certify.verify`](https://opendocs.alipay.com/open-v3/d3f37a88_datadigital.fincloud.generalsaas.face.certify.verify)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/face/certify/verify`
- 请求类型：[GeneralSaaSFaceCertifyVerifyRequest](fincloud_request.go#L23)
- 响应类型：[GeneralSaaSFaceCertifyVerifyResponse](fincloud_response.go#L23)

人脸核身服务开始认证。请根据返回结果渲染页面。

请求约束：GeneralSaaSFaceCertifyVerifyRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCertifyVerifyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `certify_id` | `CertifyId` | `*string` | JSON body | 未标注 | 本次申请操作的唯一标识，由H5人脸核身初始化接口调用后生成，后续的操作都需要用到 |

<a id="api-generalsaasfacecheckinitialize"></a>

### GeneralSaaSFaceCheckInitialize — 活体检测初始化

- 官方接口：[`datadigital.fincloud.generalsaas.face.check.initialize`](https://opendocs.alipay.com/open-v3/503aa13e_datadigital.fincloud.generalsaas.face.check.initialize)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/face/check/initialize`
- 请求类型：[GeneralSaaSFaceCheckInitializeRequest](fincloud_request.go#L29)
- 响应类型：[GeneralSaaSFaceCheckInitializeResponse](fincloud_response.go#L29)

人脸检测业务单据初始化

请求约束：GeneralSaaSFaceCheckInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceCheckInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | JSON body | 未标注 | DATA_DIGITAL_BIZ_CODE_FACE_CHECK_LIVE，代表活体检测。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | JSON body | 未标注 | 客户业务单据号。 |

<a id="api-generalsaasfacecheckquery"></a>

### GeneralSaaSFaceCheckQuery — 活体检测结果查询

- 官方接口：[`datadigital.fincloud.generalsaas.face.check.query`](https://opendocs.alipay.com/open-v3/4b93abbc_datadigital.fincloud.generalsaas.face.check.query)
- 路由：`GET /v3/datadigital/fincloud/generalsaas/face/check/query`
- 请求类型：[GeneralSaaSFaceCheckQueryRequest](fincloud_request.go#L73)
- 响应类型：[GeneralSaaSFaceCheckQueryResponse](fincloud_response.go#L39)

用于查询活体检测结果。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `certify_id` | `CertifyId` | `*string` | query | 否 | 传入初始化接口datadigital.fincloud.generalsaas.face.check.initialize 得到的certify_id |
| `need_alive_photo` | `NeedAlivePhoto` | `*string` | query | 否 | 是否需要返回活体结果图片，needAlivePhoto：Y_O （原始图片）、Y_M（虚化，背景马赛克）、N（不返图）。非必填，默认为N |
| `need_score` | `NeedScore` | `*string` | query | 否 | 该入参已经废弃。 |
| `need_quality_score` | `NeedQualityScore` | `*string` | query | 否 | 是否需要返回图片质量分，非必填，默认值为N。 N表示不需要，Y表示需要 |
| `need_attack_result` | `NeedAttackResult` | `*string` | query | 否 | 该参数已废弃 |

<a id="api-generalsaasfacesourcecertify"></a>

### GeneralSaaSFaceSourceCertify — 纯服务端人脸核身

- 官方接口：[`datadigital.fincloud.generalsaas.face.source.certify`](https://opendocs.alipay.com/open-v3/21280890_datadigital.fincloud.generalsaas.face.source.certify)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/face/source/certify`
- 请求类型：[GeneralSaaSFaceSourceCertifyRequest](fincloud_request.go#L87)
- 响应类型：[GeneralSaaSFaceSourceCertifyResponse](fincloud_response.go#L55)

上传人脸照片，调用权威核验源进行验证

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*GeneralSaaSFaceSourceCertify`](#model-generalsaasfacesourcecertify) | multipart | 未标注 | the data multipart field. |
| `file_content` | `FileContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | the file_content multipart field. |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-generalsaasfaceverificationinitialize"></a>

### GeneralSaaSFaceVerificationInitialize — 人脸核身初始化

- 官方接口：[`datadigital.fincloud.generalsaas.face.verification.initialize`](https://opendocs.alipay.com/open-v3/34fb02dc_datadigital.fincloud.generalsaas.face.verification.initialize)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/face/verification/initialize`
- 请求类型：[GeneralSaaSFaceVerificationInitializeRequest](fincloud_request.go#L37)
- 响应类型：[GeneralSaaSFaceVerificationInitializeResponse](fincloud_response.go#L69)

人脸核身初始化接口

请求约束：GeneralSaaSFaceVerificationInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasFaceVerificationInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | JSON body | 未标注 | 人脸核身具体类型目前仅支持：DATA_DIGITAL_BIZ_CODE_FACE_VERIFICATION |
| `cert_name` | `CertName` | `*string` | JSON body | 未标注 | 真实姓名 |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 证件号 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 证件类型，当前枚举支持： IDENTITY_CARD：身份证 RESIDENCE_HK_MC：港澳居民居住证 RESIDENCE_TAIWAN：台湾居民居住证 |
| `face_reserve_strategy` | `FaceReserveStrategy` | `*string` | JSON body | 未标注 | 人脸保存策略，非必填。具体取值为：reserve(保存活体人脸)/never(不保存活体人脸)，不传默认为reserve |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 认证类型，固定值为：CERT_INFO |
| `nation` | `Nation` | `*string` | JSON body | 未标注 | 国家地区代码，当证件类型为外国人永久居留证时必填。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | JSON body | 未标注 | 商户请求的唯一标识，商户要保证其唯一性，值为64位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列 |
| `phone_no` | `PhoneNo` | `*string` | JSON body | 未标注 | 手机号码 |

<a id="api-generalsaasfaceverificationquery"></a>

### GeneralSaaSFaceVerificationQuery — 人脸核身结果查询

- 官方接口：[`datadigital.fincloud.generalsaas.face.verification.query`](https://opendocs.alipay.com/open-v3/1ef1e2ed_datadigital.fincloud.generalsaas.face.verification.query)
- 路由：`GET /v3/datadigital/fincloud/generalsaas/face/verification/query`
- 请求类型：[GeneralSaaSFaceVerificationQueryRequest](fincloud_request.go#L95)
- 响应类型：[GeneralSaaSFaceVerificationQueryResponse](fincloud_response.go#L79)

查询核身结果

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `certify_id` | `CertifyId` | `*string` | query | 否 | 填入人脸核身初始化阶段获取到的certify_id |
| `need_alive_photo` | `NeedAlivePhoto` | `*string` | query | 否 | 是否需要返回活体结果图片，默认为N（不返图） |

<a id="api-generalsaasocrmobileinitialize"></a>

### GeneralSaaSOCRMobileInitialize — App端OCR初始化

- 官方接口：[`datadigital.fincloud.generalsaas.ocr.mobile.initialize`](https://opendocs.alipay.com/open-v3/74ea0784_datadigital.fincloud.generalsaas.ocr.mobile.initialize)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize`
- 请求类型：[GeneralSaaSOCRMobileInitializeRequest](fincloud_request.go#L59)
- 响应类型：[GeneralSaaSOCRMobileInitializeResponse](fincloud_response.go#L95)

OCR端云一体化识别

请求约束：GeneralSaaSOCRMobileInitializeRequest is the official V3 DatadigitalFincloudGeneralsaasOcrMobileInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | JSON body | 未标注 | bizCode，代表当前使用的能力类型。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | JSON body | 未标注 | 客户业务单据号。请保持唯一。 |

<a id="api-generalsaasocrserverdetect"></a>

### GeneralSaaSOCRServerDetect — 服务端OCR

- 官方接口：[`datadigital.fincloud.generalsaas.ocr.server.detect`](https://opendocs.alipay.com/open-v3/a6f20b9d_datadigital.fincloud.generalsaas.ocr.server.detect)
- 路由：`POST /v3/datadigital/fincloud/generalsaas/ocr/server/detect`
- 请求类型：[GeneralSaaSOCRServerDetectRequest](fincloud_request.go#L103)
- 响应类型：[GeneralSaaSOCRServerDetectResponse](fincloud_response.go#L101)

纯服务端OCR对接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*GeneralSaaSOCRServerDetect`](#model-generalsaasocrserverdetect) | multipart | 未标注 | the data multipart field. |
| `file_content` | `FileContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | the file_content multipart field. |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-generalsaasfacesourcecertify"></a>

### GeneralSaaSFaceSourceCertify

类型定义：[GeneralSaaSFaceSourceCertify](fincloud_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | 未标注 | 业务产品码 |
| `cert_name` | `CertName` | `*string` | 未标注 | 用户姓名，与身份证上的姓名相匹配 |
| `cert_no` | `CertNo` | `*string` | 未标注 | 证件号码，如大陆身份证号码 |
| `cert_type` | `CertType` | `*string` | 未标注 | 证件信息类型，大陆身份证：IDENTITY_CARD |
| `liveness_strategy` | `LivenessStrategy` | `*string` | 未标注 | 活体检测类型，当前枚举支持： OFF：关闭活体检测； CHECK：正常活体检测。 不传默认CHECK。 |
| `nation` | `Nation` | `*string` | 未标注 | 国家地区代码，当证件类型为外国人永久居留证时必填。 |
| `outer_biz_no` | `OuterBizNo` | `*string` | 未标注 | 客户业务单据号 |
| `phone` | `Phone` | `*string` | 未标注 | 手机号，特殊场景下，在验证过程中需用户输入手机来接受系统发送的安全验证短信。默认不传即可。 |
| `reserved` | `Reserved` | `*bool` | 未标注 | true：需要留底；false：无需留底。默认留底 |
| `security_strategy` | `SecurityStrategy` | `*string` | 未标注 | ON：开启安全拦截。OFF：关闭安全拦截。默认取值为ON。 |

<a id="model-generalsaasocrserverdetect"></a>

### GeneralSaaSOCRServerDetect

类型定义：[GeneralSaaSOCRServerDetect](fincloud_model.go#L29)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ocr_type` | `OcrType` | `*string` | 未标注 | 可识别OCR类型。如，ID_CARD_FRONT 身份证含照片的一面；ID_CARD_BACK 身份证带国徽的一面。具体支持的类型以接入文档为准。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | 未标注 | 客户业务单据号，请保证幂等性。 |

<a id="model-opencertifyidentifyinfo"></a>

### OpenCertifyIdentifyInfo

类型定义：[OpenCertifyIdentifyInfo](fincloud_model.go#L37)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cert_name` | `CertName` | `*string` | 未标注 | 填入真实姓名 |
| `cert_no` | `CertNo` | `*string` | 未标注 | 填入姓名相匹配的证件号码 |
| `cert_type` | `CertType` | `*string` | 未标注 | 当前仅支持IDENTITY_CARD |
| `identity_type` | `IdentityType` | `*string` | 未标注 | 当前仅支持CERT_INFO |
| `phone_no` | `PhoneNo` | `*string` | 未标注 | 选填手机号 |

<a id="model-opencertifymerchantconfigs"></a>

### OpenCertifyMerchantConfigs

类型定义：[OpenCertifyMerchantConfigs](fincloud_model.go#L51)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `face_reserve_strategy` | `FaceReserveStrategy` | `*string` | 未标注 | 不传默认为reserve |
| `return_url` | `ReturnUrl` | `*string` | 未标注 | 认证成功后需要跳转的地址，一般为商户业务页面；若无跳转地址可填空字符""。 |
