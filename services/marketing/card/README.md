# card 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/card`。本包提供 9 个支付宝 V3 REST 操作。

通过 `card.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`benefit`](benefit/README.md) | 4 | `CardBenefitCreate` / `CardBenefitDelete` / `CardBenefitModify` / `CardBenefitQuery` |
| [`template`](template/README.md) | 4 | `CardTemplateBatchQuery` / `CardTemplateCreate` / `CardTemplateModify` / `CardTemplateQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CardActivateFormQuery`](#api-cardactivateformquery) | `alipay.marketing.card.activateform.query` |
| [`CardActivateURLApply`](#api-cardactivateurlapply) | `alipay.marketing.card.activateurl.apply` |
| [`CardConsumeSync`](#api-cardconsumesync) | `alipay.marketing.card.consume.sync` |
| [`CardDelete`](#api-carddelete) | `alipay.marketing.card.delete` |
| [`CardFormTemplateSet`](#api-cardformtemplateset) | `alipay.marketing.card.formtemplate.set` |
| [`CardMessageNotify`](#api-cardmessagenotify) | `alipay.marketing.card.message.notify` |
| [`CardOpen`](#api-cardopen) | `alipay.marketing.card.open` |
| [`CardQuery`](#api-cardquery) | `alipay.marketing.card.query` |
| [`CardUpdate`](#api-cardupdate) | `alipay.marketing.card.update` |

## 使用示例

- [CardActivateFormQuery](api_card_cardactivateformquery_example_test.go)
- [CardActivateURLApply](api_card_cardactivateurlapply_example_test.go)
- [CardConsumeSync](api_card_cardconsumesync_example_test.go)
- [CardDelete](api_card_carddelete_example_test.go)
- [CardFormTemplateSet](api_card_cardformtemplateset_example_test.go)
- [CardMessageNotify](api_card_cardmessagenotify_example_test.go)
- [CardOpen](api_card_cardopen_example_test.go)
- [CardQuery](api_card_cardquery_example_test.go)
- [CardUpdate](api_card_cardupdate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-cardactivateformquery"></a>

### CardActivateFormQuery — 查询用户提交的会员卡表单信息

- 官方接口：[`alipay.marketing.card.activateform.query`](https://opendocs.alipay.com/open-v3/a195d7c2_alipay.marketing.card.activateform.query)
- 路由：`GET /v3/alipay/marketing/card/activateform/query`
- 请求类型：[CardActivateFormQueryRequest](card_request.go#L119)
- 响应类型：[CardActivateFormQueryResponse](card_response.go#L8)

会员卡开卡场景下，用户确认领卡后，跳转到商户开卡处理页面，商户通过该接口查询用户表单信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `biz_type` | `BizType` | `*string` | query | 否 | 开放表单信息查询业务类型 |
| `template_id` | `TemplateId` | `*string` | query | 否 | 会员卡模板id。使用会员卡模板创建接口(alipay.marketing.card.template.create)返回的结果 |
| `request_id` | `RequestId` | `*string` | query | 否 | 查询用户表单提交信息的请求id。 网页移动应用&小程序应用：与auth_code 一起通过alipay.marketing.card.activateurl.apply(获取会员卡领卡投放链接)接口callback地址回传。 注意：auth_code（用户授权码）用于换取用户授权令牌 accessToken，后续调用接口需传入accessToken。 |

<a id="api-cardactivateurlapply"></a>

### CardActivateURLApply — 获取会员卡领卡投放链接

- 官方接口：[`alipay.marketing.card.activateurl.apply`](https://opendocs.alipay.com/open-v3/85f9492f_alipay.marketing.card.activateurl.apply)
- 路由：`POST /v3/alipay/marketing/card/activateurl/apply`
- 请求类型：[CardActivateURLApplyRequest](card_request.go#L4)
- 响应类型：[CardActivateURLApplyResponse](card_response.go#L14)

会员卡开卡业务，开发者通过该接口获取用户开卡链接，用于会员卡投放。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `callback` | `Callback` | `*string` | JSON body | 未标注 | 会员卡开卡表单提交后回调地址。要求必须是单纯的服务端接收回调。 说明： 1.该地址不可带参数，如需回传参数，可设置out_string入参。 |
| `follow_app_id` | `FollowAppId` | `*string` | JSON body | 未标注 | 需要关注的生活号AppId。若需要在领卡页面展示“关注生活号”提示，需开通生活号并绑定会员卡。生活号快速接入详见：https://doc.open.alipay.com/docs/doc.htm?treeId=193&articleId=105933&docType=1 |
| `out_string` | `OutString` | `*string` | JSON body | 未标注 | 该值为商家拉起开卡组件的传递的 out_string 值。通常可用于区分不同业务场景，禁止将该字段作为开卡流程的必要字段，否则会导致会员推广场景下的入会失败 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 会员卡模板id。使用会员卡模板创建接口(alipay.marketing.card.template.create)返回的结果 |

<a id="api-cardconsumesync"></a>

### CardConsumeSync — 会员卡消费记录同步

- 官方接口：[`alipay.marketing.card.consume.sync`](https://opendocs.alipay.com/open-v3/e0f362b2_alipay.marketing.card.consume.sync)
- 路由：`POST /v3/alipay/marketing/card/consume/sync`
- 请求类型：[CardConsumeSyncRequest](card_request.go#L16)
- 响应类型：[CardConsumeSyncResponse](card_response.go#L20)

会员卡消费记录同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `act_pay_amount` | `ActPayAmount` | `*string` | JSON body | 未标注 | 用户实际付的现金金额 1.针对预付卡面额的核销金额在use_benefit_list展现，作为权益金额 2.权益金额不叠加在该金额上 |
| `card_info` | `CardInfo` | [`*MerchantCard`](#model-merchantcard) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `gain_benefit_list` | `GainBenefitList` | [`*[]BenefitInfoDetail`](#model-benefitinfodetail) | JSON body | 未标注 | 获取权益列表 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 备注信息，现有直接填写门店信息 |
| `shop_code` | `ShopCode` | `*string` | JSON body | 未标注 | 门店编号 |
| `swipe_cert_type` | `SwipeCertType` | `*string` | JSON body | 未标注 | 产生该笔交易时，用户出具的凭证类型。枚举支持： *ALIPAY：支付宝电子卡； *ENTITY：实体卡； *OTHER：其他。 |
| `target_card_no` | `TargetCardNo` | `*string` | JSON body | 未标注 | 支付宝业务卡号，即通过[alipay.marketing.card.open](https://opendocs.alipay.com/apis/009zw3)(会员卡开卡)接口开卡后获取的 card_info.biz_card_no 值。 |
| `target_card_no_type` | `TargetCardNoType` | `*string` | JSON body | 未标注 | 卡号ID类型。支持： *BIZ_CARD：支付宝业务卡号（商户会员卡场景使用）。 |
| `trade_amount` | `TradeAmount` | `*string` | JSON body | 未标注 | 交易金额：本次交易的实际总金额（可认为标价金额） |
| `trade_name` | `TradeName` | `*string` | JSON body | 未标注 | 交易名称。为空时支付宝将根据交易类型提供默认名称。 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝交易号 |
| `trade_time` | `TradeTime` | `*string` | JSON body | 未标注 | 线下交易时间，为用户付款的交易时间。 说明：当交易时间晚于上次消费记录同步时间时，将变更会员卡信息。 |
| `trade_type` | `TradeType` | `*string` | JSON body | 未标注 | 交易类型。枚举支持： *开卡：OPEN； *消费：TRADE； *充值：DEPOSIT； *退卡：RETURN。 |
| `use_benefit_list` | `UseBenefitList` | [`*[]BenefitInfoDetail`](#model-benefitinfodetail) | JSON body | 未标注 | 实际消耗的权益 |

<a id="api-carddelete"></a>

### CardDelete — 会员卡删卡

- 官方接口：[`alipay.marketing.card.delete`](https://opendocs.alipay.com/open-v3/fa4f5056_alipay.marketing.card.delete?scene=common&pathHash=b9e0e504)
- 路由：`DELETE /v3/alipay/marketing/card/delete`
- 请求类型：[CardDeleteRequest](card_request.go#L131)
- 响应类型：[CardDeleteResponse](card_response.go#L26)

通过 API 接口删除商家会员卡。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_serial_no` | `OutSerialNo` | `*string` | query | 否 | 商户端删卡业务流水号，商户自定义且需确保流水号唯一性。 |
| `target_card_no` | `TargetCardNo` | `*string` | query | 否 | 支付宝业务卡号。即开卡接口返回结果中的card_info.biz_card_no。 |
| `target_card_no_type` | `TargetCardNoType` | `*string` | query | 否 | 卡号ID类型 |
| `reason_code` | `ReasonCode` | `*string` | query | 否 | 删卡原因 USER_UNBUND：用户解绑（可以重新绑定） CANCEL：销户（完成销户后，就不能再重新绑定） PRESENT：转赠（可以重新绑定） |
| `ext_info` | `ExtInfo` | `*string` | query | 否 | 删卡扩展参数，json格式。 用于商户的特定业务信息的传递，只有商户与支付宝约定了传递此参数且约定了参数含义，此参数才有效。 目前支持如下key： new_card_no：新卡号 donee_user_id：受赠人userId |

<a id="api-cardformtemplateset"></a>

### CardFormTemplateSet — 会员卡开卡表单模板配置

- 官方接口：[`alipay.marketing.card.formtemplate.set`](https://opendocs.alipay.com/open-v3/c7503ca0_alipay.marketing.card.formtemplate.set?scene=common&pathHash=36809f41)
- 路由：`POST /v3/alipay/marketing/card/formtemplate/set`
- 请求类型：[CardFormTemplateSetRequest](card_request.go#L47)
- 响应类型：[CardFormTemplateSetResponse](card_response.go#L32)

会员卡开卡、用户授权确认

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `fields` | `Fields` | [`*OpenFormFieldDO`](#model-openformfielddo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。 |

<a id="api-cardmessagenotify"></a>

### CardMessageNotify — 会员卡消息通知

- 官方接口：[`alipay.marketing.card.message.notify`](https://opendocs.alipay.com/open-v3/b710cee9_alipay.marketing.card.message.notify?scene=common&pathHash=9eb33531)
- 路由：`POST /v3/alipay/marketing/card/message/notify`
- 请求类型：[CardMessageNotifyRequest](card_request.go#L54)
- 响应类型：[CardMessageNotifyResponse](card_response.go#L35)

商户触发消息通知用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `notify_info` | `NotifyInfo` | `*string` | JSON body | 未标注 | 消息通知详情，包含消息通知需要的信息，map 格式存储。 积分过期提醒： <br>expired_point 过期积分值 <br>expired_time 过期时间 |
| `notify_no` | `NotifyNo` | `*string` | JSON body | 未标注 | 消息通知幂等标识，业务流水号 |
| `notify_type` | `NotifyType` | `*string` | JSON body | 未标注 | 消息通知类型, <br>POINT_REMIND 积分过期提醒 |
| `occur_time` | `OccurTime` | `*string` | JSON body | 未标注 | 标识业务发生时间，时间格式为"yyyy-MM-dd HH:mm:ss"。 |
| `target_card_no` | `TargetCardNo` | `*string` | JSON body | 未标注 | 支付宝业务卡号，即通过 alipay.marketing.card.open 接口开卡后获取的 card_info.biz_card_no 值。 |
| `target_card_no_type` | `TargetCardNoType` | `*string` | JSON body | 未标注 | 卡号ID类型 |

<a id="api-cardopen"></a>

### CardOpen — 会员卡开卡

- 官方接口：[`alipay.marketing.card.open`](https://opendocs.alipay.com/open-v3/8e9fe25b_alipay.marketing.card.open)
- 路由：`POST /v3/alipay/marketing/card/open`
- 请求类型：[CardOpenRequest](card_request.go#L70)
- 响应类型：[CardOpenResponse](card_response.go#L41)

商户通过API接口，进行会员卡开卡。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `card_ext_info` | `CardExtInfo` | [`*MerchantCard`](#model-merchantcard) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `card_template_id` | `CardTemplateId` | `*string` | JSON body | 未标注 | 支付宝分配的卡模板Id（卡模板创建接口返回的模板ID） |
| `card_user_info` | `CardUserInfo` | [`*CardUserInfo`](#model-carduserinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `member_ext_info` | `MemberExtInfo` | [`*MerchantMenber`](#model-merchantmenber) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `open_card_channel` | `OpenCardChannel` | `*string` | JSON body | 未标注 | 领卡渠道，用于记录外部商户端领卡来源的渠道信息，渠道值可自行定义（仅限数字、字母、下划线） 可直接标识领卡渠道，也可配合open_card_channel_id标识领卡渠道类型： 例如： 线下门店领取:20161534000000000008863（直接标识领卡渠道，门店shopId） 线下扫二维码领取:QR（标识领卡类型）; 线下活动领取:20170522000000000003609（直接标识领卡渠道，商户活动ID） |
| `open_card_channel_id` | `OpenCardChannelId` | `*string` | JSON body | 未标注 | 领卡来源的渠道id，注意区别于open_card_channel领卡渠道； 一般使用场景： open_card_channel用于区分渠道类型，例如取值为"SHOP"（门店），"ACTIVITY"（活动）； 则open_card_channel_id可用于区分同渠道的不同实体，对应取各门店ID或各活动的标识ID等； |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部商户流水号。由商户自定义，需保证商户系统中唯一。 |
| `paid_outer_card_info` | `PaidOuterCardInfo` | [`*PaidOuterCardExtraInfoDTO`](#model-paidoutercardextrainfodto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-cardquery"></a>

### CardQuery — 会员卡查询

- 官方接口：[`alipay.marketing.card.query`](https://opendocs.alipay.com/open-v3/b2494988_alipay.marketing.card.query?scene=common&pathHash=842cfce1)
- 路由：`POST /v3/alipay/marketing/card/query`
- 请求类型：[CardQueryRequest](card_request.go#L88)
- 响应类型：[CardQueryResponse](card_response.go#L50)

根据卡号或者持卡人信息查询会员卡信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `card_user_info` | `CardUserInfo` | [`*CardUserInfo`](#model-carduserinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 扩展信息，暂时没有 |
| `target_card_no` | `TargetCardNo` | `*string` | JSON body | 未标注 | 操作卡号。 若target_card_no_type为 BIZ_CARD，则该值为业务卡号，即alipay.marketing.card.open(会员卡开卡)返回结果中的card_info#biz_card_no。 若target_card_no_type为 D_QR_CODE 或D_BAR_CODE，则该值为动态码值。 若target_card_no_type为 EXTERNAL_CARD，则该值为外部卡号。 |
| `target_card_no_type` | `TargetCardNoType` | `*string` | JSON body | 未标注 | 卡号ID类型 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 会员卡模板id。调用[alipay.marketing.card.template.create](https://opendocs.alipay.com/open/03sx80?scene=common&pathHash=04270a28)（会员卡模板创建接口）创建模板后同步返回。 |

<a id="api-cardupdate"></a>

### CardUpdate — 会员卡更新

- 官方接口：[`alipay.marketing.card.update`](https://opendocs.alipay.com/open-v3/c83a3e12_alipay.marketing.card.update?scene=common&pathHash=dc8a305e)
- 路由：`POST /v3/alipay/marketing/card/update`
- 请求类型：[CardUpdateRequest](card_request.go#L101)
- 响应类型：[CardUpdateResponse](card_response.go#L60)

会员卡更新

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `card_info` | `CardInfo` | [`*MerchantCard`](#model-merchantcard) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 扩展信息(暂时无用) |
| `mcard_style_info` | `McardStyleInfo` | [`*McardStylInfo`](#model-mcardstylinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_card_msg_info` | `MerchantCardMsgInfo` | [`*MerchantCardMsgInfo`](#model-merchantcardmsginfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `notify_messages` | `NotifyMessages` | [`*[]McardNotifyMessage`](#model-mcardnotifymessage) | JSON body | 未标注 | 注意：此字段已废弃。 卡信息变更通知消息 1、在列表中定义的消息，才会发送给用户，消息格式一定。 2、根据卡信息变更情况，一次可发送多个消息 |
| `occur_time` | `OccurTime` | `*string` | JSON body | 未标注 | 业务发生时间，时间格式为"yyyy-MM-dd HH:mm:ss"。 |
| `paid_outer_card_info` | `PaidOuterCardInfo` | [`*PaidOuterCardExtraInfoDTO`](#model-paidoutercardextrainfodto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `target_card_no` | `TargetCardNo` | `*string` | JSON body | 未标注 | 支付宝业务卡号：1、3.0（历史版本）版本来源于alipay.marketing.card.open(会员卡开卡)接口开卡后获取的 card_info.biz_card_no 值；2、极速版来源于spi.alipay.user.opencard.get(会员卡开通，获取会员卡信息)请求参数中的biz_card_no值。 |
| `target_card_no_type` | `TargetCardNoType` | `*string` | JSON body | 未标注 | 卡号ID类型 |

## 嵌套请求参数

<a id="model-benefitinfodetail"></a>

### BenefitInfoDetail

类型定义：[BenefitInfoDetail](card_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | PRE_FUND：实际核销或者商户赠送的金额 DISCOUNT：实际折扣掉的金额（获取权益不支持该类型） COUPON：实际核销或者商户赠送的券 |
| `benefit_type` | `BenefitType` | `*string` | 未标注 | 权益类型 |
| `count` | `Count` | `*string` | 未标注 | COUPON：当核销或者赠送券时，可以设置该值 |
| `description` | `Description` | `*string` | 未标注 | 产生核销或者赠送权益的描述 |

<a id="model-cardfronttextdto"></a>

### CardFrontTextDTO

类型定义：[CardFrontTextDTO](card_model.go#L16)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `label` | `Label` | `*string` | 未标注 | 文案标签 |
| `value` | `Value` | `*string` | 未标注 | 展示文案 |

<a id="model-carduserinfo"></a>

### CardUserInfo

类型定义：[CardUserInfo](card_model.go#L24)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `open_id` | `OpenId` | `*string` | 未标注 | 用户open_id |
| `user_uni_id` | `UserUniId` | `*string` | 未标注 | 用户唯一标识, 根据user_uni_id_type类型来定 （目前暂支持支付宝userId） 支付宝userId说明：支付宝用户号是以2088开头的16位纯数字组成 |
| `user_uni_id_type` | `UserUniIdType` | `*string` | 未标注 | 用户唯一标识类型 |

<a id="model-mcardnotifymessage"></a>

### McardNotifyMessage

类型定义：[McardNotifyMessage](card_model.go#L34)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `change_reason` | `ChangeReason` | `*string` | 未标注 | 用户提醒信息，按如下格式拼装，需要ISV提供change_reason。 积分变动模板：{change_reason}，您的积分有变动 余额变动模板：{change_reason}，您的余额有变动 等级变更无需提供原因。 |
| `ext_info` | `ExtInfo` | `*string` | 未标注 | JSON格式扩展信息，主要是发送消息中的变量 |
| `message_type` | `MessageType` | `*string` | 未标注 | 消息类型，每种消息都定义了固定消息模板 |

<a id="model-mcardstylinfo"></a>

### McardStylInfo

类型定义：[McardStylInfo](card_model.go#L44)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `background_id` | `BackgroundId` | `*string` | 未标注 | 背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif； 尺寸不小于1020*643px； 图片不得有圆角，不得拉伸变形 |
| `bg_color` | `BgColor` | `*string` | 未标注 | 背景色 |
| `logo_id` | `LogoId` | `*string` | 未标注 | logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：1M以内，格式bmp、png、jpeg、jpg、gif； 尺寸不小于500*500px的正方形； 请优先使用商家LOGO； |

<a id="model-mdcodeinfodto"></a>

### MdCodeInfoDTO

类型定义：[MdCodeInfoDTO](card_model.go#L54)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `code_status` | `CodeStatus` | `*string` | 未标注 | 本次回传动态码的状态： SUCCESS: 本次发码成功 FAIL_RETRY: 本次发码失败，且需要支付宝重试（重新通知商户发码） FAIL_NOT_RETRY: 本次发码失败，且无需支付宝重试（无需重新通知商户发码） |
| `code_value` | `CodeValue` | `*string` | 未标注 | 动态码的码值： code_status为SUCCESS时必填； 基于此码值生成条形码或二维码用于扫码核销。 |
| `expire_time` | `ExpireTime` | `*string` | 未标注 | 当前动态码的过期（失效）时间： code_status为SUCCESS时必填。 |
| `time_stamp` | `TimeStamp` | `*int64` | 未标注 | 商户回传动态码的时间戳 (单位秒)。 即商户调接口回传动态码时刻对应的long类型时间戳，用于区分不同的发码请求。 |

<a id="model-merchantcard"></a>

### MerchantCard

类型定义：[MerchantCard](card_model.go#L66)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `balance` | `Balance` | `*string` | 未标注 | 资金卡余额，单位：元，精确到小数点后两位。 |
| `biz_card_no` | `BizCardNo` | `*string` | 未标注 | 支付宝业务卡号 说明： 1、开卡成功后返回该参数，需要保存留用； 2、开卡/更新/删卡/查询卡接口请求中不需要传该参数； |
| `custom_assets` | `CustomAssets` | `*string` | 未标注 | 会员卡自定义资产值，只供特定业务使用，通常接入无需关注 |
| `external_card_no` | `ExternalCardNo` | `*string` | 未标注 | 用户在商户crm系统中的会员卡卡号，该参数必填。 |
| `front_image_id` | `FrontImageId` | `*string` | 未标注 | 卡面展示图片的图片ID，通过接口（alipay.offline.material.image.upload）上传图片 这里预期展示的是个人照片； 图片说明：1M以内，格式bmp、png、jpeg、jpg、gif； 图片尺寸为230*295px，可等比放大； |
| `front_text_list` | `FrontTextList` | [`*[]CardFrontTextDTO`](#model-cardfronttextdto) | 未标注 | 卡面文案列表，1项对应1行文案，最多只能传入4行文案；<br> 单行文案展现分为左右两部分，左边对应label字段，右边对应value；<br> 形如： 学院 新闻学院 |
| `level` | `Level` | `*string` | 未标注 | 会员卡等级（由商户自定义，并可以在卡模板创建时，定义等级信息） |
| `mdcode_info` | `MdcodeInfo` | [`*MdCodeInfoDTO`](#model-mdcodeinfodto) | 未标注 | 见字段类型及官方接口文档。 |
| `open_date` | `OpenDate` | `*string` | 未标注 | 会员卡开卡时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。 |
| `point` | `Point` | `*string` | 未标注 | 会员卡积分，积分必须为数字型（可为浮点型，带2位小数点） |
| `template_id` | `TemplateId` | `*string` | 未标注 | 会员卡更换不同的卡模板（该参数仅用在会员卡更新接口中） |
| `valid_date` | `ValidDate` | `*string` | 未标注 | 会员卡有效期结束时间，格式为yyyy-MM-dd HH:mm:ss。会员卡更新时，该时间不支持修改。 |

<a id="model-merchantcardmsginfo"></a>

### MerchantCardMsgInfo

类型定义：[MerchantCardMsgInfo](card_model.go#L93)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `changed_point` | `ChangedPoint` | `*string` | 未标注 | 本次消息通知的变动积分，积分必须为数字型（可为浮点型，带2位小数点），格式非法则不发送消息。例如： 100，则代表增加100积分。 -100，则代表减少100积分。 |

<a id="model-merchantmenber"></a>

### MerchantMenber

类型定义：[MerchantMenber](card_model.go#L99)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `birth` | `Birth` | `*string` | 未标注 | 生日 yyyy-MM-dd |
| `cell` | `Cell` | `*string` | 未标注 | 手机号 |
| `gende` | `Gende` | `*string` | 未标注 | 性别 |
| `name` | `Name` | `*string` | 未标注 | 姓名 |

<a id="model-openformfielddo"></a>

### OpenFormFieldDO

类型定义：[OpenFormFieldDO](card_model.go#L111)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `optional` | `Optional` | `*string` | 未标注 | 表单可选字段配置，common_fields属性定义一个表单字段数组，表单字段有效值列表与required字段有效值列表相同。 可选字段配置中不能含有必须字段配置的有效值。 |
| `required` | `Required` | `*string` | 未标注 | 表单必填字段配置，common_fields属性定义一个表单字段数组，字段有效值如下列表所示： OPEN_FORM_FIELD_MOBILE -- 手机号 OPEN_FORM_FIELD_GENDER -- 性别 OPEN_FORM_FIELD_NAME -- 姓名 OPEN_FORM_FIELD_BIRTHDAY -- 生日 OPEN_FORM_FIELD_IDCARD -- 身份证 OPEN_FORM_FIELD_EMAIL -- 邮箱 OPEN_FORM_FIELD_ADDRESS -- 地址 OPEN_FORM_FIELD_CITY -- 城市 OPEN_FORM_FIELD_IS_STUDENT -- 是否学生认证 OPEN_FORM_FIELD_MEMBER_GRADE -- 会员等级 |

<a id="model-paidoutercardcycleinfodto"></a>

### PaidOuterCardCycleInfoDTO

类型定义：[PaidOuterCardCycleInfoDTO](card_model.go#L119)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `alipay_deduct_agreement` | `AlipayDeductAgreement` | `*string` | 未标注 | 支付宝用户签约协议号 |
| `alipay_deduct_product_code` | `AlipayDeductProductCode` | `*string` | 未标注 | 支付宝代扣产品码 |
| `alipay_deduct_scene` | `AlipayDeductScene` | `*string` | 未标注 | 支付宝代扣场景码 |
| `close_reason` | `CloseReason` | `*string` | 未标注 | 关闭连续购买原因。 |
| `cycle_type` | `CycleType` | `*string` | 未标注 | 用户开通连续购买类型。 |
| `open_status` | `OpenStatus` | `*string` | 未标注 | 开通连续购买状态。 |

<a id="model-paidoutercardextrainfodto"></a>

### PaidOuterCardExtraInfoDTO

类型定义：[PaidOuterCardExtraInfoDTO](card_model.go#L135)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action` | `Action` | `*string` | 未标注 | 用户操作类型 |
| `cycle_info` | `CycleInfo` | [`*PaidOuterCardCycleInfoDTO`](#model-paidoutercardcycleinfodto) | 未标注 | 见字段类型及官方接口文档。 |
| `purchase_info` | `PurchaseInfo` | [`*PaidOuterCardPurchaseInfoDTO`](#model-paidoutercardpurchaseinfodto) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-paidoutercardpurchaseinfodto"></a>

### PaidOuterCardPurchaseInfoDTO

类型定义：[PaidOuterCardPurchaseInfoDTO](card_model.go#L143)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_date` | `ActionDate` | `*string` | 未标注 | 用户购买、升级、降级、续费的操作时间。格式为：yyyy-MM-dd hh:mm:ss |
| `alipay_trade_no` | `AlipayTradeNo` | `*string` | 未标注 | 支付宝交易号，由api: alipay.trade.create 返回 |
| `out_trade_no` | `OutTradeNo` | `*string` | 未标注 | 商户订单号。与创建订单api：alipay.trade.create保持一致 |
| `price` | `Price` | `*string` | 未标注 | 用户购买金额。当购买/升级/续费场景必填，单位元，精确到小数点后2位 |
| `source` | `Source` | `*string` | 未标注 | 用户操作来源 |
