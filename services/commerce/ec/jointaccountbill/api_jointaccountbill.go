package jointaccountbill

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

// EcJointAccountBillDetailBatchQuery calls alipay.commerce.ec.jointaccountbill.detail.batchquery (GET /v3/alipay/commerce/ec/jointaccountbill/detail/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05j4cf
func (s *Service) EcJointAccountBillDetailBatchQuery(ctx context.Context, req EcJointAccountBillDetailBatchQueryRequest, opts ...core.CallOption) (*core.Response[EcJointAccountBillDetailBatchQueryResponse], error) {
	return core.ExecuteJSON[EcJointAccountBillDetailBatchQueryRequest, EcJointAccountBillDetailBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.jointaccountbill.detail.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/jointaccountbill/detail/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcJointAccountBillDetailBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
