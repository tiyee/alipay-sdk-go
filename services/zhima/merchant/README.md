# zhima/merchant 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/zhima/merchant`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `merchant.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_merchant_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`SubsidiariesApply`](#api-subsidiariesapply) | [`zhima.merchant.subsidiaries.apply`](https://opendocs.alipay.com/apis/05ygh0) | 子商户信息添加 | `POST` |
| [`SubsidiariesClose`](#api-subsidiariesclose) | [`zhima.merchant.subsidiaries.close`](https://opendocs.alipay.com/apis/05yopn) | 关闭子商户信息工单 | `POST` |
| [`ZmGoCumulateQuery`](#api-zmgocumulatequery) | [`zhima.merchant.zmgo.cumulate.query`](https://opendocs.alipay.com/open-v3/62aec983_zhima.merchant.zmgo.cumulate.query) | 商家芝麻GO累计数据查询接口 | `GET` |
| [`ZmGoCumulateSync`](#api-zmgocumulatesync) | [`zhima.merchant.zmgo.cumulate.sync`](https://opendocs.alipay.com/open-v3/153c048c_zhima.merchant.zmgo.cumulate.sync) | 商家芝麻GO累计数据回传接口 | `POST` |
| [`ZmGoTemplateCreate`](#api-zmgotemplatecreate) | [`zhima.merchant.zmgo.template.create`](https://opendocs.alipay.com/open-v3/b0eba80a_zhima.merchant.zmgo.template.create) | 商户创建芝麻GO模板接口 | `POST` |
| [`ZmGoTemplateQuery`](#api-zmgotemplatequery) | [`zhima.merchant.zmgo.template.query`](https://opendocs.alipay.com/open-v3/1cbec0d3_zhima.merchant.zmgo.template.query) | 芝麻GO模板查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-subsidiariesapply"></a>

### SubsidiariesApply — 子商户信息添加

- 官方接口：[`zhima.merchant.subsidiaries.apply`](https://opendocs.alipay.com/apis/05ygh0)
- 路由：`POST /v3/zhima/merchant/subsidiaries/apply`
- 请求类型：[SubsidiariesApplyRequest](merchant_request.go#L5)
- 响应类型：[SubsidiariesApplyResponse](merchant_response.go#L5)

子商户信息添加

请求约束：SubsidiariesApplyRequest is the official V3 ZhimaMerchantSubsidiariesApplyModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 商户pid |
| `service_id` | `ServiceId` | `*string` | JSON body | 未标注 | 信用服务id |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 二级商户进件id |
| `sub_merchant_contact_number` | `SubMerchantContactNumber` | `*string` | JSON body | 未标注 | 子商户联系电话 |
| `sub_merchant_jump_link` | `SubMerchantJumpLink` | `*string` | JSON body | 未标注 | 子商户跳转链接 |
| `sub_merchant_logo_url` | `SubMerchantLogoUrl` | `*string` | JSON body | 未标注 | 子商户logo地址 |
| `sub_merchant_name` | `SubMerchantName` | `*string` | JSON body | 未标注 | 子商户名 |
| `sub_pid` | `SubPid` | `*string` | JSON body | 未标注 | 子商户id |

<a id="api-subsidiariesclose"></a>

### SubsidiariesClose — 关闭子商户信息工单

- 官方接口：[`zhima.merchant.subsidiaries.close`](https://opendocs.alipay.com/apis/05yopn)
- 路由：`POST /v3/zhima/merchant/subsidiaries/close`
- 请求类型：[SubsidiariesCloseRequest](merchant_request.go#L25)
- 响应类型：[SubsidiariesCloseResponse](merchant_response.go#L15)

关闭子商户信息工单

请求约束：SubsidiariesCloseRequest is the official V3 ZhimaMerchantSubsidiariesCloseModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 工单标识 |

<a id="api-zmgocumulatequery"></a>

### ZmGoCumulateQuery — 商家芝麻GO累计数据查询接口

- 官方接口：[`zhima.merchant.zmgo.cumulate.query`](https://opendocs.alipay.com/open-v3/62aec983_zhima.merchant.zmgo.cumulate.query)
- 路由：`GET /v3/zhima/merchant/zmgo/cumulate/query`
- 请求类型：[ZmGoCumulateQueryRequest](merchant_request.go#L79)
- 响应类型：[ZmGoCumulateQueryResponse](merchant_response.go#L29)

芝麻GO数据累计查询接口，包含汇总和明细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | query | 否 | 芝麻go协议号，唯一标识一个芝麻go协议。 |
| `user_id` | `UserId` | `*string` | query | 否 | 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `open_id` | `OpenId` | `*string` | query | 否 | 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `provider_pid` | `ProviderPid` | `*string` | query | 否 | 数据回传的商户 ID，即和用户发生业务来往的 PID 主体。商户账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `need_detail` | `NeedDetail` | `*bool` | query | 否 | 是否需要回传明细。枚举支持： * true：需要回传明细。 * false：不需要回传明细。 |
| `page_no` | `PageNo` | `*int64` | query | 否 | 分页参数，当need_detail=true时需要，代表当前页数，从1开始。 |
| `page_size` | `PageSize` | `*int64` | query | 否 | 分页参数，当need_detail=true时需要，代表每页大小，不能大于20。 |

<a id="api-zmgocumulatesync"></a>

### ZmGoCumulateSync — 商家芝麻GO累计数据回传接口

- 官方接口：[`zhima.merchant.zmgo.cumulate.sync`](https://opendocs.alipay.com/open-v3/153c048c_zhima.merchant.zmgo.cumulate.sync)
- 路由：`POST /v3/zhima/merchant/zmgo/cumulate/sync`
- 请求类型：[ZmGoCumulateSyncRequest](merchant_request.go#L31)
- 响应类型：[ZmGoCumulateSyncResponse](merchant_response.go#L49)

外部商家接入芝麻GO场景下，回传用户行为履约数据给芝麻。

请求约束：ZmGoCumulateSyncRequest is the official V3 ZhimaMerchantZmgoCumulateSyncModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_id` | `AgreementId` | `*string` | JSON body | 未标注 | 芝麻go协议号，唯一标识一个芝麻go协议。 |
| `amount_type_sync_data` | `AmountTypeSyncData` | [`*AmountTypeSyncData`](#model-amounttypesyncdata) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_action` | `BizAction` | `*string` | JSON body | 未标注 | 数据回传的动作类型。枚举如下： * POSITIVE：正向回传，即订单金额累加，优惠累加，任务次数累加。 * REVERSE：逆向回传，即订单金额退款，优惠退款，任务次数回退。 |
| `biz_time` | `BizTime` | `*string` | JSON body | 未标注 | 回传数据发生的实际时间，包括：订单交易时间，优惠核销时间，以及任务完成时间。为必传入参。 如：用户在【2019-03-08 00:00:00】核销了一个红包，但是【2019-03-10 00:00:00】才进行数据回传，此时biz_time取值红包核销的时间，也就是【2019-03-08 00:00:00】。 |
| `data_type` | `DataType` | `*string` | JSON body | 未标注 | 回传数据类型，枚举： TASK - 进度型，影响进度 DISCOUNT - 优惠型 |
| `discount_type_sync_data` | `DiscountTypeSyncData` | [`*DiscountTypeSyncData`](#model-discounttypesyncdata) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，唯一标识一笔回传数据。 |
| `provider_pid` | `ProviderPid` | `*string` | JSON body | 未标注 | 数据回传的商户 ID，即和用户发生业务来往的 PID 主体。商户账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |
| `refer_out_biz_no` | `ReferOutBizNo` | `*string` | JSON body | 未标注 | 逆向对应的正向外部业务号。逆向场景为必传入参。也就是biz_action为REVERSE时为必传入参。 逆向时需通过此字段明确对应的正向的外部业务号。 |
| `sub_biz_action` | `SubBizAction` | `*string` | JSON body | 未标注 | 数据回传动作子类型，枚举值，为必传入参。 ADD - 新增数据 UPDATE - 修改数据 备注：UPDATE只能更新金额信息。 |
| `times_type_sync_data` | `TimesTypeSyncData` | [`*TimesTypeSyncData`](#model-timestypesyncdata) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |

<a id="api-zmgotemplatecreate"></a>

### ZmGoTemplateCreate — 商户创建芝麻GO模板接口

- 官方接口：[`zhima.merchant.zmgo.template.create`](https://opendocs.alipay.com/open-v3/b0eba80a_zhima.merchant.zmgo.template.create)
- 路由：`POST /v3/zhima/merchant/zmgo/template/create`
- 请求类型：[ZmGoTemplateCreateRequest](merchant_request.go#L61)
- 响应类型：[ZmGoTemplateCreateResponse](merchant_response.go#L59)

外部商家接入芝麻GO场景下，为商户提供自助创建芝麻GO模板的能力。

请求约束：ZmGoTemplateCreateRequest is the official V3 ZhimaMerchantZmgoTemplateCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `basic_config` | `BasicConfig` | [`*ZMGOBasicConfig`](#model-zmgobasicconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `ext_config` | `ExtConfig` | [`*ZMGOExtConfig`](#model-zmgoextconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `obligation_config` | `ObligationConfig` | [`*ZMGOObligationConfig`](#model-zmgoobligationconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `open_config` | `OpenConfig` | [`*ZMGOOpenConfig`](#model-zmgoopenconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `quit_config` | `QuitConfig` | [`*ZMGOQuitConfig`](#model-zmgoquitconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `right_config` | `RightConfig` | [`*ZMGORightConfig`](#model-zmgorightconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `settlement_config` | `SettlementConfig` | [`*ZMGOSettlementConfig`](#model-zmgosettlementconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-zmgotemplatequery"></a>

### ZmGoTemplateQuery — 芝麻GO模板查询

- 官方接口：[`zhima.merchant.zmgo.template.query`](https://opendocs.alipay.com/open-v3/1cbec0d3_zhima.merchant.zmgo.template.query)
- 路由：`GET /v3/zhima/merchant/zmgo/template/query`
- 请求类型：[ZmGoTemplateQueryRequest](merchant_request.go#L97)
- 响应类型：[ZmGoTemplateQueryResponse](merchant_response.go#L65)

芝麻GO模板查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_no` | `TemplateNo` | `*string` | query | 否 | 模板ID，由创建模板接口返回结果提供 |
| `partner_id` | `PartnerId` | `*string` | query | 否 | 商户ID |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-amounttypesyncdata"></a>

### AmountTypeSyncData

类型定义：[AmountTypeSyncData](merchant_model.go#L31)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 商户回传的优惠金额，如用户享受的红包金额，单位元 |
| `discount_desc` | `DiscountDesc` | `*string` | 未标注 | 商户数据回传的优惠信息的名称。 |
| `has_alipay_trade` | `HasAlipayTrade` | `*bool` | 未标注 | 是否为支付宝交易，默认false，表示是否要做支付宝交易关联。此时tradeNo必须传支付宝交易号。 |
| `task_amount` | `TaskAmount` | `*string` | 未标注 | 用户和商户发生交易的交易单金额，单位元。 |
| `task_desc` | `TaskDesc` | `*string` | 未标注 | 任务描述 |
| `trade_no` | `TradeNo` | `*string` | 未标注 | 用户和商户发生金额类交易的交易单号。 |

<a id="model-discounttypesyncdata"></a>

### DiscountTypeSyncData

类型定义：[DiscountTypeSyncData](merchant_model.go#L47)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 商户回传的优惠金额，如用户享受的红包金额，单位元 |
| `discount_desc` | `DiscountDesc` | `*string` | 未标注 | 商户数据回传的优惠信息的名称。 |

<a id="model-timestypesyncdata"></a>

### TimesTypeSyncData

类型定义：[TimesTypeSyncData](merchant_model.go#L55)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 商户回传的优惠金额，如用户享受的红包金额，单位元 |
| `discount_desc` | `DiscountDesc` | `*string` | 未标注 | 商户数据回传的优惠信息的名称。 |
| `task_amount` | `TaskAmount` | `*string` | 未标注 | 用户和商户发生交易的交易单金额，单位元。 |
| `task_desc` | `TaskDesc` | `*string` | 未标注 | 任务描述 |
| `task_times` | `TaskTimes` | `*int64` | 未标注 | 当为次数型任务时必须传。 |

<a id="model-zmgobasicconfig"></a>

### ZMGOBasicConfig

类型定义：[ZMGOBasicConfig](merchant_model.go#L69)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `biz_type` | `BizType` | `*string` | 未标注 | 商户在芝麻GO配置的业务身份编码 |
| `contact` | `Contact` | `*string` | 未标注 | 商家客服电话 |
| `isv_pid` | `IsvPid` | `*string` | 未标注 | 运营商商户支付宝ID。若非ISV代理模式，也就是商户自运营模式，此属性取值与partner_id一致。 |
| `merchant_custom_logo` | `MerchantCustomLogo` | `*string` | 未标注 | 商户LOGO |
| `out_biz_no` | `OutBizNo` | `*string` | 未标注 | 外部业务单号，供幂等使用，需保证每次请求的值都不同 |
| `partner_id` | `PartnerId` | `*string` | 未标注 | 商户的支付宝ID，即为此商户创建芝麻GO模板 |
| `template_name` | `TemplateName` | `*string` | 未标注 | 芝麻GO模板名称 |
| `template_no` | `TemplateNo` | `*string` | 未标注 | 模板编号 |

<a id="model-zmgocycleflexconfig"></a>

### ZMGOCycleFlexConfig

类型定义：[ZMGOCycleFlexConfig](merchant_model.go#L183)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cycle_flex_withhold_fee_name` | `CycleFlexWithholdFeeName` | `*string` | 未标注 | 周期灵活扣的费用名称 |
| `cycle_flex_withhold_max_price` | `CycleFlexWithholdMaxPrice` | `*string` | 未标注 | 周期灵活扣单期扣款的最大额度 |
| `cycle_flex_withhold_total_period_count` | `CycleFlexWithholdTotalPeriodCount` | `*int64` | 未标注 | 周期灵活扣的总期数 |

<a id="model-zmgocyclewithholdconfig"></a>

### ZMGOCycleWithholdConfig

类型定义：[ZMGOCycleWithholdConfig](merchant_model.go#L193)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `deduct_plan` | `DeductPlan` | `*[]string` | 未标注 | 扣款计划 |
| `exemption_period` | `ExemptionPeriod` | `*string` | 未标注 | 豁免期天数 |
| `period` | `Period` | `*string` | 未标注 | 周期 |
| `period_type` | `PeriodType` | `*string` | 未标注 | 扣款周期类型，支持DAY-天模式，MONTH-月模式 |
| `support_cycle_withhold_high_mode` | `SupportCycleWithholdHighMode` | `*bool` | 未标注 | 是否支持周期扣高级模式，高级模式下扣款计划中内一起的扣款金额可以进行定制 |
| `support_exemption_period` | `SupportExemptionPeriod` | `*bool` | 未标注 | 是否支持豁免期，如果支持豁免期，用户开通后再豁免期内退出无需扣款 |
| `withhold_mode` | `WithholdMode` | `*string` | 未标注 | 设置扣款模式，目前支持每月固定日期或固定时间间隔 |

<a id="model-zmgoextconfig"></a>

### ZMGOExtConfig

类型定义：[ZMGOExtConfig](merchant_model.go#L89)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `text_content_fill_rule_id` | `TextContentFillRuleId` | `*string` | 未标注 | 文本内容填充规则ID |
| `text_content_fill_variable` | `TextContentFillVariable` | `*string` | 未标注 | 文本内容填充变量，变量key由运营同学提供 |

<a id="model-zmgoobligationconfig"></a>

### ZMGOObligationConfig

类型定义：[ZMGOObligationConfig](merchant_model.go#L97)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `benefit_url` | `BenefitUrl` | `*string` | 未标注 | 芝麻GO管理页做任务按钮对应链接 |
| `obligation_amount` | `ObligationAmount` | `*string` | 未标注 | 完成任务所需的最少额度，当obligation_template取值为：OBTL002时必传 |
| `obligation_template` | `ObligationTemplate` | `*string` | 未标注 | 承诺模式下，义务的类型，包含次数类型和金额类型： OBTL001("OBTL001", "在x时间内完成n次"), OBTL002("OBTL002", "在x时间内达到n元金额") |
| `obligation_times` | `ObligationTimes` | `*string` | 未标注 | 完成任务所需的最少次数，当obligation_template取值为：OBTL001时必填 |
| `promise_type_desc` | `PromiseTypeDesc` | `*string` | 未标注 | 任务描述 |
| `task_progress_redirect_schema` | `TaskProgressRedirectSchema` | `*string` | 未标注 | 芝麻GO管理页任务进度重定向链接 |

<a id="model-zmgoopenconfig"></a>

### ZMGOOpenConfig

类型定义：[ZMGOOpenConfig](merchant_model.go#L113)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `apply_button_desc` | `ApplyButtonDesc` | `*string` | 未标注 | 开通页按钮文案描述;传入内容后开通页点击开通按钮展示文案会使用该字段传入内容 |
| `appoint_date` | `AppointDate` | `*string` | 未标注 | 固定到期时间或每月的第几日。 （1）当period_mode为APPOINT_DATE时，取值固定日期，如"2022-06-01 00:00:00"，表示协议在6月1号0点到期; （2）当period_mode为FIX_DATE，取值"3"，表示协议在period_time个月后的3号到期。 |
| `card_color_scheme` | `CardColorScheme` | `*string` | 未标注 | 芝麻GO卡颜色。 C_BLACK("C_BLACK", "黑色"), C_RED("C_RED", "红色"), C_GREEN("C_GREEN", "绿色"), C_LIGHT_BLUE("C_LIGHT_BLUE", "浅蓝色"), C_BLUE("C_BLUE", "蓝色"), C_ORANGE("C_ORANGE","橙色"), C_PURPLE("C_PURPLE", "紫色"); |
| `custom_open_tip_list` | `CustomOpenTipList` | `*string` | 未标注 | 未完成任务规则详情 |
| `custom_open_tips` | `CustomOpenTips` | `*string` | 未标注 | 规则详情 |
| `freeze_amount` | `FreezeAmount` | `*string` | 未标注 | 担保金额额度，单位:元 |
| `min_sign_interval` | `MinSignInterval` | `*string` | 未标注 | 用户签约间隔时间，单位天 |
| `period_mode` | `PeriodMode` | `*string` | 未标注 | 芝麻GO义务到期的计算方式。包括： RELATIVE_DATE("RELATIVE_DATE", "相对日期"), ABSOLATE_DATE("ABSOLATE_DATE", "绝对日期(当日开始计算)"), ABSOLUTE_PLUS_1_DATE("ABSOLUTE_PLUS_1_DATE", "绝对日期(次日开始计算)"), FIX_DATE("FIX_DATE", "固定日期"), APPOINT_DATE("APPOINT_DATE", "指定日期"), |
| `period_time` | `PeriodTime` | `*string` | 未标注 | 周期时长 当period_mode为RELATIVE_DATE、ABSOLATE_DATE、ABSOLUTE_PLUS_1_DATE、FIX_DATE时必填 |
| `show_sign_success_page` | `ShowSignSuccessPage` | `*bool` | 未标注 | 是否展示签约完成页 |
| `sign_again_schema` | `SignAgainSchema` | `*string` | 未标注 | 用户签约商户活动到期 |
| `sign_success_task_button_desc` | `SignSuccessTaskButtonDesc` | `*string` | 未标注 | 签约完成页任务按钮描述，当展示签约完成页=true的时候 改内容必填 |
| `support_expire_deferral` | `SupportExpireDeferral` | `*bool` | 未标注 | 是否支持延期特权，开启后用户可以在任务到期前进行延期。(该功能已停用，不建议开启使用。) |

<a id="model-zmgoquitconfig"></a>

### ZMGOQuitConfig

类型定义：[ZMGOQuitConfig](merchant_model.go#L143)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `quit_desc` | `QuitDesc` | `*string` | 未标注 | 退出提示 |
| `quit_jump_url` | `QuitJumpUrl` | `*string` | 未标注 | 退出跳转商家页地址，当quit_type取值为：QUIT_MERCHANT_JUMP，必传 |
| `quit_type` | `QuitType` | `*string` | 未标注 | 退出芝麻GO方式，取值： （1）QUIT_COMMON("QUIT_COMMON", "标准退出"), （2）QUIT_MERCHANT_JUMP("QUIT_MERCHANT_JUMP", "商家页退出"), （3）QUIT_ONLY_TIPS("QUIT_ONLY_TIPS", "仅退出提示") |

<a id="model-zmgorightconfig"></a>

### ZMGORightConfig

类型定义：[ZMGORightConfig](merchant_model.go#L153)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cumulative_preferential_redirect_schema` | `CumulativePreferentialRedirectSchema` | `*string` | 未标注 | 芝麻GO管理页已享优惠进度的重定向链接 |
| `custom_benefit_desc` | `CustomBenefitDesc` | `*string` | 未标注 | 权益描述 |
| `custom_sub_benefit_desc` | `CustomSubBenefitDesc` | `*string` | 未标注 | 权益描述子标题 |

<a id="model-zmgosettlementconfig"></a>

### ZMGOSettlementConfig

类型定义：[ZMGOSettlementConfig](merchant_model.go#L163)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `custom_fee_name` | `CustomFeeName` | `*string` | 未标注 | 会员费扣费名称 |
| `cycle_flex_withhold_config` | `CycleFlexWithholdConfig` | [`*ZMGOCycleFlexConfig`](#model-zmgocycleflexconfig) | 未标注 | 见字段类型及官方接口文档。 |
| `cycle_withhold_config` | `CycleWithholdConfig` | [`*ZMGOCycleWithholdConfig`](#model-zmgocyclewithholdconfig) | 未标注 | 见字段类型及官方接口文档。 |
| `exp_stop_delay_days` | `ExpStopDelayDays` | `*int64` | 未标注 | 预计结算延迟天数，当exp_stop_time_mode取值RELATIVE_DATE、ABSOLATE_DATE、ABSOLUTE_PLUS_1_DATE时，必传 |
| `exp_stop_time` | `ExpStopTime` | `*string` | 未标注 | 预期结算时间，当exp_stop_time_mode取值APPOINT_DATE时，必传 |
| `exp_stop_time_mode` | `ExpStopTimeMode` | `*string` | 未标注 | 预计结算时间模式，取值：（1）RELATIVE_DATE("RELATIVE_DATE", "相对日期"), （2）ABSOLATE_DATE("ABSOLATE_DATE", "绝对日期(当日开始计算)"), （3）ABSOLUTE_PLUS_1_DATE("ABSOLUTE_PLUS_1_DATE", "绝对日期(次日开始计算)"), （4）APPOINT_DATE("APPOINT_DATE", "指定日期") |
| `fulfilltimes_custom_settlement_plan` | `FulfilltimesCustomSettlementPlan` | `*string` | 未标注 | 按履约次数灵活扣款计划。当settlement_type取值为fulfillTimesCustomSettlement时，此属性必传 |
| `settlement_type` | `SettlementType` | `*string` | 未标注 | 结算类型，取值： BENEFIT_SETTLEMENT("benefitSettlement", "按权益金额结算"), FULFILL_TIMES_CUSTOM_SETTLEMENT("fulfillTimesCustomSettlement","按履约次数扣款") |
