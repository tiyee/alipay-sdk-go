# royalty 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/royalty`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `royalty.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_royalty_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`RoyaltyRateQuery`](#api-royaltyratequery) | [`alipay.trade.royalty.rate.query`](https://opendocs.alipay.com/open-v3/d9aece9f_alipay.trade.royalty.rate.query) | 分账比例查询 | `GET` |
| [`RoyaltyRelationBatchQuery`](#api-royaltyrelationbatchquery) | [`alipay.trade.royalty.relation.batchquery`](https://opendocs.alipay.com/open-v3/1a11475c_alipay.trade.royalty.relation.batchquery) | 分账关系查询 | `POST` |
| [`RoyaltyRelationBind`](#api-royaltyrelationbind) | [`alipay.trade.royalty.relation.bind`](https://opendocs.alipay.com/open-v3/e701637d_alipay.trade.royalty.relation.bind) | 分账关系绑定 | `POST` |
| [`RoyaltyRelationUnbind`](#api-royaltyrelationunbind) | [`alipay.trade.royalty.relation.unbind`](https://opendocs.alipay.com/open-v3/302cd1f8_alipay.trade.royalty.relation.unbind) | 分账关系解绑 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-royaltyratequery"></a>

### RoyaltyRateQuery — 分账比例查询

- 官方接口：[`alipay.trade.royalty.rate.query`](https://opendocs.alipay.com/open-v3/d9aece9f_alipay.trade.royalty.rate.query)
- 路由：`GET /v3/alipay/trade/royalty/rate/query`
- 请求类型：[RoyaltyRateQueryRequest](royalty_request.go#L9)
- 响应类型：[RoyaltyRateQueryResponse](royalty_response.go#L8)

商户签约分账产品后，商户或其授权ISV可对分账比例进行查询。

请求约束：RoyaltyRateQueryRequest 是 alipay.trade.royalty.rate.query
（分账比例查询，V3 GET）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | query | 否 | 外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复。<br>SDK 字段说明：外部请求号：32 字符以内，仅可包含字母、数字、<br>下划线，需保证在商户端不重复。 |

<a id="api-royaltyrelationbatchquery"></a>

### RoyaltyRelationBatchQuery — 分账关系查询

- 官方接口：[`alipay.trade.royalty.relation.batchquery`](https://opendocs.alipay.com/open-v3/1a11475c_alipay.trade.royalty.relation.batchquery)
- 路由：`POST /v3/alipay/trade/royalty/relation/batchquery`
- 请求类型：[RoyaltyRelationBatchQueryRequest](royalty_request.go#L36)
- 响应类型：[RoyaltyRelationBatchQueryResponse](royalty_response.go#L28)

当商户签约分账产品后，授权ISV帮其进行分账关系的维护。本接口用于商户与分账方的关系查询。

请求约束：RoyaltyRelationBatchQueryRequest 是 alipay.trade.royalty.relation.batchquery
（分账关系查询）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复。 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 几页，起始页为 1。不填默认为 1。<br>SDK 字段说明：页码，起始页为 1；不填默认为 1。 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 页面大小。每页记录数，取值范围是(0,100]。不填默认为20<br>SDK 字段说明：每页记录数，取值范围 (0,100]；不填默认为 20。 |

<a id="api-royaltyrelationbind"></a>

### RoyaltyRelationBind — 分账关系绑定

- 官方接口：[`alipay.trade.royalty.relation.bind`](https://opendocs.alipay.com/open-v3/e701637d_alipay.trade.royalty.relation.bind)
- 路由：`POST /v3/alipay/trade/royalty/relation/bind`
- 请求类型：[RoyaltyRelationBindRequest](royalty_request.go#L17)
- 响应类型：[RoyaltyRelationBindResponse](royalty_response.go#L16)

当商户签约分账产品后，授权ISV帮其进行分账关系的维护。本接口用于商户与分账方的关系绑定。

请求约束：RoyaltyRelationBindRequest 是 alipay.trade.royalty.relation.bind
（分账关系绑定）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复。<br>SDK 字段说明：外部请求号：32 字符以内，仅可包含字母、数字、<br>下划线，需保证在商户端不重复。 |
| `receiver_list` | `ReceiverList` | [`[]RoyaltyEntity`](#model-royaltyentity) | JSON body | 未标注 | 分账接收方列表，单次传入最多20个<br>SDK 字段说明：分账接收方列表，单次最多 20 个。 |

注意事项：

- 绑定结果以 ResultCode 为准。

<a id="api-royaltyrelationunbind"></a>

### RoyaltyRelationUnbind — 分账关系解绑

- 官方接口：[`alipay.trade.royalty.relation.unbind`](https://opendocs.alipay.com/open-v3/302cd1f8_alipay.trade.royalty.relation.unbind)
- 路由：`POST /v3/alipay/trade/royalty/relation/unbind`
- 请求类型：[RoyaltyRelationUnbindRequest](royalty_request.go#L27)
- 响应类型：[RoyaltyRelationUnbindResponse](royalty_response.go#L22)

当商户签约分账产品后，授权ISV帮其进行分账关系的维护。本接口用于商户与分账方的关系解绑。

请求约束：RoyaltyRelationUnbindRequest 是 alipay.trade.royalty.relation.unbind
（分账关系解绑）的请求。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复。 |
| `receiver_list` | `ReceiverList` | [`[]RoyaltyEntity`](#model-royaltyentity) | JSON body | 未标注 | 分账接收方列表，单次传入最多 20 个信息。<br>SDK 字段说明：分账接收方列表，单次最多 20 个。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-royaltyentity"></a>

### RoyaltyEntity

类型定义：[RoyaltyEntity](royalty_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `type` | `Type` | [`RoyaltyEntityType`](royalty_model.go#L4) | 未标注 | 分账接收方类型。 |
| `account` | `Account` | `*string` | 未标注 | 分账接收方账号。 当分账方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字； 当分账方类型是loginName时，本参数为用户的支付宝登录号；当分账方类型是openId时，本参数传递支付宝openId信息。<br>SDK 字段说明：分账接收方账号：type=userId 时为 2088 开头的 16 位<br>用户号，type=loginName 时为登录号，type=openId 时为 openId。 |
| `account_open_id` | `AccountOpenID` | `*string` | 未标注 | 分账接收方openId，本参数为分账接收方在该应用（AppId）下的唯一用户标识，仅用于分账关系查询接口出参。<br>SDK 字段说明：分账接收方在该应用下的 OpenId；仅查询接口返回。 |
| `name` | `Name` | `*string` | 未标注 | 分账接收方真实姓名。 绑定分账关系时： 当分账方类型是userId时，本参数可以不传，若上传则进行校验不上传不会校验。 当分账方类型是loginName时，本参数必传。 解绑分账关系时：作为请求参数可不填，分账关系查询时不作为返回结果返回<br>SDK 字段说明：分账接收方真实姓名；type=loginName 绑定时必传。 |
| `memo` | `Memo` | `*string` | 未标注 | 分账关系描述 |
| `bind_login_name` | `BindLoginName` | `*string` | 未标注 | 作为查询返回结果：分账收款方绑定时的支付宝登录号。分账关系绑定（alipay.trade.royalty.relation.bind）时，通过type为loginName绑定传入的支付宝登录号，若使用userId绑定则不返回。<br>SDK 字段说明：绑定时的登录号；仅查询接口返回。 |
| `login_name` | `LoginName` | `*string` | 未标注 | 作为查询返回结果：当前userId对应的支付宝登录号。当login_name与bind_login_name不相等时，表明该支付宝账户发生了登录号变更。<br>SDK 字段说明：当前 userId 对应的登录号；仅查询接口返回。 |
