package core_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
)

func TestSignedRequestsRefuseRedirects(t *testing.T) {
	destination := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("redirect destination received a signed request")
	}))
	defer destination.Close()
	source := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, strings.Replace(destination.URL, "127.0.0.1", "localhost", 1), http.StatusTemporaryRedirect)
	}))
	defer source.Close()
	supplied := &http.Client{}
	client := newTestClient(t, source.URL, option.WithHTTPClient(supplied))
	resp, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{Options: core.NewCallSettings(core.WithAppAuthToken("dummy-token"))})
	if e, ok := compat.AsType[*core.HTTPError](err); !ok || e.Status != 307 || resp.Result.HTTPStatus != 307 {
		t.Fatalf("redirect result = %v, %v", resp, err)
	}
	if supplied.CheckRedirect != nil {
		t.Fatal("SDK mutated caller's redirect policy")
	}
}

func TestWrappedDefaultTransportAndSafeError(t *testing.T) {
	original := http.DefaultTransport
	defer func() { http.DefaultTransport = original }()
	http.DefaultTransport = contextTransport(func(*http.Request) (*http.Response, error) { return nil, context.Canceled })
	c := newTestClient(t, "https://offline.invalid")
	_, err := core.Execute(t.Context(), c, newBillOperation(t), core.Request[billRequest]{Body: billRequest{BillType: "sensitive-marker"}})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("context error lost: %v", err)
	}
	if _, ok := compat.AsType[*url.Error](err); !ok {
		t.Fatal("transport cause lost")
	}
	if strings.Contains(err.Error(), "sensitive-marker") || strings.Contains(err.Error(), "offline.invalid") {
		t.Error("transport error exposed its request URL")
	}
}

func TestSandboxUsesTLS(t *testing.T) {
	if core.GatewaySandbox != "https://openapi-sandbox.dl.alipaydev.com" {
		t.Fatalf("unexpected V3 sandbox endpoint %s", core.GatewaySandbox)
	}
}
