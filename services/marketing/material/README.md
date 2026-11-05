# marketing/material 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/material`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `material.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_material_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`MaterialCreate`](#api-materialcreate) | [`alipay.marketing.material.create`](https://opendocs.alipay.com/apis/05xk1p) | 创建支付宝素材 | `POST` |
| [`MaterialImageUpload`](#api-materialimageupload) | [`alipay.marketing.material.image.upload`](https://opendocs.alipay.com/open-v3/574939dd_alipay.marketing.material.image.upload) | 营销图片资源上传接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-materialcreate"></a>

### MaterialCreate — 创建支付宝素材

- 官方接口：[`alipay.marketing.material.create`](https://opendocs.alipay.com/apis/05xk1p)
- 路由：`POST /v3/alipay/marketing/material/create`
- 请求类型：[MaterialCreateRequest](material_request.go#L9)
- 响应类型：[MaterialCreateResponse](material_response.go#L5)

用于创建支付宝素材

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `material_fields` | `MaterialFields` | [`*[]MaterialField`](#model-materialfield) | JSON body | 未标注 | 素材字段列表，素材提报的具体内容，包含三种类型：图片、文本、链接。图片类型，则根据“图片资源上传接口”中返回的resource_id字段值进行素材字段的设值。文本或链接类型，则按照实际需要设值即可，需满足具体的规范要求。规范要求详情请参见产品文档。 |
| `material_name` | `MaterialName` | `*string` | JSON body | 未标注 | 素材名称，接入方自行生成，建议命名时遵循一定的规律以方便管理。 |
| `material_spec_id` | `MaterialSpecId` | `*string` | JSON body | 未标注 | 素材规范ID，用于标识具体场景下的素材规范，根据具体的使用场景获取对应的素材规范ID。 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务单号，用作幂等控制，相同单号会返回上一次的结果，接入方需保证单号唯一。 |

<a id="api-materialimageupload"></a>

### MaterialImageUpload — 营销图片资源上传接口

- 官方接口：[`alipay.marketing.material.image.upload`](https://opendocs.alipay.com/open-v3/574939dd_alipay.marketing.material.image.upload)
- 路由：`POST /v3/alipay/marketing/material/image`
- 请求类型：[MaterialImageUploadRequest](material_request.go#L21)
- 响应类型：[MaterialImageUploadResponse](material_response.go#L11)

营销通用上传图片接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `data` | `Data` | [`*MarketingMaterialImageUpload`](#model-marketingmaterialimageupload) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `file_content` | `FileContent` | [`*core.MultipartFile`](../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-belongmerchantinfo"></a>

### BelongMerchantInfo

类型定义：[BelongMerchantInfo](material_model.go#L14)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_type` | `BusinessType` | `*string` | 未标注 | 合作业务类型,默认为ISV_FOR_MERCHANT |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户PID,默认为当前接口调用商户。 |
| `merchant_id_type` | `MerchantIdType` | `*string` | 未标注 | 商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。 |

<a id="model-marketingmaterialimageupload"></a>

### MarketingMaterialImageUpload

类型定义：[MarketingMaterialImageUpload](material_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `belong_merchant_info` | `BelongMerchantInfo` | [`*BelongMerchantInfo`](#model-belongmerchantinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `file_key` | `FileKey` | `*string` | 未标注 | 文件业务标识。 枚举值 alipay.marketing.activity.delivery.create接口中 delivery_base_info.delivery_material.delivery_single_material.delivery_image 当delivery_booth_code=PUBLIC_UNION，上传图片接口需指定file_key=PUBLIC_UNION_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb； 当delivery_booth_code=PAYMENT_RESULT，上传图片接口需指定file_key=DELIVERY_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb。 上传图片更多要求参考文档： https://render.alipay.com/p/c/18tpirlg12e8?operateFrom=BALIPAY alipay.marketing.activity.ordervoucher.create接口中 voucher_display_info.brand_logo字段,file_key=PROMO_BRAND_LOGO，上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB voucher_display_info.voucher_image字段,file_key=PROMO_VOUCHER_IMAGE,上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB |
| `merchant_access_mode` | `MerchantAccessMode` | `*string` | 未标注 | 商户接入模式 |

<a id="model-materialfield"></a>

### MaterialField

类型定义：[MaterialField](material_model.go#L24)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `field_name` | `FieldName` | `*string` | 未标注 | 素材字段名称，来源于素材规范定义的字段名称。注意一次素材提报中字段名称不能重复。 |
| `field_value` | `FieldValue` | `*[]string` | 未标注 | 素材字段值，素材字段提报的实际值，支持多值，具体的要求请查看素材规范或对应的产品文档。 |
