# app 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/app`。本包提供 3 个支付宝 V3 REST 操作。

通过 `app.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`api`](api/README.md) | 4 | `AppAPIFieldApply` / `AppAPIFieldQuery` / `AppAPIQuery` / `AppAPISceneQuery` |
| [`members`](members/README.md) | 3 | `AppMembersCreate` / `AppMembersDelete` / `AppMembersQuery` |
| [`service`](service/README.md) | 5 | `AppServiceApply` / `AppServiceDelete` / `AppServiceListQuery` / `AppServiceQuery` / `AppServiceSchemaQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AppMessageTemplateSubscribeQuery`](#api-appmessagetemplatesubscribequery) | `alipay.open.app.messagetemplate.subscribe.query` |
| [`AppMiniTemplateMessageSend`](#api-appminitemplatemessagesend) | `alipay.open.app.mini.templatemessage.send` |
| [`AppQRCodeCreate`](#api-appqrcodecreate) | `alipay.open.app.qrcode.create` |

## 使用示例

- [AppMessageTemplateSubscribeQuery](api_app_appmessagetemplatesubscribequery_example_test.go)
- [AppMiniTemplateMessageSend](api_app_appminitemplatemessagesend_example_test.go)
- [AppQRCodeCreate](api_app_appqrcodecreate_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-appmessagetemplatesubscribequery"></a>

### AppMessageTemplateSubscribeQuery — 模板订阅关系查询

- 官方接口：[`alipay.open.app.messagetemplate.subscribe.query`](https://opendocs.alipay.com/open-v3/af42509c_alipay.open.app.messagetemplate.subscribe.query)
- 路由：`POST /v3/alipay/open/app/messagetemplate/subscribe/query`
- 请求类型：[AppMessageTemplateSubscribeQueryRequest](app_request.go#L4)
- 响应类型：[AppMessageTemplateSubscribeQueryResponse](app_response.go#L4)

商户使用该接口可以查询用户对其消息模板的订阅关系，可以根据是否订阅引导用户进行订阅和精准触达消息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝openId，用于支付宝用户在当前应用下的用户标识。 |
| `template_id_list` | `TemplateIdList` | `*[]string` | JSON body | 未标注 | 消息模板id，可以填写多个，最多不超过3个。模板id需要保持同一个应用主体，并且展示在同一个订阅组件中的模板id。 模板id获取详情参见[模板消息](https://opendocs.alipay.com/mini/01rnqx)。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 订阅消息模板用户的支付宝唯一标识，2088开头。 |

<a id="api-appminitemplatemessagesend"></a>

### AppMiniTemplateMessageSend — 小程序发送模板消息

- 官方接口：[`alipay.open.app.mini.templatemessage.send`](https://opendocs.alipay.com/open-v3/0136e987_alipay.open.app.mini.templatemessage.send)
- 路由：`POST /v3/alipay/open/app/mini/templatemessage/send`
- 请求类型：[AppMiniTemplateMessageSendRequest](app_request.go#L14)
- 响应类型：[AppMiniTemplateMessageSendResponse](app_response.go#L12)

小程序通过openapi给用户触达消息，主要为支付后的触达（通过消费id）、用户提交表单后的触达（通过formId）、刷脸后的触达（通过ftoken）以及用户订阅消息模板后的触达。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | `*string` | JSON body | 未标注 | 模板消息内容。商家/开发者将模板占位符替换为自定义内容 |
| `form_id` | `FormId` | `*string` | JSON body | 未标注 | 支付消息模板：需传入用户发生的交易行为的支付宝交易号 trade_no；表单提交模板：需传入用户在小程序触发表单提交事件获得的表单号；刷脸消息模板：需传入在IOT刷脸后得到的ftoken等，用于信息发送的校验。 |
| `page` | `Page` | `*string` | JSON body | 未标注 | 小程序的跳转页面。用于用户点击模板消息 进入小程序查看 按钮后，跳转至商家小程序对应页面。 |
| `to_open_id` | `ToOpenId` | `*string` | JSON body | 未标注 | 支付宝openId，用于支付宝用户在当前应用下的用户标识。 |
| `to_user_id` | `ToUserId` | `*string` | JSON body | 未标注 | 接收模板消息用户支付宝 user_id，可通过 [用户授权](https://opendocs.alipay.com/mini/api/openapi-authorize) 获取 |
| `user_template_id` | `UserTemplateId` | `*string` | JSON body | 未标注 | 商家在商家自运营中心选用的消息模板ID，详情参见 [选用消息模板](https://opendocs.alipay.com/mini/03l9bb?pathHash=19d2e0aa&ref=api#%E9%A2%86%E7%94%A8%E8%AE%A2%E9%98%85%E6%B6%88%E6%81%AF%E6%A8%A1%E6%9D%BF) 。 |

<a id="api-appqrcodecreate"></a>

### AppQRCodeCreate — 小程序生成推广二维码接口

- 官方接口：[`alipay.open.app.qrcode.create`](https://opendocs.alipay.com/open-v3/76c2627e_alipay.open.app.qrcode.create)
- 路由：`POST /v3/alipay/open/app/qrcode/create`
- 请求类型：[AppQRCodeCreateRequest](app_request.go#L30)
- 响应类型：[AppQRCodeCreateResponse](app_response.go#L16)

生成小程序推广二维码

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `color` | `Color` | `*string` | JSON body | 未标注 | 圆形二维码颜色（十六进制颜色色值），仅圆形二维码支持颜色设置，方形二维码默认为黑色。 |
| `describe` | `Describe` | `*string` | JSON body | 未标注 | 对应的二维码描述 |
| `query_param` | `QueryParam` | `*string` | JSON body | 未标注 | 小程序启动参数，内容按照格式为参数名=参数值&参数名=参数值。打开小程序的query ，在小程序 onLaunch的方法中获取。可查看 [如何获取各种场景的启动参数](https://opendocs.alipay.com/support/01rb2a)。 |
| `size` | `Size` | `*string` | JSON body | 未标注 | 合成后图片的大小规格，有s、m、l三档可选。 |
| `url_param` | `UrlParam` | `*string` | JSON body | 未标注 | page/component/component-pages/view/view为小程序中能访问到的页面路径 |
