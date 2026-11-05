# ebpp/bill 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/bill`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `bill.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_bill_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`BillAdd`](#api-billadd) | [`alipay.ebpp.bill.add`](https://opendocs.alipay.com/open-v3/b9117832_alipay.ebpp.bill.add) | 创建账单 | `POST` |
| [`BillGet`](#api-billget) | [`alipay.ebpp.bill.get`](https://opendocs.alipay.com/open-v3/b6ddabc9_alipay.ebpp.bill.get) | 查询账单 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-billadd"></a>

### BillAdd — 创建账单

- 官方接口：[`alipay.ebpp.bill.add`](https://opendocs.alipay.com/open-v3/b9117832_alipay.ebpp.bill.add)
- 路由：`POST /v3/alipay/ebpp/bill/add`
- 请求类型：[BillAddRequest](bill_request.go#L5)
- 响应类型：[BillAddResponse](bill_response.go#L5)

创建账单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `bank_bill_no` | `BankBillNo` | `*string` | JSON body | 未标注 | 外部订单号 |
| `bill_date` | `BillDate` | `*string` | JSON body | 未标注 | 账单的账期，格式为 yyyyMMdd。例如：202012表示2020年12月的账单。 |
| `bill_key` | `BillKey` | `*string` | JSON body | 未标注 | 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。 |
| `charge_inst` | `ChargeInst` | `*string` | JSON body | 未标注 | 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。 |
| `extend_field` | `ExtendField` | `*string` | JSON body | 未标注 | 扩展属性 |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | JSON body | 未标注 | 输出机构的业务流水号，需要保证唯一性 |
| `mobile` | `Mobile` | `*string` | JSON body | 未标注 | 用户的手机号 |
| `order_type` | `OrderType` | `*string` | JSON body | 未标注 | 支付宝订单类型。枚举支持： *JF：公共事业缴纳。 *WUYE：物业缴费。 *HK：信用卡还款。 *TX：通讯缴费。 |
| `owner_name` | `OwnerName` | `*string` | JSON body | 未标注 | 拥有该账单的用户姓名 |
| `pay_amount` | `PayAmount` | `*string` | JSON body | 未标注 | 缴费金额。用户支付的总金额。单位为 元（人民币）。取值范围为[0.01，100000000.00]，精确到小数点后两位。 |
| `service_amount` | `ServiceAmount` | `*string` | JSON body | 未标注 | 账单的服务费。 |
| `sub_order_type` | `SubOrderType` | `*string` | JSON body | 未标注 | 子业务类型是业务类型的下一级概念。枚举支持： *WATER：缴水费。 *ELECTRIC：缴电费。 *GAS：缴燃气费。 *COMMUN：缴固话宽带费。 *CATV：缴有线电视费。 *WUYE：缴物业费。 *RC：定期还车贷代扣。 *RH：定期还房贷代扣。 *RR：定期还房租代扣。 *RN：定期还网贷代扣。 *CZ：手机充值代扣。 例如：WATER表示JF下面的水费。 |
| `traffic_location` | `TrafficLocation` | `*string` | JSON body | 未标注 | 交通违章地点，sub_order_type=TRAFFIC时填写。 |
| `traffic_regulations` | `TrafficRegulations` | `*string` | JSON body | 未标注 | 违章行为，sub_order_type=TRAFFIC时填写。 |

<a id="api-billget"></a>

### BillGet — 查询账单

- 官方接口：[`alipay.ebpp.bill.get`](https://opendocs.alipay.com/open-v3/b6ddabc9_alipay.ebpp.bill.get)
- 路由：`GET /v3/alipay/ebpp/bill/get`
- 请求类型：[BillGetRequest](bill_request.go#L39)
- 响应类型：[BillGetResponse](bill_response.go#L37)

查询账单

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `order_type` | `OrderType` | `*string` | query | 否 | 支付宝订单类型。公共事业缴纳JF,信用卡还款HK |
| `merchant_order_no` | `MerchantOrderNo` | `*string` | query | 否 | 输出机构的业务流水号，需要保证唯一性。 |
