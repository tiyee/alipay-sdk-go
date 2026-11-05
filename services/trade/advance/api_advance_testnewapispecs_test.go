package advance_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/advance"
)

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*advance.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "订单咨询",
			call: func(s *advance.Service, e *recordingExecutor) error {
				_, err := s.AdvanceConsult(context.Background(), advance.AdvanceConsultRequest{
					AlipayUserID: ptr("2088101117955611")})
				return err
			},
			wantMethod: "alipay.trade.advance.consult",
			wantPath:   "/v3/alipay/trade/advance/consult",
			wantGET:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{}`)
			if err := tt.call(svc, executor); err != nil {
				t.Fatalf("call error = %v", err)
			}
			spec := executor.req.Spec
			if spec.Method != tt.wantMethod || spec.Path != tt.wantPath {
				t.Fatalf("spec = %s %s, want %s %s", spec.Method, spec.Path, tt.wantMethod, tt.wantPath)
			}
			wantMethod := core.HTTPMethodPost
			if tt.wantGET {
				wantMethod = core.HTTPMethodGet
			}
			if spec.HTTPMethod != wantMethod {
				t.Errorf("HTTPMethod = %v, want %v", spec.HTTPMethod, wantMethod)
			}
			wantMode := core.ParameterModeJSONBody
			if tt.wantGET {
				wantMode = core.ParameterModeQuery
			}
			if spec.ParameterMode != wantMode || executor.req.Params.Mode != wantMode {
				t.Errorf("parameter mode = %v / %v, want %v", spec.ParameterMode, executor.req.Params.Mode, wantMode)
			}
			if executor.result.HTTPStatus != 200 {
				t.Error("result must be passed through")
			}
		})
	}
}

func TestRiskConsultAgreementAndRepaymentAmounts(t *testing.T) {
	svc, e := newRecordingService(t)
	req := advance.AdvanceConsultRequest{ConsultScene: ptr(advance.AdvanceSceneOrderRiskEvaluation)}
	for _, agreement := range []*string{nil, ptr("")} {
		req.AgreementNo = agreement
		_, err := svc.AdvanceConsult(t.Context(), req)
		var v *core.ValidationError
		if !errors.As(err, &v) || v.Field != "agreement_no" || e.called {
			t.Fatalf("invalid agreement reached executor: %v", err)
		}
	}
	req.AgreementNo = ptr("agreement")
	req.EstimatedOrderAmount = ptr("100000000.01")
	_, err := svc.AdvanceConsult(t.Context(), req)
	var val *core.ValidationError
	if !errors.As(err, &val) || val.Field != "estimated_order_amount" || e.called {
		t.Fatalf("invalid amount reached executor: %v", err)
	}
	req.EstimatedOrderAmount = ptr("100000000.00")
	e.result = verifiedResult(`{"wait_repayment_amount":"10.00","wait_repayment_order_infos":[{"wait_repayment_amount":9007199254740993}]}`)
	resp, err := svc.AdvanceConsult(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if e.req.Params.Form.Get("agreement_no") != "agreement" || resp.Data.WaitRepaymentAmount == nil || *resp.Data.WaitRepaymentAmount != "10.00" || len(resp.Data.WaitRepaymentOrderInfos) != 1 {
		t.Fatal("missing request or response fields")
	}
	amount := resp.Data.WaitRepaymentOrderInfos[0].WaitRepaymentAmount
	if amount == nil || *amount != 9007199254740993 {
		t.Fatalf("unexpected detail amount %v", amount)
	}
}
