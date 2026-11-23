# open/public/info 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/public/info`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `info.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_info_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`PublicInfoModify`](#api-publicinfomodify) | [`alipay.open.public.info.modify`](https://opendocs.alipay.com/apis/05wz9l) | 生活号基础信息修改接口 | `POST` |
| [`PublicInfoQuery`](#api-publicinfoquery) | [`alipay.open.public.info.query`](https://opendocs.alipay.com/apis/05wygh) | 生活号基础信息查询接口 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-publicinfomodify"></a>

### PublicInfoModify — 生活号基础信息修改接口

- 官方接口：[`alipay.open.public.info.modify`](https://opendocs.alipay.com/apis/05wz9l)
- 路由：`POST /v3/alipay/open/public/info/modify`
- 请求类型：[PublicInfoModifyRequest](info_request.go#L5)
- 响应类型：[PublicInfoModifyResponse](info_response.go#L5)

提供生活号基础信息修改，同时适用于口碑商户签约当面付2.0默认创建的服务窗，需上传相关资质材料。修改名称、头像、简介会触发风控审核，审核通过后修改生效，支持isv调用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_name` | `AppName` | `*string` | JSON body | 未标注 | 生活号名称，2-20个字之间。注意： * 不得含有违反法律法规和公序良俗的相关信息； * 不得侵害他人名誉权、知识产权、商业秘密等合法权利； * 不得以太过广泛的、或产品、行业词组来命名，如：女装、皮革批发； * 不得以实名认证的媒体资质账号创建服务窗，或媒体相关名称命名服务窗，如：XX电视台、XX杂志等。 |
| `auth_pic` | `AuthPic` | `*string` | JSON body | 未标注 | 授权运营书图片地址，企业商户若为被经营方授权，需上传加盖公章的扫描件。支持 .jpg、 .jpeg、 .png 格式，需小于1M。使用 [alipay.offline.material.image.upload](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。 |
| `background_url` | `BackgroundUrl` | `*string` | JSON body | 未标注 | 背景图片地址，建议尺寸 1600 x 1000px，支持.jpg .jpeg .png格式，小于1M。使用 [alipay.offline.material.image.upload](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。 |
| `introduction` | `Introduction` | `*string` | JSON body | 未标注 | 简介 |
| `license_url` | `LicenseUrl` | `*string` | JSON body | 未标注 | 营业执照地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 [alipay.offline.material.image.upload](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。 |
| `logo_url` | `LogoUrl` | `*string` | JSON body | 未标注 | 生活号头像地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 [alipay.offline.material.image.upload](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `public_greeting` | `PublicGreeting` | `*string` | JSON body | 未标注 | 欢迎语 |
| `shop_pics` | `ShopPics` | `*[]string` | JSON body | 未标注 | 门店照片地址，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 [alipay.offline.material.image.upload](https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload)(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。 |

<a id="api-publicinfoquery"></a>

### PublicInfoQuery — 生活号基础信息查询接口

- 官方接口：[`alipay.open.public.info.query`](https://opendocs.alipay.com/apis/05wygh)
- 路由：`GET /v3/alipay/open/public/info/query`
- 请求类型：[PublicInfoQueryRequest](info_request.go#L27)
- 响应类型：[PublicInfoQueryResponse](info_response.go#L13)

提供名称、头像、简介、欢迎语、上下架、审核信息查询，支持isv调用

请求约束：PublicInfoQueryRequest is the request type for alipay.open.public.info.query.

本接口无业务请求参数；传入空的请求结构体即可。
