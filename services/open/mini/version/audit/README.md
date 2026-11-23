# audit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/version/audit`。本包提供 2 个支付宝 V3 REST 操作。

通过 `audit.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MiniVersionAuditApply`](#api-miniversionauditapply) | `alipay.open.mini.version.audit.apply` |
| [`MiniVersionAuditCancel`](#api-miniversionauditcancel) | `alipay.open.mini.version.audit.cancel` |

## 使用示例

- [MiniVersionAuditApply](api_audit_miniversionauditapply_example_test.go)
- [MiniVersionAuditCancel](api_audit_miniversionauditcancel_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniversionauditapply"></a>

### MiniVersionAuditApply — 小程序提交审核

- 官方接口：[`alipay.open.mini.version.audit.apply`](https://opendocs.alipay.com/open-v3/dd9959d2_alipay.open.mini.version.audit.apply)
- 路由：`POST /v3/alipay/open/mini/version/audit/apply`
- 请求类型：[MiniVersionAuditApplyRequest](audit_request.go#L16)
- 响应类型：[MiniVersionAuditApplyResponse](audit_response.go#L4)

小程序提交审核。
1、本接口中关于小程序基础信息的一些字段例如小程序名称、小程序logo、客服电话等，这些信息如果你前期已经设置过，平台将会在你本次调用接口的时候，将上一次设置的信息自动填入，所以允许你本次调用时不传入。如果你前期还未设置过这些信息，则本次调用是需要必传的。具体的字段要求详见每个字段的描述。你可以用[alipay.open.mini.version.detail.query](https://opendocs.alipay.com/apis/05ujoy)接口查询小程序信息之前是否已经设置。
2、小程序版本信息例如小程序版本号、小程序版本描述，这些信息是需要你每次调用接口必传的。
3、本接口涉及到“审核流程拆分”、“审核加急”、“自动上架”的功能，具体功能详见接口字段audit_rule、speed_up、auto_online描述。
4、普通模板构建的商家小程序限 5个小程序/模板/日，可以升级为标准版，标准版没有数量限制。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_logo` | `AppLogo` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*MiniVersionAuditApply`](#model-miniversionauditapply) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `eighth_special_license_pic` | `EighthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fifth_license_pic` | `FifthLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fifth_screen_shot` | `FifthScreenShot` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fifth_special_license_pic` | `FifthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `first_license_pic` | `FirstLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `first_screen_shot` | `FirstScreenShot` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `first_special_license_pic` | `FirstSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fourth_license_pic` | `FourthLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fourth_screen_shot` | `FourthScreenShot` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `fourth_special_license_pic` | `FourthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `ninth_special_license_pic` | `NinthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `out_door_pic` | `OutDoorPic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `second_license_pic` | `SecondLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `second_screen_shot` | `SecondScreenShot` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `second_special_license_pic` | `SecondSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `seventh_special_license_pic` | `SeventhSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `sixth_special_license_pic` | `SixthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `tenth_special_license_pic` | `TenthSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `test_file_name` | `TestFileName` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `third_license_pic` | `ThirdLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `third_screen_shot` | `ThirdScreenShot` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `third_special_license_pic` | `ThirdSpecialLicensePic` | [`*core.MultipartFile`](../../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-miniversionauditcancel"></a>

### MiniVersionAuditCancel — 小程序撤销审核

- 官方接口：[`alipay.open.mini.version.audit.cancel`](https://opendocs.alipay.com/open-v3/9b50e79c_alipay.open.mini.version.audit.cancel)
- 路由：`POST /v3/alipay/open/mini/version/audit/cancel`
- 请求类型：[MiniVersionAuditCancelRequest](audit_request.go#L8)
- 响应类型：[MiniVersionAuditCancelResponse](audit_response.go#L12)

三方实例化小程序撤销审核

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序审核中的版本号, 不传默认撤消正在审核中的版本。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

## 嵌套请求参数

<a id="model-miniversionauditapply"></a>

### MiniVersionAuditApply

类型定义：[MiniVersionAuditApply](audit_model.go#L8)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_category_ids` | `AppCategoryIds` | `*string` | 未标注 | 小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以通过 [https://docs.open.alipay.com/api_49/alipay.open.mini.category.query](https://docs.open.alipay.com/api_49/alipay.open.mini.category.query)接口查询，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。 |
| `app_desc` | `AppDesc` | `*string` | 未标注 | 小程序描述，长度限制 20~400 个字符（中文字符按2个字符计算）。 |
| `app_english_name` | `AppEnglishName` | `*string` | 未标注 | 小程序应用英文名称，长度限制 3～30 个字符（中文字符按2个字符计算）；仅支持英文和数字。 |
| `app_name` | `AppName` | `*string` | 未标注 | 小程序名称，长度限制 3~40 个字符（中文字符按2个字符计算），仅支持包含中文、数字、英文、下划线、+、-。 如果不填默认采用当前小程序应用名称，如无默认值则必填，不可为空。 |
| `app_slogan` | `AppSlogan` | `*string` | 未标注 | 小程序简介，一句话描述小程序功能，长度限制 10~32个字符（中文字符按2个字符计算）。 |
| `app_version` | `AppVersion` | `*string` | 未标注 | 小程序版本号，请选择开发版本执行提交审核操作。 |
| `audit_rule` | `AuditRule` | `*string` | 未标注 | 审核类型，建议设置成“BASE_PROMOTE”，小程序将会更容易上架。将小程序发布审核拆分成“准入”、“营销”。原先小程序需要两个环节都审核通过才能上架，现在小程序只需要准入通过即可上架，但是无法在营销场景（支付宝泛搜、支付宝首页推荐等）展现给用户，小程序可通过精搜或者扫码等方式使用。如果小程序需要进行营销，可发布新版本，当准入&营销都通过后，小程序也可被营销。点击查看支付宝小程序营销规范：[https://opendocs.alipay.com/b/03al2q](https://opendocs.alipay.com/b/03al2q) |
| `auto_online` | `AutoOnline` | `*string` | 未标注 | 审核通过后是否自动上架， 自动上架：true；不自动上架（默认）：false。 如果小程序上架前无需特殊处理，建议设置成true，平台将会在小程序审核通过时自动将其上架。 注意：若审核通过时，小程序仍未完成备案，则将会进行自动受限发布上架（发布上线后小程序将无法被搜索，限制小程序用户数为 50 人/天，无法进行相关业务推广，只可通过扫码使用，完成备案即可解除限制，上架后90天内不未完成备案，小程序将无法打开。）若审核通过时小程序也完成备案，则会正常上架。 |
| `bundle_id` | `BundleId` | `*string` | 未标注 | 小程序投放的端参数。例如投放到支付宝钱包是支付宝端。默认支付宝端。支持： com.alipay.alipaywallet:支付宝端； com.alipay.iot.xpaas：支付宝IoT端。 |
| `license_name` | `LicenseName` | `*string` | 未标注 | 营业执照名称，参照[https://opendocs.alipay.com/mini/operation/material](https://opendocs.alipay.com/mini/operation/material),要求填写营业执照信息。 |
| `license_no` | `LicenseNo` | `*string` | 未标注 | 营业执照证件号，部分小程序类目需要提交，可通过[https://opendocs.alipay.com/mini/03l8c8](https://opendocs.alipay.com/mini/03l8c8 )查询类目是否需要上传营业执照信息。参照[https://opendocs.alipay.com/mini/operation/material ](https://opendocs.alipay.com/mini/operation/material )要求填写营业执照信息。 |
| `license_valid_date` | `LicenseValidDate` | `*string` | 未标注 | 营业执照有效期，格式为 yyyy-MM-dd，9999-12-31表示长期有效。 |
| `memo` | `Memo` | `*string` | 未标注 | 小程序备注，小程序备注最多500字符（中文字符按2个字符计算）。 |
| `mini_category_ids` | `MiniCategoryIds` | `*string` | 未标注 | 新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 [https://docs.open.alipay.com/api_49/alipay.open.mini.category.query](https://docs.open.alipay.com/api_49/alipay.open.mini.category.query) 接口查询mini_category_list。 |
| `region_type` | `RegionType` | `*string` | 未标注 | 小程序服务区域类型 |
| `service_email` | `ServiceEmail` | `*string` | 未标注 | 小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 |
| `service_phone` | `ServicePhone` | `*string` | 未标注 | 小程序客服电话，长度限制5~30个字符，仅支持包含数字和-。如果不填默认采用当前小程序的应用客服电话，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 |
| `service_region_info` | `ServiceRegionInfo` | [`*[]RegionInfo`](../README.md#model-regioninfo) | 未标注 | 省市区信息。当region_type为LOCATION或传入city_code时，province_code不能为空；填写area_code时，province_code和city_code不能为空。只填province_code则全选该省；填写province_code和city_code则全选该市，以此类推。省市区code参见[https://gw.alipayobjects.com/os/bmw-prod/0aab0319-13de-42b9-85cf-13877a5f78ed.xlsx](https://gw.alipayobjects.com/os/bmw-prod/0aab0319-13de-42b9-85cf-13877a5f78ed.xlsx) |
| `speed_up` | `SpeedUp` | `*string` | 未标注 | 如果有审核加急权益，是否使用审核加急权益加速审核： 加速（默认）：true 不加速：false |
| `test_accout` | `TestAccout` | `*string` | 未标注 | 测试账号，是否需要填写请参见[https://opendocs.alipay.com/mini/operation/standard/case/akxg6r#3.%20%E6%B5%8B%E8%AF%95%E5%86%85%E5%AE%B9%E6%8F%90%E4%BA%A4%E4%B8%8D%E5%AE%8C%E6%95%B4](https://opendocs.alipay.com/mini/operation/standard/case/akxg6r#3.%20%E6%B5%8B%E8%AF%95%E5%86%85%E5%AE%B9%E6%8F%90%E4%BA%A4%E4%B8%8D%E5%AE%8C%E6%95%B4) |
| `test_password` | `TestPassword` | `*string` | 未标注 | 测试账号密码 |
| `version_desc` | `VersionDesc` | `*string` | 未标注 | 小程序版本描述，介绍此版本的主要变更和功能，30-500个字符（中文字符按2个字符计算）。 |
