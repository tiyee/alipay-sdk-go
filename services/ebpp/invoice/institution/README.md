# ebpp/invoice/institution 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/institution`。本目录提供 10 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `institution.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceInstitutionCreate`](#api-invoiceinstitutioncreate) | [`alipay.ebpp.invoice.institution.create`](https://opendocs.alipay.com/apis/05wvvq) | 制度新增 | `POST` |
| [`InvoiceInstitutionDelete`](#api-invoiceinstitutiondelete) | [`alipay.ebpp.invoice.institution.delete`](https://opendocs.alipay.com/apis/05wunr) | 制度删除 | `DELETE` |
| [`InvoiceInstitutionDetailInfoQuery`](#api-invoiceinstitutiondetailinfoquery) | [`alipay.ebpp.invoice.institution.detailinfo.query`](https://opendocs.alipay.com/apis/05j9jg) | 制度详情查询 | `GET` |
| [`InvoiceInstitutionExpenseRuleCreate`](#api-invoiceinstitutionexpenserulecreate) | [`alipay.ebpp.invoice.institution.expenserule.create`](https://opendocs.alipay.com/apis/05waoz) | 创建费控使用规则 | `POST` |
| [`InvoiceInstitutionExpenseRuleDelete`](#api-invoiceinstitutionexpenseruledelete) | [`alipay.ebpp.invoice.institution.expenserule.delete`](https://opendocs.alipay.com/apis/05xa73) | 删除使用规则 | `POST` |
| [`InvoiceInstitutionExpenseRuleModify`](#api-invoiceinstitutionexpenserulemodify) | [`alipay.ebpp.invoice.institution.expenserule.modify`](https://opendocs.alipay.com/apis/05wuq0) | 编辑使用规则 | `POST` |
| [`InvoiceInstitutionModify`](#api-invoiceinstitutionmodify) | [`alipay.ebpp.invoice.institution.modify`](https://opendocs.alipay.com/apis/05wvqu) | 制度编辑 | `POST` |
| [`InvoiceInstitutionPageInfoQuery`](#api-invoiceinstitutionpageinfoquery) | [`alipay.ebpp.invoice.institution.pageinfo.query`](https://opendocs.alipay.com/apis/05xo4v) | 费控制度分页查询 | `GET` |
| [`InvoiceInstitutionScopeModify`](#api-invoiceinstitutionscopemodify) | [`alipay.ebpp.invoice.institution.scope.modify`](https://opendocs.alipay.com/apis/05ip0w) | 修改制度下适配范围 | `POST` |
| [`InvoiceInstitutionScopePageInfoQuery`](#api-invoiceinstitutionscopepageinfoquery) | [`alipay.ebpp.invoice.institution.scopepageinfo.query`](https://opendocs.alipay.com/apis/05j1c2) | 分页查询制度下人员 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceinstitutioncreate"></a>

### InvoiceInstitutionCreate — 制度新增

- 官方接口：[`alipay.ebpp.invoice.institution.create`](https://opendocs.alipay.com/apis/05wvvq)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/create`
- 请求类型：[InvoiceInstitutionCreateRequest](institution_request.go#L5)
- 响应类型：[InvoiceInstitutionCreateResponse](institution_response.go#L5)

制度新增，包含制度下使用规则列表、发放规则列表和适用范围的新增。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `adapter_type` | `AdapterType` | `*string` | JSON body | 未标注 | 制度适配范围(请使用institution_scope_info中adapter_type) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `consult_mode` | `ConsultMode` | `*string` | JSON body | 未标注 | 费控咨询模式，默认为0。 |
| `currency` | `Currency` | `*string` | JSON body | 未标注 | 该字段用于管控制度下发放的资产币种，不填写默认为人民币，设置该字段后，制度下手工发放的资产币种均为设置值。 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 制度生效结束时间 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 制度生效起始时间 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `expense_type` | `ExpenseType` | `*string` | JSON body | 未标注 | 费用类型 |
| `institution_desc` | `InstitutionDesc` | `*string` | JSON body | 未标注 | 制度描述 |
| `institution_name` | `InstitutionName` | `*string` | JSON body | 未标注 | 制度名称 |
| `institution_scope_info` | `InstitutionScopeInfo` | [`*InstitutionScopeInfo`](#model-institutionscopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `issue_rule_info_list` | `IssueRuleInfoList` | [`*[]IssueRuleInfo`](#model-issueruleinfo) | JSON body | 未标注 | 发放规则列表 |
| `multi_employee_share_mode` | `MultiEmployeeShareMode` | `*string` | JSON body | 未标注 | 制度下额度是否支持多人共享。如果支持，则制度下资产所有人共享；如果不支持，则制度下的资产为个人资产，成员之间不共享。 |
| `outer_source_id` | `OuterSourceId` | `*string` | JSON body | 未标注 | 外部唯一标识，可用于防止重复创建 |
| `owner_id_list` | `OwnerIdList` | `*[]string` | JSON body | 未标注 | 适配id列表(请使用institution_scope_info中owner_id_list) |
| `owner_open_id_list` | `OwnerOpenIdList` | `*[]string` | JSON body | 未标注 | 适配开放id列表(仅存量使用员工支付宝id配置适配范围的场景需要该字段，新对接方案不推荐使用员工支付宝id) |
| `owner_type` | `OwnerType` | `*string` | JSON body | 未标注 | 适配id类型(请使用institution_scope_info中owner_type) |
| `payment_mode` | `PaymentMode` | `*string` | JSON body | 未标注 | 费控制度的出资类型 |
| `scene_type` | `SceneType` | `*string` | JSON body | 未标注 | 因公场景(该字段将废弃，不建议使用) |
| `standard_info_list` | `StandardInfoList` | [`*[]StandardInfo`](#model-standardinfo) | JSON body | 未标注 | 使用规则列表 |

<a id="api-invoiceinstitutiondelete"></a>

### InvoiceInstitutionDelete — 制度删除

- 官方接口：[`alipay.ebpp.invoice.institution.delete`](https://opendocs.alipay.com/apis/05wunr)
- 路由：`DELETE /v3/alipay/ebpp/invoice/institution/delete`
- 请求类型：[InvoiceInstitutionDeleteRequest](institution_request.go#L184)
- 响应类型：[InvoiceInstitutionDeleteResponse](institution_response.go#L15)

制度删除

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 （该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `institution_id` | `InstitutionId` | `*string` | query | 否 | 制度id |

<a id="api-invoiceinstitutiondetailinfoquery"></a>

### InvoiceInstitutionDetailInfoQuery — 制度详情查询

- 官方接口：[`alipay.ebpp.invoice.institution.detailinfo.query`](https://opendocs.alipay.com/apis/05j9jg)
- 路由：`GET /v3/alipay/ebpp/invoice/institution/detailinfo/query`
- 请求类型：[InvoiceInstitutionDetailInfoQueryRequest](institution_request.go#L196)
- 响应类型：[InvoiceInstitutionDetailInfoQueryResponse](institution_response.go#L21)

单个制度查询，返回制度详细信息，包含制度下使用规则列表和发放规则列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业共同账户id |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业ID |
| `institution_id` | `InstitutionId` | `*string` | query | 否 | 制度id |
| `outer_source_id` | `OuterSourceId` | `*string` | query | 否 | 外部唯一标识 |
| `owner_type` | `OwnerType` | `*string` | query | 否 | 适配id类型 |

<a id="api-invoiceinstitutionexpenserulecreate"></a>

### InvoiceInstitutionExpenseRuleCreate — 创建费控使用规则

- 官方接口：[`alipay.ebpp.invoice.institution.expenserule.create`](https://opendocs.alipay.com/apis/05waoz)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/expenserule/create`
- 请求类型：[InvoiceInstitutionExpenseRuleCreateRequest](institution_request.go#L50)
- 响应类型：[InvoiceInstitutionExpenseRuleCreateResponse](institution_response.go#L64)

创建费控使用规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `asset_share_source_info` | `AssetShareSourceInfo` | [`*AssetShareSourceInfo`](#model-assetsharesourceinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `consume_mode` | `ConsumeMode` | `*string` | JSON body | 未标注 | 消费模式 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业码id |
| `expense_ctrl_rule_info_list` | `ExpenseCtrlRuleInfoList` | [`*ExpenseCtrRuleInfo`](#model-expensectrruleinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `expense_type_sub_category` | `ExpenseTypeSubCategory` | `*string` | JSON body | 未标注 | 费用类型子类 |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `open_rule_id` | `OpenRuleId` | `*string` | JSON body | 未标注 | 开票规则id |
| `outer_source_id` | `OuterSourceId` | `*string` | JSON body | 未标注 | 外部唯一标识，填写该字段可用于创建幂等，防止重复创建 |
| `payment_policy` | `PaymentPolicy` | `*string` | JSON body | 未标注 | 支付策略 |
| `payment_time_mode` | `PaymentTimeMode` | `*string` | JSON body | 未标注 | 支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付 |
| `personal_qrcode_mode` | `PersonalQrcodeMode` | `*int64` | JSON body | 未标注 | 个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持） |
| `standard_condition_info_list` | `StandardConditionInfoList` | [`*[]StandardConditionInfo`](#model-standardconditioninfo) | JSON body | 未标注 | 使用规则因子列表 |
| `standard_desc` | `StandardDesc` | `*string` | JSON body | 未标注 | 使用规则描述 |
| `standard_name` | `StandardName` | `*string` | JSON body | 未标注 | 费控规则名称 |

<a id="api-invoiceinstitutionexpenseruledelete"></a>

### InvoiceInstitutionExpenseRuleDelete — 删除使用规则

- 官方接口：[`alipay.ebpp.invoice.institution.expenserule.delete`](https://opendocs.alipay.com/apis/05xa73)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/expenserule/delete`
- 请求类型：[InvoiceInstitutionExpenseRuleDeleteRequest](institution_request.go#L84)
- 响应类型：[InvoiceInstitutionExpenseRuleDeleteResponse](institution_response.go#L70)

删除制度下使用规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id（alipay.ebpp.invoice.institution.expenserule.delete）(alipay.ebpp.invoice.institution.expenserule.delete) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `standard_id_list` | `StandardIdList` | `*[]string` | JSON body | 未标注 | 删除的使用规则id列表 |

<a id="api-invoiceinstitutionexpenserulemodify"></a>

### InvoiceInstitutionExpenseRuleModify — 编辑使用规则

- 官方接口：[`alipay.ebpp.invoice.institution.expenserule.modify`](https://opendocs.alipay.com/apis/05wuq0)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/expenserule/modify`
- 请求类型：[InvoiceInstitutionExpenseRuleModifyRequest](institution_request.go#L98)
- 响应类型：[InvoiceInstitutionExpenseRuleModifyResponse](institution_response.go#L76)

编辑制度下的使用规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `action` | `Action` | `*string` | JSON body | 未标注 | 修改使用规则 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `asset_share_source_info` | `AssetShareSourceInfo` | [`*AssetShareSourceInfo`](#model-assetsharesourceinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `consume_mode` | `ConsumeMode` | `*string` | JSON body | 未标注 | 该使用规则支持的资产消费模式，不填写则为默认模式，默认模式下有余额时使用余额，没有余额则使用规则中的限额，点券模式为只能使用点券，点券+余额模式为可以使用点券和余额。 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `expense_ctrl_rule_info_list` | `ExpenseCtrlRuleInfoList` | [`*[]ExpenseCtrRuleInfo`](#model-expensectrruleinfo) | JSON body | 未标注 | 使用规则条件列表（已废弃） |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `open_rule_id` | `OpenRuleId` | `*string` | JSON body | 未标注 | 开票规则id |
| `payment_policy` | `PaymentPolicy` | `*string` | JSON body | 未标注 | 当笔消费金额大于规则可用余额时，用于控制支付策略 |
| `payment_time_mode` | `PaymentTimeMode` | `*string` | JSON body | 未标注 | 支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付 |
| `personal_qrcode_mode` | `PersonalQrcodeMode` | `*int64` | JSON body | 未标注 | 个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持） |
| `standard_condition_info_list` | `StandardConditionInfoList` | [`*[]StandardConditionInfo`](#model-standardconditioninfo) | JSON body | 未标注 | 规则条件列表 |
| `standard_desc` | `StandardDesc` | `*string` | JSON body | 未标注 | 规则描述 |
| `standard_id` | `StandardId` | `*string` | JSON body | 未标注 | 使用规则id |
| `standard_name` | `StandardName` | `*string` | JSON body | 未标注 | 规则名称 |

<a id="api-invoiceinstitutionmodify"></a>

### InvoiceInstitutionModify — 制度编辑

- 官方接口：[`alipay.ebpp.invoice.institution.modify`](https://opendocs.alipay.com/apis/05wvqu)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/modify`
- 请求类型：[InvoiceInstitutionModifyRequest](institution_request.go#L133)
- 响应类型：[InvoiceInstitutionModifyResponse](institution_response.go#L82)

制度编辑，支持制度基本信息、制度下适用范围、制度下使用规则和发放规则的修改。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id （该字段将废弃，不建议使用，可用enterprise_id字段替换）(推荐使用enterprise_id) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 （该字段将废弃，不建议使用，可用enterprise_id字段替换）(推荐使用enterprise_id) |
| `consult_mode` | `ConsultMode` | `*string` | JSON body | 未标注 | 费控咨询模式： 0-支付宝内部计算规则可用金额 1-咨询外部服务商规则可用金额 |
| `effective` | `Effective` | `*string` | JSON body | 未标注 | 制度是否启用 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 制度生效结束时间（可空），最小粒度为天，最早是当日 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 制度生效起始时间 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `institution_desc` | `InstitutionDesc` | `*string` | JSON body | 未标注 | 制度描述 |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `institution_name` | `InstitutionName` | `*string` | JSON body | 未标注 | 制度名称 |
| `modify_issue_rule_detail_info` | `ModifyIssueRuleDetailInfo` | [`*ModifyIssueRuleDetailInfo`](#model-modifyissueruledetailinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `modify_scope_info` | `ModifyScopeInfo` | [`*ModifyScopeInfo`](#model-modifyscopeinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `modify_standard_detail_info` | `ModifyStandardDetailInfo` | [`*ModifyStandardDetailInfo`](#model-modifystandarddetailinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-invoiceinstitutionpageinfoquery"></a>

### InvoiceInstitutionPageInfoQuery — 费控制度分页查询

- 官方接口：[`alipay.ebpp.invoice.institution.pageinfo.query`](https://opendocs.alipay.com/apis/05xo4v)
- 路由：`GET /v3/alipay/ebpp/invoice/institution/pageinfo/query`
- 请求类型：[InvoiceInstitutionPageInfoQueryRequest](institution_request.go#L212)
- 响应类型：[InvoiceInstitutionPageInfoQueryResponse](institution_response.go#L92)

制度分页查询，返回制度基本信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业ID |
| `page_num` | `PageNum` | `*int64` | query | 是 | 页码 |
| `page_size` | `PageSize` | `*int64` | query | 是 | 页大小 |

<a id="api-invoiceinstitutionscopemodify"></a>

### InvoiceInstitutionScopeModify — 修改制度下适配范围

- 官方接口：[`alipay.ebpp.invoice.institution.scope.modify`](https://opendocs.alipay.com/apis/05ip0w)
- 路由：`POST /v3/alipay/ebpp/invoice/institution/scope/modify`
- 请求类型：[InvoiceInstitutionScopeModifyRequest](institution_request.go#L160)
- 响应类型：[InvoiceInstitutionScopeModifyResponse](institution_response.go#L104)

修改制度下适配范围

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id |
| `adapter_type` | `AdapterType` | `*string` | JSON body | 未标注 | 制度适用范围类型 |
| `add_owner_id_list` | `AddOwnerIdList` | `*[]string` | JSON body | 未标注 | 增加适配id列表 |
| `add_owner_open_id_list` | `AddOwnerOpenIdList` | `*[]string` | JSON body | 未标注 | 增加适配开放id列表 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `owner_type` | `OwnerType` | `*string` | JSON body | 未标注 | 归属类型 |
| `remove_owner_id_list` | `RemoveOwnerIdList` | `*[]string` | JSON body | 未标注 | 删除适配id列表 |
| `remove_owner_open_id_list` | `RemoveOwnerOpenIdList` | `*[]string` | JSON body | 未标注 | 删除适配开放id列表 |

<a id="api-invoiceinstitutionscopepageinfoquery"></a>

### InvoiceInstitutionScopePageInfoQuery — 分页查询制度下人员

- 官方接口：[`alipay.ebpp.invoice.institution.scopepageinfo.query`](https://opendocs.alipay.com/apis/05j1c2)
- 路由：`GET /v3/alipay/ebpp/invoice/institution/scopepageinfo/query`
- 请求类型：[InvoiceInstitutionScopePageInfoQueryRequest](institution_request.go#L226)
- 响应类型：[InvoiceInstitutionScopePageInfoQueryResponse](institution_response.go#L110)

分页查询制度下适用范围列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（数值及布尔字段按接口文档） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `institution_id` | `InstitutionId` | `*string` | query | 否 | 制度id |
| `owner_type` | `OwnerType` | `*string` | query | 否 | 适配id类型 |
| `page_num` | `PageNum` | `*int64` | query | 是 | 页码 |
| `page_size` | `PageSize` | `*int64` | query | 是 | 页大小 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-assetsharesourceinfo"></a>

### AssetShareSourceInfo

类型定义：[AssetShareSourceInfo](institution_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `share_mode` | `ShareMode` | `*string` | 未标注 | 当前规则可用的资产来源类型，搭配source_id_list使用。 |
| `source_id_list` | `SourceIdList` | `*[]string` | 未标注 | 资产共享来源id列表，与share_mode配合设置，如果share_mode为INSTITUTION，该值则填写制度id列表 |

<a id="model-expensectrruleinfo"></a>

### ExpenseCtrRuleInfo

类型定义：[ExpenseCtrRuleInfo](institution_model.go#L13)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `rule_factor` | `RuleFactor` | `*string` | 未标注 | 费控维度 |
| `rule_id` | `RuleId` | `*string` | 未标注 | 费控条件ID |
| `rule_name` | `RuleName` | `*string` | 未标注 | 费控条件名称 |
| `rule_operator` | `RuleOperator` | `*string` | 未标注 | 费控条件操作符 枚举值： LT("&lt;","小于") LE("&lt;=","小于等于") EQ("=","等于") NE("!=","不等于") GE("&gt;=","大于等于") GT("&gt;","大于") |
| `rule_value` | `RuleValue` | `*string` | 未标注 | 费控条件值 |

<a id="model-institutionscopeinfo"></a>

### InstitutionScopeInfo

类型定义：[InstitutionScopeInfo](institution_model.go#L49)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `adapter_type` | `AdapterType` | `*string` | 未标注 | 制度适用范围，支持配置全员、指定员工和指定部门 |
| `owner_id_list` | `OwnerIdList` | `*[]string` | 未标注 | 适用范围归属id列表 |
| `owner_type` | `OwnerType` | `*string` | 未标注 | 员工id类型 |

<a id="model-issueruleinfo"></a>

### IssueRuleInfo

类型定义：[IssueRuleInfo](institution_model.go#L67)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `effective_period` | `EffectivePeriod` | `*string` | 未标注 | 生效时间段 |
| `invalid_mode` | `InvalidMode` | `*int64` | 未标注 | 累计类型，默认为0 可选值：0（不可累计）、1（可累计）、2（累计天数）、3（累计到指定日期） |
| `invalid_mode_value` | `InvalidModeValue` | `*string` | 未标注 | 累计类型值 |
| `issue_amount_value` | `IssueAmountValue` | `*string` | 未标注 | 发放金额，单位元 |
| `issue_end_date` | `IssueEndDate` | `*string` | 未标注 | 发放规则有效结束时间 |
| `issue_rule_id` | `IssueRuleId` | `*string` | 未标注 | 发放规则id |
| `issue_rule_name` | `IssueRuleName` | `*string` | 未标注 | 发放规则名称 |
| `issue_start_date` | `IssueStartDate` | `*string` | 未标注 | 发放规则有效起始时间 |
| `issue_type` | `IssueType` | `*string` | 未标注 | 发放类型 |
| `outer_source_id` | `OuterSourceId` | `*string` | 未标注 | 外部发放规则id |
| `quota_type` | `QuotaType` | `*string` | 未标注 | 额度类型 |
| `share_mode` | `ShareMode` | `*int64` | 未标注 | 是否可转赠 |
| `target_id` | `TargetId` | `*string` | 未标注 | 目标id |
| `target_type` | `TargetType` | `*string` | 未标注 | 发放规则归属的目标类型 |

<a id="model-modifyissueruledetailinfo"></a>

### ModifyIssueRuleDetailInfo

类型定义：[ModifyIssueRuleDetailInfo](institution_model.go#L99)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `add_issue_rule_list` | `AddIssueRuleList` | [`*[]IssueRuleInfo`](#model-issueruleinfo) | 未标注 | 新增发放规则列表 |
| `delete_issue_rule_id_list` | `DeleteIssueRuleIdList` | `*[]string` | 未标注 | 删除的发放规则id列表 |
| `modify_issue_rule_list` | `ModifyIssueRuleList` | [`*ModifyIssueRuleInfo`](#model-modifyissueruleinfo) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-modifyissueruleinfo"></a>

### ModifyIssueRuleInfo

类型定义：[ModifyIssueRuleInfo](institution_model.go#L108)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `effective_period` | `EffectivePeriod` | `*string` | 未标注 | 发放资产的有效时间段，该字段为空时，默认为不限制。当发放资产为余额时，仅支持设置为不限，当发放资产为点券时，可支持设置多种值。具体传值格式请参考文档 [接入文档](https://opendocs.alipay.com/pre-open/07nwcs) |
| `invalid_mode` | `InvalidMode` | `*int64` | 未标注 | 累计类型，默认为0 可选值：0（不可累计）、1（可累计）、2（累计天数）、3（累计到指定日期） |
| `invalid_mode_value` | `InvalidModeValue` | `*string` | 未标注 | 累计类型值 |
| `issue_amount_value` | `IssueAmountValue` | `*string` | 未标注 | 发放金额（不传则不修改），单位（元） |
| `issue_rule_id` | `IssueRuleId` | `*string` | 未标注 | 发放规则id |
| `issue_rule_name` | `IssueRuleName` | `*string` | 未标注 | 发放规则名称 |
| `issue_type` | `IssueType` | `*string` | 未标注 | 额度发放周期 |
| `quota_type` | `QuotaType` | `*string` | 未标注 | 发放规则发放的额度类型 |
| `share_mode` | `ShareMode` | `*int64` | 未标注 | 是否可转赠（不传则不修改），可选值：0（不可转赠）、1（可转赠），默认为0。该字段决定发放规则发放的资产是否可转赠给同制度下的其他人员 |

<a id="model-modifyscopeinfo"></a>

### ModifyScopeInfo

类型定义：[ModifyScopeInfo](institution_model.go#L130)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `adapter_type` | `AdapterType` | `*string` | 未标注 | 制度的适用范围类型 |
| `add_owner_id_list` | `AddOwnerIdList` | `*[]string` | 未标注 | 增加的归属id列表 |
| `delete_owner_id_list` | `DeleteOwnerIdList` | `*[]string` | 未标注 | 删除归属id列表 |
| `owner_type` | `OwnerType` | `*string` | 未标注 | 归属id类型 |

<a id="model-modifystandarddetailinfo"></a>

### ModifyStandardDetailInfo

类型定义：[ModifyStandardDetailInfo](institution_model.go#L142)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `add_standard_list` | `AddStandardList` | [`*[]StandardInfo`](#model-standardinfo) | 未标注 | 新增的使用规则列表 |
| `delete_standard_id_list` | `DeleteStandardIdList` | `*[]string` | 未标注 | 待删除的使用规则id列表 |
| `modify_standard_list` | `ModifyStandardList` | [`*[]ModifyStandardInfo`](#model-modifystandardinfo) | 未标注 | 待修改的使用规则信息列表 |

<a id="model-modifystandardinfo"></a>

### ModifyStandardInfo

类型定义：[ModifyStandardInfo](institution_model.go#L152)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `add_condition_list` | `AddConditionList` | [`*[]StandardConditionInfo`](#model-standardconditioninfo) | 未标注 | 要添加的条件列表 |
| `asset_share_source_info` | `AssetShareSourceInfo` | [`*AssetShareSourceInfo`](#model-assetsharesourceinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `consume_mode` | `ConsumeMode` | `*string` | 未标注 | 消费模式 |
| `delete_condition_id_list` | `DeleteConditionIdList` | `*[]string` | 未标注 | 待删除的条件id列表 |
| `modify_condition_list` | `ModifyConditionList` | [`*[]StandardConditionInfo`](#model-standardconditioninfo) | 未标注 | 修改的条件列表 |
| `open_rule_id` | `OpenRuleId` | `*string` | 未标注 | 使用规则绑定的开票规则id |
| `payment_policy` | `PaymentPolicy` | `*string` | 未标注 | 支付策略 当笔消费金额大于规则可用余额时，用于控制支付策略。COMBINATION表示支持因公资产和个人资产组合支付，PERSONAL表示整单个人支付。 |
| `payment_time_mode` | `PaymentTimeMode` | `*string` | 未标注 | 支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付 |
| `personal_qrcode_mode` | `PersonalQrcodeMode` | `*int64` | 未标注 | 个人收款码转账是否支持因公付。可选值：0（不支持）、1（支持） |
| `standard_desc` | `StandardDesc` | `*string` | 未标注 | 使用规则描述 |
| `standard_id` | `StandardId` | `*string` | 未标注 | 修改的使用规则id |
| `standard_name` | `StandardName` | `*string` | 未标注 | 使用规则名称 |

<a id="model-standardconditioninfo"></a>

### StandardConditionInfo

类型定义：[StandardConditionInfo](institution_model.go#L179)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `rule_factor` | `RuleFactor` | `*string` | 未标注 | 条件类型 |
| `rule_id` | `RuleId` | `*string` | 未标注 | 费控条件ID |
| `rule_name` | `RuleName` | `*string` | 未标注 | 费控条件名称 |
| `rule_operator` | `RuleOperator` | `*string` | 未标注 | 费控条件操作符 枚举值： LT("&lt;","小于") LE("&lt;=","小于等于") EQ("=","等于") NE("!=","不等于") GE("&gt;=","大于等于") GT("&gt;","大于")(该字段为保留字段，目前未被使用) |
| `rule_value` | `RuleValue` | `*string` | 未标注 | 费控条件值 |

<a id="model-standardinfo"></a>

### StandardInfo

类型定义：[StandardInfo](institution_model.go#L201)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `asset_share_source_info` | `AssetShareSourceInfo` | [`*AssetShareSourceInfo`](#model-assetsharesourceinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `consume_mode` | `ConsumeMode` | `*string` | 未标注 | 消费模式，不填为默认模式，枚举值：COUPON_ONLY（仅支持点券） COUPON_AND_CAP（支持点券+余额） COUNT（仅支持次卡） DEFAULT（默认模式） 点券：消费时找员工的点券，没有或者用完了不可付； 点券+余额：消费时找员工的点券，没有找员工的余额，没有或者用完了不可付； 次卡：消费时找员工的次卡，没有或者用完了不可付； 默认：有给员工设置员工余额以员工余额为准，用完为止。否则只受规则里的限额和企业账户资金上限管控； |
| `expense_type_sub_category` | `ExpenseTypeSubCategory` | `*string` | 未标注 | 费用类型子类 |
| `open_rule_id` | `OpenRuleId` | `*string` | 未标注 | 开票规则id，可通过接口alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create 创建并得到开票规则ID |
| `outer_source_id` | `OuterSourceId` | `*string` | 未标注 | 外部使用规则id，制度内使用规则该字段不允许重复 |
| `payment_policy` | `PaymentPolicy` | `*string` | 未标注 | 支付策略 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略 |
| `payment_time_mode` | `PaymentTimeMode` | `*string` | 未标注 | 支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付 |
| `personal_qrcode_mode` | `PersonalQrcodeMode` | `*int64` | 未标注 | 个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持） |
| `standard_condition_info_list` | `StandardConditionInfoList` | [`*[]StandardConditionInfo`](#model-standardconditioninfo) | 未标注 | 使用规则条件列表 |
| `standard_desc` | `StandardDesc` | `*string` | 未标注 | 使用规则描述(敏感词校验) |
| `standard_id` | `StandardId` | `*string` | 未标注 | 制度ID（创建使用规则时非必填） |
| `standard_name` | `StandardName` | `*string` | 未标注 | 规则名称 |
