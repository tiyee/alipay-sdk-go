# marketing/activity/ordervoucher 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/ordervoucher`。本目录提供 11 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `ordervoucher.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityOrderVoucherAppend`](#api-activityordervoucherappend) | [`alipay.marketing.activity.ordervoucher.append`](https://opendocs.alipay.com/apis/05vjnq) | 修改商家券活动发券数量上限 | `PATCH` |
| [`ActivityOrderVoucherAssociate`](#api-activityordervoucherassociate) | [`alipay.marketing.activity.ordervoucher.associate`](https://opendocs.alipay.com/apis/05v8k7) | 关联订单信息 | `POST` |
| [`ActivityOrderVoucherCodeCount`](#api-activityordervouchercodecount) | [`alipay.marketing.activity.ordervoucher.codecount`](https://opendocs.alipay.com/open-v3/1b98a84b_alipay.marketing.activity.ordervoucher.codecount) | 统计商家券券码数量 | `GET` |
| [`ActivityOrderVoucherCodeDeposit`](#api-activityordervouchercodedeposit) | [`alipay.marketing.activity.ordervoucher.codedeposit`](https://opendocs.alipay.com/apis/05u431) | 同步商家券券码 | `POST` |
| [`ActivityOrderVoucherCreate`](#api-activityordervouchercreate) | [`alipay.marketing.activity.ordervoucher.create`](https://opendocs.alipay.com/apis/05qgfh) | 创建商家券活动 | `POST` |
| [`ActivityOrderVoucherDisassociate`](#api-activityordervoucherdisassociate) | [`alipay.marketing.activity.ordervoucher.disassociate`](https://opendocs.alipay.com/apis/05vluj) | 取消关联订单信息 | `POST` |
| [`ActivityOrderVoucherModify`](#api-activityordervouchermodify) | [`alipay.marketing.activity.ordervoucher.modify`](https://opendocs.alipay.com/apis/05vfpx) | 修改商家券活动基本信息 | `PATCH` |
| [`ActivityOrderVoucherQuery`](#api-activityordervoucherquery) | [`alipay.marketing.activity.ordervoucher.query`](https://opendocs.alipay.com/open-v3/ab283f3f_alipay.marketing.activity.ordervoucher.query) | 查询商家券活动 | `GET` |
| [`ActivityOrderVoucherRefund`](#api-activityordervoucherrefund) | [`alipay.marketing.activity.ordervoucher.refund`](https://opendocs.alipay.com/apis/0636fj) | 取消券核销状态 | `POST` |
| [`ActivityOrderVoucherStop`](#api-activityordervoucherstop) | [`alipay.marketing.activity.ordervoucher.stop`](https://opendocs.alipay.com/apis/0630yz) | 停止商家券活动 | `PATCH` |
| [`ActivityOrderVoucherUse`](#api-activityordervoucheruse) | [`alipay.marketing.activity.ordervoucher.use`](https://opendocs.alipay.com/apis/0639qt) | 同步券核销状态 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activityordervoucherappend"></a>

### ActivityOrderVoucherAppend — 修改商家券活动发券数量上限

- 官方接口：[`alipay.marketing.activity.ordervoucher.append`](https://opendocs.alipay.com/apis/05vjnq)
- 路由：`PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}/append`
- 请求类型：[ActivityOrderVoucherAppendRequest](ordervoucher_request.go#L65)
- 响应类型：[ActivityOrderVoucherAppendResponse](ordervoucher_response.go#L9)

商家调用该接口修改商家券活动预算上限，以此来追加商家券活动预算。
修改商家券活动预算上限同步返回追加请求受理结果，后续通过alipay.marketing.activity.message.appended消息接口异步发送发券数量上限修改成功事件消息。
code_mode=MERCHANT_UPLOAD场景下，必须先导入券码后，方可追加商家券预算。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_access_mode` | `MarketingActivityOrderVoucherAppend.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherAppend.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `voucher_quantity` | `MarketingActivityOrderVoucherAppend.VoucherQuantity` | `*int64` | JSON body | 未标注 | 发行券的数量。 限制： 券总预算&lt;= 999999 特别说明： 该字段的含义是追加到的数量。 不可减少，只能增加。 code_mode= MERCHANT_UPLOAD 的情况下，voucher_quantity 需要与导入码的总数量保持一致（即历史导入数量 + 追加导入数量）。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherAppend](ordervoucher_model.go#L32) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activityordervoucherassociate"></a>

### ActivityOrderVoucherAssociate — 关联订单信息

- 官方接口：[`alipay.marketing.activity.ordervoucher.associate`](https://opendocs.alipay.com/apis/05v8k7)
- 路由：`POST /v3/alipay/marketing/activity/ordervoucher/associate`
- 请求类型：[ActivityOrderVoucherAssociateRequest](ordervoucher_request.go#L5)
- 响应类型：[ActivityOrderVoucherAssociateResponse](ordervoucher_response.go#L12)

将已发放的商家券券码与订单信息关联，用于后续参与返佣激励等操作的统计。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝支付下单时的商户订单号，欲与该商家券关联的支付宝支付 |
| `voucher_code` | `VoucherCode` | `*string` | JSON body | 未标注 | 商家券券码。 限制: 券码必须已发放给用户。 |

<a id="api-activityordervouchercodecount"></a>

### ActivityOrderVoucherCodeCount — 统计商家券券码数量

- 官方接口：[`alipay.marketing.activity.ordervoucher.codecount`](https://opendocs.alipay.com/open-v3/1b98a84b_alipay.marketing.activity.ordervoucher.codecount)
- 路由：`GET /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codecount`
- 请求类型：[ActivityOrderVoucherCodeCountRequest](ordervoucher_request.go#L72)
- 响应类型：[ActivityOrderVoucherCodeCountResponse](ordervoucher_response.go#L18)

商家调用该接口查询活动中的有效券码个数。接口同步返回查询结果。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | query | 否 | 商户接入模式 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。

<a id="api-activityordervouchercodedeposit"></a>

### ActivityOrderVoucherCodeDeposit — 同步商家券券码

- 官方接口：[`alipay.marketing.activity.ordervoucher.codedeposit`](https://opendocs.alipay.com/apis/05u431)
- 路由：`POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codedeposit`
- 请求类型：[ActivityOrderVoucherCodeDepositRequest](ordervoucher_request.go#L80)
- 响应类型：[ActivityOrderVoucherCodeDepositResponse](ordervoucher_response.go#L24)

商家调用该接口同步导入商家券券码，接口同步返回导入结果。请在接收alipay.marketing.activity.message.created消息后，再进行调用。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_access_mode` | `MarketingActivityOrderVoucherCodeDeposit.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherCodeDeposit.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `voucher_codes` | `MarketingActivityOrderVoucherCodeDeposit.VoucherCodes` | `*[]string` | JSON body | 未标注 | 券码的数量列表。接口参数为列表类型。 限制： 目前最大上传 1000 个。 单个code最长64位。 商户上传的券code列表，code允许包含的字符有0-9、a-z、A-Z、-、_、+、=、&#124;。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherCodeDeposit](ordervoucher_model.go#L42) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activityordervouchercreate"></a>

### ActivityOrderVoucherCreate — 创建商家券活动

- 官方接口：[`alipay.marketing.activity.ordervoucher.create`](https://opendocs.alipay.com/apis/05qgfh)
- 路由：`POST /v3/alipay/marketing/ordervoucher/activity`
- 请求类型：[ActivityOrderVoucherCreateRequest](ordervoucher_request.go#L19)
- 响应类型：[ActivityOrderVoucherCreateResponse](ordervoucher_response.go#L36)

商家调用该接口创建商家券活动。
创建商家券活动接口同步返回创建请求受理结果，后续通过alipay.marketing.activity.message.created消息接口异步发送活动创建成功事件消息。
如果在code_mode=MERCHANT_UPLOAD的模式下，创建商家券活动后，还需要进行券码导入操作，以及预算追加操作，然后商家券才能进行发放。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_base_info` | `ActivityBaseInfo` | [`*ActivityBaseInfo`](#model-activitybaseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_name` | `ActivityName` | `*string` | JSON body | 未标注 | 活动名称。不对用户进行展示，仅供商家在后台管理活动使用。 |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*BelongMerchantInfo`](#model-belongmerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `biz_tag` | `BizTag` | `*string` | JSON body | 未标注 | 该字段废弃！后续不要使用该字段！商家券业务标签，影响商家券对C端用户的展示形式。 |
| `code_mode` | `CodeMode` | `*string` | JSON body | 未标注 | 码模式。MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放。MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。 |
| `customer_guide` | `CustomerGuide` | [`*CustomerGuide`](#model-customerguide) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用：参数不变的情况下，再次请求返回与上一次相同的结果。 |
| `publish_end_time` | `PublishEndTime` | `*string` | JSON body | 未标注 | 券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss |
| `publish_start_time` | `PublishStartTime` | `*string` | JSON body | 未标注 | 券发放开始时间。格式为：yyyy-MM-dd HH:mm:ss |
| `voucher_available_scope_info` | `VoucherAvailableScopeInfo` | [`*VoucherAvailableScopeInfo`](#model-voucheravailablescopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_customer_guide_info` | `VoucherCustomerGuideInfo` | [`*VoucherCustomerGuideInfo`](#model-vouchercustomerguideinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_deduct_info` | `VoucherDeductInfo` | [`*VoucherDeductInfo`](#model-voucherdeductinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_info` | `VoucherDisplayInfo` | [`*VoucherDisplayInfo`](#model-voucherdisplayinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_pattern_info` | `VoucherDisplayPatternInfo` | [`*VoucherDisplayPatternInfo`](#model-voucherdisplaypatterninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode_info` | `VoucherSendModeInfo` | [`*VoucherSendModeInfo`](#model-vouchersendmodeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_rule` | `VoucherSendRule` | [`*VoucherSendRuleDetail`](#model-vouchersendruledetail) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_type` | `VoucherType` | `*string` | JSON body | 未标注 | 券类型。 兑换券仅支持一类券类型：EXCHANGE_VOUCHER: 兑换券； |
| `voucher_use_rule` | `VoucherUseRule` | [`*VoucherUseRule`](#model-voucheruserule) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule_info` | `VoucherUseRuleInfo` | [`*VoucherUseRuleInfo`](#model-voucheruseruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-activityordervoucherdisassociate"></a>

### ActivityOrderVoucherDisassociate — 取消关联订单信息

- 官方接口：[`alipay.marketing.activity.ordervoucher.disassociate`](https://opendocs.alipay.com/apis/05vluj)
- 路由：`POST /v3/alipay/marketing/activity/ordervoucher/disassociate`
- 请求类型：[ActivityOrderVoucherDisassociateRequest](ordervoucher_request.go#L51)
- 响应类型：[ActivityOrderVoucherDisassociateResponse](ordervoucher_response.go#L44)

取消商家券与订单信息的关联关系.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 该交易在支付宝系统中的交易流水号。 |
| `voucher_code` | `VoucherCode` | `*string` | JSON body | 未标注 | 券码。 限制: 券码必须已发放给用户。 |

<a id="api-activityordervouchermodify"></a>

### ActivityOrderVoucherModify — 修改商家券活动基本信息

- 官方接口：[`alipay.marketing.activity.ordervoucher.modify`](https://opendocs.alipay.com/apis/05vfpx)
- 路由：`PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}`
- 请求类型：[ActivityOrderVoucherModifyRequest](ordervoucher_request.go#L87)
- 响应类型：[ActivityOrderVoucherModifyResponse](ordervoucher_response.go#L50)

商家调用该接口修改商家券活动基本信息。
修改商家券活动基本信息接口同步返回修改请求受理结果，后续通过alipay.marketing.activity.message.modified消息接口异步发送修改活动成功事件消息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_base_info` | `MarketingActivityOrderVoucherModify.ActivityBaseInfo` | [`*ActivityBaseInfo`](#model-activitybaseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_name` | `MarketingActivityOrderVoucherModify.ActivityName` | `*string` | JSON body | 未标注 | 活动名称。 |
| `customer_guide` | `MarketingActivityOrderVoucherModify.CustomerGuide` | [`*CustomerGuideModify`](#model-customerguidemodify) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_access_mode` | `MarketingActivityOrderVoucherModify.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherModify.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 |
| `publish_end_time` | `MarketingActivityOrderVoucherModify.PublishEndTime` | `*string` | JSON body | 未标注 | 券发放结束时间。该值为空表示不修改 格式为：yyyy-MM-dd HH:mm:ss |
| `publish_start_time` | `MarketingActivityOrderVoucherModify.PublishStartTime` | `*string` | JSON body | 未标注 | 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `voucher_available_scope_info` | `MarketingActivityOrderVoucherModify.VoucherAvailableScopeInfo` | [`*VoucherAvailableScopeInfo`](#model-voucheravailablescopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_scope_modify_type` | `MarketingActivityOrderVoucherModify.VoucherAvailableScopeModifyType` | `*string` | JSON body | 未标注 | 可用范围修改类型。 |
| `voucher_customer_guide_info` | `MarketingActivityOrderVoucherModify.VoucherCustomerGuideInfo` | [`*VoucherCustomerGuideInfo`](#model-vouchercustomerguideinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_info` | `MarketingActivityOrderVoucherModify.VoucherDisplayInfo` | [`*VoucherDisplayInfoModify`](#model-voucherdisplayinfomodify) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_display_pattern_info` | `MarketingActivityOrderVoucherModify.VoucherDisplayPatternInfo` | [`*VoucherDisplayPatternInfo`](#model-voucherdisplaypatterninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode_info` | `MarketingActivityOrderVoucherModify.VoucherSendModeInfo` | [`*VoucherSendModeInfo`](#model-vouchersendmodeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_rule` | `MarketingActivityOrderVoucherModify.VoucherSendRule` | [`*VoucherSendRuleDetailModify`](#model-vouchersendruledetailmodify) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule` | `MarketingActivityOrderVoucherModify.VoucherUseRule` | [`*VoucherUseRuleModify`](#model-voucheruserulemodify) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_rule_info` | `MarketingActivityOrderVoucherModify.VoucherUseRuleInfo` | [`*VoucherUseRuleInfo`](#model-voucheruseruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id1 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherModify](ordervoucher_model.go#L52) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activityordervoucherquery"></a>

### ActivityOrderVoucherQuery — 查询商家券活动

- 官方接口：[`alipay.marketing.activity.ordervoucher.query`](https://opendocs.alipay.com/open-v3/ab283f3f_alipay.marketing.activity.ordervoucher.query)
- 路由：`GET /v3/alipay/marketing/activity/ordervoucher/query`
- 请求类型：[ActivityOrderVoucherQueryRequest](ordervoucher_request.go#L94)
- 响应类型：[ActivityOrderVoucherQueryResponse](ordervoucher_response.go#L56)

商户通过该接口查询商家券活动详情。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | query | 否 | 活动id |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | query | 否 | 商户接入模式 |

<a id="api-activityordervoucherrefund"></a>

### ActivityOrderVoucherRefund — 取消券核销状态

- 官方接口：[`alipay.marketing.activity.ordervoucher.refund`](https://opendocs.alipay.com/apis/0636fj)
- 路由：`POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/refund`
- 请求类型：[ActivityOrderVoucherRefundRequest](ordervoucher_request.go#L102)
- 响应类型：[ActivityOrderVoucherRefundResponse](ordervoucher_response.go#L91)

商户调用该接口进行优惠券退券处理。接口同步返回结果。
如果优惠券过期后，发起退券操作，该接口返回成功，优惠券退回后仍不可用。
如果优惠券在支付宝卡包已删除，发起退券操作，该接口返回成功，优惠券在支付宝卡包仍然是删除状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_dt` | `MarketingActivityOrderVoucherRefund.BizDt` | `*string` | JSON body | 未标注 | 当前订单的退券时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `merchant_access_mode` | `MarketingActivityOrderVoucherRefund.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherRefund.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一 |
| `total_fee` | `MarketingActivityOrderVoucherRefund.TotalFee` | `*string` | JSON body | 未标注 | 当前订单的退款总金额(包含优惠券金额)，单位为元。 |
| `voucher_refund_detail_info` | `MarketingActivityOrderVoucherRefund.VoucherRefundDetailInfo` | [`*VoucherRefundDetailInfo`](#model-voucherrefunddetailinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |
| `voucher_code` | `VoucherCode` | `string` | 路径 | 是 | 已核销待退券的券码 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherRefund](ordervoucher_model.go#L78) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activityordervoucherstop"></a>

### ActivityOrderVoucherStop — 停止商家券活动

- 官方接口：[`alipay.marketing.activity.ordervoucher.stop`](https://opendocs.alipay.com/apis/0630yz)
- 路由：`PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}/stop`
- 请求类型：[ActivityOrderVoucherStopRequest](ordervoucher_request.go#L111)
- 响应类型：[ActivityOrderVoucherStopResponse](ordervoucher_response.go#L98)

商家调用该接口停止优惠券。
停止优惠券接口同步返回停止请求受理结果，后续通过alipay.marketing.activity.message.stopped消息接口异步发送停止活动成功事件。
如果成功消息长时间未送达，请联系相关技术人员。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_access_mode` | `MarketingActivityOrderVoucherStop.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherStop.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherStop](ordervoucher_model.go#L91) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

<a id="api-activityordervoucheruse"></a>

### ActivityOrderVoucherUse — 同步券核销状态

- 官方接口：[`alipay.marketing.activity.ordervoucher.use`](https://opendocs.alipay.com/apis/0639qt)
- 路由：`POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/use`
- 请求类型：[ActivityOrderVoucherUseRequest](ordervoucher_request.go#L118)
- 响应类型：[ActivityOrderVoucherUseResponse](ordervoucher_response.go#L101)

商户调用该接口，通过券码进行优惠券核销，接口同步返回核销结果。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_dt` | `MarketingActivityOrderVoucherUse.BizDt` | `*string` | JSON body | 未标注 | 外部优惠券的核销时间，用于流水记录。 格式为：yyyy-MM-dd HH:mm:ss |
| `goods_detail` | `MarketingActivityOrderVoucherUse.GoodsDetail` | [`*[]GoodsDetailDTO`](#model-goodsdetaildto) | JSON body | 未标注 | 商品明细信息 |
| `merchant_access_mode` | `MarketingActivityOrderVoucherUse.MerchantAccessMode` | `*string` | JSON body | 未标注 | 商户接入模式 |
| `out_biz_no` | `MarketingActivityOrderVoucherUse.OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。 |
| `real_shop_id` | `MarketingActivityOrderVoucherUse.RealShopId` | `*string` | JSON body | 未标注 | 用户核销优惠券的代运营商业关系门店id，调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。 限制: 传入的门店id必须是创建商家券voucher_use_rule. voucher_available_scope. order_voucher_available_shop对象中real_shop_ids中的门店id |
| `store_id` | `MarketingActivityOrderVoucherUse.StoreId` | `*string` | JSON body | 未标注 | 用户核销优惠券的支付门店id,调用接口[ant.merchant.expand.shop.create](https://opendocs.alipay.com/open/03s3jk?pathHash=04312530&scene=common)创建门店返回的门店id。 限制: 传入的门店id必须是创建商家券voucher_use_rule. voucher_available_scope. order_voucher_available_shop对象中shop_ids中的门店id |
| `total_fee` | `MarketingActivityOrderVoucherUse.TotalFee` | `*string` | JSON body | 未标注 | 订单总金额 |
| `trade_channel` | `MarketingActivityOrderVoucherUse.TradeChannel` | `*string` | JSON body | 未标注 | 交易渠道。 |
| `trade_no` | `MarketingActivityOrderVoucherUse.TradeNo` | `*string` | JSON body | 未标注 | 该交易在支付宝系统中的交易流水号。 注意：若用户使用支付宝为交易的渠道，本参数必填。 |
| `voucher_use_detail_info` | `MarketingActivityOrderVoucherUse.VoucherUseDetailInfo` | [`*VoucherUseDetailInfo`](#model-voucherusedetailinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `activity_id` | `ActivityId` | `string` | 路径 | 是 | 活动id |
| `voucher_code` | `VoucherCode` | `string` | 路径 | 是 | 待核销的券码。 |

注意事项：

- 带 path 标签的字段替换路由占位符，由 SDK 转义；其余字段仍按表中位置编码。
- 请求通过内嵌类型 [MarketingActivityOrderVoucherUse](ordervoucher_model.go#L99) 展开 JSON body；表中带类型前缀的 Go 字段在该内嵌对象中赋值，编码时不增加外层对象。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-activitybaseinfo"></a>

### ActivityBaseInfo

类型定义：[ActivityBaseInfo](ordervoucher_model.go#L4)。

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

类型定义：[ActivityDeliveryExtInfo](ordervoucher_model.go#L24)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auto_delivery` | `AutoDelivery` | `*bool` | 未标注 | 是否自动推广 |
| `channel_list` | `ChannelList` | `*[]string` | 未标注 | 推广渠道列表 |

<a id="model-belongmerchantinfo"></a>

### BelongMerchantInfo

类型定义：[BelongMerchantInfo](ordervoucher_model.go#L132)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_type` | `BusinessType` | `*string` | 未标注 | 合作业务类型,默认为ISV_FOR_MERCHANT |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户PID,默认为当前接口调用商户。 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。 |

<a id="model-bizfundsettleinfo"></a>

### BizfundSettleInfo

类型定义：[BizfundSettleInfo](ordervoucher_model.go#L142)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `settle_mode` | `SettleMode` | `*string` | 未标注 | 请款&#124;账期 模式 |
| `settle_target_type` | `SettleTargetType` | `*string` | 未标注 | 结算的目标对象类型 |

<a id="model-customerdefinedeductrule"></a>

### CustomerDefineDeductRule

类型定义：[CustomerDefineDeductRule](ordervoucher_model.go#L150)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_deduct_rule_desc` | `CustomerDefineDeductRuleDesc` | `*string` | 未标注 | 自定义优惠规则描述 |

<a id="model-customerdefinevoucherinfo"></a>

### CustomerDefineVoucherInfo

类型定义：[CustomerDefineVoucherInfo](ordervoucher_model.go#L156)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `deduct_info` | `DeductInfo` | [`*DeductInfo`](#model-deductinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-customerguide"></a>

### CustomerGuide

类型定义：[CustomerGuide](ordervoucher_model.go#L161)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 券可用的小程序appId，卡包详情页可跳转到该appId |
| `mini_app_path` | `MiniAppPath` | `*string` | 未标注 | 指定跳转到mini_app_id时的具体页面路径。 限制： 1、只有mini_app_id有值时该值传入才会有效 2、该小程序路径是相对路径。详情参见 [小程序scheme链接介绍](https://opendocs.alipay.com/support/01rb18) |
| `real_shop_ids` | `RealShopIds` | `*[]string` | 未标注 | 代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。接口参数是列表类型。 |
| `service_codes` | `ServiceCodes` | `*[]string` | 未标注 | 小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 |
| `shop_ids` | `ShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |
| `store_ids` | `StoreIds` | `*[]string` | 未标注 | 该字段后续废弃。券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id。接口参数是列表类型。 |
| `voucher_send_guide` | `VoucherSendGuide` | [`*VoucherSendGuide`](#model-vouchersendguide) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_guide` | `VoucherUseGuide` | [`*VoucherUseGuide`](#model-voucheruseguide) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-customerguidemodify"></a>

### CustomerGuideModify

类型定义：[CustomerGuideModify](ordervoucher_model.go#L179)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 券可用的小程序 appId，卡包详情页可跳转到该 appId。 |
| `mini_app_path` | `MiniAppPath` | `*string` | 未标注 | 指定跳转到 mini_app_id 时的具体页面路径。 |
| `real_shop_ids` | `RealShopIds` | `*[]string` | 未标注 | 券可使用的商业关系门店列表,列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。 |
| `service_codes` | `ServiceCodes` | `*[]string` | 未标注 | 小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 |
| `shop_ids` | `ShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店 id 接口参数是列表类型。 接口参数是列表类型。 |
| `store_ids` | `StoreIds` | `*[]string` | 未标注 | 后续废弃。券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店 id 接口参数是列表类型。 接口参数是列表类型。 |
| `voucher_send_guide` | `VoucherSendGuide` | [`*VoucherSendGuideModify`](#model-vouchersendguidemodify) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_guide` | `VoucherUseGuide` | [`*VoucherUseGuideModify`](#model-voucheruseguidemodify) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-daterangeinfo"></a>

### DateRangeInfo

类型定义：[DateRangeInfo](ordervoucher_model.go#L197)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `begin_date` | `BeginDate` | `*string` | 未标注 | 开始日期 格式：yyyy-MM-dd |
| `end_date` | `EndDate` | `*string` | 未标注 | 结束日期 格式：yyyy-MM-dd |

<a id="model-dateruleinfo"></a>

### DateRuleInfo

类型定义：[DateRuleInfo](ordervoucher_model.go#L205)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `date_range_info` | `DateRangeInfo` | [`*DateRangeInfo`](#model-daterangeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-deductinfo"></a>

### DeductInfo

类型定义：[DeductInfo](ordervoucher_model.go#L211)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_deduct_rule` | `CustomerDefineDeductRule` | [`*CustomerDefineDeductRule`](#model-customerdefinedeductrule) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-discountvoucher"></a>

### DiscountVoucher

类型定义：[DiscountVoucher](ordervoucher_model.go#L216)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ceiling_amount` | `CeilingAmount` | `*string` | 未标注 | 封顶金额。 |
| `discount` | `Discount` | `*string` | 未标注 | 折扣率。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称。 |
| `origin_amount` | `OriginAmount` | `*string` | 未标注 | 原价。说明：该字段可不填，填入商品名称goods_name则必填; |

<a id="model-discountvoucherinfo"></a>

### DiscountVoucherInfo

类型定义：[DiscountVoucherInfo](ordervoucher_model.go#L230)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ceiling_amount` | `CeilingAmount` | `*string` | 未标注 | 封顶金额。 |
| `discount` | `Discount` | `*string` | 未标注 | 折扣率。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-endtimeinfo"></a>

### EndTimeInfo

类型定义：[EndTimeInfo](ordervoucher_model.go#L241)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `end_time` | `EndTime` | `*string` | 未标注 | 结束时间 格式：HH:mm:ss |
| `end_time_type` | `EndTimeType` | `*string` | 未标注 | 结束时间类型 |

<a id="model-exchangevoucher"></a>

### ExchangeVoucher

类型定义：[ExchangeVoucher](ordervoucher_model.go#L249)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 券的价值 |
| `biz_type` | `BizType` | `*string` | 未标注 | 兑换券业务类型。 枚举值 团购券：GROUP_BUY_EXCHANGE_VOUCHER 代金券：FIX_EXCHANGE_VOUCHER 注意：兑换券通过[大促活动权益报名](https://opendocs.alipay.com/pre-open/repo-00tbta)能力推广至支付宝会场时，本参数必填。 |
| `customer_service_mobile` | `CustomerServiceMobile` | `*string` | 未标注 | 客服电话 |
| `customer_service_url` | `CustomerServiceUrl` | `*string` | 未标注 | 客服链接 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `overdue_refundable` | `OverdueRefundable` | `*bool` | 未标注 | 是否支持优惠券过期后，自动退款给用户。 不填默认否，枚举值： true：是 false：否 |
| `payee_pid` | `PayeePid` | `*string` | 未标注 | 收款账号。目前的结算规则是，每核销一笔优惠券，支付宝会打款到该收款账户。 |
| `refundable` | `Refundable` | `*bool` | 未标注 | 购买的优惠券是否允许退款。不填默认否，枚举值： true：是 false：否 |
| `sale_amount` | `SaleAmount` | `*string` | 未标注 | 用户购买优惠券需要支付的金额 |
| `settle_type` | `SettleType` | `*string` | 未标注 | 解决类型。 |
| `voucher_detail_url` | `VoucherDetailUrl` | `*string` | 未标注 | 领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。 |
| `voucher_name` | `VoucherName` | `*string` | 未标注 | 对消费者展示的券(商品)名称。 |

<a id="model-exchangevoucherinfo"></a>

### ExchangeVoucherInfo

类型定义：[ExchangeVoucherInfo](ordervoucher_model.go#L277)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 券的价值 |
| `biz_type` | `BizType` | `*string` | 未标注 | 兑换券业务类型。 注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。 |
| `exchange_goods_info` | `ExchangeGoodsInfo` | [`*VoucherExchangeGoodsInfo`](#model-voucherexchangegoodsinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 优惠门槛金额，表示只有当订单金额大于等于门槛金额时券才能使用。该字段为空时表示无门槛。 门槛金额的校验由服务商(商户)核销时自行校验，支付宝侧只做展示使用。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-exchangevouchermodify"></a>

### ExchangeVoucherModify

类型定义：[ExchangeVoucherModify](ordervoucher_model.go#L289)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_detail_url` | `VoucherDetailUrl` | `*string` | 未标注 | 领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。 |
| `voucher_name` | `VoucherName` | `*string` | 未标注 | 对消费者展示的券(商品)名称。 |

<a id="model-fixvoucher"></a>

### FixVoucher

类型定义：[FixVoucher](ordervoucher_model.go#L307)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 面额，每张代金券可以抵扣的金额。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。该字段不填写，认为无门槛。 |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称。 |

<a id="model-fixvoucherinfo"></a>

### FixVoucherInfo

类型定义：[FixVoucherInfo](ordervoucher_model.go#L317)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | 未标注 | 面额，每张代金券可以抵扣的金额。 |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。该字段不填写，认为无门槛。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-goodsdetaildto"></a>

### GoodsDetailDTO

类型定义：[GoodsDetailDTO](ordervoucher_model.go#L326)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `goods_id` | `GoodsId` | `*string` | 未标注 | 商品编号 |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称 |
| `price` | `Price` | `*string` | 未标注 | 商品单价。单位 元 精确到小数点后2位。 |
| `quantity` | `Quantity` | `*int64` | 未标注 | 商品数量 |

<a id="model-holidayruleinfo"></a>

### HolidayRuleInfo

类型定义：[HolidayRuleInfo](ordervoucher_model.go#L338)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-ordervoucheravailablecitycode"></a>

### OrderVoucherAvailableCityCode

类型定义：[OrderVoucherAvailableCityCode](ordervoucher_model.go#L343)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `all_city` | `AllCity` | `*bool` | 未标注 | 是否全国。选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。 |
| `city_codes` | `CityCodes` | `*[]string` | 未标注 | 城市编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx)中内容填写。 （[参考资料](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)） |

<a id="model-ordervoucheravailablecitycodemodify"></a>

### OrderVoucherAvailableCityCodeModify

类型定义：[OrderVoucherAvailableCityCodeModify](ordervoucher_model.go#L351)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `city_codes` | `CityCodes` | `*[]string` | 未标注 | 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |

<a id="model-ordervoucheravailableshop"></a>

### OrderVoucherAvailableShop

类型定义：[OrderVoucherAvailableShop](ordervoucher_model.go#L364)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `order_voucher_merchant_all_shop` | `OrderVoucherMerchantAllShop` | [`*OrderVoucherMerchantAllShop`](#model-ordervouchermerchantallshop) | 未标注 | 见字段类型及官方接口文档。 |
| `real_shop_ids` | `RealShopIds` | `*[]string` | 未标注 | 代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。 |
| `shop_ids` | `ShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |

<a id="model-ordervoucheravailableshopmodify"></a>

### OrderVoucherAvailableShopModify

类型定义：[OrderVoucherAvailableShopModify](ordervoucher_model.go#L373)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `order_voucher_merchant_all_shop` | `OrderVoucherMerchantAllShop` | [`*OrderVoucherMerchantAllShopModify`](#model-ordervouchermerchantallshopmodify) | 未标注 | 见字段类型及官方接口文档。 |
| `real_shop_ids` | `RealShopIds` | `*[]string` | 未标注 | 券可使用的商业关系门店列表,列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。 |
| `shop_ids` | `ShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。接口参数是列表类型。 |

<a id="model-ordervouchermerchantallshop"></a>

### OrderVoucherMerchantAllShop

类型定义：[OrderVoucherMerchantAllShop](ordervoucher_model.go#L395)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `exclude_shop_ids` | `ExcludeShopIds` | `*[]string` | 未标注 | 券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |

<a id="model-ordervouchermerchantallshopmodify"></a>

### OrderVoucherMerchantAllShopModify

类型定义：[OrderVoucherMerchantAllShopModify](ordervoucher_model.go#L401)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `exclude_shop_ids` | `ExcludeShopIds` | `*[]string` | 未标注 | 券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |

<a id="model-specialvoucher"></a>

### SpecialVoucher

类型定义：[SpecialVoucher](ordervoucher_model.go#L440)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `goods_name` | `GoodsName` | `*string` | 未标注 | 商品名称。 |
| `origin_amount` | `OriginAmount` | `*string` | 未标注 | 原价。 限制： 币种为人民币，单位为元。小数点以后最多保留两位; 取值范围:0.1&lt;=x&lt;=9999。 |
| `special_amount` | `SpecialAmount` | `*string` | 未标注 | 特价，即：原价-特价=优惠金额。 |

<a id="model-specialvoucherinfo"></a>

### SpecialVoucherInfo

类型定义：[SpecialVoucherInfo](ordervoucher_model.go#L452)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `floor_amount` | `FloorAmount` | `*string` | 未标注 | 门槛金额。说明：该字段可不填，认为无门槛; |
| `special_amount` | `SpecialAmount` | `*string` | 未标注 | 特价，即：原价-特价=优惠金额。 |
| `voucher_deduct_threshold_info` | `VoucherDeductThresholdInfo` | [`*VoucherDeductThresholdInfo`](#model-voucherdeductthresholdinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-timerangeinfo"></a>

### TimeRangeInfo

类型定义：[TimeRangeInfo](ordervoucher_timerangeinfo_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `begin_time` | `BeginTime` | `*string` | 未标注 | 开始时间 格式：HH:mm:ss |
| `end_time_info` | `EndTimeInfo` | [`*EndTimeInfo`](#model-endtimeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-timerestrictinfo"></a>

### TimeRestrictInfo

类型定义：[TimeRestrictInfo](ordervoucher_timerangeinfo_model.go#L11)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `disable_period_info` | `DisablePeriodInfo` | [`*[]VoucherUseTimeRuleInfo`](#model-voucherusetimeruleinfo) | 未标注 | 不可用时间段 |
| `usable_period_info` | `UsablePeriodInfo` | [`*[]VoucherUseTimeRuleInfo`](#model-voucherusetimeruleinfo) | 未标注 | 可用时间段 |

<a id="model-voucherabsoluteperiodinfo"></a>

### VoucherAbsolutePeriodInfo

类型定义：[VoucherAbsolutePeriodInfo](ordervoucher_timerangeinfo_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_restrict_info` | `TimeRestrictInfo` | [`*TimeRestrictInfo`](#model-timerestrictinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `valid_begin_time` | `ValidBeginTime` | `*string` | 未标注 | 券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss。 |

<a id="model-voucheravailableaccountinfo"></a>

### VoucherAvailableAccountInfo

类型定义：[VoucherAvailableAccountInfo](ordervoucher_timerangeinfo_model.go#L28)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_pids` | `AvailablePids` | `*[]string` | 未标注 | 优惠券可核销的直连商户PID |
| `available_smids` | `AvailableSmids` | `*[]string` | 未标注 | 优惠券可核销的间连商户SMID |

<a id="model-voucheravailableappinfo"></a>

### VoucherAvailableAppInfo

类型定义：[VoucherAvailableAppInfo](ordervoucher_timerangeinfo_model.go#L36)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_app_ids` | `AvailableAppIds` | `*[]string` | 未标注 | 可核销的支付宝小程序id |

<a id="model-voucheravailablegeographyallshopinfo"></a>

### VoucherAvailableGeographyAllShopInfo

类型定义：[VoucherAvailableGeographyAllShopInfo](ordervoucher_timerangeinfo_model.go#L42)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_brand_id` | `AvailableBrandId` | `*string` | 未标注 | 全部门店可用的品牌ID 提示: 1.目前只支持传入当前券归属商户的品牌ID 2.品牌ID可以在登录b.alipay.com-账户中心-品牌管理里查看。如果没有品牌，可以先创建品牌。Sop： 注意事项 1.只支持传入单个品牌ID 2.传入品牌ID后，券的商户品牌和logo仍以voucher_display_pattern_info传入为准。 |
| `exclude_shop_ids` | `ExcludeShopIds` | `*[]string` | 未标注 | 券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。 列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |
| `merchant_ids` | `MerchantIds` | `*[]string` | 未标注 | 全部门店可用的商户ID |

<a id="model-voucheravailablegeographycityinfo"></a>

### VoucherAvailableGeographyCityInfo

类型定义：[VoucherAvailableGeographyCityInfo](ordervoucher_timerangeinfo_model.go#L60)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `all_city` | `AllCity` | `*bool` | 未标注 | 是否全国。 选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。 |
| `available_city_codes` | `AvailableCityCodes` | `*[]string` | 未标注 | 城市编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx)中内容填写。（[参考资料](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm)） |

<a id="model-voucheravailablegeographyscopeinfo"></a>

### VoucherAvailableGeographyScopeInfo

类型定义：[VoucherAvailableGeographyScopeInfo](ordervoucher_timerangeinfo_model.go#L68)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_geography_city_info` | `AvailableGeographyCityInfo` | [`*VoucherAvailableGeographyCityInfo`](#model-voucheravailablegeographycityinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `available_geography_scope_type` | `AvailableGeographyScopeType` | `*string` | 未标注 | 券可用地理位置类型。 |
| `available_geography_shop_info` | `AvailableGeographyShopInfo` | [`*VoucherAvailableGeographyShopInfo`](#model-voucheravailablegeographyshopinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucheravailablegeographyshopinfo"></a>

### VoucherAvailableGeographyShopInfo

类型定义：[VoucherAvailableGeographyShopInfo](ordervoucher_timerangeinfo_model.go#L81)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `available_geography_all_shop` | `AvailableGeographyAllShop` | [`*VoucherAvailableGeographyAllShopInfo`](#model-voucheravailablegeographyallshopinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `available_real_shop_ids` | `AvailableRealShopIds` | `*[]string` | 未标注 | 券可使用的物理门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。 |
| `available_shop_ids` | `AvailableShopIds` | `*[]string` | 未标注 | 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。 |

<a id="model-voucheravailablegoodsinfo"></a>

### VoucherAvailableGoodsInfo

类型定义：[VoucherAvailableGoodsInfo](ordervoucher_timerangeinfo_model.go#L103)。

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

类型定义：[VoucherAvailableItemInfo](ordervoucher_timerangeinfo_model.go#L123)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_id` | `ItemId` | `*string` | 未标注 | 内部商品ID |
| `out_item_info` | `OutItemInfo` | [`*[]VoucherAvailableOutItemInfo`](#model-voucheravailableoutiteminfo) | 未标注 | 外部商品信息 |

<a id="model-voucheravailableoutiteminfo"></a>

### VoucherAvailableOutItemInfo

类型定义：[VoucherAvailableOutItemInfo](ordervoucher_timerangeinfo_model.go#L131)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_app_id` | `ItemAppId` | `*string` | 未标注 | 小程序ID |
| `out_item_id` | `OutItemId` | `*string` | 未标注 | 外部商品ID |

<a id="model-voucheravailablescope"></a>

### VoucherAvailableScope

类型定义：[VoucherAvailableScope](ordervoucher_timerangeinfo_model.go#L139)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `order_voucher_available_city_code` | `OrderVoucherAvailableCityCode` | [`*OrderVoucherAvailableCityCode`](#model-ordervoucheravailablecitycode) | 未标注 | 见字段类型及官方接口文档。 |
| `order_voucher_available_shop` | `OrderVoucherAvailableShop` | [`*OrderVoucherAvailableShop`](#model-ordervoucheravailableshop) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_type` | `VoucherAvailableType` | `*string` | 未标注 | 可用范围类型。 |

<a id="model-voucheravailablescopeinfo"></a>

### VoucherAvailableScopeInfo

类型定义：[VoucherAvailableScopeInfo](ordervoucher_timerangeinfo_model.go#L147)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_available_account_info` | `VoucherAvailableAccountInfo` | [`*VoucherAvailableAccountInfo`](#model-voucheravailableaccountinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_app_info` | `VoucherAvailableAppInfo` | [`*VoucherAvailableAppInfo`](#model-voucheravailableappinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_geography_scope_info` | `VoucherAvailableGeographyScopeInfo` | [`*VoucherAvailableGeographyScopeInfo`](#model-voucheravailablegeographyscopeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_goods_info` | `VoucherAvailableGoodsInfo` | [`*VoucherAvailableGoodsInfo`](#model-voucheravailablegoodsinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_item_info` | `VoucherAvailableItemInfo` | [`*VoucherAvailableItemInfo`](#model-voucheravailableiteminfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucheravailablescopemodify"></a>

### VoucherAvailableScopeModify

类型定义：[VoucherAvailableScopeModify](ordervoucher_timerangeinfo_model.go#L156)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `modify_type` | `ModifyType` | `*string` | 未标注 | 可用范围修改类型。 |
| `order_voucher_available_city_code` | `OrderVoucherAvailableCityCode` | [`*OrderVoucherAvailableCityCodeModify`](#model-ordervoucheravailablecitycodemodify) | 未标注 | 见字段类型及官方接口文档。 |
| `order_voucher_available_shop` | `OrderVoucherAvailableShop` | [`*OrderVoucherAvailableShopModify`](#model-ordervoucheravailableshopmodify) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchercustomerguideinfo"></a>

### VoucherCustomerGuideInfo

类型定义：[VoucherCustomerGuideInfo](ordervoucher_timerangeinfo_model.go#L169)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_send_guide_info` | `VoucherSendGuideInfo` | [`*VoucherSendGuideInfo`](#model-vouchersendguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_use_guide_info` | `VoucherUseGuideInfo` | [`*VoucherUseGuideInfo`](#model-voucheruseguideinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherdeductinfo"></a>

### VoucherDeductInfo

类型定义：[VoucherDeductInfo](ordervoucher_timerangeinfo_model.go#L175)。

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

类型定义：[VoucherDeductThresholdInfo](ordervoucher_timerangeinfo_model.go#L186)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `customer_define_voucher_threshold_desc` | `CustomerDefineVoucherThresholdDesc` | `*string` | 未标注 | 自定义门槛描述 |
| `voucher_express_info` | `VoucherExpressInfo` | [`*VoucherExpressInfo`](#model-voucherexpressinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_goods_quantity_specification_info` | `VoucherGoodsQuantitySpecificationInfo` | [`*VoucherGoodsQuantitySpecificationInfo`](#model-vouchergoodsquantityspecificationinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherdisplayinfo"></a>

### VoucherDisplayInfo

类型定义：[VoucherDisplayInfo](ordervoucher_timerangeinfo_model.go#L194)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_logo` | `BrandLogo` | `*string` | 未标注 | 商家logo。需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值 |
| `brand_name` | `BrandName` | `*string` | 未标注 | 商户品牌名称。如果不设置品牌名称，则默认使用支付宝商家系统中维护的商家别名。如果没有维护商家别名，则使用默认名称：商家优惠。用户领取优惠券后，品牌名称会对用户进行展示 |
| `customer_service_mobile` | `CustomerServiceMobile` | `*string` | 未标注 | 客服电话 |
| `customer_service_url` | `CustomerServiceUrl` | `*string` | 未标注 | 客服链接 |
| `voucher_comment` | `VoucherComment` | `*string` | 未标注 | 券补充说明，该字段为可运营字段，可描述商品特色 会展示在用户支付宝卡包券详情页 |
| `voucher_description` | `VoucherDescription` | `*string` | 未标注 | 券详细使用说明。 |
| `voucher_detail_images` | `VoucherDetailImages` | `*[]string` | 未标注 | 券详细图列表，会展示在用户支付宝卡包券详情页。需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE ,接口返回的resource_id即为该参数的值。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。 |
| `voucher_image` | `VoucherImage` | `*string` | 未标注 | 券详情页封面图，会展示在用户支付宝卡包券详情页。需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值 |

<a id="model-voucherdisplayinfomodify"></a>

### VoucherDisplayInfoModify

类型定义：[VoucherDisplayInfoModify](ordervoucher_timerangeinfo_model.go#L214)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `brand_logo` | `BrandLogo` | `*string` | 未标注 | 商家logo 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_BRAND_LOGO，返回的resource_id即为该参数的值 |
| `brand_name` | `BrandName` | `*string` | 未标注 | 商户品牌名称。 |
| `customer_service_mobile` | `CustomerServiceMobile` | `*string` | 未标注 | 客服电话 |
| `customer_service_url` | `CustomerServiceUrl` | `*string` | 未标注 | 客服链接 |
| `voucher_comment` | `VoucherComment` | `*string` | 未标注 | 券补充说明，该字段为可运营字段，可描述商品特色 会展示在用户支付宝卡包券详情页。 |
| `voucher_description` | `VoucherDescription` | `*string` | 未标注 | 券详细使用说明。 |
| `voucher_detail_images` | `VoucherDetailImages` | `*[]string` | 未标注 | 券详细图列表，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_DETAIL_IMAGE ,接口返回的resource_id即为该参数的值 上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB。 |
| `voucher_image` | `VoucherImage` | `*string` | 未标注 | 券详情页封面图，会展示在用户支付宝卡包券详情页 需要通过 alipay.marketing.material.image.upload接口上传图片，指定file_key为PROMO_VOUCHER_IMAGE，接口返回的resource_id即为该参数的值。 |

<a id="model-voucherdisplaypatterninfo"></a>

### VoucherDisplayPatternInfo

类型定义：[VoucherDisplayPatternInfo](ordervoucher_timerangeinfo_model.go#L234)。

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

类型定义：[VoucherExchangeGoodsInfo](ordervoucher_timerangeinfo_model.go#L260)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `exchange_goods_name` | `ExchangeGoodsName` | `*string` | 未标注 | 兑换商品名称 |

<a id="model-voucherexpressinfo"></a>

### VoucherExpressInfo

类型定义：[VoucherExpressInfo](ordervoucher_timerangeinfo_model.go#L266)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `pay_express` | `PayExpress` | `*string` | 未标注 | 邮费标识，属于券的服务，用于用户门槛信息展示 |

<a id="model-vouchergoodsquantityspecificationinfo"></a>

### VoucherGoodsQuantitySpecificationInfo

类型定义：[VoucherGoodsQuantitySpecificationInfo](ordervoucher_timerangeinfo_model.go#L284)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `directional_type` | `DirectionalType` | `*string` | 未标注 | 定向类型用于规格设置，比如：FIX 表示 固定满X件，APPOINT 表示指定第N件 |
| `specification_quantity` | `SpecificationQuantity` | `*int64` | 未标注 | 数量 |
| `specification_unit` | `SpecificationUnit` | `*string` | 未标注 | 单位 |

<a id="model-voucherminiappsendguideinfo"></a>

### VoucherMiniAppSendGuideInfo

类型定义：[VoucherMiniAppSendGuideInfo](ordervoucher_timerangeinfo_model.go#L302)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_url` | `MiniAppUrl` | `*string` | 未标注 | 小程序发券链接 |

<a id="model-voucherminiappuseguideinfo"></a>

### VoucherMiniAppUseGuideInfo

类型定义：[VoucherMiniAppUseGuideInfo](ordervoucher_timerangeinfo_model.go#L308)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_service_codes` | `MiniAppServiceCodes` | `*[]string` | 未标注 | 小程序服务编码。取值来源：通过[alipay.open.app.service.apply(小程序服务创建)](https://opendocs.alipay.com/mini/03cgoq?pathHash=49cde24b&scene=common&highlight_field=service_code)接口创建服务后获取。 |
| `mini_app_url` | `MiniAppUrl` | `*string` | 未标注 | 小程序核销链接 |

<a id="model-voucherpackagemodeinfo"></a>

### VoucherPackageModeInfo

类型定义：[VoucherPackageModeInfo](ordervoucher_timerangeinfo_model.go#L316)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_package_id` | `VoucherPackageId` | `*string` | 未标注 | 券包id，对应alipay.marketing.activity.voucherpackage.query中voucher_package_id。 |

<a id="model-voucherrefunddetailinfo"></a>

### VoucherRefundDetailInfo

类型定义：[VoucherRefundDetailInfo](ordervoucher_timerangeinfo_model.go#L322)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_refund_amount` | `VoucherRefundAmount` | `*string` | 未标注 | 取消核销优惠金额, 单位:元 |
| `voucher_refund_use_times` | `VoucherRefundUseTimes` | `*int64` | 未标注 | 本次取消核销次数 |

<a id="model-voucherrelativeperiodinfo"></a>

### VoucherRelativePeriodInfo

类型定义：[VoucherRelativePeriodInfo](ordervoucher_timerangeinfo_model.go#L330)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_restrict_info` | `TimeRestrictInfo` | [`*TimeRestrictInfo`](#model-timerestrictinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后 N 天内可以使用。 可以配合wait_days_after_receive 字段使用。 |
| `valid_days_relative_type` | `ValidDaysRelativeType` | `*string` | 未标注 | 优惠券有效期的取整方式，需与 valid_days_after_receive 字段配合使用，用于确定用户领取优惠券后的具体失效时间。 取值逻辑说明：ROUND_FORWARD（向前取整），优惠券在领取当日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 示例： -若 valid_days_relative_type=ROUND_FORWARD，valid_days_after_receive=1，用户于 1 号 09:12:40 领券，则优惠券在 1 号 23:59:59 失效。 -若 valid_days_after_receive=2，则在 2 号 23:59:59 失效。ROUND_BACKWARD（向后取整），优惠券从领取次日开始计算，于第 N 天（即 valid_days_after_receive 的值）的 当天 23:59:59 失效。 示例： -若valid_days_relative_type=ROUND_BACKWARDS，valid_days_after_receive=1，领券时间是1号9点12分40秒，那么 2 号 23:59:59 失效。 -若 valid_days_after_receive=2，则在 3 号 23:59:59 失效。 |
| `wait_days_after_receive` | `WaitDaysAfterReceive` | `*int64` | 未标注 | 用户领券后需要等待 N 天，券才可以生效。 字段值为 0 表示用户领券后立刻生效。 |

<a id="model-vouchersalemodeinfo"></a>

### VoucherSaleModeInfo

类型定义：[VoucherSaleModeInfo](ordervoucher_timerangeinfo_model.go#L341)。

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

<a id="model-vouchersendguide"></a>

### VoucherSendGuide

类型定义：[VoucherSendGuide](ordervoucher_timerangeinfo_model.go#L360)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_detail_url` | `VoucherDetailUrl` | `*string` | 未标注 | 领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。 |

<a id="model-vouchersendguideinfo"></a>

### VoucherSendGuideInfo

类型定义：[VoucherSendGuideInfo](ordervoucher_timerangeinfo_model.go#L366)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_send_guide_info` | `MiniAppSendGuideInfo` | [`*VoucherMiniAppSendGuideInfo`](#model-voucherminiappsendguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `send_guide_mode` | `SendGuideMode` | `*[]string` | 未标注 | 发放引导模式 |

<a id="model-vouchersendguidemodify"></a>

### VoucherSendGuideModify

类型定义：[VoucherSendGuideModify](ordervoucher_timerangeinfo_model.go#L373)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_detail_url` | `VoucherDetailUrl` | `*string` | 未标注 | 领（购）券详情页链接，从支付宝公域跳转到服务商（商户）自定义领(购)券详情页。说明：当 voucher_type=EXCHANGE_VOUCHER 时，该字段可修改，其他不允许修改。 |

<a id="model-vouchersendmodeinfo"></a>

### VoucherSendModeInfo

类型定义：[VoucherSendModeInfo](ordervoucher_timerangeinfo_model.go#L379)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `delivery_channel_list` | `DeliveryChannelList` | `*[]string` | 未标注 | 推广渠道列表 |
| `voucher_package_mode_info` | `VoucherPackageModeInfo` | [`*VoucherPackageModeInfo`](#model-voucherpackagemodeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_sale_mode_info` | `VoucherSaleModeInfo` | [`*VoucherSaleModeInfo`](#model-vouchersalemodeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_send_mode` | `VoucherSendMode` | `*string` | 未标注 | 券发放模式。 |
| `voucher_send_rule_info` | `VoucherSendRuleInfo` | [`*VoucherSendRuleInfo`](#model-vouchersendruleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchersendruledetail"></a>

### VoucherSendRuleDetail

类型定义：[VoucherSendRuleDetail](ordervoucher_timerangeinfo_model.go#L390)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `natural_person_limit` | `NaturalPersonLimit` | `*bool` | 未标注 | 是否开启自然人领取限制。自然人表示按照身份证纬度进行领取限制。 |
| `phone_number_limit` | `PhoneNumberLimit` | `*bool` | 未标注 | 是否开启电话号码领取限制。 |
| `voucher_quantity` | `VoucherQuantity` | `*int64` | 未标注 | 发行券的数量。 |
| `voucher_quantity_limit_per_user` | `VoucherQuantityLimitPerUser` | `*int64` | 未标注 | 每人领取限制。默认按照支付宝uid进行领取限制; 不填写或填入0，默认没有领取限制。 |

<a id="model-vouchersendruledetailmodify"></a>

### VoucherSendRuleDetailModify

类型定义：[VoucherSendRuleDetailModify](ordervoucher_timerangeinfo_model.go#L402)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `natural_person_limit` | `NaturalPersonLimit` | `*bool` | 未标注 | 是否开启自然人领取限制。 自然人表示按照身份证纬度进行领取限制。 |
| `phone_number_limit` | `PhoneNumberLimit` | `*string` | 未标注 | 是否开启电话号码领取限制。 |
| `voucher_quantity_limit_per_user` | `VoucherQuantityLimitPerUser` | `*int64` | 未标注 | 每人领取限制。 默认按照支付宝uid进行领取限制; |

<a id="model-vouchersendruleinfo"></a>

### VoucherSendRuleInfo

类型定义：[VoucherSendRuleInfo](ordervoucher_timerangeinfo_model.go#L412)。

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

<a id="model-voucherusedetailinfo"></a>

### VoucherUseDetailInfo

类型定义：[VoucherUseDetailInfo](ordervoucher_timerangeinfo_model.go#L450)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `voucher_use_amount` | `VoucherUseAmount` | `*string` | 未标注 | 核销优惠金额, 单位:元 |
| `voucher_use_times` | `VoucherUseTimes` | `*int64` | 未标注 | 本次核销次数（不包含撤销核销的次数） |

<a id="model-voucheruseguide"></a>

### VoucherUseGuide

类型定义：[VoucherUseGuide](ordervoucher_voucheruseguide_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 券可用的小程序appId，卡包详情页可跳转到该appId |
| `mini_app_path` | `MiniAppPath` | `*string` | 未标注 | 指定跳转到mini_app_id时的具体页面路径。 |
| `mini_app_use_guide` | `MiniAppUseGuide` | `*bool` | 未标注 | 小程序核销。 |
| `offline_code_use_guide` | `OfflineCodeUseGuide` | `*bool` | 未标注 | 线下扫码核销。 |
| `service_codes` | `ServiceCodes` | `*[]string` | 未标注 | 小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 |

<a id="model-voucheruseguideinfo"></a>

### VoucherUseGuideInfo

类型定义：[VoucherUseGuideInfo](ordervoucher_voucheruseguide_model.go#L18)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_use_guide_info` | `MiniAppUseGuideInfo` | [`*VoucherMiniAppUseGuideInfo`](#model-voucherminiappuseguideinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `use_guide_mode` | `UseGuideMode` | `*[]string` | 未标注 | 使用引导模式 |
| `use_url` | `UseUrl` | `*string` | 未标注 | 核销跳转链接 |

<a id="model-voucheruseguidemodify"></a>

### VoucherUseGuideModify

类型定义：[VoucherUseGuideModify](ordervoucher_voucheruseguide_model.go#L27)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 券可用的小程序 appId，卡包详情页可跳转到该 appId。 |
| `mini_app_path` | `MiniAppPath` | `*string` | 未标注 | 指定跳转到 mini_app_id 时的具体页面路径。 |
| `service_codes` | `ServiceCodes` | `*[]string` | 未标注 | 小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。 |

<a id="model-voucheruserule"></a>

### VoucherUseRule

类型定义：[VoucherUseRule](ordervoucher_voucheruseguide_model.go#L37)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_voucher` | `DiscountVoucher` | [`*DiscountVoucher`](#model-discountvoucher) | 未标注 | 见字段类型及官方接口文档。 |
| `exchange_voucher` | `ExchangeVoucher` | [`*ExchangeVoucher`](#model-exchangevoucher) | 未标注 | 见字段类型及官方接口文档。 |
| `fix_voucher` | `FixVoucher` | [`*FixVoucher`](#model-fixvoucher) | 未标注 | 见字段类型及官方接口文档。 |
| `special_voucher` | `SpecialVoucher` | [`*SpecialVoucher`](#model-specialvoucher) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_scope` | `VoucherAvailableScope` | [`*VoucherAvailableScope`](#model-voucheravailablescope) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_valid_period` | `VoucherValidPeriod` | [`*VoucherValidPeriod`](#model-vouchervalidperiod) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucheruseruleinfo"></a>

### VoucherUseRuleInfo

类型定义：[VoucherUseRuleInfo](ordervoucher_voucheruseguide_model.go#L47)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `payment_restriction_list` | `PaymentRestrictionList` | `*[]string` | 未标注 | 支付方式限制 |
| `quantity_limit_per_user` | `QuantityLimitPerUser` | `*int64` | 未标注 | 核销次数限制。 |
| `quantity_limit_per_user_period_type` | `QuantityLimitPerUserPeriodType` | `*string` | 未标注 | 周期限领配置,限制每人在固定周期内领取张数(quantity_limit_per_user) |
| `voucher_max_use_times` | `VoucherMaxUseTimes` | `*int64` | 未标注 | 券最大核销次数。 |
| `voucher_use_ext_info` | `VoucherUseExtInfo` | `*string` | 未标注 | 自定义扩展信息 |
| `voucher_use_time_info` | `VoucherUseTimeInfo` | [`*VoucherUseTimeInfo`](#model-voucherusetimeinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucheruserulemodify"></a>

### VoucherUseRuleModify

类型定义：[VoucherUseRuleModify](ordervoucher_voucheruseguide_model.go#L62)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `exchange_voucher` | `ExchangeVoucher` | [`*ExchangeVoucherModify`](#model-exchangevouchermodify) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_available_scope` | `VoucherAvailableScope` | [`*VoucherAvailableScopeModify`](#model-voucheravailablescopemodify) | 未标注 | 见字段类型及官方接口文档。 |
| `voucher_valid_period` | `VoucherValidPeriod` | [`*VoucherValidPeriodModify`](#model-vouchervalidperiodmodify) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherusetimeinfo"></a>

### VoucherUseTimeInfo

类型定义：[VoucherUseTimeInfo](ordervoucher_voucheruseguide_model.go#L69)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `absolute_period_info` | `AbsolutePeriodInfo` | [`*VoucherAbsolutePeriodInfo`](#model-voucherabsoluteperiodinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `period_type` | `PeriodType` | `*string` | 未标注 | 券有效期。 |
| `relative_period_info` | `RelativePeriodInfo` | [`*VoucherRelativePeriodInfo`](#model-voucherrelativeperiodinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-voucherusetimeruleinfo"></a>

### VoucherUseTimeRuleInfo

类型定义：[VoucherUseTimeRuleInfo](ordervoucher_voucheruseguide_model.go#L77)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `date_rule_info` | `DateRuleInfo` | [`*DateRuleInfo`](#model-dateruleinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `holiday_rule_info` | `HolidayRuleInfo` | [`*HolidayRuleInfo`](#model-holidayruleinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `rule_type` | `RuleType` | `*string` | 未标注 | 规则类型 |
| `week_rule_info` | `WeekRuleInfo` | [`*WeekRuleInfo`](#model-weekruleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-vouchervalidperiod"></a>

### VoucherValidPeriod

类型定义：[VoucherValidPeriod](ordervoucher_voucheruseguide_model.go#L86)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `type` | `Type` | `*string` | 未标注 | 券有效期类型 |
| `valid_begin_time` | `ValidBeginTime` | `*string` | 未标注 | 券可使用的开始时间。格式为：yyyy-MM-dd HH:mm:ss |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后N天内可以使用。可以配合wait_days_after_receive字段使用。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。格式为yyyy-MM-dd HH:mm:ss |
| `wait_days_after_receive` | `WaitDaysAfterReceive` | `*int64` | 未标注 | 用户领券后需要等待N天，券才可以生效。如果该字段填入0或者不填写，则用户领券后立刻生效。 |

<a id="model-vouchervalidperiodmodify"></a>

### VoucherValidPeriodModify

类型定义：[VoucherValidPeriodModify](ordervoucher_voucheruseguide_model.go#L100)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `valid_begin_time` | `ValidBeginTime` | `*string` | 未标注 | 券可使用的开始时间。格式为：yyyy-MM-dd HH:mm:ss。 |
| `valid_days_after_receive` | `ValidDaysAfterReceive` | `*int64` | 未标注 | 券生效后 N 天内可以使用。 |
| `valid_end_time` | `ValidEndTime` | `*string` | 未标注 | 券可使用的结束时间。格式为yyyy-MM-dd HH:mm:ss。 |
| `wait_days_after_receive` | `WaitDaysAfterReceive` | `*int64` | 未标注 | 用户领券后需要等待 N 天，券才可以生效。 |

<a id="model-weekruleinfo"></a>

### WeekRuleInfo

类型定义：[WeekRuleInfo](ordervoucher_voucheruseguide_model.go#L112)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `time_range_info` | `TimeRangeInfo` | [`*TimeRangeInfo`](#model-timerangeinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `week_day` | `WeekDay` | `*string` | 未标注 | 星期数，以“,”连接多个星期数 |
