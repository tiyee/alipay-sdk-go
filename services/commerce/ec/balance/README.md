# balance 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/ec/balance`。本包提供 2 个支付宝 V3 REST 操作。

通过 `balance.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`EcBalanceDownloadURLQuery`](#api-ecbalancedownloadurlquery) | `alipay.commerce.ec.balance.downloadurl.query` |
| [`EcBalancePeriodModify`](#api-ecbalanceperiodmodify) | `alipay.commerce.ec.balance.period.modify` |

## 使用示例

- [EcBalanceDownloadURLQuery](api_balance_ecbalancedownloadurlquery_example_test.go)
- [EcBalancePeriodModify](api_balance_ecbalanceperiodmodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ecbalancedownloadurlquery"></a>

### EcBalanceDownloadURLQuery — 对账单文件下载接口

- 官方接口：[`alipay.commerce.ec.balance.downloadurl.query`](https://opendocs.alipay.com/apis/05j1bx)
- 路由：`GET /v3/alipay/commerce/ec/balance/downloadurl/query`
- 请求类型：[EcBalanceDownloadURLQueryRequest](balance_request.go#L16)
- 响应类型：[EcBalanceDownloadURLQueryResponse](balance_response.go#L4)

对账单文件下载接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业ID |
| `account_id` | `AccountId` | `*string` | query | 否 | 共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `bill_type` | `BillType` | `*string` | query | 否 | 账单类型 |
| `bill_date` | `BillDate` | `*string` | query | 否 | 账单时间 |

<a id="api-ecbalanceperiodmodify"></a>

### EcBalancePeriodModify — 企业码月对账单账期修改

- 官方接口：[`alipay.commerce.ec.balance.period.modify`](https://opendocs.alipay.com/apis/05j82h)
- 路由：`POST /v3/alipay/commerce/ec/balance/period/modify`
- 请求类型：[EcBalancePeriodModifyRequest](balance_request.go#L4)
- 响应类型：[EcBalancePeriodModifyResponse](balance_response.go#L10)

用于在企业码B端后台创建的企业调整月对账单账期

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id - 适用于在企业码小程序创建的共同账户，和agreement_no搭配使用 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `bill_day` | `BillDay` | `*int64` | JSON body | 未标注 | 月账单账期 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID - 适用于在企业码PC端创建的企业账号 |
