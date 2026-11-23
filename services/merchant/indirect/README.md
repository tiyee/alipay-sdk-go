# indirect 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/merchant/indirect`。本包提供 4 个支付宝 V3 REST 操作。

通过 `indirect.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`IndirectAuthOrderClose`](#api-indirectauthorderclose) | `alipay.merchant.indirect.authorder.close` |
| [`IndirectAuthOrderCreate`](#api-indirectauthordercreate) | `alipay.merchant.indirect.authorder.create` |
| [`IndirectAuthOrderQueryStatus`](#api-indirectauthorderquerystatus) | `alipay.merchant.indirect.authorder.querystatus` |
| [`IndirectSMIDBindQuery`](#api-indirectsmidbindquery) | `alipay.merchant.indirect.smidbind.query` |

## 使用示例

- [IndirectAuthOrderClose](api_indirect_indirectauthorderclose_example_test.go)
- [IndirectAuthOrderCreate](api_indirect_indirectauthordercreate_example_test.go)
- [IndirectAuthOrderQueryStatus](api_indirect_indirectauthorderquerystatus_example_test.go)
- [IndirectSMIDBindQuery](api_indirect_indirectsmidbindquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-indirectauthorderclose"></a>

### IndirectAuthOrderClose — 商家认证申请单撤销

- 官方接口：[`alipay.merchant.indirect.authorder.close`](https://opendocs.alipay.com/apis/063knt)
- 路由：`PUT /v3/alipay/merchant/indirect/authorder/close`
- 请求类型：[MerchantIndirectAuthOrderCloseRequest](indirect_request.go#L4)
- 响应类型：[MerchantIndirectAuthOrderCloseResponse](indirect_response.go#L8)

申请单处于审核中、审核失败、待确认等状态时，申请单可撤销，撤销后可重新调用提交申请单接口上传申请单信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 意愿申请单号，参数二选一 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，参数二选一，业务自定义，保证唯一 |

<a id="api-indirectauthordercreate"></a>

### IndirectAuthOrderCreate — 商家认证申请单提交

- 官方接口：[`alipay.merchant.indirect.authorder.create`](https://opendocs.alipay.com/apis/05gz1v)
- 路由：`POST /v3/alipay/merchant/indirect/authorder/create`
- 请求类型：[MerchantIndirectAuthOrderCreateRequest](indirect_request.go#L12)
- 响应类型：[MerchantIndirectAuthOrderCreateResponse](indirect_response.go#L11)

间连商户认证申请单提交，用于支付宝间连商家认证

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_identity_info` | `AuthIdentityInfo` | [`*AuthIdentityInfo`](#model-authidentityinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `benefit_person_info` | `BenefitPersonInfo` | [`*IndirectBenefitPersonInfo`](#model-indirectbenefitpersoninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `contact_person_info` | `ContactPersonInfo` | [`*IndirectContactPersonInfo`](#model-indirectcontactpersoninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `extra_credentials` | `ExtraCredentials` | [`*IndirectExtraCredentials`](#model-indirectextracredentials) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `legal_person_info` | `LegalPersonInfo` | [`*IndirectLegalPersonInfo`](#model-indirectlegalpersoninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号。支持out_biz_no和outBizNo两种格式 |
| `source` | `Source` | `*string` | JSON body | 未标注 | 1. 收单机构调用API提交申请单时，可选择是否指定单个服务商范围。非收单机构无需填写此字段。 2. 此字段填写单个服务商pid信息：填写（即：单服务商提交认证方式），查询申请单返回的认证二维码qr_code和填写服务商对应，仅能认证填写服务商下的商户；不填写（即：全服务商提交认证方式），查询申请单返回的认证二维码qr_code和收单机构对应，可认证收单机构下全部商户。 |

<a id="api-indirectauthorderquerystatus"></a>

### IndirectAuthOrderQueryStatus — 查询商家认证申请单状态

- 官方接口：[`alipay.merchant.indirect.authorder.querystatus`](https://opendocs.alipay.com/apis/063luk)
- 路由：`POST /v3/alipay/merchant/indirect/authorder/querystatus`
- 请求类型：[MerchantIndirectAuthOrderQueryStatusRequest](indirect_request.go#L25)
- 响应类型：[MerchantIndirectAuthOrderQueryStatusResponse](indirect_response.go#L19)

查询商家认证申请单状态，以及失败原因

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 商家认证申请单号，参数二选一 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，参数二选一，业务自定义，保证唯一 |
| `source` | `Source` | `*string` | JSON body | 未标注 | 1. 收单机构调用API提交申请单时，可选择是否指定单个服务商范围。非收单机构无需填写此字段。 2. 此字段填写单个服务商pid信息：填写（即：单服务商提交认证方式），查询申请单返回的认证二维码qr_code和填写服务商对应，仅能认证填写服务商下的商户；不填写（即：全服务商提交认证方式），查询申请单返回的认证二维码qr_code和收单机构对应，可认证收单机构下全部商户。 |

<a id="api-indirectsmidbindquery"></a>

### IndirectSMIDBindQuery — 获取商家认证状态

- 官方接口：[`alipay.merchant.indirect.smidbind.query`](https://opendocs.alipay.com/apis/063ec9)
- 路由：`GET /v3/alipay/merchant/indirect/smidbind/query`
- 请求类型：[MerchantIndirectSMIDBindQueryRequest](indirect_request.go#L35)
- 响应类型：[MerchantIndirectSMIDBindQueryResponse](indirect_response.go#L31)

获取商家认证状态，状态枚举：已确认/未确认/已注销

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `sub_merchant_id` | `SubMerchantId` | `*string` | query | 否 | 支付宝商户id |

## 嵌套请求参数

<a id="model-authidentityinfo"></a>

### AuthIdentityInfo

类型定义：[AuthIdentityInfo](indirect_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `certificate_info` | `CertificateInfo` | [`*IndirectCertificateInfo`](#model-indirectcertificateinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `certificate_type` | `CertificateType` | `*string` | 未标注 | 主体类型为企业/个体户/政府机关/事业单位/社会组织时，必填； 证照类型：营业执照(BUSINESS_CERT)/登记证书(REGISTER_CERT) 主体为政府机关/事业单位/社会组织时，填登记证书； 主体类型为企业/个体户时，填营业执照； |
| `employer_letter_img` | `EmployerLetterImg` | `*string` | 未标注 | 单位证明函照片（使用图片上传接口）主体类型为政府机关/事业单位时，单位证明函照片必填 |
| `financial_org_info` | `FinancialOrgInfo` | [`*IndirectFinancialOrgInfo`](#model-indirectfinancialorginfo) | 未标注 | 见字段类型及官方接口文档。 |
| `identity_type` | `IdentityType` | `*string` | 未标注 | 主体类型，枚举定义：企业(ENTERPRISE)、个体工商户(IND_BIZ)、事业单位(INST)、党政机关(GOV)、社会组织(ORG)、小微商户(MSE) |
| `is_financial_org` | `IsFinancialOrg` | `*bool` | 未标注 | 是否金融机构 |
| `qualification_info_list` | `QualificationInfoList` | [`*[]IndirectQualificationInfo`](#model-indirectqualificationinfo) | 未标注 | 经营许可证列表，填写特殊行业的经营许可证信息，一个主体最多5个行业 |
| `support_credentials` | `SupportCredentials` | [`*IndirectSupportCredentials`](#model-indirectsupportcredentials) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-indirectbenefitpersoninfo"></a>

### IndirectBenefitPersonInfo

类型定义：[IndirectBenefitPersonInfo](indirect_model.go#L29)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `card_back_img` | `CardBackImg` | `*string` | 未标注 | 受益人证件反面照（使用图片上传接口的image_id）。当证件类型为身份证时，此字段必传；当证件类型为非身份证时，此字段非必传。 |
| `card_front_img` | `CardFrontImg` | `*string` | 未标注 | 受益人证件正面照（使用图片上传接口的image_id） |
| `card_no` | `CardNo` | `*string` | 未标注 | 证件号码 |
| `card_type` | `CardType` | `*string` | 未标注 | 证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型，政府机关、小微商户仅支持身份证类型。 |
| `effect_time` | `EffectTime` | `*string` | 未标注 | 证件生效时间 |
| `expire_time` | `ExpireTime` | `*string` | 未标注 | 证件过期时间 |
| `person_name` | `PersonName` | `*string` | 未标注 | 受益人姓名 |

<a id="model-indirectcertificateinfo"></a>

### IndirectCertificateInfo

类型定义：[IndirectCertificateInfo](indirect_model.go#L47)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cert_image` | `CertImage` | `*string` | 未标注 | 证照图片（使用图片上传接口） |
| `cert_no` | `CertNo` | `*string` | 未标注 | 证件编号 |
| `cert_type` | `CertType` | `*string` | 未标注 | 登记证书类型时必填，枚举：统一社会信用代码证书（UNIT_SOCIAL_CREDIT）、慈善组织公开募捐资格证书（CHARITY_ORG_SOLICIT），社会团体法人登记证书（SOCIAL_ORG_LEGAL），民办非企业单位登记证书（CIVIL_UN_ENT），基金会法人登记证书（FOUNDATION_LEGAL_PERSON），农民专业合作社法人营业执照（FARMERS_COOPERATE），宗教活动场所登记证（RELIGION_PLACES），其他证书/批文/证明（OTHER_REG_CERT） |
| `effect_time` | `EffectTime` | `*string` | 未标注 | 证照生效时间 |
| `expire_time` | `ExpireTime` | `*string` | 未标注 | 证照过期时间，yyyy-mm-dd格式，如果是长期有效时传“forever” |
| `legal_person_name` | `LegalPersonName` | `*string` | 未标注 | 证照法人姓名 |
| `merchant_name` | `MerchantName` | `*string` | 未标注 | 证照商户名称 |
| `register_address` | `RegisterAddress` | `*string` | 未标注 | 证照注册地址 |

<a id="model-indirectcontactpersoninfo"></a>

### IndirectContactPersonInfo

类型定义：[IndirectContactPersonInfo](indirect_model.go#L67)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `contact_card_no` | `ContactCardNo` | `*string` | 未标注 | 联系人身份证件号。 可传身份证、港澳居民来往内地通行证、港澳居民居住证、台湾居民来往大陆通行证、台湾居民居住证、护照、外国人永久居住证等证件号码，号码枚举见contact_card_type。若传身份证，需保证18位，且英文字母大写 |
| `contact_cert_type` | `ContactCertType` | `*string` | 未标注 | 联系人证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型。小微商户仅支持身份证类型。 |
| `contact_name` | `ContactName` | `*string` | 未标注 | 联系人姓名 |
| `contact_phone_no` | `ContactPhoneNo` | `*string` | 未标注 | 联系人手机 |

<a id="model-indirectextracredentials"></a>

### IndirectExtraCredentials

类型定义：[IndirectExtraCredentials](indirect_model.go#L79)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `smid_list` | `SmidList` | `*[]string` | 未标注 | 待确认商户号列表，最多20个 |

<a id="model-indirectfinancialorginfo"></a>

### IndirectFinancialOrgInfo

类型定义：[IndirectFinancialOrgInfo](indirect_model.go#L85)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `financial_org_cert_img` | `FinancialOrgCertImg` | `*[]string` | 未标注 | 金融机构许可证图片，最多五张（使用图片上传接口） |
| `financial_org_type` | `FinancialOrgType` | `*string` | 未标注 | 银行业(BANK)、支付机构(PAYMENT)、保险业(INSURANCE)、交易及结算类金融机构(SETTLE)、OTHER(OTHER) |

<a id="model-indirectlegalpersoninfo"></a>

### IndirectLegalPersonInfo

类型定义：[IndirectLegalPersonInfo](indirect_model.go#L93)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `auth_letter_img` | `AuthLetterImg` | `*string` | 未标注 | 授权函照片（使用图片上传接口），当证件持有人类型为 经办人时，必须上传。其他情况，无需上传 |
| `card_back_img` | `CardBackImg` | `*string` | 未标注 | 法人/经营者证件反面照（使用图片上传接口的image_id）。当证件类型为身份证时，此字段必传；当证件类型为非身份证时，此字段非必传。 |
| `card_front_img` | `CardFrontImg` | `*string` | 未标注 | 法人/经营者证件正面照（使用图片上传接口返回的image_id） |
| `card_no` | `CardNo` | `*string` | 未标注 | 证件号码 |
| `card_type` | `CardType` | `*string` | 未标注 | 证件类型，枚举值：中国大陆居民-身份证(RESIDENT)、其他国家或地区居民-护照(PASSPORT)、中国港澳居民-来往内地通行证(PASSPORT_HK_MO)、中国台湾居民-来往大陆通行证(PASSPORT_TWN)、港澳居民居住证（RESIDENCE_PERMIT_HM）、台湾居民居住证（RESIDENCE_PERMIT_TW）、外国人永久居住证（PERMANENT_RESIDENCE_FOREIGNER）。个体户/企业/事业单位/社会组织：可选择任一证件类型，政府机关、小微商户仅支持身份证类型。 |
| `effect_time` | `EffectTime` | `*string` | 未标注 | 证件生效时间 |
| `expire_time` | `ExpireTime` | `*string` | 未标注 | 证件过期时间 |
| `is_benefit_person` | `IsBenefitPerson` | `*bool` | 未标注 | 经营者/法人是否为受益人 |
| `legal_type` | `LegalType` | `*string` | 未标注 | 证件持有人类型： - 法人 LEGAL_PERSON - 经办人 AGENT_PERSON 主体类型为政府机关、事业单位时选传： （1）若上传的是法人证件，则不需要上传该字段。 （2）若因特殊情况，无法提供法人证件时，可上传经办人。 主体类型为企业、个体户、社会组织、小微商户时，默认为经营者/法人，不需要上传该字段。 |
| `person_name` | `PersonName` | `*string` | 未标注 | 法人姓名 |

<a id="model-indirectqualificationinfo"></a>

### IndirectQualificationInfo

类型定义：[IndirectQualificationInfo](indirect_model.go#L117)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `image_list` | `ImageList` | `*[]string` | 未标注 | 行业经营许可证资质照片，一个行业类目下最多上传6张资质照片（使用图片上传接口） |
| `mcc_code` | `MccCode` | `*string` | 未标注 | 行业类目编号，支付宝商家行业二级类目code |

<a id="model-indirectsupportcredentials"></a>

### IndirectSupportCredentials

类型定义：[IndirectSupportCredentials](indirect_model.go#L125)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `city` | `City` | `*string` | 未标注 | 门店城市(该值已废弃，以city_code字段映射的城市名称为准) |
| `city_code` | `CityCode` | `*string` | 未标注 | 门店市行政区号（数字） |
| `district` | `District` | `*string` | 未标注 | 门店所在区县。(该值已废弃，以district_code字段映射的区县名称为准) |
| `district_code` | `DistrictCode` | `*string` | 未标注 | 门店街道区号（数字） |
| `merchant_type` | `MerchantType` | `*string` | 未标注 | 小微商户经营类型，枚举值：门店场所（STORE）、流动经营（STALL） |
| `province` | `Province` | `*string` | 未标注 | 门店省份(该值已废弃，以province_code字段映射的省份名称为准) |
| `province_code` | `ProvinceCode` | `*string` | 未标注 | 门店省行政区号（数字 |
| `store_address` | `StoreAddress` | `*string` | 未标注 | 门店场所填写门店详细地址，流动经营类型填“无” |
| `store_door_img` | `StoreDoorImg` | `*string` | 未标注 | 门店门头照信息或摊位照（使用图片上传接口） |
| `store_inner_img` | `StoreInnerImg` | `*string` | 未标注 | 门店店内照片或者摊位照侧面（使用图片上传接口） |
| `store_name` | `StoreName` | `*string` | 未标注 | 门店名称 |
