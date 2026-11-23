# open/public/life 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/life`。本目录提供 9 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `life.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicLifeAboardApply`](#api-publiclifeaboardapply) | [`alipay.open.public.life.aboard.apply`](https://opendocs.alipay.com/apis/05wyga) | 上架生活号 | `POST` |
| [`PublicLifeAgentCreateQuery`](#api-publiclifeagentcreatequery) | [`alipay.open.public.life.agentcreate.query`](https://opendocs.alipay.com/apis/05wuo2) | isv代创建生活号申请状态查询接口 | `GET` |
| [`PublicLifeDebarkApply`](#api-publiclifedebarkapply) | [`alipay.open.public.life.debark.apply`](https://opendocs.alipay.com/apis/05wnsi) | 下架生活号 | `POST` |
| [`PublicLifeLabelBatchQuery`](#api-publiclifelabelbatchquery) | [`alipay.open.public.life.label.batchquery`](https://opendocs.alipay.com/apis/05wui2) | 标签批量查询接口 | `POST` |
| [`PublicLifeLabelCreate`](#api-publiclifelabelcreate) | [`alipay.open.public.life.label.create`](https://opendocs.alipay.com/apis/05wh9j) | 创建标签接口 | `POST` |
| [`PublicLifeLabelDelete`](#api-publiclifelabeldelete) | [`alipay.open.public.life.label.delete`](https://opendocs.alipay.com/apis/05wvsg) | 标签删除接口 | `DELETE` |
| [`PublicLifeLabelModify`](#api-publiclifelabelmodify) | [`alipay.open.public.life.label.modify`](https://opendocs.alipay.com/apis/05xhmq) | 标签修改接口 | `POST` |
| [`PublicLifeMsgRecall`](#api-publiclifemsgrecall) | [`alipay.open.public.life.msg.recall`](https://opendocs.alipay.com/apis/05wieh) | 生活号消息撤回接口 | `POST` |
| [`PublicLifeMsgSend`](#api-publiclifemsgsend) | [`alipay.open.public.life.msg.send`](https://opendocs.alipay.com/pre-apis/05wulv) | 生活号广播消息发送接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publiclifeaboardapply"></a>

### PublicLifeAboardApply — 上架生活号

- 官方接口：[`alipay.open.public.life.aboard.apply`](https://opendocs.alipay.com/apis/05wyga)
- 路由：`POST /v3/alipay/open/public/life/aboard/apply`
- 请求类型：[PublicLifeAboardApplyRequest](life_request.go#L37)
- 响应类型：[PublicLifeAboardApplyResponse](life_response.go#L5)

上架生活号

请求约束：PublicLifeAboardApplyRequest is the request type for alipay.open.public.life.aboard.apply.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publiclifeagentcreatequery"></a>

### PublicLifeAgentCreateQuery — isv代创建生活号申请状态查询接口

- 官方接口：[`alipay.open.public.life.agentcreate.query`](https://opendocs.alipay.com/apis/05wuo2)
- 路由：`GET /v3/alipay/open/public/life/agentcreate/query`
- 请求类型：[PublicLifeAgentCreateQueryRequest](life_request.go#L40)
- 响应类型：[PublicLifeAgentCreateQueryResponse](life_response.go#L11)

isv可调用此接口查询代商户创建的生活号审核状态

请求约束：PublicLifeAgentCreateQueryRequest is the request type for alipay.open.public.life.agentcreate.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 由开发者创建的外部入驻申请单据号，需与[alipay.open.public.life.agent.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.agent.create)(isv代创建生活号接口)的out_biz_no一致。 |

<a id="api-publiclifedebarkapply"></a>

### PublicLifeDebarkApply — 下架生活号

- 官方接口：[`alipay.open.public.life.debark.apply`](https://opendocs.alipay.com/apis/05wnsi)
- 路由：`POST /v3/alipay/open/public/life/debark/apply`
- 请求类型：[PublicLifeDebarkApplyRequest](life_request.go#L46)
- 响应类型：[PublicLifeDebarkApplyResponse](life_response.go#L25)

下架生活号

请求约束：PublicLifeDebarkApplyRequest is the request type for alipay.open.public.life.debark.apply.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publiclifelabelbatchquery"></a>

### PublicLifeLabelBatchQuery — 标签批量查询接口

- 官方接口：[`alipay.open.public.life.label.batchquery`](https://opendocs.alipay.com/apis/05wui2)
- 路由：`POST /v3/alipay/open/public/life/label/batchquery`
- 请求类型：[PublicLifeLabelBatchQueryRequest](life_request.go#L49)
- 响应类型：[PublicLifeLabelBatchQueryResponse](life_response.go#L31)

开发者可调用该接口查询生活号支持的标签列表，其中包括生活号通用标签、生活号自定义标签以及数据实验室标签。

请求约束：PublicLifeLabelBatchQueryRequest is the request type for alipay.open.public.life.label.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publiclifelabelcreate"></a>

### PublicLifeLabelCreate — 创建标签接口

- 官方接口：[`alipay.open.public.life.label.create`](https://opendocs.alipay.com/apis/05wh9j)
- 路由：`POST /v3/alipay/open/public/life/label/create`
- 请求类型：[PublicLifeLabelCreateRequest](life_request.go#L9)
- 响应类型：[PublicLifeLabelCreateResponse](life_response.go#L37)

开发者可使用该API创建自定义标签

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data_type` | `DataType` | `*string` | JSON body | 未标注 | 标签值类型，不填默认为 string 类型。 注意：目前只支持 string（字符串类型）。 |
| `label_name` | `LabelName` | `*string` | JSON body | 未标注 | 自定义标签名。 注意：每个生活号最多创建 100 个自定义标签。 |

<a id="api-publiclifelabeldelete"></a>

### PublicLifeLabelDelete — 标签删除接口

- 官方接口：[`alipay.open.public.life.label.delete`](https://opendocs.alipay.com/apis/05wvsg)
- 路由：`DELETE /v3/alipay/open/public/life/label/delete`
- 请求类型：[PublicLifeLabelDeleteRequest](life_request.go#L52)
- 响应类型：[PublicLifeLabelDeleteResponse](life_response.go#L43)

开发者可调用此API删除自定义标签

请求约束：PublicLifeLabelDeleteRequest is the request type for alipay.open.public.life.label.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | query | 否 | 标签 id，只支持生活号自定义标签。通过 [alipay.open.public.life.label.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create)(创建标签接口)创建自定义标签后获取。 |

<a id="api-publiclifelabelmodify"></a>

### PublicLifeLabelModify — 标签修改接口

- 官方接口：[`alipay.open.public.life.label.modify`](https://opendocs.alipay.com/apis/05xhmq)
- 路由：`POST /v3/alipay/open/public/life/label/modify`
- 请求类型：[PublicLifeLabelModifyRequest](life_request.go#L17)
- 响应类型：[PublicLifeLabelModifyResponse](life_response.go#L47)

开发者可调用此API修改自定义标签的名称

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `label_id` | `LabelId` | `*string` | JSON body | 未标注 | 标签 id，只支持生活号自定义标签。通过 [alipay.open.public.life.label.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create)(创建标签接口)创建自定义标签后获取。 |
| `label_name` | `LabelName` | `*string` | JSON body | 未标注 | 标签名 |

<a id="api-publiclifemsgrecall"></a>

### PublicLifeMsgRecall — 生活号消息撤回接口

- 官方接口：[`alipay.open.public.life.msg.recall`](https://opendocs.alipay.com/apis/05wieh)
- 路由：`POST /v3/alipay/open/public/life/msg/recall`
- 请求类型：[PublicLifeMsgRecallRequest](life_request.go#L25)
- 响应类型：[PublicLifeMsgRecallResponse](life_response.go#L51)

生活号消息撤回接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `message_id` | `MessageId` | `*string` | JSON body | 未标注 | 消息id，可在调用[群发消息](https://opendocs.alipay.com/apis/api_6/alipay.open.public.message.total.send)、[组发消息](https://opendocs.alipay.com/apis/api_6/alipay.open.public.message.group.send)接口时会返回，请注意保存留用。 |

<a id="api-publiclifemsgsend"></a>

### PublicLifeMsgSend — 生活号广播消息发送接口

- 官方接口：[`alipay.open.public.life.msg.send`](https://opendocs.alipay.com/pre-apis/05wulv)
- 路由：`POST /v3/alipay/open/public/life/msg/send`
- 请求类型：[PublicLifeMsgSendRequest](life_request.go#L31)
- 响应类型：[PublicLifeMsgSendResponse](life_response.go#L55)

媒体资讯生活号群发一篇图文或视频类消息，生活号主页以及客户端首页展示消息详情

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cover` | `Cover` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*PublicLifeMsgSend`](#model-publiclifemsgsend) | multipart | 未标注 | 见字段类型及官方接口文档。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-publiclifemsgsend"></a>

### PublicLifeMsgSend

类型定义：[PublicLifeMsgSend](life_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `category` | `Category` | `*string` | 未标注 | 消息分类，请传入对应分类编码值 |
| `content` | `Content` | `*string` | 未标注 | 消息正文，html原文或纯文本 |
| `desc` | `Desc` | `*string` | 未标注 | 消息概述 |
| `msg_type` | `MsgType` | `*string` | 未标注 | 媒体资讯类生活号消息类型 |
| `source_ext_info` | `SourceExtInfo` | `*string` | 未标注 | 消息来源方附属信息，供搜索、推荐使用<br> publish_time（int）：消息发布时间，单位秒<br> keyword_list（String）: 文章的标签列表，英文逗号分隔<br> comment（int）：消息来源处评论次数<br> reward（int）：消息来源处打赏次数<br> is_recommended（boolean）：消息在来源处是否被推荐<br> is_news（boolean）：消息是否实时性内容<br> read（int）：消息在来源处被阅读次数<br> like（int）：消息在来源处被点赞次数<br> is_hot（boolean）：消息在来源平台是否是热门内容<br> share（int）：文章在来源平台的分享次数<br> deadline（int）：文章的失效时间，单位秒 |
| `title` | `Title` | `*string` | 未标注 | 消息标题 |
| `unique_msg_id` | `UniqueMsgId` | `*string` | 未标注 | 来源方消息唯一标识；若不为空，根据此id和生活号id对消息去重；若为空，不去重 |
| `video_length` | `VideoLength` | `*string` | 未标注 | 生活号消息视频时长，单位：秒（视频类消息必填） |
| `video_samples` | `VideoSamples` | `*[]string` | 未标注 | 视频类型消息中视频抽样关键帧截图，视频类消息选填 |
| `video_size` | `VideoSize` | `*string` | 未标注 | 视频大小，单位：KB（视频类消息必填） |
| `video_source` | `VideoSource` | `*string` | 未标注 | 视频资源来源id（视频类消息必填），取值限定youku, miaopai, taobao, sina中的一个 |
| `video_temporary_url` | `VideoTemporaryUrl` | `*string` | 未标注 | 视频的临时链接（优酷来源的视频消息，该字段不能为空） |
| `video_url` | `VideoUrl` | `*string` | 未标注 | 生活号视频类消息视频id或url（视频类消息必填，根据来源区分） |
