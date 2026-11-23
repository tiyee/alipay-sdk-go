# pcredit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/pcredit`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../README.md)创建 `*core.Client`，再通过 `pcredit.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`HuabeiAuthAccumulationQuery`](#api-huabeiauthaccumulationquery) | [`alipay.pcredit.huabei.auth.accumulation.query`](https://opendocs.alipay.com/pre-apis/05ge5u) | 花芝轻会员周期累计数据查询 | `GET` |
| [`HuabeiAuthAgreementClose`](#api-huabeiauthagreementclose) | [`alipay.pcredit.huabei.auth.agreement.close`](https://opendocs.alipay.com/pre-apis/05gb9o) | 花芝协议关闭接口 | `PUT` |
| [`HuabeiAuthAgreementQuery`](#api-huabeiauthagreementquery) | [`alipay.pcredit.huabei.auth.agreement.query`](https://opendocs.alipay.com/pre-apis/05fign) | 花芝轻会员协议查询接口 | `GET` |
| [`HuabeiAuthOrderQuery`](#api-huabeiauthorderquery) | [`alipay.pcredit.huabei.auth.order.query`](https://opendocs.alipay.com/pre-apis/05h84d) | 花呗先享订单查询接口 | `GET` |
| [`HuabeiAuthOrderUnfreeze`](#api-huabeiauthorderunfreeze) | [`alipay.pcredit.huabei.auth.order.unfreeze`](https://opendocs.alipay.com/pre-apis/05gcoh) | 花呗先享解冻或解约接口 | `PUT` |
| [`HuabeiAuthSettleApply`](#api-huabeiauthsettleapply) | [`alipay.pcredit.huabei.auth.settle.apply`](https://opendocs.alipay.com/apis/05ghz8) | 花芝轻会员结算申请 | `PUT` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-huabeiauthaccumulationquery"></a>

### HuabeiAuthAccumulationQuery — 花芝轻会员周期累计数据查询

- 官方接口：[`alipay.pcredit.huabei.auth.accumulation.query`](https://opendocs.alipay.com/pre-apis/05ge5u)
- 路由：`GET /v3/alipay/pcredit/huabei/auth/accumulation/query`
- 请求类型：[PcreditHuabeiAuthAccumulationQueryRequest](pcredit_request.go#L62)
- 响应类型：[PcreditHuabeiAuthAccumulationQueryResponse](pcredit_response.go#L5)

根据协议与对应的周期查询支付宝累计数据。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 支付宝系统中用以唯一标识用户签约记录的编号。 |
| `period` | `Period` | `*string` | query | 否 | 协议结算周期，使用支付宝端回传的周期信息。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 买家在支付宝的用户id |
| `open_id` | `OpenId` | `*string` | query | 否 | 买家在支付宝的用户id |

<a id="api-huabeiauthagreementclose"></a>

### HuabeiAuthAgreementClose — 花芝协议关闭接口

- 官方接口：[`alipay.pcredit.huabei.auth.agreement.close`](https://opendocs.alipay.com/pre-apis/05gb9o)
- 路由：`PUT /v3/alipay/pcredit/huabei/auth/agreement/close`
- 请求类型：[PcreditHuabeiAuthAgreementCloseRequest](pcredit_request.go#L5)
- 响应类型：[PcreditHuabeiAuthAgreementCloseResponse](pcredit_response.go#L17)

用户已经开通花芝先享后，商户可通过此接口提前关闭协议。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 花芝协议号，支付宝系统中用以唯一标识用户签约记录的编号，即花芝协议号。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 外部请求号 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `quit_type` | `QuitType` | `*string` | JSON body | 未标注 | 用户主动意愿退出：USER_CANCEL_QUIT; 商户结算退出：SETTLE_APPLY_QUIT; 默认值为SETTLE_APPLY_QUIT；这个字段会影响用户在芝麻信用合约的状态 |

<a id="api-huabeiauthagreementquery"></a>

### HuabeiAuthAgreementQuery — 花芝轻会员协议查询接口

- 官方接口：[`alipay.pcredit.huabei.auth.agreement.query`](https://opendocs.alipay.com/pre-apis/05fign)
- 路由：`GET /v3/alipay/pcredit/huabei/auth/agreement/query`
- 请求类型：[PcreditHuabeiAuthAgreementQueryRequest](pcredit_request.go#L74)
- 响应类型：[PcreditHuabeiAuthAgreementQueryResponse](pcredit_response.go#L23)

花芝轻会员协议查询，返回协议的状态。如果协议处于签约状态，同时返回最新的资金池余额。
有两种方式查询。优先使用agreement_no查询；其次是使用(out_sign_no, auth_scene)查询。
如果商户不能保证out_sign_no的唯一性，慎用(out_sign_no, auth_scene)的查询方式，它优先返回仍处于签约状态的协议。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员。传入该参数，会忽略其它所有参数。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 买家在支付宝的用户id |
| `open_id` | `OpenId` | `*string` | query | 否 | 买家在支付宝的用户id |
| `out_sign_no` | `OutSignNo` | `*string` | query | 否 | 外部签约号，由商户提供，花芝轻会员协议中标识用户的唯一签约号（确保在商户系统中唯一）。 |
| `auth_scene` | `AuthScene` | `*string` | query | 否 | 花芝轻会员签约场景，商户和支付宝签约时确定。该值需要与系统/页面签约接口调用时传入的值保持一致。 |

<a id="api-huabeiauthorderquery"></a>

### HuabeiAuthOrderQuery — 花呗先享订单查询接口

- 官方接口：[`alipay.pcredit.huabei.auth.order.query`](https://opendocs.alipay.com/pre-apis/05h84d)
- 路由：`GET /v3/alipay/pcredit/huabei/auth/order/query`
- 请求类型：[PcreditHuabeiAuthOrderQueryRequest](pcredit_request.go#L88)
- 响应类型：[PcreditHuabeiAuthOrderQueryResponse](pcredit_response.go#L51)

查询花呗先享冻结、解冻订单内容及状态。有3种查询方式。推荐商户优先使用auth_opt_id查询；其次是按照(alipay_user_id,out_request_no)组合方式查询；最后是单独通过out_request_no方式查询。
注意：最后一种方式，仅支持2019年2月15日开始的订单。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_opt_id` | `AuthOptId` | `*string` | query | 否 | 支付宝侧花呗冻结、解冻操作单据id。在原先的冻结或者解冻接口调用中同步返回给商户，或者通过商户通知返回给商户。按订单号查询时，此字段不可为空。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 买家在支付宝的用户id。通过userid+请求流水号组合查询时，此字段不可为空。 |
| `open_id` | `OpenId` | `*string` | query | 否 | 买家在支付宝的用户id |
| `out_request_no` | `OutRequestNo` | `*string` | query | 否 | 商户原先调用冻结、解冻接口传入的请求流水号。按照流水号查询订单时，此字段不能为空。 |

<a id="api-huabeiauthorderunfreeze"></a>

### HuabeiAuthOrderUnfreeze — 花呗先享解冻或解约接口

- 官方接口：[`alipay.pcredit.huabei.auth.order.unfreeze`](https://opendocs.alipay.com/pre-apis/05gcoh)
- 路由：`PUT /v3/alipay/pcredit/huabei/auth/order/unfreeze`
- 请求类型：[PcreditHuabeiAuthOrderUnfreezeRequest](pcredit_request.go#L21)
- 响应类型：[PcreditHuabeiAuthOrderUnfreezeResponse](pcredit_response.go#L79)

用户已经开通花呗先享协议后，商户通过此接口解冻用户资金池金额，也可以解冻并解约。
如果是解约操作，则要求传入的解冻金额必须等于用户资金池余额。
注意：商户在发起解约前，请务必保证已经结算过用户会员费，一旦解约后，无法发起结算用户会员费操作。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `need_terminated` | `NeedTerminated` | `*string` | JSON body | 未标注 | true表示需要解约。false或者不填写表示仅解冻。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 商户业务订单的简单描述，如商品名称等，长度不超过100个字母或50个汉字 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。 |
| `seller_id` | `SellerId` | `*string` | JSON body | 未标注 | 商户的支付宝用户id。如果该值为空，则默认为商户签约账号对应的支付宝用户ID。 |
| `unfreeze_amount` | `UnfreezeAmount` | `*string` | JSON body | 未标注 | 需要解冻的金额，单位为：元（人民币），精确到小数点后两位 |

<a id="api-huabeiauthsettleapply"></a>

### HuabeiAuthSettleApply — 花芝轻会员结算申请

- 官方接口：[`alipay.pcredit.huabei.auth.settle.apply`](https://opendocs.alipay.com/apis/05ghz8)
- 路由：`PUT /v3/alipay/pcredit/huabei/auth/settle/apply`
- 请求类型：[PcreditHuabeiAuthSettleApplyRequest](pcredit_request.go#L41)
- 响应类型：[PcreditHuabeiAuthSettleApplyResponse](pcredit_response.go#L105)

用户已经开通花芝轻会员协议后，商户通过此接口解冻转支付用户冻结金额。传入金额必须小于等于冻结金额

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `action_type` | `ActionType` | `*string` | JSON body | 未标注 | 操作类型，默认为 MERCHANT_SETTLE（结算）。枚举值如下： * MERCHANT_SETTLE：结算（需要主动退出在发起结算）; * QUIT_SETTLE：退出协议（解约协议）; * PERIOD_SETTLE：分阶段结算（不解约协议）。 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `extend_params` | `ExtendParams` | [`*ExtraParams`](#model-extraparams) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `need_terminated` | `NeedTerminated` | `*string` | JSON body | 未标注 | 是否解约，该参数控制结算完成之后的操作，为空则表示结算后解约。枚举值如下： * true：代表解约，默认为true。 * false：代表不解约。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户在支付宝的唯一标识 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。 |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 需要支付的金额，单位为：元（人民币），精确到小数点后两位 |
| `seller_id` | `SellerId` | `*string` | JSON body | 未标注 | 支付宝商家ID，即商家账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 说明：若该值为空，则默认为商户签约账号对应的支付宝用户ID。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-extraparams"></a>

### ExtraParams

类型定义：[ExtraParams](pcredit_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `period` | `Period` | `*string` | 未标注 | 轻会员场景下协议结算周期，由协议生效日期到失效日期拼接(yyyyMMddHHmmss-yyyyMMddHHmmss)，精确到秒。 |
| `period_summary_info` | `PeriodSummaryInfo` | `*string` | 未标注 | 轻会员费用结算原始信息。详细字段说明：total_real_pay_amount轻会员周期内累计支付宝支付金额， total_pay_count轻会员周期内累计支付宝支付次数，total_discount_amount轻会员周期内累计享受的轻会员优惠。（上面金额字段单位为元，精确小数点后两位） |
| `quit_type` | `QuitType` | `*string` | 未标注 | 用户主动意愿退出：USER_CANCEL_QUIT; 商户结算退出：SETTLE_APPLY_QUIT; 默认值为SETTLE_APPLY_QUIT；这个字段会影响用户在芝麻信用合约的状态 |
| `withhold_index` | `WithholdIndex` | `*string` | 未标注 | 代扣期数，周期扣场景PERIOD_SETTLE下需要传递 |
