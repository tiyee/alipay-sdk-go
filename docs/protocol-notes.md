# 协议说明（V3 REST、页面产物与通知）

本文件记录 SDK 实现所依据的支付宝 V3 协议规则与来源，涵盖 REST 签名与验签原文、页面/唤端产物格式及普通支付通知的独立表单验签。实现与规则不一致时以官方文档为准并修订本文件。接口覆盖见[覆盖记录](api-coverage.md)。

## 1. V3 请求签名（Authorization 头）

请求签名按支付宝 V3 [签名规则](https://opendocs.alipay.com/open-v3/054q58)构造，响应验签原文见[验签规则](https://opendocs.alipay.com/open-v3/054d0z)。

- 请求行：`{HTTP 方法} {网关基地址}{相对路由}`，相对路由以 `/v3/` 开头；JSON 接口使用 `application/json` 请求体，OpenAPI 声明文件上传的接口使用 `multipart/form-data` 请求体，query 参数按接口契约放入 URI，也适用于同时携带 JSON 正文的请求。
- Authorization 头：`ALIPAY-SHA256withRSA {authString},sign={base64 签名}`；authString = `app_id={app_id}` +（证书模式）`,app_cert_sn={sn}` + `,nonce={随机串}` + `,timestamp={毫秒时间戳}`。
- 签名原文（每行以 `\n` 结尾）：`authString`、HTTP 方法、请求 URI（不含域名，含 query）、请求体（为空时是空行）、代调用令牌（`alipay-app-auth-token`，非空才追加整行）。
- SDK 使用实际发送的 URI 签名。自定义 REST 网关可包含路径前缀及末尾斜杠；基地址中带有用户信息、查询串或片段时，Client 构造返回配置错误。查询参数中的切片按元素顺序编码为同名的多个值。
- 签名算法 SHA256WithRSA（RSA2），私钥为应用 PKCS#8 私钥。
- 证书模式另发 `alipay-root-cert-sn` 头；`alipay-request-id` 头为随机 ID，不参与签名。
- query 和 JSON body 使用传输时的编码结果签名。multipart 按[V3 文件上传签名规则](https://opendocs.alipay.com/open-v3/054q58)以 JSON 数据部分的原文字节签名，文件部分和 multipart 边界不参与签名；JSON 部分传输时不重新编码。

## 2. V3 响应验签

- 成功（2xx）：验签原文 = `alipay-timestamp`、`alipay-nonce`（响应头）与响应体三行，每行以 `\n` 结尾；签名取响应头 `alipay-signature`（base64）。**成功响应缺签名头视为协议错误**，不得作为可信结果。
- 默认 JSON 解码器在业务响应为结构体时要求响应节点具有对象语义。已验签的 `null` 返回 `*core.ProtocolError`，保留 `Result.SignatureVerified=true`，业务数据为空。带有效签名的网关业务错误也会同时设置 `Result.SignatureVerified` 与 `APIError.Verified`。
- 证书模式：按响应头 `alipay-sn`（支付宝公钥证书 SN）选择本地可信证书验签；缺失或未知 SN 失败（fail closed，不自动下载）。
- 本 SDK 明确要求带签名响应携带 `alipay-timestamp` 与 `alipay-nonce`，缺失时返回 `*core.ProtocolError`。证书模式还必须携带 `alipay-sn`，缺失或未知时返回 `*core.CertificateError`。成功响应与带签名的错误响应共用这些检查。验签字段见 [V3 验签文档](https://opendocs.alipay.com/open-v3/054d0z)；SDK 对缺失字段直接报错，官方端点是否始终提供这些字段仍待真实沙箱验收。
- 失败（非 2xx）：响应带 `alipay-signature` 时先验签（失败即失败）；未带签名按未验签诊断处理。

## 3. V3 错误模型

- 非 2xx + JSON 错误体 `{code, message, links?}`（官方 CommonErrorType）；`code` 为错误码原文（`core.ErrorCode`），网关级（如 `invalid-signature`）或业务级（如 `ACQ.TRADE_NOT_EXIST`）均按原值保留，不解析为数值。
- 成功语义由 HTTP 状态决定；2xx 响应无公共业务码字段，SDK 不使用成功码清单。
- 业务参数以 snake_case 为主（官方 openapi.yaml schema 如 `out_trade_no`、`page_data`）；金额为十进制字符串。
- 错误链：`ConfigError`/`ValidationError`/`TransportError`/`HTTPError`/`ProtocolError`/`SignatureError`/`CertificateError`/`APIError`，通过 `errors.As`（Go 1.26 起也可用 `errors.AsType[T]`） 匹配；`TransportError` 保留 `context.Canceled`、`context.DeadlineExceeded` 与底层网络错误。`APIError.Verified=false` 的错误体仅作诊断。

## 4. 网关地址

- V3 REST 基地址：生产 `https://openapi.alipay.com`，沙箱 `https://openapi-sandbox.dl.alipaydev.com`。
- 页面/唤端产物网关（`gateway.do`）：生产 `https://openapi.alipay.com/gateway.do`，沙箱 `https://openapi-sandbox.dl.alipaydev.com/gateway.do`。

来源：[沙箱快速接入](https://opendocs.alipay.com/common/097jyh.md)与[V3 基本原则](https://opendocs.alipay.com/open-v3/054oog)。REST 基地址（`Settings.Gateway`）与产物网关（`Settings.ArtifactGateway`）是两项独立配置；环境必须显式选择（`option.WithEnvironment` 成对设置，或分别用 `option.WithGateway` / `option.WithArtifactGateway` 指定），不自动回退或混用。

## 5. 密钥格式

来源：[密钥与证书格式说明](https://opendocs.alipay.com/common/058isw.md)、[自行实现签名](https://opendocs.alipay.com/common/057k53.md)。

- 应用私钥：PKCS#8（`-----BEGIN PRIVATE KEY-----`），RSA ≥2048 位；`utils.ParsePrivateKeyPKCS8` 仅接受该格式，其他格式报错并提示。
- 支付宝公钥：X.509 SubjectPublicKeyInfo（`-----BEGIN PUBLIC KEY-----`）；`utils.ParsePublicKeyPKIX`。
- 证书模式的 SN 计算与携带方式见第 6 节。
- 一个 APPID 只能配置一种加签方式；资金支出类接口官方要求证书模式。

## 6. 证书模式（公钥证书加签）

支付宝在[证书常见问题](https://opendocs.alipay.com/common/02kdne.md)中说明了单张证书 SN 的计算方法。签发机构名称的 RFC2253 格式和根证书 SN 的拼接细节已于 2026-11-05 与官方 SDK 实现核对。

**V3 携带方式**：

- 请求侧：应用证书 SN 写入 Authorization authString 的 `app_cert_sn` 字段；根证书 SN 经独立请求头 `alipay-root-cert-sn` 传递。
- 响应侧按 `alipay-sn` 选择本地可信支付宝证书。缺失或未知 SN 返回 `*core.CertificateError`。轮换过渡期可配置多张有效证书，按响应 SN 选择。

**SN 计算规则**（与 X.509 内置序列号不同，依据官方 FAQ 与 `AntCertificationUtil.getCertSN`/`getRootCertSN`）：

- 单张证书 SN = `MD5(签发机构名称 + 证书内置序列号十进制字符串)`，小写十六进制、补零至 32 位。
- 签发机构名称采用 Java `X500Principal.getName()`（RFC2253）格式：RDN 逆序输出、`,` 分隔、多值 RDN 按原始顺序用 `+` 连接；标准关键字表 CN/C/L/ST/O/OU/STREET/DC/UID，其余属性输出点分 OID（不带 `OID.` 前缀）；值中 `,=+<>#;"\` 加反斜杠转义、空字符输出 `\00`、首尾空格与 `\r` 加反斜杠，其余字符（含中文）原样保留。
- `alipay_root_cert_sn`：对根证书文件中每张证书，仅保留 RSA 家族签名算法（签名算法 OID 前缀 `1.2.840.113549.1.1`，与官方 `getRootCertSN` 一致）的证书，逐张计算 SN 后按文件顺序用 `_` 拼接；不做整体 MD5。
- 官方 Java 实现按平台默认字符集取字节（JDK 18+ 为 UTF-8），本 SDK 固定使用 UTF-8；待沙箱以真实证书复核。
- 本 SDK 的实现经 openssl 独立向量交叉验证（`internal/protocol/certs_test.go`：证书由 openssl 生成、预期 SN 由 openssl + md5 命令行独立计算）。

**通知验签**：官方 `rsaCertCheckV1` 直接使用本地配置的支付宝证书公钥，通知参数不含 SN；本 SDK 证书模式下通知验签使用支付宝证书公钥，调用方式与公钥模式一致（`core/notify.Parse` 契约不变）。

**本地装载校验**（`auth.NewCertificateCredentials`，失败返回错误，不降级为其他认证模式；Client 侧包装为 `*core.ConfigError`，运行期选择失败为 `*core.CertificateError`）：证书可解析且为 RSA 密钥；应用证书、支付宝证书、根证书在装载时刻处于有效期内；应用私钥与应用证书公钥匹配；根证书包按外层签名算法选取 RSA 根，校验其期限、CA 约束和自签名；每张平台验签叶证书通过 Go 的 X.509 路径验证链至配置的根。如果平台证书由中间 CA 签发，将该 CA 一并加入 `AlipayCertificates`，仅叶证书参与响应 SN 选择；根证书集合中至少一张 RSA 签名算法证书（否则无法计算 `alipay_root_cert_sn`）。官方 SDK 初始化不校验链与期限（仅对动态下载的证书校验），本 SDK 收紧为装载时校验，避免无效材料进入运行时。

**运行期有效期检查**覆盖内置凭证中的应用证书、全部支付宝叶证书、中间 CA 和根证书。每次签名、验签及 `Validate` 都检查构造期保存的时间快照。任一配置证书尚未生效或过期即返回 `*auth.CertValidityError`，其中 `Field` 定位配置项，`NotBefore`、`NotAfter` 给出有效期。REST、页面/App 产物与通知入口将运行期证书错误包装为 `*core.CertificateError`，可通过 `errors.As` 沿错误链识别。轮换集合中的旧证书到期同样使该凭证报错，调用方应更新或移除失效材料后重建实例。签名不匹配返回 `*core.SignatureError`。

**证书更新**：更新本地证书文件后用新参数重建 Client 实例（SN 与验签公钥在构造期计算，不支持热更新）；SDK 不实现证书自动下载，未知 SN 由调用方从开放平台获取对应证书后重建实例。

## 7. 页面/唤端产物签名（传统网关契约）

来源：官方 V3 [page.pay](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay)、[wap.pay](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay)、[app.pay](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay) 文档、[页面跳转说明](https://opendocs.alipay.com/common/0i2kxr.md)、[唤端说明](https://opendocs.alipay.com/common/0i2jez.md)与 V3 SDK `GenericExecuteApi`。

**必要性与边界**：官方 V3 的页面/唤端产物（`pageRedirectionData` 表单/URL 与 App 订单串 `orderStr`）明确使用 `gateway.do`、`method`、`version=1.0`、`biz_content`、`sign_type=RSA2` 等字段。这是产物契约要求的固定格式；SDK 据此实现本地签名产物构建（`core.NewArtifactBuilder`），不发 HTTP 请求。

- 系统参数：`method`、`app_id`、`timestamp`（官方实现固定 GMT+8）、`format=json`、`version=1.0`、`charset=UTF-8`、`sign_type=RSA2`；证书模式追加 `app_cert_sn` 与 `alipay_root_cert_sn`；`notify_url` 经产物构建的调用选项注入。
- 业务参数经 `biz_content`（JSON 文本）传递。
- 待签名串：全部参数剔除 `sign` 与空值参数，按参数名 ASCII 升序以 `k=v` 用 `&` 连接；值不做 URL 编码（`biz_content` 含中文/特殊字符时的原文形态依据官方签名规则实现，待沙箱确认）。
- page.pay / wap.pay：GET 语义产出跳转 URL（`gateway.do?` + URL 编码参数，仅值编码）；POST 语义产出 HTML Form（action 携带系统参数，隐藏域承载 `biz_content` 等，属性名值按 HTML 属性语境转义，提交时还原签名原值）。
- app.pay：产出订单串（URL 编码参数串，不含网关地址），交支付宝客户端 SDK 唤起。
- 产物构建使用独立的产物网关配置（见第 4 节），与 V3 REST 基地址分开；跳转完成、表单提交或唤起均不作为支付完成的可信证据，结果以异步通知与交易查询确认。

## 8. 普通支付通知（独立表单验签）

来源：[自行实现验签](https://opendocs.alipay.com/common/02mse7.md)、[异步通知说明](https://opendocs.alipay.com/open/194/103296)（2026-11-05 核验）。通知与 V3 REST 的验签边界独立：REST 按响应头三行原文验签（第 2 节），通知按表单参数剔除规则验签，两者互不复用。

- 通知以 POST 表单发送到 notify_url。参数表（异步通知说明）：notify_time、notify_type、notify_id、sign_type、sign、trade_no、app_id、auth_app_id、out_trade_no、out_biz_no、buyer_id（或 buyer_open_id）、buyer_logon_id、seller_id、seller_email、trade_status、total_amount、receipt_amount、invoice_amount、buyer_pay_amount、point_amount、refund_fee、send_back_fee、subject、body、gmt_create、gmt_payment、gmt_refund、gmt_close、fund_bill_list、voucher_detail_list、biz_settle_mode。
- 待验签参数 = 除 `sign`、`sign_type` 外的全部通知参数；官方原文为"将剩下参数进行 url_decode，然后进行字典排序，组成字符串"：先解码一次，再按参数名 ASCII 升序以 k=v 用 `&` 连接，值不 URL 编码。**空值参数保留**（文档未说明剔除空值，官方语言 SDK 不剔除）。
- 签名值 base64 解码后按 RSA2 验签；公钥与证书模式下通知验签方式相同（证书模式使用支付宝证书公钥，见第 6 节）。
- SDK 额外检查解码后的字段边界。普通支付通知的参数名限于 ASCII 字母、数字与下划线。签名有效但值内出现 `&字段名=` 时，`notify.Parse` 和 `ParseRequest` 返回可用 `errors.Is(err, notify.ErrAmbiguousParameters)` 识别的错误，不返回业务数据。这也会拦下带多个查询参数的某些合法 URL。此时应从商户自己的待核对订单取单号，调用交易查询核实金额、收款方与状态后再入账。
- 这套签名规则无法还原发送方最初的字段划分。下单时应限制所有商户可控的回传文本，使其在通知中解码后也不含上述分隔符。内置支付请求会检查传入的商户订单号、标题、附加信息与公用回传参数。其中页面及 App 支付的 `passback_params` 按契约接收已 URL 编码的值，SDK 也检查其解码后的内容。自定义接口需在下单时执行相同检查。历史订单、其他下单渠道或无法控制的回传字段，应以独立交易查询核实业务数据，尤其不能仅凭可选字段识别买家或入账。
- 生活号通知需在待验签串中保留 `sign_type`，属于不同策略；SDK 仅实现普通支付通知策略（`notify.PolicyTradeStatusSync`，要求 `notify_type=trade_status_sync`），其他通知类型显式拒绝，不自动猜测。
- SDK 要求通知显式携带 `sign_type=RSA2`。缺失、空值或其他算法值均返回 `*core.SignatureError`。通知参数表允许手动验签时省略该参数；本 SDK 的解析入口采用显式算法策略，不提供省略模式。
- `charset` 未出现在官方通知参数表与示例报文中；SDK 若收到非 utf-8 的 charset 值则拒绝（仅支持 UTF-8）。
- `fund_bill_list` / `voucher_detail_list` 是 JSON 文本参数，内部字段名为官方规定的驼峰形式（`fundChannel`、`voucherId`、`otherContributeDetail` 等），与同步响应的下划线字段名（`fund_channel` 等）不同，建模为通知专用 struct，不复用响应模型。
- 商户应答：处理完成后原样输出小写 `success`；未收到 success 时支付宝立即重发 3 次，再按 4m、10m、10m、1h、2h、6h、15h 重发；同一条通知重发时 `notify_id` 不变，可作为幂等依据。
- 官方要求验签通过后必须继续校验：out_trade_no 为商户系统订单、total_amount 与订单一致、seller_id（或 seller_email）为该单据收款方、app_id 为商家本身；过滤重复通知；仅 `TRADE_SUCCESS` / `TRADE_FINISHED` 表示买家付款成功。
- 通知触发条件默认值：TRADE_SUCCESS 触发；TRADE_FINISHED、WAIT_BUYER_PAY、TRADE_CLOSED 默认不触发（随商户产品配置变化）。

## 9. 已知未确认项

1. V3 生产网关对 Authorization 各字段的具体校验容差（timestamp 时钟偏移、nonce 重复策略）未见于公开文档，待沙箱验收确认。
2. 证书模式响应头需用真实沙箱报文验收。SDK 明确要求带签名响应携带 `alipay-sn`、`alipay-timestamp` 和 `alipay-nonce`。
3. 沙箱环境对 V3 证书模式的支持情况未验证；HTTPS 端点已验证连通，尚未使用沙箱凭证联调。
4. 页面/唤端产物的格式与签名字段以官方页面/唤端构建规范为准，不在 openapi.yaml 中；产物 timestamp 固定 GMT+8 是否需调整待沙箱确认；产物中的 `alipay_sdk` 版本标识参数本 SDK 未输出，如沙箱要求再补充。
5. `biz_content` 含中文/特殊字符时在待签名串中的形态（原文 vs 预编码）依据官方"值不 URL 编码"规则实现（与官方语言 SDK 一致），待沙箱验收确认。
6. 通知空值参数处理文档未明确，按官方语言 SDK 行为保留。SDK 要求 `sign_type=RSA2`，实际通知字段仍待沙箱验收。
7. 官方异步通知示例报文未携带参数表中必选的 `app_id`、`auth_app_id` 等字段（示例较旧）；字段以参数表为准建模，沙箱验收时以真实报文复核。
8. SN 计算的签发机构名称固定按 UTF-8 取字节（官方 Java 实现按平台默认字符集，JDK 18+ 即 UTF-8）；待沙箱以真实支付宝证书复核。
9. 本 SDK 在装载时对支付宝公钥证书执行根证书链校验（官方仅对动态下载的证书校验）；多级链的通过和拒绝路径已用本地证书验证，真实应用的证书配置仍待沙箱复核。
