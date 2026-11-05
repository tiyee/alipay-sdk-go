# 自定义接口扩展

状态：V2 扩展机制已有实现；本文按新的 V3 优先要求更新目标契约与账单示例，V3 相关字段、构造入口尚待迁移实现，不能作为现有可编译 API 引用。版本基线为 Go 1.27，分层依据见[整体架构](architecture.md)，选型依据见[版本清单](api-versions.md)。

## 1. 扩展目标

当 SDK 没有封装某个支付宝接口，用户可以在自己的业务项目或独立 Go 包中定义请求/响应 struct、具名枚举及 Service，通过公开 Operation 接入 SDK。无需等待 SDK 发版，不修改 SDK，不导入 `internal/`，也不注册全局方法表。

这里的“未内置”指缺少业务方法封装。先查对应官方 V3：有 V3 就采用 V3，确认只有 V2 才用 V2。目标协议尚未实现时等待或补齐对应能力，不能仅因已有 V2 扩展入口就改用 V2；文件上传或流式下载仍需专项支持。

## 2. 公开契约

所有扩展必需类型均位于公开的 `core` 包。下列为契约草案，字段名称与实际实现需在 P1 一并固定：

```go
type Protocol string

const (
	ProtocolV2 Protocol = "v2"
	ProtocolV3 Protocol = "v3"
)

type HTTPMethod string

const (
	HTTPMethodGet  HTTPMethod = "GET"
	HTTPMethodPost HTTPMethod = "POST"
)

type ParameterMode uint8

const (
	ParameterModeBizJSON ParameterMode = iota + 1
	ParameterModeForm
	ParameterModeJSON
	ParameterModeQuery
)

type ResponseCode string

const ResponseCodeSuccess ResponseCode = "10000"

type OperationSpec struct {
	Protocol      Protocol
	Method        string
	HTTPMethod    HTTPMethod
	Path          string
	Version       string
	ResponseKey   string
	ParameterMode ParameterMode
	AcceptedCodes []ResponseCode
}

type EncodedParams struct {
	Mode  ParameterMode
	JSON  jsontext.Value
	Form  url.Values
	Query url.Values
}

type Encoder[T any] interface {
	Encode(value T) (EncodedParams, error)
}

type Decoder[T any] interface {
	Decode(payload jsontext.Value) (T, error)
}

func NewJSONOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error)

func NewQueryOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error)

func NewOperation[Req, Resp any](spec OperationSpec,
	encoder Encoder[Req], decoder Decoder[Resp]) (Operation[Req, Resp], error)
```

- `Protocol` 必填，不根据请求失败结果自动切换。V3 配置官方 HTTP 方法与相对路径；V2 才使用 `Version`、`ResponseKey` 与 `AcceptedCodes`。两种协议的专属字段混用应在发请求前失败。
- `NewJSONOperation` 装配 JSON v2 默认编解码器；V3 `ParameterModeJSON` 编码为 HTTP JSON body，V2 `ParameterModeBizJSON` 才编码为 `biz_content`，要求模式与协议一致。
- `NewQueryOperation` 用于 V3 GET，根据请求 struct 的 `query` 标签生成参数；默认处理字符串、具名字符串枚举和可选指针，复杂格式通过自定义 Encoder 显式实现。
- `NewOperation` 接收用户编解码器，用于直接参数、额外的本地校验或特殊字段格式。接口不得含私有方法，外部包可以直接实现。
- `ParameterModeForm` 表示 V2 业务字段直接编码为参数；V3 GET 使用 `ParameterModeQuery`。最终 HTTP 请求由相应协议处理器构造。
- `JSON`、`Form` 与 `Query` 按模式互斥。已知业务字段始终定义为 struct，只有进入参数编码边界时才转为 `url.Values`。
- 构造时验证接口描述、编码器和解码器，并复制描述中的切片等可变数据；零值或不完整 Operation 在网络调用前失败。
- `AcceptedCodes` 只配置 V2 业务响应策略，不能取消签名验证；V3 使用对应 HTTP 和错误模型，不假定成功响应带有 `code=10000`。

## 3. 示例：用户实现账单下载地址查询

该接口不在首期内置的 5 个交易接口中。官方已提供 **V3 GET** `/v3/alipay/data/dataservice/bill/downloadurl/query`，因此扩展示例必须选择 V3，依据[官方 V3 API](https://github.com/alipay/alipay-sdk-java-all/blob/02177d34fce069f6c20d83c13f963150bb45467b/v3/docs/AlipayDataDataserviceBillDownloadurlApi.md#query)。以下仅获取下载地址，不下载文件、不记录地址到日志；旧 `examples/customapi` 中的 V2 调用需在迁移里程碑同步替换。

模型放在用户自己的包中；示例列出交易账单类型与当前文档中的账单文件结果值，其他类型按官方契约自行增加。

```go
type BillType string

const BillTypeTrade BillType = "trade"

type BillFileCode string

const BillFileCodeEmptyData BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"

type BillRequest struct {
	BillType BillType `query:"bill_type"`
	BillDate string   `query:"bill_date"`
	SMID     *string  `query:"smid"`
}

type BillResponse struct {
	BillDownloadURL *string       `json:"bill_download_url"`
	BillFileCode    *BillFileCode `json:"bill_file_code"`
}
```

V3 HTTP 与错误模型由协议处理器检查，用户模型映射对应业务数据，不添加 V2 响应外层。下载地址等字段及枚举应以 V3 schema 核对；`bill_date` 的日期范围和 `smid` 的适用条件仍需业务校验，不能仅因字段类型正确就认为参数有效。

在用户包中构造 Operation：

```go
func NewBillOperation() (core.Operation[BillRequest, BillResponse], error) {
	return core.NewQueryOperation[BillRequest, BillResponse](core.OperationSpec{
		Protocol:      core.ProtocolV3,
		Method:        "alipay.data.dataservice.bill.downloadurl.query",
		HTTPMethod:    core.HTTPMethodGet,
		Path:          "/v3/alipay/data/dataservice/bill/downloadurl/query",
		ParameterMode: core.ParameterModeQuery,
	})
}
```

直接使用 Client 泛型方法调用；Client 的环境、密钥或证书及 slog 配置沿用调用方既有配置：

```go
func QueryBill(ctx context.Context, client *core.Client,
	req BillRequest) (*core.Response[BillResponse], error) {
	operation, err := NewBillOperation()
	if err != nil {
		return nil, err
	}
	return client.Execute(ctx, operation, core.Request[BillRequest]{Body: req})
}
```

`Operation[BillRequest, BillResponse]` 同时绑定请求与响应类型。调用方检查 error 后才能使用结果数据；下载地址缺失时按业务响应处理，不能将 HTTP 成功自动视为有可下载文件。

必填整数或布尔 query 参数可声明为 `*int64` 或 `*bool` 并标注 `query:"name,required"`。传入 nil 会得到 `*core.ValidationError`，显式传入 0 或 false 会按原值编码。可选参数使用 `query:"name,omitempty"`。

## 4. 用户自己的 Service

需要组合多个接口、复用 Operation 或注入测试执行器时，用户可定义自己的 Service，依赖公开的 `core.Executor`：

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

该包只需导入 `context`、`errors` 和 `github.com/tiyee/go-alipay/core`。若字段需额外转换，可通过 `NewOperation` 注入自己的 Encoder/Decoder；协议中确有条件必填等约束时，Service 或 Encoder 在发请求前进行确定的本地校验。

## 5. 编解码与安全边界

自定义编码器负责将具体 struct 转为当前协议的 JSON body、query 或 V2 业务参数，不生成身份、时间戳或签名。Client 拦截受保护参数覆盖，按显式协议使用相应签名、证书和响应验证实现，并共享 HTTP、slog 及错误接口。

自定义解码器只接收通过响应验签及必要解密后的业务节点，不读取完整原始响应体，不重新决定应否验签。解码错误沿用 SDK 错误链，业务失败和安全失败不填充可信的 `Response[T].Data`。

用户没有配置凭证、未开通产品或不具备接口权限时，按相应配置错误或支付宝响应处理，不因为接口由用户实现而绕过认证。新协议或新的传输方式需要独立能力，不通过替换编解码器冒充已支持。

## 6. 验收要求

P1/P2 的扩展验证需在 SDK 外部包中完成，不能借助包内私有字段或测试后门：

- V3 JSON/GET query 与只有 V2 的参数扩展都能通过公开构造函数接入，V3 不套用 V2 响应结构。
- 用户可以在自己的包中实现 Encoder、Decoder 与 Service，且只依赖公开包。
- 错配的请求/响应类型不能通过编译；自定义枚举和嵌套 struct 不退化为动态业务字典。
- 修改 JSON 编码或响应字段映射不改变签名验签顺序，错误签名时不会调用用户解码器。
- 错误码、ctx、slog、凭证隔离、受保护参数和响应体关闭行为与内置接口一致。
- 示例在功能实现后转为外部包的可编译示例，离线运行默认测试，不访问支付宝或下载真实账单。
