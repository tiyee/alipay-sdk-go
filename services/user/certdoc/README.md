# certdoc 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/user/certdoc`。本包提供 2 个支付宝 V3 REST 操作。

通过 `certdoc.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CertDocCertVerifyConsult`](#api-certdoccertverifyconsult) | `alipay.user.certdoc.certverify.consult` |
| [`CertDocCertVerifyPreConsult`](#api-certdoccertverifypreconsult) | `alipay.user.certdoc.certverify.preconsult` |

## 使用示例

- [CertDocCertVerifyConsult](api_certdoc_certdoccertverifyconsult_example_test.go)
- [CertDocCertVerifyPreConsult](api_certdoc_certdoccertverifypreconsult_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-certdoccertverifyconsult"></a>

### CertDocCertVerifyConsult — 实名证件信息比对验证咨询

- 官方接口：[`alipay.user.certdoc.certverify.consult`](https://opendocs.alipay.com/apis/05wvw5)
- 路由：`GET /v3/alipay/user/certdoc/certverify/consult`
- 请求类型：[UserCertDocCertVerifyConsultRequest](certdoc_request.go#L20)
- 响应类型：[UserCertDocCertVerifyConsultResponse](certdoc_response.go#L4)

商户在预咨询接口（接口1）预先录入某用户的实名证件信息（例如姓名、身份证号等），获得用户授权后调用本咨询接口（接口2），将对比该用户的支付宝实名信息与（接口1）传入的信息是否一致。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `verify_id` | `VerifyId` | `*string` | query | 否 | 信息校验验证ID。通过[alipay.user.certdoc.certverify.preconsult](https://opendocs.alipay.com/apis/api_2/alipay.user.certdoc.certverify.preconsult)(实名证件信息比对验证预咨询)接口获取。 |

<a id="api-certdoccertverifypreconsult"></a>

### CertDocCertVerifyPreConsult — 实名证件信息比对验证预咨询

- 官方接口：[`alipay.user.certdoc.certverify.preconsult`](https://opendocs.alipay.com/apis/05wfau)
- 路由：`POST /v3/alipay/user/certdoc/certverify/preconsult`
- 请求类型：[UserCertDocCertVerifyPreConsultRequest](certdoc_request.go#L4)
- 响应类型：[UserCertDocCertVerifyPreConsultResponse](certdoc_response.go#L14)

商户在该接口（接口1）预先录入某用户的实名证件信息（例如姓名、身份证号等），获得用户授权后调用咨询接口（接口2），将对比该用户的支付宝实名信息与该（接口1）传入的信息是否一致。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 证件号 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 证件类型。暂仅支持 IDENTITY_CARD （身份证）。 |
| `ext_info` | `ExtInfo` | `*string` | JSON body | 未标注 | 拓展字段,JSON格式 |
| `logon_id` | `LogonId` | `*string` | JSON body | 未标注 | 支付宝登录名 |
| `mobile` | `Mobile` | `*string` | JSON body | 未标注 | 手机号码 |
| `user_name` | `UserName` | `*string` | JSON body | 未标注 | 真实姓名 |
