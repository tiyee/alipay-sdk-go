# userinvoiceinfo 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/boss/fnc/userinvoiceinfo`。本包提供 3 个支付宝 V3 REST 操作。

通过 `userinvoiceinfo.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`FncUserInvoiceInfoCreate`](#api-fncuserinvoiceinfocreate) | `alipay.boss.fnc.userinvoiceinfo.create` |
| [`FncUserInvoiceInfoModify`](#api-fncuserinvoiceinfomodify) | `alipay.boss.fnc.userinvoiceinfo.modify` |
| [`FncUserInvoiceInfoQuery`](#api-fncuserinvoiceinfoquery) | `alipay.boss.fnc.userinvoiceinfo.query` |

## 使用示例

- [FncUserInvoiceInfoCreate](api_userinvoiceinfo_fncuserinvoiceinfocreate_example_test.go)
- [FncUserInvoiceInfoModify](api_userinvoiceinfo_fncuserinvoiceinfomodify_example_test.go)
- [FncUserInvoiceInfoQuery](api_userinvoiceinfo_fncuserinvoiceinfoquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-fncuserinvoiceinfocreate"></a>

### FncUserInvoiceInfoCreate — 新增客户开票资料

- 官方接口：[`alipay.boss.fnc.userinvoiceinfo.create`](https://opendocs.alipay.com/pre-apis/05xgfe)
- 路由：`POST /v3/alipay/boss/fnc/userinvoiceinfo/create`
- 请求类型：[BossFncUserInvoiceInfoCreateRequest](userinvoiceinfo_request.go#L4)
- 响应类型：[BossFncUserInvoiceInfoCreateResponse](userinvoiceinfo_response.go#L8)

申请开票alipay.boss.fnc.invoice.create之前需要维护的开票资料信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `accept_electronic` | `AcceptElectronic` | `*bool` | JSON body | 未标注 | 是否接受电子票 true:是，false:否 |
| `address` | `Address` | `*string` | JSON body | 未标注 | 公司注册地址，公司营业执照上登记的住址，一般纳税人必须填写，小规模纳税人无需填写 |
| `addressing` | `Addressing` | `*bool` | JSON body | 未标注 | 开票资料pid寻址方式 true:pid寻址mid,优先以mid维度保存开票资料,若无mid则以pid维度保存开票资料 false:不寻址mid,以pid维度保存开票资料;非主站商户体系都是以pid（ipRoleId）维度保存开票资料，该值直接填为false |
| `auto` | `Auto` | `*bool` | JSON body | 未标注 | 是否自动申请开票 true:是，false:否 |
| `bank_account` | `BankAccount` | `*string` | JSON body | 未标注 | 银行账户，公司银行账号<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `bank_name` | `BankName` | `*string` | JSON body | 未标注 | 开户行，办理银行开户手续的营业网点<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `business_licence_url` | `BusinessLicenceUrl` | `*string` | JSON body | 未标注 | 营业执照地址，营业执照上传oss上的地址 |
| `hold` | `Hold` | `*bool` | JSON body | 未标注 | 是否暂停开票，商户添加的开票资料就是不暂停开票，如果后续要暂停开票，只能去结算中台处理。 |
| `ip_role_id` | `IpRoleId` | `*string` | JSON body | 未标注 | 商户的pid（ipRoleId） |
| `open_account_permit_url` | `OpenAccountPermitUrl` | `*string` | JSON body | 未标注 | 银行开户许可证附件，银行开户许可证上传oss上的地址 |
| `operator` | `Operator` | `*string` | JSON body | 未标注 | 当前操作人 |
| `operator_type` | `OperatorType` | `*string` | JSON body | 未标注 | 当前操作人类型,01:商户 02:管理员(小二) |
| `other_qualification_url` | `OtherQualificationUrl` | `*string` | JSON body | 未标注 | 其它资质证明地址，其他资质证明附件上传oss上的地址 |
| `ou` | `Ou` | `*bool` | JSON body | 未标注 | 外部只允许添加商户开票资料，该值填写为false |
| `tax_no` | `TaxNo` | `*string` | JSON body | 未标注 | 纳税人识别号，税务登记证上的号码。一般纳税人必须填写，小规模纳税人如果是商户也必须填写，个人无需填写 |
| `tax_payer_quali_valid` | `TaxPayerQualiValid` | `*string` | JSON body | 未标注 | 纳税人资格开始时间 （格式：时间戳）<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `tax_payer_qualification` | `TaxPayerQualification` | `*string` | JSON body | 未标注 | 纳税人资格种类，01：一般纳税人；02：小规模纳税人；03：国际商户 |
| `tax_qualification_url` | `TaxQualificationUrl` | `*string` | JSON body | 未标注 | 一般纳税人资格证书地址，一般纳税人资格证书上传oss上的地址 |
| `tax_reg_cert_url` | `TaxRegCertUrl` | `*string` | JSON body | 未标注 | 税务登记证地址，税务登记证上传oss上的地址 |
| `telephone` | `Telephone` | `*string` | JSON body | 未标注 | 公司注册电话（手机号和座机均可）<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 发票抬头，票面信息上的抬头信息 |
| `user_mail_info_order_list` | `UserMailInfoOrderList` | [`*[]UserMailInfoOrder`](#model-usermailinfoorder) | JSON body | 未标注 | 收件人列表,若不修改此项可以不填<br> 选择非电子票时，邮寄信息必填，且保证邮寄信息的ipRoleId和开票资料的ipRoleId相等 |

<a id="api-fncuserinvoiceinfomodify"></a>

### FncUserInvoiceInfoModify — 修改用户开票资料

- 官方接口：[`alipay.boss.fnc.userinvoiceinfo.modify`](https://opendocs.alipay.com/pre-apis/05xshh)
- 路由：`POST /v3/alipay/boss/fnc/userinvoiceinfo/modify`
- 请求类型：[BossFncUserInvoiceInfoModifyRequest](userinvoiceinfo_request.go#L52)
- 响应类型：[BossFncUserInvoiceInfoModifyResponse](userinvoiceinfo_response.go#L14)

通过查询用户的开票资料（alipay.boss.fnc.userinvoiceinfo.query）后进行开票资料的修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `accept_electronic` | `AcceptElectronic` | `*bool` | JSON body | 未标注 | 是否接受电子票 true:是，false:否 |
| `address` | `Address` | `*string` | JSON body | 未标注 | 公司注册地址，公司营业执照上登记的住址，一般纳税人必须填写，小规模纳税人无需填写 |
| `auto` | `Auto` | `*bool` | JSON body | 未标注 | 是否自动申请开票 true:是；false:否 |
| `bank_account` | `BankAccount` | `*string` | JSON body | 未标注 | 银行账户，公司银行账号<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `bank_name` | `BankName` | `*string` | JSON body | 未标注 | 开户行，办理银行开户手续的营业网点<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `business_licence_url` | `BusinessLicenceUrl` | `*string` | JSON body | 未标注 | 营业执照地址，营业执照上传oss上的地址 |
| `hold` | `Hold` | `*bool` | JSON body | 未标注 | 是否暂停开票，商户修改无效，如果要暂停开票，只能去结算中台处理。 |
| `id` | `Id` | `*string` | JSON body | 未标注 | 开票资料ID |
| `open_account_permit_url` | `OpenAccountPermitUrl` | `*string` | JSON body | 未标注 | 银行开户许可证附件，银行开户许可证上传oss上的地址 |
| `operator` | `Operator` | `*string` | JSON body | 未标注 | 当前操作人 |
| `operator_type` | `OperatorType` | `*string` | JSON body | 未标注 | 当前操作人类型,01:商户；02:管理员(小二) |
| `other_qualification_url` | `OtherQualificationUrl` | `*string` | JSON body | 未标注 | 其它资质证明地址，其他资质证明附件上传oss上的地址 |
| `tax_no` | `TaxNo` | `*string` | JSON body | 未标注 | 纳税人识别号，税务登记证上的号码。一般纳税人必须填写，小规模纳税人如果是商户也必须填写，个人无需填写 |
| `tax_payer_quali_valid` | `TaxPayerQualiValid` | `*string` | JSON body | 未标注 | 纳税人资格开始时间 （格式：时间戳）<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `tax_payer_qualification` | `TaxPayerQualification` | `*string` | JSON body | 未标注 | 纳税人资格种类，01：一般纳税人；02：小规模纳税人；03：国际商户 |
| `tax_qualification_url` | `TaxQualificationUrl` | `*string` | JSON body | 未标注 | 一般纳税人资格证书地址，一般纳税人资格证书上传oss上的地址 |
| `tax_reg_cert_url` | `TaxRegCertUrl` | `*string` | JSON body | 未标注 | 税务登记证地址，税务登记证上传oss上的地址 |
| `telephone` | `Telephone` | `*string` | JSON body | 未标注 | 公司注册电话<br> 一般纳税人必须填写，小规模纳税人无需填写 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 发票抬头，票面信息上的抬头信息 |
| `user_mail_info_order_list` | `UserMailInfoOrderList` | [`*[]UserMailInfoOrder`](#model-usermailinfoorder) | JSON body | 未标注 | 收件人列表,若不修改此项可以不填<br> 选择非电子票时，邮寄信息必填，且保证邮寄信息的ipRoleId和开票资料的ipRoleId相等 |

<a id="api-fncuserinvoiceinfoquery"></a>

### FncUserInvoiceInfoQuery — 根据PID获取相关的开票资料

- 官方接口：[`alipay.boss.fnc.userinvoiceinfo.query`](https://opendocs.alipay.com/pre-apis/05xmie)
- 路由：`GET /v3/alipay/boss/fnc/userinvoiceinfo/query`
- 请求类型：[BossFncUserInvoiceInfoQueryRequest](userinvoiceinfo_request.go#L96)
- 响应类型：[BossFncUserInvoiceInfoQueryResponse](userinvoiceinfo_response.go#L17)

根据PID获取相关的开票资料

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `pid` | `Pid` | `*string` | query | 否 | 开票pid/mid/ou，唯一标识商户信息/集团用户信息的ID |

## 嵌套请求参数

<a id="model-usermailinfoorder"></a>

### UserMailInfoOrder

类型定义：[UserMailInfoOrder](userinvoiceinfo_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `city` | `City` | `*string` | 未标注 | 联系所在城市 |
| `country` | `Country` | `*string` | 未标注 | 联系人国家 |
| `county_district` | `CountyDistrict` | `*string` | 未标注 | 联系人所在县/区 |
| `detail_address` | `DetailAddress` | `*string` | 未标注 | 联系所在详细地址 |
| `email` | `Email` | `*string` | 未标注 | 电子邮箱 |
| `ip_role_id` | `IpRoleId` | `*string` | 未标注 | 商户ipRole(pid) |
| `name` | `Name` | `*string` | 未标注 | 联系人名字 |
| `province` | `Province` | `*string` | 未标注 | 联系人省份 |
| `street` | `Street` | `*string` | 未标注 | 联系人所在街道 |
| `telephone` | `Telephone` | `*string` | 未标注 | 联系人电话 |
