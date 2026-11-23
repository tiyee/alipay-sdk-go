# open/public/template 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/template`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `template.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_template_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicTemplateMessageGet`](#api-publictemplatemessageget) | [`alipay.open.public.template.message.get`](https://opendocs.alipay.com/apis/05wbw6) | 消息模板领取接口 | `GET` |
| [`PublicTemplateMessageIndustryModify`](#api-publictemplatemessageindustrymodify) | [`alipay.open.public.template.message.industry.modify`](https://opendocs.alipay.com/apis/05wuhh) | 模板消息行业设置修改接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publictemplatemessageget"></a>

### PublicTemplateMessageGet — 消息模板领取接口

- 官方接口：[`alipay.open.public.template.message.get`](https://opendocs.alipay.com/apis/05wbw6)
- 路由：`GET /v3/alipay/open/public/template/message/get`
- 请求类型：[PublicTemplateMessageGetRequest](template_request.go#L17)
- 响应类型：[PublicTemplateMessageGetResponse](template_response.go#L5)

帮助服务窗开发者从服务窗平台（fuwu.alipay.com）公共模板库里领取一个行业的消息模板--消息模板是一种消息的样式，如消费提醒、账单提醒等，主要用于isv帮助商户领取消息模板。

请求约束：PublicTemplateMessageGetRequest is the request type for alipay.open.public.template.message.get.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | query | 否 | 消息母板id，登陆生活号后台(fuwu.alipay.com)，点击菜单“模板消息”，点击“模板库”，即可看到相应模板的消息母板id |

<a id="api-publictemplatemessageindustrymodify"></a>

### PublicTemplateMessageIndustryModify — 模板消息行业设置修改接口

- 官方接口：[`alipay.open.public.template.message.industry.modify`](https://opendocs.alipay.com/apis/05wuhh)
- 路由：`POST /v3/alipay/open/public/template/message/industry/modify`
- 请求类型：[PublicTemplateMessageIndustryModifyRequest](template_request.go#L5)
- 响应类型：[PublicTemplateMessageIndustryModifyResponse](template_response.go#L13)

用于修改服务窗消息模板设置的行业信息，如：IT科技/IT软件与服务。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `primary_industry_code` | `PrimaryIndustryCode` | `*string` | JSON body | 未标注 | 服务窗消息模板所属主行业一/二级编码，参见 [查看行业信息](https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043)。 |
| `primary_industry_name` | `PrimaryIndustryName` | `*string` | JSON body | 未标注 | 服务窗消息模板所属主行业一/二级名称，参见 [查看行业信息](https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043)。 |
| `secondary_industry_code` | `SecondaryIndustryCode` | `*string` | JSON body | 未标注 | 服务窗消息模板所属副行业一/二级编码，参见 [查看行业信息](https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043)。 |
| `secondary_industry_name` | `SecondaryIndustryName` | `*string` | JSON body | 未标注 | 服务窗消息模板所属副行业一/二级名称，参见 [查看行业信息](https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043)。 |
