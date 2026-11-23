# 接口版本说明

核验日期：2026-11-05。本文说明 SDK 的协议支持范围与官方契约清点结论；当前代码是否实现、离线或沙箱是否验证，分别见[覆盖记录](api-coverage.md)。

## 1. 支持范围：仅 V3

**本 SDK 仅支持支付宝 V3**：服务端接口使用 V3 REST（`/v3/...` 路由）；页面/唤端（page.pay、wap.pay、app.pay）按官方 V3 文档与产物构建规范提供本地签名产物；普通支付异步通知按通知自身契约处理。各调用方式均按其明确契约执行。

- 页面/唤端产物按官方 V3 文档使用 `gateway.do`、`method`、`version=1.0`、`biz_content` 等签名字段；这些字段仅用于本地签名产物构建。
- 区分官方文档版本、SDK 支持方式与实际报文。V3 文档和 V3 SDK 同时包含 REST API、页面跳转与 App 订单串等场景，不能将 V3 支持范围等同于自动生成的 REST 清单。
- REST API 按官方 HTTP 方法与 `/v3/...` 路由调用，不能将方法名的点号机械替换成斜杠，虚构 REST 路由。
- Go 标准库 `encoding/json/v2` 的"v2"与支付宝 API 协议版本无关。
- 未确认 V3 契约的接口不予实现；未实现的能力明确报为未支持。

2026-11-23 收紧认证失败处理。V3 带签名响应要求完整的时间戳、随机串及证书模式 SN，普通支付通知要求显式 `sign_type=RSA2`。内置证书凭证在每次签名和验签时检查全部配置材料的有效期，错误类型与处理方式见[协议说明](protocol-notes.md)。

## 2. 内置接口契约清点结论

2026-11-05 依照官方 OpenAPI 快照核对并补齐 `docs/apis.md` 中的 REST 操作。快照共列 746 个 REST Operation，其中 678 个属于 `alipay.*`；其余 68 个使用 `ant.*`、`zhima.*`、`zoloz.*` 等其他命名空间，现均已提供 Service 方法及请求、响应模型。

| 领域 | 包布局 | 数量 |
| --- | --- | --- |
| `alipay.boss.*` | `services/boss/<业务子域>/` | 8 |
| `alipay.commerce.*` | `services/commerce/<业务子域>/` | 55 |
| `alipay.data.*` | `services/data/<业务子域>/` | 18 |
| `alipay.ebpp.*` | `services/ebpp/<业务子域>/` | 77 |
| `alipay.eco.*` | `services/eco/<业务子域>/` | 32 |
| `alipay.fund.*` | `services/fund/<业务子域>/` | 42 |
| `alipay.iservice.*` | `services/iservice/ccm/<业务子域>/` | 45 |
| `alipay.marketing.*` | `services/marketing/<业务子域>/` | 64 |
| `alipay.merchant.*` | `services/merchant/<业务子域>/` | 18 |
| `alipay.mobile.*` | `services/mobile/<业务子域>/` | 14 |
| `alipay.offline.*` | `services/offline/<业务子域>/` | 8 |
| `alipay.open.*` | `services/open/<业务子域>/` | 241 |
| `alipay.pcredit.*` | `services/pcredit/` | 6 |
| `alipay.security.*` | `services/security/` | 3 |
| `alipay.social.*` | `services/social/` | 1 |
| `alipay.system.*` | `services/system/` | 1 |
| `alipay.trade.*` | `services/trade/` 与交易业务子包 | 27 |
| `alipay.user.*` | `services/user/<业务子域>/` | 18 |
| **合计 `alipay.*`** | | **678** |

原有 40 项的接口名、HTTP 方法和路由均在快照中找到。其余 638 项依照快照生成 Service 方法和具名模型。开放平台、发票、营销、企业码及客服等密集业务按接口前缀拆入子包，例如 `alipay.open.mini.inner.*` 位于 `services/open/mini/inner/`；请求、响应和共享模型文件与接口实现放在同一子包，不集中放在领域根目录。生成接口尚未逐项人工核验业务条件、证书要求和产品权限；这不代表商户已开通相应产品或接口可直接用于生产。完整的包路径映射见[业务包目录](service-layout.md)。逐接口的方法、路径、实现与验证状态见[接口覆盖记录](api-coverage.md)和[接口目录](apis.md)。

支付与创建交易的 `royalty_info.royalty_detail_infos[].amount_percentage` 按字符串编码，退款分账明细的同名字段按整数编码。订单咨询响应的顶层 `wait_repayment_amount` 是字符串，待还订单明细内的同名字段是整数。参见[支付](https://opendocs.alipay.com/open-v3/08c7f9f8_alipay.trade.pay)、[创建交易](https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create)和[订单咨询](https://opendocs.alipay.com/open-v3/11cb1f6e_alipay.trade.advance.consult)。

清点边界：

- 页面/唤端产物（page.pay / wap.pay / app.pay）与异步通知不在 openapi.yaml 覆盖范围（快照中无对应 path）。其依据是官方[页面构建规范](https://opendocs.alipay.com/common/0i2kxr.md)、[唤端说明](https://opendocs.alipay.com/common/0i2jez.md)、各接口 V3 文档与[通知验签文档](https://opendocs.alipay.com/common/02mse7.md)，无法以 yaml 核对，正确性依赖上述官方规范。
- openapi.yaml 快照有版本日期（2026-09-23，commit `02177d34fce069f6c20d83c13f963150bb45467b`）；官方更新契约后需以新版快照复核本清单。

### 参数位置、证书与交易通知

按 2026-09-23 版 OpenAPI 快照区分路由、query 和请求体。POST 和 DELETE 也可只传 query；同时声明 query 和 JSON 的接口将两者分别写入 URI 和正文。已逐项核对 746 个操作的 HTTP 方法、路由和请求参数模式。55 个声明 query 或省略请求体的非 GET 接口另有离线 HTTP 测试，逐项比对 URI、实际发送的正文字节和 RSA2 签名。详见[离线契约测试](../examples/generatedapi/query_contract_test.go)。

`alipay.trade.pay`、`alipay.trade.create`、`alipay.trade.precreate`、`alipay.trade.close` 和 `alipay.trade.refund.apply` 均支持在 JSON 请求体中设置 `notify_url`。`alipay.trade.query` 的结算、优惠券、单品优惠与医保字段由相应查询选项控制。`alipay.fund.trans.uni.transfer` 在发送前强制校验证书模式，依据[资金支出接口说明](https://opendocs.alipay.com/common/02kg66.md)。

2026-11-06 复核 V3 字段契约后，交易模型共用 `settle_info.settle_period_time`，预创建在顶层发送 `payment_type`。资金退回按接口文档校验 `biz_scene`，历史余额查询仅使用日期或月份。31 个操作的数值和布尔 query 参数已使用对应 Go 类型，必填参数在请求发送前检查。交易、退款、转账和账单模型已补齐快照字段，两个废弃令牌响应字段的取舍见[逐接口核查记录](checked.md)。

App 支付并签约使用 `biz_content.agreement_sign_params`。页面签名产物按 PC/WAP 参数表分别处理二维码宽度和退出地址。[App 签约](https://opendocs.alipay.com/open-v3/0jf33m) · [电脑网站支付](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay) · [手机网站支付](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay)

### 2.1 名为 query 实为 POST + JSON body 的接口

以下 4 个接口名称含 query/batchquery，但官方 V3 契约为 **POST + JSON 请求体**（已逐条与 openapi.yaml 核对，本 SDK 与官方一致）。后续维护不得按接口名想当然改为 GET + query：

| 接口 | 实际方法与路径 |
| --- | --- |
| `alipay.trade.fastpay.refund.query` | POST `/v3/alipay/trade/fastpay/refund/query` |
| `alipay.trade.settle.receivables.query` | POST `/v3/alipay/trade/settle/receivables/query` |
| `alipay.trade.batch.settle.query` | POST `/v3/alipay/trade/batch/settle/query` |
| `alipay.trade.royalty.relation.batchquery` | POST `/v3/alipay/trade/royalty/relation/batchquery` |

### 2.2 快照外接口补录

2026-11-05 在[接口目录](apis.md#快照外接口35-项v3-契约未确认)补录本次官方资料核查发现的 35 个接口，不增加 SDK 实现，也不计入上述 746 个 V3 REST 操作或 3 个已核验的页面/唤端接口：

- 6 个 OpenID 转换接口：已查阅官方接口全文，文档使用 `gateway.do`、`method`、`biz_content` 等传统网关调用方式；尚未确认 V3 契约，当前 SDK 不提供相应服务端调用。
- 29 个官方[资金操作类敏感接口列表](https://opendocs.alipay.com/common/02kkuw.md)列出的快照外接口：包含 16 个 `alipay.*`、10 个 `koubei.*` 和 3 个 `mybank.*` 接口。名称及原始文档链接有清单依据，但未取得可核验的接口正文；现行可用性、接口版本与 V3 契约均待确认。

这些补录记录用于跟踪目录缺项，不能据此推导 REST 路由、使用通用执行入口绕过协议限制，或宣称 SDK 已支持。后续取得明确官方 V3 契约后再更新分类、实现与覆盖记录。

## 3. 核验来源

- 支付宝的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)提供 OpenAPI 描述文件。本次清点采用 `info.version=2026-09-23`，746 个操作的接口名、方法、路由和文档链接见[接口目录](apis.md)。
- `alipay.fund.auth.order.app.freeze` 的[V3 接口文档](https://opendocs.alipay.com/open-v3/46bf87fe_alipay.fund.auth.order.app.freeze)演示通过 `sdkExecute` 生成 App 订单串。目录中的 POST 路由来自 OpenAPI 快照，REST 调用方式仍待核实。
- 各接口官方 V3 文档正文与[页面跳转](https://opendocs.alipay.com/common/0i2kxr.md)、[唤端](https://opendocs.alipay.com/common/0i2jez.md)、[通知验签](https://opendocs.alipay.com/common/02mse7.md)说明。

快照提交 `02177d34fce069f6c20d83c13f963150bb45467b`（2026-11-04 提交时间）。未命中 REST 目录不代表官方不支持该能力；新增接口时必须同时检查通用执行入口与具体文档。

## 4. 对架构与代码的影响

1. `core.Client` 是唯一 REST 客户端，执行 V3 REST 协议。
2. Operation 通过 `core.NewJSONOperation` / `core.NewOperation` 构造，显式声明官方方法名、HTTP 方法、相对路由与参数模式（JSON body / query / multipart）；认证要求经 `AuthRequirement` 在执行前校验。
3. 页面/唤端产物经 `core.NewArtifactBuilder` 与 `trade.PaymentProducts` 按官方产物契约构建，使用独立的产物网关配置（`gateway.do`）。异步通知按自身契约在 `core/notify` 处理，与 REST 验签边界独立。
4. 自定义接口同样只能声明 V3 规格；`examples/customapi` 演示外部包经 `NewJSONOperation` / `NewOperation` 接入 V3 GET/POST 的方式。

## 5. 持续维护

新增接口时先查官方 V3 OpenAPI 与接口正文，将方法名、HTTP 方法、路由、参数位置、证据链接和核验日期加入[覆盖记录](api-coverage.md)。快照未命中的能力先核实通用执行入口与官方文档，未确认前不实现、不宣称支持。官方快照更新后，按新版快照复核内置清单并在本文件记录复核日期。
