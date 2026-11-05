# trans 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/fund/trans`。本包提供 7 个支付宝 V3 REST 操作。

通过 `trans.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`OrderQuery`](#api-orderquery) | `alipay.fund.trans.order.query` |
| [`CommonQuery`](#api-commonquery) | `alipay.fund.trans.common.query` |
| [`PayeeBindQuery`](#api-payeebindquery) | `alipay.fund.trans.payee.bind.query` |
| [`UniTransfer`](#api-unitransfer) | `alipay.fund.trans.uni.transfer` |
| [`ToAccountTransfer`](#api-toaccounttransfer) | `alipay.fund.trans.toaccount.transfer` |
| [`Pay`](#api-pay) | `alipay.fund.trans.pay` |
| [`Refund`](#api-refund) | `alipay.fund.trans.refund` |

## 使用示例

- [OrderQuery](api_trans_orderquery_example_test.go)
- [CommonQuery](api_trans_commonquery_example_test.go)
- [PayeeBindQuery](api_trans_payeebindquery_example_test.go)
- [UniTransfer](api_trans_unitransfer_example_test.go)
- [ToAccountTransfer](api_trans_toaccounttransfer_example_test.go)
- [Pay](api_trans_pay_example_test.go)
- [Refund](api_trans_refund_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

资金支出类接口需特别核实目标产品的证书模式要求与签约权限，不能仅凭生成 Operation 的认证配置判断可调用性。依据：[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

## 接口说明与请求参数

<a id="api-orderquery"></a>

### OrderQuery — 查询转账订单接口

- 官方接口：[`alipay.fund.trans.order.query`](https://opendocs.alipay.com/open-v3/f00381dd_alipay.fund.trans.order.query)
- 路由：`GET /v3/alipay/fund/trans/order/query`
- 请求类型：[OrderQueryRequest](trans_request.go#L10)
- 响应类型：[OrderQueryResponse](trans_response.go#L4)

商户可通过该接口查询转账订单的状态、支付时间等相关信息，主要应用于B2C转账订单查询的场景

请求约束：OrderQueryRequest 是 alipay.fund.trans.order.query（转账订单查询，
V3 GET）的请求：order_id 与 out_biz_no 二选一，同时提供时优先
使用 order_id。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderID` | `*string` | query | 否 | 支付宝转账单据号：和商户转账唯一订单号不能同时为空。当和商户转账唯一订单号同时提供时，将用本参数进行查询，忽略商户转账唯一订单号。<br>SDK 字段说明：支付宝转账单据号；与 OutBizNo 不能同时为空。 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 商户转账唯一订单号：发起转账来源方定义的转账单据ID。<br> 和支付宝转账单据号不能同时为空。当和支付宝转账单据号同时提供时，将用支付宝转账单据号进行查询，忽略本参数。<br>SDK 字段说明：商户转账唯一订单号；与 OrderID 不能同时为空。 |

注意事项：

- 查询成功不代表到账；DEALING/REFUND 等状态需按业务继续跟踪。

<a id="api-commonquery"></a>

### CommonQuery — 转账业务单据查询接口

- 官方接口：[`alipay.fund.trans.common.query`](https://opendocs.alipay.com/open-v3/bcbe6351_alipay.fund.trans.common.query?scene=common&pathHash=36f664d9)
- 路由：`GET /v3/alipay/fund/trans/common/query`
- 请求类型：[CommonQueryRequest](trans_request.go#L21)
- 响应类型：[CommonQueryResponse](trans_response.go#L25)

商户可通过该接口查询转账业务单据的状态，主要应用于统一转账接口(alipay.fund.trans.uni.transfer)、无线转账接口(alipay.fund.trans.app.pay)、单笔转账到支付宝账户接口（alipay.fund.trans.toaccount.transfer）

请求约束：CommonQueryRequest 是 alipay.fund.trans.common.query（转账单据通用
查询，V3 GET）的请求：out_biz_no、order_id、pay_fund_order_id
三者不能同时为空；传入 out_biz_no 时 product_code 与 biz_scene
必传。三者同时传入时优先级：pay_fund_order_id &gt; order_id &gt; out_biz_no。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderID` | `*string` | query | 否 | 支付宝转账单据号。 本参数和out_biz_no（商户转账唯一订单号）、pay_fund_order_id（支付宝支付资金流水号）三者不能同时为空。 当三者同时传入时，将用pay_fund_order_id（支付宝支付资金流水号）进行查询，忽略其余两者； 当本参数和pay_fund_order_id（支付宝支付资金流水号）同时提供时，将用支付宝支付资金流水号进行查询，忽略本参数； 当本参数和out_biz_no（商户转账唯一订单号）同时提供时，将用本参数进行查询，忽略商户转账唯一订单号。 |
| `pay_fund_order_id` | `PayFundOrderID` | `*string` | query | 否 | 支付宝支付资金流水号。本参数和支付宝转账单据号、商户转账唯一订单号三者不能同时为空。 当本参数和out_biz_no（商户转账唯一订单号）、order_id（支付宝转账单据号）同时提供时，将用本参数进行查询，忽略其余两者； 当本参数和order_id（支付宝转账单据号）同时提供时，将用本参数进行查询，忽略支付宝转账单据号； 当本参数和out_biz_no（商户转账唯一订单号）同时提供时，将用本参数进行查询，忽略商户转账唯一订单号。 |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 商户转账唯一订单号，发起转账来源方定义的转账单据ID。 本参数和order_id（支付宝转账单据号）、pay_fund_order_id（支付宝支付资金流水号）三者不能同时为空。 当三者同时传入时，将用pay_fund_order_id（支付宝支付资金流水号）进行查询，忽略其余两者； 当本参数和支付宝转账单据号同时提供时，将用支付宝转账单据号进行查询，忽略本参数。 |
| `product_code` | `ProductCode` | `*string` | query | 否 | 销售产品码，商家和支付宝签约的产品码，如果传递了out_biz_no则该字段为必传。可传值如下： STD_RED_PACKET：现金红包 TRANS_ACCOUNT_NO_PWD：单笔无密转账到支付宝账户 TRANS_BANKCARD_NO_PWD：单笔无密转账到银行卡<br>SDK 字段说明：销售产品码；传入 out_biz_no 时必传。 |
| `biz_scene` | `BizScene` | `*string` | query | 否 | 描述特定的业务场景，如果传递了out_biz_no则该字段为必传。可取的业务场景如下：<br> PERSONAL_PAY：C2C现金红包-发红包；<br> PERSONAL_COLLECTION：C2C现金红包-领红包；<br> REFUND：C2C现金红包-红包退回；<br> DIRECT_TRANSFER：B2C现金红包、单笔无密转账<br>SDK 字段说明：业务场景；传入 out_biz_no 时必传。 |

<a id="api-payeebindquery"></a>

### PayeeBindQuery — 资金收款账号绑定关系查询

- 官方接口：[`alipay.fund.trans.payee.bind.query`](https://opendocs.alipay.com/open-v3/03fe4a0d_alipay.fund.trans.payee.bind.query?scene=common&pathHash=377d1f29)
- 路由：`GET /v3/alipay/fund/trans/payee/bind/query`
- 请求类型：[PayeeBindQueryRequest](trans_request.go#L36)
- 响应类型：[PayeeBindQueryResponse](trans_response.go#L79)

资金收款账号绑定关系查询

请求约束：PayeeBindQueryRequest 是 alipay.fund.trans.payee.bind.query
（收款账号绑定查询，V3 GET）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `identity` | `Identity` | `string` | query | 否 | 证件号 |
| `identity_type` | `IdentityType` | `*string` | query | 否 | 证件类型,暂仅支持 IDENTITY_CARD （身份证）<br>SDK 字段说明：证件类型；暂仅支持 IDENTITY_CARD（身份证）。 |

<a id="api-unitransfer"></a>

### UniTransfer — 单笔转账接口

- 官方接口：[`alipay.fund.trans.uni.transfer`](https://opendocs.alipay.com/open-v3/26a22ad7_alipay.fund.trans.uni.transfer?scene=common%2Cca56bca529e64125a2786703c6192d41%2C66dd06f5a923403393b85de68d3c0055&pathHash=9b52afe2)
- 路由：`POST /v3/alipay/fund/trans/uni/transfer`
- 请求类型：[UniTransferRequest](trans_request.go#L83)
- 响应类型：[UniTransferResponse](trans_response.go#L87)

单笔转账接口是基于支付宝的资金处理能力，为了满足支付宝商家的转账需求，针对有部分开发能力的商家，提供通过API接口完成企业自身支付宝账户到支付宝账户的转账功能。 适用于：C2C现金红包的红包领取、B2C现金红包的红包发放、B2C单笔转账到支付宝账户等。
说明：本接口为公用接口，参数说明不针对特定产品。不同产品入参请参考对应产品文档。

请求约束：UniTransferRequest 是 alipay.fund.trans.uni.transfer（单笔转账接口
或现金红包）的请求。out_biz_no、trans_amount、product_code、
payee_info 必填；产品码决定取值范围。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `string` | JSON body | 未标注 | 商家侧唯一订单号，由商家自定义。对于不同转账请求，商家需保证该订单号在自身系统唯一。<br>SDK 字段说明：商家侧唯一订单号；不同转账请求需保证唯一。 |
| `trans_amount` | `TransAmount` | `string` | JSON body | 未标注 | 订单总金额，单位为元，精确到小数点后两位，STD_RED_PACKET产品取值范围[0.01,100000000]； TRANS_ACCOUNT_NO_PWD产品取值范围[0.1,100000000]<br>SDK 字段说明：订单总金额，单位为元，精确到小数点后两位：<br>TRANS_ACCOUNT_NO_PWD 取值范围 [0.1,100000000]，<br>STD_RED_PACKET 取值范围 [0.01,100000000]。 |
| `product_code` | `ProductCode` | `string` | JSON body | 未标注 | 业务产品码， 单笔无密转账到支付宝账户固定为: TRANS_ACCOUNT_NO_PWD； 收发现金红包固定为: STD_RED_PACKET；<br>SDK 字段说明：业务产品码：单笔无密转账到支付宝账户固定为<br>TRANS_ACCOUNT_NO_PWD，收发现金红包固定为 STD_RED_PACKET。 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 描述特定的业务场景，可传的参数如下： DIRECT_TRANSFER：单笔无密转账到支付宝，B2C现金红包; PERSONAL_COLLECTION：C2C现金红包-领红包<br>SDK 字段说明：特定的业务场景：DIRECT_TRANSFER（单笔无密转账到<br>支付宝、B2C 现金红包）、PERSONAL_COLLECTION（C2C 现金红包-领红包）。 |
| `payee_info` | `PayeeInfo` | [`*Participant`](#model-participant) | JSON body | 未标注 | 收款方信息。 |
| `payer_info` | `PayerInfo` | [`*Participant`](#model-participant) | JSON body | 未标注 | 付款方信息。 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 转账业务的标题，用于在支付宝用户的账单里显示 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 业务备注 |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | 转账业务请求的扩展参数，支持传入的扩展参数如下： sub_biz_scene 子业务场景，红包业务必传，C2C现金红包、B2C现金红包均需传入，取值REDPACKET<br>SDK 字段说明：转账业务请求的扩展参数（JSON 文本，<br>如子业务场景 sub_biz_scene=REDPACKET）。 |
| `passback_params` | `PassbackParams` | `*string` | JSON body | 未标注 | 公用回传参数，异步通知时原样回传。 |
| `original_order_id` | `OriginalOrderID` | `*string` | JSON body | 未标注 | 原支付宝业务单号。 |
| `transfer_scene_name` | `TransferSceneName` | `*string` | JSON body | 未标注 | 该笔转账使用的转账场景，商户可前往【[支付宝商家平台-资金管理-转账场景](https://b.alipay.com/page/bizfund/transferScene/list)】中查看/申明转账场景。<br> 目前支持以下枚举值：现金营销、企业退款、佣金报酬、业务结算、二手回收、公益补助、行政补贴和退款、保险理赔<br> 注意：26年及以后新接入商户必须传入该字段。<br>SDK 字段说明：该笔转账使用的转账场景，见商家平台<br>资金管理-转账场景。 |
| `transfer_scene_report_infos` | `TransferSceneReportInfos` | [`[]TransferSceneReportInfo`](#model-transferscenereportinfo) | JSON body | 未标注 | 转账场景信息，包含转账场景信息类型和转账场景信息描述。需按该笔转账使用的转账场景准确填写<br> 注意：26年及以后新接入商户必须传入该字段。<br>SDK 字段说明：转账场景信息；26 年及以后新接入<br>商户必须传入，字段说明见官方转账场景信息字段说明。 |
| `mutiple_currency_detail` | `MultipleCurrencyDetail` | [`*MultipleCurrencyDetail`](trans_model.go#L103) | JSON body | 未标注 | 多币种转账金额与币种明细。 |
| `sign_data` | `SignData` | [`*SignData`](trans_model.go#L121) | JSON body | 未标注 | 转账请求的原始签名信息。 |

注意事项：

- 资金写操作：SDK 不自动重试；超时或状态未知时以订单查询确认；
该接口要求证书模式。`UniTransfer` 会在发出 HTTP 请求前拒绝公钥凭证。产品权限由支付宝网关确认。

<a id="api-toaccounttransfer"></a>

### ToAccountTransfer — 单笔转账到支付宝账户接口

- 官方接口：[`alipay.fund.trans.toaccount.transfer`](https://opendocs.alipay.com/open-v3/53a4ce8e_alipay.fund.trans.toaccount.transfer)
- 路由：`POST /v3/alipay/fund/trans/toaccount/transfer`
- 请求类型：[ToAccountTransferRequest](trans_request.go#L126)
- 响应类型：[ToAccountTransferResponse](trans_response.go#L112)

单笔转账到支付宝账户接口是基于支付宝的资金处理能力，为了满足支付宝商家向其他支付宝账户转账的需求，针对有部分开发能力的商家，提供通过API接口完成支付宝账户间的转账的功能。
该接口适用行业较广，比如商家间的货款结算，商家给个人用户发放佣金等。（目前已对该接口能力做了升级，后续建议调用新接口alipay.fund.trans.uni.transfer）

请求约束：ToAccountTransferRequest 是 alipay.fund.trans.toaccount.transfer
（单笔转账到支付宝账户）的请求。out_biz_no、amount、payee_account、
payee_type 必填；付款方和收款方不能是同一个账户。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `string` | JSON body | 未标注 | 商户转账唯一订单号。发起转账来源方定义的转账单据ID，用于将转账回执通知给来源方。<br> 不同来源方给出的ID可以重复，同一个来源方必须保证其ID的唯一性。<br> 只支持半角英文、数字，及“-”、“_”。<br>SDK 字段说明：商户转账唯一订单号；同一来源方必须保证唯一，<br>只支持半角英文、数字及 “-”“_”。 |
| `amount` | `TransAmount` | `string` | JSON body | 未标注 | 转账金额，单位：元。<br> 只支持2位小数，小数点前最大支持13位，金额必须大于等于0.1元。 <br> 最大转账金额以实际签约的限额为准。<br>SDK 字段说明：转账金额，单位为元，只支持 2 位小数，<br>金额必须大于等于 0.1 元。 |
| `payee_account` | `PayeeAccount` | `string` | JSON body | 未标注 | 收款方账户。与payee_type配合使用。付款方和收款方不能是同一个账户。<br>SDK 字段说明：收款方账户，与 PayeeType 配合使用。 |
| `payee_type` | `PayeeType` | [`PayeeAccountType`](trans_model.go#L82) | JSON body | 未标注 | 收款方账户类型。可取值： 1、ALIPAY_USERID：支付宝账号对应的支付宝唯一用户号。以2088开头的16位纯数字组成。 2、ALIPAY_LOGONID：支付宝登录号，支持邮箱和手机号格式。 2、ALIPAY_OPENID：支付宝openid |
| `payee_real_name` | `PayeeRealName` | `*string` | JSON body | 未标注 | 收款方真实姓名（最长支持100个英文/50个汉字）。<br> 如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。<br>SDK 字段说明：收款方真实姓名；非空时校验收款账户实名一致。 |
| `payer_real_name` | `PayerRealName` | `*string` | JSON body | 未标注 | 付款方真实姓名（最长支持100个英文/50个汉字）。<br> 如果本参数不为空，则会校验该账户在支付宝登记的实名是否与付款方真实姓名一致。<br>SDK 字段说明：付款方真实姓名；非空时校验付款账户实名一致。 |
| `payer_show_name` | `PayerShowName` | `*string` | JSON body | 未标注 | 付款方姓名（最长支持100个英文/50个汉字）。显示在收款方的账单详情页。如果该字段不传，则默认显示付款方的支付宝认证姓名或单位名称。<br>SDK 字段说明：付款方显示名称，展示在收款方的账单详情页。 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 转账备注（支持200个英文/100个汉字）。<br> 当付款方为企业账户，且转账金额达到（大于等于）50000元，remark不能为空。收款方可见，会展示在收款用户的收支详情中。<br>SDK 字段说明：转账备注；企业账户单笔 ≥5 万元时不能为空。 |
| `ext_param` | `ExtParam` | `*string` | JSON body | 未标注 | 扩展参数，json字符串格式，目前仅支持的key=order_title，表示收款方的转账账单标题，value可以根据自己的业务定制。<br>SDK 字段说明：扩展参数（JSON 字符串，目前仅支持 order_title）。 |

注意事项：

- 金额必须 ≥0.1 元；写操作，SDK 不自动重试。

<a id="api-pay"></a>

### Pay — 资金支付接口

- 官方接口：[`alipay.fund.trans.pay`](https://opendocs.alipay.com/pre-apis/05xggg)
- 路由：`POST /v3/alipay/fund/trans/pay`
- 请求类型：[PayRequest](trans_request.go#L151)
- 响应类型：[PayResponse](trans_response.go#L123)

付款方通过协议、二维码、声波、人脸等方式，通过本接口提交给支付宝直接完成支付。

请求约束：PayRequest 是 alipay.fund.trans.pay（资金支付/收银台转账）的请求。
out_biz_no、trans_amount、product_code、payee_info 必填。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `string` | JSON body | 未标注 | 商户端的唯一订单号，对于同一笔转账请求，商户需保证该订单号唯一。<br>SDK 字段说明：商户端的唯一订单号；同一笔转账请求需保证唯一。 |
| `trans_amount` | `TransAmount` | `string` | JSON body | 未标注 | 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]<br>SDK 字段说明：订单总金额，单位为元，精确到小数点后两位，<br>取值范围 [0.01,100000000]。 |
| `product_code` | `ProductCode` | `string` | JSON body | 未标注 | 销售产品码，商家和支付宝签约的产品码 |
| `payee_info` | `PayeeInfo` | [`*Participant`](#model-participant) | JSON body | 未标注 | 收款方信息。 |
| `payer_info` | `PayerInfo` | [`*Participant`](#model-participant) | JSON body | 未标注 | 付款方信息。 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 订单的标题，用于在收银台和消费记录展示<br>SDK 字段说明：订单标题，用于在收银台和消费记录展示。 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 业务场景，比如群收款、红包等<br>SDK 字段说明：业务场景（如群收款、红包等）。 |
| `auth_info` | `AuthInfo` | [`*AuthInfo`](#model-authinfo) | JSON body | 未标注 | 协议授权信息。 |
| `time_expire` | `TimeExpire` | `*string` | JSON body | 未标注 | 绝对超时时间，格式为yyyy-MM-dd HH:mm<br>SDK 字段说明：绝对超时时间（yyyy-MM-dd HH:mm）。 |
| `refund_time_expire` | `RefundTimeExpire` | `*string` | JSON body | 未标注 | 退款超时时间，格式yyyy-MM-dd HH:mm。到指定时间后，系统会自动触发退款，并原路退回到付款账户。如果指定了退款时间，必须早于销售方案里设置的最晚退款时间。<br>SDK 字段说明：退款超时时间（yyyy-MM-dd HH:mm）；<br>到期后系统自动原路退款。 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 备注 |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | JSON格式，传递业务扩展参数，使用前请与支付宝工程师联系！<br>SDK 字段说明：业务扩展参数（JSON 文本，使用前与支付宝确认）。 |
| `passback_params` | `PassbackParams` | `*string` | JSON body | 未标注 | 回传参数，如果请求时传递了该参数，则异步通知商户时会回传该参数。<br>SDK 字段说明：回传参数，异步通知时原样回传。 |

注意事项：

- status=UNKNOWN 时必须查询确认；写操作，SDK 不自动重试。

<a id="api-refund"></a>

### Refund — 资金退回接口

- 官方接口：[`alipay.fund.trans.refund`](https://opendocs.alipay.com/open-v3/0c6ff10a_alipay.fund.trans.refund)
- 路由：`POST /v3/alipay/fund/trans/refund`
- 请求类型：[RefundRequest](trans_request.go#L185)
- 响应类型：[RefundResponse](trans_response.go#L137)

当资金支付一段时间后，由于收款方一直没有收钱（比如抢红包），可以通过退款接口将资金退换给付款方。支付宝将会按照退款规则将支付款按原路退到付款方帐号上。
交易超过约定时间（签约时设置的可退款时间）的订单无法进行退款。支付宝退款支持单笔交易分多次退回，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。一笔退款失败后重新提交，要采用原来的退款单号。总退款金额不能超过（用户实际支付金额-已领红包）的金额。

请求约束：RefundRequest 是 alipay.fund.trans.refund（资金退回）的请求。
order_id、biz_scene、out_request_no、refund_amount 必填；退回失败重新提交
必须沿用原 out_request_no。写操作：SDK 不自动重试。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderID` | `string` | JSON body | 未标注 | 发红包时支付宝返回的支付宝订单号order_id。<br>SDK 字段说明：原转账的支付宝订单号。 |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 标识一次资金退回请求，一笔资金退回失败后重新提交，要采用原来的资金退回单号。总退款金额不能超过用户实际支付金额。<br>SDK 字段说明：资金退回请求号；重新提交必须沿用原单号。 |
| `refund_amount` | `RefundAmount` | `string` | JSON body | 未标注 | 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数<br>SDK 字段说明：退款金额，不能大于订单金额，单位为元。 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 资金退回备注 |
| `biz_scene` | `BizScene` | `string` | JSON body | 未标注 | 场景码。 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码 |
| `passback_params` | `PassbackParams` | `*string` | JSON body | 未标注 | 公用回传参数。 |
| `business_params` | `BusinessParams` | `*string` | JSON body | 未标注 | JSON格式，传递业务扩展参数，使用前请与支付宝工程师联系！。 |

注意事项：

- 退回失败重新提交必须沿用原 OutRequestNo，防止重复退回。

## 嵌套请求参数

<a id="model-authinfo"></a>

### AuthInfo

类型定义：[AuthInfo](trans_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_code` | `AuthCode` | `*string` | 未标注 | 授权号，协议授权或者cif协议对应的授权号。<br>SDK 字段说明：授权号，协议授权或 CIF 协议对应的授权号。 |
| `auth_type` | `AuthType` | `*string` | 未标注 | 授权类型 |

<a id="model-bankcardextinfo"></a>

### BankcardExtInfo

类型定义：[BankcardExtInfo](trans_model.go#L65)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `bank_code` | `BankCode` | `*string` | 未标注 | 银行支行联行号 |
| `inst_name` | `InstName` | `*string` | 未标注 | 机构名称 |
| `inst_branch_name` | `InstBranchName` | `*string` | 未标注 | 收款银行所属支行 |
| `inst_province` | `InstProvince` | `*string` | 未标注 | 银行所在省<br>SDK 字段说明：收款银行所在省。 |
| `inst_city` | `InstCity` | `*string` | 未标注 | 收款银行所在市 |
| `account_type` | `AccountType` | `*string` | 未标注 | 收款账户类型。 1：对公（在金融机构开设的公司账户） 2：对私（在金融机构开设的个人账户）<br>SDK 字段说明：收款账户类型：1 对公，2 对私。 |

<a id="model-participant"></a>

### Participant

类型定义：[Participant](trans_model.go#L42)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `identity` | `Identity` | `*string` | 未标注 | 参与方的唯一标识 |
| `identity_type` | `IdentityType` | `*string` | 未标注 | 参与方的标识类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、ALIPAY_OPEN_ID：支付宝openid 4、EXPRESS_DC_STFA：对公快捷银行卡 5、BANKCARD_ACCOUNT:银行卡账号<br>SDK 字段说明：参与方的标识类型：ALIPAY_USER_ID（会员 ID）、<br>ALIPAY_LOGON_ID（登录号）、ALIPAY_OPEN_ID（openid）、<br>EXPRESS_DC_STFA（对公快捷）等。 |
| `name` | `Name` | `*string` | 未标注 | 参与方真实姓名，如果非空，将校验收款支付宝账号姓名一致性。当identity_type=ALIPAY_LOGON_ID时，本字段必填。<br>SDK 字段说明：参与方真实姓名；非空时校验收款账号姓名一致性；<br>identity_type=ALIPAY_LOGON_ID 时必填。 |
| `cert_type` | `CertType` | `*string` | 未标注 | 参与方的证件类型。 |
| `cert_no` | `CertNo` | `*string` | 未标注 | 参与方的证件号，支持身份证号、护照号。 |
| `bankcard_ext_info` | `BankcardExtInfo` | [`*BankcardExtInfo`](#model-bankcardextinfo) | 未标注 | 银行卡扩展信息（转账到银行卡场景）。 |
| `ext_info` | `ExtInfo` | `*string` | 未标注 | 描述参与方信息的扩展属性，使用前请与支付宝工程师确认。 |
| `merchant_user_info` | `MerchantUserInfo` | `*string` | 未标注 | 商户侧的参与方信息。 |

<a id="model-transferscenereportinfo"></a>

### TransferSceneReportInfo

类型定义：[TransferSceneReportInfo](trans_model.go#L34)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `info_type` | `InfoType` | `*string` | 未标注 | 转账场景信息类型，即商户所属转账场景下的信息类型。此字段内容为固定值，需严格按照[转账场景信息字段说明](https://opendocs.alipay.com/open/0iaxid)传参。<br> 例如：transfer_scene_name为“佣金报酬”时，该值必须填写取值：佣金报酬说明<br> 注：信息类型不可重复填写，请严格按照示例传入<br>SDK 字段说明：转账场景信息类型，取值见官方转账场景信息字段说明。 |
| `info_content` | `InfoContent` | `*string` | 未标注 | 转账场景信息描述，即商户所属转账场景下的信息内容。商户可按实际业务场景自定义传参，需严格按照[转账场景信息字段说明](https://opendocs.alipay.com/open/0iaxid)传参。<br> 例如：transfer_scene_name为“佣金报酬”且info_type为“佣金报酬说明”时，该值可传入：8月家政服务报酬<br>SDK 字段说明：转账场景信息描述，按官方字段说明传参。 |
