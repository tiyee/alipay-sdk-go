# ebpp/invoice/taxno 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/taxno`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `taxno.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_taxno_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceTaxNoBatchQuery`](#api-invoicetaxnobatchquery) | [`alipay.ebpp.invoice.taxno.batchquery`](https://opendocs.alipay.com/apis/05jrk9) | 获取指定企业的发票要素列表 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicetaxnobatchquery"></a>

### InvoiceTaxNoBatchQuery — 获取指定企业的发票要素列表

- 官方接口：[`alipay.ebpp.invoice.taxno.batchquery`](https://opendocs.alipay.com/apis/05jrk9)
- 路由：`POST /v3/alipay/ebpp/invoice/taxno/batchquery`
- 请求类型：[InvoiceTaxNoBatchQueryRequest](taxno_request.go#L5)
- 响应类型：[InvoiceTaxNoBatchQueryResponse](taxno_response.go#L5)

根据购买方税号获取用户的发票关键要素列表，该接口目前适用于发票报销场景

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `enable_trade_out` | `EnableTradeOut` | `*string` | JSON body | 未标注 | 默认值为false。true为输出交易信息，false为不输出交易信息。 |
| `end_invoice_date` | `EndInvoiceDate` | `*string` | JSON body | 未标注 | 查询结束时间，精确到天（按开票日期查询） start_invoice_date和end_invoice_date传值要求 1.同时为空时，返回最近半年200条数据 2.必须同时为空 或 同时不为空 3.结束日期不能大于当前日期 4.开始时间和结束时间跨度不能超过6个月 |
| `invoice_kind_list` | `InvoiceKindList` | `*[]string` | JSON body | 未标注 | 查询票种列表 |
| `limit_size` | `LimitSize` | `*int64` | JSON body | 未标注 | 查询结果上限笔数，最大值20 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 当前页码，为空时默认第一页 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 发票要素获取应用场景。["拉"模式报销](https://opendocs.alipay.com/open/10691/bv8s88) 固定为 INVOICE_EXPENSE 表示发票报销。 |
| `start_invoice_date` | `StartInvoiceDate` | `*string` | JSON body | 未标注 | 查询起始时间，精确到天（按开票日期查询） start_invoice_date和end_invoice_date传值要求 1.同时为空时，返回最近半年200条数据 2.必须同时为空 或 同时不为空 3.结束日期不能大于当前日期 4.开始时间和结束时间跨度不能超过6个月 |
| `tax_no` | `TaxNo` | `*string` | JSON body | 未标注 | 企业税号 |
