# commerce/operation 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/operation`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `operation.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_operation_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`OperationActivityMerchantModify`](#api-operationactivitymerchantmodify) | [`alipay.commerce.operation.activity.merchant.modify`](https://opendocs.alipay.com/pre-apis/063rs2) | 商户签约信息修改 | `POST` |
| [`OperationActivityMerchantSign`](#api-operationactivitymerchantsign) | [`alipay.commerce.operation.activity.merchant.sign`](https://opendocs.alipay.com/pre-apis/06347n) | 商户签约活动 | `POST` |
| [`OperationActivityMerchantUnsign`](#api-operationactivitymerchantunsign) | [`alipay.commerce.operation.activity.merchant.unsign`](https://opendocs.alipay.com/pre-apis/06336j) | 活动商户解约 | `POST` |
| [`OperationBankActivityQuery`](#api-operationbankactivityquery) | [`alipay.commerce.operation.bank.activity.query`](https://opendocs.alipay.com/pre-apis/063ecd) | 银行活动信息查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-operationactivitymerchantmodify"></a>

### OperationActivityMerchantModify — 商户签约信息修改

- 官方接口：[`alipay.commerce.operation.activity.merchant.modify`](https://opendocs.alipay.com/pre-apis/063rs2)
- 路由：`POST /v3/alipay/commerce/operation/activity/merchant/modify`
- 请求类型：[OperationActivityMerchantModifyRequest](operation_request.go#L5)
- 响应类型：[OperationActivityMerchantModifyResponse](operation_response.go#L9)

商户签约信息修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `consumption_threshold` | `ConsumptionThreshold` | `*int64` | JSON body | 未标注 | 消费门槛，单位分 |
| `discount_amount` | `DiscountAmount` | `*int64` | JSON body | 未标注 | 优惠金额，单位分 |
| `type` | `Type` | `*string` | JSON body | 未标注 | NORMAL 常规类型，不限制, RESTRICTED 限制报名条件 |

<a id="api-operationactivitymerchantsign"></a>

### OperationActivityMerchantSign — 商户签约活动

- 官方接口：[`alipay.commerce.operation.activity.merchant.sign`](https://opendocs.alipay.com/pre-apis/06347n)
- 路由：`POST /v3/alipay/commerce/operation/activity/merchant/sign`
- 请求类型：[OperationActivityMerchantSignRequest](operation_request.go#L15)
- 响应类型：[OperationActivityMerchantSignResponse](operation_response.go#L12)

商户签约活动

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `consumption_threshold` | `ConsumptionThreshold` | `*int64` | JSON body | 未标注 | 消费门槛（单位分） |
| `discount_amount` | `DiscountAmount` | `*int64` | JSON body | 未标注 | 优惠金额，单位分 |
| `type` | `Type` | `*string` | JSON body | 未标注 | 报名类型，NORMAL 常规类型，不限制, RESTRICTED 限制报名条件 |

<a id="api-operationactivitymerchantunsign"></a>

### OperationActivityMerchantUnsign — 活动商户解约

- 官方接口：[`alipay.commerce.operation.activity.merchant.unsign`](https://opendocs.alipay.com/pre-apis/06336j)
- 路由：`POST /v3/alipay/commerce/operation/activity/merchant/unsign`
- 请求类型：[OperationActivityMerchantUnsignRequest](operation_request.go#L25)
- 响应类型：[OperationActivityMerchantUnsignResponse](operation_response.go#L15)

用户银行商户模式商户解约活动使用

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-operationbankactivityquery"></a>

### OperationBankActivityQuery — 银行活动信息查询

- 官方接口：[`alipay.commerce.operation.bank.activity.query`](https://opendocs.alipay.com/pre-apis/063ecd)
- 路由：`GET /v3/alipay/commerce/operation/bank/activity/query`
- 请求类型：[OperationBankActivityQueryRequest](operation_request.go#L28)
- 响应类型：[OperationBankActivityQueryResponse](operation_response.go#L18)

银行活动信息查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝openId |
| `user_id` | `UserId` | `*string` | query | 否 | 蚂蚁统一会员ID |
