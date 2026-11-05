# ebpp/invoice/expensecontrol 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expensecontrol`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `expensecontrol.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceExpenseControlAggregationCreate`](#api-invoiceexpensecontrolaggregationcreate) | [`alipay.ebpp.invoice.expensecontrol.aggregation.create`](https://opendocs.alipay.com/apis/05xsf0) | 创建费控规则聚合关系 | `POST` |
| [`InvoiceExpenseControlIssueBatchCancel`](#api-invoiceexpensecontrolissuebatchcancel) | [`alipay.ebpp.invoice.expensecontrol.issuebatch.cancel`](https://opendocs.alipay.com/apis/05wumt) | 作废额度发放 | `POST` |
| [`InvoiceExpenseControlIssueBatchCreate`](#api-invoiceexpensecontrolissuebatchcreate) | [`alipay.ebpp.invoice.expensecontrol.issuebatch.create`](https://opendocs.alipay.com/apis/05j9jf) | 手动发放额度接口 | `POST` |
| [`InvoiceExpenseControlQuotaCreate`](#api-invoiceexpensecontrolquotacreate) | [`alipay.ebpp.invoice.expensecontrol.quota.create`](https://opendocs.alipay.com/apis/05iz32) | 创建余额/点券 | `POST` |
| [`InvoiceExpenseControlQuotaModify`](#api-invoiceexpensecontrolquotamodify) | [`alipay.ebpp.invoice.expensecontrol.quota.modify`](https://opendocs.alipay.com/apis/05wuoa) | 修改余额/点券 | `POST` |
| [`InvoiceExpenseControlQuotaQuery`](#api-invoiceexpensecontrolquotaquery) | [`alipay.ebpp.invoice.expensecontrol.quota.query`](https://opendocs.alipay.com/apis/05imop) | 查询余额/点券 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceexpensecontrolaggregationcreate"></a>

### InvoiceExpenseControlAggregationCreate — 创建费控规则聚合关系

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.aggregation.create`](https://opendocs.alipay.com/apis/05xsf0)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/aggregation/create`
- 请求类型：[InvoiceExpenseControlAggregationCreateRequest](expensecontrol_request.go#L5)
- 响应类型：[InvoiceExpenseControlAggregationCreateResponse](expensecontrol_response.go#L5)

创建费控规则聚合关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `aggregation_name` | `AggregationName` | `*string` | JSON body | 未标注 | 聚合关系名称 特殊说明：敏感词校验 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `standard_id_list` | `StandardIdList` | `*[]string` | JSON body | 未标注 | 费控规则ID列表 |

<a id="api-invoiceexpensecontrolissuebatchcancel"></a>

### InvoiceExpenseControlIssueBatchCancel — 作废额度发放

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.issuebatch.cancel`](https://opendocs.alipay.com/apis/05wumt)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel`
- 请求类型：[InvoiceExpenseControlIssueBatchCancelRequest](expensecontrol_request.go#L17)
- 响应类型：[InvoiceExpenseControlIssueBatchCancelResponse](expensecontrol_response.go#L11)

通过发放批次id，作废当前批次下发放的额度

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `issue_batch_id` | `IssueBatchId` | `*string` | JSON body | 未标注 | 发放批次id |

<a id="api-invoiceexpensecontrolissuebatchcreate"></a>

### InvoiceExpenseControlIssueBatchCreate — 手动发放额度接口

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.issuebatch.create`](https://opendocs.alipay.com/apis/05j9jf)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create`
- 请求类型：[InvoiceExpenseControlIssueBatchCreateRequest](expensecontrol_request.go#L31)
- 响应类型：[InvoiceExpenseControlIssueBatchCreateResponse](expensecontrol_response.go#L17)

通过该接口对企业下的员工进行批量的额度发放。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `batch_no` | `BatchNo` | `*string` | JSON body | 未标注 | 发放批次号，用于幂等校验本次发放 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 额度有效结束时间 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 额度有效起始时间 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `institution_id` | `InstitutionId` | `*string` | JSON body | 未标注 | 制度id |
| `issue_desc` | `IssueDesc` | `*string` | JSON body | 未标注 | 发放说明 |
| `issue_name` | `IssueName` | `*string` | JSON body | 未标注 | 发放名称 |
| `issue_target_info_list` | `IssueTargetInfoList` | [`*[]IssueTargetInfoContent`](#model-issuetargetinfocontent) | JSON body | 未标注 | 员工发放信息列表 |
| `quota_type` | `QuotaType` | `*string` | JSON body | 未标注 | 额度类型 |
| `share_mode` | `ShareMode` | `*string` | JSON body | 未标注 | 是否可转赠 |

<a id="api-invoiceexpensecontrolquotacreate"></a>

### InvoiceExpenseControlQuotaCreate — 创建余额/点券

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.quota.create`](https://opendocs.alipay.com/apis/05iz32)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/quota/create`
- 请求类型：[InvoiceExpenseControlQuotaCreateRequest](expensecontrol_request.go#L59)
- 响应类型：[InvoiceExpenseControlQuotaCreateResponse](expensecontrol_response.go#L25)

创建余额或者点券，可支持批量创建点券

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账号id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 额度失效时间（格式：yyyy-MM-dd HH:mm:ss） |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 额度生效时间（格式：yyyy-MM-dd HH:mm:ss） |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `issue_desc` | `IssueDesc` | `*string` | JSON body | 未标注 | 批量发放时，可填写该字段 |
| `issue_name` | `IssueName` | `*string` | JSON body | 未标注 | 批量发放时，可填写该字段作为发放名称，如果未填写取outer_source_id为默认值 |
| `issue_quota_target_list` | `IssueQuotaTargetList` | [`*[]IssueTargetInfoContent`](#model-issuetargetinfocontent) | JSON body | 未标注 | 创建额度发放明细列表 |
| `outer_source_id` | `OuterSourceId` | `*string` | JSON body | 未标注 | 外部操作幂等ID，标识创建额度的唯一性，防止重复创建 |
| `owner_id` | `OwnerId` | `*string` | JSON body | 未标注 | 额度所属者ID（未切换open_id请使用此字段）：(字段升级，请使用issue_target_info_list中owner_id字段) |
| `owner_open_id` | `OwnerOpenId` | `*string` | JSON body | 未标注 | 额度所属者ID（切换open_id后请使用此字段）：(字段升级，请使用issue_quota_target_list中owner_open_id) |
| `owner_type` | `OwnerType` | `*string` | JSON body | 未标注 | 额度所属者类型(字段升级，请使用issue_quota_target_list中owner_type字段) |
| `platform` | `Platform` | `*string` | JSON body | 未标注 | 外部平台编码（通常为接入方大写英文缩写）(历史版本字段，不推荐使用) |
| `quota_type` | `QuotaType` | `*string` | JSON body | 未标注 | 创建额度类型 |
| `quota_value` | `QuotaValue` | `*string` | JSON body | 未标注 | 额度值，以（分）为单位(字段升级，请使用issue_quota_target_list中issue_quota字段) |
| `share_mode` | `ShareMode` | `*string` | JSON body | 未标注 | 0:不可转赠 1:可以转增 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 额度维度ID |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 额度维度 枚举值： INSTITUTION（制度维度）， EXPENSE_TYPE（费用类型维度） |

<a id="api-invoiceexpensecontrolquotamodify"></a>

### InvoiceExpenseControlQuotaModify — 修改余额/点券

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.quota.modify`](https://opendocs.alipay.com/apis/05wuoa)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/quota/modify`
- 请求类型：[InvoiceExpenseControlQuotaModifyRequest](expensecontrol_request.go#L99)
- 响应类型：[InvoiceExpenseControlQuotaModifyResponse](expensecontrol_response.go#L35)

修改余额或点券对应的额度，可用于增加额度、减少额度、删除额度等操作。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户ID（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `action` | `Action` | `*string` | JSON body | 未标注 | 变更模式 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `allow_negative` | `AllowNegative` | `*bool` | JSON body | 未标注 | 修改余额对应的额度时是否允许修改为负数，true表示允许，false表示不允许。默认允许 |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 变更的金额/余额，日元或者韩元以（元）为单位，其他币种以（分）为单位，当变更的资产类型为次卡时，单位为次 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `outer_source_id` | `OuterSourceId` | `*string` | JSON body | 未标注 | 外部操作幂等ID（接入方接口调用幂等控制ID） |
| `platform` | `Platform` | `*string` | JSON body | 未标注 | 外部平台编码(历史字段，不推荐使用) |
| `quota_id` | `QuotaId` | `*string` | JSON body | 未标注 | 额度ID |
| `share_mode` | `ShareMode` | `*string` | JSON body | 未标注 | 是否可转赠，1表示可转赠，0表示不可转赠 |

<a id="api-invoiceexpensecontrolquotaquery"></a>

### InvoiceExpenseControlQuotaQuery — 查询余额/点券

- 官方接口：[`alipay.ebpp.invoice.expensecontrol.quota.query`](https://opendocs.alipay.com/apis/05imop)
- 路由：`POST /v3/alipay/ebpp/invoice/expensecontrol/quota/query`
- 请求类型：[InvoiceExpenseControlQuotaQueryRequest](expensecontrol_request.go#L123)
- 响应类型：[InvoiceExpenseControlQuotaQueryResponse](expensecontrol_response.go#L41)

查询余额或者点券

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `owner_id` | `OwnerId` | `*string` | JSON body | 未标注 | 额度所属者ID |
| `owner_open_id` | `OwnerOpenId` | `*string` | JSON body | 未标注 | 额度所属者开放ID |
| `owner_type` | `OwnerType` | `*string` | JSON body | 未标注 | 额度所属者类型 |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 页码 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页条数 |
| `quota_id_list` | `QuotaIdList` | `*[]string` | JSON body | 未标注 | 额度ID列表 |
| `quota_type` | `QuotaType` | `*string` | JSON body | 未标注 | 额度类型 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 额度维度ID |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 额度维度 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-issuetargetinfocontent"></a>

### IssueTargetInfoContent

类型定义：[IssueTargetInfoContent](expensecontrol_model.go#L59)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `issue_quota` | `IssueQuota` | `*string` | 未标注 | 当发放余额或点券时，单位为元，币种以创建制度时填写的currency为准，不填写默认为人民币。当发放类型为次卡时，单位为次。 |
| `owner_id` | `OwnerId` | `*string` | 未标注 | 切换open_id前请使用：owner类型为PHONE时该字段表示员工手机号，owner类型为EMPLOYEE时该字段表示为员工支付宝uid，owner类型为ENTERPRISE_PAY_UID时该字段表示为员工企业码id，owner类型为ENTERPRISE时可不填，如果填写仅支持填企业id，owner类型为INSTITUTION时可不填，默认为制度id |
| `owner_open_id` | `OwnerOpenId` | `*string` | 未标注 | 切换open_id后请使用：owner类型为PHONE时，填写员工手机号，EMPLOYEE时填写员工open_id |
| `owner_type` | `OwnerType` | `*string` | 未标注 | owner类型，用来表示owner_id的类型 |
| `user_name` | `UserName` | `*string` | 未标注 | 员工姓名 |
