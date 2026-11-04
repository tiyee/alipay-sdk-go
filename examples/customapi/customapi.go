// Package customapi 演示调用方如何在 SDK 之外扩展未内置接口：
// 在用户自己的包中定义请求/响应 struct、具名枚举、Operation 与 Service，
// 仅依赖公开包（core、option、utils），不修改 SDK、不导入 internal/。
//
// 账单下载地址查询的接口名、参数与枚举依据
// https://opendocs.alipay.com/open/02ekfu.md。测试中的
// "custom.demo.form.api" 为编码机制演示用的占位方法名，不是真实接口。
package customapi

import (
	"context"
	"errors"

	"github.com/tiyee/go-alipay/core"
)

// BillType 是对账单类型，依据账单下载地址查询文档。
type BillType string

// BillTypeTrade 表示日账单。
const BillTypeTrade BillType = "trade"

// BillFileCode 是账单文件结果码的文档已知值；未知值原样保留。
type BillFileCode string

// BillFileCodeEmptyData 表示无数据但存在账单文件。
const BillFileCodeEmptyData BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"

// BillRequest 是账单下载地址查询请求。
// bill_date 的取值范围与 smid 的适用条件由调用方按目标账单类型校验。
type BillRequest struct {
	BillType BillType `json:"bill_type"`
	BillDate string   `json:"bill_date"`
	SMID     *string  `json:"smid,omitempty"`
}

// BillResponse 是账单下载地址查询响应；下载地址缺失按业务响应处理，
// 不代表 HTTP 或调用失败。
type BillResponse struct {
	BillDownloadURL *string       `json:"bill_download_url"`
	BillFileCode    *BillFileCode `json:"bill_file_code"`
}

// NewBillOperation 构造账单下载地址查询的 Operation。
func NewBillOperation() (core.Operation[BillRequest, BillResponse], error) {
	return core.NewJSONOperation[BillRequest, BillResponse](core.OperationSpec{
		Method:        "alipay.data.dataservice.bill.downloadurl.query",
		Version:       "1.0",
		ResponseKey:   "alipay_data_dataservice_bill_downloadurl_query_response",
		ParameterMode: core.ParameterModeBizJSON,
		AcceptedCodes: []core.ResponseCode{core.ResponseCodeSuccess},
	})
}

// QueryBill 直接通过 Client 泛型方法调用自定义接口。
func QueryBill(ctx context.Context, client *core.Client,
	req BillRequest) (*core.Response[BillResponse], error) {
	operation, err := NewBillOperation()
	if err != nil {
		return nil, err
	}
	return client.Execute(ctx, operation, core.Request[BillRequest]{Body: req})
}

// BillService 是用户自定义 Service：依赖公开的 core.Executor，
// 可注入 *core.Client 或测试替身。
type BillService struct {
	executor  core.Executor
	operation core.Operation[BillRequest, BillResponse]
}

// NewBillService 构建 Service；执行器为 nil 时返回错误。
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

// Query 查询账单下载地址。
func (s *BillService) Query(ctx context.Context,
	req BillRequest) (*core.Response[BillResponse], error) {
	return core.Execute(ctx, s.executor, s.operation,
		core.Request[BillRequest]{Body: req})
}

// FormRequest 是机制演示用的直接参数请求模型。
type FormRequest struct {
	OutTradeNo string
	Extra      string
}

// FormResponse 是机制演示用的直接参数响应模型。
type FormResponse struct {
	Code    string `json:"code"`
	TradeNo string `json:"trade_no"`
}

// FormEncoder 实现 core.Encoder：把具体 struct 编码为直接参数。
// 编码器只生成业务参数，不生成公共身份、时间戳或签名。
type FormEncoder struct{}

// Encode 实现 core.Encoder。
func (FormEncoder) Encode(value FormRequest) (core.EncodedParams, error) {
	if value.OutTradeNo == "" {
		return core.EncodedParams{}, errors.New("form encoder: out_trade_no is required")
	}
	form := map[string][]string{"out_trade_no": {value.OutTradeNo}}
	if value.Extra != "" {
		form["extra"] = []string{value.Extra}
	}
	return core.EncodedParams{
		Mode: core.ParameterModeForm,
		Form: toURLValues(form),
	}, nil
}
