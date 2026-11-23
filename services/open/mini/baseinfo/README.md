# open/mini/baseinfo 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/baseinfo`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `baseinfo.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_baseinfo_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MiniBaseInfoModify`](#api-minibaseinfomodify) | [`alipay.open.mini.baseinfo.modify`](https://opendocs.alipay.com/open-v3/762743d3_alipay.open.mini.baseinfo.modify) | 小程序修改基础信息 | `POST` |
| [`MiniBaseInfoNameCheck`](#api-minibaseinfonamecheck) | [`alipay.open.mini.baseinfo.name.check`](https://opendocs.alipay.com/open-v3/bf9816ac_alipay.open.mini.baseinfo.name.check) | 小程序名称检测 | `GET` |
| [`MiniBaseInfoQuery`](#api-minibaseinfoquery) | [`alipay.open.mini.baseinfo.query`](https://opendocs.alipay.com/open-v3/c5995fc4_alipay.open.mini.baseinfo.query) | 查询小程序基础信息 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-minibaseinfomodify"></a>

### MiniBaseInfoModify — 小程序修改基础信息

- 官方接口：[`alipay.open.mini.baseinfo.modify`](https://opendocs.alipay.com/open-v3/762743d3_alipay.open.mini.baseinfo.modify)
- 路由：`POST /v3/alipay/open/mini/baseinfo/modify`
- 请求类型：[MiniBaseInfoModifyRequest](baseinfo_request.go#L9)
- 响应类型：[MiniBaseInfoModifyResponse](baseinfo_response.go#L5)

小程序修改基础信息。
本接口暂不支持已经上过架的小程序调用（只支持小程序首次上架之前修改信息）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*MiniBaseInfoModify`](#model-minibaseinfomodify) | multipart | 未标注 | 见字段类型及官方接口文档。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-minibaseinfonamecheck"></a>

### MiniBaseInfoNameCheck — 小程序名称检测

- 官方接口：[`alipay.open.mini.baseinfo.name.check`](https://opendocs.alipay.com/open-v3/bf9816ac_alipay.open.mini.baseinfo.name.check)
- 路由：`GET /v3/alipay/open/mini/baseinfo/name/check`
- 请求类型：[MiniBaseInfoNameCheckRequest](baseinfo_request.go#L15)
- 响应类型：[MiniBaseInfoNameCheckResponse](baseinfo_response.go#L11)

小程序名称检测。包含小程序名称基础校验、名称重复、存在敏感信息、存在品牌词、名称宽泛等检测能力。若返回“check_code”：“NONE”，说明小程序名称未检测出异常。

请求约束：MiniBaseInfoNameCheckRequest is the request type for alipay.open.mini.baseinfo.name.check.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_name` | `AppName` | `*string` | query | 否 | 小程序名字，长度限制 3~30 个字符，仅支持包含简体中文、数字、英文、下划线、+、-。 |

<a id="api-minibaseinfoquery"></a>

### MiniBaseInfoQuery — 查询小程序基础信息

- 官方接口：[`alipay.open.mini.baseinfo.query`](https://opendocs.alipay.com/open-v3/c5995fc4_alipay.open.mini.baseinfo.query)
- 路由：`GET /v3/alipay/open/mini/baseinfo/query`
- 请求类型：[MiniBaseInfoQueryRequest](baseinfo_request.go#L21)
- 响应类型：[MiniBaseInfoQueryResponse](baseinfo_response.go#L19)

查询小程序基础信息

请求约束：MiniBaseInfoQueryRequest is the request type for alipay.open.mini.baseinfo.query.

本接口无业务请求参数；传入空的请求结构体即可。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-minibaseinfomodify"></a>

### MiniBaseInfoModify

类型定义：[MiniBaseInfoModify](baseinfo_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | 未标注 | 11_12;12_13。小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考[https://docs.open.alipay.com/api_49/alipay.open.mini.category.query](https://docs.open.alipay.com/api_49/alipay.open.mini.category.query)接口查询mini_category_list 。 |
| `app_desc` | `AppDesc` | `*string` | 未标注 | 小程序描述，长度限制 20~400 个字符。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过[https://opendocs.alipay.com/mini/03l21r](https://opendocs.alipay.com/mini/03l21r)查询当前小程序信息 |
| `app_english_name` | `AppEnglishName` | `*string` | 未标注 | 小程序应用英文名称。长度限制 3~30 个字符，仅支持包含中文、数字、英文、下划线、+、-。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过[alipay.open.mini.baseinfo.query](https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common)查询当前小程序信息 |
| `app_name` | `AppName` | `*string` | 未标注 | 小程序名称。长度限制 3~30 个字符，仅支持包含中文、数字、英文、下划线、+、-。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过[alipay.open.mini.baseinfo.query](https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common )查询当前小程序信息 |
| `app_slogan` | `AppSlogan` | `*string` | 未标注 | 小程序简介，一句话描述小程序功能，长度限制 10~32个字符。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过[alipay.open.mini.baseinfo.query](https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common)查询当前小程序信息 |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 [https://docs.open.alipay.com/api_49/alipay.open.mini.category.query](https://docs.open.alipay.com/api_49/alipay.open.mini.category.query) 接口查询mini_category_list。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过[https://opendocs.alipay.com/mini/03l21r](https://opendocs.alipay.com/mini/03l21r)查询当前小程序信息 注意：个人开发者不得使用企业类目。 |
| `service_email` | `ServiceEmail` | `*string` | 未标注 | 小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 注意：2021年7月1日后，该字段将逐步灰度为可选字段，请按可选开发。 |
| `service_phone` | `ServicePhone` | `*string` | 未标注 | 小程序客服电话，长度限制5~30个字符，仅支持包含数字和-。如果不填默认采用当前小程序的应用客服电话，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 注意：2021年7月1日后，该字段将逐步灰度为必填字段，请按必填开发。 |
