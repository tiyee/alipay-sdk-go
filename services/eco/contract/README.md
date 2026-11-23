# eco/contract 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/contract`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `contract.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_contract_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ContractSignFlowsCreate`](#api-contractsignflowscreate) | [`alipay.eco.contract.signflows.create`](https://opendocs.alipay.com/pre-apis/0681tj) | 创建电子合同签署流程(E签宝) | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-contractsignflowscreate"></a>

### ContractSignFlowsCreate — 创建电子合同签署流程(E签宝)

- 官方接口：[`alipay.eco.contract.signflows.create`](https://opendocs.alipay.com/pre-apis/0681tj)
- 路由：`POST /v3/alipay/eco/contract/signflows/create`
- 请求类型：[ContractSignFlowsCreateRequest](contract_request.go#L5)
- 响应类型：[ContractSignFlowsCreateResponse](contract_response.go#L5)

根据配置好的合同模板，调用接口填充内容并指定签署人创建电子合同签署流程。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `target_app_id` | `TargetAppId` | `*string` | query | 否 | 目标isv应用ID |
| `attachments` | `Attachments` | [`*[]Attachment`](#model-attachment) | JSON body | 未标注 | 附件信息（附件可作为合同的辅助证明材料，无需签署。需通过文件直传接口上传附件，获取文件id后添加）。注意事项参见 [文件流上传方法](https://opendocs.alipay.com/mini/00kr2w)。 |
| `business_scene` | `BusinessScene` | `*string` | JSON body | 未标注 | 流程主题（简要概括合同的签署内容） |
| `config_info` | `ConfigInfo` | [`*ConfigInfoBean`](#model-configinfobean) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `template_infos` | `TemplateInfos` | [`*[]TemplateInfoBean`](#model-templateinfobean) | JSON body | 未标注 | 模板信息（基于合同模板填充内容生成待签文件，并指定签署人） |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-attachment"></a>

### Attachment

类型定义：[Attachment](contract_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `attachment_name` | `AttachmentName` | `*string` | 未标注 | 流程附件名称 |
| `file_id` | `FileId` | `*string` | 未标注 | 流程附件id |

<a id="model-attachmentexplainbean"></a>

### AttachmentExplainBean

类型定义：[AttachmentExplainBean](contract_model.go#L13)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `description` | `Description` | `*string` | 未标注 | 页面展示的任务描（collect_attachement值为true时必填） |
| `details` | `Details` | [`*[]DetailBean`](#model-detailbean) | 未标注 | 附件列表，用于定义每个附件code的描述及页面展示顺序 |
| `title` | `Title` | `*string` | 未标注 | 页面展示的任务标题（collect_attachement值为true时必填） |

<a id="model-configinfobean"></a>

### ConfigInfoBean

类型定义：[ConfigInfoBean](contract_model.go#L23)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `attachment_explain` | `AttachmentExplain` | [`*AttachmentExplainBean`](#model-attachmentexplainbean) | 未标注 | 见字段类型及官方接口文档。 |
| `auto_archive` | `AutoArchive` | `*bool` | 未标注 | 签署完成是否自动归档，默认true，如果false，则归档需要调用签署流程归档接口 |
| `collect_attachement` | `CollectAttachement` | `*bool` | 未标注 | 是否收集附件（需签署人上传身份证或者其他文件的，需设置为true） |
| `company_no` | `CompanyNo` | `*string` | 未标注 | 物流公司编号（目前仅支持顺丰"SF"） 注：避免影响生产订单，联调及测试环境请传入SF_TEST。 |
| `contract_validity` | `ContractValidity` | `*int64` | 未标注 | 合同过期时间：该参数设置的时间若到期，则会触发流程文件过期通知 |
| `jump_url` | `JumpUrl` | `*string` | 未标注 | 签署完成后跳转商户小程序的地址（signModel字段值为2时 选填） |
| `merchant_mini_sign_url` | `MerchantMiniSignUrl` | `*string` | 未标注 | 商户小程序签署地址（signModel字段值为1时 必填） |
| `notice_developer_url` | `NoticeDeveloperUrl` | `*string` | 未标注 | https://esign.cn/callback |
| `platform_order_no` | `PlatformOrderNo` | `*string` | 未标注 | 平台订单号，可将商户订单与电子合同进行关联。用户通过支付宝官方小程序 合同管家 查看合同时，可快速进入商家的小程序及对应的订单详情页，为商家的小程序提供了流量入口，提高用户的活跃度和留存率； |
| `serial_no` | `SerialNo` | `*string` | 未标注 | 12或15位物流单号 |
| `sign_model` | `SignModel` | `*int64` | 未标注 | 快递扫码签收方式： 1-商户小程序：扫码后跳转商户小程序进行签收； 0-e签宝小程序：如商户没有支付宝小程序，可使用e签宝小程序完成签收。 |
| `specify_attachment_infos` | `SpecifyAttachmentInfos` | [`*[]SpecifyAttachmentInfo`](#model-specifyattachmentinfo) | 未标注 | 指定签署人需要上传的附件列表，例如：A、B、C三方签署，A需上传附件1，B需指定附件2，C需上传附件1,2,3 |

<a id="model-detailbean"></a>

### DetailBean

类型定义：[DetailBean](contract_model.go#L50)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `code` | `Code` | `*string` | 未标注 | 自定义附件标识，后续下载附件使用此标识区分。 |
| `description` | `Description` | `*string` | 未标注 | 页面展示的附件中文描述（collectAttachement值为true时必填） |
| `order` | `Order` | `*int64` | 未标注 | 展示顺序（collectAttachement值为true时必填） |

<a id="model-fillcontent"></a>

### FillContent

类型定义：[FillContent](contract_model.go#L60)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `struct_key` | `StructKey` | `*string` | 未标注 | 模板组件自定义key |
| `value` | `Value` | `*string` | 未标注 | 模板值,该值长度取决于配置模板时该字段的限制 |

<a id="model-org"></a>

### Org

类型定义：[Org](contract_model.go#L68)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `org_id_number` | `OrgIdNumber` | `*string` | 未标注 | 企业证件号 |
| `org_id_type` | `OrgIdType` | `*string` | 未标注 | 企业证件类型，（证件号不为空时必填，详见企业证件类型） |
| `org_legal_id_number` | `OrgLegalIdNumber` | `*string` | 未标注 | 法定代表人证件号 |
| `org_legal_id_type` | `OrgLegalIdType` | `*string` | 未标注 | 法定代表人证件类型 |
| `org_legal_name` | `OrgLegalName` | `*string` | 未标注 | 法定代表人名称 |
| `org_name` | `OrgName` | `*string` | 未标注 | 企业名称 |
| `third_party_user_id` | `ThirdPartyUserId` | `*string` | 未标注 | 企业唯一标识：可传入平台机构id、企业证件号、企业邮箱等（个人用户与企业的唯一标识不可重复） |

<a id="model-signfieldbean"></a>

### SignFieldBean

类型定义：[SignFieldBean](contract_model.go#L86)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `sign_field_type` | `SignFieldType` | `*string` | 未标注 | 签署类型： platform-平台自动签：无需指定签署人。创建流程后，系统将自动盖上商户的默认企业印章。 person-个人签署：需要指定签署人个人信息。创建流程后，需通过签署插件完成签署。 org-企业签署：需要指定签署企业信息与经办人个人信息。创建流程后，需经办人代企业完成签署。（企业暂不支持插件签署，可通过获取签署地址接口获取企业签署地址） |
| `signer` | `Signer` | [`*SignerBean`](#model-signerbean) | 未标注 | 见字段类型及官方接口文档。 |
| `struct_key` | `StructKey` | `*string` | 未标注 | 签署区key值 |

<a id="model-signerbean"></a>

### SignerBean

类型定义：[SignerBean](contract_model.go#L95)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `email` | `Email` | `*string` | 未标注 | 个人邮箱 |
| `id_number` | `IdNumber` | `*string` | 未标注 | 个人证件号 |
| `id_type` | `IdType` | `*string` | 未标注 | 个人证件类型（证件号不为空时必填，详见个人证件类型说明 ） |
| `mobile` | `Mobile` | `*string` | 未标注 | 个人手机号 |
| `name` | `Name` | `*string` | 未标注 | 个人姓名 |
| `org` | `Org` | [`*Org`](#model-org) | 未标注 | 见字段类型及官方接口文档。 |
| `third_party_user_id` | `ThirdPartyUserId` | `*string` | 未标注 | 个人唯一标识：可传入平台的个人用户id、支付宝userid、证件号、手机号、邮箱等。 |

<a id="model-specifyattachmentinfo"></a>

### SpecifyAttachmentInfo

类型定义：[SpecifyAttachmentInfo](contract_model.go#L112)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `allow_more_uploads` | `AllowMoreUploads` | `*bool` | 未标注 | 是否允许上传更多附件，默认值为false。开启后，除了“collect_common_codes”中指定的附件，签署人可自由上传至多10张图片。 |
| `allow_more_uploads_max_count` | `AllowMoreUploadsMaxCount` | `*int64` | 未标注 | 允许上传更多附件时，可上传的最大附件数，一次流程中累计的上传附件数至多30（含“附件列表”数量） |
| `collect_cert_types` | `CollectCertTypes` | `*[]string` | 未标注 | 需要签署人上传的证件照片，可进行ORC识别，目前仅支持IDCARD |
| `collect_common_codes` | `CollectCommonCodes` | `*[]string` | 未标注 | 需要签署人上传的其他附件列表，即attachment_explain.details中定义的code列表 |
| `third_party_user_id` | `ThirdPartyUserId` | `*string` | 未标注 | 个人唯一标识，对应签署区中指定的签署人 |

<a id="model-templateinfobean"></a>

### TemplateInfoBean

类型定义：[TemplateInfoBean](contract_model.go#L126)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `fill_contents` | `FillContents` | [`*[]FillContent`](#model-fillcontent) | 未标注 | 模板填充项，根据组件key值传入填写内容 |
| `name` | `Name` | `*string` | 未标注 | 签署文件名称 |
| `signfields` | `Signfields` | [`*[]SignFieldBean`](#model-signfieldbean) | 未标注 | 签署区，根据签署区key值传入对应的签署人信息 |
| `template_id` | `TemplateId` | `*string` | 未标注 | 模板id，通过创建合同模板获取 |
