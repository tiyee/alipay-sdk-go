# open/mini/innerbaseinfo 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/innerbaseinfo`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `innerbaseinfo.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_innerbaseinfo_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniInnerBaseInfoAppLogoUpload`](#api-miniinnerbaseinfoapplogoupload) | [`alipay.open.mini.innerbaseinfo.applogo.upload`](https://opendocs.alipay.com/pre-apis/05xdzx) | 内部小程序-小程序logo图片上传 | `POST` |
| [`MiniInnerBaseInfoQuery`](#api-miniinnerbaseinfoquery) | [`alipay.open.mini.innerbaseinfo.query`](https://opendocs.alipay.com/pre-apis/05xiqj) | 内部小程序-应用信息查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerbaseinfoapplogoupload"></a>

### MiniInnerBaseInfoAppLogoUpload — 内部小程序-小程序logo图片上传

- 官方接口：[`alipay.open.mini.innerbaseinfo.applogo.upload`](https://opendocs.alipay.com/pre-apis/05xdzx)
- 路由：`POST /v3/alipay/open/mini/innerbaseinfo/applogo/upload`
- 请求类型：[MiniInnerBaseInfoAppLogoUploadRequest](innerbaseinfo_request.go#L9)
- 响应类型：[MiniInnerBaseInfoAppLogoUploadResponse](innerbaseinfo_response.go#L5)

上传小程序logo图片

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-miniinnerbaseinfoquery"></a>

### MiniInnerBaseInfoQuery — 内部小程序-应用信息查询

- 官方接口：[`alipay.open.mini.innerbaseinfo.query`](https://opendocs.alipay.com/pre-apis/05xiqj)
- 路由：`GET /v3/alipay/open/mini/innerbaseinfo/query`
- 请求类型：[MiniInnerBaseInfoQueryRequest](innerbaseinfo_request.go#L14)
- 响应类型：[MiniInnerBaseInfoQueryResponse](innerbaseinfo_response.go#L11)

查询小程序应用信息

请求约束：MiniInnerBaseInfoQueryRequest is the request type for alipay.open.mini.innerbaseinfo.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。 |
| `inst_code` | `InstCode` | `*string` | query | 否 | 租户code，alipay or taobao |
| `mini_app_name` | `MiniAppName` | `*string` | query | 否 | 小程序name，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。 |
| `app_sub_type` | `AppSubType` | `*string` | query | 否 | 小程序类型，TINYAPP_TEMPLATE，TINYAPP_NORMAL，TINYAPP_PLUGIN，使用mini_app_name查询的时候，该字段要求必传。 |
