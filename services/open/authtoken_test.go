package open_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/open"
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

func newService(t *testing.T) (*open.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{result: &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(`{}`),
	}}
	svc, err := open.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestAuthTokenAppSpecAndMapping(t *testing.T) {
	svc, executor := newService(t)
	executor.result.Node = []byte(`{"app_auth_token":"aat-1","app_refresh_token":"art-1","auth_app_id":"2021000000000002","user_id":"2088101117955611","tokens":[{"app_auth_token":"aat-2","auth_app_id":"2021000000000003"}]}`)
	resp, err := svc.AuthTokenApp(context.Background(), open.AuthTokenAppRequest{
		GrantType: open.GrantTypeAuthorizationCode,
		Code:      ptr("app-auth-code"),
	})
	if err != nil {
		t.Fatalf("AuthTokenApp() error = %v", err)
	}
	if executor.req.Spec.Method != "alipay.open.auth.token.app" ||
		executor.req.Spec.Path != "/v3/alipay/open/auth/token/app" {
		t.Fatalf("spec = %+v", executor.req.Spec)
	}
	if resp.Data.AppAuthToken == nil || *resp.Data.AppAuthToken != "aat-1" {
		t.Fatalf("app_auth_token missing: %+v", resp.Data)
	}
	if len(resp.Data.Tokens) != 1 || *resp.Data.Tokens[0].AppAuthToken != "aat-2" {
		t.Fatalf("tokens missing: %+v", resp.Data)
	}
}

func TestAuthTokenAppValidations(t *testing.T) {
	svc, executor := newService(t)
	if _, err := svc.AuthTokenApp(context.Background(), open.AuthTokenAppRequest{
		GrantType: open.GrantTypeRefreshToken}); err == nil {
		t.Fatal("refresh_token is required")
	}
	if executor.called {
		t.Error("validation failure must not send request")
	}
}

func ptr[T any](v T) *T { return &v }
