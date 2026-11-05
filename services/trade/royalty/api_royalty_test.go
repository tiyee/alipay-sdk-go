package royalty_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/royalty"
)

// recordingExecutor 记录 WireRequest，返回预置结果，
// 用于验证接口名、参数位置与响应映射（不发起网络请求）。
type recordingExecutor struct {
	req    core.WireRequest
	called bool

	result *core.APIResult
	err    error
}

func (e *recordingExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.called = true
	e.req = req
	return e.result, e.err
}

func verifiedResult(node string) *core.APIResult {
	return &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(node),
	}
}

func newRecordingService(t *testing.T) (*royalty.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := royalty.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := royalty.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func ptr[T any](v T) *T { return &v }

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*royalty.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "分账比例查询",
			call: func(s *royalty.Service, e *recordingExecutor) error {
				_, err := s.RoyaltyRateQuery(context.Background(), royalty.RoyaltyRateQueryRequest{
					OutRequestNo: "R1"})
				return err
			},
			wantMethod: "alipay.trade.royalty.rate.query",
			wantPath:   "/v3/alipay/trade/royalty/rate/query",
			wantGET:    true,
		},
		{
			name: "分账关系绑定",
			call: func(s *royalty.Service, e *recordingExecutor) error {
				_, err := s.RoyaltyRelationBind(context.Background(), royalty.RoyaltyRelationBindRequest{
					OutRequestNo: "R1",
					ReceiverList: []royalty.RoyaltyEntity{{Type: royalty.RoyaltyEntityTypeUserID, Account: ptr("2088001")}}})
				return err
			},
			wantMethod: "alipay.trade.royalty.relation.bind",
			wantPath:   "/v3/alipay/trade/royalty/relation/bind",
		},
		{
			name: "分账关系解绑",
			call: func(s *royalty.Service, e *recordingExecutor) error {
				_, err := s.RoyaltyRelationUnbind(context.Background(), royalty.RoyaltyRelationUnbindRequest{
					OutRequestNo: "R1",
					ReceiverList: []royalty.RoyaltyEntity{{Type: royalty.RoyaltyEntityTypeUserID, Account: ptr("2088001")}}})
				return err
			},
			wantMethod: "alipay.trade.royalty.relation.unbind",
			wantPath:   "/v3/alipay/trade/royalty/relation/unbind",
		},
		{
			name: "分账关系查询",
			call: func(s *royalty.Service, e *recordingExecutor) error {
				_, err := s.RoyaltyRelationBatchQuery(context.Background(), royalty.RoyaltyRelationBatchQueryRequest{
					OutRequestNo: "R1"})
				return err
			},
			wantMethod: "alipay.trade.royalty.relation.batchquery",
			wantPath:   "/v3/alipay/trade/royalty/relation/batchquery",
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

// TestNewAPIQuerySerialization 验证 GET 接口的 query 参数编码：
// 非空指针参与编码，空 omitempty 字段省略。
func TestNewAPIQuerySerialization(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"user_id":"2088101117955611","max_ratio":100}`)
	resp, err := svc.RoyaltyRateQuery(context.Background(), royalty.RoyaltyRateQueryRequest{
		OutRequestNo: "20261105001"})
	if err != nil {
		t.Fatalf("RoyaltyRateQuery() error = %v", err)
	}
	if got := executor.req.Params.Form.Encode(); got != "out_request_no=20261105001" {
		t.Fatalf("query = %q", got)
	}
	if resp.Data == nil || resp.Data.MaxRatio == nil || *resp.Data.MaxRatio != 100 {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
}

// TestNewAPIValidations 覆盖新接口的请求校验分支。
func TestNewAPIValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*royalty.Service) error
	}{
		{
			name: "分账关系绑定超限",
			call: func(s *royalty.Service) error {
				list := make([]royalty.RoyaltyEntity, 21)
				for i := range list {
					list[i] = royalty.RoyaltyEntity{Type: royalty.RoyaltyEntityTypeUserID, Account: ptr("2088001")}
				}
				_, err := s.RoyaltyRelationBind(context.Background(), royalty.RoyaltyRelationBindRequest{
					OutRequestNo: "R1", ReceiverList: list})
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			err := tt.call(svc)
			var valErr *core.ValidationError
			if !errors.As(err, &valErr) {
				t.Fatalf("error = %v, want *core.ValidationError", err)
			}
			if executor.called {
				t.Error("validation failure must not send request")
			}
		})
	}
}
