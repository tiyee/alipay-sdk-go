# isv 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/sp/isv`。本包提供 3 个支付宝 V3 REST 操作。

通过 `isv.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SpISVBillSync`](#api-spisvbillsync) | `alipay.open.sp.isv.bill.sync` |
| [`SpISVSignAuthCreate`](#api-spisvsignauthcreate) | `alipay.open.sp.isv.signauth.create` |
| [`SpISVSignAuthQuery`](#api-spisvsignauthquery) | `alipay.open.sp.isv.signauth.query` |

## 使用示例

- [SpISVBillSync](api_isv_spisvbillsync_example_test.go)
- [SpISVSignAuthCreate](api_isv_spisvsignauthcreate_example_test.go)
- [SpISVSignAuthQuery](api_isv_spisvsignauthquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-spisvbillsync"></a>

### SpISVBillSync — 服务商账单数据回传

- 官方接口：[`alipay.open.sp.isv.bill.sync`](https://opendocs.alipay.com/apis/05vidn)
- 路由：`POST /v3/alipay/open/sp/isv/bill/sync`
- 请求类型：[SpISVBillSyncRequest](isv_request.go#L4)
- 响应类型：[SpISVBillSyncResponse](isv_response.go#L4)

针对SSB业务中，S1回传S2推广过程中，获得的分佣账单数据

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `end_time` | `EndTime` | `*string` | JSON body | 未标注 | 账单结束时间 |
| `generate_time` | `GenerateTime` | `*string` | JSON body | 未标注 | 出账时间，回传出账信息时，必须传递 |
| `item_code` | `ItemCode` | `*string` | JSON body | 未标注 | 服务优选商品编码 |
| `out_biz_no` | `OutBizNo` | `*string` | JSON body | 未标注 | 外部业务号，传isv系统生成的账单号，需要保证唯一 |
| `pay_time` | `PayTime` | `*string` | JSON body | 未标注 | 账单支付时间，回传支付数时，必须传递 |
| `price` | `Price` | `*string` | JSON body | 未标注 | 账单金额，这里填写的是整数，单位为分，比如1元，那么输入100 |
| `promotor_pid` | `PromotorPid` | `*string` | JSON body | 未标注 | 推广服务商(S2)pid |
| `start_time` | `StartTime` | `*string` | JSON body | 未标注 | 账单开始时间 |
| `status` | `Status` | `*string` | JSON body | 未标注 | GENERATE_BILL代表”已出账“，PAID_BILL代表”已支付“，选择这两者之一 |
| `sub_promotor_pid` | `SubPromotorPid` | `*string` | JSON body | 未标注 | S1回传的S2的操作员工子账号，如果是子账号操作，请回传 |

<a id="api-spisvsignauthcreate"></a>

### SpISVSignAuthCreate — 创建签约授权一体化订单

- 官方接口：[`alipay.open.sp.isv.signauth.create`](https://opendocs.alipay.com/open-v3/e4f4b975_alipay.open.sp.isv.signauth.create)
- 路由：`POST /v3/alipay/open/sp/isv/signauth/create`
- 请求类型：[SpISVSignAuthCreateRequest](isv_request.go#L28)
- 响应类型：[SpISVSignAuthCreateResponse](isv_response.go#L10)

针对签约授权一体化场景，创建签约授权一体化订单绑定签约单（可选）、绑定代运营授权单、应用授权单，跟踪整个签约授权一体化订单的状态

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `isv_auth_scene_infos` | `IsvAuthSceneInfos` | [`*[]ISVAuthSceneInfo`](#model-isvauthsceneinfo) | JSON body | 未标注 | 代运营授权场景信息 |
| `merchant_logon_id` | `MerchantLogonId` | `*string` | JSON body | 未标注 | 商户登录账号 支持手机号和邮箱账号；不支持pid |
| `need_app_auth` | `NeedAppAuth` | `*string` | JSON body | 未标注 | 1表示需要，0表示不需要 不传参数默认是需要（1） 是否创建三方应用授权子任务，也就是是否需要给三方应用授权 |
| `sign_order_no` | `SignOrderNo` | `*string` | JSON body | 未标注 | 签约单号 可通过alipay.open.agent.confirm接口获取签约单号 |

<a id="api-spisvsignauthquery"></a>

### SpISVSignAuthQuery — 查询签约授权一体化订单状态详情

- 官方接口：[`alipay.open.sp.isv.signauth.query`](https://opendocs.alipay.com/open-v3/a322ed32_alipay.open.sp.isv.signauth.query)
- 路由：`GET /v3/alipay/open/sp/isv/signauth/query`
- 请求类型：[SpISVSignAuthQueryRequest](isv_request.go#L40)
- 响应类型：[SpISVSignAuthQueryResponse](isv_response.go#L16)

查询签约授权一体化订单状态详情（包含子任务信息）

请求约束：SpISVSignAuthQueryRequest is the request type for alipay.open.sp.isv.signauth.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderId` | `*string` | query | 否 | 签约授权一体化订单ID 可通过alipay.open.sp.isv.signauth.create 接口获取 |

## 嵌套请求参数

<a id="model-isvauthsceneinfo"></a>

### ISVAuthSceneInfo

类型定义：[ISVAuthSceneInfo](isv_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `scene_code` | `SceneCode` | `*string` | 未标注 | 运营场景编码 OPERATION_POINTS：管理运营积分 SHOP_MANAGE：管理门店信息 MINI_APP_OPER：运营支付宝小程序 PROMOTION_MANAGE：运营营销活动 |
| `scene_permissions` | `ScenePermissions` | `*string` | 未标注 | 运营场景下的权限编码，多个权限编码以,隔开 1、管理门店信息：SHOP_MANAGE；基础权限（升级）：SHOP_MANAGE_BASE 2、运营营销活动：PROMOTION_MANAGE ；基础权限（升级）：PROMOTION_MANAGE_BASE 3、运营支付宝小程序：MINI_APP_OPER；基础权限（升级）：MINI_APP_OPER_BASE 4、管理运营积分：OPERATION_POINTS；基础权限（升级）：OPERATION_POINTS_BASE |
