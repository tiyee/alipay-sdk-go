# open/mini/template 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/template`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `template.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_template_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniTemplateMarketingCreate`](#api-minitemplatemarketingcreate) | [`alipay.open.mini.template.marketing.create`](https://opendocs.alipay.com/open-v3/5c89642c_alipay.open.mini.template.marketing.create) | 小程序消息运营位创建 | `POST` |
| [`MiniTemplateUsageQuery`](#api-minitemplateusagequery) | [`alipay.open.mini.template.usage.query`](https://opendocs.alipay.com/open-v3/883d84eb_alipay.open.mini.template.usage.query) | 查询使用模板的小程序列表 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-minitemplatemarketingcreate"></a>

### MiniTemplateMarketingCreate — 小程序消息运营位创建

- 官方接口：[`alipay.open.mini.template.marketing.create`](https://opendocs.alipay.com/open-v3/5c89642c_alipay.open.mini.template.marketing.create)
- 路由：`POST /v3/alipay/open/mini/template/marketing/create`
- 请求类型：[MiniTemplateMarketingCreateRequest](template_request.go#L5)
- 响应类型：[MiniTemplateMarketingCreateResponse](template_response.go#L5)

小程序消息运营位创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | JSON body | 未标注 | 营销活动id |
| `gmt_end` | `GmtEnd` | `*string` | JSON body | 未标注 | 活动结束时间 |
| `gmt_start` | `GmtStart` | `*string` | JSON body | 未标注 | 活动开始时间 |
| `template_ids` | `TemplateIds` | `*[]string` | JSON body | 未标注 | 消息模板id列表，最多50个模板id |
| `title` | `Title` | `*string` | JSON body | 未标注 | 消息运营位名称，不填默认使用券名称 |

<a id="api-minitemplateusagequery"></a>

### MiniTemplateUsageQuery — 查询使用模板的小程序列表

- 官方接口：[`alipay.open.mini.template.usage.query`](https://opendocs.alipay.com/open-v3/883d84eb_alipay.open.mini.template.usage.query)
- 路由：`GET /v3/alipay/open/mini/template/usage/query`
- 请求类型：[MiniTemplateUsageQueryRequest](template_request.go#L19)
- 响应类型：[MiniTemplateUsageQueryResponse](template_response.go#L11)

查询使用模板的小程序列表

请求约束：MiniTemplateUsageQueryRequest is the request type for alipay.open.mini.template.usage.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | query | 否 | 小程序模板APPID |
| `page_num` | `PageNum` | `*string` | query | 否 | 查询的页数，起始为 1（第一页） 。默认第一页。 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页的数量，最多查询50个，默认查询10个 |
| `template_version` | `TemplateVersion` | `*string` | query | 否 | 模板小程序的版本号 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |
