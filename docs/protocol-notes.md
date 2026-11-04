# 协议说明（现有 V2 传统网关实现）

本文件记录既有 V2 代码依据，不表示继续默认选用 V2。目标策略为“有官方 V3 使用 V3，确认只有 V2 才使用 V2”；REST 迁移与页面/唤端的调用方式见[版本清单](api-versions.md)。V3 REST 必须另行核验鉴权、响应和错误契约；有 V3 文档的页面签名产物则按其实际格式处理，不能仅根据目录版本判断报文。

状态：2026-11-05 依据官方文档核验。本文件记录 SDK 实现所依据的协议规则与来源；实现与规则不一致时以官方文档为准并修订本文件。接口覆盖见[覆盖记录](api-coverage.md)。

## 1. 请求签名

来源：[自行实现签名](https://opendocs.alipay.com/common/057k53.md)、[接口加签方式](https://opendocs.alipay.com/common/02kf5p.md)。

- 待签名串 = 全部请求参数（公共参数 + 业务参数），**剔除 sign 参数与空值参数**，按参数名 ASCII 升序以 `k=v` 用 `&` 连接；值不做整体 URL 编码。
- `sign_type` 参与签名（不剔除）；RSA2 对应 SHA256WithRSA，签名值 base64 编码，密钥 ≥2048 位。
- `biz_content` 作为单个参数整体参与签名，JSON 内部不展开、不重排。
- timestamp 格式 `yyyy-MM-dd HH:mm:ss`；时区要求文档未明确（官方 SDK 使用服务器本地时间）。
- 发送：POST `application/x-www-form-urlencoded` 到网关；签名与传输使用同一份参数值，不在签名后重新编码业务 JSON。
- 官方文档示例将公共参数置于 URL query、biz_content 置于表单；官方语言 SDK 实现统一以 form body 发送全部参数。本 SDK 采用后者，沙箱验收时确认。

## 2. 公共请求参数

来源：各接口文档公共参数表（如 [02ekfq](https://opendocs.alipay.com/open/02ekfq.md)）。

| 参数 | 必选性 | 本 SDK 取值 |
| --- | --- | --- |
| app_id | 必选 | 配置注入 |
| method | 必选 | OperationSpec.Method |
| format | 可选 | 固定 JSON |
| charset | 必选 | 固定 utf-8 |
| sign_type | 必选 | RSA2 |
| sign | 必选 | 请求时计算 |
| timestamp | 必选 | 每次请求生成 |
| version | 必选 | OperationSpec.Version（本 SDK 接口均为 1.0） |
| notify_url | 可选 | CallOption（请求级） |
| app_auth_token | 可选 | CallOption（请求级，代调用） |
| biz_content | 必选（本组接口） | 编码器输出 |

`notify_url`、`app_auth_token` 属于调用级参数，不进入 Client 共享配置，避免并发调用串用。

## 3. 同步响应验签

来源：[自行实现验签](https://opendocs.alipay.com/common/02mse7.md)。

- 响应结构：`{"<method>_response":{...},"sign":"..."}`；节点名为方法名中 `.` 替换为 `_`。
- 待验签原文 = `xxx_response` 节点的**原始 JSON 字节**：必须保留首尾大括号、双引号、原始转义（如 `\/`）与内部空白；不重新序列化、不压缩、不重排。
- 验签算法：RSA2（SHA256WithRSA），sign 值 base64 解码后校验。
- 成功与业务失败（code≠10000）的业务节点均带 sign；缺失 sign 或验签失败一律按 `*core.SignatureError` 处理，不产生可信业务数据。
- `error_response`（网关层拒绝，如 isv.invalid-signature）在验签文档中未说明是否签名；本 SDK 一律按**未验签错误**处理：返回 `*core.APIError{Verified:false}`，字段仅作诊断。
- 提取实现（`internal/protocol`）使用 `encoding/json/jsontext` 流式读取：拒绝重复的响应节点/ sign、拒绝非对象顶层结构与尾部内容；`ReadValue` 返回的字节引用解码器内部缓冲，必须立即复制以保留原文。

## 4. 错误模型

来源：[公共错误码](https://opendocs.alipay.com/common/02km9f.md)。

- 业务节点与 `error_response` 均含 `code`/`msg`，业务错误含 `sub_code`/`sub_msg`。
- 公共 code：10000 接口调用成功；20000 服务不可用；20001 授权权限不足；40001 缺少必选参数；40002 非法参数；40003 条件异常；40004 业务处理失败；40005 频次超限；40006 权限不足。
- code 不在 OperationSpec.AcceptedCodes 内时返回 `*core.APIError`（已验签时 Verified=true）；HTTP 200 不等于业务成功，也不等于全部子状态成功。
- 错误链：`ConfigError`/`ValidationError`/`TransportError`/`HTTPError`/`ProtocolError`/`SignatureError`/`APIError`，通过 `errors.AsType[T]` 匹配；`TransportError` 保留 `context.Canceled`、`context.DeadlineExceeded` 与底层网络错误。

## 5. 网关地址

来源：[沙箱快速接入](https://opendocs.alipay.com/common/097jyh.md)。

- 生产：`https://openapi.alipay.com/gateway.do`
- 沙箱：`https://openapi-sandbox.dl.alipaydev.com/gateway.do`

环境必须显式选择（`option.WithEnvironment` 或 `option.WithGateway`），不自动回退或混用。

## 6. 密钥格式

来源：[密钥与证书格式说明](https://opendocs.alipay.com/common/058isw.md)、[自行实现签名](https://opendocs.alipay.com/common/057k53.md)。

- 应用私钥：PKCS#8（`-----BEGIN PRIVATE KEY-----`），RSA ≥2048 位；`utils.ParsePrivateKeyPKCS8` 仅接受该格式，其他格式报错并提示。
- 支付宝公钥：X.509 SubjectPublicKeyInfo（`-----BEGIN PUBLIC KEY-----`）；`utils.ParsePublicKeyPKIX`。
- 证书模式（app_cert_sn / alipay_root_cert_sn / SN 选择）为 P5 范围，尚未实现。
- 一个 APPID 只能配置一种加签方式；资金支出类接口官方要求证书模式。

## 7. 接口契约要点（实现依据）

各接口完整字段表记录于核验底稿（P0），要点如下：

- **alipay.trade.query**（[02ekfq](https://opendocs.alipay.com/open/02ekfq.md)）：out_trade_no / trade_no 二选一，同传优先 trade_no；trade_status 枚举 WAIT_BUYER_PAY / TRADE_CLOSED / TRADE_SUCCESS / TRADE_FINISHED；trade_no、buyer_logon_id 在未生成真实交易时不返回；fund_bill_list 标注必选但注明仅在签约指定或 query_options 指定时返回，建模为可选。
- **alipay.trade.precreate**（[02np92](https://opendocs.alipay.com/open/02np92)，页面更新时间 2026-08-17）：必选仅 out_trade_no（字母/数字/下划线）、total_amount（[0.01,1e8]，两位小数）、subject（禁用 / = &）；二维码有效期固定 2 小时；time_expire 优先于 timeout_express；响应 out_trade_no 必选，qr_code/share_code/prepay_id 可选。
- **alipay.trade.refund**（[02ekfs](https://opendocs.alipay.com/open/02ekfs.md)）：out_request_no 部分退款必选，交易号下唯一，重试不得变更；同笔交易退款间隔 ≥3s；累计退款不超原交易金额；code=10000 仅代表请求成功，fund_change=Y 才为退款成功，N 或未返回需退款查询确认；文档参数表**不含** refund_currency（不得凭其他 SDK 记忆添加）。
- **alipay.trade.fastpay.refund.query**（[02ekft](https://opendocs.alipay.com/open/02ekft.md)）：out_request_no 必选；code=10000 仅代表查询成功；refund_status 仅有 REFUND_SUCCESS，未返回即未收到或失败；建议退款后 10s 以上再查询。
- **alipay.trade.close**（[02ivbu](https://opendocs.alipay.com/open/02ivbu.md)）：trade_no（≥16 位）/ out_trade_no 二选一，同传以 trade_no 为准；仅等待买家付款状态可关单；响应仅 trade_no/out_trade_no。

## 8. 异步通知（普通支付通知）

来源：[自行实现验签](https://opendocs.alipay.com/common/02mse7.md)、[异步通知说明](https://opendocs.alipay.com/open/194/103296)（2026-11-05 核验）。

- 通知以 POST 表单发送到 notify_url。参数表（异步通知说明）：notify_time、notify_type、notify_id、sign_type、sign、trade_no、app_id、auth_app_id、out_trade_no、out_biz_no、buyer_id（或 buyer_open_id）、buyer_logon_id、seller_id、seller_email、trade_status、total_amount、receipt_amount、invoice_amount、buyer_pay_amount、point_amount、refund_fee、send_back_fee、subject、body、gmt_create、gmt_payment、gmt_refund、gmt_close、fund_bill_list、voucher_detail_list、biz_settle_mode。
- 待验签参数 = 除 `sign`、`sign_type` 外的全部通知参数；官方原文为"将剩下参数进行 url_decode，然后进行字典排序，组成字符串"：先解码一次，再按参数名 ASCII 升序以 k=v 用 `&` 连接，值不 URL 编码。**空值参数保留**（文档未说明剔除空值，官方语言 SDK 不剔除），与请求待签名串剔除空值的规则不同。
- 签名值 base64 解码后按 RSA2 验签；公钥与证书模式下通知验签方式相同（证书模式整体属 P5）。
- 生活号通知需在待验签串中保留 `sign_type`，属于不同策略；SDK 仅实现普通支付通知策略（`notify.PolicyTradeStatusSync`，要求 `notify_type=trade_status_sync`），其他通知类型显式拒绝，不自动猜测。
- `sign_type` 出现其他算法值（如 RSA1）时拒绝验签；缺失时按验签器算法（RSA2）验签——文档参数表注明"如果开发者手动验签，不使用 SDK 验签，可以不传此参数"。
- `charset` 未出现在官方通知参数表与示例报文中；SDK 若收到非 utf-8 的 charset 值则拒绝（仅支持 UTF-8）。
- `fund_bill_list` / `voucher_detail_list` 是 JSON 文本参数，内部字段名为官方规定的驼峰形式（`fundChannel`、`voucherId`、`otherContributeDetail` 等），与同步响应的下划线字段名（`fund_channel` 等）不同，建模为通知专用 struct，不复用响应模型。
- 商户应答：处理完成后原样输出小写 `success`；未收到 success 时支付宝立即重发 3 次，再按 4m、10m、10m、1h、2h、6h、15h 重发；同一条通知重发时 `notify_id` 不变，可作为幂等依据。
- 官方要求验签通过后必须继续校验：out_trade_no 为商户系统订单、total_amount 与订单一致、seller_id（或 seller_email）为该单据收款方、app_id 为商家本身；过滤重复通知；仅 `TRADE_SUCCESS` / `TRADE_FINISHED` 表示买家付款成功。
- 通知触发条件默认值：TRADE_SUCCESS 触发；TRADE_FINISHED、WAIT_BUYER_PAY、TRADE_CLOSED 默认不触发（随商户产品配置变化）。

## 9. 已知未确认项

1. biz_content 含中文/特殊字符时，官方示例中该值在待签名串里呈现为原文还是预编码形式存在转述分歧；本 SDK 按"待签名串值不做 URL 编码"实现（与官方语言 SDK 一致），待沙箱验收确认。
2. `error_response` 是否签名（见第 3 节，按未验签处理）。
3. 退款查询 `deposit_back_info_list` 对应的 query_options 取值原文不一致（枚举仅列 `deposit_back_info`）。
4. 异步通知验签的参数剔除规则已按验签文档实现：普通支付通知剔除 `sign` 与 `sign_type`、保留空值参数；生活号策略未实现。空值参数处理文档未明确，按官方语言 SDK 行为保留，待沙箱验收确认。
5. timestamp 时区要求未明确，SDK 使用服务器本地时间并建议校准服务器时钟。
6. `sign_type` 缺失时按验签器算法（RSA2）验签，依据通知参数表"手动验签可以不传此参数"的表述；出现其他算法值一律拒绝。待沙箱验收确认实际通知是否始终携带 `sign_type`。
7. 官方异步通知示例报文未携带参数表中必选的 `app_id`、`auth_app_id` 等字段（示例较旧）；字段以参数表为准建模，沙箱验收时以真实报文复核。
