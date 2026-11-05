# marketing/activity/voucher 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/voucher`。本目录提供 7 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `voucher.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityVoucherAppend`](#api-activityvoucherappend) | [`alipay.marketing.activity.voucher.append`](https://opendocs.alipay.com/apis/05vsc4) | 追加支付券预算 | `POST` |
| [`ActivityVoucherCreate`](#api-activityvouchercreate) | [`alipay.marketing.activity.voucher.create`](https://opendocs.alipay.com/apis/05vb41) | 创建支付券 | `POST` |
| [`ActivityVoucherModify`](#api-activityvouchermodify) | [`alipay.marketing.activity.voucher.modify`](https://opendocs.alipay.com/apis/05v8k9) | 修改支付券基本信息 | `POST` |
| [`ActivityVoucherPublish`](#api-activityvoucherpublish) | [`alipay.marketing.activity.voucher.publish`](https://opendocs.alipay.com/apis/05vsc3) | 激活支付券 | `POST` |
| [`ActivityVoucherQuery`](#api-activityvoucherquery) | [`alipay.marketing.activity.voucher.query`](https://opendocs.alipay.com/apis/05veld) | 查询支付券详情 | `GET` |
| [`ActivityVoucherSend`](#api-activityvouchersend) | [`alipay.marketing.activity.voucher.send`](https://opendocs.alipay.com/apis/05vidi) | 发放支付券 | `POST` |
| [`ActivityVoucherStop`](#api-activityvoucherstop) | [`alipay.marketing.activity.voucher.stop`](https://opendocs.alipay.com/apis/05vc1k) | 停止支付券 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activityvoucherappend"></a>

### ActivityVoucherAppend — 追加支付券预算

- 官方接口：[`alipay.marketing.activity.voucher.append`](https://opendocs.alipay.com/apis/05vsc4)
- 路由：`POST /v3/alipay/marketing/activity/voucher/append`
- 请求类型：[ActivityVoucherAppendRequest](voucher_request.go#L5)
- 响应类型：[ActivityVoucherAppendResponse](voucher_response.go#L9)

支付券激活成功后，调用该接口追加支付券预算。追加支付券预算接口同步返回请求受理结果，后续通过alipay.marketing.activity.message.appended消息接口异步发送追加支付券预算成功事件消息。
如果成功消息长时间未送达，请联系相关技术人员。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `voucher_quantity` | `VoucherQuantity` | `*int64` | JSON body | 未标注 | 发行券的数量。。 限制： 券总预算&lt;= 99999999 特别说明： 该字段的含义是追加到的数量。 不可减少，只能增加。 |

<a id="api-activityvouchercreate"></a>

### ActivityVoucherCreate — 创建支付券

- 官方接口：[`alipay.marketing.activity.voucher.create`](https://opendocs.alipay.com/apis/05vb41)
- 路由：`POST /v3/alipay/marketing/activity/voucher/create`
- 请求类型：[ActivityVoucherCreateRequest](voucher_request.go#L17)
- 响应类型：[ActivityVoucherCreateResponse](voucher_response.go#L15)

商户通过调用此接口创建支付券，包括预充值&免充值支付券，接口同步返回创建处理结果。创建完成后将获得活动id，可用于各个营销场景的活动投放。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_base_info` | `ActivityBaseInfo` | [`*ActivityBaseInfo`](#model-activitybaseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_name` | `ActivityName` | `*string` | JSON body | 未标注 | 活动名称。 |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*PaymentVoucherBelongMerchantInfo`](#model-paymentvoucherbelongmerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `publish_end_time` | `PublishEndTime` | `*string` | JSON body | 未标注 | 券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss 限制： 券发放结束时间 publish_end_time 与 券发放开始时间 publish_start_time 间隔必须小于等于180天 |
| `publish_start_time` | `PublishStartTime` | `*string` | JSON body | 未标注 | 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `voucher_available_scope_info` | `VoucherAvailableScopeInfo` | [`*VoucherAvailableScopeInfo`](#model-voucheravailablescopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_budget_info` | `VoucherBudgetInfo` | [`*PaymentVoucherBudgetInfo`](#model-paymentvoucherbudgetinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_budget_supply_info` | `VoucherBudgetSupplyInfo` | [`*VoucherBudgetSupplyInfo`](#model-voucherbudgetsupplyinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_customer_guide_info` | `VoucherCustomerGuideInfo` | [`*VoucherCustomerGuideInfo`](#model-vouchercustomerguideinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_deduct_info` | `VoucherDeductInfo` | [`*VoucherDeductInfo`](#model-voucherdeductinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_info` | `VoucherDisplayInfo` | [`*PaymentVoucherDisplayInfo`](#model-paymentvoucherdisplayinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_pattern_info` | `VoucherDisplayPatternInfo` | [`*VoucherDisplayPatternInfo`](#model-voucherdisplaypatterninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode_info` | `VoucherSendModeInfo` | [`*VoucherSendModeInfo`](#model-vouchersendmodeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_rule` | `VoucherSendRule` | [`*PaymentVoucherSendRule`](#model-paymentvouchersendrule) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_type` | `VoucherType` | `*string` | JSON body | 未标注 | 券类型 |
| `voucher_use_rule` | `VoucherUseRule` | [`*PaymentVoucherUseRule`](#model-paymentvoucheruserule) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule_info` | `VoucherUseRuleInfo` | [`*VoucherUseRuleInfo`](#model-voucheruseruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-activityvouchermodify"></a>

### ActivityVoucherModify — 修改支付券基本信息

- 官方接口：[`alipay.marketing.activity.voucher.modify`](https://opendocs.alipay.com/apis/05v8k9)
- 路由：`POST /v3/alipay/marketing/activity/voucher/modify`
- 请求类型：[ActivityVoucherModifyRequest](voucher_request.go#L46)
- 响应类型：[ActivityVoucherModifyResponse](voucher_response.go#L21)

商户可以通过该接口修改支付券基本信息。 接口同步返回修改请求受理结果，后续通过alipay.marketing.activity.message.modified消息接口异步发送修改活动成功事件消息。
特别注意：修改支付券接口返回结果并不代表活动修改成功，必须等待消息回调后才能确保成功。 如果成功消息长时间未送达，请联系相关技术人员。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_base_info` | `ActivityBaseInfo` | [`*ActivityBaseInfo`](#model-activitybaseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 |
| `publish_end_time` | `PublishEndTime` | `*string` | JSON body | 未标注 | 券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss |
| `voucher_available_scope_info` | `VoucherAvailableScopeInfo` | [`*VoucherAvailableScopeInfo`](#model-voucheravailablescopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode_info` | `VoucherSendModeInfo` | [`*VoucherSendModeInfo`](#model-vouchersendmodeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule` | `VoucherUseRule` | [`*PaymentVoucherUseRuleModify`](#model-paymentvoucheruserulemodify) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule_info` | `VoucherUseRuleInfo` | [`*VoucherUseRuleInfo`](#model-voucheruseruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-activityvoucherpublish"></a>

### ActivityVoucherPublish — 激活支付券

- 官方接口：[`alipay.marketing.activity.voucher.publish`](https://opendocs.alipay.com/apis/05vsc3)
- 路由：`POST /v3/alipay/marketing/activity/voucher/publish`
- 请求类型：[ActivityVoucherPublishRequest](voucher_request.go#L63)
- 响应类型：[ActivityVoucherPublishResponse](voucher_response.go#L24)

支付券创建成功后，通过调用此接口激活支付券。 激活支付券接口同步返回请求受理结果，后续通过alipay.marketing.activity.message.created消息接口异步发送激活支付券成功事件消息。 如果成功消息长时间未送达，请联系相关技术人员。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 用作幂等控制 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |

<a id="api-activityvoucherquery"></a>

### ActivityVoucherQuery — 查询支付券详情

- 官方接口：[`alipay.marketing.activity.voucher.query`](https://opendocs.alipay.com/apis/05veld)
- 路由：`GET /v3/alipay/marketing/activity/voucher/query`
- 请求类型：[ActivityVoucherQueryRequest](voucher_request.go#L97)
- 响应类型：[ActivityVoucherQueryResponse](voucher_response.go#L30)

通过此接口可查询支付券活动信息，包括活动的配置信息以及活动概况数据。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | query | 否 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | query | 否 | 商户接入模式。 |

<a id="api-activityvouchersend"></a>

### ActivityVoucherSend — 发放支付券

- 官方接口：[`alipay.marketing.activity.voucher.send`](https://opendocs.alipay.com/apis/05vidi)
- 路由：`POST /v3/alipay/marketing/activity/voucher/send`
- 请求类型：[ActivityVoucherSendRequest](voucher_request.go#L73)
- 响应类型：[ActivityVoucherSendResponse](voucher_response.go#L62)

API完成制券后，可使用发放支付券接口发券。通过调用此接口可发放支付券给指定用户，发券场景可以是小程序、H5、APP等。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 领券的支付宝用户openId |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 领券的支付宝user_id账号 |

<a id="api-activityvoucherstop"></a>

### ActivityVoucherStop — 停止支付券

- 官方接口：[`alipay.marketing.activity.voucher.stop`](https://opendocs.alipay.com/apis/05vc1k)
- 路由：`POST /v3/alipay/marketing/activity/voucher/stop`
- 请求类型：[ActivityVoucherStopRequest](voucher_request.go#L87)
- 响应类型：[ActivityVoucherStopResponse](voucher_response.go#L68)

商家调用该接口停止支付券。 停止支付券接口同步返回请求受理结果，后续通过alipay.marketing.activity.message.stopped消息接口异步发送停止活动成功事件。 如果成功消息长时间未送达，请联系相关技术人员。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 64 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-activitybaseinfo"></a>

### ActivityBaseInfo

类型定义：[ActivityBaseInfo](voucher_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `activity_delivery_ext_info` | `ActivityDeliveryExtInfo` | [`*ActivityDeliveryExtInfo`](#model-activitydeliveryextinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `activity_id` | `ActivityId` | `*string` | 未标注 | 活动id |
| `activity_name` | `ActivityName` | `*string` | 未标注 | 活动名称。 |
| `activity_operation_status` | `ActivityOperationStatus` | `*string` | 未标注 | 活动操作状态。 |
| `activity_product_type` | `ActivityProductType` | `*string` | 未标注 | 券活动产品类型 |
| `activity_status` | `ActivityStatus` | `*string` | 未标注 | 活动状态 。 ACTIVE:活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。 PAUSE:活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。 |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*BelongMerchantInfo`](#model-belongmerchantinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `code_mode` | `CodeMode` | `*string` | 未标注 | 码模式。 MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放 MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。 |
| `out_activity_id` | `OutActivityId` | `*string` | 未标注 | 外部活动ID |

<a id="model-activitydeliveryextinfo"></a>

### ActivityDeliveryExtInfo

类型定义：[ActivityDeliveryExtInfo](voucher_model.go#L24)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auto_delivery` | `AutoDelivery` | `*bool` | 未标注 | 是否自动推广 |
| `channel_list` | `ChannelList` | `*[]string` | 未标注 | 推广渠道列表 |

<a id="model-belongmerchantinfo"></a>

### BelongMerchantInfo

类型定义：[BelongMerchantInfo](voucher_model.go#L42)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_type` | `BusinessType` | `*string` | 未标注 | 合作业务类型,默认为ISV_FOR_MERCHANT |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户PID,默认为当前接口调用商户。 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。 |

<a id="model-bizfundsettleinfo"></a>

### BizfundSettleInfo

类型定义：[BizfundSettleInfo](voucher_model.go#L52)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `settle_mode` | `SettleMode` | `*string` | 未标注 | 请款&#124;账期 模式 |
| `settle_target_type` | `SettleTargetType` | `*string` | 未标注 | 结算的目标对象类型 |

<a id="model-customerdefinedeductrule"></a>

### CustomerDefineDeductRule

类型定义：[CustomerDefineDeductRule](voucher_model.go#L60)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_deduct_rule_desc` | `CustomerDefineDeductRuleDesc` | `*string` | 未标注 | 自定义优惠规则描述 |

<a id="model-customerdefinevoucherinfo"></a>

### CustomerDefineVoucherInfo

类型定义：[CustomerDefineVoucherInfo](voucher_model.go#L66)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `deduct_info` | `DeductInfo` | [`*DeductInfo`](#model-deductinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-daterangeinfo"></a>

### DateRangeInfo

类型定义：[DateRangeInfo](voucher_model.go#L71)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `begin_date` | `BeginDate` | `*string` | 未标注 | 开始日期 格式：yyyy-MM-dd |
| `end_date` | `EndDate` | `*string` | 未标注 | 结束日期 格式：yyyy-MM-dd |

<a id="model-dateruleinfo"></a>

### DateRuleInfo

类型定义：[DateRuleInfo](voucher_model.go#L79)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `date_range_info` | `DateRangeInfo` | [`*DateRangeInfo`](#model-daterangeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deductinfo"></a>

### DeductInfo

类型定义：[DeductInfo](voucher_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_deduct_rule` | `CustomerDefineDeductRule` | [`*CustomerDefineDeductRule`](#model-customerdefinedeductrule) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-discountvoucherinfo"></a>

### DiscountVoucherInfo

类型定义：[DiscountVoucherInfo](voucher_model.go#L90)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ceiling_amount` | `CeilingAmount` | `*string` | 未标注 | 封顶金额。 |
| `discount` | `Discount` | `*string` | 未标注 | 折扣率。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-endtimeinfo"></a>

### EndTimeInfo

类型定义：[EndTimeInfo](voucher_model.go#L101)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `end_time` | `EndTime` | `*string` | 未标注 | 结束时间 格式：HH:mm:ss |
| `end_time_type` | `EndTimeType` | `*string` | 未标注 | 结束时间类型 |

<a id="model-exchangevoucherinfo"></a>

### ExchangeVoucherInfo

类型定义：[ExchangeVoucherInfo](voucher_model.go#L109)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 券的价值 |
| `biz_type` | `BizType` | `*string` | 未标注 | 兑换券业务类型。 注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。 |
| `exchange_goods_info` | `ExchangeGoodsInfo` | [`*VoucherExchangeGoodsInfo`](#model-voucherexchangegoodsinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 优惠门槛金额，表示只有当订单金额大于等于门槛金额时券才能使用。该字段为空时表示无门槛。 门槛金额的校验由服务商(商户)核销时自行校验，支付宝侧只做展示使用。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-fixvoucherinfo"></a>

### FixVoucherInfo

类型定义：[FixVoucherInfo](voucher_model.go#L121)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 面额，每张代金券可以抵扣的金额。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。该字段不填写，认为无门槛。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-holidayruleinfo"></a>

### HolidayRuleInfo

类型定义：[HolidayRuleInfo](voucher_model.go#L130)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-paymentfixvoucher"></a>

### PaymentFixVoucher

类型定义：[PaymentFixVoucher](voucher_model.go#L135)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 面额，每张优惠券可以抵扣的金额。 限制： 1、币种为人民币，单位为元。小数点以后最多保留两位。 2、取值范围:0.1&lt;=x&lt;=3000 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。该字段不填写，认为无门槛。 |

<a id="model-paymentvoucheralipaybalancerechargeinfo"></a>

### PaymentVoucherAlipayBalanceRechargeInfo

类型定义：[PaymentVoucherAlipayBalanceRechargeInfo](voucher_model.go#L143)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `logon_id` | `LogonId` | `*string` | 未标注 | 出资的支付宝登录账号 限制: 1、登录账号和用户ID必须且只能二选一 |
| `user_id` | `UserId` | `*string` | 未标注 | 出资的支付宝用户id 限制: 1、登录账号和用户ID必须且只能二选一 |

<a id="model-paymentvoucheravailablegoods"></a>

### PaymentVoucherAvailableGoods

类型定义：[PaymentVoucherAvailableGoods](voucher_model.go#L151)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_description` | `GoodsDescription` | `*string` | 未标注 | 商品描述信息。 用于券面展示，向用户介绍商品 |
| `goods_ids` | `GoodsIds` | `*[]string` | 未标注 | 可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。 |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称 |

<a id="model-paymentvoucheravailablegoodsmodify"></a>

### PaymentVoucherAvailableGoodsModify

类型定义：[PaymentVoucherAvailableGoodsModify](voucher_model.go#L161)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_ids` | `GoodsIds` | `*[]string` | 未标注 | 可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。 限制: 1、核销商品范围只能增大不能减小 |

<a id="model-paymentvoucheravailablemerchant"></a>

### PaymentVoucherAvailableMerchant

类型定义：[PaymentVoucherAvailableMerchant](voucher_model.go#L167)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_pids` | `AvailablePids` | `*[]string` | 未标注 | 优惠券可核销的直连商户PID 限制： 1、available_pids和available_smids至少二选一。 |
| `available_smids` | `AvailableSmids` | `*[]string` | 未标注 | 优惠券可核销的间连商户SMID 限制： 1、available_pids和available_smids至少二选一。 |

<a id="model-paymentvoucheravailablemerchantmodify"></a>

### PaymentVoucherAvailableMerchantModify

类型定义：[PaymentVoucherAvailableMerchantModify](voucher_model.go#L175)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_pids` | `AvailablePids` | `*[]string` | 未标注 | 优惠券可以核销的直连商户PID。默认不修改。 限制： 1、核销商户范围只能增大不能减小。 |
| `available_smids` | `AvailableSmids` | `*[]string` | 未标注 | 优惠券可以核销的间连商户SMID。默认不修改。 限制： 1、核销商户范围只能增大不能减小。 |

<a id="model-paymentvoucherbelongmerchantinfo"></a>

### PaymentVoucherBelongMerchantInfo

类型定义：[PaymentVoucherBelongMerchantInfo](voucher_model.go#L183)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户支付宝id，默认品牌名和品牌logo将从该商户信息中获取 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户支付宝id类型。 |

<a id="model-paymentvoucherbudgetinfo"></a>

### PaymentVoucherBudgetInfo

类型定义：[PaymentVoucherBudgetInfo](voucher_model.go#L191)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 营销资金预算,单位元。 总预算=优惠金额*总发券张数 |
| `budget_type` | `BudgetType` | `*string` | 未标注 | 营销资金预算类型 1、免充值：制券方无需提前充值资金，用户核销支付券时，直接从订单原价中扣除优惠减价金额，最终只将用户实际支付的金额结算给核销商户，商户实收少于订单原价。 2.预充值: 制券方需将优惠预算提前充值到支付宝指定营销账户中，用户核销支付券时，系统从该账户可用余额中扣除优惠减价部分对应的资金，连同用户实际支付的资金，一并结算给核销商户，不影响实收。 |
| `recharge_info` | `RechargeInfo` | [`*PaymentVoucherRechargeInfo`](#model-paymentvoucherrechargeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-paymentvoucherdisplayinfo"></a>

### PaymentVoucherDisplayInfo

类型定义：[PaymentVoucherDisplayInfo](voucher_model.go#L200)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_name` | `BrandName` | `*string` | 未标注 | 商户品牌名称。 |
| `customer_service_mobile` | `CustomerServiceMobile` | `*string` | 未标注 | 商家客服电话，会展示在支付宝卡包券详情页。 |
| `voucher_description` | `VoucherDescription` | `*string` | 未标注 | 用于说明详细的活动规则，会展示在支付宝卡包券详情页。 |
| `voucher_detail_images` | `VoucherDetailImages` | `*[]string` | 未标注 | 券详细图列表，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB 限制 1.voucher_image填入，该值才能填入； 2.最多3张； |
| `voucher_image` | `VoucherImage` | `*string` | 未标注 | 券详情页封面图，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值 限制: 该字段在兑换券场景下必传 上传图片尺寸670*335，支持格式：png、jpg、jpeg、bmp，大小不超过2MB 限制: 1.指定可核销商品(available_goods)时必传 |

<a id="model-paymentvoucherrechargeinfo"></a>

### PaymentVoucherRechargeInfo

类型定义：[PaymentVoucherRechargeInfo](voucher_model.go#L214)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `alipay_balance_recharge_info` | `AlipayBalanceRechargeInfo` | [`*PaymentVoucherAlipayBalanceRechargeInfo`](#model-paymentvoucheralipaybalancerechargeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `recharge_type` | `RechargeType` | `*string` | 未标注 | 预充值方式 |

<a id="model-paymentvouchersendrule"></a>

### PaymentVoucherSendRule

类型定义：[PaymentVoucherSendRule](voucher_model.go#L221)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `max_quantity_by_day` | `MaxQuantityByDay` | `*int64` | 未标注 | 设置此字段，允许指定单天最大发券数量。 限制: 每天发放张数*活动天数应小于等于优惠券发放总量 |
| `natural_person_limit` | `NaturalPersonLimit` | `*bool` | 未标注 | 限制相同身份证号领取次数(voucher_quantity_limit_per_user)。默认false不限制。 枚举值 true：是 false：否 |
| `phone_number_limit` | `PhoneNumberLimit` | `*bool` | 未标注 | 限制相同手机号领取次数(voucher_quantity_limit_per_user)。默认false不限制 枚举值 true：是 false：否 |
| `real_name_limit` | `RealNameLimit` | `*bool` | 未标注 | 限制支付宝实名用户才能领取支付券,默认为false表示不限制 枚举值 true\false |
| `voucher_quantity` | `VoucherQuantity` | `*int64` | 未标注 | 发行券的总数量。 限制： 1、发放总个数最少1个 2、发放总个数最多99999999个 |
| `voucher_quantity_limit_per_user` | `VoucherQuantityLimitPerUser` | `*int64` | 未标注 | 每人领取限制。 默认按照支付宝账号进行领取限制; 不填写或填入0表示没有领取限制. |
| `voucher_quantity_limit_per_user_period_type` | `VoucherQuantityLimitPerUserPeriodType` | `*string` | 未标注 | 周期限领配置,限制每人在固定周期内领取张数(voucher_quantity_limit_per_user),默认LIFE_CYCLE |

<a id="model-paymentvoucheruserule"></a>

### PaymentVoucherUseRule

类型定义：[PaymentVoucherUseRule](voucher_model.go#L239)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_app_ids` | `AvailableAppIds` | `*[]string` | 未标注 | 可核销的支付宝小程序id 限制： 1、必须是支付宝小程序 2、如果包含重复的小程序id会自动进行去重操作。 |
| `available_goods` | `AvailableGoods` | [`*PaymentVoucherAvailableGoods`](#model-paymentvoucheravailablegoods) | 未标注 | 见字段类型及官方接口文档。 |
| `available_merchant` | `AvailableMerchant` | [`*PaymentVoucherAvailableMerchant`](#model-paymentvoucheravailablemerchant) | 未标注 | 见字段类型及官方接口文档。 |
| `available_store_ids` | `AvailableStoreIds` | `*[]string` | 未标注 | 可核销支付门店id。 限制： 1、available_store_ids中的门店id必须是支付宝门店id。 2、available_store_ids如果包含重复的门店id会自动进行去重操作。 |
| `fix_voucher` | `FixVoucher` | [`*PaymentFixVoucher`](#model-paymentfixvoucher) | 未标注 | 见字段类型及官方接口文档。 |
| `unavailable_goods_ids` | `UnavailableGoodsIds` | `*[]string` | 未标注 | 不可核销商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，不可以使用优惠券。 限制: 1、available_goods与unavailable_goods_ids只能二选一或者同时为空 |
| `use_mode` | `UseMode` | `*string` | 未标注 | 核销跳转方式,默认CAN_USE(详情页展示可使用按钮) |
| `use_url` | `UseUrl` | `*string` | 未标注 | 核销跳转链接 限制: 1、只有use_mode为SELF_LINK时该值有效且必传。 2、核销跳转链接必须是支付宝小程序链接。 |
| `voucher_quantity_limit_per_user` | `VoucherQuantityLimitPerUser` | `*int64` | 未标注 | 每个支付宝账号的核销次数限制。 默认不限制 |
| `voucher_quantity_limit_per_user_period_type` | `VoucherQuantityLimitPerUserPeriodType` | `*string` | 未标注 | 核销次数限制周期配置。默认值LIFE_CYCLE表示限制每个支付宝账号在整个活动期间核销次数。 |
| `voucher_valid_period` | `VoucherValidPeriod` | [`*PaymentVoucherValidPeriod`](#model-paymentvouchervalidperiod) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-paymentvoucheruserulemodify"></a>

### PaymentVoucherUseRuleModify

类型定义：[PaymentVoucherUseRuleModify](voucher_model.go#L275)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_app_ids` | `AvailableAppIds` | `*string` | 未标注 | 可核销的支付宝小程序id 限制： 1、必须是支付宝小程序 2、如果包含重复的小程序id会自动进行去重操作。 3、可核销小程序范围只能扩大不能缩小 |
| `available_goods` | `AvailableGoods` | [`*PaymentVoucherAvailableGoodsModify`](#model-paymentvoucheravailablegoodsmodify) | 未标注 | 见字段类型及官方接口文档。 |
| `available_merchant` | `AvailableMerchant` | [`*PaymentVoucherAvailableMerchantModify`](#model-paymentvoucheravailablemerchantmodify) | 未标注 | 见字段类型及官方接口文档。 |
| `available_store_ids` | `AvailableStoreIds` | `*string` | 未标注 | 可核销支付门店id。 限制： 1、available_store_ids中的门店id必须是支付宝门店id。 2、available_store_ids如果包含重复的门店id会自动进行去重操作。 3、可核销门店范围只能扩大不能缩小 |
| `voucher_valid_period` | `VoucherValidPeriod` | [`*PaymentVoucherValidPeriodModify`](#model-paymentvouchervalidperiodmodify) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-paymentvouchervalidperiod"></a>

### PaymentVoucherValidPeriod

类型定义：[PaymentVoucherValidPeriod](voucher_model.go#L286)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `type` | `Type` | `*string` | 未标注 | 券有效期类型 |
| `valid_begin_time` | `ValidBeginTime` | `*string` | 未标注 | 券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss 限制： type为ABSOLUTE时该字段必填。 |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后N天内可以使用。 可以配合wait_days_after_receive字段使用。 限制： 1.type为RELATIVE时必填。 2.valid_days_after_receive必须大于0。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss 限制： 1、type为ABSOLUTE必填。 2、券可使用的结束时间valid_end_time必须大于券的发放结束时间publish_end_time |
| `wait_days_after_receive` | `WaitDaysAfterReceive` | `*int64` | 未标注 | 用户领券后需要等待N天，券才可以生效。默认用户领券后立刻生效。 限制： 1、type为RELATIVE时可选。 |

<a id="model-paymentvouchervalidperiodmodify"></a>

### PaymentVoucherValidPeriodModify

类型定义：[PaymentVoucherValidPeriodModify](voucher_model.go#L300)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后N天内可以使用。 限制： type为RELATIVE时可修改。 valid_days_after_receive必须大于0。 修改该时间，只允许延长，不允许缩短。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss 限制： type为ABSOLUTE可修改。 券可使用的结束时间valid_end_time 必须大于 券的发放结束时间 publish_end_time 修改券可使用的结束时间，只能延长，不允许缩短。 |

<a id="model-specialvoucherinfo"></a>

### SpecialVoucherInfo

类型定义：[SpecialVoucherInfo](voucher_model.go#L308)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `special_amount` | `SpecialAmount` | `*string` | 未标注 | 特价，即：原价-特价=优惠金额。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-timerangeinfo"></a>

### TimeRangeInfo

类型定义：[TimeRangeInfo](voucher_model.go#L317)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `begin_time` | `BeginTime` | `*string` | 未标注 | 开始时间 格式：HH:mm:ss |
| `end_time_info` | `EndTimeInfo` | [`*EndTimeInfo`](#model-endtimeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-timerestrictinfo"></a>

### TimeRestrictInfo

类型定义：[TimeRestrictInfo](voucher_model.go#L324)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `disable_period_info` | `DisablePeriodInfo` | [`*[]VoucherUseTimeRuleInfo`](#model-voucherusetimeruleinfo) | 未标注 | 不可用时间段 |
| `usable_period_info` | `UsablePeriodInfo` | [`*[]VoucherUseTimeRuleInfo`](#model-voucherusetimeruleinfo) | 未标注 | 可用时间段 |

<a id="model-voucherabsoluteperiodinfo"></a>

### VoucherAbsolutePeriodInfo

类型定义：[VoucherAbsolutePeriodInfo](voucher_model.go#L332)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_restrict_info` | `TimeRestrictInfo` | [`*TimeRestrictInfo`](#model-timerestrictinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `valid_begin_time` | `ValidBeginTime` | `*string` | 未标注 | 券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss。 |

<a id="model-voucheravailableaccountinfo"></a>

### VoucherAvailableAccountInfo

类型定义：[VoucherAvailableAccountInfo](voucher_model.go#L341)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_pids` | `AvailablePids` | `*[]string` | 未标注 | 优惠券可核销的直连商户PID |
| `available_smids` | `AvailableSmids` | `*[]string` | 未标注 | 优惠券可核销的间连商户SMID |

<a id="model-voucheravailableappinfo"></a>

### VoucherAvailableAppInfo

类型定义：[VoucherAvailableAppInfo](voucher_model.go#L349)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_app_ids` | `AvailableAppIds` | `*[]string` | 未标注 | 可核销的支付宝小程序id |

<a id="model-voucheravailablegeographyallshopinfo"></a>

### VoucherAvailableGeographyAllShopInfo

类型定义：[VoucherAvailableGeographyAllShopInfo](voucher_model.go#L355)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_brand_id` | `AvailableBrandId` | `*string` | 未标注 | 全部门店可用的品牌ID 提示: 1.目前只支持传入当前券归属商户的品牌ID 2.品牌ID可以在登录b.alipay.com-账户中心-品牌管理里查看。如果没有品牌，可以先创建品牌。Sop： 注意事项 1.只支持传入单个品牌ID 2.传入品牌ID后，券的商户品牌和logo仍以voucher_display_pattern_info传入为准。 |
| `exclude_shop_ids` | `ExcludeShopIds` | `*[]string` | 未标注 | 券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。 列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |
| `merchant_ids` | `MerchantIds` | `*[]string` | 未标注 | 全部门店可用的商户ID |

<a id="model-voucheravailablegeographycityinfo"></a>

### VoucherAvailableGeographyCityInfo

类型定义：[VoucherAvailableGeographyCityInfo](voucher_model.go#L365)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `all_city` | `AllCity` | `*bool` | 未标注 | 是否全国。 选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。 |
| `available_city_codes` | `AvailableCityCodes` | `*[]string` | 未标注 | 城市编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx)中内容填写。（[参考资料](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm)） |

<a id="model-voucheravailablegeographyscopeinfo"></a>

### VoucherAvailableGeographyScopeInfo

类型定义：[VoucherAvailableGeographyScopeInfo](voucher_model.go#L373)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_geography_city_info` | `AvailableGeographyCityInfo` | [`*VoucherAvailableGeographyCityInfo`](#model-voucheravailablegeographycityinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `available_geography_scope_type` | `AvailableGeographyScopeType` | `*string` | 未标注 | 券可用地理位置类型。 |
| `available_geography_shop_info` | `AvailableGeographyShopInfo` | [`*VoucherAvailableGeographyShopInfo`](#model-voucheravailablegeographyshopinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucheravailablegeographyshopinfo"></a>

### VoucherAvailableGeographyShopInfo

类型定义：[VoucherAvailableGeographyShopInfo](voucher_model.go#L381)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_geography_all_shop` | `AvailableGeographyAllShop` | [`*VoucherAvailableGeographyAllShopInfo`](#model-voucheravailablegeographyallshopinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `available_real_shop_ids` | `AvailableRealShopIds` | `*[]string` | 未标注 | 券可使用的物理门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。 |
| `available_shop_ids` | `AvailableShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |

<a id="model-voucheravailablegoodsinfo"></a>

### VoucherAvailableGoodsInfo

类型定义：[VoucherAvailableGoodsInfo](voucher_model.go#L390)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_goods_sku_ids` | `AvailableGoodsSkuIds` | `*[]string` | 未标注 | 可优惠商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，可以使用优惠券。 |
| `exclude_goods_sku_ids` | `ExcludeGoodsSkuIds` | `*[]string` | 未标注 | 不可核销商品编码，商家自定义商品编码。当用户支付时，交易中的商品编码和支付券配置的商品编码有任一匹配时，不可以使用优惠券。 |
| `goods_description` | `GoodsDescription` | `*string` | 未标注 | 商品描述信息。 用于券面展示，向用户介绍商品 |
| `goods_detail` | `GoodsDetail` | `*string` | 未标注 | 商品详情，会展示在特定渠道(如服务市场团购业务插件的套餐明细)。 |
| `goods_detail_images` | `GoodsDetailImages` | `*[]string` | 未标注 | 商品详情图片列表，会展示在特定渠道(如服务市场团购业务插件的头图) |
| `goods_detail_rich_description` | `GoodsDetailRichDescription` | `*string` | 未标注 | 商品详情富文本描述 |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称。 |
| `origin_amount` | `OriginAmount` | `*string` | 未标注 | 原价。说明：该字段可不填，填入商品名称goods_name则必填; |

<a id="model-voucheravailableiteminfo"></a>

### VoucherAvailableItemInfo

类型定义：[VoucherAvailableItemInfo](voucher_model.go#L410)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | 未标注 | 内部商品ID |
| `out_item_info` | `OutItemInfo` | [`*[]VoucherAvailableOutItemInfo`](#model-voucheravailableoutiteminfo) | 未标注 | 外部商品信息 |

<a id="model-voucheravailableoutiteminfo"></a>

### VoucherAvailableOutItemInfo

类型定义：[VoucherAvailableOutItemInfo](voucher_model.go#L418)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_app_id` | `ItemAppId` | `*string` | 未标注 | 小程序ID |
| `out_item_id` | `OutItemId` | `*string` | 未标注 | 外部商品ID |

<a id="model-voucheravailablescopeinfo"></a>

### VoucherAvailableScopeInfo

类型定义：[VoucherAvailableScopeInfo](voucher_model.go#L426)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_available_account_info` | `VoucherAvailableAccountInfo` | [`*VoucherAvailableAccountInfo`](#model-voucheravailableaccountinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_app_info` | `VoucherAvailableAppInfo` | [`*VoucherAvailableAppInfo`](#model-voucheravailableappinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_geography_scope_info` | `VoucherAvailableGeographyScopeInfo` | [`*VoucherAvailableGeographyScopeInfo`](#model-voucheravailablegeographyscopeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_goods_info` | `VoucherAvailableGoodsInfo` | [`*VoucherAvailableGoodsInfo`](#model-voucheravailablegoodsinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_item_info` | `VoucherAvailableItemInfo` | [`*VoucherAvailableItemInfo`](#model-voucheravailableiteminfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherbalancerechargeinfo"></a>

### VoucherBalanceRechargeInfo

类型定义：[VoucherBalanceRechargeInfo](voucher_model.go#L435)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 支付宝余额充值金额 限制： 1.币种为人民币，单位元。 2. 总预算=优惠金额*总发券张数 |
| `logon_id` | `LogonId` | `*string` | 未标注 | 出资的商户支付宝登录账号 |
| `partner_id` | `PartnerId` | `*string` | 未标注 | 出资的商户支付宝ID |

<a id="model-voucherbudgetsupplyinfo"></a>

### VoucherBudgetSupplyInfo

类型定义：[VoucherBudgetSupplyInfo](voucher_model.go#L445)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `budget_type` | `BudgetType` | `*string` | 未标注 | 营销资金预算类型 1、免充值：制券方无需提前充值资金，用户核销支付券时，直接从订单原价中扣除优惠减价金额，最终只将用户实际支付的金额结算给核销商户，商户实收少于订单原价。 2.预充值: 制券方需将优惠预算提前充值到支付宝指定营销账户中，用户核销支付券时，系统从该账户可用余额中扣除优惠减价部分对应的资金，连同用户实际支付的资金，一并结算给核销商户，不影响实收。 |
| `voucher_recharge_info` | `VoucherRechargeInfo` | [`*VoucherRechargeInfo`](#model-voucherrechargeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchercustomerguideinfo"></a>

### VoucherCustomerGuideInfo

类型定义：[VoucherCustomerGuideInfo](voucher_model.go#L452)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_send_guide_info` | `VoucherSendGuideInfo` | [`*VoucherSendGuideInfo`](#model-vouchersendguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_guide_info` | `VoucherUseGuideInfo` | [`*VoucherUseGuideInfo`](#model-voucheruseguideinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherdeductinfo"></a>

### VoucherDeductInfo

类型定义：[VoucherDeductInfo](voucher_voucherdeductinfo_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_voucher_info` | `CustomerDefineVoucherInfo` | [`*CustomerDefineVoucherInfo`](#model-customerdefinevoucherinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `discount_voucher_info` | `DiscountVoucherInfo` | [`*DiscountVoucherInfo`](#model-discountvoucherinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `exchange_voucher_info` | `ExchangeVoucherInfo` | [`*ExchangeVoucherInfo`](#model-exchangevoucherinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `fix_voucher_info` | `FixVoucherInfo` | [`*FixVoucherInfo`](#model-fixvoucherinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `special_voucher_info` | `SpecialVoucherInfo` | [`*SpecialVoucherInfo`](#model-specialvoucherinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_type` | `VoucherType` | `*string` | 未标注 | 券类型。 |

<a id="model-voucherdeductthresholdinfo"></a>

### VoucherDeductThresholdInfo

类型定义：[VoucherDeductThresholdInfo](voucher_voucherdeductinfo_model.go#L15)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_voucher_threshold_desc` | `CustomerDefineVoucherThresholdDesc` | `*string` | 未标注 | 自定义门槛描述 |
| `voucher_express_info` | `VoucherExpressInfo` | [`*VoucherExpressInfo`](#model-voucherexpressinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_goods_quantity_specification_info` | `VoucherGoodsQuantitySpecificationInfo` | [`*VoucherGoodsQuantitySpecificationInfo`](#model-vouchergoodsquantityspecificationinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherdisplaypatterninfo"></a>

### VoucherDisplayPatternInfo

类型定义：[VoucherDisplayPatternInfo](voucher_voucherdeductinfo_model.go#L23)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_logo` | `BrandLogo` | `*string` | 未标注 | 商家logo 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值。 |
| `brand_logo_url` | `BrandLogoUrl` | `*string` | 未标注 | 商家品牌logo链接 |
| `brand_name` | `BrandName` | `*string` | 未标注 | 商户品牌名称。 |
| `customer_service_mobile` | `CustomerServiceMobile` | `*string` | 未标注 | 客服电话。 |
| `customer_service_url` | `CustomerServiceUrl` | `*string` | 未标注 | 客服链接。 |
| `voucher_description` | `VoucherDescription` | `*string` | 未标注 | 券详细使用说明。 |
| `voucher_detail_image_urls` | `VoucherDetailImageUrls` | `*[]string` | 未标注 | 券详情链接 |
| `voucher_detail_images` | `VoucherDetailImages` | `*[]string` | 未标注 | 券详细图列表，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。 |
| `voucher_image` | `VoucherImage` | `*string` | 未标注 | 券详情页封面图，会展示在用户支付宝卡包券详情页。 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值。 |
| `voucher_image_url` | `VoucherImageUrl` | `*string` | 未标注 | 券封面链接 |
| `voucher_name` | `VoucherName` | `*string` | 未标注 | 券名称。 |

<a id="model-voucherexchangegoodsinfo"></a>

### VoucherExchangeGoodsInfo

类型定义：[VoucherExchangeGoodsInfo](voucher_voucherdeductinfo_model.go#L49)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `exchange_goods_name` | `ExchangeGoodsName` | `*string` | 未标注 | 兑换商品名称 |

<a id="model-voucherexpressinfo"></a>

### VoucherExpressInfo

类型定义：[VoucherExpressInfo](voucher_voucherdeductinfo_model.go#L55)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `pay_express` | `PayExpress` | `*string` | 未标注 | 邮费标识，属于券的服务，用于用户门槛信息展示 |

<a id="model-vouchergoodsquantityspecificationinfo"></a>

### VoucherGoodsQuantitySpecificationInfo

类型定义：[VoucherGoodsQuantitySpecificationInfo](voucher_voucherdeductinfo_model.go#L61)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `directional_type` | `DirectionalType` | `*string` | 未标注 | 定向类型用于规格设置，比如：FIX 表示 固定满X件，APPOINT 表示指定第N件 |
| `specification_quantity` | `SpecificationQuantity` | `*int64` | 未标注 | 数量 |
| `specification_unit` | `SpecificationUnit` | `*string` | 未标注 | 单位 |

<a id="model-voucherminiappsendguideinfo"></a>

### VoucherMiniAppSendGuideInfo

类型定义：[VoucherMiniAppSendGuideInfo](voucher_voucherdeductinfo_model.go#L79)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_url` | `MiniAppUrl` | `*string` | 未标注 | 小程序发券链接 |

<a id="model-voucherminiappuseguideinfo"></a>

### VoucherMiniAppUseGuideInfo

类型定义：[VoucherMiniAppUseGuideInfo](voucher_voucherdeductinfo_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_service_codes` | `MiniAppServiceCodes` | `*[]string` | 未标注 | 小程序服务编码。取值来源：通过[alipay.open.app.service.apply(小程序服务创建)](https://opendocs.alipay.com/mini/03cgoq?pathHash=49cde24b&scene=common&highlight_field=service_code)接口创建服务后获取。 |
| `mini_app_url` | `MiniAppUrl` | `*string` | 未标注 | 小程序核销链接 |

<a id="model-voucherpackagemodeinfo"></a>

### VoucherPackageModeInfo

类型定义：[VoucherPackageModeInfo](voucher_voucherdeductinfo_model.go#L93)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_package_id` | `VoucherPackageId` | `*string` | 未标注 | 券包id，对应alipay.marketing.activity.voucherpackage.query中voucher_package_id。 |

<a id="model-voucherrechargeinfo"></a>

### VoucherRechargeInfo

类型定义：[VoucherRechargeInfo](voucher_voucherdeductinfo_model.go#L99)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `recharge_type` | `RechargeType` | `*string` | 未标注 | 预充值方式。 |
| `voucher_balance_recharge_info` | `VoucherBalanceRechargeInfo` | [`*VoucherBalanceRechargeInfo`](#model-voucherbalancerechargeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherrelativeperiodinfo"></a>

### VoucherRelativePeriodInfo

类型定义：[VoucherRelativePeriodInfo](voucher_voucherdeductinfo_model.go#L106)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_restrict_info` | `TimeRestrictInfo` | [`*TimeRestrictInfo`](#model-timerestrictinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后 N 天内可以使用。 可以配合wait_days_after_receive 字段使用。 |
| `valid_days_relative_type` | `ValidDaysRelativeType` | `*string` | 未标注 | 优惠券有效期的取整方式，需与 valid_days_after_receive 字段配合使用，用于确定用户领取优惠券后的具体失效时间。 取值逻辑说明：ROUND_FORWARD（向前取整），优惠券在领取当日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 示例： -若 valid_days_relative_type=ROUND_FORWARD，valid_days_after_receive=1，用户于 1 号 09:12:40 领券，则优惠券在 1 号 23:59:59 失效。 -若 valid_days_after_receive=2，则在 2 号 23:59:59 失效。ROUND_BACKWARD（向后取整），优惠券从领取次日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 示例： -若valid_days_relative_type=ROUND_BACKWARDS，valid_days_after_receive=1，领券时间是1号9点12分40秒，那么 2 号 23:59:59 失效。 -若 valid_days_after_receive=2，则在 3 号 23:59:59 失效。 |
| `wait_days_after_receive` | `WaitDaysAfterReceive` | `*int64` | 未标注 | 用户领券后需要等待 N 天，券才可以生效。 字段值为 0 表示用户领券后立刻生效。 |

<a id="model-vouchersalemodeinfo"></a>

### VoucherSaleModeInfo

类型定义：[VoucherSaleModeInfo](voucher_voucherdeductinfo_model.go#L117)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `fund_custody_mode` | `FundCustodyMode` | `*string` | 未标注 | 资金托管方式 |
| `overdue_refundable` | `OverdueRefundable` | `*bool` | 未标注 | 是否支持优惠券过期后，自动退款给用户。 不填默认否，枚举值： true：是 false：否。自动退款功能需要服务商在优惠券过期时，主动调用alipay.marketing.activity.order.refund接口进行退款。 如果配置优惠券时选择了过期自动退款，但是实际券过期后，服务商没有进行退款，那么用户投诉后，需要服务商进行解决。 |
| `overdue_refundable_need_confirm` | `OverdueRefundableNeedConfirm` | `*bool` | 未标注 | 在支持过期自动退款后，支付宝发起自动退款前 是否需要向商户确认退款。 不填默认否，枚举值： true：是 false：否。 如果选择了 true:是 ，则需要商户配合接入 过期自动退款前是否需要向商户确认SPI(spi.alipay.marketing.activity.order.refundconfirm) 以及 退款确认结果API(alipay.marketing.activity.order.refundconfirmcommit) |
| `payee_pid` | `PayeePid` | `*string` | 未标注 | 收款账号。 目前的结算规则是，每核销一笔优惠券，支付宝会打款到该收款账户。 |
| `refundable` | `Refundable` | `*bool` | 未标注 | 购买的优惠券是否允许退款。 不填默认否，枚举值： true：是 false：否 |
| `refundable_need_confirm` | `RefundableNeedConfirm` | `*bool` | 未标注 | 在支持随时退后，支付宝发起退款前用于判断是否需要向商户确认退款。 不填默认否，枚举值： true：是 false：否。 如果选择了 true:是 ，则需要商户配合接入退款前是否需要向商户确认SPI(spi.alipay.marketing.activity.order.refundconfirm) |
| `sale_amount` | `SaleAmount` | `*string` | 未标注 | 用户购买优惠券需要支付的金额 |
| `settle_info` | `SettleInfo` | [`*BizfundSettleInfo`](#model-bizfundsettleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchersendguideinfo"></a>

### VoucherSendGuideInfo

类型定义：[VoucherSendGuideInfo](voucher_voucherdeductinfo_model.go#L136)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_send_guide_info` | `MiniAppSendGuideInfo` | [`*VoucherMiniAppSendGuideInfo`](#model-voucherminiappsendguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `send_guide_mode` | `SendGuideMode` | `*[]string` | 未标注 | 发放引导模式 |

<a id="model-vouchersendmodeinfo"></a>

### VoucherSendModeInfo

类型定义：[VoucherSendModeInfo](voucher_voucherdeductinfo_model.go#L143)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_channel_list` | `DeliveryChannelList` | `*[]string` | 未标注 | 推广渠道列表 |
| `voucher_package_mode_info` | `VoucherPackageModeInfo` | [`*VoucherPackageModeInfo`](#model-voucherpackagemodeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_sale_mode_info` | `VoucherSaleModeInfo` | [`*VoucherSaleModeInfo`](#model-vouchersalemodeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode` | `VoucherSendMode` | `*string` | 未标注 | 券发放模式。 |
| `voucher_send_rule_info` | `VoucherSendRuleInfo` | [`*VoucherSendRuleInfo`](#model-vouchersendruleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchersendruleinfo"></a>

### VoucherSendRuleInfo

类型定义：[VoucherSendRuleInfo](voucher_voucherdeductinfo_model.go#L154)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `max_quantity_by_day` | `MaxQuantityByDay` | `*int64` | 未标注 | 设置此字段，允许指定单天最大发券数量。 |
| `natural_person_limit` | `NaturalPersonLimit` | `*bool` | 未标注 | 是否开启自然人领取限制。 自然人表示按照身份证纬度进行领取限制。 |
| `phone_number_limit` | `PhoneNumberLimit` | `*bool` | 未标注 | 是否开启电话号码领取限制。 |
| `phone_number_need_input_limit` | `PhoneNumberNeedInputLimit` | `*bool` | 未标注 | 下单时是否需要用户填写手机号码 |
| `publish_end_time` | `PublishEndTime` | `*string` | 未标注 | 券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `publish_start_time` | `PublishStartTime` | `*string` | 未标注 | 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `quantity` | `Quantity` | `*int64` | 未标注 | 发行券的数量。 |
| `quantity_day_limit_per_user` | `QuantityDayLimitPerUser` | `*int64` | 未标注 | 每人每日可领取数量限制，填0或者不填表示不限制。 |
| `quantity_limit_per_user` | `QuantityLimitPerUser` | `*int64` | 未标注 | 每人领取限制。 默认按照支付宝账号进行领取限制; 不填写或填入0表示没有领取限制. |
| `quantity_limit_per_user_period_type` | `QuantityLimitPerUserPeriodType` | `*string` | 未标注 | 周期限领配置,限制每人在固定周期内领取张数(voucher_quantity_limit_per_user),默认LIFE_CYCLE |
| `quantity_month_limit_per_user` | `QuantityMonthLimitPerUser` | `*int64` | 未标注 | 每人每月可领取数量限制，填0或者不填表示不限制 |
| `quantity_week_limit_per_user` | `QuantityWeekLimitPerUser` | `*int64` | 未标注 | 每人每周可领取数量限制，填0或者不填表示不限制 |
| `real_name_limit` | `RealNameLimit` | `*bool` | 未标注 | 限制支付宝实名用户才能领取支付券,默认为false表示不限制 枚举值 true\false |

<a id="model-voucheruseguideinfo"></a>

### VoucherUseGuideInfo

类型定义：[VoucherUseGuideInfo](voucher_voucherdeductinfo_model.go#L192)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_use_guide_info` | `MiniAppUseGuideInfo` | [`*VoucherMiniAppUseGuideInfo`](#model-voucherminiappuseguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `use_guide_mode` | `UseGuideMode` | `*[]string` | 未标注 | 使用引导模式 |
| `use_url` | `UseUrl` | `*string` | 未标注 | 核销跳转链接 |

<a id="model-voucheruseruleinfo"></a>

### VoucherUseRuleInfo

类型定义：[VoucherUseRuleInfo](voucher_voucherdeductinfo_model.go#L201)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `payment_restriction_list` | `PaymentRestrictionList` | `*[]string` | 未标注 | 支付方式限制 |
| `quantity_limit_per_user` | `QuantityLimitPerUser` | `*int64` | 未标注 | 核销次数限制。 |
| `quantity_limit_per_user_period_type` | `QuantityLimitPerUserPeriodType` | `*string` | 未标注 | 周期限领配置,限制每人在固定周期内领取张数(quantity_limit_per_user) |
| `voucher_max_use_times` | `VoucherMaxUseTimes` | `*int64` | 未标注 | 券最大核销次数。 |
| `voucher_use_ext_info` | `VoucherUseExtInfo` | `*string` | 未标注 | 自定义扩展信息 |
| `voucher_use_time_info` | `VoucherUseTimeInfo` | [`*VoucherUseTimeInfo`](#model-voucherusetimeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherusetimeinfo"></a>

### VoucherUseTimeInfo

类型定义：[VoucherUseTimeInfo](voucher_voucherdeductinfo_model.go#L216)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `absolute_period_info` | `AbsolutePeriodInfo` | [`*VoucherAbsolutePeriodInfo`](#model-voucherabsoluteperiodinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `period_type` | `PeriodType` | `*string` | 未标注 | 券有效期。 |
| `relative_period_info` | `RelativePeriodInfo` | [`*VoucherRelativePeriodInfo`](#model-voucherrelativeperiodinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherusetimeruleinfo"></a>

### VoucherUseTimeRuleInfo

类型定义：[VoucherUseTimeRuleInfo](voucher_voucherdeductinfo_model.go#L224)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `date_rule_info` | `DateRuleInfo` | [`*DateRuleInfo`](#model-dateruleinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `holiday_rule_info` | `HolidayRuleInfo` | [`*HolidayRuleInfo`](#model-holidayruleinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `rule_type` | `RuleType` | `*string` | 未标注 | 规则类型 |
| `week_rule_info` | `WeekRuleInfo` | [`*WeekRuleInfo`](#model-weekruleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-weekruleinfo"></a>

### WeekRuleInfo

类型定义：[WeekRuleInfo](voucher_voucherdeductinfo_model.go#L233)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `week_day` | `WeekDay` | `*string` | 未标注 | 星期数，以“,”连接多个星期数 |
