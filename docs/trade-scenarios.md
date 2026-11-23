# 交易场景说明

本文说明当前已实现的支付宝 V3 交易链路。

状态：2026-11-05。本文档面向使用本 SDK 的开发者，说明交易闭环各场景的接入顺序、SDK 入口与必须由商户完成的校验。协议依据见[协议说明](protocol-notes.md)，能力边界见[接口覆盖记录](api-coverage.md)。本文描述的能力覆盖 RSA2 公钥模式、本地证书模式与 UTF-8 JSON（证书模式规则见[协议说明](protocol-notes.md)第 6 节）；页面/App 支付产物构建见 `trade.PaymentProducts` 与 `examples/pagepay`；内容加密未实现。

## 1. 场景总览

```text
预下单(precreate) → 展示二维码
        │
        ├─ 用户扫码支付 ──→ 异步通知(notify_url) → 验签+业务校验+幂等 → 应答 success
        │            └──→ 查询确认(trade.query)：未收到通知时的兜底，官方强制要求
        │
        ├─ 买家未付款 → 查询确认 WAIT_BUYER_PAY → 关单(trade.close) → 可用新订单号重新预下单
        │
        └─ 支付成功后退款(refund) → 未确认时退款查询(fastpay.refund.query)
```

官方要求商户**同时接入支付结果异步通知与统一收单交易查询接口**：未收到异步通知时必须调用 `alipay.trade.query` 查询支付结果；针对同一业务订单更换 `out_trade_no` 重新发起支付前，先查询前一订单号——`TRADE_SUCCESS` 则不要再重复支付，`WAIT_BUYER_PAY` 则先关单再重新预下单（依据[异步通知说明](https://opendocs.alipay.com/open/194/103296)）。

## 2. 预下单（当面付扫码）

入口：`trade.Service.PreCreate`，V3 POST `/v3/alipay/trade/precreate`。

- 必选参数：`out_trade_no`（字母/数字/下划线）、`total_amount`（两位小数，[0.01,1e8]）、`subject`（禁用 `/` `=` `&`）。金额使用十进制字符串，不得用浮点数。
- 通知地址：V3 中 `notify_url` 属于业务字段（官方 V3 schema 的 precreate、pay、create、close、refund.apply 等请求模型含该字段）；当前 SDK 请求模型尚未建模 `notify_url`，REST 调用路径也不提供该调用选项（`core.WithNotifyURL` 仅用于页面产物构建）。接入回调请先在开放平台按应用/产品配置通知地址，SDK 建模该字段后可按接口传入。
- `qr_code` 有效期固定 2 小时，同一订单号不能重复申请。
- HTTP 成功仅表示预下单受理成功，不代表用户已支付；支付结果以下述通知或查询为准。

## 3. 支付完成异步通知（普通支付通知）

SDK 入口：`notify.NewHandler(credentials, notify.PolicyTradeStatusSync)` 构建 Handler，再调用其 `ParseRequest`（原始 HTTP 请求）或 `Parse`（已解码一次的表单）并以 `trade.NotificationDecoder()` 完成领域映射；验签器由凭证派生（公钥模式用支付宝公钥，证书模式用支付宝证书公钥）。通知按自身契约走独立表单验签，与 REST 响应头验签边界分开。完整最小示例见 `examples/tradenotify`；参数表依据[异步通知说明](https://opendocs.alipay.com/open/194/103296)。

处理顺序及每一步的责任边界：

1. **验签**（SDK）：按普通支付通知规则剔除 `sign`、`sign_type` 构建待验签串（保留空值参数），用可信支付宝公钥 RSA2 验签。验签失败应答 `fail`。
2. **身份校验**（商户）：`app_id` 必须为本应用；`seller_id`（或 `seller_email`）必须为该订单的收款方。
3. **订单校验**（商户）：`out_trade_no` 必须是系统中创建的订单号；`total_amount` 必须与订单创建时的金额一致。任何一项不通过都应忽略该通知。
4. **幂等处理**（商户）：同一条通知重发时 `notify_id` 不变，以其（或订单状态机）去重；重复通知直接应答 `success`。
5. **状态语义**：只有 `TRADE_SUCCESS` / `TRADE_FINISHED` 表示买家付款成功；`WAIT_BUYER_PAY`、`TRADE_CLOSED` 与未知新状态按各自业务语义处理。`Notification.Paid()` 辅助方法内置该判断。
6. **应答**：处理完成后原样输出小写 `success`（无空格、HTML 或跳转）。未收到 `success` 时支付宝先立即重发 3 次，再按 4m、10m、10m、1h、2h、6h、15h 重发。

注意：验签只证明通知来自支付宝，不能替代上述 2–4 的校验；SDK 示例不会在仅验签后确认订单。

## 4. 查询确认

入口：`trade.Service.Query`，V3 POST `/v3/alipay/trade/query`。

- HTTP 成功仅表示查询成功；支付成功以 `trade_status` 为准。
- 通知未到达、超时或对结果有疑义时以查询结果做兜底；建议以查询作为订单状态的最终依据。

## 5. 未支付订单关闭

入口：`trade.Service.Close`，V3 POST `/v3/alipay/trade/close`。

- 仅 `WAIT_BUYER_PAY` 状态的交易可关单；SDK 不依据本地状态自动发起关单。可编译示例 `ExampleService_Close` 演示"先查询确认状态、再关单"的顺序。
- 关单响应只包含 `trade_no` / `out_trade_no`。关单成功后可使用新的 `out_trade_no` 重新预下单。

## 6. 退款与退款查询

导入 `github.com/tiyee/alipay-sdk-go/services/trade/refund`。入口：`refund.Service.Refund` / `QueryRefund`，V3 POST `/v3/alipay/trade/refund`、POST `/v3/alipay/trade/fastpay/refund/query`（名为 query 实为 POST+JSON body）。

- 部分退款必须提供 `out_request_no`，同一交易号下唯一；**重试时不得变更**，否则可能重复退款。同笔交易退款间隔 ≥3 秒，累计退款不超过原交易金额。
- HTTP 成功仅代表退款请求成功：`fund_change=Y` 才是退款成功；`N` 或未返回时必须通过退款查询确认，官方建议间隔 10 秒以上。
- 退款查询 HTTP 成功同样只代表查询成功；仅当 `refund_status=REFUND_SUCCESS` 才是退款成功，字段未返回表示退款请求未收到或失败。
- SDK 默认不自动重试任何写操作；退款超时后应查询确认，不要盲目重发。

## 7. 类型与错误速查

- 所有金额为十进制字符串；时间为官方格式的字符串，SDK 不解析时间格式。
- 交易状态、退款状态、通知类型等均为具名枚举类型，响应中的未知值保留原值，可在 `switch` 的 `default` 分支识别。
- 错误通过 `errors.As`（Go 1.26 起也可用 `errors.AsType[T]`） 匹配：`*core.SignatureError`（验签失败）、`*core.ProtocolError`（解码/协议异常）、`*core.ValidationError`（必选通知参数缺失）、`*core.APIError`（V3 网关业务错误，`Verified=false` 时仅作诊断）等；错误链保留原始原因。
