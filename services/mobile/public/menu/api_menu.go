package menu

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

// PublicMenuAdd calls alipay.mobile.public.menu.add (POST /v3/alipay/mobile/public/menu/add).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wuom
func (s *Service) PublicMenuAdd(ctx context.Context, req MobilePublicMenuAddRequest, opts ...core.CallOption) (*core.Response[MobilePublicMenuAddResponse], error) {
	return core.ExecuteJSON[MobilePublicMenuAddRequest, MobilePublicMenuAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.menu.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/menu/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMenuAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuGet calls alipay.mobile.public.menu.get (GET /v3/alipay/mobile/public/menu/get).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wyfs
func (s *Service) PublicMenuGet(ctx context.Context, req MobilePublicMenuGetRequest, opts ...core.CallOption) (*core.Response[MobilePublicMenuGetResponse], error) {
	return core.ExecuteJSON[MobilePublicMenuGetRequest, MobilePublicMenuGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.menu.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/mobile/public/menu/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMenuGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuUpdate calls alipay.mobile.public.menu.update (POST /v3/alipay/mobile/public/menu/update).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05x1fv
func (s *Service) PublicMenuUpdate(ctx context.Context, req MobilePublicMenuUpdateRequest, opts ...core.CallOption) (*core.Response[MobilePublicMenuUpdateResponse], error) {
	return core.ExecuteJSON[MobilePublicMenuUpdateRequest, MobilePublicMenuUpdateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.menu.update", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/menu/update",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMenuUpdateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
