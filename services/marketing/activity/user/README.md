# marketing/activity/user 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/user`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `user.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_user_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityUserBatchQueryVoucher`](#api-activityuserbatchqueryvoucher) | [`alipay.marketing.activity.user.batchqueryvoucher`](https://opendocs.alipay.com/open-v3/63264abf_alipay.marketing.activity.user.batchqueryvoucher) | 条件查询用户券 | `GET` |
| [`ActivityUserQueryVoucher`](#api-activityuserqueryvoucher) | [`alipay.marketing.activity.user.queryvoucher`](https://opendocs.alipay.com/open-v3/e75be5aa_alipay.marketing.activity.user.queryvoucher) | 查询用户券详情 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activityuserbatchqueryvoucher"></a>

### ActivityUserBatchQueryVoucher — 条件查询用户券

- 官方接口：[`alipay.marketing.activity.user.batchqueryvoucher`](https://opendocs.alipay.com/open-v3/63264abf_alipay.marketing.activity.user.batchqueryvoucher)
- 路由：`GET /v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher`
- 请求类型：[ActivityUserBatchQueryVoucherRequest](user_request.go#L5)
- 响应类型：[ActivityUserBatchQueryVoucherResponse](user_response.go#L5)

可通过该接口查询用户在某商户可用的全部券，可用于商户的小程序/H5中，用户"我的优惠券"或"提交订单页"展示优惠信息。只能查询商户配的券，无法查询平台优惠券。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `user_id` | `UserId` | `*string` | query | 否 | 支付宝用户 id |
| `open_id` | `OpenId` | `*string` | query | 否 | 支付宝用户openId |
| `activity_id` | `ActivityId` | `*string` | query | 否 | 活动 id |
| `belong_merchant_id` | `BelongMerchantId` | `*string` | query | 否 | 券归属商户 |
| `sender_merchant_id` | `SenderMerchantId` | `*string` | query | 否 | 券发放商户 |
| `voucher_status` | `VoucherStatus` | `*string` | query | 否 | 券状态，其中已过期状态可查询6个月内的数据 |
| `page_num` | `PageNum` | `*string` | query | 否 | 分页查询页码 |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | query | 否 | 商户接入模式 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页查询单页数据条数 |

<a id="api-activityuserqueryvoucher"></a>

### ActivityUserQueryVoucher — 查询用户券详情

- 官方接口：[`alipay.marketing.activity.user.queryvoucher`](https://opendocs.alipay.com/open-v3/e75be5aa_alipay.marketing.activity.user.queryvoucher)
- 路由：`GET /v3/alipay/marketing/activity/ordervoucher/user/voucher`
- 请求类型：[ActivityUserQueryVoucherRequest](user_request.go#L27)
- 响应类型：[ActivityUserQueryVoucherResponse](user_response.go#L17)

通过此接口可以查询用户券信息，包括用户券的基础信息、状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_id` | `MerchantId` | `*string` | query | 否 | 商户 PID，默认为当前接口调用商户。 |
| `user_id` | `UserId` | `*string` | query | 否 | 支付宝user_id账号 |
| `open_id` | `OpenId` | `*string` | query | 否 | 领券的支付宝用户openId |
| `activity_id` | `ActivityId` | `*string` | query | 否 | 活动id |
| `voucher_id` | `VoucherId` | `*string` | query | 否 | 用户券 id。支付宝为用户优惠券唯一分配的 id。 |
| `voucher_code` | `VoucherCode` | `*string` | query | 否 | 用户领取的商家券券码。 |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | query | 否 | 商户接入模式。 |
