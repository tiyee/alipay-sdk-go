# project 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules/project`。本包提供 5 个支付宝 V3 REST 操作。

通过 `project.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`InvoiceExpenseRulesProjectEmployeeModify`](#api-invoiceexpenserulesprojectemployeemodify) | `alipay.ebpp.invoice.expenserules.projectemployee.modify` |
| [`InvoiceExpenseRulesProjectInfoModify`](#api-invoiceexpenserulesprojectinfomodify) | `alipay.ebpp.invoice.expenserules.projectinfo.modify` |
| [`InvoiceExpenseRulesProjectRuleCreate`](#api-invoiceexpenserulesprojectrulecreate) | `alipay.ebpp.invoice.expenserules.projectrule.create` |
| [`InvoiceExpenseRulesProjectRuleQuery`](#api-invoiceexpenserulesprojectrulequery) | `alipay.ebpp.invoice.expenserules.projectrule.query` |
| [`InvoiceExpenseRulesProjectRulesModify`](#api-invoiceexpenserulesprojectrulesmodify) | `alipay.ebpp.invoice.expenserules.projectrules.modify` |

## 使用示例

- [InvoiceExpenseRulesProjectEmployeeModify](api_project_invoiceexpenserulesprojectemployeemodify_example_test.go)
- [InvoiceExpenseRulesProjectInfoModify](api_project_invoiceexpenserulesprojectinfomodify_example_test.go)
- [InvoiceExpenseRulesProjectRuleCreate](api_project_invoiceexpenserulesprojectrulecreate_example_test.go)
- [InvoiceExpenseRulesProjectRuleQuery](api_project_invoiceexpenserulesprojectrulequery_example_test.go)
- [InvoiceExpenseRulesProjectRulesModify](api_project_invoiceexpenserulesprojectrulesmodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceexpenserulesprojectemployeemodify"></a>

### InvoiceExpenseRulesProjectEmployeeModify — 修改项目的员工列表

- 官方接口：[`alipay.ebpp.invoice.expenserules.projectemployee.modify`](https://opendocs.alipay.com/apis/05wvx5)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/projectemployee/modify`
- 请求类型：[InvoiceExpenseRulesProjectEmployeeModifyRequest](project_request.go#L4)
- 响应类型：[InvoiceExpenseRulesProjectEmployeeModifyResponse](project_response.go#L4)

修改项目的员工列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户ID |
| `add_employee_list` | `AddEmployeeList` | `*[]string` | JSON body | 未标注 | 切换open_id前使用此字段： 需要添加的员工UID列表 约束：一次最多50个 |
| `add_employee_open_id_list` | `AddEmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换open_id后使用此字段： 需要添加的员工open_id列表 约束：一次最多50个 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `project_id` | `ProjectId` | `*string` | JSON body | 未标注 | 项目ID |
| `remove_employee_list` | `RemoveEmployeeList` | `*[]string` | JSON body | 未标注 | 切换open_id前使用此字段： 需要移除的员工UID列表 约束：一次最多50个 |
| `remove_employee_open_id_list` | `RemoveEmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换open_id后前使用此字段： 需要移除的员工open_id列表 约束：一次最多50个 |

<a id="api-invoiceexpenserulesprojectinfomodify"></a>

### InvoiceExpenseRulesProjectInfoModify — 修改项目基础信息

- 官方接口：[`alipay.ebpp.invoice.expenserules.projectinfo.modify`](https://opendocs.alipay.com/apis/05wuq6)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/projectinfo/modify`
- 请求类型：[InvoiceExpenseRulesProjectInfoModifyRequest](project_request.go#L22)
- 响应类型：[InvoiceExpenseRulesProjectInfoModifyResponse](project_response.go#L10)

修改项目基础信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 有效期截止（yyyy-MM-dd HH:mm:ss） 特殊说明： 1）与起始时间必须同时传，且大于起始时间 2）不传则默认不修改 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 有效期起始（yyyy-MM-dd HH:mm:ss） 特殊说明： 1）与截止时间必须同时传，且小于截止时间 2）不传入则默认不修改 |
| `project_id` | `ProjectId` | `*string` | JSON body | 未标注 | 项目ID |
| `project_name` | `ProjectName` | `*string` | JSON body | 未标注 | 项目名 特殊说明：不传入则默认不修改 |

<a id="api-invoiceexpenserulesprojectrulecreate"></a>

### InvoiceExpenseRulesProjectRuleCreate — 创建项目费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.projectrule.create`](https://opendocs.alipay.com/apis/05isio)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/projectrule/create`
- 请求类型：[InvoiceExpenseRulesProjectRuleCreateRequest](project_request.go#L38)
- 响应类型：[InvoiceExpenseRulesProjectRuleCreateResponse](project_response.go#L16)

创建项目费控规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账户id |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 有效期截止（yyyy-MM-dd HH:mm:ss，必须大于起始时间） |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 有效期起始（yyyy-MM-dd HH:mm:ss，必须小于截止时间） |
| `employee_list` | `EmployeeList` | `*[]string` | JSON body | 未标注 | 切换open_id前使用此字段：员工支付宝ID列表 特殊说明：单次传入的最大员工数为50，后续通过修改项目的员工列表接口进行员工调整 |
| `employee_open_id_list` | `EmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换open_id后使用此字段：员工open_id列表 特殊说明：单次传入的最大员工数为50，后续通过修改项目的员工列表接口进行员工调整 |
| `expense_ctrl_rule_info_group_list` | `ExpenseCtrlRuleInfoGroupList` | [`*[]ExpenseCtrRuleGroupInfo`](#model-expensectrrulegroupinfo) | JSON body | 未标注 | 费控规则列表 |
| `project_name` | `ProjectName` | `*string` | JSON body | 未标注 | 项目名称 约束： 1）敏感词校验 2）不能重复 |

<a id="api-invoiceexpenserulesprojectrulequery"></a>

### InvoiceExpenseRulesProjectRuleQuery — 查询项目费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.projectrule.query`](https://opendocs.alipay.com/apis/05ip0x)
- 路由：`GET /v3/alipay/ebpp/invoice/expenserules/projectrule/query`
- 请求类型：[InvoiceExpenseRulesProjectRuleQueryRequest](project_request.go#L72)
- 响应类型：[InvoiceExpenseRulesProjectRuleQueryResponse](project_response.go#L22)

查询项目费控规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 共同账号ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `project_id` | `ProjectId` | `*string` | query | 否 | 项目id（可空） |
| `employee_id` | `EmployeeId` | `*string` | query | 否 | 员工UID（可空） |
| `employee_open_id` | `EmployeeOpenId` | `*string` | query | 否 | 切换open_id后请使用：员工open_id（可空） |
| `page_num` | `PageNum` | `*string` | query | 否 | 页数 |
| `page_size` | `PageSize` | `*string` | query | 否 | 页大小 |

<a id="api-invoiceexpenserulesprojectrulesmodify"></a>

### InvoiceExpenseRulesProjectRulesModify — 修改项目费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.projectrules.modify`](https://opendocs.alipay.com/apis/05xnct)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/projectrules/modify`
- 请求类型：[InvoiceExpenseRulesProjectRulesModifyRequest](project_request.go#L58)
- 响应类型：[InvoiceExpenseRulesProjectRulesModifyResponse](project_response.go#L36)

修改项目费控规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `action` | `Action` | `*string` | JSON body | 未标注 | 修改操作 枚举值：MODIFY_RULE（修改费控条件），仅支持MODIFY_RULE |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `expense_ctrl_rule_info_group_list` | `ExpenseCtrlRuleInfoGroupList` | [`*[]ExpenseCtrRuleGroupInfo`](#model-expensectrrulegroupinfo) | JSON body | 未标注 | 费控规则列表 |
| `project_id` | `ProjectId` | `*string` | JSON body | 未标注 | 项目ID |

## 嵌套请求参数

<a id="model-expensectrrulegroupinfo"></a>

### ExpenseCtrRuleGroupInfo

类型定义：[ExpenseCtrRuleGroupInfo](project_model.go#L8)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `expense_ctrl_rule_info_list` | `ExpenseCtrlRuleInfoList` | [`*[]ExpenseCtrRuleInfo`](../README.md#model-expensectrruleinfo) | 未标注 | 费控条件列表 |
| `expense_type` | `ExpenseType` | `*string` | 未标注 | 费用类型 特殊说明：仅支持费用类型DEFAULT |
| `group_id` | `GroupId` | `*string` | 未标注 | 费控规则ID |
| `group_name` | `GroupName` | `*string` | 未标注 | 费控规则名称 约束： 1）敏感词校验 2）不能重复 |
