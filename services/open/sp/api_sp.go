package sp

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

// SpImageUpload calls alipay.open.sp.image.upload (POST /v3/alipay/open/sp/image/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/59220073_alipay.open.sp.image.upload
func (s *Service) SpImageUpload(ctx context.Context, req SpImageUploadRequest, opts ...core.CallOption) (*core.Response[SpImageUploadResponse], error) {
	return core.ExecuteJSON[SpImageUploadRequest, SpImageUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.image.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/image/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpImageUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpLeadsExpandCreate calls alipay.open.sp.leads.expand.create (POST /v3/alipay/open/sp/leads/expand/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05vh1e
func (s *Service) SpLeadsExpandCreate(ctx context.Context, req SpLeadsExpandCreateRequest, opts ...core.CallOption) (*core.Response[SpLeadsExpandCreateResponse], error) {
	return core.ExecuteJSON[SpLeadsExpandCreateRequest, SpLeadsExpandCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.leads.expand.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/leads/expand/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpLeadsExpandCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpMerchantInconsistentApprove calls alipay.open.sp.merchant.inconsistent.approve (POST /v3/alipay/open/sp/merchant/inconsistent/approve).
// Official V3 contract: https://opendocs.alipay.com/apis/05v8ke
func (s *Service) SpMerchantInconsistentApprove(ctx context.Context, req SpMerchantInconsistentApproveRequest, opts ...core.CallOption) (*core.Response[SpMerchantInconsistentApproveResponse], error) {
	return core.ExecuteJSON[SpMerchantInconsistentApproveRequest, SpMerchantInconsistentApproveResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.merchant.inconsistent.approve", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/merchant/inconsistent/approve",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpMerchantInconsistentApproveRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
