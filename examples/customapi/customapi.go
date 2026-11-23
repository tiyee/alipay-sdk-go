// Package customapi 演示调用方如何在 SDK 之外扩展未内置的 V3 接口：
// 在用户自己的包中定义请求/响应 struct、具名枚举、OperationSpec 与
// 编解码器，仅依赖公开包（core、option、utils），不修改 SDK、
// 不导入 internal/。
//
// 账单下载地址查询为 V3 GET 接口，接口名、参数与枚举依据官方 V3 文档
// https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49
// 通过 NewJSONOperation 使用默认编解码器。第二个演示（query_values.go）
// 用 core.NewOperation 接入自实现的 query 编码器；其中
// "custom.demo.query.api" 与 "/v3/custom/demo/query" 为编码机制演示
// 用的占位接口名与路由，不是真实支付宝接口。
package customapi

import (
	"context"
	"errors"

	"github.com/tiyee/alipay-sdk-go/core"
)

// BillType 是对账单类型，依据账单下载地址查询文档。
type BillType string

// BillTypeTrade 表示日账单。
const BillTypeTrade BillType = "trade"

// BillFileCode 是账单文件结果码的文档已知值；未知值原样保留。
type BillFileCode string

// BillFileCodeEmptyData 表示无数据但存在账单文件。
const BillFileCodeEmptyData BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"

// BillRequest 是账单下载地址查询请求（V3 GET query 参数）。
// bill_date 的取值范围与 smid 的适用条件由调用方按目标账单类型校验。
type BillRequest struct {
	BillType BillType `query:"bill_type"`
	BillDate string   `query:"bill_date"`
	SMID     *string  `query:"smid,omitempty"`
}

// BillResponse 是账单下载地址查询响应；下载地址缺失按业务响应处理，
// 不代表 HTTP 或调用失败。
type BillResponse struct {
	BillDownloadURL *string       `json:"bill_download_url"`
	BillFileCode    *BillFileCode `json:"bill_file_code"`
}

// NewBillOperation 构造账单下载地址查询的 V3 GET Operation。
// NewJSONOperation 绑定默认编解码器：GET 接口按契约使用 query 编码，
// QueryEncoder 依据 `query` 标签展开参数，响应由 JSONDecoder 解码。
func NewBillOperation() (core.Operation[BillRequest, BillResponse], error) {
	return core.NewJSONOperation[BillRequest, BillResponse](core.OperationSpec{
		Method:          "alipay.data.dataservice.bill.downloadurl.query",
		HTTPMethod:      core.HTTPMethodGet,
		Path:            "/v3/alipay/data/dataservice/bill/downloadurl/query",
		ParameterMode:   core.ParameterModeQuery,
		AuthRequirement: core.AuthRequirementAny,
	})
}

// QueryBill 通过包级泛型入口调用自定义接口。
func QueryBill(ctx context.Context, client *core.Client,
	req BillRequest) (*core.Response[BillResponse], error) {
	operation, err := NewBillOperation()
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, client, operation, core.Request[BillRequest]{Body: req})
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
