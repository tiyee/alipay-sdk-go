# account 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/account`。本包提供 1 个支付宝 V3 REST 操作。

通过 `account.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AccountQuery`](#api-accountquery) | `alipay.fund.account.query` |

## 使用示例

- [AccountQuery](api_account_accountquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-accountquery"></a>

### AccountQuery — 支付宝资金账户资产查询接口

- 官方接口：[`alipay.fund.account.query`](https://opendocs.alipay.com/open-v3/77e2b925_alipay.fund.account.query)
- 路由：`GET /v3/alipay/fund/account/query`
- 请求类型：[FundAccountQueryRequest](account_request.go#L4)
- 响应类型：[FundAccountQueryResponse](account_response.go#L8)

可查询请求方的支付宝账户余额信息。

请求约束：FundAccountQueryRequest 是 alipay.fund.account.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_user_id` | `MerchantUserId` | `*string` | query | 否 | 商户会员的唯一标识。如果传入的user_id为虚拟账户userId，此字段必传并比对一致性。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 支付宝会员 id。 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | query | 否 | 支付宝openId |
| `account_product_code` | `AccountProductCode` | `*string` | query | 否 | 开户产品码。如果查询托管子户余额，必传且必须传入与开户时传入的值一致。 |
| `account_type` | `AccountType` | `*string` | query | 否 | 查询的账号类型，查询余额账户值为ACCTRANS_ACCOUNT。必填。 |
| `account_scene_code` | `AccountSceneCode` | `*string` | query | 否 | 开户场景码，与开户产品码不可同时传递。 |
| `ext_info` | `ExtInfo` | `*string` | query | 否 | JSON格式，传递业务扩展参数。 |
