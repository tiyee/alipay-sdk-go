# 接口版本选型与 V2 保留清单

核验日期：2026-11-05。本文决定接口的目标协议；当前代码是否实现、离线或沙箱是否验证，分别见[覆盖记录](api-coverage.md)。

## 1. 选型规则

**同一能力有官方 V3 接口时使用 V3；确认只有 V2 的接口才使用 V2。** 此规则同时适用于内置 Service 与用户自定义 API。

- 本文的 V2 指传统 AOP 网关协议，通常通过 `gateway.do`、`method`、`biz_content` 调用；其公共参数 `version=1.0` 不代表它是 V1 协议。
- V3 指官方 OpenAPI 描述中的实际 HTTP 方法与 `/v3/...` 路由。不能只看文档页面是否位于 `open-v3` 路径，也不能将路径中的点号机械替换为斜杠来猜路由。
- Go 标准库 `encoding/json/v2` 的“v2”与支付宝 API 协议版本无关。
- 协议在 Operation 定义时显式固定。V3 调用失败、权限不足、超时或验签失败时，不自动回退到 V2，尤其不能跨协议重试支付、退款和转账。
- 文档暂不可读、SDK 未封装或目录中未检索到，都不能独立证明接口只有 V2。证据不足时标记“待核验”，补充官方资料后再实现。

## 2. 核验来源与范围

本轮覆盖本项目已实现或明确规划的 16 个常用接口：**13 个选用 V3，3 个保留 V2**。这是项目选型清单，不是支付宝全部接口的穷举。

主要依据为支付宝官方仓库的固定版本：

- [官方 V3 OpenAPI 描述](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/api/openapi.yaml)：`info.version=2026-09-23`，本次按 `operationId` 核对，包含 746 个操作。
- [同一版本的官方 V3 API 目录](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/README.md)：用于交叉确认方法与路由。
- 同一官方仓库的 V2 Request 源码，以及支付宝[页面跳转规范](https://opendocs.alipay.com/common/0i2kxr.md)、[唤端规范](https://opendocs.alipay.com/common/0i2jez.md)：确认 V2 页面/订单串场景。

上述提交时间为 2026-11-04，commit 为 `02177d34fce069f6c20d83c13f963150bb45467b`。未命中 V3 目录只是辅助证据；下面保留 V2 的三个接口还具有 V2 官方实现与对应签名产物规范。

## 3. 保留 V2 的接口

| 接口 | 用途 | 采用方式 | 官方依据 |
| --- | --- | --- | --- |
| `alipay.trade.page.pay` | 电脑网站支付 | V2 签名 URL / HTML Form；不作为服务端 V3 REST 请求发送 | [V2 Request](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v2/src/main/java/com/alipay/api/request/AlipayTradePagePayRequest.java)、[页面跳转规范](https://opendocs.alipay.com/common/0i2kxr.md) |
| `alipay.trade.wap.pay` | 手机网站支付 | V2 签名 URL / HTML Form | [V2 Request](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v2/src/main/java/com/alipay/api/request/AlipayTradeWapPayRequest.java)、[接口说明](https://opendocs.alipay.com/open/02ivbs.md?scene=21) |
| `alipay.trade.app.pay` | App 支付 | V2 签名订单字符串，交给 App 支付 SDK 唤起收银台 | [V2 Request](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v2/src/main/java/com/alipay/api/request/AlipayTradeAppPayRequest.java)、[服务端示例](https://opendocs.alipay.com/common/02kiq5.md) |

这三个接口在本轮官方 V3 描述中没有对应 REST 操作。页面/唤端本身生成签名产物，不通过“新增一个 `/v3/...` 路由”替代。付款后的交易查询、退款、退款查询等仍应采用其已有的 V3 接口，不能因为下单产物是 V2 就将整条交易链固定为 V2。

普通支付异步通知是独立的回调报文及验签契约，不计入上述三个 V2 调用接口，也不因服务端交易 API 迁移到 V3 就自动改变其通知格式。

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
