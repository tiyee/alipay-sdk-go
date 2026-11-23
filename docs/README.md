# 开发文档

本目录记录 SDK 当前设计、支持范围和后续工作。SDK 首版尚未发布，仅支持支付宝 V3；当前提供官方快照中全部 746 个 REST 操作入口，其中 68 个属于非 `alipay.*` 命名空间。生成接口尚未逐项核验业务规则或进行沙箱验收。接口能力与限制以[覆盖记录](api-coverage.md)为准；沙箱验收步骤与未确认项见[沙箱验收指引](sandbox-verification.md)。

| 文档 | 内容 |
| --- | --- |
| [整体架构](architecture.md) | 系统边界、目录、依赖、公开 API、V3 REST 协议流程、页面产物与通知、错误和安全设计 |
| [开发计划](development-plan.md) | 阶段任务、依赖、接口优先级、验收标准、工期估算和待核实项 |
| [交易场景说明](trade-scenarios.md) | 预下单、支付通知、查询确认、关单、退款与退款查询的接入顺序与商户校验清单 |
| [自定义接口扩展](custom-apis.md) | 用户定义请求/响应、枚举、Operation 与独立 Service，复用 SDK 执行能力 |
| [支付宝接口目录](apis.md) | 官方 V3 REST 接口快照（746 项）及已核验的页面和唤端接口 |
| [支付宝业务包目录](service-layout.md) | 各业务子包、调用迁移路径与文件拆分约定 |
| [接口版本说明](api-versions.md) | V3-only 支持范围、接口覆盖统计、拆包布局及快照依据 |
| [Credentials 认证接口方案](credentials-design.md) | Credentials 作为 NewClient 首参，公钥/证书的 Signer、Verifier、能力校验与实施记录 |
| [接口使用示例](api-examples.md) | 各接口完整平台证书接入示例、配置、类型化请求及编译与调用方式 |
| [沙箱验收指引](sandbox-verification.md) | 沙箱验收原则、前置条件、用例矩阵与当前未验证记录 |
| [CHANGELOG](../CHANGELOG.md) | 当前候选能力与已知限制 |

建议先阅读整体架构，再按需查阅其余文档。状态以代码、测试和验收证据为准，不能仅凭计划中的勾选宣布接口可用。

## 规划基线

- 编写日期：2026-11-05。
- 模块路径为 `github.com/tiyee/alipay-sdk-go`，与当前 Git remote 一致。
- 项目统一使用 alipay 命名，采用 `core/`、`services/`、`utils/` 分层，协议实现依据支付宝官方文档。
- 数据模型优先使用具名 struct，通用调用、编解码和通知结果通过泛型扩展；枚举采用具名类型与带类型的常量。
- 未内置 API 可由用户在独立包中实现，使用公开的泛型 Operation、编解码器及 Client，不需要修改 SDK。
- 日志统一采用标准库 `log/slog`，支持任意 `slog.Handler` 实现注入，未配置时保持静默。
- SDK 仅支持支付宝 V3：服务端接口使用 V3 REST；页面/唤端按官方 V3 文档与产物构建规范提供本地签名产物；普通支付通知按自身契约处理。未确认 V3 契约的接口不予实现。
- 模块基线为 `go 1.24.0`，CI 验证各支持系列及最新稳定版。
- 版本相关实现通过 Go 编译标签选择，具体规则见[整体架构](architecture.md)。
- V3 Client、Service、离线测试与 CI 已就位。真实网关调用与沙箱验收尚未实施，验收安排见沙箱验收指引。

## 官方资料入口

先通过 [索引版](https://opendocs.alipay.com/common/llms.txt) 定位，再参考 [摘要版](https://opendocs.alipay.com/common/llms-full.txt) 和具体文档全文。两个汇总文件均不能替代目标 API 的完整参数表；本次查阅遇到的文档可访问性问题见开发计划。
