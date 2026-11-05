# sp 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/sp`。本包提供 3 个支付宝 V3 REST 操作。

通过 `sp.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`blueseaactivity`](blueseaactivity/README.md) | 3 | `SpBlueSeaActivityCreate` / `SpBlueSeaActivityModify` / `SpBlueSeaActivityQuery` |
| [`isv`](isv/README.md) | 3 | `SpISVBillSync` / `SpISVSignAuthCreate` / `SpISVSignAuthQuery` |
| [`operation`](operation/README.md) | 3 | `SpOperationApply` / `SpOperationQRCodeQuery` / `SpOperationResultQuery` |
| [`oppor`](oppor/README.md) | 2 | `SpOpporDetailQuery` / `SpOpporPageQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SpImageUpload`](#api-spimageupload) | `alipay.open.sp.image.upload` |
| [`SpLeadsExpandCreate`](#api-spleadsexpandcreate) | `alipay.open.sp.leads.expand.create` |
| [`SpMerchantInconsistentApprove`](#api-spmerchantinconsistentapprove) | `alipay.open.sp.merchant.inconsistent.approve` |

## 使用示例

- [SpImageUpload](api_sp_spimageupload_example_test.go)
- [SpLeadsExpandCreate](api_sp_spleadsexpandcreate_example_test.go)
- [SpMerchantInconsistentApprove](api_sp_spmerchantinconsistentapprove_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-spimageupload"></a>

### SpImageUpload — 图片上传接口

- 官方接口：[`alipay.open.sp.image.upload`](https://opendocs.alipay.com/open-v3/59220073_alipay.open.sp.image.upload)
- 路由：`POST /v3/alipay/open/sp/image/upload`
- 请求类型：[SpImageUploadRequest](sp_request.go#L68)
- 响应类型：[SpImageUploadResponse](sp_response.go#L4)

图片上传

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `image_content` | `ImageContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-spleadsexpandcreate"></a>

### SpLeadsExpandCreate — 在线写入商机库

- 官方接口：[`alipay.open.sp.leads.expand.create`](https://opendocs.alipay.com/apis/05vh1e)
- 路由：`POST /v3/alipay/open/sp/leads/expand/create`
- 请求类型：[SpLeadsExpandCreateRequest](sp_request.go#L8)
- 响应类型：[SpLeadsExpandCreateResponse](sp_response.go#L10)

在线写入商机库

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | JSON body | 未标注 | 拓展助手商机地址 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 拓展助手商机地址市code |
| `city_name` | `CityName` | `*string` | JSON body | 未标注 | 拓展助手商机地址市名称 |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 拓展助手商机地址区code |
| `district_name` | `DistrictName` | `*string` | JSON body | 未标注 | 拓展助手商机地址区名称 |
| `email` | `Email` | `*string` | JSON body | 未标注 | 拓展助手商机邮箱 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 拓展助手商机地址纬度 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 拓展助手商机地址经度 |
| `mcc` | `Mcc` | `*string` | JSON body | 未标注 | mcc行业 |
| `merchant_pid` | `MerchantPid` | `*string` | JSON body | 未标注 | 2088账号 |
| `name` | `Name` | `*string` | JSON body | 未标注 | 拓展助手商机名称 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部幂等唯一键 |
| `phone` | `Phone` | `*string` | JSON body | 未标注 | 拓展助手商机联系电话 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 拓展助手商机地址省code |
| `province_name` | `ProvinceName` | `*string` | JSON body | 未标注 | 拓展助手商机地址省名称 |
| `scene` | `Scene` | `*string` | JSON body | 未标注 | 拓展助手场景，参数约定 |

<a id="api-spmerchantinconsistentapprove"></a>

### SpMerchantInconsistentApprove — 非同人商家进件资料审批

- 官方接口：[`alipay.open.sp.merchant.inconsistent.approve`](https://opendocs.alipay.com/apis/05v8ke)
- 路由：`POST /v3/alipay/open/sp/merchant/inconsistent/approve`
- 请求类型：[SpMerchantInconsistentApproveRequest](sp_request.go#L44)
- 响应类型：[SpMerchantInconsistentApproveResponse](sp_response.go#L16)

针对服务优选扫码点餐小程序，用于收集并审核非同人商家的进件资料。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `back_card` | `BackCard` | `*string` | JSON body | 未标注 | 身份证背面照，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 |
| `commitment_letter` | `CommitmentLetter` | `*string` | JSON body | 未标注 | 实际经营人承诺函照片，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。 |
| `front_card` | `FrontCard` | `*string` | JSON body | 未标注 | 身份证正面照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式. 请传[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 |
| `handheld_business_license` | `HandheldBusinessLicense` | `*string` | JSON body | 未标注 | 手持营业执照合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。 |
| `handheld_card` | `HandheldCard` | `*string` | JSON body | 未标注 | 手持身份证合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。 |
| `item_code` | `ItemCode` | `*string` | JSON body | 未标注 | 服务优选商品编码 |
| `merchant_pid` | `MerchantPid` | `*string` | JSON body | 未标注 | 商家pid |
| `mini_appid` | `MiniAppid` | `*string` | JSON body | 未标注 | 商家小程序appId |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，需不重复 |
| `promotor_pid` | `PromotorPid` | `*string` | JSON body | 未标注 | 推广服务商(S2)pid |
