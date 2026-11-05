# api 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/app/api`。本包提供 4 个支付宝 V3 REST 操作。

通过 `api.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AppAPIFieldApply`](#api-appapifieldapply) | `alipay.open.app.api.field.apply` |
| [`AppAPIFieldQuery`](#api-appapifieldquery) | `alipay.open.app.api.field.query` |
| [`AppAPIQuery`](#api-appapiquery) | `alipay.open.app.api.query` |
| [`AppAPISceneQuery`](#api-appapiscenequery) | `alipay.open.app.api.scene.query` |

## 使用示例

- [AppAPIFieldApply](api_api_appapifieldapply_example_test.go)
- [AppAPIFieldQuery](api_api_appapifieldquery_example_test.go)
- [AppAPIQuery](api_api_appapiquery_example_test.go)
- [AppAPISceneQuery](api_api_appapiscenequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-appapifieldapply"></a>

### AppAPIFieldApply — 申请获取接口用户敏感信息字段

- 官方接口：[`alipay.open.app.api.field.apply`](https://opendocs.alipay.com/open-v3/db6692ea_alipay.open.app.api.field.apply)
- 路由：`POST /v3/alipay/open/app/api/field/apply`
- 请求类型：[AppAPIFieldApplyRequest](api_request.go#L8)
- 响应类型：[AppAPIFieldApplyResponse](api_response.go#L4)

申请获取接口出参中用户敏感信息字段，应用使用视频或图片（至少选择一项上传），单个图片或视频不超过10M，最多可上传5张图片示例。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*AppAPIFieldApply`](#model-appapifieldapply) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `picture_1` | `Picture1` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `picture_2` | `Picture2` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `picture_3` | `Picture3` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `picture_4` | `Picture4` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `picture_5` | `Picture5` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `video` | `Video` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-appapifieldquery"></a>

### AppAPIFieldQuery — 用户信息申请记录查询

- 官方接口：[`alipay.open.app.api.field.query`](https://opendocs.alipay.com/open-v3/0cfbef83_alipay.open.app.api.field.query)
- 路由：`GET /v3/alipay/open/app/api/field/query`
- 请求类型：[AppAPIFieldQueryRequest](api_request.go#L19)
- 响应类型：[AppAPIFieldQueryResponse](api_response.go#L8)

查询用户信息申请记录，开发者发起请求时会基于当前发起请求的app_id进行查询，若为isv发起的三方代调用则以授权的商户app_id进行查询。

请求约束：AppAPIFieldQueryRequest is the request type for alipay.open.app.api.field.query.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-appapiquery"></a>

### AppAPIQuery — 查询应用可申请的接口出参敏感字段列表

- 官方接口：[`alipay.open.app.api.query`](https://opendocs.alipay.com/open-v3/d091ad0b_alipay.open.app.api.query)
- 路由：`GET /v3/alipay/open/app/api/query`
- 请求类型：[AppAPIQueryRequest](api_request.go#L22)
- 响应类型：[AppAPIQueryResponse](api_response.go#L13)

通过appId查询该应用下可申请的接口敏感字段信息

请求约束：AppAPIQueryRequest is the request type for alipay.open.app.api.query.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-appapiscenequery"></a>

### AppAPISceneQuery — 查询接口字段使用场景

- 官方接口：[`alipay.open.app.api.scene.query`](https://opendocs.alipay.com/open-v3/2bc4666c_alipay.open.app.api.scene.query)
- 路由：`GET /v3/alipay/open/app/api/scene/query`
- 请求类型：[AppAPISceneQueryRequest](api_request.go#L25)
- 响应类型：[AppAPISceneQueryResponse](api_response.go#L19)

在申请接口出参中用户敏感信息字段时，本接口用于获取所申请接口字段的使用场景

请求约束：AppAPISceneQueryRequest is the request type for alipay.open.app.api.scene.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `field_name` | `FieldName` | `*string` | query | 否 | 接口出参字段英文名。 |
| `api_name` | `ApiName` | `*string` | query | 否 | 接口英文名。 |

## 嵌套请求参数

<a id="model-appapifieldapply"></a>

### AppAPIFieldApply

类型定义：[AppAPIFieldApply](api_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_field_apply` | `AuthFieldApply` | [`*AuthFieldApply`](#model-authfieldapply) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-authfieldapply"></a>

### AuthFieldApply

类型定义：[AuthFieldApply](api_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `api_name` | `ApiName` | `*string` | 未标注 | 接口英文名称，通过alipay.open.app.api.query接口查询获得。 |
| `customer_answer` | `CustomerAnswer` | `*string` | 未标注 | 贵公司是否有自己的客服渠道，能及时响应和处理舆情数量是多少？ |
| `field_name` | `FieldName` | `*string` | 未标注 | 字段英文名称，通过alipay.open.app.api.query接口查询获得。 |
| `memo` | `Memo` | `*string` | 未标注 | 备注 |
| `package_code` | `PackageCode` | `*string` | 未标注 | 功能code，通过alipay.open.app.api.query接口查询获得 |
| `qps_answer` | `QpsAnswer` | `*string` | 未标注 | 接入后一年内预计接口秒级调用量峰值是多少？（最高峰值：1000QPS） |
| `scene_code` | `SceneCode` | `*string` | 未标注 | 场景code，alipay.open.app.api.scene.query 接口查询获得 |
| `tiny_app_template_id` | `TinyAppTemplateId` | `*string` | 未标注 | 当为使用使用模板的小程序申请时可传入所使用的小程序模板id |
