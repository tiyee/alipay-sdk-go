# alipaypoint 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/user/alipaypoint`。本包提供 2 个支付宝 V3 REST 操作。

通过 `alipaypoint.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AlipayPointBudgetLibQuery`](#api-alipaypointbudgetlibquery) | `alipay.user.alipaypoint.budgetlib.query` |
| [`AlipayPointSend`](#api-alipaypointsend) | `alipay.user.alipaypoint.send` |

## 使用示例

- [AlipayPointBudgetLibQuery](api_alipaypoint_alipaypointbudgetlibquery_example_test.go)
- [AlipayPointSend](api_alipaypoint_alipaypointsend_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-alipaypointbudgetlibquery"></a>

### AlipayPointBudgetLibQuery — 查询集分宝预算库详情

- 官方接口：[`alipay.user.alipaypoint.budgetlib.query`](https://opendocs.alipay.com/apis/05wmz4)
- 路由：`GET /v3/alipay/user/alipaypoint/budgetlib/query`
- 请求类型：[UserAlipayPointBudgetLibQueryRequest](alipaypoint_request.go#L22)
- 响应类型：[UserAlipayPointBudgetLibQueryResponse](alipaypoint_response.go#L4)

用于查询集分宝预算库详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `budget_code` | `BudgetCode` | `*string` | query | 否 | 集分宝预算库编码 |

<a id="api-alipaypointsend"></a>

### AlipayPointSend — 集分宝发放接口

- 官方接口：[`alipay.user.alipaypoint.send`](https://opendocs.alipay.com/apis/05wvr1)
- 路由：`POST /v3/alipay/user/alipaypoint/send`
- 请求类型：[UserAlipayPointSendRequest](alipaypoint_request.go#L4)
- 响应类型：[UserAlipayPointSendResponse](alipaypoint_response.go#L22)

用于发放商户采购的集分宝给到支付宝用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `budget_code` | `BudgetCode` | `*string` | JSON body | 未标注 | 签约商户的集分宝的预算库，扣除此预算库的集分宝发放给用户。会校验budgetcode和业务方appId的签约商户pid的关联关系，若无关则发放失败。 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 商户关于该笔发放的描述或者信息补充，仅存储，无实际校验功能，该信息会在"集分宝"小程序的"集分宝明细"中展示给用户。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 被发放集分宝用户ID，商户app_id维度下的用户标识，与user_account字段二选一。 建议优先使该字段指定发放用户，接口性能更佳 |
| `partner_biz_no` | `PartnerBizNo` | `*string` | JSON body | 未标注 | 接入方自己交易的唯一流水ID号，不同交易请保证幂等号唯一性，集分宝服务将依据该字段来进行幂等控制，重试请求不要更换幂等号，需要慎重传递，否则可能会造成损失。 |
| `point_amount` | `PointAmount` | `*int64` | JSON body | 未标注 | 发放给用户的集分宝个数。个数区间为 [1,10000000]，需为整数。 |
| `user_account` | `UserAccount` | `*string` | JSON body | 未标注 | 被发放集分宝用户的支付宝登录号，邮箱地址或者手机号均可，与user_id字段二选一，在有user_id时，优先使用user_id字段。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 被发放集分宝用户的蚂蚁统一会员ID，与user_account字段二选一。 建议优先使用user_id，接口性能更佳，user_id和user_account都传递时，系统优先使用本字段。 |
