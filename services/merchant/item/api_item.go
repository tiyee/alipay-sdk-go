package item

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

// ItemFileUpload calls alipay.merchant.item.file.upload (POST /v3/alipay/merchant/item/file/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/1f740bc5_alipay.merchant.item.file.upload
func (s *Service) ItemFileUpload(ctx context.Context, req MerchantItemFileUploadRequest, opts ...core.CallOption) (*core.Response[MerchantItemFileUploadResponse], error) {
	return core.ExecuteJSON[MerchantItemFileUploadRequest, MerchantItemFileUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.item.file.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/item/file/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantItemFileUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
