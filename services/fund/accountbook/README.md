# accountbook 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/accountbook`。本包提供 5 个支付宝 V3 REST 操作。

通过 `accountbook.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AccountBookCreate`](#api-accountbookcreate) | `alipay.fund.accountbook.create` |
| [`AccountBookNotifyQuery`](#api-accountbooknotifyquery) | `alipay.fund.accountbook.notify.query` |
| [`AccountBookNotifySubscribe`](#api-accountbooknotifysubscribe) | `alipay.fund.accountbook.notify.subscribe` |
| [`AccountBookNotifyUnsubscribe`](#api-accountbooknotifyunsubscribe) | `alipay.fund.accountbook.notify.unsubscribe` |
| [`AccountBookQuery`](#api-accountbookquery) | `alipay.fund.accountbook.query` |

## 使用示例

- [AccountBookCreate](api_accountbook_accountbookcreate_example_test.go)
- [AccountBookNotifyQuery](api_accountbook_accountbooknotifyquery_example_test.go)
- [AccountBookNotifySubscribe](api_accountbook_accountbooknotifysubscribe_example_test.go)
- [AccountBookNotifyUnsubscribe](api_accountbook_accountbooknotifyunsubscribe_example_test.go)
- [AccountBookQuery](api_accountbook_accountbookquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-accountbookcreate"></a>

### AccountBookCreate — 资金记账本开通

- 官方接口：[`alipay.fund.accountbook.create`](https://opendocs.alipay.com/apis/0cjl4u)
- 路由：`POST /v3/alipay/fund/accountbook/create`
- 请求类型：[FundAccountBookCreateRequest](accountbook_request.go#L4)
- 响应类型：[FundAccountBookCreateResponse](accountbook_response.go#L8)

资金记账本开通

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | JSON格式，传递业务扩展参数 |
| `merchant_user_id` | `MerchantUserId` | `*string` | JSON body | 未标注 | 外部商户系统会员的唯一标识，自定义传入 |
| `merchant_user_type` | `MerchantUserType` | `*string` | JSON body | 未标注 | 外部商户用户类型:BUSINESS_EMPLOYEE |
| `scene_code` | `SceneCode` | `*string` | JSON body | 未标注 | 资金记账本的业务场景 |

<a id="api-accountbooknotifyquery"></a>

### AccountBookNotifyQuery — 记账本通知订阅关系查询接口

- 官方接口：[`alipay.fund.accountbook.notify.query`](https://opendocs.alipay.com/open-v3/495b0416_alipay.fund.accountbook.notify.query)
- 路由：`GET /v3/alipay/fund/accountbook/notify/query`
- 请求类型：[FundAccountBookNotifyQueryRequest](accountbook_request.go#L40)
- 响应类型：[FundAccountBookNotifyQueryResponse](accountbook_response.go#L15)

ISV通过此接口查询指定记账本的入金和出金通知的订阅关系

请求约束：FundAccountBookNotifyQueryRequest 是 alipay.fund.accountbook.notify.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `product_code` | `ProductCode` | `*string` | query | 否 | 产品码。固定为SATF_FUND_BOOK |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 场景码。固定为DEFAULT |
| `account_book_id` | `AccountBookId` | `*string` | query | 否 | 记账本ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传； |

<a id="api-accountbooknotifysubscribe"></a>

### AccountBookNotifySubscribe — 记账本通知订阅接口

- 官方接口：[`alipay.fund.accountbook.notify.subscribe`](https://opendocs.alipay.com/open-v3/37df5c87_alipay.fund.accountbook.notify.subscribe)
- 路由：`POST /v3/alipay/fund/accountbook/notify/subscribe`
- 请求类型：[FundAccountBookNotifySubscribeRequest](accountbook_request.go#L16)
- 响应类型：[FundAccountBookNotifySubscribeResponse](accountbook_response.go#L27)

ISV通过此接口订阅指定记账本的入金和出金通知。目前只有非OpenAPI方式的充值通知（比如大额来账、贷款入金）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_book_id` | `AccountBookId` | `*string` | JSON body | 未标注 | 记账本ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传； |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码。固定为DEFAULT |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码。固定为SATF_FUND_BOOK |

<a id="api-accountbooknotifyunsubscribe"></a>

### AccountBookNotifyUnsubscribe — 记账本通知取消订阅接口

- 官方接口：[`alipay.fund.accountbook.notify.unsubscribe`](https://opendocs.alipay.com/open-v3/1923216a_alipay.fund.accountbook.notify.unsubscribe)
- 路由：`POST /v3/alipay/fund/accountbook/notify/unsubscribe`
- 请求类型：[FundAccountBookNotifyUnsubscribeRequest](accountbook_request.go#L28)
- 响应类型：[FundAccountBookNotifyUnsubscribeResponse](accountbook_response.go#L39)

ISV通过此接口取消订阅指定记账本的入金和出金通知

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_book_id` | `AccountBookId` | `*string` | JSON body | 未标注 | 记账本ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传； |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码。固定为DEFAULT |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码。固定为SATF_FUND_BOOK |

<a id="api-accountbookquery"></a>

### AccountBookQuery — 资金记账本的信息查询

- 官方接口：[`alipay.fund.accountbook.query`](https://opendocs.alipay.com/open-v3/f072b3ba_alipay.fund.accountbook.query)
- 路由：`GET /v3/alipay/fund/accountbook/query`
- 请求类型：[FundAccountBookQueryRequest](accountbook_request.go#L52)
- 响应类型：[FundAccountBookQueryResponse](accountbook_response.go#L51)

资金记账本的信息查询

请求约束：FundAccountBookQueryRequest 是 alipay.fund.accountbook.query 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_book_id` | `AccountBookId` | `*string` | query | 否 | 记账账簿id |
| `merchant_user_id` | `MerchantUserId` | `*string` | query | 否 | 商户会员的唯一标识，如果传入account_book_id此字段必传并比对一致性。 |
| `scene_code` | `SceneCode` | `*string` | query | 否 | 资金记账本的开通场景码 |
| `ext_info` | `ExtInfo` | `*string` | query | 否 | 传递业务扩展参数。 注意：ext_info字段本质上是String类型，所以传递 的是转义后的json字符串。 |
