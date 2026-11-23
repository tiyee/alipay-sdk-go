# Changelog

本文件记录面向使用者的变更与已知限制。当前版本为**候选状态**，尚未打 tag 或发布模块；发布属于后续明确的发布操作。

## 未发布（v0.1.0 候选，2026-11-05）

首个测试版本候选：V3-only 单协议改造 + P7 交易模块补齐 + P8 授权/资金/账单 + 本地证书模式 + 工程整理。**未通过沙箱验收。**

### 支持范围与核心能力

- `core.NewClient(credentials, opts...)` 是唯一 REST 客户端，执行 V3 REST：以 Authorization 头签名请求、验签响应头，并解析 V3 错误体。
- `core.NewJSONOperation[Req, Resp](spec)` 支持独立 query、JSON、query 与 JSON 组合，以及 multipart；`core.NewOperation` 接入自定义编解码器。仅有 V3 REST Operation。
- 页面/App 产物使用独立配置本地签名构建；普通支付通知按自身契约解析和验签。两者都不扩展为传统网关服务端调用能力。
- `OperationSpec` 仅包含 V3 REST 所需路由、方法、参数位置与认证要求；错误类型对应 V3 错误体。
- `WithAppAuthToken` 经 `alipay-app-auth-token` 头传递并参与 V3 请求签名；`WithNotifyURL` 用于页面/唤端产物。

### 新增（V3 REST 与 P7/P8）

- **V3 REST 协议**：请求签名、响应验签与错误解析的文档依据见[协议说明](docs/protocol-notes.md)。接口快照版本和提交记录见[接口版本说明](docs/api-versions.md)。
- **services/trade（P7）**：27 个 REST 接口——query、precreate、refund、fastpay.refund.query、close、pay、create、cancel、order.pay、refund.apply、order.settle（+query）、settle.confirm、settle.receivables.query、batch.settle（+query）、royalty.rate.query、royalty.relation.bind/unbind/batchquery、orderinfo.sync、customs.declare（+query）、overseas.settle、advance.consult、batch.transfer（+query）。
- **页面/App 支付产物构建（P7）**：`core.NewArtifactBuilder` 与 `trade.PaymentProducts`——page.pay/wap.pay 的跳转 URL 与自动提交 HTML Form、app.pay 订单串；含支付并签约参数（SignParams）。构建不发 HTTP 请求，签名可经密钥工具独立核对。
- **授权/资金/账单（P8）**：`services/system`（alipay.system.oauth.token）、`services/open`（alipay.open.auth.token.app）、`services/fund`（fund.trans 7 接口）、`services/data`（对账单下载地址、余额、历史余额、账务明细）。
- **core.Client 通用能力**：Credentials 作为构造首参（公钥/证书两种内置实现，第三方可实现 `auth.Credentials` 接入）；HTTP 执行（ctx 传递、超时、响应上限、拒绝重定向、不修改全局 Client）；错误分类（config/validation/transport/canceled/http/protocol/signature/certificate/api）与`errors.As` 匹配；`slog` 日志（默认静默、筛选字段、级别约定）。
- **core/notify**：异步通知解析与验签，仅支持普通支付通知策略（剔除 `sign`、`sign_type`，保留空值参数）；有界读取、单次解码、泛型 `Parse` / `ParseRequest` 映射；`services/trade.NotificationDecoder()` 提供领域模型。通知链路与 REST 验签边界独立。
- **自定义接口扩展**：公开 `Operation[Req, Resp]`、`Encoder[T]` / `Decoder[T]`、`Executor.Do` 与包级 `core.Execute`（Go 1.27 起另提供 `Client.Execute[Req, Resp]` 泛型方法）；用户在自己的包内实现接口，无需修改 SDK、不导入 `internal/`、不注册全局方法表。
- **utils**：PKCS#8 私钥、PKIX 公钥与证书 PEM 解析、证书有效期与根证书链校验、密钥文本清理。
- **证书模式**：`auth.NewCertificateCredentials` 装配（有效期、私钥匹配、根证书链校验）；官方规则 SN 计算（openssl 独立向量交叉验证）；请求携带 `app_cert_sn` / `alipay_root_cert_sn`（证书模式另发 `alipay-root-cert-sn` 头），响应按 `alipay-sn` 选择可信证书验签；未知 SN 显式失败（fail closed），不自动下载证书。
- **示例**：`examples/certmode`（证书模式接入与更新）、`examples/tradenotify`（通知最小接入）、`examples/customapi`（V3 自定义接口）、`examples/pagepay`（页面/App 产物）、`examples/oauth`（授权令牌）、`examples/logging`（slog 静默/标准/自定义 Handler）。
- **文档**：架构、开发计划、接口覆盖记录、协议说明、交易场景、自定义接口、接口版本说明、沙箱验收指引、支付宝接口目录（docs/apis.md）。

### 已知限制

- **未沙箱验收**：全部能力仅有离线验证证据（`httptest` 端到端、mock、独立密码学向量），已通过无凭证的请求确认 HTTPS 沙箱路由可达，尚未执行业务联调；公钥/证书配置与沙箱验收步骤见[沙箱验收指引](docs/sandbox-verification.md)。缺少沙箱结果的能力不应用于生产判断。
- **生产网关未验证**：V3 REST、页面/App 产物构建与授权、资金、账单接口目前仅有离线验证记录，仍待沙箱交叉复核。协议依据见[协议说明](docs/protocol-notes.md)。
- 当前提供快照中的 746 个 REST 操作。手写响应模型覆盖常用业务字段，生成接口的业务约束和权限仍需逐项确认。快照以外的接口见[接口目录](docs/apis.md)。
- **协议范围**：仅 UTF-8 JSON + RSA2；不支持 GBK、XML、RSA(v1) 签名；内容加密无入口。
- **通知范围**：仅普通支付通知策略；生活号等其他通知类型显式拒绝。验签不替代商户侧的应用、收款方、订单、金额校验与幂等。
- **证书模式**：SN 的签发机构名固定按 UTF-8 取字节（官方 Java 按平台默认字符集）；装载时对支付宝证书执行根证书链校验；证书更新需重建实例，无自动下载。以上行为与真实证书的兼容性待沙箱复核。
- **重试**：默认不自动重试任何请求；写操作超时后可能已生效，重试需调用方保证幂等。
- 其他文档级未确认项（precreate 文档抓取方式、退款查询间隔建议分歧、通知空值与 `sign_type` 处理、证书模式响应 `alipay-sn` 的官方强制性与回退行为等）见[接口覆盖记录](docs/api-coverage.md)"未确认项"，均待沙箱以真实报文复核。

### 版本与 API

- 当前仍处于首版发布前，公开 API 按 V3 能力设计；不提供历史 API 兼容层。
- `core.NewClient(credentials, opts...)` 以 `auth.Credentials` 为第一个必选参数；公钥与证书凭证分别经 `auth.NewPublicKeyCredentials` / `auth.NewCertificateCredentials` 构造，公共 ClientOption 不覆盖认证策略。
- 模块最低支持 Go 1.24.0。各版本共享包级调用入口和 `json.RawMessage` 接口，版本相关的优化由编译标签自动选择。CI 覆盖 1.24 至 1.27 各系列及最新稳定版。
