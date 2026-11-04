# 自定义接口扩展

状态：公开 API 设计草案，随 P1/P2 实现；当前仓库尚无 SDK 源码，以下示例不能作为现有可编译能力引用。版本基线为 Go 1.27，分层依据见[整体架构](architecture.md)。

## 1. 扩展目标

当 SDK 没有封装某个支付宝接口，用户可以在自己的业务项目或独立 Go 包中定义请求/响应 struct、具名枚举及 Service，通过公开 Operation 接入 SDK。无需等待 SDK 发版，不修改 SDK，不导入 `internal/`，也不注册全局方法表。

这里的“未内置”指缺少业务方法封装。目标 API 仍需使用 SDK 已实现的网关协议、认证、编码和响应处理能力；扩展入口不会自动获得尚未实现的 V3、文件上传或流式下载支持。

## 2. 公开契约

所有扩展必需类型均位于公开的 `core` 包。下列为契约草案，字段名称与实际实现需在 P1 一并固定：

```go
type ParameterMode uint8

const (
	ParameterModeBizJSON ParameterMode = iota + 1
	ParameterModeForm
)

type ResponseCode string

const ResponseCodeSuccess ResponseCode = "10000"

type OperationSpec struct {
	Method        string
	Version       string
	ResponseKey   string
	ParameterMode ParameterMode
	AcceptedCodes []ResponseCode
}

type EncodedParams struct {
	Mode ParameterMode
	JSON jsontext.Value
	Form url.Values
}

type Encoder[T any] interface {
	Encode(value T) (EncodedParams, error)
}

type Decoder[T any] interface {
	Decode(payload jsontext.Value) (T, error)
}

func NewJSONOperation[Req, Resp any](spec OperationSpec) (Operation[Req, Resp], error)

func NewOperation[Req, Resp any](spec OperationSpec,
	encoder Encoder[Req], decoder Decoder[Resp]) (Operation[Req, Resp], error)
```

- `NewJSONOperation` 装配 JSON v2 默认编解码器，适用于 `biz_content` 为 JSON 的接口；要求参数模式与构造入口一致。
- `NewOperation` 接收用户编解码器，用于直接参数、额外的本地校验或特殊字段格式。接口不得含私有方法，外部包可以直接实现。
- `ParameterModeForm` 表示业务字段直接编码为参数，不自动封装到 `biz_content`；最终传输仍由 Client 按网关协议构造。
- `JSON` 与 `Form` 按模式互斥使用。已知业务字段始终定义为 struct，只有进入表单编码边界时才转为 `url.Values`。
- 构造时验证接口描述、编码器和解码器，并复制描述中的切片等可变数据；零值或不完整 Operation 在网络调用前失败。
- `AcceptedCodes` 只配置接口业务响应策略，不能取消签名验证；处理中的响应也不代表商户业务终态成功。

## 3. 示例：用户实现账单下载地址查询

该接口不在首期内置的 5 个交易接口中，适合作为用户扩展示例。方法名、参数、枚举和响应字段依据[支付宝账单下载地址查询文档](https://opendocs.alipay.com/open/02ekfu.md)。以下仅获取下载地址，不下载账单文件，不记录地址到日志。

模型放在用户自己的包中；示例列出交易账单类型与当前文档中的账单文件结果值，其他类型按官方契约自行增加。

```go
type BillType string

const BillTypeTrade BillType = "trade"

type BillFileCode string

const BillFileCodeEmptyData BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"

type BillRequest struct {
	BillType BillType `json:"bill_type"`
	BillDate string   `json:"bill_date"`
	SMID     *string  `json:"smid,omitzero"`
}

type BillResponse struct {
	BillDownloadURL *string       `json:"bill_download_url"`
	BillFileCode    *BillFileCode `json:"bill_file_code"`
}
```

通用响应码由 Client 统一检查；用户业务模型只需映射业务数据。下载地址与文件结果字段均按可选字段处理，未知结果枚举保留原始字符串值。`bill_date` 的日期范围及 `smid` 的适用条件由用户按目标账单类型校验，不能仅因字段类型正确就认为业务参数有效。

在用户包中构造 Operation：

```go
func NewBillOperation() (core.Operation[BillRequest, BillResponse], error) {
	return core.NewJSONOperation[BillRequest, BillResponse](core.OperationSpec{
		Method:        "alipay.data.dataservice.bill.downloadurl.query",
		Version:       "1.0",
		ResponseKey:   "alipay_data_dataservice_bill_downloadurl_query_response",
		ParameterMode: core.ParameterModeBizJSON,
		AcceptedCodes: []core.ResponseCode{core.ResponseCodeSuccess},
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

自定义编码器负责将具体 struct 转为业务 JSON 或直接参数，不生成公共身份、证书序列号、时间戳或签名。Client 会拦截对受保护公共参数的覆盖，使用同一套 HTTP、签名、证书、响应验证、slog 和错误处理。

自定义解码器只接收通过响应验签及必要解密后的业务节点，不读取完整原始响应体，不重新决定应否验签。解码错误沿用 SDK 错误链，业务失败和安全失败不填充可信的 `Response[T].Data`。

用户没有配置凭证、未开通产品或不具备接口权限时，按相应配置错误或支付宝响应处理，不因为接口由用户实现而绕过认证。新协议或新的传输方式需要独立能力，不通过替换编解码器冒充已支持。

## 6. 验收要求

P1/P2 的扩展验证需在 SDK 外部包中完成，不能借助包内私有字段或测试后门：

- JSON 与直接参数两类自定义 API 都能通过公开构造函数接入。
- 用户可以在自己的包中实现 Encoder、Decoder 与 Service，且只依赖公开包。
- 错配的请求/响应类型不能通过编译；自定义枚举和嵌套 struct 不退化为动态业务字典。
- 修改 JSON 编码或响应字段映射不改变签名验签顺序，错误签名时不会调用用户解码器。
- 错误码、ctx、slog、凭证隔离、受保护参数和响应体关闭行为与内置接口一致。
- 示例在功能实现后转为外部包的可编译示例，离线运行默认测试，不访问支付宝或下载真实账单。
