# live 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/merchant/live`。本包提供 2 个支付宝 V3 REST 操作。

通过 `live.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`LiveChannelQuery`](#api-livechannelquery) | `alipay.merchant.live.channel.query` |
| [`LiveItemPlanModify`](#api-liveitemplanmodify) | `alipay.merchant.live.itemplan.modify` |

## 使用示例

- [LiveChannelQuery](api_live_livechannelquery_example_test.go)
- [LiveItemPlanModify](api_live_liveitemplanmodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-livechannelquery"></a>

### LiveChannelQuery — 渠道信息查询

- 官方接口：[`alipay.merchant.live.channel.query`](https://opendocs.alipay.com/apis/069r05)
- 路由：`GET /v3/alipay/merchant/live/channel/query`
- 请求类型：[MerchantLiveChannelQueryRequest](live_request.go#L30)
- 响应类型：[MerchantLiveChannelQueryResponse](live_response.go#L4)

根据直播等内容域跳转小程序下单时携带的orderStartId，回查业务主体信息(如操作人、主播、直播、文章等，用于外部业务信息补全、返佣等)。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `secret` | `Secret` | `*string` | query | 否 | 渠道下单参数(orderStartId) |

<a id="api-liveitemplanmodify"></a>

### LiveItemPlanModify — 支付宝货品推广计划创建

- 官方接口：[`alipay.merchant.live.itemplan.modify`](https://opendocs.alipay.com/apis/069o8w)
- 路由：`POST /v3/alipay/merchant/live/itemplan/modify`
- 请求类型：[MerchantLiveItemPlanModifyRequest](live_request.go#L4)
- 响应类型：[MerchantLiveItemPlanModifyResponse](live_response.go#L16)

支付宝货品推广计划创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `content` | `Content` | [`*[]ItemPlanContentDO`](#model-itemplancontentdo) | JSON body | 未标注 | 投放内容，数组，内容为具体条目 |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 可空，计划截止时间 |
| `op_type` | `OpType` | `*string` | JSON body | 未标注 | 针对投放内容的操作， 包含3种类型： 覆盖:OVERRIDE 追加:APPEND 删除:DELETE |
| `operator_appid` | `OperatorAppid` | `*string` | JSON body | 未标注 | 投放方标识，一般为小程序id |
| `operator_type` | `OperatorType` | `*string` | JSON body | 未标注 | 投放方类型，生活号：LIFE_APP 达人：CONTENT_USER 小程序：TINYAPP |
| `plan_name` | `PlanName` | `*string` | JSON body | 未标注 | 投放计划名，可空 |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 可空，计划生效时间 |
| `status` | `Status` | `*int64` | JSON body | 未标注 | 计划状态 0待生效 1生效中 -1已作废 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 投放目标，一般为生活号id，多个逗号隔开 |
| `target_range` | `TargetRange` | `*string` | JSON body | 未标注 | 投放范围， 指主体内的具体子渠道，为空则任意子渠道 |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 投放目标类型 生活号：LIFE_APP 达人：CONTENT_USER 小程序:TINYAPP |

## 嵌套请求参数

<a id="model-itemplancontentdo"></a>

### ItemPlanContentDO

类型定义：[ItemPlanContentDO](live_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `codes` | `Codes` | `*[]string` | 未标注 | 对应服务库编码 |
| `type` | `Type` | `*string` | 未标注 | 条目类型，必须， 商品：ITEM |
