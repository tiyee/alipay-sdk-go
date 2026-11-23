# customs 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/customs`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `customs.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_customs_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CustomsDeclare`](#api-customsdeclare) | [`alipay.trade.customs.declare`](https://opendocs.alipay.com/pre-apis/05wyhb) | 统一收单报关接口 | `POST` |
| [`CustomsQuery`](#api-customsquery) | [`alipay.trade.customs.query`](https://opendocs.alipay.com/pre-apis/05wyjm) | 查询报关详细信息 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-customsdeclare"></a>

### CustomsDeclare — 统一收单报关接口

- 官方接口：[`alipay.trade.customs.declare`](https://opendocs.alipay.com/pre-apis/05wyhb)
- 路由：`POST /v3/alipay/trade/customs/declare`
- 请求类型：[CustomsDeclareRequest](customs_request.go#L13)
- 响应类型：[CustomsDeclareResponse](customs_response.go#L4)

境外商户的商品进口到国内时需要向海关中心报送支付、物流等信息，由海关审核，如果交易是由支付宝支付的话，需要商户可以调用此报关接口来请求支付宝将其指定交易报送支付单给海关。

请求约束：CustomsDeclareRequest 是 alipay.trade.customs.declare（统一收单报关）
的请求。out_request_no、amount、customs_place、merchant_customs_code、
merchant_customs_name 与交易号必填。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 报关流水号。商户生成的用于唯一标识一次报关操作的业务编号。 建议生成规则：yyyymmdd型8位日期拼接4位序列号。每个报关请求号仅允许传入：数字、英文字母、下划线”_”、短横线”－” 。长度6-64位前后不能有空格<br>SDK 字段说明：报关流水号：商户生成的唯一业务编号，6-64 位，<br>允许数字、字母、下划线与短横线。 |
| `trade_no` | `TradeNo` | `string` | JSON body | 未标注 | 支付宝交易号。该交易在支付宝系统中的交易流水号，最长64位。<br>SDK 字段说明：支付宝交易号（与 OutTradeNo 二选一）。 |
| `out_trade_no` | `OutTradeNo` | `string` | JSON body | 未标注 | 国际站外部订单号 |
| `amount` | `Amount` | `string` | JSON body | 未标注 | 报关金额，单位为人民币“元”，精确到小数点后2位。<br>SDK 字段说明：报关金额，单位为人民币元，精确到小数点后 2 位。 |
| `customs_place` | `CustomsPlace` | `string` | JSON body | 未标注 | 海关编号（大小写皆可）。参见“ [海关编号](https://opendocs.alipay.com/pre-open/01x3kh#%E6%B5%B7%E5%85%B3%E7%BC%96%E5%8F%B7)”。<br>SDK 字段说明：海关编号，取值见官方海关编号列表。 |
| `merchant_customs_code` | `MerchantCustomsCode` | `string` | JSON body | 未标注 | 商户在海关备案的编号。 |
| `merchant_customs_name` | `MerchantCustomsName` | `string` | JSON body | 未标注 | 商户海关备案名称。 |
| `is_split` | `IsSplit` | `*string` | JSON body | 未标注 | 商户控制本单是否拆单的报关参数。<br> 仅当该参数传值为T或者t时，才会触发拆单。<br>SDK 字段说明：控制本单是否拆单；仅传 T 或 t 时触发拆单。 |
| `sub_out_biz_no` | `SubOutBizNo` | `*string` | JSON body | 未标注 | 拆单报关的商户子订单号。 用于区别拆单时不同子单。拆单时必须传入，否则会报INVALID_PARAMETER错误码。<br>SDK 字段说明：拆单报关的商户子订单号；拆单时必须传入。 |
| `buyer_info` | `BuyerInfo` | [`*CustomsDeclareBuyerInfo`](#model-customsdeclarebuyerinfo) | JSON body | 未标注 | 订购人信息。 |
| `declare_mode` | `DeclareMode` | `*int64` | JSON body | 未标注 | 报关模式，默认可空，1表示需要强校验买家和支付人的身份信息。<br>SDK 字段说明：报关模式：1 表示需要强校验买家和支付人的身份信息。 |
| `deposit_trade_no` | `DepositTradeNo` | `*string` | JSON body | 未标注 | 定金在支付宝系统中的交易流水号，最长64位。当该参数不为空时，代表为”定金+尾款“报关场景。<br>SDK 字段说明：定金交易流水号；“定金+尾款”报关场景传入。 |
| `merchant_id` | `MerchantID` | `*string` | JSON body | 未标注 | 国际站商户标识 |

注意事项：

- 报关为异步流程：受理成功不代表报关完成，以 CustomsQuery 查询为准。

<a id="api-customsquery"></a>

### CustomsQuery — 查询报关详细信息

- 官方接口：[`alipay.trade.customs.query`](https://opendocs.alipay.com/pre-apis/05wyjm)
- 路由：`GET /v3/alipay/trade/customs/query`
- 请求类型：[CustomsQueryRequest](customs_request.go#L45)
- 响应类型：[CustomsQueryResponse](customs_response.go#L27)

境外商户的商品进口到国内时需要向海关中心报送支付、物流等信息，由海关审核，如果交易是由支付宝支付的话，需要商户调用支付宝报关接口来请求支付宝将其指定交易报送支付单给海关。商户可以通过该接口主动查询报关状态，了解报关详情以及进度。

请求约束：CustomsQueryRequest 是 alipay.trade.customs.query（报关查询，V3 GET）
的请求：报关请求号逗号分隔，单次最多 10 个。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_nos` | `OutRequestNos` | `string` | query | 否 | 报关请求号。需要查询的商户端报关请求号，支持批量查询，<br> 多个值用英文半角逗号分隔，单次请求最多10个;<br>SDK 字段说明：需要查询的商户端报关请求号，多个值用英文半角<br>逗号分隔，单次最多 10 个。 |

注意事项：

- 单次最多查询 10 个报关请求号。

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-customsdeclarebuyerinfo"></a>

### CustomsDeclareBuyerInfo

类型定义：[CustomsDeclareBuyerInfo](customs_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `buyer_name` | `BuyerName` | `*string` | 未标注 | 订购人姓名。即订购人留在商户处的姓名信息 |
| `buyer_cert_no` | `BuyerCertNo` | `*string` | 未标注 | 订购人身份证号。即订购人留在商户处的身份证信息 |
