# 沙箱验收指引与记录

状态：**未验收**（2026-11-05）。当前缺少真实沙箱应用与凭证，本文件先固定验收步骤、用例矩阵与记录模板；已完成的验证仍以[接口覆盖记录](api-coverage.md)为准，未标注"沙箱验证"的能力不应用于生产判断。

## 1. 验收原则

- **显式启用沙箱**：代码中使用 `option.WithEnvironment(core.EnvironmentSandbox)`（或示例中的 `ALIPAY_ENV=sandbox`）；SDK 不提供自动回退或环境混用。
- 沙箱 V3 REST 基地址为 `https://openapi-sandbox.dl.alipaydev.com`；页面/唤端产物网关为 `https://openapi-sandbox.dl.alipaydev.com/gateway.do`。与 [官方沙箱接入说明](https://opendocs.alipay.com/common/097jyh.md)核对，验收时以官方页面为准。
- 生产凭证、生产 APPID 不用于沙箱；沙箱结果不外推为生产可用性。
- **公钥模式与证书模式分别配置、分别记录**，不混用同一份材料。
- 记录只保存接口名、订单号、时间戳、码值与结论，不记录私钥、完整报文或个人信息。

## 2. 前置条件

| 条件 | 说明 |
| --- | --- |
| 沙箱应用 | 开放平台"沙箱环境"提供的沙箱 APPID 与沙箱买家账号 |
| 公钥模式材料 | 沙箱应用配置 RSA2 应用公钥，取得支付宝公钥（PKIX PEM） |
| 证书模式材料 | 沙箱应用配置证书加签方式，下载应用证书、支付宝证书与根证书 |
| 产品权限 | 确认沙箱应用已具备当面付（precreate）等目标产品权限；沙箱权限以官方页面为准 |
| 通知回调端点 | 可公网访问的 HTTPS 端点（如内网穿透），仅用于 S5 通知用例 |

### 通过各领域 API 示例测试调用沙箱

`services/` 下的 API 示例统一使用 `api_*_example_test.go`。每个 REST `ExampleService_*` 都完整展示平台证书装载、Client 与 Service 创建、具体请求及响应处理；全部 746 个 REST 方法均有对应示例。示例无 `Output` 注释，默认 `go test ./...` 只编译、不执行其函数体。配置与复制使用方式见[接口使用示例](api-examples.md)。其中已有的 `TestSandboxService_*` 通过共享证书装载入口调用对应的真实 Service 方法；未显式启用时跳过，不访问支付宝。

按[官方沙箱快速接入的“配置证书模式”](https://opendocs.alipay.com/common/097jyh.md)在沙箱控制台启用证书模式，下载应用公钥证书、支付宝公钥证书和支付宝根证书，并取得匹配的应用私钥。平台证书用于验签，仍需应用私钥签名；无需配置 `ALIPAY_PUBLIC_KEY`。

先在仓库外准备请求文件，例如 `/absolute/path/query.json`：

```json
{"out_trade_no":"your-sandbox-order-no"}
```

从仓库根目录运行：

```sh
export ALIPAY_SANDBOX=1
export ALIPAY_ENV=sandbox
export ALIPAY_SANDBOX_API='alipay.trade.query'
export ALIPAY_APP_ID='your-sandbox-app-id'
export ALIPAY_APP_PRIVATE_KEY='/absolute/path/app-private-key.pem'
export ALIPAY_APP_CERT='/absolute/path/appPublicCert.crt'
export ALIPAY_CERT='/absolute/path/alipayPublicCert.crt'
export ALIPAY_ROOT_CERT='/absolute/path/alipayRootCert.crt'
export ALIPAY_SANDBOX_REQUEST='/absolute/path/query.json'
go test ./services/trade -run '^TestSandboxService_Query$' -count=1 -v
```

- `ALIPAY_SANDBOX=1` 显式启用真实调用；`ALIPAY_SANDBOX_API` 必须与测试中的官方接口名完全相等。不匹配的测试跳过，因此运行后应确认目标测试实际执行而非 `SKIP`。配置不完整或调用失败会使目标测试失败。
- `ALIPAY_ENV` 可不传，入口始终固定 `core.EnvironmentSandbox`；显式传入 `production` 或其他值会失败。不接受网关覆盖，也不回退到生产或其他协议。
- `ALIPAY_SANDBOX_REQUEST` 是该方法请求 struct 的 JSON 对象，不包含公共参数或 `biz_content` 包装。带 `json` 标签的字段使用标签名；只有 `query` 或 `path` 标签的字段使用 Go 字段名，例如 `{"PageSize":10}`。未知字段、重复键、类型错误与 `null` 请求被拒绝，防止配置拼错后静默发送默认值。
- multipart 请求仍按请求模型填写嵌套对象，例如 `file_content` 为 `{"Filename":"example.txt","ContentType":"text/plain","Content":"aGVsbG8="}`；`Content` 是文件内容的 base64 字符串。
- `ALIPAY_CERT` 可指向含多张 PEM 证书的文件，也可用逗号分隔多个证书文件。多级链须一并提供签发平台叶证书的中间 CA。应用证书、私钥匹配、有效期和支付宝证书根信任校验均复用 `auth.NewCertificateCredentials`，校验失败不降级。
- 需要代商户调用时额外配置 `ALIPAY_APP_AUTH_TOKEN`；它只作为本次调用选项传入。请求文件与凭证放在仓库外，不记录或提交令牌、授权码等敏感值。
- 其他领域选择对应目录、`TestSandboxService_<方法名>` 和官方接口名，并按该领域的请求模型提供真实沙箱参数。例如资金查询使用 `./services/fund`、`TestSandboxService_OrderQuery`、`alipay.fund.trans.order.query`。每个领域 README 的 API 示例链接均指向对应入口。

测试成功只表示该次调用返回了已验签的类型化响应，仍需按接口定义确认业务状态。证书与入口可用于各领域示例，不代表沙箱支持所有产品；权限或产品不支持会明确失败，不能记为验收通过。写操作使用沙箱专用且稳定的业务请求号，入口不自动重试。`-count=1` 避免真实调用被 Go 测试缓存跳过。

通知示例 `api_notify_example_test.go` 使用平台证书验签，从 `ALIPAY_NOTIFY_FORM` 指向的原始通知表单文件演示 `ParseRequest`，默认仅编译。真实通知需使用公网回调端点与[完整通知示例](../examples/tradenotify/main.go)按 S5 单独验证，不能用平台公钥证书伪造通知签名。

## 3. 用例矩阵

每条用例在公钥模式与证书模式下各执行一次并分别记录。V3 的成功语义由 HTTP 状态决定；HTTP 成功不代表业务终态，具体判定见各用例期望。

| 编号 | 场景 | 接口链 | 期望结果 |
| --- | --- | --- | --- |
| S1 | 预下单 | `precreate` | HTTP 2xx，返回 `qr_code`；同一 `out_trade_no` 重复申请应被拒绝（记录错误码原文） |
| S2 | 未支付查单 | `query` | HTTP 2xx 且 `trade_status=WAIT_BUYER_PAY`；查询成功≠支付成功 |
| S3 | 未支付关单 | `close` | 关单成功；对 `WAIT_BUYER_PAY` 之外状态关单返回对应业务错误（`*core.APIError`） |
| S4 | 支付后查单 | 沙箱买家付款 → `query` | `trade_status=TRADE_SUCCESS`，金额与预下单一致 |
| S5 | 支付通知 | 异步通知 → 本 SDK 验签 | 验签通过；商户侧校验与幂等按[交易场景说明](trade-scenarios.md)第 3 节执行；应答 `success` 后支付宝停止重发 |
| S6 | 退款 | `refund` | HTTP 2xx 且 `fund_change=Y` 表示退款成功；同笔交易间隔 ≥3 秒，部分退款 `out_request_no` 保持不变 |
| S7 | 退款查询 | `fastpay.refund.query` | 仅 `refund_status=REFUND_SUCCESS` 才是退款成功 |
| S8 | 业务错误分支 | 对不存在订单 `query` | `*core.APIError`，`Code=ACQ.TRADE_NOT_EXIST`；记录错误响应是否携带 `alipay-signature` 头及 `Verified` 取值 |
| S9 | 本地校验分支 | `subject` 含 `/` `=` `&`、金额非两位小数 | `*core.ValidationError`，请求未发送到网关 |

补充核对项（伴随上述用例）：

- 请求侧：Authorization 头结构与[自行实现签名](https://opendocs.alipay.com/common/057k53.md)的 V3 规则一致；证书模式 `app_cert_sn`（authString 内）与 `alipay-root-cert-sn` 头被网关接受。
- 响应侧：证书模式响应头 `alipay-sn` 的真实取值能命中本地证书；篡改检测（可临时改错本地公钥验证验签失败路径）。
- 取消/超时：ctx 取消后请求返回 `*core.TransportError`，且资金写操作以查单确认结果。

## 4. 必须复核的未确认项

以下决策目前依据官方文档与官方 SDK 行为实现，验收时以沙箱真实报文逐条复核（详见[接口覆盖记录](api-coverage.md)"未确认项"与[协议说明](protocol-notes.md)）：

1. **证书模式响应头验收**。检查成功响应及带签名错误响应中的 `alipay-sn`、`alipay-timestamp` 与 `alipay-nonce`。SDK 对缺失字段、未知证书和签名失败直接报错，调用方需排查网关报文或更新本地证书。
2. **沙箱 V3 基地址连通性**（本次改造后新增）：2026-11-05 使用 TLS 验证成功的 GET 请求访问 `/v3/alipay/fund/trans/order/query`，返回 HTTP 400 和 `missing-timestamp`。公钥和证书模式的业务验收仍需沙箱凭证。
3. **页面/唤端产物的实际可提交性**（本次改造后新增）：page.pay/wap.pay 产物在沙箱经 `gateway.do` 提交后能否正常渲染并完成支付；app.pay 订单串能否在支付宝客户端/沙箱工具唤起；产物字段（含证书参数、`alipay_sdk` 缺失）是否被沙箱接受。
4. `biz_content` 含中文时待签名串的形态（本 SDK 按"值不 URL 编码"实现，适用于页面产物）。
5. 通知空值参数是否参与待验签串，以及实际通知是否完整提供 `sign_type=RSA2`。SDK 对缺失或其他算法值直接报错。
6. `precreate` 字段契约（文档以浏览器渲染页面抓取为准，非 `.md` 端点）。
7. 证书 SN 计算的签发机构名按 UTF-8 取字节；装载时对支付宝证书执行根证书链校验，多级链使用显式配置的中间 CA。
8. 退款查询间隔建议（GoDoc 取更保守的 10 秒）。

## 5. 当前记录

| 用例 | 公钥模式 | 证书模式 | 记录 |
| --- | --- | --- | --- |
| S1–S9 全部 | 未验证 | 未验证 | 2026-11-05：缺少沙箱应用与凭证，未执行 |
| S5 通知回调 | 未验证 | 未验证 | 通知需要可公网访问的回调端点，单独列为未验证项 |
| 第 4 节第 1–3 条（改造后新增） | 未验证 | 未验证 | 2026-11-05：V3-only 改造完成后待核验 |
| 补充核对项 | 未验证 | 未验证 | — |

验收完成后回填：本表、[接口覆盖记录](api-coverage.md)各接口"沙箱"状态、[开发计划](development-plan.md)第 16 节遗留项，以及 [CHANGELOG](../CHANGELOG.md) 的沙箱验收状态。
