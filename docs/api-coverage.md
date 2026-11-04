# 接口覆盖记录

状态：随实现更新。日期：2026-11-05。协议依据的详细记录见[协议说明](protocol-notes.md)。

覆盖表按计划要求区分"文档核验 / 已实现 / 离线测试 / 沙箱验证"，不以单一"已支持"列合并。

**选型已调整为 V3 优先。以下 5 个交易接口当前实现均为 V2，目标全部迁移到官方已有 V3；不能将表中的“已实现”理解成 V3 已实现。** REST 路由、页面/唤端方式及证据见[版本清单](api-versions.md)。

## 覆盖表

| 接口名 | 当前协议/参数版本 | 具体官方链接 | 文档核验日期 | 请求编码 | 支持授权模式 | 实现状态 | 离线测试 | 沙箱结果 | 限制 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| alipay.trade.query | V2（version=1.0） | [02ekfq](https://opendocs.alipay.com/open/02ekfq.md) | 2026-11-05 | biz_content JSON | 直连；app_auth_token 代调用 | 已实现 | 已离线验证（httptest 端到端 + mock） | 未验证 | 查询成功≠支付成功；query_options 部分扩展响应结构本页未定义 |
| alipay.trade.precreate | V2（version=1.0） | [02np92](https://opendocs.alipay.com/open/02np92) | 2026-11-05 | biz_content JSON | 直连；app_auth_token 代调用 | 已实现 | 已离线验证（httptest 端到端 + mock） | 未验证 | 二维码有效期固定 2 小时；同一订单号不能重复申请；subject 禁用 /=& |
| alipay.trade.refund | V2（version=1.0） | [02ekfs](https://opendocs.alipay.com/open/02ekfs.md) | 2026-11-05 | biz_content JSON | 直连；app_auth_token 代调用 | 已实现 | 已离线验证（mock；错误码语义见下） | 未验证 | code=10000≠退款成功，以 fund_change 为准；SDK 不自动重试 |
| alipay.trade.fastpay.refund.query | V2（version=1.0） | [02ekft](https://opendocs.alipay.com/open/02ekft.md) | 2026-11-05 | biz_content JSON | 直连；app_auth_token 代调用 | 已实现 | 已离线验证（mock） | 未验证 | code=10000 仅代表查询成功；refund_status 仅 REFUND_SUCCESS 一个枚举值 |
| alipay.trade.close | V2（version=1.0） | [02ivbu](https://opendocs.alipay.com/open/02ivbu.md) | 2026-11-05 | biz_content JSON | 直连；app_auth_token 代调用 | 已实现 | 已离线验证（mock） | 未验证 | 仅 WAIT_BUYER_PAY 状态可关单；SDK 不自动关单 |
| 支付异步通知（普通支付通知） | - | [103296](https://opendocs.alipay.com/open/194/103296)（参数表）、[02mse7](https://opendocs.alipay.com/common/02mse7.md)（验签规则） | 2026-11-05 | 表单 | 直连 | 已实现 | 已离线验证（签名链路端到端 + 特殊字符/错误分支/日志脱敏） | 未验证 | 仅普通支付通知策略（剔除 sign、sign_type，保留空值）；生活号等策略未实现；success/fail 应答与业务校验由调用方完成 |

## 协议能力覆盖

| 能力 | 状态 | 说明 |
| --- | --- | --- |
| RSA2 签名（PKCS#8 私钥） | 已实现 | SHA256WithRSA，签名值 base64；openssl 独立向量交叉验证 |
| RSA2 同步响应验签 | 已实现 | 仅对 xxx_response 节点原始字节验签，不重新序列化 |
| 公钥模式 | 已实现 | 应用私钥 + 支付宝公钥（PKIX PEM） |
| V3 HTTP/鉴权/响应与交易路由 | 未实现（下一优先里程碑） | 5 个交易 API 与账单 GET 示例应迁移到 V3；不自动回退到 V2 |
| 证书模式 | 未实现（P5） | app_cert_sn / alipay_root_cert_sn 规则待实现 |
| 内容加密 | 未实现（P8） | 请求该能力时当前无入口；不将签名私钥用作加密密钥 |
| UTF-8 JSON | 已支持 | charset 固定 utf-8、format 固定 JSON |
| GBK / XML / RSA(v1) 签名 | 不支持 | 首期明确列为未支持 |
| 页面/App 支付产物构建 | 未实现（P7） | 已核验官方 V3 文档与 pageExecute/sdkExecute 方式；不归类为无 V3 支持 |
| 自定义接口扩展（JSON） | 已实现 | 见 examples/customapi（真实接口示例） |
| 自定义接口扩展（直接参数） | 已实现 | 机制经外部测试包验证；`custom.demo.form.api` 为占位方法名，非真实接口 |

## 测试证据索引

- 密码学：`core/auth/signers/rsa2_test.go`（openssl 已知答案向量、篡改/错钥/错签失败）。
- 协议：`internal/protocol/protocol_test.go`（排序签名原文、原始节点提取、重复节点拒绝）。
- Client：`core/client_test.go`（httptest 服务端独立核对签名、响应原文验签、篡改检测、error_response 未验签标记、业务错误分类、超时/取消/超限/重定向、日志级别与脱敏、并发隔离）。
- 交易 Service：`services/trade/api_trade_test.go`（接口名/版本/节点名、字段序列化、金额字符串透传、枚举未知值保留、校验规则、端到端业务错误）。
- 通知：`core/notify/notify_test.go`（单次解码与 `+`/`%2B`/`%26`/中文/空值一致性、重复参数拒绝、缺签/错签/错公钥/错 sign_type、未支持策略与 notify_type 拒绝、请求体大小上限、日志脱敏与级别、共享验签器并发）；`services/trade/api_notify_test.go`（官方参数表映射、必选参数校验、未知枚举保留、端到端签名链路）；`internal/protocol/protocol_test.go`（通知待验签串构造规则）。
- 通知示例：`services/trade/api_notify_example_test.go`（业务校验与幂等接入位置）；`examples/tradenotify`（最小 net/http 接入）。
- 扩展：`examples/customapi/customapi_test.go`（外部包仅依赖公开包；错签时不调用用户解码器；受保护参数拦截）。

## 未确认项

- precreate 文档 URL（02np92）为浏览器渲染抓取，`.md` 端点多次返回空正文；字段以页面原文为准，沙箱验收时复核。
- 退款查询的 `deposit_back_info_list` 要求的 query_options 取值与枚举列表存在原文不一致（枚举仅列 `deposit_back_info`），已按原文建模并标注。
- 退款查询间隔建议在两处文档不一致（5s / 10s），GoDoc 取更保守的 10s。
- biz_content 含中文时在待签名串中的形态（原文 vs 预编码）以官方示例存在转述分歧，本 SDK 按官方规则"值不 URL 编码"实现，待沙箱确认。
- `error_response` 是否携带签名在验签文档中未明确；本 SDK 一律按未验签错误处理（`APIError.Verified=false`），不作为可信数据。
- 异步通知的空值参数处理文档未明确，按官方语言 SDK 行为保留在待验签串中；`sign_type` 缺失时按 RSA2 验签（依据通知参数表"手动验签可以不传此参数"）；官方示例报文未携带参数表必选的 `app_id` 等字段（示例较旧）。以上均待沙箱验收时以真实报文复核。
