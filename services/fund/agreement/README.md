# agreement 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/agreement`。本包提供 2 个支付宝 V3 REST 操作。

通过 `agreement.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AgreementQuotaModify`](#api-agreementquotamodify) | `alipay.fund.agreement.quota.modify` |
| [`AgreementQuotaQuery`](#api-agreementquotaquery) | `alipay.fund.agreement.quota.query` |

## 使用示例

- [AgreementQuotaModify](api_agreement_agreementquotamodify_example_test.go)
- [AgreementQuotaQuery](api_agreement_agreementquotaquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-agreementquotamodify"></a>

### AgreementQuotaModify — 额度设置接口

- 官方接口：[`alipay.fund.agreement.quota.modify`](https://opendocs.alipay.com/apis/05xo4q)
- 路由：`POST /v3/alipay/fund/agreement/quota/modify`
- 请求类型：[FundAgreementQuotaModifyRequest](agreement_request.go#L4)
- 响应类型：[FundAgreementQuotaModifyResponse](agreement_response.go#L4)

基于扣款合约设置付款方额度

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_quota_modify_list` | `AgreementQuotaModifyList` | [`*[]AgreementQuotaModifyList`](#model-agreementquotamodifylist) | JSON body | 未标注 | 协议额度设置信息 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码，非固定值，使用的对应的资金场景码保持一致 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 额度产品码，非固定值，使用的对应的资金产品码保持一致 |

<a id="api-agreementquotaquery"></a>

### AgreementQuotaQuery — 额度查询接口

- 官方接口：[`alipay.fund.agreement.quota.query`](https://opendocs.alipay.com/apis/05ximq)
- 路由：`POST /v3/alipay/fund/agreement/quota/query`
- 请求类型：[FundAgreementQuotaQueryRequest](agreement_request.go#L14)
- 响应类型：[FundAgreementQuotaQueryResponse](agreement_response.go#L10)

查询基于扣款协议的额度配置查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no_list` | `AgreementNoList` | `*[]string` | JSON body | 未标注 | 授权协议号 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 场景码，非固定值，使用的对应的资金场景码保持一致 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 额度产品码，非固定值，使用的对应的资金产品码保持一致 |

## 嵌套请求参数

<a id="model-agreementquotamodifylist"></a>

### AgreementQuotaModifyList

类型定义：[AgreementQuotaModifyList](agreement_model.go#L14)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | 未标注 | 授权协议号 |
| `quota_details` | `QuotaDetails` | [`*[]ModifyQuotaDetails`](#model-modifyquotadetails) | 未标注 | 设置额度明细 |

<a id="model-modifyquotadetails"></a>

### ModifyQuotaDetails

类型定义：[ModifyQuotaDetails](agreement_model.go#L22)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `quota_amount` | `QuotaAmount` | `*string` | 未标注 | 具体额度，单位：元。如金额为空，表示删除已有的额度设置，无已有额度设置则忽略。 |
| `quota_dimension` | `QuotaDimension` | `*string` | 未标注 | 额度维度 MONTH/DAY/SINGLE 分别代表月、日、单次 |
| `role` | `Role` | `*string` | 未标注 | PAYER/PAYEE 额度管控的角色，收or付款方，目前只支持付款方 |
