# open/mini/isv 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/isv`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `isv.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_isv_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniISVCreate`](#api-miniisvcreate) | [`alipay.open.mini.isv.create`](https://opendocs.alipay.com/open-v3/719ad502_alipay.open.mini.isv.create) | isv服务商代商户创建小程序 | `POST` |
| [`MiniISVFastRegisterCreate`](#api-miniisvfastregistercreate) | [`alipay.open.mini.isv.fastregister.create`](https://opendocs.alipay.com/open-v3/58d4f210_alipay.open.mini.isv.fastregister.create) | ISV代创建试用小程序 | `POST` |
| [`MiniISVFastRegisterQuery`](#api-miniisvfastregisterquery) | [`alipay.open.mini.isv.fastregister.query`](https://opendocs.alipay.com/open-v3/033e1882_alipay.open.mini.isv.fastregister.query) | 查询试用小程序 | `GET` |
| [`MiniISVQuery`](#api-miniisvquery) | [`alipay.open.mini.isv.query`](https://opendocs.alipay.com/open-v3/f6bafc97_alipay.open.mini.isv.query) | isv查询代商家创建小程序记录 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniisvcreate"></a>

### MiniISVCreate — isv服务商代商户创建小程序

- 官方接口：[`alipay.open.mini.isv.create`](https://opendocs.alipay.com/open-v3/719ad502_alipay.open.mini.isv.create)
- 路由：`POST /v3/alipay/open/mini/isv/create`
- 请求类型：[MiniISVCreateRequest](isv_request.go#L5)
- 响应类型：[MiniISVCreateResponse](isv_response.go#L5)

服务商代商家创建小程序，待商家确认后服务商代创建的小程序才会生效，若商家超过确认时间则服务商创建的小程序会自动作废释放占用的小程序名称。
将 is_individual 设置为 true，并且同时传入 license_pic。将同时对具备个体工商户营业执照的商家个人账号认证成为个体工商户账号（创建小程序+升级个体工商户账号一体化）。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `create_mini_request` | `CreateMiniRequest` | [`*CreateMiniRequest`](isv_request.go#L22) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-miniisvfastregistercreate"></a>

### MiniISVFastRegisterCreate — ISV代创建试用小程序

- 官方接口：[`alipay.open.mini.isv.fastregister.create`](https://opendocs.alipay.com/open-v3/58d4f210_alipay.open.mini.isv.fastregister.create)
- 路由：`POST /v3/alipay/open/mini/isv/fastregister/create`
- 请求类型：[MiniISVFastRegisterCreateRequest](isv_request.go#L10)
- 响应类型：[MiniISVFastRegisterCreateResponse](isv_response.go#L11)

ISV代商户创建试用小程序

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_name` | `AppName` | `*string` | JSON body | 未标注 | 小程序名称，昵称半自动设定，强制后缀“的试用小程序”。且该参数会进行关键字检查，如果命中品牌关键字则会报错。 如遇到品牌大客户要用试用小程序，建议用户先换个名字，认证后再修改成品牌名。 |
| `auth_notify_url` | `AuthNotifyUrl` | `*string` | JSON body | 未标注 | 授权确认成功后回调h5链接或小程序页面 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 开发者外部订单号，通过开发者账号+外部订单号做业务幂等 |
| `uid` | `Uid` | `*string` | JSON body | 未标注 | 商户uid |

<a id="api-miniisvfastregisterquery"></a>

### MiniISVFastRegisterQuery — 查询试用小程序

- 官方接口：[`alipay.open.mini.isv.fastregister.query`](https://opendocs.alipay.com/open-v3/033e1882_alipay.open.mini.isv.fastregister.query)
- 路由：`GET /v3/alipay/open/mini/isv/fastregister/query`
- 请求类型：[MiniISVFastRegisterQueryRequest](isv_request.go#L46)
- 响应类型：[MiniISVFastRegisterQueryResponse](isv_response.go#L19)

查询试用小程序

请求约束：MiniISVFastRegisterQueryRequest is the request type for alipay.open.mini.isv.fastregister.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_no` | `OrderNo` | `*string` | query | 否 | 代创建试用小程序单号 |

<a id="api-miniisvquery"></a>

### MiniISVQuery — isv查询代商家创建小程序记录

- 官方接口：[`alipay.open.mini.isv.query`](https://opendocs.alipay.com/open-v3/f6bafc97_alipay.open.mini.isv.query)
- 路由：`GET /v3/alipay/open/mini/isv/query`
- 请求类型：[MiniISVQueryRequest](isv_request.go#L52)
- 响应类型：[MiniISVQueryResponse](isv_response.go#L35)

isv查询代商家创建小程序记录

请求约束：MiniISVQueryRequest is the request type for alipay.open.mini.isv.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_no` | `OrderNo` | `*string` | query | 否 | 服务商代商家创建小程序，由支付宝开放平台返回的订单号用于后续查询处理进展。 |
