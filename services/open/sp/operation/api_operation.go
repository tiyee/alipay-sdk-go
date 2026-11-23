package operation

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor core.Executor
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}

	return &Service{
		executor: executor,
	}, nil
}

// SpOperationApply calls alipay.open.sp.operation.apply (POST /v3/alipay/open/sp/operation/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a8869af2_alipay.open.sp.operation.apply?scene=common&pathHash=6c69600a
func (s *Service) SpOperationApply(ctx context.Context, req SpOperationApplyRequest, opts ...core.CallOption) (*core.Response[SpOperationApplyResponse], error) {
	return core.ExecuteJSON[SpOperationApplyRequest, SpOperationApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.operation.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/operation/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpOperationApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpOperationQRCodeQuery calls alipay.open.sp.operation.qrcode.query (GET /v3/alipay/open/sp/operation/qrcode/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8245f76c_alipay.open.sp.operation.qrcode.query?scene=common&pathHash=c82f7840
func (s *Service) SpOperationQRCodeQuery(ctx context.Context, req SpOperationQRCodeQueryRequest, opts ...core.CallOption) (*core.Response[SpOperationQRCodeQueryResponse], error) {
	return core.ExecuteJSON[SpOperationQRCodeQueryRequest, SpOperationQRCodeQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.operation.qrcode.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/operation/qrcode/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpOperationQRCodeQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpOperationResultQuery calls alipay.open.sp.operation.result.query (GET /v3/alipay/open/sp/operation/result/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/229831e2_alipay.open.sp.operation.result.query?scene=common&pathHash=4c2b6906
func (s *Service) SpOperationResultQuery(ctx context.Context, req SpOperationResultQueryRequest, opts ...core.CallOption) (*core.Response[SpOperationResultQueryResponse], error) {
	return core.ExecuteJSON[SpOperationResultQueryRequest, SpOperationResultQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.operation.result.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/operation/result/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpOperationResultQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
