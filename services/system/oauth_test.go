package system_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/system"
)

type recordingExecutor struct {
	req    core.WireRequest
	called bool
	result *core.APIResult
}

func (e *recordingExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.called = true
	e.req = req
	return e.result, nil
}

func newService(t *testing.T) (*system.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{result: &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(`{}`),
	}}
	svc, err := system.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestOAuthTokenSpecAndMapping(t *testing.T) {
	svc, executor := newService(t)
	executor.result.Node = []byte(`{"access_token":"at-1","refresh_token":"rt-1","expires_in":"3600","user_id":"2088101117955611","open_id":"open-1"}`)
	resp, err := svc.OAuthToken(context.Background(), system.OAuthTokenRequest{
		GrantType: system.GrantTypeAuthorizationCode,
		Code:      ptr("auth-code-1"),
	})
	if err != nil {
		t.Fatalf("OAuthToken() error = %v", err)
	}
	spec := executor.req.Spec
	if spec.Method != "alipay.system.oauth.token" || spec.Path != "/v3/alipay/system/oauth/token" {
		t.Fatalf("spec = %s %s", spec.Method, spec.Path)
	}
	if spec.HTTPMethod != core.HTTPMethodPost ||
		spec.ParameterMode != core.ParameterModeJSONBody {
		t.Fatalf("unexpected protocol metadata: %+v", spec)
	}
	if string(executor.req.Params.JSON) != `{"grant_type":"authorization_code","code":"auth-code-1"}` {
		t.Fatalf("body = %s", executor.req.Params.JSON)
	}
	if resp.Data.AccessToken != "at-1" || resp.Data.RefreshToken == nil || *resp.Data.RefreshToken != "rt-1" {
		t.Fatalf("unexpected tokens: %+v", resp.Data)
	}
	if resp.Data.OpenID == nil || *resp.Data.OpenID != "open-1" {
		t.Fatalf("open_id missing: %+v", resp.Data)
	}
}

func TestOAuthTokenValidations(t *testing.T) {
	tests := []struct {
		name string
		req  system.OAuthTokenRequest
	}{
		{"缺少 grant_type", system.OAuthTokenRequest{Code: ptr("c")}},
		{"未知 grant_type", system.OAuthTokenRequest{GrantType: system.GrantType("password")}},
		{"授权码方式缺少 code", system.OAuthTokenRequest{GrantType: system.GrantTypeAuthorizationCode}},
		{"刷新方式缺少 refresh_token", system.OAuthTokenRequest{GrantType: system.GrantTypeRefreshToken}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newService(t)
			_, err := svc.OAuthToken(context.Background(), tt.req)
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

func TestNewServiceAcceptsPlainExecutor(t *testing.T) {
	// SDK 已是 V3-only：Executor 不再携带协议元数据，
	// 任何实现 Do 的执行器均可接入 Service。
	if _, err := system.NewService(&fakeExecutor{}); err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
}

type fakeExecutor struct{}

func (e *fakeExecutor) Do(context.Context, core.WireRequest) (*core.APIResult, error) {
	return nil, nil
}

func ptr[T any](v T) *T { return &v }
