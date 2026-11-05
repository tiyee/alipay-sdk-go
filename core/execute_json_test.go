package core_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
)

type executeJSONRequest struct {
	Value string `json:"value" query:"value"`
}

type executeJSONResponse struct {
	Value string `json:"value"`
}

type executeJSONExecutor func(context.Context, core.WireRequest) (*core.APIResult, error)

func (f executeJSONExecutor) Do(ctx context.Context, req core.WireRequest) (*core.APIResult, error) {
	return f(ctx, req)
}

func TestExecuteJSONForwardsRequestAndDecodesResponse(t *testing.T) {
	for _, tc := range []struct {
		name   string
		method core.HTTPMethod
		mode   core.ParameterMode
	}{
		{"json", core.HTTPMethodPost, core.ParameterModeJSONBody},
		{"query", core.HTTPMethodGet, core.ParameterModeQuery},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			spec := core.OperationSpec{
				Method: "custom.demo.query", HTTPMethod: tc.method, Path: "/v3/custom/demo/query",
				ParameterMode: tc.mode, AuthRequirement: core.AuthRequirementCertificate,
			}
			settings := core.NewCallSettings(core.WithAppAuthToken("test-token"))
			result := &core.APIResult{HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{"value":"response"}`)}
			called := false
			executor := executeJSONExecutor(func(gotCtx context.Context, wire core.WireRequest) (*core.APIResult, error) {
				called = true
				if gotCtx != ctx || wire.Spec != spec || wire.Settings != settings || wire.Params.Mode != tc.mode {
					t.Fatalf("context, spec, settings or mode changed: %+v", wire)
				}
				switch tc.mode {
				case core.ParameterModeJSONBody:
					if string(wire.Params.JSON) != `{"value":"中文 +%"}` {
						t.Fatalf("JSON = %s", wire.Params.JSON)
					}
				case core.ParameterModeQuery:
					if wire.Params.Form.Get("value") != "中文 +%" {
						t.Fatalf("query = %v", wire.Params.Form)
					}
				}
				return result, nil
			})
			resp, err := core.ExecuteJSON[executeJSONRequest, executeJSONResponse](ctx, executor, spec,
				core.Request[executeJSONRequest]{Body: executeJSONRequest{Value: "中文 +%"}, Options: settings})
			if err != nil {
				t.Fatal(err)
			}
			if !called || resp == nil || resp.Data == nil || resp.Data.Value != "response" || resp.Result != result {
				t.Fatalf("response = %+v, executor called = %v", resp, called)
			}
		})
	}
}

func TestExecuteJSONRejectsInvalidSpecBeforeExecution(t *testing.T) {
	for _, tc := range []struct {
		name string
		spec core.OperationSpec
	}{
		{"missing method", core.OperationSpec{HTTPMethod: core.HTTPMethodPost, Path: "/v3/custom/demo", AuthRequirement: core.AuthRequirementAny}},
		{"missing auth requirement", core.OperationSpec{Method: "custom.demo", HTTPMethod: core.HTTPMethodPost, Path: "/v3/custom/demo"}},
		{"invalid mode", core.OperationSpec{Method: "custom.demo", HTTPMethod: core.HTTPMethodPost, Path: "/v3/custom/demo", ParameterMode: core.ParameterMode(255), AuthRequirement: core.AuthRequirementAny}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			executor := &mockExecutor{}
			resp, err := core.ExecuteJSON[executeJSONRequest, executeJSONResponse](context.Background(), executor, tc.spec,
				core.Request[executeJSONRequest]{})
			if _, ok := compat.AsType[*core.ValidationError](err); !ok || resp != nil || executor.called {
				t.Fatalf("response = %+v, error = %v, executor called = %v", resp, err, executor.called)
			}
		})
	}
}

func TestExecuteJSONPreservesFailureSemantics(t *testing.T) {
	spec := core.OperationSpec{
		Method: "custom.demo", HTTPMethod: core.HTTPMethodPost, Path: "/v3/custom/demo",
		AuthRequirement: core.AuthRequirementAny,
	}
	t.Run("executor error preserves cause and metadata", func(t *testing.T) {
		cause := errors.New("test transport error")
		result := &core.APIResult{HTTPStatus: 503}
		executor := &mockExecutor{result: result, err: cause}
		resp, err := core.ExecuteJSON[executeJSONRequest, executeJSONResponse](context.Background(), executor, spec,
			core.Request[executeJSONRequest]{})
		if !errors.Is(err, cause) || resp == nil || resp.Data != nil || resp.Result != result {
			t.Fatalf("response = %+v, error = %v", resp, err)
		}
	})
	t.Run("unsigned response is rejected before decoding", func(t *testing.T) {
		result := &core.APIResult{HTTPStatus: 200, Node: []byte(`invalid JSON`)}
		executor := &mockExecutor{result: result}
		resp, err := core.ExecuteJSON[executeJSONRequest, executeJSONResponse](context.Background(), executor, spec,
			core.Request[executeJSONRequest]{})
		if _, ok := compat.AsType[*core.SignatureError](err); !ok || resp == nil || resp.Data != nil || resp.Result != result {
			t.Fatalf("response = %+v, error = %v", resp, err)
		}
	})
}

// ExampleExecuteJSON demonstrates a V3 request using the default codecs.
// Official contract: https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query
func ExampleExecuteJSON() {
	type QueryRequest struct {
		OutTradeNo string `json:"out_trade_no"`
	}
	type QueryResponse struct {
		TradeNo string `json:"trade_no"`
	}
	query := func(ctx context.Context, executor core.Executor, outTradeNo string) (*core.Response[QueryResponse], error) {
		return core.ExecuteJSON[QueryRequest, QueryResponse](ctx, executor, core.OperationSpec{
			Method: "alipay.trade.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/query",
			ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
		}, core.Request[QueryRequest]{Body: QueryRequest{OutTradeNo: outTradeNo}})
	}
	_ = query // Pass a configured *core.Client or another core.Executor.
}
