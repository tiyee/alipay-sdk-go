# open/public/topic 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/topic`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `topic.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_topic_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicTopicBatchQuery`](#api-publictopicbatchquery) | [`alipay.open.public.topic.batchquery`](https://opendocs.alipay.com/apis/05wvr7) | 营销位查询接口 | `POST` |
| [`PublicTopicCreate`](#api-publictopiccreate) | [`alipay.open.public.topic.create`](https://opendocs.alipay.com/apis/05wuj6) | 营销位添加接口 | `POST` |
| [`PublicTopicDelete`](#api-publictopicdelete) | [`alipay.open.public.topic.delete`](https://opendocs.alipay.com/apis/05wyj8) | 营销位删除接口 | `DELETE` |
| [`PublicTopicModify`](#api-publictopicmodify) | [`alipay.open.public.topic.modify`](https://opendocs.alipay.com/apis/05wuk1) | 营销位修改接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publictopicbatchquery"></a>

### PublicTopicBatchQuery — 营销位查询接口

- 官方接口：[`alipay.open.public.topic.batchquery`](https://opendocs.alipay.com/apis/05wvr7)
- 路由：`POST /v3/alipay/open/public/topic/batchquery`
- 请求类型：[PublicTopicBatchQueryRequest](topic_request.go#L38)
- 响应类型：[PublicTopicBatchQueryResponse](topic_response.go#L5)

通过此接口,开发者可以查询已经在生活号主页配置的营销位具体信息，包括标题、描述、链接等信息。

请求约束：PublicTopicBatchQueryRequest is the request type for alipay.open.public.topic.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publictopiccreate"></a>

### PublicTopicCreate — 营销位添加接口

- 官方接口：[`alipay.open.public.topic.create`](https://opendocs.alipay.com/apis/05wuj6)
- 路由：`POST /v3/alipay/open/public/topic/create`
- 请求类型：[PublicTopicCreateRequest](topic_request.go#L5)
- 响应类型：[PublicTopicCreateResponse](topic_response.go#L11)

开发者可以通过此接口，在生活号主页中配置营销位，支持展示4-8个营销位，满足商家对业务的需求，需要注意的是，当某个营销位需要显示为头图时，必须要配置以下三个参数：营销位描述 (sub_title)、营销位图片（img_url）、营销位链接地址（link_url），如不需要展示头图时，以上三个参数均不要传。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `img_url` | `ImgUrl` | `*string` | JSON body | 未标注 | 营销位图片url, 尺寸为996*450，最大不超过5M，支持格式:.jpg、.png ，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url。营销位需要展示头图时，必须填写该参数。 |
| `link_type` | `LinkType` | `*string` | JSON body | 未标注 | 跳转类型，网页:HTTP、小程序:APP，不传默认HTTP |
| `link_url` | `LinkUrl` | `*string` | JSON body | 未标注 | 营销位跳转地址，点击营销位头图跳到的链接url。营销位需要展示头图时，必须填写该参数。 |
| `sub_title` | `SubTitle` | `*string` | JSON body | 未标注 | 营销位描述。营销位需要展示头图时，必须填写该参数。 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 营销位名称 |
| `topic_items` | `TopicItems` | [`*[]TopicItem`](#model-topicitem) | JSON body | 未标注 | 营销位内容，数量限制：大于4个，小于8个 |

<a id="api-publictopicdelete"></a>

### PublicTopicDelete — 营销位删除接口

- 官方接口：[`alipay.open.public.topic.delete`](https://opendocs.alipay.com/apis/05wyj8)
- 路由：`DELETE /v3/alipay/open/public/topic/delete`
- 请求类型：[PublicTopicDeleteRequest](topic_request.go#L41)
- 响应类型：[PublicTopicDeleteResponse](topic_response.go#L17)

通过此接口,开发者可以根据营销位的分组group删除指定的营销位。

请求约束：PublicTopicDeleteRequest is the request type for alipay.open.public.topic.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `topic_id` | `TopicId` | `*string` | query | 否 | 营销位id |

<a id="api-publictopicmodify"></a>

### PublicTopicModify — 营销位修改接口

- 官方接口：[`alipay.open.public.topic.modify`](https://opendocs.alipay.com/apis/05wuk1)
- 路由：`POST /v3/alipay/open/public/topic/modify`
- 请求类型：[PublicTopicModifyRequest](topic_request.go#L21)
- 响应类型：[PublicTopicModifyResponse](topic_response.go#L21)

通过此接口,开发者可以修改已经在生活号主页配置的营销位具体信息，包括标题、描述、链接等信息。需要注意的是，当某个营销位需要显示为头图时，必须要配置以下三个参数：营销位描述 (sub_title)、营销位图片（img_url）、营销位链接地址（link_url），如不需要展示头图时，以上三个参数均不要传。请注意修改是针对整套营销位，会覆盖之前的营销位。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `img_url` | `ImgUrl` | `*string` | JSON body | 未标注 | 营销位图片url, 尺寸为996*450，最大不超过5M，支持格式:.jpg、.png ，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url。营销位需要展示头图时，必须填写该参数。 |
| `link_type` | `LinkType` | `*string` | JSON body | 未标注 | 跳转类型，网页:HTTP、小程序:APP，不传默认HTTP |
| `link_url` | `LinkUrl` | `*string` | JSON body | 未标注 | 营销位跳转地址，点击营销位头图跳到的链接url。营销位需要展示头图时，必须填写该参数。 |
| `sub_title` | `SubTitle` | `*string` | JSON body | 未标注 | 营销位描述。营销位需要展示头图时，必须填写该参数 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 营销位名称 |
| `topic_id` | `TopicId` | `*string` | JSON body | 未标注 | 营销位id |
| `topic_items` | `TopicItems` | [`*TopicItem`](#model-topicitem) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-topicitem"></a>

### TopicItem

类型定义：[TopicItem](topic_model.go#L23)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `img_url` | `ImgUrl` | `*string` | 未标注 | 内容图片url, 尺寸为300*300，最大不超过3M，支持格式:.jpg、.png ，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `link_type` | `LinkType` | `*string` | 未标注 | 跳转类型，网页:HTTP、小程序:APP，不传默认HTTP |
| `link_url` | `LinkUrl` | `*string` | 未标注 | 跳转链接，点击营销位内容图片跳到的链接url |
| `sub_title` | `SubTitle` | `*string` | 未标注 | 内容说明 |
| `title` | `Title` | `*string` | 未标注 | 内容标题 |
