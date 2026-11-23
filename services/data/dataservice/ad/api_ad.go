package ad

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

// DataServiceAdConversionUpload calls alipay.data.dataservice.ad.conversion.upload (POST /v3/alipay/data/dataservice/ad/conversion/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/32a2992f_alipay.data.dataservice.ad.conversion.upload
func (s *Service) DataServiceAdConversionUpload(ctx context.Context, req DataDataServiceAdConversionUploadRequest, opts ...core.CallOption) (*core.Response[DataDataServiceAdConversionUploadResponse], error) {
	return core.ExecuteJSON[DataDataServiceAdConversionUploadRequest, DataDataServiceAdConversionUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.dataservice.ad.conversion.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/data/dataservice/ad/conversion/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataDataServiceAdConversionUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// DataServiceAdPromotePageBatchQuery calls alipay.data.dataservice.ad.promotepage.batchquery (GET /v3/alipay/data/dataservice/ad/promotepage/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/17d7cd6f_alipay.data.dataservice.ad.promotepage.batchquery
func (s *Service) DataServiceAdPromotePageBatchQuery(ctx context.Context, req DataDataServiceAdPromotePageBatchQueryRequest, opts ...core.CallOption) (*core.Response[DataDataServiceAdPromotePageBatchQueryResponse], error) {
	return core.ExecuteJSON[DataDataServiceAdPromotePageBatchQueryRequest, DataDataServiceAdPromotePageBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.dataservice.ad.promotepage.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/dataservice/ad/promotepage/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataDataServiceAdPromotePageBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// DataServiceAdPromotePageDownload calls alipay.data.dataservice.ad.promotepage.download (GET /v3/alipay/data/dataservice/ad/promotepage/download).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2887027f_alipay.data.dataservice.ad.promotepage.download
func (s *Service) DataServiceAdPromotePageDownload(ctx context.Context, req DataDataServiceAdPromotePageDownloadRequest, opts ...core.CallOption) (*core.Response[DataDataServiceAdPromotePageDownloadResponse], error) {
	return core.ExecuteJSON[DataDataServiceAdPromotePageDownloadRequest, DataDataServiceAdPromotePageDownloadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.dataservice.ad.promotepage.download", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/dataservice/ad/promotepage/download",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataDataServiceAdPromotePageDownloadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
