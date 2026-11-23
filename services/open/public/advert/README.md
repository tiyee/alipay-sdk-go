# open/public/advert 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/advert`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `advert.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_advert_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicAdvertBatchQuery`](#api-publicadvertbatchquery) | [`alipay.open.public.advert.batchquery`](https://opendocs.alipay.com/apis/05wggg) | 生活号广告位查询接口 | `POST` |
| [`PublicAdvertCreate`](#api-publicadvertcreate) | [`alipay.open.public.advert.create`](https://opendocs.alipay.com/apis/05xmcx) | 生活号广告位添加接口 | `POST` |
| [`PublicAdvertDelete`](#api-publicadvertdelete) | [`alipay.open.public.advert.delete`](https://opendocs.alipay.com/apis/05wvvm) | 生活号广告位删除接口 | `DELETE` |
| [`PublicAdvertModify`](#api-publicadvertmodify) | [`alipay.open.public.advert.modify`](https://opendocs.alipay.com/apis/05x1fw) | 生活号广告位修改接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicadvertbatchquery"></a>

### PublicAdvertBatchQuery — 生活号广告位查询接口

- 官方接口：[`alipay.open.public.advert.batchquery`](https://opendocs.alipay.com/apis/05wggg)
- 路由：`POST /v3/alipay/open/public/advert/batchquery`
- 请求类型：[PublicAdvertBatchQueryRequest](advert_request.go#L19)
- 响应类型：[PublicAdvertBatchQueryResponse](advert_response.go#L5)

可以通过此接口查询所有广告位的配置信息

请求约束：PublicAdvertBatchQueryRequest is the request type for alipay.open.public.advert.batchquery.

本接口无业务请求参数；传入空的请求结构体即可。

<a id="api-publicadvertcreate"></a>

### PublicAdvertCreate — 生活号广告位添加接口

- 官方接口：[`alipay.open.public.advert.create`](https://opendocs.alipay.com/apis/05xmcx)
- 路由：`POST /v3/alipay/open/public/advert/create`
- 请求类型：[PublicAdvertCreateRequest](advert_request.go#L5)
- 响应类型：[PublicAdvertCreateResponse](advert_response.go#L13)

可通过本接口在生活号主页添加广告位，根据业务需求配置不同的图片和跳转链接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `advert_items` | `AdvertItems` | [`*[]AdvertItem`](#model-advertitem) | JSON body | 未标注 | 广告位轮播内容列表。数量限制：大于1个，小于5个。广告位轮播内容顺序：与接口传入的顺序排列一致。 |

<a id="api-publicadvertdelete"></a>

### PublicAdvertDelete — 生活号广告位删除接口

- 官方接口：[`alipay.open.public.advert.delete`](https://opendocs.alipay.com/apis/05wvvm)
- 路由：`DELETE /v3/alipay/open/public/advert/delete`
- 请求类型：[PublicAdvertDeleteRequest](advert_request.go#L22)
- 响应类型：[PublicAdvertDeleteResponse](advert_response.go#L21)

可以通过此接口删除指定广告位的配置信息

请求约束：PublicAdvertDeleteRequest is the request type for alipay.open.public.advert.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `advert_id` | `AdvertId` | `*string` | query | 否 | 待删除的广告位id，删除通投广告位需传递此参数 |
| `advert_group` | `AdvertGroup` | `*string` | query | 否 | 待删除的广告位分组标识，删除个性化广告位需传递此参数。 |

<a id="api-publicadvertmodify"></a>

### PublicAdvertModify — 生活号广告位修改接口

- 官方接口：[`alipay.open.public.advert.modify`](https://opendocs.alipay.com/apis/05x1fw)
- 路由：`POST /v3/alipay/open/public/advert/modify`
- 请求类型：[PublicAdvertModifyRequest](advert_request.go#L11)
- 响应类型：[PublicAdvertModifyResponse](advert_response.go#L25)

可以通过此接口修改指定广告位的配置信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `advert_id` | `AdvertId` | `*string` | JSON body | 未标注 | 广告位id，通过[alipay.open.public.advert.create](https://opendocs.alipay.com/apis/api_6/alipay.open.public.advert.create)(生活号广告位添加接口)添加广告位获取。 |
| `advert_items` | `AdvertItems` | [`*[]AdvertItem`](#model-advertitem) | JSON body | 未标注 | 广告位轮播内容列表。数量限制：大于1个，小于5个。广告位轮播内容顺序：与接口传入的顺序排列一致。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-advertitem"></a>

### AdvertItem

类型定义：[AdvertItem](advert_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `img_url` | `ImgUrl` | `*string` | 未标注 | 广告图片url, 尺寸为996*240，最大不超过5M，支持格式:.jpg、.png ，请先调用[ 图片上传接口](https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload)获得图片url |
| `link_type` | `LinkType` | `*string` | 未标注 | 跳转类型，网页:HTTP、小程序:APP，不传默认HTTP |
| `link_url` | `LinkUrl` | `*string` | 未标注 | 跳转链接，点击广告图片跳到的链接url |
| `name` | `Name` | `*string` | 未标注 | 广告帧名称，可选 |
