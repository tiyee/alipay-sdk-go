# 接口文档、调用方式与协议选型

核验日期：2026-11-05。本文决定接口的目标协议；当前代码是否实现、离线或沙箱是否验证，分别见[覆盖记录](api-coverage.md)。

## 1. 选型规则

**同一能力有官方 V3 接口时使用 V3；确认只有 V2 的接口才使用 V2。** 此规则同时适用于内置 Service 与用户自定义 API。

- 本文的 V2 指传统 AOP 网关协议，通常通过 `gateway.do`、`method`、`biz_content` 调用；其公共参数 `version=1.0` 不代表它是 V1 协议。
- 区分官方文档版本、SDK 支持方式与实际报文。V3 文档和 V3 SDK 同时包含 REST API、页面跳转与 App 订单串等场景，不能将 V3 支持范围等同于自动生成的 REST 清单。
- REST API 按官方 HTTP 方法与 `/v3/...` 路由调用；页面/唤端按该接口的 V3 文档与 SDK 示例构建签名产物。不能将方法名的点号机械替换成斜杠，虚构 REST 路由。
- Go 标准库 `encoding/json/v2` 的“v2”与支付宝 API 协议版本无关。
- 协议在 Operation 定义时显式固定。V3 调用失败、权限不足、超时或验签失败时，不自动回退到 V2，尤其不能跨协议重试支付、退款和转账。
- 文档暂不可读、SDK 未封装或目录中未检索到，都不能独立证明接口只有 V2。证据不足时标记“待核验”，补充官方资料后再实现。

## 2. 核验来源与范围

本轮覆盖本项目已实现或明确规划的 16 个常用接口：**13 个 V3 REST API，另有 3 个支付页面/唤端接口须按各自官方 V3 文档及通用执行入口接入**。后者不归类为“没有 V3 文档”或“只有 V2”。这是项目选型清单，不是支付宝全部接口的穷举。

主要依据为支付宝官方仓库的固定版本：

- [官方 V3 OpenAPI 描述](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/api/openapi.yaml)：`info.version=2026-09-23`，本次按 `operationId` 核对，包含 746 个操作。
- [同一版本的官方 V3 API 目录](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/README.md)：用于交叉确认方法与路由。
- [官方 V3 SDK 的 GenericExecuteApi](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/src/main/java/com/alipay/v3/util/GenericExecuteApi.java)：提供 `pageExecute` 和 `sdkExecute`，不在自动生成的 REST API 列表中。
- 下表的官方 V3 页面正文及其示例：用于确认接口字段、调用方式、响应产物与签名参数，而非只根据文档 URL 判断协议。

上述 SDK 提交时间为 2026-11-04，commit 为 `02177d34fce069f6c20d83c13f963150bb45467b`。未命中 REST 目录不代表 V3 SDK 不支持；必须同时检查通用执行入口与具体文档。

## 3. 有官方 V3 文档的页面与唤端接口

| 接口 | 用途 | 官方接入方式 | 官方 V3 文档 |
| --- | --- | --- | --- |
| `alipay.trade.page.pay` | 电脑网站支付 | V3 SDK `GenericExecuteApi.pageExecute`；POST 生成 HTML Form，GET 生成支付 URL | [统一收单下单并支付页面接口](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay) |
| `alipay.trade.wap.pay` | 手机网站支付 | V3 SDK `GenericExecuteApi.pageExecute`；POST 生成 HTML Form，GET 生成支付 URL | [手机网站支付接口](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay) |
| `alipay.trade.app.pay` | App 支付及支付并签约场景 | 按 V3 SDK `GenericExecuteApi.sdkExecute` 的订单字符串构建方式接入 | [App 支付](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay)、[商家扣款场景](https://opendocs.alipay.com/open-v3/0jf33m) |

`page.pay` 文档更新时间为 2026-09-13，`wap.pay` 为 2026-08-17。本次已读取其 V3 页面正文：二者都明确为页面跳转接口，Java 示例使用 `com.alipay.v3.util.GenericExecuteApi.pageExecute`，响应为 `pageRedirectionData`。POST 返回可在浏览器渲染的表单，GET 返回供打开或重定向的 URL。`app.pay` 的 V3 文档更新于 2026-08-21，明确为签名数据准备接口，通过 `sdkExecute` 生成 `orderStr`。

同日核验的 [0jf33m](https://opendocs.alipay.com/open-v3/0jf33m) 是 `alipay.trade.app.pay` 在商家扣款产品中的 V3 场景文档，更新时间同为 2026-08-21。正文明确用于 App 支付并签约，Java 示例使用 V3 SDK `GenericExecuteApi.sdkExecute`；新增的是同一接口的场景依据，因此接口总数仍为 16，不重复计数。

该场景包含 `agreement_sign_params` 以及嵌套的 `access_params`、`period_rule_params` 等签约数据。实现时采用独立具名 struct、具名枚举与明确的条件校验，不照搬普通 App 支付场景的必填关系，也不以动态字典替代模型。签约能力按场景排入实现范围，不将存在官方文档视为代码已经支持。

这些 V3 文档的页面产物仍展示 `gateway.do`、`method`、`version=1.0`、`biz_content` 和签名参数，应按文档实现；这描述实际签名格式，不能据此宣称没有 V3 文档或 V3 SDK 支持。SDK 应提供独立的页面/订单串构建入口，而非强行套用服务端 REST JSON 执行器。

付款后的交易查询、退款和退款查询采用第 4 节的 V3 REST。普通支付异步通知按自身回调契约验签，不能根据文档目录或交易 API 版本自动改变格式。

### 3.1 “仅 V2”状态的判定

本轮不再将上述三个支付接口列入“仅 V2”清单。只有取得明确的官方证据、确认没有相应 V3 文档和支持路径后，才将其他接口记录为仅 V2；不依据 REST 清单缺失作此判定。

## 4. 已有 V3、必须选用 V3 的接口

| 接口 | HTTP 方法 | V3 路由 | 官方 API 文档 |
| --- | --- | --- | --- |
| `alipay.trade.query` | POST | `/v3/alipay/trade/query` | [Trade.query](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#query) |
| `alipay.trade.precreate` | POST | `/v3/alipay/trade/precreate` | [Trade.precreate](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#precreate) |
| `alipay.trade.close` | POST | `/v3/alipay/trade/close` | [Trade.close](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#close) |
| `alipay.trade.refund` | POST | `/v3/alipay/trade/refund` | [Trade.refund](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#refund) |
| `alipay.trade.fastpay.refund.query` | POST | `/v3/alipay/trade/fastpay/refund/query` | [Refund.query](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeFastpayRefundApi.md#query) |
| `alipay.trade.pay` | POST | `/v3/alipay/trade/pay` | [Trade.pay](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#pay) |
| `alipay.trade.create` | POST | `/v3/alipay/trade/create` | [Trade.create](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#create) |
| `alipay.trade.cancel` | POST | `/v3/alipay/trade/cancel` | [Trade.cancel](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayTradeApi.md#cancel) |
| `alipay.data.dataservice.bill.downloadurl.query` | GET | `/v3/alipay/data/dataservice/bill/downloadurl/query` | [Bill.query](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayDataDataserviceBillDownloadurlApi.md#query) |
| `alipay.system.oauth.token` | POST | `/v3/alipay/system/oauth/token` | [Oauth.token](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipaySystemOauthApi.md#token) |
| `alipay.open.auth.token.app` | POST | `/v3/alipay/open/auth/token/app` | [AuthToken.app](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayOpenAuthTokenApi.md#app) |
| `alipay.fund.trans.uni.transfer` | POST | `/v3/alipay/fund/trans/uni/transfer` | [Fund.transfer](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayFundTransUniApi.md#transfer) |
| `alipay.fund.trans.common.query` | GET | `/v3/alipay/fund/trans/common/query` | [Fund.query](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayFundTransCommonApi.md#query) |

表中方法与路由直接来自同一份官方 OpenAPI，不从 V2 方法名推导。GET 使用 query 参数，POST 使用对应 V3 请求体；不得把 V2 的 `biz_content` 封装、`xxx_response` 节点或 `code=10000` 成功判定原样套入 V3。

## 5. 对架构与计划的影响

1. Client 共享配置、凭证装载、HTTP 生命周期、slog、类型化编解码与错误接口；V2/V3 分别实现请求签名、响应校验、证书定位和报文结构。
2. Operation 新增明确的协议枚举及 V3 HTTP 方法、相对路由与参数位置。协议类型、HTTP 方法、参数模式均采用具名类型和带类型常量。
3. 已存在的 5 个交易方法目前使用 V2；目标全部改为 V3。新增 V3 实现、契约测试与 Service 路由迁移完成前，覆盖表必须继续标为“当前 V2，V3 待迁移”，不能声称已支持 V3。
4. 下一开发里程碑先实现 V3 协议与上述交易迁移，再继续证书完善和测试版本交付；取消“先发布完整 V2，再考虑 V3”的安排。
5. 自定义账单地址查询示例改用 V3 GET。用户自定义接口也必须显式选择协议，不能使用通用扩展入口绕回已有 V3 接口的 V2 封装。

## 6. 持续维护

新增接口时先查官方 V3 OpenAPI 与接口正文，将方法名、协议、HTTP 方法、路由、证据链接和核验日期加入本清单或领域覆盖表。确认只有 V2 时记录具体依据；未确认时不默认归类为 V2。

官方增加对应 V3 操作后，将该接口移出 V2 保留清单，安排独立迁移与模型、签名、错误语义回归；不在运行时悄悄切换，不因迁移失败跨协议重试。旧调用的兼容安排通过明确版本和迁移说明管理。
