# zhima/credit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/zhima/credit`。本目录提供 13 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `credit.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_credit_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PayAfterUseCreditAgreementQuery`](#api-payafterusecreditagreementquery) | [`zhima.credit.payafteruse.creditagreement.query`](https://opendocs.alipay.com/open-v3/eaff9140_zhima.credit.payafteruse.creditagreement.query) | 查询服务开通/授权信息 | `GET` |
| [`PayAfterUseCreditAgreementTransfer`](#api-payafterusecreditagreementtransfer) | [`zhima.credit.payafteruse.creditagreement.transfer`](https://opendocs.alipay.com/apis/05vdam) | 芝麻信用服务授权 | `POST` |
| [`PayAfterUseCreditBizOrderFinish`](#api-payafterusecreditbizorderfinish) | [`zhima.credit.payafteruse.creditbizorder.finish`](https://opendocs.alipay.com/open-v3/7972514c_zhima.credit.payafteruse.creditbizorder.finish?scene=common&pathHash=ba647372) | 结束信用服务订单 | `POST` |
| [`PayAfterUseCreditBizOrderOrder`](#api-payafterusecreditbizorderorder) | [`zhima.credit.payafteruse.creditbizorder.order`](https://opendocs.alipay.com/open-v3/51ba4ec1_zhima.credit.payafteruse.creditbizorder.order) | 芝麻信用服务下单（免用户确认场景） | `POST` |
| [`PayAfterUseCreditBizOrderQuery`](#api-payafterusecreditbizorderquery) | [`zhima.credit.payafteruse.creditbizorder.query`](https://opendocs.alipay.com/open-v3/0f11f105_zhima.credit.payafteruse.creditbizorder.query?scene=common&pathHash=203d9491) | 信用服务订单查询 | `GET` |
| [`PeZmGoAgreementQuery`](#api-pezmgoagreementquery) | [`zhima.credit.pe.zmgo.agreement.query`](https://opendocs.alipay.com/open-v3/849a96d9_zhima.credit.pe.zmgo.agreement.query) | 芝麻Go协议查询接口 | `GET` |
| [`PeZmGoAgreementUnsign`](#api-pezmgoagreementunsign) | [`zhima.credit.pe.zmgo.agreement.unsign`](https://opendocs.alipay.com/open-v3/4d77e884_zhima.credit.pe.zmgo.agreement.unsign) | 芝麻GO协议解约 | `POST` |
| [`PeZmGoBizOptClose`](#api-pezmgobizoptclose) | [`zhima.credit.pe.zmgo.bizopt.close`](https://opendocs.alipay.com/apis/05vluc) | 芝麻GO签约关单 | `POST` |
| [`PeZmGoCumulationSync`](#api-pezmgocumulationsync) | [`zhima.credit.pe.zmgo.cumulation.sync`](https://opendocs.alipay.com/apis/05v7c8) | 芝麻go用户数据回传 | `POST` |
| [`PeZmGoPreOrderCreate`](#api-pezmgopreordercreate) | [`zhima.credit.pe.zmgo.preorder.create`](https://opendocs.alipay.com/open-v3/3a385c23_zhima.credit.pe.zmgo.preorder.create) | 芝麻GO签约预创单 | `PUT` |
| [`PeZmGoSettleApply`](#api-pezmgosettleapply) | [`zhima.credit.pe.zmgo.settle.apply`](https://opendocs.alipay.com/open-v3/3a0505ce_zhima.credit.pe.zmgo.settle.apply) | 芝麻GO结算申请 | `POST` |
| [`PeZmGoSettleRefund`](#api-pezmgosettlerefund) | [`zhima.credit.pe.zmgo.settle.refund`](https://opendocs.alipay.com/open-v3/a66ab3b0_zhima.credit.pe.zmgo.settle.refund) | 芝麻GO结算退款接口 | `POST` |
| [`PeZmGoSettleUnfreeze`](#api-pezmgosettleunfreeze) | [`zhima.credit.pe.zmgo.settle.unfreeze`](https://opendocs.alipay.com/apis/05gzu9) | 芝麻GO解冻接口 | `PUT` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-payafterusecreditagreementquery"></a>

### PayAfterUseCreditAgreementQuery — 查询服务开通/授权信息

- 官方接口：[`zhima.credit.payafteruse.creditagreement.query`](https://opendocs.alipay.com/open-v3/eaff9140_zhima.credit.payafteruse.creditagreement.query)
- 路由：`GET /v3/zhima/credit/payafteruse/creditagreement/query`
- 请求类型：[PayAfterUseCreditAgreementQueryRequest](credit_request.go#L235)
- 响应类型：[PayAfterUseCreditAgreementQueryResponse](credit_response.go#L5)

使用商户外部协议号或开通/授权协议号，查询服务开通/授权信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_agreement_no` | `OutAgreementNo` | `*string` | query | 否 | 商户外部协议号 |
| `credit_agreement_id` | `CreditAgreementId` | `*string` | query | 否 | 芝麻开通/授权协议号，out_agreement_no与credit_agreement_id必填一个，推荐使用credit_agreement_id。 |
| `product_code` | `ProductCode` | `*string` | query | 否 | 当产品是芝麻守约时必填，需传入CREDIT_AGREEMENT |
| `service_id` | `ServiceId` | `*string` | query | 否 | 芝麻信用服务id，上线前提前分配给商户 |
| `extra_param` | `ExtraParam` | `*string` | query | 否 | 扩展参数 |

<a id="api-payafterusecreditagreementtransfer"></a>

### PayAfterUseCreditAgreementTransfer — 芝麻信用服务授权

- 官方接口：[`zhima.credit.payafteruse.creditagreement.transfer`](https://opendocs.alipay.com/apis/05vdam)
- 路由：`POST /v3/zhima/credit/payafteruse/creditagreement/transfer`
- 请求类型：[PayAfterUseCreditAgreementTransferRequest](credit_request.go#L5)
- 响应类型：[PayAfterUseCreditAgreementTransferResponse](credit_response.go#L23)

其他产品功能（例如信用代扣），通过该接口，可以转换为芝麻产品（芝麻先享产品，芝麻风险评估产品），完成信用服务授权

请求约束：PayAfterUseCreditAgreementTransferRequest is the official V3 ZhimaCreditPayafteruseCreditagreementTransferModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 芝麻外部类目 |
| `credit_agreement_id` | `CreditAgreementId` | `*string` | JSON body | 未标注 | 芝麻开通/授权协议号，在用户开通芝麻先享服务后，由芝麻生成并返回给商户；只有当opertaion_type传入AGREEMENT_TRANSFER值是，该字段必传；其他场景不需要传入 |
| `deduct_agreement_no` | `DeductAgreementNo` | `*string` | JSON body | 未标注 | 用户的代扣协议号 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 业务扩展参数，Json格式 |
| `operation_type` | `OperationType` | `*string` | JSON body | 未标注 | AGREEMENT_TRANSFER代表是存量先享协议的主体切换功能；其他场景不需要传入；默认为空； |
| `out_agreement_no` | `OutAgreementNo` | `*string` | JSON body | 未标注 | 商户外部协议号，不同的支付宝用户需要传递不同的外部单号 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 商户签约的芝麻产品的产品码 |
| `zm_service_id` | `ZmServiceId` | `*string` | JSON body | 未标注 | 芝麻信用服务ID |

<a id="api-payafterusecreditbizorderfinish"></a>

### PayAfterUseCreditBizOrderFinish — 结束信用服务订单

- 官方接口：[`zhima.credit.payafteruse.creditbizorder.finish`](https://opendocs.alipay.com/open-v3/7972514c_zhima.credit.payafteruse.creditbizorder.finish?scene=common&pathHash=ba647372)
- 路由：`POST /v3/zhima/credit/payafteruse/creditbizorder/finish`
- 请求类型：[PayAfterUseCreditBizOrderFinishRequest](credit_request.go#L25)
- 响应类型：[PayAfterUseCreditBizOrderFinishResponse](credit_response.go#L33)

结束信用服务订单

请求约束：PayAfterUseCreditBizOrderFinishRequest is the official V3 ZhimaCreditPayafteruseCreditbizorderFinishModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `credit_biz_order_id` | `CreditBizOrderId` | `*string` | JSON body | 未标注 | 信用服务订单号 |
| `is_fulfilled` | `IsFulfilled` | `*string` | JSON body | 未标注 | 字符串类型，用户此订单是否守约。 传"true"时，用户在芝麻信用-守约记录中，该笔订单是已守约状态； 传"false"时，用户在芝麻信用-守约记录中，该笔订单是已取消状态。 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户外部请求号 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 商户对本次操作的附言描述 |

<a id="api-payafterusecreditbizorderorder"></a>

### PayAfterUseCreditBizOrderOrder — 芝麻信用服务下单（免用户确认场景）

- 官方接口：[`zhima.credit.payafteruse.creditbizorder.order`](https://opendocs.alipay.com/open-v3/51ba4ec1_zhima.credit.payafteruse.creditbizorder.order)
- 路由：`POST /v3/zhima/credit/payafteruse/creditbizorder/order`
- 请求类型：[PayAfterUseCreditBizOrderOrderRequest](credit_request.go#L37)
- 响应类型：[PayAfterUseCreditBizOrderOrderResponse](credit_response.go#L43)

芝麻信用产品免密下单，不需要唤起支付宝APP，通过服务端调用完成下单。
涉及芝麻信用服务产品、芝麻风险评估产品

请求约束：PayAfterUseCreditBizOrderOrderRequest is the official V3 ZhimaCreditPayafteruseCreditbizorderOrderModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount_type` | `AmountType` | `*string` | JSON body | 未标注 | 只有当传递了order_amount时，该参数才有意义； 1）该参数不传时，默认为ORDER_AMOUNT。 2）传ORDER_AMOUNT时，表示order_amount传入的金额为后付金额，在发起扣款时，最大扣款支付金额为order_amount传入的值（取值单位为元）； 3）传RISK_AMOUNT时，表示order_amount传入的金额为风险预估金额，在发起扣款时，最大扣款支付金额为商户签约时约定的上限额度（取值单位为元）。 |
| `body` | `Body` | `*string` | JSON body | 未标注 | 订单描述 |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 芝麻外部类目，芝麻先享接入无差异化风控诉求可不传 |
| `commercial_sub_mode` | `CommercialSubMode` | `*string` | JSON body | 未标注 | 业务子模式。默认的单次付模式无需传入，阶段付模式传入以区分是分次还是分期子模式。 |
| `credit_agreement_id` | `CreditAgreementId` | `*string` | JSON body | 未标注 | 芝麻开通协议号 |
| `credit_commercial_mode` | `CreditCommercialMode` | `*string` | JSON body | 未标注 | 信用业务模式，不填默认为单次扣款模式。阶段付模式为STAGE_PAYMENT，其它模式请根据对应的技术支持文档传入 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 业务扩展参数 |
| `order_amount` | `OrderAmount` | `*string` | JSON body | 未标注 | 订单金额，该金额为当前订单扣款的累计最大额度。例如，下单时传递10.00，则扣款时最大支付金额为10元。若该参数不传，则默认使用商户签约时约定的上限额度。芝麻速办业务场景（极速回收、极速返利、极速退款等）商户请求时，order_amount必传，且amount_type类型需传递ORDER_AMOUNT。 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 商户外部订单号，保证不重复 |
| `payment_total_times` | `PaymentTotalTimes` | `*string` | JSON body | 未标注 | 多阶段订单次数，业务模式为阶段付模式下时需传入 |
| `product_code` | `ProductCode` | `*string` | JSON body | 未标注 | 产品码，不传默认为CREDIT_PAY_AFTER_USE |
| `stage_period_type` | `StagePeriodType` | `*string` | JSON body | 未标注 | 阶段付分期类型。阶段付模式，且子业务模式为分期模式下需要传入，分次不需要 |
| `subject` | `Subject` | `*string` | JSON body | 未标注 | 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。 |

<a id="api-payafterusecreditbizorderquery"></a>

### PayAfterUseCreditBizOrderQuery — 信用服务订单查询

- 官方接口：[`zhima.credit.payafteruse.creditbizorder.query`](https://opendocs.alipay.com/open-v3/0f11f105_zhima.credit.payafteruse.creditbizorder.query?scene=common&pathHash=203d9491)
- 路由：`GET /v3/zhima/credit/payafteruse/creditbizorder/query`
- 请求类型：[PayAfterUseCreditBizOrderQueryRequest](credit_request.go#L249)
- 响应类型：[PayAfterUseCreditBizOrderQueryResponse](credit_response.go#L51)

信用服务订单查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `credit_biz_order_id` | `CreditBizOrderId` | `*string` | query | 否 | 信用服务订单号，out_order_no与credit_biz_order_id至少传一个 |
| `out_order_no` | `OutOrderNo` | `*string` | query | 否 | 商户外部单号，out_order_no与credit_biz_order_id至少传一个 |

<a id="api-pezmgoagreementquery"></a>

### PeZmGoAgreementQuery — 芝麻Go协议查询接口

- 官方接口：[`zhima.credit.pe.zmgo.agreement.query`](https://opendocs.alipay.com/open-v3/849a96d9_zhima.credit.pe.zmgo.agreement.query)
- 路由：`GET /v3/zhima/credit/pe/zmgo/agreement/query`
- 请求类型：[PeZmGoAgreementQueryRequest](credit_request.go#L257)
- 响应类型：[PeZmGoAgreementQueryResponse](credit_response.go#L73)

用户已经开通芝麻GO后，通过此接口查询协议。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | query | 否 | 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员。传入该参数，会忽略其它所有参数。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | query | 否 | 买家在支付宝的用户id |
| `open_id` | `OpenId` | `*string` | query | 否 | 买家在支付宝的用户id |

<a id="api-pezmgoagreementunsign"></a>

### PeZmGoAgreementUnsign — 芝麻GO协议解约

- 官方接口：[`zhima.credit.pe.zmgo.agreement.unsign`](https://opendocs.alipay.com/open-v3/4d77e884_zhima.credit.pe.zmgo.agreement.unsign)
- 路由：`POST /v3/zhima/credit/pe/zmgo/agreement/unsign`
- 请求类型：[PeZmGoAgreementUnsignRequest](credit_request.go#L67)
- 响应类型：[PeZmGoAgreementUnsignResponse](credit_response.go#L93)

芝麻GO协议解约

请求约束：PeZmGoAgreementUnsignRequest is the official V3 ZhimaCreditPeZmgoAgreementUnsignModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号。 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 支付宝的用户id |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝的用户id |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `quit_type` | `QuitType` | `*string` | JSON body | 未标注 | quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT |

<a id="api-pezmgobizoptclose"></a>

### PeZmGoBizOptClose — 芝麻GO签约关单

- 官方接口：[`zhima.credit.pe.zmgo.bizopt.close`](https://opendocs.alipay.com/apis/05vluc)
- 路由：`POST /v3/zhima/credit/pe/zmgo/bizopt/close`
- 请求类型：[PeZmGoBizOptCloseRequest](credit_request.go#L81)
- 响应类型：[PeZmGoBizOptCloseResponse](credit_response.go#L101)

签约芝麻GO时，在签约申请时会生初始状态的签约单，如果后续不进行签约确认，可通过此接口关闭签约单

请求约束：PeZmGoBizOptCloseRequest is the official V3 ZhimaCreditPeZmgoBizoptCloseModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 支付宝用户Id |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户Id |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次操作的请求流水号，由商户生成，请确保唯一性 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户Id |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 芝麻GO配置模板Id |

<a id="api-pezmgocumulationsync"></a>

### PeZmGoCumulationSync — 芝麻go用户数据回传

- 官方接口：[`zhima.credit.pe.zmgo.cumulation.sync`](https://opendocs.alipay.com/apis/05v7c8)
- 路由：`POST /v3/zhima/credit/pe/zmgo/cumulation/sync`
- 请求类型：[PeZmGoCumulationSyncRequest](credit_request.go#L95)
- 响应类型：[PeZmGoCumulationSyncResponse](credit_response.go#L115)

芝麻go商户通过此api接口，将用户行为数据回传到芝麻侧，数据包括：交易型数据（订单信息、优惠信息）和非交易型数据（用户完成任务次数）。芝麻侧将回传的数据在C端向用户展示，并且在结算时对结算金额进行校验（可选）。

请求约束：PeZmGoCumulationSyncRequest is the official V3 ZhimaCreditPeZmgoCumulationSyncModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 芝麻go协议号，唯一标识一个芝麻go协议。为必传入参。 |
| `amount_type_data` | `AmountTypeData` | [`*AmountTypeData`](#model-amounttypedata) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_action` | `BizAction` | `*string` | JSON body | 未标注 | 数据回传的动作类型，为枚举值，为必传入参。 可取值： ("ORDER_PAID", "正向支付")，("ORDER_REFUND", "逆向退款")。 其中ORDER_PAID表示回传为正向的，即订单金额累加，优惠累加，任务次数累加；而ORDER_REFUND表示回传为逆向的，即订单金额退款，优惠退款，任务次数回退。 |
| `biz_time` | `BizTime` | `*string` | JSON body | 未标注 | 交易发生时间，包括：订单交易时间，优惠核销时间，以及任务完成时间。为必传入参。 如：用户在【2019-03-08 00:00:00】核销了一个红包，但是【2019-03-10 00:00:00】才进行数据回传，此时biz_time取值红包核销的时间，也就是【2019-03-08 00:00:00】。 |
| `cumulate_data_type` | `CumulateDataType` | `*string` | JSON body | 未标注 | 累计类型。为枚举值，为必传入参。 可取值：("AMOUNT", "金额类型")，("TASK", "任务类型")，("AMOUNT_TASK", "额度和任务类型")。 其中，交易解耦模式仅允许AMOUNT或TASK。而交易绑定模式下，仅允许AMOUNT_TASK。 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 扩展字段，格式为Json格式的字符串，非必传。 |
| `has_alipay_trade` | `HasAlipayTrade` | `*bool` | JSON body | 未标注 | 是否为支付宝交易。仅在累计模式为【交易绑定模式】下被消费。当芝麻Go模板配置的累计模式为交易绑定模式时，此时需传入扩展参数是否为支付宝交易。若是支付宝交易，传值true，否则为false。若不传此字段，默认为false。 若为true，则会对交易号trade_no进行校验，校验是否为支付宝交易，以及用户userId和商户partnerId是否一致，若校验不通过，累计失败。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户userId，蚂蚁统一会员ID，为必传入参。 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号。需保证唯一，为必传入参。 若接口调用失败，提示信息为更换幂等号，商户更换此字段重新发起调用。 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID，为必传入参。 |
| `pay_out_biz_no` | `PayOutBizNo` | `*string` | JSON body | 未标注 | 逆向对应的正向外部业务号。逆向场景为必传入参。也就是biz_action为ORDER_REFUND时为必传入参。 逆向时需通过此字段明确对应的正向的外部业务号。 |
| `request_from` | `RequestFrom` | `*string` | JSON body | 未标注 | 取固定值ExternalMerchantSource，为必传入参。表示数据回传来源为外部商户，数据回传内部逻辑需感知。 |
| `task_type_data` | `TaskTypeData` | [`*TaskTypeData`](#model-tasktypedata) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户userId，蚂蚁统一会员ID，为必传入参。 |

<a id="api-pezmgopreordercreate"></a>

### PeZmGoPreOrderCreate — 芝麻GO签约预创单

- 官方接口：[`zhima.credit.pe.zmgo.preorder.create`](https://opendocs.alipay.com/open-v3/3a385c23_zhima.credit.pe.zmgo.preorder.create)
- 路由：`PUT /v3/zhima/credit/pe/zmgo/preorder/create`
- 请求类型：[PeZmGoPreOrderCreateRequest](credit_request.go#L127)
- 响应类型：[PeZmGoPreOrderCreateResponse](credit_response.go#L129)

芝麻GO签约三件套前预创建单据，预传入签约基本参数和自定义配置，并支持动态模板配置能力。

请求约束：PeZmGoPreOrderCreateRequest is the official V3 ZhimaCreditPeZmgoPreorderCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `alipay_open_id` | `AlipayOpenId` | `*string` | JSON body | 未标注 | 支付宝用户ID。同一个外部请求号，调用方要确保支付宝userId取值不变 |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 支付宝用户ID。同一个外部请求号，调用方要确保支付宝userId取值不变 |
| `biz_time` | `BizTime` | `*string` | JSON body | 未标注 | 业务发生时间，如2016-07-06 00:00:02 签约超时关单时间依赖这个时间 这个时间会做幂等校验，如果创单幂等，但时间不同会失败 |
| `custom_temp_conf` | `CustomTempConf` | `*string` | JSON body | 未标注 | 废弃字段，勿用 |
| `expire_aisle_data` | `ExpireAisleData` | `*string` | JSON body | 未标注 | 协议到期后通知 |
| `ext_template_conf` | `ExtTemplateConf` | [`*ExtTemplateConf`](#model-exttemplateconf) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `extend_params` | `ExtendParams` | [`*PreOrderExtInfo`](#model-preorderextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `freeze_amount` | `FreezeAmount` | `*string` | JSON body | 未标注 | 冻结金额，单位：元 |
| `isv_pid` | `IsvPid` | `*string` | JSON body | 未标注 | ISV商户ID |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 外部请求号，唯一 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `partner_user_identifier` | `PartnerUserIdentifier` | `*string` | JSON body | 未标注 | 商户侧用户唯一标识。ex.商户侧uid |
| `pay_aisle_data` | `PayAisleData` | `*string` | JSON body | 未标注 | 结算成功后通知 |
| `sign_aisle_data` | `SignAisleData` | `*string` | JSON body | 未标注 | 签约成功后通知 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 签约芝麻GO的模板ID |
| `timeout_express` | `TimeoutExpress` | `*string` | JSON body | 未标注 | 超时关单时间 默认1H。1m～15d。m表示分钟，h表示小时，d表示天。该参数数值不接受小数点， 如1.5h，可转换为90m 签约超时关单时间依赖这个时间 |

<a id="api-pezmgosettleapply"></a>

### PeZmGoSettleApply — 芝麻GO结算申请

- 官方接口：[`zhima.credit.pe.zmgo.settle.apply`](https://opendocs.alipay.com/open-v3/3a0505ce_zhima.credit.pe.zmgo.settle.apply)
- 路由：`POST /v3/zhima/credit/pe/zmgo/settle/apply`
- 请求类型：[PeZmGoSettleApplyRequest](credit_request.go#L163)
- 响应类型：[PeZmGoSettleApplyResponse](credit_response.go#L139)

用户已经开通芝麻GO后，商户通过此接口解冻转支付用户冻结金额。传入金额必须小于等于冻结金额

请求约束：PeZmGoSettleApplyRequest is the official V3 ZhimaCreditPeZmgoSettleApplyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `action_type` | `ActionType` | `*string` | JSON body | 未标注 | DEFAULT_SETTLE 默认结算模式; PAY_TO_ZERO 支付转0元； |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号。 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 需要支付的金额，单位为：元（人民币），精确到小数点后两位 |
| `settle_extend_params` | `SettleExtendParams` | [`*SettleExtraParams`](#model-settleextraparams) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `total_discount_amount` | `TotalDiscountAmount` | `*string` | JSON body | 未标注 | 芝麻GO任务开始到到期过程中和任务相关享受的总优惠金额，单位为：元（人民币），精确到小数点后两位。 |
| `total_real_pay_amount` | `TotalRealPayAmount` | `*string` | JSON body | 未标注 | 芝麻GO任务开始到到期过程中和任务相关支付的总金额，单位为：元（人民币），精确到小数点后两位。 |
| `total_task_count` | `TotalTaskCount` | `*string` | JSON body | 未标注 | 芝麻GO任务开始到到期过程中和任务相关完成的任务次数 |
| `withhold_plan_no` | `WithholdPlanNo` | `*string` | JSON body | 未标注 | 芝麻GO结算时，需要传入的扣款单据号。来源于协议到期通知消息，外部商户接入时，该参数必填 |

<a id="api-pezmgosettlerefund"></a>

### PeZmGoSettleRefund — 芝麻GO结算退款接口

- 官方接口：[`zhima.credit.pe.zmgo.settle.refund`](https://opendocs.alipay.com/open-v3/a66ab3b0_zhima.credit.pe.zmgo.settle.refund)
- 路由：`POST /v3/zhima/credit/pe/zmgo/settle/refund`
- 请求类型：[PeZmGoSettleRefundRequest](credit_request.go#L191)
- 响应类型：[PeZmGoSettleRefundResponse](credit_response.go#L153)

芝麻GO结算金额逆向退款接口: 对zhima.credit.pe.zmgo.settle.apply(芝麻GO结算申请接口) 产生的结算交易发起退款操作

请求约束：PeZmGoSettleRefundRequest is the official V3 ZhimaCreditPeZmgoSettleRefundModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 签约芝麻GO的业务协议号，在签约成功、结算成功等通知中回回传给商户。 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | JSON body | 未标注 | 支付宝用户ID |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 支付宝用户ID |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 本次退款说明 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 外部请求号 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `refund_amount` | `RefundAmount` | `*string` | JSON body | 未标注 | 退款金额，单位为：元（人民币），精确到小数点后两位 |
| `refund_type` | `RefundType` | `*string` | JSON body | 未标注 | 退款类型： MEMBER_FEE_REFUND 会员费结算退款 ，WITHHOLD_FEE_REFUND 周期扣款金额退款。 不填写默认为 MEMBER_FEE_REFUND |
| `withhold_plan_no` | `WithholdPlanNo` | `*string` | JSON body | 未标注 | 芝麻侧扣款计划单据号，在结算扣款成功的通知消息中会回传。 |

<a id="api-pezmgosettleunfreeze"></a>

### PeZmGoSettleUnfreeze — 芝麻GO解冻接口

- 官方接口：[`zhima.credit.pe.zmgo.settle.unfreeze`](https://opendocs.alipay.com/apis/05gzu9)
- 路由：`PUT /v3/zhima/credit/pe/zmgo/settle/unfreeze`
- 请求类型：[PeZmGoSettleUnfreezeRequest](credit_request.go#L213)
- 响应类型：[PeZmGoSettleUnfreezeResponse](credit_response.go#L169)

用户已经开通芝麻GO后，商户通过此接口解冻转支付用户冻结金额。传入金额必须小于等于冻结金额

请求约束：PeZmGoSettleUnfreezeRequest is the official V3 ZhimaCreditPeZmgoSettleUnfreezeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号 |
| `alipay_open_id` | `AlipayOpenId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `alipay_user_id` | `AlipayUserId` | `*string` | JSON body | 未标注 | 买家在支付宝的用户id |
| `biz_time` | `BizTime` | `*string` | JSON body | 未标注 | 解冻成功时间 |
| `order_title` | `OrderTitle` | `*string` | JSON body | 未标注 | 解冻的描述 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。由商户传入，最终返回给商户。 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | 商户ID |
| `unfreeze_amount` | `UnfreezeAmount` | `*string` | JSON body | 未标注 | 解冻金额 |
| `unfreeze_extend_params` | `UnfreezeExtendParams` | [`*UnfreezeExtendParams`](#model-unfreezeextendparams) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-amounttypedata"></a>

### AmountTypeData

类型定义：[AmountTypeData](credit_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `name` | `Name` | `*string` | 未标注 | 累计类型为金额类型的累计名称。为必传入参。 |
| `out_discount_infos` | `OutDiscountInfos` | [`*[]ZMGoOutDiscountInfo`](#model-zmgooutdiscountinfo) | 未标注 | 优惠信息。为List类型，支持传入多笔优惠，大部分场景传入一笔。包含：优惠名称，优惠类型，优惠金额（单位元）。 字段必传场景： 1.当累计模式为【交易绑定模式】时，为必传入参。 2.当累计模式为【交易解耦模式】，同时需要累计优惠信息时，为必传入参。 |
| `trade_info` | `TradeInfo` | [`*ZMGoTradeInfo`](#model-zmgotradeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-exttemplateconf"></a>

### ExtTemplateConf

类型定义：[ExtTemplateConf](credit_model.go#L21)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `buyer_id` | `BuyerId` | `*string` | 未标注 | 淘宝uid |
| `xxhm_info_id` | `XxhmInfoId` | `*string` | 未标注 | 先享后买id |

<a id="model-preorderextinfo"></a>

### PreOrderExtInfo

类型定义：[PreOrderExtInfo](credit_model.go#L29)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_id` | `BrandId` | `*string` | 未标注 | 品牌id |
| `brand_name` | `BrandName` | `*string` | 未标注 | 品牌名称 |
| `channel` | `Channel` | `*string` | 未标注 | 请求渠道 |
| `order_feature` | `OrderFeature` | `*string` | 未标注 | 单据属性,风控系统使用 |

<a id="model-settleextraparams"></a>

### SettleExtraParams

类型定义：[SettleExtraParams](credit_model.go#L41)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `quit_type` | `QuitType` | `*string` | 未标注 | quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT |
| `scene_pay_link_info` | `ScenePayLinkInfo` | `*string` | 未标注 | 商链通权益抵扣信息, 当前只有运营商使用, 未使用权益抵扣忽略该字段 fundRuleList为对应的出资信息，目前该列表不支持多对象，长度限制为1。其中fundAgreementNo为本笔交易实际出资的协议，由(商链通二方通用版本接入文档)2.3.2获取，fundAmount为该出资协议出资的金额。 bizScene和subBizScene为(商链通二方通用版本接入文档)2.1.3中支付宝侧分配的信息。 bizMode=MERCHANT_ORDER为固定值。 |
| `settle_adjust_reason` | `SettleAdjustReason` | `*string` | 未标注 | action_type选择PAY_TO_ZERO时必填， "SERVICE_CANCELED":"服务已取消", "OTHER_CHANNEL_PERFORMANCE":"户已通过其他方式履约" |

<a id="model-tasktypedata"></a>

### TaskTypeData

类型定义：[TaskTypeData](credit_model.go#L15)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `name` | `Name` | `*string` | 未标注 | 商户数据回传的任务名称，供C端展示。当累计的数据类型为任务类型时，此为必传字段 |

<a id="model-unfreezeextendparams"></a>

### UnfreezeExtendParams

类型定义：[UnfreezeExtendParams](credit_model.go#L51)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `quit_type` | `QuitType` | `*string` | 未标注 | 退出类型 |
| `total_discount_amount` | `TotalDiscountAmount` | `*string` | 未标注 | 芝麻Go开始到到期过程中享受的优惠总金额 |
| `total_real_pay_amount` | `TotalRealPayAmount` | `*string` | 未标注 | 芝麻Go开始到到期过程中任务相关支付的总金额。 |
| `total_task_count` | `TotalTaskCount` | `*string` | 未标注 | 芝麻Go开始到到期过程中和任务相关的完成次数 |

<a id="model-zmgooutdiscountinfo"></a>

### ZMGoOutDiscountInfo

类型定义：[ZMGoOutDiscountInfo](credit_model.go#L63)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 商户回传的优惠金额，如用户享受的红包金额，单位元 |
| `discount_name` | `DiscountName` | `*string` | 未标注 | 商户数据回传的优惠信息的名称 |
| `discount_type` | `DiscountType` | `*string` | 未标注 | 优惠信息的类型。为枚举值，商户可任意选择一个值作为入参传入。可取值： ("exclusiveBenefit","专享优惠"), ("exclusiveDiscount","专享折扣"), ("appreciationBenefit", "增值权益"), ("memberPoint", "会员积分") |

<a id="model-zmgotradeinfo"></a>

### ZMGoTradeInfo

类型定义：[ZMGoTradeInfo](credit_model.go#L73)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 用户和商户发生交易的交易单金额，单位元 |
| `trade_no` | `TradeNo` | `*string` | 未标注 | 用户和商户发生交易的交易单号，仅在商户需要回传交易信息时传入。 |
