# certify 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/user/certify`。本包提供 2 个支付宝 V3 REST 操作。

通过 `certify.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CertifyOpenInitialize`](#api-certifyopeninitialize) | `alipay.user.certify.open.initialize` |
| [`CertifyOpenQuery`](#api-certifyopenquery) | `alipay.user.certify.open.query` |

## 使用示例

- [CertifyOpenInitialize](api_certify_certifyopeninitialize_example_test.go)
- [CertifyOpenQuery](api_certify_certifyopenquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-certifyopeninitialize"></a>

### CertifyOpenInitialize — 身份认证初始化服务

- 官方接口：[`alipay.user.certify.open.initialize`](https://opendocs.alipay.com/open-v3/013ee3e1_alipay.user.certify.open.initialize)
- 路由：`POST /v3/alipay/user/certify/open/initialize`
- 请求类型：[UserCertifyOpenInitializeRequest](certify_request.go#L4)
- 响应类型：[UserCertifyOpenInitializeResponse](certify_response.go#L4)

支付宝开放认证初始化服务，用于创建开放认证流程。注意：港澳台同胞身份验证服务仅支持在支付宝客户端内使用，外部 APP 使用该项服务需唤起支付宝客户端。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | JSON body | 未标注 | 认证场景码。入参支持的认证场景码和商户签约的认证场景相关，取值如下: FACE：多因子人脸认证 CERT_PHOTO：多因子证照认证 CERT_PHOTO_FACE ：多因子证照和人脸认证 SMART_FACE：多因子快捷认证 |
| `face_contrast_picture` | `FaceContrastPicture` | `*string` | JSON body | 未标注 | 自定义人脸比对图片的base64编码格式的string字符串 |
| `identity_param` | `IdentityParam` | [`*OpenCertifyIdentityParam`](#model-opencertifyidentityparam) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_config` | `MerchantConfig` | [`*OpenCertifyMerchantConfig`](#model-opencertifymerchantconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `outer_order_no` | `OuterOrderNo` | `*string` | JSON body | 未标注 | 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列 |

<a id="api-certifyopenquery"></a>

### CertifyOpenQuery — 身份认证记录查询

- 官方接口：[`alipay.user.certify.open.query`](https://opendocs.alipay.com/open-v3/c166c117_alipay.user.certify.open.query)
- 路由：`GET /v3/alipay/user/certify/open/query`
- 请求类型：[UserCertifyOpenQueryRequest](certify_request.go#L16)
- 响应类型：[UserCertifyOpenQueryResponse](certify_response.go#L10)

商户在开放认证完成后，调用本接口查询认证状态和相关数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `certify_id` | `CertifyId` | `*string` | query | 否 | 本次申请操作的唯一标识，通过alipay.user.certify.open.initialize(身份认证初始化服务)接口同步响应获取。 |

## 嵌套请求参数

<a id="model-opencertifyidentityparam"></a>

### OpenCertifyIdentityParam

类型定义：[OpenCertifyIdentityParam](certify_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cert_digest` | `CertDigest` | `*string` | 未标注 | cert_digest： 身份信息摘要 |
| `cert_name` | `CertName` | `*string` | 未标注 | cert_name： 填入真实姓名 |
| `cert_no` | `CertNo` | `*string` | 未标注 | cert_no： 填入姓名相匹配的证件号码 |
| `cert_type` | `CertType` | `*string` | 未标注 | cert_type： 1、若为身份证，填IDENTITY_CARD； 2、若为港澳居民来往内地通行证，填HOME_VISIT_PERMIT_HK_MC； 3、若为台湾居民来往内地通行证，填HOME_VISIT_PERMIT_TAIWAN； 4、若为港澳居民居住证，填RESIDENCE_PERMIT_HK_MC； 5、若为台湾居民居住证，填RESIDENCE_PERMIT_TAIWAN； 6、外国人永久居留证，填PERMANENT_RESIDENCE_FOREIGNER； |
| `facial_picture_front` | `FacialPictureFront` | `*string` | 未标注 | facial_picture_front： 活体人脸图片base64串，主要用于离线人脸比对 |
| `identity_type` | `IdentityType` | `*string` | 未标注 | identity_type 1.若本人验证，使用CERT_INFO； 2.若代他人验证，使用AGENT_CERT_INFO； |
| `open_id` | `OpenId` | `*string` | 未标注 | open_id： 开放平台用户ID |
| `phone_no` | `PhoneNo` | `*string` | 未标注 | phone_no： 填写姓名相匹配的手机号 |
| `user_id` | `UserId` | `*string` | 未标注 | user_id： 支付宝用户ID |

<a id="model-opencertifymerchantconfig"></a>

### OpenCertifyMerchantConfig

类型定义：[OpenCertifyMerchantConfig](certify_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_scope` | `AuthScope` | `*string` | 未标注 | 用于开放认证授权 |
| `auth_type` | `AuthType` | `*string` | 未标注 | 用于指定授权类型，与auth_scope配合使用 |
| `face_reserve_strategy` | `FaceReserveStrategy` | `*string` | 未标注 | 不传默认为reserve |
| `facial_picture_level` | `FacialPictureLevel` | `*string` | 未标注 | 若有特殊人脸等级采集要求，可指定等级 |
| `linked_merchant_app_id` | `LinkedMerchantAppId` | `*string` | 未标注 | 用于授权二级商户操作 |
| `linked_merchant_logo_url` | `LinkedMerchantLogoUrl` | `*string` | 未标注 | 用于指定展示的商户logo |
| `linked_merchant_name` | `LinkedMerchantName` | `*string` | 未标注 | 用于指定展示的商户名称 |
| `out_put_facial_picture` | `OutPutFacialPicture` | `*bool` | 未标注 | 在拥有该权限前提下，用于商户控制是否透出活体人脸 |
| `return_url` | `ReturnUrl` | `*string` | 未标注 | 认证成功后需要跳转的地址，一般为商户业务页面；若无跳转地址可填空字符""; |
