# scene 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules/scene`。本包提供 3 个支付宝 V3 REST 操作。

通过 `scene.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`InvoiceExpenseRulesSceneRuleCreate`](#api-invoiceexpenserulesscenerulecreate) | `alipay.ebpp.invoice.expenserules.scenerule.create` |
| [`InvoiceExpenseRulesSceneRuleModify`](#api-invoiceexpenserulesscenerulemodify) | `alipay.ebpp.invoice.expenserules.scenerule.modify` |
| [`InvoiceExpenseRulesSceneRuleQuery`](#api-invoiceexpenserulesscenerulequery) | `alipay.ebpp.invoice.expenserules.scenerule.query` |

## 使用示例

- [InvoiceExpenseRulesSceneRuleCreate](api_scene_invoiceexpenserulesscenerulecreate_example_test.go)
- [InvoiceExpenseRulesSceneRuleModify](api_scene_invoiceexpenserulesscenerulemodify_example_test.go)
- [InvoiceExpenseRulesSceneRuleQuery](api_scene_invoiceexpenserulesscenerulequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceexpenserulesscenerulecreate"></a>

### InvoiceExpenseRulesSceneRuleCreate — 创建费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.scenerule.create`](https://opendocs.alipay.com/apis/05j82p)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/scenerule/create`
- 请求类型：[InvoiceExpenseRulesSceneRuleCreateRequest](scene_request.go#L8)
- 响应类型：[InvoiceExpenseRulesSceneRuleCreateResponse](scene_response.go#L4)

创建费控规则，例如：创建一条有效期为2021-08-01 00:00:00到2031-08-01 00:00:00的福利餐费控规则，限制对应员工的单日消费金额不超过100元；

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账号ID |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 有效期截止 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 有效期起始 |
| `employee_list` | `EmployeeList` | `*[]string` | JSON body | 未标注 | 切换open_id前请使用：员工支付宝uid列表 特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整 |
| `employee_open_id_list` | `EmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换open_id后请使用：员工open_id/企业码员工ID列表 特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业码企业id |
| `expense_ctrl_rule_info_list` | `ExpenseCtrlRuleInfoList` | [`*[]ExpenseCtrRuleInfo`](../README.md#model-expensectrruleinfo) | JSON body | 未标注 | 费控条件列表 特殊说明： 1）至少存在日额度（QUOTA_DAY）、月额度（QUOTA_MONTH）、有效期总额度（QUOTA_TOTAL）三者中的一个额度条件； 2）如果费用类型为MEAL，费控维度商户（MEAL_MERCHANT）和商户类型（MCC）对应的费控条件必须存在其一且不能同时存在； 3）如果费用类型为METRO，费控维度地铁卡类型（CARD_TYPE）对应的费控条件必须存在； 4）如果因公场景为OVERTIME，费控维度时间段（ALARM_CLOCK_TIME）对应的费控条件必须存在； 5）不能存在重复的费控维度对应的费控条件； 6）非MEAL费用类型，商户仅支持MERCHANT，不支持MEAL_MERCHANT |
| `expense_type` | `ExpenseType` | `*string` | JSON body | 未标注 | 费用类型 |
| `payment_policy` | `PaymentPolicy` | `*string` | JSON body | 未标注 | 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）, COMBINATION（因公账户和个人账户组合支付） |
| `scene_type` | `SceneType` | `*string` | JSON body | 未标注 | 因公场景 |
| `standard_desc` | `StandardDesc` | `*string` | JSON body | 未标注 | 费控规则说明 特殊说明：敏感词校验 |
| `standard_name` | `StandardName` | `*string` | JSON body | 未标注 | 费控规则名称 特殊说明： 1）敏感词校验 2）不能重复，若需要重复联系支持人员定向放开 |

<a id="api-invoiceexpenserulesscenerulemodify"></a>

### InvoiceExpenseRulesSceneRuleModify — 修改费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.scenerule.modify`](https://opendocs.alipay.com/apis/05xk1w)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/scenerule/modify`
- 请求类型：[InvoiceExpenseRulesSceneRuleModifyRequest](scene_request.go#L38)
- 响应类型：[InvoiceExpenseRulesSceneRuleModifyResponse](scene_response.go#L10)

修改费控规则

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业ID |
| `action` | `Action` | `*string` | JSON body | 未标注 | 修改操作 枚举值： ADD_RULE（新增费控条件），DELETE_RULE（删除费控条件），MODIFY_RULE（修改费控条件），MODIFY_STANDARD（修改费控规则基本信息） |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `effective_end_date` | `EffectiveEndDate` | `*string` | JSON body | 未标注 | 有效期截止 特殊说明：修改费控规则基本信息时必传 |
| `effective_start_date` | `EffectiveStartDate` | `*string` | JSON body | 未标注 | 有效期起始 特殊说明：修改费控规则基本信息时必传 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业码企业id |
| `expense_ctrl_rule_info_list` | `ExpenseCtrlRuleInfoList` | [`*[]ExpenseCtrRuleInfo`](../README.md#model-expensectrruleinfo) | JSON body | 未标注 | 费控规则条件列表 特殊说明：新增费控条件/修改费控条件/删除费控条件操作时必填 |
| `payment_policy` | `PaymentPolicy` | `*string` | JSON body | 未标注 | 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）, COMBINATION（因公账户和个人账户组合支付） |
| `standard_desc` | `StandardDesc` | `*string` | JSON body | 未标注 | 费控规则说明 特殊说明： 1）敏感词校验 2）修改费控规则基本信息时必传 |
| `standard_id` | `StandardId` | `*string` | JSON body | 未标注 | 费控规则ID |
| `standard_name` | `StandardName` | `*string` | JSON body | 未标注 | 费控规则名称 特殊说明： 1）敏感词校验 2）修改费控规则基本信息时必传 3）不能重复，若需要重复联系支持人员定向放开 |

<a id="api-invoiceexpenserulesscenerulequery"></a>

### InvoiceExpenseRulesSceneRuleQuery — 查询费控规则信息

- 官方接口：[`alipay.ebpp.invoice.expenserules.scenerule.query`](https://opendocs.alipay.com/apis/05x1fs)
- 路由：`GET /v3/alipay/ebpp/invoice/expenserules/scenerule/query`
- 请求类型：[InvoiceExpenseRulesSceneRuleQueryRequest](scene_request.go#L64)
- 响应类型：[InvoiceExpenseRulesSceneRuleQueryResponse](scene_response.go#L16)

查询费控规则信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业码企业id |
| `standard_id` | `StandardId` | `*string` | query | 否 | 费控规则ID |
| `page_num` | `PageNum` | `*string` | query | 否 | 页数 |
| `page_size` | `PageSize` | `*string` | query | 否 | 页大小 |
