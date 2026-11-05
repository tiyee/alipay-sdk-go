# eco/signflows 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/signflows`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `signflows.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_signflows_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`SignFlowsDetailQuery`](#api-signflowsdetailquery) | [`alipay.eco.signflows.detail.query`](https://opendocs.alipay.com/pre-apis/05wylm) | 获取签署流程合同与附件的下载地址（E签宝） | `GET` |
| [`SignFlowsURLQuery`](#api-signflowsurlquery) | [`alipay.eco.signflows.url.query`](https://opendocs.alipay.com/pre-apis/05wyc8) | 获取签署地址 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-signflowsdetailquery"></a>

### SignFlowsDetailQuery — 获取签署流程合同与附件的下载地址（E签宝）

- 官方接口：[`alipay.eco.signflows.detail.query`](https://opendocs.alipay.com/pre-apis/05wylm)
- 路由：`GET /v3/alipay/eco/signflows/detail/query`
- 请求类型：[SignFlowsDetailQueryRequest](signflows_request.go#L5)
- 响应类型：[SignFlowsDetailQueryResponse](signflows_response.go#L5)

获取签署流程合同与附件的下载地址（E签宝）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `flow_id` | `FlowId` | `*string` | query | 否 | 流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。 |

<a id="api-signflowsurlquery"></a>

### SignFlowsURLQuery — 获取签署地址

- 官方接口：[`alipay.eco.signflows.url.query`](https://opendocs.alipay.com/pre-apis/05wyc8)
- 路由：`GET /v3/alipay/eco/signflows/url/query`
- 请求类型：[SignFlowsURLQueryRequest](signflows_request.go#L13)
- 响应类型：[SignFlowsURLQueryResponse](signflows_response.go#L11)

创建流程后，获取指定签署人的签署链接地址，可在应用内集成H5签署页，或者通过短信发送签署链接。
传入个人唯一标识，则获取的签署任务链接仅包含个人人的签署任务；如同时传入企业唯一标识，则获取的签署任务链接包含企业与个人的签署任务。
预览链接：支持签署人先查看合同原文，后进行登录并完成签署。适用于应用内集成场景。
签署链接：签署人需要登录后查看合同原文并签署。适用用短信发送场景。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `flow_id` | `FlowId` | `*string` | query | 否 | 流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。 |
| `third_party_user_id` | `ThirdPartyUserId` | `*string` | query | 否 | 创建流程时指定个人唯一标识 |
| `org_third_party_user_id` | `OrgThirdPartyUserId` | `*string` | query | 否 | 创建流程时指定企业唯一标识 |
