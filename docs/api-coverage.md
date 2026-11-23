# 接口覆盖记录

状态：随实现更新。日期：2026-11-05。SDK 仅支持支付宝 V3；`docs/apis.md` 中 746 个 REST 接口均已有 Service 方法，包括 678 个 `alipay.*` 和 68 个其他命名空间操作。原有 40 个 REST 接口完成离线验证；其余 `alipay.*` 生成接口尚未逐项完成功能测试或沙箱验收。新增的其他命名空间接口均有操作规格测试，尚未逐项核验业务条件及沙箱行为。授权模式、产品权限和接口级条件仍须按具体官方文档确认。沙箱验收步骤见[沙箱验收指引](sandbox-verification.md)，协议依据见[协议说明](protocol-notes.md)。

`monitor.heartbeat.syn` 位于 `services/monitor`，依据[官方 V3 接口文档](https://opendocs.alipay.com/apis/05vc16)实现；已有离线 Service 测试，尚未进行沙箱验证。

其余 67 项基于相同快照生成，分别位于 `services/alios/open`（1）、`services/ant/merchant`（25）、`services/datadigital/fincloud`（10）、`services/grandsecurity/bizrisks`（1）、`services/zhima/credit`（13）、`services/zhima/customer`（6）、`services/zhima/merchant`（6）及 `services/zoloz/authentication`（5）。各操作的官方文档链接、方法及路径见[接口目录](apis.md)和方法 GoDoc。`grandsecurity.bizrisks.fact.check` 快照声明 JSON 响应但字段说明提及流式数据，当前只支持快照声明的普通 JSON 响应；流式返回与验签规则尚未确认。

本文覆盖表区分"已实现 / 离线验证 / 沙箱验证"。原有 40 个 REST 接口与页面产物、通知能力已离线验证（httptest 端到端 + 独立签名复算 + mock 规格测试），沙箱未验证。生成接口按官方 OpenAPI 快照提供请求/响应模型与共享执行入口；仓库测试确认全部包可编译，但尚未逐项验证业务规则、商户权限、证书要求或沙箱行为。

契约清点结论：原有 40 个内置 REST Operation 已逐条与官方 openapi.yaml 快照核对，40 项的方法名、HTTP 方法和路由均与快照匹配。其余 638 个生成操作采用同一快照的操作名、HTTP 方法、路由、参数位置和成功响应 schema；全部 678 个 `alipay.*` 操作均有对应 Service 方法。生成接口的业务规则与权限限制仍须查阅各方法 GoDoc 链接指向的具体文档。清点边界与"名为 query 实为 POST+body"的特殊接口见[接口版本说明](api-versions.md)第 2 节。

## 1. 交易 `services/trade/`（含业务子包，共 27 项）

包与方法的对应关系见[交易目录](../services/trade/README.md)。`order`、`refund`、`batch`、`royalty`、`settle`、`customs`、`advance`、`overseas`、`orderinfo` 各自提供 `NewService(executor)`，共享类型保留在 `trade` 根包。

以下交易接口实现统一位于 `api_trade.go`，请求与响应分别位于 `trade_request.go`、`trade_response.go`，共享模型位于 `trade_model.go`。

### 查询、预下单与退款（5 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.trade.query | POST `/v3/alipay/trade/query` | [接口文档](https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query) | 查询成功≠支付成功，以 TradeStatus 为准；可通过 query_options 指定额外返回字段 |
| alipay.trade.precreate | POST `/v3/alipay/trade/precreate` | [接口文档](https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate) | 二维码有效期固定 2 小时；同一订单号不能重复申请；subject 禁用 /=& |
| alipay.trade.refund | POST `/v3/alipay/trade/refund` | [接口文档](https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund?scene=common&pathHash=00ab12b7) | 请求成功≠退款成功，以 fund_change 为准；重试须保持 out_request_no 不变；SDK 不自动重试 |
| alipay.trade.fastpay.refund.query | POST `/v3/alipay/trade/fastpay/refund/query` | [接口文档](https://opendocs.alipay.com/open-v3/b9ef37bd_alipay.trade.fastpay.refund.query?scene=common&pathHash=f071344e) | 名为 query 实为 POST+JSON body；请求成功仅代表查询成功；refund_status 未返回表示退款未收到或失败；官方建议间隔 10s |
| alipay.trade.close | POST `/v3/alipay/trade/close` | [接口文档](https://opendocs.alipay.com/open-v3/48ea518b_alipay.trade.close?scene=common&pathHash=2c6cdbea) | 仅 WAIT_BUYER_PAY 状态可关单；SDK 不自动关单 |

### 支付、创建与撤销（5 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.trade.pay | POST `/v3/alipay/trade/pay` | [接口文档](https://opendocs.alipay.com/open-v3/08c7f9f8_alipay.trade.pay) | 当面付需 scene+auth_code；异步支付以通知/查询确认终态；写操作不自动重试 |
| alipay.trade.create | POST `/v3/alipay/trade/create` | [接口文档](https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create) | 返回交易号仅表示订单创建成功；含支付并签约参数（SignParams） |
| alipay.trade.cancel | POST `/v3/alipay/trade/cancel` | [接口文档](https://opendocs.alipay.com/open-v3/cd7d54d2_alipay.trade.cancel?scene=common&pathHash=cb83add6) | retry_flag=Y 需调用方重试；SDK 不自动重试 |
| alipay.trade.order.pay | POST `/v3/alipay/trade/order/pay` | [接口文档](https://opendocs.alipay.com/open-v3/c2751e3b_alipay.trade.order.pay) | out_request_no 幂等；异步支付以通知/查询确认 |
| alipay.trade.refund.apply | POST `/v3/alipay/trade/refund/apply` | [接口文档](https://opendocs.alipay.com/apis/05wbvf) | REFUND_PROCESSING 非终态，需查询确认；SDK 不自动重试 |

### 信息同步、报关与跨境（7 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.trade.orderinfo.sync | POST `/v3/alipay/trade/orderinfo/sync` | [接口文档](https://opendocs.alipay.com/open-v3/f8b82729_alipay.trade.orderinfo.sync) | 相同 out_request_no 重复请求不更新订单信息（网关幂等） |
| alipay.trade.customs.declare | POST `/v3/alipay/trade/customs/declare` | [接口文档](https://opendocs.alipay.com/pre-apis/05wyhb) | 报关为异步流程，受理成功≠报关完成，以报关查询为准 |
| alipay.trade.customs.query | GET `/v3/alipay/trade/customs/query` | [接口文档](https://opendocs.alipay.com/pre-apis/05wyjm) | 单次最多 10 个报关请求号 |
| alipay.trade.overseas.settle | POST `/v3/alipay/trade/overseas/settle` | [接口文档](https://opendocs.alipay.com/apis/05v7c4) | out_request_no 幂等；写操作不自动重试 |
| alipay.trade.advance.consult | GET `/v3/alipay/trade/advance/consult` | [接口文档](https://opendocs.alipay.com/open-v3/11cb1f6e_alipay.trade.advance.consult) | 风险预评估场景需 consult_scene+agreement_no；user_out_binding_phone 需按官方要求加密 |
| alipay.trade.batch.transfer | POST `/v3/alipay/trade/batch/transfer` | [接口文档](https://opendocs.alipay.com/pre-apis/05vd9z) | 多个分账条款要求出账方唯一；处理结果以查询为准 |
| alipay.trade.batch.transfer.query | GET `/v3/alipay/trade/batch/transfer/query` | [接口文档](https://opendocs.alipay.com/pre-apis/05vsbp) | settle_no 与 out_request_no 必填其一 |

### 分账（4 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.trade.royalty.rate.query | GET `/v3/alipay/trade/royalty/rate/query` | [接口文档](https://opendocs.alipay.com/open-v3/d9aece9f_alipay.trade.royalty.rate.query) | - |
| alipay.trade.royalty.relation.bind | POST `/v3/alipay/trade/royalty/relation/bind` | [接口文档](https://opendocs.alipay.com/open-v3/e701637d_alipay.trade.royalty.relation.bind) | 单次最多 20 个收款方；结果以 result_code 为准 |
| alipay.trade.royalty.relation.unbind | POST `/v3/alipay/trade/royalty/relation/unbind` | [接口文档](https://opendocs.alipay.com/open-v3/302cd1f8_alipay.trade.royalty.relation.unbind) | 同上 |
| alipay.trade.royalty.relation.batchquery | POST `/v3/alipay/trade/royalty/relation/batchquery` | [接口文档](https://opendocs.alipay.com/open-v3/1a11475c_alipay.trade.royalty.relation.batchquery) | 名为 batchquery 实为 POST+JSON body；page_size 取值 (0,100] |

### 结算（6 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.trade.order.settle | POST `/v3/alipay/trade/order/settle` | [接口文档](https://opendocs.alipay.com/open-v3/34402227_alipay.trade.order.settle) | out_request_no 幂等；分账结果以查询接口确认 |
| alipay.trade.order.settle.query | GET `/v3/alipay/trade/order/settle/query` | [接口文档](https://opendocs.alipay.com/open-v3/3af46ab7_alipay.trade.order.settle.query) | settle_no 或（out_request_no+trade_no）二选一 |
| alipay.trade.settle.confirm | POST `/v3/alipay/trade/settle/confirm` | [接口文档](https://opendocs.alipay.com/apis/0hp5md) | out_request_no 幂等；写操作不自动重试 |
| alipay.trade.settle.receivables.query | POST `/v3/alipay/trade/settle/receivables/query` | [接口文档](https://opendocs.alipay.com/pre-apis/05vlu7) | 名为 query 实为 POST+JSON body；直付通账期查询单笔交易时 trade_no 必传 |
| alipay.trade.batch.settle | POST `/v3/alipay/trade/batch/settle` | [接口文档](https://opendocs.alipay.com/pre-apis/05vfpn) | 重复 out_request_no 返回既有单据状态；处理结果以批量结算查询为准 |
| alipay.trade.batch.settle.query | POST `/v3/alipay/trade/batch/settle/query` | [接口文档](https://opendocs.alipay.com/pre-apis/05vel6) | 名为 query 实为 POST+JSON body；settle_no 与 out_request_no 必填其一 |

## 2. 资金转账 `services/fund/trans/`（7 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.fund.trans.uni.transfer | POST `/v3/alipay/fund/trans/uni/transfer` | [接口文档](https://opendocs.alipay.com/open-v3/26a22ad7_alipay.fund.trans.uni.transfer?scene=common%2Cca56bca529e64125a2786703c6192d41%2C66dd06f5a923403393b85de68d3c0055&pathHash=9b52afe2) | 资金支出类：发送前强制校验证书模式，产品权限由网关确认；写操作不自动重试，超时以查询确认 |
| alipay.fund.trans.toaccount.transfer | POST `/v3/alipay/fund/trans/toaccount/transfer` | [接口文档](https://opendocs.alipay.com/open-v3/53a4ce8e_alipay.fund.trans.toaccount.transfer) | 同上；金额下限 0.1 元 |
| alipay.fund.trans.pay | POST `/v3/alipay/fund/trans/pay` | [接口文档](https://opendocs.alipay.com/pre-apis/05xggg) | status=UNKNOWN 必须查询确认 |
| alipay.fund.trans.refund | POST `/v3/alipay/fund/trans/refund` | [接口文档](https://opendocs.alipay.com/open-v3/0c6ff10a_alipay.fund.trans.refund) | 退回失败重提必须沿用原 out_request_no |
| alipay.fund.trans.order.query | GET `/v3/alipay/fund/trans/order/query` | [接口文档](https://opendocs.alipay.com/open-v3/f00381dd_alipay.fund.trans.order.query) | 到账状态需持续跟踪（银行卡产品可能由成功转退票） |
| alipay.fund.trans.common.query | GET `/v3/alipay/fund/trans/common/query` | [接口文档](https://opendocs.alipay.com/open-v3/bcbe6351_alipay.fund.trans.common.query?scene=common&pathHash=36f664d9) | 三单号优先级 pay_fund_order_id > order_id > out_biz_no；out_biz_no 必带 product_code/biz_scene |
| alipay.fund.trans.payee.bind.query | GET `/v3/alipay/fund/trans/payee/bind/query` | [接口文档](https://opendocs.alipay.com/open-v3/03fe4a0d_alipay.fund.trans.payee.bind.query?scene=common&pathHash=377d1f29) | 暂仅支持身份证 |

## 3. 账单 `services/data/bill/`（4 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.data.dataservice.bill.downloadurl.query | GET `/v3/alipay/data/dataservice/bill/downloadurl/query` | [接口文档](https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49) | 下载地址 30 秒失效；账单文件内容校验由调用方完成 |
| alipay.data.bill.balance.query | GET `/v3/alipay/data/bill/balance/query` | [接口文档](https://opendocs.alipay.com/open-v3/9b66d9d8_alipay.data.bill.balance.query) | - |
| alipay.data.bill.balancehis.query | GET `/v3/alipay/data/bill/balancehis/query` | [接口文档](https://opendocs.alipay.com/open-v3/029a9c24_alipay.data.bill.balancehis.query) | biz_date 与 biz_month 互斥 |
| alipay.data.bill.accountlog.query | GET `/v3/alipay/data/bill/accountlog/query` | [接口文档](https://opendocs.alipay.com/open-v3/4cce9615_alipay.data.bill.accountlog.query) | 时间范围间隔 ≤31 天；支付宝订单号与商户订单号互斥 |

## 4. 授权（各 1 项）

| 接口 | 方法与路径 | 官方文档 | 限制/备注 |
| --- | --- | --- | --- |
| alipay.system.oauth.token | POST `/v3/alipay/system/oauth/token` | [接口文档](https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token) | `services/system`；令牌为敏感凭证，调用方安全保存；SDK 不缓存令牌 |
| alipay.open.auth.token.app | POST `/v3/alipay/open/auth/token/app` | [接口文档](https://opendocs.alipay.com/open-v3/4b00cc9c_alipay.open.auth.token.app) | `services/open`；应用授权令牌经 `core.WithAppAuthToken` 代调用，不修改共享 Client |

## 5. 非 REST 能力（不在 openapi.yaml 覆盖范围）

以下能力不发起服务端 REST 调用或属于支付宝→商户方向，契约依据是官方页面/唤端产物构建规范与通知验签文档，无法以 openapi.yaml 核对：

| 能力 | 形态 | 官方依据 | 状态 |
| --- | --- | --- | --- |
| alipay.trade.page.pay（产物构建） | 跳转 URL（GET 语义）或自动提交 HTML Form（POST 语义），经 `core.NewArtifactBuilder` 构建 | [官方 V3 文档](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay)、[页面跳转说明](https://opendocs.alipay.com/common/0i2kxr.md) | 已实现，产物签名独立复算验证；沙箱可提交性未验证 |
| alipay.trade.wap.pay（产物构建） | 同上 | [官方 V3 文档](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay) | 同上 |
| alipay.trade.app.pay（产物构建） | 订单串（orderStr），交支付宝客户端 SDK；含支付并签约参数 | [官方 V3 文档](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay)、[唤端说明](https://opendocs.alipay.com/common/0i2jez.md)、[支付并签约](https://opendocs.alipay.com/open-v3/0jf33m) | 已实现，订单串签名独立复算；沙箱可提交性未验证 |
| 普通支付异步通知（notify_type=trade_status_sync） | 表单解码与验签（`core/notify` + `services/trade.NotificationDecoder`） | [异步通知说明](https://opendocs.alipay.com/open/194/103296)、[02mse7 验签规则](https://opendocs.alipay.com/common/02mse7.md) | 已实现，签名链路端到端验证；真实通知报文未验证 |

## 6. 协议能力覆盖

| 能力 | 状态 | 说明 |
| --- | --- | --- |
| V3 REST 协议 | 已实现（离线验证） | Authorization 头 ALIPAY-SHA256withRSA 签名、响应头验签（alipay-signature/alipay-timestamp/alipay-nonce，证书模式按 alipay-sn 选择可信证书）、非 2xx 错误体 `{code,message,links}` 解析 |
| multipart/form-data 请求编码 | 已实现并离线验证 | `core.MultipartBodyEncoder` 将文件与 JSON 数据编码为 multipart 请求；仅 JSON 数据部分参与签名，文件和边界不参与签名。 |
| RSA2 签名（PKCS#8 私钥） | 已实现（离线验证） | SHA256WithRSA，签名值 base64；openssl 独立向量交叉验证 |
| 公钥模式 | 已实现（离线验证） | 应用私钥 + 支付宝公钥（PKIX PEM） |
| 证书模式 | 已实现（离线验证） | `auth.NewCertificateCredentials` 装载应用/支付宝/根证书；请求 Authorization 携带 app_cert_sn、另发 alipay-root-cert-sn 头；响应按 alipay-sn 选择可信证书验签；SN 规则经 openssl 独立向量交叉验证；不自动下载证书，未知 SN 显式失败 |
| Credentials 认证契约 | 已实现（离线验证） | `core.NewClient(credentials, opts...)` 首参必选；公钥/证书两种内置凭证 + 外部可实现；`OperationSpec.AuthRequirement` 在统一 Do 路径校验；`notify.NewHandler` 以 Credentials 驱动通知验签 |
| 页面/App 支付产物构建 | 已实现（离线验证） | `core.NewArtifactBuilder` + `trade.PaymentProducts`：page.pay/wap.pay URL 与 HTML Form、app.pay 订单串；签名经独立复算；不发 HTTP 请求 |
| 异步通知 | 已实现（离线验证） | 仅普通支付通知策略；独立表单解码与验签，与 REST 验签边界分开 |
| 自定义接口扩展 | 已实现 | `NewJSONOperation`（默认编解码）/ `NewOperation`（自定义 Encoder/Decoder）；见 examples/customapi |
| 内容加密 | 未实现 | 请求该能力时当前无入口；不将签名私钥用作加密密钥 |
| UTF-8 JSON | 已支持 | charset 固定 utf-8；金额为十进制字符串 |
| GBK / XML / RSA(v1) 签名 | 不支持 | 当前仅支持 UTF-8 JSON 与 RSA2 |

## 7. 测试证据索引

- 密码学：`core/auth/signers/rsa2_test.go`（openssl 已知答案向量、篡改/错钥/错签失败）。
- V3 协议：`internal/protocol/v3_test.go`（Authorization authString/签名原文/验签原文固定向量、错误体解析与缺字段拒绝、常量锁定）；`internal/protocol/protocol_test.go`（通知待验签串构造规则等）。
- REST Client：`core/client_test.go`（服务端独立复算 Authorization 签名、GET query 签名、篡改/错钥/缺签失败、签名与未签名错误体分类、代调用令牌入签、notify_url 选项拒绝、规格与参数模式推断校验、调用方参数不可变、超时/取消/超限、日志级别与脱敏、并发隔离、泛型入口一致性）。
- 非 GET 的 query、query 与 JSON 组合、空请求体：55 个接口的本地 HTTP 与签名回归，见[离线契约测试](../examples/generatedapi/query_contract_test.go)。
- URL 路径参数：`core/path_params_test.go`（GET query 与 JSON body 分离、路径值转义、路径模板缺失或多余值拒绝）；快照中的 16 个路径参数接口均有对应 `path` 字段。
- 证书模式：`core/client_cert_test.go`（请求携带 SN 且参与签名核对、按 alipay-sn 验签成功、未知/缺失 SN 显式失败、篡改失败、轮换多证书选择）；`services/trade/cert_mode_test.go`（同一 trade.Service 与 core/notify 复用证书 Client；通知用支付宝证书公钥验签；notify_url 选项拒绝）。
- 构造边界：`core/client_credentials_test.go`（nil/typed-nil 凭证、校验失败、模式与能力矛盾、第三方 Credentials 实现、Settings 快照只读）；`core/errors_test.go`、`core/option/option_test.go`（错误类型与选项行为）。
- 证书装载与验签选择：`core/auth/credentials_test.go`（私钥与证书公钥匹配、有效期、根证书链、非 CA 根、多支付宝证书轮换、身份副本不可变、矛盾配置拒绝）；`core/auth/resolver.go` 按 SN 选择证书，普通支付通知使用唯一配置证书。运行期有效期由 `core/auth/validity.go` 检查，边界、快照隔离和各入口到期失败见 `core/auth/validity*_test.go` 与 `core/certificate_validity_go125_test.go`。
- 证书 SN 规则：`internal/protocol/certs_test.go`（openssl 独立已知答案向量）；`internal/protocol/cert_params_test.go`（标识提取与受保护参数）；`utils/certs_test.go`（PEM 解析与链校验）。
- 交易 Service：`services/trade/api_trade_*_test.go` 覆盖基础交易、信息同步、跨境与咨询；`services/trade/{order,refund,batch,royalty,settle,customs,advance,overseas,orderinfo}/api_<包名>*_test.go` 覆盖迁移后的操作规格、请求序列化、校验与状态映射。端到端签名与业务错误测试保留在根包；各包的 `api_<包名>_example_test.go` 提供可编译示例。
- 各领域 API 示例恢复 `api_*_example_test.go`，由默认测试编译；`TestSandboxService_*` 提供显式启用的沙箱证书模式入口，运行方式见[沙箱验收指引](sandbox-verification.md)。`internal/sandbox/sandbox_test.go` 离线覆盖开关隔离、请求解码、证书材料装载、证书身份签名与响应篡改拒绝；这些检查不表示真实沙箱验收通过。
- 产物构建：`services/trade/pagepay_test.go`（表单/URL/订单串签名独立复算、biz_content 隐藏域与双引号转义、篡改失败、证书参数、校验分支）；实现位于 `core/artifact.go`。
- 通知：`core/notify/notify_test.go`（单次解码与 `+`/`%2B`/`%26`/中文/空值一致性、重复参数拒绝、缺签/错签/错公钥/错 sign_type、未支持策略与 notify_type 拒绝、请求体大小上限、日志脱敏与级别、共享验签器并发）；`services/trade/api_notify_test.go`（官方参数表映射、必选参数校验、未知枚举保留、端到端签名链路）；`services/trade/api_notify_example_test.go`（业务校验与幂等接入位置）。
- 其他领域：`services/fund/trans/api_trans_*_test.go`（转账规格/序列化/校验 + httptest 端到端）、`services/system/oauth_test.go`、`services/open/authtoken_test.go`、`services/data/bill/api_bill_*_test.go`（规格、映射与校验）。
- 示例：`examples/certmode`（证书模式最小接入）、`examples/logging`（默认静默、标准/自定义 slog Handler、日志资源由调用方关闭）、`examples/customapi/customapi_test.go`（外部包仅依赖公开包；错签时不调用用户解码器；受保护参数拦截）。

## 8. 未确认项

REST 契约清点无不一致项；以下为实现与验证层面的未确认项，均待沙箱以真实报文复核：

### 协议与网关侧

- V3 生产网关对 Authorization 各字段的具体校验容差（timestamp 时钟偏移、nonce 重复策略）未见于公开文档，待沙箱确认。
- V3 沙箱基地址使用 HTTPS（`https://openapi-sandbox.dl.alipaydev.com`），无凭证请求已确认路由可达；沙箱对 V3 及证书模式的支持情况未验证。
- 证书模式响应头的真实沙箱验收仍待完成。SDK 要求带签名响应携带 `alipay-sn`、`alipay-timestamp` 和 `alipay-nonce`，缺失时明确报错。
- 2026-11-06 按固定快照复核请求与成功响应字段，两个令牌接口省略的废弃字段及逐接口结果见[核查记录](checked.md)。网关新增的响应字段按解码规则忽略，可通过已验签的响应原文读取。

### 页面产物与通知

- 页面/唤端产物不在 openapi.yaml 覆盖范围，产物格式与签名字段以官方页面/唤端构建规范为准，沙箱实际可提交性未验证；产物中的 `alipay_sdk` 参数（官方 SDK 携带的版本标识）本 SDK 未输出，如沙箱要求再补充。
- 产物 timestamp 按官方实现固定 GMT+8；是否需调整待沙箱确认。
- 产物构建与通知沿用的 `biz_content` 签名形态（含中文/特殊字符时值不 URL 编码）依据官方签名规则实现，待沙箱确认。
- 通知空值参数处理文档未明确，按官方语言 SDK 行为保留在待验签串中；SDK 要求显式携带 `sign_type=RSA2`；官方示例报文未携带参数表必选的 `app_id` 等字段（示例较旧）。以上均待沙箱以真实报文复核。

### 文档与证书材料

- precreate 文档 URL（02np92）为浏览器渲染抓取，`.md` 端点多次返回空正文；字段以页面原文为准，沙箱验收时复核。
- 退款查询的 `deposit_back_info_list` 要求的 query_options 取值与枚举列表存在原文不一致（枚举仅列 `deposit_back_info`），已按原文建模并标注。
- 退款查询间隔建议在两处文档不一致（5s / 10s），GoDoc 取更保守的 10s。
- 证书模式：SN 计算的签发机构名固定按 UTF-8 取字节（官方 Java 按平台默认字符集，JDK 18+ 即 UTF-8）；装载时对支付宝证书执行根证书链校验（官方仅对下载证书校验），多级链需要显式提供中间 CA，并已通过离线验证。`app_cert_sn`/`alipay_root_cert_sn` 的网关侧接受性与 `alipay-sn` 的真实取值尚未经沙箱确认。

## 9. 当前 `alipay.*` 接口实现

全部 678 个 `alipay.*` REST 操作基于官方 V3 OpenAPI 快照（`info.version=2026-09-23`）。每个公开方法的 GoDoc 链接到官方接口说明；复杂领域按业务前缀拆分 Service 子包，请求、响应与共享模型和接口实现放在同一子包目录。

| 领域 | 接口数 | 领域 | 接口数 |
| --- | ---: | --- | ---: |
| `services/boss` | 8 | `services/commerce` | 55 |
| `services/data` | 18 | `services/ebpp` | 77 |
| `services/eco` | 32 | `services/fund` | 42 |
| `services/iservice` | 45 | `services/marketing` | 64 |
| `services/merchant` | 18 | `services/mobile` | 14 |
| `services/offline` | 8 | `services/open` | 241 |
| `services/pcredit` | 6 | `services/security` | 3 |
| `services/social` | 1 | `services/system` | 1 |
| `services/trade` | 27 | `services/user` | 18 |
| **合计** | **678** | | |

这 638 个生成接口目前未逐项核验产品权限、证书要求、字段条件校验及业务状态语义，也未进行逐接口沙箱验证。官方 OpenAPI 快照未完整表达这些约束，使用前须查看对应方法 GoDoc 链接的具体官方文档。SDK 不自动重试写操作；未知或处理中业务状态应由调用方按接口文档处理。
