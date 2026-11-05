# article 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/article`。本包提供 5 个支付宝 V3 REST 操作。

通过 `article.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CcmSwArticleBatchQuery`](#api-ccmswarticlebatchquery) | `alipay.iservice.ccm.sw.article.batchquery` |
| [`CcmSwArticleCreate`](#api-ccmswarticlecreate) | `alipay.iservice.ccm.sw.article.create` |
| [`CcmSwArticleDelete`](#api-ccmswarticledelete) | `alipay.iservice.ccm.sw.article.delete` |
| [`CcmSwArticleGet`](#api-ccmswarticleget) | `alipay.iservice.ccm.sw.article.get` |
| [`CcmSwArticleModify`](#api-ccmswarticlemodify) | `alipay.iservice.ccm.sw.article.modify` |

## 使用示例

- [CcmSwArticleBatchQuery](api_article_ccmswarticlebatchquery_example_test.go)
- [CcmSwArticleCreate](api_article_ccmswarticlecreate_example_test.go)
- [CcmSwArticleDelete](api_article_ccmswarticledelete_example_test.go)
- [CcmSwArticleGet](api_article_ccmswarticleget_example_test.go)
- [CcmSwArticleModify](api_article_ccmswarticlemodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmswarticlebatchquery"></a>

### CcmSwArticleBatchQuery — 知识库-问答-批量查询

- 官方接口：[`alipay.iservice.ccm.sw.article.batchquery`](https://opendocs.alipay.com/open-v3/031db27a_alipay.iservice.ccm.sw.article.batchquery)
- 路由：`POST /v3/alipay/iservice/ccm/sw/article/batchquery`
- 请求类型：[CcmSwArticleBatchQueryRequest](article_request.go#L4)
- 响应类型：[CcmSwArticleBatchQueryResponse](article_response.go#L10)

知识库-问答-批量查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `category_id` | `CategoryId` | `*int64` | JSON body | 未标注 | 所属类目ID，如果search_all_category为true则不用填 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 结束时间，并且开始时间不能为空 |
| `ids` | `Ids` | `*[]int64` | JSON body | 未标注 | 文章ID集合 |
| `keyword` | `Keyword` | `*string` | JSON body | 未标注 | 关键字 |
| `keywords` | `Keywords` | `*[]string` | JSON body | 未标注 | 标签 |
| `library_id` | `LibraryId` | `*int64` | JSON body | 未标注 | 知识库ID |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 页数，page_size不能为空 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 页大小，page_num不能为空 |
| `search_all_category` | `SearchAllCategory` | `*bool` | JSON body | 未标注 | 是否搜索所有类目，如果为true则不用填写category_id值 |
| `search_category_type` | `SearchCategoryType` | `*string` | JSON body | 未标注 | Current（搜索当前节点）； Children（搜索当前节点以及子节点） |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 开始时间，并且结束时间不能为空 |
| `status` | `Status` | `*string` | JSON body | 未标注 | 状态，PUBLISHED（已发布），UNPUBLISH（未发布），EXPIRED（失效），DELETED（已删除） |

<a id="api-ccmswarticlecreate"></a>

### CcmSwArticleCreate — 知识库-问答-创建

- 官方接口：[`alipay.iservice.ccm.sw.article.create`](https://opendocs.alipay.com/open-v3/7c58db8d_alipay.iservice.ccm.sw.article.create)
- 路由：`PUT /v3/alipay/iservice/ccm/sw/article/create`
- 请求类型：[CcmSwArticleCreateRequest](article_request.go#L34)
- 响应类型：[CcmSwArticleCreateResponse](article_response.go#L22)

知识库-问答-创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `attachments` | `Attachments` | [`*[]ArticleAttachmentInfo`](#model-articleattachmentinfo) | JSON body | 未标注 | 附件列表 |
| `category_id` | `CategoryId` | `*int64` | JSON body | 未标注 | 所属类目ID，如果search_all_category为true则不用填 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `content` | `Content` | `*string` | JSON body | 未标注 | 内容 |
| `extend_titles` | `ExtendTitles` | `*[]string` | JSON body | 未标注 | 扩展标题 |
| `keywords` | `Keywords` | `*[]string` | JSON body | 未标注 | 标签 |
| `library_id` | `LibraryId` | `*string` | JSON body | 未标注 | 知识库ID |
| `publish_end` | `PublishEnd` | `*string` | JSON body | 未标注 | 默认为 2999-12-31 23:59:59 |
| `publish_start` | `PublishStart` | `*string` | JSON body | 未标注 | 默认为当前时间 |
| `related_articles` | `RelatedArticles` | [`*[]BaseArticleInfo`](#model-basearticleinfo) | JSON body | 未标注 | 关联知识点 |
| `scene_codes` | `SceneCodes` | `*[]string` | JSON body | 未标注 | 场景ID。KNOWLEDGE（内部知识库）；ROBOT（机器人）;HELP（帮助中心）；WHELP（无线帮助中心） |
| `title` | `Title` | `*string` | JSON body | 未标注 | 标题 |

<a id="api-ccmswarticledelete"></a>

### CcmSwArticleDelete — 知识库-问答-删除

- 官方接口：[`alipay.iservice.ccm.sw.article.delete`](https://opendocs.alipay.com/open-v3/d2475255_alipay.iservice.ccm.sw.article.delete)
- 路由：`DELETE /v3/alipay/iservice/ccm/sw/article/delete`
- 请求类型：[CcmSwArticleDeleteRequest](article_request.go#L90)
- 响应类型：[CcmSwArticleDeleteResponse](article_response.go#L28)

知识库-问答-删除

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-ccmswarticleget"></a>

### CcmSwArticleGet — 知识库-问答-获取

- 官方接口：[`alipay.iservice.ccm.sw.article.get`](https://opendocs.alipay.com/open-v3/c2addca8_alipay.iservice.ccm.sw.article.get)
- 路由：`GET /v3/alipay/iservice/ccm/sw/article/get`
- 请求类型：[CcmSwArticleGetRequest](article_request.go#L93)
- 响应类型：[CcmSwArticleGetResponse](article_response.go#L31)

知识库-问答-获取

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | query | 否 | 子部门ID，不传为默认部门 |
| `id` | `Id` | `*int64` | query | 是 | 文章ID |

<a id="api-ccmswarticlemodify"></a>

### CcmSwArticleModify — 知识库-问答-修改

- 官方接口：[`alipay.iservice.ccm.sw.article.modify`](https://opendocs.alipay.com/open-v3/cc0b0135_alipay.iservice.ccm.sw.article.modify)
- 路由：`PATCH /v3/alipay/iservice/ccm/sw/article/modify`
- 请求类型：[CcmSwArticleModifyRequest](article_request.go#L62)
- 响应类型：[CcmSwArticleModifyResponse](article_response.go#L83)

知识库-问答-修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `attachments` | `Attachments` | [`*[]ArticleAttachmentInfo`](#model-articleattachmentinfo) | JSON body | 未标注 | 附件列表 |
| `category_id` | `CategoryId` | `*int64` | JSON body | 未标注 | 所属类目ID，如果search_all_category为true则不用填 |
| `ccs_instance_id` | `CcsInstanceId` | `*string` | JSON body | 未标注 | 子部门ID，不传为默认部门 |
| `content` | `Content` | `*string` | JSON body | 未标注 | 内容 |
| `extend_titles` | `ExtendTitles` | `*[]string` | JSON body | 未标注 | 扩展标题 |
| `id` | `Id` | `*int64` | JSON body | 未标注 | 文章ID |
| `keywords` | `Keywords` | `*[]string` | JSON body | 未标注 | 标签 |
| `publish_end` | `PublishEnd` | `*string` | JSON body | 未标注 | 有效期结束时间 |
| `publish_start` | `PublishStart` | `*string` | JSON body | 未标注 | 有效期开始时间 |
| `related_articles` | `RelatedArticles` | [`*[]BaseArticleInfo`](#model-basearticleinfo) | JSON body | 未标注 | 关联知识点列表 |
| `scene_codes` | `SceneCodes` | `*[]string` | JSON body | 未标注 | 场景ID。1（内部知识库）；2（机器人）;3（帮助中心）；4（无线帮助中心） |
| `title` | `Title` | `*string` | JSON body | 未标注 | 标题 |

## 嵌套请求参数

<a id="model-articleattachmentinfo"></a>

### ArticleAttachmentInfo

类型定义：[ArticleAttachmentInfo](article_model.go#L8)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `file_name` | `FileName` | `*string` | 未标注 | 文章附件 |
| `id` | `Id` | `*int64` | 未标注 | 附件id |
| `key` | `Key` | `*string` | 未标注 | 附件完整Key |
| `url` | `Url` | `*string` | 未标注 | 附件的url |

<a id="model-basearticleinfo"></a>

### BaseArticleInfo

类型定义：[BaseArticleInfo](article_model.go#L70)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `id` | `Id` | `*int64` | 未标注 | 知识点id |
| `title` | `Title` | `*string` | 未标注 | 知识点标题 |
