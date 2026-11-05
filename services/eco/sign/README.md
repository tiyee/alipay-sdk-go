# eco/sign 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/sign`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `sign.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_sign_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`SignFlowCancel`](#api-signflowcancel) | [`alipay.eco.sign.flow.cancel`](https://opendocs.alipay.com/pre-apis/05wvs6) | 签署流程撤销（E签宝） | `POST` |
| [`SignFlowCreate`](#api-signflowcreate) | [`alipay.eco.sign.flow.create`](https://opendocs.alipay.com/pre-apis/05xsf2) | 创建流程（E签宝） | `POST` |
| [`SignFlowFinish`](#api-signflowfinish) | [`alipay.eco.sign.flow.finish`](https://opendocs.alipay.com/pre-apis/05wuhl) | 签署流程归档 | `POST` |
| [`SignFlowQuery`](#api-signflowquery) | [`alipay.eco.sign.flow.query`](https://opendocs.alipay.com/pre-apis/05wvxl) | 签署流程查询（E签宝） | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-signflowcancel"></a>

### SignFlowCancel — 签署流程撤销（E签宝）

- 官方接口：[`alipay.eco.sign.flow.cancel`](https://opendocs.alipay.com/pre-apis/05wvs6)
- 路由：`POST /v3/alipay/eco/sign/flow/cancel`
- 请求类型：[SignFlowCancelRequest](sign_request.go#L5)
- 响应类型：[SignFlowCancelResponse](sign_response.go#L9)

签署流程撤销（E签宝）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `flow_id` | `FlowId` | `*string` | JSON body | 未标注 | 流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。 |
| `revoke_reason` | `RevokeReason` | `*string` | JSON body | 未标注 | 撤销原因，默认为"撤销"。商家可自定义。 |

<a id="api-signflowcreate"></a>

### SignFlowCreate — 创建流程（E签宝）

- 官方接口：[`alipay.eco.sign.flow.create`](https://opendocs.alipay.com/pre-apis/05xsf2)
- 路由：`POST /v3/alipay/eco/sign/flow/create`
- 请求类型：[SignFlowCreateRequest](sign_request.go#L15)
- 响应类型：[SignFlowCreateResponse](sign_response.go#L12)

创建流程（E签宝）。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `attachments` | `Attachments` | [`*[]Attachment`](#model-attachment) | JSON body | 未标注 | 附件信息 |
| `business_scene` | `BusinessScene` | `*string` | JSON body | 未标注 | 流程主题 |
| `config_info` | `ConfigInfo` | [`*ConfigInfo`](#model-configinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `template_infos` | `TemplateInfos` | [`*[]TemplateInfo`](#model-templateinfo) | JSON body | 未标注 | 模板信息 |

<a id="api-signflowfinish"></a>

### SignFlowFinish — 签署流程归档

- 官方接口：[`alipay.eco.sign.flow.finish`](https://opendocs.alipay.com/pre-apis/05wuhl)
- 路由：`POST /v3/alipay/eco/sign/flow/finish`
- 请求类型：[SignFlowFinishRequest](sign_request.go#L28)
- 响应类型：[SignFlowFinishResponse](sign_response.go#L18)

手动归档签署流程，归档后所有资源不可修改，归档前签署流程中的所有签署人必须都签署完成。如果创建流程时，设置了自动归档，则无需调用本接口，签署完成后系统会自动调用。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `flow_id` | `FlowId` | `*string` | JSON body | 未标注 | 流程id |

<a id="api-signflowquery"></a>

### SignFlowQuery — 签署流程查询（E签宝）

- 官方接口：[`alipay.eco.sign.flow.query`](https://opendocs.alipay.com/pre-apis/05wvxl)
- 路由：`GET /v3/alipay/eco/sign/flow/query`
- 请求类型：[SignFlowQueryRequest](sign_request.go#L36)
- 响应类型：[SignFlowQueryResponse](sign_response.go#L21)

签署流程查询（E签宝）。可通过此接口查询流程、签署人的签署状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `flow_id` | `FlowId` | `*string` | query | 否 | 流程id，通过 创建电子合同签署流程(alipay.eco.contract.signflows.create)接口获取。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-attachment"></a>

### Attachment

类型定义：[Attachment](sign_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `attachment_name` | `AttachmentName` | `*string` | 未标注 | 流程附件名称 |
| `file_id` | `FileId` | `*string` | 未标注 | 流程附件id |

<a id="model-attachmentexplain"></a>

### AttachmentExplain

类型定义：[AttachmentExplain](sign_model.go#L13)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `allow_more_uploads` | `AllowMoreUploads` | `*bool` | 未标注 | 允许上传更多附件。开启后，除了“附件列表”中指定的内容，签署人可自由上传之多10张图片。默认值为false |
| `description` | `Description` | `*string` | 未标注 | 描述 |
| `details` | `Details` | [`*[]Detail`](#model-detail) | 未标注 | 附件列表 |
| `title` | `Title` | `*string` | 未标注 | 附件标题（collect_attachement值为true时必填 |

<a id="model-configinfo"></a>

### ConfigInfo

类型定义：[ConfigInfo](sign_model.go#L25)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `attachment_explain` | `AttachmentExplain` | [`*AttachmentExplain`](#model-attachmentexplain) | 未标注 | 见字段类型及官方接口文档。 |
| `collect_attachement` | `CollectAttachement` | `*bool` | 未标注 | 是否收集附件 |
| `collect_cert_types` | `CollectCertTypes` | `*[]string` | 未标注 | 收集证件类型列表 |
| `company_no` | `CompanyNo` | `*string` | 未标注 | 公司编号-SF |
| `contract_validity` | `ContractValidity` | `*int64` | 未标注 | 文档过期时间 |
| `jump_url` | `JumpUrl` | `*string` | 未标注 | 支付宝小程序跳转 |
| `merchant_mini_sign_url` | `MerchantMiniSignUrl` | `*string` | 未标注 | 商户小程序签署地址（signModel字段值为1时 必填） |
| `notice_developer_url` | `NoticeDeveloperUrl` | `*string` | 未标注 | 回调地址，签署的过程和签署完成都会回调。 |
| `platform_order_no` | `PlatformOrderNo` | `*string` | 未标注 | 平台订单号 |
| `serial_no` | `SerialNo` | `*string` | 未标注 | 物流单号 |
| `sign_model` | `SignModel` | `*int64` | 未标注 | 签署方式 0-e签宝小程序签署 1-商户小程序签署 默认0 |

<a id="model-detail"></a>

### Detail

类型定义：[Detail](sign_model.go#L50)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `code` | `Code` | `*string` | 未标注 | 英文描述（collectAttachement值为true时必填） 以下是系统固定code,对上传的图片做相应的校验。 身份证校验情况： 身份证正面照片：ESIGN_IDCARD 身份证背面照片：ESIGN_BACK_IDCARD |
| `description` | `Description` | `*string` | 未标注 | 请上传手机正面照 |
| `order` | `Order` | `*int64` | 未标注 | 顺序（collectAttachement值为true时必填） |

<a id="model-fillcontent"></a>

### FillContent

类型定义：[FillContent](sign_model.go#L60)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `struct_key` | `StructKey` | `*string` | 未标注 | 模板组件自定义key |
| `value` | `Value` | `*string` | 未标注 | 模板值,该值长度取决于配置模板时该字段的限制 |

<a id="model-signfield"></a>

### SignField

类型定义：[SignField](sign_model.go#L80)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auto_execute` | `AutoExecute` | `*string` | 未标注 | 平台自动签 |
| `signer` | `Signer` | [`*Signer`](#model-signer) | 未标注 | 见字段类型及官方接口文档。 |
| `struct_key` | `StructKey` | `*string` | 未标注 | 模板组件id |

<a id="model-signer"></a>

### Signer

类型定义：[Signer](sign_model.go#L89)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `email` | `Email` | `*string` | 未标注 | 邮箱 |
| `id_number` | `IdNumber` | `*string` | 未标注 | 证件号 |
| `id_type` | `IdType` | `*string` | 未标注 | 证件类型,详见个人证件类型说明（http://open.esign.cn/docs/xy/%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/%E9%99%84%E5%BD%95/%E4%B8%AA%E4%BA%BA%E8%AF%81%E4%BB%B6%E7%B1%BB%E5%9E%8B.html），默认CRED_PSN_CH_IDCARD (证件号不为空，则必填) |
| `mobile` | `Mobile` | `*string` | 未标注 | 手机号 |
| `name` | `Name` | `*string` | 未标注 | 姓名 |
| `user_id` | `UserId` | `*string` | 未标注 | 支付宝userId |

<a id="model-templateinfo"></a>

### TemplateInfo

类型定义：[TemplateInfo](sign_model.go#L105)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `fill_contents` | `FillContents` | [`*[]FillContent`](#model-fillcontent) | 未标注 | 模板填充项 |
| `name` | `Name` | `*string` | 未标注 | 文件名称 |
| `signfields` | `Signfields` | [`*SignField`](#model-signfield) | 未标注 | 见字段类型及官方接口文档。 |
| `template_id` | `TemplateId` | `*string` | 未标注 | 模板id |
