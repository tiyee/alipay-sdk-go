# eco/doc 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/doc`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `doc.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_doc_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`DocTemplateCreate`](#api-doctemplatecreate) | [`alipay.eco.doc.template.create`](https://opendocs.alipay.com/pre-apis/05whak) | 创建合同模板（E签宝） | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-doctemplatecreate"></a>

### DocTemplateCreate — 创建合同模板（E签宝）

- 官方接口：[`alipay.eco.doc.template.create`](https://opendocs.alipay.com/pre-apis/05whak)
- 路由：`POST /v3/alipay/eco/doc/template/create`
- 请求类型：[DocTemplateCreateRequest](doc_request.go#L5)
- 响应类型：[DocTemplateCreateResponse](doc_response.go#L5)

通过此接口创建合同模板后，通过返回的文件直传地址上传模板文件（WORD/PDF）。同时记录接口返回的模板ID（templateId），后续发起合同签署需要使用到。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `content_md_5` | `ContentMd5` | `*string` | JSON body | 未标注 | 模版文件的md5值。示例代码参见 [创建合同模板](https://opendocs.alipay.com/mini/00arl2#%E5%88%9B%E5%BB%BA%E5%90%88%E5%90%8C%E6%A8%A1%E6%9D%BF) |
| `content_type` | `ContentType` | `*string` | JSON body | 未标注 | 目标文件的MIME类型。示例代码参见 [创建合同模板](https://opendocs.alipay.com/mini/00arl2#%E5%88%9B%E5%BB%BA%E5%90%88%E5%90%8C%E6%A8%A1%E6%9D%BF) |
| `convert_to_pdf` | `ConvertToPdf` | `*string` | JSON body | 未标注 | 是否需要转成pdf，如果模板文件为.doc/.docx 传true，为pdf传false |
| `file_name` | `FileName` | `*string` | JSON body | 未标注 | 文件名称，必须带上文件扩展名如 .pdf、.doc、.docx，否则会导致后续发起流程校验异常。 |
