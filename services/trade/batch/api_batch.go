package batch

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 batch 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor           core.Executor
	batchTransfer      core.Operation[BatchTransferRequest, BatchTransferResponse]
	batchTransferQuery core.Operation[BatchTransferQueryRequest, BatchTransferQueryResponse]
	batchSettle        core.Operation[BatchSettleRequest, BatchSettleResponse]
	batchSettleQuery   core.Operation[BatchSettleQueryRequest, BatchSettleQueryResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	batchTransfer, err := core.NewJSONOperation[BatchTransferRequest, BatchTransferResponse](core.OperationSpec{
		Method: "alipay.trade.batch.transfer", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/batch/transfer",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	batchTransferQuery, err := core.NewJSONOperation[BatchTransferQueryRequest, BatchTransferQueryResponse](core.OperationSpec{
		Method: "alipay.trade.batch.transfer.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/trade/batch/transfer/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	batchSettle, err := core.NewJSONOperation[BatchSettleRequest, BatchSettleResponse](core.OperationSpec{
		Method: "alipay.trade.batch.settle", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/batch/settle",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	batchSettleQuery, err := core.NewJSONOperation[BatchSettleQueryRequest, BatchSettleQueryResponse](core.OperationSpec{
		Method: "alipay.trade.batch.settle.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/batch/settle/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:           executor,
		batchTransfer:      batchTransfer,
		batchTransferQuery: batchTransferQuery,
		batchSettle:        batchSettle,
		batchSettleQuery:   batchSettleQuery,
	}, nil
}

// BatchTransfer 调用 alipay.trade.batch.transfer 发起请款分账
// （V3 POST /v3/alipay/trade/batch/transfer）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05vd9z
// 多个分账条款要求出账方只能有一个；处理结果以查询接口为准。
func (s *Service) BatchTransfer(ctx context.Context, req BatchTransferRequest,
	opts ...core.CallOption) (*core.Response[BatchTransferResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.batchTransfer,
		core.Request[BatchTransferRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BatchTransferQuery 调用 alipay.trade.batch.transfer.query 查询请款
// 分账结果（V3 GET /v3/alipay/trade/batch/transfer/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05vsbp
func (s *Service) BatchTransferQuery(ctx context.Context, req BatchTransferQueryRequest,
	opts ...core.CallOption) (*core.Response[BatchTransferQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.batchTransferQuery,
		core.Request[BatchTransferQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BatchSettle 调用 alipay.trade.batch.settle 发起批量结算
// （V3 POST /v3/alipay/trade/batch/settle）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05vfpn
// 重复的 OutRequestNo 返回既有单据状态；受理结果以 ResultCode 为准，
// 处理结果以批量结算查询为准。
func (s *Service) BatchSettle(ctx context.Context, req BatchSettleRequest,
	opts ...core.CallOption) (*core.Response[BatchSettleResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.batchSettle,
		core.Request[BatchSettleRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BatchSettleQuery 调用 alipay.trade.batch.settle.query 查询批量结算
// 结果（V3 POST /v3/alipay/trade/batch/settle/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05vel6
func (s *Service) BatchSettleQuery(ctx context.Context, req BatchSettleQueryRequest,
	opts ...core.CallOption) (*core.Response[BatchSettleQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.batchSettleQuery,
		core.Request[BatchSettleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
