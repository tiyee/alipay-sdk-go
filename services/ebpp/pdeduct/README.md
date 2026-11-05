# pdeduct 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/pdeduct`。本包提供 3 个支付宝 V3 REST 操作。

通过 `pdeduct.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`sign`](sign/README.md) | 4 | `PdeductSignAdd` / `PdeductSignCancel` / `PdeductSignQuery` / `PdeductSignValidate` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`PdeductAsyncPay`](#api-pdeductasyncpay) | `alipay.ebpp.pdeduct.async.pay` |
| [`PdeductBillPayStatus`](#api-pdeductbillpaystatus) | `alipay.ebpp.pdeduct.bill.pay.status` |
| [`PdeductPay`](#api-pdeductpay) | `alipay.ebpp.pdeduct.pay` |

## 使用示例

- [PdeductAsyncPay](api_pdeduct_pdeductasyncpay_example_test.go)
- [PdeductBillPayStatus](api_pdeduct_pdeductbillpaystatus_example_test.go)
- [PdeductPay](api_pdeduct_pdeductpay_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-pdeductasyncpay"></a>

### PdeductAsyncPay — 公共事业缴费直连代扣异步扣款支付接口

- 官方接口：[`alipay.ebpp.pdeduct.async.pay`](https://opendocs.alipay.com/pre-apis/05xmda)
- 路由：`POST /v3/alipay/ebpp/pdeduct/async/pay`
- 请求类型：[PdeductAsyncPayRequest](pdeduct_request.go#L4)
- 响应类型：[PdeductAsyncPayResponse](pdeduct_response.go#L4)

直连代扣机构根据用户个人签约协议，按期按账单请求从用户账户扣款的接口。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_channel` | `AgentChannel` | `*string` | JSON body | 未标注 | 分配给外部机构发起扣款时的渠道码。朗新为LANGXIN |
| `agent_code` | `AgentCode` | `*string` | JSON body | 未标注 | 二级渠道码，预留字段 |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝代扣协议Id |
| `bill_date` | `BillDate` | `*string` | JSON body | 未标注 | 账期 |
| `bill_key` | `BillKey` | `*string` | JSON body | 未标注 | 户号 |
| `extend_field` | `ExtendField` | `*string` | JSON body | 未标注 | 扩展字段 |
| `fine_amount` | `FineAmount` | `*string` | JSON body | 未标注 | 滞纳金 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注信息 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户UserId在应用AppId下的唯一用户标识 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户外部业务流水号 |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 扣款金额，支付总金额，包含滞纳金 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 商户partnerId |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户ID |

<a id="api-pdeductbillpaystatus"></a>

### PdeductBillPayStatus — 缴费直连代扣订单支付状态查询

- 官方接口：[`alipay.ebpp.pdeduct.bill.pay.status`](https://opendocs.alipay.com/open-v3/c7416677_alipay.ebpp.pdeduct.bill.pay.status)
- 路由：`GET /v3/alipay/ebpp/pdeduct/bill/pay/status`
- 请求类型：[PdeductBillPayStatusRequest](pdeduct_request.go#L64)
- 响应类型：[PdeductBillPayStatusResponse](pdeduct_response.go#L14)

缴费直连代扣

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | query | 否 | 支付宝用户ID |
| `open_id` | `OpenId` | `*string` | query | 否 | 用户UserId在应用AppId下的唯一用户标识 |
| `out_order_no` | `OutOrderNo` | `*string` | query | 否 | 商户代扣业务流水 |

<a id="api-pdeductpay"></a>

### PdeductPay — 公共事业缴费直连代扣扣款支付接口

- 官方接口：[`alipay.ebpp.pdeduct.pay`](https://opendocs.alipay.com/open-v3/fea30a1b_alipay.ebpp.pdeduct.pay)
- 路由：`POST /v3/alipay/ebpp/pdeduct/pay`
- 请求类型：[PdeductPayRequest](pdeduct_request.go#L34)
- 响应类型：[PdeductPayResponse](pdeduct_response.go#L32)

直连代扣机构根据用户个人签约协议，按期按账单请求从用户账户扣款的接口。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_channel` | `AgentChannel` | `*string` | JSON body | 未标注 | 渠道码，如用户通过机构通过服务窗进来签约则是PUBLICFORM，此值可随意传，只要不空就行 |
| `agent_code` | `AgentCode` | `*string` | JSON body | 未标注 | 二级渠道码，预留字段 |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝代扣协议Id |
| `bill_date` | `BillDate` | `*string` | JSON body | 未标注 | 账期 |
| `bill_key` | `BillKey` | `*string` | JSON body | 未标注 | 户号，缴费单位用于标识每一户的唯一性的 |
| `extend_field` | `ExtendField` | `*string` | JSON body | 未标注 | 扩展参数。必须以key value形式定义， 转为json为格式：{"key1":"value1","key2":"value2", "key3":"value3","key4":"value4"} 后端会直接转换为MAP对象，转换异常会报参数格式错误。 instUniqueNo,机构唯一流水号 |
| `fine_amount` | `FineAmount` | `*string` | JSON body | 未标注 | 滞纳金 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注信息 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户UserId在应用AppId下的唯一用户标识 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户外部业务流水号 |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 扣款金额，支付总金额，包含滞纳金 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 商户PartnerId |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户ID |
