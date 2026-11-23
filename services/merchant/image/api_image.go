package image

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

// ImageUpload calls alipay.merchant.image.upload (POST /v3/alipay/merchant/image/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/132c6b5b_alipay.merchant.image.upload
func (s *Service) ImageUpload(ctx context.Context, req MerchantImageUploadRequest, opts ...core.CallOption) (*core.Response[MerchantImageUploadResponse], error) {
	return core.ExecuteJSON[MerchantImageUploadRequest, MerchantImageUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.image.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/image/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantImageUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
