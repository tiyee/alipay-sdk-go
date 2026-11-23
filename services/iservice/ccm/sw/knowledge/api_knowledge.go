package knowledge

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

// CcmSwKnowledgeModify calls alipay.iservice.ccm.sw.knowledge.modify (POST /v3/alipay/iservice/ccm/sw/knowledge/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c77df3ee_alipay.iservice.ccm.sw.knowledge.modify
func (s *Service) CcmSwKnowledgeModify(ctx context.Context, req CcmSwKnowledgeModifyRequest, opts ...core.CallOption) (*core.Response[CcmSwKnowledgeModifyResponse], error) {
	return core.ExecuteJSON[CcmSwKnowledgeModifyRequest, CcmSwKnowledgeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.knowledge.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/sw/knowledge/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwKnowledgeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
