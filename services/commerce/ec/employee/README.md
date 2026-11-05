# employee 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/ec/employee`。本包提供 9 个支付宝 V3 REST 操作。

通过 `employee.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`EcEmployeeAdd`](#api-ecemployeeadd) | `alipay.commerce.ec.employee.add` |
| [`EcEmployeeDelete`](#api-ecemployeedelete) | `alipay.commerce.ec.employee.delete` |
| [`EcEmployeeIDListQuery`](#api-ecemployeeidlistquery) | `alipay.commerce.ec.employee.idlist.query` |
| [`EcEmployeeInfoModify`](#api-ecemployeeinfomodify) | `alipay.commerce.ec.employee.info.modify` |
| [`EcEmployeeInfoQuery`](#api-ecemployeeinfoquery) | `alipay.commerce.ec.employee.info.query` |
| [`EcEmployeeInviteQuery`](#api-ecemployeeinvitequery) | `alipay.commerce.ec.employee.invite.query` |
| [`EcEmployeeTitleCreate`](#api-ecemployeetitlecreate) | `alipay.commerce.ec.employee.title.create` |
| [`EcEmployeeTitleDelete`](#api-ecemployeetitledelete) | `alipay.commerce.ec.employee.title.delete` |
| [`EcEmployeeTitleModify`](#api-ecemployeetitlemodify) | `alipay.commerce.ec.employee.title.modify` |

## 使用示例

- [EcEmployeeAdd](api_employee_ecemployeeadd_example_test.go)
- [EcEmployeeDelete](api_employee_ecemployeedelete_example_test.go)
- [EcEmployeeIDListQuery](api_employee_ecemployeeidlistquery_example_test.go)
- [EcEmployeeInfoModify](api_employee_ecemployeeinfomodify_example_test.go)
- [EcEmployeeInfoQuery](api_employee_ecemployeeinfoquery_example_test.go)
- [EcEmployeeInviteQuery](api_employee_ecemployeeinvitequery_example_test.go)
- [EcEmployeeTitleCreate](api_employee_ecemployeetitlecreate_example_test.go)
- [EcEmployeeTitleDelete](api_employee_ecemployeetitledelete_example_test.go)
- [EcEmployeeTitleModify](api_employee_ecemployeetitlemodify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ecemployeeadd"></a>

### EcEmployeeAdd — 添加员工

- 官方接口：[`alipay.commerce.ec.employee.add`](https://opendocs.alipay.com/apis/05lo9i)
- 路由：`POST /v3/alipay/commerce/ec/employee/add`
- 请求类型：[EcEmployeeAddRequest](employee_request.go#L4)
- 响应类型：[EcEmployeeAddResponse](employee_response.go#L8)

企业添加员工，并返回员工签约激活链接和吱口令

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `accounting_entity_ids` | `AccountingEntityIds` | `*[]string` | JSON body | 未标注 | 员工所属核算主体，核算主体可用于管控不同员工的出资方式，建议和不同出资账户关联 |
| `create_share_code` | `CreateShareCode` | `*bool` | JSON body | 未标注 | 是否需要生成吱口令，默认不生成 |
| `department_ids` | `DepartmentIds` | `*[]string` | JSON body | 未标注 | 员工所属部门，支持多个部门，不传默认为根部门 |
| `employee_cert_no` | `EmployeeCertNo` | `*string` | JSON body | 未标注 | 员工证件号码 |
| `employee_cert_type` | `EmployeeCertType` | `*string` | JSON body | 未标注 | 员工证件类型 |
| `employee_email` | `EmployeeEmail` | `*string` | JSON body | 未标注 | 员工邮箱 |
| `employee_mobile` | `EmployeeMobile` | `*string` | JSON body | 未标注 | 员工手机号 |
| `employee_name` | `EmployeeName` | `*string` | JSON body | 未标注 | 员工姓名 |
| `employee_no` | `EmployeeNo` | `*string` | JSON body | 未标注 | 员工编号/工号 |
| `encrypt_cert_no` | `EncryptCertNo` | `*string` | JSON body | 未标注 | 加密证件号，需与证件类型搭配使用，可替代证件号码(employee_cert_no)字段 |
| `encrypt_mobile` | `EncryptMobile` | `*string` | JSON body | 未标注 | 加密手机号，可替代手机号(employee_mobile)字段 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 员工所属企业ID |
| `free_sign_token` | `FreeSignToken` | `*string` | JSON body | 未标注 | 员工签约授权密钥；<br> 该密钥需联系企业码业务小二申请，当传入密钥有效，则员工添加时会默认签约企业因公付，并激活员工。 |
| `identity` | `Identity` | `*string` | JSON body | 未标注 | 员工身份唯一标识 |
| `identity_open_id` | `IdentityOpenId` | `*string` | JSON body | 未标注 | 当身份类型为支付宝用户ID(identity_type=ALIPAY_USER_ID)时，传入user_id对应的open_id |
| `identity_type` | `IdentityType` | `*string` | JSON body | 未标注 | 员工身份类型 |
| `iot_check_type` | `IotCheckType` | `*string` | JSON body | 未标注 | 员工签约开通企业刷脸付时，刷脸核身的核验方式 |
| `iot_vid` | `IotVid` | `*string` | JSON body | 未标注 | 员工在企业人脸库的人脸唯一标识（该字段已作废，不支持自定义传入）(iot不支持自定义传入) |
| `label_names` | `LabelNames` | `*[]string` | JSON body | 未标注 | 员工标签，用于员工的打标分类，后续费控管理可使用标签进行控制，支持输入多个标签，如“差旅员工，用餐员工”等；若企业需要将离职员工统一管理且不变更员工状态，可打标“离职员工” |
| `profiles` | `Profiles` | `*string` | JSON body | 未标注 | 个性化信息 [详见文档](https://opendocs.alipay.com/pre-open/0ceh47?pathHash=14fac87c) |
| `role_list` | `RoleList` | `*[]string` | JSON body | 未标注 | 默认角色列表，目前只支持默认为 USER(目前仅支持普通员工，作为默认值，无需调用方传入) |
| `sign_return_url` | `SignReturnUrl` | `*string` | JSON body | 未标注 | 员工签约后回跳地址；支持 alipay scheme 协议地址，不传默认跳转企业码小程序首页 |
| `sign_url_carry_info` | `SignUrlCarryInfo` | `*bool` | JSON body | 未标注 | 签约链接页面是否需要自动回显员工的身份信息（手机或邮箱），默认不显示 |
| `withholding_sign_str` | `WithholdingSignStr` | `*string` | JSON body | 未标注 | 代扣签约字符串，员工需签约开通代扣时必传，否则无需传值；另外，如需参与企业码针对代扣的营销活动的服务商，必须要将该字段上传 |

<a id="api-ecemployeedelete"></a>

### EcEmployeeDelete — 删除员工

- 官方接口：[`alipay.commerce.ec.employee.delete`](https://opendocs.alipay.com/apis/05j06m)
- 路由：`POST /v3/alipay/commerce/ec/employee/delete`
- 请求类型：[EcEmployeeDeleteRequest](employee_request.go#L56)
- 响应类型：[EcEmployeeDeleteResponse](employee_response.go#L20)

删除员工并解绑因公付

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `employee_id` | `EmployeeId` | `*string` | JSON body | 未标注 | 员工id |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |

<a id="api-ecemployeeidlistquery"></a>

### EcEmployeeIDListQuery — 查询部门下员工id列表

- 官方接口：[`alipay.commerce.ec.employee.idlist.query`](https://opendocs.alipay.com/apis/05udr4)
- 路由：`GET /v3/alipay/commerce/ec/employee/idlist/query`
- 请求类型：[EcEmployeeIDListQueryRequest](employee_request.go#L130)
- 响应类型：[EcEmployeeIDListQueryResponse](employee_response.go#L23)

查询某个部门下员工id列表

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `department_id` | `DepartmentId` | `*string` | query | 否 | 部门id |
| `page_num` | `PageNum` | `*string` | query | 否 | 查询页数 |
| `page_size` | `PageSize` | `*string` | query | 否 | 每页查询大小，限制最大为1000 |

<a id="api-ecemployeeinfomodify"></a>

### EcEmployeeInfoModify — 修改员工基础信息

- 官方接口：[`alipay.commerce.ec.employee.info.modify`](https://opendocs.alipay.com/apis/05j1c0)
- 路由：`PUT /v3/alipay/commerce/ec/employee/info`
- 请求类型：[EcEmployeeInfoModifyRequest](employee_request.go#L64)
- 响应类型：[EcEmployeeInfoModifyResponse](employee_response.go#L35)

修改员工基础信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `accounting_entity_ids` | `AccountingEntityIds` | `*[]string` | JSON body | 未标注 | 员工所属核算主体，核算主体可用于管控不同员工的出资方式，建议和不同出资账户关联，一个员工只能有一个核算主体。 |
| `department_ids` | `DepartmentIds` | `*[]string` | JSON body | 未标注 | 员工所属部门。如果不传值，则不更新所属部门 |
| `employee_email` | `EmployeeEmail` | `*string` | JSON body | 未标注 | 员工邮箱。如果不传值，则不更新邮箱。 |
| `employee_id` | `EmployeeId` | `*string` | JSON body | 未标注 | 员工id |
| `employee_mobile` | `EmployeeMobile` | `*string` | JSON body | 未标注 | 员工手机号。如果不传值，则不更新手机号。 |
| `employee_name` | `EmployeeName` | `*string` | JSON body | 未标注 | 员工姓名 |
| `employee_no` | `EmployeeNo` | `*string` | JSON body | 未标注 | 员工工号。如果不传值，则不更新工号。 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业id |
| `label_names` | `LabelNames` | `*[]string` | JSON body | 未标注 | 员工标签，用于员工的打标分类，后续费控管理可使用标签进行控制，支持输入多个标签，如“差旅员工，用餐员工”等；<br> 使用场景：1. 请严格对标签分类，不要滥用员工标签，否则影响员工和费控管理； 2. 一个员工最多10个标签，如无必要请勿使用；3. 标签名只能包含字母（大小写）、数字或中文字符 |
| `profiles` | `Profiles` | `*string` | JSON body | 未标注 | 个性化信息。如果不传值，则不更新个性化信息。 [详见文档](https://opendocs.alipay.com/pre-open/0ceh47?pathHash=14fac87c) |
| `role_list` | `RoleList` | `*[]string` | JSON body | 未标注 | 角色列表，目前只支持填一种角色。如果不传值，则不更新角色。 |

<a id="api-ecemployeeinfoquery"></a>

### EcEmployeeInfoQuery — 查询员工详情

- 官方接口：[`alipay.commerce.ec.employee.info.query`](https://opendocs.alipay.com/apis/05uhw6)
- 路由：`GET /v3/alipay/commerce/ec/employee/info/query`
- 请求类型：[EcEmployeeInfoQueryRequest](employee_request.go#L142)
- 响应类型：[EcEmployeeInfoQueryResponse](employee_response.go#L41)

根据员工id或员工uid查询企业下某个员工的详细信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `open_id` | `OpenId` | `*string` | query | 否 | 蚂蚁统一会员ID |
| `user_id` | `UserId` | `*string` | query | 否 | 支付宝用户ID |
| `employee_id` | `EmployeeId` | `*string` | query | 否 | 员工id |
| `mobile` | `Mobile` | `*string` | query | 否 | 员工手机号码 |
| `alipay_logon_id` | `AlipayLogonId` | `*string` | query | 否 | 支付宝登录号（手机号或邮箱） |
| `employee_cert_type` | `EmployeeCertType` | `*string` | query | 否 | 员工证件类型 |
| `employee_cert_no` | `EmployeeCertNo` | `*string` | query | 否 | 员工证件号码 |
| `employee_email` | `EmployeeEmail` | `*string` | query | 否 | 员工邮箱 |

<a id="api-ecemployeeinvitequery"></a>

### EcEmployeeInviteQuery — 获取员工签约激活链接

- 官方接口：[`alipay.commerce.ec.employee.invite.query`](https://opendocs.alipay.com/apis/05j9je)
- 路由：`GET /v3/alipay/commerce/ec/employee/invite`
- 请求类型：[EcEmployeeInviteQueryRequest](employee_request.go#L164)
- 响应类型：[EcEmployeeInviteQueryResponse](employee_response.go#L46)

获取员工签约激活链接

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `enterprise_id` | `EnterpriseId` | `*string` | query | 否 | 企业id |
| `employee_id` | `EmployeeId` | `*string` | query | 否 | 企业码员工ID，签约适用对象为指定员工时填写 |
| `page_content_code` | `PageContentCode` | `*string` | query | 否 | 用于展示不同的页面文案，使用标准页面文案是无需传值 |
| `withholding_sign_str` | `WithholdingSignStr` | `*string` | query | 否 | 代扣签约串。需要签约代扣协议时，必传，返回的签约链接会拼接上代扣签约串。 |
| `create_share_code` | `CreateShareCode` | `*string` | query | 否 | 是否生成签约吱口令，不传默认为N |

<a id="api-ecemployeetitlecreate"></a>

### EcEmployeeTitleCreate — 员工抬头：新增员工抬头关系

- 官方接口：[`alipay.commerce.ec.employee.title.create`](https://opendocs.alipay.com/apis/05ip0u)
- 路由：`POST /v3/alipay/commerce/ec/employee/title`
- 请求类型：[EcEmployeeTitleCreateRequest](employee_request.go#L90)
- 响应类型：[EcEmployeeTitleCreateResponse](employee_response.go#L58)

批量新增员工抬头关系，单次新增最大一千条

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `employee_title_list` | `EmployeeTitleList` | [`*[]EmployeeTitleDTO`](#model-employeetitledto) | JSON body | 未标注 | 员工与抬头关系列表 |

<a id="api-ecemployeetitledelete"></a>

### EcEmployeeTitleDelete — 员工抬头：失效员工抬头关系

- 官方接口：[`alipay.commerce.ec.employee.title.delete`](https://opendocs.alipay.com/apis/05imon)
- 路由：`POST /v3/alipay/commerce/ec/employee/title/delete`
- 请求类型：[EcEmployeeTitleDeleteRequest](employee_request.go#L96)
- 响应类型：[EcEmployeeTitleDeleteResponse](employee_response.go#L63)

对员工抬头关系做失效处理。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `employee_title_list` | `EmployeeTitleList` | [`*[]EmployeeTitleDTO`](#model-employeetitledto) | JSON body | 未标注 | 员工抬头列表 |

<a id="api-ecemployeetitlemodify"></a>

### EcEmployeeTitleModify — 修改员工抬头信息

- 官方接口：[`alipay.commerce.ec.employee.title.modify`](https://opendocs.alipay.com/apis/05j4cg)
- 路由：`PUT /v3/alipay/commerce/ec/employee/title`
- 请求类型：[EcEmployeeTitleModifyRequest](employee_request.go#L102)
- 响应类型：[EcEmployeeTitleModifyResponse](employee_response.go#L68)

修改员工抬头信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | JSON body | 未标注 | 共同账号id，与enterprise_id两者必填其一 |
| `create_by` | `CreateBy` | `*string` | JSON body | 未标注 | 创建人 |
| `employee_id` | `EmployeeId` | `*string` | JSON body | 未标注 | 企业域定义的员工id，与user_id两者必填其一 |
| `enterprise_id` | `EnterpriseId` | `*string` | JSON body | 未标注 | 企业域定义下的企业id，与accountId两者必填其一 |
| `modify_by` | `ModifyBy` | `*string` | JSON body | 未标注 | 修改人 |
| `multi_title_list` | `MultiTitleList` | [`*[]EcEmployeeTitleModifyTaxNoAndTitleId`](#model-ecemployeetitlemodifytaxnoandtitleid) | JSON body | 未标注 | 绑定的抬头ID列表 |
| `new_title_id` | `NewTitleId` | `*string` | JSON body | 未标注 | 修改后的抬头id |
| `old_title_id` | `OldTitleId` | `*string` | JSON body | 未标注 | 修改前的抬头id(数据库自行查询是否有有效抬头) |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝open_id，与employee_id两者必填其一 |
| `tax_register_no` | `TaxRegisterNo` | `*string` | JSON body | 未标注 | 税号 |
| `title_tag` | `TitleTag` | `*string` | JSON body | 未标注 | 员工抬头关系标 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 支付宝id2088，与employee_id两者必填其一 |

## 嵌套请求参数

<a id="model-ecemployeetitlemodifytaxnoandtitleid"></a>

### EcEmployeeTitleModifyTaxNoAndTitleId

类型定义：[EcEmployeeTitleModifyTaxNoAndTitleId](employee_model.go#L18)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `tax_register_no` | `TaxRegisterNo` | `*string` | 未标注 | 税号 |
| `title_id` | `TitleId` | `*string` | 未标注 | 抬头ID |

<a id="model-employeetitledto"></a>

### EmployeeTitleDTO

类型定义：[EmployeeTitleDTO](employee_model.go#L98)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `account_id` | `AccountId` | `*string` | 未标注 | 共同账户id,与enterprise_id两者必填其一 |
| `create_by` | `CreateBy` | `*string` | 未标注 | 创建人 |
| `employee_id` | `EmployeeId` | `*string` | 未标注 | 企业域定义的员工id，与user_id两者必填其一 |
| `enterprise_id` | `EnterpriseId` | `*string` | 未标注 | 企业域定义的企业id，与account_id两者必填其一 |
| `modify_by` | `ModifyBy` | `*string` | 未标注 | 修改人 |
| `open_id` | `OpenId` | `*string` | 未标注 | 员工支付宝账号open_id,与employee_id两者必填其一 |
| `tax_register_no` | `TaxRegisterNo` | `*string` | 未标注 | 税号 |
| `title_id` | `TitleId` | `*string` | 未标注 | 抬头id |
| `title_tag` | `TitleTag` | `*string` | 未标注 | 关系标记 |
| `user_id` | `UserId` | `*string` | 未标注 | 员工支付宝账号2088,与employee_id两者必填其一 |
