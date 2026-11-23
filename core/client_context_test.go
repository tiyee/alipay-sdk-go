package core_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
)

type contextTransport func(*http.Request) (*http.Response, error)

func (f contextTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// Cancellation occurs inside the injected transport, without sleeps or external I/O.
func TestClientCancellationAndDeadlinePropagation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const timeout = time.Minute
	called := false
	transport := contextTransport(func(r *http.Request) (*http.Response, error) {
		called = true
		deadline, ok := r.Context().Deadline()
		if !ok || time.Until(deadline) > timeout {
			t.Fatal("request timeout was not propagated")
		}
		cancel()
		<-r.Context().Done()
		return nil, r.Context().Err()
	})
	client := newTestClient(t, "http://alipay.test", option.WithHTTPClient(&http.Client{Transport: transport}), option.WithRequestTimeout(timeout))
	_, err := core.Execute(ctx, client, newQueryOperation(t), core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "canceled"}})
	var transportErr *core.TransportError
	if !called || !errors.Is(err, context.Canceled) || !errors.As(err, &transportErr) || transportErr.Kind() != core.ErrorKindCanceled {
		t.Fatalf("error = %v; called = %v", err, called)
	}
}
