# twostage 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/user/twostage`。本包提供 2 个支付宝 V3 REST 操作。

通过 `twostage.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`TwoStageCommonUse`](#api-twostagecommonuse) | `alipay.user.twostage.common.use` |
| [`TwoStageIndirectUse`](#api-twostageindirectuse) | `alipay.user.twostage.indirect.use` |

## 使用示例

- [TwoStageCommonUse](api_twostage_twostagecommonuse_example_test.go)
- [TwoStageIndirectUse](api_twostage_twostageindirectuse_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-twostagecommonuse"></a>

### TwoStageCommonUse — 通用当面付二阶段接口

- 官方接口：[`alipay.user.twostage.common.use`](https://opendocs.alipay.com/apis/05xim7)
- 路由：`POST /v3/alipay/user/twostage/common/use`
- 请求类型：[UserTwoStageCommonUseRequest](twostage_request.go#L4)
- 响应类型：[UserTwoStageCommonUseResponse](twostage_response.go#L4)

商家通过小程序/卡券开放权益(优惠券)，用户在小程序/卡券授权后，打开付款码时通过该接口获取用户信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `dynamic_id` | `DynamicId` | `*string` | JSON body | 未标注 | 商户扫描用户的付款码值。 |
| `pay_check_strategy` | `PayCheckStrategy` | `*string` | JSON body | 未标注 | 传1表示 校验pid(pay_pid)和来支付的时候的pid一致性； 传2表示校验pid(pay_pid)和来支付的时候的pid一致性、校验scene_no和来支付时DYNAMIC_TOKEN_OUT_BIZ_NO属性一致； 不传值和value=2的表现一致 |
| `pay_pid` | `PayPid` | `*string` | JSON body | 未标注 | 商家进行二阶段支付的PID信息。 |
| `sence_no` | `SenceNo` | `*string` | JSON body | 未标注 | 外部业务号，用于标识这笔解码请求，对同一个码的重复解码请求，sence_no必须与上一次保持一致，每次请求的sence_no必须不一样，如alipay.user.twostage.common.use接口配合alipay.trade.pay（统一收单交易支付接口）一并使用时，alipay.trade.pay接口的extend_params属性中必须设置DYNAMIC_TOKEN_OUT_BIZ_NO，且值必须与sence_no保持一致。 |

<a id="api-twostageindirectuse"></a>

### TwoStageIndirectUse — 双联二阶段

- 官方接口：[`alipay.user.twostage.indirect.use`](https://opendocs.alipay.com/pre-apis/05x9jk)
- 路由：`POST /v3/alipay/user/twostage/indirect/use`
- 请求类型：[UserTwoStageIndirectUseRequest](twostage_request.go#L16)
- 响应类型：[UserTwoStageIndirectUseResponse](twostage_response.go#L13)

间连模式通过双联向银行输出付款码查UID接口.
双联二阶段里，授权之后调用该接口，进行解码并校验双联pid、银行orgId、isv sourceId及学校smid之间的级联关系；检查授权关系；解码并返回用户uid

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `dynamic_id` | `DynamicId` | `*string` | JSON body | 未标注 | 商户扫描用户的付款码值。18~24位，25~30开头，例如28开头的18位的数字；或人脸支付的ftoken等。 |
| `org_pid` | `OrgPid` | `*string` | JSON body | 未标注 | 进件信息中，SMID对应的银行机构的PID信息，一般为2088开头的16位数字。 |
| `pay_smid` | `PaySmid` | `*string` | JSON body | 未标注 | 进件信息中，二级商户ID（ sub_merchant_id)信息，一般为2088开头的16位数字。 |
| `sence_no` | `SenceNo` | `*string` | JSON body | 未标注 | 外部业务号，用于标识这笔解码请求，对同一个码的重复解码请求，sence_no必须与上一次保持一致，每次请求的sence_no必须不一样，如alipay.user.twostage.common.use接口配合alipay.trade.pay（统一收单交易支付接口）一并使用时，alipay.trade.pay接口的extend_params属性中必须设置DYNAMIC_TOKEN_OUT_BIZ_NO，且值必须与sence_no保持一致 |
| `source_pid` | `SourcePid` | `*string` | JSON body | 未标注 | 进件信息中，SMID对应渠道的PID信息，一般为2088开头的16位数字。 |
