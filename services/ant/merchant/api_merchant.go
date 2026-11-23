// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service invokes V3 operations in this domain.
type Service struct{ executor core.Executor }

// NewService creates a service using the supplied executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

func execute[Req, Resp any](s *Service, ctx context.Context, method string, httpMethod core.HTTPMethod, path string, mode core.ParameterMode, req Req, opts []core.CallOption) (*core.Response[Resp], error) {
	op, err := core.NewJSONOperation[Req, Resp](core.OperationSpec{Method: method, HTTPMethod: httpMethod, Path: path, ParameterMode: mode, AuthRequirement: core.AuthRequirementAny})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, op, core.Request[Req]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ExpandAppRecommendAccountCreate calls ant.merchant.expand.apprecommend.account.create.
// Official V3 contract: https://opendocs.alipay.com/open-v3/32d136f0_ant.merchant.expand.apprecommend.account.create
func (s *Service) ExpandAppRecommendAccountCreate(ctx context.Context, req ExpandAppRecommendAccountCreateRequest, opts ...core.CallOption) (*core.Response[ExpandAppRecommendAccountCreateResponse], error) {
	return execute[ExpandAppRecommendAccountCreateRequest, ExpandAppRecommendAccountCreateResponse](s, ctx, "ant.merchant.expand.apprecommend.account.create", core.HTTPMethodPost, "/v3/ant/merchant/expand/apprecommend/account/create", core.ParameterModeJSONBody, req, opts)
}

// ExpandAppRecommendAccountDelete calls ant.merchant.expand.apprecommend.account.delete.
// Official V3 contract: https://opendocs.alipay.com/open-v3/130a15de_ant.merchant.expand.apprecommend.account.delete
func (s *Service) ExpandAppRecommendAccountDelete(ctx context.Context, req ExpandAppRecommendAccountDeleteRequest, opts ...core.CallOption) (*core.Response[ExpandAppRecommendAccountDeleteResponse], error) {
	return execute[ExpandAppRecommendAccountDeleteRequest, ExpandAppRecommendAccountDeleteResponse](s, ctx, "ant.merchant.expand.apprecommend.account.delete", core.HTTPMethodDelete, "/v3/ant/merchant/expand/apprecommend/account/delete", core.ParameterModeQuery, req, opts)
}

// ExpandAppRecommendAccountQuery calls ant.merchant.expand.apprecommend.account.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/cb07737c_ant.merchant.expand.apprecommend.account.query
func (s *Service) ExpandAppRecommendAccountQuery(ctx context.Context, req ExpandAppRecommendAccountQueryRequest, opts ...core.CallOption) (*core.Response[ExpandAppRecommendAccountQueryResponse], error) {
	return execute[ExpandAppRecommendAccountQueryRequest, ExpandAppRecommendAccountQueryResponse](s, ctx, "ant.merchant.expand.apprecommend.account.query", core.HTTPMethodGet, "/v3/ant/merchant/expand/apprecommend/account/query", core.ParameterModeQuery, req, opts)
}

// ExpandAppRecommendAvailableQuery calls ant.merchant.expand.apprecommend.available.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/fc63d6c8_ant.merchant.expand.apprecommend.available.query
func (s *Service) ExpandAppRecommendAvailableQuery(ctx context.Context, req ExpandAppRecommendAvailableQueryRequest, opts ...core.CallOption) (*core.Response[ExpandAppRecommendAvailableQueryResponse], error) {
	return execute[ExpandAppRecommendAvailableQueryRequest, ExpandAppRecommendAvailableQueryResponse](s, ctx, "ant.merchant.expand.apprecommend.available.query", core.HTTPMethodGet, "/v3/ant/merchant/expand/apprecommend/available/query", core.ParameterModeQuery, req, opts)
}

// ExpandInfoQuery calls ant.merchant.expand.info.query.
// Official V3 contract: https://opendocs.alipay.com/apis/05vb3x
func (s *Service) ExpandInfoQuery(ctx context.Context, req ExpandInfoQueryRequest, opts ...core.CallOption) (*core.Response[ExpandInfoQueryResponse], error) {
	return execute[ExpandInfoQueryRequest, ExpandInfoQueryResponse](s, ctx, "ant.merchant.expand.info.query", core.HTTPMethodGet, "/v3/ant/merchant/expand/info/query", core.ParameterModeQuery, req, opts)
}

// ExpandItemCreate calls ant.merchant.expand.item.create.
// Official V3 contract: https://opendocs.alipay.com/open-v3/bf0fb8ad_ant.merchant.expand.item.create
func (s *Service) ExpandItemCreate(ctx context.Context, req ExpandItemCreateRequest, opts ...core.CallOption) (*core.Response[ExpandItemCreateResponse], error) {
	return execute[ExpandItemCreateRequest, ExpandItemCreateResponse](s, ctx, "ant.merchant.expand.item.create", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/create", core.ParameterModeJSONBody, req, opts)
}

// ExpandItemDelete calls ant.merchant.expand.item.delete.
// Official V3 contract: https://opendocs.alipay.com/open-v3/ac2de50f_ant.merchant.expand.item.delete
func (s *Service) ExpandItemDelete(ctx context.Context, req ExpandItemDeleteRequest, opts ...core.CallOption) (*core.Response[ExpandItemDeleteResponse], error) {
	return execute[ExpandItemDeleteRequest, ExpandItemDeleteResponse](s, ctx, "ant.merchant.expand.item.delete", core.HTTPMethodDelete, "/v3/ant/merchant/expand/item/delete", core.ParameterModeQuery, req, opts)
}

// ExpandItemModify calls ant.merchant.expand.item.modify.
// Official V3 contract: https://opendocs.alipay.com/open-v3/60d35f53_ant.merchant.expand.item.modify
func (s *Service) ExpandItemModify(ctx context.Context, req ExpandItemModifyRequest, opts ...core.CallOption) (*core.Response[ExpandItemModifyResponse], error) {
	return execute[ExpandItemModifyRequest, ExpandItemModifyResponse](s, ctx, "ant.merchant.expand.item.modify", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/modify", core.ParameterModeJSONBody, req, opts)
}

// ExpandItemOpenBatchQuery calls ant.merchant.expand.item.open.batchquery.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vjno
func (s *Service) ExpandItemOpenBatchQuery(ctx context.Context, req ExpandItemOpenBatchQueryRequest, opts ...core.CallOption) (*core.Response[ExpandItemOpenBatchQueryResponse], error) {
	return execute[ExpandItemOpenBatchQueryRequest, ExpandItemOpenBatchQueryResponse](s, ctx, "ant.merchant.expand.item.open.batchquery", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/open/batchquery", core.ParameterModeJSONBody, req, opts)
}

// ExpandItemOpenCreate calls ant.merchant.expand.item.open.create.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vn67
func (s *Service) ExpandItemOpenCreate(ctx context.Context, req ExpandItemOpenCreateRequest, opts ...core.CallOption) (*core.Response[ExpandItemOpenCreateResponse], error) {
	return execute[ExpandItemOpenCreateRequest, ExpandItemOpenCreateResponse](s, ctx, "ant.merchant.expand.item.open.create", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/open/create", core.ParameterModeJSONBody, req, opts)
}

// ExpandItemOpenDelete calls ant.merchant.expand.item.open.delete.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05v7ca
func (s *Service) ExpandItemOpenDelete(ctx context.Context, req ExpandItemOpenDeleteRequest, opts ...core.CallOption) (*core.Response[ExpandItemOpenDeleteResponse], error) {
	return execute[ExpandItemOpenDeleteRequest, ExpandItemOpenDeleteResponse](s, ctx, "ant.merchant.expand.item.open.delete", core.HTTPMethodDelete, "/v3/ant/merchant/expand/item/open/delete", core.ParameterModeQuery, req, opts)
}

// ExpandItemOpenModify calls ant.merchant.expand.item.open.modify.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vda6
func (s *Service) ExpandItemOpenModify(ctx context.Context, req ExpandItemOpenModifyRequest, opts ...core.CallOption) (*core.Response[ExpandItemOpenModifyResponse], error) {
	return execute[ExpandItemOpenModifyRequest, ExpandItemOpenModifyResponse](s, ctx, "ant.merchant.expand.item.open.modify", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/open/modify", core.ParameterModeJSONBody, req, opts)
}

// ExpandItemOpenQuery calls ant.merchant.expand.item.open.query.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vlue
func (s *Service) ExpandItemOpenQuery(ctx context.Context, req ExpandItemOpenQueryRequest, opts ...core.CallOption) (*core.Response[ExpandItemOpenQueryResponse], error) {
	return execute[ExpandItemOpenQueryRequest, ExpandItemOpenQueryResponse](s, ctx, "ant.merchant.expand.item.open.query", core.HTTPMethodGet, "/v3/ant/merchant/expand/item/open/query", core.ParameterModeQuery, req, opts)
}

// ExpandItemQuery calls ant.merchant.expand.item.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/aad82813_ant.merchant.expand.item.query
func (s *Service) ExpandItemQuery(ctx context.Context, req ExpandItemQueryRequest, opts ...core.CallOption) (*core.Response[ExpandItemQueryResponse], error) {
	return execute[ExpandItemQueryRequest, ExpandItemQueryResponse](s, ctx, "ant.merchant.expand.item.query", core.HTTPMethodGet, "/v3/ant/merchant/expand/item/query", core.ParameterModeQuery, req, opts)
}

// ExpandItemStatusModify calls ant.merchant.expand.item.status.modify.
// Official V3 contract: https://opendocs.alipay.com/open-v3/5cb8e04e_ant.merchant.expand.item.status.modify
func (s *Service) ExpandItemStatusModify(ctx context.Context, req ExpandItemStatusModifyRequest, opts ...core.CallOption) (*core.Response[ExpandItemStatusModifyResponse], error) {
	return execute[ExpandItemStatusModifyRequest, ExpandItemStatusModifyResponse](s, ctx, "ant.merchant.expand.item.status.modify", core.HTTPMethodPost, "/v3/ant/merchant/expand/item/status/modify", core.ParameterModeJSONBody, req, opts)
}

// ExpandShopConsult calls ant.merchant.expand.shop.consult.
// Official V3 contract: https://opendocs.alipay.com/apis/05vda1
func (s *Service) ExpandShopConsult(ctx context.Context, req ExpandShopConsultRequest, opts ...core.CallOption) (*core.Response[ExpandShopConsultResponse], error) {
	return execute[ExpandShopConsultRequest, ExpandShopConsultResponse](s, ctx, "ant.merchant.expand.shop.consult", core.HTTPMethodPost, "/v3/ant/merchant/expand/shop/consult", core.ParameterModeJSONBody, req, opts)
}

// ExpandShopReceiptAccountSave calls ant.merchant.expand.shop.receiptaccount.save.
// Official V3 contract: https://opendocs.alipay.com/open-v3/74728d59_ant.merchant.expand.shop.receiptaccount.save
func (s *Service) ExpandShopReceiptAccountSave(ctx context.Context, req ExpandShopReceiptAccountSaveRequest, opts ...core.CallOption) (*core.Response[ExpandShopReceiptAccountSaveResponse], error) {
	return execute[ExpandShopReceiptAccountSaveRequest, ExpandShopReceiptAccountSaveResponse](s, ctx, "ant.merchant.expand.shop.receiptaccount.save", core.HTTPMethodPost, "/v3/ant/merchant/expand/shop/receiptaccount/save", core.ParameterModeJSONBody, req, opts)
}

// ExpandIndirectImageUpload calls ant.merchant.expand.indirect.image.upload.
// Official V3 contract: https://opendocs.alipay.com/open-v3/af8e3a42_ant.merchant.expand.indirect.image.upload
func (s *Service) ExpandIndirectImageUpload(ctx context.Context, req ExpandIndirectImageUploadRequest, opts ...core.CallOption) (*core.Response[ExpandIndirectImageUploadResponse], error) {
	return execute[ExpandIndirectImageUploadRequest, ExpandIndirectImageUploadResponse](s, ctx, "ant.merchant.expand.indirect.image.upload", core.HTTPMethodPost, "/v3/ant/merchant/indirect/image", core.ParameterModeMultipart, req, opts)
}

// ExpandMCCQuery calls ant.merchant.expand.mcc.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/b00ab2d2_ant.merchant.expand.mcc.query
func (s *Service) ExpandMCCQuery(ctx context.Context, req ExpandMCCQueryRequest, opts ...core.CallOption) (*core.Response[ExpandMCCQueryResponse], error) {
	return execute[ExpandMCCQueryRequest, ExpandMCCQueryResponse](s, ctx, "ant.merchant.expand.mcc.query", core.HTTPMethodGet, "/v3/ant/merchant/mcc/query", core.ParameterModeQuery, req, opts)
}

// ExpandOrderQuery calls ant.merchant.expand.order.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/dd6e9adb_ant.merchant.expand.order.query
func (s *Service) ExpandOrderQuery(ctx context.Context, req ExpandOrderQueryRequest, opts ...core.CallOption) (*core.Response[ExpandOrderQueryResponse], error) {
	return execute[ExpandOrderQueryRequest, ExpandOrderQueryResponse](s, ctx, "ant.merchant.expand.order.query", core.HTTPMethodGet, "/v3/ant/merchant/order/{order_id}", core.ParameterModeQuery, req, opts)
}

// ExpandShopQuery calls ant.merchant.expand.shop.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/b80b5948_ant.merchant.expand.shop.query
func (s *Service) ExpandShopQuery(ctx context.Context, req ExpandShopQueryRequest, opts ...core.CallOption) (*core.Response[ExpandShopQueryResponse], error) {
	return execute[ExpandShopQueryRequest, ExpandShopQueryResponse](s, ctx, "ant.merchant.expand.shop.query", core.HTTPMethodGet, "/v3/ant/merchant/shop", core.ParameterModeQuery, req, opts)
}

// ExpandShopModify calls ant.merchant.expand.shop.modify.
// Official V3 contract: https://opendocs.alipay.com/open-v3/b1a37213_ant.merchant.expand.shop.modify
func (s *Service) ExpandShopModify(ctx context.Context, req ExpandShopModifyRequest, opts ...core.CallOption) (*core.Response[ExpandShopModifyResponse], error) {
	return execute[ExpandShopModifyRequest, ExpandShopModifyResponse](s, ctx, "ant.merchant.expand.shop.modify", core.HTTPMethodPatch, "/v3/ant/merchant/shop", core.ParameterModeJSONBody, req, opts)
}

// ExpandShopCreate calls ant.merchant.expand.shop.create.
// Official V3 contract: https://opendocs.alipay.com/open-v3/99b91969_ant.merchant.expand.shop.create
func (s *Service) ExpandShopCreate(ctx context.Context, req ExpandShopCreateRequest, opts ...core.CallOption) (*core.Response[ExpandShopCreateResponse], error) {
	return execute[ExpandShopCreateRequest, ExpandShopCreateResponse](s, ctx, "ant.merchant.expand.shop.create", core.HTTPMethodPost, "/v3/ant/merchant/shop", core.ParameterModeJSONBody, req, opts)
}

// ExpandShopClose calls ant.merchant.expand.shop.close.
// Official V3 contract: https://opendocs.alipay.com/open-v3/4950631b_ant.merchant.expand.shop.close
func (s *Service) ExpandShopClose(ctx context.Context, req ExpandShopCloseRequest, opts ...core.CallOption) (*core.Response[ExpandShopCloseResponse], error) {
	return execute[ExpandShopCloseRequest, ExpandShopCloseResponse](s, ctx, "ant.merchant.expand.shop.close", core.HTTPMethodPatch, "/v3/ant/merchant/shop/close", core.ParameterModeJSONBody, req, opts)
}

// ExpandShopPageQuery calls ant.merchant.expand.shop.page.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/fe443839_ant.merchant.expand.shop.page.query
func (s *Service) ExpandShopPageQuery(ctx context.Context, req ExpandShopPageQueryRequest, opts ...core.CallOption) (*core.Response[ExpandShopPageQueryResponse], error) {
	return execute[ExpandShopPageQueryRequest, ExpandShopPageQueryResponse](s, ctx, "ant.merchant.expand.shop.page.query", core.HTTPMethodGet, "/v3/ant/merchant/shop/pagequery", core.ParameterModeQuery, req, opts)
}
