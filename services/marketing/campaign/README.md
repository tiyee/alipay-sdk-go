# marketing/campaign 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/campaign`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `campaign.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CampaignCashCreate`](#api-campaigncashcreate) | [`alipay.marketing.campaign.cash.create`](https://opendocs.alipay.com/open-v3/e572d9c1_alipay.marketing.campaign.cash.create?scene=common&pathHash=5197b580) | 创建现金活动 | `POST` |
| [`CampaignCashDetailQuery`](#api-campaigncashdetailquery) | [`alipay.marketing.campaign.cash.detail.query`](https://opendocs.alipay.com/open-v3/7fede5af_alipay.marketing.campaign.cash.detail.query?scene=common&pathHash=430db29a) | 现金活动详情查询 | `GET` |
| [`CampaignCashListQuery`](#api-campaigncashlistquery) | [`alipay.marketing.campaign.cash.list.query`](https://opendocs.alipay.com/open-v3/b0871b97_alipay.marketing.campaign.cash.list.query?scene=common&pathHash=e06176c5) | 现金活动列表查询 | `GET` |
| [`CampaignCashStatusModify`](#api-campaigncashstatusmodify) | [`alipay.marketing.campaign.cash.status.modify`](https://opendocs.alipay.com/open-v3/4f41b60d_alipay.marketing.campaign.cash.status.modify?scene=common&pathHash=7b368c26) | 更改现金活动状态 | `POST` |
| [`CampaignCashTrigger`](#api-campaigncashtrigger) | [`alipay.marketing.campaign.cash.trigger`](https://opendocs.alipay.com/open-v3/06823eee_alipay.marketing.campaign.cash.trigger?scene=common&pathHash=e770d208) | 触发现金红包活动 | `POST` |
| [`CampaignOrderVoucherConsult`](#api-campaignordervoucherconsult) | [`alipay.marketing.campaign.order.voucher.consult`](https://opendocs.alipay.com/open-v3/bb374dd7_alipay.marketing.campaign.order.voucher.consult) | 订单优惠前置咨询 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-campaigncashcreate"></a>

### CampaignCashCreate — 创建现金活动

- 官方接口：[`alipay.marketing.campaign.cash.create`](https://opendocs.alipay.com/open-v3/e572d9c1_alipay.marketing.campaign.cash.create?scene=common&pathHash=5197b580)
- 路由：`POST /v3/alipay/marketing/campaign/cash/create`
- 请求类型：[CampaignCashCreateRequest](campaign_request.go#L5)
- 响应类型：[CampaignCashCreateResponse](campaign_response.go#L9)

商户通过开放平台创建商户活动

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `coupon_name` | `CouponName` | `*string` | JSON body | 未标注 | 红包名称,商户在查询列表、详情看到的名字,同时也会显示在商户付款页面。 |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 活动结束时间,必须大于活动开始时间, 基本格式:yyyy-MM-dd HH:mm:ss,活动有效时间最长为6个月，过期后需要创建新的活动。 |
| `merchant_link` | `MerchantLink` | `*string` | JSON body | 未标注 | 商户打款后的跳转链接，从支付宝收银台打款成功后的跳转链接。不填时，打款后停留在支付宝支付成功页。商户实际跳转会自动添加crowdNo作为跳转参数。示例: http://www.yourhomepage.com?crowdNo=XXX |
| `prize_msg` | `PrizeMsg` | `*string` | JSON body | 未标注 | 活动文案，用户在账单、红包中看到的账单描述、红包描述。 |
| `prize_type` | `PrizeType` | `*string` | JSON body | 未标注 | 现金红包的发放形式,。枚举支持： *fixed：固定金额。 *random：随机金额。选择随机金额时，单个红包的金额在平均金额的0.5~1.5倍之间浮动。 |
| `send_freqency` | `SendFreqency` | `*string` | JSON body | 未标注 | 用户在当前活动参与次数、频率限制。支持日(D)、周(W)、月(M)、终身(L)维度的限制，整个字段不填时默认值为L1（即用户终生仅能参与1次）。 注意： 1. 终身(L)限制必选且最大值为 100。 2. 日(D)、周(W)、月(M)限制可选且最多只能选择一个（即限制日领取次数就不能再限制月领取次数），频率设置必须小于等于终身(L)的次数。 3. 多个配置之间使用"&#124;"进行分隔。如 D3&#124;L4（表示限制用户终生只能参与4次，单日只能参与3次）。 4. 允许多次领取时，活动触发接口需要传入out_biz_no来配合。 |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 活动开始时间，必须大于活动创建的时间。入参支持如下两种形式： 1. 填入固定时间，时间格式为 yyyy-MM-dd HH:mm:ss，如 2020-12-10 22:28:30 2. 填字符串 NowTime，表示创建即生效。 |
| `total_money` | `TotalMoney` | `*string` | JSON body | 未标注 | 活动发放的现金总金额,最小金额1.00元,最大金额10000000.00元。每个红包的最大金额不允许超过200元,最小金额不得低于0.20元。 实际的金额限制可能会根据业务进行动态调整。 |
| `total_num` | `TotalNum` | `*string` | JSON body | 未标注 | 红包发放个数，最小为 1 个，最大10000000个。 说明：不同的发放形式（即prize_type）会使得发放个数含义不同： 1. 若prize_type 为 fixed（固定金额），则每个用户领取的红包金额为total_money除以total_num得到固定金额。 2. 若prize_type为 random（随机金额），则每个用户领取的红包金额为total_money除以total_num得到的平均金额值的0.5~1.5倍。由于金额是随机的，在红包金额全部被领取完时，有可能total_num有所剩余、或者大于设置值的情况。 |

<a id="api-campaigncashdetailquery"></a>

### CampaignCashDetailQuery — 现金活动详情查询

- 官方接口：[`alipay.marketing.campaign.cash.detail.query`](https://opendocs.alipay.com/open-v3/7fede5af_alipay.marketing.campaign.cash.detail.query?scene=common&pathHash=430db29a)
- 路由：`GET /v3/alipay/marketing/campaign/cash/detail/query`
- 请求类型：[CampaignCashDetailQueryRequest](campaign_request.go#L77)
- 响应类型：[CampaignCashDetailQueryResponse](campaign_response.go#L19)

商户通过开放平台查询自己创建的现金活动详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `crowd_no` | `CrowdNo` | `*string` | query | 否 | 现金活动号，通过[alipay.marketing.campaign.cash.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.create)(创建现金活动)接口创建现金活动获取。 |

<a id="api-campaigncashlistquery"></a>

### CampaignCashListQuery — 现金活动列表查询

- 官方接口：[`alipay.marketing.campaign.cash.list.query`](https://opendocs.alipay.com/open-v3/b0871b97_alipay.marketing.campaign.cash.list.query?scene=common&pathHash=e06176c5)
- 路由：`GET /v3/alipay/marketing/campaign/cash/list/query`
- 请求类型：[CampaignCashListQueryRequest](campaign_request.go#L83)
- 响应类型：[CampaignCashListQueryResponse](campaign_response.go#L47)

商户通过开放平台查询自己创建的现金活动列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `camp_status` | `CampStatus` | `*string` | query | 否 | 要查询的活动状态,不填默认为 ALL 返回所有类型。枚举支持: *ALL：所有类型的活动。 *CREATED：已创建未打款。 *PAID：已打款。 *READY：活动已开始。 *PAUSE：活动已暂停。 *CLOSED：活动已结束。 *SETTLE：活动已清算。 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页查询时每页返回的列表大小，每页数据最大为 50。 |
| `page_index` | `PageIndex` | `*string` | query | 否 | 分页查询时的页码，从1开始 |

<a id="api-campaigncashstatusmodify"></a>

### CampaignCashStatusModify — 更改现金活动状态

- 官方接口：[`alipay.marketing.campaign.cash.status.modify`](https://opendocs.alipay.com/open-v3/4f41b60d_alipay.marketing.campaign.cash.status.modify?scene=common&pathHash=7b368c26)
- 路由：`POST /v3/alipay/marketing/campaign/cash/status/modify`
- 请求类型：[CampaignCashStatusModifyRequest](campaign_request.go#L27)
- 响应类型：[CampaignCashStatusModifyResponse](campaign_response.go#L59)

商户更改现金活动状态，启动/暂停/终止

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `camp_status` | `CampStatus` | `*string` | JSON body | 未标注 | 金活动修改后的状态。支持修改为： *PAUSE：活动暂停。 *READY：活动开始。 *CLOSED：活动结束。 |
| `crowd_no` | `CrowdNo` | `*string` | JSON body | 未标注 | 现金活动号，通过[alipay.marketing.campaign.cash.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.create)(创建现金活动)接口创建现金活动获取。 |

<a id="api-campaigncashtrigger"></a>

### CampaignCashTrigger — 触发现金红包活动

- 官方接口：[`alipay.marketing.campaign.cash.trigger`](https://opendocs.alipay.com/open-v3/06823eee_alipay.marketing.campaign.cash.trigger?scene=common&pathHash=e770d208)
- 路由：`POST /v3/alipay/marketing/campaign/cash/trigger`
- 请求类型：[CampaignCashTriggerRequest](campaign_request.go#L35)
- 响应类型：[CampaignCashTriggerResponse](campaign_response.go#L62)

用户通过开放平台参与现金活动

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `crowd_no` | `CrowdNo` | `*string` | JSON body | 未标注 | 现金活动号，通过[alipay.marketing.campaign.cash.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.create)(创建现金活动)接口创建现金活动获取。 |
| `login_id` | `LoginId` | `*string` | JSON body | 未标注 | 用户登录账号名，一般为邮箱或手机号。 注意：user_id与login_id二选一必填，同时传入时以user_id为准。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | open_id |
| `order_price` | `OrderPrice` | `*string` | JSON body | 未标注 | 此字段如果传入金额，就忽略prize_type算法，按照传入的金额发奖。如果不传或者小于等于0，则按照活动创建时指定的prize_type为fixed或者random算法发奖 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 领取红包的外部业务号，只由可由字母、数字、下划线组成。同一个活动中不可重复，相同的外部业务号会被幂等并返回之前的结果。不填时，系统会生成一个默认固定的外部业务号。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 用户支付宝唯一标识，2088开头。 注意：user_id与login_id二选一必填，同时传入时以user_id为准。 |

<a id="api-campaignordervoucherconsult"></a>

### CampaignOrderVoucherConsult — 订单优惠前置咨询

- 官方接口：[`alipay.marketing.campaign.order.voucher.consult`](https://opendocs.alipay.com/open-v3/bb374dd7_alipay.marketing.campaign.order.voucher.consult)
- 路由：`POST /v3/alipay/marketing/campaign/order/voucher/consult`
- 请求类型：[CampaignOrderVoucherConsultRequest](campaign_request.go#L51)
- 响应类型：[CampaignOrderVoucherConsultResponse](campaign_response.go#L86)

本接口用于前置优惠咨询，包括本订单可享的优惠券列表，注意该列表并不等价于最终核销时使用的券列表，而是包含了咨询的那一刻该订单能用的所有券，且券满足叠加使用下的互斥等规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `business_param` | `BusinessParam` | `*string` | JSON body | 未标注 | 业务参数大字段，优惠咨询的控制参数，json格式；目前支持传入useBigAmountSkipOrderThold为N来控制不使用大金额跳过优惠的订单门槛检查；默认不传； |
| `item_consult_list` | `ItemConsultList` | [`*[]ItemConsultRequest`](campaign_request.go#L67) | JSON body | 未标注 | 商品咨询请求列表（当需要咨询单品券时必传，如果某商品不希望参与本次单品优惠咨询则不传递对应信息即可） |
| `order_amount` | `OrderAmount` | `*string` | JSON body | 未标注 | 订单金额（如同时享受商户自有优惠请先扣除后传入），单位为元，最多2位小数 |
| `scene_code` | `SceneCode` | `*[]string` | JSON body | 未标注 | 场景码：默认(DEFAULT) |
| `specified_app_id` | `SpecifiedAppId` | `*string` | JSON body | 未标注 | 券指定的核销appid（如果配券时指定了核销范围为线上小程序及相应的appid则此处必传） |
