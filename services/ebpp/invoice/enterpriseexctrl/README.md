# ebpp/invoice/enterpriseexctrl 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/enterpriseexctrl`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `enterpriseexctrl.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_enterpriseexctrl_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceEnterpriseExCtrlEmployerTitleBatchQuery`](#api-invoiceenterpriseexctrlemployertitlebatchquery) | [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery`](https://opendocs.alipay.com/apis/05wvun) | 企业抬头：分页查询企业抬头信息 | `POST` |
| [`InvoiceEnterpriseExCtrlEmployerTitleCreate`](#api-invoiceenterpriseexctrlemployertitlecreate) | [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.create`](https://opendocs.alipay.com/apis/05wydt) | 企业抬头：新增企业抬头 | `POST` |
| [`InvoiceEnterpriseExCtrlEmployerTitleModify`](#api-invoiceenterpriseexctrlemployertitlemodify) | [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify`](https://opendocs.alipay.com/apis/05lhsq) | 企业抬头：修改企业抬头 | `PUT` |
| [`InvoiceEnterpriseExCtrlEmployerTitleQuery`](#api-invoiceenterpriseexctrlemployertitlequery) | [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.query`](https://opendocs.alipay.com/apis/05x1fx) | 企业抬头：查询企业抬头 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceenterpriseexctrlemployertitlebatchquery"></a>

### InvoiceEnterpriseExCtrlEmployerTitleBatchQuery — 企业抬头：分页查询企业抬头信息

- 官方接口：[`alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery`](https://opendocs.alipay.com/apis/05wvun)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery`
- 请求类型：[InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest](enterpriseexctrl_request.go#L5)
- 响应类型：[InvoiceEnterpriseExCtrlEmployerTitleBatchQueryResponse](enterpriseexctrl_response.go#L5)

分页查询企业抬头列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 页码 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页行数 |
| `tax_register_no` | `TaxRegisterNo` | `*string` | JSON body | 未标注 | 税号 |

<a id="api-invoiceenterpriseexctrlemployertitlecreate"></a>

### InvoiceEnterpriseExCtrlEmployerTitleCreate — 企业抬头：新增企业抬头

- 官方接口：[`alipay.ebpp.invoice.enterpriseexctrl.employertitle.create`](https://opendocs.alipay.com/apis/05wydt)
- 路由：`POST /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create`
- 请求类型：[InvoiceEnterpriseExCtrlEmployerTitleCreateRequest](enterpriseexctrl_request.go#L21)
- 响应类型：[InvoiceEnterpriseExCtrlEmployerTitleCreateResponse](enterpriseexctrl_response.go#L17)

新增一条企业开票抬头

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id |
| `address` | `Address` | `*string` | JSON body | 未标注 | 企业地址 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `bank_account` | `BankAccount` | `*string` | JSON body | 未标注 | 开户行账号 |
| `bank_name` | `BankName` | `*string` | JSON body | 未标注 | 开户行名称 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `tax_register_no` | `TaxRegisterNo` | `*string` | JSON body | 未标注 | 税号 |
| `telephone` | `Telephone` | `*string` | JSON body | 未标注 | 电话 |
| `title_code` | `TitleCode` | `*string` | JSON body | 未标注 | 抬头编码 |
| `title_name` | `TitleName` | `*string` | JSON body | 未标注 | 企业抬头名称 |

<a id="api-invoiceenterpriseexctrlemployertitlemodify"></a>

### InvoiceEnterpriseExCtrlEmployerTitleModify — 企业抬头：修改企业抬头

- 官方接口：[`alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify`](https://opendocs.alipay.com/apis/05lhsq)
- 路由：`PUT /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle`
- 请求类型：[InvoiceEnterpriseExCtrlEmployerTitleModifyRequest](enterpriseexctrl_request.go#L45)
- 响应类型：[InvoiceEnterpriseExCtrlEmployerTitleModifyResponse](enterpriseexctrl_response.go#L29)

修改企业开票抬头

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id |
| `address` | `Address` | `*string` | JSON body | 未标注 | 详细地址 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `bank_account` | `BankAccount` | `*string` | JSON body | 未标注 | 开户行账号 |
| `bank_name` | `BankName` | `*string` | JSON body | 未标注 | 开户行 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `tax_register_no` | `TaxRegisterNo` | `*string` | JSON body | 未标注 | 税号 |
| `telephone` | `Telephone` | `*string` | JSON body | 未标注 | 电话 |
| `title_code` | `TitleCode` | `*string` | JSON body | 未标注 | 抬头编码 |
| `title_id` | `TitleId` | `*string` | JSON body | 未标注 | 抬头ID |
| `title_name` | `TitleName` | `*string` | JSON body | 未标注 | 企业抬头名称 |

<a id="api-invoiceenterpriseexctrlemployertitlequery"></a>

### InvoiceEnterpriseExCtrlEmployerTitleQuery — 企业抬头：查询企业抬头

- 官方接口：[`alipay.ebpp.invoice.enterpriseexctrl.employertitle.query`](https://opendocs.alipay.com/apis/05x1fx)
- 路由：`GET /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query`
- 请求类型：[InvoiceEnterpriseExCtrlEmployerTitleQueryRequest](enterpriseexctrl_request.go#L71)
- 响应类型：[InvoiceEnterpriseExCtrlEmployerTitleQueryResponse](enterpriseexctrl_response.go#L41)

根据抬头ID查询企业抬头详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业共同账户id |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `title_id` | `TitleId` | `*string` | query | 否 | 抬头ID |
| `tax_register_no` | `TaxRegisterNo` | `*string` | query | 否 | 税号 |
