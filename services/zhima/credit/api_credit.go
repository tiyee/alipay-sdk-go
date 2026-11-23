// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package credit

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

// PayAfterUseCreditAgreementQuery calls zhima.credit.payafteruse.creditagreement.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/eaff9140_zhima.credit.payafteruse.creditagreement.query
func (s *Service) PayAfterUseCreditAgreementQuery(ctx context.Context, req PayAfterUseCreditAgreementQueryRequest, opts ...core.CallOption) (*core.Response[PayAfterUseCreditAgreementQueryResponse], error) {
	return execute[PayAfterUseCreditAgreementQueryRequest, PayAfterUseCreditAgreementQueryResponse](s, ctx, "zhima.credit.payafteruse.creditagreement.query", core.HTTPMethodGet, "/v3/zhima/credit/payafteruse/creditagreement/query", core.ParameterModeQuery, req, opts)
}

// PayAfterUseCreditAgreementTransfer calls zhima.credit.payafteruse.creditagreement.transfer.
// Official V3 contract: https://opendocs.alipay.com/apis/05vdam
func (s *Service) PayAfterUseCreditAgreementTransfer(ctx context.Context, req PayAfterUseCreditAgreementTransferRequest, opts ...core.CallOption) (*core.Response[PayAfterUseCreditAgreementTransferResponse], error) {
	return execute[PayAfterUseCreditAgreementTransferRequest, PayAfterUseCreditAgreementTransferResponse](s, ctx, "zhima.credit.payafteruse.creditagreement.transfer", core.HTTPMethodPost, "/v3/zhima/credit/payafteruse/creditagreement/transfer", core.ParameterModeJSONBody, req, opts)
}

// PayAfterUseCreditBizOrderFinish calls zhima.credit.payafteruse.creditbizorder.finish.
// Official V3 contract: https://opendocs.alipay.com/open-v3/7972514c_zhima.credit.payafteruse.creditbizorder.finish?scene=common&pathHash=ba647372
func (s *Service) PayAfterUseCreditBizOrderFinish(ctx context.Context, req PayAfterUseCreditBizOrderFinishRequest, opts ...core.CallOption) (*core.Response[PayAfterUseCreditBizOrderFinishResponse], error) {
	return execute[PayAfterUseCreditBizOrderFinishRequest, PayAfterUseCreditBizOrderFinishResponse](s, ctx, "zhima.credit.payafteruse.creditbizorder.finish", core.HTTPMethodPost, "/v3/zhima/credit/payafteruse/creditbizorder/finish", core.ParameterModeJSONBody, req, opts)
}

// PayAfterUseCreditBizOrderOrder calls zhima.credit.payafteruse.creditbizorder.order.
// Official V3 contract: https://opendocs.alipay.com/open-v3/51ba4ec1_zhima.credit.payafteruse.creditbizorder.order
func (s *Service) PayAfterUseCreditBizOrderOrder(ctx context.Context, req PayAfterUseCreditBizOrderOrderRequest, opts ...core.CallOption) (*core.Response[PayAfterUseCreditBizOrderOrderResponse], error) {
	return execute[PayAfterUseCreditBizOrderOrderRequest, PayAfterUseCreditBizOrderOrderResponse](s, ctx, "zhima.credit.payafteruse.creditbizorder.order", core.HTTPMethodPost, "/v3/zhima/credit/payafteruse/creditbizorder/order", core.ParameterModeJSONBody, req, opts)
}

// PayAfterUseCreditBizOrderQuery calls zhima.credit.payafteruse.creditbizorder.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/0f11f105_zhima.credit.payafteruse.creditbizorder.query?scene=common&pathHash=203d9491
func (s *Service) PayAfterUseCreditBizOrderQuery(ctx context.Context, req PayAfterUseCreditBizOrderQueryRequest, opts ...core.CallOption) (*core.Response[PayAfterUseCreditBizOrderQueryResponse], error) {
	return execute[PayAfterUseCreditBizOrderQueryRequest, PayAfterUseCreditBizOrderQueryResponse](s, ctx, "zhima.credit.payafteruse.creditbizorder.query", core.HTTPMethodGet, "/v3/zhima/credit/payafteruse/creditbizorder/query", core.ParameterModeQuery, req, opts)
}

// PeZmGoAgreementQuery calls zhima.credit.pe.zmgo.agreement.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/849a96d9_zhima.credit.pe.zmgo.agreement.query
func (s *Service) PeZmGoAgreementQuery(ctx context.Context, req PeZmGoAgreementQueryRequest, opts ...core.CallOption) (*core.Response[PeZmGoAgreementQueryResponse], error) {
	return execute[PeZmGoAgreementQueryRequest, PeZmGoAgreementQueryResponse](s, ctx, "zhima.credit.pe.zmgo.agreement.query", core.HTTPMethodGet, "/v3/zhima/credit/pe/zmgo/agreement/query", core.ParameterModeQuery, req, opts)
}

// PeZmGoAgreementUnsign calls zhima.credit.pe.zmgo.agreement.unsign.
// Official V3 contract: https://opendocs.alipay.com/open-v3/4d77e884_zhima.credit.pe.zmgo.agreement.unsign
func (s *Service) PeZmGoAgreementUnsign(ctx context.Context, req PeZmGoAgreementUnsignRequest, opts ...core.CallOption) (*core.Response[PeZmGoAgreementUnsignResponse], error) {
	return execute[PeZmGoAgreementUnsignRequest, PeZmGoAgreementUnsignResponse](s, ctx, "zhima.credit.pe.zmgo.agreement.unsign", core.HTTPMethodPost, "/v3/zhima/credit/pe/zmgo/agreement/unsign", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoBizOptClose calls zhima.credit.pe.zmgo.bizopt.close.
// Official V3 contract: https://opendocs.alipay.com/apis/05vluc
func (s *Service) PeZmGoBizOptClose(ctx context.Context, req PeZmGoBizOptCloseRequest, opts ...core.CallOption) (*core.Response[PeZmGoBizOptCloseResponse], error) {
	return execute[PeZmGoBizOptCloseRequest, PeZmGoBizOptCloseResponse](s, ctx, "zhima.credit.pe.zmgo.bizopt.close", core.HTTPMethodPost, "/v3/zhima/credit/pe/zmgo/bizopt/close", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoCumulationSync calls zhima.credit.pe.zmgo.cumulation.sync.
// Official V3 contract: https://opendocs.alipay.com/apis/05v7c8
func (s *Service) PeZmGoCumulationSync(ctx context.Context, req PeZmGoCumulationSyncRequest, opts ...core.CallOption) (*core.Response[PeZmGoCumulationSyncResponse], error) {
	return execute[PeZmGoCumulationSyncRequest, PeZmGoCumulationSyncResponse](s, ctx, "zhima.credit.pe.zmgo.cumulation.sync", core.HTTPMethodPost, "/v3/zhima/credit/pe/zmgo/cumulation/sync", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoPreOrderCreate calls zhima.credit.pe.zmgo.preorder.create.
// Official V3 contract: https://opendocs.alipay.com/open-v3/3a385c23_zhima.credit.pe.zmgo.preorder.create
func (s *Service) PeZmGoPreOrderCreate(ctx context.Context, req PeZmGoPreOrderCreateRequest, opts ...core.CallOption) (*core.Response[PeZmGoPreOrderCreateResponse], error) {
	return execute[PeZmGoPreOrderCreateRequest, PeZmGoPreOrderCreateResponse](s, ctx, "zhima.credit.pe.zmgo.preorder.create", core.HTTPMethodPut, "/v3/zhima/credit/pe/zmgo/preorder/create", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoSettleApply calls zhima.credit.pe.zmgo.settle.apply.
// Official V3 contract: https://opendocs.alipay.com/open-v3/3a0505ce_zhima.credit.pe.zmgo.settle.apply
func (s *Service) PeZmGoSettleApply(ctx context.Context, req PeZmGoSettleApplyRequest, opts ...core.CallOption) (*core.Response[PeZmGoSettleApplyResponse], error) {
	return execute[PeZmGoSettleApplyRequest, PeZmGoSettleApplyResponse](s, ctx, "zhima.credit.pe.zmgo.settle.apply", core.HTTPMethodPost, "/v3/zhima/credit/pe/zmgo/settle/apply", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoSettleRefund calls zhima.credit.pe.zmgo.settle.refund.
// Official V3 contract: https://opendocs.alipay.com/open-v3/a66ab3b0_zhima.credit.pe.zmgo.settle.refund
func (s *Service) PeZmGoSettleRefund(ctx context.Context, req PeZmGoSettleRefundRequest, opts ...core.CallOption) (*core.Response[PeZmGoSettleRefundResponse], error) {
	return execute[PeZmGoSettleRefundRequest, PeZmGoSettleRefundResponse](s, ctx, "zhima.credit.pe.zmgo.settle.refund", core.HTTPMethodPost, "/v3/zhima/credit/pe/zmgo/settle/refund", core.ParameterModeJSONBody, req, opts)
}

// PeZmGoSettleUnfreeze calls zhima.credit.pe.zmgo.settle.unfreeze.
// Official V3 contract: https://opendocs.alipay.com/apis/05gzu9
func (s *Service) PeZmGoSettleUnfreeze(ctx context.Context, req PeZmGoSettleUnfreezeRequest, opts ...core.CallOption) (*core.Response[PeZmGoSettleUnfreezeResponse], error) {
	return execute[PeZmGoSettleUnfreezeRequest, PeZmGoSettleUnfreezeResponse](s, ctx, "zhima.credit.pe.zmgo.settle.unfreeze", core.HTTPMethodPut, "/v3/zhima/credit/pe/zmgo/settle/unfreeze", core.ParameterModeJSONBody, req, opts)
}
