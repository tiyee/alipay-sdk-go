# agent 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/agent`。本包提供 5 个支付宝 V3 REST 操作。

通过 `agent.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`sign`](sign/README.md) | 6 | `AgentCommonSign` / `AgentCommonSignConfirm` / `AgentFaceToFaceSign` / `AgentMobilePaySign` / `AgentOfflinePaymentSign` / `AgentSignStatusQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AgentCancel`](#api-agentcancel) | `alipay.open.agent.cancel` |
| [`AgentConfirm`](#api-agentconfirm) | `alipay.open.agent.confirm` |
| [`AgentCreate`](#api-agentcreate) | `alipay.open.agent.create` |
| [`AgentMiniCreate`](#api-agentminicreate) | `alipay.open.agent.mini.create` |
| [`AgentOrderQuery`](#api-agentorderquery) | `alipay.open.agent.order.query` |

## 使用示例

- [AgentCancel](api_agent_agentcancel_example_test.go)
- [AgentConfirm](api_agent_agentconfirm_example_test.go)
- [AgentCreate](api_agent_agentcreate_example_test.go)
- [AgentMiniCreate](api_agent_agentminicreate_example_test.go)
- [AgentOrderQuery](api_agent_agentorderquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-agentcancel"></a>

### AgentCancel — 取消代商户签约、创建应用事务

- 官方接口：[`alipay.open.agent.cancel`](https://opendocs.alipay.com/open-v3/8e66406f_alipay.open.agent.cancel)
- 路由：`POST /v3/alipay/open/agent/cancel`
- 请求类型：[AgentCancelRequest](agent_request.go#L8)
- 响应类型：[AgentCancelResponse](agent_response.go#L4)

取消代商户签约、创建应用事务

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | JSON body | 未标注 | ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。 |

<a id="api-agentconfirm"></a>

### AgentConfirm — 提交代商户签约、创建应用事务

- 官方接口：[`alipay.open.agent.confirm`](https://opendocs.alipay.com/open-v3/e853dfa1_alipay.open.agent.confirm)
- 路由：`POST /v3/alipay/open/agent/confirm`
- 请求类型：[AgentConfirmRequest](agent_request.go#L14)
- 响应类型：[AgentConfirmResponse](agent_response.go#L8)

提交事务。只允许提交init状态的事务，submit|cancel|timeout 状态的都是终态，不允许提交，且不允许提交空事务，需要先调用代创建小程序、代签约当面付等业务接口，再提交事务。服务市场订购及授权，使用订单授权凭证order_ticket开启的事务，提交后会有应用授权令牌返回。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | JSON body | 未标注 | ISV 代商户操作事务编号，通过调用alipay.open.agent.create(开启代商户签约、创建应用事务)接口返回，详见 https://opendocs.alipay.com/apis/api_50/alipay.open.agent.create/ 。 |

<a id="api-agentcreate"></a>

### AgentCreate — 开启代商户签约、创建应用事务

- 官方接口：[`alipay.open.agent.create`](https://opendocs.alipay.com/open-v3/b4c46445_alipay.open.agent.create)
- 路由：`POST /v3/alipay/open/agent/create`
- 请求类型：[AgentCreateRequest](agent_request.go#L20)
- 响应类型：[AgentCreateResponse](agent_response.go#L28)

在 ISV 代商户进行应用创建、产品签约时，用于开启一个操作事务，必须是第一个调用的接口。场景1：ISV 代商户进行应用创建、产品签约，最后提交事务后需要商户确认才能完成流程；场景2：服务市场订购及授权，使用订单授权凭证order_ticket开启预授权模式，该模式下提交事务后无需商户确认。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account` | `Account` | `*string` | JSON body | 未标注 | isv代操作的商户账号，可以是支付宝账号，也可以是pid（2088开头） |
| `contact_info` | `ContactInfo` | [`*Contact`](#model-contact) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `order_ticket` | `OrderTicket` | `*string` | JSON body | 未标注 | 订单授权凭证。若传入本参数，则对应事务提交后进入预授权模式。 |
| `sign_mode` | `SignMode` | `*string` | JSON body | 未标注 | 签约模式：目前只支持批量签约并不通知商户确认 使用场景：批量签约并授权场景，先调用签约接口进行签约，然后调用授权接口授权，在批量模式下，一次性签约多个产品，要么都成功要么都失败，同时在授权接口传入orderNo，会一起发送签约确认和授权确认，商户只需要确认一次；如果不是批量签约并授权场景，不要传此参数，否则商户无法收到确认签约通知； |

<a id="api-agentminicreate"></a>

### AgentMiniCreate — 代商家创建小程序应用

- 官方接口：[`alipay.open.agent.mini.create`](https://opendocs.alipay.com/open-v3/e39a93e8_alipay.open.agent.mini.create)
- 路由：`POST /v3/alipay/open/agent/mini/create`
- 请求类型：[AgentMiniCreateRequest](agent_request.go#L31)
- 响应类型：[AgentMiniCreateResponse](agent_response.go#L36)

ISV 代商家创建小程序应用。该接口后续将不再更新迭代，请使用新接口对接开发。新接口介绍文档：https://opendocs.alipay.com/isv/03kqzh ；新接口对接文档：https://opendocs.alipay.com/mini/03l21u 。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*AgentMiniCreate`](#model-agentminicreate) | multipart | 未标注 | 见字段类型及官方接口文档。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-agentorderquery"></a>

### AgentOrderQuery — 查询申请单状态

- 官方接口：[`alipay.open.agent.order.query`](https://opendocs.alipay.com/open-v3/ceae4d66_alipay.open.agent.order.query)
- 路由：`GET /v3/alipay/open/agent/order/query`
- 请求类型：[AgentOrderQueryRequest](agent_request.go#L37)
- 响应类型：[AgentOrderQueryResponse](agent_response.go#L40)

代签约产品提交后，查询对应申请单的情况

请求约束：AgentOrderQueryRequest is the request type for alipay.open.agent.order.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | query | 否 | ISV 代商户操作事务编号，通过调用alipay.open.agent.create(开启代商户签约、创建应用事务)接口返回，详见 https://opendocs.alipay.com/apis/api_50/alipay.open.agent.create/ 。 |

## 嵌套请求参数

<a id="model-agentminicreate"></a>

### AgentMiniCreate

类型定义：[AgentMiniCreate](agent_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | 未标注 | 小程序类目，选填，可以支持最多三个类目，最少一个，参数格式：一级类目1_二级类目1;一级类目2_二级类目2。 类目可以通过以下接口查询：alipay.open.mini.category.query |
| `app_desc` | `AppDesc` | `*string` | 未标注 | 商家小程序描述信息，简要描述小程序主要功能（20-200个字），例：xx小程序提供了xx功能，主要解决了XX问题。 |
| `app_english_name` | `AppEnglishName` | `*string` | 未标注 | 小程序英文名称，长度3~20个字符 |
| `app_name` | `AppName` | `*string` | 未标注 | 代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间，一个中文字等于2个字符，更多名称规则见：https://docs.alipay.com/mini/operation/name |
| `app_slogan` | `AppSlogan` | `*string` | 未标注 | 代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符） |
| `batch_no` | `BatchNo` | `*string` | 未标注 | ISV 代商家操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。 |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list。使用后不再读取app_category_ids值，老前台类目将废弃 |
| `service_email` | `ServiceEmail` | `*string` | 未标注 | 商家小程序客服邮箱<br> 商家小程序客服电话和邮箱，可以二选一填写，但不能同时为空 |
| `service_phone` | `ServicePhone` | `*string` | 未标注 | 商家小程序的客服电话，推荐填写<br> 商家小程序客服电话和邮箱，可以二选一填写，但不能同时为空 |

<a id="model-contact"></a>

### Contact

类型定义：[Contact](agent_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `contact_email` | `ContactEmail` | `*string` | 未标注 | 联系人邮箱 |
| `contact_mobile` | `ContactMobile` | `*string` | 未标注 | 联系人手机号码 |
| `contact_name` | `ContactName` | `*string` | 未标注 | 联系人名称 |
