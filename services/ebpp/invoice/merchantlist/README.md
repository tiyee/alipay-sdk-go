# ebpp/invoice/merchantlist 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/merchantlist`。本目录提供 1 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `merchantlist.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_merchantlist_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`InvoiceMerchantListEnterApply`](#api-invoicemerchantlistenterapply) | [`alipay.ebpp.invoice.merchantlist.enter.apply`](https://opendocs.alipay.com/apis/05vxtl) | 商户批量入驻申请接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-invoicemerchantlistenterapply"></a>

### InvoiceMerchantListEnterApply — 商户批量入驻申请接口

- 官方接口：[`alipay.ebpp.invoice.merchantlist.enter.apply`](https://opendocs.alipay.com/apis/05vxtl)
- 路由：`POST /v3/alipay/ebpp/invoice/merchantlist/enter/apply`
- 请求类型：[InvoiceMerchantListEnterApplyRequest](merchantlist_request.go#L5)
- 响应类型：[InvoiceMerchantListEnterApplyResponse](merchantlist_response.go#L5)

针对已经在支付宝入驻的服务方，使用该接口导入服务方后面的开票商户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `merchant_base` | `MerchantBase` | [`*MerchantBaseEnterOpen`](#model-merchantbaseenteropen) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `sub_merchant_common_info` | `SubMerchantCommonInfo` | [`*SubMerchantCommonEnterOpen`](#model-submerchantcommonenteropen) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `sub_merchant_list` | `SubMerchantList` | [`*[]SubMerchantEnterOpen`](#model-submerchantenteropen) | JSON body | 未标注 | 商户门店列表信息，最多传入30个门店信息 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-merchantbaseenteropen"></a>

### MerchantBaseEnterOpen

类型定义：[MerchantBaseEnterOpen](merchantlist_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `login_id` | `LoginId` | `*string` | 未标注 | 登陆账号 |
| `logo_info` | `LogoInfo` | `*string` | 未标注 | 商户logo文件流信息，该字段为文件流经base64编码后得到的字符串。 |
| `m_name` | `MName` | `*string` | 未标注 | 商户品牌全称。 |
| `m_short_name` | `MShortName` | `*string` | 未标注 | 商户品牌简称 |

<a id="model-submerchantcommonenteropen"></a>

### SubMerchantCommonEnterOpen

类型定义：[SubMerchantCommonEnterOpen](merchantlist_model.go#L17)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `extend_fields` | `ExtendFields` | `*string` | 未标注 | 扩展字段，为json字符串格式；入驻支付即开票的产品该字段必填，用户白名单，由此字段传入；支付即开票商户配置维度，如果不传默认为SMID_MODE，目前仅支持SMID_MODE和PID_MODE |
| `product_code` | `ProductCode` | `*string` | 未标注 | 商户门店入驻的产品码,STANDARD_INVOICE:扫码开票，INVOICE_RETURN:企业自建-发票回传，INVOICE_EXPENSE:发票报销，PAYMENT_OPEN:支付即开票，PAY_FEE_OPEN:缴费后开票。 |
| `s_short_name` | `SShortName` | `*string` | 未标注 | 商户门店所属的服务方简称。 |

<a id="model-submerchantenteropen"></a>

### SubMerchantEnterOpen

类型定义：[SubMerchantEnterOpen](merchantlist_model.go#L27)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `extend_fields` | `ExtendFields` | `*string` | 未标注 | 扩展字段，为json字符串格式；入驻缴费后开票的产品如果需要传入自定义开票链接，由此字段传入。 |
| `industry` | `Industry` | `*string` | 未标注 | 商户行业，枚举值如下：subway：地铁，etc：ETC，online_car_hailing：网约车，catering：餐饮，hotel：酒店，convenience_store：便利店，oil：加油，insurance：保险，service_provider：运营商，life_payment：生活缴费，political：政务，other：其他 |
| `pid` | `Pid` | `*string` | 未标注 | 开票商户pid，入驻支付即开票场景的时候，非登录账号入驻模式，则该字段必传 |
| `register_no` | `RegisterNo` | `*string` | 未标注 | 商户门店税号。 |
| `sub_m_name` | `SubMName` | `*string` | 未标注 | 商户门店全称。 |
| `sub_m_short_name` | `SubMShortName` | `*string` | 未标注 | 商户门店简称 |
