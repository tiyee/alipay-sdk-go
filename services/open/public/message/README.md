# open/public/message 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/message`。本目录提供 7 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `message.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicMessageContentCreate`](#api-publicmessagecontentcreate) | [`alipay.open.public.message.content.create`](https://opendocs.alipay.com/apis/05xhmr) | 创建图文消息内容接口 | `POST` |
| [`PublicMessageContentModify`](#api-publicmessagecontentmodify) | [`alipay.open.public.message.content.modify`](https://opendocs.alipay.com/apis/05wuj5) | 更新图文消息内容接口 | `POST` |
| [`PublicMessageCustomSend`](#api-publicmessagecustomsend) | [`alipay.open.public.message.custom.send`](https://opendocs.alipay.com/apis/05wcpx) | 异步单发消息 | `POST` |
| [`PublicMessageGroupSend`](#api-publicmessagegroupsend) | [`alipay.open.public.message.group.send`](https://opendocs.alipay.com/apis/05wyfq) | 分组消息发送接口 | `POST` |
| [`PublicMessageLabelSend`](#api-publicmessagelabelsend) | [`alipay.open.public.message.label.send`](https://opendocs.alipay.com/pre-apis/05xlk0) | 根据标签组发消息接口 | `POST` |
| [`PublicMessageSingleSend`](#api-publicmessagesinglesend) | [`alipay.open.public.message.single.send`](https://opendocs.alipay.com/apis/06bpqn) | 单发模板消息 | `POST` |
| [`PublicMessageTotalSend`](#api-publicmessagetotalsend) | [`alipay.open.public.message.total.send`](https://opendocs.alipay.com/apis/05wnte) | 群发消息 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicmessagecontentcreate"></a>

### PublicMessageContentCreate — 创建图文消息内容接口

- 官方接口：[`alipay.open.public.message.content.create`](https://opendocs.alipay.com/apis/05xhmr)
- 路由：`POST /v3/alipay/open/public/message/content/create`
- 请求类型：[PublicMessageContentCreateRequest](message_request.go#L5)
- 响应类型：[PublicMessageContentCreateResponse](message_response.go#L5)

开发者可使用该接口将富文本内容生成生活号图文消息内容链接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `benefit` | `Benefit` | `*string` | JSON body | 未标注 | 活动利益点，最多10个字符。仅 ctype 为 activity 类型时才需传入。 |
| `content` | `Content` | `*string` | JSON body | 未标注 | 消息正文（支持富文本） |
| `could_comment` | `CouldComment` | `*string` | JSON body | 未标注 | 是否允许评论。枚举支持： *T：允许。 *F：不允许，默认不允许。 |
| `cover` | `Cover` | `*string` | JSON body | 未标注 | 封面图url，需传入 [图片上传接口](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)上传图片获取的 image_url。 注意：尺寸为 996*450，最大不超过3M，支持格式：jpg、.png 。 |
| `ctype` | `Ctype` | `*string` | JSON body | 未标注 | 图文类型，不填默认普通图文。还支持 activity（活动图文）。 |
| `ext_tags` | `ExtTags` | `*string` | JSON body | 未标注 | 关键词列表，英文逗号分隔，最多不超过5个 |
| `login_ids` | `LoginIds` | `*string` | JSON body | 未标注 | 可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 标题 |

<a id="api-publicmessagecontentmodify"></a>

### PublicMessageContentModify — 更新图文消息内容接口

- 官方接口：[`alipay.open.public.message.content.modify`](https://opendocs.alipay.com/apis/05wuj5)
- 路由：`POST /v3/alipay/open/public/message/content/modify`
- 请求类型：[PublicMessageContentModifyRequest](message_request.go#L25)
- 响应类型：[PublicMessageContentModifyResponse](message_response.go#L13)

开发者可使用该接口修改已经创建的生活号图文消息素材里的内容

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `benefit` | `Benefit` | `*string` | JSON body | 未标注 | 活动利益点，图文类型ctype为activity类型时才需要传，最多10个字符 |
| `content` | `Content` | `*string` | JSON body | 未标注 | 消息正文（支持富文本） |
| `content_id` | `ContentId` | `*string` | JSON body | 未标注 | 内容id |
| `could_comment` | `CouldComment` | `*string` | JSON body | 未标注 | 是否允许评论 T:允许 F:不允许，默认不允许 |
| `cover` | `Cover` | `*string` | JSON body | 未标注 | 封面图url, 尺寸为996*450，最大不超过3M，支持格式:.jpg、.png ，请先调用&lt;ahref="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"&gt; 图片上传接口获得图片url。 |
| `ctype` | `Ctype` | `*string` | JSON body | 未标注 | 图文类型 activity: 活动图文，不填默认普通图文 |
| `ext_tags` | `ExtTags` | `*string` | JSON body | 未标注 | 关键词列表，英文逗号分隔，最多不超过5个 |
| `login_ids` | `LoginIds` | `*string` | JSON body | 未标注 | 可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 标题 |

<a id="api-publicmessagecustomsend"></a>

### PublicMessageCustomSend — 异步单发消息

- 官方接口：[`alipay.open.public.message.custom.send`](https://opendocs.alipay.com/apis/05wcpx)
- 路由：`POST /v3/alipay/open/public/message/custom/send`
- 请求类型：[PublicMessageCustomSendRequest](message_request.go#L47)
- 响应类型：[PublicMessageCustomSendResponse](message_response.go#L21)

开发者可使用该接口向指定关注用户推送客服消息，一般用于开发者使用接口与用户进行实时信息交互，例如自动回答用户包含某些指定关键字的咨询等。注意：用户必须在48小时之内与该生活号有交互行为才能收到消息（交互行为包括进入生活号、点击菜单、向生活号发送消息等），每个生活号每天只能向1个用户推送100条消息，超过这个数目会报错频率超限

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `articles` | `Articles` | [`*[]Article`](#model-article) | JSON body | 未标注 | msg_type为image-text，本参数必填。 |
| `chat` | `Chat` | `*string` | JSON body | 未标注 | 是否为聊天消息。枚举支持： *0：代表非聊天消息，消息显示在生活号主页，默认为 0。 *1：代表是聊天消息，消息显示在咨询反馈列表页。 |
| `event_type` | `EventType` | `*string` | JSON body | 未标注 | 触发消息的事件类型，默认为空。代表商户未改造。如果是follow，代表关注消息。click代表菜单点击，enter_ppchat代表进入事件；请注意事件类型的大小写 |
| `msg_type` | `MsgType` | `*string` | JSON body | 未标注 | 消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝用户的唯一标识 |
| `text` | `Text` | [`*Text`](#model-text) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `to_user_id` | `ToUserId` | `*string` | JSON body | 未标注 | 消息接收用户的userid |

<a id="api-publicmessagegroupsend"></a>

### PublicMessageGroupSend — 分组消息发送接口

- 官方接口：[`alipay.open.public.message.group.send`](https://opendocs.alipay.com/apis/05wyfq)
- 路由：`POST /v3/alipay/open/public/message/group/send`
- 请求类型：[PublicMessageGroupSendRequest](message_request.go#L64)
- 响应类型：[PublicMessageGroupSendResponse](message_response.go#L25)

开发者可调用此接口向指定分组中的关注用户发送消息，该接口为异步接口，一般请求成功后一段时间用户才能收到消息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `articles` | `Articles` | [`*[]Article`](#model-article) | JSON body | 未标注 | 图文消息，当 msg_type为image-text，本参数必填。 |
| `group_id` | `GroupId` | `*string` | JSON body | 未标注 | 分组ID，通过[alipay.open.public.group.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.group.create)（用户分组创建接口) 创建用户分组获取。 |
| `image` | `Image` | [`*Image`](#model-image) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `msg_type` | `MsgType` | `*string` | JSON body | 未标注 | 消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。 |
| `text` | `Text` | [`*Text`](#model-text) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-publicmessagelabelsend"></a>

### PublicMessageLabelSend — 根据标签组发消息接口

- 官方接口：[`alipay.open.public.message.label.send`](https://opendocs.alipay.com/pre-apis/05xlk0)
- 路由：`POST /v3/alipay/open/public/message/label/send`
- 请求类型：[PublicMessageLabelSendRequest](message_request.go#L76)
- 响应类型：[PublicMessageLabelSendResponse](message_response.go#L31)

开发者可以通过标签运算圈定一批用户，并且向这批用户推送消息。关于标签组发有以下注意点：
除了商户自定义标签，还支持支付宝用户标签，我们将陆续开放用户画像标签，目前支持的支付宝标签及说明请查看支付宝开放标签；
用户及标签数据有1天的缓存时间，即昨天开发者给M用户打了a标签，今天给有a标签的用户组发消息，M能够收到消息；但是如果开发者今天给M用户打了a标签，然后向有a标签的用户组发消息，M不能收到消息；
与群发消息接口类似，该接口调用之后，消息发送系统会处理一段时间，调用接口到发送成功存在一定的时延，几分钟到半个小时不等；
对于同一个用户而言，一周只能收到一条群发或者标签组发的消息；
消息组发接口一周最多调用5次。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `filter` | `Filter` | [`*Filter`](#model-filter) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `material` | `Material` | [`*Material`](#model-material) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-publicmessagesinglesend"></a>

### PublicMessageSingleSend — 单发模板消息

- 官方接口：[`alipay.open.public.message.single.send`](https://opendocs.alipay.com/apis/06bpqn)
- 路由：`POST /v3/alipay/open/public/message/single/send`
- 请求类型：[PublicMessageSingleSendRequest](message_request.go#L82)
- 响应类型：[PublicMessageSingleSendResponse](message_response.go#L35)

商户可使用该接口向用户推送制定模板的消息，如动账提醒、停车通知等消息，消息模板可在生活号后台(fuwu.alipay.com)领取，点击“模板消息”菜单，即可在模板库中领取行业对应的模板

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template` | `Template` | [`*Template`](#model-template) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `to_open_id` | `ToOpenId` | `*string` | JSON body | 未标注 | 支付宝openId，用于支付宝用户在当前应用下的用户标识。 |
| `to_user_id` | `ToUserId` | `*string` | JSON body | 未标注 | 消息接收用户的支付宝用户id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 |

<a id="api-publicmessagetotalsend"></a>

### PublicMessageTotalSend — 群发消息

- 官方接口：[`alipay.open.public.message.total.send`](https://opendocs.alipay.com/apis/05wnte)
- 路由：`POST /v3/alipay/open/public/message/total/send`
- 请求类型：[PublicMessageTotalSendRequest](message_request.go#L91)
- 响应类型：[PublicMessageTotalSendResponse](message_response.go#L39)

群发消息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `articles` | `Articles` | [`*[]Article`](#model-article) | JSON body | 未标注 | 图文消息，当 msg_type为image-text，本参数必填。 |
| `msg_type` | `MsgType` | `*string` | JSON body | 未标注 | 消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。 |
| `text` | `Text` | [`*Text`](#model-text) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-article"></a>

### Article

类型定义：[Article](message_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_name` | `ActionName` | `*string` | 未标注 | 链接文字 |
| `desc` | `Desc` | `*string` | 未标注 | 图文消息内容 |
| `image_url` | `ImageUrl` | `*string` | 未标注 | 图片链接，对于多条图文消息的第一条消息，该字段不能为空; 请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `title` | `Title` | `*string` | 未标注 | 图文消息标题 |
| `url` | `Url` | `*string` | 未标注 | 点击图文消息跳转的链接 |

<a id="model-context"></a>

### Context

类型定义：[Context](message_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `action_name` | `ActionName` | `*string` | 未标注 | 底部链接描述文字，如“查看详情”，最多能传8个汉字或16个英文字符，长度超出会报错 |
| `first` | `First` | [`*Keyword`](#model-keyword) | 未标注 | 见字段类型及官方接口文档。 |
| `head_color` | `HeadColor` | `*string` | 未标注 | 顶部色条的色值 |
| `keyword1` | `Keyword1` | [`*Keyword`](#model-keyword) | 未标注 | 见字段类型及官方接口文档。 |
| `keyword2` | `Keyword2` | [`*Keyword`](#model-keyword) | 未标注 | 见字段类型及官方接口文档。 |
| `remark` | `Remark` | [`*Keyword`](#model-keyword) | 未标注 | 见字段类型及官方接口文档。 |
| `url` | `Url` | `*string` | 未标注 | 点击消息后承接页的地址 |

<a id="model-filter"></a>

### Filter

类型定义：[Filter](message_model.go#L33)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `context` | `Context` | [`*LabelContext`](#model-labelcontext) | 未标注 | 见字段类型及官方接口文档。 |
| `template` | `Template` | `*string` | 未标注 | 过滤器模板，${a}是一个变量，会被context参数中的a参数替换，从而展开为最终的表达式，template最多支持两个参数，支持and及or连接符。<br> and：同时满足条件；<br> or：只需满足其中一个条件 |

<a id="model-image"></a>

### Image

类型定义：[Image](message_model.go#L40)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `url` | `Url` | `*string` | 未标注 | 图片url，请先调用alipay.offline.material.image.upload 图片上传接口获得图片url |

<a id="model-keyword"></a>

### Keyword

类型定义：[Keyword](message_model.go#L46)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `color` | `Color` | `*string` | 未标注 | 当前文字颜色 |
| `value` | `Value` | `*string` | 未标注 | 模板中占位符的值 |

<a id="model-labelcontext"></a>

### LabelContext

类型定义：[LabelContext](message_model.go#L54)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `a` | `A` | [`*LabelFilter`](#model-labelfilter) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-labelfilter"></a>

### LabelFilter

类型定义：[LabelFilter](message_model.go#L59)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `column_name` | `ColumnName` | `*string` | 未标注 | 标签组名，商户自定义的标签固定为label_id_list，支付宝开放的标签详见[支付宝开放标签](http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/FirstPartOfTagsFromAlipay.xlsx) |
| `op` | `Op` | `*string` | 未标注 | 操作符，支持=、!=、in三个操作符；其中in表示是某几个标签中的一个 |
| `values` | `Values` | `*[]string` | 未标注 | 标签数组，用于组装最后的表达式 |

<a id="model-material"></a>

### Material

类型定义：[Material](message_model.go#L69)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `articles` | `Articles` | [`*[]Article`](#model-article) | 未标注 | 图文消息子消息项集合，单条消息最多6个子项，否则会发送失败 |
| `msg_type` | `MsgType` | `*string` | 未标注 | 消息类型，text：文本类型，image-text：图文类型。当消息类型为text时，text参数必传，当消息类型为image-text时，articles参数必传 |
| `text` | `Text` | [`*Text`](#model-text) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-template"></a>

### Template

类型定义：[Template](message_model.go#L78)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `context` | `Context` | [`*Context`](#model-context) | 未标注 | 见字段类型及官方接口文档。 |
| `template_id` | `TemplateId` | `*string` | 未标注 | 消息模板ID |

<a id="model-text"></a>

### Text

类型定义：[Text](message_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `content` | `Content` | `*string` | 未标注 | 文本消息的内容 |
| `title` | `Title` | `*string` | 未标注 | 文本消息的标题 |
