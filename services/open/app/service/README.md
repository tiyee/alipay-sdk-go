# service 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/app/service`。本包提供 5 个支付宝 V3 REST 操作。

通过 `service.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AppServiceApply`](#api-appserviceapply) | `alipay.open.app.service.apply` |
| [`AppServiceDelete`](#api-appservicedelete) | `alipay.open.app.service.delete` |
| [`AppServiceListQuery`](#api-appservicelistquery) | `alipay.open.app.service.list.query` |
| [`AppServiceQuery`](#api-appservicequery) | `alipay.open.app.service.query` |
| [`AppServiceSchemaQuery`](#api-appserviceschemaquery) | `alipay.open.app.service.schema.query` |

## 使用示例

- [AppServiceApply](api_service_appserviceapply_example_test.go)
- [AppServiceDelete](api_service_appservicedelete_example_test.go)
- [AppServiceListQuery](api_service_appservicelistquery_example_test.go)
- [AppServiceQuery](api_service_appservicequery_example_test.go)
- [AppServiceSchemaQuery](api_service_appserviceschemaquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-appserviceapply"></a>

### AppServiceApply — 服务提报申请

- 官方接口：[`alipay.open.app.service.apply`](https://opendocs.alipay.com/open-v3/8b18b82d_alipay.open.app.service.apply)
- 路由：`POST /v3/alipay/open/app/service/apply`
- 请求类型：[AppServiceApplyRequest](service_request.go#L4)
- 响应类型：[AppServiceApplyResponse](service_response.go#L4)

服务提报申请

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `category_id` | `CategoryId` | `*string` | JSON body | 未标注 | 行业类目id,获取请参考[各个行业场景服务接入资料](https://opendocs.alipay.com/mini/03ci0w?pathHash=ed3c875c) |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务编号,平台会基于appId+out_biz_no做幂等控制，如果出现幂等，会返回幂等的service_code |
| `schema_version` | `SchemaVersion` | `*string` | JSON body | 未标注 | 服务schema版本 |
| `service_code` | `ServiceCode` | `*string` | JSON body | 未标注 | 服务编码, 首次提报时设置空值，支付宝侧返回ServiceCode；当传入serviceCode，则对已提报服务做编辑操作。 |
| `service_xml` | `ServiceXml` | `*string` | JSON body | 未标注 | 服务xml,格式请参考[Schema 规则介绍](https://opendocs.alipay.com/mini/03cj40?pathHash=a5de4bd9) |
| `template_type` | `TemplateType` | `*string` | JSON body | 未标注 | 默认值:DEFAULT。 |
| `user_service_delivery_type` | `UserServiceDeliveryType` | `*string` | JSON body | 未标注 | 商户提报服务履约类型 |

<a id="api-appservicedelete"></a>

### AppServiceDelete — 服务失效接口

- 官方接口：[`alipay.open.app.service.delete`](https://opendocs.alipay.com/open-v3/5b7ade20_alipay.open.app.service.delete)
- 路由：`DELETE /v3/alipay/open/app/service/delete`
- 请求类型：[AppServiceDeleteRequest](service_request.go#L22)
- 响应类型：[AppServiceDeleteResponse](service_response.go#L10)

服务失效

请求约束：AppServiceDeleteRequest is the request type for alipay.open.app.service.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_code` | `ServiceCode` | `*string` | query | 否 | 服务id |

<a id="api-appservicelistquery"></a>

### AppServiceListQuery — 服务批量查询

- 官方接口：[`alipay.open.app.service.list.query`](https://opendocs.alipay.com/open-v3/77d4e708_alipay.open.app.service.list.query)
- 路由：`GET /v3/alipay/open/app/service/list/query`
- 请求类型：[AppServiceListQueryRequest](service_request.go#L28)
- 响应类型：[AppServiceListQueryResponse](service_response.go#L14)

服务批量分页查询

请求约束：AppServiceListQueryRequest is the request type for alipay.open.app.service.list.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_name` | `ServiceName` | `*string` | query | 否 | 服务名称 |
| `category_id` | `CategoryId` | `*string` | query | 否 | 类目 |
| `page_size` | `PageSize` | `*int64` | query | 否 | 每页记录数 |
| `page_num` | `PageNum` | `*int64` | query | 否 | 页码 |

<a id="api-appservicequery"></a>

### AppServiceQuery — 服务查询

- 官方接口：[`alipay.open.app.service.query`](https://opendocs.alipay.com/open-v3/86469046_alipay.open.app.service.query)
- 路由：`GET /v3/alipay/open/app/service/query`
- 请求类型：[AppServiceQueryRequest](service_request.go#L40)
- 响应类型：[AppServiceQueryResponse](service_response.go#L26)

服务查询

请求约束：AppServiceQueryRequest is the request type for alipay.open.app.service.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_code` | `ServiceCode` | `*string` | query | 否 | 服务id |

<a id="api-appserviceschemaquery"></a>

### AppServiceSchemaQuery — 服务schema下发

- 官方接口：[`alipay.open.app.service.schema.query`](https://opendocs.alipay.com/open-v3/afe544b5_alipay.open.app.service.schema.query)
- 路由：`GET /v3/alipay/open/app/service/schema/query`
- 请求类型：[AppServiceSchemaQueryRequest](service_request.go#L46)
- 响应类型：[AppServiceSchemaQueryResponse](service_response.go#L42)

服务schema查询下发

请求约束：AppServiceSchemaQueryRequest is the request type for alipay.open.app.service.schema.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `category_id` | `CategoryId` | `*string` | query | 否 | 类目id, 获取请参考[各个行业场景服务接入资料](https://opendocs.alipay.com/mini/03ci0w?pathHash=ed3c875c) |
| `template_type` | `TemplateType` | `*string` | query | 否 | 服务模版类型，默认值:DEFAULT |
