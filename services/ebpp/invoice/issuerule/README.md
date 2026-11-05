# ebpp/invoice/issuerule 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/issuerule`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `issuerule.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_issuerule_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceIssueRuleCreate`](#api-invoiceissuerulecreate) | [`alipay.ebpp.invoice.issuerule.create`](https://opendocs.alipay.com/apis/05xoz2) | 新增发放规则 | `POST` |
| [`InvoiceIssueRuleDelete`](#api-invoiceissueruledelete) | [`alipay.ebpp.invoice.issuerule.delete`](https://opendocs.alipay.com/apis/05wol0) | 删除发放规则 | `POST` |
| [`InvoiceIssueRuleModify`](#api-invoiceissuerulemodify) | [`alipay.ebpp.invoice.issuerule.modify`](https://opendocs.alipay.com/apis/05wyj4) | 编辑发放规则 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceissuerulecreate"></a>

### InvoiceIssueRuleCreate — 新增发放规则

- 官方接口：[`alipay.ebpp.invoice.issuerule.create`](https://opendocs.alipay.com/apis/05xoz2)
- 路由：`POST /v3/alipay/ebpp/invoice/issuerule/create`
- 请求类型：[InvoiceIssueRuleCreateRequest](issuerule_request.go#L5)
- 响应类型：[InvoiceIssueRuleCreateResponse](issuerule_response.go#L5)

新增发放规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `effective` | `Effective` | `*string` | JSON body | 未标注 | 是否启用，"0"为停用，"1"为启用； 如果target_type指定INSTITUTION，则与制度一致,无需填写 |
| `effective_period` | `EffectivePeriod` | `*string` | JSON body | 未标注 | 生效周期，时间不限格式："{\"all\":true}" |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `invalid_mode` | `InvalidMode` | `*int64` | JSON body | 未标注 | 是否可累计：0（不可累计）、1（可累计），2（累计天数）、3（累计到指定日期）余额只允许可累计，按城市发放只允许不可累计，设置了生效周期只允许不可累计 |
| `invalid_mode_value` | `InvalidModeValue` | `*string` | JSON body | 未标注 | 累计值（预留字段） |
| `issue_amount_value` | `IssueAmountValue` | `*string` | JSON body | 未标注 | 发放金额，单位（元）； 按城市不同发放标准示例值："[{\"cities\":[\"321000\",\"321100\"],\"amount\":\"50\"},{\"cities\":[\"321200\",\"321300\"],\"amount\":\"150\"}]" 统一标准发放示例值："200" |
| `issue_end_date` | `IssueEndDate` | `*string` | JSON body | 未标注 | 发放规则的截止时间 如果target_type指定INSTITUTION，则使用制度有效期结束时间，无需填写 |
| `issue_rule_name` | `IssueRuleName` | `*string` | JSON body | 未标注 | 发放规则名称 |
| `issue_start_date` | `IssueStartDate` | `*string` | JSON body | 未标注 | 发放规则的开始时间 如果target_type指定INSTITUTION，则使用制度有效期起始时间，无需填写 |
| `issue_type` | `IssueType` | `*string` | JSON body | 未标注 | 发放类型 |
| `outer_source_id` | `OuterSourceId` | `*string` | JSON body | 未标注 | 填写该字段可用于创建幂等，防止重复创建 |
| `quota_type` | `QuotaType` | `*string` | JSON body | 未标注 | 额度类型，枚举：COUPON（点券）、CAP（余额） |
| `share_mode` | `ShareMode` | `*int64` | JSON body | 未标注 | 是否可转赠：0（不可转赠）、1（可转赠） |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 发放规则关联的目标值 target_type为制度时：制度id |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 发放规则关联的目标类型 |

<a id="api-invoiceissueruledelete"></a>

### InvoiceIssueRuleDelete — 删除发放规则

- 官方接口：[`alipay.ebpp.invoice.issuerule.delete`](https://opendocs.alipay.com/apis/05wol0)
- 路由：`POST /v3/alipay/ebpp/invoice/issuerule/delete`
- 请求类型：[InvoiceIssueRuleDeleteRequest](issuerule_request.go#L43)
- 响应类型：[InvoiceIssueRuleDeleteResponse](issuerule_response.go#L11)

删除发放规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `issue_rule_id_list` | `IssueRuleIdList` | `*[]string` | JSON body | 未标注 | 需要删除的发放规则id列表 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 目标值id |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 发放规则关联的目标类型 |

<a id="api-invoiceissuerulemodify"></a>

### InvoiceIssueRuleModify — 编辑发放规则

- 官方接口：[`alipay.ebpp.invoice.issuerule.modify`](https://opendocs.alipay.com/apis/05wyj4)
- 路由：`POST /v3/alipay/ebpp/invoice/issuerule/modify`
- 请求类型：[InvoiceIssueRuleModifyRequest](issuerule_request.go#L59)
- 响应类型：[InvoiceIssueRuleModifyResponse](issuerule_response.go#L17)

编辑发放规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `action` | `Action` | `*string` | JSON body | 未标注 | 修改类型，枚举： MODIFY_BASIC_INFO（修改发放规则基本信息） |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换) |
| `effective` | `Effective` | `*string` | JSON body | 未标注 | 是否启用，"0"为停用，"1"为启用； |
| `effective_period` | `EffectivePeriod` | `*string` | JSON body | 未标注 | 生效周期（不传则不修改），不限："{\"all\":true}" |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业ID |
| `invalid_mode` | `InvalidMode` | `*int64` | JSON body | 未标注 | 是否可累计（不传则不修改）可选值：0（不可累计）、1（可累计），2（累计天数），3（累计到固定时间） |
| `invalid_mode_value` | `InvalidModeValue` | `*string` | JSON body | 未标注 | 累计值 |
| `issue_amount_value` | `IssueAmountValue` | `*string` | JSON body | 未标注 | 发放金额（不传则不修改），单位（元） |
| `issue_end_date` | `IssueEndDate` | `*string` | JSON body | 未标注 | 发放规则的截止时间 |
| `issue_rule_id` | `IssueRuleId` | `*string` | JSON body | 未标注 | 发放规则id |
| `issue_rule_name` | `IssueRuleName` | `*string` | JSON body | 未标注 | 发放规则名称（不传则不修改） |
| `issue_start_date` | `IssueStartDate` | `*string` | JSON body | 未标注 | 发放规则的开始时间 |
| `issue_type` | `IssueType` | `*string` | JSON body | 未标注 | 发放类型 |
| `quota_type` | `QuotaType` | `*string` | JSON body | 未标注 | 额度类型 |
| `share_mode` | `ShareMode` | `*int64` | JSON body | 未标注 | 是否可转赠（不传则不修改），可选值：0（不可转赠）、1（可转赠），默认为0，按城市发放不支持可转赠 |
| `target_id` | `TargetId` | `*string` | JSON body | 未标注 | 发放规则关联的目标值 target_type为制度时：制度id |
| `target_type` | `TargetType` | `*string` | JSON body | 未标注 | 发放规则关联的目标类型 |
