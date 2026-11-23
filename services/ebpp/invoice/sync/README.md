# ebpp/invoice/sync 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/sync`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `sync.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_sync_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceSyncSimpleSend`](#api-invoicesyncsimplesend) | [`alipay.ebpp.invoice.sync.simple.send`](https://opendocs.alipay.com/apis/05kxx8) | 简单模式发票回传接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicesyncsimplesend"></a>

### InvoiceSyncSimpleSend — 简单模式发票回传接口

- 官方接口：[`alipay.ebpp.invoice.sync.simple.send`](https://opendocs.alipay.com/apis/05kxx8)
- 路由：`POST /v3/alipay/ebpp/invoice/sync/simple/send`
- 请求类型：[InvoiceSyncSimpleSendRequest](sync_request.go#L5)
- 响应类型：[InvoiceSyncSimpleSendResponse](sync_response.go#L9)

商户或ISV将发票信息发送至发票管家。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `invoice_info` | `InvoiceInfo` | [`*InvoicePDFSyn`](#model-invoicepdfsyn) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `m_short_name` | `MShortName` | `*string` | JSON body | 未标注 | 开票商户品牌简称，与商户入驻时的品牌简称保持一致。详情参见 [电子发票](https://opendocs.alipay.com/open/10691/welcome-to-lark) |
| `sub_m_short_name` | `SubMShortName` | `*string` | JSON body | 未标注 | 开票商户门店简称，与商户入驻时的门店简称保持一致。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-invoicepdfsyn"></a>

### InvoicePDFSyn

类型定义：[InvoicePDFSyn](sync_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `apply_id` | `ApplyId` | `*string` | 未标注 | 支付宝开票申请id |
| `extend_fields` | `ExtendFields` | `*string` | 未标注 | 预留的扩展字段，格式如：key1=value1\nkey2=value2\nkey3=value3，字段之间以\n(换行符)分隔。 |
| `file_base` | `FileBase` | `*string` | 未标注 | base64化的字符串，与file_download_url不可同时为空。 |
| `file_download_type` | `FileDownloadType` | `*string` | 未标注 | 文件类型，必填。 PDF类型文件填写PDF， OFD类型文件填写ofd， JPG类型文件填写JPG(JPG文件请先询问对接人当前是否支持) |
| `file_download_url` | `FileDownloadUrl` | `*string` | 未标注 | 发票文件下载地址，和file_base不可同时为空。 |
| `login_id` | `LoginId` | `*string` | 未标注 | 支付宝登录账号 |
| `open_id` | `OpenId` | `*string` | 未标注 | 支付宝用户id，当apply_id为空时，open_id必填 |
| `out_invoice_id` | `OutInvoiceId` | `*string` | 未标注 | 外部ISV的唯一发票标识 |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝用户userId |
| `zip` | `Zip` | `*string` | 未标注 | 压缩方式 |
