// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package singlearticle

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// PublicSingleArticleDataBatchQuery calls alipay.open.public.singlearticle.data.batchquery (POST /v3/alipay/open/public/singlearticle/data/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuky
func (s *Service) PublicSingleArticleDataBatchQuery(ctx context.Context, req PublicSingleArticleDataBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicSingleArticleDataBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicSingleArticleDataBatchQueryRequest, PublicSingleArticleDataBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.singlearticle.data.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/singlearticle/data/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicSingleArticleDataBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
