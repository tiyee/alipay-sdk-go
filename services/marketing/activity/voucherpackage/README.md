# marketing/activity/voucherpackage 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/activity/voucherpackage`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `voucherpackage.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_voucherpackage_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`ActivityVoucherPackageBatchQuery`](#api-activityvoucherpackagebatchquery) | [`alipay.marketing.activity.voucherpackage.batchquery`](https://opendocs.alipay.com/apis/05velf) | 券包批量查询 | `POST` |
| [`ActivityVoucherPackageConsult`](#api-activityvoucherpackageconsult) | [`alipay.marketing.activity.voucherpackage.consult`](https://opendocs.alipay.com/apis/05vc1l) | 券包购买咨询 | `POST` |
| [`ActivityVoucherPackageQuery`](#api-activityvoucherpackagequery) | [`alipay.marketing.activity.voucherpackage.query`](https://opendocs.alipay.com/apis/05v9l4) | 券包详情查询 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-activityvoucherpackagebatchquery"></a>

### ActivityVoucherPackageBatchQuery — 券包批量查询

- 官方接口：[`alipay.marketing.activity.voucherpackage.batchquery`](https://opendocs.alipay.com/apis/05velf)
- 路由：`POST /v3/alipay/marketing/activity/voucherpackage/batchquery`
- 请求类型：[ActivityVoucherPackageBatchQueryRequest](voucherpackage_request.go#L5)
- 响应类型：[ActivityVoucherPackageBatchQueryResponse](voucherpackage_response.go#L5)

券包批量查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_num` | `PageNum` | `*int64` | JSON body | 未标注 | 分页参数，需要查询的页码，起始页是1 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 分页参数，每页记录数，最大不可超过30 |
| `voucher_package_purchase_start_time` | `VoucherPackagePurchaseStartTime` | `*string` | JSON body | 未标注 | 券包购买开始时间。 格式为：yyyy-MM-dd HH:mm:ss |
| `voucher_package_status` | `VoucherPackageStatus` | `*string` | JSON body | 未标注 | 券包状态 |

<a id="api-activityvoucherpackageconsult"></a>

### ActivityVoucherPackageConsult — 券包购买咨询

- 官方接口：[`alipay.marketing.activity.voucherpackage.consult`](https://opendocs.alipay.com/apis/05vc1l)
- 路由：`POST /v3/alipay/marketing/activity/voucherpackage/consult`
- 请求类型：[ActivityVoucherPackageConsultRequest](voucherpackage_request.go#L17)
- 响应类型：[ActivityVoucherPackageConsultResponse](voucherpackage_response.go#L17)

券包购买咨询，咨询当前用户是否可以购买指定的券包

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `voucher_package_id_list` | `VoucherPackageIdList` | `*[]string` | JSON body | 未标注 | 券包id列表 |

<a id="api-activityvoucherpackagequery"></a>

### ActivityVoucherPackageQuery — 券包详情查询

- 官方接口：[`alipay.marketing.activity.voucherpackage.query`](https://opendocs.alipay.com/apis/05v9l4)
- 路由：`GET /v3/alipay/marketing/activity/voucherpackage/query`
- 请求类型：[ActivityVoucherPackageQueryRequest](voucherpackage_request.go#L25)
- 响应类型：[ActivityVoucherPackageQueryResponse](voucherpackage_response.go#L23)

查询券包详情，包括券包信息和券包下面券信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `voucher_package_id` | `VoucherPackageId` | `*string` | query | 否 | 券包id |
