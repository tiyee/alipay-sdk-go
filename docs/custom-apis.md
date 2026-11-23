# 自定义接口扩展

状态：与当前 V3-only 实现一致（2026-11-05）。本文说明如何调用 SDK 未内置的支付宝 V3 接口：在用户自己的包中定义请求/响应 struct、具名枚举与 Operation，仅依赖公开包（`core`、`core/option`、`utils` 等），不修改 SDK、不导入 `internal/`、不注册全局方法表。可运行示例见 `examples/customapi`，分层依据见[整体架构](architecture.md)。

## 1. 扩展目标与前提

SDK 当前提供官方快照中的全部 746 个 REST 接口，其中 678 个属于 `alipay.*`。详见[接口版本说明](api-versions.md)。未内置接口按以下前提扩展：

- 先核实目标接口的官方 V3 契约（方法名、HTTP 方法、`/v3/...` 路由、参数位置、认证要求）。SDK 仅接入已确认 V3 契约的接口；未确认时不构造猜测路由。
- 文件上传使用 `ParameterModeMultipart`。流式下载和内容加密尚待实现。
- 用户负责接口文档、业务模型和校验；SDK 统一执行公共身份与签名、验签、证书、HTTP、slog 与错误处理。

## 2. 公开契约

所有扩展必需类型均位于公开的 `core` 包：

```go
type HTTPMethod string

const (
	HTTPMethodGet  HTTPMethod = "GET"
	HTTPMethodPost HTTPMethod = "POST"
	HTTPMethodPut HTTPMethod = "PUT"
	HTTPMethodDelete HTTPMethod = "DELETE"
	HTTPMethodPatch HTTPMethod = "PATCH"
)

type ParameterMode uint8

const (
	ParameterModeJSONBody ParameterMode = iota + 1 // JSON 正文
	ParameterModeQuery                             // query 或空请求体
	ParameterModeMultipart                         // multipart/form-data
	ParameterModeJSONBodyWithQuery                  // query 与独立 JSON 正文
)

type AuthRequirement uint8

const (
	AuthRequirementAny         AuthRequirement = iota + 1 // 公钥与证书模式均可
	AuthRequirementCertificate                            // 官方要求证书模式的接口
)

type OperationSpec struct {
	Method          string // 官方接口名，如 "alipay.data.dataservice.bill.downloadurl.query"
	HTTPMethod      HTTPMethod
	Path            string // "/v3/..." 相对路由
	ParameterMode   ParameterMode // 零值时按 HTTPMethod 推断（GET→Query，POST→JSONBody）
	AuthRequirement AuthRequirement
}

type Encoder[T any] interface {
	Encode(value T) (EncodedParams, error)
}

type Decoder[T any] interface {
	Decode(payload json.RawMessage) (T, error)
}

// 默认编解码入口。
func NewJSONOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error)

// 自定义编解码器入口。
func NewOperation[Req, Resp any](spec OperationSpec,
	encoder Encoder[Req], decoder Decoder[Resp]) (Operation[Req, Resp], error)
```

- `NewJSONOperation` 绑定 SDK 默认编解码器：依据显式的参数模式分别编码 query、JSON 或 multipart。混合模式中的 query 字段同时标注 `query:"name" json:"-"`，以便独立发送和签名。`ParameterMode` 留零值时按 `HTTPMethod` 推断。
- `NewOperation` 接收用户实现的 `Encoder[Req]` / `Decoder[Resp]`，用于直接参数、额外本地校验或特殊字段格式。接口不包含私有方法，外部包可以直接实现。
- 编码器只生成业务参数：V3 的公共身份参数和签名头由 Client 设置，编码器只能提供业务 query、路径占位值与正文，签名头和凭证由 Client 单独构建。
- 同一输入产生的编码结果同时用于签名原文与传输，编码器输出必须确定；JSON 使用 `json.RawMessage` 保留一份编码结果。
- 构造时验证规格、编码器和解码器；零值或无效 Operation 在网络调用前失败。`AuthRequirement` 为 `Certificate` 时，公钥模式 Client 在执行前返回错误。

## 3. 示例：调用内置清单外的 V3 接口

以账单下载地址查询为例（该接口同时有内置实现，此处演示扩展机制本身；真实调用建议直接使用 `services/data`）。官方契约为 **V3 GET** `/v3/alipay/data/dataservice/bill/downloadurl/query`，依据[官方 V3 API 文档](https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49)。以下仅获取下载地址，不下载文件、不记录地址到日志。

模型放在用户自己的包中；枚举使用具名类型与带类型常量：

```go
type BillType string

const BillTypeTrade BillType = "trade"

type BillFileCode string

const BillFileCodeEmptyData BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"

type BillRequest struct {
	BillType BillType `query:"bill_type"`
	BillDate string   `query:"bill_date"`
	SMID     *string  `query:"smid,omitempty"`
}

type BillResponse struct {
	BillDownloadURL *string       `json:"bill_download_url"`
	BillFileCode    *BillFileCode `json:"bill_file_code"`
}
```

在用户包中用 `NewJSONOperation` 构造 Operation：

```go
func NewBillOperation() (core.Operation[BillRequest, BillResponse], error) {
	return core.NewJSONOperation[BillRequest, BillResponse](core.OperationSpec{
		Method:          "alipay.data.dataservice.bill.downloadurl.query",
		HTTPMethod:      core.HTTPMethodGet,
		Path:            "/v3/alipay/data/dataservice/bill/downloadurl/query",
		ParameterMode:   core.ParameterModeQuery,
		AuthRequirement: core.AuthRequirementAny,
	})
}
```

所有支持版本均可将 Client 传给包级 `core.Execute`。Go 1.27 起也可写作 `client.Execute(ctx, operation, req)`。

```go
func QueryBill(ctx context.Context, client *core.Client,
	req BillRequest) (*core.Response[BillResponse], error) {
	operation, err := NewBillOperation()
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, client, operation, core.Request[BillRequest]{Body: req})
}
```

`Operation[BillRequest, BillResponse]` 在编译期绑定请求与响应类型。调用方检查 error 后才能使用结果数据；下载地址缺失时按业务响应处理，不能将 HTTP 成功自动视为有可下载文件。

字段映射与必填条件以官方 V3 文档为准：`bill_date` 的日期范围、`smid` 的适用条件等业务约束由调用方或自定义 Encoder 在发请求前做确定的本地校验。

## 4. 自定义编解码器

字段名不是下划线风格、需要条件必填或特殊格式时，用 `NewOperation` 接入自己的 Encoder/Decoder。以下摘自 `examples/customapi/query_values.go`（方法名与路由为演示编码机制的占位值，不是真实接口）：

```go
// DirectQueryEncoder 实现 core.Encoder：把具体 struct 编码为 V3 query 参数。
// Query 模式填 Form，JSONBody 模式填 JSON，混合模式两者均填。
type DirectQueryEncoder struct{}

func (DirectQueryEncoder) Encode(v DirectParamsRequest) (core.EncodedParams, error) {
	q := url.Values{}
	q.Set("out_trade_no", v.OutTradeNo)
	if v.Extra != "" {
		q.Set("extra", v.Extra)
	}
	return core.EncodedParams{Mode: core.ParameterModeQuery, Form: q}, nil
}

func NewDirectParamsOperation() (core.Operation[DirectParamsRequest, DirectParamsResponse], error) {
	return core.NewOperation[DirectParamsRequest, DirectParamsResponse](
		core.OperationSpec{
			Method:          "custom.demo.query.api",
			HTTPMethod:      core.HTTPMethodGet,
			Path:            "/v3/custom/demo/query",
			ParameterMode:   core.ParameterModeQuery,
			AuthRequirement: core.AuthRequirementAny,
		}, DirectQueryEncoder{}, core.JSONDecoder[DirectParamsResponse]())
}
```

仅有 JSON 正文的接口编码器返回 `core.EncodedParams{Mode: core.ParameterModeJSONBody, JSON: ...}`（`JSON` 为 `json.RawMessage`，可用 `core.JSONBodyEncoder[T]()` 作为默认实现）；响应默认用 `core.JSONDecoder[T]()` 解码，特殊响应结构可实现自己的 Decoder。SDK 另提供 `core.QueryEncoder[T]()` 作为 `query` 标签编码的公开实现，可作为自定义编码器的参照。

必填整数或布尔 query 参数可声明为 `*int64` 或 `*bool` 并标注 `query:"name,required"`。传入 nil 会得到 `*core.ValidationError`，显式传入 0 或 false 会按原值编码。可选参数使用 `query:"name,omitempty"`。

## 5. 用户自己的 Service

需要组合多个接口、复用 Operation 或注入测试执行器时，用户可定义自己的 Service，依赖公开的非泛型 `core.Executor`（`*core.Client` 与测试替身均可注入）：

```go
type BillService struct {
	executor  core.Executor
	operation core.Operation[BillRequest, BillResponse]
}

func NewBillService(executor core.Executor) (*BillService, error) {
	if executor == nil {
		return nil, errors.New("bill service: executor is required")
	}
	operation, err := NewBillOperation()
	if err != nil {
		return nil, err
	}
	return &BillService{executor: executor, operation: operation}, nil
}

func (s *BillService) Query(ctx context.Context,
	req BillRequest) (*core.Response[BillResponse], error) {
	return core.Execute(ctx, s.executor, s.operation,
		core.Request[BillRequest]{Body: req})
}
```

该包只需导入 `context`、`errors` 和 `github.com/tiyee/alipay-sdk-go/core`。调用级参数（如代调用令牌 `core.WithAppAuthToken`）经 `core.Request.Body` 之外的 CallOption 传入 `core.Execute` 的 Options，不修改共享 Client。

## 6. 编解码与安全边界

- 自定义编码器负责将具体 struct 转为 V3 JSON body 或 query 参数，不生成身份、时间戳或签名；SDK 拦截受保护参数覆盖。
- 自定义解码器只接收通过响应验签后的业务数据（HTTP 成功的 2xx 响应体），不读取完整原始响应体，不重新决定应否验签。验签失败或错误签名时不会调用用户解码器。
- 解码错误沿用 SDK 错误链；业务失败（非 2xx 错误体）与安全失败不填充可信的 `Response[T].Data`。
- 用户没有配置凭证、未开通产品或不具备接口权限时，按相应配置错误或支付宝错误响应处理，不因为接口由用户实现而绕过认证。`AuthRequirement` 声明为 `Certificate` 的接口在公钥模式下于执行前失败。
- 上传使用现有 multipart 编码器。流式下载和内容加密需要相应协议支持。

## 7. 验收要求

扩展机制经外部测试包（`examples/customapi` 及其测试）验证，不借助 SDK 内部字段或测试后门：

- 外部包仅依赖公开包即可定义模型、Encoder/Decoder、Operation 与 Service。
- 错配的请求/响应类型不能通过编译；自定义枚举和嵌套 struct 不退化为动态业务字典。
- 修改 JSON 编码或响应字段映射不改变签名验签顺序；错误签名时不会调用用户解码器。
- 错误分类、ctx、slog、凭证隔离、受保护参数和响应体关闭行为与内置接口一致。
- 示例离线运行默认测试，不访问支付宝或下载真实账单。
