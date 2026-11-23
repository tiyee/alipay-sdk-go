# operation 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/sp/operation`。本包提供 3 个支付宝 V3 REST 操作。

通过 `operation.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SpOperationApply`](#api-spoperationapply) | `alipay.open.sp.operation.apply` |
| [`SpOperationQRCodeQuery`](#api-spoperationqrcodequery) | `alipay.open.sp.operation.qrcode.query` |
| [`SpOperationResultQuery`](#api-spoperationresultquery) | `alipay.open.sp.operation.result.query` |

## 使用示例

- [SpOperationApply](api_operation_spoperationapply_example_test.go)
- [SpOperationQRCodeQuery](api_operation_spoperationqrcodequery_example_test.go)
- [SpOperationResultQuery](api_operation_spoperationresultquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-spoperationapply"></a>

### SpOperationApply — 向商户发起代运营操作

- 官方接口：[`alipay.open.sp.operation.apply`](https://opendocs.alipay.com/open-v3/a8869af2_alipay.open.sp.operation.apply?scene=common&pathHash=6c69600a)
- 路由：`POST /v3/alipay/open/sp/operation/apply`
- 请求类型：[SpOperationApplyRequest](operation_request.go#L4)
- 响应类型：[SpOperationApplyResponse](operation_response.go#L4)

服务商向商户发起代运营操作（账号绑定或者代运营授权）

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `access_product_code` | `AccessProductCode` | `*string` | JSON body | 未标注 | 接入的产品编号。 枚举如下： * OPENAPI_BIND_DEFAULT：操作类型为账号绑定。 * OPENAPI_AUTH_DEFAULT：操作类型为代运营授权时。 |
| `alipay_account` | `AlipayAccount` | `*string` | JSON body | 未标注 | 支付宝登录账号，通常为手机号或者邮箱。 间连场景必填。 直连场景选填，特别注意merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。 |
| `isv_scene_permissions` | `IsvScenePermissions` | `*string` | JSON body | 未标注 | 场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2; 说明： * 本参数和access_product_code只需要传一个。 * 场景 + 权限Code含义： SHOP_MANAGE:SHOP_MANAGE_BASE：管理门店信息 MINI_APP_OPER:MINI_APP_OPER_BASE：运营支付宝小程序 PROMOTION_MANAGE:PROMOTION_MANAGE_BASE：运营营销活动 OPERATION_POINTS:OPERATION_POINTS_BASE：管理运营积分 INCENTIVE_POINT_MANAGE:INCENTIVE_POINT_MANAGE_BASE：管理激励点数 |
| `merchant_no` | `MerchantNo` | `*string` | JSON body | 未标注 | 支付宝商户号。 间连场景，merchant_no必填，传入商户smid，特别注意仅支持2088开头的间连商户。 直连场景，merchant_no选填，传入商户支付宝pid，特别注意merchant_no和alipay_account不能同时为空，优先取merchant_no。 |
| `operate_type` | `OperateType` | `*string` | JSON body | 未标注 | 代运营操作类型。枚举如下： * ACCOUNT_BIND：代表绑定支付宝账号，仅对于间连商户。 * OPERATION_AUTH：代表代运营授权，支持间连和直连商户，其中间连场景包含绑定支付宝账号。 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部操作流水，ISV自定义。每次操作需要确保唯一。 |

<a id="api-spoperationqrcodequery"></a>

### SpOperationQRCodeQuery — 查询代运营授权或者账号绑定二维码

- 官方接口：[`alipay.open.sp.operation.qrcode.query`](https://opendocs.alipay.com/open-v3/8245f76c_alipay.open.sp.operation.qrcode.query?scene=common&pathHash=c82f7840)
- 路由：`GET /v3/alipay/open/sp/operation/qrcode/query`
- 请求类型：[SpOperationQRCodeQueryRequest](operation_request.go#L20)
- 响应类型：[SpOperationQRCodeQueryResponse](operation_response.go#L13)

面向服务商，查询发起代运营授权或者账号绑定的二维码。代间联商户代运营时，推荐使用接口同步返回的recommend_accounts字段中的账号扫码确认。

请求约束：SpOperationQRCodeQueryRequest is the request type for alipay.open.sp.operation.qrcode.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_biz_no` | `OutBizNo` | `*string` | query | 否 | 外部操作流水，由服务商自定义，需确保每次操作唯一。仅支持数字、字母、下划线组合。 |
| `operate_type` | `OperateType` | `*string` | query | 否 | 代运营操作类型。取值如下： * ACCOUNT_BIND：账号绑定，仅支持间连商户。 * OPERATION_AUTH：代运营授权，支持间连、直连商户。 |
| `merchant_no` | `MerchantNo` | `*string` | query | 否 | 支付宝商户号。注意仅支持 2088 开头的间连商户。 若被代运营者是间连商户，则 merchant_no 必填； 若为直连商户，则 merchant_no 和 alipay_account 不能同时为空，都有值优先取 merchant_no。 |
| `alipay_account` | `AlipayAccount` | `*string` | query | 否 | 支付宝登录账号。通常为手机号或者邮箱。 若被代运营者是间连商户，该字段无需填写。 若为直连商户，则 merchant_no 和 alipay_account 不能同时为空，都有值优先取 merchant_no。 |
| `access_product_code` | `AccessProductCode` | `*string` | query | 否 | 接入的产品编号。 操作类型为绑定的，填OPENAPI_BIND_DEFAULT; 操作类型为授权的，填OPENAPI_AUTH_DEFAULT |
| `isv_scene_permissions` | `IsvScenePermissions` | `*string` | query | 否 | 场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2; |

<a id="api-spoperationresultquery"></a>

### SpOperationResultQuery — 查询服务商代运营操作结果

- 官方接口：[`alipay.open.sp.operation.result.query`](https://opendocs.alipay.com/open-v3/229831e2_alipay.open.sp.operation.result.query?scene=common&pathHash=4c2b6906)
- 路由：`GET /v3/alipay/open/sp/operation/result/query`
- 请求类型：[SpOperationResultQueryRequest](operation_request.go#L36)
- 响应类型：[SpOperationResultQueryResponse](operation_response.go#L25)

查询服务商代运营操作（代运营授权或者账号绑定）结果。提供两种查询方式：
方式一：operate_type + batch_no，提供基于操作批次的精确查询，只查询本操作批次的代运营操作结果，不反映最终代运营操作结果。 方式一查询方式优先级更高。
方式二：operate_type + access_product_code + merchant_no + alipay_account 参数组合，可以查询代运营操作最终结果，任意一个操作批次成功都会体现在最终结果中。
方式三：operate_type + isv_scene_permissions + merchant_no + alipay_account 参数组合，可以查询代运营操作最终结果，任意一个操作批次成功都会体现在最终结果中（仅支持查询授权结果）。

请求约束：SpOperationResultQueryRequest is the request type for alipay.open.sp.operation.result.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `operate_type` | `OperateType` | `*string` | query | 否 | 代运营操作类型。取值如下 ACCOUNT_BIND：账号绑定； OPERATION_AUTH：代运营授权。 |
| `batch_no` | `BatchNo` | `*string` | query | 否 | 支付宝操作批次号。 batch_no提供基于操作批次的精确查询，只查询本操作批次的代运营操作结果，不反映最终代运营操作结果。 注意： 1.access_product_code、merchant_no、alipay_account三个参数组合与传入 batch_no 为两种查询方式，不可同时为空且以 batch_no 优先。 2.access_product_code、merchant_no、alipay_account三个参数组合可以查询最终代运营操作结果。 |
| `access_product_code` | `AccessProductCode` | `*string` | query | 否 | 接入的产品编号。未传batch_no时必填，枚举如下： OPENAPI_BIND_DEFAULT：操作类型为账号绑定； OPENAPI_AUTH_DEFAULT：操作类型为代运营授权。 |
| `merchant_no` | `MerchantNo` | `*string` | query | 否 | 支付宝商户号。支持间连、直连商户，值为2088开头。 未传batch_no时需注意： 若查询间连商户则 merchant_no 必填； 若查询直连商户，则merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。 |
| `alipay_account` | `AlipayAccount` | `*string` | query | 否 | 支付宝登录账号，通常为手机号或者邮箱。未传batch_no时需注意： 1.当查询merchant_no为间连商户时： * 若 operate_type=ACCOUNT_BIND，则alipay_account必填。 * 若 operate_type=OPERATION_AUTH，则alipay_account无需传入。 2.当查询merchant_no为直连商户时： * merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。 |
| `isv_scene_permissions` | `IsvScenePermissions` | `*string` | query | 否 | 场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2; |
