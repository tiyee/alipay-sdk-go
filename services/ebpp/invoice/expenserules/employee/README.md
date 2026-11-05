# employee 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules/employee`。本包提供 3 个支付宝 V3 REST 操作。

通过 `employee.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`InvoiceExpenseRulesEmployeeQuery`](#api-invoiceexpenserulesemployeequery) | `alipay.ebpp.invoice.expenserules.employee.query` |
| [`InvoiceExpenseRulesEmployeeRulesQuery`](#api-invoiceexpenserulesemployeerulesquery) | `alipay.ebpp.invoice.expenserules.employeerules.query` |
| [`InvoiceExpenseRulesGroupEmployeeModify`](#api-invoiceexpenserulesgroupemployeemodify) | `alipay.ebpp.invoice.expenserules.groupemployee.modify` |

## 使用示例

- [InvoiceExpenseRulesEmployeeQuery](api_employee_invoiceexpenserulesemployeequery_example_test.go)
- [InvoiceExpenseRulesEmployeeRulesQuery](api_employee_invoiceexpenserulesemployeerulesquery_example_test.go)
- [InvoiceExpenseRulesGroupEmployeeModify](api_employee_invoiceexpenserulesgroupemployeemodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoiceexpenserulesemployeequery"></a>

### InvoiceExpenseRulesEmployeeQuery — 查询费控规则员工

- 官方接口：[`alipay.ebpp.invoice.expenserules.employee.query`](https://opendocs.alipay.com/apis/05j4cj)
- 路由：`GET /v3/alipay/ebpp/invoice/expenserules/employee/query`
- 请求类型：[InvoiceExpenseRulesEmployeeQueryRequest](employee_request.go#L24)
- 响应类型：[InvoiceExpenseRulesEmployeeQueryResponse](employee_response.go#L4)

查询费控规则员工

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 共同账号ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业码企业id |
| `standard_id` | `StandardId` | `*string` | query | 否 | 费控规则ID |
| `page_num` | `PageNum` | `*string` | query | 否 | 页数 |
| `page_size` | `PageSize` | `*string` | query | 否 | 页大小 |

<a id="api-invoiceexpenserulesemployeerulesquery"></a>

### InvoiceExpenseRulesEmployeeRulesQuery — 查询员工绑定的所有费控规则

- 官方接口：[`alipay.ebpp.invoice.expenserules.employeerules.query`](https://opendocs.alipay.com/apis/05j4ci)
- 路由：`GET /v3/alipay/ebpp/invoice/expenserules/employeerules/query`
- 请求类型：[InvoiceExpenseRulesEmployeeRulesQueryRequest](employee_request.go#L40)
- 响应类型：[InvoiceExpenseRulesEmployeeRulesQueryResponse](employee_response.go#L18)

分页查询指定员工绑定的所有费控规则，默认返回所有有效和无效规则，查询结果按费控规则的创建时间倒序

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | query | 否 | 企业ID |
| `agreement_no` | `AgreementNo` | `*string` | query | 否 | 授权签约协议号 |
| `employee_id` | `EmployeeId` | `*string` | query | 否 | 未切换open_id时使用此字段：员工ID |
| `employee_open_id` | `EmployeeOpenId` | `*string` | query | 否 | 切换open_id后使用此字段：员工的open_id |
| `page_num` | `PageNum` | `*string` | query | 否 | 页码，默认值为1 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页行数，默认值为20 |

<a id="api-invoiceexpenserulesgroupemployeemodify"></a>

### InvoiceExpenseRulesGroupEmployeeModify — 修改费控规则员工

- 官方接口：[`alipay.ebpp.invoice.expenserules.groupemployee.modify`](https://opendocs.alipay.com/apis/05j82o)
- 路由：`POST /v3/alipay/ebpp/invoice/expenserules/groupemployee/modify`
- 请求类型：[InvoiceExpenseRulesGroupEmployeeModifyRequest](employee_request.go#L4)
- 响应类型：[InvoiceExpenseRulesGroupEmployeeModifyResponse](employee_response.go#L30)

修改费控规则员工

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 企业共同账户id |
| `add_employee_list` | `AddEmployeeList` | `*[]string` | JSON body | 未标注 | 未切换 open_id 时请使用此字段： 需要添加的员工UID列表 特殊说明：一次最多50个 |
| `add_employee_open_id_list` | `AddEmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换 open_id 后请使用此字段： 需要添加的open_id/企业码员工ID列表 特殊说明：一次最多50个 |
| `agreement_no` | `AgreementNo` | `*string` | JSON body | 未标注 | 授权签约协议号 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业码企业id |
| `group_id_list` | `GroupIdList` | `*[]string` | JSON body | 未标注 | 费控规则ID列表 |
| `remove_employee_list` | `RemoveEmployeeList` | `*[]string` | JSON body | 未标注 | 未切换 open_id 时请使用此字段： 需要移除的员工UID列表 特殊说明：一次最多50个 |
| `remove_employee_open_id_list` | `RemoveEmployeeOpenIdList` | `*[]string` | JSON body | 未标注 | 切换 open_id 后请使用此字段： 需要移除的open_id/企业码员工ID列表 特殊说明：一次最多50个 |
