# marketing/recruit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/recruit`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `recruit.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](#使用示例)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`RecruitEnrollClose`](#api-recruitenrollclose) | [`alipay.marketing.recruit.enroll.close`](https://opendocs.alipay.com/pre-apis/05xe2w) | 下线报名 | `POST` |
| [`RecruitEnrollCreate`](#api-recruitenrollcreate) | [`alipay.marketing.recruit.enroll.create`](https://opendocs.alipay.com/pre-apis/05xqxy) | 招商报名提交 | `POST` |
| [`RecruitEnrollQuery`](#api-recruitenrollquery) | [`alipay.marketing.recruit.enroll.query`](https://opendocs.alipay.com/pre-apis/05xk7o) | 报名详情查询 | `GET` |
| [`RecruitPlanQuery`](#api-recruitplanquery) | [`alipay.marketing.recruit.plan.query`](https://opendocs.alipay.com/pre-apis/05xhns) | 招商方案详情查询 | `POST` |
| [`RecruitPlanListQuery`](#api-recruitplanlistquery) | [`alipay.marketing.recruit.planlist.query`](https://opendocs.alipay.com/pre-apis/05xk3f) | 查询可报名的方案列表 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-recruitenrollclose"></a>

### RecruitEnrollClose — 下线报名

- 官方接口：[`alipay.marketing.recruit.enroll.close`](https://opendocs.alipay.com/pre-apis/05xe2w)
- 路由：`POST /v3/alipay/marketing/recruit/enroll/close`
- 请求类型：[RecruitEnrollCloseRequest](recruit_request.go#L5)
- 响应类型：[RecruitEnrollCloseResponse](recruit_response.go#L9)

下线报名

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enroll_id` | `EnrollId` | `*string` | JSON body | 未标注 | 活动报名ID |

<a id="api-recruitenrollcreate"></a>

### RecruitEnrollCreate — 招商报名提交

- 官方接口：[`alipay.marketing.recruit.enroll.create`](https://opendocs.alipay.com/pre-apis/05xqxy)
- 路由：`POST /v3/alipay/marketing/recruit/enroll/create`
- 请求类型：[RecruitEnrollCreateRequest](recruit_request.go#L11)
- 响应类型：[RecruitEnrollCreateResponse](recruit_response.go#L12)

创建报名后立即提交，审核通过后会以消息的形式通知调用方（需要接入消息接口alipay.marketing.enroll.status.changed）。在消息通知前可以尝试调用“报名详情查询接口（alipay.marketing.enroll.detail.query）”了解报名状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enroll_info` | `EnrollInfo` | [`*RecruitEnrollInfo`](#model-recruitenrollinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `enroll_scene_type` | `EnrollSceneType` | `*string` | JSON body | 未标注 | 活动报名场景配置。 枚举值： VOUCHER：券报名场景的报名； MINI_APP：小程序报名场景的报名； |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部操作流水号。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。 |
| `plan_id` | `PlanId` | `*string` | JSON body | 未标注 | 要报名的招商方案ID |

<a id="api-recruitenrollquery"></a>

### RecruitEnrollQuery — 报名详情查询

- 官方接口：[`alipay.marketing.recruit.enroll.query`](https://opendocs.alipay.com/pre-apis/05xk7o)
- 路由：`GET /v3/alipay/marketing/recruit/enroll/query`
- 请求类型：[RecruitEnrollQueryRequest](recruit_request.go#L40)
- 响应类型：[RecruitEnrollQueryResponse](recruit_response.go#L20)

报名详情查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 外部操作流水号，创建招商报名时传入。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。 |
| `enroll_id` | `EnrollId` | `*string` | query | 否 | 报名ID，此参数和out_biz_no至少传一个，优先取enroll_id |

<a id="api-recruitplanquery"></a>

### RecruitPlanQuery — 招商方案详情查询

- 官方接口：[`alipay.marketing.recruit.plan.query`](https://opendocs.alipay.com/pre-apis/05xhns)
- 路由：`POST /v3/alipay/marketing/recruit/plan/query`
- 请求类型：[RecruitPlanQueryRequest](recruit_request.go#L22)
- 响应类型：[RecruitPlanQueryResponse](recruit_response.go#L39)

招商方案详情查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enroll_merchant` | `EnrollMerchant` | [`*RecruitEnrollMerchant`](#model-recruitenrollmerchant) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `plan_id` | `PlanId` | `*string` | JSON body | 未标注 | 招商方案ID |

<a id="api-recruitplanlistquery"></a>

### RecruitPlanListQuery — 查询可报名的方案列表

- 官方接口：[`alipay.marketing.recruit.planlist.query`](https://opendocs.alipay.com/pre-apis/05xk3f)
- 路由：`POST /v3/alipay/marketing/recruit/planlist/query`
- 请求类型：[RecruitPlanListQueryRequest](recruit_request.go#L29)
- 响应类型：[RecruitPlanListQueryResponse](recruit_response.go#L61)

查询可报名的方案列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enroll_merchant` | `EnrollMerchant` | [`*RecruitEnrollMerchant`](#model-recruitenrollmerchant) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `enroll_scene_type` | `EnrollSceneType` | `*string` | JSON body | 未标注 | 活动报名场景配置。 枚举值： VOUCHER：券报名场景查询； MINI_APP：小程序报名场景查询； VOUCHER,MINI_APP：同时进行券报名场景和小程序报名场景的查询； |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 第几页，默认1（需要在1和999之间） |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页记录条数，默认20，需要在1和100之间 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-recruitenrollinfo"></a>

### RecruitEnrollInfo

类型定义：[RecruitEnrollInfo](recruit_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cities` | `Cities` | `*[]string` | 未标注 | 报名提交的活动城市码，详见[活动城市码](https://opendocs.alipay.com/pre-open/02r07u) |
| `enroll_merchant` | `EnrollMerchant` | [`*RecruitEnrollMerchant`](#model-recruitenrollmerchant) | 未标注 | 见字段类型及官方接口文档。 |
| `materials` | `Materials` | [`*[]RecruitMaterial`](#model-recruitmaterial) | 未标注 | 报名提交的素材 |
| `mini_apps` | `MiniApps` | [`*[]RecruitMiniApp`](#model-recruitminiapp) | 未标注 | 报名提交的小程序信息，是否必选取决于方案要求 |
| `vouchers` | `Vouchers` | [`*[]RecruitVoucher`](#model-recruitvoucher) | 未标注 | 报名提交的券信息，是否必选取决于方案要求 |

<a id="model-recruitenrollmerchant"></a>

### RecruitEnrollMerchant

类型定义：[RecruitEnrollMerchant](recruit_model.go#L18)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `merchant_uid` | `MerchantUid` | `*string` | 未标注 | 商户唯一ID(商户pid) |

<a id="model-recruitmaterial"></a>

### RecruitMaterial

类型定义：[RecruitMaterial](recruit_model.go#L45)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `data` | `Data` | `*string` | 未标注 | 提报的素材内容，json字符串，使用时需要把此字符串解析成json对象。 字段来源于“招商方案查询接口（https://opendocs.alipay.com/pre-open/02bhl7）”中返回的素材字段field。如：activityImage、contentId等。 字段类型mediaType包含： IMAGE，图片类型 TEXT，文本类型 DATASOURCE，数据源类型， 如：内容ID（contentId）字段类型为DATASOURCE 字段取值： 图片（IMAGE）： aftsFileId：“营销图片资源上传接口（https://opendocs.alipay.com/pre-open/02bhlj）”返回的图片唯一资源ID 文本（TEXT）： text：文本内容 数据源（DATASOURCE）： text：当前仅支持内容ID值，可通过“内容批量查询接口（https://opendocs.alipay.com/pre-apis/02lu12，仅供受邀用户使用）”查询内容ID。 |
| `description` | `Description` | `*string` | 未标注 | 对这个素材的描述，用于后续管理该素材 |
| `material_audit_info` | `MaterialAuditInfo` | `*string` | 未标注 | 素材审核信息，展示给商家的素材审核不通过的意见。素材审核通过时该字段为空。 报名详情查询接口不用填写该字段。 |
| `name` | `Name` | `*string` | 未标注 | 素材名名称 |
| `subject_id` | `SubjectId` | `*string` | 未标注 | 关联的报名主体ID。 enroll_scene_type是VOUCHER时，报名主体ID是activity_id；enroll_scene_type是MINI_APP时，报名主体ID是mini_app_id |

<a id="model-recruitminiapp"></a>

### RecruitMiniApp

类型定义：[RecruitMiniApp](recruit_model.go#L59)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 小程序ID |

<a id="model-recruitvoucher"></a>

### RecruitVoucher

类型定义：[RecruitVoucher](recruit_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `activity_id` | `ActivityId` | `*string` | 未标注 | 券活动id。支持商家券（创建商家券请参考https://opendocs.alipay.com/apis/01xm17）和支付券（创建支付券请参考https://opendocs.alipay.com/pre-apis/027185，仅供受邀用户使用），取值均为返回结果的activity_id字段。 |

## 使用示例

- [RecruitEnrollQuery](api_recruit_recruitenrollquery_example_test.go)
- [RecruitEnrollClose](api_recruit_recruitenrollclose_example_test.go)
- [RecruitEnrollCreate](api_recruit_recruitenrollcreate_example_test.go)
- [RecruitPlanQuery](api_recruit_recruitplanquery_example_test.go)
- [RecruitPlanListQuery](api_recruit_recruitplanlistquery_example_test.go)
