// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package articlesummary

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

// PublicArticleSummaryDataBatchQuery calls alipay.open.public.articlesummary.data.batchquery (POST /v3/alipay/open/public/articlesummary/data/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wz9t
func (s *Service) PublicArticleSummaryDataBatchQuery(ctx context.Context, req PublicArticleSummaryDataBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicArticleSummaryDataBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicArticleSummaryDataBatchQueryRequest, PublicArticleSummaryDataBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.articlesummary.data.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/articlesummary/data/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicArticleSummaryDataBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
