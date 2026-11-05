# audit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/innerversion/audit`。本包提供 2 个支付宝 V3 REST 操作。

通过 `audit.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MiniInnerVersionAuditCancel`](#api-miniinnerversionauditcancel) | `alipay.open.mini.innerversion.audit.cancel` |
| [`MiniInnerVersionAuditSubmit`](#api-miniinnerversionauditsubmit) | `alipay.open.mini.innerversion.audit.submit` |

## 使用示例

- [MiniInnerVersionAuditCancel](api_audit_miniinnerversionauditcancel_example_test.go)
- [MiniInnerVersionAuditSubmit](api_audit_miniinnerversionauditsubmit_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerversionauditcancel"></a>

### MiniInnerVersionAuditCancel — 内部链路撤销审核

- 官方接口：[`alipay.open.mini.innerversion.audit.cancel`](https://opendocs.alipay.com/pre-apis/05xaab)
- 路由：`POST /v3/alipay/open/mini/innerversion/audit/cancel`
- 请求类型：[MiniInnerVersionAuditCancelRequest](audit_request.go#L4)
- 响应类型：[MiniInnerVersionAuditCancelResponse](audit_response.go#L4)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属PID |

<a id="api-miniinnerversionauditsubmit"></a>

### MiniInnerVersionAuditSubmit — 小程序版本提交审核

- 官方接口：[`alipay.open.mini.innerversion.audit.submit`](https://opendocs.alipay.com/pre-apis/05xnee)
- 路由：`POST /v3/alipay/open/mini/innerversion/audit/submit`
- 请求类型：[MiniInnerVersionAuditSubmitRequest](audit_request.go#L18)
- 响应类型：[MiniInnerVersionAuditSubmitResponse](audit_response.go#L8)

小程序版本提交审核，审核申请完结后会通过From蚂蚁消息通知到开发者。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | JSON body | 未标注 | 小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考开放服务类目，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。 |
| `app_desc` | `AppDesc` | `*string` | JSON body | 未标注 | 小程序应用描述，20-200个字，如果不填默认采用当前小程序的应用描述 |
| `app_english_name` | `AppEnglishName` | `*string` | JSON body | 未标注 | 小程序应用英文名称，如果不填默认采用当前小程序应用英文名称，3～30个字符。插件不填 |
| `app_logo` | `AppLogo` | `*string` | JSON body | 未标注 | 小程序logo图标，图片格式必须为：png、jpeg、jpg，比例必须为1:1，建议上传像素为180*180，不超过256kb， 如果不填默认采用当前小程序应用logo图标 |
| `app_name` | `AppName` | `*string` | JSON body | 未标注 | 小程序应用名称，如果不填默认采用当前小程序应用名称，名称3-40个字符，一个中文算两个字符 |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源类型，新接入方需要向支付宝申请专用来源，否则不予接入，申请方式请参见接入手册。 |
| `app_slogan` | `AppSlogan` | `*string` | JSON body | 未标注 | 小程序应用简介，一句话描述小程序功能，如果不填默认采用当前小程序应用简介，10~32个字符 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 需要提交审核的小程序版本号，格式为: x.y.z，其中x、y、z均为整型数字，一个应用最多只能有一个审核中、审核通过或者审核驳回的版本。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端ID，多端场景下区分不同端 |
| `license_info` | `LicenseInfo` | [`*AuditLicenseInfo`](#model-auditlicenseinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 小程序备注 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，特殊场景专用，普通业务方无需关注该参数。 |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | JSON body | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 [alipay.open.mini.category.query](https://opendocs.alipay.com/mini/03l8c6)（小程序类目树查询接口）查询mini_category_list，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。使用后不再读取app_category_ids值，老前台类目将废弃 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 小程序所属PID |
| `out_door_pic` | `OutDoorPic` | `*string` | JSON body | 未标注 | 门头照图片地址，部分小程序类目需要提交，参照 开放服务类目 中是否需要营业执照信息，如果不填默认采用当前小程序门头照图片 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属PID |
| `screen_shot_list` | `ScreenShotList` | `*[]string` | JSON body | 未标注 | 小程序应用截图列表，逗号分割，需要2-5张图片，单张图片不能超过4MB，图片格式只支持jpg，png；通过模板实例化小程序不需要传递此参数。 |
| `service_email` | `ServiceEmail` | `*string` | JSON body | 未标注 | 小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，和客服电话至少填写一个 |
| `service_phone` | `ServicePhone` | `*string` | JSON body | 未标注 | 小程序客服电话，如果不填默认采用当前小程序的应用客服电话，和客服邮箱至少填写一个。插件不填 |
| `special_license_pic_list` | `SpecialLicensePicList` | `*string` | JSON body | 未标注 | 特殊资质图片地址列表，逗号分隔；部分类目需要特殊资质，如果需要特殊资质，最少一张，最多三张。模板和插件不需要特殊资质 |
| `version_desc` | `VersionDesc` | `*string` | JSON body | 未标注 | 小程序版本变更描述，30-500个字符，区分于app_desc |

## 嵌套请求参数

<a id="model-auditlicenseinfo"></a>

### AuditLicenseInfo

类型定义：[AuditLicenseInfo](audit_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `license_name` | `LicenseName` | `*string` | 未标注 | 营业执照名称，需要与营业执照保持一致 |
| `license_no` | `LicenseNo` | `*string` | 未标注 | 营业执照号，部分小程序类目需要提交，参照 开放服务类目 中是否需要营业执照信息，如果不填默认采用当前小程序应用营业执照号。 |
| `license_pic_list` | `LicensePicList` | `*[]string` | 未标注 | 营业执照照片地址列表 |
| `license_valid_date` | `LicenseValidDate` | `*string` | 未标注 | 营业执有效期，格式为yyyy-MM-dd，9999-12-31表示长期 |
| `out_door_pic` | `OutDoorPic` | `*string` | 未标注 | 门头照图片地址，部分小程序类目需要提交，参照 开放服务类目 中是否需要营业执照信息，如果不填默认采用当前小程序门头照图片 |
