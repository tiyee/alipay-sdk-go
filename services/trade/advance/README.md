# advance 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/advance`。本包提供 1 个支付宝 V3 REST 操作。

通过 `advance.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AdvanceConsult`](#api-advanceconsult) | `alipay.trade.advance.consult` |

## 使用示例

- [AdvanceConsult](api_advance_advanceconsult_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-advanceconsult"></a>

### AdvanceConsult — 订单咨询服务

- 官方接口：[`alipay.trade.advance.consult`](https://opendocs.alipay.com/open-v3/11cb1f6e_alipay.trade.advance.consult)
- 路由：`GET /v3/alipay/trade/advance/consult`
- 请求类型：[AdvanceConsultRequest](advance_request.go#L10)
- 响应类型：[AdvanceConsultResponse](advance_response.go#L4)

商户通过此接口咨询，当前用户是否满足垫资服务条件 或者 针对单笔订单的风险评估进行预咨询服务。

请求约束：AdvanceConsultRequest 是 alipay.trade.advance.consult（订单咨询，
V3 GET）的请求：垫资咨询场景 alipay_user_id 必选；单笔订单风险
预评估场景（consult_scene=ORDER_RISK_EVALUATION）必传非空协议号。若传预估金额，其范围为 0.01 至 100000000 元。响应顶层的待还总额按字符串读取，订单明细内的待还金额按整数读取。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `alipay_user_id` | `AlipayUserID` | `*string` | query | 否 | 支付宝用户 id，2088开头的16为数字。 除单笔订单风险预评估场景(即consult_scene指定ORDER_RISK_EVALUATION)外，其他场景必选。<br>SDK 字段说明：支付宝用户 id（2088 开头 16 位数字）；<br>除订单风险预评估场景外必选。 |
| `consult_scene` | `ConsultScene` | `*string` | query | 否 | 订单咨询类型，用以选择咨询的服务。不传时默认为垫资咨询，其他的场景需要按照对接的服务传入指定的值。 ORDER_RISK_EVALUATION：表示单笔订单风险预评估。<br>SDK 字段说明：订单咨询类型；默认为垫资咨询，<br>ORDER_RISK_EVALUATION 表示单笔订单风险预评估。 |
| `industry_product_code` | `IndustryProductCode` | `*string` | query | 否 | 行业产品信息，咨询时会从该产品对应的销售方案中获取相关垫资规则配置。 |
| `sub_merchant_id` | `SubMerchantID` | `*string` | query | 否 | 子商户id<br>SDK 字段说明：子商户 id。 |
| `sub_merchant_type` | `SubMerchantType` | `*string` | query | 否 | 子商户类型 |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，在进行单笔订单风险评估预咨询时必传。<br>SDK 字段说明：用户签约协议号；单笔订单风险预评估时必传。 |
| `estimated_order_amount` | `EstimatedOrderAmount` | `*string` | query | 否 | 预估订单总金额，单位为元，精确到小数点后两位，单笔订单风险预评估时使用，取值范围[0.01,100000000]。<br>SDK 字段说明：预估订单总金额；单笔订单风险预评估时使用。 |
| `out_trade_no` | `OutTradeNo` | `*string` | query | 否 | 商户请求时的外部订单号，在花芝场景下非空。 |
| `need_user_risk_prediction` | `NeedUserRiskPrediction` | `*string` | query | 否 | 是否需要进行用户风险预测<br>SDK 字段说明：标识是否需要进行用户风险预测。 |
| `trade_digest` | `TradeDigest` | `*string` | query | 否 | 交易信息摘要，需要进行用户风险预测时可传，用于结合场景的更加精细的预测。具体信息结构服务接入前咨询约定。<br>SDK 字段说明：交易信息摘要；需要用户风险预测时可传。 |
| `user_out_binding_phone` | `UserOutBindingPhone` | `*string` | query | 否 | 用户在外部域账号系统绑定的手机号，需要进行用户风险预测时必传，用于预测用户是否面临被二次放号的场景。使用时需要对请求加密，加密后请求在公网传输时为加密文本。详见：https://opendocs.alipay.com/open/common/104567；https://opendocs.alipay.com/isv/grefvl/getaes<br>SDK 字段说明：用户在外部域账号系统绑定的手机号；<br>需要用户风险预测时必传，按官方要求加密后传输。 |

注意事项：

- 垫资咨询场景 AlipayUserID 必选；风险预评估场景需 ConsultScene 与
AgreementNo 等参数。
