package credit

import (
	"context"
	"github.com/tiyee/alipay-sdk-go/core"
	"testing"
)

type recordingExecutor struct{ req core.WireRequest }

func (e *recordingExecutor) Do(_ context.Context, r core.WireRequest) (*core.APIResult, error) {
	e.req = r
	return &core.APIResult{HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{}`)}, nil
}
func TestOperationSpecs(t *testing.T) {
	cases := []struct {
		operation, method, path string
		mode                    core.ParameterMode
		call                    func(*Service) error
	}{
		{"zhima.credit.payafteruse.creditagreement.query", "GET", "/v3/zhima/credit/payafteruse/creditagreement/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.PayAfterUseCreditAgreementQuery(context.Background(), PayAfterUseCreditAgreementQueryRequest{})
			return err
		}},
		{"zhima.credit.payafteruse.creditagreement.transfer", "POST", "/v3/zhima/credit/payafteruse/creditagreement/transfer", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PayAfterUseCreditAgreementTransfer(context.Background(), PayAfterUseCreditAgreementTransferRequest{})
			return err
		}},
		{"zhima.credit.payafteruse.creditbizorder.finish", "POST", "/v3/zhima/credit/payafteruse/creditbizorder/finish", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PayAfterUseCreditBizOrderFinish(context.Background(), PayAfterUseCreditBizOrderFinishRequest{})
			return err
		}},
		{"zhima.credit.payafteruse.creditbizorder.order", "POST", "/v3/zhima/credit/payafteruse/creditbizorder/order", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PayAfterUseCreditBizOrderOrder(context.Background(), PayAfterUseCreditBizOrderOrderRequest{})
			return err
		}},
		{"zhima.credit.payafteruse.creditbizorder.query", "GET", "/v3/zhima/credit/payafteruse/creditbizorder/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.PayAfterUseCreditBizOrderQuery(context.Background(), PayAfterUseCreditBizOrderQueryRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.agreement.query", "GET", "/v3/zhima/credit/pe/zmgo/agreement/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.PeZmGoAgreementQuery(context.Background(), PeZmGoAgreementQueryRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.agreement.unsign", "POST", "/v3/zhima/credit/pe/zmgo/agreement/unsign", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoAgreementUnsign(context.Background(), PeZmGoAgreementUnsignRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.bizopt.close", "POST", "/v3/zhima/credit/pe/zmgo/bizopt/close", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoBizOptClose(context.Background(), PeZmGoBizOptCloseRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.cumulation.sync", "POST", "/v3/zhima/credit/pe/zmgo/cumulation/sync", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoCumulationSync(context.Background(), PeZmGoCumulationSyncRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.preorder.create", "PUT", "/v3/zhima/credit/pe/zmgo/preorder/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoPreOrderCreate(context.Background(), PeZmGoPreOrderCreateRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.settle.apply", "POST", "/v3/zhima/credit/pe/zmgo/settle/apply", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoSettleApply(context.Background(), PeZmGoSettleApplyRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.settle.refund", "POST", "/v3/zhima/credit/pe/zmgo/settle/refund", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoSettleRefund(context.Background(), PeZmGoSettleRefundRequest{})
			return err
		}},
		{"zhima.credit.pe.zmgo.settle.unfreeze", "PUT", "/v3/zhima/credit/pe/zmgo/settle/unfreeze", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.PeZmGoSettleUnfreeze(context.Background(), PeZmGoSettleUnfreezeRequest{})
			return err
		}}}
	for _, tc := range cases {
		t.Run(tc.operation, func(t *testing.T) {
			e := &recordingExecutor{}
			s, err := NewService(e)
			if err != nil {
				t.Fatal(err)
			}
			if err = tc.call(s); err != nil {
				t.Fatal(err)
			}
			if e.req.Spec.Method != tc.operation || e.req.Spec.HTTPMethod.String() != tc.method || e.req.Spec.Path != tc.path || e.req.Spec.ParameterMode != tc.mode || e.req.Params.Mode != tc.mode {
				t.Fatalf("unexpected spec: %+v", e.req.Spec)
			}
		})
	}
}
