# Credentials 认证凭证接口实现方案

状态：实现方案，尚未落地；本文不表示已支持证书模式或 V3。日期：2026-11-05。基线为 Go 1.27，关联[整体架构](architecture.md)、[开发计划](development-plan.md)和[接口版本选型](api-versions.md)。

## 1. 设计结论与当前问题

**将 `Credentials` 接口作为 `NewClient` 的第一个必选参数。** 公钥与证书模式分别实现此接口，提供各自的 Signer、Verifier 与凭证校验能力。Client 复用同一套执行流程，泛型继续用于 `Operation[Req, Resp]`、请求/响应、编解码和通知结果。

```go
func NewClient(credentials auth.Credentials, opts ...ClientOption) (*Client, error)
```

当前 `core.Settings` 同时保存 Signer、Verifier 和 AuthMode，`option.WithVerifier` 还会将模式设置为公钥，存在认证选项相互覆盖的设计问题。本方案把这些配置集中到凭证构造器，公共选项只管理 APPID、环境、HTTP、超时和 slog 等资源。

本方案采用普通的、可由外部包实现的 Credentials 接口，不使用封闭的类型集合，也不把 Client 改为泛型类型。两种模式的 Client 都是 `*core.Client`，因此模式约束在构造与执行阶段检查，不能宣称基础 Client 已实现编译期模式隔离。

## 2. 包布局与依赖

```text
core/auth/
  credentials.go             # Credentials、CertificateProvider 与公共类型
  public_key.go              # 公钥配置、凭证工厂及组合策略
  certificate.go             # 证书配置、凭证工厂及组合策略（P5）
  identity.go                # 证书身份标识
  resolver.go                # 本地可信证书查找（P5）
  signers/                   # RSA2 等实际签名实现
  verifiers/                 # 固定公钥与证书选择验签实现
core/
  settings.go                # 公共配置与认证快照绑定
  client.go                  # 统一 Client 与普通 Do
  execute.go                 # 泛型 Execute 及泛型方法委托
core/option/                 # 公共配置选项，不设置认证模式
core/notify/                 # Credentials 驱动的通知 Handler
internal/protocol/           # 协议原文、报文封装与标识提取
```

按阶段创建实际文件，不生成空包。依赖方向为 `core -> core/auth`；`core/auth` 不反向导入 `core`。认证层返回自身错误或标准错误链，由 Client 包装为配置、签名、证书或验签错误。

## 3. Credentials 与算法接口

以下为目标 API 草案，尚不能作为现有 SDK 的可编译使用示例。

```go
// core/auth 包
type Mode uint8

const (
	ModeUnset Mode = iota
	ModePublicKey
	ModeCertificate
)

type SignatureAlgorithm string

const SignatureAlgorithmRSA2 SignatureAlgorithm = "RSA2"

type Signature struct {
	Algorithm SignatureAlgorithm
	Value     string
}

type VerificationInput struct {
	Payload   []byte
	Signature Signature
	KeyID     string
}

type Signer interface {
	Algorithm() SignatureAlgorithm
	Sign(ctx context.Context, payload []byte) (Signature, error)
}

type Verifier interface {
	Verify(ctx context.Context, input VerificationInput) error
}

type Credentials interface {
	Mode() Mode
	Signer() Signer
	Verifier() Verifier
	Validate() error
}
```

职责划分：

- `Credentials` 组合并提供认证对象，本身不构造 HTTP 请求，也不解释交易状态。
- `Signer` 只处理协议模块已经准备好的签名原文；两种模式均可复用 RSA2 签名器，不为模式差异重复实现相同算法。
- `Verifier` 实现模式差异：公钥验签器固定使用可信支付宝公钥；证书验签器根据标识选择本地可信证书，再执行对应校验与验签。
- `Mode()` 是具体实现派生的只读信息，不是调用方可设置的策略切换开关。
- `Validate()` 检查已装载状态和材料一致性，不隐式下载证书或调用远程接口。

`VerificationInput.KeyID` 来自协议模块提取的响应或通知标识，仅用于查找已信任材料，不作为远端下载地址或任意公钥。是否必需、取自哪个字段或请求头，由具体协议明确规定。

Signer/Verifier 的算法检查必须与协议一致，不能因为签名结果包含算法字段就接受任意算法。签名值、原文、身份凭证和业务报文不进入默认日志。

## 4. 公钥与证书实现

### 4.1 独立配置和构造函数

```go
type PublicKeyConfig struct {
	AppSigner       crypto.Signer
	AlipayPublicKey *rsa.PublicKey
}

type CertificateConfig struct {
	AppSigner          crypto.Signer
	AppCertificate     *x509.Certificate
	AlipayCertificates []*x509.Certificate
	RootCertificates   []*x509.Certificate
}

func NewPublicKeyCredentials(config PublicKeyConfig) (*PublicKeyCredentials, error)
func NewCertificateCredentials(config CertificateConfig) (*CertificateCredentials, error)
```

构造后凭证的内部字段不导出；公钥配置没有证书字段，证书配置不提供“失败后改用公钥”的备用项。应用私钥可由 `utils` 解析为标准库类型，再作为 `crypto.Signer` 传入；其他符合该契约的签名实现也可通过适配器接入。

标准库 `crypto.Signer` 适配为本方案的 Signer。构造时检查其 Public 返回的是符合 RSA2 要求的 RSA 公钥，并在证书模式校验与应用证书的匹配关系。标准库 Signer 本身没有 context 取消保证，不能承诺任意底层签名操作都可中断。

两种凭证的组合如下：

| 实现 | Mode | Signer | Verifier |
| --- | --- | --- | --- |
| PublicKeyCredentials | ModePublicKey | 应用 RSA2 签名器或合规适配器 | 使用固定可信支付宝公钥的验签器 |
| CertificateCredentials | ModeCertificate | 与应用证书匹配的应用签名器 | 使用本地可信证书索引的验签器 |

### 4.2 证书身份能力

证书请求还需要身份元数据，单独定义能力接口，避免要求公钥实现返回无意义的证书配置：

```go
type CertificateIdentity struct {
	AppCertificateSN  string
	RootCertificateSN string
}

type CertificateProvider interface {
	Credentials
	CertificateIdentity() CertificateIdentity
}
```

CertificateCredentials 实现 CertificateProvider；PublicKeyCredentials 只实现 Credentials。Client 构造时若得到 ModeCertificate，必须同时具备证书身份能力、通过验证并提供完整材料；模式与能力不一致则失败，不自动推断或修正模式。

身份数据只是已核验的只读标识，不包含私钥或可信材料的可变引用。标识计算依据[官方证书说明](https://opendocs.alipay.com/common/02kdne.md)，签名依据[官方规范](https://opendocs.alipay.com/common/057k53.md)。V2/V3 分别决定元数据的映射方式，不把 V2 参数名原样套入 V3。

### 4.3 可变数据与第三方实现

复制公钥和证书的可变输入；证书可基于独立 Raw 数据重新解析为内部快照。外部 Signer/Verifier 需要保证并发安全和状态稳定，不声称可以深拷贝任意硬件或远程实现。

Credentials 允许外部包实现，便于自定义签名、凭证加载和测试。此类实现属于调用方信任的扩展；SDK 校验非空、算法、模式及能力一致性，不能仅凭接口实现就证明第三方内部确实执行了正确密码学操作。证书模式的默认实现必须执行完整的可信材料检查。

## 5. NewClient 构造与公共配置

公开入口保持普通 Client，并把 Credentials 放在第一个参数：

```go
// core 包
func NewClient(credentials auth.Credentials, opts ...ClientOption) (*Client, error)
```

构造顺序：

1. 拒绝 nil 和 typed nil Credentials，调用其 Validate。
2. 获取 Mode、Signer、Verifier，拒绝未定义模式、nil/typed nil 算法对象和算法配置不一致。
3. 证书模式读取 CertificateProvider 的身份副本，检查模式与能力一致性；不下载证书。
4. 处理公共选项并校验 APPID、环境、HTTP、超时等配置。
5. 绑定只读认证快照与公共资源，成功后才返回 Client。

Client 保存已经校验的模式、Signer、Verifier 和证书身份快照，执行时不反复调用 Credentials getter 选择另一套对象。签名器和验签器可以按自身契约实现并发安全的内部机制，但不能在调用方原地修改配置后悄悄切换认证模式。

目标 Settings/ClientOption 不再含可设置的 Signer、Verifier、AuthMode 或 Credentials；公共选项无法覆盖首参确定的认证策略。`Settings()` 不返回凭证或签名器，读取快照不写回 Client 字段。

典型调用形态：

```go
credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
	AppSigner:       appPrivateKey,
	AlipayPublicKey: alipayPublicKey,
})
if err != nil {
	return err
}

client, err := core.NewClient(credentials,
	option.WithAppID(appID),
	option.WithEnvironment(core.EnvironmentSandbox),
	option.WithLogHandler(handler),
)
```

切换为证书配置时，只替换凭证构造过程，业务 Service 与公共配置无需改成另一套类型。构造器不修改全局 HTTP 或 slog 状态，默认使用 `slog.DiscardHandler`。

## 6. 调用流程与能力校验

请求链路：校验 Operation 的认证要求 → 协议层构造参数、身份与签名原文 → 快照 Signer 签名 → HTTP 发送 → 协议层提取原始响应、签名和标识 → 快照 Verifier 验签 → 按需解密 → 泛型业务解码。

- 公钥验签器只使用配置的可信公钥。
- 证书验签器按协议标识查找本地可信证书，校验对应信任、时效与标识；未知证书失败，不降级到公钥或无验签。
- 签名原文、V3 REST 鉴权、传统格式页面产物与通知解析由对应协议/调用方式负责，Credentials 不猜测 API 的传输方式。
- 认证模式、V2/V3 协议和 REST/页面/App 调用方式分别建模，遵守既有 V3 优先规则。

公共 Client 类型本身不区分公钥与证书，因此必须证书的接口依靠执行前校验。目标 OperationSpec 增加具名 AuthRequirement：

```go
type AuthRequirement uint8

const (
	AuthRequirementAny AuthRequirement = iota + 1
	AuthRequirementCertificate
)
```

内置 Operation 按官方契约明确填写要求，Client.Do 在发请求前检查认证快照是否满足。自定义 Operation 同样声明要求，构造时拒绝无效枚举；不能通过通用 Executor 入口省略已声明的校验。

SDK 不能证明任意自定义 Operation 声明都符合尚未收录的接口契约，最终权限仍由支付宝判断。若以后需要编译期的证书专用入口，可另设计具名能力封装；不把当前普通 Client 方案宣称为已有静态隔离。

## 7. Service、泛型与通知

普通领域 Service 和用户自己的 Service 继续接收 `core.Executor`，不持有 Signer/Verifier，也不重复证书选择流程。Client 的普通 Do 仍实现 Executor，泛型方法继续委托共享流程：

```go
func (c *Client) Execute[Req, Resp any](ctx context.Context,
	operation Operation[Req, Resp], req Request[Req]) (*Response[Resp], error) {
	return Execute(ctx, c, operation, req)
}
```

泛型用于请求/响应配对、编码器、解码器和通知数据，而不用于选择认证策略。已知数据保持具名 struct、嵌套具体类型和具名枚举，协议边界才使用有明确值类型的参数集合。

通知建议增加普通 Handler，以 Credentials 作为构造首参，复用模式对应的 Verifier：

```go
func NewHandler(credentials auth.Credentials,
	policy Policy, opts ...ParseOption) (*Handler, error)

func (h *Handler) ParseRequest[T any](ctx context.Context,
	req *http.Request, decoder Decoder[T]) (*Notification[T], error)
```

通知 Handler 绑定自身验证所需的快照，不依赖 Client 的私有 HTTP 执行逻辑。通知策略和认证模式分开；证书标识来自该通知类型的实际契约，缺少必要信息时明确失败，不盲试未确认密钥。原有 Parse 函数与领域通知示例需配套迁移。

## 8. 迁移与实施顺序

这是明确的公开 API 变更：NewClient 增加必选首参，认证选项迁移到凭证工厂，Signer/Verifier 接口增加上下文和确定的验证输入。Client 仍是普通类型，领域 Executor 与请求/响应泛型保持职责不变。

| 当前入口 | 目标入口 |
| --- | --- |
| `core.NewClient(opts...)` | `core.NewClient(credentials, opts...)` |
| `option.WithPrivateKey` | 将解析后的应用密钥传入凭证配置 AppSigner |
| `option.WithAlipayPublicKey` | PublicKeyConfig 的 AlipayPublicKey |
| `option.WithSigner` / `WithVerifier` | 凭证实现负责组合，公共 ClientOption 不再设置 |
| `Settings.AuthMode` | Credentials 派生并绑定的只读模式信息 |
| 直接给通知函数传 Verifier | Credentials 首参的 Handler 与类型化 Decoder |

建议在 V3 迁移和首次稳定发布前统一调整，提供调用迁移示例，不宣称原构造调用可无感兼容。旧自定义签名器若只实现旧接口，应增加适配器或按新接口迁移，不能用无类型参数猜测旧/新配置。

实施步骤：

1. **C1：接口与公钥迁移。** 定义 Credentials/Signer/Verifier 新契约，实现公钥凭证和适配器，调整 NewClient 首参与公共配置，更新全部构造用例与扩展示例，验证现有公钥行为等价。
2. **C2：衔接 V3。** 将签名器、验签器与认证身份接入独立 V3 处理器，完成既有交易与账单示例迁移，不跨协议或跨模式降级。
3. **C3：P5 证书模式。** 实现 CertificateCredentials、CertificateProvider、证书校验及可信索引，启用能力检查与对应通知场景，再更新覆盖状态。
4. **C4：整理交付。** 完成第三方 Credentials 示例、两种模式的行为一致性与错误用例、通知迁移、证书更新说明及版本策略。

C1 可与 V3 里程碑合并实施；C3 在 P5 完成。首期本地换证采用重建 Credentials、Client 与 Handler，不引入全局下载器和隐式更新协程；正在运行的请求使用各自的认证快照。

## 9. 验证与验收

| 方面 | 验证要求 |
| --- | --- |
| 接口实现 | 公钥与证书凭证满足 Credentials，外部包可以实现自定义 Credentials |
| 配置隔离 | 两套配置工厂不能混传，公共选项不能覆盖模式、Signer 或 Verifier |
| 构造边界 | nil、typed nil、缺失算法对象、未知模式、模式/证书能力矛盾均在构造阶段失败 |
| 快照稳定 | Getter 不被执行流程反复读取，外部可变输入不污染已构造实例 |
| 公钥回归 | 原有签名、验签、ctx、日志、网络错误和资源关闭用例继续通过 |
| 证书回归 | 错根、错角色、私钥不匹配、过期、未知标识、多证书选择均有独立用例 |
| 能力要求 | 证书要求在统一 Do 路径校验，不因 Executor 或自定义 API 绕过 |
| 泛型数据流 | Operation 类型配对、外部编解码器和非泛型 Executor 保持可用 |
| 通知与并发 | 模式对应的 Verifier、通知策略、共享快照、脱敏和竞态检测 |

使用外部测试包验证公开 Credentials 扩展；行为测试使用专用虚构密钥、独立签名向量和 httptest。实现后执行格式、`go test ./...`、`go vet ./...` 与 `go test -race ./...`。沙箱按模式和协议分别记录，没有条件时明确未验证。

编写方案时只验证接口与调用形态，不把文档或类型编译检查当作已完成密码学实现、证书验证或真实接口验收。
