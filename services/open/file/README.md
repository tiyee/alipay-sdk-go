# open/file 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/file`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `file.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_file_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`FileUpload`](#api-fileupload) | [`alipay.open.file.upload`](https://opendocs.alipay.com/open-v3/5aa91070_alipay.open.file.upload?scene=common&pathHash=e40a9831) | 支付宝文件上传接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-fileupload"></a>

### FileUpload — 支付宝文件上传接口

- 官方接口：[`alipay.open.file.upload`](https://opendocs.alipay.com/open-v3/5aa91070_alipay.open.file.upload?scene=common&pathHash=e40a9831)
- 路由：`POST /v3/alipay/open/file/upload`
- 请求类型：[FileUploadRequest](file_request.go#L9)
- 响应类型：[FileUploadResponse](file_response.go#L5)

支付宝通用文件上传接口，开发者可调用此接口上传文件，提供给其他业务使用，例如：上传素材（视频/图片），供给到生活号+/小程序等业务使用。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*FileUpload`](#model-fileupload) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `file_content` | `FileContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-fileupload"></a>

### FileUpload

类型定义：[FileUpload](file_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | 未标注 | 素材使用场景，请关联其他开放平台接口联合使用。如果素材用于内容发布，请必传入content_creation |
| `extra_parameters` | `ExtraParameters` | [`*FileUploadExtraParam`](#model-fileuploadextraparam) | 未标注 | 见字段类型及官方接口文档。 |
| `md5` | `Md5` | `*string` | 未标注 | 文件的md5值，如果不传则不校验md5 |

<a id="model-fileuploadextraparam"></a>

### FileUploadExtraParam

类型定义：[FileUploadExtraParam](file_model.go#L14)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `extern_upload` | `ExternUpload` | `*string` | 未标注 | 是否同步上传到其他平台。素材类型为视频且需要同步优酷，请必传入参数youku，素材接口将代上传到优酷账号；素材类型为图片，无需传入参数 |
| `file_encrypt_type` | `FileEncryptType` | `*string` | 未标注 | 文件加密类型 |
